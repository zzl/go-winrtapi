package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type ProviderI2cBusSpeed int32

const (
	ProviderI2cBusSpeed_StandardMode ProviderI2cBusSpeed = 0
	ProviderI2cBusSpeed_FastMode     ProviderI2cBusSpeed = 1
)

// enum
type ProviderI2cSharingMode int32

const (
	ProviderI2cSharingMode_Exclusive ProviderI2cSharingMode = 0
	ProviderI2cSharingMode_Shared    ProviderI2cSharingMode = 1
)

// enum
type ProviderI2cTransferStatus int32

const (
	ProviderI2cTransferStatus_FullTransfer                ProviderI2cTransferStatus = 0
	ProviderI2cTransferStatus_PartialTransfer             ProviderI2cTransferStatus = 1
	ProviderI2cTransferStatus_SlaveAddressNotAcknowledged ProviderI2cTransferStatus = 2
)

// structs

type ProviderI2cTransferResult struct {
	Status           ProviderI2cTransferStatus
	BytesTransferred uint32
}

// interfaces

// 61C2BB82-4510-4163-A87C-4E15A9558980
var IID_II2cControllerProvider = syscall.GUID{0x61C2BB82, 0x4510, 0x4163,
	[8]byte{0xA8, 0x7C, 0x4E, 0x15, 0xA9, 0x55, 0x89, 0x80}}

type II2cControllerProviderInterface interface {
	win32.IInspectableInterface
	GetDeviceProvider(settings *IProviderI2cConnectionSettings) *II2cDeviceProvider
}

type II2cControllerProviderVtbl struct {
	win32.IInspectableVtbl
	GetDeviceProvider uintptr
}

type II2cControllerProvider struct {
	win32.IInspectable
}

func (this *II2cControllerProvider) Vtbl() *II2cControllerProviderVtbl {
	return (*II2cControllerProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *II2cControllerProvider) GetDeviceProvider(settings *IProviderI2cConnectionSettings) *II2cDeviceProvider {
	var _result *II2cDeviceProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(settings)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AD342654-57E8-453E-8329-D1E447D103A9
var IID_II2cDeviceProvider = syscall.GUID{0xAD342654, 0x57E8, 0x453E,
	[8]byte{0x83, 0x29, 0xD1, 0xE4, 0x47, 0xD1, 0x03, 0xA9}}

type II2cDeviceProviderInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Write(bufferLength uint32, buffer *byte)
	WritePartial(bufferLength uint32, buffer *byte) ProviderI2cTransferResult
	Read(bufferLength uint32, buffer *byte)
	ReadPartial(bufferLength uint32, buffer *byte) ProviderI2cTransferResult
	WriteRead(writeBufferLength uint32, writeBuffer *byte, readBufferLength uint32, readBuffer *byte)
	WriteReadPartial(writeBufferLength uint32, writeBuffer *byte, readBufferLength uint32, readBuffer *byte) ProviderI2cTransferResult
}

type II2cDeviceProviderVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId     uintptr
	Write            uintptr
	WritePartial     uintptr
	Read             uintptr
	ReadPartial      uintptr
	WriteRead        uintptr
	WriteReadPartial uintptr
}

type II2cDeviceProvider struct {
	win32.IInspectable
}

func (this *II2cDeviceProvider) Vtbl() *II2cDeviceProviderVtbl {
	return (*II2cDeviceProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *II2cDeviceProvider) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *II2cDeviceProvider) Write(bufferLength uint32, buffer *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Write, uintptr(unsafe.Pointer(this)), uintptr(bufferLength), uintptr(unsafe.Pointer(buffer)))
	_ = _hr
}

func (this *II2cDeviceProvider) WritePartial(bufferLength uint32, buffer *byte) ProviderI2cTransferResult {
	var _result ProviderI2cTransferResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WritePartial, uintptr(unsafe.Pointer(this)), uintptr(bufferLength), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *II2cDeviceProvider) Read(bufferLength uint32, buffer *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Read, uintptr(unsafe.Pointer(this)), uintptr(bufferLength), uintptr(unsafe.Pointer(buffer)))
	_ = _hr
}

func (this *II2cDeviceProvider) ReadPartial(bufferLength uint32, buffer *byte) ProviderI2cTransferResult {
	var _result ProviderI2cTransferResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadPartial, uintptr(unsafe.Pointer(this)), uintptr(bufferLength), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *II2cDeviceProvider) WriteRead(writeBufferLength uint32, writeBuffer *byte, readBufferLength uint32, readBuffer *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteRead, uintptr(unsafe.Pointer(this)), uintptr(writeBufferLength), uintptr(unsafe.Pointer(writeBuffer)), uintptr(readBufferLength), uintptr(unsafe.Pointer(readBuffer)))
	_ = _hr
}

func (this *II2cDeviceProvider) WriteReadPartial(writeBufferLength uint32, writeBuffer *byte, readBufferLength uint32, readBuffer *byte) ProviderI2cTransferResult {
	var _result ProviderI2cTransferResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteReadPartial, uintptr(unsafe.Pointer(this)), uintptr(writeBufferLength), uintptr(unsafe.Pointer(writeBuffer)), uintptr(readBufferLength), uintptr(unsafe.Pointer(readBuffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6F13083E-BF62-4FE2-A95A-F08999669818
var IID_II2cProvider = syscall.GUID{0x6F13083E, 0xBF62, 0x4FE2,
	[8]byte{0xA9, 0x5A, 0xF0, 0x89, 0x99, 0x66, 0x98, 0x18}}

type II2cProviderInterface interface {
	win32.IInspectableInterface
	GetControllersAsync() *IAsyncOperation[*IVectorView[*II2cControllerProvider]]
}

type II2cProviderVtbl struct {
	win32.IInspectableVtbl
	GetControllersAsync uintptr
}

type II2cProvider struct {
	win32.IInspectable
}

func (this *II2cProvider) Vtbl() *II2cProviderVtbl {
	return (*II2cProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *II2cProvider) GetControllersAsync() *IAsyncOperation[*IVectorView[*II2cControllerProvider]] {
	var _result *IAsyncOperation[*IVectorView[*II2cControllerProvider]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetControllersAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E9DB4E34-E510-44B7-809D-F2F85B555339
var IID_IProviderI2cConnectionSettings = syscall.GUID{0xE9DB4E34, 0xE510, 0x44B7,
	[8]byte{0x80, 0x9D, 0xF2, 0xF8, 0x5B, 0x55, 0x53, 0x39}}

type IProviderI2cConnectionSettingsInterface interface {
	win32.IInspectableInterface
	Get_SlaveAddress() int32
	Put_SlaveAddress(value int32)
	Get_BusSpeed() ProviderI2cBusSpeed
	Put_BusSpeed(value ProviderI2cBusSpeed)
	Get_SharingMode() ProviderI2cSharingMode
	Put_SharingMode(value ProviderI2cSharingMode)
}

type IProviderI2cConnectionSettingsVtbl struct {
	win32.IInspectableVtbl
	Get_SlaveAddress uintptr
	Put_SlaveAddress uintptr
	Get_BusSpeed     uintptr
	Put_BusSpeed     uintptr
	Get_SharingMode  uintptr
	Put_SharingMode  uintptr
}

type IProviderI2cConnectionSettings struct {
	win32.IInspectable
}

func (this *IProviderI2cConnectionSettings) Vtbl() *IProviderI2cConnectionSettingsVtbl {
	return (*IProviderI2cConnectionSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProviderI2cConnectionSettings) Get_SlaveAddress() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SlaveAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProviderI2cConnectionSettings) Put_SlaveAddress(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SlaveAddress, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IProviderI2cConnectionSettings) Get_BusSpeed() ProviderI2cBusSpeed {
	var _result ProviderI2cBusSpeed
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BusSpeed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProviderI2cConnectionSettings) Put_BusSpeed(value ProviderI2cBusSpeed) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BusSpeed, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IProviderI2cConnectionSettings) Get_SharingMode() ProviderI2cSharingMode {
	var _result ProviderI2cSharingMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SharingMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProviderI2cConnectionSettings) Put_SharingMode(value ProviderI2cSharingMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SharingMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// classes

type ProviderI2cConnectionSettings struct {
	RtClass
	*IProviderI2cConnectionSettings
}
