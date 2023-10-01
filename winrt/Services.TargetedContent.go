package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type TargetedContentAppInstallationState int32

const (
	TargetedContentAppInstallationState_NotApplicable TargetedContentAppInstallationState = 0
	TargetedContentAppInstallationState_NotInstalled  TargetedContentAppInstallationState = 1
	TargetedContentAppInstallationState_Installed     TargetedContentAppInstallationState = 2
)

// enum
type TargetedContentAvailability int32

const (
	TargetedContentAvailability_None    TargetedContentAvailability = 0
	TargetedContentAvailability_Partial TargetedContentAvailability = 1
	TargetedContentAvailability_All     TargetedContentAvailability = 2
)

// enum
type TargetedContentInteraction int32

const (
	TargetedContentInteraction_Impression   TargetedContentInteraction = 0
	TargetedContentInteraction_ClickThrough TargetedContentInteraction = 1
	TargetedContentInteraction_Hover        TargetedContentInteraction = 2
	TargetedContentInteraction_Like         TargetedContentInteraction = 3
	TargetedContentInteraction_Dislike      TargetedContentInteraction = 4
	TargetedContentInteraction_Dismiss      TargetedContentInteraction = 5
	TargetedContentInteraction_Ineligible   TargetedContentInteraction = 6
	TargetedContentInteraction_Accept       TargetedContentInteraction = 7
	TargetedContentInteraction_Decline      TargetedContentInteraction = 8
	TargetedContentInteraction_Defer        TargetedContentInteraction = 9
	TargetedContentInteraction_Canceled     TargetedContentInteraction = 10
	TargetedContentInteraction_Conversion   TargetedContentInteraction = 11
	TargetedContentInteraction_Opportunity  TargetedContentInteraction = 12
)

// enum
type TargetedContentObjectKind int32

const (
	TargetedContentObjectKind_Collection TargetedContentObjectKind = 0
	TargetedContentObjectKind_Item       TargetedContentObjectKind = 1
	TargetedContentObjectKind_Value      TargetedContentObjectKind = 2
)

// enum
type TargetedContentValueKind int32

const (
	TargetedContentValueKind_String     TargetedContentValueKind = 0
	TargetedContentValueKind_Uri        TargetedContentValueKind = 1
	TargetedContentValueKind_Number     TargetedContentValueKind = 2
	TargetedContentValueKind_Boolean    TargetedContentValueKind = 3
	TargetedContentValueKind_File       TargetedContentValueKind = 4
	TargetedContentValueKind_ImageFile  TargetedContentValueKind = 5
	TargetedContentValueKind_Action     TargetedContentValueKind = 6
	TargetedContentValueKind_Strings    TargetedContentValueKind = 7
	TargetedContentValueKind_Uris       TargetedContentValueKind = 8
	TargetedContentValueKind_Numbers    TargetedContentValueKind = 9
	TargetedContentValueKind_Booleans   TargetedContentValueKind = 10
	TargetedContentValueKind_Files      TargetedContentValueKind = 11
	TargetedContentValueKind_ImageFiles TargetedContentValueKind = 12
	TargetedContentValueKind_Actions    TargetedContentValueKind = 13
)

// structs

type TargetedContentContract struct {
}

// interfaces

// D75B691E-6CD6-4CA0-9D8F-4728B0B7E6B6
var IID_ITargetedContentAction = syscall.GUID{0xD75B691E, 0x6CD6, 0x4CA0,
	[8]byte{0x9D, 0x8F, 0x47, 0x28, 0xB0, 0xB7, 0xE6, 0xB6}}

type ITargetedContentActionInterface interface {
	win32.IInspectableInterface
	InvokeAsync() *IAsyncAction
}

type ITargetedContentActionVtbl struct {
	win32.IInspectableVtbl
	InvokeAsync uintptr
}

type ITargetedContentAction struct {
	win32.IInspectable
}

func (this *ITargetedContentAction) Vtbl() *ITargetedContentActionVtbl {
	return (*ITargetedContentActionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITargetedContentAction) InvokeAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InvokeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E0F59D26-5927-4450-965C-1CEB7BECDE65
var IID_ITargetedContentAvailabilityChangedEventArgs = syscall.GUID{0xE0F59D26, 0x5927, 0x4450,
	[8]byte{0x96, 0x5C, 0x1C, 0xEB, 0x7B, 0xEC, 0xDE, 0x65}}

type ITargetedContentAvailabilityChangedEventArgsInterface interface {
	win32.IInspectableInterface
	GetDeferral() *IDeferral
}

type ITargetedContentAvailabilityChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	GetDeferral uintptr
}

type ITargetedContentAvailabilityChangedEventArgs struct {
	win32.IInspectable
}

func (this *ITargetedContentAvailabilityChangedEventArgs) Vtbl() *ITargetedContentAvailabilityChangedEventArgsVtbl {
	return (*ITargetedContentAvailabilityChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITargetedContentAvailabilityChangedEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 99D488C9-587E-4586-8EF7-B54CA9453A16
var IID_ITargetedContentChangedEventArgs = syscall.GUID{0x99D488C9, 0x587E, 0x4586,
	[8]byte{0x8E, 0xF7, 0xB5, 0x4C, 0xA9, 0x45, 0x3A, 0x16}}

type ITargetedContentChangedEventArgsInterface interface {
	win32.IInspectableInterface
	GetDeferral() *IDeferral
	Get_HasPreviousContentExpired() bool
}

type ITargetedContentChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	GetDeferral                   uintptr
	Get_HasPreviousContentExpired uintptr
}

type ITargetedContentChangedEventArgs struct {
	win32.IInspectable
}

func (this *ITargetedContentChangedEventArgs) Vtbl() *ITargetedContentChangedEventArgsVtbl {
	return (*ITargetedContentChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITargetedContentChangedEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITargetedContentChangedEventArgs) Get_HasPreviousContentExpired() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasPreviousContentExpired, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 2D4B66C5-F163-44BA-9F6E-E1A4C2BB559D
var IID_ITargetedContentCollection = syscall.GUID{0x2D4B66C5, 0xF163, 0x44BA,
	[8]byte{0x9F, 0x6E, 0xE1, 0xA4, 0xC2, 0xBB, 0x55, 0x9D}}

type ITargetedContentCollectionInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	ReportInteraction(interaction TargetedContentInteraction)
	ReportCustomInteraction(customInteractionName string)
	Get_Path() string
	Get_Properties() *IMapView[string, *ITargetedContentValue]
	Get_Collections() *IVectorView[*ITargetedContentCollection]
	Get_Items() *IVectorView[*ITargetedContentItem]
}

type ITargetedContentCollectionVtbl struct {
	win32.IInspectableVtbl
	Get_Id                  uintptr
	ReportInteraction       uintptr
	ReportCustomInteraction uintptr
	Get_Path                uintptr
	Get_Properties          uintptr
	Get_Collections         uintptr
	Get_Items               uintptr
}

type ITargetedContentCollection struct {
	win32.IInspectable
}

func (this *ITargetedContentCollection) Vtbl() *ITargetedContentCollectionVtbl {
	return (*ITargetedContentCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITargetedContentCollection) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ITargetedContentCollection) ReportInteraction(interaction TargetedContentInteraction) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReportInteraction, uintptr(unsafe.Pointer(this)), uintptr(interaction))
	_ = _hr
}

func (this *ITargetedContentCollection) ReportCustomInteraction(customInteractionName string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReportCustomInteraction, uintptr(unsafe.Pointer(this)), NewHStr(customInteractionName).Ptr)
	_ = _hr
}

func (this *ITargetedContentCollection) Get_Path() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Path, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ITargetedContentCollection) Get_Properties() *IMapView[string, *ITargetedContentValue] {
	var _result *IMapView[string, *ITargetedContentValue]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITargetedContentCollection) Get_Collections() *IVectorView[*ITargetedContentCollection] {
	var _result *IVectorView[*ITargetedContentCollection]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Collections, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITargetedContentCollection) Get_Items() *IVectorView[*ITargetedContentItem] {
	var _result *IVectorView[*ITargetedContentItem]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Items, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BC2494C9-8837-47C2-850F-D79D64595926
var IID_ITargetedContentContainer = syscall.GUID{0xBC2494C9, 0x8837, 0x47C2,
	[8]byte{0x85, 0x0F, 0xD7, 0x9D, 0x64, 0x59, 0x59, 0x26}}

type ITargetedContentContainerInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Get_Timestamp() DateTime
	Get_Availability() TargetedContentAvailability
	Get_Content() *ITargetedContentCollection
	SelectSingleObject(path string) *ITargetedContentObject
}

type ITargetedContentContainerVtbl struct {
	win32.IInspectableVtbl
	Get_Id             uintptr
	Get_Timestamp      uintptr
	Get_Availability   uintptr
	Get_Content        uintptr
	SelectSingleObject uintptr
}

type ITargetedContentContainer struct {
	win32.IInspectable
}

func (this *ITargetedContentContainer) Vtbl() *ITargetedContentContainerVtbl {
	return (*ITargetedContentContainerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITargetedContentContainer) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ITargetedContentContainer) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITargetedContentContainer) Get_Availability() TargetedContentAvailability {
	var _result TargetedContentAvailability
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Availability, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITargetedContentContainer) Get_Content() *ITargetedContentCollection {
	var _result *ITargetedContentCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Content, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITargetedContentContainer) SelectSingleObject(path string) *ITargetedContentObject {
	var _result *ITargetedContentObject
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SelectSingleObject, uintptr(unsafe.Pointer(this)), NewHStr(path).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5B47E7FB-2140-4C1F-A736-C59583F227D8
var IID_ITargetedContentContainerStatics = syscall.GUID{0x5B47E7FB, 0x2140, 0x4C1F,
	[8]byte{0xA7, 0x36, 0xC5, 0x95, 0x83, 0xF2, 0x27, 0xD8}}

type ITargetedContentContainerStaticsInterface interface {
	win32.IInspectableInterface
	GetAsync(contentId string) *IAsyncOperation[*ITargetedContentContainer]
}

type ITargetedContentContainerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetAsync uintptr
}

type ITargetedContentContainerStatics struct {
	win32.IInspectable
}

func (this *ITargetedContentContainerStatics) Vtbl() *ITargetedContentContainerStaticsVtbl {
	return (*ITargetedContentContainerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITargetedContentContainerStatics) GetAsync(contentId string) *IAsyncOperation[*ITargetedContentContainer] {
	var _result *IAsyncOperation[*ITargetedContentContainer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAsync, uintptr(unsafe.Pointer(this)), NewHStr(contentId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A7A585D9-779F-4B1E-BBB1-8EAF53FBEAB2
var IID_ITargetedContentImage = syscall.GUID{0xA7A585D9, 0x779F, 0x4B1E,
	[8]byte{0xBB, 0xB1, 0x8E, 0xAF, 0x53, 0xFB, 0xEA, 0xB2}}

type ITargetedContentImageInterface interface {
	win32.IInspectableInterface
	Get_Height() uint32
	Get_Width() uint32
}

type ITargetedContentImageVtbl struct {
	win32.IInspectableVtbl
	Get_Height uintptr
	Get_Width  uintptr
}

type ITargetedContentImage struct {
	win32.IInspectable
}

func (this *ITargetedContentImage) Vtbl() *ITargetedContentImageVtbl {
	return (*ITargetedContentImageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITargetedContentImage) Get_Height() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Height, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITargetedContentImage) Get_Width() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Width, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 38168DC4-276C-4C32-96BA-565C6E406E74
var IID_ITargetedContentItem = syscall.GUID{0x38168DC4, 0x276C, 0x4C32,
	[8]byte{0x96, 0xBA, 0x56, 0x5C, 0x6E, 0x40, 0x6E, 0x74}}

type ITargetedContentItemInterface interface {
	win32.IInspectableInterface
	Get_Path() string
	ReportInteraction(interaction TargetedContentInteraction)
	ReportCustomInteraction(customInteractionName string)
	Get_State() *ITargetedContentItemState
	Get_Properties() *IMapView[string, *ITargetedContentValue]
	Get_Collections() *IVectorView[*ITargetedContentCollection]
}

type ITargetedContentItemVtbl struct {
	win32.IInspectableVtbl
	Get_Path                uintptr
	ReportInteraction       uintptr
	ReportCustomInteraction uintptr
	Get_State               uintptr
	Get_Properties          uintptr
	Get_Collections         uintptr
}

type ITargetedContentItem struct {
	win32.IInspectable
}

func (this *ITargetedContentItem) Vtbl() *ITargetedContentItemVtbl {
	return (*ITargetedContentItemVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITargetedContentItem) Get_Path() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Path, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ITargetedContentItem) ReportInteraction(interaction TargetedContentInteraction) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReportInteraction, uintptr(unsafe.Pointer(this)), uintptr(interaction))
	_ = _hr
}

func (this *ITargetedContentItem) ReportCustomInteraction(customInteractionName string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReportCustomInteraction, uintptr(unsafe.Pointer(this)), NewHStr(customInteractionName).Ptr)
	_ = _hr
}

func (this *ITargetedContentItem) Get_State() *ITargetedContentItemState {
	var _result *ITargetedContentItemState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITargetedContentItem) Get_Properties() *IMapView[string, *ITargetedContentValue] {
	var _result *IMapView[string, *ITargetedContentValue]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITargetedContentItem) Get_Collections() *IVectorView[*ITargetedContentCollection] {
	var _result *IVectorView[*ITargetedContentCollection]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Collections, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 73935454-4C65-4B47-A441-472DE53C79B6
var IID_ITargetedContentItemState = syscall.GUID{0x73935454, 0x4C65, 0x4B47,
	[8]byte{0xA4, 0x41, 0x47, 0x2D, 0xE5, 0x3C, 0x79, 0xB6}}

type ITargetedContentItemStateInterface interface {
	win32.IInspectableInterface
	Get_ShouldDisplay() bool
	Get_AppInstallationState() TargetedContentAppInstallationState
}

type ITargetedContentItemStateVtbl struct {
	win32.IInspectableVtbl
	Get_ShouldDisplay        uintptr
	Get_AppInstallationState uintptr
}

type ITargetedContentItemState struct {
	win32.IInspectable
}

func (this *ITargetedContentItemState) Vtbl() *ITargetedContentItemStateVtbl {
	return (*ITargetedContentItemStateVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITargetedContentItemState) Get_ShouldDisplay() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShouldDisplay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITargetedContentItemState) Get_AppInstallationState() TargetedContentAppInstallationState {
	var _result TargetedContentAppInstallationState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppInstallationState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 041D7969-2212-42D1-9DFA-88A8E3033AA3
var IID_ITargetedContentObject = syscall.GUID{0x041D7969, 0x2212, 0x42D1,
	[8]byte{0x9D, 0xFA, 0x88, 0xA8, 0xE3, 0x03, 0x3A, 0xA3}}

type ITargetedContentObjectInterface interface {
	win32.IInspectableInterface
	Get_ObjectKind() TargetedContentObjectKind
	Get_Collection() *ITargetedContentCollection
	Get_Item() *ITargetedContentItem
	Get_Value() *ITargetedContentValue
}

type ITargetedContentObjectVtbl struct {
	win32.IInspectableVtbl
	Get_ObjectKind uintptr
	Get_Collection uintptr
	Get_Item       uintptr
	Get_Value      uintptr
}

type ITargetedContentObject struct {
	win32.IInspectable
}

func (this *ITargetedContentObject) Vtbl() *ITargetedContentObjectVtbl {
	return (*ITargetedContentObjectVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITargetedContentObject) Get_ObjectKind() TargetedContentObjectKind {
	var _result TargetedContentObjectKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ObjectKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITargetedContentObject) Get_Collection() *ITargetedContentCollection {
	var _result *ITargetedContentCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Collection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITargetedContentObject) Get_Item() *ITargetedContentItem {
	var _result *ITargetedContentItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Item, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITargetedContentObject) Get_Value() *ITargetedContentValue {
	var _result *ITargetedContentValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9A1CEF3D-8073-4416-8DF2-546835A6414F
var IID_ITargetedContentStateChangedEventArgs = syscall.GUID{0x9A1CEF3D, 0x8073, 0x4416,
	[8]byte{0x8D, 0xF2, 0x54, 0x68, 0x35, 0xA6, 0x41, 0x4F}}

type ITargetedContentStateChangedEventArgsInterface interface {
	win32.IInspectableInterface
	GetDeferral() *IDeferral
}

type ITargetedContentStateChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	GetDeferral uintptr
}

type ITargetedContentStateChangedEventArgs struct {
	win32.IInspectable
}

func (this *ITargetedContentStateChangedEventArgs) Vtbl() *ITargetedContentStateChangedEventArgsVtbl {
	return (*ITargetedContentStateChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITargetedContentStateChangedEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 882C2C49-C652-4C7A-ACAD-1F7FA2986C73
var IID_ITargetedContentSubscription = syscall.GUID{0x882C2C49, 0xC652, 0x4C7A,
	[8]byte{0xAC, 0xAD, 0x1F, 0x7F, 0xA2, 0x98, 0x6C, 0x73}}

type ITargetedContentSubscriptionInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	GetContentContainerAsync() *IAsyncOperation[*ITargetedContentContainer]
	Add_ContentChanged(handler TypedEventHandler[*ITargetedContentSubscription, *ITargetedContentChangedEventArgs]) EventRegistrationToken
	Remove_ContentChanged(cookie EventRegistrationToken)
	Add_AvailabilityChanged(handler TypedEventHandler[*ITargetedContentSubscription, *ITargetedContentAvailabilityChangedEventArgs]) EventRegistrationToken
	Remove_AvailabilityChanged(cookie EventRegistrationToken)
	Add_StateChanged(handler TypedEventHandler[*ITargetedContentSubscription, *ITargetedContentStateChangedEventArgs]) EventRegistrationToken
	Remove_StateChanged(cookie EventRegistrationToken)
}

type ITargetedContentSubscriptionVtbl struct {
	win32.IInspectableVtbl
	Get_Id                     uintptr
	GetContentContainerAsync   uintptr
	Add_ContentChanged         uintptr
	Remove_ContentChanged      uintptr
	Add_AvailabilityChanged    uintptr
	Remove_AvailabilityChanged uintptr
	Add_StateChanged           uintptr
	Remove_StateChanged        uintptr
}

type ITargetedContentSubscription struct {
	win32.IInspectable
}

func (this *ITargetedContentSubscription) Vtbl() *ITargetedContentSubscriptionVtbl {
	return (*ITargetedContentSubscriptionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITargetedContentSubscription) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ITargetedContentSubscription) GetContentContainerAsync() *IAsyncOperation[*ITargetedContentContainer] {
	var _result *IAsyncOperation[*ITargetedContentContainer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetContentContainerAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITargetedContentSubscription) Add_ContentChanged(handler TypedEventHandler[*ITargetedContentSubscription, *ITargetedContentChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ContentChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITargetedContentSubscription) Remove_ContentChanged(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ContentChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *ITargetedContentSubscription) Add_AvailabilityChanged(handler TypedEventHandler[*ITargetedContentSubscription, *ITargetedContentAvailabilityChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AvailabilityChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITargetedContentSubscription) Remove_AvailabilityChanged(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AvailabilityChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *ITargetedContentSubscription) Add_StateChanged(handler TypedEventHandler[*ITargetedContentSubscription, *ITargetedContentStateChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITargetedContentSubscription) Remove_StateChanged(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

// 61EE6AD0-2C83-421B-8467-413EAF1AEB97
var IID_ITargetedContentSubscriptionOptions = syscall.GUID{0x61EE6AD0, 0x2C83, 0x421B,
	[8]byte{0x84, 0x67, 0x41, 0x3E, 0xAF, 0x1A, 0xEB, 0x97}}

type ITargetedContentSubscriptionOptionsInterface interface {
	win32.IInspectableInterface
	Get_SubscriptionId() string
	Get_AllowPartialContentAvailability() bool
	Put_AllowPartialContentAvailability(value bool)
	Get_CloudQueryParameters() *IMap[string, string]
	Get_LocalFilters() *IVector[string]
	Update()
}

type ITargetedContentSubscriptionOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_SubscriptionId                  uintptr
	Get_AllowPartialContentAvailability uintptr
	Put_AllowPartialContentAvailability uintptr
	Get_CloudQueryParameters            uintptr
	Get_LocalFilters                    uintptr
	Update                              uintptr
}

type ITargetedContentSubscriptionOptions struct {
	win32.IInspectable
}

func (this *ITargetedContentSubscriptionOptions) Vtbl() *ITargetedContentSubscriptionOptionsVtbl {
	return (*ITargetedContentSubscriptionOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITargetedContentSubscriptionOptions) Get_SubscriptionId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SubscriptionId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ITargetedContentSubscriptionOptions) Get_AllowPartialContentAvailability() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllowPartialContentAvailability, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITargetedContentSubscriptionOptions) Put_AllowPartialContentAvailability(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AllowPartialContentAvailability, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ITargetedContentSubscriptionOptions) Get_CloudQueryParameters() *IMap[string, string] {
	var _result *IMap[string, string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CloudQueryParameters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITargetedContentSubscriptionOptions) Get_LocalFilters() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalFilters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITargetedContentSubscriptionOptions) Update() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Update, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// FADDFE80-360D-4916-B53C-7EA27090D02A
var IID_ITargetedContentSubscriptionStatics = syscall.GUID{0xFADDFE80, 0x360D, 0x4916,
	[8]byte{0xB5, 0x3C, 0x7E, 0xA2, 0x70, 0x90, 0xD0, 0x2A}}

type ITargetedContentSubscriptionStaticsInterface interface {
	win32.IInspectableInterface
	GetAsync(subscriptionId string) *IAsyncOperation[*ITargetedContentSubscription]
	GetOptions(subscriptionId string) *ITargetedContentSubscriptionOptions
}

type ITargetedContentSubscriptionStaticsVtbl struct {
	win32.IInspectableVtbl
	GetAsync   uintptr
	GetOptions uintptr
}

type ITargetedContentSubscriptionStatics struct {
	win32.IInspectable
}

func (this *ITargetedContentSubscriptionStatics) Vtbl() *ITargetedContentSubscriptionStaticsVtbl {
	return (*ITargetedContentSubscriptionStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITargetedContentSubscriptionStatics) GetAsync(subscriptionId string) *IAsyncOperation[*ITargetedContentSubscription] {
	var _result *IAsyncOperation[*ITargetedContentSubscription]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAsync, uintptr(unsafe.Pointer(this)), NewHStr(subscriptionId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITargetedContentSubscriptionStatics) GetOptions(subscriptionId string) *ITargetedContentSubscriptionOptions {
	var _result *ITargetedContentSubscriptionOptions
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetOptions, uintptr(unsafe.Pointer(this)), NewHStr(subscriptionId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AAFDE4B3-4215-4BF8-867F-43F04865F9BF
var IID_ITargetedContentValue = syscall.GUID{0xAAFDE4B3, 0x4215, 0x4BF8,
	[8]byte{0x86, 0x7F, 0x43, 0xF0, 0x48, 0x65, 0xF9, 0xBF}}

type ITargetedContentValueInterface interface {
	win32.IInspectableInterface
	Get_ValueKind() TargetedContentValueKind
	Get_Path() string
	Get_String() string
	Get_Uri() *IUriRuntimeClass
	Get_Number() float64
	Get_Boolean() bool
	Get_File() *IRandomAccessStreamReference
	Get_ImageFile() *ITargetedContentImage
	Get_Action() *ITargetedContentAction
	Get_Strings() *IVectorView[string]
	Get_Uris() *IVectorView[*IUriRuntimeClass]
	Get_Numbers() *IVectorView[float64]
	Get_Booleans() *IVectorView[bool]
	Get_Files() *IVectorView[*IRandomAccessStreamReference]
	Get_ImageFiles() *IVectorView[*ITargetedContentImage]
	Get_Actions() *IVectorView[*ITargetedContentAction]
}

type ITargetedContentValueVtbl struct {
	win32.IInspectableVtbl
	Get_ValueKind  uintptr
	Get_Path       uintptr
	Get_String     uintptr
	Get_Uri        uintptr
	Get_Number     uintptr
	Get_Boolean    uintptr
	Get_File       uintptr
	Get_ImageFile  uintptr
	Get_Action     uintptr
	Get_Strings    uintptr
	Get_Uris       uintptr
	Get_Numbers    uintptr
	Get_Booleans   uintptr
	Get_Files      uintptr
	Get_ImageFiles uintptr
	Get_Actions    uintptr
}

type ITargetedContentValue struct {
	win32.IInspectable
}

func (this *ITargetedContentValue) Vtbl() *ITargetedContentValueVtbl {
	return (*ITargetedContentValueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITargetedContentValue) Get_ValueKind() TargetedContentValueKind {
	var _result TargetedContentValueKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ValueKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITargetedContentValue) Get_Path() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Path, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ITargetedContentValue) Get_String() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_String, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ITargetedContentValue) Get_Uri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Uri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITargetedContentValue) Get_Number() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Number, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITargetedContentValue) Get_Boolean() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Boolean, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITargetedContentValue) Get_File() *IRandomAccessStreamReference {
	var _result *IRandomAccessStreamReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_File, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITargetedContentValue) Get_ImageFile() *ITargetedContentImage {
	var _result *ITargetedContentImage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ImageFile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITargetedContentValue) Get_Action() *ITargetedContentAction {
	var _result *ITargetedContentAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Action, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITargetedContentValue) Get_Strings() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Strings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITargetedContentValue) Get_Uris() *IVectorView[*IUriRuntimeClass] {
	var _result *IVectorView[*IUriRuntimeClass]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Uris, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITargetedContentValue) Get_Numbers() *IVectorView[float64] {
	var _result *IVectorView[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Numbers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITargetedContentValue) Get_Booleans() *IVectorView[bool] {
	var _result *IVectorView[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Booleans, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITargetedContentValue) Get_Files() *IVectorView[*IRandomAccessStreamReference] {
	var _result *IVectorView[*IRandomAccessStreamReference]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Files, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITargetedContentValue) Get_ImageFiles() *IVectorView[*ITargetedContentImage] {
	var _result *IVectorView[*ITargetedContentImage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ImageFiles, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITargetedContentValue) Get_Actions() *IVectorView[*ITargetedContentAction] {
	var _result *IVectorView[*ITargetedContentAction]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Actions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type TargetedContentAction struct {
	RtClass
	*ITargetedContentAction
}

type TargetedContentAvailabilityChangedEventArgs struct {
	RtClass
	*ITargetedContentAvailabilityChangedEventArgs
}

type TargetedContentChangedEventArgs struct {
	RtClass
	*ITargetedContentChangedEventArgs
}

type TargetedContentCollection struct {
	RtClass
	*ITargetedContentCollection
}

type TargetedContentContainer struct {
	RtClass
	*ITargetedContentContainer
}

func NewITargetedContentContainerStatics() *ITargetedContentContainerStatics {
	var p *ITargetedContentContainerStatics
	hs := NewHStr("Windows.Services.TargetedContent.TargetedContentContainer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITargetedContentContainerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type TargetedContentFile struct {
	RtClass
	*IRandomAccessStreamReference
}

type TargetedContentImage struct {
	RtClass
	*ITargetedContentImage
}

type TargetedContentItem struct {
	RtClass
	*ITargetedContentItem
}

type TargetedContentItemState struct {
	RtClass
	*ITargetedContentItemState
}

type TargetedContentObject struct {
	RtClass
	*ITargetedContentObject
}

type TargetedContentStateChangedEventArgs struct {
	RtClass
	*ITargetedContentStateChangedEventArgs
}

type TargetedContentSubscription struct {
	RtClass
	*ITargetedContentSubscription
}

func NewITargetedContentSubscriptionStatics() *ITargetedContentSubscriptionStatics {
	var p *ITargetedContentSubscriptionStatics
	hs := NewHStr("Windows.Services.TargetedContent.TargetedContentSubscription")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITargetedContentSubscriptionStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type TargetedContentSubscriptionOptions struct {
	RtClass
	*ITargetedContentSubscriptionOptions
}

type TargetedContentValue struct {
	RtClass
	*ITargetedContentValue
}
