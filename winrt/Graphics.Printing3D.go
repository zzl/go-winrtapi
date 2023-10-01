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
type Print3DTaskCompletion int32

const (
	Print3DTaskCompletion_Abandoned Print3DTaskCompletion = 0
	Print3DTaskCompletion_Canceled  Print3DTaskCompletion = 1
	Print3DTaskCompletion_Failed    Print3DTaskCompletion = 2
	Print3DTaskCompletion_Slicing   Print3DTaskCompletion = 3
	Print3DTaskCompletion_Submitted Print3DTaskCompletion = 4
)

// enum
type Print3DTaskDetail int32

const (
	Print3DTaskDetail_Unknown                  Print3DTaskDetail = 0
	Print3DTaskDetail_ModelExceedsPrintBed     Print3DTaskDetail = 1
	Print3DTaskDetail_UploadFailed             Print3DTaskDetail = 2
	Print3DTaskDetail_InvalidMaterialSelection Print3DTaskDetail = 3
	Print3DTaskDetail_InvalidModel             Print3DTaskDetail = 4
	Print3DTaskDetail_ModelNotManifold         Print3DTaskDetail = 5
	Print3DTaskDetail_InvalidPrintTicket       Print3DTaskDetail = 6
)

// enum
type Printing3DBufferFormat int32

const (
	Printing3DBufferFormat_Unknown           Printing3DBufferFormat = 0
	Printing3DBufferFormat_R32G32B32A32Float Printing3DBufferFormat = 2
	Printing3DBufferFormat_R32G32B32A32UInt  Printing3DBufferFormat = 3
	Printing3DBufferFormat_R32G32B32Float    Printing3DBufferFormat = 6
	Printing3DBufferFormat_R32G32B32UInt     Printing3DBufferFormat = 7
	Printing3DBufferFormat_Printing3DDouble  Printing3DBufferFormat = 500
	Printing3DBufferFormat_Printing3DUInt    Printing3DBufferFormat = 501
)

// enum
type Printing3DMeshVerificationMode int32

const (
	Printing3DMeshVerificationMode_FindFirstError Printing3DMeshVerificationMode = 0
	Printing3DMeshVerificationMode_FindAllErrors  Printing3DMeshVerificationMode = 1
)

// enum
type Printing3DModelUnit int32

const (
	Printing3DModelUnit_Meter      Printing3DModelUnit = 0
	Printing3DModelUnit_Micron     Printing3DModelUnit = 1
	Printing3DModelUnit_Millimeter Printing3DModelUnit = 2
	Printing3DModelUnit_Centimeter Printing3DModelUnit = 3
	Printing3DModelUnit_Inch       Printing3DModelUnit = 4
	Printing3DModelUnit_Foot       Printing3DModelUnit = 5
)

// enum
type Printing3DObjectType int32

const (
	Printing3DObjectType_Model   Printing3DObjectType = 0
	Printing3DObjectType_Support Printing3DObjectType = 1
	Printing3DObjectType_Others  Printing3DObjectType = 2
)

// enum
type Printing3DPackageCompression int32

const (
	Printing3DPackageCompression_Low    Printing3DPackageCompression = 0
	Printing3DPackageCompression_Medium Printing3DPackageCompression = 1
	Printing3DPackageCompression_High   Printing3DPackageCompression = 2
)

// enum
type Printing3DTextureEdgeBehavior int32

const (
	Printing3DTextureEdgeBehavior_None   Printing3DTextureEdgeBehavior = 0
	Printing3DTextureEdgeBehavior_Wrap   Printing3DTextureEdgeBehavior = 1
	Printing3DTextureEdgeBehavior_Mirror Printing3DTextureEdgeBehavior = 2
	Printing3DTextureEdgeBehavior_Clamp  Printing3DTextureEdgeBehavior = 3
)

// structs

type Printing3DBufferDescription struct {
	Format Printing3DBufferFormat
	Stride uint32
}

type Printing3DContract struct {
}

// func types

// E9175E70-C917-46DE-BB51-D9A94DB3711F
type Print3DTaskSourceRequestedHandler func(args *IPrint3DTaskSourceRequestedArgs) com.Error

// interfaces

// 4D2FCB0A-7366-4971-8BD5-17C4E3E8C6C0
var IID_IPrint3DManager = syscall.GUID{0x4D2FCB0A, 0x7366, 0x4971,
	[8]byte{0x8B, 0xD5, 0x17, 0xC4, 0xE3, 0xE8, 0xC6, 0xC0}}

type IPrint3DManagerInterface interface {
	win32.IInspectableInterface
	Add_TaskRequested(eventHandler TypedEventHandler[*IPrint3DManager, *IPrint3DTaskRequestedEventArgs]) EventRegistrationToken
	Remove_TaskRequested(token EventRegistrationToken)
}

type IPrint3DManagerVtbl struct {
	win32.IInspectableVtbl
	Add_TaskRequested    uintptr
	Remove_TaskRequested uintptr
}

type IPrint3DManager struct {
	win32.IInspectable
}

func (this *IPrint3DManager) Vtbl() *IPrint3DManagerVtbl {
	return (*IPrint3DManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrint3DManager) Add_TaskRequested(eventHandler TypedEventHandler[*IPrint3DManager, *IPrint3DTaskRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_TaskRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(eventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrint3DManager) Remove_TaskRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_TaskRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 0EF1CAFE-A9AD-4C08-A917-1D1F863EABCB
var IID_IPrint3DManagerStatics = syscall.GUID{0x0EF1CAFE, 0xA9AD, 0x4C08,
	[8]byte{0xA9, 0x17, 0x1D, 0x1F, 0x86, 0x3E, 0xAB, 0xCB}}

type IPrint3DManagerStaticsInterface interface {
	win32.IInspectableInterface
	GetForCurrentView() *IPrint3DManager
	ShowPrintUIAsync() *IAsyncOperation[bool]
}

type IPrint3DManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForCurrentView uintptr
	ShowPrintUIAsync  uintptr
}

type IPrint3DManagerStatics struct {
	win32.IInspectable
}

func (this *IPrint3DManagerStatics) Vtbl() *IPrint3DManagerStaticsVtbl {
	return (*IPrint3DManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrint3DManagerStatics) GetForCurrentView() *IPrint3DManager {
	var _result *IPrint3DManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrint3DManagerStatics) ShowPrintUIAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowPrintUIAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8CE3D080-2118-4C28-80DE-F426D70191AE
var IID_IPrint3DTask = syscall.GUID{0x8CE3D080, 0x2118, 0x4C28,
	[8]byte{0x80, 0xDE, 0xF4, 0x26, 0xD7, 0x01, 0x91, 0xAE}}

type IPrint3DTaskInterface interface {
	win32.IInspectableInterface
	Get_Source() *IPrinting3D3MFPackage
	Add_Submitting(eventHandler TypedEventHandler[*IPrint3DTask, interface{}]) EventRegistrationToken
	Remove_Submitting(eventCookie EventRegistrationToken)
	Add_Completed(eventHandler TypedEventHandler[*IPrint3DTask, *IPrint3DTaskCompletedEventArgs]) EventRegistrationToken
	Remove_Completed(eventCookie EventRegistrationToken)
	Add_SourceChanged(eventHandler TypedEventHandler[*IPrint3DTask, *IPrint3DTaskSourceChangedEventArgs]) EventRegistrationToken
	Remove_SourceChanged(eventCookie EventRegistrationToken)
}

type IPrint3DTaskVtbl struct {
	win32.IInspectableVtbl
	Get_Source           uintptr
	Add_Submitting       uintptr
	Remove_Submitting    uintptr
	Add_Completed        uintptr
	Remove_Completed     uintptr
	Add_SourceChanged    uintptr
	Remove_SourceChanged uintptr
}

type IPrint3DTask struct {
	win32.IInspectable
}

func (this *IPrint3DTask) Vtbl() *IPrint3DTaskVtbl {
	return (*IPrint3DTaskVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrint3DTask) Get_Source() *IPrinting3D3MFPackage {
	var _result *IPrinting3D3MFPackage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrint3DTask) Add_Submitting(eventHandler TypedEventHandler[*IPrint3DTask, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Submitting, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(eventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrint3DTask) Remove_Submitting(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Submitting, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

func (this *IPrint3DTask) Add_Completed(eventHandler TypedEventHandler[*IPrint3DTask, *IPrint3DTaskCompletedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Completed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(eventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrint3DTask) Remove_Completed(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Completed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

func (this *IPrint3DTask) Add_SourceChanged(eventHandler TypedEventHandler[*IPrint3DTask, *IPrint3DTaskSourceChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SourceChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(eventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrint3DTask) Remove_SourceChanged(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SourceChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

// CC1914AF-2614-4F1D-ACCC-D6FC4FDA5455
var IID_IPrint3DTaskCompletedEventArgs = syscall.GUID{0xCC1914AF, 0x2614, 0x4F1D,
	[8]byte{0xAC, 0xCC, 0xD6, 0xFC, 0x4F, 0xDA, 0x54, 0x55}}

type IPrint3DTaskCompletedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Completion() Print3DTaskCompletion
	Get_ExtendedStatus() Print3DTaskDetail
}

type IPrint3DTaskCompletedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Completion     uintptr
	Get_ExtendedStatus uintptr
}

type IPrint3DTaskCompletedEventArgs struct {
	win32.IInspectable
}

func (this *IPrint3DTaskCompletedEventArgs) Vtbl() *IPrint3DTaskCompletedEventArgsVtbl {
	return (*IPrint3DTaskCompletedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrint3DTaskCompletedEventArgs) Get_Completion() Print3DTaskCompletion {
	var _result Print3DTaskCompletion
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Completion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrint3DTaskCompletedEventArgs) Get_ExtendedStatus() Print3DTaskDetail {
	var _result Print3DTaskDetail
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 2595C46F-2245-4C5A-8731-0D604DC6BC3C
var IID_IPrint3DTaskRequest = syscall.GUID{0x2595C46F, 0x2245, 0x4C5A,
	[8]byte{0x87, 0x31, 0x0D, 0x60, 0x4D, 0xC6, 0xBC, 0x3C}}

type IPrint3DTaskRequestInterface interface {
	win32.IInspectableInterface
	CreateTask(title string, printerId string, handler Print3DTaskSourceRequestedHandler) *IPrint3DTask
}

type IPrint3DTaskRequestVtbl struct {
	win32.IInspectableVtbl
	CreateTask uintptr
}

type IPrint3DTaskRequest struct {
	win32.IInspectable
}

func (this *IPrint3DTaskRequest) Vtbl() *IPrint3DTaskRequestVtbl {
	return (*IPrint3DTaskRequestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrint3DTaskRequest) CreateTask(title string, printerId string, handler Print3DTaskSourceRequestedHandler) *IPrint3DTask {
	var _result *IPrint3DTask
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateTask, uintptr(unsafe.Pointer(this)), NewHStr(title).Ptr, NewHStr(printerId).Ptr, uintptr(unsafe.Pointer(NewOneArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 150CB77F-18C5-40D7-9F40-FAB3096E05A9
var IID_IPrint3DTaskRequestedEventArgs = syscall.GUID{0x150CB77F, 0x18C5, 0x40D7,
	[8]byte{0x9F, 0x40, 0xFA, 0xB3, 0x09, 0x6E, 0x05, 0xA9}}

type IPrint3DTaskRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Request() *IPrint3DTaskRequest
}

type IPrint3DTaskRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Request uintptr
}

type IPrint3DTaskRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IPrint3DTaskRequestedEventArgs) Vtbl() *IPrint3DTaskRequestedEventArgsVtbl {
	return (*IPrint3DTaskRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrint3DTaskRequestedEventArgs) Get_Request() *IPrint3DTaskRequest {
	var _result *IPrint3DTaskRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Request, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5BCD34AF-24E9-4C10-8D07-14C346BA3FCF
var IID_IPrint3DTaskSourceChangedEventArgs = syscall.GUID{0x5BCD34AF, 0x24E9, 0x4C10,
	[8]byte{0x8D, 0x07, 0x14, 0xC3, 0x46, 0xBA, 0x3F, 0xCF}}

type IPrint3DTaskSourceChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Source() *IPrinting3D3MFPackage
}

type IPrint3DTaskSourceChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Source uintptr
}

type IPrint3DTaskSourceChangedEventArgs struct {
	win32.IInspectable
}

func (this *IPrint3DTaskSourceChangedEventArgs) Vtbl() *IPrint3DTaskSourceChangedEventArgsVtbl {
	return (*IPrint3DTaskSourceChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrint3DTaskSourceChangedEventArgs) Get_Source() *IPrinting3D3MFPackage {
	var _result *IPrinting3D3MFPackage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C77C9ABA-24AF-424D-A3BF-92250C355602
var IID_IPrint3DTaskSourceRequestedArgs = syscall.GUID{0xC77C9ABA, 0x24AF, 0x424D,
	[8]byte{0xA3, 0xBF, 0x92, 0x25, 0x0C, 0x35, 0x56, 0x02}}

type IPrint3DTaskSourceRequestedArgsInterface interface {
	win32.IInspectableInterface
	SetSource(source *IPrinting3D3MFPackage)
}

type IPrint3DTaskSourceRequestedArgsVtbl struct {
	win32.IInspectableVtbl
	SetSource uintptr
}

type IPrint3DTaskSourceRequestedArgs struct {
	win32.IInspectable
}

func (this *IPrint3DTaskSourceRequestedArgs) Vtbl() *IPrint3DTaskSourceRequestedArgsVtbl {
	return (*IPrint3DTaskSourceRequestedArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrint3DTaskSourceRequestedArgs) SetSource(source *IPrinting3D3MFPackage) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)))
	_ = _hr
}

// F64DD5C8-2AB7-45A9-A1B7-267E948D5B18
var IID_IPrinting3D3MFPackage = syscall.GUID{0xF64DD5C8, 0x2AB7, 0x45A9,
	[8]byte{0xA1, 0xB7, 0x26, 0x7E, 0x94, 0x8D, 0x5B, 0x18}}

type IPrinting3D3MFPackageInterface interface {
	win32.IInspectableInterface
	SaveAsync() *IAsyncOperation[*IRandomAccessStream]
	Get_PrintTicket() *IRandomAccessStream
	Put_PrintTicket(value *IRandomAccessStream)
	Get_ModelPart() *IRandomAccessStream
	Put_ModelPart(value *IRandomAccessStream)
	Get_Thumbnail() *IPrinting3DTextureResource
	Put_Thumbnail(value *IPrinting3DTextureResource)
	Get_Textures() *IVector[*IPrinting3DTextureResource]
	LoadModelFromPackageAsync(value *IRandomAccessStream) *IAsyncOperation[*IPrinting3DModel]
	SaveModelToPackageAsync(value *IPrinting3DModel) *IAsyncAction
}

type IPrinting3D3MFPackageVtbl struct {
	win32.IInspectableVtbl
	SaveAsync                 uintptr
	Get_PrintTicket           uintptr
	Put_PrintTicket           uintptr
	Get_ModelPart             uintptr
	Put_ModelPart             uintptr
	Get_Thumbnail             uintptr
	Put_Thumbnail             uintptr
	Get_Textures              uintptr
	LoadModelFromPackageAsync uintptr
	SaveModelToPackageAsync   uintptr
}

type IPrinting3D3MFPackage struct {
	win32.IInspectable
}

func (this *IPrinting3D3MFPackage) Vtbl() *IPrinting3D3MFPackageVtbl {
	return (*IPrinting3D3MFPackageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3D3MFPackage) SaveAsync() *IAsyncOperation[*IRandomAccessStream] {
	var _result *IAsyncOperation[*IRandomAccessStream]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SaveAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3D3MFPackage) Get_PrintTicket() *IRandomAccessStream {
	var _result *IRandomAccessStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrintTicket, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3D3MFPackage) Put_PrintTicket(value *IRandomAccessStream) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PrintTicket, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPrinting3D3MFPackage) Get_ModelPart() *IRandomAccessStream {
	var _result *IRandomAccessStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ModelPart, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3D3MFPackage) Put_ModelPart(value *IRandomAccessStream) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ModelPart, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPrinting3D3MFPackage) Get_Thumbnail() *IPrinting3DTextureResource {
	var _result *IPrinting3DTextureResource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Thumbnail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3D3MFPackage) Put_Thumbnail(value *IPrinting3DTextureResource) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Thumbnail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPrinting3D3MFPackage) Get_Textures() *IVector[*IPrinting3DTextureResource] {
	var _result *IVector[*IPrinting3DTextureResource]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Textures, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3D3MFPackage) LoadModelFromPackageAsync(value *IRandomAccessStream) *IAsyncOperation[*IPrinting3DModel] {
	var _result *IAsyncOperation[*IPrinting3DModel]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LoadModelFromPackageAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3D3MFPackage) SaveModelToPackageAsync(value *IPrinting3DModel) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SaveModelToPackageAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 965C7AC4-93CB-4430-92B8-789CD454F883
var IID_IPrinting3D3MFPackage2 = syscall.GUID{0x965C7AC4, 0x93CB, 0x4430,
	[8]byte{0x92, 0xB8, 0x78, 0x9C, 0xD4, 0x54, 0xF8, 0x83}}

type IPrinting3D3MFPackage2Interface interface {
	win32.IInspectableInterface
	Get_Compression() Printing3DPackageCompression
	Put_Compression(value Printing3DPackageCompression)
}

type IPrinting3D3MFPackage2Vtbl struct {
	win32.IInspectableVtbl
	Get_Compression uintptr
	Put_Compression uintptr
}

type IPrinting3D3MFPackage2 struct {
	win32.IInspectable
}

func (this *IPrinting3D3MFPackage2) Vtbl() *IPrinting3D3MFPackage2Vtbl {
	return (*IPrinting3D3MFPackage2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3D3MFPackage2) Get_Compression() Printing3DPackageCompression {
	var _result Printing3DPackageCompression
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Compression, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrinting3D3MFPackage2) Put_Compression(value Printing3DPackageCompression) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Compression, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 7058D9AF-7A9A-4787-B817-F6F459214823
var IID_IPrinting3D3MFPackageStatics = syscall.GUID{0x7058D9AF, 0x7A9A, 0x4787,
	[8]byte{0xB8, 0x17, 0xF6, 0xF4, 0x59, 0x21, 0x48, 0x23}}

type IPrinting3D3MFPackageStaticsInterface interface {
	win32.IInspectableInterface
	LoadAsync(value *IRandomAccessStream) *IAsyncOperation[*IPrinting3D3MFPackage]
}

type IPrinting3D3MFPackageStaticsVtbl struct {
	win32.IInspectableVtbl
	LoadAsync uintptr
}

type IPrinting3D3MFPackageStatics struct {
	win32.IInspectable
}

func (this *IPrinting3D3MFPackageStatics) Vtbl() *IPrinting3D3MFPackageStaticsVtbl {
	return (*IPrinting3D3MFPackageStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3D3MFPackageStatics) LoadAsync(value *IRandomAccessStream) *IAsyncOperation[*IPrinting3D3MFPackage] {
	var _result *IAsyncOperation[*IPrinting3D3MFPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LoadAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D0F0E743-C50C-4BCB-9D04-FC16ADCEA2C9
var IID_IPrinting3DBaseMaterial = syscall.GUID{0xD0F0E743, 0xC50C, 0x4BCB,
	[8]byte{0x9D, 0x04, 0xFC, 0x16, 0xAD, 0xCE, 0xA2, 0xC9}}

type IPrinting3DBaseMaterialInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Put_Name(value string)
	Get_Color() *IPrinting3DColorMaterial
	Put_Color(value *IPrinting3DColorMaterial)
}

type IPrinting3DBaseMaterialVtbl struct {
	win32.IInspectableVtbl
	Get_Name  uintptr
	Put_Name  uintptr
	Get_Color uintptr
	Put_Color uintptr
}

type IPrinting3DBaseMaterial struct {
	win32.IInspectable
}

func (this *IPrinting3DBaseMaterial) Vtbl() *IPrinting3DBaseMaterialVtbl {
	return (*IPrinting3DBaseMaterialVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DBaseMaterial) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPrinting3DBaseMaterial) Put_Name(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Name, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPrinting3DBaseMaterial) Get_Color() *IPrinting3DColorMaterial {
	var _result *IPrinting3DColorMaterial
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Color, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DBaseMaterial) Put_Color(value *IPrinting3DColorMaterial) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Color, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 94F070B8-2515-4A8D-A1F0-D0FC13D06021
var IID_IPrinting3DBaseMaterialGroup = syscall.GUID{0x94F070B8, 0x2515, 0x4A8D,
	[8]byte{0xA1, 0xF0, 0xD0, 0xFC, 0x13, 0xD0, 0x60, 0x21}}

type IPrinting3DBaseMaterialGroupInterface interface {
	win32.IInspectableInterface
	Get_Bases() *IVector[*IPrinting3DBaseMaterial]
	Get_MaterialGroupId() uint32
}

type IPrinting3DBaseMaterialGroupVtbl struct {
	win32.IInspectableVtbl
	Get_Bases           uintptr
	Get_MaterialGroupId uintptr
}

type IPrinting3DBaseMaterialGroup struct {
	win32.IInspectable
}

func (this *IPrinting3DBaseMaterialGroup) Vtbl() *IPrinting3DBaseMaterialGroupVtbl {
	return (*IPrinting3DBaseMaterialGroupVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DBaseMaterialGroup) Get_Bases() *IVector[*IPrinting3DBaseMaterial] {
	var _result *IVector[*IPrinting3DBaseMaterial]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bases, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DBaseMaterialGroup) Get_MaterialGroupId() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaterialGroupId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5C1546DC-8697-4193-976B-84BB4116E5BF
var IID_IPrinting3DBaseMaterialGroupFactory = syscall.GUID{0x5C1546DC, 0x8697, 0x4193,
	[8]byte{0x97, 0x6B, 0x84, 0xBB, 0x41, 0x16, 0xE5, 0xBF}}

type IPrinting3DBaseMaterialGroupFactoryInterface interface {
	win32.IInspectableInterface
	Create(MaterialGroupId uint32) *IPrinting3DBaseMaterialGroup
}

type IPrinting3DBaseMaterialGroupFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IPrinting3DBaseMaterialGroupFactory struct {
	win32.IInspectable
}

func (this *IPrinting3DBaseMaterialGroupFactory) Vtbl() *IPrinting3DBaseMaterialGroupFactoryVtbl {
	return (*IPrinting3DBaseMaterialGroupFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DBaseMaterialGroupFactory) Create(MaterialGroupId uint32) *IPrinting3DBaseMaterialGroup {
	var _result *IPrinting3DBaseMaterialGroup
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(MaterialGroupId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 815A47BC-374A-476D-BE92-3ECFD1CB9776
var IID_IPrinting3DBaseMaterialStatics = syscall.GUID{0x815A47BC, 0x374A, 0x476D,
	[8]byte{0xBE, 0x92, 0x3E, 0xCF, 0xD1, 0xCB, 0x97, 0x76}}

type IPrinting3DBaseMaterialStaticsInterface interface {
	win32.IInspectableInterface
	Get_Abs() string
	Get_Pla() string
}

type IPrinting3DBaseMaterialStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Abs uintptr
	Get_Pla uintptr
}

type IPrinting3DBaseMaterialStatics struct {
	win32.IInspectable
}

func (this *IPrinting3DBaseMaterialStatics) Vtbl() *IPrinting3DBaseMaterialStaticsVtbl {
	return (*IPrinting3DBaseMaterialStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DBaseMaterialStatics) Get_Abs() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Abs, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPrinting3DBaseMaterialStatics) Get_Pla() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Pla, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// E1899928-7CE7-4285-A35D-F145C9510C7B
var IID_IPrinting3DColorMaterial = syscall.GUID{0xE1899928, 0x7CE7, 0x4285,
	[8]byte{0xA3, 0x5D, 0xF1, 0x45, 0xC9, 0x51, 0x0C, 0x7B}}

type IPrinting3DColorMaterialInterface interface {
	win32.IInspectableInterface
	Get_Value() uint32
	Put_Value(value uint32)
}

type IPrinting3DColorMaterialVtbl struct {
	win32.IInspectableVtbl
	Get_Value uintptr
	Put_Value uintptr
}

type IPrinting3DColorMaterial struct {
	win32.IInspectable
}

func (this *IPrinting3DColorMaterial) Vtbl() *IPrinting3DColorMaterialVtbl {
	return (*IPrinting3DColorMaterialVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DColorMaterial) Get_Value() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrinting3DColorMaterial) Put_Value(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Value, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// FAB0E852-0AEF-44E9-9DDD-36EEEA5ACD44
var IID_IPrinting3DColorMaterial2 = syscall.GUID{0xFAB0E852, 0x0AEF, 0x44E9,
	[8]byte{0x9D, 0xDD, 0x36, 0xEE, 0xEA, 0x5A, 0xCD, 0x44}}

type IPrinting3DColorMaterial2Interface interface {
	win32.IInspectableInterface
	Get_Color() unsafe.Pointer
	Put_Color(value unsafe.Pointer)
}

type IPrinting3DColorMaterial2Vtbl struct {
	win32.IInspectableVtbl
	Get_Color uintptr
	Put_Color uintptr
}

type IPrinting3DColorMaterial2 struct {
	win32.IInspectable
}

func (this *IPrinting3DColorMaterial2) Vtbl() *IPrinting3DColorMaterial2Vtbl {
	return (*IPrinting3DColorMaterial2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DColorMaterial2) Get_Color() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Color, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrinting3DColorMaterial2) Put_Color(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Color, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 001A6BD0-AADF-4226-AFE9-F369A0B45004
var IID_IPrinting3DColorMaterialGroup = syscall.GUID{0x001A6BD0, 0xAADF, 0x4226,
	[8]byte{0xAF, 0xE9, 0xF3, 0x69, 0xA0, 0xB4, 0x50, 0x04}}

type IPrinting3DColorMaterialGroupInterface interface {
	win32.IInspectableInterface
	Get_Colors() *IVector[*IPrinting3DColorMaterial]
	Get_MaterialGroupId() uint32
}

type IPrinting3DColorMaterialGroupVtbl struct {
	win32.IInspectableVtbl
	Get_Colors          uintptr
	Get_MaterialGroupId uintptr
}

type IPrinting3DColorMaterialGroup struct {
	win32.IInspectable
}

func (this *IPrinting3DColorMaterialGroup) Vtbl() *IPrinting3DColorMaterialGroupVtbl {
	return (*IPrinting3DColorMaterialGroupVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DColorMaterialGroup) Get_Colors() *IVector[*IPrinting3DColorMaterial] {
	var _result *IVector[*IPrinting3DColorMaterial]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Colors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DColorMaterialGroup) Get_MaterialGroupId() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaterialGroupId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 71D38D6D-B1EA-4A5B-BC54-19C65F3DF044
var IID_IPrinting3DColorMaterialGroupFactory = syscall.GUID{0x71D38D6D, 0xB1EA, 0x4A5B,
	[8]byte{0xBC, 0x54, 0x19, 0xC6, 0x5F, 0x3D, 0xF0, 0x44}}

type IPrinting3DColorMaterialGroupFactoryInterface interface {
	win32.IInspectableInterface
	Create(MaterialGroupId uint32) *IPrinting3DColorMaterialGroup
}

type IPrinting3DColorMaterialGroupFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IPrinting3DColorMaterialGroupFactory struct {
	win32.IInspectable
}

func (this *IPrinting3DColorMaterialGroupFactory) Vtbl() *IPrinting3DColorMaterialGroupFactoryVtbl {
	return (*IPrinting3DColorMaterialGroupFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DColorMaterialGroupFactory) Create(MaterialGroupId uint32) *IPrinting3DColorMaterialGroup {
	var _result *IPrinting3DColorMaterialGroup
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(MaterialGroupId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7E287845-BF7F-4CDB-A27F-30A01437FEDE
var IID_IPrinting3DComponent = syscall.GUID{0x7E287845, 0xBF7F, 0x4CDB,
	[8]byte{0xA2, 0x7F, 0x30, 0xA0, 0x14, 0x37, 0xFE, 0xDE}}

type IPrinting3DComponentInterface interface {
	win32.IInspectableInterface
	Get_Mesh() *IPrinting3DMesh
	Put_Mesh(value *IPrinting3DMesh)
	Get_Components() *IVector[*IPrinting3DComponentWithMatrix]
	Get_Thumbnail() *IPrinting3DTextureResource
	Put_Thumbnail(value *IPrinting3DTextureResource)
	Get_Type() Printing3DObjectType
	Put_Type(value Printing3DObjectType)
	Get_Name() string
	Put_Name(value string)
	Get_PartNumber() string
	Put_PartNumber(value string)
}

type IPrinting3DComponentVtbl struct {
	win32.IInspectableVtbl
	Get_Mesh       uintptr
	Put_Mesh       uintptr
	Get_Components uintptr
	Get_Thumbnail  uintptr
	Put_Thumbnail  uintptr
	Get_Type       uintptr
	Put_Type       uintptr
	Get_Name       uintptr
	Put_Name       uintptr
	Get_PartNumber uintptr
	Put_PartNumber uintptr
}

type IPrinting3DComponent struct {
	win32.IInspectable
}

func (this *IPrinting3DComponent) Vtbl() *IPrinting3DComponentVtbl {
	return (*IPrinting3DComponentVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DComponent) Get_Mesh() *IPrinting3DMesh {
	var _result *IPrinting3DMesh
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mesh, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DComponent) Put_Mesh(value *IPrinting3DMesh) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Mesh, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPrinting3DComponent) Get_Components() *IVector[*IPrinting3DComponentWithMatrix] {
	var _result *IVector[*IPrinting3DComponentWithMatrix]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Components, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DComponent) Get_Thumbnail() *IPrinting3DTextureResource {
	var _result *IPrinting3DTextureResource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Thumbnail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DComponent) Put_Thumbnail(value *IPrinting3DTextureResource) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Thumbnail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPrinting3DComponent) Get_Type() Printing3DObjectType {
	var _result Printing3DObjectType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Type, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrinting3DComponent) Put_Type(value Printing3DObjectType) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Type, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrinting3DComponent) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPrinting3DComponent) Put_Name(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Name, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPrinting3DComponent) Get_PartNumber() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PartNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPrinting3DComponent) Put_PartNumber(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PartNumber, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 3279F335-0EF0-456B-9A21-49BEBE8B51C2
var IID_IPrinting3DComponentWithMatrix = syscall.GUID{0x3279F335, 0x0EF0, 0x456B,
	[8]byte{0x9A, 0x21, 0x49, 0xBE, 0xBE, 0x8B, 0x51, 0xC2}}

type IPrinting3DComponentWithMatrixInterface interface {
	win32.IInspectableInterface
	Get_Component() *IPrinting3DComponent
	Put_Component(value *IPrinting3DComponent)
	Get_Matrix() Matrix4x4
	Put_Matrix(value Matrix4x4)
}

type IPrinting3DComponentWithMatrixVtbl struct {
	win32.IInspectableVtbl
	Get_Component uintptr
	Put_Component uintptr
	Get_Matrix    uintptr
	Put_Matrix    uintptr
}

type IPrinting3DComponentWithMatrix struct {
	win32.IInspectable
}

func (this *IPrinting3DComponentWithMatrix) Vtbl() *IPrinting3DComponentWithMatrixVtbl {
	return (*IPrinting3DComponentWithMatrixVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DComponentWithMatrix) Get_Component() *IPrinting3DComponent {
	var _result *IPrinting3DComponent
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Component, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DComponentWithMatrix) Put_Component(value *IPrinting3DComponent) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Component, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPrinting3DComponentWithMatrix) Get_Matrix() Matrix4x4 {
	var _result Matrix4x4
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Matrix, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrinting3DComponentWithMatrix) Put_Matrix(value Matrix4x4) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Matrix, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

// 462238DD-562E-4F6C-882D-F4D841FD63C7
var IID_IPrinting3DCompositeMaterial = syscall.GUID{0x462238DD, 0x562E, 0x4F6C,
	[8]byte{0x88, 0x2D, 0xF4, 0xD8, 0x41, 0xFD, 0x63, 0xC7}}

type IPrinting3DCompositeMaterialInterface interface {
	win32.IInspectableInterface
	Get_Values() *IVector[float64]
}

type IPrinting3DCompositeMaterialVtbl struct {
	win32.IInspectableVtbl
	Get_Values uintptr
}

type IPrinting3DCompositeMaterial struct {
	win32.IInspectable
}

func (this *IPrinting3DCompositeMaterial) Vtbl() *IPrinting3DCompositeMaterialVtbl {
	return (*IPrinting3DCompositeMaterialVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DCompositeMaterial) Get_Values() *IVector[float64] {
	var _result *IVector[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Values, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8D946A5B-40F1-496D-A5FB-340A5A678E30
var IID_IPrinting3DCompositeMaterialGroup = syscall.GUID{0x8D946A5B, 0x40F1, 0x496D,
	[8]byte{0xA5, 0xFB, 0x34, 0x0A, 0x5A, 0x67, 0x8E, 0x30}}

type IPrinting3DCompositeMaterialGroupInterface interface {
	win32.IInspectableInterface
	Get_Composites() *IVector[*IPrinting3DCompositeMaterial]
	Get_MaterialGroupId() uint32
	Get_MaterialIndices() *IVector[uint32]
}

type IPrinting3DCompositeMaterialGroupVtbl struct {
	win32.IInspectableVtbl
	Get_Composites      uintptr
	Get_MaterialGroupId uintptr
	Get_MaterialIndices uintptr
}

type IPrinting3DCompositeMaterialGroup struct {
	win32.IInspectable
}

func (this *IPrinting3DCompositeMaterialGroup) Vtbl() *IPrinting3DCompositeMaterialGroupVtbl {
	return (*IPrinting3DCompositeMaterialGroupVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DCompositeMaterialGroup) Get_Composites() *IVector[*IPrinting3DCompositeMaterial] {
	var _result *IVector[*IPrinting3DCompositeMaterial]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Composites, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DCompositeMaterialGroup) Get_MaterialGroupId() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaterialGroupId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrinting3DCompositeMaterialGroup) Get_MaterialIndices() *IVector[uint32] {
	var _result *IVector[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaterialIndices, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 06E86D62-7D3B-41E1-944C-BAFDE4555483
var IID_IPrinting3DCompositeMaterialGroup2 = syscall.GUID{0x06E86D62, 0x7D3B, 0x41E1,
	[8]byte{0x94, 0x4C, 0xBA, 0xFD, 0xE4, 0x55, 0x54, 0x83}}

type IPrinting3DCompositeMaterialGroup2Interface interface {
	win32.IInspectableInterface
	Get_BaseMaterialGroup() *IPrinting3DBaseMaterialGroup
	Put_BaseMaterialGroup(value *IPrinting3DBaseMaterialGroup)
}

type IPrinting3DCompositeMaterialGroup2Vtbl struct {
	win32.IInspectableVtbl
	Get_BaseMaterialGroup uintptr
	Put_BaseMaterialGroup uintptr
}

type IPrinting3DCompositeMaterialGroup2 struct {
	win32.IInspectable
}

func (this *IPrinting3DCompositeMaterialGroup2) Vtbl() *IPrinting3DCompositeMaterialGroup2Vtbl {
	return (*IPrinting3DCompositeMaterialGroup2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DCompositeMaterialGroup2) Get_BaseMaterialGroup() *IPrinting3DBaseMaterialGroup {
	var _result *IPrinting3DBaseMaterialGroup
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BaseMaterialGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DCompositeMaterialGroup2) Put_BaseMaterialGroup(value *IPrinting3DBaseMaterialGroup) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BaseMaterialGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// D08ECD13-92FF-43AA-A627-8D43C22C817E
var IID_IPrinting3DCompositeMaterialGroupFactory = syscall.GUID{0xD08ECD13, 0x92FF, 0x43AA,
	[8]byte{0xA6, 0x27, 0x8D, 0x43, 0xC2, 0x2C, 0x81, 0x7E}}

type IPrinting3DCompositeMaterialGroupFactoryInterface interface {
	win32.IInspectableInterface
	Create(MaterialGroupId uint32) *IPrinting3DCompositeMaterialGroup
}

type IPrinting3DCompositeMaterialGroupFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IPrinting3DCompositeMaterialGroupFactory struct {
	win32.IInspectable
}

func (this *IPrinting3DCompositeMaterialGroupFactory) Vtbl() *IPrinting3DCompositeMaterialGroupFactoryVtbl {
	return (*IPrinting3DCompositeMaterialGroupFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DCompositeMaterialGroupFactory) Create(MaterialGroupId uint32) *IPrinting3DCompositeMaterialGroup {
	var _result *IPrinting3DCompositeMaterialGroup
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(MaterialGroupId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BBFED397-2D74-46F7-BE85-99A67BBB6629
var IID_IPrinting3DFaceReductionOptions = syscall.GUID{0xBBFED397, 0x2D74, 0x46F7,
	[8]byte{0xBE, 0x85, 0x99, 0xA6, 0x7B, 0xBB, 0x66, 0x29}}

type IPrinting3DFaceReductionOptionsInterface interface {
	win32.IInspectableInterface
	Get_MaxReductionArea() float64
	Put_MaxReductionArea(value float64)
	Get_TargetTriangleCount() uint32
	Put_TargetTriangleCount(value uint32)
	Get_MaxEdgeLength() float64
	Put_MaxEdgeLength(value float64)
}

type IPrinting3DFaceReductionOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_MaxReductionArea    uintptr
	Put_MaxReductionArea    uintptr
	Get_TargetTriangleCount uintptr
	Put_TargetTriangleCount uintptr
	Get_MaxEdgeLength       uintptr
	Put_MaxEdgeLength       uintptr
}

type IPrinting3DFaceReductionOptions struct {
	win32.IInspectable
}

func (this *IPrinting3DFaceReductionOptions) Vtbl() *IPrinting3DFaceReductionOptionsVtbl {
	return (*IPrinting3DFaceReductionOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DFaceReductionOptions) Get_MaxReductionArea() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxReductionArea, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrinting3DFaceReductionOptions) Put_MaxReductionArea(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxReductionArea, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrinting3DFaceReductionOptions) Get_TargetTriangleCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TargetTriangleCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrinting3DFaceReductionOptions) Put_TargetTriangleCount(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TargetTriangleCount, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrinting3DFaceReductionOptions) Get_MaxEdgeLength() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxEdgeLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrinting3DFaceReductionOptions) Put_MaxEdgeLength(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxEdgeLength, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 378DB256-ED62-4952-B85B-03567D7C465E
var IID_IPrinting3DMaterial = syscall.GUID{0x378DB256, 0xED62, 0x4952,
	[8]byte{0xB8, 0x5B, 0x03, 0x56, 0x7D, 0x7C, 0x46, 0x5E}}

type IPrinting3DMaterialInterface interface {
	win32.IInspectableInterface
	Get_BaseGroups() *IVector[*IPrinting3DBaseMaterialGroup]
	Get_ColorGroups() *IVector[*IPrinting3DColorMaterialGroup]
	Get_Texture2CoordGroups() *IVector[*IPrinting3DTexture2CoordMaterialGroup]
	Get_CompositeGroups() *IVector[*IPrinting3DCompositeMaterialGroup]
	Get_MultiplePropertyGroups() *IVector[*IPrinting3DMultiplePropertyMaterialGroup]
}

type IPrinting3DMaterialVtbl struct {
	win32.IInspectableVtbl
	Get_BaseGroups             uintptr
	Get_ColorGroups            uintptr
	Get_Texture2CoordGroups    uintptr
	Get_CompositeGroups        uintptr
	Get_MultiplePropertyGroups uintptr
}

type IPrinting3DMaterial struct {
	win32.IInspectable
}

func (this *IPrinting3DMaterial) Vtbl() *IPrinting3DMaterialVtbl {
	return (*IPrinting3DMaterialVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DMaterial) Get_BaseGroups() *IVector[*IPrinting3DBaseMaterialGroup] {
	var _result *IVector[*IPrinting3DBaseMaterialGroup]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BaseGroups, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DMaterial) Get_ColorGroups() *IVector[*IPrinting3DColorMaterialGroup] {
	var _result *IVector[*IPrinting3DColorMaterialGroup]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ColorGroups, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DMaterial) Get_Texture2CoordGroups() *IVector[*IPrinting3DTexture2CoordMaterialGroup] {
	var _result *IVector[*IPrinting3DTexture2CoordMaterialGroup]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Texture2CoordGroups, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DMaterial) Get_CompositeGroups() *IVector[*IPrinting3DCompositeMaterialGroup] {
	var _result *IVector[*IPrinting3DCompositeMaterialGroup]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CompositeGroups, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DMaterial) Get_MultiplePropertyGroups() *IVector[*IPrinting3DMultiplePropertyMaterialGroup] {
	var _result *IVector[*IPrinting3DMultiplePropertyMaterialGroup]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MultiplePropertyGroups, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 192E90DC-0228-2E01-BC20-C5290CBF32C4
var IID_IPrinting3DMesh = syscall.GUID{0x192E90DC, 0x0228, 0x2E01,
	[8]byte{0xBC, 0x20, 0xC5, 0x29, 0x0C, 0xBF, 0x32, 0xC4}}

type IPrinting3DMeshInterface interface {
	win32.IInspectableInterface
	Get_VertexCount() uint32
	Put_VertexCount(value uint32)
	Get_IndexCount() uint32
	Put_IndexCount(value uint32)
	Get_VertexPositionsDescription() Printing3DBufferDescription
	Put_VertexPositionsDescription(value Printing3DBufferDescription)
	Get_VertexNormalsDescription() Printing3DBufferDescription
	Put_VertexNormalsDescription(value Printing3DBufferDescription)
	Get_TriangleIndicesDescription() Printing3DBufferDescription
	Put_TriangleIndicesDescription(value Printing3DBufferDescription)
	Get_TriangleMaterialIndicesDescription() Printing3DBufferDescription
	Put_TriangleMaterialIndicesDescription(value Printing3DBufferDescription)
	GetVertexPositions() *IBuffer
	CreateVertexPositions(value uint32)
	GetVertexNormals() *IBuffer
	CreateVertexNormals(value uint32)
	GetTriangleIndices() *IBuffer
	CreateTriangleIndices(value uint32)
	GetTriangleMaterialIndices() *IBuffer
	CreateTriangleMaterialIndices(value uint32)
	Get_BufferDescriptionSet() *IPropertySet
	Get_BufferSet() *IPropertySet
	VerifyAsync(value Printing3DMeshVerificationMode) *IAsyncOperation[*IPrinting3DMeshVerificationResult]
}

type IPrinting3DMeshVtbl struct {
	win32.IInspectableVtbl
	Get_VertexCount                        uintptr
	Put_VertexCount                        uintptr
	Get_IndexCount                         uintptr
	Put_IndexCount                         uintptr
	Get_VertexPositionsDescription         uintptr
	Put_VertexPositionsDescription         uintptr
	Get_VertexNormalsDescription           uintptr
	Put_VertexNormalsDescription           uintptr
	Get_TriangleIndicesDescription         uintptr
	Put_TriangleIndicesDescription         uintptr
	Get_TriangleMaterialIndicesDescription uintptr
	Put_TriangleMaterialIndicesDescription uintptr
	GetVertexPositions                     uintptr
	CreateVertexPositions                  uintptr
	GetVertexNormals                       uintptr
	CreateVertexNormals                    uintptr
	GetTriangleIndices                     uintptr
	CreateTriangleIndices                  uintptr
	GetTriangleMaterialIndices             uintptr
	CreateTriangleMaterialIndices          uintptr
	Get_BufferDescriptionSet               uintptr
	Get_BufferSet                          uintptr
	VerifyAsync                            uintptr
}

type IPrinting3DMesh struct {
	win32.IInspectable
}

func (this *IPrinting3DMesh) Vtbl() *IPrinting3DMeshVtbl {
	return (*IPrinting3DMeshVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DMesh) Get_VertexCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VertexCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrinting3DMesh) Put_VertexCount(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_VertexCount, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrinting3DMesh) Get_IndexCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IndexCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrinting3DMesh) Put_IndexCount(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IndexCount, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrinting3DMesh) Get_VertexPositionsDescription() Printing3DBufferDescription {
	var _result Printing3DBufferDescription
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VertexPositionsDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrinting3DMesh) Put_VertexPositionsDescription(value Printing3DBufferDescription) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_VertexPositionsDescription, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IPrinting3DMesh) Get_VertexNormalsDescription() Printing3DBufferDescription {
	var _result Printing3DBufferDescription
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VertexNormalsDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrinting3DMesh) Put_VertexNormalsDescription(value Printing3DBufferDescription) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_VertexNormalsDescription, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IPrinting3DMesh) Get_TriangleIndicesDescription() Printing3DBufferDescription {
	var _result Printing3DBufferDescription
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TriangleIndicesDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrinting3DMesh) Put_TriangleIndicesDescription(value Printing3DBufferDescription) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TriangleIndicesDescription, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IPrinting3DMesh) Get_TriangleMaterialIndicesDescription() Printing3DBufferDescription {
	var _result Printing3DBufferDescription
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TriangleMaterialIndicesDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrinting3DMesh) Put_TriangleMaterialIndicesDescription(value Printing3DBufferDescription) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TriangleMaterialIndicesDescription, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IPrinting3DMesh) GetVertexPositions() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetVertexPositions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DMesh) CreateVertexPositions(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateVertexPositions, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrinting3DMesh) GetVertexNormals() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetVertexNormals, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DMesh) CreateVertexNormals(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateVertexNormals, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrinting3DMesh) GetTriangleIndices() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTriangleIndices, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DMesh) CreateTriangleIndices(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateTriangleIndices, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrinting3DMesh) GetTriangleMaterialIndices() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTriangleMaterialIndices, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DMesh) CreateTriangleMaterialIndices(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateTriangleMaterialIndices, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrinting3DMesh) Get_BufferDescriptionSet() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BufferDescriptionSet, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DMesh) Get_BufferSet() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BufferSet, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DMesh) VerifyAsync(value Printing3DMeshVerificationMode) *IAsyncOperation[*IPrinting3DMeshVerificationResult] {
	var _result *IAsyncOperation[*IPrinting3DMeshVerificationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().VerifyAsync, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 195671BA-E93A-4E8A-A46F-DEA8E852197E
var IID_IPrinting3DMeshVerificationResult = syscall.GUID{0x195671BA, 0xE93A, 0x4E8A,
	[8]byte{0xA4, 0x6F, 0xDE, 0xA8, 0xE8, 0x52, 0x19, 0x7E}}

type IPrinting3DMeshVerificationResultInterface interface {
	win32.IInspectableInterface
	Get_IsValid() bool
	Get_NonmanifoldTriangles() *IVectorView[uint32]
	Get_ReversedNormalTriangles() *IVectorView[uint32]
}

type IPrinting3DMeshVerificationResultVtbl struct {
	win32.IInspectableVtbl
	Get_IsValid                 uintptr
	Get_NonmanifoldTriangles    uintptr
	Get_ReversedNormalTriangles uintptr
}

type IPrinting3DMeshVerificationResult struct {
	win32.IInspectable
}

func (this *IPrinting3DMeshVerificationResult) Vtbl() *IPrinting3DMeshVerificationResultVtbl {
	return (*IPrinting3DMeshVerificationResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DMeshVerificationResult) Get_IsValid() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsValid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrinting3DMeshVerificationResult) Get_NonmanifoldTriangles() *IVectorView[uint32] {
	var _result *IVectorView[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NonmanifoldTriangles, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DMeshVerificationResult) Get_ReversedNormalTriangles() *IVectorView[uint32] {
	var _result *IVectorView[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReversedNormalTriangles, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2D012EF0-52FB-919A-77B0-4B1A3B80324F
var IID_IPrinting3DModel = syscall.GUID{0x2D012EF0, 0x52FB, 0x919A,
	[8]byte{0x77, 0xB0, 0x4B, 0x1A, 0x3B, 0x80, 0x32, 0x4F}}

type IPrinting3DModelInterface interface {
	win32.IInspectableInterface
	Get_Unit() Printing3DModelUnit
	Put_Unit(value Printing3DModelUnit)
	Get_Textures() *IVector[*IPrinting3DModelTexture]
	Get_Meshes() *IVector[*IPrinting3DMesh]
	Get_Components() *IVector[*IPrinting3DComponent]
	Get_Material() *IPrinting3DMaterial
	Put_Material(value *IPrinting3DMaterial)
	Get_Build() *IPrinting3DComponent
	Put_Build(value *IPrinting3DComponent)
	Get_Version() string
	Put_Version(value string)
	Get_RequiredExtensions() *IVector[string]
	Get_Metadata() *IMap[string, string]
	RepairAsync() *IAsyncAction
	Clone() *IPrinting3DModel
}

type IPrinting3DModelVtbl struct {
	win32.IInspectableVtbl
	Get_Unit               uintptr
	Put_Unit               uintptr
	Get_Textures           uintptr
	Get_Meshes             uintptr
	Get_Components         uintptr
	Get_Material           uintptr
	Put_Material           uintptr
	Get_Build              uintptr
	Put_Build              uintptr
	Get_Version            uintptr
	Put_Version            uintptr
	Get_RequiredExtensions uintptr
	Get_Metadata           uintptr
	RepairAsync            uintptr
	Clone                  uintptr
}

type IPrinting3DModel struct {
	win32.IInspectable
}

func (this *IPrinting3DModel) Vtbl() *IPrinting3DModelVtbl {
	return (*IPrinting3DModelVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DModel) Get_Unit() Printing3DModelUnit {
	var _result Printing3DModelUnit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Unit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrinting3DModel) Put_Unit(value Printing3DModelUnit) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Unit, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrinting3DModel) Get_Textures() *IVector[*IPrinting3DModelTexture] {
	var _result *IVector[*IPrinting3DModelTexture]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Textures, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DModel) Get_Meshes() *IVector[*IPrinting3DMesh] {
	var _result *IVector[*IPrinting3DMesh]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Meshes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DModel) Get_Components() *IVector[*IPrinting3DComponent] {
	var _result *IVector[*IPrinting3DComponent]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Components, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DModel) Get_Material() *IPrinting3DMaterial {
	var _result *IPrinting3DMaterial
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Material, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DModel) Put_Material(value *IPrinting3DMaterial) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Material, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPrinting3DModel) Get_Build() *IPrinting3DComponent {
	var _result *IPrinting3DComponent
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Build, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DModel) Put_Build(value *IPrinting3DComponent) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Build, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPrinting3DModel) Get_Version() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Version, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPrinting3DModel) Put_Version(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Version, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPrinting3DModel) Get_RequiredExtensions() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequiredExtensions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DModel) Get_Metadata() *IMap[string, string] {
	var _result *IMap[string, string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Metadata, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DModel) RepairAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RepairAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DModel) Clone() *IPrinting3DModel {
	var _result *IPrinting3DModel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C92069C7-C841-47F3-A84E-A149FD08B657
var IID_IPrinting3DModel2 = syscall.GUID{0xC92069C7, 0xC841, 0x47F3,
	[8]byte{0xA8, 0x4E, 0xA1, 0x49, 0xFD, 0x08, 0xB6, 0x57}}

type IPrinting3DModel2Interface interface {
	win32.IInspectableInterface
	TryPartialRepairAsync() *IAsyncOperation[bool]
	TryPartialRepairWithTimeAsync(maxWaitTime TimeSpan) *IAsyncOperation[bool]
	TryReduceFacesAsync() *IAsyncOperationWithProgress[bool, float64]
	TryReduceFacesWithOptionsAsync(printing3DFaceReductionOptions *IPrinting3DFaceReductionOptions) *IAsyncOperationWithProgress[bool, float64]
	TryReduceFacesWithOptionsAndTimeAsync(printing3DFaceReductionOptions *IPrinting3DFaceReductionOptions, maxWait TimeSpan) *IAsyncOperationWithProgress[bool, float64]
	RepairWithProgressAsync() *IAsyncOperationWithProgress[bool, float64]
}

type IPrinting3DModel2Vtbl struct {
	win32.IInspectableVtbl
	TryPartialRepairAsync                 uintptr
	TryPartialRepairWithTimeAsync         uintptr
	TryReduceFacesAsync                   uintptr
	TryReduceFacesWithOptionsAsync        uintptr
	TryReduceFacesWithOptionsAndTimeAsync uintptr
	RepairWithProgressAsync               uintptr
}

type IPrinting3DModel2 struct {
	win32.IInspectable
}

func (this *IPrinting3DModel2) Vtbl() *IPrinting3DModel2Vtbl {
	return (*IPrinting3DModel2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DModel2) TryPartialRepairAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryPartialRepairAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DModel2) TryPartialRepairWithTimeAsync(maxWaitTime TimeSpan) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryPartialRepairWithTimeAsync, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&maxWaitTime)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DModel2) TryReduceFacesAsync() *IAsyncOperationWithProgress[bool, float64] {
	var _result *IAsyncOperationWithProgress[bool, float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryReduceFacesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DModel2) TryReduceFacesWithOptionsAsync(printing3DFaceReductionOptions *IPrinting3DFaceReductionOptions) *IAsyncOperationWithProgress[bool, float64] {
	var _result *IAsyncOperationWithProgress[bool, float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryReduceFacesWithOptionsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(printing3DFaceReductionOptions)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DModel2) TryReduceFacesWithOptionsAndTimeAsync(printing3DFaceReductionOptions *IPrinting3DFaceReductionOptions, maxWait TimeSpan) *IAsyncOperationWithProgress[bool, float64] {
	var _result *IAsyncOperationWithProgress[bool, float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryReduceFacesWithOptionsAndTimeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(printing3DFaceReductionOptions)), *(*uintptr)(unsafe.Pointer(&maxWait)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DModel2) RepairWithProgressAsync() *IAsyncOperationWithProgress[bool, float64] {
	var _result *IAsyncOperationWithProgress[bool, float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RepairWithProgressAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5DAFCF01-B59D-483C-97BB-A4D546D1C75C
var IID_IPrinting3DModelTexture = syscall.GUID{0x5DAFCF01, 0xB59D, 0x483C,
	[8]byte{0x97, 0xBB, 0xA4, 0xD5, 0x46, 0xD1, 0xC7, 0x5C}}

type IPrinting3DModelTextureInterface interface {
	win32.IInspectableInterface
	Get_TextureResource() *IPrinting3DTextureResource
	Put_TextureResource(value *IPrinting3DTextureResource)
	Get_TileStyleU() Printing3DTextureEdgeBehavior
	Put_TileStyleU(value Printing3DTextureEdgeBehavior)
	Get_TileStyleV() Printing3DTextureEdgeBehavior
	Put_TileStyleV(value Printing3DTextureEdgeBehavior)
}

type IPrinting3DModelTextureVtbl struct {
	win32.IInspectableVtbl
	Get_TextureResource uintptr
	Put_TextureResource uintptr
	Get_TileStyleU      uintptr
	Put_TileStyleU      uintptr
	Get_TileStyleV      uintptr
	Put_TileStyleV      uintptr
}

type IPrinting3DModelTexture struct {
	win32.IInspectable
}

func (this *IPrinting3DModelTexture) Vtbl() *IPrinting3DModelTextureVtbl {
	return (*IPrinting3DModelTextureVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DModelTexture) Get_TextureResource() *IPrinting3DTextureResource {
	var _result *IPrinting3DTextureResource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TextureResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DModelTexture) Put_TextureResource(value *IPrinting3DTextureResource) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TextureResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPrinting3DModelTexture) Get_TileStyleU() Printing3DTextureEdgeBehavior {
	var _result Printing3DTextureEdgeBehavior
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TileStyleU, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrinting3DModelTexture) Put_TileStyleU(value Printing3DTextureEdgeBehavior) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TileStyleU, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrinting3DModelTexture) Get_TileStyleV() Printing3DTextureEdgeBehavior {
	var _result Printing3DTextureEdgeBehavior
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TileStyleV, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrinting3DModelTexture) Put_TileStyleV(value Printing3DTextureEdgeBehavior) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TileStyleV, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 25A6254B-C6E9-484D-A214-A25E5776BA62
var IID_IPrinting3DMultiplePropertyMaterial = syscall.GUID{0x25A6254B, 0xC6E9, 0x484D,
	[8]byte{0xA2, 0x14, 0xA2, 0x5E, 0x57, 0x76, 0xBA, 0x62}}

type IPrinting3DMultiplePropertyMaterialInterface interface {
	win32.IInspectableInterface
	Get_MaterialIndices() *IVector[uint32]
}

type IPrinting3DMultiplePropertyMaterialVtbl struct {
	win32.IInspectableVtbl
	Get_MaterialIndices uintptr
}

type IPrinting3DMultiplePropertyMaterial struct {
	win32.IInspectable
}

func (this *IPrinting3DMultiplePropertyMaterial) Vtbl() *IPrinting3DMultiplePropertyMaterialVtbl {
	return (*IPrinting3DMultiplePropertyMaterialVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DMultiplePropertyMaterial) Get_MaterialIndices() *IVector[uint32] {
	var _result *IVector[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaterialIndices, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F0950519-AEB9-4515-A39B-A088FBBB277C
var IID_IPrinting3DMultiplePropertyMaterialGroup = syscall.GUID{0xF0950519, 0xAEB9, 0x4515,
	[8]byte{0xA3, 0x9B, 0xA0, 0x88, 0xFB, 0xBB, 0x27, 0x7C}}

type IPrinting3DMultiplePropertyMaterialGroupInterface interface {
	win32.IInspectableInterface
	Get_MultipleProperties() *IVector[*IPrinting3DMultiplePropertyMaterial]
	Get_MaterialGroupIndices() *IVector[uint32]
	Get_MaterialGroupId() uint32
}

type IPrinting3DMultiplePropertyMaterialGroupVtbl struct {
	win32.IInspectableVtbl
	Get_MultipleProperties   uintptr
	Get_MaterialGroupIndices uintptr
	Get_MaterialGroupId      uintptr
}

type IPrinting3DMultiplePropertyMaterialGroup struct {
	win32.IInspectable
}

func (this *IPrinting3DMultiplePropertyMaterialGroup) Vtbl() *IPrinting3DMultiplePropertyMaterialGroupVtbl {
	return (*IPrinting3DMultiplePropertyMaterialGroupVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DMultiplePropertyMaterialGroup) Get_MultipleProperties() *IVector[*IPrinting3DMultiplePropertyMaterial] {
	var _result *IVector[*IPrinting3DMultiplePropertyMaterial]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MultipleProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DMultiplePropertyMaterialGroup) Get_MaterialGroupIndices() *IVector[uint32] {
	var _result *IVector[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaterialGroupIndices, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DMultiplePropertyMaterialGroup) Get_MaterialGroupId() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaterialGroupId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 323E196E-D4C6-451E-A814-4D78A210FE53
var IID_IPrinting3DMultiplePropertyMaterialGroupFactory = syscall.GUID{0x323E196E, 0xD4C6, 0x451E,
	[8]byte{0xA8, 0x14, 0x4D, 0x78, 0xA2, 0x10, 0xFE, 0x53}}

type IPrinting3DMultiplePropertyMaterialGroupFactoryInterface interface {
	win32.IInspectableInterface
	Create(MaterialGroupId uint32) *IPrinting3DMultiplePropertyMaterialGroup
}

type IPrinting3DMultiplePropertyMaterialGroupFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IPrinting3DMultiplePropertyMaterialGroupFactory struct {
	win32.IInspectable
}

func (this *IPrinting3DMultiplePropertyMaterialGroupFactory) Vtbl() *IPrinting3DMultiplePropertyMaterialGroupFactoryVtbl {
	return (*IPrinting3DMultiplePropertyMaterialGroupFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DMultiplePropertyMaterialGroupFactory) Create(MaterialGroupId uint32) *IPrinting3DMultiplePropertyMaterialGroup {
	var _result *IPrinting3DMultiplePropertyMaterialGroup
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(MaterialGroupId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8D844BFB-07E9-4986-9833-8DD3D48C6859
var IID_IPrinting3DTexture2CoordMaterial = syscall.GUID{0x8D844BFB, 0x07E9, 0x4986,
	[8]byte{0x98, 0x33, 0x8D, 0xD3, 0xD4, 0x8C, 0x68, 0x59}}

type IPrinting3DTexture2CoordMaterialInterface interface {
	win32.IInspectableInterface
	Get_Texture() *IPrinting3DModelTexture
	Put_Texture(value *IPrinting3DModelTexture)
	Get_U() float64
	Put_U(value float64)
	Get_V() float64
	Put_V(value float64)
}

type IPrinting3DTexture2CoordMaterialVtbl struct {
	win32.IInspectableVtbl
	Get_Texture uintptr
	Put_Texture uintptr
	Get_U       uintptr
	Put_U       uintptr
	Get_V       uintptr
	Put_V       uintptr
}

type IPrinting3DTexture2CoordMaterial struct {
	win32.IInspectable
}

func (this *IPrinting3DTexture2CoordMaterial) Vtbl() *IPrinting3DTexture2CoordMaterialVtbl {
	return (*IPrinting3DTexture2CoordMaterialVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DTexture2CoordMaterial) Get_Texture() *IPrinting3DModelTexture {
	var _result *IPrinting3DModelTexture
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Texture, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DTexture2CoordMaterial) Put_Texture(value *IPrinting3DModelTexture) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Texture, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPrinting3DTexture2CoordMaterial) Get_U() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_U, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrinting3DTexture2CoordMaterial) Put_U(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_U, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrinting3DTexture2CoordMaterial) Get_V() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_V, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrinting3DTexture2CoordMaterial) Put_V(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_V, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 627D7CA7-6D90-4FB9-9FC4-9FEFF3DFA892
var IID_IPrinting3DTexture2CoordMaterialGroup = syscall.GUID{0x627D7CA7, 0x6D90, 0x4FB9,
	[8]byte{0x9F, 0xC4, 0x9F, 0xEF, 0xF3, 0xDF, 0xA8, 0x92}}

type IPrinting3DTexture2CoordMaterialGroupInterface interface {
	win32.IInspectableInterface
	Get_Texture2Coords() *IVector[*IPrinting3DTexture2CoordMaterial]
	Get_MaterialGroupId() uint32
}

type IPrinting3DTexture2CoordMaterialGroupVtbl struct {
	win32.IInspectableVtbl
	Get_Texture2Coords  uintptr
	Get_MaterialGroupId uintptr
}

type IPrinting3DTexture2CoordMaterialGroup struct {
	win32.IInspectable
}

func (this *IPrinting3DTexture2CoordMaterialGroup) Vtbl() *IPrinting3DTexture2CoordMaterialGroupVtbl {
	return (*IPrinting3DTexture2CoordMaterialGroupVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DTexture2CoordMaterialGroup) Get_Texture2Coords() *IVector[*IPrinting3DTexture2CoordMaterial] {
	var _result *IVector[*IPrinting3DTexture2CoordMaterial]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Texture2Coords, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DTexture2CoordMaterialGroup) Get_MaterialGroupId() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaterialGroupId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 69FBDBBA-B12E-429B-8386-DF5284F6E80F
var IID_IPrinting3DTexture2CoordMaterialGroup2 = syscall.GUID{0x69FBDBBA, 0xB12E, 0x429B,
	[8]byte{0x83, 0x86, 0xDF, 0x52, 0x84, 0xF6, 0xE8, 0x0F}}

type IPrinting3DTexture2CoordMaterialGroup2Interface interface {
	win32.IInspectableInterface
	Get_Texture() *IPrinting3DModelTexture
	Put_Texture(value *IPrinting3DModelTexture)
}

type IPrinting3DTexture2CoordMaterialGroup2Vtbl struct {
	win32.IInspectableVtbl
	Get_Texture uintptr
	Put_Texture uintptr
}

type IPrinting3DTexture2CoordMaterialGroup2 struct {
	win32.IInspectable
}

func (this *IPrinting3DTexture2CoordMaterialGroup2) Vtbl() *IPrinting3DTexture2CoordMaterialGroup2Vtbl {
	return (*IPrinting3DTexture2CoordMaterialGroup2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DTexture2CoordMaterialGroup2) Get_Texture() *IPrinting3DModelTexture {
	var _result *IPrinting3DModelTexture
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Texture, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DTexture2CoordMaterialGroup2) Put_Texture(value *IPrinting3DModelTexture) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Texture, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// CBB049B0-468A-4C6F-B2A2-8EB8BA8DEA48
var IID_IPrinting3DTexture2CoordMaterialGroupFactory = syscall.GUID{0xCBB049B0, 0x468A, 0x4C6F,
	[8]byte{0xB2, 0xA2, 0x8E, 0xB8, 0xBA, 0x8D, 0xEA, 0x48}}

type IPrinting3DTexture2CoordMaterialGroupFactoryInterface interface {
	win32.IInspectableInterface
	Create(MaterialGroupId uint32) *IPrinting3DTexture2CoordMaterialGroup
}

type IPrinting3DTexture2CoordMaterialGroupFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IPrinting3DTexture2CoordMaterialGroupFactory struct {
	win32.IInspectable
}

func (this *IPrinting3DTexture2CoordMaterialGroupFactory) Vtbl() *IPrinting3DTexture2CoordMaterialGroupFactoryVtbl {
	return (*IPrinting3DTexture2CoordMaterialGroupFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DTexture2CoordMaterialGroupFactory) Create(MaterialGroupId uint32) *IPrinting3DTexture2CoordMaterialGroup {
	var _result *IPrinting3DTexture2CoordMaterialGroup
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(MaterialGroupId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A70DF32D-6AB1-44AE-BC45-A27382C0D38C
var IID_IPrinting3DTextureResource = syscall.GUID{0xA70DF32D, 0x6AB1, 0x44AE,
	[8]byte{0xBC, 0x45, 0xA2, 0x73, 0x82, 0xC0, 0xD3, 0x8C}}

type IPrinting3DTextureResourceInterface interface {
	win32.IInspectableInterface
	Get_TextureData() *IRandomAccessStreamWithContentType
	Put_TextureData(value *IRandomAccessStreamWithContentType)
	Get_Name() string
	Put_Name(value string)
}

type IPrinting3DTextureResourceVtbl struct {
	win32.IInspectableVtbl
	Get_TextureData uintptr
	Put_TextureData uintptr
	Get_Name        uintptr
	Put_Name        uintptr
}

type IPrinting3DTextureResource struct {
	win32.IInspectable
}

func (this *IPrinting3DTextureResource) Vtbl() *IPrinting3DTextureResourceVtbl {
	return (*IPrinting3DTextureResourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrinting3DTextureResource) Get_TextureData() *IRandomAccessStreamWithContentType {
	var _result *IRandomAccessStreamWithContentType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TextureData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrinting3DTextureResource) Put_TextureData(value *IRandomAccessStreamWithContentType) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TextureData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPrinting3DTextureResource) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPrinting3DTextureResource) Put_Name(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Name, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// classes

type Print3DManager struct {
	RtClass
	*IPrint3DManager
}

func NewIPrint3DManagerStatics() *IPrint3DManagerStatics {
	var p *IPrint3DManagerStatics
	hs := NewHStr("Windows.Graphics.Printing3D.Print3DManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPrint3DManagerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type Print3DTask struct {
	RtClass
	*IPrint3DTask
}

type Print3DTaskCompletedEventArgs struct {
	RtClass
	*IPrint3DTaskCompletedEventArgs
}

type Print3DTaskRequest struct {
	RtClass
	*IPrint3DTaskRequest
}

type Print3DTaskSourceChangedEventArgs struct {
	RtClass
	*IPrint3DTaskSourceChangedEventArgs
}

type Print3DTaskSourceRequestedArgs struct {
	RtClass
	*IPrint3DTaskSourceRequestedArgs
}

type Printing3D3MFPackage struct {
	RtClass
	*IPrinting3D3MFPackage
}

func NewPrinting3D3MFPackage() *Printing3D3MFPackage {
	hs := NewHStr("Windows.Graphics.Printing3D.Printing3D3MFPackage")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &Printing3D3MFPackage{
		RtClass:               RtClass{PInspect: p},
		IPrinting3D3MFPackage: (*IPrinting3D3MFPackage)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewIPrinting3D3MFPackageStatics() *IPrinting3D3MFPackageStatics {
	var p *IPrinting3D3MFPackageStatics
	hs := NewHStr("Windows.Graphics.Printing3D.Printing3D3MFPackage")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPrinting3D3MFPackageStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type Printing3DBaseMaterial struct {
	RtClass
	*IPrinting3DBaseMaterial
}

func NewPrinting3DBaseMaterial() *Printing3DBaseMaterial {
	hs := NewHStr("Windows.Graphics.Printing3D.Printing3DBaseMaterial")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &Printing3DBaseMaterial{
		RtClass:                 RtClass{PInspect: p},
		IPrinting3DBaseMaterial: (*IPrinting3DBaseMaterial)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewIPrinting3DBaseMaterialStatics() *IPrinting3DBaseMaterialStatics {
	var p *IPrinting3DBaseMaterialStatics
	hs := NewHStr("Windows.Graphics.Printing3D.Printing3DBaseMaterial")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPrinting3DBaseMaterialStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type Printing3DBaseMaterialGroup struct {
	RtClass
	*IPrinting3DBaseMaterialGroup
}

func NewPrinting3DBaseMaterialGroup_Create(MaterialGroupId uint32) *Printing3DBaseMaterialGroup {
	hs := NewHStr("Windows.Graphics.Printing3D.Printing3DBaseMaterialGroup")
	var pFac *IPrinting3DBaseMaterialGroupFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPrinting3DBaseMaterialGroupFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IPrinting3DBaseMaterialGroup
	p = pFac.Create(MaterialGroupId)
	result := &Printing3DBaseMaterialGroup{
		RtClass:                      RtClass{PInspect: &p.IInspectable},
		IPrinting3DBaseMaterialGroup: p,
	}
	com.AddToScope(result)
	return result
}

type Printing3DColorMaterial struct {
	RtClass
	*IPrinting3DColorMaterial
}

func NewPrinting3DColorMaterial() *Printing3DColorMaterial {
	hs := NewHStr("Windows.Graphics.Printing3D.Printing3DColorMaterial")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &Printing3DColorMaterial{
		RtClass:                  RtClass{PInspect: p},
		IPrinting3DColorMaterial: (*IPrinting3DColorMaterial)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type Printing3DColorMaterialGroup struct {
	RtClass
	*IPrinting3DColorMaterialGroup
}

func NewPrinting3DColorMaterialGroup_Create(MaterialGroupId uint32) *Printing3DColorMaterialGroup {
	hs := NewHStr("Windows.Graphics.Printing3D.Printing3DColorMaterialGroup")
	var pFac *IPrinting3DColorMaterialGroupFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPrinting3DColorMaterialGroupFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IPrinting3DColorMaterialGroup
	p = pFac.Create(MaterialGroupId)
	result := &Printing3DColorMaterialGroup{
		RtClass:                       RtClass{PInspect: &p.IInspectable},
		IPrinting3DColorMaterialGroup: p,
	}
	com.AddToScope(result)
	return result
}

type Printing3DComponent struct {
	RtClass
	*IPrinting3DComponent
}

func NewPrinting3DComponent() *Printing3DComponent {
	hs := NewHStr("Windows.Graphics.Printing3D.Printing3DComponent")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &Printing3DComponent{
		RtClass:              RtClass{PInspect: p},
		IPrinting3DComponent: (*IPrinting3DComponent)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type Printing3DComponentWithMatrix struct {
	RtClass
	*IPrinting3DComponentWithMatrix
}

func NewPrinting3DComponentWithMatrix() *Printing3DComponentWithMatrix {
	hs := NewHStr("Windows.Graphics.Printing3D.Printing3DComponentWithMatrix")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &Printing3DComponentWithMatrix{
		RtClass:                        RtClass{PInspect: p},
		IPrinting3DComponentWithMatrix: (*IPrinting3DComponentWithMatrix)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type Printing3DCompositeMaterial struct {
	RtClass
	*IPrinting3DCompositeMaterial
}

func NewPrinting3DCompositeMaterial() *Printing3DCompositeMaterial {
	hs := NewHStr("Windows.Graphics.Printing3D.Printing3DCompositeMaterial")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &Printing3DCompositeMaterial{
		RtClass:                      RtClass{PInspect: p},
		IPrinting3DCompositeMaterial: (*IPrinting3DCompositeMaterial)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type Printing3DCompositeMaterialGroup struct {
	RtClass
	*IPrinting3DCompositeMaterialGroup
}

func NewPrinting3DCompositeMaterialGroup_Create(MaterialGroupId uint32) *Printing3DCompositeMaterialGroup {
	hs := NewHStr("Windows.Graphics.Printing3D.Printing3DCompositeMaterialGroup")
	var pFac *IPrinting3DCompositeMaterialGroupFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPrinting3DCompositeMaterialGroupFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IPrinting3DCompositeMaterialGroup
	p = pFac.Create(MaterialGroupId)
	result := &Printing3DCompositeMaterialGroup{
		RtClass:                           RtClass{PInspect: &p.IInspectable},
		IPrinting3DCompositeMaterialGroup: p,
	}
	com.AddToScope(result)
	return result
}

type Printing3DFaceReductionOptions struct {
	RtClass
	*IPrinting3DFaceReductionOptions
}

func NewPrinting3DFaceReductionOptions() *Printing3DFaceReductionOptions {
	hs := NewHStr("Windows.Graphics.Printing3D.Printing3DFaceReductionOptions")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &Printing3DFaceReductionOptions{
		RtClass:                         RtClass{PInspect: p},
		IPrinting3DFaceReductionOptions: (*IPrinting3DFaceReductionOptions)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type Printing3DMaterial struct {
	RtClass
	*IPrinting3DMaterial
}

func NewPrinting3DMaterial() *Printing3DMaterial {
	hs := NewHStr("Windows.Graphics.Printing3D.Printing3DMaterial")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &Printing3DMaterial{
		RtClass:             RtClass{PInspect: p},
		IPrinting3DMaterial: (*IPrinting3DMaterial)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type Printing3DMesh struct {
	RtClass
	*IPrinting3DMesh
}

func NewPrinting3DMesh() *Printing3DMesh {
	hs := NewHStr("Windows.Graphics.Printing3D.Printing3DMesh")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &Printing3DMesh{
		RtClass:         RtClass{PInspect: p},
		IPrinting3DMesh: (*IPrinting3DMesh)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type Printing3DMeshVerificationResult struct {
	RtClass
	*IPrinting3DMeshVerificationResult
}

type Printing3DModel struct {
	RtClass
	*IPrinting3DModel
}

func NewPrinting3DModel() *Printing3DModel {
	hs := NewHStr("Windows.Graphics.Printing3D.Printing3DModel")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &Printing3DModel{
		RtClass:          RtClass{PInspect: p},
		IPrinting3DModel: (*IPrinting3DModel)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type Printing3DModelTexture struct {
	RtClass
	*IPrinting3DModelTexture
}

func NewPrinting3DModelTexture() *Printing3DModelTexture {
	hs := NewHStr("Windows.Graphics.Printing3D.Printing3DModelTexture")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &Printing3DModelTexture{
		RtClass:                 RtClass{PInspect: p},
		IPrinting3DModelTexture: (*IPrinting3DModelTexture)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type Printing3DMultiplePropertyMaterial struct {
	RtClass
	*IPrinting3DMultiplePropertyMaterial
}

func NewPrinting3DMultiplePropertyMaterial() *Printing3DMultiplePropertyMaterial {
	hs := NewHStr("Windows.Graphics.Printing3D.Printing3DMultiplePropertyMaterial")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &Printing3DMultiplePropertyMaterial{
		RtClass:                             RtClass{PInspect: p},
		IPrinting3DMultiplePropertyMaterial: (*IPrinting3DMultiplePropertyMaterial)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type Printing3DMultiplePropertyMaterialGroup struct {
	RtClass
	*IPrinting3DMultiplePropertyMaterialGroup
}

func NewPrinting3DMultiplePropertyMaterialGroup_Create(MaterialGroupId uint32) *Printing3DMultiplePropertyMaterialGroup {
	hs := NewHStr("Windows.Graphics.Printing3D.Printing3DMultiplePropertyMaterialGroup")
	var pFac *IPrinting3DMultiplePropertyMaterialGroupFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPrinting3DMultiplePropertyMaterialGroupFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IPrinting3DMultiplePropertyMaterialGroup
	p = pFac.Create(MaterialGroupId)
	result := &Printing3DMultiplePropertyMaterialGroup{
		RtClass:                                  RtClass{PInspect: &p.IInspectable},
		IPrinting3DMultiplePropertyMaterialGroup: p,
	}
	com.AddToScope(result)
	return result
}

type Printing3DTexture2CoordMaterial struct {
	RtClass
	*IPrinting3DTexture2CoordMaterial
}

func NewPrinting3DTexture2CoordMaterial() *Printing3DTexture2CoordMaterial {
	hs := NewHStr("Windows.Graphics.Printing3D.Printing3DTexture2CoordMaterial")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &Printing3DTexture2CoordMaterial{
		RtClass:                          RtClass{PInspect: p},
		IPrinting3DTexture2CoordMaterial: (*IPrinting3DTexture2CoordMaterial)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type Printing3DTexture2CoordMaterialGroup struct {
	RtClass
	*IPrinting3DTexture2CoordMaterialGroup
}

func NewPrinting3DTexture2CoordMaterialGroup_Create(MaterialGroupId uint32) *Printing3DTexture2CoordMaterialGroup {
	hs := NewHStr("Windows.Graphics.Printing3D.Printing3DTexture2CoordMaterialGroup")
	var pFac *IPrinting3DTexture2CoordMaterialGroupFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPrinting3DTexture2CoordMaterialGroupFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IPrinting3DTexture2CoordMaterialGroup
	p = pFac.Create(MaterialGroupId)
	result := &Printing3DTexture2CoordMaterialGroup{
		RtClass:                               RtClass{PInspect: &p.IInspectable},
		IPrinting3DTexture2CoordMaterialGroup: p,
	}
	com.AddToScope(result)
	return result
}

type Printing3DTextureResource struct {
	RtClass
	*IPrinting3DTextureResource
}

func NewPrinting3DTextureResource() *Printing3DTextureResource {
	hs := NewHStr("Windows.Graphics.Printing3D.Printing3DTextureResource")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &Printing3DTextureResource{
		RtClass:                    RtClass{PInspect: p},
		IPrinting3DTextureResource: (*IPrinting3DTextureResource)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}
