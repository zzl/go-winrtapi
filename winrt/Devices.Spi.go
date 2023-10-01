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
type SpiMode int32

const (
	SpiMode_Mode0 SpiMode = 0
	SpiMode_Mode1 SpiMode = 1
	SpiMode_Mode2 SpiMode = 2
	SpiMode_Mode3 SpiMode = 3
)

// enum
type SpiSharingMode int32

const (
	SpiSharingMode_Exclusive SpiSharingMode = 0
	SpiSharingMode_Shared    SpiSharingMode = 1
)

// interfaces

// 9929444A-54F2-48C6-B952-9C32FC02C669
var IID_ISpiBusInfo = syscall.GUID{0x9929444A, 0x54F2, 0x48C6,
	[8]byte{0xB9, 0x52, 0x9C, 0x32, 0xFC, 0x02, 0xC6, 0x69}}

type ISpiBusInfoInterface interface {
	win32.IInspectableInterface
	Get_ChipSelectLineCount() int32
	Get_MinClockFrequency() int32
	Get_MaxClockFrequency() int32
	Get_SupportedDataBitLengths() *IVectorView[int32]
}

type ISpiBusInfoVtbl struct {
	win32.IInspectableVtbl
	Get_ChipSelectLineCount     uintptr
	Get_MinClockFrequency       uintptr
	Get_MaxClockFrequency       uintptr
	Get_SupportedDataBitLengths uintptr
}

type ISpiBusInfo struct {
	win32.IInspectable
}

func (this *ISpiBusInfo) Vtbl() *ISpiBusInfoVtbl {
	return (*ISpiBusInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpiBusInfo) Get_ChipSelectLineCount() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ChipSelectLineCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpiBusInfo) Get_MinClockFrequency() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinClockFrequency, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpiBusInfo) Get_MaxClockFrequency() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxClockFrequency, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpiBusInfo) Get_SupportedDataBitLengths() *IVectorView[int32] {
	var _result *IVectorView[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedDataBitLengths, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5283A37F-F935-4B9F-A7A7-3A7890AFA5CE
var IID_ISpiConnectionSettings = syscall.GUID{0x5283A37F, 0xF935, 0x4B9F,
	[8]byte{0xA7, 0xA7, 0x3A, 0x78, 0x90, 0xAF, 0xA5, 0xCE}}

type ISpiConnectionSettingsInterface interface {
	win32.IInspectableInterface
	Get_ChipSelectLine() int32
	Put_ChipSelectLine(value int32)
	Get_Mode() SpiMode
	Put_Mode(value SpiMode)
	Get_DataBitLength() int32
	Put_DataBitLength(value int32)
	Get_ClockFrequency() int32
	Put_ClockFrequency(value int32)
	Get_SharingMode() SpiSharingMode
	Put_SharingMode(value SpiSharingMode)
}

type ISpiConnectionSettingsVtbl struct {
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

type ISpiConnectionSettings struct {
	win32.IInspectable
}

func (this *ISpiConnectionSettings) Vtbl() *ISpiConnectionSettingsVtbl {
	return (*ISpiConnectionSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpiConnectionSettings) Get_ChipSelectLine() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ChipSelectLine, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpiConnectionSettings) Put_ChipSelectLine(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ChipSelectLine, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpiConnectionSettings) Get_Mode() SpiMode {
	var _result SpiMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpiConnectionSettings) Put_Mode(value SpiMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Mode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpiConnectionSettings) Get_DataBitLength() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DataBitLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpiConnectionSettings) Put_DataBitLength(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DataBitLength, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpiConnectionSettings) Get_ClockFrequency() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ClockFrequency, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpiConnectionSettings) Put_ClockFrequency(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ClockFrequency, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpiConnectionSettings) Get_SharingMode() SpiSharingMode {
	var _result SpiSharingMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SharingMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpiConnectionSettings) Put_SharingMode(value SpiSharingMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SharingMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// FF99081E-10C4-44B7-9FEA-A748B5A46F31
var IID_ISpiConnectionSettingsFactory = syscall.GUID{0xFF99081E, 0x10C4, 0x44B7,
	[8]byte{0x9F, 0xEA, 0xA7, 0x48, 0xB5, 0xA4, 0x6F, 0x31}}

type ISpiConnectionSettingsFactoryInterface interface {
	win32.IInspectableInterface
	Create(chipSelectLine int32) *ISpiConnectionSettings
}

type ISpiConnectionSettingsFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ISpiConnectionSettingsFactory struct {
	win32.IInspectable
}

func (this *ISpiConnectionSettingsFactory) Vtbl() *ISpiConnectionSettingsFactoryVtbl {
	return (*ISpiConnectionSettingsFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpiConnectionSettingsFactory) Create(chipSelectLine int32) *ISpiConnectionSettings {
	var _result *ISpiConnectionSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(chipSelectLine), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A8D3C829-9895-4159-A934-8741F1EE6D27
var IID_ISpiController = syscall.GUID{0xA8D3C829, 0x9895, 0x4159,
	[8]byte{0xA9, 0x34, 0x87, 0x41, 0xF1, 0xEE, 0x6D, 0x27}}

type ISpiControllerInterface interface {
	win32.IInspectableInterface
	GetDevice(settings *ISpiConnectionSettings) *ISpiDevice
}

type ISpiControllerVtbl struct {
	win32.IInspectableVtbl
	GetDevice uintptr
}

type ISpiController struct {
	win32.IInspectable
}

func (this *ISpiController) Vtbl() *ISpiControllerVtbl {
	return (*ISpiControllerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpiController) GetDevice(settings *ISpiConnectionSettings) *ISpiDevice {
	var _result *ISpiDevice
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDevice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(settings)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0D5229E2-138B-4E48-B964-4F2F79B9C5A2
var IID_ISpiControllerStatics = syscall.GUID{0x0D5229E2, 0x138B, 0x4E48,
	[8]byte{0xB9, 0x64, 0x4F, 0x2F, 0x79, 0xB9, 0xC5, 0xA2}}

type ISpiControllerStaticsInterface interface {
	win32.IInspectableInterface
	GetDefaultAsync() *IAsyncOperation[*ISpiController]
	GetControllersAsync(provider *ISpiProvider) *IAsyncOperation[*IVectorView[*ISpiController]]
}

type ISpiControllerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefaultAsync     uintptr
	GetControllersAsync uintptr
}

type ISpiControllerStatics struct {
	win32.IInspectable
}

func (this *ISpiControllerStatics) Vtbl() *ISpiControllerStaticsVtbl {
	return (*ISpiControllerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpiControllerStatics) GetDefaultAsync() *IAsyncOperation[*ISpiController] {
	var _result *IAsyncOperation[*ISpiController]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpiControllerStatics) GetControllersAsync(provider *ISpiProvider) *IAsyncOperation[*IVectorView[*ISpiController]] {
	var _result *IAsyncOperation[*IVectorView[*ISpiController]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetControllersAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(provider)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 05D5356D-11B6-4D39-84D5-95DFB4C9F2CE
var IID_ISpiDevice = syscall.GUID{0x05D5356D, 0x11B6, 0x4D39,
	[8]byte{0x84, 0xD5, 0x95, 0xDF, 0xB4, 0xC9, 0xF2, 0xCE}}

type ISpiDeviceInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_ConnectionSettings() *ISpiConnectionSettings
	Write(bufferLength uint32, buffer *byte)
	Read(bufferLength uint32, buffer *byte)
	TransferSequential(writeBufferLength uint32, writeBuffer *byte, readBufferLength uint32, readBuffer *byte)
	TransferFullDuplex(writeBufferLength uint32, writeBuffer *byte, readBufferLength uint32, readBuffer *byte)
}

type ISpiDeviceVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId           uintptr
	Get_ConnectionSettings uintptr
	Write                  uintptr
	Read                   uintptr
	TransferSequential     uintptr
	TransferFullDuplex     uintptr
}

type ISpiDevice struct {
	win32.IInspectable
}

func (this *ISpiDevice) Vtbl() *ISpiDeviceVtbl {
	return (*ISpiDeviceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpiDevice) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISpiDevice) Get_ConnectionSettings() *ISpiConnectionSettings {
	var _result *ISpiConnectionSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConnectionSettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpiDevice) Write(bufferLength uint32, buffer *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Write, uintptr(unsafe.Pointer(this)), uintptr(bufferLength), uintptr(unsafe.Pointer(buffer)))
	_ = _hr
}

func (this *ISpiDevice) Read(bufferLength uint32, buffer *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Read, uintptr(unsafe.Pointer(this)), uintptr(bufferLength), uintptr(unsafe.Pointer(buffer)))
	_ = _hr
}

func (this *ISpiDevice) TransferSequential(writeBufferLength uint32, writeBuffer *byte, readBufferLength uint32, readBuffer *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TransferSequential, uintptr(unsafe.Pointer(this)), uintptr(writeBufferLength), uintptr(unsafe.Pointer(writeBuffer)), uintptr(readBufferLength), uintptr(unsafe.Pointer(readBuffer)))
	_ = _hr
}

func (this *ISpiDevice) TransferFullDuplex(writeBufferLength uint32, writeBuffer *byte, readBufferLength uint32, readBuffer *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TransferFullDuplex, uintptr(unsafe.Pointer(this)), uintptr(writeBufferLength), uintptr(unsafe.Pointer(writeBuffer)), uintptr(readBufferLength), uintptr(unsafe.Pointer(readBuffer)))
	_ = _hr
}

// A278E559-5720-4D3F-BD93-56F5FF5A5879
var IID_ISpiDeviceStatics = syscall.GUID{0xA278E559, 0x5720, 0x4D3F,
	[8]byte{0xBD, 0x93, 0x56, 0xF5, 0xFF, 0x5A, 0x58, 0x79}}

type ISpiDeviceStaticsInterface interface {
	win32.IInspectableInterface
	GetDeviceSelector() string
	GetDeviceSelectorFromFriendlyName(friendlyName string) string
	GetBusInfo(busId string) *ISpiBusInfo
	FromIdAsync(busId string, settings *ISpiConnectionSettings) *IAsyncOperation[*ISpiDevice]
}

type ISpiDeviceStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector                 uintptr
	GetDeviceSelectorFromFriendlyName uintptr
	GetBusInfo                        uintptr
	FromIdAsync                       uintptr
}

type ISpiDeviceStatics struct {
	win32.IInspectable
}

func (this *ISpiDeviceStatics) Vtbl() *ISpiDeviceStaticsVtbl {
	return (*ISpiDeviceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpiDeviceStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISpiDeviceStatics) GetDeviceSelectorFromFriendlyName(friendlyName string) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorFromFriendlyName, uintptr(unsafe.Pointer(this)), NewHStr(friendlyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISpiDeviceStatics) GetBusInfo(busId string) *ISpiBusInfo {
	var _result *ISpiBusInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetBusInfo, uintptr(unsafe.Pointer(this)), NewHStr(busId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpiDeviceStatics) FromIdAsync(busId string, settings *ISpiConnectionSettings) *IAsyncOperation[*ISpiDevice] {
	var _result *IAsyncOperation[*ISpiDevice]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(busId).Ptr, uintptr(unsafe.Pointer(settings)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type SpiBusInfo struct {
	RtClass
	*ISpiBusInfo
}

type SpiConnectionSettings struct {
	RtClass
	*ISpiConnectionSettings
}

func NewSpiConnectionSettings_Create(chipSelectLine int32) *SpiConnectionSettings {
	hs := NewHStr("Windows.Devices.Spi.SpiConnectionSettings")
	var pFac *ISpiConnectionSettingsFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISpiConnectionSettingsFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISpiConnectionSettings
	p = pFac.Create(chipSelectLine)
	result := &SpiConnectionSettings{
		RtClass:                RtClass{PInspect: &p.IInspectable},
		ISpiConnectionSettings: p,
	}
	com.AddToScope(result)
	return result
}

type SpiController struct {
	RtClass
	*ISpiController
}

func NewISpiControllerStatics() *ISpiControllerStatics {
	var p *ISpiControllerStatics
	hs := NewHStr("Windows.Devices.Spi.SpiController")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISpiControllerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SpiDevice struct {
	RtClass
	*ISpiDevice
}

func NewISpiDeviceStatics() *ISpiDeviceStatics {
	var p *ISpiDeviceStatics
	hs := NewHStr("Windows.Devices.Spi.SpiDevice")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISpiDeviceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
