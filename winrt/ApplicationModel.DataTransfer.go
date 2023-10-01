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
type ClipboardHistoryItemsResultStatus int32

const (
	ClipboardHistoryItemsResultStatus_Success                  ClipboardHistoryItemsResultStatus = 0
	ClipboardHistoryItemsResultStatus_AccessDenied             ClipboardHistoryItemsResultStatus = 1
	ClipboardHistoryItemsResultStatus_ClipboardHistoryDisabled ClipboardHistoryItemsResultStatus = 2
)

// enum
// flags
type DataPackageOperation uint32

const (
	DataPackageOperation_None DataPackageOperation = 0
	DataPackageOperation_Copy DataPackageOperation = 1
	DataPackageOperation_Move DataPackageOperation = 2
	DataPackageOperation_Link DataPackageOperation = 4
)

// enum
type SetHistoryItemAsContentStatus int32

const (
	SetHistoryItemAsContentStatus_Success      SetHistoryItemAsContentStatus = 0
	SetHistoryItemAsContentStatus_AccessDenied SetHistoryItemAsContentStatus = 1
	SetHistoryItemAsContentStatus_ItemDeleted  SetHistoryItemAsContentStatus = 2
)

// enum
type ShareUITheme int32

const (
	ShareUITheme_Default ShareUITheme = 0
	ShareUITheme_Light   ShareUITheme = 1
	ShareUITheme_Dark    ShareUITheme = 2
)

// func types

// E7ECD720-F2F4-4A2D-920E-170A2F482A27
type DataProviderHandler func(request *IDataProviderRequest) com.Error

// E7F9D9BA-E1BA-4E4D-BD65-D43845D3212F
type ShareProviderHandler func(operation *IShareProviderOperation) com.Error

// interfaces

// E888A98C-AD4B-5447-A056-AB3556276D2B
var IID_IClipboardContentOptions = syscall.GUID{0xE888A98C, 0xAD4B, 0x5447,
	[8]byte{0xA0, 0x56, 0xAB, 0x35, 0x56, 0x27, 0x6D, 0x2B}}

type IClipboardContentOptionsInterface interface {
	win32.IInspectableInterface
	Get_IsRoamable() bool
	Put_IsRoamable(value bool)
	Get_IsAllowedInHistory() bool
	Put_IsAllowedInHistory(value bool)
	Get_RoamingFormats() *IVector[string]
	Get_HistoryFormats() *IVector[string]
}

type IClipboardContentOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_IsRoamable         uintptr
	Put_IsRoamable         uintptr
	Get_IsAllowedInHistory uintptr
	Put_IsAllowedInHistory uintptr
	Get_RoamingFormats     uintptr
	Get_HistoryFormats     uintptr
}

type IClipboardContentOptions struct {
	win32.IInspectable
}

func (this *IClipboardContentOptions) Vtbl() *IClipboardContentOptionsVtbl {
	return (*IClipboardContentOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClipboardContentOptions) Get_IsRoamable() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsRoamable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClipboardContentOptions) Put_IsRoamable(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsRoamable, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IClipboardContentOptions) Get_IsAllowedInHistory() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAllowedInHistory, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClipboardContentOptions) Put_IsAllowedInHistory(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsAllowedInHistory, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IClipboardContentOptions) Get_RoamingFormats() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RoamingFormats, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClipboardContentOptions) Get_HistoryFormats() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HistoryFormats, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C0BE453F-8EA2-53CE-9ABA-8D2212573452
var IID_IClipboardHistoryChangedEventArgs = syscall.GUID{0xC0BE453F, 0x8EA2, 0x53CE,
	[8]byte{0x9A, 0xBA, 0x8D, 0x22, 0x12, 0x57, 0x34, 0x52}}

type IClipboardHistoryChangedEventArgsInterface interface {
	win32.IInspectableInterface
}

type IClipboardHistoryChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
}

type IClipboardHistoryChangedEventArgs struct {
	win32.IInspectable
}

func (this *IClipboardHistoryChangedEventArgs) Vtbl() *IClipboardHistoryChangedEventArgsVtbl {
	return (*IClipboardHistoryChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 0173BD8A-AFFF-5C50-AB92-3D19F481EC58
var IID_IClipboardHistoryItem = syscall.GUID{0x0173BD8A, 0xAFFF, 0x5C50,
	[8]byte{0xAB, 0x92, 0x3D, 0x19, 0xF4, 0x81, 0xEC, 0x58}}

type IClipboardHistoryItemInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Get_Timestamp() DateTime
	Get_Content() *IDataPackageView
}

type IClipboardHistoryItemVtbl struct {
	win32.IInspectableVtbl
	Get_Id        uintptr
	Get_Timestamp uintptr
	Get_Content   uintptr
}

type IClipboardHistoryItem struct {
	win32.IInspectable
}

func (this *IClipboardHistoryItem) Vtbl() *IClipboardHistoryItemVtbl {
	return (*IClipboardHistoryItemVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClipboardHistoryItem) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IClipboardHistoryItem) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClipboardHistoryItem) Get_Content() *IDataPackageView {
	var _result *IDataPackageView
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Content, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E6DFDEE6-0EE2-52E3-852B-F295DB65939A
var IID_IClipboardHistoryItemsResult = syscall.GUID{0xE6DFDEE6, 0x0EE2, 0x52E3,
	[8]byte{0x85, 0x2B, 0xF2, 0x95, 0xDB, 0x65, 0x93, 0x9A}}

type IClipboardHistoryItemsResultInterface interface {
	win32.IInspectableInterface
	Get_Status() ClipboardHistoryItemsResultStatus
	Get_Items() *IVectorView[*IClipboardHistoryItem]
}

type IClipboardHistoryItemsResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
	Get_Items  uintptr
}

type IClipboardHistoryItemsResult struct {
	win32.IInspectable
}

func (this *IClipboardHistoryItemsResult) Vtbl() *IClipboardHistoryItemsResultVtbl {
	return (*IClipboardHistoryItemsResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClipboardHistoryItemsResult) Get_Status() ClipboardHistoryItemsResultStatus {
	var _result ClipboardHistoryItemsResultStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClipboardHistoryItemsResult) Get_Items() *IVectorView[*IClipboardHistoryItem] {
	var _result *IVectorView[*IClipboardHistoryItem]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Items, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C627E291-34E2-4963-8EED-93CBB0EA3D70
var IID_IClipboardStatics = syscall.GUID{0xC627E291, 0x34E2, 0x4963,
	[8]byte{0x8E, 0xED, 0x93, 0xCB, 0xB0, 0xEA, 0x3D, 0x70}}

type IClipboardStaticsInterface interface {
	win32.IInspectableInterface
	GetContent() *IDataPackageView
	SetContent(content *IDataPackage)
	Flush()
	Clear()
	Add_ContentChanged(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_ContentChanged(token EventRegistrationToken)
}

type IClipboardStaticsVtbl struct {
	win32.IInspectableVtbl
	GetContent            uintptr
	SetContent            uintptr
	Flush                 uintptr
	Clear                 uintptr
	Add_ContentChanged    uintptr
	Remove_ContentChanged uintptr
}

type IClipboardStatics struct {
	win32.IInspectable
}

func (this *IClipboardStatics) Vtbl() *IClipboardStaticsVtbl {
	return (*IClipboardStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClipboardStatics) GetContent() *IDataPackageView {
	var _result *IDataPackageView
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetContent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClipboardStatics) SetContent(content *IDataPackage) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetContent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(content)))
	_ = _hr
}

func (this *IClipboardStatics) Flush() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Flush, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IClipboardStatics) Clear() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Clear, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IClipboardStatics) Add_ContentChanged(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ContentChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClipboardStatics) Remove_ContentChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ContentChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// D2AC1B6A-D29F-554B-B303-F0452345FE02
var IID_IClipboardStatics2 = syscall.GUID{0xD2AC1B6A, 0xD29F, 0x554B,
	[8]byte{0xB3, 0x03, 0xF0, 0x45, 0x23, 0x45, 0xFE, 0x02}}

type IClipboardStatics2Interface interface {
	win32.IInspectableInterface
	GetHistoryItemsAsync() *IAsyncOperation[*IClipboardHistoryItemsResult]
	ClearHistory() bool
	DeleteItemFromHistory(item *IClipboardHistoryItem) bool
	SetHistoryItemAsContent(item *IClipboardHistoryItem) SetHistoryItemAsContentStatus
	IsHistoryEnabled() bool
	IsRoamingEnabled() bool
	SetContentWithOptions(content *IDataPackage, options *IClipboardContentOptions) bool
	Add_HistoryChanged(handler EventHandler[*IClipboardHistoryChangedEventArgs]) EventRegistrationToken
	Remove_HistoryChanged(token EventRegistrationToken)
	Add_RoamingEnabledChanged(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_RoamingEnabledChanged(token EventRegistrationToken)
	Add_HistoryEnabledChanged(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_HistoryEnabledChanged(token EventRegistrationToken)
}

type IClipboardStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetHistoryItemsAsync         uintptr
	ClearHistory                 uintptr
	DeleteItemFromHistory        uintptr
	SetHistoryItemAsContent      uintptr
	IsHistoryEnabled             uintptr
	IsRoamingEnabled             uintptr
	SetContentWithOptions        uintptr
	Add_HistoryChanged           uintptr
	Remove_HistoryChanged        uintptr
	Add_RoamingEnabledChanged    uintptr
	Remove_RoamingEnabledChanged uintptr
	Add_HistoryEnabledChanged    uintptr
	Remove_HistoryEnabledChanged uintptr
}

type IClipboardStatics2 struct {
	win32.IInspectable
}

func (this *IClipboardStatics2) Vtbl() *IClipboardStatics2Vtbl {
	return (*IClipboardStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClipboardStatics2) GetHistoryItemsAsync() *IAsyncOperation[*IClipboardHistoryItemsResult] {
	var _result *IAsyncOperation[*IClipboardHistoryItemsResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetHistoryItemsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClipboardStatics2) ClearHistory() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearHistory, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClipboardStatics2) DeleteItemFromHistory(item *IClipboardHistoryItem) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DeleteItemFromHistory, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(item)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClipboardStatics2) SetHistoryItemAsContent(item *IClipboardHistoryItem) SetHistoryItemAsContentStatus {
	var _result SetHistoryItemAsContentStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetHistoryItemAsContent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(item)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClipboardStatics2) IsHistoryEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsHistoryEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClipboardStatics2) IsRoamingEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsRoamingEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClipboardStatics2) SetContentWithOptions(content *IDataPackage, options *IClipboardContentOptions) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetContentWithOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(content)), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClipboardStatics2) Add_HistoryChanged(handler EventHandler[*IClipboardHistoryChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_HistoryChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClipboardStatics2) Remove_HistoryChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_HistoryChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IClipboardStatics2) Add_RoamingEnabledChanged(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_RoamingEnabledChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClipboardStatics2) Remove_RoamingEnabledChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_RoamingEnabledChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IClipboardStatics2) Add_HistoryEnabledChanged(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_HistoryEnabledChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClipboardStatics2) Remove_HistoryEnabledChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_HistoryEnabledChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 61EBF5C7-EFEA-4346-9554-981D7E198FFE
var IID_IDataPackage = syscall.GUID{0x61EBF5C7, 0xEFEA, 0x4346,
	[8]byte{0x95, 0x54, 0x98, 0x1D, 0x7E, 0x19, 0x8F, 0xFE}}

type IDataPackageInterface interface {
	win32.IInspectableInterface
	GetView() *IDataPackageView
	Get_Properties() *IDataPackagePropertySet
	Get_RequestedOperation() DataPackageOperation
	Put_RequestedOperation(value DataPackageOperation)
	Add_OperationCompleted(handler TypedEventHandler[*IDataPackage, *IOperationCompletedEventArgs]) EventRegistrationToken
	Remove_OperationCompleted(token EventRegistrationToken)
	Add_Destroyed(handler TypedEventHandler[*IDataPackage, interface{}]) EventRegistrationToken
	Remove_Destroyed(token EventRegistrationToken)
	SetData(formatId string, value interface{})
	SetDataProvider(formatId string, delayRenderer DataProviderHandler)
	SetText(value string)
	SetUri(value *IUriRuntimeClass)
	SetHtmlFormat(value string)
	Get_ResourceMap() *IMap[string, *IRandomAccessStreamReference]
	SetRtf(value string)
	SetBitmap(value *IRandomAccessStreamReference)
	SetStorageItemsReadOnly(value *IIterable[*IStorageItem])
	SetStorageItems(value *IIterable[*IStorageItem], readOnly bool)
}

type IDataPackageVtbl struct {
	win32.IInspectableVtbl
	GetView                   uintptr
	Get_Properties            uintptr
	Get_RequestedOperation    uintptr
	Put_RequestedOperation    uintptr
	Add_OperationCompleted    uintptr
	Remove_OperationCompleted uintptr
	Add_Destroyed             uintptr
	Remove_Destroyed          uintptr
	SetData                   uintptr
	SetDataProvider           uintptr
	SetText                   uintptr
	SetUri                    uintptr
	SetHtmlFormat             uintptr
	Get_ResourceMap           uintptr
	SetRtf                    uintptr
	SetBitmap                 uintptr
	SetStorageItemsReadOnly   uintptr
	SetStorageItems           uintptr
}

type IDataPackage struct {
	win32.IInspectable
}

func (this *IDataPackage) Vtbl() *IDataPackageVtbl {
	return (*IDataPackageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataPackage) GetView() *IDataPackageView {
	var _result *IDataPackageView
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackage) Get_Properties() *IDataPackagePropertySet {
	var _result *IDataPackagePropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackage) Get_RequestedOperation() DataPackageOperation {
	var _result DataPackageOperation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestedOperation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataPackage) Put_RequestedOperation(value DataPackageOperation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RequestedOperation, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDataPackage) Add_OperationCompleted(handler TypedEventHandler[*IDataPackage, *IOperationCompletedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_OperationCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataPackage) Remove_OperationCompleted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_OperationCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IDataPackage) Add_Destroyed(handler TypedEventHandler[*IDataPackage, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Destroyed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataPackage) Remove_Destroyed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Destroyed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IDataPackage) SetData(formatId string, value interface{}) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetData, uintptr(unsafe.Pointer(this)), NewHStr(formatId).Ptr, *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IDataPackage) SetDataProvider(formatId string, delayRenderer DataProviderHandler) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetDataProvider, uintptr(unsafe.Pointer(this)), NewHStr(formatId).Ptr, uintptr(unsafe.Pointer(NewOneArgFuncDelegate(delayRenderer))))
	_ = _hr
}

func (this *IDataPackage) SetText(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetText, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IDataPackage) SetUri(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IDataPackage) SetHtmlFormat(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetHtmlFormat, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IDataPackage) Get_ResourceMap() *IMap[string, *IRandomAccessStreamReference] {
	var _result *IMap[string, *IRandomAccessStreamReference]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceMap, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackage) SetRtf(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetRtf, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IDataPackage) SetBitmap(value *IRandomAccessStreamReference) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetBitmap, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IDataPackage) SetStorageItemsReadOnly(value *IIterable[*IStorageItem]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetStorageItemsReadOnly, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IDataPackage) SetStorageItems(value *IIterable[*IStorageItem], readOnly bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetStorageItems, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)), uintptr(*(*byte)(unsafe.Pointer(&readOnly))))
	_ = _hr
}

// 041C1FE9-2409-45E1-A538-4C53EEEE04A7
var IID_IDataPackage2 = syscall.GUID{0x041C1FE9, 0x2409, 0x45E1,
	[8]byte{0xA5, 0x38, 0x4C, 0x53, 0xEE, 0xEE, 0x04, 0xA7}}

type IDataPackage2Interface interface {
	win32.IInspectableInterface
	SetApplicationLink(value *IUriRuntimeClass)
	SetWebLink(value *IUriRuntimeClass)
}

type IDataPackage2Vtbl struct {
	win32.IInspectableVtbl
	SetApplicationLink uintptr
	SetWebLink         uintptr
}

type IDataPackage2 struct {
	win32.IInspectable
}

func (this *IDataPackage2) Vtbl() *IDataPackage2Vtbl {
	return (*IDataPackage2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataPackage2) SetApplicationLink(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetApplicationLink, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IDataPackage2) SetWebLink(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetWebLink, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 88F31F5D-787B-4D32-965A-A9838105A056
var IID_IDataPackage3 = syscall.GUID{0x88F31F5D, 0x787B, 0x4D32,
	[8]byte{0x96, 0x5A, 0xA9, 0x83, 0x81, 0x05, 0xA0, 0x56}}

type IDataPackage3Interface interface {
	win32.IInspectableInterface
	Add_ShareCompleted(handler TypedEventHandler[*IDataPackage, *IShareCompletedEventArgs]) EventRegistrationToken
	Remove_ShareCompleted(token EventRegistrationToken)
}

type IDataPackage3Vtbl struct {
	win32.IInspectableVtbl
	Add_ShareCompleted    uintptr
	Remove_ShareCompleted uintptr
}

type IDataPackage3 struct {
	win32.IInspectable
}

func (this *IDataPackage3) Vtbl() *IDataPackage3Vtbl {
	return (*IDataPackage3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataPackage3) Add_ShareCompleted(handler TypedEventHandler[*IDataPackage, *IShareCompletedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ShareCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataPackage3) Remove_ShareCompleted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ShareCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 13A24EC8-9382-536F-852A-3045E1B29A3B
var IID_IDataPackage4 = syscall.GUID{0x13A24EC8, 0x9382, 0x536F,
	[8]byte{0x85, 0x2A, 0x30, 0x45, 0xE1, 0xB2, 0x9A, 0x3B}}

type IDataPackage4Interface interface {
	win32.IInspectableInterface
	Add_ShareCanceled(handler TypedEventHandler[*IDataPackage, interface{}]) EventRegistrationToken
	Remove_ShareCanceled(token EventRegistrationToken)
}

type IDataPackage4Vtbl struct {
	win32.IInspectableVtbl
	Add_ShareCanceled    uintptr
	Remove_ShareCanceled uintptr
}

type IDataPackage4 struct {
	win32.IInspectable
}

func (this *IDataPackage4) Vtbl() *IDataPackage4Vtbl {
	return (*IDataPackage4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataPackage4) Add_ShareCanceled(handler TypedEventHandler[*IDataPackage, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ShareCanceled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataPackage4) Remove_ShareCanceled(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ShareCanceled, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// CD1C93EB-4C4C-443A-A8D3-F5C241E91689
var IID_IDataPackagePropertySet = syscall.GUID{0xCD1C93EB, 0x4C4C, 0x443A,
	[8]byte{0xA8, 0xD3, 0xF5, 0xC2, 0x41, 0xE9, 0x16, 0x89}}

type IDataPackagePropertySetInterface interface {
	win32.IInspectableInterface
	Get_Title() string
	Put_Title(value string)
	Get_Description() string
	Put_Description(value string)
	Get_Thumbnail() *IRandomAccessStreamReference
	Put_Thumbnail(value *IRandomAccessStreamReference)
	Get_FileTypes() *IVector[string]
	Get_ApplicationName() string
	Put_ApplicationName(value string)
	Get_ApplicationListingUri() *IUriRuntimeClass
	Put_ApplicationListingUri(value *IUriRuntimeClass)
}

type IDataPackagePropertySetVtbl struct {
	win32.IInspectableVtbl
	Get_Title                 uintptr
	Put_Title                 uintptr
	Get_Description           uintptr
	Put_Description           uintptr
	Get_Thumbnail             uintptr
	Put_Thumbnail             uintptr
	Get_FileTypes             uintptr
	Get_ApplicationName       uintptr
	Put_ApplicationName       uintptr
	Get_ApplicationListingUri uintptr
	Put_ApplicationListingUri uintptr
}

type IDataPackagePropertySet struct {
	win32.IInspectable
}

func (this *IDataPackagePropertySet) Vtbl() *IDataPackagePropertySetVtbl {
	return (*IDataPackagePropertySetVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataPackagePropertySet) Get_Title() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDataPackagePropertySet) Put_Title(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Title, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IDataPackagePropertySet) Get_Description() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Description, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDataPackagePropertySet) Put_Description(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Description, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IDataPackagePropertySet) Get_Thumbnail() *IRandomAccessStreamReference {
	var _result *IRandomAccessStreamReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Thumbnail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackagePropertySet) Put_Thumbnail(value *IRandomAccessStreamReference) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Thumbnail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IDataPackagePropertySet) Get_FileTypes() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FileTypes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackagePropertySet) Get_ApplicationName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ApplicationName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDataPackagePropertySet) Put_ApplicationName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ApplicationName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IDataPackagePropertySet) Get_ApplicationListingUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ApplicationListingUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackagePropertySet) Put_ApplicationListingUri(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ApplicationListingUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// EB505D4A-9800-46AA-B181-7B6F0F2B919A
var IID_IDataPackagePropertySet2 = syscall.GUID{0xEB505D4A, 0x9800, 0x46AA,
	[8]byte{0xB1, 0x81, 0x7B, 0x6F, 0x0F, 0x2B, 0x91, 0x9A}}

type IDataPackagePropertySet2Interface interface {
	win32.IInspectableInterface
	Get_ContentSourceWebLink() *IUriRuntimeClass
	Put_ContentSourceWebLink(value *IUriRuntimeClass)
	Get_ContentSourceApplicationLink() *IUriRuntimeClass
	Put_ContentSourceApplicationLink(value *IUriRuntimeClass)
	Get_PackageFamilyName() string
	Put_PackageFamilyName(value string)
	Get_Square30x30Logo() *IRandomAccessStreamReference
	Put_Square30x30Logo(value *IRandomAccessStreamReference)
	Get_LogoBackgroundColor() unsafe.Pointer
	Put_LogoBackgroundColor(value unsafe.Pointer)
}

type IDataPackagePropertySet2Vtbl struct {
	win32.IInspectableVtbl
	Get_ContentSourceWebLink         uintptr
	Put_ContentSourceWebLink         uintptr
	Get_ContentSourceApplicationLink uintptr
	Put_ContentSourceApplicationLink uintptr
	Get_PackageFamilyName            uintptr
	Put_PackageFamilyName            uintptr
	Get_Square30x30Logo              uintptr
	Put_Square30x30Logo              uintptr
	Get_LogoBackgroundColor          uintptr
	Put_LogoBackgroundColor          uintptr
}

type IDataPackagePropertySet2 struct {
	win32.IInspectable
}

func (this *IDataPackagePropertySet2) Vtbl() *IDataPackagePropertySet2Vtbl {
	return (*IDataPackagePropertySet2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataPackagePropertySet2) Get_ContentSourceWebLink() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentSourceWebLink, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackagePropertySet2) Put_ContentSourceWebLink(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ContentSourceWebLink, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IDataPackagePropertySet2) Get_ContentSourceApplicationLink() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentSourceApplicationLink, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackagePropertySet2) Put_ContentSourceApplicationLink(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ContentSourceApplicationLink, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IDataPackagePropertySet2) Get_PackageFamilyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PackageFamilyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDataPackagePropertySet2) Put_PackageFamilyName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PackageFamilyName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IDataPackagePropertySet2) Get_Square30x30Logo() *IRandomAccessStreamReference {
	var _result *IRandomAccessStreamReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Square30x30Logo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackagePropertySet2) Put_Square30x30Logo(value *IRandomAccessStreamReference) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Square30x30Logo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IDataPackagePropertySet2) Get_LogoBackgroundColor() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LogoBackgroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataPackagePropertySet2) Put_LogoBackgroundColor(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LogoBackgroundColor, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 9E87FD9B-5205-401B-874A-455653BD39E8
var IID_IDataPackagePropertySet3 = syscall.GUID{0x9E87FD9B, 0x5205, 0x401B,
	[8]byte{0x87, 0x4A, 0x45, 0x56, 0x53, 0xBD, 0x39, 0xE8}}

type IDataPackagePropertySet3Interface interface {
	win32.IInspectableInterface
	Get_EnterpriseId() string
	Put_EnterpriseId(value string)
}

type IDataPackagePropertySet3Vtbl struct {
	win32.IInspectableVtbl
	Get_EnterpriseId uintptr
	Put_EnterpriseId uintptr
}

type IDataPackagePropertySet3 struct {
	win32.IInspectable
}

func (this *IDataPackagePropertySet3) Vtbl() *IDataPackagePropertySet3Vtbl {
	return (*IDataPackagePropertySet3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataPackagePropertySet3) Get_EnterpriseId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EnterpriseId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDataPackagePropertySet3) Put_EnterpriseId(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_EnterpriseId, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 6390EBF5-1739-4C74-B22F-865FAB5E8545
var IID_IDataPackagePropertySet4 = syscall.GUID{0x6390EBF5, 0x1739, 0x4C74,
	[8]byte{0xB2, 0x2F, 0x86, 0x5F, 0xAB, 0x5E, 0x85, 0x45}}

type IDataPackagePropertySet4Interface interface {
	win32.IInspectableInterface
	Get_ContentSourceUserActivityJson() string
	Put_ContentSourceUserActivityJson(value string)
}

type IDataPackagePropertySet4Vtbl struct {
	win32.IInspectableVtbl
	Get_ContentSourceUserActivityJson uintptr
	Put_ContentSourceUserActivityJson uintptr
}

type IDataPackagePropertySet4 struct {
	win32.IInspectable
}

func (this *IDataPackagePropertySet4) Vtbl() *IDataPackagePropertySet4Vtbl {
	return (*IDataPackagePropertySet4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataPackagePropertySet4) Get_ContentSourceUserActivityJson() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentSourceUserActivityJson, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDataPackagePropertySet4) Put_ContentSourceUserActivityJson(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ContentSourceUserActivityJson, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// B94CEC01-0C1A-4C57-BE55-75D01289735D
var IID_IDataPackagePropertySetView = syscall.GUID{0xB94CEC01, 0x0C1A, 0x4C57,
	[8]byte{0xBE, 0x55, 0x75, 0xD0, 0x12, 0x89, 0x73, 0x5D}}

type IDataPackagePropertySetViewInterface interface {
	win32.IInspectableInterface
	Get_Title() string
	Get_Description() string
	Get_Thumbnail() *IRandomAccessStreamReference
	Get_FileTypes() *IVectorView[string]
	Get_ApplicationName() string
	Get_ApplicationListingUri() *IUriRuntimeClass
}

type IDataPackagePropertySetViewVtbl struct {
	win32.IInspectableVtbl
	Get_Title                 uintptr
	Get_Description           uintptr
	Get_Thumbnail             uintptr
	Get_FileTypes             uintptr
	Get_ApplicationName       uintptr
	Get_ApplicationListingUri uintptr
}

type IDataPackagePropertySetView struct {
	win32.IInspectable
}

func (this *IDataPackagePropertySetView) Vtbl() *IDataPackagePropertySetViewVtbl {
	return (*IDataPackagePropertySetViewVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataPackagePropertySetView) Get_Title() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDataPackagePropertySetView) Get_Description() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Description, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDataPackagePropertySetView) Get_Thumbnail() *IRandomAccessStreamReference {
	var _result *IRandomAccessStreamReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Thumbnail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackagePropertySetView) Get_FileTypes() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FileTypes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackagePropertySetView) Get_ApplicationName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ApplicationName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDataPackagePropertySetView) Get_ApplicationListingUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ApplicationListingUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6054509B-8EBE-4FEB-9C1E-75E69DE54B84
var IID_IDataPackagePropertySetView2 = syscall.GUID{0x6054509B, 0x8EBE, 0x4FEB,
	[8]byte{0x9C, 0x1E, 0x75, 0xE6, 0x9D, 0xE5, 0x4B, 0x84}}

type IDataPackagePropertySetView2Interface interface {
	win32.IInspectableInterface
	Get_PackageFamilyName() string
	Get_ContentSourceWebLink() *IUriRuntimeClass
	Get_ContentSourceApplicationLink() *IUriRuntimeClass
	Get_Square30x30Logo() *IRandomAccessStreamReference
	Get_LogoBackgroundColor() unsafe.Pointer
}

type IDataPackagePropertySetView2Vtbl struct {
	win32.IInspectableVtbl
	Get_PackageFamilyName            uintptr
	Get_ContentSourceWebLink         uintptr
	Get_ContentSourceApplicationLink uintptr
	Get_Square30x30Logo              uintptr
	Get_LogoBackgroundColor          uintptr
}

type IDataPackagePropertySetView2 struct {
	win32.IInspectable
}

func (this *IDataPackagePropertySetView2) Vtbl() *IDataPackagePropertySetView2Vtbl {
	return (*IDataPackagePropertySetView2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataPackagePropertySetView2) Get_PackageFamilyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PackageFamilyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDataPackagePropertySetView2) Get_ContentSourceWebLink() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentSourceWebLink, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackagePropertySetView2) Get_ContentSourceApplicationLink() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentSourceApplicationLink, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackagePropertySetView2) Get_Square30x30Logo() *IRandomAccessStreamReference {
	var _result *IRandomAccessStreamReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Square30x30Logo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackagePropertySetView2) Get_LogoBackgroundColor() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LogoBackgroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// DB764CE5-D174-495C-84FC-1A51F6AB45D7
var IID_IDataPackagePropertySetView3 = syscall.GUID{0xDB764CE5, 0xD174, 0x495C,
	[8]byte{0x84, 0xFC, 0x1A, 0x51, 0xF6, 0xAB, 0x45, 0xD7}}

type IDataPackagePropertySetView3Interface interface {
	win32.IInspectableInterface
	Get_EnterpriseId() string
}

type IDataPackagePropertySetView3Vtbl struct {
	win32.IInspectableVtbl
	Get_EnterpriseId uintptr
}

type IDataPackagePropertySetView3 struct {
	win32.IInspectable
}

func (this *IDataPackagePropertySetView3) Vtbl() *IDataPackagePropertySetView3Vtbl {
	return (*IDataPackagePropertySetView3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataPackagePropertySetView3) Get_EnterpriseId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EnterpriseId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 4474C80D-D16F-40AE-9580-6F8562B94235
var IID_IDataPackagePropertySetView4 = syscall.GUID{0x4474C80D, 0xD16F, 0x40AE,
	[8]byte{0x95, 0x80, 0x6F, 0x85, 0x62, 0xB9, 0x42, 0x35}}

type IDataPackagePropertySetView4Interface interface {
	win32.IInspectableInterface
	Get_ContentSourceUserActivityJson() string
}

type IDataPackagePropertySetView4Vtbl struct {
	win32.IInspectableVtbl
	Get_ContentSourceUserActivityJson uintptr
}

type IDataPackagePropertySetView4 struct {
	win32.IInspectable
}

func (this *IDataPackagePropertySetView4) Vtbl() *IDataPackagePropertySetView4Vtbl {
	return (*IDataPackagePropertySetView4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataPackagePropertySetView4) Get_ContentSourceUserActivityJson() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentSourceUserActivityJson, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 6F0A9445-3760-50BB-8523-C4202DED7D78
var IID_IDataPackagePropertySetView5 = syscall.GUID{0x6F0A9445, 0x3760, 0x50BB,
	[8]byte{0x85, 0x23, 0xC4, 0x20, 0x2D, 0xED, 0x7D, 0x78}}

type IDataPackagePropertySetView5Interface interface {
	win32.IInspectableInterface
	Get_IsFromRoamingClipboard() bool
}

type IDataPackagePropertySetView5Vtbl struct {
	win32.IInspectableVtbl
	Get_IsFromRoamingClipboard uintptr
}

type IDataPackagePropertySetView5 struct {
	win32.IInspectable
}

func (this *IDataPackagePropertySetView5) Vtbl() *IDataPackagePropertySetView5Vtbl {
	return (*IDataPackagePropertySetView5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataPackagePropertySetView5) Get_IsFromRoamingClipboard() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsFromRoamingClipboard, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7B840471-5900-4D85-A90B-10CB85FE3552
var IID_IDataPackageView = syscall.GUID{0x7B840471, 0x5900, 0x4D85,
	[8]byte{0xA9, 0x0B, 0x10, 0xCB, 0x85, 0xFE, 0x35, 0x52}}

type IDataPackageViewInterface interface {
	win32.IInspectableInterface
	Get_Properties() *IDataPackagePropertySetView
	Get_RequestedOperation() DataPackageOperation
	ReportOperationCompleted(value DataPackageOperation)
	Get_AvailableFormats() *IVectorView[string]
	Contains(formatId string) bool
	GetDataAsync(formatId string) *IAsyncOperation[interface{}]
	GetTextAsync() *IAsyncOperation[string]
	GetCustomTextAsync(formatId string) *IAsyncOperation[string]
	GetUriAsync() *IAsyncOperation[*IUriRuntimeClass]
	GetHtmlFormatAsync() *IAsyncOperation[string]
	GetResourceMapAsync() *IAsyncOperation[*IMapView[string, *IRandomAccessStreamReference]]
	GetRtfAsync() *IAsyncOperation[string]
	GetBitmapAsync() *IAsyncOperation[*IRandomAccessStreamReference]
	GetStorageItemsAsync() *IAsyncOperation[*IVectorView[*IStorageItem]]
}

type IDataPackageViewVtbl struct {
	win32.IInspectableVtbl
	Get_Properties           uintptr
	Get_RequestedOperation   uintptr
	ReportOperationCompleted uintptr
	Get_AvailableFormats     uintptr
	Contains                 uintptr
	GetDataAsync             uintptr
	GetTextAsync             uintptr
	GetCustomTextAsync       uintptr
	GetUriAsync              uintptr
	GetHtmlFormatAsync       uintptr
	GetResourceMapAsync      uintptr
	GetRtfAsync              uintptr
	GetBitmapAsync           uintptr
	GetStorageItemsAsync     uintptr
}

type IDataPackageView struct {
	win32.IInspectable
}

func (this *IDataPackageView) Vtbl() *IDataPackageViewVtbl {
	return (*IDataPackageViewVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataPackageView) Get_Properties() *IDataPackagePropertySetView {
	var _result *IDataPackagePropertySetView
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackageView) Get_RequestedOperation() DataPackageOperation {
	var _result DataPackageOperation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestedOperation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataPackageView) ReportOperationCompleted(value DataPackageOperation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReportOperationCompleted, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDataPackageView) Get_AvailableFormats() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AvailableFormats, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackageView) Contains(formatId string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Contains, uintptr(unsafe.Pointer(this)), NewHStr(formatId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataPackageView) GetDataAsync(formatId string) *IAsyncOperation[interface{}] {
	var _result *IAsyncOperation[interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDataAsync, uintptr(unsafe.Pointer(this)), NewHStr(formatId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackageView) GetTextAsync() *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTextAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackageView) GetCustomTextAsync(formatId string) *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCustomTextAsync, uintptr(unsafe.Pointer(this)), NewHStr(formatId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackageView) GetUriAsync() *IAsyncOperation[*IUriRuntimeClass] {
	var _result *IAsyncOperation[*IUriRuntimeClass]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetUriAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackageView) GetHtmlFormatAsync() *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetHtmlFormatAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackageView) GetResourceMapAsync() *IAsyncOperation[*IMapView[string, *IRandomAccessStreamReference]] {
	var _result *IAsyncOperation[*IMapView[string, *IRandomAccessStreamReference]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetResourceMapAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackageView) GetRtfAsync() *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetRtfAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackageView) GetBitmapAsync() *IAsyncOperation[*IRandomAccessStreamReference] {
	var _result *IAsyncOperation[*IRandomAccessStreamReference]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetBitmapAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackageView) GetStorageItemsAsync() *IAsyncOperation[*IVectorView[*IStorageItem]] {
	var _result *IAsyncOperation[*IVectorView[*IStorageItem]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStorageItemsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 40ECBA95-2450-4C1D-B6B4-ED45463DEE9C
var IID_IDataPackageView2 = syscall.GUID{0x40ECBA95, 0x2450, 0x4C1D,
	[8]byte{0xB6, 0xB4, 0xED, 0x45, 0x46, 0x3D, 0xEE, 0x9C}}

type IDataPackageView2Interface interface {
	win32.IInspectableInterface
	GetApplicationLinkAsync() *IAsyncOperation[*IUriRuntimeClass]
	GetWebLinkAsync() *IAsyncOperation[*IUriRuntimeClass]
}

type IDataPackageView2Vtbl struct {
	win32.IInspectableVtbl
	GetApplicationLinkAsync uintptr
	GetWebLinkAsync         uintptr
}

type IDataPackageView2 struct {
	win32.IInspectable
}

func (this *IDataPackageView2) Vtbl() *IDataPackageView2Vtbl {
	return (*IDataPackageView2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataPackageView2) GetApplicationLinkAsync() *IAsyncOperation[*IUriRuntimeClass] {
	var _result *IAsyncOperation[*IUriRuntimeClass]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetApplicationLinkAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackageView2) GetWebLinkAsync() *IAsyncOperation[*IUriRuntimeClass] {
	var _result *IAsyncOperation[*IUriRuntimeClass]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetWebLinkAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D37771A8-DDAD-4288-8428-D1CAE394128B
var IID_IDataPackageView3 = syscall.GUID{0xD37771A8, 0xDDAD, 0x4288,
	[8]byte{0x84, 0x28, 0xD1, 0xCA, 0xE3, 0x94, 0x12, 0x8B}}

type IDataPackageView3Interface interface {
	win32.IInspectableInterface
	RequestAccessAsync() *IAsyncOperation[ProtectionPolicyEvaluationResult]
	RequestAccessWithEnterpriseIdAsync(enterpriseId string) *IAsyncOperation[ProtectionPolicyEvaluationResult]
	UnlockAndAssumeEnterpriseIdentity() ProtectionPolicyEvaluationResult
}

type IDataPackageView3Vtbl struct {
	win32.IInspectableVtbl
	RequestAccessAsync                 uintptr
	RequestAccessWithEnterpriseIdAsync uintptr
	UnlockAndAssumeEnterpriseIdentity  uintptr
}

type IDataPackageView3 struct {
	win32.IInspectable
}

func (this *IDataPackageView3) Vtbl() *IDataPackageView3Vtbl {
	return (*IDataPackageView3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataPackageView3) RequestAccessAsync() *IAsyncOperation[ProtectionPolicyEvaluationResult] {
	var _result *IAsyncOperation[ProtectionPolicyEvaluationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackageView3) RequestAccessWithEnterpriseIdAsync(enterpriseId string) *IAsyncOperation[ProtectionPolicyEvaluationResult] {
	var _result *IAsyncOperation[ProtectionPolicyEvaluationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessWithEnterpriseIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(enterpriseId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPackageView3) UnlockAndAssumeEnterpriseIdentity() ProtectionPolicyEvaluationResult {
	var _result ProtectionPolicyEvaluationResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UnlockAndAssumeEnterpriseIdentity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// DFE96F1F-E042-4433-A09F-26D6FFDA8B85
var IID_IDataPackageView4 = syscall.GUID{0xDFE96F1F, 0xE042, 0x4433,
	[8]byte{0xA0, 0x9F, 0x26, 0xD6, 0xFF, 0xDA, 0x8B, 0x85}}

type IDataPackageView4Interface interface {
	win32.IInspectableInterface
	SetAcceptedFormatId(formatId string)
}

type IDataPackageView4Vtbl struct {
	win32.IInspectableVtbl
	SetAcceptedFormatId uintptr
}

type IDataPackageView4 struct {
	win32.IInspectable
}

func (this *IDataPackageView4) Vtbl() *IDataPackageView4Vtbl {
	return (*IDataPackageView4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataPackageView4) SetAcceptedFormatId(formatId string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetAcceptedFormatId, uintptr(unsafe.Pointer(this)), NewHStr(formatId).Ptr)
	_ = _hr
}

// C2CF2373-2D26-43D9-B69D-DCB86D03F6DA
var IID_IDataProviderDeferral = syscall.GUID{0xC2CF2373, 0x2D26, 0x43D9,
	[8]byte{0xB6, 0x9D, 0xDC, 0xB8, 0x6D, 0x03, 0xF6, 0xDA}}

type IDataProviderDeferralInterface interface {
	win32.IInspectableInterface
	Complete()
}

type IDataProviderDeferralVtbl struct {
	win32.IInspectableVtbl
	Complete uintptr
}

type IDataProviderDeferral struct {
	win32.IInspectable
}

func (this *IDataProviderDeferral) Vtbl() *IDataProviderDeferralVtbl {
	return (*IDataProviderDeferralVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataProviderDeferral) Complete() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Complete, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// EBBC7157-D3C8-47DA-ACDE-F82388D5F716
var IID_IDataProviderRequest = syscall.GUID{0xEBBC7157, 0xD3C8, 0x47DA,
	[8]byte{0xAC, 0xDE, 0xF8, 0x23, 0x88, 0xD5, 0xF7, 0x16}}

type IDataProviderRequestInterface interface {
	win32.IInspectableInterface
	Get_FormatId() string
	Get_Deadline() DateTime
	GetDeferral() *IDataProviderDeferral
	SetData(value interface{})
}

type IDataProviderRequestVtbl struct {
	win32.IInspectableVtbl
	Get_FormatId uintptr
	Get_Deadline uintptr
	GetDeferral  uintptr
	SetData      uintptr
}

type IDataProviderRequest struct {
	win32.IInspectable
}

func (this *IDataProviderRequest) Vtbl() *IDataProviderRequestVtbl {
	return (*IDataProviderRequestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataProviderRequest) Get_FormatId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FormatId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDataProviderRequest) Get_Deadline() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Deadline, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataProviderRequest) GetDeferral() *IDataProviderDeferral {
	var _result *IDataProviderDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataProviderRequest) SetData(value interface{}) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetData, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

// 4341AE3B-FC12-4E53-8C02-AC714C415A27
var IID_IDataRequest = syscall.GUID{0x4341AE3B, 0xFC12, 0x4E53,
	[8]byte{0x8C, 0x02, 0xAC, 0x71, 0x4C, 0x41, 0x5A, 0x27}}

type IDataRequestInterface interface {
	win32.IInspectableInterface
	Get_Data() *IDataPackage
	Put_Data(value *IDataPackage)
	Get_Deadline() DateTime
	FailWithDisplayText(value string)
	GetDeferral() *IDataRequestDeferral
}

type IDataRequestVtbl struct {
	win32.IInspectableVtbl
	Get_Data            uintptr
	Put_Data            uintptr
	Get_Deadline        uintptr
	FailWithDisplayText uintptr
	GetDeferral         uintptr
}

type IDataRequest struct {
	win32.IInspectable
}

func (this *IDataRequest) Vtbl() *IDataRequestVtbl {
	return (*IDataRequestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataRequest) Get_Data() *IDataPackage {
	var _result *IDataPackage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataRequest) Put_Data(value *IDataPackage) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IDataRequest) Get_Deadline() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Deadline, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataRequest) FailWithDisplayText(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FailWithDisplayText, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IDataRequest) GetDeferral() *IDataRequestDeferral {
	var _result *IDataRequestDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6DC4B89F-0386-4263-87C1-ED7DCE30890E
var IID_IDataRequestDeferral = syscall.GUID{0x6DC4B89F, 0x0386, 0x4263,
	[8]byte{0x87, 0xC1, 0xED, 0x7D, 0xCE, 0x30, 0x89, 0x0E}}

type IDataRequestDeferralInterface interface {
	win32.IInspectableInterface
	Complete()
}

type IDataRequestDeferralVtbl struct {
	win32.IInspectableVtbl
	Complete uintptr
}

type IDataRequestDeferral struct {
	win32.IInspectable
}

func (this *IDataRequestDeferral) Vtbl() *IDataRequestDeferralVtbl {
	return (*IDataRequestDeferralVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataRequestDeferral) Complete() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Complete, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// CB8BA807-6AC5-43C9-8AC5-9BA232163182
var IID_IDataRequestedEventArgs = syscall.GUID{0xCB8BA807, 0x6AC5, 0x43C9,
	[8]byte{0x8A, 0xC5, 0x9B, 0xA2, 0x32, 0x16, 0x31, 0x82}}

type IDataRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Request() *IDataRequest
}

type IDataRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Request uintptr
}

type IDataRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IDataRequestedEventArgs) Vtbl() *IDataRequestedEventArgsVtbl {
	return (*IDataRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataRequestedEventArgs) Get_Request() *IDataRequest {
	var _result *IDataRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Request, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A5CAEE9B-8708-49D1-8D36-67D25A8DA00C
var IID_IDataTransferManager = syscall.GUID{0xA5CAEE9B, 0x8708, 0x49D1,
	[8]byte{0x8D, 0x36, 0x67, 0xD2, 0x5A, 0x8D, 0xA0, 0x0C}}

type IDataTransferManagerInterface interface {
	win32.IInspectableInterface
	Add_DataRequested(handler TypedEventHandler[*IDataTransferManager, *IDataRequestedEventArgs]) EventRegistrationToken
	Remove_DataRequested(token EventRegistrationToken)
	Add_TargetApplicationChosen(handler TypedEventHandler[*IDataTransferManager, *ITargetApplicationChosenEventArgs]) EventRegistrationToken
	Remove_TargetApplicationChosen(token EventRegistrationToken)
}

type IDataTransferManagerVtbl struct {
	win32.IInspectableVtbl
	Add_DataRequested              uintptr
	Remove_DataRequested           uintptr
	Add_TargetApplicationChosen    uintptr
	Remove_TargetApplicationChosen uintptr
}

type IDataTransferManager struct {
	win32.IInspectable
}

func (this *IDataTransferManager) Vtbl() *IDataTransferManagerVtbl {
	return (*IDataTransferManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataTransferManager) Add_DataRequested(handler TypedEventHandler[*IDataTransferManager, *IDataRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DataRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataTransferManager) Remove_DataRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DataRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IDataTransferManager) Add_TargetApplicationChosen(handler TypedEventHandler[*IDataTransferManager, *ITargetApplicationChosenEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_TargetApplicationChosen, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataTransferManager) Remove_TargetApplicationChosen(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_TargetApplicationChosen, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 30AE7D71-8BA8-4C02-8E3F-DDB23B388715
var IID_IDataTransferManager2 = syscall.GUID{0x30AE7D71, 0x8BA8, 0x4C02,
	[8]byte{0x8E, 0x3F, 0xDD, 0xB2, 0x3B, 0x38, 0x87, 0x15}}

type IDataTransferManager2Interface interface {
	win32.IInspectableInterface
	Add_ShareProvidersRequested(handler TypedEventHandler[*IDataTransferManager, *IShareProvidersRequestedEventArgs]) EventRegistrationToken
	Remove_ShareProvidersRequested(token EventRegistrationToken)
}

type IDataTransferManager2Vtbl struct {
	win32.IInspectableVtbl
	Add_ShareProvidersRequested    uintptr
	Remove_ShareProvidersRequested uintptr
}

type IDataTransferManager2 struct {
	win32.IInspectable
}

func (this *IDataTransferManager2) Vtbl() *IDataTransferManager2Vtbl {
	return (*IDataTransferManager2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataTransferManager2) Add_ShareProvidersRequested(handler TypedEventHandler[*IDataTransferManager, *IShareProvidersRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ShareProvidersRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataTransferManager2) Remove_ShareProvidersRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ShareProvidersRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// A9DA01AA-E00E-4CFE-AA44-2DD932DCA3D8
var IID_IDataTransferManagerStatics = syscall.GUID{0xA9DA01AA, 0xE00E, 0x4CFE,
	[8]byte{0xAA, 0x44, 0x2D, 0xD9, 0x32, 0xDC, 0xA3, 0xD8}}

type IDataTransferManagerStaticsInterface interface {
	win32.IInspectableInterface
	ShowShareUI()
	GetForCurrentView() *IDataTransferManager
}

type IDataTransferManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	ShowShareUI       uintptr
	GetForCurrentView uintptr
}

type IDataTransferManagerStatics struct {
	win32.IInspectable
}

func (this *IDataTransferManagerStatics) Vtbl() *IDataTransferManagerStaticsVtbl {
	return (*IDataTransferManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataTransferManagerStatics) ShowShareUI() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowShareUI, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IDataTransferManagerStatics) GetForCurrentView() *IDataTransferManager {
	var _result *IDataTransferManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C54EC2EC-9F97-4D63-9868-395E271AD8F5
var IID_IDataTransferManagerStatics2 = syscall.GUID{0xC54EC2EC, 0x9F97, 0x4D63,
	[8]byte{0x98, 0x68, 0x39, 0x5E, 0x27, 0x1A, 0xD8, 0xF5}}

type IDataTransferManagerStatics2Interface interface {
	win32.IInspectableInterface
	IsSupported() bool
}

type IDataTransferManagerStatics2Vtbl struct {
	win32.IInspectableVtbl
	IsSupported uintptr
}

type IDataTransferManagerStatics2 struct {
	win32.IInspectable
}

func (this *IDataTransferManagerStatics2) Vtbl() *IDataTransferManagerStatics2Vtbl {
	return (*IDataTransferManagerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataTransferManagerStatics2) IsSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 05845473-6C82-4F5C-AC23-62E458361FAC
var IID_IDataTransferManagerStatics3 = syscall.GUID{0x05845473, 0x6C82, 0x4F5C,
	[8]byte{0xAC, 0x23, 0x62, 0xE4, 0x58, 0x36, 0x1F, 0xAC}}

type IDataTransferManagerStatics3Interface interface {
	win32.IInspectableInterface
	ShowShareUIWithOptions(options *IShareUIOptions)
}

type IDataTransferManagerStatics3Vtbl struct {
	win32.IInspectableVtbl
	ShowShareUIWithOptions uintptr
}

type IDataTransferManagerStatics3 struct {
	win32.IInspectable
}

func (this *IDataTransferManagerStatics3) Vtbl() *IDataTransferManagerStatics3Vtbl {
	return (*IDataTransferManagerStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataTransferManagerStatics3) ShowShareUIWithOptions(options *IShareUIOptions) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowShareUIWithOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(options)))
	_ = _hr
}

// E22E7749-DD70-446F-AEFC-61CEE59F655E
var IID_IHtmlFormatHelperStatics = syscall.GUID{0xE22E7749, 0xDD70, 0x446F,
	[8]byte{0xAE, 0xFC, 0x61, 0xCE, 0xE5, 0x9F, 0x65, 0x5E}}

type IHtmlFormatHelperStaticsInterface interface {
	win32.IInspectableInterface
	GetStaticFragment(htmlFormat string) string
	CreateHtmlFormat(htmlFragment string) string
}

type IHtmlFormatHelperStaticsVtbl struct {
	win32.IInspectableVtbl
	GetStaticFragment uintptr
	CreateHtmlFormat  uintptr
}

type IHtmlFormatHelperStatics struct {
	win32.IInspectable
}

func (this *IHtmlFormatHelperStatics) Vtbl() *IHtmlFormatHelperStaticsVtbl {
	return (*IHtmlFormatHelperStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHtmlFormatHelperStatics) GetStaticFragment(htmlFormat string) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStaticFragment, uintptr(unsafe.Pointer(this)), NewHStr(htmlFormat).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHtmlFormatHelperStatics) CreateHtmlFormat(htmlFragment string) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateHtmlFormat, uintptr(unsafe.Pointer(this)), NewHStr(htmlFragment).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// E7AF329D-051D-4FAB-B1A9-47FD77F70A41
var IID_IOperationCompletedEventArgs = syscall.GUID{0xE7AF329D, 0x051D, 0x4FAB,
	[8]byte{0xB1, 0xA9, 0x47, 0xFD, 0x77, 0xF7, 0x0A, 0x41}}

type IOperationCompletedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Operation() DataPackageOperation
}

type IOperationCompletedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Operation uintptr
}

type IOperationCompletedEventArgs struct {
	win32.IInspectable
}

func (this *IOperationCompletedEventArgs) Vtbl() *IOperationCompletedEventArgsVtbl {
	return (*IOperationCompletedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IOperationCompletedEventArgs) Get_Operation() DataPackageOperation {
	var _result DataPackageOperation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Operation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 858FA073-1E19-4105-B2F7-C8478808D562
var IID_IOperationCompletedEventArgs2 = syscall.GUID{0x858FA073, 0x1E19, 0x4105,
	[8]byte{0xB2, 0xF7, 0xC8, 0x47, 0x88, 0x08, 0xD5, 0x62}}

type IOperationCompletedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_AcceptedFormatId() string
}

type IOperationCompletedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_AcceptedFormatId uintptr
}

type IOperationCompletedEventArgs2 struct {
	win32.IInspectable
}

func (this *IOperationCompletedEventArgs2) Vtbl() *IOperationCompletedEventArgs2Vtbl {
	return (*IOperationCompletedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IOperationCompletedEventArgs2) Get_AcceptedFormatId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AcceptedFormatId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 4574C442-F913-4F60-9DF7-CC4060AB1916
var IID_IShareCompletedEventArgs = syscall.GUID{0x4574C442, 0xF913, 0x4F60,
	[8]byte{0x9D, 0xF7, 0xCC, 0x40, 0x60, 0xAB, 0x19, 0x16}}

type IShareCompletedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ShareTarget() *IShareTargetInfo
}

type IShareCompletedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ShareTarget uintptr
}

type IShareCompletedEventArgs struct {
	win32.IInspectable
}

func (this *IShareCompletedEventArgs) Vtbl() *IShareCompletedEventArgsVtbl {
	return (*IShareCompletedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IShareCompletedEventArgs) Get_ShareTarget() *IShareTargetInfo {
	var _result *IShareTargetInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShareTarget, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2FABE026-443E-4CDA-AF25-8D81070EFD80
var IID_IShareProvider = syscall.GUID{0x2FABE026, 0x443E, 0x4CDA,
	[8]byte{0xAF, 0x25, 0x8D, 0x81, 0x07, 0x0E, 0xFD, 0x80}}

type IShareProviderInterface interface {
	win32.IInspectableInterface
	Get_Title() string
	Get_DisplayIcon() *IRandomAccessStreamReference
	Get_BackgroundColor() unsafe.Pointer
	Get_Tag() interface{}
	Put_Tag(value interface{})
}

type IShareProviderVtbl struct {
	win32.IInspectableVtbl
	Get_Title           uintptr
	Get_DisplayIcon     uintptr
	Get_BackgroundColor uintptr
	Get_Tag             uintptr
	Put_Tag             uintptr
}

type IShareProvider struct {
	win32.IInspectable
}

func (this *IShareProvider) Vtbl() *IShareProviderVtbl {
	return (*IShareProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IShareProvider) Get_Title() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IShareProvider) Get_DisplayIcon() *IRandomAccessStreamReference {
	var _result *IRandomAccessStreamReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayIcon, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IShareProvider) Get_BackgroundColor() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BackgroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IShareProvider) Get_Tag() interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Tag, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IShareProvider) Put_Tag(value interface{}) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Tag, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

// 172A174C-E79E-4F6D-B07D-128F469E0296
var IID_IShareProviderFactory = syscall.GUID{0x172A174C, 0xE79E, 0x4F6D,
	[8]byte{0xB0, 0x7D, 0x12, 0x8F, 0x46, 0x9E, 0x02, 0x96}}

type IShareProviderFactoryInterface interface {
	win32.IInspectableInterface
	Create(title string, displayIcon *IRandomAccessStreamReference, backgroundColor unsafe.Pointer, handler ShareProviderHandler) *IShareProvider
}

type IShareProviderFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IShareProviderFactory struct {
	win32.IInspectable
}

func (this *IShareProviderFactory) Vtbl() *IShareProviderFactoryVtbl {
	return (*IShareProviderFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IShareProviderFactory) Create(title string, displayIcon *IRandomAccessStreamReference, backgroundColor unsafe.Pointer, handler ShareProviderHandler) *IShareProvider {
	var _result *IShareProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(title).Ptr, uintptr(unsafe.Pointer(displayIcon)), uintptr(backgroundColor), uintptr(unsafe.Pointer(NewOneArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 19CEF937-D435-4179-B6AF-14E0492B69F6
var IID_IShareProviderOperation = syscall.GUID{0x19CEF937, 0xD435, 0x4179,
	[8]byte{0xB6, 0xAF, 0x14, 0xE0, 0x49, 0x2B, 0x69, 0xF6}}

type IShareProviderOperationInterface interface {
	win32.IInspectableInterface
	Get_Data() *IDataPackageView
	Get_Provider() *IShareProvider
	ReportCompleted()
}

type IShareProviderOperationVtbl struct {
	win32.IInspectableVtbl
	Get_Data        uintptr
	Get_Provider    uintptr
	ReportCompleted uintptr
}

type IShareProviderOperation struct {
	win32.IInspectable
}

func (this *IShareProviderOperation) Vtbl() *IShareProviderOperationVtbl {
	return (*IShareProviderOperationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IShareProviderOperation) Get_Data() *IDataPackageView {
	var _result *IDataPackageView
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IShareProviderOperation) Get_Provider() *IShareProvider {
	var _result *IShareProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Provider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IShareProviderOperation) ReportCompleted() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReportCompleted, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// F888F356-A3F8-4FCE-85E4-8826E63BE799
var IID_IShareProvidersRequestedEventArgs = syscall.GUID{0xF888F356, 0xA3F8, 0x4FCE,
	[8]byte{0x85, 0xE4, 0x88, 0x26, 0xE6, 0x3B, 0xE7, 0x99}}

type IShareProvidersRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Providers() *IVector[*IShareProvider]
	Get_Data() *IDataPackageView
	GetDeferral() *IDeferral
}

type IShareProvidersRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Providers uintptr
	Get_Data      uintptr
	GetDeferral   uintptr
}

type IShareProvidersRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IShareProvidersRequestedEventArgs) Vtbl() *IShareProvidersRequestedEventArgsVtbl {
	return (*IShareProvidersRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IShareProvidersRequestedEventArgs) Get_Providers() *IVector[*IShareProvider] {
	var _result *IVector[*IShareProvider]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Providers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IShareProvidersRequestedEventArgs) Get_Data() *IDataPackageView {
	var _result *IDataPackageView
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IShareProvidersRequestedEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 385BE607-C6E8-4114-B294-28F3BB6F9904
var IID_IShareTargetInfo = syscall.GUID{0x385BE607, 0xC6E8, 0x4114,
	[8]byte{0xB2, 0x94, 0x28, 0xF3, 0xBB, 0x6F, 0x99, 0x04}}

type IShareTargetInfoInterface interface {
	win32.IInspectableInterface
	Get_AppUserModelId() string
	Get_ShareProvider() *IShareProvider
}

type IShareTargetInfoVtbl struct {
	win32.IInspectableVtbl
	Get_AppUserModelId uintptr
	Get_ShareProvider  uintptr
}

type IShareTargetInfo struct {
	win32.IInspectable
}

func (this *IShareTargetInfo) Vtbl() *IShareTargetInfoVtbl {
	return (*IShareTargetInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IShareTargetInfo) Get_AppUserModelId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppUserModelId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IShareTargetInfo) Get_ShareProvider() *IShareProvider {
	var _result *IShareProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShareProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 72FA8A80-342F-4D90-9551-2AE04E37680C
var IID_IShareUIOptions = syscall.GUID{0x72FA8A80, 0x342F, 0x4D90,
	[8]byte{0x95, 0x51, 0x2A, 0xE0, 0x4E, 0x37, 0x68, 0x0C}}

type IShareUIOptionsInterface interface {
	win32.IInspectableInterface
	Get_Theme() ShareUITheme
	Put_Theme(value ShareUITheme)
	Get_SelectionRect() *IReference[Rect]
	Put_SelectionRect(value *IReference[Rect])
}

type IShareUIOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_Theme         uintptr
	Put_Theme         uintptr
	Get_SelectionRect uintptr
	Put_SelectionRect uintptr
}

type IShareUIOptions struct {
	win32.IInspectable
}

func (this *IShareUIOptions) Vtbl() *IShareUIOptionsVtbl {
	return (*IShareUIOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IShareUIOptions) Get_Theme() ShareUITheme {
	var _result ShareUITheme
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Theme, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IShareUIOptions) Put_Theme(value ShareUITheme) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Theme, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IShareUIOptions) Get_SelectionRect() *IReference[Rect] {
	var _result *IReference[Rect]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectionRect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IShareUIOptions) Put_SelectionRect(value *IReference[Rect]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SelectionRect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// C6132ADA-34B1-4849-BD5F-D09FEE3158C5
var IID_ISharedStorageAccessManagerStatics = syscall.GUID{0xC6132ADA, 0x34B1, 0x4849,
	[8]byte{0xBD, 0x5F, 0xD0, 0x9F, 0xEE, 0x31, 0x58, 0xC5}}

type ISharedStorageAccessManagerStaticsInterface interface {
	win32.IInspectableInterface
	AddFile(file *IStorageFile) string
	RedeemTokenForFileAsync(token string) *IAsyncOperation[*IStorageFile]
	RemoveFile(token string)
}

type ISharedStorageAccessManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	AddFile                 uintptr
	RedeemTokenForFileAsync uintptr
	RemoveFile              uintptr
}

type ISharedStorageAccessManagerStatics struct {
	win32.IInspectable
}

func (this *ISharedStorageAccessManagerStatics) Vtbl() *ISharedStorageAccessManagerStaticsVtbl {
	return (*ISharedStorageAccessManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISharedStorageAccessManagerStatics) AddFile(file *IStorageFile) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddFile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISharedStorageAccessManagerStatics) RedeemTokenForFileAsync(token string) *IAsyncOperation[*IStorageFile] {
	var _result *IAsyncOperation[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RedeemTokenForFileAsync, uintptr(unsafe.Pointer(this)), NewHStr(token).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISharedStorageAccessManagerStatics) RemoveFile(token string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveFile, uintptr(unsafe.Pointer(this)), NewHStr(token).Ptr)
	_ = _hr
}

// 7ED681A1-A880-40C9-B4ED-0BEE1E15F549
var IID_IStandardDataFormatsStatics = syscall.GUID{0x7ED681A1, 0xA880, 0x40C9,
	[8]byte{0xB4, 0xED, 0x0B, 0xEE, 0x1E, 0x15, 0xF5, 0x49}}

type IStandardDataFormatsStaticsInterface interface {
	win32.IInspectableInterface
	Get_Text() string
	Get_Uri() string
	Get_Html() string
	Get_Rtf() string
	Get_Bitmap() string
	Get_StorageItems() string
}

type IStandardDataFormatsStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Text         uintptr
	Get_Uri          uintptr
	Get_Html         uintptr
	Get_Rtf          uintptr
	Get_Bitmap       uintptr
	Get_StorageItems uintptr
}

type IStandardDataFormatsStatics struct {
	win32.IInspectable
}

func (this *IStandardDataFormatsStatics) Vtbl() *IStandardDataFormatsStaticsVtbl {
	return (*IStandardDataFormatsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStandardDataFormatsStatics) Get_Text() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Text, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStandardDataFormatsStatics) Get_Uri() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Uri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStandardDataFormatsStatics) Get_Html() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Html, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStandardDataFormatsStatics) Get_Rtf() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Rtf, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStandardDataFormatsStatics) Get_Bitmap() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bitmap, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStandardDataFormatsStatics) Get_StorageItems() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StorageItems, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 42A254F4-9D76-42E8-861B-47C25DD0CF71
var IID_IStandardDataFormatsStatics2 = syscall.GUID{0x42A254F4, 0x9D76, 0x42E8,
	[8]byte{0x86, 0x1B, 0x47, 0xC2, 0x5D, 0xD0, 0xCF, 0x71}}

type IStandardDataFormatsStatics2Interface interface {
	win32.IInspectableInterface
	Get_WebLink() string
	Get_ApplicationLink() string
}

type IStandardDataFormatsStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_WebLink         uintptr
	Get_ApplicationLink uintptr
}

type IStandardDataFormatsStatics2 struct {
	win32.IInspectable
}

func (this *IStandardDataFormatsStatics2) Vtbl() *IStandardDataFormatsStatics2Vtbl {
	return (*IStandardDataFormatsStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStandardDataFormatsStatics2) Get_WebLink() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WebLink, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStandardDataFormatsStatics2) Get_ApplicationLink() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ApplicationLink, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 3B57B069-01D4-474C-8B5F-BC8E27F38B21
var IID_IStandardDataFormatsStatics3 = syscall.GUID{0x3B57B069, 0x01D4, 0x474C,
	[8]byte{0x8B, 0x5F, 0xBC, 0x8E, 0x27, 0xF3, 0x8B, 0x21}}

type IStandardDataFormatsStatics3Interface interface {
	win32.IInspectableInterface
	Get_UserActivityJsonArray() string
}

type IStandardDataFormatsStatics3Vtbl struct {
	win32.IInspectableVtbl
	Get_UserActivityJsonArray uintptr
}

type IStandardDataFormatsStatics3 struct {
	win32.IInspectable
}

func (this *IStandardDataFormatsStatics3) Vtbl() *IStandardDataFormatsStatics3Vtbl {
	return (*IStandardDataFormatsStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStandardDataFormatsStatics3) Get_UserActivityJsonArray() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserActivityJsonArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// CA6FB8AC-2987-4EE3-9C54-D8AFBCB86C1D
var IID_ITargetApplicationChosenEventArgs = syscall.GUID{0xCA6FB8AC, 0x2987, 0x4EE3,
	[8]byte{0x9C, 0x54, 0xD8, 0xAF, 0xBC, 0xB8, 0x6C, 0x1D}}

type ITargetApplicationChosenEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ApplicationName() string
}

type ITargetApplicationChosenEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ApplicationName uintptr
}

type ITargetApplicationChosenEventArgs struct {
	win32.IInspectable
}

func (this *ITargetApplicationChosenEventArgs) Vtbl() *ITargetApplicationChosenEventArgsVtbl {
	return (*ITargetApplicationChosenEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITargetApplicationChosenEventArgs) Get_ApplicationName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ApplicationName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// classes

type Clipboard struct {
	RtClass
}

func NewIClipboardStatics() *IClipboardStatics {
	var p *IClipboardStatics
	hs := NewHStr("Windows.ApplicationModel.DataTransfer.Clipboard")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IClipboardStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIClipboardStatics2() *IClipboardStatics2 {
	var p *IClipboardStatics2
	hs := NewHStr("Windows.ApplicationModel.DataTransfer.Clipboard")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IClipboardStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type DataPackage struct {
	RtClass
	*IDataPackage
}

func NewDataPackage() *DataPackage {
	hs := NewHStr("Windows.ApplicationModel.DataTransfer.DataPackage")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &DataPackage{
		RtClass:      RtClass{PInspect: p},
		IDataPackage: (*IDataPackage)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type DataPackagePropertySet struct {
	RtClass
	*IDataPackagePropertySet
}

type DataPackagePropertySetView struct {
	RtClass
	*IDataPackagePropertySetView
}

type DataPackageView struct {
	RtClass
	*IDataPackageView
}

type DataProviderDeferral struct {
	RtClass
	*IDataProviderDeferral
}

type DataProviderRequest struct {
	RtClass
	*IDataProviderRequest
}

type HtmlFormatHelper struct {
	RtClass
}

func NewIHtmlFormatHelperStatics() *IHtmlFormatHelperStatics {
	var p *IHtmlFormatHelperStatics
	hs := NewHStr("Windows.ApplicationModel.DataTransfer.HtmlFormatHelper")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHtmlFormatHelperStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type OperationCompletedEventArgs struct {
	RtClass
	*IOperationCompletedEventArgs
}

type ShareUIOptions struct {
	RtClass
	*IShareUIOptions
}

func NewShareUIOptions() *ShareUIOptions {
	hs := NewHStr("Windows.ApplicationModel.DataTransfer.ShareUIOptions")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &ShareUIOptions{
		RtClass:         RtClass{PInspect: p},
		IShareUIOptions: (*IShareUIOptions)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type StandardDataFormats struct {
	RtClass
}

func NewIStandardDataFormatsStatics2() *IStandardDataFormatsStatics2 {
	var p *IStandardDataFormatsStatics2
	hs := NewHStr("Windows.ApplicationModel.DataTransfer.StandardDataFormats")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IStandardDataFormatsStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIStandardDataFormatsStatics() *IStandardDataFormatsStatics {
	var p *IStandardDataFormatsStatics
	hs := NewHStr("Windows.ApplicationModel.DataTransfer.StandardDataFormats")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IStandardDataFormatsStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIStandardDataFormatsStatics3() *IStandardDataFormatsStatics3 {
	var p *IStandardDataFormatsStatics3
	hs := NewHStr("Windows.ApplicationModel.DataTransfer.StandardDataFormats")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IStandardDataFormatsStatics3, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
