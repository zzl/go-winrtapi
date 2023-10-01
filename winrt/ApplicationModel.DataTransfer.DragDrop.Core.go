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
// flags
type CoreDragUIContentMode uint32

const (
	CoreDragUIContentMode_Auto     CoreDragUIContentMode = 0
	CoreDragUIContentMode_Deferred CoreDragUIContentMode = 1
)

// interfaces

// 7D56D344-8464-4FAF-AA49-37EA6E2D7BD1
var IID_ICoreDragDropManager = syscall.GUID{0x7D56D344, 0x8464, 0x4FAF,
	[8]byte{0xAA, 0x49, 0x37, 0xEA, 0x6E, 0x2D, 0x7B, 0xD1}}

type ICoreDragDropManagerInterface interface {
	win32.IInspectableInterface
	Add_TargetRequested(value TypedEventHandler[*ICoreDragDropManager, *ICoreDropOperationTargetRequestedEventArgs]) EventRegistrationToken
	Remove_TargetRequested(value EventRegistrationToken)
	Get_AreConcurrentOperationsEnabled() bool
	Put_AreConcurrentOperationsEnabled(value bool)
}

type ICoreDragDropManagerVtbl struct {
	win32.IInspectableVtbl
	Add_TargetRequested                uintptr
	Remove_TargetRequested             uintptr
	Get_AreConcurrentOperationsEnabled uintptr
	Put_AreConcurrentOperationsEnabled uintptr
}

type ICoreDragDropManager struct {
	win32.IInspectable
}

func (this *ICoreDragDropManager) Vtbl() *ICoreDragDropManagerVtbl {
	return (*ICoreDragDropManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreDragDropManager) Add_TargetRequested(value TypedEventHandler[*ICoreDragDropManager, *ICoreDropOperationTargetRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_TargetRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreDragDropManager) Remove_TargetRequested(value EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_TargetRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ICoreDragDropManager) Get_AreConcurrentOperationsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AreConcurrentOperationsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreDragDropManager) Put_AreConcurrentOperationsEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AreConcurrentOperationsEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 9542FDCA-DA12-4C1C-8D06-041DB29733C3
var IID_ICoreDragDropManagerStatics = syscall.GUID{0x9542FDCA, 0xDA12, 0x4C1C,
	[8]byte{0x8D, 0x06, 0x04, 0x1D, 0xB2, 0x97, 0x33, 0xC3}}

type ICoreDragDropManagerStaticsInterface interface {
	win32.IInspectableInterface
	GetForCurrentView() *ICoreDragDropManager
}

type ICoreDragDropManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForCurrentView uintptr
}

type ICoreDragDropManagerStatics struct {
	win32.IInspectable
}

func (this *ICoreDragDropManagerStatics) Vtbl() *ICoreDragDropManagerStaticsVtbl {
	return (*ICoreDragDropManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreDragDropManagerStatics) GetForCurrentView() *ICoreDragDropManager {
	var _result *ICoreDragDropManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 48353A8B-CB50-464E-9575-CD4E3A7AB028
var IID_ICoreDragInfo = syscall.GUID{0x48353A8B, 0xCB50, 0x464E,
	[8]byte{0x95, 0x75, 0xCD, 0x4E, 0x3A, 0x7A, 0xB0, 0x28}}

type ICoreDragInfoInterface interface {
	win32.IInspectableInterface
	Get_Data() *IDataPackageView
	Get_Modifiers() unsafe.Pointer
	Get_Position() Point
}

type ICoreDragInfoVtbl struct {
	win32.IInspectableVtbl
	Get_Data      uintptr
	Get_Modifiers uintptr
	Get_Position  uintptr
}

type ICoreDragInfo struct {
	win32.IInspectable
}

func (this *ICoreDragInfo) Vtbl() *ICoreDragInfoVtbl {
	return (*ICoreDragInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreDragInfo) Get_Data() *IDataPackageView {
	var _result *IDataPackageView
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICoreDragInfo) Get_Modifiers() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Modifiers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreDragInfo) Get_Position() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C54691E5-E6FB-4D74-B4B1-8A3C17F25E9E
var IID_ICoreDragInfo2 = syscall.GUID{0xC54691E5, 0xE6FB, 0x4D74,
	[8]byte{0xB4, 0xB1, 0x8A, 0x3C, 0x17, 0xF2, 0x5E, 0x9E}}

type ICoreDragInfo2Interface interface {
	win32.IInspectableInterface
	Get_AllowedOperations() DataPackageOperation
}

type ICoreDragInfo2Vtbl struct {
	win32.IInspectableVtbl
	Get_AllowedOperations uintptr
}

type ICoreDragInfo2 struct {
	win32.IInspectable
}

func (this *ICoreDragInfo2) Vtbl() *ICoreDragInfo2Vtbl {
	return (*ICoreDragInfo2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreDragInfo2) Get_AllowedOperations() DataPackageOperation {
	var _result DataPackageOperation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllowedOperations, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// CC06DE4F-6DB0-4E62-AB1B-A74A02DC6D85
var IID_ICoreDragOperation = syscall.GUID{0xCC06DE4F, 0x6DB0, 0x4E62,
	[8]byte{0xAB, 0x1B, 0xA7, 0x4A, 0x02, 0xDC, 0x6D, 0x85}}

type ICoreDragOperationInterface interface {
	win32.IInspectableInterface
	Get_Data() *IDataPackage
	SetPointerId(pointerId uint32)
	SetDragUIContentFromSoftwareBitmap(softwareBitmap *ISoftwareBitmap)
	SetDragUIContentFromSoftwareBitmapWithAnchorPoint(softwareBitmap *ISoftwareBitmap, anchorPoint Point)
	Get_DragUIContentMode() CoreDragUIContentMode
	Put_DragUIContentMode(value CoreDragUIContentMode)
	StartAsync() *IAsyncOperation[DataPackageOperation]
}

type ICoreDragOperationVtbl struct {
	win32.IInspectableVtbl
	Get_Data                                          uintptr
	SetPointerId                                      uintptr
	SetDragUIContentFromSoftwareBitmap                uintptr
	SetDragUIContentFromSoftwareBitmapWithAnchorPoint uintptr
	Get_DragUIContentMode                             uintptr
	Put_DragUIContentMode                             uintptr
	StartAsync                                        uintptr
}

type ICoreDragOperation struct {
	win32.IInspectable
}

func (this *ICoreDragOperation) Vtbl() *ICoreDragOperationVtbl {
	return (*ICoreDragOperationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreDragOperation) Get_Data() *IDataPackage {
	var _result *IDataPackage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICoreDragOperation) SetPointerId(pointerId uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPointerId, uintptr(unsafe.Pointer(this)), uintptr(pointerId))
	_ = _hr
}

func (this *ICoreDragOperation) SetDragUIContentFromSoftwareBitmap(softwareBitmap *ISoftwareBitmap) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetDragUIContentFromSoftwareBitmap, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(softwareBitmap)))
	_ = _hr
}

func (this *ICoreDragOperation) SetDragUIContentFromSoftwareBitmapWithAnchorPoint(softwareBitmap *ISoftwareBitmap, anchorPoint Point) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetDragUIContentFromSoftwareBitmapWithAnchorPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(softwareBitmap)), *(*uintptr)(unsafe.Pointer(&anchorPoint)))
	_ = _hr
}

func (this *ICoreDragOperation) Get_DragUIContentMode() CoreDragUIContentMode {
	var _result CoreDragUIContentMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DragUIContentMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreDragOperation) Put_DragUIContentMode(value CoreDragUIContentMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DragUIContentMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICoreDragOperation) StartAsync() *IAsyncOperation[DataPackageOperation] {
	var _result *IAsyncOperation[DataPackageOperation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 824B1E2C-D99A-4FC3-8507-6C182F33B46A
var IID_ICoreDragOperation2 = syscall.GUID{0x824B1E2C, 0xD99A, 0x4FC3,
	[8]byte{0x85, 0x07, 0x6C, 0x18, 0x2F, 0x33, 0xB4, 0x6A}}

type ICoreDragOperation2Interface interface {
	win32.IInspectableInterface
	Get_AllowedOperations() DataPackageOperation
	Put_AllowedOperations(value DataPackageOperation)
}

type ICoreDragOperation2Vtbl struct {
	win32.IInspectableVtbl
	Get_AllowedOperations uintptr
	Put_AllowedOperations uintptr
}

type ICoreDragOperation2 struct {
	win32.IInspectable
}

func (this *ICoreDragOperation2) Vtbl() *ICoreDragOperation2Vtbl {
	return (*ICoreDragOperation2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreDragOperation2) Get_AllowedOperations() DataPackageOperation {
	var _result DataPackageOperation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllowedOperations, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreDragOperation2) Put_AllowedOperations(value DataPackageOperation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AllowedOperations, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 89A85064-3389-4F4F-8897-7E8A3FFB3C93
var IID_ICoreDragUIOverride = syscall.GUID{0x89A85064, 0x3389, 0x4F4F,
	[8]byte{0x88, 0x97, 0x7E, 0x8A, 0x3F, 0xFB, 0x3C, 0x93}}

type ICoreDragUIOverrideInterface interface {
	win32.IInspectableInterface
	SetContentFromSoftwareBitmap(softwareBitmap *ISoftwareBitmap)
	SetContentFromSoftwareBitmapWithAnchorPoint(softwareBitmap *ISoftwareBitmap, anchorPoint Point)
	Get_IsContentVisible() bool
	Put_IsContentVisible(value bool)
	Get_Caption() string
	Put_Caption(value string)
	Get_IsCaptionVisible() bool
	Put_IsCaptionVisible(value bool)
	Get_IsGlyphVisible() bool
	Put_IsGlyphVisible(value bool)
	Clear()
}

type ICoreDragUIOverrideVtbl struct {
	win32.IInspectableVtbl
	SetContentFromSoftwareBitmap                uintptr
	SetContentFromSoftwareBitmapWithAnchorPoint uintptr
	Get_IsContentVisible                        uintptr
	Put_IsContentVisible                        uintptr
	Get_Caption                                 uintptr
	Put_Caption                                 uintptr
	Get_IsCaptionVisible                        uintptr
	Put_IsCaptionVisible                        uintptr
	Get_IsGlyphVisible                          uintptr
	Put_IsGlyphVisible                          uintptr
	Clear                                       uintptr
}

type ICoreDragUIOverride struct {
	win32.IInspectable
}

func (this *ICoreDragUIOverride) Vtbl() *ICoreDragUIOverrideVtbl {
	return (*ICoreDragUIOverrideVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreDragUIOverride) SetContentFromSoftwareBitmap(softwareBitmap *ISoftwareBitmap) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetContentFromSoftwareBitmap, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(softwareBitmap)))
	_ = _hr
}

func (this *ICoreDragUIOverride) SetContentFromSoftwareBitmapWithAnchorPoint(softwareBitmap *ISoftwareBitmap, anchorPoint Point) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetContentFromSoftwareBitmapWithAnchorPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(softwareBitmap)), *(*uintptr)(unsafe.Pointer(&anchorPoint)))
	_ = _hr
}

func (this *ICoreDragUIOverride) Get_IsContentVisible() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsContentVisible, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreDragUIOverride) Put_IsContentVisible(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsContentVisible, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICoreDragUIOverride) Get_Caption() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Caption, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICoreDragUIOverride) Put_Caption(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Caption, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICoreDragUIOverride) Get_IsCaptionVisible() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCaptionVisible, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreDragUIOverride) Put_IsCaptionVisible(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsCaptionVisible, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICoreDragUIOverride) Get_IsGlyphVisible() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsGlyphVisible, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreDragUIOverride) Put_IsGlyphVisible(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsGlyphVisible, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICoreDragUIOverride) Clear() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Clear, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// D9126196-4C5B-417D-BB37-76381DEF8DB4
var IID_ICoreDropOperationTarget = syscall.GUID{0xD9126196, 0x4C5B, 0x417D,
	[8]byte{0xBB, 0x37, 0x76, 0x38, 0x1D, 0xEF, 0x8D, 0xB4}}

type ICoreDropOperationTargetInterface interface {
	win32.IInspectableInterface
	EnterAsync(dragInfo *ICoreDragInfo, dragUIOverride *ICoreDragUIOverride) *IAsyncOperation[DataPackageOperation]
	OverAsync(dragInfo *ICoreDragInfo, dragUIOverride *ICoreDragUIOverride) *IAsyncOperation[DataPackageOperation]
	LeaveAsync(dragInfo *ICoreDragInfo) *IAsyncAction
	DropAsync(dragInfo *ICoreDragInfo) *IAsyncOperation[DataPackageOperation]
}

type ICoreDropOperationTargetVtbl struct {
	win32.IInspectableVtbl
	EnterAsync uintptr
	OverAsync  uintptr
	LeaveAsync uintptr
	DropAsync  uintptr
}

type ICoreDropOperationTarget struct {
	win32.IInspectable
}

func (this *ICoreDropOperationTarget) Vtbl() *ICoreDropOperationTargetVtbl {
	return (*ICoreDropOperationTargetVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreDropOperationTarget) EnterAsync(dragInfo *ICoreDragInfo, dragUIOverride *ICoreDragUIOverride) *IAsyncOperation[DataPackageOperation] {
	var _result *IAsyncOperation[DataPackageOperation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnterAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(dragInfo)), uintptr(unsafe.Pointer(dragUIOverride)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICoreDropOperationTarget) OverAsync(dragInfo *ICoreDragInfo, dragUIOverride *ICoreDragUIOverride) *IAsyncOperation[DataPackageOperation] {
	var _result *IAsyncOperation[DataPackageOperation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OverAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(dragInfo)), uintptr(unsafe.Pointer(dragUIOverride)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICoreDropOperationTarget) LeaveAsync(dragInfo *ICoreDragInfo) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LeaveAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(dragInfo)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICoreDropOperationTarget) DropAsync(dragInfo *ICoreDragInfo) *IAsyncOperation[DataPackageOperation] {
	var _result *IAsyncOperation[DataPackageOperation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DropAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(dragInfo)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2ACA929A-5E28-4EA6-829E-29134E665D6D
var IID_ICoreDropOperationTargetRequestedEventArgs = syscall.GUID{0x2ACA929A, 0x5E28, 0x4EA6,
	[8]byte{0x82, 0x9E, 0x29, 0x13, 0x4E, 0x66, 0x5D, 0x6D}}

type ICoreDropOperationTargetRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	SetTarget(target *ICoreDropOperationTarget)
}

type ICoreDropOperationTargetRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	SetTarget uintptr
}

type ICoreDropOperationTargetRequestedEventArgs struct {
	win32.IInspectable
}

func (this *ICoreDropOperationTargetRequestedEventArgs) Vtbl() *ICoreDropOperationTargetRequestedEventArgsVtbl {
	return (*ICoreDropOperationTargetRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreDropOperationTargetRequestedEventArgs) SetTarget(target *ICoreDropOperationTarget) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetTarget, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(target)))
	_ = _hr
}

// classes

type CoreDragDropManager struct {
	RtClass
	*ICoreDragDropManager
}

func NewICoreDragDropManagerStatics() *ICoreDragDropManagerStatics {
	var p *ICoreDragDropManagerStatics
	hs := NewHStr("Windows.ApplicationModel.DataTransfer.DragDrop.Core.CoreDragDropManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICoreDragDropManagerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type CoreDragInfo struct {
	RtClass
	*ICoreDragInfo
}

type CoreDragOperation struct {
	RtClass
	*ICoreDragOperation
}

func NewCoreDragOperation() *CoreDragOperation {
	hs := NewHStr("Windows.ApplicationModel.DataTransfer.DragDrop.Core.CoreDragOperation")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &CoreDragOperation{
		RtClass:            RtClass{PInspect: p},
		ICoreDragOperation: (*ICoreDragOperation)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type CoreDragUIOverride struct {
	RtClass
	*ICoreDragUIOverride
}

type CoreDropOperationTargetRequestedEventArgs struct {
	RtClass
	*ICoreDropOperationTargetRequestedEventArgs
}
