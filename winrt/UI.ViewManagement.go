package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type ApplicationViewBoundsMode int32

const (
	ApplicationViewBoundsMode_UseVisible    ApplicationViewBoundsMode = 0
	ApplicationViewBoundsMode_UseCoreWindow ApplicationViewBoundsMode = 1
)

// enum
type ApplicationViewMode int32

const (
	ApplicationViewMode_Default        ApplicationViewMode = 0
	ApplicationViewMode_CompactOverlay ApplicationViewMode = 1
)

// enum
type ApplicationViewOrientation int32

const (
	ApplicationViewOrientation_Landscape ApplicationViewOrientation = 0
	ApplicationViewOrientation_Portrait  ApplicationViewOrientation = 1
)

// enum
type ApplicationViewState int32

const (
	ApplicationViewState_FullScreenLandscape ApplicationViewState = 0
	ApplicationViewState_Filled              ApplicationViewState = 1
	ApplicationViewState_Snapped             ApplicationViewState = 2
	ApplicationViewState_FullScreenPortrait  ApplicationViewState = 3
)

// enum
// flags
type ApplicationViewSwitchingOptions uint32

const (
	ApplicationViewSwitchingOptions_Default          ApplicationViewSwitchingOptions = 0
	ApplicationViewSwitchingOptions_SkipAnimation    ApplicationViewSwitchingOptions = 1
	ApplicationViewSwitchingOptions_ConsolidateViews ApplicationViewSwitchingOptions = 2
)

// enum
type ApplicationViewWindowingMode int32

const (
	ApplicationViewWindowingMode_Auto                    ApplicationViewWindowingMode = 0
	ApplicationViewWindowingMode_PreferredLaunchViewSize ApplicationViewWindowingMode = 1
	ApplicationViewWindowingMode_FullScreen              ApplicationViewWindowingMode = 2
	ApplicationViewWindowingMode_CompactOverlay          ApplicationViewWindowingMode = 3
	ApplicationViewWindowingMode_Maximized               ApplicationViewWindowingMode = 4
)

// enum
type FullScreenSystemOverlayMode int32

const (
	FullScreenSystemOverlayMode_Standard FullScreenSystemOverlayMode = 0
	FullScreenSystemOverlayMode_Minimal  FullScreenSystemOverlayMode = 1
)

// enum
type HandPreference int32

const (
	HandPreference_LeftHanded  HandPreference = 0
	HandPreference_RightHanded HandPreference = 1
)

// enum
type UIColorType int32

const (
	UIColorType_Background   UIColorType = 0
	UIColorType_Foreground   UIColorType = 1
	UIColorType_AccentDark3  UIColorType = 2
	UIColorType_AccentDark2  UIColorType = 3
	UIColorType_AccentDark1  UIColorType = 4
	UIColorType_Accent       UIColorType = 5
	UIColorType_AccentLight1 UIColorType = 6
	UIColorType_AccentLight2 UIColorType = 7
	UIColorType_AccentLight3 UIColorType = 8
	UIColorType_Complement   UIColorType = 9
)

// enum
type UIElementType int32

const (
	UIElementType_ActiveCaption        UIElementType = 0
	UIElementType_Background           UIElementType = 1
	UIElementType_ButtonFace           UIElementType = 2
	UIElementType_ButtonText           UIElementType = 3
	UIElementType_CaptionText          UIElementType = 4
	UIElementType_GrayText             UIElementType = 5
	UIElementType_Highlight            UIElementType = 6
	UIElementType_HighlightText        UIElementType = 7
	UIElementType_Hotlight             UIElementType = 8
	UIElementType_InactiveCaption      UIElementType = 9
	UIElementType_InactiveCaptionText  UIElementType = 10
	UIElementType_Window               UIElementType = 11
	UIElementType_WindowText           UIElementType = 12
	UIElementType_AccentColor          UIElementType = 1000
	UIElementType_TextHigh             UIElementType = 1001
	UIElementType_TextMedium           UIElementType = 1002
	UIElementType_TextLow              UIElementType = 1003
	UIElementType_TextContrastWithHigh UIElementType = 1004
	UIElementType_NonTextHigh          UIElementType = 1005
	UIElementType_NonTextMediumHigh    UIElementType = 1006
	UIElementType_NonTextMedium        UIElementType = 1007
	UIElementType_NonTextMediumLow     UIElementType = 1008
	UIElementType_NonTextLow           UIElementType = 1009
	UIElementType_PageBackground       UIElementType = 1010
	UIElementType_PopupBackground      UIElementType = 1011
	UIElementType_OverlayOutsidePopup  UIElementType = 1012
)

// enum
type UserInteractionMode int32

const (
	UserInteractionMode_Mouse UserInteractionMode = 0
	UserInteractionMode_Touch UserInteractionMode = 1
)

// enum
type ViewSizePreference int32

const (
	ViewSizePreference_Default    ViewSizePreference = 0
	ViewSizePreference_UseLess    ViewSizePreference = 1
	ViewSizePreference_UseHalf    ViewSizePreference = 2
	ViewSizePreference_UseMore    ViewSizePreference = 3
	ViewSizePreference_UseMinimum ViewSizePreference = 4
	ViewSizePreference_UseNone    ViewSizePreference = 5
	ViewSizePreference_Custom     ViewSizePreference = 6
)

// structs

type ViewManagementViewScalingContract struct {
}

// interfaces

// FE0E8147-C4C0-4562-B962-1327B52AD5B9
var IID_IAccessibilitySettings = syscall.GUID{0xFE0E8147, 0xC4C0, 0x4562,
	[8]byte{0xB9, 0x62, 0x13, 0x27, 0xB5, 0x2A, 0xD5, 0xB9}}

type IAccessibilitySettingsInterface interface {
	win32.IInspectableInterface
	Get_HighContrast() bool
	Get_HighContrastScheme() string
	Add_HighContrastChanged(handler TypedEventHandler[*IAccessibilitySettings, interface{}]) EventRegistrationToken
	Remove_HighContrastChanged(cookie EventRegistrationToken)
}

type IAccessibilitySettingsVtbl struct {
	win32.IInspectableVtbl
	Get_HighContrast           uintptr
	Get_HighContrastScheme     uintptr
	Add_HighContrastChanged    uintptr
	Remove_HighContrastChanged uintptr
}

type IAccessibilitySettings struct {
	win32.IInspectable
}

func (this *IAccessibilitySettings) Vtbl() *IAccessibilitySettingsVtbl {
	return (*IAccessibilitySettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccessibilitySettings) Get_HighContrast() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HighContrast, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAccessibilitySettings) Get_HighContrastScheme() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HighContrastScheme, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAccessibilitySettings) Add_HighContrastChanged(handler TypedEventHandler[*IAccessibilitySettings, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_HighContrastChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAccessibilitySettings) Remove_HighContrastChanged(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_HighContrastChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

// DCA71BB6-7350-492B-AAC7-C8A13D7224AD
var IID_IActivationViewSwitcher = syscall.GUID{0xDCA71BB6, 0x7350, 0x492B,
	[8]byte{0xAA, 0xC7, 0xC8, 0xA1, 0x3D, 0x72, 0x24, 0xAD}}

type IActivationViewSwitcherInterface interface {
	win32.IInspectableInterface
	ShowAsStandaloneAsync(viewId int32) *IAsyncAction
	ShowAsStandaloneWithSizePreferenceAsync(viewId int32, sizePreference ViewSizePreference) *IAsyncAction
	IsViewPresentedOnActivationVirtualDesktop(viewId int32) bool
}

type IActivationViewSwitcherVtbl struct {
	win32.IInspectableVtbl
	ShowAsStandaloneAsync                     uintptr
	ShowAsStandaloneWithSizePreferenceAsync   uintptr
	IsViewPresentedOnActivationVirtualDesktop uintptr
}

type IActivationViewSwitcher struct {
	win32.IInspectable
}

func (this *IActivationViewSwitcher) Vtbl() *IActivationViewSwitcherVtbl {
	return (*IActivationViewSwitcherVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IActivationViewSwitcher) ShowAsStandaloneAsync(viewId int32) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowAsStandaloneAsync, uintptr(unsafe.Pointer(this)), uintptr(viewId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IActivationViewSwitcher) ShowAsStandaloneWithSizePreferenceAsync(viewId int32, sizePreference ViewSizePreference) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowAsStandaloneWithSizePreferenceAsync, uintptr(unsafe.Pointer(this)), uintptr(viewId), uintptr(sizePreference), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IActivationViewSwitcher) IsViewPresentedOnActivationVirtualDesktop(viewId int32) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsViewPresentedOnActivationVirtualDesktop, uintptr(unsafe.Pointer(this)), uintptr(viewId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D222D519-4361-451E-96C4-60F4F9742DB0
var IID_IApplicationView = syscall.GUID{0xD222D519, 0x4361, 0x451E,
	[8]byte{0x96, 0xC4, 0x60, 0xF4, 0xF9, 0x74, 0x2D, 0xB0}}

type IApplicationViewInterface interface {
	win32.IInspectableInterface
	Get_Orientation() ApplicationViewOrientation
	Get_AdjacentToLeftDisplayEdge() bool
	Get_AdjacentToRightDisplayEdge() bool
	Get_IsFullScreen() bool
	Get_IsOnLockScreen() bool
	Get_IsScreenCaptureEnabled() bool
	Put_IsScreenCaptureEnabled(value bool)
	Put_Title(value string)
	Get_Title() string
	Get_Id() int32
	Add_Consolidated(handler TypedEventHandler[*IApplicationView, *IApplicationViewConsolidatedEventArgs]) EventRegistrationToken
	Remove_Consolidated(token EventRegistrationToken)
}

type IApplicationViewVtbl struct {
	win32.IInspectableVtbl
	Get_Orientation                uintptr
	Get_AdjacentToLeftDisplayEdge  uintptr
	Get_AdjacentToRightDisplayEdge uintptr
	Get_IsFullScreen               uintptr
	Get_IsOnLockScreen             uintptr
	Get_IsScreenCaptureEnabled     uintptr
	Put_IsScreenCaptureEnabled     uintptr
	Put_Title                      uintptr
	Get_Title                      uintptr
	Get_Id                         uintptr
	Add_Consolidated               uintptr
	Remove_Consolidated            uintptr
}

type IApplicationView struct {
	win32.IInspectable
}

func (this *IApplicationView) Vtbl() *IApplicationViewVtbl {
	return (*IApplicationViewVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationView) Get_Orientation() ApplicationViewOrientation {
	var _result ApplicationViewOrientation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Orientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationView) Get_AdjacentToLeftDisplayEdge() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AdjacentToLeftDisplayEdge, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationView) Get_AdjacentToRightDisplayEdge() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AdjacentToRightDisplayEdge, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationView) Get_IsFullScreen() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsFullScreen, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationView) Get_IsOnLockScreen() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsOnLockScreen, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationView) Get_IsScreenCaptureEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsScreenCaptureEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationView) Put_IsScreenCaptureEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsScreenCaptureEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IApplicationView) Put_Title(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Title, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IApplicationView) Get_Title() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IApplicationView) Get_Id() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationView) Add_Consolidated(handler TypedEventHandler[*IApplicationView, *IApplicationViewConsolidatedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Consolidated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationView) Remove_Consolidated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Consolidated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// E876B196-A545-40DC-B594-450CBA68CC00
var IID_IApplicationView2 = syscall.GUID{0xE876B196, 0xA545, 0x40DC,
	[8]byte{0xB5, 0x94, 0x45, 0x0C, 0xBA, 0x68, 0xCC, 0x00}}

type IApplicationView2Interface interface {
	win32.IInspectableInterface
	Get_SuppressSystemOverlays() bool
	Put_SuppressSystemOverlays(value bool)
	Get_VisibleBounds() Rect
	Add_VisibleBoundsChanged(handler TypedEventHandler[*IApplicationView, interface{}]) EventRegistrationToken
	Remove_VisibleBoundsChanged(token EventRegistrationToken)
	SetDesiredBoundsMode(boundsMode ApplicationViewBoundsMode) bool
	Get_DesiredBoundsMode() ApplicationViewBoundsMode
}

type IApplicationView2Vtbl struct {
	win32.IInspectableVtbl
	Get_SuppressSystemOverlays  uintptr
	Put_SuppressSystemOverlays  uintptr
	Get_VisibleBounds           uintptr
	Add_VisibleBoundsChanged    uintptr
	Remove_VisibleBoundsChanged uintptr
	SetDesiredBoundsMode        uintptr
	Get_DesiredBoundsMode       uintptr
}

type IApplicationView2 struct {
	win32.IInspectable
}

func (this *IApplicationView2) Vtbl() *IApplicationView2Vtbl {
	return (*IApplicationView2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationView2) Get_SuppressSystemOverlays() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SuppressSystemOverlays, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationView2) Put_SuppressSystemOverlays(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SuppressSystemOverlays, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IApplicationView2) Get_VisibleBounds() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VisibleBounds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationView2) Add_VisibleBoundsChanged(handler TypedEventHandler[*IApplicationView, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_VisibleBoundsChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationView2) Remove_VisibleBoundsChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_VisibleBoundsChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IApplicationView2) SetDesiredBoundsMode(boundsMode ApplicationViewBoundsMode) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetDesiredBoundsMode, uintptr(unsafe.Pointer(this)), uintptr(boundsMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationView2) Get_DesiredBoundsMode() ApplicationViewBoundsMode {
	var _result ApplicationViewBoundsMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesiredBoundsMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 903C9CE5-793A-4FDF-A2B2-AF1AC21E3108
var IID_IApplicationView3 = syscall.GUID{0x903C9CE5, 0x793A, 0x4FDF,
	[8]byte{0xA2, 0xB2, 0xAF, 0x1A, 0xC2, 0x1E, 0x31, 0x08}}

type IApplicationView3Interface interface {
	win32.IInspectableInterface
	Get_TitleBar() *IApplicationViewTitleBar
	Get_FullScreenSystemOverlayMode() FullScreenSystemOverlayMode
	Put_FullScreenSystemOverlayMode(value FullScreenSystemOverlayMode)
	Get_IsFullScreenMode() bool
	TryEnterFullScreenMode() bool
	ExitFullScreenMode()
	ShowStandardSystemOverlays()
	TryResizeView(value Size) bool
	SetPreferredMinSize(minSize Size)
}

type IApplicationView3Vtbl struct {
	win32.IInspectableVtbl
	Get_TitleBar                    uintptr
	Get_FullScreenSystemOverlayMode uintptr
	Put_FullScreenSystemOverlayMode uintptr
	Get_IsFullScreenMode            uintptr
	TryEnterFullScreenMode          uintptr
	ExitFullScreenMode              uintptr
	ShowStandardSystemOverlays      uintptr
	TryResizeView                   uintptr
	SetPreferredMinSize             uintptr
}

type IApplicationView3 struct {
	win32.IInspectable
}

func (this *IApplicationView3) Vtbl() *IApplicationView3Vtbl {
	return (*IApplicationView3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationView3) Get_TitleBar() *IApplicationViewTitleBar {
	var _result *IApplicationViewTitleBar
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TitleBar, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationView3) Get_FullScreenSystemOverlayMode() FullScreenSystemOverlayMode {
	var _result FullScreenSystemOverlayMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FullScreenSystemOverlayMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationView3) Put_FullScreenSystemOverlayMode(value FullScreenSystemOverlayMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FullScreenSystemOverlayMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IApplicationView3) Get_IsFullScreenMode() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsFullScreenMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationView3) TryEnterFullScreenMode() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryEnterFullScreenMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationView3) ExitFullScreenMode() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ExitFullScreenMode, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IApplicationView3) ShowStandardSystemOverlays() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowStandardSystemOverlays, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IApplicationView3) TryResizeView(value Size) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryResizeView, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationView3) SetPreferredMinSize(minSize Size) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPreferredMinSize, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&minSize)))
	_ = _hr
}

// 15E5CBEC-9E0F-46B5-BC3F-9BF653E74B5E
var IID_IApplicationView4 = syscall.GUID{0x15E5CBEC, 0x9E0F, 0x46B5,
	[8]byte{0xBC, 0x3F, 0x9B, 0xF6, 0x53, 0xE7, 0x4B, 0x5E}}

type IApplicationView4Interface interface {
	win32.IInspectableInterface
	Get_ViewMode() ApplicationViewMode
	IsViewModeSupported(viewMode ApplicationViewMode) bool
	TryEnterViewModeAsync(viewMode ApplicationViewMode) *IAsyncOperation[bool]
	TryEnterViewModeWithPreferencesAsync(viewMode ApplicationViewMode, viewModePreferences *IViewModePreferences) *IAsyncOperation[bool]
	TryConsolidateAsync() *IAsyncOperation[bool]
}

type IApplicationView4Vtbl struct {
	win32.IInspectableVtbl
	Get_ViewMode                         uintptr
	IsViewModeSupported                  uintptr
	TryEnterViewModeAsync                uintptr
	TryEnterViewModeWithPreferencesAsync uintptr
	TryConsolidateAsync                  uintptr
}

type IApplicationView4 struct {
	win32.IInspectable
}

func (this *IApplicationView4) Vtbl() *IApplicationView4Vtbl {
	return (*IApplicationView4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationView4) Get_ViewMode() ApplicationViewMode {
	var _result ApplicationViewMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ViewMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationView4) IsViewModeSupported(viewMode ApplicationViewMode) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsViewModeSupported, uintptr(unsafe.Pointer(this)), uintptr(viewMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationView4) TryEnterViewModeAsync(viewMode ApplicationViewMode) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryEnterViewModeAsync, uintptr(unsafe.Pointer(this)), uintptr(viewMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationView4) TryEnterViewModeWithPreferencesAsync(viewMode ApplicationViewMode, viewModePreferences *IViewModePreferences) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryEnterViewModeWithPreferencesAsync, uintptr(unsafe.Pointer(this)), uintptr(viewMode), uintptr(unsafe.Pointer(viewModePreferences)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationView4) TryConsolidateAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryConsolidateAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A0369647-5FAF-5AA6-9C38-BEFBB12A071E
var IID_IApplicationView7 = syscall.GUID{0xA0369647, 0x5FAF, 0x5AA6,
	[8]byte{0x9C, 0x38, 0xBE, 0xFB, 0xB1, 0x2A, 0x07, 0x1E}}

type IApplicationView7Interface interface {
	win32.IInspectableInterface
	Get_PersistedStateId() string
	Put_PersistedStateId(value string)
}

type IApplicationView7Vtbl struct {
	win32.IInspectableVtbl
	Get_PersistedStateId uintptr
	Put_PersistedStateId uintptr
}

type IApplicationView7 struct {
	win32.IInspectable
}

func (this *IApplicationView7) Vtbl() *IApplicationView7Vtbl {
	return (*IApplicationView7Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationView7) Get_PersistedStateId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PersistedStateId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IApplicationView7) Put_PersistedStateId(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PersistedStateId, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 9C6516F9-021A-5F01-93E5-9BDAD2647574
var IID_IApplicationView9 = syscall.GUID{0x9C6516F9, 0x021A, 0x5F01,
	[8]byte{0x93, 0xE5, 0x9B, 0xDA, 0xD2, 0x64, 0x75, 0x74}}

type IApplicationView9Interface interface {
	win32.IInspectableInterface
	Get_WindowingEnvironment() unsafe.Pointer
	GetDisplayRegions() *IVectorView[unsafe.Pointer]
}

type IApplicationView9Vtbl struct {
	win32.IInspectableVtbl
	Get_WindowingEnvironment uintptr
	GetDisplayRegions        uintptr
}

type IApplicationView9 struct {
	win32.IInspectable
}

func (this *IApplicationView9) Vtbl() *IApplicationView9Vtbl {
	return (*IApplicationView9Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationView9) Get_WindowingEnvironment() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WindowingEnvironment, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationView9) GetDisplayRegions() *IVectorView[unsafe.Pointer] {
	var _result *IVectorView[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDisplayRegions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 514449EC-7EA2-4DE7-A6A6-7DFBAAEBB6FB
var IID_IApplicationViewConsolidatedEventArgs = syscall.GUID{0x514449EC, 0x7EA2, 0x4DE7,
	[8]byte{0xA6, 0xA6, 0x7D, 0xFB, 0xAA, 0xEB, 0xB6, 0xFB}}

type IApplicationViewConsolidatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_IsUserInitiated() bool
}

type IApplicationViewConsolidatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_IsUserInitiated uintptr
}

type IApplicationViewConsolidatedEventArgs struct {
	win32.IInspectable
}

func (this *IApplicationViewConsolidatedEventArgs) Vtbl() *IApplicationViewConsolidatedEventArgsVtbl {
	return (*IApplicationViewConsolidatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationViewConsolidatedEventArgs) Get_IsUserInitiated() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsUserInitiated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1C199ECC-6DC1-40F4-AFEE-07D9EA296430
var IID_IApplicationViewConsolidatedEventArgs2 = syscall.GUID{0x1C199ECC, 0x6DC1, 0x40F4,
	[8]byte{0xAF, 0xEE, 0x07, 0xD9, 0xEA, 0x29, 0x64, 0x30}}

type IApplicationViewConsolidatedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_IsAppInitiated() bool
}

type IApplicationViewConsolidatedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsAppInitiated uintptr
}

type IApplicationViewConsolidatedEventArgs2 struct {
	win32.IInspectable
}

func (this *IApplicationViewConsolidatedEventArgs2) Vtbl() *IApplicationViewConsolidatedEventArgs2Vtbl {
	return (*IApplicationViewConsolidatedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationViewConsolidatedEventArgs2) Get_IsAppInitiated() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAppInitiated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// BC792EBD-64FE-4B65-A0C0-901CE2B68636
var IID_IApplicationViewFullscreenStatics = syscall.GUID{0xBC792EBD, 0x64FE, 0x4B65,
	[8]byte{0xA0, 0xC0, 0x90, 0x1C, 0xE2, 0xB6, 0x86, 0x36}}

type IApplicationViewFullscreenStaticsInterface interface {
	win32.IInspectableInterface
	TryUnsnapToFullscreen() bool
}

type IApplicationViewFullscreenStaticsVtbl struct {
	win32.IInspectableVtbl
	TryUnsnapToFullscreen uintptr
}

type IApplicationViewFullscreenStatics struct {
	win32.IInspectable
}

func (this *IApplicationViewFullscreenStatics) Vtbl() *IApplicationViewFullscreenStaticsVtbl {
	return (*IApplicationViewFullscreenStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationViewFullscreenStatics) TryUnsnapToFullscreen() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryUnsnapToFullscreen, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C446FB5D-4793-4896-A8E2-BE57A8BB0F50
var IID_IApplicationViewInteropStatics = syscall.GUID{0xC446FB5D, 0x4793, 0x4896,
	[8]byte{0xA8, 0xE2, 0xBE, 0x57, 0xA8, 0xBB, 0x0F, 0x50}}

type IApplicationViewInteropStaticsInterface interface {
	win32.IInspectableInterface
	GetApplicationViewIdForWindow(window unsafe.Pointer) int32
}

type IApplicationViewInteropStaticsVtbl struct {
	win32.IInspectableVtbl
	GetApplicationViewIdForWindow uintptr
}

type IApplicationViewInteropStatics struct {
	win32.IInspectable
}

func (this *IApplicationViewInteropStatics) Vtbl() *IApplicationViewInteropStaticsVtbl {
	return (*IApplicationViewInteropStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationViewInteropStatics) GetApplicationViewIdForWindow(window unsafe.Pointer) int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetApplicationViewIdForWindow, uintptr(unsafe.Pointer(this)), uintptr(window), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1D0DDC23-23F3-4B2D-84FE-74BF37B48B66
var IID_IApplicationViewScaling = syscall.GUID{0x1D0DDC23, 0x23F3, 0x4B2D,
	[8]byte{0x84, 0xFE, 0x74, 0xBF, 0x37, 0xB4, 0x8B, 0x66}}

type IApplicationViewScalingInterface interface {
	win32.IInspectableInterface
}

type IApplicationViewScalingVtbl struct {
	win32.IInspectableVtbl
}

type IApplicationViewScaling struct {
	win32.IInspectable
}

func (this *IApplicationViewScaling) Vtbl() *IApplicationViewScalingVtbl {
	return (*IApplicationViewScalingVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// B08FECF0-B946-45C8-A5E3-71F5AA78F861
var IID_IApplicationViewScalingStatics = syscall.GUID{0xB08FECF0, 0xB946, 0x45C8,
	[8]byte{0xA5, 0xE3, 0x71, 0xF5, 0xAA, 0x78, 0xF8, 0x61}}

type IApplicationViewScalingStaticsInterface interface {
	win32.IInspectableInterface
	Get_DisableLayoutScaling() bool
	TrySetDisableLayoutScaling(disableLayoutScaling bool) bool
}

type IApplicationViewScalingStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_DisableLayoutScaling   uintptr
	TrySetDisableLayoutScaling uintptr
}

type IApplicationViewScalingStatics struct {
	win32.IInspectable
}

func (this *IApplicationViewScalingStatics) Vtbl() *IApplicationViewScalingStaticsVtbl {
	return (*IApplicationViewScalingStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationViewScalingStatics) Get_DisableLayoutScaling() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisableLayoutScaling, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationViewScalingStatics) TrySetDisableLayoutScaling(disableLayoutScaling bool) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySetDisableLayoutScaling, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&disableLayoutScaling))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 010A6306-C433-44E5-A9F2-BD84D4030A95
var IID_IApplicationViewStatics = syscall.GUID{0x010A6306, 0xC433, 0x44E5,
	[8]byte{0xA9, 0xF2, 0xBD, 0x84, 0xD4, 0x03, 0x0A, 0x95}}

type IApplicationViewStaticsInterface interface {
	win32.IInspectableInterface
	Get_Value() ApplicationViewState
	TryUnsnap() bool
}

type IApplicationViewStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Value uintptr
	TryUnsnap uintptr
}

type IApplicationViewStatics struct {
	win32.IInspectable
}

func (this *IApplicationViewStatics) Vtbl() *IApplicationViewStaticsVtbl {
	return (*IApplicationViewStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationViewStatics) Get_Value() ApplicationViewState {
	var _result ApplicationViewState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationViewStatics) TryUnsnap() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryUnsnap, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// AF338AE5-CF64-423C-85E5-F3E72448FB23
var IID_IApplicationViewStatics2 = syscall.GUID{0xAF338AE5, 0xCF64, 0x423C,
	[8]byte{0x85, 0xE5, 0xF3, 0xE7, 0x24, 0x48, 0xFB, 0x23}}

type IApplicationViewStatics2Interface interface {
	win32.IInspectableInterface
	GetForCurrentView() *IApplicationView
	Get_TerminateAppOnFinalViewClose() bool
	Put_TerminateAppOnFinalViewClose(value bool)
}

type IApplicationViewStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetForCurrentView                uintptr
	Get_TerminateAppOnFinalViewClose uintptr
	Put_TerminateAppOnFinalViewClose uintptr
}

type IApplicationViewStatics2 struct {
	win32.IInspectable
}

func (this *IApplicationViewStatics2) Vtbl() *IApplicationViewStatics2Vtbl {
	return (*IApplicationViewStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationViewStatics2) GetForCurrentView() *IApplicationView {
	var _result *IApplicationView
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationViewStatics2) Get_TerminateAppOnFinalViewClose() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TerminateAppOnFinalViewClose, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationViewStatics2) Put_TerminateAppOnFinalViewClose(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TerminateAppOnFinalViewClose, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// A28D7594-8C41-4E13-9719-5164796FE4C7
var IID_IApplicationViewStatics3 = syscall.GUID{0xA28D7594, 0x8C41, 0x4E13,
	[8]byte{0x97, 0x19, 0x51, 0x64, 0x79, 0x6F, 0xE4, 0xC7}}

type IApplicationViewStatics3Interface interface {
	win32.IInspectableInterface
	Get_PreferredLaunchWindowingMode() ApplicationViewWindowingMode
	Put_PreferredLaunchWindowingMode(value ApplicationViewWindowingMode)
	Get_PreferredLaunchViewSize() Size
	Put_PreferredLaunchViewSize(value Size)
}

type IApplicationViewStatics3Vtbl struct {
	win32.IInspectableVtbl
	Get_PreferredLaunchWindowingMode uintptr
	Put_PreferredLaunchWindowingMode uintptr
	Get_PreferredLaunchViewSize      uintptr
	Put_PreferredLaunchViewSize      uintptr
}

type IApplicationViewStatics3 struct {
	win32.IInspectable
}

func (this *IApplicationViewStatics3) Vtbl() *IApplicationViewStatics3Vtbl {
	return (*IApplicationViewStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationViewStatics3) Get_PreferredLaunchWindowingMode() ApplicationViewWindowingMode {
	var _result ApplicationViewWindowingMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PreferredLaunchWindowingMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationViewStatics3) Put_PreferredLaunchWindowingMode(value ApplicationViewWindowingMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PreferredLaunchWindowingMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IApplicationViewStatics3) Get_PreferredLaunchViewSize() Size {
	var _result Size
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PreferredLaunchViewSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationViewStatics3) Put_PreferredLaunchViewSize(value Size) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PreferredLaunchViewSize, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

// 08FD8D33-2611-5336-A315-D98E6366C9DB
var IID_IApplicationViewStatics4 = syscall.GUID{0x08FD8D33, 0x2611, 0x5336,
	[8]byte{0xA3, 0x15, 0xD9, 0x8E, 0x63, 0x66, 0xC9, 0xDB}}

type IApplicationViewStatics4Interface interface {
	win32.IInspectableInterface
	ClearAllPersistedState()
	ClearPersistedState(key string)
}

type IApplicationViewStatics4Vtbl struct {
	win32.IInspectableVtbl
	ClearAllPersistedState uintptr
	ClearPersistedState    uintptr
}

type IApplicationViewStatics4 struct {
	win32.IInspectable
}

func (this *IApplicationViewStatics4) Vtbl() *IApplicationViewStatics4Vtbl {
	return (*IApplicationViewStatics4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationViewStatics4) ClearAllPersistedState() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearAllPersistedState, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IApplicationViewStatics4) ClearPersistedState(key string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearPersistedState, uintptr(unsafe.Pointer(this)), NewHStr(key).Ptr)
	_ = _hr
}

// 975F2F1E-E656-4C5E-A0A1-717C6FFA7D64
var IID_IApplicationViewSwitcherStatics = syscall.GUID{0x975F2F1E, 0xE656, 0x4C5E,
	[8]byte{0xA0, 0xA1, 0x71, 0x7C, 0x6F, 0xFA, 0x7D, 0x64}}

type IApplicationViewSwitcherStaticsInterface interface {
	win32.IInspectableInterface
	DisableShowingMainViewOnActivation()
	TryShowAsStandaloneAsync(viewId int32) *IAsyncOperation[bool]
	TryShowAsStandaloneWithSizePreferenceAsync(viewId int32, sizePreference ViewSizePreference) *IAsyncOperation[bool]
	TryShowAsStandaloneWithAnchorViewAndSizePreferenceAsync(viewId int32, sizePreference ViewSizePreference, anchorViewId int32, anchorSizePreference ViewSizePreference) *IAsyncOperation[bool]
	SwitchAsync(viewId int32) *IAsyncAction
	SwitchFromViewAsync(toViewId int32, fromViewId int32) *IAsyncAction
	SwitchFromViewWithOptionsAsync(toViewId int32, fromViewId int32, options ApplicationViewSwitchingOptions) *IAsyncAction
	PrepareForCustomAnimatedSwitchAsync(toViewId int32, fromViewId int32, options ApplicationViewSwitchingOptions) *IAsyncOperation[bool]
}

type IApplicationViewSwitcherStaticsVtbl struct {
	win32.IInspectableVtbl
	DisableShowingMainViewOnActivation                      uintptr
	TryShowAsStandaloneAsync                                uintptr
	TryShowAsStandaloneWithSizePreferenceAsync              uintptr
	TryShowAsStandaloneWithAnchorViewAndSizePreferenceAsync uintptr
	SwitchAsync                                             uintptr
	SwitchFromViewAsync                                     uintptr
	SwitchFromViewWithOptionsAsync                          uintptr
	PrepareForCustomAnimatedSwitchAsync                     uintptr
}

type IApplicationViewSwitcherStatics struct {
	win32.IInspectable
}

func (this *IApplicationViewSwitcherStatics) Vtbl() *IApplicationViewSwitcherStaticsVtbl {
	return (*IApplicationViewSwitcherStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationViewSwitcherStatics) DisableShowingMainViewOnActivation() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DisableShowingMainViewOnActivation, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IApplicationViewSwitcherStatics) TryShowAsStandaloneAsync(viewId int32) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryShowAsStandaloneAsync, uintptr(unsafe.Pointer(this)), uintptr(viewId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationViewSwitcherStatics) TryShowAsStandaloneWithSizePreferenceAsync(viewId int32, sizePreference ViewSizePreference) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryShowAsStandaloneWithSizePreferenceAsync, uintptr(unsafe.Pointer(this)), uintptr(viewId), uintptr(sizePreference), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationViewSwitcherStatics) TryShowAsStandaloneWithAnchorViewAndSizePreferenceAsync(viewId int32, sizePreference ViewSizePreference, anchorViewId int32, anchorSizePreference ViewSizePreference) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryShowAsStandaloneWithAnchorViewAndSizePreferenceAsync, uintptr(unsafe.Pointer(this)), uintptr(viewId), uintptr(sizePreference), uintptr(anchorViewId), uintptr(anchorSizePreference), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationViewSwitcherStatics) SwitchAsync(viewId int32) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SwitchAsync, uintptr(unsafe.Pointer(this)), uintptr(viewId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationViewSwitcherStatics) SwitchFromViewAsync(toViewId int32, fromViewId int32) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SwitchFromViewAsync, uintptr(unsafe.Pointer(this)), uintptr(toViewId), uintptr(fromViewId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationViewSwitcherStatics) SwitchFromViewWithOptionsAsync(toViewId int32, fromViewId int32, options ApplicationViewSwitchingOptions) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SwitchFromViewWithOptionsAsync, uintptr(unsafe.Pointer(this)), uintptr(toViewId), uintptr(fromViewId), uintptr(options), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationViewSwitcherStatics) PrepareForCustomAnimatedSwitchAsync(toViewId int32, fromViewId int32, options ApplicationViewSwitchingOptions) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PrepareForCustomAnimatedSwitchAsync, uintptr(unsafe.Pointer(this)), uintptr(toViewId), uintptr(fromViewId), uintptr(options), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 60E995CD-4FC2-48C4-B8E3-395F2B9F0FC1
var IID_IApplicationViewSwitcherStatics2 = syscall.GUID{0x60E995CD, 0x4FC2, 0x48C4,
	[8]byte{0xB8, 0xE3, 0x39, 0x5F, 0x2B, 0x9F, 0x0F, 0xC1}}

type IApplicationViewSwitcherStatics2Interface interface {
	win32.IInspectableInterface
	DisableSystemViewActivationPolicy()
}

type IApplicationViewSwitcherStatics2Vtbl struct {
	win32.IInspectableVtbl
	DisableSystemViewActivationPolicy uintptr
}

type IApplicationViewSwitcherStatics2 struct {
	win32.IInspectable
}

func (this *IApplicationViewSwitcherStatics2) Vtbl() *IApplicationViewSwitcherStatics2Vtbl {
	return (*IApplicationViewSwitcherStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationViewSwitcherStatics2) DisableSystemViewActivationPolicy() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DisableSystemViewActivationPolicy, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 92059420-80A7-486D-B21F-C7A4A242A383
var IID_IApplicationViewSwitcherStatics3 = syscall.GUID{0x92059420, 0x80A7, 0x486D,
	[8]byte{0xB2, 0x1F, 0xC7, 0xA4, 0xA2, 0x42, 0xA3, 0x83}}

type IApplicationViewSwitcherStatics3Interface interface {
	win32.IInspectableInterface
	TryShowAsViewModeAsync(viewId int32, viewMode ApplicationViewMode) *IAsyncOperation[bool]
	TryShowAsViewModeWithPreferencesAsync(viewId int32, viewMode ApplicationViewMode, viewModePreferences *IViewModePreferences) *IAsyncOperation[bool]
}

type IApplicationViewSwitcherStatics3Vtbl struct {
	win32.IInspectableVtbl
	TryShowAsViewModeAsync                uintptr
	TryShowAsViewModeWithPreferencesAsync uintptr
}

type IApplicationViewSwitcherStatics3 struct {
	win32.IInspectable
}

func (this *IApplicationViewSwitcherStatics3) Vtbl() *IApplicationViewSwitcherStatics3Vtbl {
	return (*IApplicationViewSwitcherStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationViewSwitcherStatics3) TryShowAsViewModeAsync(viewId int32, viewMode ApplicationViewMode) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryShowAsViewModeAsync, uintptr(unsafe.Pointer(this)), uintptr(viewId), uintptr(viewMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationViewSwitcherStatics3) TryShowAsViewModeWithPreferencesAsync(viewId int32, viewMode ApplicationViewMode, viewModePreferences *IViewModePreferences) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryShowAsViewModeWithPreferencesAsync, uintptr(unsafe.Pointer(this)), uintptr(viewId), uintptr(viewMode), uintptr(unsafe.Pointer(viewModePreferences)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 00924AC0-932B-4A6B-9C4B-DC38C82478CE
var IID_IApplicationViewTitleBar = syscall.GUID{0x00924AC0, 0x932B, 0x4A6B,
	[8]byte{0x9C, 0x4B, 0xDC, 0x38, 0xC8, 0x24, 0x78, 0xCE}}

type IApplicationViewTitleBarInterface interface {
	win32.IInspectableInterface
	Put_ForegroundColor(value *IReference[unsafe.Pointer])
	Get_ForegroundColor() *IReference[unsafe.Pointer]
	Put_BackgroundColor(value *IReference[unsafe.Pointer])
	Get_BackgroundColor() *IReference[unsafe.Pointer]
	Put_ButtonForegroundColor(value *IReference[unsafe.Pointer])
	Get_ButtonForegroundColor() *IReference[unsafe.Pointer]
	Put_ButtonBackgroundColor(value *IReference[unsafe.Pointer])
	Get_ButtonBackgroundColor() *IReference[unsafe.Pointer]
	Put_ButtonHoverForegroundColor(value *IReference[unsafe.Pointer])
	Get_ButtonHoverForegroundColor() *IReference[unsafe.Pointer]
	Put_ButtonHoverBackgroundColor(value *IReference[unsafe.Pointer])
	Get_ButtonHoverBackgroundColor() *IReference[unsafe.Pointer]
	Put_ButtonPressedForegroundColor(value *IReference[unsafe.Pointer])
	Get_ButtonPressedForegroundColor() *IReference[unsafe.Pointer]
	Put_ButtonPressedBackgroundColor(value *IReference[unsafe.Pointer])
	Get_ButtonPressedBackgroundColor() *IReference[unsafe.Pointer]
	Put_InactiveForegroundColor(value *IReference[unsafe.Pointer])
	Get_InactiveForegroundColor() *IReference[unsafe.Pointer]
	Put_InactiveBackgroundColor(value *IReference[unsafe.Pointer])
	Get_InactiveBackgroundColor() *IReference[unsafe.Pointer]
	Put_ButtonInactiveForegroundColor(value *IReference[unsafe.Pointer])
	Get_ButtonInactiveForegroundColor() *IReference[unsafe.Pointer]
	Put_ButtonInactiveBackgroundColor(value *IReference[unsafe.Pointer])
	Get_ButtonInactiveBackgroundColor() *IReference[unsafe.Pointer]
}

type IApplicationViewTitleBarVtbl struct {
	win32.IInspectableVtbl
	Put_ForegroundColor               uintptr
	Get_ForegroundColor               uintptr
	Put_BackgroundColor               uintptr
	Get_BackgroundColor               uintptr
	Put_ButtonForegroundColor         uintptr
	Get_ButtonForegroundColor         uintptr
	Put_ButtonBackgroundColor         uintptr
	Get_ButtonBackgroundColor         uintptr
	Put_ButtonHoverForegroundColor    uintptr
	Get_ButtonHoverForegroundColor    uintptr
	Put_ButtonHoverBackgroundColor    uintptr
	Get_ButtonHoverBackgroundColor    uintptr
	Put_ButtonPressedForegroundColor  uintptr
	Get_ButtonPressedForegroundColor  uintptr
	Put_ButtonPressedBackgroundColor  uintptr
	Get_ButtonPressedBackgroundColor  uintptr
	Put_InactiveForegroundColor       uintptr
	Get_InactiveForegroundColor       uintptr
	Put_InactiveBackgroundColor       uintptr
	Get_InactiveBackgroundColor       uintptr
	Put_ButtonInactiveForegroundColor uintptr
	Get_ButtonInactiveForegroundColor uintptr
	Put_ButtonInactiveBackgroundColor uintptr
	Get_ButtonInactiveBackgroundColor uintptr
}

type IApplicationViewTitleBar struct {
	win32.IInspectable
}

func (this *IApplicationViewTitleBar) Vtbl() *IApplicationViewTitleBarVtbl {
	return (*IApplicationViewTitleBarVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationViewTitleBar) Put_ForegroundColor(value *IReference[unsafe.Pointer]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ForegroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IApplicationViewTitleBar) Get_ForegroundColor() *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ForegroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationViewTitleBar) Put_BackgroundColor(value *IReference[unsafe.Pointer]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BackgroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IApplicationViewTitleBar) Get_BackgroundColor() *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BackgroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationViewTitleBar) Put_ButtonForegroundColor(value *IReference[unsafe.Pointer]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ButtonForegroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IApplicationViewTitleBar) Get_ButtonForegroundColor() *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ButtonForegroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationViewTitleBar) Put_ButtonBackgroundColor(value *IReference[unsafe.Pointer]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ButtonBackgroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IApplicationViewTitleBar) Get_ButtonBackgroundColor() *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ButtonBackgroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationViewTitleBar) Put_ButtonHoverForegroundColor(value *IReference[unsafe.Pointer]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ButtonHoverForegroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IApplicationViewTitleBar) Get_ButtonHoverForegroundColor() *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ButtonHoverForegroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationViewTitleBar) Put_ButtonHoverBackgroundColor(value *IReference[unsafe.Pointer]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ButtonHoverBackgroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IApplicationViewTitleBar) Get_ButtonHoverBackgroundColor() *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ButtonHoverBackgroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationViewTitleBar) Put_ButtonPressedForegroundColor(value *IReference[unsafe.Pointer]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ButtonPressedForegroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IApplicationViewTitleBar) Get_ButtonPressedForegroundColor() *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ButtonPressedForegroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationViewTitleBar) Put_ButtonPressedBackgroundColor(value *IReference[unsafe.Pointer]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ButtonPressedBackgroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IApplicationViewTitleBar) Get_ButtonPressedBackgroundColor() *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ButtonPressedBackgroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationViewTitleBar) Put_InactiveForegroundColor(value *IReference[unsafe.Pointer]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InactiveForegroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IApplicationViewTitleBar) Get_InactiveForegroundColor() *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InactiveForegroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationViewTitleBar) Put_InactiveBackgroundColor(value *IReference[unsafe.Pointer]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InactiveBackgroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IApplicationViewTitleBar) Get_InactiveBackgroundColor() *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InactiveBackgroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationViewTitleBar) Put_ButtonInactiveForegroundColor(value *IReference[unsafe.Pointer]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ButtonInactiveForegroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IApplicationViewTitleBar) Get_ButtonInactiveForegroundColor() *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ButtonInactiveForegroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationViewTitleBar) Put_ButtonInactiveBackgroundColor(value *IReference[unsafe.Pointer]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ButtonInactiveBackgroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IApplicationViewTitleBar) Get_ButtonInactiveBackgroundColor() *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ButtonInactiveBackgroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8574BC63-3C17-408E-9408-8A1A9EA81BFA
var IID_IApplicationViewTransferContext = syscall.GUID{0x8574BC63, 0x3C17, 0x408E,
	[8]byte{0x94, 0x08, 0x8A, 0x1A, 0x9E, 0xA8, 0x1B, 0xFA}}

type IApplicationViewTransferContextInterface interface {
	win32.IInspectableInterface
	Get_ViewId() int32
	Put_ViewId(value int32)
}

type IApplicationViewTransferContextVtbl struct {
	win32.IInspectableVtbl
	Get_ViewId uintptr
	Put_ViewId uintptr
}

type IApplicationViewTransferContext struct {
	win32.IInspectable
}

func (this *IApplicationViewTransferContext) Vtbl() *IApplicationViewTransferContextVtbl {
	return (*IApplicationViewTransferContextVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationViewTransferContext) Get_ViewId() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ViewId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationViewTransferContext) Put_ViewId(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ViewId, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 15A89D92-DD79-4B0B-BC47-D5F195F14661
var IID_IApplicationViewTransferContextStatics = syscall.GUID{0x15A89D92, 0xDD79, 0x4B0B,
	[8]byte{0xBC, 0x47, 0xD5, 0xF1, 0x95, 0xF1, 0x46, 0x61}}

type IApplicationViewTransferContextStaticsInterface interface {
	win32.IInspectableInterface
	Get_DataPackageFormatId() string
}

type IApplicationViewTransferContextStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_DataPackageFormatId uintptr
}

type IApplicationViewTransferContextStatics struct {
	win32.IInspectable
}

func (this *IApplicationViewTransferContextStatics) Vtbl() *IApplicationViewTransferContextStaticsVtbl {
	return (*IApplicationViewTransferContextStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationViewTransferContextStatics) Get_DataPackageFormatId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DataPackageFormatId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// BD55D512-9DC1-44FC-8501-666625DF60DC
var IID_IApplicationViewWithContext = syscall.GUID{0xBD55D512, 0x9DC1, 0x44FC,
	[8]byte{0x85, 0x01, 0x66, 0x66, 0x25, 0xDF, 0x60, 0xDC}}

type IApplicationViewWithContextInterface interface {
	win32.IInspectableInterface
	Get_UIContext() unsafe.Pointer
}

type IApplicationViewWithContextVtbl struct {
	win32.IInspectableVtbl
	Get_UIContext uintptr
}

type IApplicationViewWithContext struct {
	win32.IInspectable
}

func (this *IApplicationViewWithContext) Vtbl() *IApplicationViewWithContextVtbl {
	return (*IApplicationViewWithContextVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationViewWithContext) Get_UIContext() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UIContext, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 640ADA70-06F3-4C87-A678-9829C9127C28
var IID_IInputPane = syscall.GUID{0x640ADA70, 0x06F3, 0x4C87,
	[8]byte{0xA6, 0x78, 0x98, 0x29, 0xC9, 0x12, 0x7C, 0x28}}

type IInputPaneInterface interface {
	win32.IInspectableInterface
	Add_Showing(handler TypedEventHandler[*IInputPane, *IInputPaneVisibilityEventArgs]) EventRegistrationToken
	Remove_Showing(token EventRegistrationToken)
	Add_Hiding(handler TypedEventHandler[*IInputPane, *IInputPaneVisibilityEventArgs]) EventRegistrationToken
	Remove_Hiding(token EventRegistrationToken)
	Get_OccludedRect() Rect
}

type IInputPaneVtbl struct {
	win32.IInspectableVtbl
	Add_Showing      uintptr
	Remove_Showing   uintptr
	Add_Hiding       uintptr
	Remove_Hiding    uintptr
	Get_OccludedRect uintptr
}

type IInputPane struct {
	win32.IInspectable
}

func (this *IInputPane) Vtbl() *IInputPaneVtbl {
	return (*IInputPaneVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInputPane) Add_Showing(handler TypedEventHandler[*IInputPane, *IInputPaneVisibilityEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Showing, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInputPane) Remove_Showing(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Showing, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IInputPane) Add_Hiding(handler TypedEventHandler[*IInputPane, *IInputPaneVisibilityEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Hiding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInputPane) Remove_Hiding(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Hiding, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IInputPane) Get_OccludedRect() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OccludedRect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 8A6B3F26-7090-4793-944C-C3F2CDE26276
var IID_IInputPane2 = syscall.GUID{0x8A6B3F26, 0x7090, 0x4793,
	[8]byte{0x94, 0x4C, 0xC3, 0xF2, 0xCD, 0xE2, 0x62, 0x76}}

type IInputPane2Interface interface {
	win32.IInspectableInterface
	TryShow() bool
	TryHide() bool
}

type IInputPane2Vtbl struct {
	win32.IInspectableVtbl
	TryShow uintptr
	TryHide uintptr
}

type IInputPane2 struct {
	win32.IInspectable
}

func (this *IInputPane2) Vtbl() *IInputPane2Vtbl {
	return (*IInputPane2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInputPane2) TryShow() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryShow, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInputPane2) TryHide() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryHide, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 088BB24F-962F-489D-AA6E-C6BE1A0A6E52
var IID_IInputPaneControl = syscall.GUID{0x088BB24F, 0x962F, 0x489D,
	[8]byte{0xAA, 0x6E, 0xC6, 0xBE, 0x1A, 0x0A, 0x6E, 0x52}}

type IInputPaneControlInterface interface {
	win32.IInspectableInterface
	Get_Visible() bool
	Put_Visible(value bool)
}

type IInputPaneControlVtbl struct {
	win32.IInspectableVtbl
	Get_Visible uintptr
	Put_Visible uintptr
}

type IInputPaneControl struct {
	win32.IInspectable
}

func (this *IInputPaneControl) Vtbl() *IInputPaneControlVtbl {
	return (*IInputPaneControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInputPaneControl) Get_Visible() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Visible, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInputPaneControl) Put_Visible(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Visible, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 95F4AF3A-EF47-424A-9741-FD2815EBA2BD
var IID_IInputPaneStatics = syscall.GUID{0x95F4AF3A, 0xEF47, 0x424A,
	[8]byte{0x97, 0x41, 0xFD, 0x28, 0x15, 0xEB, 0xA2, 0xBD}}

type IInputPaneStaticsInterface interface {
	win32.IInspectableInterface
	GetForCurrentView() *IInputPane
}

type IInputPaneStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForCurrentView uintptr
}

type IInputPaneStatics struct {
	win32.IInspectable
}

func (this *IInputPaneStatics) Vtbl() *IInputPaneStaticsVtbl {
	return (*IInputPaneStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInputPaneStatics) GetForCurrentView() *IInputPane {
	var _result *IInputPane
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1B63529B-D9EC-4531-8445-71BAB9FB828E
var IID_IInputPaneStatics2 = syscall.GUID{0x1B63529B, 0xD9EC, 0x4531,
	[8]byte{0x84, 0x45, 0x71, 0xBA, 0xB9, 0xFB, 0x82, 0x8E}}

type IInputPaneStatics2Interface interface {
	win32.IInspectableInterface
	GetForUIContext(context unsafe.Pointer) *IInputPane
}

type IInputPaneStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetForUIContext uintptr
}

type IInputPaneStatics2 struct {
	win32.IInspectable
}

func (this *IInputPaneStatics2) Vtbl() *IInputPaneStatics2Vtbl {
	return (*IInputPaneStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInputPaneStatics2) GetForUIContext(context unsafe.Pointer) *IInputPane {
	var _result *IInputPane
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForUIContext, uintptr(unsafe.Pointer(this)), uintptr(context), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D243E016-D907-4FCC-BB8D-F77BAA5028F1
var IID_IInputPaneVisibilityEventArgs = syscall.GUID{0xD243E016, 0xD907, 0x4FCC,
	[8]byte{0xBB, 0x8D, 0xF7, 0x7B, 0xAA, 0x50, 0x28, 0xF1}}

type IInputPaneVisibilityEventArgsInterface interface {
	win32.IInspectableInterface
	Get_OccludedRect() Rect
	Put_EnsuredFocusedElementInView(value bool)
	Get_EnsuredFocusedElementInView() bool
}

type IInputPaneVisibilityEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_OccludedRect                uintptr
	Put_EnsuredFocusedElementInView uintptr
	Get_EnsuredFocusedElementInView uintptr
}

type IInputPaneVisibilityEventArgs struct {
	win32.IInspectable
}

func (this *IInputPaneVisibilityEventArgs) Vtbl() *IInputPaneVisibilityEventArgsVtbl {
	return (*IInputPaneVisibilityEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInputPaneVisibilityEventArgs) Get_OccludedRect() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OccludedRect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInputPaneVisibilityEventArgs) Put_EnsuredFocusedElementInView(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_EnsuredFocusedElementInView, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IInputPaneVisibilityEventArgs) Get_EnsuredFocusedElementInView() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EnsuredFocusedElementInView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// B65F913D-E2F0-4FFD-BA95-34241647E45C
var IID_IProjectionManagerStatics = syscall.GUID{0xB65F913D, 0xE2F0, 0x4FFD,
	[8]byte{0xBA, 0x95, 0x34, 0x24, 0x16, 0x47, 0xE4, 0x5C}}

type IProjectionManagerStaticsInterface interface {
	win32.IInspectableInterface
	StartProjectingAsync(projectionViewId int32, anchorViewId int32) *IAsyncAction
	SwapDisplaysForViewsAsync(projectionViewId int32, anchorViewId int32) *IAsyncAction
	StopProjectingAsync(projectionViewId int32, anchorViewId int32) *IAsyncAction
	Get_ProjectionDisplayAvailable() bool
	Add_ProjectionDisplayAvailableChanged(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_ProjectionDisplayAvailableChanged(token EventRegistrationToken)
}

type IProjectionManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	StartProjectingAsync                     uintptr
	SwapDisplaysForViewsAsync                uintptr
	StopProjectingAsync                      uintptr
	Get_ProjectionDisplayAvailable           uintptr
	Add_ProjectionDisplayAvailableChanged    uintptr
	Remove_ProjectionDisplayAvailableChanged uintptr
}

type IProjectionManagerStatics struct {
	win32.IInspectable
}

func (this *IProjectionManagerStatics) Vtbl() *IProjectionManagerStaticsVtbl {
	return (*IProjectionManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProjectionManagerStatics) StartProjectingAsync(projectionViewId int32, anchorViewId int32) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartProjectingAsync, uintptr(unsafe.Pointer(this)), uintptr(projectionViewId), uintptr(anchorViewId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProjectionManagerStatics) SwapDisplaysForViewsAsync(projectionViewId int32, anchorViewId int32) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SwapDisplaysForViewsAsync, uintptr(unsafe.Pointer(this)), uintptr(projectionViewId), uintptr(anchorViewId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProjectionManagerStatics) StopProjectingAsync(projectionViewId int32, anchorViewId int32) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopProjectingAsync, uintptr(unsafe.Pointer(this)), uintptr(projectionViewId), uintptr(anchorViewId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProjectionManagerStatics) Get_ProjectionDisplayAvailable() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProjectionDisplayAvailable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProjectionManagerStatics) Add_ProjectionDisplayAvailableChanged(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ProjectionDisplayAvailableChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProjectionManagerStatics) Remove_ProjectionDisplayAvailableChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ProjectionDisplayAvailableChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// F33D2F43-2749-4CDE-B977-C0C41E7415D1
var IID_IProjectionManagerStatics2 = syscall.GUID{0xF33D2F43, 0x2749, 0x4CDE,
	[8]byte{0xB9, 0x77, 0xC0, 0xC4, 0x1E, 0x74, 0x15, 0xD1}}

type IProjectionManagerStatics2Interface interface {
	win32.IInspectableInterface
	StartProjectingWithDeviceInfoAsync(projectionViewId int32, anchorViewId int32, displayDeviceInfo *IDeviceInformation) *IAsyncAction
	RequestStartProjectingAsync(projectionViewId int32, anchorViewId int32, selection Rect) *IAsyncOperation[bool]
	RequestStartProjectingWithPlacementAsync(projectionViewId int32, anchorViewId int32, selection Rect, prefferedPlacement Placement) *IAsyncOperation[bool]
	GetDeviceSelector() string
}

type IProjectionManagerStatics2Vtbl struct {
	win32.IInspectableVtbl
	StartProjectingWithDeviceInfoAsync       uintptr
	RequestStartProjectingAsync              uintptr
	RequestStartProjectingWithPlacementAsync uintptr
	GetDeviceSelector                        uintptr
}

type IProjectionManagerStatics2 struct {
	win32.IInspectable
}

func (this *IProjectionManagerStatics2) Vtbl() *IProjectionManagerStatics2Vtbl {
	return (*IProjectionManagerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProjectionManagerStatics2) StartProjectingWithDeviceInfoAsync(projectionViewId int32, anchorViewId int32, displayDeviceInfo *IDeviceInformation) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartProjectingWithDeviceInfoAsync, uintptr(unsafe.Pointer(this)), uintptr(projectionViewId), uintptr(anchorViewId), uintptr(unsafe.Pointer(displayDeviceInfo)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProjectionManagerStatics2) RequestStartProjectingAsync(projectionViewId int32, anchorViewId int32, selection Rect) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestStartProjectingAsync, uintptr(unsafe.Pointer(this)), uintptr(projectionViewId), uintptr(anchorViewId), uintptr(unsafe.Pointer(&selection)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProjectionManagerStatics2) RequestStartProjectingWithPlacementAsync(projectionViewId int32, anchorViewId int32, selection Rect, prefferedPlacement Placement) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestStartProjectingWithPlacementAsync, uintptr(unsafe.Pointer(this)), uintptr(projectionViewId), uintptr(anchorViewId), uintptr(unsafe.Pointer(&selection)), uintptr(prefferedPlacement), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProjectionManagerStatics2) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 0FFCC5BF-98D0-4864-B1E8-B3F4020BE8B4
var IID_IStatusBar = syscall.GUID{0x0FFCC5BF, 0x98D0, 0x4864,
	[8]byte{0xB1, 0xE8, 0xB3, 0xF4, 0x02, 0x0B, 0xE8, 0xB4}}

type IStatusBarInterface interface {
	win32.IInspectableInterface
	ShowAsync() *IAsyncAction
	HideAsync() *IAsyncAction
	Get_BackgroundOpacity() float64
	Put_BackgroundOpacity(value float64)
	Get_ForegroundColor() *IReference[unsafe.Pointer]
	Put_ForegroundColor(value *IReference[unsafe.Pointer])
	Get_BackgroundColor() *IReference[unsafe.Pointer]
	Put_BackgroundColor(value *IReference[unsafe.Pointer])
	Get_ProgressIndicator() *IStatusBarProgressIndicator
	Get_OccludedRect() Rect
	Add_Showing(eventHandler TypedEventHandler[*IStatusBar, interface{}]) EventRegistrationToken
	Remove_Showing(token EventRegistrationToken)
	Add_Hiding(eventHandler TypedEventHandler[*IStatusBar, interface{}]) EventRegistrationToken
	Remove_Hiding(token EventRegistrationToken)
}

type IStatusBarVtbl struct {
	win32.IInspectableVtbl
	ShowAsync             uintptr
	HideAsync             uintptr
	Get_BackgroundOpacity uintptr
	Put_BackgroundOpacity uintptr
	Get_ForegroundColor   uintptr
	Put_ForegroundColor   uintptr
	Get_BackgroundColor   uintptr
	Put_BackgroundColor   uintptr
	Get_ProgressIndicator uintptr
	Get_OccludedRect      uintptr
	Add_Showing           uintptr
	Remove_Showing        uintptr
	Add_Hiding            uintptr
	Remove_Hiding         uintptr
}

type IStatusBar struct {
	win32.IInspectable
}

func (this *IStatusBar) Vtbl() *IStatusBarVtbl {
	return (*IStatusBarVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStatusBar) ShowAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStatusBar) HideAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().HideAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStatusBar) Get_BackgroundOpacity() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BackgroundOpacity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStatusBar) Put_BackgroundOpacity(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BackgroundOpacity, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IStatusBar) Get_ForegroundColor() *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ForegroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStatusBar) Put_ForegroundColor(value *IReference[unsafe.Pointer]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ForegroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IStatusBar) Get_BackgroundColor() *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BackgroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStatusBar) Put_BackgroundColor(value *IReference[unsafe.Pointer]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BackgroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IStatusBar) Get_ProgressIndicator() *IStatusBarProgressIndicator {
	var _result *IStatusBarProgressIndicator
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProgressIndicator, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStatusBar) Get_OccludedRect() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OccludedRect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStatusBar) Add_Showing(eventHandler TypedEventHandler[*IStatusBar, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Showing, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(eventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStatusBar) Remove_Showing(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Showing, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IStatusBar) Add_Hiding(eventHandler TypedEventHandler[*IStatusBar, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Hiding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(eventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStatusBar) Remove_Hiding(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Hiding, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 76CB2670-A3D7-49CF-8200-4F3EEDCA27BB
var IID_IStatusBarProgressIndicator = syscall.GUID{0x76CB2670, 0xA3D7, 0x49CF,
	[8]byte{0x82, 0x00, 0x4F, 0x3E, 0xED, 0xCA, 0x27, 0xBB}}

type IStatusBarProgressIndicatorInterface interface {
	win32.IInspectableInterface
	ShowAsync() *IAsyncAction
	HideAsync() *IAsyncAction
	Get_Text() string
	Put_Text(value string)
	Get_ProgressValue() *IReference[float64]
	Put_ProgressValue(value *IReference[float64])
}

type IStatusBarProgressIndicatorVtbl struct {
	win32.IInspectableVtbl
	ShowAsync         uintptr
	HideAsync         uintptr
	Get_Text          uintptr
	Put_Text          uintptr
	Get_ProgressValue uintptr
	Put_ProgressValue uintptr
}

type IStatusBarProgressIndicator struct {
	win32.IInspectable
}

func (this *IStatusBarProgressIndicator) Vtbl() *IStatusBarProgressIndicatorVtbl {
	return (*IStatusBarProgressIndicatorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStatusBarProgressIndicator) ShowAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStatusBarProgressIndicator) HideAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().HideAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStatusBarProgressIndicator) Get_Text() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Text, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStatusBarProgressIndicator) Put_Text(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Text, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IStatusBarProgressIndicator) Get_ProgressValue() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProgressValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStatusBarProgressIndicator) Put_ProgressValue(value *IReference[float64]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProgressValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 8B463FDF-422F-4561-8806-FB1289CADFB7
var IID_IStatusBarStatics = syscall.GUID{0x8B463FDF, 0x422F, 0x4561,
	[8]byte{0x88, 0x06, 0xFB, 0x12, 0x89, 0xCA, 0xDF, 0xB7}}

type IStatusBarStaticsInterface interface {
	win32.IInspectableInterface
	GetForCurrentView() *IStatusBar
}

type IStatusBarStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForCurrentView uintptr
}

type IStatusBarStatics struct {
	win32.IInspectable
}

func (this *IStatusBarStatics) Vtbl() *IStatusBarStaticsVtbl {
	return (*IStatusBarStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStatusBarStatics) GetForCurrentView() *IStatusBar {
	var _result *IStatusBar
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 85361600-1C63-4627-BCB1-3A89E0BC9C55
var IID_IUISettings = syscall.GUID{0x85361600, 0x1C63, 0x4627,
	[8]byte{0xBC, 0xB1, 0x3A, 0x89, 0xE0, 0xBC, 0x9C, 0x55}}

type IUISettingsInterface interface {
	win32.IInspectableInterface
	Get_HandPreference() HandPreference
	Get_CursorSize() Size
	Get_ScrollBarSize() Size
	Get_ScrollBarArrowSize() Size
	Get_ScrollBarThumbBoxSize() Size
	Get_MessageDuration() uint32
	Get_AnimationsEnabled() bool
	Get_CaretBrowsingEnabled() bool
	Get_CaretBlinkRate() uint32
	Get_CaretWidth() uint32
	Get_DoubleClickTime() uint32
	Get_MouseHoverTime() uint32
	UIElementColor(desiredElement UIElementType) unsafe.Pointer
}

type IUISettingsVtbl struct {
	win32.IInspectableVtbl
	Get_HandPreference        uintptr
	Get_CursorSize            uintptr
	Get_ScrollBarSize         uintptr
	Get_ScrollBarArrowSize    uintptr
	Get_ScrollBarThumbBoxSize uintptr
	Get_MessageDuration       uintptr
	Get_AnimationsEnabled     uintptr
	Get_CaretBrowsingEnabled  uintptr
	Get_CaretBlinkRate        uintptr
	Get_CaretWidth            uintptr
	Get_DoubleClickTime       uintptr
	Get_MouseHoverTime        uintptr
	UIElementColor            uintptr
}

type IUISettings struct {
	win32.IInspectable
}

func (this *IUISettings) Vtbl() *IUISettingsVtbl {
	return (*IUISettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUISettings) Get_HandPreference() HandPreference {
	var _result HandPreference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HandPreference, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUISettings) Get_CursorSize() Size {
	var _result Size
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CursorSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUISettings) Get_ScrollBarSize() Size {
	var _result Size
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScrollBarSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUISettings) Get_ScrollBarArrowSize() Size {
	var _result Size
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScrollBarArrowSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUISettings) Get_ScrollBarThumbBoxSize() Size {
	var _result Size
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScrollBarThumbBoxSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUISettings) Get_MessageDuration() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MessageDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUISettings) Get_AnimationsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AnimationsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUISettings) Get_CaretBrowsingEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CaretBrowsingEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUISettings) Get_CaretBlinkRate() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CaretBlinkRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUISettings) Get_CaretWidth() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CaretWidth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUISettings) Get_DoubleClickTime() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DoubleClickTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUISettings) Get_MouseHoverTime() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MouseHoverTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUISettings) UIElementColor(desiredElement UIElementType) unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UIElementColor, uintptr(unsafe.Pointer(this)), uintptr(desiredElement), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// BAD82401-2721-44F9-BB91-2BB228BE442F
var IID_IUISettings2 = syscall.GUID{0xBAD82401, 0x2721, 0x44F9,
	[8]byte{0xBB, 0x91, 0x2B, 0xB2, 0x28, 0xBE, 0x44, 0x2F}}

type IUISettings2Interface interface {
	win32.IInspectableInterface
	Get_TextScaleFactor() float64
	Add_TextScaleFactorChanged(handler TypedEventHandler[*IUISettings, interface{}]) EventRegistrationToken
	Remove_TextScaleFactorChanged(cookie EventRegistrationToken)
}

type IUISettings2Vtbl struct {
	win32.IInspectableVtbl
	Get_TextScaleFactor           uintptr
	Add_TextScaleFactorChanged    uintptr
	Remove_TextScaleFactorChanged uintptr
}

type IUISettings2 struct {
	win32.IInspectable
}

func (this *IUISettings2) Vtbl() *IUISettings2Vtbl {
	return (*IUISettings2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUISettings2) Get_TextScaleFactor() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TextScaleFactor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUISettings2) Add_TextScaleFactorChanged(handler TypedEventHandler[*IUISettings, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_TextScaleFactorChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUISettings2) Remove_TextScaleFactorChanged(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_TextScaleFactorChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

// 03021BE4-5254-4781-8194-5168F7D06D7B
var IID_IUISettings3 = syscall.GUID{0x03021BE4, 0x5254, 0x4781,
	[8]byte{0x81, 0x94, 0x51, 0x68, 0xF7, 0xD0, 0x6D, 0x7B}}

type IUISettings3Interface interface {
	win32.IInspectableInterface
	GetColorValue(desiredColor UIColorType) unsafe.Pointer
	Add_ColorValuesChanged(handler TypedEventHandler[*IUISettings, interface{}]) EventRegistrationToken
	Remove_ColorValuesChanged(cookie EventRegistrationToken)
}

type IUISettings3Vtbl struct {
	win32.IInspectableVtbl
	GetColorValue             uintptr
	Add_ColorValuesChanged    uintptr
	Remove_ColorValuesChanged uintptr
}

type IUISettings3 struct {
	win32.IInspectable
}

func (this *IUISettings3) Vtbl() *IUISettings3Vtbl {
	return (*IUISettings3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUISettings3) GetColorValue(desiredColor UIColorType) unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetColorValue, uintptr(unsafe.Pointer(this)), uintptr(desiredColor), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUISettings3) Add_ColorValuesChanged(handler TypedEventHandler[*IUISettings, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ColorValuesChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUISettings3) Remove_ColorValuesChanged(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ColorValuesChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

// 52BB3002-919B-4D6B-9B78-8DD66FF4B93B
var IID_IUISettings4 = syscall.GUID{0x52BB3002, 0x919B, 0x4D6B,
	[8]byte{0x9B, 0x78, 0x8D, 0xD6, 0x6F, 0xF4, 0xB9, 0x3B}}

type IUISettings4Interface interface {
	win32.IInspectableInterface
	Get_AdvancedEffectsEnabled() bool
	Add_AdvancedEffectsEnabledChanged(handler TypedEventHandler[*IUISettings, interface{}]) EventRegistrationToken
	Remove_AdvancedEffectsEnabledChanged(cookie EventRegistrationToken)
}

type IUISettings4Vtbl struct {
	win32.IInspectableVtbl
	Get_AdvancedEffectsEnabled           uintptr
	Add_AdvancedEffectsEnabledChanged    uintptr
	Remove_AdvancedEffectsEnabledChanged uintptr
}

type IUISettings4 struct {
	win32.IInspectable
}

func (this *IUISettings4) Vtbl() *IUISettings4Vtbl {
	return (*IUISettings4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUISettings4) Get_AdvancedEffectsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AdvancedEffectsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUISettings4) Add_AdvancedEffectsEnabledChanged(handler TypedEventHandler[*IUISettings, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AdvancedEffectsEnabledChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUISettings4) Remove_AdvancedEffectsEnabledChanged(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AdvancedEffectsEnabledChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

// 5349D588-0CB5-5F05-BD34-706B3231F0BD
var IID_IUISettings5 = syscall.GUID{0x5349D588, 0x0CB5, 0x5F05,
	[8]byte{0xBD, 0x34, 0x70, 0x6B, 0x32, 0x31, 0xF0, 0xBD}}

type IUISettings5Interface interface {
	win32.IInspectableInterface
	Get_AutoHideScrollBars() bool
	Add_AutoHideScrollBarsChanged(handler TypedEventHandler[*IUISettings, *IUISettingsAutoHideScrollBarsChangedEventArgs]) EventRegistrationToken
	Remove_AutoHideScrollBarsChanged(token EventRegistrationToken)
}

type IUISettings5Vtbl struct {
	win32.IInspectableVtbl
	Get_AutoHideScrollBars           uintptr
	Add_AutoHideScrollBarsChanged    uintptr
	Remove_AutoHideScrollBarsChanged uintptr
}

type IUISettings5 struct {
	win32.IInspectable
}

func (this *IUISettings5) Vtbl() *IUISettings5Vtbl {
	return (*IUISettings5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUISettings5) Get_AutoHideScrollBars() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutoHideScrollBars, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUISettings5) Add_AutoHideScrollBarsChanged(handler TypedEventHandler[*IUISettings, *IUISettingsAutoHideScrollBarsChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AutoHideScrollBarsChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUISettings5) Remove_AutoHideScrollBarsChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AutoHideScrollBarsChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// AEF19BD7-FE31-5A04-ADA4-469AAEC6DFA9
var IID_IUISettings6 = syscall.GUID{0xAEF19BD7, 0xFE31, 0x5A04,
	[8]byte{0xAD, 0xA4, 0x46, 0x9A, 0xAE, 0xC6, 0xDF, 0xA9}}

type IUISettings6Interface interface {
	win32.IInspectableInterface
	Add_AnimationsEnabledChanged(handler TypedEventHandler[*IUISettings, *IUISettingsAnimationsEnabledChangedEventArgs]) EventRegistrationToken
	Remove_AnimationsEnabledChanged(token EventRegistrationToken)
	Add_MessageDurationChanged(handler TypedEventHandler[*IUISettings, *IUISettingsMessageDurationChangedEventArgs]) EventRegistrationToken
	Remove_MessageDurationChanged(token EventRegistrationToken)
}

type IUISettings6Vtbl struct {
	win32.IInspectableVtbl
	Add_AnimationsEnabledChanged    uintptr
	Remove_AnimationsEnabledChanged uintptr
	Add_MessageDurationChanged      uintptr
	Remove_MessageDurationChanged   uintptr
}

type IUISettings6 struct {
	win32.IInspectable
}

func (this *IUISettings6) Vtbl() *IUISettings6Vtbl {
	return (*IUISettings6Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUISettings6) Add_AnimationsEnabledChanged(handler TypedEventHandler[*IUISettings, *IUISettingsAnimationsEnabledChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AnimationsEnabledChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUISettings6) Remove_AnimationsEnabledChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AnimationsEnabledChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IUISettings6) Add_MessageDurationChanged(handler TypedEventHandler[*IUISettings, *IUISettingsMessageDurationChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_MessageDurationChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUISettings6) Remove_MessageDurationChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_MessageDurationChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 0C7B4B3D-2EA1-533E-894D-415BC5243C29
var IID_IUISettingsAnimationsEnabledChangedEventArgs = syscall.GUID{0x0C7B4B3D, 0x2EA1, 0x533E,
	[8]byte{0x89, 0x4D, 0x41, 0x5B, 0xC5, 0x24, 0x3C, 0x29}}

type IUISettingsAnimationsEnabledChangedEventArgsInterface interface {
	win32.IInspectableInterface
}

type IUISettingsAnimationsEnabledChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
}

type IUISettingsAnimationsEnabledChangedEventArgs struct {
	win32.IInspectable
}

func (this *IUISettingsAnimationsEnabledChangedEventArgs) Vtbl() *IUISettingsAnimationsEnabledChangedEventArgsVtbl {
	return (*IUISettingsAnimationsEnabledChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 87AFD4B2-9146-5F02-8F6B-06D454174C0F
var IID_IUISettingsAutoHideScrollBarsChangedEventArgs = syscall.GUID{0x87AFD4B2, 0x9146, 0x5F02,
	[8]byte{0x8F, 0x6B, 0x06, 0xD4, 0x54, 0x17, 0x4C, 0x0F}}

type IUISettingsAutoHideScrollBarsChangedEventArgsInterface interface {
	win32.IInspectableInterface
}

type IUISettingsAutoHideScrollBarsChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
}

type IUISettingsAutoHideScrollBarsChangedEventArgs struct {
	win32.IInspectable
}

func (this *IUISettingsAutoHideScrollBarsChangedEventArgs) Vtbl() *IUISettingsAutoHideScrollBarsChangedEventArgsVtbl {
	return (*IUISettingsAutoHideScrollBarsChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 338AAD52-4A5D-5B59-8002-D930F608FD6E
var IID_IUISettingsMessageDurationChangedEventArgs = syscall.GUID{0x338AAD52, 0x4A5D, 0x5B59,
	[8]byte{0x80, 0x02, 0xD9, 0x30, 0xF6, 0x08, 0xFD, 0x6E}}

type IUISettingsMessageDurationChangedEventArgsInterface interface {
	win32.IInspectableInterface
}

type IUISettingsMessageDurationChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
}

type IUISettingsMessageDurationChangedEventArgs struct {
	win32.IInspectable
}

func (this *IUISettingsMessageDurationChangedEventArgs) Vtbl() *IUISettingsMessageDurationChangedEventArgsVtbl {
	return (*IUISettingsMessageDurationChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// C63657F6-8850-470D-88F8-455E16EA2C26
var IID_IUIViewSettings = syscall.GUID{0xC63657F6, 0x8850, 0x470D,
	[8]byte{0x88, 0xF8, 0x45, 0x5E, 0x16, 0xEA, 0x2C, 0x26}}

type IUIViewSettingsInterface interface {
	win32.IInspectableInterface
	Get_UserInteractionMode() UserInteractionMode
}

type IUIViewSettingsVtbl struct {
	win32.IInspectableVtbl
	Get_UserInteractionMode uintptr
}

type IUIViewSettings struct {
	win32.IInspectable
}

func (this *IUIViewSettings) Vtbl() *IUIViewSettingsVtbl {
	return (*IUIViewSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIViewSettings) Get_UserInteractionMode() UserInteractionMode {
	var _result UserInteractionMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserInteractionMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 595C97A5-F8F6-41CF-B0FB-AACDB81FD5F6
var IID_IUIViewSettingsStatics = syscall.GUID{0x595C97A5, 0xF8F6, 0x41CF,
	[8]byte{0xB0, 0xFB, 0xAA, 0xCD, 0xB8, 0x1F, 0xD5, 0xF6}}

type IUIViewSettingsStaticsInterface interface {
	win32.IInspectableInterface
	GetForCurrentView() *IUIViewSettings
}

type IUIViewSettingsStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForCurrentView uintptr
}

type IUIViewSettingsStatics struct {
	win32.IInspectable
}

func (this *IUIViewSettingsStatics) Vtbl() *IUIViewSettingsStaticsVtbl {
	return (*IUIViewSettingsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIViewSettingsStatics) GetForCurrentView() *IUIViewSettings {
	var _result *IUIViewSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 878FCD3A-0B99-42C9-84D0-D3F1D403554B
var IID_IViewModePreferences = syscall.GUID{0x878FCD3A, 0x0B99, 0x42C9,
	[8]byte{0x84, 0xD0, 0xD3, 0xF1, 0xD4, 0x03, 0x55, 0x4B}}

type IViewModePreferencesInterface interface {
	win32.IInspectableInterface
	Get_ViewSizePreference() ViewSizePreference
	Put_ViewSizePreference(value ViewSizePreference)
	Get_CustomSize() Size
	Put_CustomSize(value Size)
}

type IViewModePreferencesVtbl struct {
	win32.IInspectableVtbl
	Get_ViewSizePreference uintptr
	Put_ViewSizePreference uintptr
	Get_CustomSize         uintptr
	Put_CustomSize         uintptr
}

type IViewModePreferences struct {
	win32.IInspectable
}

func (this *IViewModePreferences) Vtbl() *IViewModePreferencesVtbl {
	return (*IViewModePreferencesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IViewModePreferences) Get_ViewSizePreference() ViewSizePreference {
	var _result ViewSizePreference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ViewSizePreference, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IViewModePreferences) Put_ViewSizePreference(value ViewSizePreference) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ViewSizePreference, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IViewModePreferences) Get_CustomSize() Size {
	var _result Size
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CustomSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IViewModePreferences) Put_CustomSize(value Size) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CustomSize, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

// 69B60A65-5DE5-40D8-8306-3833DF7A2274
var IID_IViewModePreferencesStatics = syscall.GUID{0x69B60A65, 0x5DE5, 0x40D8,
	[8]byte{0x83, 0x06, 0x38, 0x33, 0xDF, 0x7A, 0x22, 0x74}}

type IViewModePreferencesStaticsInterface interface {
	win32.IInspectableInterface
	CreateDefault(mode ApplicationViewMode) *IViewModePreferences
}

type IViewModePreferencesStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateDefault uintptr
}

type IViewModePreferencesStatics struct {
	win32.IInspectable
}

func (this *IViewModePreferencesStatics) Vtbl() *IViewModePreferencesStaticsVtbl {
	return (*IViewModePreferencesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IViewModePreferencesStatics) CreateDefault(mode ApplicationViewMode) *IViewModePreferences {
	var _result *IViewModePreferences
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDefault, uintptr(unsafe.Pointer(this)), uintptr(mode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type InputPane struct {
	RtClass
	*IInputPane
}

func NewIInputPaneStatics2() *IInputPaneStatics2 {
	var p *IInputPaneStatics2
	hs := NewHStr("Windows.UI.ViewManagement.InputPane")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IInputPaneStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIInputPaneStatics() *IInputPaneStatics {
	var p *IInputPaneStatics
	hs := NewHStr("Windows.UI.ViewManagement.InputPane")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IInputPaneStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type UIViewSettings struct {
	RtClass
	*IUIViewSettings
}

func NewIUIViewSettingsStatics() *IUIViewSettingsStatics {
	var p *IUIViewSettingsStatics
	hs := NewHStr("Windows.UI.ViewManagement.UIViewSettings")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IUIViewSettingsStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
