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
type AsyncStatus int32

const (
	AsyncStatus_Canceled  AsyncStatus = 2
	AsyncStatus_Completed AsyncStatus = 1
	AsyncStatus_Error     AsyncStatus = 3
	AsyncStatus_Started   AsyncStatus = 0
)

// enum
type PropertyType int32

const (
	PropertyType_Empty            PropertyType = 0
	PropertyType_UInt8            PropertyType = 1
	PropertyType_Int16            PropertyType = 2
	PropertyType_UInt16           PropertyType = 3
	PropertyType_Int32            PropertyType = 4
	PropertyType_UInt32           PropertyType = 5
	PropertyType_Int64            PropertyType = 6
	PropertyType_UInt64           PropertyType = 7
	PropertyType_Single           PropertyType = 8
	PropertyType_Double           PropertyType = 9
	PropertyType_Char16           PropertyType = 10
	PropertyType_Boolean          PropertyType = 11
	PropertyType_String           PropertyType = 12
	PropertyType_Inspectable      PropertyType = 13
	PropertyType_DateTime         PropertyType = 14
	PropertyType_TimeSpan         PropertyType = 15
	PropertyType_Guid             PropertyType = 16
	PropertyType_Point            PropertyType = 17
	PropertyType_Size             PropertyType = 18
	PropertyType_Rect             PropertyType = 19
	PropertyType_OtherType        PropertyType = 20
	PropertyType_UInt8Array       PropertyType = 1025
	PropertyType_Int16Array       PropertyType = 1026
	PropertyType_UInt16Array      PropertyType = 1027
	PropertyType_Int32Array       PropertyType = 1028
	PropertyType_UInt32Array      PropertyType = 1029
	PropertyType_Int64Array       PropertyType = 1030
	PropertyType_UInt64Array      PropertyType = 1031
	PropertyType_SingleArray      PropertyType = 1032
	PropertyType_DoubleArray      PropertyType = 1033
	PropertyType_Char16Array      PropertyType = 1034
	PropertyType_BooleanArray     PropertyType = 1035
	PropertyType_StringArray      PropertyType = 1036
	PropertyType_InspectableArray PropertyType = 1037
	PropertyType_DateTimeArray    PropertyType = 1038
	PropertyType_TimeSpanArray    PropertyType = 1039
	PropertyType_GuidArray        PropertyType = 1040
	PropertyType_PointArray       PropertyType = 1041
	PropertyType_SizeArray        PropertyType = 1042
	PropertyType_RectArray        PropertyType = 1043
	PropertyType_OtherTypeArray   PropertyType = 1044
)

// structs

type DateTime struct {
	UniversalTime int64
}

type EventRegistrationToken struct {
	Value int64
}

type FoundationContract struct {
}

type HResult struct {
	Value int32
}

type Point struct {
	X float32
	Y float32
}

type Rect struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

type Size struct {
	Width  float32
	Height float32
}

type TimeSpan struct {
	Duration int64
}

type UniversalApiContract struct {
}

// func types

//A4ED5C81-76C9-40BD-8BE6-B1D90FB20AE7
type AsyncActionCompletedHandler func(asyncInfo *IAsyncAction, asyncStatus AsyncStatus) com.Error

//6D844858-0CFF-4590-AE89-95A5A5C8B4B8
type AsyncActionProgressHandler[TProgress any] func(asyncInfo *IAsyncActionWithProgress[TProgress], progressInfo TProgress) com.Error

//9C029F91-CC84-44FD-AC26-0A6C4E555281
type AsyncActionWithProgressCompletedHandler[TProgress any] func(asyncInfo *IAsyncActionWithProgress[TProgress], asyncStatus AsyncStatus) com.Error

//FCDCF02C-E5D8-4478-915A-4D90B74B83A5
type AsyncOperationCompletedHandler[TResult any] func(asyncInfo *IAsyncOperation[TResult], asyncStatus AsyncStatus) com.Error

//55690902-0AAB-421A-8778-F8CE5026D758
type AsyncOperationProgressHandler[TResult any, TProgress any] func(asyncInfo *IAsyncOperationWithProgress[TResult, TProgress], progressInfo TProgress) com.Error

//E85DF41D-6AA7-46E3-A8E2-F009D840C627
type AsyncOperationWithProgressCompletedHandler[TResult any, TProgress any] func(asyncInfo *IAsyncOperationWithProgress[TResult, TProgress], asyncStatus AsyncStatus) com.Error

//ED32A372-F3C8-4FAA-9CFB-470148DA3888
type DeferralCompletedHandler func() com.Error

//9DE1C535-6AE1-11E0-84E1-18A905BCC53F
type EventHandler[T any] func(sender interface{}, args T) com.Error

//9DE1C534-6AE1-11E0-84E1-18A905BCC53F
type TypedEventHandler[TSender any, TResult any] func(sender TSender, args TResult) com.Error

// interfaces

// 5A648006-843A-4DA9-865B-9D26E5DFAD7B
var IID_IAsyncAction = syscall.GUID{0x5A648006, 0x843A, 0x4DA9,
	[8]byte{0x86, 0x5B, 0x9D, 0x26, 0xE5, 0xDF, 0xAD, 0x7B}}

type IAsyncActionInterface interface {
	win32.IInspectableInterface
	Put_Completed(handler AsyncActionCompletedHandler)
	Get_Completed() AsyncActionCompletedHandler
	GetResults()
}

type IAsyncActionVtbl struct {
	win32.IInspectableVtbl
	Put_Completed uintptr
	Get_Completed uintptr
	GetResults    uintptr
}

type IAsyncAction struct {
	win32.IInspectable
}

func (this *IAsyncAction) Vtbl() *IAsyncActionVtbl {
	return (*IAsyncActionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAsyncAction) Put_Completed(handler AsyncActionCompletedHandler) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Completed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))))
	_ = _hr
}

func (this *IAsyncAction) Get_Completed() AsyncActionCompletedHandler {
	var _result AsyncActionCompletedHandler
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Completed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAsyncAction) GetResults() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetResults, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 1F6DB258-E803-48A1-9546-EB7353398884
var IID_IAsyncActionWithProgress = syscall.GUID{0x1F6DB258, 0xE803, 0x48A1,
	[8]byte{0x95, 0x46, 0xEB, 0x73, 0x53, 0x39, 0x88, 0x84}}

type IAsyncActionWithProgressInterface[TProgress any] interface {
	win32.IInspectableInterface
	Put_Progress(handler AsyncActionProgressHandler[TProgress])
	Get_Progress() AsyncActionProgressHandler[TProgress]
	Put_Completed(handler AsyncActionWithProgressCompletedHandler[TProgress])
	Get_Completed() AsyncActionWithProgressCompletedHandler[TProgress]
	GetResults()
}

type IAsyncActionWithProgressVtbl struct {
	win32.IInspectableVtbl
	Put_Progress  uintptr
	Get_Progress  uintptr
	Put_Completed uintptr
	Get_Completed uintptr
	GetResults    uintptr
}

type IAsyncActionWithProgress[TProgress any] struct {
	win32.IInspectable
}

func (this *IAsyncActionWithProgress[TProgress]) Vtbl() *IAsyncActionWithProgressVtbl {
	return (*IAsyncActionWithProgressVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAsyncActionWithProgress[TProgress]) Put_Progress(handler AsyncActionProgressHandler[TProgress]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Progress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))))
	_ = _hr
}

func (this *IAsyncActionWithProgress[TProgress]) Get_Progress() AsyncActionProgressHandler[TProgress] {
	var _result AsyncActionProgressHandler[TProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Progress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAsyncActionWithProgress[TProgress]) Put_Completed(handler AsyncActionWithProgressCompletedHandler[TProgress]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Completed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))))
	_ = _hr
}

func (this *IAsyncActionWithProgress[TProgress]) Get_Completed() AsyncActionWithProgressCompletedHandler[TProgress] {
	var _result AsyncActionWithProgressCompletedHandler[TProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Completed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAsyncActionWithProgress[TProgress]) GetResults() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetResults, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 00000036-0000-0000-C000-000000000046
var IID_IAsyncInfo = syscall.GUID{0x00000036, 0x0000, 0x0000,
	[8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

type IAsyncInfoInterface interface {
	win32.IInspectableInterface
	Get_Id() uint32
	Get_Status() AsyncStatus
	Get_ErrorCode() HResult
	Cancel()
	Close()
}

type IAsyncInfoVtbl struct {
	win32.IInspectableVtbl
	Get_Id        uintptr
	Get_Status    uintptr
	Get_ErrorCode uintptr
	Cancel        uintptr
	Close         uintptr
}

type IAsyncInfo struct {
	win32.IInspectable
}

func (this *IAsyncInfo) Vtbl() *IAsyncInfoVtbl {
	return (*IAsyncInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAsyncInfo) Get_Id() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAsyncInfo) Get_Status() AsyncStatus {
	var _result AsyncStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAsyncInfo) Get_ErrorCode() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAsyncInfo) Cancel() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Cancel, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAsyncInfo) Close() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Close, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// B5D036D7-E297-498F-BA60-0289E76E23DD
var IID_IAsyncOperationWithProgress = syscall.GUID{0xB5D036D7, 0xE297, 0x498F,
	[8]byte{0xBA, 0x60, 0x02, 0x89, 0xE7, 0x6E, 0x23, 0xDD}}

type IAsyncOperationWithProgressInterface[TResult any, TProgress any] interface {
	win32.IInspectableInterface
	Put_Progress(handler AsyncOperationProgressHandler[TResult, TProgress])
	Get_Progress() AsyncOperationProgressHandler[TResult, TProgress]
	Put_Completed(handler AsyncOperationWithProgressCompletedHandler[TResult, TProgress])
	Get_Completed() AsyncOperationWithProgressCompletedHandler[TResult, TProgress]
	GetResults() TResult
}

type IAsyncOperationWithProgressVtbl struct {
	win32.IInspectableVtbl
	Put_Progress  uintptr
	Get_Progress  uintptr
	Put_Completed uintptr
	Get_Completed uintptr
	GetResults    uintptr
}

type IAsyncOperationWithProgress[TResult any, TProgress any] struct {
	win32.IInspectable
}

func (this *IAsyncOperationWithProgress[TResult, TProgress]) Vtbl() *IAsyncOperationWithProgressVtbl {
	return (*IAsyncOperationWithProgressVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAsyncOperationWithProgress[TResult, TProgress]) Put_Progress(handler AsyncOperationProgressHandler[TResult, TProgress]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Progress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))))
	_ = _hr
}

func (this *IAsyncOperationWithProgress[TResult, TProgress]) Get_Progress() AsyncOperationProgressHandler[TResult, TProgress] {
	var _result AsyncOperationProgressHandler[TResult, TProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Progress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAsyncOperationWithProgress[TResult, TProgress]) Put_Completed(handler AsyncOperationWithProgressCompletedHandler[TResult, TProgress]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Completed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))))
	_ = _hr
}

func (this *IAsyncOperationWithProgress[TResult, TProgress]) Get_Completed() AsyncOperationWithProgressCompletedHandler[TResult, TProgress] {
	var _result AsyncOperationWithProgressCompletedHandler[TResult, TProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Completed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAsyncOperationWithProgress[TResult, TProgress]) GetResults() TResult {
	var _result TResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetResults, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return PostProcessGenericResult(_result)
}

// 9FC2B0BB-E446-44E2-AA61-9CAB8F636AF2
var IID_IAsyncOperation = syscall.GUID{0x9FC2B0BB, 0xE446, 0x44E2,
	[8]byte{0xAA, 0x61, 0x9C, 0xAB, 0x8F, 0x63, 0x6A, 0xF2}}

type IAsyncOperationInterface[TResult any] interface {
	win32.IInspectableInterface
	Put_Completed(handler AsyncOperationCompletedHandler[TResult])
	Get_Completed() AsyncOperationCompletedHandler[TResult]
	GetResults() TResult
}

type IAsyncOperationVtbl struct {
	win32.IInspectableVtbl
	Put_Completed uintptr
	Get_Completed uintptr
	GetResults    uintptr
}

type IAsyncOperation[TResult any] struct {
	win32.IInspectable
}

func (this *IAsyncOperation[TResult]) Vtbl() *IAsyncOperationVtbl {
	return (*IAsyncOperationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAsyncOperation[TResult]) Put_Completed(handler AsyncOperationCompletedHandler[TResult]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Completed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))))
	_ = _hr
}

func (this *IAsyncOperation[TResult]) Get_Completed() AsyncOperationCompletedHandler[TResult] {
	var _result AsyncOperationCompletedHandler[TResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Completed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAsyncOperation[TResult]) GetResults() TResult {
	var _result TResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetResults, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return PostProcessGenericResult(_result)
}

// 30D5A829-7FA4-4026-83BB-D75BAE4EA99E
var IID_IClosable = syscall.GUID{0x30D5A829, 0x7FA4, 0x4026,
	[8]byte{0x83, 0xBB, 0xD7, 0x5B, 0xAE, 0x4E, 0xA9, 0x9E}}

type IClosableInterface interface {
	win32.IInspectableInterface
	Close()
}

type IClosableVtbl struct {
	win32.IInspectableVtbl
	Close uintptr
}

type IClosable struct {
	win32.IInspectable
}

func (this *IClosable) Vtbl() *IClosableVtbl {
	return (*IClosableVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClosable) Close() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Close, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// D6269732-3B7F-46A7-B40B-4FDCA2A2C693
var IID_IDeferral = syscall.GUID{0xD6269732, 0x3B7F, 0x46A7,
	[8]byte{0xB4, 0x0B, 0x4F, 0xDC, 0xA2, 0xA2, 0xC6, 0x93}}

type IDeferralInterface interface {
	win32.IInspectableInterface
	Complete()
}

type IDeferralVtbl struct {
	win32.IInspectableVtbl
	Complete uintptr
}

type IDeferral struct {
	win32.IInspectable
}

func (this *IDeferral) Vtbl() *IDeferralVtbl {
	return (*IDeferralVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeferral) Complete() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Complete, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 65A1ECC5-3FB5-4832-8CA9-F061B281D13A
var IID_IDeferralFactory = syscall.GUID{0x65A1ECC5, 0x3FB5, 0x4832,
	[8]byte{0x8C, 0xA9, 0xF0, 0x61, 0xB2, 0x81, 0xD1, 0x3A}}

type IDeferralFactoryInterface interface {
	win32.IInspectableInterface
	Create(handler DeferralCompletedHandler) *IDeferral
}

type IDeferralFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IDeferralFactory struct {
	win32.IInspectable
}

func (this *IDeferralFactory) Vtbl() *IDeferralFactoryVtbl {
	return (*IDeferralFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeferralFactory) Create(handler DeferralCompletedHandler) *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewNoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4EDB8EE2-96DD-49A7-94F7-4607DDAB8E3C
var IID_IGetActivationFactory = syscall.GUID{0x4EDB8EE2, 0x96DD, 0x49A7,
	[8]byte{0x94, 0xF7, 0x46, 0x07, 0xDD, 0xAB, 0x8E, 0x3C}}

type IGetActivationFactoryInterface interface {
	win32.IInspectableInterface
	GetActivationFactory(activatableClassId string) interface{}
}

type IGetActivationFactoryVtbl struct {
	win32.IInspectableVtbl
	GetActivationFactory uintptr
}

type IGetActivationFactory struct {
	win32.IInspectable
}

func (this *IGetActivationFactory) Vtbl() *IGetActivationFactoryVtbl {
	return (*IGetActivationFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGetActivationFactory) GetActivationFactory(activatableClassId string) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetActivationFactory, uintptr(unsafe.Pointer(this)), NewHStr(activatableClassId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 59C7966B-AE52-5283-AD7F-A1B9E9678ADD
var IID_IGuidHelperStatics = syscall.GUID{0x59C7966B, 0xAE52, 0x5283,
	[8]byte{0xAD, 0x7F, 0xA1, 0xB9, 0xE9, 0x67, 0x8A, 0xDD}}

type IGuidHelperStaticsInterface interface {
	win32.IInspectableInterface
	CreateNewGuid() syscall.GUID
	Get_Empty() syscall.GUID
	Equals(target syscall.GUID, value syscall.GUID) bool
}

type IGuidHelperStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateNewGuid uintptr
	Get_Empty     uintptr
	Equals        uintptr
}

type IGuidHelperStatics struct {
	win32.IInspectable
}

func (this *IGuidHelperStatics) Vtbl() *IGuidHelperStaticsVtbl {
	return (*IGuidHelperStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGuidHelperStatics) CreateNewGuid() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateNewGuid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGuidHelperStatics) Get_Empty() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Empty, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGuidHelperStatics) Equals(target syscall.GUID, value syscall.GUID) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Equals, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&target)), uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// FBC4DD2A-245B-11E4-AF98-689423260CF8
var IID_IMemoryBuffer = syscall.GUID{0xFBC4DD2A, 0x245B, 0x11E4,
	[8]byte{0xAF, 0x98, 0x68, 0x94, 0x23, 0x26, 0x0C, 0xF8}}

type IMemoryBufferInterface interface {
	win32.IInspectableInterface
	CreateReference() *IMemoryBufferReference
}

type IMemoryBufferVtbl struct {
	win32.IInspectableVtbl
	CreateReference uintptr
}

type IMemoryBuffer struct {
	win32.IInspectable
}

func (this *IMemoryBuffer) Vtbl() *IMemoryBufferVtbl {
	return (*IMemoryBufferVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMemoryBuffer) CreateReference() *IMemoryBufferReference {
	var _result *IMemoryBufferReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateReference, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FBC4DD2B-245B-11E4-AF98-689423260CF8
var IID_IMemoryBufferFactory = syscall.GUID{0xFBC4DD2B, 0x245B, 0x11E4,
	[8]byte{0xAF, 0x98, 0x68, 0x94, 0x23, 0x26, 0x0C, 0xF8}}

type IMemoryBufferFactoryInterface interface {
	win32.IInspectableInterface
	Create(capacity uint32) *IMemoryBuffer
}

type IMemoryBufferFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IMemoryBufferFactory struct {
	win32.IInspectable
}

func (this *IMemoryBufferFactory) Vtbl() *IMemoryBufferFactoryVtbl {
	return (*IMemoryBufferFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMemoryBufferFactory) Create(capacity uint32) *IMemoryBuffer {
	var _result *IMemoryBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(capacity), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FBC4DD29-245B-11E4-AF98-689423260CF8
var IID_IMemoryBufferReference = syscall.GUID{0xFBC4DD29, 0x245B, 0x11E4,
	[8]byte{0xAF, 0x98, 0x68, 0x94, 0x23, 0x26, 0x0C, 0xF8}}

type IMemoryBufferReferenceInterface interface {
	win32.IInspectableInterface
	Get_Capacity() uint32
	Add_Closed(handler TypedEventHandler[*IMemoryBufferReference, interface{}]) EventRegistrationToken
	Remove_Closed(cookie EventRegistrationToken)
}

type IMemoryBufferReferenceVtbl struct {
	win32.IInspectableVtbl
	Get_Capacity  uintptr
	Add_Closed    uintptr
	Remove_Closed uintptr
}

type IMemoryBufferReference struct {
	win32.IInspectable
}

func (this *IMemoryBufferReference) Vtbl() *IMemoryBufferReferenceVtbl {
	return (*IMemoryBufferReferenceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMemoryBufferReference) Get_Capacity() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Capacity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMemoryBufferReference) Add_Closed(handler TypedEventHandler[*IMemoryBufferReference, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Closed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMemoryBufferReference) Remove_Closed(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Closed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

// 4BD682DD-7554-40E9-9A9B-82654EDE7E62
var IID_IPropertyValue = syscall.GUID{0x4BD682DD, 0x7554, 0x40E9,
	[8]byte{0x9A, 0x9B, 0x82, 0x65, 0x4E, 0xDE, 0x7E, 0x62}}

type IPropertyValueInterface interface {
	win32.IInspectableInterface
	Get_Type() PropertyType
	Get_IsNumericScalar() bool
	GetUInt8() byte
	GetInt16() int16
	GetUInt16() uint16
	GetInt32() int32
	GetUInt32() uint32
	GetInt64() int64
	GetUInt64() uint64
	GetSingle() float32
	GetDouble() float64
	GetChar16() uint16
	GetBoolean() bool
	GetString() string
	GetGuid() syscall.GUID
	GetDateTime() DateTime
	GetTimeSpan() TimeSpan
	GetPoint() Point
	GetSize() Size
	GetRect() Rect
	GetUInt8Array(valueLength uint32, value *byte)
	GetInt16Array(valueLength uint32, value *int16)
	GetUInt16Array(valueLength uint32, value *uint16)
	GetInt32Array(valueLength uint32, value *int32)
	GetUInt32Array(valueLength uint32, value *uint32)
	GetInt64Array(valueLength uint32, value *int64)
	GetUInt64Array(valueLength uint32, value *uint64)
	GetSingleArray(valueLength uint32, value *float32)
	GetDoubleArray(valueLength uint32, value *float64)
	GetChar16Array(valueLength uint32, value *uint16)
	GetBooleanArray(valueLength uint32, value *bool)
	GetStringArray(valueLength uint32, value *string)
	GetInspectableArray(valueLength uint32, value *interface{})
	GetGuidArray(valueLength uint32, value *syscall.GUID)
	GetDateTimeArray(valueLength uint32, value *DateTime)
	GetTimeSpanArray(valueLength uint32, value *TimeSpan)
	GetPointArray(valueLength uint32, value *Point)
	GetSizeArray(valueLength uint32, value *Size)
	GetRectArray(valueLength uint32, value *Rect)
}

type IPropertyValueVtbl struct {
	win32.IInspectableVtbl
	Get_Type            uintptr
	Get_IsNumericScalar uintptr
	GetUInt8            uintptr
	GetInt16            uintptr
	GetUInt16           uintptr
	GetInt32            uintptr
	GetUInt32           uintptr
	GetInt64            uintptr
	GetUInt64           uintptr
	GetSingle           uintptr
	GetDouble           uintptr
	GetChar16           uintptr
	GetBoolean          uintptr
	GetString           uintptr
	GetGuid             uintptr
	GetDateTime         uintptr
	GetTimeSpan         uintptr
	GetPoint            uintptr
	GetSize             uintptr
	GetRect             uintptr
	GetUInt8Array       uintptr
	GetInt16Array       uintptr
	GetUInt16Array      uintptr
	GetInt32Array       uintptr
	GetUInt32Array      uintptr
	GetInt64Array       uintptr
	GetUInt64Array      uintptr
	GetSingleArray      uintptr
	GetDoubleArray      uintptr
	GetChar16Array      uintptr
	GetBooleanArray     uintptr
	GetStringArray      uintptr
	GetInspectableArray uintptr
	GetGuidArray        uintptr
	GetDateTimeArray    uintptr
	GetTimeSpanArray    uintptr
	GetPointArray       uintptr
	GetSizeArray        uintptr
	GetRectArray        uintptr
}

type IPropertyValue struct {
	win32.IInspectable
}

func (this *IPropertyValue) Vtbl() *IPropertyValueVtbl {
	return (*IPropertyValueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertyValue) Get_Type() PropertyType {
	var _result PropertyType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Type, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValue) Get_IsNumericScalar() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsNumericScalar, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValue) GetUInt8() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetUInt8, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValue) GetInt16() int16 {
	var _result int16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetInt16, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValue) GetUInt16() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetUInt16, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValue) GetInt32() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetInt32, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValue) GetUInt32() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetUInt32, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValue) GetInt64() int64 {
	var _result int64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetInt64, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValue) GetUInt64() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetUInt64, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValue) GetSingle() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSingle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValue) GetDouble() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDouble, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValue) GetChar16() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetChar16, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValue) GetBoolean() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetBoolean, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValue) GetString() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPropertyValue) GetGuid() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetGuid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValue) GetDateTime() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDateTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValue) GetTimeSpan() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTimeSpan, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValue) GetPoint() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValue) GetSize() Size {
	var _result Size
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValue) GetRect() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetRect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValue) GetUInt8Array(valueLength uint32, value *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetUInt8Array, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPropertyValue) GetInt16Array(valueLength uint32, value *int16) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetInt16Array, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPropertyValue) GetUInt16Array(valueLength uint32, value *uint16) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetUInt16Array, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPropertyValue) GetInt32Array(valueLength uint32, value *int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetInt32Array, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPropertyValue) GetUInt32Array(valueLength uint32, value *uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetUInt32Array, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPropertyValue) GetInt64Array(valueLength uint32, value *int64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetInt64Array, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPropertyValue) GetUInt64Array(valueLength uint32, value *uint64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetUInt64Array, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPropertyValue) GetSingleArray(valueLength uint32, value *float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSingleArray, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPropertyValue) GetDoubleArray(valueLength uint32, value *float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDoubleArray, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPropertyValue) GetChar16Array(valueLength uint32, value *uint16) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetChar16Array, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPropertyValue) GetBooleanArray(valueLength uint32, value *bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetBooleanArray, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPropertyValue) GetStringArray(valueLength uint32, value *string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStringArray, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPropertyValue) GetInspectableArray(valueLength uint32, value *interface{}) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetInspectableArray, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPropertyValue) GetGuidArray(valueLength uint32, value *syscall.GUID) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetGuidArray, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPropertyValue) GetDateTimeArray(valueLength uint32, value *DateTime) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDateTimeArray, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPropertyValue) GetTimeSpanArray(valueLength uint32, value *TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTimeSpanArray, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPropertyValue) GetPointArray(valueLength uint32, value *Point) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPointArray, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPropertyValue) GetSizeArray(valueLength uint32, value *Size) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSizeArray, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPropertyValue) GetRectArray(valueLength uint32, value *Rect) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetRectArray, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 629BDBC8-D932-4FF4-96B9-8D96C5C1E858
var IID_IPropertyValueStatics = syscall.GUID{0x629BDBC8, 0xD932, 0x4FF4,
	[8]byte{0x96, 0xB9, 0x8D, 0x96, 0xC5, 0xC1, 0xE8, 0x58}}

type IPropertyValueStaticsInterface interface {
	win32.IInspectableInterface
	CreateEmpty() interface{}
	CreateUInt8(value byte) interface{}
	CreateInt16(value int16) interface{}
	CreateUInt16(value uint16) interface{}
	CreateInt32(value int32) interface{}
	CreateUInt32(value uint32) interface{}
	CreateInt64(value int64) interface{}
	CreateUInt64(value uint64) interface{}
	CreateSingle(value float32) interface{}
	CreateDouble(value float64) interface{}
	CreateChar16(value uint16) interface{}
	CreateBoolean(value bool) interface{}
	CreateString(value string) interface{}
	CreateInspectable(value interface{}) interface{}
	CreateGuid(value syscall.GUID) interface{}
	CreateDateTime(value DateTime) interface{}
	CreateTimeSpan(value TimeSpan) interface{}
	CreatePoint(value Point) interface{}
	CreateSize(value Size) interface{}
	CreateRect(value Rect) interface{}
	CreateUInt8Array(valueLength uint32, value *byte) interface{}
	CreateInt16Array(valueLength uint32, value *int16) interface{}
	CreateUInt16Array(valueLength uint32, value *uint16) interface{}
	CreateInt32Array(valueLength uint32, value *int32) interface{}
	CreateUInt32Array(valueLength uint32, value *uint32) interface{}
	CreateInt64Array(valueLength uint32, value *int64) interface{}
	CreateUInt64Array(valueLength uint32, value *uint64) interface{}
	CreateSingleArray(valueLength uint32, value *float32) interface{}
	CreateDoubleArray(valueLength uint32, value *float64) interface{}
	CreateChar16Array(valueLength uint32, value *uint16) interface{}
	CreateBooleanArray(valueLength uint32, value *bool) interface{}
	CreateStringArray(valueLength uint32, value *string) interface{}
	CreateInspectableArray(valueLength uint32, value *interface{}) interface{}
	CreateGuidArray(valueLength uint32, value *syscall.GUID) interface{}
	CreateDateTimeArray(valueLength uint32, value *DateTime) interface{}
	CreateTimeSpanArray(valueLength uint32, value *TimeSpan) interface{}
	CreatePointArray(valueLength uint32, value *Point) interface{}
	CreateSizeArray(valueLength uint32, value *Size) interface{}
	CreateRectArray(valueLength uint32, value *Rect) interface{}
}

type IPropertyValueStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateEmpty            uintptr
	CreateUInt8            uintptr
	CreateInt16            uintptr
	CreateUInt16           uintptr
	CreateInt32            uintptr
	CreateUInt32           uintptr
	CreateInt64            uintptr
	CreateUInt64           uintptr
	CreateSingle           uintptr
	CreateDouble           uintptr
	CreateChar16           uintptr
	CreateBoolean          uintptr
	CreateString           uintptr
	CreateInspectable      uintptr
	CreateGuid             uintptr
	CreateDateTime         uintptr
	CreateTimeSpan         uintptr
	CreatePoint            uintptr
	CreateSize             uintptr
	CreateRect             uintptr
	CreateUInt8Array       uintptr
	CreateInt16Array       uintptr
	CreateUInt16Array      uintptr
	CreateInt32Array       uintptr
	CreateUInt32Array      uintptr
	CreateInt64Array       uintptr
	CreateUInt64Array      uintptr
	CreateSingleArray      uintptr
	CreateDoubleArray      uintptr
	CreateChar16Array      uintptr
	CreateBooleanArray     uintptr
	CreateStringArray      uintptr
	CreateInspectableArray uintptr
	CreateGuidArray        uintptr
	CreateDateTimeArray    uintptr
	CreateTimeSpanArray    uintptr
	CreatePointArray       uintptr
	CreateSizeArray        uintptr
	CreateRectArray        uintptr
}

type IPropertyValueStatics struct {
	win32.IInspectable
}

func (this *IPropertyValueStatics) Vtbl() *IPropertyValueStaticsVtbl {
	return (*IPropertyValueStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertyValueStatics) CreateEmpty() interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateEmpty, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateUInt8(value byte) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateUInt8, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateInt16(value int16) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateInt16, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateUInt16(value uint16) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateUInt16, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateInt32(value int32) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateInt32, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateUInt32(value uint32) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateUInt32, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateInt64(value int64) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateInt64, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateUInt64(value uint64) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateUInt64, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateSingle(value float32) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSingle, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateDouble(value float64) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDouble, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateChar16(value uint16) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateChar16, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateBoolean(value bool) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateBoolean, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateString(value string) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateString, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateInspectable(value interface{}) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateInspectable, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateGuid(value syscall.GUID) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateGuid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateDateTime(value DateTime) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDateTime, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateTimeSpan(value TimeSpan) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateTimeSpan, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreatePoint(value Point) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreatePoint, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateSize(value Size) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSize, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateRect(value Rect) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateRect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateUInt8Array(valueLength uint32, value *byte) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateUInt8Array, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateInt16Array(valueLength uint32, value *int16) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateInt16Array, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateUInt16Array(valueLength uint32, value *uint16) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateUInt16Array, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateInt32Array(valueLength uint32, value *int32) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateInt32Array, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateUInt32Array(valueLength uint32, value *uint32) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateUInt32Array, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateInt64Array(valueLength uint32, value *int64) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateInt64Array, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateUInt64Array(valueLength uint32, value *uint64) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateUInt64Array, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateSingleArray(valueLength uint32, value *float32) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSingleArray, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateDoubleArray(valueLength uint32, value *float64) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDoubleArray, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateChar16Array(valueLength uint32, value *uint16) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateChar16Array, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateBooleanArray(valueLength uint32, value *bool) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateBooleanArray, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateStringArray(valueLength uint32, value *string) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateStringArray, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateInspectableArray(valueLength uint32, value *interface{}) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateInspectableArray, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateGuidArray(valueLength uint32, value *syscall.GUID) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateGuidArray, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateDateTimeArray(valueLength uint32, value *DateTime) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDateTimeArray, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateTimeSpanArray(valueLength uint32, value *TimeSpan) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateTimeSpanArray, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreatePointArray(valueLength uint32, value *Point) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreatePointArray, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateSizeArray(valueLength uint32, value *Size) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSizeArray, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyValueStatics) CreateRectArray(valueLength uint32, value *Rect) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateRectArray, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 61C17707-2D65-11E0-9AE8-D48564015472
var IID_IReferenceArray = syscall.GUID{0x61C17707, 0x2D65, 0x11E0,
	[8]byte{0x9A, 0xE8, 0xD4, 0x85, 0x64, 0x01, 0x54, 0x72}}

type IReferenceArrayInterface[T any] interface {
	win32.IInspectableInterface
	Get_Value() []T
}

type IReferenceArrayVtbl struct {
	win32.IInspectableVtbl
	Get_Value uintptr
}

type IReferenceArray[T any] struct {
	win32.IInspectable
}

func (this *IReferenceArray[T]) Vtbl() *IReferenceArrayVtbl {
	return (*IReferenceArrayVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IReferenceArray[T]) Get_Value() []T {
	var _result []T
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 61C17706-2D65-11E0-9AE8-D48564015472
var IID_IReference = syscall.GUID{0x61C17706, 0x2D65, 0x11E0,
	[8]byte{0x9A, 0xE8, 0xD4, 0x85, 0x64, 0x01, 0x54, 0x72}}

type IReferenceInterface[T any] interface {
	win32.IInspectableInterface
	Get_Value() T
}

type IReferenceVtbl struct {
	win32.IInspectableVtbl
	Get_Value uintptr
}

type IReference[T any] struct {
	win32.IInspectable
}

func (this *IReference[T]) Vtbl() *IReferenceVtbl {
	return (*IReferenceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IReference[T]) Get_Value() T {
	var _result T
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return PostProcessGenericResult(_result)
}

// 96369F54-8EB6-48F0-ABCE-C1B211E627C3
var IID_IStringable = syscall.GUID{0x96369F54, 0x8EB6, 0x48F0,
	[8]byte{0xAB, 0xCE, 0xC1, 0xB2, 0x11, 0xE6, 0x27, 0xC3}}

type IStringableInterface interface {
	win32.IInspectableInterface
	ToString() string
}

type IStringableVtbl struct {
	win32.IInspectableVtbl
	ToString uintptr
}

type IStringable struct {
	win32.IInspectable
}

func (this *IStringable) Vtbl() *IStringableVtbl {
	return (*IStringableVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStringable) ToString() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ToString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// C1D432BA-C824-4452-A7FD-512BC3BBE9A1
var IID_IUriEscapeStatics = syscall.GUID{0xC1D432BA, 0xC824, 0x4452,
	[8]byte{0xA7, 0xFD, 0x51, 0x2B, 0xC3, 0xBB, 0xE9, 0xA1}}

type IUriEscapeStaticsInterface interface {
	win32.IInspectableInterface
	UnescapeComponent(toUnescape string) string
	EscapeComponent(toEscape string) string
}

type IUriEscapeStaticsVtbl struct {
	win32.IInspectableVtbl
	UnescapeComponent uintptr
	EscapeComponent   uintptr
}

type IUriEscapeStatics struct {
	win32.IInspectable
}

func (this *IUriEscapeStatics) Vtbl() *IUriEscapeStaticsVtbl {
	return (*IUriEscapeStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUriEscapeStatics) UnescapeComponent(toUnescape string) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UnescapeComponent, uintptr(unsafe.Pointer(this)), NewHStr(toUnescape).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUriEscapeStatics) EscapeComponent(toEscape string) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EscapeComponent, uintptr(unsafe.Pointer(this)), NewHStr(toEscape).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 9E365E57-48B2-4160-956F-C7385120BBFC
var IID_IUriRuntimeClass = syscall.GUID{0x9E365E57, 0x48B2, 0x4160,
	[8]byte{0x95, 0x6F, 0xC7, 0x38, 0x51, 0x20, 0xBB, 0xFC}}

type IUriRuntimeClassInterface interface {
	win32.IInspectableInterface
	Get_AbsoluteUri() string
	Get_DisplayUri() string
	Get_Domain() string
	Get_Extension() string
	Get_Fragment() string
	Get_Host() string
	Get_Password() string
	Get_Path() string
	Get_Query() string
	Get_QueryParsed() *IWwwFormUrlDecoderRuntimeClass
	Get_RawUri() string
	Get_SchemeName() string
	Get_UserName() string
	Get_Port() int32
	Get_Suspicious() bool
	Equals(pUri *IUriRuntimeClass) bool
	CombineUri(relativeUri string) *IUriRuntimeClass
}

type IUriRuntimeClassVtbl struct {
	win32.IInspectableVtbl
	Get_AbsoluteUri uintptr
	Get_DisplayUri  uintptr
	Get_Domain      uintptr
	Get_Extension   uintptr
	Get_Fragment    uintptr
	Get_Host        uintptr
	Get_Password    uintptr
	Get_Path        uintptr
	Get_Query       uintptr
	Get_QueryParsed uintptr
	Get_RawUri      uintptr
	Get_SchemeName  uintptr
	Get_UserName    uintptr
	Get_Port        uintptr
	Get_Suspicious  uintptr
	Equals          uintptr
	CombineUri      uintptr
}

type IUriRuntimeClass struct {
	win32.IInspectable
}

func (this *IUriRuntimeClass) Vtbl() *IUriRuntimeClassVtbl {
	return (*IUriRuntimeClassVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUriRuntimeClass) Get_AbsoluteUri() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AbsoluteUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUriRuntimeClass) Get_DisplayUri() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUriRuntimeClass) Get_Domain() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Domain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUriRuntimeClass) Get_Extension() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Extension, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUriRuntimeClass) Get_Fragment() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Fragment, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUriRuntimeClass) Get_Host() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Host, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUriRuntimeClass) Get_Password() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Password, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUriRuntimeClass) Get_Path() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Path, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUriRuntimeClass) Get_Query() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Query, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUriRuntimeClass) Get_QueryParsed() *IWwwFormUrlDecoderRuntimeClass {
	var _result *IWwwFormUrlDecoderRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_QueryParsed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUriRuntimeClass) Get_RawUri() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RawUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUriRuntimeClass) Get_SchemeName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SchemeName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUriRuntimeClass) Get_UserName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUriRuntimeClass) Get_Port() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Port, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUriRuntimeClass) Get_Suspicious() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Suspicious, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUriRuntimeClass) Equals(pUri *IUriRuntimeClass) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Equals, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pUri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUriRuntimeClass) CombineUri(relativeUri string) *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CombineUri, uintptr(unsafe.Pointer(this)), NewHStr(relativeUri).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 44A9796F-723E-4FDF-A218-033E75B0C084
var IID_IUriRuntimeClassFactory = syscall.GUID{0x44A9796F, 0x723E, 0x4FDF,
	[8]byte{0xA2, 0x18, 0x03, 0x3E, 0x75, 0xB0, 0xC0, 0x84}}

type IUriRuntimeClassFactoryInterface interface {
	win32.IInspectableInterface
	CreateUri(uri string) *IUriRuntimeClass
	CreateWithRelativeUri(baseUri string, relativeUri string) *IUriRuntimeClass
}

type IUriRuntimeClassFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateUri             uintptr
	CreateWithRelativeUri uintptr
}

type IUriRuntimeClassFactory struct {
	win32.IInspectable
}

func (this *IUriRuntimeClassFactory) Vtbl() *IUriRuntimeClassFactoryVtbl {
	return (*IUriRuntimeClassFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUriRuntimeClassFactory) CreateUri(uri string) *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateUri, uintptr(unsafe.Pointer(this)), NewHStr(uri).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUriRuntimeClassFactory) CreateWithRelativeUri(baseUri string, relativeUri string) *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithRelativeUri, uintptr(unsafe.Pointer(this)), NewHStr(baseUri).Ptr, NewHStr(relativeUri).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 758D9661-221C-480F-A339-50656673F46F
var IID_IUriRuntimeClassWithAbsoluteCanonicalUri = syscall.GUID{0x758D9661, 0x221C, 0x480F,
	[8]byte{0xA3, 0x39, 0x50, 0x65, 0x66, 0x73, 0xF4, 0x6F}}

type IUriRuntimeClassWithAbsoluteCanonicalUriInterface interface {
	win32.IInspectableInterface
	Get_AbsoluteCanonicalUri() string
	Get_DisplayIri() string
}

type IUriRuntimeClassWithAbsoluteCanonicalUriVtbl struct {
	win32.IInspectableVtbl
	Get_AbsoluteCanonicalUri uintptr
	Get_DisplayIri           uintptr
}

type IUriRuntimeClassWithAbsoluteCanonicalUri struct {
	win32.IInspectable
}

func (this *IUriRuntimeClassWithAbsoluteCanonicalUri) Vtbl() *IUriRuntimeClassWithAbsoluteCanonicalUriVtbl {
	return (*IUriRuntimeClassWithAbsoluteCanonicalUriVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUriRuntimeClassWithAbsoluteCanonicalUri) Get_AbsoluteCanonicalUri() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AbsoluteCanonicalUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUriRuntimeClassWithAbsoluteCanonicalUri) Get_DisplayIri() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayIri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 125E7431-F678-4E8E-B670-20A9B06C512D
var IID_IWwwFormUrlDecoderEntry = syscall.GUID{0x125E7431, 0xF678, 0x4E8E,
	[8]byte{0xB6, 0x70, 0x20, 0xA9, 0xB0, 0x6C, 0x51, 0x2D}}

type IWwwFormUrlDecoderEntryInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Get_Value() string
}

type IWwwFormUrlDecoderEntryVtbl struct {
	win32.IInspectableVtbl
	Get_Name  uintptr
	Get_Value uintptr
}

type IWwwFormUrlDecoderEntry struct {
	win32.IInspectable
}

func (this *IWwwFormUrlDecoderEntry) Vtbl() *IWwwFormUrlDecoderEntryVtbl {
	return (*IWwwFormUrlDecoderEntryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWwwFormUrlDecoderEntry) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWwwFormUrlDecoderEntry) Get_Value() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// D45A0451-F225-4542-9296-0E1DF5D254DF
var IID_IWwwFormUrlDecoderRuntimeClass = syscall.GUID{0xD45A0451, 0xF225, 0x4542,
	[8]byte{0x92, 0x96, 0x0E, 0x1D, 0xF5, 0xD2, 0x54, 0xDF}}

type IWwwFormUrlDecoderRuntimeClassInterface interface {
	win32.IInspectableInterface
	GetFirstValueByName(name string) string
}

type IWwwFormUrlDecoderRuntimeClassVtbl struct {
	win32.IInspectableVtbl
	GetFirstValueByName uintptr
}

type IWwwFormUrlDecoderRuntimeClass struct {
	win32.IInspectable
}

func (this *IWwwFormUrlDecoderRuntimeClass) Vtbl() *IWwwFormUrlDecoderRuntimeClassVtbl {
	return (*IWwwFormUrlDecoderRuntimeClassVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWwwFormUrlDecoderRuntimeClass) GetFirstValueByName(name string) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFirstValueByName, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 5B8C6B3D-24AE-41B5-A1BF-F0C3D544845B
var IID_IWwwFormUrlDecoderRuntimeClassFactory = syscall.GUID{0x5B8C6B3D, 0x24AE, 0x41B5,
	[8]byte{0xA1, 0xBF, 0xF0, 0xC3, 0xD5, 0x44, 0x84, 0x5B}}

type IWwwFormUrlDecoderRuntimeClassFactoryInterface interface {
	win32.IInspectableInterface
	CreateWwwFormUrlDecoder(query string) *IWwwFormUrlDecoderRuntimeClass
}

type IWwwFormUrlDecoderRuntimeClassFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateWwwFormUrlDecoder uintptr
}

type IWwwFormUrlDecoderRuntimeClassFactory struct {
	win32.IInspectable
}

func (this *IWwwFormUrlDecoderRuntimeClassFactory) Vtbl() *IWwwFormUrlDecoderRuntimeClassFactoryVtbl {
	return (*IWwwFormUrlDecoderRuntimeClassFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWwwFormUrlDecoderRuntimeClassFactory) CreateWwwFormUrlDecoder(query string) *IWwwFormUrlDecoderRuntimeClass {
	var _result *IWwwFormUrlDecoderRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWwwFormUrlDecoder, uintptr(unsafe.Pointer(this)), NewHStr(query).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type Deferral struct {
	RtClass
	*IDeferral
}

func NewDeferral_Create(handler DeferralCompletedHandler) *Deferral {
	hs := NewHStr("Windows.Foundation.Deferral")
	var pFac *IDeferralFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDeferralFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IDeferral
	p = pFac.Create(handler)
	result := &Deferral{
		RtClass:   RtClass{PInspect: &p.IInspectable},
		IDeferral: p,
	}
	com.AddToScope(result)
	return result
}

type MemoryBuffer struct {
	RtClass
	*IMemoryBuffer
}

func NewMemoryBuffer_Create(capacity uint32) *MemoryBuffer {
	hs := NewHStr("Windows.Foundation.MemoryBuffer")
	var pFac *IMemoryBufferFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMemoryBufferFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IMemoryBuffer
	p = pFac.Create(capacity)
	result := &MemoryBuffer{
		RtClass:       RtClass{PInspect: &p.IInspectable},
		IMemoryBuffer: p,
	}
	com.AddToScope(result)
	return result
}

type PropertyValue struct {
	RtClass
}

func NewIPropertyValueStatics() *IPropertyValueStatics {
	var p *IPropertyValueStatics
	hs := NewHStr("Windows.Foundation.PropertyValue")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPropertyValueStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type Uri struct {
	RtClass
	*IUriRuntimeClass
}

func NewUri_CreateUri(uri string) *Uri {
	hs := NewHStr("Windows.Foundation.Uri")
	var pFac *IUriRuntimeClassFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IUriRuntimeClassFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IUriRuntimeClass
	p = pFac.CreateUri(uri)
	result := &Uri{
		RtClass:          RtClass{PInspect: &p.IInspectable},
		IUriRuntimeClass: p,
	}
	com.AddToScope(result)
	return result
}

func NewUri_CreateWithRelativeUri(baseUri string, relativeUri string) *Uri {
	hs := NewHStr("Windows.Foundation.Uri")
	var pFac *IUriRuntimeClassFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IUriRuntimeClassFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IUriRuntimeClass
	p = pFac.CreateWithRelativeUri(baseUri, relativeUri)
	result := &Uri{
		RtClass:          RtClass{PInspect: &p.IInspectable},
		IUriRuntimeClass: p,
	}
	com.AddToScope(result)
	return result
}

func NewIUriEscapeStatics() *IUriEscapeStatics {
	var p *IUriEscapeStatics
	hs := NewHStr("Windows.Foundation.Uri")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IUriEscapeStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type WwwFormUrlDecoder struct {
	RtClass
	*IWwwFormUrlDecoderRuntimeClass
}

func NewWwwFormUrlDecoder_CreateWwwFormUrlDecoder(query string) *WwwFormUrlDecoder {
	hs := NewHStr("Windows.Foundation.WwwFormUrlDecoder")
	var pFac *IWwwFormUrlDecoderRuntimeClassFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWwwFormUrlDecoderRuntimeClassFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IWwwFormUrlDecoderRuntimeClass
	p = pFac.CreateWwwFormUrlDecoder(query)
	result := &WwwFormUrlDecoder{
		RtClass:                        RtClass{PInspect: &p.IInspectable},
		IWwwFormUrlDecoderRuntimeClass: p,
	}
	com.AddToScope(result)
	return result
}

type WwwFormUrlDecoderEntry struct {
	RtClass
	*IWwwFormUrlDecoderEntry
}
