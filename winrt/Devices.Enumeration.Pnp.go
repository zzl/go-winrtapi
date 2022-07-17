package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type PnpObjectType int32

const (
	PnpObjectType_Unknown                      PnpObjectType = 0
	PnpObjectType_DeviceInterface              PnpObjectType = 1
	PnpObjectType_DeviceContainer              PnpObjectType = 2
	PnpObjectType_Device                       PnpObjectType = 3
	PnpObjectType_DeviceInterfaceClass         PnpObjectType = 4
	PnpObjectType_AssociationEndpoint          PnpObjectType = 5
	PnpObjectType_AssociationEndpointContainer PnpObjectType = 6
	PnpObjectType_AssociationEndpointService   PnpObjectType = 7
	PnpObjectType_DevicePanel                  PnpObjectType = 8
)

// interfaces

// 95C66258-733B-4A8F-93A3-DB078AC870C1
var IID_IPnpObject = syscall.GUID{0x95C66258, 0x733B, 0x4A8F,
	[8]byte{0x93, 0xA3, 0xDB, 0x07, 0x8A, 0xC8, 0x70, 0xC1}}

type IPnpObjectInterface interface {
	win32.IInspectableInterface
	Get_Type() PnpObjectType
	Get_Id() string
	Get_Properties() *IMapView[string, interface{}]
	Update(updateInfo *IPnpObjectUpdate)
}

type IPnpObjectVtbl struct {
	win32.IInspectableVtbl
	Get_Type       uintptr
	Get_Id         uintptr
	Get_Properties uintptr
	Update         uintptr
}

type IPnpObject struct {
	win32.IInspectable
}

func (this *IPnpObject) Vtbl() *IPnpObjectVtbl {
	return (*IPnpObjectVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPnpObject) Get_Type() PnpObjectType {
	var _result PnpObjectType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Type, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPnpObject) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPnpObject) Get_Properties() *IMapView[string, interface{}] {
	var _result *IMapView[string, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPnpObject) Update(updateInfo *IPnpObjectUpdate) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Update, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(updateInfo)))
	_ = _hr
}

// B3C32A3D-D168-4660-BBF3-A733B14B6E01
var IID_IPnpObjectStatics = syscall.GUID{0xB3C32A3D, 0xD168, 0x4660,
	[8]byte{0xBB, 0xF3, 0xA7, 0x33, 0xB1, 0x4B, 0x6E, 0x01}}

type IPnpObjectStaticsInterface interface {
	win32.IInspectableInterface
	CreateFromIdAsync(type_ PnpObjectType, id string, requestedProperties *IIterable[string]) *IAsyncOperation[*IPnpObject]
	FindAllAsync(type_ PnpObjectType, requestedProperties *IIterable[string]) *IAsyncOperation[*IVectorView[*IPnpObject]]
	FindAllAsyncAqsFilter(type_ PnpObjectType, requestedProperties *IIterable[string], aqsFilter string) *IAsyncOperation[*IVectorView[*IPnpObject]]
	CreateWatcher(type_ PnpObjectType, requestedProperties *IIterable[string]) *IPnpObjectWatcher
	CreateWatcherAqsFilter(type_ PnpObjectType, requestedProperties *IIterable[string], aqsFilter string) *IPnpObjectWatcher
}

type IPnpObjectStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateFromIdAsync      uintptr
	FindAllAsync           uintptr
	FindAllAsyncAqsFilter  uintptr
	CreateWatcher          uintptr
	CreateWatcherAqsFilter uintptr
}

type IPnpObjectStatics struct {
	win32.IInspectable
}

func (this *IPnpObjectStatics) Vtbl() *IPnpObjectStaticsVtbl {
	return (*IPnpObjectStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPnpObjectStatics) CreateFromIdAsync(type_ PnpObjectType, id string, requestedProperties *IIterable[string]) *IAsyncOperation[*IPnpObject] {
	var _result *IAsyncOperation[*IPnpObject]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromIdAsync, uintptr(unsafe.Pointer(this)), uintptr(type_), NewHStr(id).Ptr, uintptr(unsafe.Pointer(requestedProperties)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPnpObjectStatics) FindAllAsync(type_ PnpObjectType, requestedProperties *IIterable[string]) *IAsyncOperation[*IVectorView[*IPnpObject]] {
	var _result *IAsyncOperation[*IVectorView[*IPnpObject]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllAsync, uintptr(unsafe.Pointer(this)), uintptr(type_), uintptr(unsafe.Pointer(requestedProperties)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPnpObjectStatics) FindAllAsyncAqsFilter(type_ PnpObjectType, requestedProperties *IIterable[string], aqsFilter string) *IAsyncOperation[*IVectorView[*IPnpObject]] {
	var _result *IAsyncOperation[*IVectorView[*IPnpObject]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllAsyncAqsFilter, uintptr(unsafe.Pointer(this)), uintptr(type_), uintptr(unsafe.Pointer(requestedProperties)), NewHStr(aqsFilter).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPnpObjectStatics) CreateWatcher(type_ PnpObjectType, requestedProperties *IIterable[string]) *IPnpObjectWatcher {
	var _result *IPnpObjectWatcher
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWatcher, uintptr(unsafe.Pointer(this)), uintptr(type_), uintptr(unsafe.Pointer(requestedProperties)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPnpObjectStatics) CreateWatcherAqsFilter(type_ PnpObjectType, requestedProperties *IIterable[string], aqsFilter string) *IPnpObjectWatcher {
	var _result *IPnpObjectWatcher
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWatcherAqsFilter, uintptr(unsafe.Pointer(this)), uintptr(type_), uintptr(unsafe.Pointer(requestedProperties)), NewHStr(aqsFilter).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6F59E812-001E-4844-BCC6-432886856A17
var IID_IPnpObjectUpdate = syscall.GUID{0x6F59E812, 0x001E, 0x4844,
	[8]byte{0xBC, 0xC6, 0x43, 0x28, 0x86, 0x85, 0x6A, 0x17}}

type IPnpObjectUpdateInterface interface {
	win32.IInspectableInterface
	Get_Type() PnpObjectType
	Get_Id() string
	Get_Properties() *IMapView[string, interface{}]
}

type IPnpObjectUpdateVtbl struct {
	win32.IInspectableVtbl
	Get_Type       uintptr
	Get_Id         uintptr
	Get_Properties uintptr
}

type IPnpObjectUpdate struct {
	win32.IInspectable
}

func (this *IPnpObjectUpdate) Vtbl() *IPnpObjectUpdateVtbl {
	return (*IPnpObjectUpdateVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPnpObjectUpdate) Get_Type() PnpObjectType {
	var _result PnpObjectType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Type, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPnpObjectUpdate) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPnpObjectUpdate) Get_Properties() *IMapView[string, interface{}] {
	var _result *IMapView[string, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 83C95CA8-4772-4A7A-ACA8-E48C42A89C44
var IID_IPnpObjectWatcher = syscall.GUID{0x83C95CA8, 0x4772, 0x4A7A,
	[8]byte{0xAC, 0xA8, 0xE4, 0x8C, 0x42, 0xA8, 0x9C, 0x44}}

type IPnpObjectWatcherInterface interface {
	win32.IInspectableInterface
	Add_Added(handler TypedEventHandler[*IPnpObjectWatcher, *IPnpObject]) EventRegistrationToken
	Remove_Added(token EventRegistrationToken)
	Add_Updated(handler TypedEventHandler[*IPnpObjectWatcher, *IPnpObjectUpdate]) EventRegistrationToken
	Remove_Updated(token EventRegistrationToken)
	Add_Removed(handler TypedEventHandler[*IPnpObjectWatcher, *IPnpObjectUpdate]) EventRegistrationToken
	Remove_Removed(token EventRegistrationToken)
	Add_EnumerationCompleted(handler TypedEventHandler[*IPnpObjectWatcher, interface{}]) EventRegistrationToken
	Remove_EnumerationCompleted(token EventRegistrationToken)
	Add_Stopped(handler TypedEventHandler[*IPnpObjectWatcher, interface{}]) EventRegistrationToken
	Remove_Stopped(token EventRegistrationToken)
	Get_Status() DeviceWatcherStatus
	Start()
	Stop()
}

type IPnpObjectWatcherVtbl struct {
	win32.IInspectableVtbl
	Add_Added                   uintptr
	Remove_Added                uintptr
	Add_Updated                 uintptr
	Remove_Updated              uintptr
	Add_Removed                 uintptr
	Remove_Removed              uintptr
	Add_EnumerationCompleted    uintptr
	Remove_EnumerationCompleted uintptr
	Add_Stopped                 uintptr
	Remove_Stopped              uintptr
	Get_Status                  uintptr
	Start                       uintptr
	Stop                        uintptr
}

type IPnpObjectWatcher struct {
	win32.IInspectable
}

func (this *IPnpObjectWatcher) Vtbl() *IPnpObjectWatcherVtbl {
	return (*IPnpObjectWatcherVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPnpObjectWatcher) Add_Added(handler TypedEventHandler[*IPnpObjectWatcher, *IPnpObject]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Added, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPnpObjectWatcher) Remove_Added(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Added, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPnpObjectWatcher) Add_Updated(handler TypedEventHandler[*IPnpObjectWatcher, *IPnpObjectUpdate]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Updated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPnpObjectWatcher) Remove_Updated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Updated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPnpObjectWatcher) Add_Removed(handler TypedEventHandler[*IPnpObjectWatcher, *IPnpObjectUpdate]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Removed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPnpObjectWatcher) Remove_Removed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Removed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPnpObjectWatcher) Add_EnumerationCompleted(handler TypedEventHandler[*IPnpObjectWatcher, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_EnumerationCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPnpObjectWatcher) Remove_EnumerationCompleted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_EnumerationCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPnpObjectWatcher) Add_Stopped(handler TypedEventHandler[*IPnpObjectWatcher, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Stopped, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPnpObjectWatcher) Remove_Stopped(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Stopped, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPnpObjectWatcher) Get_Status() DeviceWatcherStatus {
	var _result DeviceWatcherStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPnpObjectWatcher) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IPnpObjectWatcher) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// classes

type PnpObject struct {
	RtClass
	*IPnpObject
}

func NewIPnpObjectStatics() *IPnpObjectStatics {
	var p *IPnpObjectStatics
	hs := NewHStr("Windows.Devices.Enumeration.Pnp.PnpObject")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPnpObjectStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type PnpObjectCollection struct {
	RtClass
	*IVectorView[*IPnpObject]
}

type PnpObjectUpdate struct {
	RtClass
	*IPnpObjectUpdate
}

type PnpObjectWatcher struct {
	RtClass
	*IPnpObjectWatcher
}
