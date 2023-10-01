package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type GipFirmwareUpdateStatus int32

const (
	GipFirmwareUpdateStatus_Completed GipFirmwareUpdateStatus = 0
	GipFirmwareUpdateStatus_UpToDate  GipFirmwareUpdateStatus = 1
	GipFirmwareUpdateStatus_Failed    GipFirmwareUpdateStatus = 2
)

// enum
type GipMessageClass int32

const (
	GipMessageClass_Command         GipMessageClass = 0
	GipMessageClass_LowLatency      GipMessageClass = 1
	GipMessageClass_StandardLatency GipMessageClass = 2
)

// enum
type XusbDeviceSubtype int32

const (
	XusbDeviceSubtype_Unknown         XusbDeviceSubtype = 0
	XusbDeviceSubtype_Gamepad         XusbDeviceSubtype = 1
	XusbDeviceSubtype_ArcadePad       XusbDeviceSubtype = 2
	XusbDeviceSubtype_ArcadeStick     XusbDeviceSubtype = 3
	XusbDeviceSubtype_FlightStick     XusbDeviceSubtype = 4
	XusbDeviceSubtype_Wheel           XusbDeviceSubtype = 5
	XusbDeviceSubtype_Guitar          XusbDeviceSubtype = 6
	XusbDeviceSubtype_GuitarAlternate XusbDeviceSubtype = 7
	XusbDeviceSubtype_GuitarBass      XusbDeviceSubtype = 8
	XusbDeviceSubtype_DrumKit         XusbDeviceSubtype = 9
	XusbDeviceSubtype_DancePad        XusbDeviceSubtype = 10
)

// enum
type XusbDeviceType int32

const (
	XusbDeviceType_Unknown XusbDeviceType = 0
	XusbDeviceType_Gamepad XusbDeviceType = 1
)

// structs

type GameControllerVersionInfo struct {
	Major    uint16
	Minor    uint16
	Build    uint16
	Revision uint16
}

type GipFirmwareUpdateProgress struct {
	PercentCompleted   float64
	CurrentComponentId uint32
}

// interfaces

// 69A0AE5E-758E-4CBE-ACE6-62155FE9126F
var IID_ICustomGameControllerFactory = syscall.GUID{0x69A0AE5E, 0x758E, 0x4CBE,
	[8]byte{0xAC, 0xE6, 0x62, 0x15, 0x5F, 0xE9, 0x12, 0x6F}}

type ICustomGameControllerFactoryInterface interface {
	win32.IInspectableInterface
	CreateGameController(provider *IGameControllerProvider) interface{}
	OnGameControllerAdded(value *IGameController)
	OnGameControllerRemoved(value *IGameController)
}

type ICustomGameControllerFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateGameController    uintptr
	OnGameControllerAdded   uintptr
	OnGameControllerRemoved uintptr
}

type ICustomGameControllerFactory struct {
	win32.IInspectable
}

func (this *ICustomGameControllerFactory) Vtbl() *ICustomGameControllerFactoryVtbl {
	return (*ICustomGameControllerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICustomGameControllerFactory) CreateGameController(provider *IGameControllerProvider) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateGameController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(provider)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICustomGameControllerFactory) OnGameControllerAdded(value *IGameController) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OnGameControllerAdded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICustomGameControllerFactory) OnGameControllerRemoved(value *IGameController) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OnGameControllerRemoved, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 36CB66E3-D0A1-4986-A24C-40B137DEBA9E
var IID_IGameControllerFactoryManagerStatics = syscall.GUID{0x36CB66E3, 0xD0A1, 0x4986,
	[8]byte{0xA2, 0x4C, 0x40, 0xB1, 0x37, 0xDE, 0xBA, 0x9E}}

type IGameControllerFactoryManagerStaticsInterface interface {
	win32.IInspectableInterface
	RegisterCustomFactoryForGipInterface(factory *ICustomGameControllerFactory, interfaceId syscall.GUID)
	RegisterCustomFactoryForHardwareId(factory *ICustomGameControllerFactory, hardwareVendorId uint16, hardwareProductId uint16)
	RegisterCustomFactoryForXusbType(factory *ICustomGameControllerFactory, xusbType XusbDeviceType, xusbSubtype XusbDeviceSubtype)
}

type IGameControllerFactoryManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	RegisterCustomFactoryForGipInterface uintptr
	RegisterCustomFactoryForHardwareId   uintptr
	RegisterCustomFactoryForXusbType     uintptr
}

type IGameControllerFactoryManagerStatics struct {
	win32.IInspectable
}

func (this *IGameControllerFactoryManagerStatics) Vtbl() *IGameControllerFactoryManagerStaticsVtbl {
	return (*IGameControllerFactoryManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameControllerFactoryManagerStatics) RegisterCustomFactoryForGipInterface(factory *ICustomGameControllerFactory, interfaceId syscall.GUID) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterCustomFactoryForGipInterface, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(factory)), uintptr(unsafe.Pointer(&interfaceId)))
	_ = _hr
}

func (this *IGameControllerFactoryManagerStatics) RegisterCustomFactoryForHardwareId(factory *ICustomGameControllerFactory, hardwareVendorId uint16, hardwareProductId uint16) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterCustomFactoryForHardwareId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(factory)), uintptr(hardwareVendorId), uintptr(hardwareProductId))
	_ = _hr
}

func (this *IGameControllerFactoryManagerStatics) RegisterCustomFactoryForXusbType(factory *ICustomGameControllerFactory, xusbType XusbDeviceType, xusbSubtype XusbDeviceSubtype) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterCustomFactoryForXusbType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(factory)), uintptr(xusbType), uintptr(xusbSubtype))
	_ = _hr
}

// EACE5644-19DF-4115-B32A-2793E2AEA3BB
var IID_IGameControllerFactoryManagerStatics2 = syscall.GUID{0xEACE5644, 0x19DF, 0x4115,
	[8]byte{0xB3, 0x2A, 0x27, 0x93, 0xE2, 0xAE, 0xA3, 0xBB}}

type IGameControllerFactoryManagerStatics2Interface interface {
	win32.IInspectableInterface
	TryGetFactoryControllerFromGameController(factory *ICustomGameControllerFactory, gameController *IGameController) *IGameController
}

type IGameControllerFactoryManagerStatics2Vtbl struct {
	win32.IInspectableVtbl
	TryGetFactoryControllerFromGameController uintptr
}

type IGameControllerFactoryManagerStatics2 struct {
	win32.IInspectable
}

func (this *IGameControllerFactoryManagerStatics2) Vtbl() *IGameControllerFactoryManagerStatics2Vtbl {
	return (*IGameControllerFactoryManagerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameControllerFactoryManagerStatics2) TryGetFactoryControllerFromGameController(factory *ICustomGameControllerFactory, gameController *IGameController) *IGameController {
	var _result *IGameController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetFactoryControllerFromGameController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(factory)), uintptr(unsafe.Pointer(gameController)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1FF6F922-C640-4C78-A820-9A715C558BCB
var IID_IGameControllerInputSink = syscall.GUID{0x1FF6F922, 0xC640, 0x4C78,
	[8]byte{0xA8, 0x20, 0x9A, 0x71, 0x5C, 0x55, 0x8B, 0xCB}}

type IGameControllerInputSinkInterface interface {
	win32.IInspectableInterface
	OnInputResumed(timestamp uint64)
	OnInputSuspended(timestamp uint64)
}

type IGameControllerInputSinkVtbl struct {
	win32.IInspectableVtbl
	OnInputResumed   uintptr
	OnInputSuspended uintptr
}

type IGameControllerInputSink struct {
	win32.IInspectable
}

func (this *IGameControllerInputSink) Vtbl() *IGameControllerInputSinkVtbl {
	return (*IGameControllerInputSinkVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameControllerInputSink) OnInputResumed(timestamp uint64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OnInputResumed, uintptr(unsafe.Pointer(this)), uintptr(timestamp))
	_ = _hr
}

func (this *IGameControllerInputSink) OnInputSuspended(timestamp uint64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OnInputSuspended, uintptr(unsafe.Pointer(this)), uintptr(timestamp))
	_ = _hr
}

// E6D73982-2996-4559-B16C-3E57D46E58D6
var IID_IGameControllerProvider = syscall.GUID{0xE6D73982, 0x2996, 0x4559,
	[8]byte{0xB1, 0x6C, 0x3E, 0x57, 0xD4, 0x6E, 0x58, 0xD6}}

type IGameControllerProviderInterface interface {
	win32.IInspectableInterface
	Get_FirmwareVersionInfo() GameControllerVersionInfo
	Get_HardwareProductId() uint16
	Get_HardwareVendorId() uint16
	Get_HardwareVersionInfo() GameControllerVersionInfo
	Get_IsConnected() bool
}

type IGameControllerProviderVtbl struct {
	win32.IInspectableVtbl
	Get_FirmwareVersionInfo uintptr
	Get_HardwareProductId   uintptr
	Get_HardwareVendorId    uintptr
	Get_HardwareVersionInfo uintptr
	Get_IsConnected         uintptr
}

type IGameControllerProvider struct {
	win32.IInspectable
}

func (this *IGameControllerProvider) Vtbl() *IGameControllerProviderVtbl {
	return (*IGameControllerProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameControllerProvider) Get_FirmwareVersionInfo() GameControllerVersionInfo {
	var _result GameControllerVersionInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FirmwareVersionInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGameControllerProvider) Get_HardwareProductId() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HardwareProductId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGameControllerProvider) Get_HardwareVendorId() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HardwareVendorId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGameControllerProvider) Get_HardwareVersionInfo() GameControllerVersionInfo {
	var _result GameControllerVersionInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HardwareVersionInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGameControllerProvider) Get_IsConnected() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsConnected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6B794D32-8553-4292-8E03-E16651A2F8BC
var IID_IGipFirmwareUpdateResult = syscall.GUID{0x6B794D32, 0x8553, 0x4292,
	[8]byte{0x8E, 0x03, 0xE1, 0x66, 0x51, 0xA2, 0xF8, 0xBC}}

type IGipFirmwareUpdateResultInterface interface {
	win32.IInspectableInterface
	Get_ExtendedErrorCode() uint32
	Get_FinalComponentId() uint32
	Get_Status() GipFirmwareUpdateStatus
}

type IGipFirmwareUpdateResultVtbl struct {
	win32.IInspectableVtbl
	Get_ExtendedErrorCode uintptr
	Get_FinalComponentId  uintptr
	Get_Status            uintptr
}

type IGipFirmwareUpdateResult struct {
	win32.IInspectable
}

func (this *IGipFirmwareUpdateResult) Vtbl() *IGipFirmwareUpdateResultVtbl {
	return (*IGipFirmwareUpdateResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGipFirmwareUpdateResult) Get_ExtendedErrorCode() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedErrorCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGipFirmwareUpdateResult) Get_FinalComponentId() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FinalComponentId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGipFirmwareUpdateResult) Get_Status() GipFirmwareUpdateStatus {
	var _result GipFirmwareUpdateStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A2108ABF-09F1-43BC-A140-80F899EC36FB
var IID_IGipGameControllerInputSink = syscall.GUID{0xA2108ABF, 0x09F1, 0x43BC,
	[8]byte{0xA1, 0x40, 0x80, 0xF8, 0x99, 0xEC, 0x36, 0xFB}}

type IGipGameControllerInputSinkInterface interface {
	win32.IInspectableInterface
	OnKeyReceived(timestamp uint64, keyCode byte, isPressed bool)
	OnMessageReceived(timestamp uint64, messageClass GipMessageClass, messageId byte, sequenceId byte, messageBufferLength uint32, messageBuffer *byte)
}

type IGipGameControllerInputSinkVtbl struct {
	win32.IInspectableVtbl
	OnKeyReceived     uintptr
	OnMessageReceived uintptr
}

type IGipGameControllerInputSink struct {
	win32.IInspectable
}

func (this *IGipGameControllerInputSink) Vtbl() *IGipGameControllerInputSinkVtbl {
	return (*IGipGameControllerInputSinkVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGipGameControllerInputSink) OnKeyReceived(timestamp uint64, keyCode byte, isPressed bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OnKeyReceived, uintptr(unsafe.Pointer(this)), uintptr(timestamp), uintptr(keyCode), uintptr(*(*byte)(unsafe.Pointer(&isPressed))))
	_ = _hr
}

func (this *IGipGameControllerInputSink) OnMessageReceived(timestamp uint64, messageClass GipMessageClass, messageId byte, sequenceId byte, messageBufferLength uint32, messageBuffer *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OnMessageReceived, uintptr(unsafe.Pointer(this)), uintptr(timestamp), uintptr(messageClass), uintptr(messageId), uintptr(sequenceId), uintptr(messageBufferLength), uintptr(unsafe.Pointer(messageBuffer)))
	_ = _hr
}

// DBCF1E19-1AF5-45A8-BF02-A0EE50C823FC
var IID_IGipGameControllerProvider = syscall.GUID{0xDBCF1E19, 0x1AF5, 0x45A8,
	[8]byte{0xBF, 0x02, 0xA0, 0xEE, 0x50, 0xC8, 0x23, 0xFC}}

type IGipGameControllerProviderInterface interface {
	win32.IInspectableInterface
	SendMessage(messageClass GipMessageClass, messageId byte, messageBufferLength uint32, messageBuffer *byte)
	SendReceiveMessage(messageClass GipMessageClass, messageId byte, requestMessageBufferLength uint32, requestMessageBuffer *byte, responseMessageBufferLength uint32, responseMessageBuffer *byte)
	UpdateFirmwareAsync(firmwareImage *IInputStream) *IAsyncOperationWithProgress[*IGipFirmwareUpdateResult, GipFirmwareUpdateProgress]
}

type IGipGameControllerProviderVtbl struct {
	win32.IInspectableVtbl
	SendMessage         uintptr
	SendReceiveMessage  uintptr
	UpdateFirmwareAsync uintptr
}

type IGipGameControllerProvider struct {
	win32.IInspectable
}

func (this *IGipGameControllerProvider) Vtbl() *IGipGameControllerProviderVtbl {
	return (*IGipGameControllerProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGipGameControllerProvider) SendMessage(messageClass GipMessageClass, messageId byte, messageBufferLength uint32, messageBuffer *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendMessage, uintptr(unsafe.Pointer(this)), uintptr(messageClass), uintptr(messageId), uintptr(messageBufferLength), uintptr(unsafe.Pointer(messageBuffer)))
	_ = _hr
}

func (this *IGipGameControllerProvider) SendReceiveMessage(messageClass GipMessageClass, messageId byte, requestMessageBufferLength uint32, requestMessageBuffer *byte, responseMessageBufferLength uint32, responseMessageBuffer *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendReceiveMessage, uintptr(unsafe.Pointer(this)), uintptr(messageClass), uintptr(messageId), uintptr(requestMessageBufferLength), uintptr(unsafe.Pointer(requestMessageBuffer)), uintptr(responseMessageBufferLength), uintptr(unsafe.Pointer(responseMessageBuffer)))
	_ = _hr
}

func (this *IGipGameControllerProvider) UpdateFirmwareAsync(firmwareImage *IInputStream) *IAsyncOperationWithProgress[*IGipFirmwareUpdateResult, GipFirmwareUpdateProgress] {
	var _result *IAsyncOperationWithProgress[*IGipFirmwareUpdateResult, GipFirmwareUpdateProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UpdateFirmwareAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(firmwareImage)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F754C322-182D-40E4-A126-FCEE4FFA1E31
var IID_IHidGameControllerInputSink = syscall.GUID{0xF754C322, 0x182D, 0x40E4,
	[8]byte{0xA1, 0x26, 0xFC, 0xEE, 0x4F, 0xFA, 0x1E, 0x31}}

type IHidGameControllerInputSinkInterface interface {
	win32.IInspectableInterface
	OnInputReportReceived(timestamp uint64, reportId byte, reportBufferLength uint32, reportBuffer *byte)
}

type IHidGameControllerInputSinkVtbl struct {
	win32.IInspectableVtbl
	OnInputReportReceived uintptr
}

type IHidGameControllerInputSink struct {
	win32.IInspectable
}

func (this *IHidGameControllerInputSink) Vtbl() *IHidGameControllerInputSinkVtbl {
	return (*IHidGameControllerInputSinkVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHidGameControllerInputSink) OnInputReportReceived(timestamp uint64, reportId byte, reportBufferLength uint32, reportBuffer *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OnInputReportReceived, uintptr(unsafe.Pointer(this)), uintptr(timestamp), uintptr(reportId), uintptr(reportBufferLength), uintptr(unsafe.Pointer(reportBuffer)))
	_ = _hr
}

// 95CE3AF4-ABF0-4B68-A081-3B7DE73FF0E7
var IID_IHidGameControllerProvider = syscall.GUID{0x95CE3AF4, 0xABF0, 0x4B68,
	[8]byte{0xA0, 0x81, 0x3B, 0x7D, 0xE7, 0x3F, 0xF0, 0xE7}}

type IHidGameControllerProviderInterface interface {
	win32.IInspectableInterface
	Get_UsageId() uint16
	Get_UsagePage() uint16
	GetFeatureReport(reportId byte, reportBufferLength uint32, reportBuffer *byte)
	SendFeatureReport(reportId byte, reportBufferLength uint32, reportBuffer *byte)
	SendOutputReport(reportId byte, reportBufferLength uint32, reportBuffer *byte)
}

type IHidGameControllerProviderVtbl struct {
	win32.IInspectableVtbl
	Get_UsageId       uintptr
	Get_UsagePage     uintptr
	GetFeatureReport  uintptr
	SendFeatureReport uintptr
	SendOutputReport  uintptr
}

type IHidGameControllerProvider struct {
	win32.IInspectable
}

func (this *IHidGameControllerProvider) Vtbl() *IHidGameControllerProviderVtbl {
	return (*IHidGameControllerProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHidGameControllerProvider) Get_UsageId() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UsageId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidGameControllerProvider) Get_UsagePage() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UsagePage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidGameControllerProvider) GetFeatureReport(reportId byte, reportBufferLength uint32, reportBuffer *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFeatureReport, uintptr(unsafe.Pointer(this)), uintptr(reportId), uintptr(reportBufferLength), uintptr(unsafe.Pointer(reportBuffer)))
	_ = _hr
}

func (this *IHidGameControllerProvider) SendFeatureReport(reportId byte, reportBufferLength uint32, reportBuffer *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendFeatureReport, uintptr(unsafe.Pointer(this)), uintptr(reportId), uintptr(reportBufferLength), uintptr(unsafe.Pointer(reportBuffer)))
	_ = _hr
}

func (this *IHidGameControllerProvider) SendOutputReport(reportId byte, reportBufferLength uint32, reportBuffer *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendOutputReport, uintptr(unsafe.Pointer(this)), uintptr(reportId), uintptr(reportBufferLength), uintptr(unsafe.Pointer(reportBuffer)))
	_ = _hr
}

// B2AC1D95-6ECB-42B3-8AAB-025401CA4712
var IID_IXusbGameControllerInputSink = syscall.GUID{0xB2AC1D95, 0x6ECB, 0x42B3,
	[8]byte{0x8A, 0xAB, 0x02, 0x54, 0x01, 0xCA, 0x47, 0x12}}

type IXusbGameControllerInputSinkInterface interface {
	win32.IInspectableInterface
	OnInputReceived(timestamp uint64, reportId byte, inputBufferLength uint32, inputBuffer *byte)
}

type IXusbGameControllerInputSinkVtbl struct {
	win32.IInspectableVtbl
	OnInputReceived uintptr
}

type IXusbGameControllerInputSink struct {
	win32.IInspectable
}

func (this *IXusbGameControllerInputSink) Vtbl() *IXusbGameControllerInputSinkVtbl {
	return (*IXusbGameControllerInputSinkVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXusbGameControllerInputSink) OnInputReceived(timestamp uint64, reportId byte, inputBufferLength uint32, inputBuffer *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OnInputReceived, uintptr(unsafe.Pointer(this)), uintptr(timestamp), uintptr(reportId), uintptr(inputBufferLength), uintptr(unsafe.Pointer(inputBuffer)))
	_ = _hr
}

// 6E2971EB-0EFB-48B4-808B-837643B2F216
var IID_IXusbGameControllerProvider = syscall.GUID{0x6E2971EB, 0x0EFB, 0x48B4,
	[8]byte{0x80, 0x8B, 0x83, 0x76, 0x43, 0xB2, 0xF2, 0x16}}

type IXusbGameControllerProviderInterface interface {
	win32.IInspectableInterface
	SetVibration(lowFrequencyMotorSpeed float64, highFrequencyMotorSpeed float64)
}

type IXusbGameControllerProviderVtbl struct {
	win32.IInspectableVtbl
	SetVibration uintptr
}

type IXusbGameControllerProvider struct {
	win32.IInspectable
}

func (this *IXusbGameControllerProvider) Vtbl() *IXusbGameControllerProviderVtbl {
	return (*IXusbGameControllerProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXusbGameControllerProvider) SetVibration(lowFrequencyMotorSpeed float64, highFrequencyMotorSpeed float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetVibration, uintptr(unsafe.Pointer(this)), uintptr(lowFrequencyMotorSpeed), uintptr(highFrequencyMotorSpeed))
	_ = _hr
}

// classes

type GameControllerFactoryManager struct {
	RtClass
}

func NewIGameControllerFactoryManagerStatics() *IGameControllerFactoryManagerStatics {
	var p *IGameControllerFactoryManagerStatics
	hs := NewHStr("Windows.Gaming.Input.Custom.GameControllerFactoryManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGameControllerFactoryManagerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIGameControllerFactoryManagerStatics2() *IGameControllerFactoryManagerStatics2 {
	var p *IGameControllerFactoryManagerStatics2
	hs := NewHStr("Windows.Gaming.Input.Custom.GameControllerFactoryManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGameControllerFactoryManagerStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type GipFirmwareUpdateResult struct {
	RtClass
	*IGipFirmwareUpdateResult
}

type GipGameControllerProvider struct {
	RtClass
	*IGipGameControllerProvider
}

type HidGameControllerProvider struct {
	RtClass
	*IHidGameControllerProvider
}

type XusbGameControllerProvider struct {
	RtClass
	*IXusbGameControllerProvider
}
