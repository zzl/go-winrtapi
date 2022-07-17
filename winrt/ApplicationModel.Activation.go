package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type ActivationKind int32

const (
	ActivationKind_Launch                              ActivationKind = 0
	ActivationKind_Search                              ActivationKind = 1
	ActivationKind_ShareTarget                         ActivationKind = 2
	ActivationKind_File                                ActivationKind = 3
	ActivationKind_Protocol                            ActivationKind = 4
	ActivationKind_FileOpenPicker                      ActivationKind = 5
	ActivationKind_FileSavePicker                      ActivationKind = 6
	ActivationKind_CachedFileUpdater                   ActivationKind = 7
	ActivationKind_ContactPicker                       ActivationKind = 8
	ActivationKind_Device                              ActivationKind = 9
	ActivationKind_PrintTaskSettings                   ActivationKind = 10
	ActivationKind_CameraSettings                      ActivationKind = 11
	ActivationKind_RestrictedLaunch                    ActivationKind = 12
	ActivationKind_AppointmentsProvider                ActivationKind = 13
	ActivationKind_Contact                             ActivationKind = 14
	ActivationKind_LockScreenCall                      ActivationKind = 15
	ActivationKind_VoiceCommand                        ActivationKind = 16
	ActivationKind_LockScreen                          ActivationKind = 17
	ActivationKind_PickerReturned                      ActivationKind = 1000
	ActivationKind_WalletAction                        ActivationKind = 1001
	ActivationKind_PickFileContinuation                ActivationKind = 1002
	ActivationKind_PickSaveFileContinuation            ActivationKind = 1003
	ActivationKind_PickFolderContinuation              ActivationKind = 1004
	ActivationKind_WebAuthenticationBrokerContinuation ActivationKind = 1005
	ActivationKind_WebAccountProvider                  ActivationKind = 1006
	ActivationKind_ComponentUI                         ActivationKind = 1007
	ActivationKind_ProtocolForResults                  ActivationKind = 1009
	ActivationKind_ToastNotification                   ActivationKind = 1010
	ActivationKind_Print3DWorkflow                     ActivationKind = 1011
	ActivationKind_DialReceiver                        ActivationKind = 1012
	ActivationKind_DevicePairing                       ActivationKind = 1013
	ActivationKind_UserDataAccountsProvider            ActivationKind = 1014
	ActivationKind_FilePickerExperience                ActivationKind = 1015
	ActivationKind_LockScreenComponent                 ActivationKind = 1016
	ActivationKind_ContactPanel                        ActivationKind = 1017
	ActivationKind_PrintWorkflowForegroundTask         ActivationKind = 1018
	ActivationKind_GameUIProvider                      ActivationKind = 1019
	ActivationKind_StartupTask                         ActivationKind = 1020
	ActivationKind_CommandLineLaunch                   ActivationKind = 1021
	ActivationKind_BarcodeScannerProvider              ActivationKind = 1022
	ActivationKind_PrintSupportJobUI                   ActivationKind = 1023
	ActivationKind_PrintSupportSettingsUI              ActivationKind = 1024
	ActivationKind_PhoneCallActivation                 ActivationKind = 1025
	ActivationKind_VpnForeground                       ActivationKind = 1026
)

// enum
type ApplicationExecutionState int32

const (
	ApplicationExecutionState_NotRunning   ApplicationExecutionState = 0
	ApplicationExecutionState_Running      ApplicationExecutionState = 1
	ApplicationExecutionState_Suspended    ApplicationExecutionState = 2
	ApplicationExecutionState_Terminated   ApplicationExecutionState = 3
	ApplicationExecutionState_ClosedByUser ApplicationExecutionState = 4
)

// structs

type ActivatedEventsContract struct {
}

type ActivationCameraSettingsContract struct {
}

type ContactActivatedEventsContract struct {
}

type WebUISearchActivatedEventsContract struct {
}

// interfaces

// CF651713-CD08-4FD8-B697-A281B6544E2E
var IID_IActivatedEventArgs = syscall.GUID{0xCF651713, 0xCD08, 0x4FD8,
	[8]byte{0xB6, 0x97, 0xA2, 0x81, 0xB6, 0x54, 0x4E, 0x2E}}

type IActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Kind() ActivationKind
	Get_PreviousExecutionState() ApplicationExecutionState
	Get_SplashScreen() *ISplashScreen
}

type IActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Kind                   uintptr
	Get_PreviousExecutionState uintptr
	Get_SplashScreen           uintptr
}

type IActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IActivatedEventArgs) Vtbl() *IActivatedEventArgsVtbl {
	return (*IActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IActivatedEventArgs) Get_Kind() ActivationKind {
	var _result ActivationKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Kind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IActivatedEventArgs) Get_PreviousExecutionState() ApplicationExecutionState {
	var _result ApplicationExecutionState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PreviousExecutionState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IActivatedEventArgs) Get_SplashScreen() *ISplashScreen {
	var _result *ISplashScreen
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SplashScreen, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1CF09B9E-9962-4936-80FF-AFC8E8AE5C8C
var IID_IActivatedEventArgsWithUser = syscall.GUID{0x1CF09B9E, 0x9962, 0x4936,
	[8]byte{0x80, 0xFF, 0xAF, 0xC8, 0xE8, 0xAE, 0x5C, 0x8C}}

type IActivatedEventArgsWithUserInterface interface {
	win32.IInspectableInterface
	Get_User() *IUser
}

type IActivatedEventArgsWithUserVtbl struct {
	win32.IInspectableVtbl
	Get_User uintptr
}

type IActivatedEventArgsWithUser struct {
	win32.IInspectable
}

func (this *IActivatedEventArgsWithUser) Vtbl() *IActivatedEventArgsWithUserVtbl {
	return (*IActivatedEventArgsWithUserVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IActivatedEventArgsWithUser) Get_User() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_User, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 930CEF4B-B829-40FC-88F4-8513E8A64738
var IID_IApplicationViewActivatedEventArgs = syscall.GUID{0x930CEF4B, 0xB829, 0x40FC,
	[8]byte{0x88, 0xF4, 0x85, 0x13, 0xE8, 0xA6, 0x47, 0x38}}

type IApplicationViewActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_CurrentlyShownApplicationViewId() int32
}

type IApplicationViewActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_CurrentlyShownApplicationViewId uintptr
}

type IApplicationViewActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IApplicationViewActivatedEventArgs) Vtbl() *IApplicationViewActivatedEventArgsVtbl {
	return (*IApplicationViewActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationViewActivatedEventArgs) Get_CurrentlyShownApplicationViewId() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentlyShownApplicationViewId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3364C405-933C-4E7D-A034-500FB8DCD9F3
var IID_IAppointmentsProviderActivatedEventArgs = syscall.GUID{0x3364C405, 0x933C, 0x4E7D,
	[8]byte{0xA0, 0x34, 0x50, 0x0F, 0xB8, 0xDC, 0xD9, 0xF3}}

type IAppointmentsProviderActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Verb() string
}

type IAppointmentsProviderActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Verb uintptr
}

type IAppointmentsProviderActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IAppointmentsProviderActivatedEventArgs) Vtbl() *IAppointmentsProviderActivatedEventArgsVtbl {
	return (*IAppointmentsProviderActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppointmentsProviderActivatedEventArgs) Get_Verb() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Verb, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// A2861367-CEE5-4E4D-9ED7-41C34EC18B02
var IID_IAppointmentsProviderAddAppointmentActivatedEventArgs = syscall.GUID{0xA2861367, 0xCEE5, 0x4E4D,
	[8]byte{0x9E, 0xD7, 0x41, 0xC3, 0x4E, 0xC1, 0x8B, 0x02}}

type IAppointmentsProviderAddAppointmentActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_AddAppointmentOperation() unsafe.Pointer
}

type IAppointmentsProviderAddAppointmentActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_AddAppointmentOperation uintptr
}

type IAppointmentsProviderAddAppointmentActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IAppointmentsProviderAddAppointmentActivatedEventArgs) Vtbl() *IAppointmentsProviderAddAppointmentActivatedEventArgsVtbl {
	return (*IAppointmentsProviderAddAppointmentActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppointmentsProviderAddAppointmentActivatedEventArgs) Get_AddAppointmentOperation() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AddAppointmentOperation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 751F3AB8-0B8E-451C-9F15-966E699BAC25
var IID_IAppointmentsProviderRemoveAppointmentActivatedEventArgs = syscall.GUID{0x751F3AB8, 0x0B8E, 0x451C,
	[8]byte{0x9F, 0x15, 0x96, 0x6E, 0x69, 0x9B, 0xAC, 0x25}}

type IAppointmentsProviderRemoveAppointmentActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_RemoveAppointmentOperation() unsafe.Pointer
}

type IAppointmentsProviderRemoveAppointmentActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_RemoveAppointmentOperation uintptr
}

type IAppointmentsProviderRemoveAppointmentActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IAppointmentsProviderRemoveAppointmentActivatedEventArgs) Vtbl() *IAppointmentsProviderRemoveAppointmentActivatedEventArgsVtbl {
	return (*IAppointmentsProviderRemoveAppointmentActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppointmentsProviderRemoveAppointmentActivatedEventArgs) Get_RemoveAppointmentOperation() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemoveAppointmentOperation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1551B7D4-A981-4067-8A62-0524E4ADE121
var IID_IAppointmentsProviderReplaceAppointmentActivatedEventArgs = syscall.GUID{0x1551B7D4, 0xA981, 0x4067,
	[8]byte{0x8A, 0x62, 0x05, 0x24, 0xE4, 0xAD, 0xE1, 0x21}}

type IAppointmentsProviderReplaceAppointmentActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ReplaceAppointmentOperation() unsafe.Pointer
}

type IAppointmentsProviderReplaceAppointmentActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ReplaceAppointmentOperation uintptr
}

type IAppointmentsProviderReplaceAppointmentActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IAppointmentsProviderReplaceAppointmentActivatedEventArgs) Vtbl() *IAppointmentsProviderReplaceAppointmentActivatedEventArgsVtbl {
	return (*IAppointmentsProviderReplaceAppointmentActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppointmentsProviderReplaceAppointmentActivatedEventArgs) Get_ReplaceAppointmentOperation() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReplaceAppointmentOperation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3958F065-9841-4CA5-999B-885198B9EF2A
var IID_IAppointmentsProviderShowAppointmentDetailsActivatedEventArgs = syscall.GUID{0x3958F065, 0x9841, 0x4CA5,
	[8]byte{0x99, 0x9B, 0x88, 0x51, 0x98, 0xB9, 0xEF, 0x2A}}

type IAppointmentsProviderShowAppointmentDetailsActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_InstanceStartDate() *IReference[DateTime]
	Get_LocalId() string
	Get_RoamingId() string
}

type IAppointmentsProviderShowAppointmentDetailsActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_InstanceStartDate uintptr
	Get_LocalId           uintptr
	Get_RoamingId         uintptr
}

type IAppointmentsProviderShowAppointmentDetailsActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IAppointmentsProviderShowAppointmentDetailsActivatedEventArgs) Vtbl() *IAppointmentsProviderShowAppointmentDetailsActivatedEventArgsVtbl {
	return (*IAppointmentsProviderShowAppointmentDetailsActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppointmentsProviderShowAppointmentDetailsActivatedEventArgs) Get_InstanceStartDate() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InstanceStartDate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppointmentsProviderShowAppointmentDetailsActivatedEventArgs) Get_LocalId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppointmentsProviderShowAppointmentDetailsActivatedEventArgs) Get_RoamingId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RoamingId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 9BAEABA6-0E0B-49AA-BABC-12B1DC774986
var IID_IAppointmentsProviderShowTimeFrameActivatedEventArgs = syscall.GUID{0x9BAEABA6, 0x0E0B, 0x49AA,
	[8]byte{0xBA, 0xBC, 0x12, 0xB1, 0xDC, 0x77, 0x49, 0x86}}

type IAppointmentsProviderShowTimeFrameActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_TimeToShow() DateTime
	Get_Duration() TimeSpan
}

type IAppointmentsProviderShowTimeFrameActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_TimeToShow uintptr
	Get_Duration   uintptr
}

type IAppointmentsProviderShowTimeFrameActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IAppointmentsProviderShowTimeFrameActivatedEventArgs) Vtbl() *IAppointmentsProviderShowTimeFrameActivatedEventArgsVtbl {
	return (*IAppointmentsProviderShowTimeFrameActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppointmentsProviderShowTimeFrameActivatedEventArgs) Get_TimeToShow() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimeToShow, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppointmentsProviderShowTimeFrameActivatedEventArgs) Get_Duration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// AB14BEE0-E760-440E-A91C-44796DE3A92D
var IID_IBackgroundActivatedEventArgs = syscall.GUID{0xAB14BEE0, 0xE760, 0x440E,
	[8]byte{0xA9, 0x1C, 0x44, 0x79, 0x6D, 0xE3, 0xA9, 0x2D}}

type IBackgroundActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_TaskInstance() *IBackgroundTaskInstance
}

type IBackgroundActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_TaskInstance uintptr
}

type IBackgroundActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IBackgroundActivatedEventArgs) Vtbl() *IBackgroundActivatedEventArgsVtbl {
	return (*IBackgroundActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundActivatedEventArgs) Get_TaskInstance() *IBackgroundTaskInstance {
	var _result *IBackgroundTaskInstance
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TaskInstance, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6772797C-99BF-4349-AF22-E4123560371C
var IID_IBarcodeScannerPreviewActivatedEventArgs = syscall.GUID{0x6772797C, 0x99BF, 0x4349,
	[8]byte{0xAF, 0x22, 0xE4, 0x12, 0x35, 0x60, 0x37, 0x1C}}

type IBarcodeScannerPreviewActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ConnectionId() string
}

type IBarcodeScannerPreviewActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ConnectionId uintptr
}

type IBarcodeScannerPreviewActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IBarcodeScannerPreviewActivatedEventArgs) Vtbl() *IBarcodeScannerPreviewActivatedEventArgsVtbl {
	return (*IBarcodeScannerPreviewActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarcodeScannerPreviewActivatedEventArgs) Get_ConnectionId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConnectionId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// D06EB1C7-3805-4ECB-B757-6CF15E26FEF3
var IID_ICachedFileUpdaterActivatedEventArgs = syscall.GUID{0xD06EB1C7, 0x3805, 0x4ECB,
	[8]byte{0xB7, 0x57, 0x6C, 0xF1, 0x5E, 0x26, 0xFE, 0xF3}}

type ICachedFileUpdaterActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_CachedFileUpdaterUI() *ICachedFileUpdaterUI
}

type ICachedFileUpdaterActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_CachedFileUpdaterUI uintptr
}

type ICachedFileUpdaterActivatedEventArgs struct {
	win32.IInspectable
}

func (this *ICachedFileUpdaterActivatedEventArgs) Vtbl() *ICachedFileUpdaterActivatedEventArgsVtbl {
	return (*ICachedFileUpdaterActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICachedFileUpdaterActivatedEventArgs) Get_CachedFileUpdaterUI() *ICachedFileUpdaterUI {
	var _result *ICachedFileUpdaterUI
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedFileUpdaterUI, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FB67A508-2DAD-490A-9170-DCA036EB114B
var IID_ICameraSettingsActivatedEventArgs = syscall.GUID{0xFB67A508, 0x2DAD, 0x490A,
	[8]byte{0x91, 0x70, 0xDC, 0xA0, 0x36, 0xEB, 0x11, 0x4B}}

type ICameraSettingsActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_VideoDeviceController() interface{}
	Get_VideoDeviceExtension() interface{}
}

type ICameraSettingsActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_VideoDeviceController uintptr
	Get_VideoDeviceExtension  uintptr
}

type ICameraSettingsActivatedEventArgs struct {
	win32.IInspectable
}

func (this *ICameraSettingsActivatedEventArgs) Vtbl() *ICameraSettingsActivatedEventArgsVtbl {
	return (*ICameraSettingsActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICameraSettingsActivatedEventArgs) Get_VideoDeviceController() interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoDeviceController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICameraSettingsActivatedEventArgs) Get_VideoDeviceExtension() interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoDeviceExtension, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4506472C-006A-48EB-8AFB-D07AB25E3366
var IID_ICommandLineActivatedEventArgs = syscall.GUID{0x4506472C, 0x006A, 0x48EB,
	[8]byte{0x8A, 0xFB, 0xD0, 0x7A, 0xB2, 0x5E, 0x33, 0x66}}

type ICommandLineActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Operation() *ICommandLineActivationOperation
}

type ICommandLineActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Operation uintptr
}

type ICommandLineActivatedEventArgs struct {
	win32.IInspectable
}

func (this *ICommandLineActivatedEventArgs) Vtbl() *ICommandLineActivatedEventArgsVtbl {
	return (*ICommandLineActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICommandLineActivatedEventArgs) Get_Operation() *ICommandLineActivationOperation {
	var _result *ICommandLineActivationOperation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Operation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 994B2841-C59E-4F69-BCFD-B61ED4E622EB
var IID_ICommandLineActivationOperation = syscall.GUID{0x994B2841, 0xC59E, 0x4F69,
	[8]byte{0xBC, 0xFD, 0xB6, 0x1E, 0xD4, 0xE6, 0x22, 0xEB}}

type ICommandLineActivationOperationInterface interface {
	win32.IInspectableInterface
	Get_Arguments() string
	Get_CurrentDirectoryPath() string
	Put_ExitCode(value int32)
	Get_ExitCode() int32
	GetDeferral() *IDeferral
}

type ICommandLineActivationOperationVtbl struct {
	win32.IInspectableVtbl
	Get_Arguments            uintptr
	Get_CurrentDirectoryPath uintptr
	Put_ExitCode             uintptr
	Get_ExitCode             uintptr
	GetDeferral              uintptr
}

type ICommandLineActivationOperation struct {
	win32.IInspectable
}

func (this *ICommandLineActivationOperation) Vtbl() *ICommandLineActivationOperationVtbl {
	return (*ICommandLineActivationOperationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICommandLineActivationOperation) Get_Arguments() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Arguments, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICommandLineActivationOperation) Get_CurrentDirectoryPath() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentDirectoryPath, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICommandLineActivationOperation) Put_ExitCode(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ExitCode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICommandLineActivationOperation) Get_ExitCode() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExitCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommandLineActivationOperation) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D627A1C4-C025-4C41-9DEF-F1EAFAD075E7
var IID_IContactActivatedEventArgs = syscall.GUID{0xD627A1C4, 0xC025, 0x4C41,
	[8]byte{0x9D, 0xEF, 0xF1, 0xEA, 0xFA, 0xD0, 0x75, 0xE7}}

type IContactActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Verb() string
}

type IContactActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Verb uintptr
}

type IContactActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IContactActivatedEventArgs) Vtbl() *IContactActivatedEventArgsVtbl {
	return (*IContactActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IContactActivatedEventArgs) Get_Verb() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Verb, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// C2DF14C7-30EB-41C6-B3BC-5B1694F9DAB3
var IID_IContactCallActivatedEventArgs = syscall.GUID{0xC2DF14C7, 0x30EB, 0x41C6,
	[8]byte{0xB3, 0xBC, 0x5B, 0x16, 0x94, 0xF9, 0xDA, 0xB3}}

type IContactCallActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ServiceId() string
	Get_ServiceUserId() string
	Get_Contact() unsafe.Pointer
}

type IContactCallActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ServiceId     uintptr
	Get_ServiceUserId uintptr
	Get_Contact       uintptr
}

type IContactCallActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IContactCallActivatedEventArgs) Vtbl() *IContactCallActivatedEventArgsVtbl {
	return (*IContactCallActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IContactCallActivatedEventArgs) Get_ServiceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IContactCallActivatedEventArgs) Get_ServiceUserId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceUserId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IContactCallActivatedEventArgs) Get_Contact() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Contact, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// B32BF870-EEE7-4AD2-AAF1-A87EFFCF00A4
var IID_IContactMapActivatedEventArgs = syscall.GUID{0xB32BF870, 0xEEE7, 0x4AD2,
	[8]byte{0xAA, 0xF1, 0xA8, 0x7E, 0xFF, 0xCF, 0x00, 0xA4}}

type IContactMapActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Address() unsafe.Pointer
	Get_Contact() unsafe.Pointer
}

type IContactMapActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Address uintptr
	Get_Contact uintptr
}

type IContactMapActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IContactMapActivatedEventArgs) Vtbl() *IContactMapActivatedEventArgsVtbl {
	return (*IContactMapActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IContactMapActivatedEventArgs) Get_Address() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Address, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IContactMapActivatedEventArgs) Get_Contact() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Contact, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// DE598DB2-0E03-43B0-BF56-BCC40B3162DF
var IID_IContactMessageActivatedEventArgs = syscall.GUID{0xDE598DB2, 0x0E03, 0x43B0,
	[8]byte{0xBF, 0x56, 0xBC, 0xC4, 0x0B, 0x31, 0x62, 0xDF}}

type IContactMessageActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ServiceId() string
	Get_ServiceUserId() string
	Get_Contact() unsafe.Pointer
}

type IContactMessageActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ServiceId     uintptr
	Get_ServiceUserId uintptr
	Get_Contact       uintptr
}

type IContactMessageActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IContactMessageActivatedEventArgs) Vtbl() *IContactMessageActivatedEventArgsVtbl {
	return (*IContactMessageActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IContactMessageActivatedEventArgs) Get_ServiceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IContactMessageActivatedEventArgs) Get_ServiceUserId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceUserId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IContactMessageActivatedEventArgs) Get_Contact() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Contact, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 52BB63E4-D3D4-4B63-8051-4AF2082CAB80
var IID_IContactPanelActivatedEventArgs = syscall.GUID{0x52BB63E4, 0xD3D4, 0x4B63,
	[8]byte{0x80, 0x51, 0x4A, 0xF2, 0x08, 0x2C, 0xAB, 0x80}}

type IContactPanelActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ContactPanel() unsafe.Pointer
	Get_Contact() unsafe.Pointer
}

type IContactPanelActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ContactPanel uintptr
	Get_Contact      uintptr
}

type IContactPanelActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IContactPanelActivatedEventArgs) Vtbl() *IContactPanelActivatedEventArgsVtbl {
	return (*IContactPanelActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IContactPanelActivatedEventArgs) Get_ContactPanel() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContactPanel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IContactPanelActivatedEventArgs) Get_Contact() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Contact, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// CE57AAE7-6449-45A7-971F-D113BE7A8936
var IID_IContactPickerActivatedEventArgs = syscall.GUID{0xCE57AAE7, 0x6449, 0x45A7,
	[8]byte{0x97, 0x1F, 0xD1, 0x13, 0xBE, 0x7A, 0x89, 0x36}}

type IContactPickerActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ContactPickerUI() unsafe.Pointer
}

type IContactPickerActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ContactPickerUI uintptr
}

type IContactPickerActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IContactPickerActivatedEventArgs) Vtbl() *IContactPickerActivatedEventArgsVtbl {
	return (*IContactPickerActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IContactPickerActivatedEventArgs) Get_ContactPickerUI() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContactPickerUI, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// B35A3C67-F1E7-4655-AD6E-4857588F552F
var IID_IContactPostActivatedEventArgs = syscall.GUID{0xB35A3C67, 0xF1E7, 0x4655,
	[8]byte{0xAD, 0x6E, 0x48, 0x57, 0x58, 0x8F, 0x55, 0x2F}}

type IContactPostActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ServiceId() string
	Get_ServiceUserId() string
	Get_Contact() unsafe.Pointer
}

type IContactPostActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ServiceId     uintptr
	Get_ServiceUserId uintptr
	Get_Contact       uintptr
}

type IContactPostActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IContactPostActivatedEventArgs) Vtbl() *IContactPostActivatedEventArgsVtbl {
	return (*IContactPostActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IContactPostActivatedEventArgs) Get_ServiceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IContactPostActivatedEventArgs) Get_ServiceUserId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceUserId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IContactPostActivatedEventArgs) Get_Contact() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Contact, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 61079DB8-E3E7-4B4F-858D-5C63A96EF684
var IID_IContactVideoCallActivatedEventArgs = syscall.GUID{0x61079DB8, 0xE3E7, 0x4B4F,
	[8]byte{0x85, 0x8D, 0x5C, 0x63, 0xA9, 0x6E, 0xF6, 0x84}}

type IContactVideoCallActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ServiceId() string
	Get_ServiceUserId() string
	Get_Contact() unsafe.Pointer
}

type IContactVideoCallActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ServiceId     uintptr
	Get_ServiceUserId uintptr
	Get_Contact       uintptr
}

type IContactVideoCallActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IContactVideoCallActivatedEventArgs) Vtbl() *IContactVideoCallActivatedEventArgsVtbl {
	return (*IContactVideoCallActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IContactVideoCallActivatedEventArgs) Get_ServiceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IContactVideoCallActivatedEventArgs) Get_ServiceUserId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceUserId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IContactVideoCallActivatedEventArgs) Get_Contact() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Contact, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4580DCA8-5750-4916-AA52-C0829521EB94
var IID_IContactsProviderActivatedEventArgs = syscall.GUID{0x4580DCA8, 0x5750, 0x4916,
	[8]byte{0xAA, 0x52, 0xC0, 0x82, 0x95, 0x21, 0xEB, 0x94}}

type IContactsProviderActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Verb() string
}

type IContactsProviderActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Verb uintptr
}

type IContactsProviderActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IContactsProviderActivatedEventArgs) Vtbl() *IContactsProviderActivatedEventArgsVtbl {
	return (*IContactsProviderActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IContactsProviderActivatedEventArgs) Get_Verb() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Verb, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// E58106B5-155F-4A94-A742-C7E08F4E188C
var IID_IContinuationActivatedEventArgs = syscall.GUID{0xE58106B5, 0x155F, 0x4A94,
	[8]byte{0xA7, 0x42, 0xC7, 0xE0, 0x8F, 0x4E, 0x18, 0x8C}}

type IContinuationActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ContinuationData() *IPropertySet
}

type IContinuationActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ContinuationData uintptr
}

type IContinuationActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IContinuationActivatedEventArgs) Vtbl() *IContinuationActivatedEventArgsVtbl {
	return (*IContinuationActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IContinuationActivatedEventArgs) Get_ContinuationData() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContinuationData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CD50B9A9-CE10-44D2-8234-C355A073EF33
var IID_IDeviceActivatedEventArgs = syscall.GUID{0xCD50B9A9, 0xCE10, 0x44D2,
	[8]byte{0x82, 0x34, 0xC3, 0x55, 0xA0, 0x73, 0xEF, 0x33}}

type IDeviceActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_DeviceInformationId() string
	Get_Verb() string
}

type IDeviceActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceInformationId uintptr
	Get_Verb                uintptr
}

type IDeviceActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IDeviceActivatedEventArgs) Vtbl() *IDeviceActivatedEventArgsVtbl {
	return (*IDeviceActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceActivatedEventArgs) Get_DeviceInformationId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceInformationId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDeviceActivatedEventArgs) Get_Verb() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Verb, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// EBA0D1E4-ECC6-4148-94ED-F4B37EC05B3E
var IID_IDevicePairingActivatedEventArgs = syscall.GUID{0xEBA0D1E4, 0xECC6, 0x4148,
	[8]byte{0x94, 0xED, 0xF4, 0xB3, 0x7E, 0xC0, 0x5B, 0x3E}}

type IDevicePairingActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_DeviceInformation() *IDeviceInformation
}

type IDevicePairingActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceInformation uintptr
}

type IDevicePairingActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IDevicePairingActivatedEventArgs) Vtbl() *IDevicePairingActivatedEventArgsVtbl {
	return (*IDevicePairingActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDevicePairingActivatedEventArgs) Get_DeviceInformation() *IDeviceInformation {
	var _result *IDeviceInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FB777ED7-85EE-456E-A44D-85D730E70AED
var IID_IDialReceiverActivatedEventArgs = syscall.GUID{0xFB777ED7, 0x85EE, 0x456E,
	[8]byte{0xA4, 0x4D, 0x85, 0xD7, 0x30, 0xE7, 0x0A, 0xED}}

type IDialReceiverActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_AppName() string
}

type IDialReceiverActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_AppName uintptr
}

type IDialReceiverActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IDialReceiverActivatedEventArgs) Vtbl() *IDialReceiverActivatedEventArgsVtbl {
	return (*IDialReceiverActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDialReceiverActivatedEventArgs) Get_AppName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// BB2AFC33-93B1-42ED-8B26-236DD9C78496
var IID_IFileActivatedEventArgs = syscall.GUID{0xBB2AFC33, 0x93B1, 0x42ED,
	[8]byte{0x8B, 0x26, 0x23, 0x6D, 0xD9, 0xC7, 0x84, 0x96}}

type IFileActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Files() *IVectorView[*IStorageItem]
	Get_Verb() string
}

type IFileActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Files uintptr
	Get_Verb  uintptr
}

type IFileActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IFileActivatedEventArgs) Vtbl() *IFileActivatedEventArgsVtbl {
	return (*IFileActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFileActivatedEventArgs) Get_Files() *IVectorView[*IStorageItem] {
	var _result *IVectorView[*IStorageItem]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Files, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileActivatedEventArgs) Get_Verb() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Verb, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 2D60F06B-D25F-4D25-8653-E1C5E1108309
var IID_IFileActivatedEventArgsWithCallerPackageFamilyName = syscall.GUID{0x2D60F06B, 0xD25F, 0x4D25,
	[8]byte{0x86, 0x53, 0xE1, 0xC5, 0xE1, 0x10, 0x83, 0x09}}

type IFileActivatedEventArgsWithCallerPackageFamilyNameInterface interface {
	win32.IInspectableInterface
	Get_CallerPackageFamilyName() string
}

type IFileActivatedEventArgsWithCallerPackageFamilyNameVtbl struct {
	win32.IInspectableVtbl
	Get_CallerPackageFamilyName uintptr
}

type IFileActivatedEventArgsWithCallerPackageFamilyName struct {
	win32.IInspectable
}

func (this *IFileActivatedEventArgsWithCallerPackageFamilyName) Vtbl() *IFileActivatedEventArgsWithCallerPackageFamilyNameVtbl {
	return (*IFileActivatedEventArgsWithCallerPackageFamilyNameVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFileActivatedEventArgsWithCallerPackageFamilyName) Get_CallerPackageFamilyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CallerPackageFamilyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 433BA1A4-E1E2-48FD-B7FC-B5D6EEE65033
var IID_IFileActivatedEventArgsWithNeighboringFiles = syscall.GUID{0x433BA1A4, 0xE1E2, 0x48FD,
	[8]byte{0xB7, 0xFC, 0xB5, 0xD6, 0xEE, 0xE6, 0x50, 0x33}}

type IFileActivatedEventArgsWithNeighboringFilesInterface interface {
	win32.IInspectableInterface
	Get_NeighboringFilesQuery() unsafe.Pointer
}

type IFileActivatedEventArgsWithNeighboringFilesVtbl struct {
	win32.IInspectableVtbl
	Get_NeighboringFilesQuery uintptr
}

type IFileActivatedEventArgsWithNeighboringFiles struct {
	win32.IInspectable
}

func (this *IFileActivatedEventArgsWithNeighboringFiles) Vtbl() *IFileActivatedEventArgsWithNeighboringFilesVtbl {
	return (*IFileActivatedEventArgsWithNeighboringFilesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFileActivatedEventArgsWithNeighboringFiles) Get_NeighboringFilesQuery() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NeighboringFilesQuery, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 72827082-5525-4BF2-BC09-1F5095D4964D
var IID_IFileOpenPickerActivatedEventArgs = syscall.GUID{0x72827082, 0x5525, 0x4BF2,
	[8]byte{0xBC, 0x09, 0x1F, 0x50, 0x95, 0xD4, 0x96, 0x4D}}

type IFileOpenPickerActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_FileOpenPickerUI() unsafe.Pointer
}

type IFileOpenPickerActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_FileOpenPickerUI uintptr
}

type IFileOpenPickerActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IFileOpenPickerActivatedEventArgs) Vtbl() *IFileOpenPickerActivatedEventArgsVtbl {
	return (*IFileOpenPickerActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFileOpenPickerActivatedEventArgs) Get_FileOpenPickerUI() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FileOpenPickerUI, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5E731F66-8D1F-45FB-AF1D-73205C8FC7A1
var IID_IFileOpenPickerActivatedEventArgs2 = syscall.GUID{0x5E731F66, 0x8D1F, 0x45FB,
	[8]byte{0xAF, 0x1D, 0x73, 0x20, 0x5C, 0x8F, 0xC7, 0xA1}}

type IFileOpenPickerActivatedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_CallerPackageFamilyName() string
}

type IFileOpenPickerActivatedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_CallerPackageFamilyName uintptr
}

type IFileOpenPickerActivatedEventArgs2 struct {
	win32.IInspectable
}

func (this *IFileOpenPickerActivatedEventArgs2) Vtbl() *IFileOpenPickerActivatedEventArgs2Vtbl {
	return (*IFileOpenPickerActivatedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFileOpenPickerActivatedEventArgs2) Get_CallerPackageFamilyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CallerPackageFamilyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// F0FA3F3A-D4E8-4AD3-9C34-2308F32FCEC9
var IID_IFileOpenPickerContinuationEventArgs = syscall.GUID{0xF0FA3F3A, 0xD4E8, 0x4AD3,
	[8]byte{0x9C, 0x34, 0x23, 0x08, 0xF3, 0x2F, 0xCE, 0xC9}}

type IFileOpenPickerContinuationEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Files() *IVectorView[*IStorageFile]
}

type IFileOpenPickerContinuationEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Files uintptr
}

type IFileOpenPickerContinuationEventArgs struct {
	win32.IInspectable
}

func (this *IFileOpenPickerContinuationEventArgs) Vtbl() *IFileOpenPickerContinuationEventArgsVtbl {
	return (*IFileOpenPickerContinuationEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFileOpenPickerContinuationEventArgs) Get_Files() *IVectorView[*IStorageFile] {
	var _result *IVectorView[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Files, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 81C19CF1-74E6-4387-82EB-BB8FD64B4346
var IID_IFileSavePickerActivatedEventArgs = syscall.GUID{0x81C19CF1, 0x74E6, 0x4387,
	[8]byte{0x82, 0xEB, 0xBB, 0x8F, 0xD6, 0x4B, 0x43, 0x46}}

type IFileSavePickerActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_FileSavePickerUI() unsafe.Pointer
}

type IFileSavePickerActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_FileSavePickerUI uintptr
}

type IFileSavePickerActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IFileSavePickerActivatedEventArgs) Vtbl() *IFileSavePickerActivatedEventArgsVtbl {
	return (*IFileSavePickerActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFileSavePickerActivatedEventArgs) Get_FileSavePickerUI() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FileSavePickerUI, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6B73FE13-2CF2-4D48-8CBC-AF67D23F1CE7
var IID_IFileSavePickerActivatedEventArgs2 = syscall.GUID{0x6B73FE13, 0x2CF2, 0x4D48,
	[8]byte{0x8C, 0xBC, 0xAF, 0x67, 0xD2, 0x3F, 0x1C, 0xE7}}

type IFileSavePickerActivatedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_CallerPackageFamilyName() string
	Get_EnterpriseId() string
}

type IFileSavePickerActivatedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_CallerPackageFamilyName uintptr
	Get_EnterpriseId            uintptr
}

type IFileSavePickerActivatedEventArgs2 struct {
	win32.IInspectable
}

func (this *IFileSavePickerActivatedEventArgs2) Vtbl() *IFileSavePickerActivatedEventArgs2Vtbl {
	return (*IFileSavePickerActivatedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFileSavePickerActivatedEventArgs2) Get_CallerPackageFamilyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CallerPackageFamilyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IFileSavePickerActivatedEventArgs2) Get_EnterpriseId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EnterpriseId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 2C846FE1-3BAD-4F33-8C8B-E46FAE824B4B
var IID_IFileSavePickerContinuationEventArgs = syscall.GUID{0x2C846FE1, 0x3BAD, 0x4F33,
	[8]byte{0x8C, 0x8B, 0xE4, 0x6F, 0xAE, 0x82, 0x4B, 0x4B}}

type IFileSavePickerContinuationEventArgsInterface interface {
	win32.IInspectableInterface
	Get_File() *IStorageFile
}

type IFileSavePickerContinuationEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_File uintptr
}

type IFileSavePickerContinuationEventArgs struct {
	win32.IInspectable
}

func (this *IFileSavePickerContinuationEventArgs) Vtbl() *IFileSavePickerContinuationEventArgsVtbl {
	return (*IFileSavePickerContinuationEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFileSavePickerContinuationEventArgs) Get_File() *IStorageFile {
	var _result *IStorageFile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_File, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 51882366-9F4B-498F-BEB0-42684F6E1C29
var IID_IFolderPickerContinuationEventArgs = syscall.GUID{0x51882366, 0x9F4B, 0x498F,
	[8]byte{0xBE, 0xB0, 0x42, 0x68, 0x4F, 0x6E, 0x1C, 0x29}}

type IFolderPickerContinuationEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Folder() *IStorageFolder
}

type IFolderPickerContinuationEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Folder uintptr
}

type IFolderPickerContinuationEventArgs struct {
	win32.IInspectable
}

func (this *IFolderPickerContinuationEventArgs) Vtbl() *IFolderPickerContinuationEventArgsVtbl {
	return (*IFolderPickerContinuationEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFolderPickerContinuationEventArgs) Get_Folder() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Folder, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FBC93E26-A14A-4B4F-82B0-33BED920AF52
var IID_ILaunchActivatedEventArgs = syscall.GUID{0xFBC93E26, 0xA14A, 0x4B4F,
	[8]byte{0x82, 0xB0, 0x33, 0xBE, 0xD9, 0x20, 0xAF, 0x52}}

type ILaunchActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Arguments() string
	Get_TileId() string
}

type ILaunchActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Arguments uintptr
	Get_TileId    uintptr
}

type ILaunchActivatedEventArgs struct {
	win32.IInspectable
}

func (this *ILaunchActivatedEventArgs) Vtbl() *ILaunchActivatedEventArgsVtbl {
	return (*ILaunchActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILaunchActivatedEventArgs) Get_Arguments() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Arguments, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILaunchActivatedEventArgs) Get_TileId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TileId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 0FD37EBC-9DC9-46B5-9ACE-BD95D4565345
var IID_ILaunchActivatedEventArgs2 = syscall.GUID{0x0FD37EBC, 0x9DC9, 0x46B5,
	[8]byte{0x9A, 0xCE, 0xBD, 0x95, 0xD4, 0x56, 0x53, 0x45}}

type ILaunchActivatedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_TileActivatedInfo() *ITileActivatedInfo
}

type ILaunchActivatedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_TileActivatedInfo uintptr
}

type ILaunchActivatedEventArgs2 struct {
	win32.IInspectable
}

func (this *ILaunchActivatedEventArgs2) Vtbl() *ILaunchActivatedEventArgs2Vtbl {
	return (*ILaunchActivatedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILaunchActivatedEventArgs2) Get_TileActivatedInfo() *ITileActivatedInfo {
	var _result *ITileActivatedInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TileActivatedInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3CA77966-6108-4A41-8220-EE7D133C8532
var IID_ILockScreenActivatedEventArgs = syscall.GUID{0x3CA77966, 0x6108, 0x4A41,
	[8]byte{0x82, 0x20, 0xEE, 0x7D, 0x13, 0x3C, 0x85, 0x32}}

type ILockScreenActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Info() interface{}
}

type ILockScreenActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Info uintptr
}

type ILockScreenActivatedEventArgs struct {
	win32.IInspectable
}

func (this *ILockScreenActivatedEventArgs) Vtbl() *ILockScreenActivatedEventArgsVtbl {
	return (*ILockScreenActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILockScreenActivatedEventArgs) Get_Info() interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Info, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 06F37FBE-B5F2-448B-B13E-E328AC1C516A
var IID_ILockScreenCallActivatedEventArgs = syscall.GUID{0x06F37FBE, 0xB5F2, 0x448B,
	[8]byte{0xB1, 0x3E, 0xE3, 0x28, 0xAC, 0x1C, 0x51, 0x6A}}

type ILockScreenCallActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_CallUI() unsafe.Pointer
}

type ILockScreenCallActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_CallUI uintptr
}

type ILockScreenCallActivatedEventArgs struct {
	win32.IInspectable
}

func (this *ILockScreenCallActivatedEventArgs) Vtbl() *ILockScreenCallActivatedEventArgsVtbl {
	return (*ILockScreenCallActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILockScreenCallActivatedEventArgs) Get_CallUI() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CallUI, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 54615221-A3C1-4CED-B62F-8C60523619AD
var IID_IPhoneCallActivatedEventArgs = syscall.GUID{0x54615221, 0xA3C1, 0x4CED,
	[8]byte{0xB6, 0x2F, 0x8C, 0x60, 0x52, 0x36, 0x19, 0xAD}}

type IPhoneCallActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_LineId() syscall.GUID
}

type IPhoneCallActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_LineId uintptr
}

type IPhoneCallActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IPhoneCallActivatedEventArgs) Vtbl() *IPhoneCallActivatedEventArgsVtbl {
	return (*IPhoneCallActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPhoneCallActivatedEventArgs) Get_LineId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LineId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 360DEFB9-A9D3-4984-A4ED-9EC734604921
var IID_IPickerReturnedActivatedEventArgs = syscall.GUID{0x360DEFB9, 0xA9D3, 0x4984,
	[8]byte{0xA4, 0xED, 0x9E, 0xC7, 0x34, 0x60, 0x49, 0x21}}

type IPickerReturnedActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_PickerOperationId() string
}

type IPickerReturnedActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_PickerOperationId uintptr
}

type IPickerReturnedActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IPickerReturnedActivatedEventArgs) Vtbl() *IPickerReturnedActivatedEventArgsVtbl {
	return (*IPickerReturnedActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPickerReturnedActivatedEventArgs) Get_PickerOperationId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PickerOperationId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 0C44717B-19F7-48D6-B046-CF22826EAA74
var IID_IPrelaunchActivatedEventArgs = syscall.GUID{0x0C44717B, 0x19F7, 0x48D6,
	[8]byte{0xB0, 0x46, 0xCF, 0x22, 0x82, 0x6E, 0xAA, 0x74}}

type IPrelaunchActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_PrelaunchActivated() bool
}

type IPrelaunchActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_PrelaunchActivated uintptr
}

type IPrelaunchActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IPrelaunchActivatedEventArgs) Vtbl() *IPrelaunchActivatedEventArgsVtbl {
	return (*IPrelaunchActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrelaunchActivatedEventArgs) Get_PrelaunchActivated() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrelaunchActivated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3F57E78B-F2AC-4619-8302-EF855E1C9B90
var IID_IPrint3DWorkflowActivatedEventArgs = syscall.GUID{0x3F57E78B, 0xF2AC, 0x4619,
	[8]byte{0x83, 0x02, 0xEF, 0x85, 0x5E, 0x1C, 0x9B, 0x90}}

type IPrint3DWorkflowActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Workflow() unsafe.Pointer
}

type IPrint3DWorkflowActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Workflow uintptr
}

type IPrint3DWorkflowActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IPrint3DWorkflowActivatedEventArgs) Vtbl() *IPrint3DWorkflowActivatedEventArgsVtbl {
	return (*IPrint3DWorkflowActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrint3DWorkflowActivatedEventArgs) Get_Workflow() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Workflow, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// EE30A0C9-CE56-4865-BA8E-8954AC271107
var IID_IPrintTaskSettingsActivatedEventArgs = syscall.GUID{0xEE30A0C9, 0xCE56, 0x4865,
	[8]byte{0xBA, 0x8E, 0x89, 0x54, 0xAC, 0x27, 0x11, 0x07}}

type IPrintTaskSettingsActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Configuration() unsafe.Pointer
}

type IPrintTaskSettingsActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Configuration uintptr
}

type IPrintTaskSettingsActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IPrintTaskSettingsActivatedEventArgs) Vtbl() *IPrintTaskSettingsActivatedEventArgsVtbl {
	return (*IPrintTaskSettingsActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrintTaskSettingsActivatedEventArgs) Get_Configuration() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Configuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6095F4DD-B7C0-46AB-81FE-D90F36D00D24
var IID_IProtocolActivatedEventArgs = syscall.GUID{0x6095F4DD, 0xB7C0, 0x46AB,
	[8]byte{0x81, 0xFE, 0xD9, 0x0F, 0x36, 0xD0, 0x0D, 0x24}}

type IProtocolActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Uri() *IUriRuntimeClass
}

type IProtocolActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Uri uintptr
}

type IProtocolActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IProtocolActivatedEventArgs) Vtbl() *IProtocolActivatedEventArgsVtbl {
	return (*IProtocolActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProtocolActivatedEventArgs) Get_Uri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Uri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D84A0C12-5C8F-438C-83CB-C28FCC0B2FDB
var IID_IProtocolActivatedEventArgsWithCallerPackageFamilyNameAndData = syscall.GUID{0xD84A0C12, 0x5C8F, 0x438C,
	[8]byte{0x83, 0xCB, 0xC2, 0x8F, 0xCC, 0x0B, 0x2F, 0xDB}}

type IProtocolActivatedEventArgsWithCallerPackageFamilyNameAndDataInterface interface {
	win32.IInspectableInterface
	Get_CallerPackageFamilyName() string
	Get_Data() *IPropertySet
}

type IProtocolActivatedEventArgsWithCallerPackageFamilyNameAndDataVtbl struct {
	win32.IInspectableVtbl
	Get_CallerPackageFamilyName uintptr
	Get_Data                    uintptr
}

type IProtocolActivatedEventArgsWithCallerPackageFamilyNameAndData struct {
	win32.IInspectable
}

func (this *IProtocolActivatedEventArgsWithCallerPackageFamilyNameAndData) Vtbl() *IProtocolActivatedEventArgsWithCallerPackageFamilyNameAndDataVtbl {
	return (*IProtocolActivatedEventArgsWithCallerPackageFamilyNameAndDataVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProtocolActivatedEventArgsWithCallerPackageFamilyNameAndData) Get_CallerPackageFamilyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CallerPackageFamilyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IProtocolActivatedEventArgsWithCallerPackageFamilyNameAndData) Get_Data() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E75132C2-7AE7-4517-80AC-DBE8D7CC5B9C
var IID_IProtocolForResultsActivatedEventArgs = syscall.GUID{0xE75132C2, 0x7AE7, 0x4517,
	[8]byte{0x80, 0xAC, 0xDB, 0xE8, 0xD7, 0xCC, 0x5B, 0x9C}}

type IProtocolForResultsActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ProtocolForResultsOperation() *IProtocolForResultsOperation
}

type IProtocolForResultsActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ProtocolForResultsOperation uintptr
}

type IProtocolForResultsActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IProtocolForResultsActivatedEventArgs) Vtbl() *IProtocolForResultsActivatedEventArgsVtbl {
	return (*IProtocolForResultsActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProtocolForResultsActivatedEventArgs) Get_ProtocolForResultsOperation() *IProtocolForResultsOperation {
	var _result *IProtocolForResultsOperation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtocolForResultsOperation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E0B7AC81-BFC3-4344-A5DA-19FD5A27BAAE
var IID_IRestrictedLaunchActivatedEventArgs = syscall.GUID{0xE0B7AC81, 0xBFC3, 0x4344,
	[8]byte{0xA5, 0xDA, 0x19, 0xFD, 0x5A, 0x27, 0xBA, 0xAE}}

type IRestrictedLaunchActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_SharedContext() interface{}
}

type IRestrictedLaunchActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_SharedContext uintptr
}

type IRestrictedLaunchActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IRestrictedLaunchActivatedEventArgs) Vtbl() *IRestrictedLaunchActivatedEventArgsVtbl {
	return (*IRestrictedLaunchActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRestrictedLaunchActivatedEventArgs) Get_SharedContext() interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SharedContext, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 8CB36951-58C8-43E3-94BC-41D33F8B630E
var IID_ISearchActivatedEventArgs = syscall.GUID{0x8CB36951, 0x58C8, 0x43E3,
	[8]byte{0x94, 0xBC, 0x41, 0xD3, 0x3F, 0x8B, 0x63, 0x0E}}

type ISearchActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_QueryText() string
	Get_Language() string
}

type ISearchActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_QueryText uintptr
	Get_Language  uintptr
}

type ISearchActivatedEventArgs struct {
	win32.IInspectable
}

func (this *ISearchActivatedEventArgs) Vtbl() *ISearchActivatedEventArgsVtbl {
	return (*ISearchActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISearchActivatedEventArgs) Get_QueryText() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_QueryText, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISearchActivatedEventArgs) Get_Language() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Language, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// C09F33DA-08AB-4931-9B7C-451025F21F81
var IID_ISearchActivatedEventArgsWithLinguisticDetails = syscall.GUID{0xC09F33DA, 0x08AB, 0x4931,
	[8]byte{0x9B, 0x7C, 0x45, 0x10, 0x25, 0xF2, 0x1F, 0x81}}

type ISearchActivatedEventArgsWithLinguisticDetailsInterface interface {
	win32.IInspectableInterface
	Get_LinguisticDetails() unsafe.Pointer
}

type ISearchActivatedEventArgsWithLinguisticDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_LinguisticDetails uintptr
}

type ISearchActivatedEventArgsWithLinguisticDetails struct {
	win32.IInspectable
}

func (this *ISearchActivatedEventArgsWithLinguisticDetails) Vtbl() *ISearchActivatedEventArgsWithLinguisticDetailsVtbl {
	return (*ISearchActivatedEventArgsWithLinguisticDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISearchActivatedEventArgsWithLinguisticDetails) Get_LinguisticDetails() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LinguisticDetails, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4BDAF9C8-CDB2-4ACB-BFC3-6648563378EC
var IID_IShareTargetActivatedEventArgs = syscall.GUID{0x4BDAF9C8, 0xCDB2, 0x4ACB,
	[8]byte{0xBF, 0xC3, 0x66, 0x48, 0x56, 0x33, 0x78, 0xEC}}

type IShareTargetActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ShareOperation() unsafe.Pointer
}

type IShareTargetActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ShareOperation uintptr
}

type IShareTargetActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IShareTargetActivatedEventArgs) Vtbl() *IShareTargetActivatedEventArgsVtbl {
	return (*IShareTargetActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IShareTargetActivatedEventArgs) Get_ShareOperation() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShareOperation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// CA4D975C-D4D6-43F0-97C0-0833C6391C24
var IID_ISplashScreen = syscall.GUID{0xCA4D975C, 0xD4D6, 0x43F0,
	[8]byte{0x97, 0xC0, 0x08, 0x33, 0xC6, 0x39, 0x1C, 0x24}}

type ISplashScreenInterface interface {
	win32.IInspectableInterface
	Get_ImageLocation() Rect
	Add_Dismissed(handler TypedEventHandler[*ISplashScreen, interface{}]) EventRegistrationToken
	Remove_Dismissed(cookie EventRegistrationToken)
}

type ISplashScreenVtbl struct {
	win32.IInspectableVtbl
	Get_ImageLocation uintptr
	Add_Dismissed     uintptr
	Remove_Dismissed  uintptr
}

type ISplashScreen struct {
	win32.IInspectable
}

func (this *ISplashScreen) Vtbl() *ISplashScreenVtbl {
	return (*ISplashScreenVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISplashScreen) Get_ImageLocation() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ImageLocation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISplashScreen) Add_Dismissed(handler TypedEventHandler[*ISplashScreen, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Dismissed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISplashScreen) Remove_Dismissed(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Dismissed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

// 03B11A58-5276-4D91-8621-54611864D5FA
var IID_IStartupTaskActivatedEventArgs = syscall.GUID{0x03B11A58, 0x5276, 0x4D91,
	[8]byte{0x86, 0x21, 0x54, 0x61, 0x18, 0x64, 0xD5, 0xFA}}

type IStartupTaskActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_TaskId() string
}

type IStartupTaskActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_TaskId uintptr
}

type IStartupTaskActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IStartupTaskActivatedEventArgs) Vtbl() *IStartupTaskActivatedEventArgsVtbl {
	return (*IStartupTaskActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStartupTaskActivatedEventArgs) Get_TaskId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TaskId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 80E4A3B1-3980-4F17-B738-89194E0B8F65
var IID_ITileActivatedInfo = syscall.GUID{0x80E4A3B1, 0x3980, 0x4F17,
	[8]byte{0xB7, 0x38, 0x89, 0x19, 0x4E, 0x0B, 0x8F, 0x65}}

type ITileActivatedInfoInterface interface {
	win32.IInspectableInterface
	Get_RecentlyShownNotifications() *IVectorView[*IShownTileNotification]
}

type ITileActivatedInfoVtbl struct {
	win32.IInspectableVtbl
	Get_RecentlyShownNotifications uintptr
}

type ITileActivatedInfo struct {
	win32.IInspectable
}

func (this *ITileActivatedInfo) Vtbl() *ITileActivatedInfoVtbl {
	return (*ITileActivatedInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITileActivatedInfo) Get_RecentlyShownNotifications() *IVectorView[*IShownTileNotification] {
	var _result *IVectorView[*IShownTileNotification]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RecentlyShownNotifications, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 92A86F82-5290-431D-BE85-C4AAEEB8685F
var IID_IToastNotificationActivatedEventArgs = syscall.GUID{0x92A86F82, 0x5290, 0x431D,
	[8]byte{0xBE, 0x85, 0xC4, 0xAA, 0xEE, 0xB8, 0x68, 0x5F}}

type IToastNotificationActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Argument() string
	Get_UserInput() *IPropertySet
}

type IToastNotificationActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Argument  uintptr
	Get_UserInput uintptr
}

type IToastNotificationActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IToastNotificationActivatedEventArgs) Vtbl() *IToastNotificationActivatedEventArgsVtbl {
	return (*IToastNotificationActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastNotificationActivatedEventArgs) Get_Argument() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Argument, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IToastNotificationActivatedEventArgs) Get_UserInput() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1BC9F723-8EF1-4A51-A63A-FE711EEAB607
var IID_IUserDataAccountProviderActivatedEventArgs = syscall.GUID{0x1BC9F723, 0x8EF1, 0x4A51,
	[8]byte{0xA6, 0x3A, 0xFE, 0x71, 0x1E, 0xEA, 0xB6, 0x07}}

type IUserDataAccountProviderActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Operation() unsafe.Pointer
}

type IUserDataAccountProviderActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Operation uintptr
}

type IUserDataAccountProviderActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IUserDataAccountProviderActivatedEventArgs) Vtbl() *IUserDataAccountProviderActivatedEventArgsVtbl {
	return (*IUserDataAccountProviderActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUserDataAccountProviderActivatedEventArgs) Get_Operation() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Operation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 33F288A6-5C2C-4D27-BAC7-7536088F1219
var IID_IViewSwitcherProvider = syscall.GUID{0x33F288A6, 0x5C2C, 0x4D27,
	[8]byte{0xBA, 0xC7, 0x75, 0x36, 0x08, 0x8F, 0x12, 0x19}}

type IViewSwitcherProviderInterface interface {
	win32.IInspectableInterface
	Get_ViewSwitcher() *IActivationViewSwitcher
}

type IViewSwitcherProviderVtbl struct {
	win32.IInspectableVtbl
	Get_ViewSwitcher uintptr
}

type IViewSwitcherProvider struct {
	win32.IInspectable
}

func (this *IViewSwitcherProvider) Vtbl() *IViewSwitcherProviderVtbl {
	return (*IViewSwitcherProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IViewSwitcherProvider) Get_ViewSwitcher() *IActivationViewSwitcher {
	var _result *IActivationViewSwitcher
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ViewSwitcher, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AB92DCFD-8D43-4DE6-9775-20704B581B00
var IID_IVoiceCommandActivatedEventArgs = syscall.GUID{0xAB92DCFD, 0x8D43, 0x4DE6,
	[8]byte{0x97, 0x75, 0x20, 0x70, 0x4B, 0x58, 0x1B, 0x00}}

type IVoiceCommandActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Result() *ISpeechRecognitionResult
}

type IVoiceCommandActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Result uintptr
}

type IVoiceCommandActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IVoiceCommandActivatedEventArgs) Vtbl() *IVoiceCommandActivatedEventArgsVtbl {
	return (*IVoiceCommandActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVoiceCommandActivatedEventArgs) Get_Result() *ISpeechRecognitionResult {
	var _result *ISpeechRecognitionResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Result, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FCFC027B-1A1A-4D22-923F-AE6F45FA52D9
var IID_IWalletActionActivatedEventArgs = syscall.GUID{0xFCFC027B, 0x1A1A, 0x4D22,
	[8]byte{0x92, 0x3F, 0xAE, 0x6F, 0x45, 0xFA, 0x52, 0xD9}}

type IWalletActionActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ItemId() string
	Get_ActionKind() unsafe.Pointer
	Get_ActionId() string
}

type IWalletActionActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ItemId     uintptr
	Get_ActionKind uintptr
	Get_ActionId   uintptr
}

type IWalletActionActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IWalletActionActivatedEventArgs) Vtbl() *IWalletActionActivatedEventArgsVtbl {
	return (*IWalletActionActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWalletActionActivatedEventArgs) Get_ItemId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ItemId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWalletActionActivatedEventArgs) Get_ActionKind() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActionKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWalletActionActivatedEventArgs) Get_ActionId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActionId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 72B71774-98EA-4CCF-9752-46D9051004F1
var IID_IWebAccountProviderActivatedEventArgs = syscall.GUID{0x72B71774, 0x98EA, 0x4CCF,
	[8]byte{0x97, 0x52, 0x46, 0xD9, 0x05, 0x10, 0x04, 0xF1}}

type IWebAccountProviderActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Operation() *IWebAccountProviderOperation
}

type IWebAccountProviderActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Operation uintptr
}

type IWebAccountProviderActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IWebAccountProviderActivatedEventArgs) Vtbl() *IWebAccountProviderActivatedEventArgsVtbl {
	return (*IWebAccountProviderActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountProviderActivatedEventArgs) Get_Operation() *IWebAccountProviderOperation {
	var _result *IWebAccountProviderOperation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Operation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 75DDA3D4-7714-453D-B7FF-B95E3A1709DA
var IID_IWebAuthenticationBrokerContinuationEventArgs = syscall.GUID{0x75DDA3D4, 0x7714, 0x453D,
	[8]byte{0xB7, 0xFF, 0xB9, 0x5E, 0x3A, 0x17, 0x09, 0xDA}}

type IWebAuthenticationBrokerContinuationEventArgsInterface interface {
	win32.IInspectableInterface
	Get_WebAuthenticationResult() unsafe.Pointer
}

type IWebAuthenticationBrokerContinuationEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_WebAuthenticationResult uintptr
}

type IWebAuthenticationBrokerContinuationEventArgs struct {
	win32.IInspectable
}

func (this *IWebAuthenticationBrokerContinuationEventArgs) Vtbl() *IWebAuthenticationBrokerContinuationEventArgsVtbl {
	return (*IWebAuthenticationBrokerContinuationEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAuthenticationBrokerContinuationEventArgs) Get_WebAuthenticationResult() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WebAuthenticationResult, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type LockScreenActivatedEventArgs struct {
	RtClass
	*ILockScreenActivatedEventArgs
}
