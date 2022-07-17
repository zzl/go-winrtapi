package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"log"
	"syscall"
	"unsafe"
)

// structs

type DevicesLowLevelContract struct {
}

// interfaces

// A73E561C-AAC1-4EC7-A852-479F7060D01F
var IID_ILowLevelDevicesAggregateProvider = syscall.GUID{0xA73E561C, 0xAAC1, 0x4EC7,
	[8]byte{0xA8, 0x52, 0x47, 0x9F, 0x70, 0x60, 0xD0, 0x1F}}

type ILowLevelDevicesAggregateProviderInterface interface {
	win32.IInspectableInterface
	Get_AdcControllerProvider() unsafe.Pointer
	Get_PwmControllerProvider() unsafe.Pointer
	Get_GpioControllerProvider() *IGpioControllerProvider
	Get_I2cControllerProvider() *II2cControllerProvider
	Get_SpiControllerProvider() *ISpiControllerProvider
}

type ILowLevelDevicesAggregateProviderVtbl struct {
	win32.IInspectableVtbl
	Get_AdcControllerProvider  uintptr
	Get_PwmControllerProvider  uintptr
	Get_GpioControllerProvider uintptr
	Get_I2cControllerProvider  uintptr
	Get_SpiControllerProvider  uintptr
}

type ILowLevelDevicesAggregateProvider struct {
	win32.IInspectable
}

func (this *ILowLevelDevicesAggregateProvider) Vtbl() *ILowLevelDevicesAggregateProviderVtbl {
	return (*ILowLevelDevicesAggregateProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILowLevelDevicesAggregateProvider) Get_AdcControllerProvider() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AdcControllerProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILowLevelDevicesAggregateProvider) Get_PwmControllerProvider() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PwmControllerProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILowLevelDevicesAggregateProvider) Get_GpioControllerProvider() *IGpioControllerProvider {
	var _result *IGpioControllerProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GpioControllerProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILowLevelDevicesAggregateProvider) Get_I2cControllerProvider() *II2cControllerProvider {
	var _result *II2cControllerProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_I2cControllerProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILowLevelDevicesAggregateProvider) Get_SpiControllerProvider() *ISpiControllerProvider {
	var _result *ISpiControllerProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SpiControllerProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9AC4AAF6-3473-465E-96D5-36281A2C57AF
var IID_ILowLevelDevicesAggregateProviderFactory = syscall.GUID{0x9AC4AAF6, 0x3473, 0x465E,
	[8]byte{0x96, 0xD5, 0x36, 0x28, 0x1A, 0x2C, 0x57, 0xAF}}

type ILowLevelDevicesAggregateProviderFactoryInterface interface {
	win32.IInspectableInterface
	Create(adc unsafe.Pointer, pwm unsafe.Pointer, gpio *IGpioControllerProvider, i2c *II2cControllerProvider, spi *ISpiControllerProvider) *ILowLevelDevicesAggregateProvider
}

type ILowLevelDevicesAggregateProviderFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ILowLevelDevicesAggregateProviderFactory struct {
	win32.IInspectable
}

func (this *ILowLevelDevicesAggregateProviderFactory) Vtbl() *ILowLevelDevicesAggregateProviderFactoryVtbl {
	return (*ILowLevelDevicesAggregateProviderFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILowLevelDevicesAggregateProviderFactory) Create(adc unsafe.Pointer, pwm unsafe.Pointer, gpio *IGpioControllerProvider, i2c *II2cControllerProvider, spi *ISpiControllerProvider) *ILowLevelDevicesAggregateProvider {
	var _result *ILowLevelDevicesAggregateProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(adc), uintptr(pwm), uintptr(unsafe.Pointer(gpio)), uintptr(unsafe.Pointer(i2c)), uintptr(unsafe.Pointer(spi)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2EC23DD4-179B-45DE-9B39-3AE02527DE52
var IID_ILowLevelDevicesController = syscall.GUID{0x2EC23DD4, 0x179B, 0x45DE,
	[8]byte{0x9B, 0x39, 0x3A, 0xE0, 0x25, 0x27, 0xDE, 0x52}}

type ILowLevelDevicesControllerInterface interface {
	win32.IInspectableInterface
}

type ILowLevelDevicesControllerVtbl struct {
	win32.IInspectableVtbl
}

type ILowLevelDevicesController struct {
	win32.IInspectable
}

func (this *ILowLevelDevicesController) Vtbl() *ILowLevelDevicesControllerVtbl {
	return (*ILowLevelDevicesControllerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 093E926A-FCCB-4394-A697-19DE637C2DB3
var IID_ILowLevelDevicesControllerStatics = syscall.GUID{0x093E926A, 0xFCCB, 0x4394,
	[8]byte{0xA6, 0x97, 0x19, 0xDE, 0x63, 0x7C, 0x2D, 0xB3}}

type ILowLevelDevicesControllerStaticsInterface interface {
	win32.IInspectableInterface
	Get_DefaultProvider() *ILowLevelDevicesAggregateProvider
	Put_DefaultProvider(value *ILowLevelDevicesAggregateProvider)
}

type ILowLevelDevicesControllerStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_DefaultProvider uintptr
	Put_DefaultProvider uintptr
}

type ILowLevelDevicesControllerStatics struct {
	win32.IInspectable
}

func (this *ILowLevelDevicesControllerStatics) Vtbl() *ILowLevelDevicesControllerStaticsVtbl {
	return (*ILowLevelDevicesControllerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILowLevelDevicesControllerStatics) Get_DefaultProvider() *ILowLevelDevicesAggregateProvider {
	var _result *ILowLevelDevicesAggregateProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DefaultProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILowLevelDevicesControllerStatics) Put_DefaultProvider(value *ILowLevelDevicesAggregateProvider) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DefaultProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// classes

type LowLevelDevicesAggregateProvider struct {
	RtClass
	*ILowLevelDevicesAggregateProvider
}

func NewLowLevelDevicesAggregateProvider_Create(adc unsafe.Pointer, pwm unsafe.Pointer, gpio *IGpioControllerProvider, i2c *II2cControllerProvider, spi *ISpiControllerProvider) *LowLevelDevicesAggregateProvider {
	hs := NewHStr("Windows.Devices.LowLevelDevicesAggregateProvider")
	var pFac *ILowLevelDevicesAggregateProviderFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ILowLevelDevicesAggregateProviderFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ILowLevelDevicesAggregateProvider
	p = pFac.Create(adc, pwm, gpio, i2c, spi)
	result := &LowLevelDevicesAggregateProvider{
		RtClass:                           RtClass{PInspect: &p.IInspectable},
		ILowLevelDevicesAggregateProvider: p,
	}
	com.AddToScope(result)
	return result
}

type LowLevelDevicesController struct {
	RtClass
	*ILowLevelDevicesController
}

func NewILowLevelDevicesControllerStatics() *ILowLevelDevicesControllerStatics {
	var p *ILowLevelDevicesControllerStatics
	hs := NewHStr("Windows.Devices.LowLevelDevicesController")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ILowLevelDevicesControllerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
