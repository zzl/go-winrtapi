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
type ProviderSpiMode int32

const (
	ProviderSpiMode_Mode0 ProviderSpiMode = 0
	ProviderSpiMode_Mode1 ProviderSpiMode = 1
	ProviderSpiMode_Mode2 ProviderSpiMode = 2
	ProviderSpiMode_Mode3 ProviderSpiMode = 3
)

// enum
type ProviderSpiSharingMode int32

const (
	ProviderSpiSharingMode_Exclusive ProviderSpiSharingMode = 0
	ProviderSpiSharingMode_Shared    ProviderSpiSharingMode = 1
)

// interfaces

// F6034550-A542-4EC0-9601-A4DD68F8697B
var IID_IProviderSpiConnectionSettings = syscall.GUID{0xF6034550, 0xA542, 0x4EC0,
	[8]byte{0x96, 0x01, 0xA4, 0xDD, 0x68, 0xF8, 0x69, 0x7B}}

type IProviderSpiConnectionSettingsInterface interface {
	win32.IInspectableInterface
	Get_ChipSelectLine() int32
	Put_ChipSelectLine(value int32)
	Get_Mode() ProviderSpiMode
	Put_Mode(value ProviderSpiMode)
	Get_DataBitLength() int32
	Put_DataBitLength(value int32)
	Get_ClockFrequency() int32
	Put_ClockFrequency(value int32)
	Get_SharingMode() ProviderSpiSharingMode
	Put_SharingMode(value ProviderSpiSharingMode)
}

type IProviderSpiConnectionSettingsVtbl struct {
	win32.IInspectableVtbl
	Get_ChipSelectLine uintptr
	Put_ChipSelectLine uintptr
	Get_Mode           uintptr
	Put_Mode           uintptr
	Get_DataBitLength  uintptr
	Put_DataBitLength  uintptr
	Get_ClockFrequency uintptr
	Put_ClockFrequency uintptr
	Get_SharingMode    uintptr
	Put_SharingMode    uintptr
}

type IProviderSpiConnectionSettings struct {
	win32.IInspectable
}

func (this *IProviderSpiConnectionSettings) Vtbl() *IProviderSpiConnectionSettingsVtbl {
	return (*IProviderSpiConnectionSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProviderSpiConnectionSettings) Get_ChipSelectLine() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ChipSelectLine, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProviderSpiConnectionSettings) Put_ChipSelectLine(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ChipSelectLine, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IProviderSpiConnectionSettings) Get_Mode() ProviderSpiMode {
	var _result ProviderSpiMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProviderSpiConnectionSettings) Put_Mode(value ProviderSpiMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Mode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IProviderSpiConnectionSettings) Get_DataBitLength() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DataBitLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProviderSpiConnectionSettings) Put_DataBitLength(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DataBitLength, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IProviderSpiConnectionSettings) Get_ClockFrequency() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ClockFrequency, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProviderSpiConnectionSettings) Put_ClockFrequency(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ClockFrequency, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IProviderSpiConnectionSettings) Get_SharingMode() ProviderSpiSharingMode {
	var _result ProviderSpiSharingMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SharingMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProviderSpiConnectionSettings) Put_SharingMode(value ProviderSpiSharingMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SharingMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 66456B5A-0C79-43E3-9F3C-E59780AC18FA
var IID_IProviderSpiConnectionSettingsFactory = syscall.GUID{0x66456B5A, 0x0C79, 0x43E3,
	[8]byte{0x9F, 0x3C, 0xE5, 0x97, 0x80, 0xAC, 0x18, 0xFA}}

type IProviderSpiConnectionSettingsFactoryInterface interface {
	win32.IInspectableInterface
	Create(chipSelectLine int32) *IProviderSpiConnectionSettings
}

type IProviderSpiConnectionSettingsFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IProviderSpiConnectionSettingsFactory struct {
	win32.IInspectable
}

func (this *IProviderSpiConnectionSettingsFactory) Vtbl() *IProviderSpiConnectionSettingsFactoryVtbl {
	return (*IProviderSpiConnectionSettingsFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProviderSpiConnectionSettingsFactory) Create(chipSelectLine int32) *IProviderSpiConnectionSettings {
	var _result *IProviderSpiConnectionSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(chipSelectLine), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C1686504-02CE-4226-A385-4F11FB04B41B
var IID_ISpiControllerProvider = syscall.GUID{0xC1686504, 0x02CE, 0x4226,
	[8]byte{0xA3, 0x85, 0x4F, 0x11, 0xFB, 0x04, 0xB4, 0x1B}}

type ISpiControllerProviderInterface interface {
	win32.IInspectableInterface
	GetDeviceProvider(settings *IProviderSpiConnectionSettings) *ISpiDeviceProvider
}

type ISpiControllerProviderVtbl struct {
	win32.IInspectableVtbl
	GetDeviceProvider uintptr
}

type ISpiControllerProvider struct {
	win32.IInspectable
}

func (this *ISpiControllerProvider) Vtbl() *ISpiControllerProviderVtbl {
	return (*ISpiControllerProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpiControllerProvider) GetDeviceProvider(settings *IProviderSpiConnectionSettings) *ISpiDeviceProvider {
	var _result *ISpiDeviceProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(settings)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0D1C3443-304B-405C-B4F7-F5AB1074461E
var IID_ISpiDeviceProvider = syscall.GUID{0x0D1C3443, 0x304B, 0x405C,
	[8]byte{0xB4, 0xF7, 0xF5, 0xAB, 0x10, 0x74, 0x46, 0x1E}}

type ISpiDeviceProviderInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_ConnectionSettings() *IProviderSpiConnectionSettings
	Write(bufferLength uint32, buffer *byte)
	Read(bufferLength uint32, buffer *byte)
	TransferSequential(writeBufferLength uint32, writeBuffer *byte, readBufferLength uint32, readBuffer *byte)
	TransferFullDuplex(writeBufferLength uint32, writeBuffer *byte, readBufferLength uint32, readBuffer *byte)
}

type ISpiDeviceProviderVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId           uintptr
	Get_ConnectionSettings uintptr
	Write                  uintptr
	Read                   uintptr
	TransferSequential     uintptr
	TransferFullDuplex     uintptr
}

type ISpiDeviceProvider struct {
	win32.IInspectable
}

func (this *ISpiDeviceProvider) Vtbl() *ISpiDeviceProviderVtbl {
	return (*ISpiDeviceProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpiDeviceProvider) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISpiDeviceProvider) Get_ConnectionSettings() *IProviderSpiConnectionSettings {
	var _result *IProviderSpiConnectionSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConnectionSettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpiDeviceProvider) Write(bufferLength uint32, buffer *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Write, uintptr(unsafe.Pointer(this)), uintptr(bufferLength), uintptr(unsafe.Pointer(buffer)))
	_ = _hr
}

func (this *ISpiDeviceProvider) Read(bufferLength uint32, buffer *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Read, uintptr(unsafe.Pointer(this)), uintptr(bufferLength), uintptr(unsafe.Pointer(buffer)))
	_ = _hr
}

func (this *ISpiDeviceProvider) TransferSequential(writeBufferLength uint32, writeBuffer *byte, readBufferLength uint32, readBuffer *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TransferSequential, uintptr(unsafe.Pointer(this)), uintptr(writeBufferLength), uintptr(unsafe.Pointer(writeBuffer)), uintptr(readBufferLength), uintptr(unsafe.Pointer(readBuffer)))
	_ = _hr
}

func (this *ISpiDeviceProvider) TransferFullDuplex(writeBufferLength uint32, writeBuffer *byte, readBufferLength uint32, readBuffer *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TransferFullDuplex, uintptr(unsafe.Pointer(this)), uintptr(writeBufferLength), uintptr(unsafe.Pointer(writeBuffer)), uintptr(readBufferLength), uintptr(unsafe.Pointer(readBuffer)))
	_ = _hr
}

// 96B461E2-77D4-48CE-AAA0-75715A8362CF
var IID_ISpiProvider = syscall.GUID{0x96B461E2, 0x77D4, 0x48CE,
	[8]byte{0xAA, 0xA0, 0x75, 0x71, 0x5A, 0x83, 0x62, 0xCF}}

type ISpiProviderInterface interface {
	win32.IInspectableInterface
	GetControllersAsync() *IAsyncOperation[*IVectorView[*ISpiControllerProvider]]
}

type ISpiProviderVtbl struct {
	win32.IInspectableVtbl
	GetControllersAsync uintptr
}

type ISpiProvider struct {
	win32.IInspectable
}

func (this *ISpiProvider) Vtbl() *ISpiProviderVtbl {
	return (*ISpiProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpiProvider) GetControllersAsync() *IAsyncOperation[*IVectorView[*ISpiControllerProvider]] {
	var _result *IAsyncOperation[*IVectorView[*ISpiControllerProvider]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetControllersAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type ProviderSpiConnectionSettings struct {
	RtClass
	*IProviderSpiConnectionSettings
}

func NewProviderSpiConnectionSettings_Create(chipSelectLine int32) *ProviderSpiConnectionSettings {
	hs := NewHStr("Windows.Devices.Spi.Provider.ProviderSpiConnectionSettings")
	var pFac *IProviderSpiConnectionSettingsFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IProviderSpiConnectionSettingsFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IProviderSpiConnectionSettings
	p = pFac.Create(chipSelectLine)
	result := &ProviderSpiConnectionSettings{
		RtClass:                        RtClass{PInspect: &p.IInspectable},
		IProviderSpiConnectionSettings: p,
	}
	com.AddToScope(result)
	return result
}
