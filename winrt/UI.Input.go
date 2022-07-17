package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type CrossSlidingState int32

const (
	CrossSlidingState_Started            CrossSlidingState = 0
	CrossSlidingState_Dragging           CrossSlidingState = 1
	CrossSlidingState_Selecting          CrossSlidingState = 2
	CrossSlidingState_SelectSpeedBumping CrossSlidingState = 3
	CrossSlidingState_SpeedBumping       CrossSlidingState = 4
	CrossSlidingState_Rearranging        CrossSlidingState = 5
	CrossSlidingState_Completed          CrossSlidingState = 6
)

// enum
type DraggingState int32

const (
	DraggingState_Started    DraggingState = 0
	DraggingState_Continuing DraggingState = 1
	DraggingState_Completed  DraggingState = 2
)

// enum
type EdgeGestureKind int32

const (
	EdgeGestureKind_Touch    EdgeGestureKind = 0
	EdgeGestureKind_Keyboard EdgeGestureKind = 1
	EdgeGestureKind_Mouse    EdgeGestureKind = 2
)

// enum
type GazeInputAccessStatus int32

const (
	GazeInputAccessStatus_Unspecified    GazeInputAccessStatus = 0
	GazeInputAccessStatus_Allowed        GazeInputAccessStatus = 1
	GazeInputAccessStatus_DeniedByUser   GazeInputAccessStatus = 2
	GazeInputAccessStatus_DeniedBySystem GazeInputAccessStatus = 3
)

// enum
// flags
type GestureSettings uint32

const (
	GestureSettings_None                              GestureSettings = 0
	GestureSettings_Tap                               GestureSettings = 1
	GestureSettings_DoubleTap                         GestureSettings = 2
	GestureSettings_Hold                              GestureSettings = 4
	GestureSettings_HoldWithMouse                     GestureSettings = 8
	GestureSettings_RightTap                          GestureSettings = 16
	GestureSettings_Drag                              GestureSettings = 32
	GestureSettings_ManipulationTranslateX            GestureSettings = 64
	GestureSettings_ManipulationTranslateY            GestureSettings = 128
	GestureSettings_ManipulationTranslateRailsX       GestureSettings = 256
	GestureSettings_ManipulationTranslateRailsY       GestureSettings = 512
	GestureSettings_ManipulationRotate                GestureSettings = 1024
	GestureSettings_ManipulationScale                 GestureSettings = 2048
	GestureSettings_ManipulationTranslateInertia      GestureSettings = 4096
	GestureSettings_ManipulationRotateInertia         GestureSettings = 8192
	GestureSettings_ManipulationScaleInertia          GestureSettings = 16384
	GestureSettings_CrossSlide                        GestureSettings = 32768
	GestureSettings_ManipulationMultipleFingerPanning GestureSettings = 65536
)

// enum
type HoldingState int32

const (
	HoldingState_Started   HoldingState = 0
	HoldingState_Completed HoldingState = 1
	HoldingState_Canceled  HoldingState = 2
)

// enum
type InputActivationState int32

const (
	InputActivationState_None                   InputActivationState = 0
	InputActivationState_Deactivated            InputActivationState = 1
	InputActivationState_ActivatedNotForeground InputActivationState = 2
	InputActivationState_ActivatedInForeground  InputActivationState = 3
)

// enum
type PointerUpdateKind int32

const (
	PointerUpdateKind_Other                PointerUpdateKind = 0
	PointerUpdateKind_LeftButtonPressed    PointerUpdateKind = 1
	PointerUpdateKind_LeftButtonReleased   PointerUpdateKind = 2
	PointerUpdateKind_RightButtonPressed   PointerUpdateKind = 3
	PointerUpdateKind_RightButtonReleased  PointerUpdateKind = 4
	PointerUpdateKind_MiddleButtonPressed  PointerUpdateKind = 5
	PointerUpdateKind_MiddleButtonReleased PointerUpdateKind = 6
	PointerUpdateKind_XButton1Pressed      PointerUpdateKind = 7
	PointerUpdateKind_XButton1Released     PointerUpdateKind = 8
	PointerUpdateKind_XButton2Pressed      PointerUpdateKind = 9
	PointerUpdateKind_XButton2Released     PointerUpdateKind = 10
)

// enum
type RadialControllerMenuKnownIcon int32

const (
	RadialControllerMenuKnownIcon_Scroll            RadialControllerMenuKnownIcon = 0
	RadialControllerMenuKnownIcon_Zoom              RadialControllerMenuKnownIcon = 1
	RadialControllerMenuKnownIcon_UndoRedo          RadialControllerMenuKnownIcon = 2
	RadialControllerMenuKnownIcon_Volume            RadialControllerMenuKnownIcon = 3
	RadialControllerMenuKnownIcon_NextPreviousTrack RadialControllerMenuKnownIcon = 4
	RadialControllerMenuKnownIcon_Ruler             RadialControllerMenuKnownIcon = 5
	RadialControllerMenuKnownIcon_InkColor          RadialControllerMenuKnownIcon = 6
	RadialControllerMenuKnownIcon_InkThickness      RadialControllerMenuKnownIcon = 7
	RadialControllerMenuKnownIcon_PenType           RadialControllerMenuKnownIcon = 8
)

// enum
type RadialControllerSystemMenuItemKind int32

const (
	RadialControllerSystemMenuItemKind_Scroll            RadialControllerSystemMenuItemKind = 0
	RadialControllerSystemMenuItemKind_Zoom              RadialControllerSystemMenuItemKind = 1
	RadialControllerSystemMenuItemKind_UndoRedo          RadialControllerSystemMenuItemKind = 2
	RadialControllerSystemMenuItemKind_Volume            RadialControllerSystemMenuItemKind = 3
	RadialControllerSystemMenuItemKind_NextPreviousTrack RadialControllerSystemMenuItemKind = 4
)

// structs

type CrossSlideThresholds struct {
	SelectionStart float32
	SpeedBumpStart float32
	SpeedBumpEnd   float32
	RearrangeStart float32
}

type ManipulationDelta struct {
	Translation Point
	Scale       float32
	Rotation    float32
	Expansion   float32
}

type ManipulationVelocities struct {
	Linear    Point
	Angular   float32
	Expansion float32
}

// interfaces

// 9B822734-A3C1-542A-B2F4-0E32B773FB07
var IID_IAttachableInputObject = syscall.GUID{0x9B822734, 0xA3C1, 0x542A,
	[8]byte{0xB2, 0xF4, 0x0E, 0x32, 0xB7, 0x73, 0xFB, 0x07}}

type IAttachableInputObjectInterface interface {
	win32.IInspectableInterface
}

type IAttachableInputObjectVtbl struct {
	win32.IInspectableVtbl
}

type IAttachableInputObject struct {
	win32.IInspectable
}

func (this *IAttachableInputObject) Vtbl() *IAttachableInputObjectVtbl {
	return (*IAttachableInputObjectVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// A4C54C4E-42BC-58FA-A640-EA1516F4C06B
var IID_IAttachableInputObjectFactory = syscall.GUID{0xA4C54C4E, 0x42BC, 0x58FA,
	[8]byte{0xA6, 0x40, 0xEA, 0x15, 0x16, 0xF4, 0xC0, 0x6B}}

type IAttachableInputObjectFactoryInterface interface {
	win32.IInspectableInterface
}

type IAttachableInputObjectFactoryVtbl struct {
	win32.IInspectableVtbl
}

type IAttachableInputObjectFactory struct {
	win32.IInspectable
}

func (this *IAttachableInputObjectFactory) Vtbl() *IAttachableInputObjectFactoryVtbl {
	return (*IAttachableInputObjectFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// E9374738-6F88-41D9-8720-78E08E398349
var IID_ICrossSlidingEventArgs = syscall.GUID{0xE9374738, 0x6F88, 0x41D9,
	[8]byte{0x87, 0x20, 0x78, 0xE0, 0x8E, 0x39, 0x83, 0x49}}

type ICrossSlidingEventArgsInterface interface {
	win32.IInspectableInterface
	Get_PointerDeviceType() PointerDeviceType
	Get_Position() Point
	Get_CrossSlidingState() CrossSlidingState
}

type ICrossSlidingEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_PointerDeviceType uintptr
	Get_Position          uintptr
	Get_CrossSlidingState uintptr
}

type ICrossSlidingEventArgs struct {
	win32.IInspectable
}

func (this *ICrossSlidingEventArgs) Vtbl() *ICrossSlidingEventArgsVtbl {
	return (*ICrossSlidingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICrossSlidingEventArgs) Get_PointerDeviceType() PointerDeviceType {
	var _result PointerDeviceType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PointerDeviceType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICrossSlidingEventArgs) Get_Position() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICrossSlidingEventArgs) Get_CrossSlidingState() CrossSlidingState {
	var _result CrossSlidingState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CrossSlidingState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// EEFB7D48-C070-59F3-8DAB-BCAF621D8687
var IID_ICrossSlidingEventArgs2 = syscall.GUID{0xEEFB7D48, 0xC070, 0x59F3,
	[8]byte{0x8D, 0xAB, 0xBC, 0xAF, 0x62, 0x1D, 0x86, 0x87}}

type ICrossSlidingEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_ContactCount() uint32
}

type ICrossSlidingEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_ContactCount uintptr
}

type ICrossSlidingEventArgs2 struct {
	win32.IInspectable
}

func (this *ICrossSlidingEventArgs2) Vtbl() *ICrossSlidingEventArgs2Vtbl {
	return (*ICrossSlidingEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICrossSlidingEventArgs2) Get_ContactCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContactCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1C905384-083C-4BD3-B559-179CDDEB33EC
var IID_IDraggingEventArgs = syscall.GUID{0x1C905384, 0x083C, 0x4BD3,
	[8]byte{0xB5, 0x59, 0x17, 0x9C, 0xDD, 0xEB, 0x33, 0xEC}}

type IDraggingEventArgsInterface interface {
	win32.IInspectableInterface
	Get_PointerDeviceType() PointerDeviceType
	Get_Position() Point
	Get_DraggingState() DraggingState
}

type IDraggingEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_PointerDeviceType uintptr
	Get_Position          uintptr
	Get_DraggingState     uintptr
}

type IDraggingEventArgs struct {
	win32.IInspectable
}

func (this *IDraggingEventArgs) Vtbl() *IDraggingEventArgsVtbl {
	return (*IDraggingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDraggingEventArgs) Get_PointerDeviceType() PointerDeviceType {
	var _result PointerDeviceType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PointerDeviceType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDraggingEventArgs) Get_Position() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDraggingEventArgs) Get_DraggingState() DraggingState {
	var _result DraggingState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DraggingState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 71EFDBF9-382A-55CA-B4B9-008123C1BF1A
var IID_IDraggingEventArgs2 = syscall.GUID{0x71EFDBF9, 0x382A, 0x55CA,
	[8]byte{0xB4, 0xB9, 0x00, 0x81, 0x23, 0xC1, 0xBF, 0x1A}}

type IDraggingEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_ContactCount() uint32
}

type IDraggingEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_ContactCount uintptr
}

type IDraggingEventArgs2 struct {
	win32.IInspectable
}

func (this *IDraggingEventArgs2) Vtbl() *IDraggingEventArgs2Vtbl {
	return (*IDraggingEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDraggingEventArgs2) Get_ContactCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContactCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 580D5292-2AB1-49AA-A7F0-33BD3F8DF9F1
var IID_IEdgeGesture = syscall.GUID{0x580D5292, 0x2AB1, 0x49AA,
	[8]byte{0xA7, 0xF0, 0x33, 0xBD, 0x3F, 0x8D, 0xF9, 0xF1}}

type IEdgeGestureInterface interface {
	win32.IInspectableInterface
	Add_Starting(handler TypedEventHandler[*IEdgeGesture, *IEdgeGestureEventArgs]) EventRegistrationToken
	Remove_Starting(token EventRegistrationToken)
	Add_Completed(handler TypedEventHandler[*IEdgeGesture, *IEdgeGestureEventArgs]) EventRegistrationToken
	Remove_Completed(token EventRegistrationToken)
	Add_Canceled(handler TypedEventHandler[*IEdgeGesture, *IEdgeGestureEventArgs]) EventRegistrationToken
	Remove_Canceled(token EventRegistrationToken)
}

type IEdgeGestureVtbl struct {
	win32.IInspectableVtbl
	Add_Starting     uintptr
	Remove_Starting  uintptr
	Add_Completed    uintptr
	Remove_Completed uintptr
	Add_Canceled     uintptr
	Remove_Canceled  uintptr
}

type IEdgeGesture struct {
	win32.IInspectable
}

func (this *IEdgeGesture) Vtbl() *IEdgeGestureVtbl {
	return (*IEdgeGestureVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEdgeGesture) Add_Starting(handler TypedEventHandler[*IEdgeGesture, *IEdgeGestureEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Starting, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IEdgeGesture) Remove_Starting(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Starting, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IEdgeGesture) Add_Completed(handler TypedEventHandler[*IEdgeGesture, *IEdgeGestureEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Completed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IEdgeGesture) Remove_Completed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Completed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IEdgeGesture) Add_Canceled(handler TypedEventHandler[*IEdgeGesture, *IEdgeGestureEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Canceled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IEdgeGesture) Remove_Canceled(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Canceled, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 44FA4A24-2D09-42E1-8B5E-368208796A4C
var IID_IEdgeGestureEventArgs = syscall.GUID{0x44FA4A24, 0x2D09, 0x42E1,
	[8]byte{0x8B, 0x5E, 0x36, 0x82, 0x08, 0x79, 0x6A, 0x4C}}

type IEdgeGestureEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Kind() EdgeGestureKind
}

type IEdgeGestureEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Kind uintptr
}

type IEdgeGestureEventArgs struct {
	win32.IInspectable
}

func (this *IEdgeGestureEventArgs) Vtbl() *IEdgeGestureEventArgsVtbl {
	return (*IEdgeGestureEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEdgeGestureEventArgs) Get_Kind() EdgeGestureKind {
	var _result EdgeGestureKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Kind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// BC6A8519-18EE-4043-9839-4FC584D60A14
var IID_IEdgeGestureStatics = syscall.GUID{0xBC6A8519, 0x18EE, 0x4043,
	[8]byte{0x98, 0x39, 0x4F, 0xC5, 0x84, 0xD6, 0x0A, 0x14}}

type IEdgeGestureStaticsInterface interface {
	win32.IInspectableInterface
	GetForCurrentView() *IEdgeGesture
}

type IEdgeGestureStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForCurrentView uintptr
}

type IEdgeGestureStatics struct {
	win32.IInspectable
}

func (this *IEdgeGestureStatics) Vtbl() *IEdgeGestureStaticsVtbl {
	return (*IEdgeGestureStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEdgeGestureStatics) GetForCurrentView() *IEdgeGesture {
	var _result *IEdgeGesture
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B47A37BF-3D6B-4F88-83E8-6DCB4012FFB0
var IID_IGestureRecognizer = syscall.GUID{0xB47A37BF, 0x3D6B, 0x4F88,
	[8]byte{0x83, 0xE8, 0x6D, 0xCB, 0x40, 0x12, 0xFF, 0xB0}}

type IGestureRecognizerInterface interface {
	win32.IInspectableInterface
	Get_GestureSettings() GestureSettings
	Put_GestureSettings(value GestureSettings)
	Get_IsInertial() bool
	Get_IsActive() bool
	Get_ShowGestureFeedback() bool
	Put_ShowGestureFeedback(value bool)
	Get_PivotCenter() Point
	Put_PivotCenter(value Point)
	Get_PivotRadius() float32
	Put_PivotRadius(value float32)
	Get_InertiaTranslationDeceleration() float32
	Put_InertiaTranslationDeceleration(value float32)
	Get_InertiaRotationDeceleration() float32
	Put_InertiaRotationDeceleration(value float32)
	Get_InertiaExpansionDeceleration() float32
	Put_InertiaExpansionDeceleration(value float32)
	Get_InertiaTranslationDisplacement() float32
	Put_InertiaTranslationDisplacement(value float32)
	Get_InertiaRotationAngle() float32
	Put_InertiaRotationAngle(value float32)
	Get_InertiaExpansion() float32
	Put_InertiaExpansion(value float32)
	Get_ManipulationExact() bool
	Put_ManipulationExact(value bool)
	Get_CrossSlideThresholds() CrossSlideThresholds
	Put_CrossSlideThresholds(value CrossSlideThresholds)
	Get_CrossSlideHorizontally() bool
	Put_CrossSlideHorizontally(value bool)
	Get_CrossSlideExact() bool
	Put_CrossSlideExact(value bool)
	Get_AutoProcessInertia() bool
	Put_AutoProcessInertia(value bool)
	Get_MouseWheelParameters() *IMouseWheelParameters
	CanBeDoubleTap(value *IPointerPoint) bool
	ProcessDownEvent(value *IPointerPoint)
	ProcessMoveEvents(value *IVector[*IPointerPoint])
	ProcessUpEvent(value *IPointerPoint)
	ProcessMouseWheelEvent(value *IPointerPoint, isShiftKeyDown bool, isControlKeyDown bool)
	ProcessInertia()
	CompleteGesture()
	Add_Tapped(handler TypedEventHandler[*IGestureRecognizer, *ITappedEventArgs]) EventRegistrationToken
	Remove_Tapped(token EventRegistrationToken)
	Add_RightTapped(handler TypedEventHandler[*IGestureRecognizer, *IRightTappedEventArgs]) EventRegistrationToken
	Remove_RightTapped(token EventRegistrationToken)
	Add_Holding(handler TypedEventHandler[*IGestureRecognizer, *IHoldingEventArgs]) EventRegistrationToken
	Remove_Holding(token EventRegistrationToken)
	Add_Dragging(handler TypedEventHandler[*IGestureRecognizer, *IDraggingEventArgs]) EventRegistrationToken
	Remove_Dragging(token EventRegistrationToken)
	Add_ManipulationStarted(handler TypedEventHandler[*IGestureRecognizer, *IManipulationStartedEventArgs]) EventRegistrationToken
	Remove_ManipulationStarted(token EventRegistrationToken)
	Add_ManipulationUpdated(handler TypedEventHandler[*IGestureRecognizer, *IManipulationUpdatedEventArgs]) EventRegistrationToken
	Remove_ManipulationUpdated(token EventRegistrationToken)
	Add_ManipulationInertiaStarting(handler TypedEventHandler[*IGestureRecognizer, *IManipulationInertiaStartingEventArgs]) EventRegistrationToken
	Remove_ManipulationInertiaStarting(token EventRegistrationToken)
	Add_ManipulationCompleted(handler TypedEventHandler[*IGestureRecognizer, *IManipulationCompletedEventArgs]) EventRegistrationToken
	Remove_ManipulationCompleted(token EventRegistrationToken)
	Add_CrossSliding(handler TypedEventHandler[*IGestureRecognizer, *ICrossSlidingEventArgs]) EventRegistrationToken
	Remove_CrossSliding(token EventRegistrationToken)
}

type IGestureRecognizerVtbl struct {
	win32.IInspectableVtbl
	Get_GestureSettings                uintptr
	Put_GestureSettings                uintptr
	Get_IsInertial                     uintptr
	Get_IsActive                       uintptr
	Get_ShowGestureFeedback            uintptr
	Put_ShowGestureFeedback            uintptr
	Get_PivotCenter                    uintptr
	Put_PivotCenter                    uintptr
	Get_PivotRadius                    uintptr
	Put_PivotRadius                    uintptr
	Get_InertiaTranslationDeceleration uintptr
	Put_InertiaTranslationDeceleration uintptr
	Get_InertiaRotationDeceleration    uintptr
	Put_InertiaRotationDeceleration    uintptr
	Get_InertiaExpansionDeceleration   uintptr
	Put_InertiaExpansionDeceleration   uintptr
	Get_InertiaTranslationDisplacement uintptr
	Put_InertiaTranslationDisplacement uintptr
	Get_InertiaRotationAngle           uintptr
	Put_InertiaRotationAngle           uintptr
	Get_InertiaExpansion               uintptr
	Put_InertiaExpansion               uintptr
	Get_ManipulationExact              uintptr
	Put_ManipulationExact              uintptr
	Get_CrossSlideThresholds           uintptr
	Put_CrossSlideThresholds           uintptr
	Get_CrossSlideHorizontally         uintptr
	Put_CrossSlideHorizontally         uintptr
	Get_CrossSlideExact                uintptr
	Put_CrossSlideExact                uintptr
	Get_AutoProcessInertia             uintptr
	Put_AutoProcessInertia             uintptr
	Get_MouseWheelParameters           uintptr
	CanBeDoubleTap                     uintptr
	ProcessDownEvent                   uintptr
	ProcessMoveEvents                  uintptr
	ProcessUpEvent                     uintptr
	ProcessMouseWheelEvent             uintptr
	ProcessInertia                     uintptr
	CompleteGesture                    uintptr
	Add_Tapped                         uintptr
	Remove_Tapped                      uintptr
	Add_RightTapped                    uintptr
	Remove_RightTapped                 uintptr
	Add_Holding                        uintptr
	Remove_Holding                     uintptr
	Add_Dragging                       uintptr
	Remove_Dragging                    uintptr
	Add_ManipulationStarted            uintptr
	Remove_ManipulationStarted         uintptr
	Add_ManipulationUpdated            uintptr
	Remove_ManipulationUpdated         uintptr
	Add_ManipulationInertiaStarting    uintptr
	Remove_ManipulationInertiaStarting uintptr
	Add_ManipulationCompleted          uintptr
	Remove_ManipulationCompleted       uintptr
	Add_CrossSliding                   uintptr
	Remove_CrossSliding                uintptr
}

type IGestureRecognizer struct {
	win32.IInspectable
}

func (this *IGestureRecognizer) Vtbl() *IGestureRecognizerVtbl {
	return (*IGestureRecognizerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGestureRecognizer) Get_GestureSettings() GestureSettings {
	var _result GestureSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GestureSettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Put_GestureSettings(value GestureSettings) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_GestureSettings, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGestureRecognizer) Get_IsInertial() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsInertial, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Get_IsActive() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsActive, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Get_ShowGestureFeedback() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShowGestureFeedback, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Put_ShowGestureFeedback(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ShowGestureFeedback, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IGestureRecognizer) Get_PivotCenter() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PivotCenter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Put_PivotCenter(value Point) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PivotCenter, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IGestureRecognizer) Get_PivotRadius() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PivotRadius, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Put_PivotRadius(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PivotRadius, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGestureRecognizer) Get_InertiaTranslationDeceleration() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InertiaTranslationDeceleration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Put_InertiaTranslationDeceleration(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InertiaTranslationDeceleration, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGestureRecognizer) Get_InertiaRotationDeceleration() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InertiaRotationDeceleration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Put_InertiaRotationDeceleration(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InertiaRotationDeceleration, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGestureRecognizer) Get_InertiaExpansionDeceleration() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InertiaExpansionDeceleration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Put_InertiaExpansionDeceleration(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InertiaExpansionDeceleration, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGestureRecognizer) Get_InertiaTranslationDisplacement() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InertiaTranslationDisplacement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Put_InertiaTranslationDisplacement(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InertiaTranslationDisplacement, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGestureRecognizer) Get_InertiaRotationAngle() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InertiaRotationAngle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Put_InertiaRotationAngle(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InertiaRotationAngle, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGestureRecognizer) Get_InertiaExpansion() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InertiaExpansion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Put_InertiaExpansion(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InertiaExpansion, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGestureRecognizer) Get_ManipulationExact() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ManipulationExact, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Put_ManipulationExact(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ManipulationExact, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IGestureRecognizer) Get_CrossSlideThresholds() CrossSlideThresholds {
	var _result CrossSlideThresholds
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CrossSlideThresholds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Put_CrossSlideThresholds(value CrossSlideThresholds) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CrossSlideThresholds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IGestureRecognizer) Get_CrossSlideHorizontally() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CrossSlideHorizontally, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Put_CrossSlideHorizontally(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CrossSlideHorizontally, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IGestureRecognizer) Get_CrossSlideExact() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CrossSlideExact, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Put_CrossSlideExact(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CrossSlideExact, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IGestureRecognizer) Get_AutoProcessInertia() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutoProcessInertia, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Put_AutoProcessInertia(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AutoProcessInertia, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IGestureRecognizer) Get_MouseWheelParameters() *IMouseWheelParameters {
	var _result *IMouseWheelParameters
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MouseWheelParameters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGestureRecognizer) CanBeDoubleTap(value *IPointerPoint) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CanBeDoubleTap, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) ProcessDownEvent(value *IPointerPoint) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ProcessDownEvent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IGestureRecognizer) ProcessMoveEvents(value *IVector[*IPointerPoint]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ProcessMoveEvents, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IGestureRecognizer) ProcessUpEvent(value *IPointerPoint) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ProcessUpEvent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IGestureRecognizer) ProcessMouseWheelEvent(value *IPointerPoint, isShiftKeyDown bool, isControlKeyDown bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ProcessMouseWheelEvent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)), uintptr(*(*byte)(unsafe.Pointer(&isShiftKeyDown))), uintptr(*(*byte)(unsafe.Pointer(&isControlKeyDown))))
	_ = _hr
}

func (this *IGestureRecognizer) ProcessInertia() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ProcessInertia, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IGestureRecognizer) CompleteGesture() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CompleteGesture, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IGestureRecognizer) Add_Tapped(handler TypedEventHandler[*IGestureRecognizer, *ITappedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Tapped, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Remove_Tapped(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Tapped, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGestureRecognizer) Add_RightTapped(handler TypedEventHandler[*IGestureRecognizer, *IRightTappedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_RightTapped, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Remove_RightTapped(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_RightTapped, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGestureRecognizer) Add_Holding(handler TypedEventHandler[*IGestureRecognizer, *IHoldingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Holding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Remove_Holding(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Holding, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGestureRecognizer) Add_Dragging(handler TypedEventHandler[*IGestureRecognizer, *IDraggingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Dragging, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Remove_Dragging(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Dragging, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGestureRecognizer) Add_ManipulationStarted(handler TypedEventHandler[*IGestureRecognizer, *IManipulationStartedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ManipulationStarted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Remove_ManipulationStarted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ManipulationStarted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGestureRecognizer) Add_ManipulationUpdated(handler TypedEventHandler[*IGestureRecognizer, *IManipulationUpdatedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ManipulationUpdated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Remove_ManipulationUpdated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ManipulationUpdated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGestureRecognizer) Add_ManipulationInertiaStarting(handler TypedEventHandler[*IGestureRecognizer, *IManipulationInertiaStartingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ManipulationInertiaStarting, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Remove_ManipulationInertiaStarting(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ManipulationInertiaStarting, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGestureRecognizer) Add_ManipulationCompleted(handler TypedEventHandler[*IGestureRecognizer, *IManipulationCompletedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ManipulationCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Remove_ManipulationCompleted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ManipulationCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGestureRecognizer) Add_CrossSliding(handler TypedEventHandler[*IGestureRecognizer, *ICrossSlidingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_CrossSliding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer) Remove_CrossSliding(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_CrossSliding, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// D646097F-6EF7-5746-8BA8-8FF2206E6F3B
var IID_IGestureRecognizer2 = syscall.GUID{0xD646097F, 0x6EF7, 0x5746,
	[8]byte{0x8B, 0xA8, 0x8F, 0xF2, 0x20, 0x6E, 0x6F, 0x3B}}

type IGestureRecognizer2Interface interface {
	win32.IInspectableInterface
	Get_TapMinContactCount() uint32
	Put_TapMinContactCount(value uint32)
	Get_TapMaxContactCount() uint32
	Put_TapMaxContactCount(value uint32)
	Get_HoldMinContactCount() uint32
	Put_HoldMinContactCount(value uint32)
	Get_HoldMaxContactCount() uint32
	Put_HoldMaxContactCount(value uint32)
	Get_HoldRadius() float32
	Put_HoldRadius(value float32)
	Get_HoldStartDelay() TimeSpan
	Put_HoldStartDelay(value TimeSpan)
	Get_TranslationMinContactCount() uint32
	Put_TranslationMinContactCount(value uint32)
	Get_TranslationMaxContactCount() uint32
	Put_TranslationMaxContactCount(value uint32)
}

type IGestureRecognizer2Vtbl struct {
	win32.IInspectableVtbl
	Get_TapMinContactCount         uintptr
	Put_TapMinContactCount         uintptr
	Get_TapMaxContactCount         uintptr
	Put_TapMaxContactCount         uintptr
	Get_HoldMinContactCount        uintptr
	Put_HoldMinContactCount        uintptr
	Get_HoldMaxContactCount        uintptr
	Put_HoldMaxContactCount        uintptr
	Get_HoldRadius                 uintptr
	Put_HoldRadius                 uintptr
	Get_HoldStartDelay             uintptr
	Put_HoldStartDelay             uintptr
	Get_TranslationMinContactCount uintptr
	Put_TranslationMinContactCount uintptr
	Get_TranslationMaxContactCount uintptr
	Put_TranslationMaxContactCount uintptr
}

type IGestureRecognizer2 struct {
	win32.IInspectable
}

func (this *IGestureRecognizer2) Vtbl() *IGestureRecognizer2Vtbl {
	return (*IGestureRecognizer2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGestureRecognizer2) Get_TapMinContactCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TapMinContactCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer2) Put_TapMinContactCount(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TapMinContactCount, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGestureRecognizer2) Get_TapMaxContactCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TapMaxContactCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer2) Put_TapMaxContactCount(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TapMaxContactCount, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGestureRecognizer2) Get_HoldMinContactCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HoldMinContactCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer2) Put_HoldMinContactCount(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_HoldMinContactCount, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGestureRecognizer2) Get_HoldMaxContactCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HoldMaxContactCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer2) Put_HoldMaxContactCount(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_HoldMaxContactCount, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGestureRecognizer2) Get_HoldRadius() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HoldRadius, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer2) Put_HoldRadius(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_HoldRadius, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGestureRecognizer2) Get_HoldStartDelay() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HoldStartDelay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer2) Put_HoldStartDelay(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_HoldStartDelay, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IGestureRecognizer2) Get_TranslationMinContactCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TranslationMinContactCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer2) Put_TranslationMinContactCount(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TranslationMinContactCount, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGestureRecognizer2) Get_TranslationMaxContactCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TranslationMaxContactCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGestureRecognizer2) Put_TranslationMaxContactCount(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TranslationMaxContactCount, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 2BF755C5-E799-41B4-BB40-242F40959B71
var IID_IHoldingEventArgs = syscall.GUID{0x2BF755C5, 0xE799, 0x41B4,
	[8]byte{0xBB, 0x40, 0x24, 0x2F, 0x40, 0x95, 0x9B, 0x71}}

type IHoldingEventArgsInterface interface {
	win32.IInspectableInterface
	Get_PointerDeviceType() PointerDeviceType
	Get_Position() Point
	Get_HoldingState() HoldingState
}

type IHoldingEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_PointerDeviceType uintptr
	Get_Position          uintptr
	Get_HoldingState      uintptr
}

type IHoldingEventArgs struct {
	win32.IInspectable
}

func (this *IHoldingEventArgs) Vtbl() *IHoldingEventArgsVtbl {
	return (*IHoldingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHoldingEventArgs) Get_PointerDeviceType() PointerDeviceType {
	var _result PointerDeviceType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PointerDeviceType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHoldingEventArgs) Get_Position() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHoldingEventArgs) Get_HoldingState() HoldingState {
	var _result HoldingState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HoldingState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 141DA9EA-4C79-5674-AFEA-493FDEB91F19
var IID_IHoldingEventArgs2 = syscall.GUID{0x141DA9EA, 0x4C79, 0x5674,
	[8]byte{0xAF, 0xEA, 0x49, 0x3F, 0xDE, 0xB9, 0x1F, 0x19}}

type IHoldingEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_ContactCount() uint32
	Get_CurrentContactCount() uint32
}

type IHoldingEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_ContactCount        uintptr
	Get_CurrentContactCount uintptr
}

type IHoldingEventArgs2 struct {
	win32.IInspectable
}

func (this *IHoldingEventArgs2) Vtbl() *IHoldingEventArgs2Vtbl {
	return (*IHoldingEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHoldingEventArgs2) Get_ContactCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContactCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHoldingEventArgs2) Get_CurrentContactCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentContactCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5D6D4ED2-28C7-5AE3-AA74-C918A9F243CA
var IID_IInputActivationListener = syscall.GUID{0x5D6D4ED2, 0x28C7, 0x5AE3,
	[8]byte{0xAA, 0x74, 0xC9, 0x18, 0xA9, 0xF2, 0x43, 0xCA}}

type IInputActivationListenerInterface interface {
	win32.IInspectableInterface
	Get_State() InputActivationState
	Add_InputActivationChanged(handler TypedEventHandler[*IInputActivationListener, *IInputActivationListenerActivationChangedEventArgs]) EventRegistrationToken
	Remove_InputActivationChanged(token EventRegistrationToken)
}

type IInputActivationListenerVtbl struct {
	win32.IInspectableVtbl
	Get_State                     uintptr
	Add_InputActivationChanged    uintptr
	Remove_InputActivationChanged uintptr
}

type IInputActivationListener struct {
	win32.IInspectable
}

func (this *IInputActivationListener) Vtbl() *IInputActivationListenerVtbl {
	return (*IInputActivationListenerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInputActivationListener) Get_State() InputActivationState {
	var _result InputActivationState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInputActivationListener) Add_InputActivationChanged(handler TypedEventHandler[*IInputActivationListener, *IInputActivationListenerActivationChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_InputActivationChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInputActivationListener) Remove_InputActivationChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_InputActivationChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 7699B465-1DCF-5791-B4B9-6CAFBEED2056
var IID_IInputActivationListenerActivationChangedEventArgs = syscall.GUID{0x7699B465, 0x1DCF, 0x5791,
	[8]byte{0xB4, 0xB9, 0x6C, 0xAF, 0xBE, 0xED, 0x20, 0x56}}

type IInputActivationListenerActivationChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_State() InputActivationState
}

type IInputActivationListenerActivationChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_State uintptr
}

type IInputActivationListenerActivationChangedEventArgs struct {
	win32.IInspectable
}

func (this *IInputActivationListenerActivationChangedEventArgs) Vtbl() *IInputActivationListenerActivationChangedEventArgsVtbl {
	return (*IInputActivationListenerActivationChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInputActivationListenerActivationChangedEventArgs) Get_State() InputActivationState {
	var _result InputActivationState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// B4BAF068-8F49-446C-8DB5-8C0FFE85CC9E
var IID_IKeyboardDeliveryInterceptor = syscall.GUID{0xB4BAF068, 0x8F49, 0x446C,
	[8]byte{0x8D, 0xB5, 0x8C, 0x0F, 0xFE, 0x85, 0xCC, 0x9E}}

type IKeyboardDeliveryInterceptorInterface interface {
	win32.IInspectableInterface
	Get_IsInterceptionEnabledWhenInForeground() bool
	Put_IsInterceptionEnabledWhenInForeground(value bool)
	Add_KeyDown(handler TypedEventHandler[*IKeyboardDeliveryInterceptor, unsafe.Pointer]) EventRegistrationToken
	Remove_KeyDown(token EventRegistrationToken)
	Add_KeyUp(handler TypedEventHandler[*IKeyboardDeliveryInterceptor, unsafe.Pointer]) EventRegistrationToken
	Remove_KeyUp(token EventRegistrationToken)
}

type IKeyboardDeliveryInterceptorVtbl struct {
	win32.IInspectableVtbl
	Get_IsInterceptionEnabledWhenInForeground uintptr
	Put_IsInterceptionEnabledWhenInForeground uintptr
	Add_KeyDown                               uintptr
	Remove_KeyDown                            uintptr
	Add_KeyUp                                 uintptr
	Remove_KeyUp                              uintptr
}

type IKeyboardDeliveryInterceptor struct {
	win32.IInspectable
}

func (this *IKeyboardDeliveryInterceptor) Vtbl() *IKeyboardDeliveryInterceptorVtbl {
	return (*IKeyboardDeliveryInterceptorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyboardDeliveryInterceptor) Get_IsInterceptionEnabledWhenInForeground() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsInterceptionEnabledWhenInForeground, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IKeyboardDeliveryInterceptor) Put_IsInterceptionEnabledWhenInForeground(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsInterceptionEnabledWhenInForeground, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IKeyboardDeliveryInterceptor) Add_KeyDown(handler TypedEventHandler[*IKeyboardDeliveryInterceptor, unsafe.Pointer]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_KeyDown, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IKeyboardDeliveryInterceptor) Remove_KeyDown(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_KeyDown, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IKeyboardDeliveryInterceptor) Add_KeyUp(handler TypedEventHandler[*IKeyboardDeliveryInterceptor, unsafe.Pointer]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_KeyUp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IKeyboardDeliveryInterceptor) Remove_KeyUp(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_KeyUp, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// F9F63BA2-CEBA-4755-8A7E-14C0FFECD239
var IID_IKeyboardDeliveryInterceptorStatics = syscall.GUID{0xF9F63BA2, 0xCEBA, 0x4755,
	[8]byte{0x8A, 0x7E, 0x14, 0xC0, 0xFF, 0xEC, 0xD2, 0x39}}

type IKeyboardDeliveryInterceptorStaticsInterface interface {
	win32.IInspectableInterface
	GetForCurrentView() *IKeyboardDeliveryInterceptor
}

type IKeyboardDeliveryInterceptorStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForCurrentView uintptr
}

type IKeyboardDeliveryInterceptorStatics struct {
	win32.IInspectable
}

func (this *IKeyboardDeliveryInterceptorStatics) Vtbl() *IKeyboardDeliveryInterceptorStaticsVtbl {
	return (*IKeyboardDeliveryInterceptorStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyboardDeliveryInterceptorStatics) GetForCurrentView() *IKeyboardDeliveryInterceptor {
	var _result *IKeyboardDeliveryInterceptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B34AB22B-D19B-46FF-9F38-DEC7754BB9E7
var IID_IManipulationCompletedEventArgs = syscall.GUID{0xB34AB22B, 0xD19B, 0x46FF,
	[8]byte{0x9F, 0x38, 0xDE, 0xC7, 0x75, 0x4B, 0xB9, 0xE7}}

type IManipulationCompletedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_PointerDeviceType() PointerDeviceType
	Get_Position() Point
	Get_Cumulative() ManipulationDelta
	Get_Velocities() ManipulationVelocities
}

type IManipulationCompletedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_PointerDeviceType uintptr
	Get_Position          uintptr
	Get_Cumulative        uintptr
	Get_Velocities        uintptr
}

type IManipulationCompletedEventArgs struct {
	win32.IInspectable
}

func (this *IManipulationCompletedEventArgs) Vtbl() *IManipulationCompletedEventArgsVtbl {
	return (*IManipulationCompletedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IManipulationCompletedEventArgs) Get_PointerDeviceType() PointerDeviceType {
	var _result PointerDeviceType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PointerDeviceType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IManipulationCompletedEventArgs) Get_Position() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IManipulationCompletedEventArgs) Get_Cumulative() ManipulationDelta {
	var _result ManipulationDelta
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Cumulative, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IManipulationCompletedEventArgs) Get_Velocities() ManipulationVelocities {
	var _result ManipulationVelocities
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Velocities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F0C0DCE7-30A9-5B96-886F-6560A85E4757
var IID_IManipulationCompletedEventArgs2 = syscall.GUID{0xF0C0DCE7, 0x30A9, 0x5B96,
	[8]byte{0x88, 0x6F, 0x65, 0x60, 0xA8, 0x5E, 0x47, 0x57}}

type IManipulationCompletedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_ContactCount() uint32
	Get_CurrentContactCount() uint32
}

type IManipulationCompletedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_ContactCount        uintptr
	Get_CurrentContactCount uintptr
}

type IManipulationCompletedEventArgs2 struct {
	win32.IInspectable
}

func (this *IManipulationCompletedEventArgs2) Vtbl() *IManipulationCompletedEventArgs2Vtbl {
	return (*IManipulationCompletedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IManipulationCompletedEventArgs2) Get_ContactCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContactCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IManipulationCompletedEventArgs2) Get_CurrentContactCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentContactCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// DD37A898-26BF-467A-9CE5-CCF3FB11371E
var IID_IManipulationInertiaStartingEventArgs = syscall.GUID{0xDD37A898, 0x26BF, 0x467A,
	[8]byte{0x9C, 0xE5, 0xCC, 0xF3, 0xFB, 0x11, 0x37, 0x1E}}

type IManipulationInertiaStartingEventArgsInterface interface {
	win32.IInspectableInterface
	Get_PointerDeviceType() PointerDeviceType
	Get_Position() Point
	Get_Delta() ManipulationDelta
	Get_Cumulative() ManipulationDelta
	Get_Velocities() ManipulationVelocities
}

type IManipulationInertiaStartingEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_PointerDeviceType uintptr
	Get_Position          uintptr
	Get_Delta             uintptr
	Get_Cumulative        uintptr
	Get_Velocities        uintptr
}

type IManipulationInertiaStartingEventArgs struct {
	win32.IInspectable
}

func (this *IManipulationInertiaStartingEventArgs) Vtbl() *IManipulationInertiaStartingEventArgsVtbl {
	return (*IManipulationInertiaStartingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IManipulationInertiaStartingEventArgs) Get_PointerDeviceType() PointerDeviceType {
	var _result PointerDeviceType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PointerDeviceType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IManipulationInertiaStartingEventArgs) Get_Position() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IManipulationInertiaStartingEventArgs) Get_Delta() ManipulationDelta {
	var _result ManipulationDelta
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Delta, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IManipulationInertiaStartingEventArgs) Get_Cumulative() ManipulationDelta {
	var _result ManipulationDelta
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Cumulative, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IManipulationInertiaStartingEventArgs) Get_Velocities() ManipulationVelocities {
	var _result ManipulationVelocities
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Velocities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C25409B8-F9FA-5A45-BD97-DCBBB2201860
var IID_IManipulationInertiaStartingEventArgs2 = syscall.GUID{0xC25409B8, 0xF9FA, 0x5A45,
	[8]byte{0xBD, 0x97, 0xDC, 0xBB, 0xB2, 0x20, 0x18, 0x60}}

type IManipulationInertiaStartingEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_ContactCount() uint32
}

type IManipulationInertiaStartingEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_ContactCount uintptr
}

type IManipulationInertiaStartingEventArgs2 struct {
	win32.IInspectable
}

func (this *IManipulationInertiaStartingEventArgs2) Vtbl() *IManipulationInertiaStartingEventArgs2Vtbl {
	return (*IManipulationInertiaStartingEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IManipulationInertiaStartingEventArgs2) Get_ContactCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContactCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// DDEC873E-CFCE-4932-8C1D-3C3D011A34C0
var IID_IManipulationStartedEventArgs = syscall.GUID{0xDDEC873E, 0xCFCE, 0x4932,
	[8]byte{0x8C, 0x1D, 0x3C, 0x3D, 0x01, 0x1A, 0x34, 0xC0}}

type IManipulationStartedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_PointerDeviceType() PointerDeviceType
	Get_Position() Point
	Get_Cumulative() ManipulationDelta
}

type IManipulationStartedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_PointerDeviceType uintptr
	Get_Position          uintptr
	Get_Cumulative        uintptr
}

type IManipulationStartedEventArgs struct {
	win32.IInspectable
}

func (this *IManipulationStartedEventArgs) Vtbl() *IManipulationStartedEventArgsVtbl {
	return (*IManipulationStartedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IManipulationStartedEventArgs) Get_PointerDeviceType() PointerDeviceType {
	var _result PointerDeviceType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PointerDeviceType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IManipulationStartedEventArgs) Get_Position() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IManipulationStartedEventArgs) Get_Cumulative() ManipulationDelta {
	var _result ManipulationDelta
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Cumulative, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 2DA3DB4E-E583-5055-AFAA-16FD986531A6
var IID_IManipulationStartedEventArgs2 = syscall.GUID{0x2DA3DB4E, 0xE583, 0x5055,
	[8]byte{0xAF, 0xAA, 0x16, 0xFD, 0x98, 0x65, 0x31, 0xA6}}

type IManipulationStartedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_ContactCount() uint32
}

type IManipulationStartedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_ContactCount uintptr
}

type IManipulationStartedEventArgs2 struct {
	win32.IInspectable
}

func (this *IManipulationStartedEventArgs2) Vtbl() *IManipulationStartedEventArgs2Vtbl {
	return (*IManipulationStartedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IManipulationStartedEventArgs2) Get_ContactCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContactCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// CB354CE5-ABB8-4F9F-B3CE-8181AA61AD82
var IID_IManipulationUpdatedEventArgs = syscall.GUID{0xCB354CE5, 0xABB8, 0x4F9F,
	[8]byte{0xB3, 0xCE, 0x81, 0x81, 0xAA, 0x61, 0xAD, 0x82}}

type IManipulationUpdatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_PointerDeviceType() PointerDeviceType
	Get_Position() Point
	Get_Delta() ManipulationDelta
	Get_Cumulative() ManipulationDelta
	Get_Velocities() ManipulationVelocities
}

type IManipulationUpdatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_PointerDeviceType uintptr
	Get_Position          uintptr
	Get_Delta             uintptr
	Get_Cumulative        uintptr
	Get_Velocities        uintptr
}

type IManipulationUpdatedEventArgs struct {
	win32.IInspectable
}

func (this *IManipulationUpdatedEventArgs) Vtbl() *IManipulationUpdatedEventArgsVtbl {
	return (*IManipulationUpdatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IManipulationUpdatedEventArgs) Get_PointerDeviceType() PointerDeviceType {
	var _result PointerDeviceType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PointerDeviceType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IManipulationUpdatedEventArgs) Get_Position() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IManipulationUpdatedEventArgs) Get_Delta() ManipulationDelta {
	var _result ManipulationDelta
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Delta, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IManipulationUpdatedEventArgs) Get_Cumulative() ManipulationDelta {
	var _result ManipulationDelta
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Cumulative, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IManipulationUpdatedEventArgs) Get_Velocities() ManipulationVelocities {
	var _result ManipulationVelocities
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Velocities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F3DFB96A-3306-5903-A1C5-FF9757A8689E
var IID_IManipulationUpdatedEventArgs2 = syscall.GUID{0xF3DFB96A, 0x3306, 0x5903,
	[8]byte{0xA1, 0xC5, 0xFF, 0x97, 0x57, 0xA8, 0x68, 0x9E}}

type IManipulationUpdatedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_ContactCount() uint32
	Get_CurrentContactCount() uint32
}

type IManipulationUpdatedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_ContactCount        uintptr
	Get_CurrentContactCount uintptr
}

type IManipulationUpdatedEventArgs2 struct {
	win32.IInspectable
}

func (this *IManipulationUpdatedEventArgs2) Vtbl() *IManipulationUpdatedEventArgs2Vtbl {
	return (*IManipulationUpdatedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IManipulationUpdatedEventArgs2) Get_ContactCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContactCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IManipulationUpdatedEventArgs2) Get_CurrentContactCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentContactCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// EAD0CA44-9DED-4037-8149-5E4CC2564468
var IID_IMouseWheelParameters = syscall.GUID{0xEAD0CA44, 0x9DED, 0x4037,
	[8]byte{0x81, 0x49, 0x5E, 0x4C, 0xC2, 0x56, 0x44, 0x68}}

type IMouseWheelParametersInterface interface {
	win32.IInspectableInterface
	Get_CharTranslation() Point
	Put_CharTranslation(value Point)
	Get_DeltaScale() float32
	Put_DeltaScale(value float32)
	Get_DeltaRotationAngle() float32
	Put_DeltaRotationAngle(value float32)
	Get_PageTranslation() Point
	Put_PageTranslation(value Point)
}

type IMouseWheelParametersVtbl struct {
	win32.IInspectableVtbl
	Get_CharTranslation    uintptr
	Put_CharTranslation    uintptr
	Get_DeltaScale         uintptr
	Put_DeltaScale         uintptr
	Get_DeltaRotationAngle uintptr
	Put_DeltaRotationAngle uintptr
	Get_PageTranslation    uintptr
	Put_PageTranslation    uintptr
}

type IMouseWheelParameters struct {
	win32.IInspectable
}

func (this *IMouseWheelParameters) Vtbl() *IMouseWheelParametersVtbl {
	return (*IMouseWheelParametersVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMouseWheelParameters) Get_CharTranslation() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CharTranslation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMouseWheelParameters) Put_CharTranslation(value Point) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CharTranslation, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IMouseWheelParameters) Get_DeltaScale() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeltaScale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMouseWheelParameters) Put_DeltaScale(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DeltaScale, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMouseWheelParameters) Get_DeltaRotationAngle() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeltaRotationAngle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMouseWheelParameters) Put_DeltaRotationAngle(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DeltaRotationAngle, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMouseWheelParameters) Get_PageTranslation() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PageTranslation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMouseWheelParameters) Put_PageTranslation(value Point) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PageTranslation, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

// E995317D-7296-42D9-8233-C5BE73B74A4A
var IID_IPointerPoint = syscall.GUID{0xE995317D, 0x7296, 0x42D9,
	[8]byte{0x82, 0x33, 0xC5, 0xBE, 0x73, 0xB7, 0x4A, 0x4A}}

type IPointerPointInterface interface {
	win32.IInspectableInterface
	Get_PointerDevice() *IPointerDevice
	Get_Position() Point
	Get_RawPosition() Point
	Get_PointerId() uint32
	Get_FrameId() uint32
	Get_Timestamp() uint64
	Get_IsInContact() bool
	Get_Properties() *IPointerPointProperties
}

type IPointerPointVtbl struct {
	win32.IInspectableVtbl
	Get_PointerDevice uintptr
	Get_Position      uintptr
	Get_RawPosition   uintptr
	Get_PointerId     uintptr
	Get_FrameId       uintptr
	Get_Timestamp     uintptr
	Get_IsInContact   uintptr
	Get_Properties    uintptr
}

type IPointerPoint struct {
	win32.IInspectable
}

func (this *IPointerPoint) Vtbl() *IPointerPointVtbl {
	return (*IPointerPointVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPointerPoint) Get_PointerDevice() *IPointerDevice {
	var _result *IPointerDevice
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PointerDevice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPointerPoint) Get_Position() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPoint) Get_RawPosition() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RawPosition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPoint) Get_PointerId() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PointerId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPoint) Get_FrameId() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPoint) Get_Timestamp() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPoint) Get_IsInContact() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsInContact, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPoint) Get_Properties() *IPointerPointProperties {
	var _result *IPointerPointProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C79D8A4B-C163-4EE7-803F-67CE79F9972D
var IID_IPointerPointProperties = syscall.GUID{0xC79D8A4B, 0xC163, 0x4EE7,
	[8]byte{0x80, 0x3F, 0x67, 0xCE, 0x79, 0xF9, 0x97, 0x2D}}

type IPointerPointPropertiesInterface interface {
	win32.IInspectableInterface
	Get_Pressure() float32
	Get_IsInverted() bool
	Get_IsEraser() bool
	Get_Orientation() float32
	Get_XTilt() float32
	Get_YTilt() float32
	Get_Twist() float32
	Get_ContactRect() Rect
	Get_ContactRectRaw() Rect
	Get_TouchConfidence() bool
	Get_IsLeftButtonPressed() bool
	Get_IsRightButtonPressed() bool
	Get_IsMiddleButtonPressed() bool
	Get_MouseWheelDelta() int32
	Get_IsHorizontalMouseWheel() bool
	Get_IsPrimary() bool
	Get_IsInRange() bool
	Get_IsCanceled() bool
	Get_IsBarrelButtonPressed() bool
	Get_IsXButton1Pressed() bool
	Get_IsXButton2Pressed() bool
	Get_PointerUpdateKind() PointerUpdateKind
	HasUsage(usagePage uint32, usageId uint32) bool
	GetUsageValue(usagePage uint32, usageId uint32) int32
}

type IPointerPointPropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_Pressure               uintptr
	Get_IsInverted             uintptr
	Get_IsEraser               uintptr
	Get_Orientation            uintptr
	Get_XTilt                  uintptr
	Get_YTilt                  uintptr
	Get_Twist                  uintptr
	Get_ContactRect            uintptr
	Get_ContactRectRaw         uintptr
	Get_TouchConfidence        uintptr
	Get_IsLeftButtonPressed    uintptr
	Get_IsRightButtonPressed   uintptr
	Get_IsMiddleButtonPressed  uintptr
	Get_MouseWheelDelta        uintptr
	Get_IsHorizontalMouseWheel uintptr
	Get_IsPrimary              uintptr
	Get_IsInRange              uintptr
	Get_IsCanceled             uintptr
	Get_IsBarrelButtonPressed  uintptr
	Get_IsXButton1Pressed      uintptr
	Get_IsXButton2Pressed      uintptr
	Get_PointerUpdateKind      uintptr
	HasUsage                   uintptr
	GetUsageValue              uintptr
}

type IPointerPointProperties struct {
	win32.IInspectable
}

func (this *IPointerPointProperties) Vtbl() *IPointerPointPropertiesVtbl {
	return (*IPointerPointPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPointerPointProperties) Get_Pressure() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Pressure, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPointProperties) Get_IsInverted() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsInverted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPointProperties) Get_IsEraser() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEraser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPointProperties) Get_Orientation() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Orientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPointProperties) Get_XTilt() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_XTilt, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPointProperties) Get_YTilt() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_YTilt, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPointProperties) Get_Twist() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Twist, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPointProperties) Get_ContactRect() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContactRect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPointProperties) Get_ContactRectRaw() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContactRectRaw, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPointProperties) Get_TouchConfidence() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TouchConfidence, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPointProperties) Get_IsLeftButtonPressed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsLeftButtonPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPointProperties) Get_IsRightButtonPressed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsRightButtonPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPointProperties) Get_IsMiddleButtonPressed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsMiddleButtonPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPointProperties) Get_MouseWheelDelta() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MouseWheelDelta, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPointProperties) Get_IsHorizontalMouseWheel() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsHorizontalMouseWheel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPointProperties) Get_IsPrimary() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPrimary, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPointProperties) Get_IsInRange() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsInRange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPointProperties) Get_IsCanceled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCanceled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPointProperties) Get_IsBarrelButtonPressed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsBarrelButtonPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPointProperties) Get_IsXButton1Pressed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsXButton1Pressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPointProperties) Get_IsXButton2Pressed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsXButton2Pressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPointProperties) Get_PointerUpdateKind() PointerUpdateKind {
	var _result PointerUpdateKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PointerUpdateKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPointProperties) HasUsage(usagePage uint32, usageId uint32) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().HasUsage, uintptr(unsafe.Pointer(this)), uintptr(usagePage), uintptr(usageId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPointProperties) GetUsageValue(usagePage uint32, usageId uint32) int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetUsageValue, uintptr(unsafe.Pointer(this)), uintptr(usagePage), uintptr(usageId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 22C3433A-C83B-41C0-A296-5E232D64D6AF
var IID_IPointerPointProperties2 = syscall.GUID{0x22C3433A, 0xC83B, 0x41C0,
	[8]byte{0xA2, 0x96, 0x5E, 0x23, 0x2D, 0x64, 0xD6, 0xAF}}

type IPointerPointProperties2Interface interface {
	win32.IInspectableInterface
	Get_ZDistance() *IReference[float32]
}

type IPointerPointProperties2Vtbl struct {
	win32.IInspectableVtbl
	Get_ZDistance uintptr
}

type IPointerPointProperties2 struct {
	win32.IInspectable
}

func (this *IPointerPointProperties2) Vtbl() *IPointerPointProperties2Vtbl {
	return (*IPointerPointProperties2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPointerPointProperties2) Get_ZDistance() *IReference[float32] {
	var _result *IReference[float32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ZDistance, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A506638D-2A1A-413E-BC75-9F38381CC069
var IID_IPointerPointStatics = syscall.GUID{0xA506638D, 0x2A1A, 0x413E,
	[8]byte{0xBC, 0x75, 0x9F, 0x38, 0x38, 0x1C, 0xC0, 0x69}}

type IPointerPointStaticsInterface interface {
	win32.IInspectableInterface
	GetCurrentPoint(pointerId uint32) *IPointerPoint
	GetIntermediatePoints(pointerId uint32) *IVector[*IPointerPoint]
	GetCurrentPointTransformed(pointerId uint32, transform *IPointerPointTransform) *IPointerPoint
	GetIntermediatePointsTransformed(pointerId uint32, transform *IPointerPointTransform) *IVector[*IPointerPoint]
}

type IPointerPointStaticsVtbl struct {
	win32.IInspectableVtbl
	GetCurrentPoint                  uintptr
	GetIntermediatePoints            uintptr
	GetCurrentPointTransformed       uintptr
	GetIntermediatePointsTransformed uintptr
}

type IPointerPointStatics struct {
	win32.IInspectable
}

func (this *IPointerPointStatics) Vtbl() *IPointerPointStaticsVtbl {
	return (*IPointerPointStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPointerPointStatics) GetCurrentPoint(pointerId uint32) *IPointerPoint {
	var _result *IPointerPoint
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentPoint, uintptr(unsafe.Pointer(this)), uintptr(pointerId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPointerPointStatics) GetIntermediatePoints(pointerId uint32) *IVector[*IPointerPoint] {
	var _result *IVector[*IPointerPoint]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetIntermediatePoints, uintptr(unsafe.Pointer(this)), uintptr(pointerId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPointerPointStatics) GetCurrentPointTransformed(pointerId uint32, transform *IPointerPointTransform) *IPointerPoint {
	var _result *IPointerPoint
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentPointTransformed, uintptr(unsafe.Pointer(this)), uintptr(pointerId), uintptr(unsafe.Pointer(transform)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPointerPointStatics) GetIntermediatePointsTransformed(pointerId uint32, transform *IPointerPointTransform) *IVector[*IPointerPoint] {
	var _result *IVector[*IPointerPoint]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetIntermediatePointsTransformed, uintptr(unsafe.Pointer(this)), uintptr(pointerId), uintptr(unsafe.Pointer(transform)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4D5FE14F-B87C-4028-BC9C-59E9947FB056
var IID_IPointerPointTransform = syscall.GUID{0x4D5FE14F, 0xB87C, 0x4028,
	[8]byte{0xBC, 0x9C, 0x59, 0xE9, 0x94, 0x7F, 0xB0, 0x56}}

type IPointerPointTransformInterface interface {
	win32.IInspectableInterface
	Get_Inverse() *IPointerPointTransform
	TryTransform(inPoint Point, outPoint Point) bool
	TransformBounds(rect Rect) Rect
}

type IPointerPointTransformVtbl struct {
	win32.IInspectableVtbl
	Get_Inverse     uintptr
	TryTransform    uintptr
	TransformBounds uintptr
}

type IPointerPointTransform struct {
	win32.IInspectable
}

func (this *IPointerPointTransform) Vtbl() *IPointerPointTransformVtbl {
	return (*IPointerPointTransformVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPointerPointTransform) Get_Inverse() *IPointerPointTransform {
	var _result *IPointerPointTransform
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Inverse, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPointerPointTransform) TryTransform(inPoint Point, outPoint Point) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryTransform, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&inPoint)), *(*uintptr)(unsafe.Pointer(&outPoint)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerPointTransform) TransformBounds(rect Rect) Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TransformBounds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&rect)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4D1E6461-84F7-499D-BD91-2A36E2B7AAA2
var IID_IPointerVisualizationSettings = syscall.GUID{0x4D1E6461, 0x84F7, 0x499D,
	[8]byte{0xBD, 0x91, 0x2A, 0x36, 0xE2, 0xB7, 0xAA, 0xA2}}

type IPointerVisualizationSettingsInterface interface {
	win32.IInspectableInterface
	Put_IsContactFeedbackEnabled(value bool)
	Get_IsContactFeedbackEnabled() bool
	Put_IsBarrelButtonFeedbackEnabled(value bool)
	Get_IsBarrelButtonFeedbackEnabled() bool
}

type IPointerVisualizationSettingsVtbl struct {
	win32.IInspectableVtbl
	Put_IsContactFeedbackEnabled      uintptr
	Get_IsContactFeedbackEnabled      uintptr
	Put_IsBarrelButtonFeedbackEnabled uintptr
	Get_IsBarrelButtonFeedbackEnabled uintptr
}

type IPointerVisualizationSettings struct {
	win32.IInspectable
}

func (this *IPointerVisualizationSettings) Vtbl() *IPointerVisualizationSettingsVtbl {
	return (*IPointerVisualizationSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPointerVisualizationSettings) Put_IsContactFeedbackEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsContactFeedbackEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPointerVisualizationSettings) Get_IsContactFeedbackEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsContactFeedbackEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerVisualizationSettings) Put_IsBarrelButtonFeedbackEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsBarrelButtonFeedbackEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPointerVisualizationSettings) Get_IsBarrelButtonFeedbackEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsBarrelButtonFeedbackEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 68870EDB-165B-4214-B4F3-584ECA8C8A69
var IID_IPointerVisualizationSettingsStatics = syscall.GUID{0x68870EDB, 0x165B, 0x4214,
	[8]byte{0xB4, 0xF3, 0x58, 0x4E, 0xCA, 0x8C, 0x8A, 0x69}}

type IPointerVisualizationSettingsStaticsInterface interface {
	win32.IInspectableInterface
	GetForCurrentView() *IPointerVisualizationSettings
}

type IPointerVisualizationSettingsStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForCurrentView uintptr
}

type IPointerVisualizationSettingsStatics struct {
	win32.IInspectable
}

func (this *IPointerVisualizationSettingsStatics) Vtbl() *IPointerVisualizationSettingsStaticsVtbl {
	return (*IPointerVisualizationSettingsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPointerVisualizationSettingsStatics) GetForCurrentView() *IPointerVisualizationSettings {
	var _result *IPointerVisualizationSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3055D1C8-DF51-43D4-B23B-0E1037467A09
var IID_IRadialController = syscall.GUID{0x3055D1C8, 0xDF51, 0x43D4,
	[8]byte{0xB2, 0x3B, 0x0E, 0x10, 0x37, 0x46, 0x7A, 0x09}}

type IRadialControllerInterface interface {
	win32.IInspectableInterface
	Get_Menu() *IRadialControllerMenu
	Get_RotationResolutionInDegrees() float64
	Put_RotationResolutionInDegrees(value float64)
	Get_UseAutomaticHapticFeedback() bool
	Put_UseAutomaticHapticFeedback(value bool)
	Add_ScreenContactStarted(handler TypedEventHandler[*IRadialController, *IRadialControllerScreenContactStartedEventArgs]) EventRegistrationToken
	Remove_ScreenContactStarted(cookie EventRegistrationToken)
	Add_ScreenContactEnded(handler TypedEventHandler[*IRadialController, interface{}]) EventRegistrationToken
	Remove_ScreenContactEnded(cookie EventRegistrationToken)
	Add_ScreenContactContinued(handler TypedEventHandler[*IRadialController, *IRadialControllerScreenContactContinuedEventArgs]) EventRegistrationToken
	Remove_ScreenContactContinued(cookie EventRegistrationToken)
	Add_ControlLost(handler TypedEventHandler[*IRadialController, interface{}]) EventRegistrationToken
	Remove_ControlLost(cookie EventRegistrationToken)
	Add_RotationChanged(handler TypedEventHandler[*IRadialController, *IRadialControllerRotationChangedEventArgs]) EventRegistrationToken
	Remove_RotationChanged(token EventRegistrationToken)
	Add_ButtonClicked(handler TypedEventHandler[*IRadialController, *IRadialControllerButtonClickedEventArgs]) EventRegistrationToken
	Remove_ButtonClicked(token EventRegistrationToken)
	Add_ControlAcquired(handler TypedEventHandler[*IRadialController, *IRadialControllerControlAcquiredEventArgs]) EventRegistrationToken
	Remove_ControlAcquired(cookie EventRegistrationToken)
}

type IRadialControllerVtbl struct {
	win32.IInspectableVtbl
	Get_Menu                        uintptr
	Get_RotationResolutionInDegrees uintptr
	Put_RotationResolutionInDegrees uintptr
	Get_UseAutomaticHapticFeedback  uintptr
	Put_UseAutomaticHapticFeedback  uintptr
	Add_ScreenContactStarted        uintptr
	Remove_ScreenContactStarted     uintptr
	Add_ScreenContactEnded          uintptr
	Remove_ScreenContactEnded       uintptr
	Add_ScreenContactContinued      uintptr
	Remove_ScreenContactContinued   uintptr
	Add_ControlLost                 uintptr
	Remove_ControlLost              uintptr
	Add_RotationChanged             uintptr
	Remove_RotationChanged          uintptr
	Add_ButtonClicked               uintptr
	Remove_ButtonClicked            uintptr
	Add_ControlAcquired             uintptr
	Remove_ControlAcquired          uintptr
}

type IRadialController struct {
	win32.IInspectable
}

func (this *IRadialController) Vtbl() *IRadialControllerVtbl {
	return (*IRadialControllerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialController) Get_Menu() *IRadialControllerMenu {
	var _result *IRadialControllerMenu
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Menu, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRadialController) Get_RotationResolutionInDegrees() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RotationResolutionInDegrees, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRadialController) Put_RotationResolutionInDegrees(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RotationResolutionInDegrees, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IRadialController) Get_UseAutomaticHapticFeedback() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UseAutomaticHapticFeedback, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRadialController) Put_UseAutomaticHapticFeedback(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_UseAutomaticHapticFeedback, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IRadialController) Add_ScreenContactStarted(handler TypedEventHandler[*IRadialController, *IRadialControllerScreenContactStartedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ScreenContactStarted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRadialController) Remove_ScreenContactStarted(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ScreenContactStarted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IRadialController) Add_ScreenContactEnded(handler TypedEventHandler[*IRadialController, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ScreenContactEnded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRadialController) Remove_ScreenContactEnded(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ScreenContactEnded, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IRadialController) Add_ScreenContactContinued(handler TypedEventHandler[*IRadialController, *IRadialControllerScreenContactContinuedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ScreenContactContinued, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRadialController) Remove_ScreenContactContinued(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ScreenContactContinued, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IRadialController) Add_ControlLost(handler TypedEventHandler[*IRadialController, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ControlLost, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRadialController) Remove_ControlLost(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ControlLost, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IRadialController) Add_RotationChanged(handler TypedEventHandler[*IRadialController, *IRadialControllerRotationChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_RotationChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRadialController) Remove_RotationChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_RotationChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IRadialController) Add_ButtonClicked(handler TypedEventHandler[*IRadialController, *IRadialControllerButtonClickedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ButtonClicked, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRadialController) Remove_ButtonClicked(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ButtonClicked, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IRadialController) Add_ControlAcquired(handler TypedEventHandler[*IRadialController, *IRadialControllerControlAcquiredEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ControlAcquired, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRadialController) Remove_ControlAcquired(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ControlAcquired, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

// 3D577EFF-4CEE-11E6-B535-001BDC06AB3B
var IID_IRadialController2 = syscall.GUID{0x3D577EFF, 0x4CEE, 0x11E6,
	[8]byte{0xB5, 0x35, 0x00, 0x1B, 0xDC, 0x06, 0xAB, 0x3B}}

type IRadialController2Interface interface {
	win32.IInspectableInterface
	Add_ButtonPressed(handler TypedEventHandler[*IRadialController, *IRadialControllerButtonPressedEventArgs]) EventRegistrationToken
	Remove_ButtonPressed(token EventRegistrationToken)
	Add_ButtonHolding(handler TypedEventHandler[*IRadialController, *IRadialControllerButtonHoldingEventArgs]) EventRegistrationToken
	Remove_ButtonHolding(token EventRegistrationToken)
	Add_ButtonReleased(handler TypedEventHandler[*IRadialController, *IRadialControllerButtonReleasedEventArgs]) EventRegistrationToken
	Remove_ButtonReleased(token EventRegistrationToken)
}

type IRadialController2Vtbl struct {
	win32.IInspectableVtbl
	Add_ButtonPressed     uintptr
	Remove_ButtonPressed  uintptr
	Add_ButtonHolding     uintptr
	Remove_ButtonHolding  uintptr
	Add_ButtonReleased    uintptr
	Remove_ButtonReleased uintptr
}

type IRadialController2 struct {
	win32.IInspectable
}

func (this *IRadialController2) Vtbl() *IRadialController2Vtbl {
	return (*IRadialController2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialController2) Add_ButtonPressed(handler TypedEventHandler[*IRadialController, *IRadialControllerButtonPressedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ButtonPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRadialController2) Remove_ButtonPressed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ButtonPressed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IRadialController2) Add_ButtonHolding(handler TypedEventHandler[*IRadialController, *IRadialControllerButtonHoldingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ButtonHolding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRadialController2) Remove_ButtonHolding(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ButtonHolding, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IRadialController2) Add_ButtonReleased(handler TypedEventHandler[*IRadialController, *IRadialControllerButtonReleasedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ButtonReleased, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRadialController2) Remove_ButtonReleased(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ButtonReleased, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 206AA438-E651-11E5-BF62-2C27D7404E85
var IID_IRadialControllerButtonClickedEventArgs = syscall.GUID{0x206AA438, 0xE651, 0x11E5,
	[8]byte{0xBF, 0x62, 0x2C, 0x27, 0xD7, 0x40, 0x4E, 0x85}}

type IRadialControllerButtonClickedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Contact() *IRadialControllerScreenContact
}

type IRadialControllerButtonClickedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Contact uintptr
}

type IRadialControllerButtonClickedEventArgs struct {
	win32.IInspectable
}

func (this *IRadialControllerButtonClickedEventArgs) Vtbl() *IRadialControllerButtonClickedEventArgsVtbl {
	return (*IRadialControllerButtonClickedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerButtonClickedEventArgs) Get_Contact() *IRadialControllerScreenContact {
	var _result *IRadialControllerScreenContact
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Contact, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3D577EF3-3CEE-11E6-B535-001BDC06AB3B
var IID_IRadialControllerButtonClickedEventArgs2 = syscall.GUID{0x3D577EF3, 0x3CEE, 0x11E6,
	[8]byte{0xB5, 0x35, 0x00, 0x1B, 0xDC, 0x06, 0xAB, 0x3B}}

type IRadialControllerButtonClickedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_SimpleHapticsController() *ISimpleHapticsController
}

type IRadialControllerButtonClickedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_SimpleHapticsController uintptr
}

type IRadialControllerButtonClickedEventArgs2 struct {
	win32.IInspectable
}

func (this *IRadialControllerButtonClickedEventArgs2) Vtbl() *IRadialControllerButtonClickedEventArgs2Vtbl {
	return (*IRadialControllerButtonClickedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerButtonClickedEventArgs2) Get_SimpleHapticsController() *ISimpleHapticsController {
	var _result *ISimpleHapticsController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SimpleHapticsController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3D577EEE-3CEE-11E6-B535-001BDC06AB3B
var IID_IRadialControllerButtonHoldingEventArgs = syscall.GUID{0x3D577EEE, 0x3CEE, 0x11E6,
	[8]byte{0xB5, 0x35, 0x00, 0x1B, 0xDC, 0x06, 0xAB, 0x3B}}

type IRadialControllerButtonHoldingEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Contact() *IRadialControllerScreenContact
	Get_SimpleHapticsController() *ISimpleHapticsController
}

type IRadialControllerButtonHoldingEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Contact                 uintptr
	Get_SimpleHapticsController uintptr
}

type IRadialControllerButtonHoldingEventArgs struct {
	win32.IInspectable
}

func (this *IRadialControllerButtonHoldingEventArgs) Vtbl() *IRadialControllerButtonHoldingEventArgsVtbl {
	return (*IRadialControllerButtonHoldingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerButtonHoldingEventArgs) Get_Contact() *IRadialControllerScreenContact {
	var _result *IRadialControllerScreenContact
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Contact, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRadialControllerButtonHoldingEventArgs) Get_SimpleHapticsController() *ISimpleHapticsController {
	var _result *ISimpleHapticsController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SimpleHapticsController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3D577EED-4CEE-11E6-B535-001BDC06AB3B
var IID_IRadialControllerButtonPressedEventArgs = syscall.GUID{0x3D577EED, 0x4CEE, 0x11E6,
	[8]byte{0xB5, 0x35, 0x00, 0x1B, 0xDC, 0x06, 0xAB, 0x3B}}

type IRadialControllerButtonPressedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Contact() *IRadialControllerScreenContact
	Get_SimpleHapticsController() *ISimpleHapticsController
}

type IRadialControllerButtonPressedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Contact                 uintptr
	Get_SimpleHapticsController uintptr
}

type IRadialControllerButtonPressedEventArgs struct {
	win32.IInspectable
}

func (this *IRadialControllerButtonPressedEventArgs) Vtbl() *IRadialControllerButtonPressedEventArgsVtbl {
	return (*IRadialControllerButtonPressedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerButtonPressedEventArgs) Get_Contact() *IRadialControllerScreenContact {
	var _result *IRadialControllerScreenContact
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Contact, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRadialControllerButtonPressedEventArgs) Get_SimpleHapticsController() *ISimpleHapticsController {
	var _result *ISimpleHapticsController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SimpleHapticsController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3D577EEF-3CEE-11E6-B535-001BDC06AB3B
var IID_IRadialControllerButtonReleasedEventArgs = syscall.GUID{0x3D577EEF, 0x3CEE, 0x11E6,
	[8]byte{0xB5, 0x35, 0x00, 0x1B, 0xDC, 0x06, 0xAB, 0x3B}}

type IRadialControllerButtonReleasedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Contact() *IRadialControllerScreenContact
	Get_SimpleHapticsController() *ISimpleHapticsController
}

type IRadialControllerButtonReleasedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Contact                 uintptr
	Get_SimpleHapticsController uintptr
}

type IRadialControllerButtonReleasedEventArgs struct {
	win32.IInspectable
}

func (this *IRadialControllerButtonReleasedEventArgs) Vtbl() *IRadialControllerButtonReleasedEventArgsVtbl {
	return (*IRadialControllerButtonReleasedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerButtonReleasedEventArgs) Get_Contact() *IRadialControllerScreenContact {
	var _result *IRadialControllerScreenContact
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Contact, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRadialControllerButtonReleasedEventArgs) Get_SimpleHapticsController() *ISimpleHapticsController {
	var _result *ISimpleHapticsController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SimpleHapticsController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A6B79ECB-6A52-4430-910C-56370A9D6B42
var IID_IRadialControllerConfiguration = syscall.GUID{0xA6B79ECB, 0x6A52, 0x4430,
	[8]byte{0x91, 0x0C, 0x56, 0x37, 0x0A, 0x9D, 0x6B, 0x42}}

type IRadialControllerConfigurationInterface interface {
	win32.IInspectableInterface
	SetDefaultMenuItems(buttons *IIterable[RadialControllerSystemMenuItemKind])
	ResetToDefaultMenuItems()
	TrySelectDefaultMenuItem(type_ RadialControllerSystemMenuItemKind) bool
}

type IRadialControllerConfigurationVtbl struct {
	win32.IInspectableVtbl
	SetDefaultMenuItems      uintptr
	ResetToDefaultMenuItems  uintptr
	TrySelectDefaultMenuItem uintptr
}

type IRadialControllerConfiguration struct {
	win32.IInspectable
}

func (this *IRadialControllerConfiguration) Vtbl() *IRadialControllerConfigurationVtbl {
	return (*IRadialControllerConfigurationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerConfiguration) SetDefaultMenuItems(buttons *IIterable[RadialControllerSystemMenuItemKind]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetDefaultMenuItems, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(buttons)))
	_ = _hr
}

func (this *IRadialControllerConfiguration) ResetToDefaultMenuItems() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ResetToDefaultMenuItems, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IRadialControllerConfiguration) TrySelectDefaultMenuItem(type_ RadialControllerSystemMenuItemKind) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySelectDefaultMenuItem, uintptr(unsafe.Pointer(this)), uintptr(type_), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3D577EF7-3CEE-11E6-B535-001BDC06AB3B
var IID_IRadialControllerConfiguration2 = syscall.GUID{0x3D577EF7, 0x3CEE, 0x11E6,
	[8]byte{0xB5, 0x35, 0x00, 0x1B, 0xDC, 0x06, 0xAB, 0x3B}}

type IRadialControllerConfiguration2Interface interface {
	win32.IInspectableInterface
	Put_ActiveControllerWhenMenuIsSuppressed(value *IRadialController)
	Get_ActiveControllerWhenMenuIsSuppressed() *IRadialController
	Put_IsMenuSuppressed(value bool)
	Get_IsMenuSuppressed() bool
}

type IRadialControllerConfiguration2Vtbl struct {
	win32.IInspectableVtbl
	Put_ActiveControllerWhenMenuIsSuppressed uintptr
	Get_ActiveControllerWhenMenuIsSuppressed uintptr
	Put_IsMenuSuppressed                     uintptr
	Get_IsMenuSuppressed                     uintptr
}

type IRadialControllerConfiguration2 struct {
	win32.IInspectable
}

func (this *IRadialControllerConfiguration2) Vtbl() *IRadialControllerConfiguration2Vtbl {
	return (*IRadialControllerConfiguration2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerConfiguration2) Put_ActiveControllerWhenMenuIsSuppressed(value *IRadialController) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ActiveControllerWhenMenuIsSuppressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IRadialControllerConfiguration2) Get_ActiveControllerWhenMenuIsSuppressed() *IRadialController {
	var _result *IRadialController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActiveControllerWhenMenuIsSuppressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRadialControllerConfiguration2) Put_IsMenuSuppressed(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsMenuSuppressed, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IRadialControllerConfiguration2) Get_IsMenuSuppressed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsMenuSuppressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 79B6B0E5-069A-4486-A99D-8DB772B9642F
var IID_IRadialControllerConfigurationStatics = syscall.GUID{0x79B6B0E5, 0x069A, 0x4486,
	[8]byte{0xA9, 0x9D, 0x8D, 0xB7, 0x72, 0xB9, 0x64, 0x2F}}

type IRadialControllerConfigurationStaticsInterface interface {
	win32.IInspectableInterface
	GetForCurrentView() *IRadialControllerConfiguration
}

type IRadialControllerConfigurationStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForCurrentView uintptr
}

type IRadialControllerConfigurationStatics struct {
	win32.IInspectable
}

func (this *IRadialControllerConfigurationStatics) Vtbl() *IRadialControllerConfigurationStaticsVtbl {
	return (*IRadialControllerConfigurationStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerConfigurationStatics) GetForCurrentView() *IRadialControllerConfiguration {
	var _result *IRadialControllerConfiguration
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 53E08B17-E205-48D3-9CAF-80FF47C4D7C7
var IID_IRadialControllerConfigurationStatics2 = syscall.GUID{0x53E08B17, 0xE205, 0x48D3,
	[8]byte{0x9C, 0xAF, 0x80, 0xFF, 0x47, 0xC4, 0xD7, 0xC7}}

type IRadialControllerConfigurationStatics2Interface interface {
	win32.IInspectableInterface
	Put_AppController(value *IRadialController)
	Get_AppController() *IRadialController
	Put_IsAppControllerEnabled(value bool)
	Get_IsAppControllerEnabled() bool
}

type IRadialControllerConfigurationStatics2Vtbl struct {
	win32.IInspectableVtbl
	Put_AppController          uintptr
	Get_AppController          uintptr
	Put_IsAppControllerEnabled uintptr
	Get_IsAppControllerEnabled uintptr
}

type IRadialControllerConfigurationStatics2 struct {
	win32.IInspectable
}

func (this *IRadialControllerConfigurationStatics2) Vtbl() *IRadialControllerConfigurationStatics2Vtbl {
	return (*IRadialControllerConfigurationStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerConfigurationStatics2) Put_AppController(value *IRadialController) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AppController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IRadialControllerConfigurationStatics2) Get_AppController() *IRadialController {
	var _result *IRadialController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRadialControllerConfigurationStatics2) Put_IsAppControllerEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsAppControllerEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IRadialControllerConfigurationStatics2) Get_IsAppControllerEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAppControllerEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 206AA439-E651-11E5-BF62-2C27D7404E85
var IID_IRadialControllerControlAcquiredEventArgs = syscall.GUID{0x206AA439, 0xE651, 0x11E5,
	[8]byte{0xBF, 0x62, 0x2C, 0x27, 0xD7, 0x40, 0x4E, 0x85}}

type IRadialControllerControlAcquiredEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Contact() *IRadialControllerScreenContact
}

type IRadialControllerControlAcquiredEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Contact uintptr
}

type IRadialControllerControlAcquiredEventArgs struct {
	win32.IInspectable
}

func (this *IRadialControllerControlAcquiredEventArgs) Vtbl() *IRadialControllerControlAcquiredEventArgsVtbl {
	return (*IRadialControllerControlAcquiredEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerControlAcquiredEventArgs) Get_Contact() *IRadialControllerScreenContact {
	var _result *IRadialControllerScreenContact
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Contact, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3D577EF4-3CEE-11E6-B535-001BDC06AB3B
var IID_IRadialControllerControlAcquiredEventArgs2 = syscall.GUID{0x3D577EF4, 0x3CEE, 0x11E6,
	[8]byte{0xB5, 0x35, 0x00, 0x1B, 0xDC, 0x06, 0xAB, 0x3B}}

type IRadialControllerControlAcquiredEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_IsButtonPressed() bool
	Get_SimpleHapticsController() *ISimpleHapticsController
}

type IRadialControllerControlAcquiredEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsButtonPressed         uintptr
	Get_SimpleHapticsController uintptr
}

type IRadialControllerControlAcquiredEventArgs2 struct {
	win32.IInspectable
}

func (this *IRadialControllerControlAcquiredEventArgs2) Vtbl() *IRadialControllerControlAcquiredEventArgs2Vtbl {
	return (*IRadialControllerControlAcquiredEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerControlAcquiredEventArgs2) Get_IsButtonPressed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsButtonPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRadialControllerControlAcquiredEventArgs2) Get_SimpleHapticsController() *ISimpleHapticsController {
	var _result *ISimpleHapticsController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SimpleHapticsController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8506B35D-F640-4412-ABA0-BAD077E5EA8A
var IID_IRadialControllerMenu = syscall.GUID{0x8506B35D, 0xF640, 0x4412,
	[8]byte{0xAB, 0xA0, 0xBA, 0xD0, 0x77, 0xE5, 0xEA, 0x8A}}

type IRadialControllerMenuInterface interface {
	win32.IInspectableInterface
	Get_Items() *IVector[*IRadialControllerMenuItem]
	Get_IsEnabled() bool
	Put_IsEnabled(value bool)
	GetSelectedMenuItem() *IRadialControllerMenuItem
	SelectMenuItem(menuItem *IRadialControllerMenuItem)
	TrySelectPreviouslySelectedMenuItem() bool
}

type IRadialControllerMenuVtbl struct {
	win32.IInspectableVtbl
	Get_Items                           uintptr
	Get_IsEnabled                       uintptr
	Put_IsEnabled                       uintptr
	GetSelectedMenuItem                 uintptr
	SelectMenuItem                      uintptr
	TrySelectPreviouslySelectedMenuItem uintptr
}

type IRadialControllerMenu struct {
	win32.IInspectable
}

func (this *IRadialControllerMenu) Vtbl() *IRadialControllerMenuVtbl {
	return (*IRadialControllerMenuVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerMenu) Get_Items() *IVector[*IRadialControllerMenuItem] {
	var _result *IVector[*IRadialControllerMenuItem]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Items, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRadialControllerMenu) Get_IsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRadialControllerMenu) Put_IsEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IRadialControllerMenu) GetSelectedMenuItem() *IRadialControllerMenuItem {
	var _result *IRadialControllerMenuItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSelectedMenuItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRadialControllerMenu) SelectMenuItem(menuItem *IRadialControllerMenuItem) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SelectMenuItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(menuItem)))
	_ = _hr
}

func (this *IRadialControllerMenu) TrySelectPreviouslySelectedMenuItem() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySelectPreviouslySelectedMenuItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C80FC98D-AD0B-4C9C-8F2F-136A2373A6BA
var IID_IRadialControllerMenuItem = syscall.GUID{0xC80FC98D, 0xAD0B, 0x4C9C,
	[8]byte{0x8F, 0x2F, 0x13, 0x6A, 0x23, 0x73, 0xA6, 0xBA}}

type IRadialControllerMenuItemInterface interface {
	win32.IInspectableInterface
	Get_DisplayText() string
	Get_Tag() interface{}
	Put_Tag(value interface{})
	Add_Invoked(handler TypedEventHandler[*IRadialControllerMenuItem, interface{}]) EventRegistrationToken
	Remove_Invoked(token EventRegistrationToken)
}

type IRadialControllerMenuItemVtbl struct {
	win32.IInspectableVtbl
	Get_DisplayText uintptr
	Get_Tag         uintptr
	Put_Tag         uintptr
	Add_Invoked     uintptr
	Remove_Invoked  uintptr
}

type IRadialControllerMenuItem struct {
	win32.IInspectable
}

func (this *IRadialControllerMenuItem) Vtbl() *IRadialControllerMenuItemVtbl {
	return (*IRadialControllerMenuItemVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerMenuItem) Get_DisplayText() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayText, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IRadialControllerMenuItem) Get_Tag() interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Tag, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRadialControllerMenuItem) Put_Tag(value interface{}) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Tag, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IRadialControllerMenuItem) Add_Invoked(handler TypedEventHandler[*IRadialControllerMenuItem, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Invoked, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRadialControllerMenuItem) Remove_Invoked(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Invoked, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 249E0887-D842-4524-9DF8-E0D647EDC887
var IID_IRadialControllerMenuItemStatics = syscall.GUID{0x249E0887, 0xD842, 0x4524,
	[8]byte{0x9D, 0xF8, 0xE0, 0xD6, 0x47, 0xED, 0xC8, 0x87}}

type IRadialControllerMenuItemStaticsInterface interface {
	win32.IInspectableInterface
	CreateFromIcon(displayText string, icon *IRandomAccessStreamReference) *IRadialControllerMenuItem
	CreateFromKnownIcon(displayText string, value RadialControllerMenuKnownIcon) *IRadialControllerMenuItem
}

type IRadialControllerMenuItemStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateFromIcon      uintptr
	CreateFromKnownIcon uintptr
}

type IRadialControllerMenuItemStatics struct {
	win32.IInspectable
}

func (this *IRadialControllerMenuItemStatics) Vtbl() *IRadialControllerMenuItemStaticsVtbl {
	return (*IRadialControllerMenuItemStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerMenuItemStatics) CreateFromIcon(displayText string, icon *IRandomAccessStreamReference) *IRadialControllerMenuItem {
	var _result *IRadialControllerMenuItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromIcon, uintptr(unsafe.Pointer(this)), NewHStr(displayText).Ptr, uintptr(unsafe.Pointer(icon)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRadialControllerMenuItemStatics) CreateFromKnownIcon(displayText string, value RadialControllerMenuKnownIcon) *IRadialControllerMenuItem {
	var _result *IRadialControllerMenuItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromKnownIcon, uintptr(unsafe.Pointer(this)), NewHStr(displayText).Ptr, uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0CBB70BE-7E3E-48BD-BE04-2C7FCAA9C1FF
var IID_IRadialControllerMenuItemStatics2 = syscall.GUID{0x0CBB70BE, 0x7E3E, 0x48BD,
	[8]byte{0xBE, 0x04, 0x2C, 0x7F, 0xCA, 0xA9, 0xC1, 0xFF}}

type IRadialControllerMenuItemStatics2Interface interface {
	win32.IInspectableInterface
	CreateFromFontGlyph(displayText string, glyph string, fontFamily string) *IRadialControllerMenuItem
	CreateFromFontGlyphWithUri(displayText string, glyph string, fontFamily string, fontUri *IUriRuntimeClass) *IRadialControllerMenuItem
}

type IRadialControllerMenuItemStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateFromFontGlyph        uintptr
	CreateFromFontGlyphWithUri uintptr
}

type IRadialControllerMenuItemStatics2 struct {
	win32.IInspectable
}

func (this *IRadialControllerMenuItemStatics2) Vtbl() *IRadialControllerMenuItemStatics2Vtbl {
	return (*IRadialControllerMenuItemStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerMenuItemStatics2) CreateFromFontGlyph(displayText string, glyph string, fontFamily string) *IRadialControllerMenuItem {
	var _result *IRadialControllerMenuItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromFontGlyph, uintptr(unsafe.Pointer(this)), NewHStr(displayText).Ptr, NewHStr(glyph).Ptr, NewHStr(fontFamily).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRadialControllerMenuItemStatics2) CreateFromFontGlyphWithUri(displayText string, glyph string, fontFamily string, fontUri *IUriRuntimeClass) *IRadialControllerMenuItem {
	var _result *IRadialControllerMenuItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromFontGlyphWithUri, uintptr(unsafe.Pointer(this)), NewHStr(displayText).Ptr, NewHStr(glyph).Ptr, NewHStr(fontFamily).Ptr, uintptr(unsafe.Pointer(fontUri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 206AA435-E651-11E5-BF62-2C27D7404E85
var IID_IRadialControllerRotationChangedEventArgs = syscall.GUID{0x206AA435, 0xE651, 0x11E5,
	[8]byte{0xBF, 0x62, 0x2C, 0x27, 0xD7, 0x40, 0x4E, 0x85}}

type IRadialControllerRotationChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_RotationDeltaInDegrees() float64
	Get_Contact() *IRadialControllerScreenContact
}

type IRadialControllerRotationChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_RotationDeltaInDegrees uintptr
	Get_Contact                uintptr
}

type IRadialControllerRotationChangedEventArgs struct {
	win32.IInspectable
}

func (this *IRadialControllerRotationChangedEventArgs) Vtbl() *IRadialControllerRotationChangedEventArgsVtbl {
	return (*IRadialControllerRotationChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerRotationChangedEventArgs) Get_RotationDeltaInDegrees() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RotationDeltaInDegrees, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRadialControllerRotationChangedEventArgs) Get_Contact() *IRadialControllerScreenContact {
	var _result *IRadialControllerScreenContact
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Contact, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3D577EEC-4CEE-11E6-B535-001BDC06AB3B
var IID_IRadialControllerRotationChangedEventArgs2 = syscall.GUID{0x3D577EEC, 0x4CEE, 0x11E6,
	[8]byte{0xB5, 0x35, 0x00, 0x1B, 0xDC, 0x06, 0xAB, 0x3B}}

type IRadialControllerRotationChangedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_IsButtonPressed() bool
	Get_SimpleHapticsController() *ISimpleHapticsController
}

type IRadialControllerRotationChangedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsButtonPressed         uintptr
	Get_SimpleHapticsController uintptr
}

type IRadialControllerRotationChangedEventArgs2 struct {
	win32.IInspectable
}

func (this *IRadialControllerRotationChangedEventArgs2) Vtbl() *IRadialControllerRotationChangedEventArgs2Vtbl {
	return (*IRadialControllerRotationChangedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerRotationChangedEventArgs2) Get_IsButtonPressed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsButtonPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRadialControllerRotationChangedEventArgs2) Get_SimpleHapticsController() *ISimpleHapticsController {
	var _result *ISimpleHapticsController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SimpleHapticsController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 206AA434-E651-11E5-BF62-2C27D7404E85
var IID_IRadialControllerScreenContact = syscall.GUID{0x206AA434, 0xE651, 0x11E5,
	[8]byte{0xBF, 0x62, 0x2C, 0x27, 0xD7, 0x40, 0x4E, 0x85}}

type IRadialControllerScreenContactInterface interface {
	win32.IInspectableInterface
	Get_Bounds() Rect
	Get_Position() Point
}

type IRadialControllerScreenContactVtbl struct {
	win32.IInspectableVtbl
	Get_Bounds   uintptr
	Get_Position uintptr
}

type IRadialControllerScreenContact struct {
	win32.IInspectable
}

func (this *IRadialControllerScreenContact) Vtbl() *IRadialControllerScreenContactVtbl {
	return (*IRadialControllerScreenContactVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerScreenContact) Get_Bounds() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bounds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRadialControllerScreenContact) Get_Position() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 206AA437-E651-11E5-BF62-2C27D7404E85
var IID_IRadialControllerScreenContactContinuedEventArgs = syscall.GUID{0x206AA437, 0xE651, 0x11E5,
	[8]byte{0xBF, 0x62, 0x2C, 0x27, 0xD7, 0x40, 0x4E, 0x85}}

type IRadialControllerScreenContactContinuedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Contact() *IRadialControllerScreenContact
}

type IRadialControllerScreenContactContinuedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Contact uintptr
}

type IRadialControllerScreenContactContinuedEventArgs struct {
	win32.IInspectable
}

func (this *IRadialControllerScreenContactContinuedEventArgs) Vtbl() *IRadialControllerScreenContactContinuedEventArgsVtbl {
	return (*IRadialControllerScreenContactContinuedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerScreenContactContinuedEventArgs) Get_Contact() *IRadialControllerScreenContact {
	var _result *IRadialControllerScreenContact
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Contact, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3D577EF1-3CEE-11E6-B535-001BDC06AB3B
var IID_IRadialControllerScreenContactContinuedEventArgs2 = syscall.GUID{0x3D577EF1, 0x3CEE, 0x11E6,
	[8]byte{0xB5, 0x35, 0x00, 0x1B, 0xDC, 0x06, 0xAB, 0x3B}}

type IRadialControllerScreenContactContinuedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_IsButtonPressed() bool
	Get_SimpleHapticsController() *ISimpleHapticsController
}

type IRadialControllerScreenContactContinuedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsButtonPressed         uintptr
	Get_SimpleHapticsController uintptr
}

type IRadialControllerScreenContactContinuedEventArgs2 struct {
	win32.IInspectable
}

func (this *IRadialControllerScreenContactContinuedEventArgs2) Vtbl() *IRadialControllerScreenContactContinuedEventArgs2Vtbl {
	return (*IRadialControllerScreenContactContinuedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerScreenContactContinuedEventArgs2) Get_IsButtonPressed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsButtonPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRadialControllerScreenContactContinuedEventArgs2) Get_SimpleHapticsController() *ISimpleHapticsController {
	var _result *ISimpleHapticsController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SimpleHapticsController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3D577EF2-3CEE-11E6-B535-001BDC06AB3B
var IID_IRadialControllerScreenContactEndedEventArgs = syscall.GUID{0x3D577EF2, 0x3CEE, 0x11E6,
	[8]byte{0xB5, 0x35, 0x00, 0x1B, 0xDC, 0x06, 0xAB, 0x3B}}

type IRadialControllerScreenContactEndedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_IsButtonPressed() bool
	Get_SimpleHapticsController() *ISimpleHapticsController
}

type IRadialControllerScreenContactEndedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_IsButtonPressed         uintptr
	Get_SimpleHapticsController uintptr
}

type IRadialControllerScreenContactEndedEventArgs struct {
	win32.IInspectable
}

func (this *IRadialControllerScreenContactEndedEventArgs) Vtbl() *IRadialControllerScreenContactEndedEventArgsVtbl {
	return (*IRadialControllerScreenContactEndedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerScreenContactEndedEventArgs) Get_IsButtonPressed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsButtonPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRadialControllerScreenContactEndedEventArgs) Get_SimpleHapticsController() *ISimpleHapticsController {
	var _result *ISimpleHapticsController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SimpleHapticsController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 206AA436-E651-11E5-BF62-2C27D7404E85
var IID_IRadialControllerScreenContactStartedEventArgs = syscall.GUID{0x206AA436, 0xE651, 0x11E5,
	[8]byte{0xBF, 0x62, 0x2C, 0x27, 0xD7, 0x40, 0x4E, 0x85}}

type IRadialControllerScreenContactStartedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Contact() *IRadialControllerScreenContact
}

type IRadialControllerScreenContactStartedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Contact uintptr
}

type IRadialControllerScreenContactStartedEventArgs struct {
	win32.IInspectable
}

func (this *IRadialControllerScreenContactStartedEventArgs) Vtbl() *IRadialControllerScreenContactStartedEventArgsVtbl {
	return (*IRadialControllerScreenContactStartedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerScreenContactStartedEventArgs) Get_Contact() *IRadialControllerScreenContact {
	var _result *IRadialControllerScreenContact
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Contact, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3D577EF0-3CEE-11E6-B535-001BDC06AB3B
var IID_IRadialControllerScreenContactStartedEventArgs2 = syscall.GUID{0x3D577EF0, 0x3CEE, 0x11E6,
	[8]byte{0xB5, 0x35, 0x00, 0x1B, 0xDC, 0x06, 0xAB, 0x3B}}

type IRadialControllerScreenContactStartedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_IsButtonPressed() bool
	Get_SimpleHapticsController() *ISimpleHapticsController
}

type IRadialControllerScreenContactStartedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsButtonPressed         uintptr
	Get_SimpleHapticsController uintptr
}

type IRadialControllerScreenContactStartedEventArgs2 struct {
	win32.IInspectable
}

func (this *IRadialControllerScreenContactStartedEventArgs2) Vtbl() *IRadialControllerScreenContactStartedEventArgs2Vtbl {
	return (*IRadialControllerScreenContactStartedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerScreenContactStartedEventArgs2) Get_IsButtonPressed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsButtonPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRadialControllerScreenContactStartedEventArgs2) Get_SimpleHapticsController() *ISimpleHapticsController {
	var _result *ISimpleHapticsController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SimpleHapticsController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FADED0B7-B84C-4894-87AA-8F25AA5F288B
var IID_IRadialControllerStatics = syscall.GUID{0xFADED0B7, 0xB84C, 0x4894,
	[8]byte{0x87, 0xAA, 0x8F, 0x25, 0xAA, 0x5F, 0x28, 0x8B}}

type IRadialControllerStaticsInterface interface {
	win32.IInspectableInterface
	IsSupported() bool
	CreateForCurrentView() *IRadialController
}

type IRadialControllerStaticsVtbl struct {
	win32.IInspectableVtbl
	IsSupported          uintptr
	CreateForCurrentView uintptr
}

type IRadialControllerStatics struct {
	win32.IInspectable
}

func (this *IRadialControllerStatics) Vtbl() *IRadialControllerStaticsVtbl {
	return (*IRadialControllerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerStatics) IsSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRadialControllerStatics) CreateForCurrentView() *IRadialController {
	var _result *IRadialController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4CBF40BD-AF7A-4A36-9476-B1DCE141709A
var IID_IRightTappedEventArgs = syscall.GUID{0x4CBF40BD, 0xAF7A, 0x4A36,
	[8]byte{0x94, 0x76, 0xB1, 0xDC, 0xE1, 0x41, 0x70, 0x9A}}

type IRightTappedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_PointerDeviceType() PointerDeviceType
	Get_Position() Point
}

type IRightTappedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_PointerDeviceType uintptr
	Get_Position          uintptr
}

type IRightTappedEventArgs struct {
	win32.IInspectable
}

func (this *IRightTappedEventArgs) Vtbl() *IRightTappedEventArgsVtbl {
	return (*IRightTappedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRightTappedEventArgs) Get_PointerDeviceType() PointerDeviceType {
	var _result PointerDeviceType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PointerDeviceType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRightTappedEventArgs) Get_Position() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 61C7B7BB-9F57-5857-A33C-C58C3DFA959E
var IID_IRightTappedEventArgs2 = syscall.GUID{0x61C7B7BB, 0x9F57, 0x5857,
	[8]byte{0xA3, 0x3C, 0xC5, 0x8C, 0x3D, 0xFA, 0x95, 0x9E}}

type IRightTappedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_ContactCount() uint32
}

type IRightTappedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_ContactCount uintptr
}

type IRightTappedEventArgs2 struct {
	win32.IInspectable
}

func (this *IRightTappedEventArgs2) Vtbl() *IRightTappedEventArgs2Vtbl {
	return (*IRightTappedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRightTappedEventArgs2) Get_ContactCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContactCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 59B893A9-73BC-52B5-BA41-82511B2CB46C
var IID_ISystemButtonEventController = syscall.GUID{0x59B893A9, 0x73BC, 0x52B5,
	[8]byte{0xBA, 0x41, 0x82, 0x51, 0x1B, 0x2C, 0xB4, 0x6C}}

type ISystemButtonEventControllerInterface interface {
	win32.IInspectableInterface
	Add_SystemFunctionButtonPressed(handler TypedEventHandler[*ISystemButtonEventController, *ISystemFunctionButtonEventArgs]) EventRegistrationToken
	Remove_SystemFunctionButtonPressed(token EventRegistrationToken)
	Add_SystemFunctionButtonReleased(handler TypedEventHandler[*ISystemButtonEventController, *ISystemFunctionButtonEventArgs]) EventRegistrationToken
	Remove_SystemFunctionButtonReleased(token EventRegistrationToken)
	Add_SystemFunctionLockChanged(handler TypedEventHandler[*ISystemButtonEventController, *ISystemFunctionLockChangedEventArgs]) EventRegistrationToken
	Remove_SystemFunctionLockChanged(token EventRegistrationToken)
	Add_SystemFunctionLockIndicatorChanged(handler TypedEventHandler[*ISystemButtonEventController, *ISystemFunctionLockIndicatorChangedEventArgs]) EventRegistrationToken
	Remove_SystemFunctionLockIndicatorChanged(token EventRegistrationToken)
}

type ISystemButtonEventControllerVtbl struct {
	win32.IInspectableVtbl
	Add_SystemFunctionButtonPressed           uintptr
	Remove_SystemFunctionButtonPressed        uintptr
	Add_SystemFunctionButtonReleased          uintptr
	Remove_SystemFunctionButtonReleased       uintptr
	Add_SystemFunctionLockChanged             uintptr
	Remove_SystemFunctionLockChanged          uintptr
	Add_SystemFunctionLockIndicatorChanged    uintptr
	Remove_SystemFunctionLockIndicatorChanged uintptr
}

type ISystemButtonEventController struct {
	win32.IInspectable
}

func (this *ISystemButtonEventController) Vtbl() *ISystemButtonEventControllerVtbl {
	return (*ISystemButtonEventControllerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemButtonEventController) Add_SystemFunctionButtonPressed(handler TypedEventHandler[*ISystemButtonEventController, *ISystemFunctionButtonEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SystemFunctionButtonPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemButtonEventController) Remove_SystemFunctionButtonPressed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SystemFunctionButtonPressed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISystemButtonEventController) Add_SystemFunctionButtonReleased(handler TypedEventHandler[*ISystemButtonEventController, *ISystemFunctionButtonEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SystemFunctionButtonReleased, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemButtonEventController) Remove_SystemFunctionButtonReleased(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SystemFunctionButtonReleased, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISystemButtonEventController) Add_SystemFunctionLockChanged(handler TypedEventHandler[*ISystemButtonEventController, *ISystemFunctionLockChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SystemFunctionLockChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemButtonEventController) Remove_SystemFunctionLockChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SystemFunctionLockChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISystemButtonEventController) Add_SystemFunctionLockIndicatorChanged(handler TypedEventHandler[*ISystemButtonEventController, *ISystemFunctionLockIndicatorChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SystemFunctionLockIndicatorChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemButtonEventController) Remove_SystemFunctionLockIndicatorChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SystemFunctionLockIndicatorChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 632FB07B-20BD-5E15-AF4A-00DBF2064FFA
var IID_ISystemButtonEventControllerStatics = syscall.GUID{0x632FB07B, 0x20BD, 0x5E15,
	[8]byte{0xAF, 0x4A, 0x00, 0xDB, 0xF2, 0x06, 0x4F, 0xFA}}

type ISystemButtonEventControllerStaticsInterface interface {
	win32.IInspectableInterface
	CreateForDispatcherQueue(queue *IDispatcherQueue) *ISystemButtonEventController
}

type ISystemButtonEventControllerStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateForDispatcherQueue uintptr
}

type ISystemButtonEventControllerStatics struct {
	win32.IInspectable
}

func (this *ISystemButtonEventControllerStatics) Vtbl() *ISystemButtonEventControllerStaticsVtbl {
	return (*ISystemButtonEventControllerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemButtonEventControllerStatics) CreateForDispatcherQueue(queue *IDispatcherQueue) *ISystemButtonEventController {
	var _result *ISystemButtonEventController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateForDispatcherQueue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(queue)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4833896F-80D1-5DD6-92A7-62A508FFEF5A
var IID_ISystemFunctionButtonEventArgs = syscall.GUID{0x4833896F, 0x80D1, 0x5DD6,
	[8]byte{0x92, 0xA7, 0x62, 0xA5, 0x08, 0xFF, 0xEF, 0x5A}}

type ISystemFunctionButtonEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() uint64
	Get_Handled() bool
	Put_Handled(value bool)
}

type ISystemFunctionButtonEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp uintptr
	Get_Handled   uintptr
	Put_Handled   uintptr
}

type ISystemFunctionButtonEventArgs struct {
	win32.IInspectable
}

func (this *ISystemFunctionButtonEventArgs) Vtbl() *ISystemFunctionButtonEventArgsVtbl {
	return (*ISystemFunctionButtonEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemFunctionButtonEventArgs) Get_Timestamp() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemFunctionButtonEventArgs) Get_Handled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Handled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemFunctionButtonEventArgs) Put_Handled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Handled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// CD040608-FCF9-585C-BEAB-F1D2EAF364AB
var IID_ISystemFunctionLockChangedEventArgs = syscall.GUID{0xCD040608, 0xFCF9, 0x585C,
	[8]byte{0xBE, 0xAB, 0xF1, 0xD2, 0xEA, 0xF3, 0x64, 0xAB}}

type ISystemFunctionLockChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() uint64
	Get_IsLocked() bool
	Get_Handled() bool
	Put_Handled(value bool)
}

type ISystemFunctionLockChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp uintptr
	Get_IsLocked  uintptr
	Get_Handled   uintptr
	Put_Handled   uintptr
}

type ISystemFunctionLockChangedEventArgs struct {
	win32.IInspectable
}

func (this *ISystemFunctionLockChangedEventArgs) Vtbl() *ISystemFunctionLockChangedEventArgsVtbl {
	return (*ISystemFunctionLockChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemFunctionLockChangedEventArgs) Get_Timestamp() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemFunctionLockChangedEventArgs) Get_IsLocked() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsLocked, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemFunctionLockChangedEventArgs) Get_Handled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Handled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemFunctionLockChangedEventArgs) Put_Handled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Handled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// B212B94E-7A6F-58AE-B304-BAE61D0371B9
var IID_ISystemFunctionLockIndicatorChangedEventArgs = syscall.GUID{0xB212B94E, 0x7A6F, 0x58AE,
	[8]byte{0xB3, 0x04, 0xBA, 0xE6, 0x1D, 0x03, 0x71, 0xB9}}

type ISystemFunctionLockIndicatorChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() uint64
	Get_IsIndicatorOn() bool
	Get_Handled() bool
	Put_Handled(value bool)
}

type ISystemFunctionLockIndicatorChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp     uintptr
	Get_IsIndicatorOn uintptr
	Get_Handled       uintptr
	Put_Handled       uintptr
}

type ISystemFunctionLockIndicatorChangedEventArgs struct {
	win32.IInspectable
}

func (this *ISystemFunctionLockIndicatorChangedEventArgs) Vtbl() *ISystemFunctionLockIndicatorChangedEventArgsVtbl {
	return (*ISystemFunctionLockIndicatorChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemFunctionLockIndicatorChangedEventArgs) Get_Timestamp() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemFunctionLockIndicatorChangedEventArgs) Get_IsIndicatorOn() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsIndicatorOn, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemFunctionLockIndicatorChangedEventArgs) Get_Handled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Handled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemFunctionLockIndicatorChangedEventArgs) Put_Handled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Handled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// CFA126E4-253A-4C3C-953B-395C37AED309
var IID_ITappedEventArgs = syscall.GUID{0xCFA126E4, 0x253A, 0x4C3C,
	[8]byte{0x95, 0x3B, 0x39, 0x5C, 0x37, 0xAE, 0xD3, 0x09}}

type ITappedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_PointerDeviceType() PointerDeviceType
	Get_Position() Point
	Get_TapCount() uint32
}

type ITappedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_PointerDeviceType uintptr
	Get_Position          uintptr
	Get_TapCount          uintptr
}

type ITappedEventArgs struct {
	win32.IInspectable
}

func (this *ITappedEventArgs) Vtbl() *ITappedEventArgsVtbl {
	return (*ITappedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITappedEventArgs) Get_PointerDeviceType() PointerDeviceType {
	var _result PointerDeviceType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PointerDeviceType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITappedEventArgs) Get_Position() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITappedEventArgs) Get_TapCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TapCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 294388F2-177E-51D5-BE56-EE0866FA968C
var IID_ITappedEventArgs2 = syscall.GUID{0x294388F2, 0x177E, 0x51D5,
	[8]byte{0xBE, 0x56, 0xEE, 0x08, 0x66, 0xFA, 0x96, 0x8C}}

type ITappedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_ContactCount() uint32
}

type ITappedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_ContactCount uintptr
}

type ITappedEventArgs2 struct {
	win32.IInspectable
}

func (this *ITappedEventArgs2) Vtbl() *ITappedEventArgs2Vtbl {
	return (*ITappedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITappedEventArgs2) Get_ContactCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContactCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type RadialController struct {
	RtClass
	*IRadialController
}

func NewIRadialControllerStatics() *IRadialControllerStatics {
	var p *IRadialControllerStatics
	hs := NewHStr("Windows.UI.Input.RadialController")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IRadialControllerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type RadialControllerButtonClickedEventArgs struct {
	RtClass
	*IRadialControllerButtonClickedEventArgs
}

type RadialControllerButtonHoldingEventArgs struct {
	RtClass
	*IRadialControllerButtonHoldingEventArgs
}

type RadialControllerButtonPressedEventArgs struct {
	RtClass
	*IRadialControllerButtonPressedEventArgs
}

type RadialControllerButtonReleasedEventArgs struct {
	RtClass
	*IRadialControllerButtonReleasedEventArgs
}

type RadialControllerConfiguration struct {
	RtClass
	*IRadialControllerConfiguration
}

func NewIRadialControllerConfigurationStatics() *IRadialControllerConfigurationStatics {
	var p *IRadialControllerConfigurationStatics
	hs := NewHStr("Windows.UI.Input.RadialControllerConfiguration")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IRadialControllerConfigurationStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIRadialControllerConfigurationStatics2() *IRadialControllerConfigurationStatics2 {
	var p *IRadialControllerConfigurationStatics2
	hs := NewHStr("Windows.UI.Input.RadialControllerConfiguration")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IRadialControllerConfigurationStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type RadialControllerControlAcquiredEventArgs struct {
	RtClass
	*IRadialControllerControlAcquiredEventArgs
}

type RadialControllerMenu struct {
	RtClass
	*IRadialControllerMenu
}

type RadialControllerMenuItem struct {
	RtClass
	*IRadialControllerMenuItem
}

func NewIRadialControllerMenuItemStatics2() *IRadialControllerMenuItemStatics2 {
	var p *IRadialControllerMenuItemStatics2
	hs := NewHStr("Windows.UI.Input.RadialControllerMenuItem")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IRadialControllerMenuItemStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIRadialControllerMenuItemStatics() *IRadialControllerMenuItemStatics {
	var p *IRadialControllerMenuItemStatics
	hs := NewHStr("Windows.UI.Input.RadialControllerMenuItem")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IRadialControllerMenuItemStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type RadialControllerRotationChangedEventArgs struct {
	RtClass
	*IRadialControllerRotationChangedEventArgs
}

type RadialControllerScreenContact struct {
	RtClass
	*IRadialControllerScreenContact
}

type RadialControllerScreenContactContinuedEventArgs struct {
	RtClass
	*IRadialControllerScreenContactContinuedEventArgs
}

type RadialControllerScreenContactEndedEventArgs struct {
	RtClass
	*IRadialControllerScreenContactEndedEventArgs
}

type RadialControllerScreenContactStartedEventArgs struct {
	RtClass
	*IRadialControllerScreenContactStartedEventArgs
}
