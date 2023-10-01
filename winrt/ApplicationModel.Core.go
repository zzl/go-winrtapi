package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type AppRestartFailureReason int32

const (
	AppRestartFailureReason_RestartPending  AppRestartFailureReason = 0
	AppRestartFailureReason_NotInForeground AppRestartFailureReason = 1
	AppRestartFailureReason_InvalidUser     AppRestartFailureReason = 2
	AppRestartFailureReason_Other           AppRestartFailureReason = 3
)

// interfaces

// EF00F07F-2108-490A-877A-8A9F17C25FAD
var IID_IAppListEntry = syscall.GUID{0xEF00F07F, 0x2108, 0x490A,
	[8]byte{0x87, 0x7A, 0x8A, 0x9F, 0x17, 0xC2, 0x5F, 0xAD}}

type IAppListEntryInterface interface {
	win32.IInspectableInterface
	Get_DisplayInfo() *IAppDisplayInfo
	LaunchAsync() *IAsyncOperation[bool]
}

type IAppListEntryVtbl struct {
	win32.IInspectableVtbl
	Get_DisplayInfo uintptr
	LaunchAsync     uintptr
}

type IAppListEntry struct {
	win32.IInspectable
}

func (this *IAppListEntry) Vtbl() *IAppListEntryVtbl {
	return (*IAppListEntryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppListEntry) Get_DisplayInfo() *IAppDisplayInfo {
	var _result *IAppDisplayInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppListEntry) LaunchAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D0A618AD-BF35-42AC-AC06-86EEEB41D04B
var IID_IAppListEntry2 = syscall.GUID{0xD0A618AD, 0xBF35, 0x42AC,
	[8]byte{0xAC, 0x06, 0x86, 0xEE, 0xEB, 0x41, 0xD0, 0x4B}}

type IAppListEntry2Interface interface {
	win32.IInspectableInterface
	Get_AppUserModelId() string
}

type IAppListEntry2Vtbl struct {
	win32.IInspectableVtbl
	Get_AppUserModelId uintptr
}

type IAppListEntry2 struct {
	win32.IInspectable
}

func (this *IAppListEntry2) Vtbl() *IAppListEntry2Vtbl {
	return (*IAppListEntry2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppListEntry2) Get_AppUserModelId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppUserModelId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 6099F28D-FC32-470A-BC69-4B061A76EF2E
var IID_IAppListEntry3 = syscall.GUID{0x6099F28D, 0xFC32, 0x470A,
	[8]byte{0xBC, 0x69, 0x4B, 0x06, 0x1A, 0x76, 0xEF, 0x2E}}

type IAppListEntry3Interface interface {
	win32.IInspectableInterface
	LaunchForUserAsync(user *IUser) *IAsyncOperation[bool]
}

type IAppListEntry3Vtbl struct {
	win32.IInspectableVtbl
	LaunchForUserAsync uintptr
}

type IAppListEntry3 struct {
	win32.IInspectable
}

func (this *IAppListEntry3) Vtbl() *IAppListEntry3Vtbl {
	return (*IAppListEntry3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppListEntry3) LaunchForUserAsync(user *IUser) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2A131ED2-56F5-487C-8697-5166F3B33DA0
var IID_IAppListEntry4 = syscall.GUID{0x2A131ED2, 0x56F5, 0x487C,
	[8]byte{0x86, 0x97, 0x51, 0x66, 0xF3, 0xB3, 0x3D, 0xA0}}

type IAppListEntry4Interface interface {
	win32.IInspectableInterface
	Get_AppInfo() *IAppInfo
}

type IAppListEntry4Vtbl struct {
	win32.IInspectableVtbl
	Get_AppInfo uintptr
}

type IAppListEntry4 struct {
	win32.IInspectable
}

func (this *IAppListEntry4) Vtbl() *IAppListEntry4Vtbl {
	return (*IAppListEntry4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppListEntry4) Get_AppInfo() *IAppInfo {
	var _result *IAppInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0AACF7A4-5E1D-49DF-8034-FB6A68BC5ED1
var IID_ICoreApplication = syscall.GUID{0x0AACF7A4, 0x5E1D, 0x49DF,
	[8]byte{0x80, 0x34, 0xFB, 0x6A, 0x68, 0xBC, 0x5E, 0xD1}}

type ICoreApplicationInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Add_Suspending(handler EventHandler[*ISuspendingEventArgs]) EventRegistrationToken
	Remove_Suspending(token EventRegistrationToken)
	Add_Resuming(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_Resuming(token EventRegistrationToken)
	Get_Properties() *IPropertySet
	GetCurrentView() *ICoreApplicationView
	Run(viewSource *IFrameworkViewSource)
	RunWithActivationFactories(activationFactoryCallback *IGetActivationFactory)
}

type ICoreApplicationVtbl struct {
	win32.IInspectableVtbl
	Get_Id                     uintptr
	Add_Suspending             uintptr
	Remove_Suspending          uintptr
	Add_Resuming               uintptr
	Remove_Resuming            uintptr
	Get_Properties             uintptr
	GetCurrentView             uintptr
	Run                        uintptr
	RunWithActivationFactories uintptr
}

type ICoreApplication struct {
	win32.IInspectable
}

func (this *ICoreApplication) Vtbl() *ICoreApplicationVtbl {
	return (*ICoreApplicationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreApplication) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICoreApplication) Add_Suspending(handler EventHandler[*ISuspendingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Suspending, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreApplication) Remove_Suspending(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Suspending, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICoreApplication) Add_Resuming(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Resuming, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreApplication) Remove_Resuming(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Resuming, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICoreApplication) Get_Properties() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICoreApplication) GetCurrentView() *ICoreApplicationView {
	var _result *ICoreApplicationView
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICoreApplication) Run(viewSource *IFrameworkViewSource) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Run, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(viewSource)))
	_ = _hr
}

func (this *ICoreApplication) RunWithActivationFactories(activationFactoryCallback *IGetActivationFactory) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RunWithActivationFactories, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(activationFactoryCallback)))
	_ = _hr
}

// 998681FB-1AB6-4B7F-BE4A-9A0645224C04
var IID_ICoreApplication2 = syscall.GUID{0x998681FB, 0x1AB6, 0x4B7F,
	[8]byte{0xBE, 0x4A, 0x9A, 0x06, 0x45, 0x22, 0x4C, 0x04}}

type ICoreApplication2Interface interface {
	win32.IInspectableInterface
	Add_BackgroundActivated(handler EventHandler[*IBackgroundActivatedEventArgs]) EventRegistrationToken
	Remove_BackgroundActivated(token EventRegistrationToken)
	Add_LeavingBackground(handler EventHandler[*ILeavingBackgroundEventArgs]) EventRegistrationToken
	Remove_LeavingBackground(token EventRegistrationToken)
	Add_EnteredBackground(handler EventHandler[*IEnteredBackgroundEventArgs]) EventRegistrationToken
	Remove_EnteredBackground(token EventRegistrationToken)
	EnablePrelaunch(value bool)
}

type ICoreApplication2Vtbl struct {
	win32.IInspectableVtbl
	Add_BackgroundActivated    uintptr
	Remove_BackgroundActivated uintptr
	Add_LeavingBackground      uintptr
	Remove_LeavingBackground   uintptr
	Add_EnteredBackground      uintptr
	Remove_EnteredBackground   uintptr
	EnablePrelaunch            uintptr
}

type ICoreApplication2 struct {
	win32.IInspectable
}

func (this *ICoreApplication2) Vtbl() *ICoreApplication2Vtbl {
	return (*ICoreApplication2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreApplication2) Add_BackgroundActivated(handler EventHandler[*IBackgroundActivatedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_BackgroundActivated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreApplication2) Remove_BackgroundActivated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_BackgroundActivated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICoreApplication2) Add_LeavingBackground(handler EventHandler[*ILeavingBackgroundEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_LeavingBackground, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreApplication2) Remove_LeavingBackground(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_LeavingBackground, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICoreApplication2) Add_EnteredBackground(handler EventHandler[*IEnteredBackgroundEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_EnteredBackground, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreApplication2) Remove_EnteredBackground(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_EnteredBackground, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICoreApplication2) EnablePrelaunch(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnablePrelaunch, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// FEEC0D39-598B-4507-8A67-772632580A57
var IID_ICoreApplication3 = syscall.GUID{0xFEEC0D39, 0x598B, 0x4507,
	[8]byte{0x8A, 0x67, 0x77, 0x26, 0x32, 0x58, 0x0A, 0x57}}

type ICoreApplication3Interface interface {
	win32.IInspectableInterface
	RequestRestartAsync(launchArguments string) *IAsyncOperation[AppRestartFailureReason]
	RequestRestartForUserAsync(user *IUser, launchArguments string) *IAsyncOperation[AppRestartFailureReason]
}

type ICoreApplication3Vtbl struct {
	win32.IInspectableVtbl
	RequestRestartAsync        uintptr
	RequestRestartForUserAsync uintptr
}

type ICoreApplication3 struct {
	win32.IInspectable
}

func (this *ICoreApplication3) Vtbl() *ICoreApplication3Vtbl {
	return (*ICoreApplication3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreApplication3) RequestRestartAsync(launchArguments string) *IAsyncOperation[AppRestartFailureReason] {
	var _result *IAsyncOperation[AppRestartFailureReason]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestRestartAsync, uintptr(unsafe.Pointer(this)), NewHStr(launchArguments).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICoreApplication3) RequestRestartForUserAsync(user *IUser, launchArguments string) *IAsyncOperation[AppRestartFailureReason] {
	var _result *IAsyncOperation[AppRestartFailureReason]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestRestartForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), NewHStr(launchArguments).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CF86461D-261E-4B72-9ACD-44ED2ACE6A29
var IID_ICoreApplicationExit = syscall.GUID{0xCF86461D, 0x261E, 0x4B72,
	[8]byte{0x9A, 0xCD, 0x44, 0xED, 0x2A, 0xCE, 0x6A, 0x29}}

type ICoreApplicationExitInterface interface {
	win32.IInspectableInterface
	Exit()
	Add_Exiting(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_Exiting(token EventRegistrationToken)
}

type ICoreApplicationExitVtbl struct {
	win32.IInspectableVtbl
	Exit           uintptr
	Add_Exiting    uintptr
	Remove_Exiting uintptr
}

type ICoreApplicationExit struct {
	win32.IInspectable
}

func (this *ICoreApplicationExit) Vtbl() *ICoreApplicationExitVtbl {
	return (*ICoreApplicationExitVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreApplicationExit) Exit() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Exit, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ICoreApplicationExit) Add_Exiting(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Exiting, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreApplicationExit) Remove_Exiting(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Exiting, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// F0E24AB0-DD09-42E1-B0BC-E0E131F78D7E
var IID_ICoreApplicationUnhandledError = syscall.GUID{0xF0E24AB0, 0xDD09, 0x42E1,
	[8]byte{0xB0, 0xBC, 0xE0, 0xE1, 0x31, 0xF7, 0x8D, 0x7E}}

type ICoreApplicationUnhandledErrorInterface interface {
	win32.IInspectableInterface
	Add_UnhandledErrorDetected(handler EventHandler[*IUnhandledErrorDetectedEventArgs]) EventRegistrationToken
	Remove_UnhandledErrorDetected(token EventRegistrationToken)
}

type ICoreApplicationUnhandledErrorVtbl struct {
	win32.IInspectableVtbl
	Add_UnhandledErrorDetected    uintptr
	Remove_UnhandledErrorDetected uintptr
}

type ICoreApplicationUnhandledError struct {
	win32.IInspectable
}

func (this *ICoreApplicationUnhandledError) Vtbl() *ICoreApplicationUnhandledErrorVtbl {
	return (*ICoreApplicationUnhandledErrorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreApplicationUnhandledError) Add_UnhandledErrorDetected(handler EventHandler[*IUnhandledErrorDetectedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_UnhandledErrorDetected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreApplicationUnhandledError) Remove_UnhandledErrorDetected(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_UnhandledErrorDetected, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 518DC408-C077-475B-809E-0BC0C57E4B74
var IID_ICoreApplicationUseCount = syscall.GUID{0x518DC408, 0xC077, 0x475B,
	[8]byte{0x80, 0x9E, 0x0B, 0xC0, 0xC5, 0x7E, 0x4B, 0x74}}

type ICoreApplicationUseCountInterface interface {
	win32.IInspectableInterface
	IncrementApplicationUseCount()
	DecrementApplicationUseCount()
}

type ICoreApplicationUseCountVtbl struct {
	win32.IInspectableVtbl
	IncrementApplicationUseCount uintptr
	DecrementApplicationUseCount uintptr
}

type ICoreApplicationUseCount struct {
	win32.IInspectable
}

func (this *ICoreApplicationUseCount) Vtbl() *ICoreApplicationUseCountVtbl {
	return (*ICoreApplicationUseCountVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreApplicationUseCount) IncrementApplicationUseCount() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IncrementApplicationUseCount, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ICoreApplicationUseCount) DecrementApplicationUseCount() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DecrementApplicationUseCount, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 638BB2DB-451D-4661-B099-414F34FFB9F1
var IID_ICoreApplicationView = syscall.GUID{0x638BB2DB, 0x451D, 0x4661,
	[8]byte{0xB0, 0x99, 0x41, 0x4F, 0x34, 0xFF, 0xB9, 0xF1}}

type ICoreApplicationViewInterface interface {
	win32.IInspectableInterface
	Get_CoreWindow() unsafe.Pointer
	Add_Activated(handler TypedEventHandler[*ICoreApplicationView, *IActivatedEventArgs]) EventRegistrationToken
	Remove_Activated(token EventRegistrationToken)
	Get_IsMain() bool
	Get_IsHosted() bool
}

type ICoreApplicationViewVtbl struct {
	win32.IInspectableVtbl
	Get_CoreWindow   uintptr
	Add_Activated    uintptr
	Remove_Activated uintptr
	Get_IsMain       uintptr
	Get_IsHosted     uintptr
}

type ICoreApplicationView struct {
	win32.IInspectable
}

func (this *ICoreApplicationView) Vtbl() *ICoreApplicationViewVtbl {
	return (*ICoreApplicationViewVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreApplicationView) Get_CoreWindow() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CoreWindow, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreApplicationView) Add_Activated(handler TypedEventHandler[*ICoreApplicationView, *IActivatedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Activated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreApplicationView) Remove_Activated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Activated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICoreApplicationView) Get_IsMain() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsMain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreApplicationView) Get_IsHosted() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsHosted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 68EB7ADF-917F-48EB-9AEB-7DE53E086AB1
var IID_ICoreApplicationView2 = syscall.GUID{0x68EB7ADF, 0x917F, 0x48EB,
	[8]byte{0x9A, 0xEB, 0x7D, 0xE5, 0x3E, 0x08, 0x6A, 0xB1}}

type ICoreApplicationView2Interface interface {
	win32.IInspectableInterface
	Get_Dispatcher() unsafe.Pointer
}

type ICoreApplicationView2Vtbl struct {
	win32.IInspectableVtbl
	Get_Dispatcher uintptr
}

type ICoreApplicationView2 struct {
	win32.IInspectable
}

func (this *ICoreApplicationView2) Vtbl() *ICoreApplicationView2Vtbl {
	return (*ICoreApplicationView2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreApplicationView2) Get_Dispatcher() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Dispatcher, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 07EBE1B3-A4CF-4550-AB70-B07E85330BC8
var IID_ICoreApplicationView3 = syscall.GUID{0x07EBE1B3, 0xA4CF, 0x4550,
	[8]byte{0xAB, 0x70, 0xB0, 0x7E, 0x85, 0x33, 0x0B, 0xC8}}

type ICoreApplicationView3Interface interface {
	win32.IInspectableInterface
	Get_IsComponent() bool
	Get_TitleBar() *ICoreApplicationViewTitleBar
	Add_HostedViewClosing(handler TypedEventHandler[*ICoreApplicationView, *IHostedViewClosingEventArgs]) EventRegistrationToken
	Remove_HostedViewClosing(token EventRegistrationToken)
}

type ICoreApplicationView3Vtbl struct {
	win32.IInspectableVtbl
	Get_IsComponent          uintptr
	Get_TitleBar             uintptr
	Add_HostedViewClosing    uintptr
	Remove_HostedViewClosing uintptr
}

type ICoreApplicationView3 struct {
	win32.IInspectable
}

func (this *ICoreApplicationView3) Vtbl() *ICoreApplicationView3Vtbl {
	return (*ICoreApplicationView3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreApplicationView3) Get_IsComponent() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsComponent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreApplicationView3) Get_TitleBar() *ICoreApplicationViewTitleBar {
	var _result *ICoreApplicationViewTitleBar
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TitleBar, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICoreApplicationView3) Add_HostedViewClosing(handler TypedEventHandler[*ICoreApplicationView, *IHostedViewClosingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_HostedViewClosing, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreApplicationView3) Remove_HostedViewClosing(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_HostedViewClosing, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 2BC095A8-8EF0-446D-9E60-3A3E0428C671
var IID_ICoreApplicationView5 = syscall.GUID{0x2BC095A8, 0x8EF0, 0x446D,
	[8]byte{0x9E, 0x60, 0x3A, 0x3E, 0x04, 0x28, 0xC6, 0x71}}

type ICoreApplicationView5Interface interface {
	win32.IInspectableInterface
	Get_Properties() *IPropertySet
}

type ICoreApplicationView5Vtbl struct {
	win32.IInspectableVtbl
	Get_Properties uintptr
}

type ICoreApplicationView5 struct {
	win32.IInspectable
}

func (this *ICoreApplicationView5) Vtbl() *ICoreApplicationView5Vtbl {
	return (*ICoreApplicationView5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreApplicationView5) Get_Properties() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C119D49A-0679-49BA-803F-B79C5CF34CCA
var IID_ICoreApplicationView6 = syscall.GUID{0xC119D49A, 0x0679, 0x49BA,
	[8]byte{0x80, 0x3F, 0xB7, 0x9C, 0x5C, 0xF3, 0x4C, 0xCA}}

type ICoreApplicationView6Interface interface {
	win32.IInspectableInterface
	Get_DispatcherQueue() *IDispatcherQueue
}

type ICoreApplicationView6Vtbl struct {
	win32.IInspectableVtbl
	Get_DispatcherQueue uintptr
}

type ICoreApplicationView6 struct {
	win32.IInspectable
}

func (this *ICoreApplicationView6) Vtbl() *ICoreApplicationView6Vtbl {
	return (*ICoreApplicationView6Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreApplicationView6) Get_DispatcherQueue() *IDispatcherQueue {
	var _result *IDispatcherQueue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DispatcherQueue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 006D35E3-E1F1-431B-9508-29B96926AC53
var IID_ICoreApplicationViewTitleBar = syscall.GUID{0x006D35E3, 0xE1F1, 0x431B,
	[8]byte{0x95, 0x08, 0x29, 0xB9, 0x69, 0x26, 0xAC, 0x53}}

type ICoreApplicationViewTitleBarInterface interface {
	win32.IInspectableInterface
	Put_ExtendViewIntoTitleBar(value bool)
	Get_ExtendViewIntoTitleBar() bool
	Get_SystemOverlayLeftInset() float64
	Get_SystemOverlayRightInset() float64
	Get_Height() float64
	Add_LayoutMetricsChanged(handler TypedEventHandler[*ICoreApplicationViewTitleBar, interface{}]) EventRegistrationToken
	Remove_LayoutMetricsChanged(token EventRegistrationToken)
	Get_IsVisible() bool
	Add_IsVisibleChanged(handler TypedEventHandler[*ICoreApplicationViewTitleBar, interface{}]) EventRegistrationToken
	Remove_IsVisibleChanged(token EventRegistrationToken)
}

type ICoreApplicationViewTitleBarVtbl struct {
	win32.IInspectableVtbl
	Put_ExtendViewIntoTitleBar  uintptr
	Get_ExtendViewIntoTitleBar  uintptr
	Get_SystemOverlayLeftInset  uintptr
	Get_SystemOverlayRightInset uintptr
	Get_Height                  uintptr
	Add_LayoutMetricsChanged    uintptr
	Remove_LayoutMetricsChanged uintptr
	Get_IsVisible               uintptr
	Add_IsVisibleChanged        uintptr
	Remove_IsVisibleChanged     uintptr
}

type ICoreApplicationViewTitleBar struct {
	win32.IInspectable
}

func (this *ICoreApplicationViewTitleBar) Vtbl() *ICoreApplicationViewTitleBarVtbl {
	return (*ICoreApplicationViewTitleBarVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreApplicationViewTitleBar) Put_ExtendViewIntoTitleBar(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ExtendViewIntoTitleBar, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICoreApplicationViewTitleBar) Get_ExtendViewIntoTitleBar() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendViewIntoTitleBar, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreApplicationViewTitleBar) Get_SystemOverlayLeftInset() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SystemOverlayLeftInset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreApplicationViewTitleBar) Get_SystemOverlayRightInset() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SystemOverlayRightInset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreApplicationViewTitleBar) Get_Height() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Height, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreApplicationViewTitleBar) Add_LayoutMetricsChanged(handler TypedEventHandler[*ICoreApplicationViewTitleBar, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_LayoutMetricsChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreApplicationViewTitleBar) Remove_LayoutMetricsChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_LayoutMetricsChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICoreApplicationViewTitleBar) Get_IsVisible() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsVisible, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreApplicationViewTitleBar) Add_IsVisibleChanged(handler TypedEventHandler[*ICoreApplicationViewTitleBar, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_IsVisibleChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreApplicationViewTitleBar) Remove_IsVisibleChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_IsVisibleChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 1ADA0E3E-E4A2-4123-B451-DC96BF800419
var IID_ICoreImmersiveApplication = syscall.GUID{0x1ADA0E3E, 0xE4A2, 0x4123,
	[8]byte{0xB4, 0x51, 0xDC, 0x96, 0xBF, 0x80, 0x04, 0x19}}

type ICoreImmersiveApplicationInterface interface {
	win32.IInspectableInterface
	Get_Views() *IVectorView[*ICoreApplicationView]
	CreateNewView(runtimeType string, entryPoint string) *ICoreApplicationView
	Get_MainView() *ICoreApplicationView
}

type ICoreImmersiveApplicationVtbl struct {
	win32.IInspectableVtbl
	Get_Views     uintptr
	CreateNewView uintptr
	Get_MainView  uintptr
}

type ICoreImmersiveApplication struct {
	win32.IInspectable
}

func (this *ICoreImmersiveApplication) Vtbl() *ICoreImmersiveApplicationVtbl {
	return (*ICoreImmersiveApplicationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreImmersiveApplication) Get_Views() *IVectorView[*ICoreApplicationView] {
	var _result *IVectorView[*ICoreApplicationView]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Views, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICoreImmersiveApplication) CreateNewView(runtimeType string, entryPoint string) *ICoreApplicationView {
	var _result *ICoreApplicationView
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateNewView, uintptr(unsafe.Pointer(this)), NewHStr(runtimeType).Ptr, NewHStr(entryPoint).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICoreImmersiveApplication) Get_MainView() *ICoreApplicationView {
	var _result *ICoreApplicationView
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MainView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 828E1E36-E9E3-4CFC-9B66-48B78EA9BB2C
var IID_ICoreImmersiveApplication2 = syscall.GUID{0x828E1E36, 0xE9E3, 0x4CFC,
	[8]byte{0x9B, 0x66, 0x48, 0xB7, 0x8E, 0xA9, 0xBB, 0x2C}}

type ICoreImmersiveApplication2Interface interface {
	win32.IInspectableInterface
	CreateNewViewFromMainView() *ICoreApplicationView
}

type ICoreImmersiveApplication2Vtbl struct {
	win32.IInspectableVtbl
	CreateNewViewFromMainView uintptr
}

type ICoreImmersiveApplication2 struct {
	win32.IInspectable
}

func (this *ICoreImmersiveApplication2) Vtbl() *ICoreImmersiveApplication2Vtbl {
	return (*ICoreImmersiveApplication2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreImmersiveApplication2) CreateNewViewFromMainView() *ICoreApplicationView {
	var _result *ICoreApplicationView
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateNewViewFromMainView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 34A05B2F-EE0D-41E5-8314-CF10C91BF0AF
var IID_ICoreImmersiveApplication3 = syscall.GUID{0x34A05B2F, 0xEE0D, 0x41E5,
	[8]byte{0x83, 0x14, 0xCF, 0x10, 0xC9, 0x1B, 0xF0, 0xAF}}

type ICoreImmersiveApplication3Interface interface {
	win32.IInspectableInterface
	CreateNewViewWithViewSource(viewSource *IFrameworkViewSource) *ICoreApplicationView
}

type ICoreImmersiveApplication3Vtbl struct {
	win32.IInspectableVtbl
	CreateNewViewWithViewSource uintptr
}

type ICoreImmersiveApplication3 struct {
	win32.IInspectable
}

func (this *ICoreImmersiveApplication3) Vtbl() *ICoreImmersiveApplication3Vtbl {
	return (*ICoreImmersiveApplication3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreImmersiveApplication3) CreateNewViewWithViewSource(viewSource *IFrameworkViewSource) *ICoreApplicationView {
	var _result *ICoreApplicationView
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateNewViewWithViewSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(viewSource)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FAAB5CD0-8924-45AC-AD0F-A08FAE5D0324
var IID_IFrameworkView = syscall.GUID{0xFAAB5CD0, 0x8924, 0x45AC,
	[8]byte{0xAD, 0x0F, 0xA0, 0x8F, 0xAE, 0x5D, 0x03, 0x24}}

type IFrameworkViewInterface interface {
	win32.IInspectableInterface
	Initialize(applicationView *ICoreApplicationView)
	SetWindow(window unsafe.Pointer)
	Load(entryPoint string)
	Run()
	Uninitialize()
}

type IFrameworkViewVtbl struct {
	win32.IInspectableVtbl
	Initialize   uintptr
	SetWindow    uintptr
	Load         uintptr
	Run          uintptr
	Uninitialize uintptr
}

type IFrameworkView struct {
	win32.IInspectable
}

func (this *IFrameworkView) Vtbl() *IFrameworkViewVtbl {
	return (*IFrameworkViewVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFrameworkView) Initialize(applicationView *ICoreApplicationView) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Initialize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(applicationView)))
	_ = _hr
}

func (this *IFrameworkView) SetWindow(window unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetWindow, uintptr(unsafe.Pointer(this)), uintptr(window))
	_ = _hr
}

func (this *IFrameworkView) Load(entryPoint string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Load, uintptr(unsafe.Pointer(this)), NewHStr(entryPoint).Ptr)
	_ = _hr
}

func (this *IFrameworkView) Run() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Run, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IFrameworkView) Uninitialize() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Uninitialize, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// CD770614-65C4-426C-9494-34FC43554862
var IID_IFrameworkViewSource = syscall.GUID{0xCD770614, 0x65C4, 0x426C,
	[8]byte{0x94, 0x94, 0x34, 0xFC, 0x43, 0x55, 0x48, 0x62}}

type IFrameworkViewSourceInterface interface {
	win32.IInspectableInterface
	CreateView() *IFrameworkView
}

type IFrameworkViewSourceVtbl struct {
	win32.IInspectableVtbl
	CreateView uintptr
}

type IFrameworkViewSource struct {
	win32.IInspectable
}

func (this *IFrameworkViewSource) Vtbl() *IFrameworkViewSourceVtbl {
	return (*IFrameworkViewSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFrameworkViewSource) CreateView() *IFrameworkView {
	var _result *IFrameworkView
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D238943C-B24E-4790-ACB5-3E4243C4FF87
var IID_IHostedViewClosingEventArgs = syscall.GUID{0xD238943C, 0xB24E, 0x4790,
	[8]byte{0xAC, 0xB5, 0x3E, 0x42, 0x43, 0xC4, 0xFF, 0x87}}

type IHostedViewClosingEventArgsInterface interface {
	win32.IInspectableInterface
	GetDeferral() *IDeferral
}

type IHostedViewClosingEventArgsVtbl struct {
	win32.IInspectableVtbl
	GetDeferral uintptr
}

type IHostedViewClosingEventArgs struct {
	win32.IInspectable
}

func (this *IHostedViewClosingEventArgs) Vtbl() *IHostedViewClosingEventArgsVtbl {
	return (*IHostedViewClosingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHostedViewClosingEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9459B726-53B5-4686-9EAF-FA8162DC3980
var IID_IUnhandledError = syscall.GUID{0x9459B726, 0x53B5, 0x4686,
	[8]byte{0x9E, 0xAF, 0xFA, 0x81, 0x62, 0xDC, 0x39, 0x80}}

type IUnhandledErrorInterface interface {
	win32.IInspectableInterface
	Get_Handled() bool
	Propagate()
}

type IUnhandledErrorVtbl struct {
	win32.IInspectableVtbl
	Get_Handled uintptr
	Propagate   uintptr
}

type IUnhandledError struct {
	win32.IInspectable
}

func (this *IUnhandledError) Vtbl() *IUnhandledErrorVtbl {
	return (*IUnhandledErrorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUnhandledError) Get_Handled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Handled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUnhandledError) Propagate() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Propagate, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 679AB78B-B336-4822-AC40-0D750F0B7A2B
var IID_IUnhandledErrorDetectedEventArgs = syscall.GUID{0x679AB78B, 0xB336, 0x4822,
	[8]byte{0xAC, 0x40, 0x0D, 0x75, 0x0F, 0x0B, 0x7A, 0x2B}}

type IUnhandledErrorDetectedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_UnhandledError() *IUnhandledError
}

type IUnhandledErrorDetectedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_UnhandledError uintptr
}

type IUnhandledErrorDetectedEventArgs struct {
	win32.IInspectable
}

func (this *IUnhandledErrorDetectedEventArgs) Vtbl() *IUnhandledErrorDetectedEventArgsVtbl {
	return (*IUnhandledErrorDetectedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUnhandledErrorDetectedEventArgs) Get_UnhandledError() *IUnhandledError {
	var _result *IUnhandledError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UnhandledError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type AppListEntry struct {
	RtClass
	*IAppListEntry
}
