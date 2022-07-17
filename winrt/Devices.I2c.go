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
type I2cBusSpeed int32

const (
	I2cBusSpeed_StandardMode I2cBusSpeed = 0
	I2cBusSpeed_FastMode     I2cBusSpeed = 1
)

// enum
type I2cSharingMode int32

const (
	I2cSharingMode_Exclusive I2cSharingMode = 0
	I2cSharingMode_Shared    I2cSharingMode = 1
)

// enum
type I2cTransferStatus int32

const (
	I2cTransferStatus_FullTransfer                I2cTransferStatus = 0
	I2cTransferStatus_PartialTransfer             I2cTransferStatus = 1
	I2cTransferStatus_SlaveAddressNotAcknowledged I2cTransferStatus = 2
	I2cTransferStatus_ClockStretchTimeout         I2cTransferStatus = 3
	I2cTransferStatus_UnknownError                I2cTransferStatus = 4
)

// structs

type I2cTransferResult struct {
	Status           I2cTransferStatus
	BytesTransferred uint32
}

// interfaces

// F2DB1307-AB6F-4639-A767-54536DC3460F
var IID_II2cConnectionSettings = syscall.GUID{0xF2DB1307, 0xAB6F, 0x4639,
	[8]byte{0xA7, 0x67, 0x54, 0x53, 0x6D, 0xC3, 0x46, 0x0F}}

type II2cConnectionSettingsInterface interface {
	win32.IInspectableInterface
	Get_SlaveAddress() int32
	Put_SlaveAddress(value int32)
	Get_BusSpeed() I2cBusSpeed
	Put_BusSpeed(value I2cBusSpeed)
	Get_SharingMode() I2cSharingMode
	Put_SharingMode(value I2cSharingMode)
}

type II2cConnectionSettingsVtbl struct {
	win32.IInspectableVtbl
	Get_SlaveAddress uintptr
	Put_SlaveAddress uintptr
	Get_BusSpeed     uintptr
	Put_BusSpeed     uintptr
	Get_SharingMode  uintptr
	Put_SharingMode  uintptr
}

type II2cConnectionSettings struct {
	win32.IInspectable
}

func (this *II2cConnectionSettings) Vtbl() *II2cConnectionSettingsVtbl {
	return (*II2cConnectionSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *II2cConnectionSettings) Get_SlaveAddress() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SlaveAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *II2cConnectionSettings) Put_SlaveAddress(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SlaveAddress, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *II2cConnectionSettings) Get_BusSpeed() I2cBusSpeed {
	var _result I2cBusSpeed
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BusSpeed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *II2cConnectionSettings) Put_BusSpeed(value I2cBusSpeed) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BusSpeed, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *II2cConnectionSettings) Get_SharingMode() I2cSharingMode {
	var _result I2cSharingMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SharingMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *II2cConnectionSettings) Put_SharingMode(value I2cSharingMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SharingMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 81B586B3-9693-41B1-A243-DED4F6E66926
var IID_II2cConnectionSettingsFactory = syscall.GUID{0x81B586B3, 0x9693, 0x41B1,
	[8]byte{0xA2, 0x43, 0xDE, 0xD4, 0xF6, 0xE6, 0x69, 0x26}}

type II2cConnectionSettingsFactoryInterface interface {
	win32.IInspectableInterface
	Create(slaveAddress int32) *II2cConnectionSettings
}

type II2cConnectionSettingsFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type II2cConnectionSettingsFactory struct {
	win32.IInspectable
}

func (this *II2cConnectionSettingsFactory) Vtbl() *II2cConnectionSettingsFactoryVtbl {
	return (*II2cConnectionSettingsFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *II2cConnectionSettingsFactory) Create(slaveAddress int32) *II2cConnectionSettings {
	var _result *II2cConnectionSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(slaveAddress), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C48AB1B2-87A0-4166-8E3E-B4B8F97CD729
var IID_II2cController = syscall.GUID{0xC48AB1B2, 0x87A0, 0x4166,
	[8]byte{0x8E, 0x3E, 0xB4, 0xB8, 0xF9, 0x7C, 0xD7, 0x29}}

type II2cControllerInterface interface {
	win32.IInspectableInterface
	GetDevice(settings *II2cConnectionSettings) *II2cDevice
}

type II2cControllerVtbl struct {
	win32.IInspectableVtbl
	GetDevice uintptr
}

type II2cController struct {
	win32.IInspectable
}

func (this *II2cController) Vtbl() *II2cControllerVtbl {
	return (*II2cControllerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *II2cController) GetDevice(settings *II2cConnectionSettings) *II2cDevice {
	var _result *II2cDevice
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDevice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(settings)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 40FC0365-5F05-4E7E-84BD-100DB8E0AEC5
var IID_II2cControllerStatics = syscall.GUID{0x40FC0365, 0x5F05, 0x4E7E,
	[8]byte{0x84, 0xBD, 0x10, 0x0D, 0xB8, 0xE0, 0xAE, 0xC5}}

type II2cControllerStaticsInterface interface {
	win32.IInspectableInterface
	GetControllersAsync(provider *II2cProvider) *IAsyncOperation[*IVectorView[*II2cController]]
	GetDefaultAsync() *IAsyncOperation[*II2cController]
}

type II2cControllerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetControllersAsync uintptr
	GetDefaultAsync     uintptr
}

type II2cControllerStatics struct {
	win32.IInspectable
}

func (this *II2cControllerStatics) Vtbl() *II2cControllerStaticsVtbl {
	return (*II2cControllerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *II2cControllerStatics) GetControllersAsync(provider *II2cProvider) *IAsyncOperation[*IVectorView[*II2cController]] {
	var _result *IAsyncOperation[*IVectorView[*II2cController]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetControllersAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(provider)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *II2cControllerStatics) GetDefaultAsync() *IAsyncOperation[*II2cController] {
	var _result *IAsyncOperation[*II2cController]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8636C136-B9C5-4F70-9449-CC46DC6F57EB
var IID_II2cDevice = syscall.GUID{0x8636C136, 0xB9C5, 0x4F70,
	[8]byte{0x94, 0x49, 0xCC, 0x46, 0xDC, 0x6F, 0x57, 0xEB}}

type II2cDeviceInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_ConnectionSettings() *II2cConnectionSettings
	Write(bufferLength uint32, buffer *byte)
	WritePartial(bufferLength uint32, buffer *byte) I2cTransferResult
	Read(bufferLength uint32, buffer *byte)
	ReadPartial(bufferLength uint32, buffer *byte) I2cTransferResult
	WriteRead(writeBufferLength uint32, writeBuffer *byte, readBufferLength uint32, readBuffer *byte)
	WriteReadPartial(writeBufferLength uint32, writeBuffer *byte, readBufferLength uint32, readBuffer *byte) I2cTransferResult
}

type II2cDeviceVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId           uintptr
	Get_ConnectionSettings uintptr
	Write                  uintptr
	WritePartial           uintptr
	Read                   uintptr
	ReadPartial            uintptr
	WriteRead              uintptr
	WriteReadPartial       uintptr
}

type II2cDevice struct {
	win32.IInspectable
}

func (this *II2cDevice) Vtbl() *II2cDeviceVtbl {
	return (*II2cDeviceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *II2cDevice) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *II2cDevice) Get_ConnectionSettings() *II2cConnectionSettings {
	var _result *II2cConnectionSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConnectionSettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *II2cDevice) Write(bufferLength uint32, buffer *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Write, uintptr(unsafe.Pointer(this)), uintptr(bufferLength), uintptr(unsafe.Pointer(buffer)))
	_ = _hr
}

func (this *II2cDevice) WritePartial(bufferLength uint32, buffer *byte) I2cTransferResult {
	var _result I2cTransferResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WritePartial, uintptr(unsafe.Pointer(this)), uintptr(bufferLength), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *II2cDevice) Read(bufferLength uint32, buffer *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Read, uintptr(unsafe.Pointer(this)), uintptr(bufferLength), uintptr(unsafe.Pointer(buffer)))
	_ = _hr
}

func (this *II2cDevice) ReadPartial(bufferLength uint32, buffer *byte) I2cTransferResult {
	var _result I2cTransferResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadPartial, uintptr(unsafe.Pointer(this)), uintptr(bufferLength), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *II2cDevice) WriteRead(writeBufferLength uint32, writeBuffer *byte, readBufferLength uint32, readBuffer *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteRead, uintptr(unsafe.Pointer(this)), uintptr(writeBufferLength), uintptr(unsafe.Pointer(writeBuffer)), uintptr(readBufferLength), uintptr(unsafe.Pointer(readBuffer)))
	_ = _hr
}

func (this *II2cDevice) WriteReadPartial(writeBufferLength uint32, writeBuffer *byte, readBufferLength uint32, readBuffer *byte) I2cTransferResult {
	var _result I2cTransferResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteReadPartial, uintptr(unsafe.Pointer(this)), uintptr(writeBufferLength), uintptr(unsafe.Pointer(writeBuffer)), uintptr(readBufferLength), uintptr(unsafe.Pointer(readBuffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 91A33BE3-7334-4512-96BC-FBAE9459F5F6
var IID_II2cDeviceStatics = syscall.GUID{0x91A33BE3, 0x7334, 0x4512,
	[8]byte{0x96, 0xBC, 0xFB, 0xAE, 0x94, 0x59, 0xF5, 0xF6}}

type II2cDeviceStaticsInterface interface {
	win32.IInspectableInterface
	GetDeviceSelector() string
	GetDeviceSelectorFromFriendlyName(friendlyName string) string
	FromIdAsync(deviceId string, settings *II2cConnectionSettings) *IAsyncOperation[*II2cDevice]
}

type II2cDeviceStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector                 uintptr
	GetDeviceSelectorFromFriendlyName uintptr
	FromIdAsync                       uintptr
}

type II2cDeviceStatics struct {
	win32.IInspectable
}

func (this *II2cDeviceStatics) Vtbl() *II2cDeviceStaticsVtbl {
	return (*II2cDeviceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *II2cDeviceStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *II2cDeviceStatics) GetDeviceSelectorFromFriendlyName(friendlyName string) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorFromFriendlyName, uintptr(unsafe.Pointer(this)), NewHStr(friendlyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *II2cDeviceStatics) FromIdAsync(deviceId string, settings *II2cConnectionSettings) *IAsyncOperation[*II2cDevice] {
	var _result *IAsyncOperation[*II2cDevice]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(settings)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type I2cConnectionSettings struct {
	RtClass
	*II2cConnectionSettings
}

func NewI2cConnectionSettings_Create(slaveAddress int32) *I2cConnectionSettings {
	hs := NewHStr("Windows.Devices.I2c.I2cConnectionSettings")
	var pFac *II2cConnectionSettingsFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_II2cConnectionSettingsFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *II2cConnectionSettings
	p = pFac.Create(slaveAddress)
	result := &I2cConnectionSettings{
		RtClass:                RtClass{PInspect: &p.IInspectable},
		II2cConnectionSettings: p,
	}
	com.AddToScope(result)
	return result
}

type I2cController struct {
	RtClass
	*II2cController
}

func NewII2cControllerStatics() *II2cControllerStatics {
	var p *II2cControllerStatics
	hs := NewHStr("Windows.Devices.I2c.I2cController")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_II2cControllerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type I2cDevice struct {
	RtClass
	*II2cDevice
}

func NewII2cDeviceStatics() *II2cDeviceStatics {
	var p *II2cDeviceStatics
	hs := NewHStr("Windows.Devices.I2c.I2cDevice")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_II2cDeviceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
