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
type DeviceAccessMode int32

const (
	DeviceAccessMode_Read      DeviceAccessMode = 0
	DeviceAccessMode_Write     DeviceAccessMode = 1
	DeviceAccessMode_ReadWrite DeviceAccessMode = 2
)

// enum
type DeviceSharingMode int32

const (
	DeviceSharingMode_Shared    DeviceSharingMode = 0
	DeviceSharingMode_Exclusive DeviceSharingMode = 1
)

// enum
type IOControlAccessMode int32

const (
	IOControlAccessMode_Any       IOControlAccessMode = 0
	IOControlAccessMode_Read      IOControlAccessMode = 1
	IOControlAccessMode_Write     IOControlAccessMode = 2
	IOControlAccessMode_ReadWrite IOControlAccessMode = 3
)

// enum
type IOControlBufferingMethod int32

const (
	IOControlBufferingMethod_Buffered     IOControlBufferingMethod = 0
	IOControlBufferingMethod_DirectInput  IOControlBufferingMethod = 1
	IOControlBufferingMethod_DirectOutput IOControlBufferingMethod = 2
	IOControlBufferingMethod_Neither      IOControlBufferingMethod = 3
)

// structs

type CustomDeviceContract struct {
}

// interfaces

// DD30251F-C48B-43BD-BCB1-DEC88F15143E
var IID_ICustomDevice = syscall.GUID{0xDD30251F, 0xC48B, 0x43BD,
	[8]byte{0xBC, 0xB1, 0xDE, 0xC8, 0x8F, 0x15, 0x14, 0x3E}}

type ICustomDeviceInterface interface {
	win32.IInspectableInterface
	Get_InputStream() *IInputStream
	Get_OutputStream() *IOutputStream
	SendIOControlAsync(ioControlCode *IIOControlCode, inputBuffer *IBuffer, outputBuffer *IBuffer) *IAsyncOperation[uint32]
	TrySendIOControlAsync(ioControlCode *IIOControlCode, inputBuffer *IBuffer, outputBuffer *IBuffer) *IAsyncOperation[bool]
}

type ICustomDeviceVtbl struct {
	win32.IInspectableVtbl
	Get_InputStream       uintptr
	Get_OutputStream      uintptr
	SendIOControlAsync    uintptr
	TrySendIOControlAsync uintptr
}

type ICustomDevice struct {
	win32.IInspectable
}

func (this *ICustomDevice) Vtbl() *ICustomDeviceVtbl {
	return (*ICustomDeviceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICustomDevice) Get_InputStream() *IInputStream {
	var _result *IInputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InputStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICustomDevice) Get_OutputStream() *IOutputStream {
	var _result *IOutputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutputStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICustomDevice) SendIOControlAsync(ioControlCode *IIOControlCode, inputBuffer *IBuffer, outputBuffer *IBuffer) *IAsyncOperation[uint32] {
	var _result *IAsyncOperation[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendIOControlAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ioControlCode)), uintptr(unsafe.Pointer(inputBuffer)), uintptr(unsafe.Pointer(outputBuffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICustomDevice) TrySendIOControlAsync(ioControlCode *IIOControlCode, inputBuffer *IBuffer, outputBuffer *IBuffer) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySendIOControlAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ioControlCode)), uintptr(unsafe.Pointer(inputBuffer)), uintptr(unsafe.Pointer(outputBuffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C8220312-EF4C-46B1-A58E-EEB308DC8917
var IID_ICustomDeviceStatics = syscall.GUID{0xC8220312, 0xEF4C, 0x46B1,
	[8]byte{0xA5, 0x8E, 0xEE, 0xB3, 0x08, 0xDC, 0x89, 0x17}}

type ICustomDeviceStaticsInterface interface {
	win32.IInspectableInterface
	GetDeviceSelector(classGuid syscall.GUID) string
	FromIdAsync(deviceId string, desiredAccess DeviceAccessMode, sharingMode DeviceSharingMode) *IAsyncOperation[*ICustomDevice]
}

type ICustomDeviceStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector uintptr
	FromIdAsync       uintptr
}

type ICustomDeviceStatics struct {
	win32.IInspectable
}

func (this *ICustomDeviceStatics) Vtbl() *ICustomDeviceStaticsVtbl {
	return (*ICustomDeviceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICustomDeviceStatics) GetDeviceSelector(classGuid syscall.GUID) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&classGuid)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICustomDeviceStatics) FromIdAsync(deviceId string, desiredAccess DeviceAccessMode, sharingMode DeviceSharingMode) *IAsyncOperation[*ICustomDevice] {
	var _result *IAsyncOperation[*ICustomDevice]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(desiredAccess), uintptr(sharingMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0E9559E7-60C8-4375-A761-7F8808066C60
var IID_IIOControlCode = syscall.GUID{0x0E9559E7, 0x60C8, 0x4375,
	[8]byte{0xA7, 0x61, 0x7F, 0x88, 0x08, 0x06, 0x6C, 0x60}}

type IIOControlCodeInterface interface {
	win32.IInspectableInterface
	Get_AccessMode() IOControlAccessMode
	Get_BufferingMethod() IOControlBufferingMethod
	Get_Function() uint16
	Get_DeviceType() uint16
	Get_ControlCode() uint32
}

type IIOControlCodeVtbl struct {
	win32.IInspectableVtbl
	Get_AccessMode      uintptr
	Get_BufferingMethod uintptr
	Get_Function        uintptr
	Get_DeviceType      uintptr
	Get_ControlCode     uintptr
}

type IIOControlCode struct {
	win32.IInspectable
}

func (this *IIOControlCode) Vtbl() *IIOControlCodeVtbl {
	return (*IIOControlCodeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IIOControlCode) Get_AccessMode() IOControlAccessMode {
	var _result IOControlAccessMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AccessMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IIOControlCode) Get_BufferingMethod() IOControlBufferingMethod {
	var _result IOControlBufferingMethod
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BufferingMethod, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IIOControlCode) Get_Function() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Function, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IIOControlCode) Get_DeviceType() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IIOControlCode) Get_ControlCode() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ControlCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 856A7CF0-4C11-44AE-AFC6-B8D4A212788F
var IID_IIOControlCodeFactory = syscall.GUID{0x856A7CF0, 0x4C11, 0x44AE,
	[8]byte{0xAF, 0xC6, 0xB8, 0xD4, 0xA2, 0x12, 0x78, 0x8F}}

type IIOControlCodeFactoryInterface interface {
	win32.IInspectableInterface
	CreateIOControlCode(deviceType uint16, function uint16, accessMode IOControlAccessMode, bufferingMethod IOControlBufferingMethod) *IIOControlCode
}

type IIOControlCodeFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateIOControlCode uintptr
}

type IIOControlCodeFactory struct {
	win32.IInspectable
}

func (this *IIOControlCodeFactory) Vtbl() *IIOControlCodeFactoryVtbl {
	return (*IIOControlCodeFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IIOControlCodeFactory) CreateIOControlCode(deviceType uint16, function uint16, accessMode IOControlAccessMode, bufferingMethod IOControlBufferingMethod) *IIOControlCode {
	var _result *IIOControlCode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateIOControlCode, uintptr(unsafe.Pointer(this)), uintptr(deviceType), uintptr(function), uintptr(accessMode), uintptr(bufferingMethod), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EE5479C2-5448-45DA-AD1B-24948C239094
var IID_IKnownDeviceTypesStatics = syscall.GUID{0xEE5479C2, 0x5448, 0x45DA,
	[8]byte{0xAD, 0x1B, 0x24, 0x94, 0x8C, 0x23, 0x90, 0x94}}

type IKnownDeviceTypesStaticsInterface interface {
	win32.IInspectableInterface
	Get_Unknown() uint16
}

type IKnownDeviceTypesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Unknown uintptr
}

type IKnownDeviceTypesStatics struct {
	win32.IInspectable
}

func (this *IKnownDeviceTypesStatics) Vtbl() *IKnownDeviceTypesStaticsVtbl {
	return (*IKnownDeviceTypesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownDeviceTypesStatics) Get_Unknown() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Unknown, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type CustomDevice struct {
	RtClass
	*ICustomDevice
}

func NewICustomDeviceStatics() *ICustomDeviceStatics {
	var p *ICustomDeviceStatics
	hs := NewHStr("Windows.Devices.Custom.CustomDevice")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICustomDeviceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type IOControlCode struct {
	RtClass
	*IIOControlCode
}

func NewIOControlCode_CreateIOControlCode(deviceType uint16, function uint16, accessMode IOControlAccessMode, bufferingMethod IOControlBufferingMethod) *IOControlCode {
	hs := NewHStr("Windows.Devices.Custom.IOControlCode")
	var pFac *IIOControlCodeFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IIOControlCodeFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IIOControlCode
	p = pFac.CreateIOControlCode(deviceType, function, accessMode, bufferingMethod)
	result := &IOControlCode{
		RtClass:        RtClass{PInspect: &p.IInspectable},
		IIOControlCode: p,
	}
	com.AddToScope(result)
	return result
}

type KnownDeviceTypes struct {
	RtClass
}

func NewIKnownDeviceTypesStatics() *IKnownDeviceTypesStatics {
	var p *IKnownDeviceTypesStatics
	hs := NewHStr("Windows.Devices.Custom.KnownDeviceTypes")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKnownDeviceTypesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
