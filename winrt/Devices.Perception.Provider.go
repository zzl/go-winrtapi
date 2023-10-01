package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"log"
	"syscall"
	"unsafe"
)

// func types

// 74816D2A-2090-4670-8C48-EF39E7FF7C26
type PerceptionStartFaceAuthenticationHandler func(sender *IPerceptionFaceAuthenticationGroup, pResult *bool) com.Error

// 387EE6AA-89CD-481E-AADE-DD92F70B2AD7
type PerceptionStopFaceAuthenticationHandler func(sender *IPerceptionFaceAuthenticationGroup) com.Error

// interfaces

// 3AE651D6-9669-4106-9FAE-4835C1B96104
var IID_IKnownPerceptionFrameKindStatics = syscall.GUID{0x3AE651D6, 0x9669, 0x4106,
	[8]byte{0x9F, 0xAE, 0x48, 0x35, 0xC1, 0xB9, 0x61, 0x04}}

type IKnownPerceptionFrameKindStaticsInterface interface {
	win32.IInspectableInterface
	Get_Color() string
	Get_Depth() string
	Get_Infrared() string
}

type IKnownPerceptionFrameKindStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Color    uintptr
	Get_Depth    uintptr
	Get_Infrared uintptr
}

type IKnownPerceptionFrameKindStatics struct {
	win32.IInspectable
}

func (this *IKnownPerceptionFrameKindStatics) Vtbl() *IKnownPerceptionFrameKindStaticsVtbl {
	return (*IKnownPerceptionFrameKindStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownPerceptionFrameKindStatics) Get_Color() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Color, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownPerceptionFrameKindStatics) Get_Depth() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Depth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownPerceptionFrameKindStatics) Get_Infrared() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Infrared, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 172C4882-2FD9-4C4E-BA34-FDF20A73DDE5
var IID_IPerceptionControlGroup = syscall.GUID{0x172C4882, 0x2FD9, 0x4C4E,
	[8]byte{0xBA, 0x34, 0xFD, 0xF2, 0x0A, 0x73, 0xDD, 0xE5}}

type IPerceptionControlGroupInterface interface {
	win32.IInspectableInterface
	Get_FrameProviderIds() *IVectorView[string]
}

type IPerceptionControlGroupVtbl struct {
	win32.IInspectableVtbl
	Get_FrameProviderIds uintptr
}

type IPerceptionControlGroup struct {
	win32.IInspectable
}

func (this *IPerceptionControlGroup) Vtbl() *IPerceptionControlGroupVtbl {
	return (*IPerceptionControlGroupVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionControlGroup) Get_FrameProviderIds() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameProviderIds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2F1AF2E0-BAF1-453B-BED4-CD9D4619154C
var IID_IPerceptionControlGroupFactory = syscall.GUID{0x2F1AF2E0, 0xBAF1, 0x453B,
	[8]byte{0xBE, 0xD4, 0xCD, 0x9D, 0x46, 0x19, 0x15, 0x4C}}

type IPerceptionControlGroupFactoryInterface interface {
	win32.IInspectableInterface
	Create(ids *IIterable[string]) *IPerceptionControlGroup
}

type IPerceptionControlGroupFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IPerceptionControlGroupFactory struct {
	win32.IInspectable
}

func (this *IPerceptionControlGroupFactory) Vtbl() *IPerceptionControlGroupFactoryVtbl {
	return (*IPerceptionControlGroupFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionControlGroupFactory) Create(ids *IIterable[string]) *IPerceptionControlGroup {
	var _result *IPerceptionControlGroup
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ids)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B4131A82-DFF5-4047-8A19-3B4D805F7176
var IID_IPerceptionCorrelation = syscall.GUID{0xB4131A82, 0xDFF5, 0x4047,
	[8]byte{0x8A, 0x19, 0x3B, 0x4D, 0x80, 0x5F, 0x71, 0x76}}

type IPerceptionCorrelationInterface interface {
	win32.IInspectableInterface
	Get_TargetId() string
	Get_Position() Vector3
	Get_Orientation() Quaternion
}

type IPerceptionCorrelationVtbl struct {
	win32.IInspectableVtbl
	Get_TargetId    uintptr
	Get_Position    uintptr
	Get_Orientation uintptr
}

type IPerceptionCorrelation struct {
	win32.IInspectable
}

func (this *IPerceptionCorrelation) Vtbl() *IPerceptionCorrelationVtbl {
	return (*IPerceptionCorrelationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionCorrelation) Get_TargetId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TargetId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPerceptionCorrelation) Get_Position() Vector3 {
	var _result Vector3
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionCorrelation) Get_Orientation() Quaternion {
	var _result Quaternion
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Orientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D4A6C425-2884-4A8F-8134-2835D7286CBF
var IID_IPerceptionCorrelationFactory = syscall.GUID{0xD4A6C425, 0x2884, 0x4A8F,
	[8]byte{0x81, 0x34, 0x28, 0x35, 0xD7, 0x28, 0x6C, 0xBF}}

type IPerceptionCorrelationFactoryInterface interface {
	win32.IInspectableInterface
	Create(targetId string, position Vector3, orientation Quaternion) *IPerceptionCorrelation
}

type IPerceptionCorrelationFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IPerceptionCorrelationFactory struct {
	win32.IInspectable
}

func (this *IPerceptionCorrelationFactory) Vtbl() *IPerceptionCorrelationFactoryVtbl {
	return (*IPerceptionCorrelationFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionCorrelationFactory) Create(targetId string, position Vector3, orientation Quaternion) *IPerceptionCorrelation {
	var _result *IPerceptionCorrelation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(targetId).Ptr, uintptr(unsafe.Pointer(&position)), uintptr(unsafe.Pointer(&orientation)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 752A0906-36A7-47BB-9B79-56CC6B746770
var IID_IPerceptionCorrelationGroup = syscall.GUID{0x752A0906, 0x36A7, 0x47BB,
	[8]byte{0x9B, 0x79, 0x56, 0xCC, 0x6B, 0x74, 0x67, 0x70}}

type IPerceptionCorrelationGroupInterface interface {
	win32.IInspectableInterface
	Get_RelativeLocations() *IVectorView[*IPerceptionCorrelation]
}

type IPerceptionCorrelationGroupVtbl struct {
	win32.IInspectableVtbl
	Get_RelativeLocations uintptr
}

type IPerceptionCorrelationGroup struct {
	win32.IInspectable
}

func (this *IPerceptionCorrelationGroup) Vtbl() *IPerceptionCorrelationGroupVtbl {
	return (*IPerceptionCorrelationGroupVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionCorrelationGroup) Get_RelativeLocations() *IVectorView[*IPerceptionCorrelation] {
	var _result *IVectorView[*IPerceptionCorrelation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RelativeLocations, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7DFE2088-63DF-48ED-83B1-4AB829132995
var IID_IPerceptionCorrelationGroupFactory = syscall.GUID{0x7DFE2088, 0x63DF, 0x48ED,
	[8]byte{0x83, 0xB1, 0x4A, 0xB8, 0x29, 0x13, 0x29, 0x95}}

type IPerceptionCorrelationGroupFactoryInterface interface {
	win32.IInspectableInterface
	Create(relativeLocations *IIterable[*IPerceptionCorrelation]) *IPerceptionCorrelationGroup
}

type IPerceptionCorrelationGroupFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IPerceptionCorrelationGroupFactory struct {
	win32.IInspectable
}

func (this *IPerceptionCorrelationGroupFactory) Vtbl() *IPerceptionCorrelationGroupFactoryVtbl {
	return (*IPerceptionCorrelationGroupFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionCorrelationGroupFactory) Create(relativeLocations *IIterable[*IPerceptionCorrelation]) *IPerceptionCorrelationGroup {
	var _result *IPerceptionCorrelationGroup
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(relativeLocations)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E8019814-4A91-41B0-83A6-881A1775353E
var IID_IPerceptionFaceAuthenticationGroup = syscall.GUID{0xE8019814, 0x4A91, 0x41B0,
	[8]byte{0x83, 0xA6, 0x88, 0x1A, 0x17, 0x75, 0x35, 0x3E}}

type IPerceptionFaceAuthenticationGroupInterface interface {
	win32.IInspectableInterface
	Get_FrameProviderIds() *IVectorView[string]
}

type IPerceptionFaceAuthenticationGroupVtbl struct {
	win32.IInspectableVtbl
	Get_FrameProviderIds uintptr
}

type IPerceptionFaceAuthenticationGroup struct {
	win32.IInspectable
}

func (this *IPerceptionFaceAuthenticationGroup) Vtbl() *IPerceptionFaceAuthenticationGroupVtbl {
	return (*IPerceptionFaceAuthenticationGroupVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionFaceAuthenticationGroup) Get_FrameProviderIds() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameProviderIds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E68A05D4-B60C-40F4-BCB9-F24D46467320
var IID_IPerceptionFaceAuthenticationGroupFactory = syscall.GUID{0xE68A05D4, 0xB60C, 0x40F4,
	[8]byte{0xBC, 0xB9, 0xF2, 0x4D, 0x46, 0x46, 0x73, 0x20}}

type IPerceptionFaceAuthenticationGroupFactoryInterface interface {
	win32.IInspectableInterface
	Create(ids *IIterable[string], startHandler PerceptionStartFaceAuthenticationHandler, stopHandler PerceptionStopFaceAuthenticationHandler) *IPerceptionFaceAuthenticationGroup
}

type IPerceptionFaceAuthenticationGroupFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IPerceptionFaceAuthenticationGroupFactory struct {
	win32.IInspectable
}

func (this *IPerceptionFaceAuthenticationGroupFactory) Vtbl() *IPerceptionFaceAuthenticationGroupFactoryVtbl {
	return (*IPerceptionFaceAuthenticationGroupFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionFaceAuthenticationGroupFactory) Create(ids *IIterable[string], startHandler PerceptionStartFaceAuthenticationHandler, stopHandler PerceptionStopFaceAuthenticationHandler) *IPerceptionFaceAuthenticationGroup {
	var _result *IPerceptionFaceAuthenticationGroup
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ids)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(startHandler))), uintptr(unsafe.Pointer(NewOneArgFuncDelegate(stopHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7CFE7825-54BB-4D9D-BEC5-8EF66151D2AC
var IID_IPerceptionFrame = syscall.GUID{0x7CFE7825, 0x54BB, 0x4D9D,
	[8]byte{0xBE, 0xC5, 0x8E, 0xF6, 0x61, 0x51, 0xD2, 0xAC}}

type IPerceptionFrameInterface interface {
	win32.IInspectableInterface
	Get_RelativeTime() TimeSpan
	Put_RelativeTime(value TimeSpan)
	Get_Properties() *IPropertySet
	Get_FrameData() *IMemoryBuffer
}

type IPerceptionFrameVtbl struct {
	win32.IInspectableVtbl
	Get_RelativeTime uintptr
	Put_RelativeTime uintptr
	Get_Properties   uintptr
	Get_FrameData    uintptr
}

type IPerceptionFrame struct {
	win32.IInspectable
}

func (this *IPerceptionFrame) Vtbl() *IPerceptionFrameVtbl {
	return (*IPerceptionFrameVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionFrame) Get_RelativeTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RelativeTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionFrame) Put_RelativeTime(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RelativeTime, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IPerceptionFrame) Get_Properties() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionFrame) Get_FrameData() *IMemoryBuffer {
	var _result *IMemoryBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 794F7AB9-B37D-3B33-A10D-30626419CE65
var IID_IPerceptionFrameProvider = syscall.GUID{0x794F7AB9, 0xB37D, 0x3B33,
	[8]byte{0xA1, 0x0D, 0x30, 0x62, 0x64, 0x19, 0xCE, 0x65}}

type IPerceptionFrameProviderInterface interface {
	win32.IInspectableInterface
	Get_FrameProviderInfo() *IPerceptionFrameProviderInfo
	Get_Available() bool
	Get_Properties() *IPropertySet
	Start()
	Stop()
	SetProperty(value *IPerceptionPropertyChangeRequest)
}

type IPerceptionFrameProviderVtbl struct {
	win32.IInspectableVtbl
	Get_FrameProviderInfo uintptr
	Get_Available         uintptr
	Get_Properties        uintptr
	Start                 uintptr
	Stop                  uintptr
	SetProperty           uintptr
}

type IPerceptionFrameProvider struct {
	win32.IInspectable
}

func (this *IPerceptionFrameProvider) Vtbl() *IPerceptionFrameProviderVtbl {
	return (*IPerceptionFrameProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionFrameProvider) Get_FrameProviderInfo() *IPerceptionFrameProviderInfo {
	var _result *IPerceptionFrameProviderInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameProviderInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionFrameProvider) Get_Available() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Available, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionFrameProvider) Get_Properties() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionFrameProvider) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IPerceptionFrameProvider) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IPerceptionFrameProvider) SetProperty(value *IPerceptionPropertyChangeRequest) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetProperty, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// CCA959E8-797E-4E83-9B87-036A74142FC4
var IID_IPerceptionFrameProviderInfo = syscall.GUID{0xCCA959E8, 0x797E, 0x4E83,
	[8]byte{0x9B, 0x87, 0x03, 0x6A, 0x74, 0x14, 0x2F, 0xC4}}

type IPerceptionFrameProviderInfoInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Put_Id(value string)
	Get_DisplayName() string
	Put_DisplayName(value string)
	Get_DeviceKind() string
	Put_DeviceKind(value string)
	Get_FrameKind() string
	Put_FrameKind(value string)
	Get_Hidden() bool
	Put_Hidden(value bool)
}

type IPerceptionFrameProviderInfoVtbl struct {
	win32.IInspectableVtbl
	Get_Id          uintptr
	Put_Id          uintptr
	Get_DisplayName uintptr
	Put_DisplayName uintptr
	Get_DeviceKind  uintptr
	Put_DeviceKind  uintptr
	Get_FrameKind   uintptr
	Put_FrameKind   uintptr
	Get_Hidden      uintptr
	Put_Hidden      uintptr
}

type IPerceptionFrameProviderInfo struct {
	win32.IInspectable
}

func (this *IPerceptionFrameProviderInfo) Vtbl() *IPerceptionFrameProviderInfoVtbl {
	return (*IPerceptionFrameProviderInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionFrameProviderInfo) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPerceptionFrameProviderInfo) Put_Id(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Id, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPerceptionFrameProviderInfo) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPerceptionFrameProviderInfo) Put_DisplayName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DisplayName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPerceptionFrameProviderInfo) Get_DeviceKind() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPerceptionFrameProviderInfo) Put_DeviceKind(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DeviceKind, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPerceptionFrameProviderInfo) Get_FrameKind() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPerceptionFrameProviderInfo) Put_FrameKind(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FrameKind, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPerceptionFrameProviderInfo) Get_Hidden() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Hidden, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionFrameProviderInfo) Put_Hidden(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Hidden, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// A959CE07-EAD3-33DF-8EC1-B924ABE019C4
var IID_IPerceptionFrameProviderManager = syscall.GUID{0xA959CE07, 0xEAD3, 0x33DF,
	[8]byte{0x8E, 0xC1, 0xB9, 0x24, 0xAB, 0xE0, 0x19, 0xC4}}

type IPerceptionFrameProviderManagerInterface interface {
	win32.IInspectableInterface
	GetFrameProvider(frameProviderInfo *IPerceptionFrameProviderInfo) *IPerceptionFrameProvider
}

type IPerceptionFrameProviderManagerVtbl struct {
	win32.IInspectableVtbl
	GetFrameProvider uintptr
}

type IPerceptionFrameProviderManager struct {
	win32.IInspectable
}

func (this *IPerceptionFrameProviderManager) Vtbl() *IPerceptionFrameProviderManagerVtbl {
	return (*IPerceptionFrameProviderManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionFrameProviderManager) GetFrameProvider(frameProviderInfo *IPerceptionFrameProviderInfo) *IPerceptionFrameProvider {
	var _result *IPerceptionFrameProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFrameProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(frameProviderInfo)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AE8386E6-CAD9-4359-8F96-8EAE51810526
var IID_IPerceptionFrameProviderManagerServiceStatics = syscall.GUID{0xAE8386E6, 0xCAD9, 0x4359,
	[8]byte{0x8F, 0x96, 0x8E, 0xAE, 0x51, 0x81, 0x05, 0x26}}

type IPerceptionFrameProviderManagerServiceStaticsInterface interface {
	win32.IInspectableInterface
	RegisterFrameProviderInfo(manager *IPerceptionFrameProviderManager, frameProviderInfo *IPerceptionFrameProviderInfo)
	UnregisterFrameProviderInfo(manager *IPerceptionFrameProviderManager, frameProviderInfo *IPerceptionFrameProviderInfo)
	RegisterFaceAuthenticationGroup(manager *IPerceptionFrameProviderManager, faceAuthenticationGroup *IPerceptionFaceAuthenticationGroup)
	UnregisterFaceAuthenticationGroup(manager *IPerceptionFrameProviderManager, faceAuthenticationGroup *IPerceptionFaceAuthenticationGroup)
	RegisterControlGroup(manager *IPerceptionFrameProviderManager, controlGroup *IPerceptionControlGroup)
	UnregisterControlGroup(manager *IPerceptionFrameProviderManager, controlGroup *IPerceptionControlGroup)
	RegisterCorrelationGroup(manager *IPerceptionFrameProviderManager, correlationGroup *IPerceptionCorrelationGroup)
	UnregisterCorrelationGroup(manager *IPerceptionFrameProviderManager, correlationGroup *IPerceptionCorrelationGroup)
	UpdateAvailabilityForProvider(provider *IPerceptionFrameProvider, available bool)
	PublishFrameForProvider(provider *IPerceptionFrameProvider, frame *IPerceptionFrame)
}

type IPerceptionFrameProviderManagerServiceStaticsVtbl struct {
	win32.IInspectableVtbl
	RegisterFrameProviderInfo         uintptr
	UnregisterFrameProviderInfo       uintptr
	RegisterFaceAuthenticationGroup   uintptr
	UnregisterFaceAuthenticationGroup uintptr
	RegisterControlGroup              uintptr
	UnregisterControlGroup            uintptr
	RegisterCorrelationGroup          uintptr
	UnregisterCorrelationGroup        uintptr
	UpdateAvailabilityForProvider     uintptr
	PublishFrameForProvider           uintptr
}

type IPerceptionFrameProviderManagerServiceStatics struct {
	win32.IInspectable
}

func (this *IPerceptionFrameProviderManagerServiceStatics) Vtbl() *IPerceptionFrameProviderManagerServiceStaticsVtbl {
	return (*IPerceptionFrameProviderManagerServiceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionFrameProviderManagerServiceStatics) RegisterFrameProviderInfo(manager *IPerceptionFrameProviderManager, frameProviderInfo *IPerceptionFrameProviderInfo) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterFrameProviderInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(manager)), uintptr(unsafe.Pointer(frameProviderInfo)))
	_ = _hr
}

func (this *IPerceptionFrameProviderManagerServiceStatics) UnregisterFrameProviderInfo(manager *IPerceptionFrameProviderManager, frameProviderInfo *IPerceptionFrameProviderInfo) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UnregisterFrameProviderInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(manager)), uintptr(unsafe.Pointer(frameProviderInfo)))
	_ = _hr
}

func (this *IPerceptionFrameProviderManagerServiceStatics) RegisterFaceAuthenticationGroup(manager *IPerceptionFrameProviderManager, faceAuthenticationGroup *IPerceptionFaceAuthenticationGroup) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterFaceAuthenticationGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(manager)), uintptr(unsafe.Pointer(faceAuthenticationGroup)))
	_ = _hr
}

func (this *IPerceptionFrameProviderManagerServiceStatics) UnregisterFaceAuthenticationGroup(manager *IPerceptionFrameProviderManager, faceAuthenticationGroup *IPerceptionFaceAuthenticationGroup) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UnregisterFaceAuthenticationGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(manager)), uintptr(unsafe.Pointer(faceAuthenticationGroup)))
	_ = _hr
}

func (this *IPerceptionFrameProviderManagerServiceStatics) RegisterControlGroup(manager *IPerceptionFrameProviderManager, controlGroup *IPerceptionControlGroup) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterControlGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(manager)), uintptr(unsafe.Pointer(controlGroup)))
	_ = _hr
}

func (this *IPerceptionFrameProviderManagerServiceStatics) UnregisterControlGroup(manager *IPerceptionFrameProviderManager, controlGroup *IPerceptionControlGroup) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UnregisterControlGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(manager)), uintptr(unsafe.Pointer(controlGroup)))
	_ = _hr
}

func (this *IPerceptionFrameProviderManagerServiceStatics) RegisterCorrelationGroup(manager *IPerceptionFrameProviderManager, correlationGroup *IPerceptionCorrelationGroup) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterCorrelationGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(manager)), uintptr(unsafe.Pointer(correlationGroup)))
	_ = _hr
}

func (this *IPerceptionFrameProviderManagerServiceStatics) UnregisterCorrelationGroup(manager *IPerceptionFrameProviderManager, correlationGroup *IPerceptionCorrelationGroup) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UnregisterCorrelationGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(manager)), uintptr(unsafe.Pointer(correlationGroup)))
	_ = _hr
}

func (this *IPerceptionFrameProviderManagerServiceStatics) UpdateAvailabilityForProvider(provider *IPerceptionFrameProvider, available bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UpdateAvailabilityForProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(provider)), uintptr(*(*byte)(unsafe.Pointer(&available))))
	_ = _hr
}

func (this *IPerceptionFrameProviderManagerServiceStatics) PublishFrameForProvider(provider *IPerceptionFrameProvider, frame *IPerceptionFrame) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PublishFrameForProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(provider)), uintptr(unsafe.Pointer(frame)))
	_ = _hr
}

// 3C5AEB51-350B-4DF8-9414-59E09815510B
var IID_IPerceptionPropertyChangeRequest = syscall.GUID{0x3C5AEB51, 0x350B, 0x4DF8,
	[8]byte{0x94, 0x14, 0x59, 0xE0, 0x98, 0x15, 0x51, 0x0B}}

type IPerceptionPropertyChangeRequestInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Get_Value() interface{}
	Get_Status() PerceptionFrameSourcePropertyChangeStatus
	Put_Status(value PerceptionFrameSourcePropertyChangeStatus)
	GetDeferral() *IDeferral
}

type IPerceptionPropertyChangeRequestVtbl struct {
	win32.IInspectableVtbl
	Get_Name    uintptr
	Get_Value   uintptr
	Get_Status  uintptr
	Put_Status  uintptr
	GetDeferral uintptr
}

type IPerceptionPropertyChangeRequest struct {
	win32.IInspectable
}

func (this *IPerceptionPropertyChangeRequest) Vtbl() *IPerceptionPropertyChangeRequestVtbl {
	return (*IPerceptionPropertyChangeRequestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionPropertyChangeRequest) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPerceptionPropertyChangeRequest) Get_Value() interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionPropertyChangeRequest) Get_Status() PerceptionFrameSourcePropertyChangeStatus {
	var _result PerceptionFrameSourcePropertyChangeStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionPropertyChangeRequest) Put_Status(value PerceptionFrameSourcePropertyChangeStatus) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Status, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPerceptionPropertyChangeRequest) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4C38A7DA-FDD8-4ED4-A039-2A6F9B235038
var IID_IPerceptionVideoFrameAllocator = syscall.GUID{0x4C38A7DA, 0xFDD8, 0x4ED4,
	[8]byte{0xA0, 0x39, 0x2A, 0x6F, 0x9B, 0x23, 0x50, 0x38}}

type IPerceptionVideoFrameAllocatorInterface interface {
	win32.IInspectableInterface
	AllocateFrame() *IPerceptionFrame
	CopyFromVideoFrame(frame *IVideoFrame) *IPerceptionFrame
}

type IPerceptionVideoFrameAllocatorVtbl struct {
	win32.IInspectableVtbl
	AllocateFrame      uintptr
	CopyFromVideoFrame uintptr
}

type IPerceptionVideoFrameAllocator struct {
	win32.IInspectable
}

func (this *IPerceptionVideoFrameAllocator) Vtbl() *IPerceptionVideoFrameAllocatorVtbl {
	return (*IPerceptionVideoFrameAllocatorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionVideoFrameAllocator) AllocateFrame() *IPerceptionFrame {
	var _result *IPerceptionFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AllocateFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionVideoFrameAllocator) CopyFromVideoFrame(frame *IVideoFrame) *IPerceptionFrame {
	var _result *IPerceptionFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CopyFromVideoFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(frame)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1A58B0E1-E91A-481E-B876-A89E2BBC6B33
var IID_IPerceptionVideoFrameAllocatorFactory = syscall.GUID{0x1A58B0E1, 0xE91A, 0x481E,
	[8]byte{0xB8, 0x76, 0xA8, 0x9E, 0x2B, 0xBC, 0x6B, 0x33}}

type IPerceptionVideoFrameAllocatorFactoryInterface interface {
	win32.IInspectableInterface
	Create(maxOutstandingFrameCountForWrite uint32, format BitmapPixelFormat, resolution Size, alpha BitmapAlphaMode) *IPerceptionVideoFrameAllocator
}

type IPerceptionVideoFrameAllocatorFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IPerceptionVideoFrameAllocatorFactory struct {
	win32.IInspectable
}

func (this *IPerceptionVideoFrameAllocatorFactory) Vtbl() *IPerceptionVideoFrameAllocatorFactoryVtbl {
	return (*IPerceptionVideoFrameAllocatorFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionVideoFrameAllocatorFactory) Create(maxOutstandingFrameCountForWrite uint32, format BitmapPixelFormat, resolution Size, alpha BitmapAlphaMode) *IPerceptionVideoFrameAllocator {
	var _result *IPerceptionVideoFrameAllocator
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(maxOutstandingFrameCountForWrite), uintptr(format), *(*uintptr)(unsafe.Pointer(&resolution)), uintptr(alpha), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type KnownPerceptionFrameKind struct {
	RtClass
}

func NewIKnownPerceptionFrameKindStatics() *IKnownPerceptionFrameKindStatics {
	var p *IKnownPerceptionFrameKindStatics
	hs := NewHStr("Windows.Devices.Perception.Provider.KnownPerceptionFrameKind")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKnownPerceptionFrameKindStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type PerceptionControlGroup struct {
	RtClass
	*IPerceptionControlGroup
}

func NewPerceptionControlGroup_Create(ids *IIterable[string]) *PerceptionControlGroup {
	hs := NewHStr("Windows.Devices.Perception.Provider.PerceptionControlGroup")
	var pFac *IPerceptionControlGroupFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPerceptionControlGroupFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IPerceptionControlGroup
	p = pFac.Create(ids)
	result := &PerceptionControlGroup{
		RtClass:                 RtClass{PInspect: &p.IInspectable},
		IPerceptionControlGroup: p,
	}
	com.AddToScope(result)
	return result
}

type PerceptionCorrelation struct {
	RtClass
	*IPerceptionCorrelation
}

func NewPerceptionCorrelation_Create(targetId string, position Vector3, orientation Quaternion) *PerceptionCorrelation {
	hs := NewHStr("Windows.Devices.Perception.Provider.PerceptionCorrelation")
	var pFac *IPerceptionCorrelationFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPerceptionCorrelationFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IPerceptionCorrelation
	p = pFac.Create(targetId, position, orientation)
	result := &PerceptionCorrelation{
		RtClass:                RtClass{PInspect: &p.IInspectable},
		IPerceptionCorrelation: p,
	}
	com.AddToScope(result)
	return result
}

type PerceptionCorrelationGroup struct {
	RtClass
	*IPerceptionCorrelationGroup
}

func NewPerceptionCorrelationGroup_Create(relativeLocations *IIterable[*IPerceptionCorrelation]) *PerceptionCorrelationGroup {
	hs := NewHStr("Windows.Devices.Perception.Provider.PerceptionCorrelationGroup")
	var pFac *IPerceptionCorrelationGroupFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPerceptionCorrelationGroupFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IPerceptionCorrelationGroup
	p = pFac.Create(relativeLocations)
	result := &PerceptionCorrelationGroup{
		RtClass:                     RtClass{PInspect: &p.IInspectable},
		IPerceptionCorrelationGroup: p,
	}
	com.AddToScope(result)
	return result
}

type PerceptionFaceAuthenticationGroup struct {
	RtClass
	*IPerceptionFaceAuthenticationGroup
}

func NewPerceptionFaceAuthenticationGroup_Create(ids *IIterable[string], startHandler PerceptionStartFaceAuthenticationHandler, stopHandler PerceptionStopFaceAuthenticationHandler) *PerceptionFaceAuthenticationGroup {
	hs := NewHStr("Windows.Devices.Perception.Provider.PerceptionFaceAuthenticationGroup")
	var pFac *IPerceptionFaceAuthenticationGroupFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPerceptionFaceAuthenticationGroupFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IPerceptionFaceAuthenticationGroup
	p = pFac.Create(ids, startHandler, stopHandler)
	result := &PerceptionFaceAuthenticationGroup{
		RtClass:                            RtClass{PInspect: &p.IInspectable},
		IPerceptionFaceAuthenticationGroup: p,
	}
	com.AddToScope(result)
	return result
}

type PerceptionFrame struct {
	RtClass
	*IPerceptionFrame
}

type PerceptionFrameProviderInfo struct {
	RtClass
	*IPerceptionFrameProviderInfo
}

func NewPerceptionFrameProviderInfo() *PerceptionFrameProviderInfo {
	hs := NewHStr("Windows.Devices.Perception.Provider.PerceptionFrameProviderInfo")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &PerceptionFrameProviderInfo{
		RtClass:                      RtClass{PInspect: p},
		IPerceptionFrameProviderInfo: (*IPerceptionFrameProviderInfo)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type PerceptionFrameProviderManagerService struct {
	RtClass
}

func NewIPerceptionFrameProviderManagerServiceStatics() *IPerceptionFrameProviderManagerServiceStatics {
	var p *IPerceptionFrameProviderManagerServiceStatics
	hs := NewHStr("Windows.Devices.Perception.Provider.PerceptionFrameProviderManagerService")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPerceptionFrameProviderManagerServiceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type PerceptionPropertyChangeRequest struct {
	RtClass
	*IPerceptionPropertyChangeRequest
}

type PerceptionVideoFrameAllocator struct {
	RtClass
	*IPerceptionVideoFrameAllocator
}

func NewPerceptionVideoFrameAllocator_Create(maxOutstandingFrameCountForWrite uint32, format BitmapPixelFormat, resolution Size, alpha BitmapAlphaMode) *PerceptionVideoFrameAllocator {
	hs := NewHStr("Windows.Devices.Perception.Provider.PerceptionVideoFrameAllocator")
	var pFac *IPerceptionVideoFrameAllocatorFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPerceptionVideoFrameAllocatorFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IPerceptionVideoFrameAllocator
	p = pFac.Create(maxOutstandingFrameCountForWrite, format, resolution, alpha)
	result := &PerceptionVideoFrameAllocator{
		RtClass:                        RtClass{PInspect: &p.IInspectable},
		IPerceptionVideoFrameAllocator: p,
	}
	com.AddToScope(result)
	return result
}
