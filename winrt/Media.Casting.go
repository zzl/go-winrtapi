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
type CastingConnectionErrorStatus int32

const (
	CastingConnectionErrorStatus_Succeeded               CastingConnectionErrorStatus = 0
	CastingConnectionErrorStatus_DeviceDidNotRespond     CastingConnectionErrorStatus = 1
	CastingConnectionErrorStatus_DeviceError             CastingConnectionErrorStatus = 2
	CastingConnectionErrorStatus_DeviceLocked            CastingConnectionErrorStatus = 3
	CastingConnectionErrorStatus_ProtectedPlaybackFailed CastingConnectionErrorStatus = 4
	CastingConnectionErrorStatus_InvalidCastingSource    CastingConnectionErrorStatus = 5
	CastingConnectionErrorStatus_Unknown                 CastingConnectionErrorStatus = 6
)

// enum
type CastingConnectionState int32

const (
	CastingConnectionState_Disconnected  CastingConnectionState = 0
	CastingConnectionState_Connected     CastingConnectionState = 1
	CastingConnectionState_Rendering     CastingConnectionState = 2
	CastingConnectionState_Disconnecting CastingConnectionState = 3
	CastingConnectionState_Connecting    CastingConnectionState = 4
)

// enum
// flags
type CastingPlaybackTypes uint32

const (
	CastingPlaybackTypes_None    CastingPlaybackTypes = 0
	CastingPlaybackTypes_Audio   CastingPlaybackTypes = 1
	CastingPlaybackTypes_Video   CastingPlaybackTypes = 2
	CastingPlaybackTypes_Picture CastingPlaybackTypes = 4
)

// interfaces

// CD951653-C2F1-4498-8B78-5FB4CD3640DD
var IID_ICastingConnection = syscall.GUID{0xCD951653, 0xC2F1, 0x4498,
	[8]byte{0x8B, 0x78, 0x5F, 0xB4, 0xCD, 0x36, 0x40, 0xDD}}

type ICastingConnectionInterface interface {
	win32.IInspectableInterface
	Get_State() CastingConnectionState
	Get_Device() *ICastingDevice
	Get_Source() *ICastingSource
	Put_Source(value *ICastingSource)
	Add_StateChanged(handler TypedEventHandler[*ICastingConnection, interface{}]) EventRegistrationToken
	Remove_StateChanged(token EventRegistrationToken)
	Add_ErrorOccurred(handler TypedEventHandler[*ICastingConnection, *ICastingConnectionErrorOccurredEventArgs]) EventRegistrationToken
	Remove_ErrorOccurred(token EventRegistrationToken)
	RequestStartCastingAsync(value *ICastingSource) *IAsyncOperation[CastingConnectionErrorStatus]
	DisconnectAsync() *IAsyncOperation[CastingConnectionErrorStatus]
}

type ICastingConnectionVtbl struct {
	win32.IInspectableVtbl
	Get_State                uintptr
	Get_Device               uintptr
	Get_Source               uintptr
	Put_Source               uintptr
	Add_StateChanged         uintptr
	Remove_StateChanged      uintptr
	Add_ErrorOccurred        uintptr
	Remove_ErrorOccurred     uintptr
	RequestStartCastingAsync uintptr
	DisconnectAsync          uintptr
}

type ICastingConnection struct {
	win32.IInspectable
}

func (this *ICastingConnection) Vtbl() *ICastingConnectionVtbl {
	return (*ICastingConnectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICastingConnection) Get_State() CastingConnectionState {
	var _result CastingConnectionState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICastingConnection) Get_Device() *ICastingDevice {
	var _result *ICastingDevice
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Device, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICastingConnection) Get_Source() *ICastingSource {
	var _result *ICastingSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICastingConnection) Put_Source(value *ICastingSource) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICastingConnection) Add_StateChanged(handler TypedEventHandler[*ICastingConnection, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICastingConnection) Remove_StateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICastingConnection) Add_ErrorOccurred(handler TypedEventHandler[*ICastingConnection, *ICastingConnectionErrorOccurredEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ErrorOccurred, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICastingConnection) Remove_ErrorOccurred(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ErrorOccurred, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICastingConnection) RequestStartCastingAsync(value *ICastingSource) *IAsyncOperation[CastingConnectionErrorStatus] {
	var _result *IAsyncOperation[CastingConnectionErrorStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestStartCastingAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICastingConnection) DisconnectAsync() *IAsyncOperation[CastingConnectionErrorStatus] {
	var _result *IAsyncOperation[CastingConnectionErrorStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DisconnectAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A7FB3C69-8719-4F00-81FB-961863C79A32
var IID_ICastingConnectionErrorOccurredEventArgs = syscall.GUID{0xA7FB3C69, 0x8719, 0x4F00,
	[8]byte{0x81, 0xFB, 0x96, 0x18, 0x63, 0xC7, 0x9A, 0x32}}

type ICastingConnectionErrorOccurredEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ErrorStatus() CastingConnectionErrorStatus
	Get_Message() string
}

type ICastingConnectionErrorOccurredEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ErrorStatus uintptr
	Get_Message     uintptr
}

type ICastingConnectionErrorOccurredEventArgs struct {
	win32.IInspectable
}

func (this *ICastingConnectionErrorOccurredEventArgs) Vtbl() *ICastingConnectionErrorOccurredEventArgsVtbl {
	return (*ICastingConnectionErrorOccurredEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICastingConnectionErrorOccurredEventArgs) Get_ErrorStatus() CastingConnectionErrorStatus {
	var _result CastingConnectionErrorStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICastingConnectionErrorOccurredEventArgs) Get_Message() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Message, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// DE721C83-4A43-4AD1-A6D2-2492A796C3F2
var IID_ICastingDevice = syscall.GUID{0xDE721C83, 0x4A43, 0x4AD1,
	[8]byte{0xA6, 0xD2, 0x24, 0x92, 0xA7, 0x96, 0xC3, 0xF2}}

type ICastingDeviceInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Get_FriendlyName() string
	Get_Icon() *IRandomAccessStreamWithContentType
	GetSupportedCastingPlaybackTypesAsync() *IAsyncOperation[CastingPlaybackTypes]
	CreateCastingConnection() *ICastingConnection
}

type ICastingDeviceVtbl struct {
	win32.IInspectableVtbl
	Get_Id                                uintptr
	Get_FriendlyName                      uintptr
	Get_Icon                              uintptr
	GetSupportedCastingPlaybackTypesAsync uintptr
	CreateCastingConnection               uintptr
}

type ICastingDevice struct {
	win32.IInspectable
}

func (this *ICastingDevice) Vtbl() *ICastingDeviceVtbl {
	return (*ICastingDeviceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICastingDevice) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICastingDevice) Get_FriendlyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FriendlyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICastingDevice) Get_Icon() *IRandomAccessStreamWithContentType {
	var _result *IRandomAccessStreamWithContentType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Icon, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICastingDevice) GetSupportedCastingPlaybackTypesAsync() *IAsyncOperation[CastingPlaybackTypes] {
	var _result *IAsyncOperation[CastingPlaybackTypes]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSupportedCastingPlaybackTypesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICastingDevice) CreateCastingConnection() *ICastingConnection {
	var _result *ICastingConnection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCastingConnection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DCD39924-0591-49BE-AACB-4B82EE756A95
var IID_ICastingDevicePicker = syscall.GUID{0xDCD39924, 0x0591, 0x49BE,
	[8]byte{0xAA, 0xCB, 0x4B, 0x82, 0xEE, 0x75, 0x6A, 0x95}}

type ICastingDevicePickerInterface interface {
	win32.IInspectableInterface
	Get_Filter() *ICastingDevicePickerFilter
	Get_Appearance() *IDevicePickerAppearance
	Add_CastingDeviceSelected(handler TypedEventHandler[*ICastingDevicePicker, *ICastingDeviceSelectedEventArgs]) EventRegistrationToken
	Remove_CastingDeviceSelected(token EventRegistrationToken)
	Add_CastingDevicePickerDismissed(handler TypedEventHandler[*ICastingDevicePicker, interface{}]) EventRegistrationToken
	Remove_CastingDevicePickerDismissed(token EventRegistrationToken)
	Show(selection Rect)
	ShowWithPlacement(selection Rect, preferredPlacement Placement)
	Hide()
}

type ICastingDevicePickerVtbl struct {
	win32.IInspectableVtbl
	Get_Filter                          uintptr
	Get_Appearance                      uintptr
	Add_CastingDeviceSelected           uintptr
	Remove_CastingDeviceSelected        uintptr
	Add_CastingDevicePickerDismissed    uintptr
	Remove_CastingDevicePickerDismissed uintptr
	Show                                uintptr
	ShowWithPlacement                   uintptr
	Hide                                uintptr
}

type ICastingDevicePicker struct {
	win32.IInspectable
}

func (this *ICastingDevicePicker) Vtbl() *ICastingDevicePickerVtbl {
	return (*ICastingDevicePickerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICastingDevicePicker) Get_Filter() *ICastingDevicePickerFilter {
	var _result *ICastingDevicePickerFilter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Filter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICastingDevicePicker) Get_Appearance() *IDevicePickerAppearance {
	var _result *IDevicePickerAppearance
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Appearance, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICastingDevicePicker) Add_CastingDeviceSelected(handler TypedEventHandler[*ICastingDevicePicker, *ICastingDeviceSelectedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_CastingDeviceSelected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICastingDevicePicker) Remove_CastingDeviceSelected(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_CastingDeviceSelected, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICastingDevicePicker) Add_CastingDevicePickerDismissed(handler TypedEventHandler[*ICastingDevicePicker, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_CastingDevicePickerDismissed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICastingDevicePicker) Remove_CastingDevicePickerDismissed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_CastingDevicePickerDismissed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICastingDevicePicker) Show(selection Rect) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Show, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&selection)))
	_ = _hr
}

func (this *ICastingDevicePicker) ShowWithPlacement(selection Rect, preferredPlacement Placement) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowWithPlacement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&selection)), uintptr(preferredPlacement))
	_ = _hr
}

func (this *ICastingDevicePicker) Hide() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Hide, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// BE8C619C-B563-4354-AE33-9FDAAD8C6291
var IID_ICastingDevicePickerFilter = syscall.GUID{0xBE8C619C, 0xB563, 0x4354,
	[8]byte{0xAE, 0x33, 0x9F, 0xDA, 0xAD, 0x8C, 0x62, 0x91}}

type ICastingDevicePickerFilterInterface interface {
	win32.IInspectableInterface
	Get_SupportsAudio() bool
	Put_SupportsAudio(value bool)
	Get_SupportsVideo() bool
	Put_SupportsVideo(value bool)
	Get_SupportsPictures() bool
	Put_SupportsPictures(value bool)
	Get_SupportedCastingSources() *IVector[*ICastingSource]
}

type ICastingDevicePickerFilterVtbl struct {
	win32.IInspectableVtbl
	Get_SupportsAudio           uintptr
	Put_SupportsAudio           uintptr
	Get_SupportsVideo           uintptr
	Put_SupportsVideo           uintptr
	Get_SupportsPictures        uintptr
	Put_SupportsPictures        uintptr
	Get_SupportedCastingSources uintptr
}

type ICastingDevicePickerFilter struct {
	win32.IInspectable
}

func (this *ICastingDevicePickerFilter) Vtbl() *ICastingDevicePickerFilterVtbl {
	return (*ICastingDevicePickerFilterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICastingDevicePickerFilter) Get_SupportsAudio() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportsAudio, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICastingDevicePickerFilter) Put_SupportsAudio(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SupportsAudio, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICastingDevicePickerFilter) Get_SupportsVideo() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportsVideo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICastingDevicePickerFilter) Put_SupportsVideo(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SupportsVideo, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICastingDevicePickerFilter) Get_SupportsPictures() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportsPictures, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICastingDevicePickerFilter) Put_SupportsPictures(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SupportsPictures, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICastingDevicePickerFilter) Get_SupportedCastingSources() *IVector[*ICastingSource] {
	var _result *IVector[*ICastingSource]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedCastingSources, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DC439E86-DD57-4D0D-9400-AF45E4FB3663
var IID_ICastingDeviceSelectedEventArgs = syscall.GUID{0xDC439E86, 0xDD57, 0x4D0D,
	[8]byte{0x94, 0x00, 0xAF, 0x45, 0xE4, 0xFB, 0x36, 0x63}}

type ICastingDeviceSelectedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_SelectedCastingDevice() *ICastingDevice
}

type ICastingDeviceSelectedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_SelectedCastingDevice uintptr
}

type ICastingDeviceSelectedEventArgs struct {
	win32.IInspectable
}

func (this *ICastingDeviceSelectedEventArgs) Vtbl() *ICastingDeviceSelectedEventArgsVtbl {
	return (*ICastingDeviceSelectedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICastingDeviceSelectedEventArgs) Get_SelectedCastingDevice() *ICastingDevice {
	var _result *ICastingDevice
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectedCastingDevice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E7D958D7-4D13-4237-A365-4C4F6A4CFD2F
var IID_ICastingDeviceStatics = syscall.GUID{0xE7D958D7, 0x4D13, 0x4237,
	[8]byte{0xA3, 0x65, 0x4C, 0x4F, 0x6A, 0x4C, 0xFD, 0x2F}}

type ICastingDeviceStaticsInterface interface {
	win32.IInspectableInterface
	GetDeviceSelector(type_ CastingPlaybackTypes) string
	GetDeviceSelectorFromCastingSourceAsync(castingSource *ICastingSource) *IAsyncOperation[string]
	FromIdAsync(value string) *IAsyncOperation[*ICastingDevice]
	DeviceInfoSupportsCastingAsync(device *IDeviceInformation) *IAsyncOperation[bool]
}

type ICastingDeviceStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector                       uintptr
	GetDeviceSelectorFromCastingSourceAsync uintptr
	FromIdAsync                             uintptr
	DeviceInfoSupportsCastingAsync          uintptr
}

type ICastingDeviceStatics struct {
	win32.IInspectable
}

func (this *ICastingDeviceStatics) Vtbl() *ICastingDeviceStaticsVtbl {
	return (*ICastingDeviceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICastingDeviceStatics) GetDeviceSelector(type_ CastingPlaybackTypes) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(type_), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICastingDeviceStatics) GetDeviceSelectorFromCastingSourceAsync(castingSource *ICastingSource) *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorFromCastingSourceAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(castingSource)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICastingDeviceStatics) FromIdAsync(value string) *IAsyncOperation[*ICastingDevice] {
	var _result *IAsyncOperation[*ICastingDevice]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICastingDeviceStatics) DeviceInfoSupportsCastingAsync(device *IDeviceInformation) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DeviceInfoSupportsCastingAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F429EA72-3467-47E6-A027-522923E9D727
var IID_ICastingSource = syscall.GUID{0xF429EA72, 0x3467, 0x47E6,
	[8]byte{0xA0, 0x27, 0x52, 0x29, 0x23, 0xE9, 0xD7, 0x27}}

type ICastingSourceInterface interface {
	win32.IInspectableInterface
	Get_PreferredSourceUri() *IUriRuntimeClass
	Put_PreferredSourceUri(value *IUriRuntimeClass)
}

type ICastingSourceVtbl struct {
	win32.IInspectableVtbl
	Get_PreferredSourceUri uintptr
	Put_PreferredSourceUri uintptr
}

type ICastingSource struct {
	win32.IInspectable
}

func (this *ICastingSource) Vtbl() *ICastingSourceVtbl {
	return (*ICastingSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICastingSource) Get_PreferredSourceUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PreferredSourceUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICastingSource) Put_PreferredSourceUri(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PreferredSourceUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// classes

type CastingConnection struct {
	RtClass
	*ICastingConnection
}

type CastingConnectionErrorOccurredEventArgs struct {
	RtClass
	*ICastingConnectionErrorOccurredEventArgs
}

type CastingDevice struct {
	RtClass
	*ICastingDevice
}

func NewICastingDeviceStatics() *ICastingDeviceStatics {
	var p *ICastingDeviceStatics
	hs := NewHStr("Windows.Media.Casting.CastingDevice")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICastingDeviceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type CastingDevicePicker struct {
	RtClass
	*ICastingDevicePicker
}

func NewCastingDevicePicker() *CastingDevicePicker {
	hs := NewHStr("Windows.Media.Casting.CastingDevicePicker")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &CastingDevicePicker{
		RtClass:              RtClass{PInspect: p},
		ICastingDevicePicker: (*ICastingDevicePicker)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type CastingDevicePickerFilter struct {
	RtClass
	*ICastingDevicePickerFilter
}

type CastingDeviceSelectedEventArgs struct {
	RtClass
	*ICastingDeviceSelectedEventArgs
}

type CastingSource struct {
	RtClass
	*ICastingSource
}
