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
type DialAppLaunchResult int32

const (
	DialAppLaunchResult_Launched       DialAppLaunchResult = 0
	DialAppLaunchResult_FailedToLaunch DialAppLaunchResult = 1
	DialAppLaunchResult_NotFound       DialAppLaunchResult = 2
	DialAppLaunchResult_NetworkFailure DialAppLaunchResult = 3
)

// enum
type DialAppState int32

const (
	DialAppState_Unknown        DialAppState = 0
	DialAppState_Stopped        DialAppState = 1
	DialAppState_Running        DialAppState = 2
	DialAppState_NetworkFailure DialAppState = 3
)

// enum
type DialAppStopResult int32

const (
	DialAppStopResult_Stopped               DialAppStopResult = 0
	DialAppStopResult_StopFailed            DialAppStopResult = 1
	DialAppStopResult_OperationNotSupported DialAppStopResult = 2
	DialAppStopResult_NetworkFailure        DialAppStopResult = 3
)

// enum
type DialDeviceDisplayStatus int32

const (
	DialDeviceDisplayStatus_None          DialDeviceDisplayStatus = 0
	DialDeviceDisplayStatus_Connecting    DialDeviceDisplayStatus = 1
	DialDeviceDisplayStatus_Connected     DialDeviceDisplayStatus = 2
	DialDeviceDisplayStatus_Disconnecting DialDeviceDisplayStatus = 3
	DialDeviceDisplayStatus_Disconnected  DialDeviceDisplayStatus = 4
	DialDeviceDisplayStatus_Error         DialDeviceDisplayStatus = 5
)

// interfaces

// 555FFBD3-45B7-49F3-BBD7-302DB6084646
var IID_IDialApp = syscall.GUID{0x555FFBD3, 0x45B7, 0x49F3,
	[8]byte{0xBB, 0xD7, 0x30, 0x2D, 0xB6, 0x08, 0x46, 0x46}}

type IDialAppInterface interface {
	win32.IInspectableInterface
	Get_AppName() string
	RequestLaunchAsync(appArgument string) *IAsyncOperation[DialAppLaunchResult]
	StopAsync() *IAsyncOperation[DialAppStopResult]
	GetAppStateAsync() *IAsyncOperation[*IDialAppStateDetails]
}

type IDialAppVtbl struct {
	win32.IInspectableVtbl
	Get_AppName        uintptr
	RequestLaunchAsync uintptr
	StopAsync          uintptr
	GetAppStateAsync   uintptr
}

type IDialApp struct {
	win32.IInspectable
}

func (this *IDialApp) Vtbl() *IDialAppVtbl {
	return (*IDialAppVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDialApp) Get_AppName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDialApp) RequestLaunchAsync(appArgument string) *IAsyncOperation[DialAppLaunchResult] {
	var _result *IAsyncOperation[DialAppLaunchResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestLaunchAsync, uintptr(unsafe.Pointer(this)), NewHStr(appArgument).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDialApp) StopAsync() *IAsyncOperation[DialAppStopResult] {
	var _result *IAsyncOperation[DialAppStopResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDialApp) GetAppStateAsync() *IAsyncOperation[*IDialAppStateDetails] {
	var _result *IAsyncOperation[*IDialAppStateDetails]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAppStateAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DDC4A4A1-F5DE-400D-BEA4-8C8466BB2961
var IID_IDialAppStateDetails = syscall.GUID{0xDDC4A4A1, 0xF5DE, 0x400D,
	[8]byte{0xBE, 0xA4, 0x8C, 0x84, 0x66, 0xBB, 0x29, 0x61}}

type IDialAppStateDetailsInterface interface {
	win32.IInspectableInterface
	Get_State() DialAppState
	Get_FullXml() string
}

type IDialAppStateDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_State   uintptr
	Get_FullXml uintptr
}

type IDialAppStateDetails struct {
	win32.IInspectable
}

func (this *IDialAppStateDetails) Vtbl() *IDialAppStateDetailsVtbl {
	return (*IDialAppStateDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDialAppStateDetails) Get_State() DialAppState {
	var _result DialAppState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDialAppStateDetails) Get_FullXml() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FullXml, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// FFF0EDAF-759F-41D2-A20A-7F29CE0B3784
var IID_IDialDevice = syscall.GUID{0xFFF0EDAF, 0x759F, 0x41D2,
	[8]byte{0xA2, 0x0A, 0x7F, 0x29, 0xCE, 0x0B, 0x37, 0x84}}

type IDialDeviceInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	GetDialApp(appName string) *IDialApp
}

type IDialDeviceVtbl struct {
	win32.IInspectableVtbl
	Get_Id     uintptr
	GetDialApp uintptr
}

type IDialDevice struct {
	win32.IInspectable
}

func (this *IDialDevice) Vtbl() *IDialDeviceVtbl {
	return (*IDialDeviceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDialDevice) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDialDevice) GetDialApp(appName string) *IDialApp {
	var _result *IDialApp
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDialApp, uintptr(unsafe.Pointer(this)), NewHStr(appName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BAB7F3D5-5BFB-4EBA-8B32-B57C5C5EE5C9
var IID_IDialDevice2 = syscall.GUID{0xBAB7F3D5, 0x5BFB, 0x4EBA,
	[8]byte{0x8B, 0x32, 0xB5, 0x7C, 0x5C, 0x5E, 0xE5, 0xC9}}

type IDialDevice2Interface interface {
	win32.IInspectableInterface
	Get_FriendlyName() string
	Get_Thumbnail() *IRandomAccessStreamReference
}

type IDialDevice2Vtbl struct {
	win32.IInspectableVtbl
	Get_FriendlyName uintptr
	Get_Thumbnail    uintptr
}

type IDialDevice2 struct {
	win32.IInspectable
}

func (this *IDialDevice2) Vtbl() *IDialDevice2Vtbl {
	return (*IDialDevice2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDialDevice2) Get_FriendlyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FriendlyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDialDevice2) Get_Thumbnail() *IRandomAccessStreamReference {
	var _result *IRandomAccessStreamReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Thumbnail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BA7E520A-FF59-4F4B-BDAC-D89F495AD6E1
var IID_IDialDevicePicker = syscall.GUID{0xBA7E520A, 0xFF59, 0x4F4B,
	[8]byte{0xBD, 0xAC, 0xD8, 0x9F, 0x49, 0x5A, 0xD6, 0xE1}}

type IDialDevicePickerInterface interface {
	win32.IInspectableInterface
	Get_Filter() *IDialDevicePickerFilter
	Get_Appearance() *IDevicePickerAppearance
	Add_DialDeviceSelected(handler TypedEventHandler[*IDialDevicePicker, *IDialDeviceSelectedEventArgs]) EventRegistrationToken
	Remove_DialDeviceSelected(token EventRegistrationToken)
	Add_DisconnectButtonClicked(handler TypedEventHandler[*IDialDevicePicker, *IDialDisconnectButtonClickedEventArgs]) EventRegistrationToken
	Remove_DisconnectButtonClicked(token EventRegistrationToken)
	Add_DialDevicePickerDismissed(handler TypedEventHandler[*IDialDevicePicker, interface{}]) EventRegistrationToken
	Remove_DialDevicePickerDismissed(token EventRegistrationToken)
	Show(selection Rect)
	ShowWithPlacement(selection Rect, preferredPlacement Placement)
	PickSingleDialDeviceAsync(selection Rect) *IAsyncOperation[*IDialDevice]
	PickSingleDialDeviceAsyncWithPlacement(selection Rect, preferredPlacement Placement) *IAsyncOperation[*IDialDevice]
	Hide()
	SetDisplayStatus(device *IDialDevice, status DialDeviceDisplayStatus)
}

type IDialDevicePickerVtbl struct {
	win32.IInspectableVtbl
	Get_Filter                             uintptr
	Get_Appearance                         uintptr
	Add_DialDeviceSelected                 uintptr
	Remove_DialDeviceSelected              uintptr
	Add_DisconnectButtonClicked            uintptr
	Remove_DisconnectButtonClicked         uintptr
	Add_DialDevicePickerDismissed          uintptr
	Remove_DialDevicePickerDismissed       uintptr
	Show                                   uintptr
	ShowWithPlacement                      uintptr
	PickSingleDialDeviceAsync              uintptr
	PickSingleDialDeviceAsyncWithPlacement uintptr
	Hide                                   uintptr
	SetDisplayStatus                       uintptr
}

type IDialDevicePicker struct {
	win32.IInspectable
}

func (this *IDialDevicePicker) Vtbl() *IDialDevicePickerVtbl {
	return (*IDialDevicePickerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDialDevicePicker) Get_Filter() *IDialDevicePickerFilter {
	var _result *IDialDevicePickerFilter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Filter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDialDevicePicker) Get_Appearance() *IDevicePickerAppearance {
	var _result *IDevicePickerAppearance
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Appearance, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDialDevicePicker) Add_DialDeviceSelected(handler TypedEventHandler[*IDialDevicePicker, *IDialDeviceSelectedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DialDeviceSelected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDialDevicePicker) Remove_DialDeviceSelected(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DialDeviceSelected, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IDialDevicePicker) Add_DisconnectButtonClicked(handler TypedEventHandler[*IDialDevicePicker, *IDialDisconnectButtonClickedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DisconnectButtonClicked, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDialDevicePicker) Remove_DisconnectButtonClicked(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DisconnectButtonClicked, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IDialDevicePicker) Add_DialDevicePickerDismissed(handler TypedEventHandler[*IDialDevicePicker, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DialDevicePickerDismissed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDialDevicePicker) Remove_DialDevicePickerDismissed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DialDevicePickerDismissed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IDialDevicePicker) Show(selection Rect) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Show, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&selection)))
	_ = _hr
}

func (this *IDialDevicePicker) ShowWithPlacement(selection Rect, preferredPlacement Placement) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowWithPlacement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&selection)), uintptr(preferredPlacement))
	_ = _hr
}

func (this *IDialDevicePicker) PickSingleDialDeviceAsync(selection Rect) *IAsyncOperation[*IDialDevice] {
	var _result *IAsyncOperation[*IDialDevice]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PickSingleDialDeviceAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&selection)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDialDevicePicker) PickSingleDialDeviceAsyncWithPlacement(selection Rect, preferredPlacement Placement) *IAsyncOperation[*IDialDevice] {
	var _result *IAsyncOperation[*IDialDevice]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PickSingleDialDeviceAsyncWithPlacement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&selection)), uintptr(preferredPlacement), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDialDevicePicker) Hide() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Hide, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IDialDevicePicker) SetDisplayStatus(device *IDialDevice, status DialDeviceDisplayStatus) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetDisplayStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(device)), uintptr(status))
	_ = _hr
}

// C17C93BA-86C0-485D-B8D6-0F9A8F641590
var IID_IDialDevicePickerFilter = syscall.GUID{0xC17C93BA, 0x86C0, 0x485D,
	[8]byte{0xB8, 0xD6, 0x0F, 0x9A, 0x8F, 0x64, 0x15, 0x90}}

type IDialDevicePickerFilterInterface interface {
	win32.IInspectableInterface
	Get_SupportedAppNames() *IVector[string]
}

type IDialDevicePickerFilterVtbl struct {
	win32.IInspectableVtbl
	Get_SupportedAppNames uintptr
}

type IDialDevicePickerFilter struct {
	win32.IInspectable
}

func (this *IDialDevicePickerFilter) Vtbl() *IDialDevicePickerFilterVtbl {
	return (*IDialDevicePickerFilterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDialDevicePickerFilter) Get_SupportedAppNames() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedAppNames, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 480B92AD-AC76-47EB-9C06-A19304DA0247
var IID_IDialDeviceSelectedEventArgs = syscall.GUID{0x480B92AD, 0xAC76, 0x47EB,
	[8]byte{0x9C, 0x06, 0xA1, 0x93, 0x04, 0xDA, 0x02, 0x47}}

type IDialDeviceSelectedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_SelectedDialDevice() *IDialDevice
}

type IDialDeviceSelectedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_SelectedDialDevice uintptr
}

type IDialDeviceSelectedEventArgs struct {
	win32.IInspectable
}

func (this *IDialDeviceSelectedEventArgs) Vtbl() *IDialDeviceSelectedEventArgsVtbl {
	return (*IDialDeviceSelectedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDialDeviceSelectedEventArgs) Get_SelectedDialDevice() *IDialDevice {
	var _result *IDialDevice
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectedDialDevice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AA69CC95-01F8-4758-8461-2BBD1CDC3CF3
var IID_IDialDeviceStatics = syscall.GUID{0xAA69CC95, 0x01F8, 0x4758,
	[8]byte{0x84, 0x61, 0x2B, 0xBD, 0x1C, 0xDC, 0x3C, 0xF3}}

type IDialDeviceStaticsInterface interface {
	win32.IInspectableInterface
	GetDeviceSelector(appName string) string
	FromIdAsync(value string) *IAsyncOperation[*IDialDevice]
	DeviceInfoSupportsDialAsync(device *IDeviceInformation) *IAsyncOperation[bool]
}

type IDialDeviceStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector           uintptr
	FromIdAsync                 uintptr
	DeviceInfoSupportsDialAsync uintptr
}

type IDialDeviceStatics struct {
	win32.IInspectable
}

func (this *IDialDeviceStatics) Vtbl() *IDialDeviceStaticsVtbl {
	return (*IDialDeviceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDialDeviceStatics) GetDeviceSelector(appName string) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), NewHStr(appName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDialDeviceStatics) FromIdAsync(value string) *IAsyncOperation[*IDialDevice] {
	var _result *IAsyncOperation[*IDialDevice]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDialDeviceStatics) DeviceInfoSupportsDialAsync(device *IDeviceInformation) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DeviceInfoSupportsDialAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 52765152-9C81-4E55-ADC2-0EBE99CDE3B6
var IID_IDialDisconnectButtonClickedEventArgs = syscall.GUID{0x52765152, 0x9C81, 0x4E55,
	[8]byte{0xAD, 0xC2, 0x0E, 0xBE, 0x99, 0xCD, 0xE3, 0xB6}}

type IDialDisconnectButtonClickedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Device() *IDialDevice
}

type IDialDisconnectButtonClickedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Device uintptr
}

type IDialDisconnectButtonClickedEventArgs struct {
	win32.IInspectable
}

func (this *IDialDisconnectButtonClickedEventArgs) Vtbl() *IDialDisconnectButtonClickedEventArgsVtbl {
	return (*IDialDisconnectButtonClickedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDialDisconnectButtonClickedEventArgs) Get_Device() *IDialDevice {
	var _result *IDialDevice
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Device, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FD3E7C57-5045-470E-B304-4DD9B13E7D11
var IID_IDialReceiverApp = syscall.GUID{0xFD3E7C57, 0x5045, 0x470E,
	[8]byte{0xB3, 0x04, 0x4D, 0xD9, 0xB1, 0x3E, 0x7D, 0x11}}

type IDialReceiverAppInterface interface {
	win32.IInspectableInterface
	GetAdditionalDataAsync() *IAsyncOperation[*IMap[string, string]]
	SetAdditionalDataAsync(additionalData *IIterable[*IKeyValuePair[string, string]]) *IAsyncAction
}

type IDialReceiverAppVtbl struct {
	win32.IInspectableVtbl
	GetAdditionalDataAsync uintptr
	SetAdditionalDataAsync uintptr
}

type IDialReceiverApp struct {
	win32.IInspectable
}

func (this *IDialReceiverApp) Vtbl() *IDialReceiverAppVtbl {
	return (*IDialReceiverAppVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDialReceiverApp) GetAdditionalDataAsync() *IAsyncOperation[*IMap[string, string]] {
	var _result *IAsyncOperation[*IMap[string, string]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAdditionalDataAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDialReceiverApp) SetAdditionalDataAsync(additionalData *IIterable[*IKeyValuePair[string, string]]) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetAdditionalDataAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(additionalData)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 530C5805-9130-42AC-A504-1977DCB2EA8A
var IID_IDialReceiverApp2 = syscall.GUID{0x530C5805, 0x9130, 0x42AC,
	[8]byte{0xA5, 0x04, 0x19, 0x77, 0xDC, 0xB2, 0xEA, 0x8A}}

type IDialReceiverApp2Interface interface {
	win32.IInspectableInterface
	GetUniqueDeviceNameAsync() *IAsyncOperation[string]
}

type IDialReceiverApp2Vtbl struct {
	win32.IInspectableVtbl
	GetUniqueDeviceNameAsync uintptr
}

type IDialReceiverApp2 struct {
	win32.IInspectable
}

func (this *IDialReceiverApp2) Vtbl() *IDialReceiverApp2Vtbl {
	return (*IDialReceiverApp2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDialReceiverApp2) GetUniqueDeviceNameAsync() *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetUniqueDeviceNameAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 53183A3C-4C36-4D02-B28A-F2A9DA38EC52
var IID_IDialReceiverAppStatics = syscall.GUID{0x53183A3C, 0x4C36, 0x4D02,
	[8]byte{0xB2, 0x8A, 0xF2, 0xA9, 0xDA, 0x38, 0xEC, 0x52}}

type IDialReceiverAppStaticsInterface interface {
	win32.IInspectableInterface
	Get_Current() *IDialReceiverApp
}

type IDialReceiverAppStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Current uintptr
}

type IDialReceiverAppStatics struct {
	win32.IInspectable
}

func (this *IDialReceiverAppStatics) Vtbl() *IDialReceiverAppStaticsVtbl {
	return (*IDialReceiverAppStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDialReceiverAppStatics) Get_Current() *IDialReceiverApp {
	var _result *IDialReceiverApp
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Current, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type DialDevicePicker struct {
	RtClass
	*IDialDevicePicker
}

func NewDialDevicePicker() *DialDevicePicker {
	hs := NewHStr("Windows.Media.DialProtocol.DialDevicePicker")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &DialDevicePicker{
		RtClass:           RtClass{PInspect: p},
		IDialDevicePicker: (*IDialDevicePicker)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type DialDevicePickerFilter struct {
	RtClass
	*IDialDevicePickerFilter
}

type DialDeviceSelectedEventArgs struct {
	RtClass
	*IDialDeviceSelectedEventArgs
}

type DialDisconnectButtonClickedEventArgs struct {
	RtClass
	*IDialDisconnectButtonClickedEventArgs
}
