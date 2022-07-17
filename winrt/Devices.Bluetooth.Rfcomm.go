package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// interfaces

// AE81FF1F-C5A1-4C40-8C28-F3EFD69062F3
var IID_IRfcommDeviceService = syscall.GUID{0xAE81FF1F, 0xC5A1, 0x4C40,
	[8]byte{0x8C, 0x28, 0xF3, 0xEF, 0xD6, 0x90, 0x62, 0xF3}}

type IRfcommDeviceServiceInterface interface {
	win32.IInspectableInterface
	Get_ConnectionHostName() *IHostName
	Get_ConnectionServiceName() string
	Get_ServiceId() *IRfcommServiceId
	Get_ProtectionLevel() SocketProtectionLevel
	Get_MaxProtectionLevel() SocketProtectionLevel
	GetSdpRawAttributesAsync() *IAsyncOperation[*IMapView[uint32, *IBuffer]]
	GetSdpRawAttributesWithCacheModeAsync(cacheMode BluetoothCacheMode) *IAsyncOperation[*IMapView[uint32, *IBuffer]]
}

type IRfcommDeviceServiceVtbl struct {
	win32.IInspectableVtbl
	Get_ConnectionHostName                uintptr
	Get_ConnectionServiceName             uintptr
	Get_ServiceId                         uintptr
	Get_ProtectionLevel                   uintptr
	Get_MaxProtectionLevel                uintptr
	GetSdpRawAttributesAsync              uintptr
	GetSdpRawAttributesWithCacheModeAsync uintptr
}

type IRfcommDeviceService struct {
	win32.IInspectable
}

func (this *IRfcommDeviceService) Vtbl() *IRfcommDeviceServiceVtbl {
	return (*IRfcommDeviceServiceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRfcommDeviceService) Get_ConnectionHostName() *IHostName {
	var _result *IHostName
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConnectionHostName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRfcommDeviceService) Get_ConnectionServiceName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConnectionServiceName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IRfcommDeviceService) Get_ServiceId() *IRfcommServiceId {
	var _result *IRfcommServiceId
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRfcommDeviceService) Get_ProtectionLevel() SocketProtectionLevel {
	var _result SocketProtectionLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRfcommDeviceService) Get_MaxProtectionLevel() SocketProtectionLevel {
	var _result SocketProtectionLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRfcommDeviceService) GetSdpRawAttributesAsync() *IAsyncOperation[*IMapView[uint32, *IBuffer]] {
	var _result *IAsyncOperation[*IMapView[uint32, *IBuffer]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSdpRawAttributesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRfcommDeviceService) GetSdpRawAttributesWithCacheModeAsync(cacheMode BluetoothCacheMode) *IAsyncOperation[*IMapView[uint32, *IBuffer]] {
	var _result *IAsyncOperation[*IMapView[uint32, *IBuffer]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSdpRawAttributesWithCacheModeAsync, uintptr(unsafe.Pointer(this)), uintptr(cacheMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 536CED14-EBCD-49FE-BF9F-40EFC689B20D
var IID_IRfcommDeviceService2 = syscall.GUID{0x536CED14, 0xEBCD, 0x49FE,
	[8]byte{0xBF, 0x9F, 0x40, 0xEF, 0xC6, 0x89, 0xB2, 0x0D}}

type IRfcommDeviceService2Interface interface {
	win32.IInspectableInterface
	Get_Device() *IBluetoothDevice
}

type IRfcommDeviceService2Vtbl struct {
	win32.IInspectableVtbl
	Get_Device uintptr
}

type IRfcommDeviceService2 struct {
	win32.IInspectable
}

func (this *IRfcommDeviceService2) Vtbl() *IRfcommDeviceService2Vtbl {
	return (*IRfcommDeviceService2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRfcommDeviceService2) Get_Device() *IBluetoothDevice {
	var _result *IBluetoothDevice
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Device, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1C22ACE6-DD44-4D23-866D-8F3486EE6490
var IID_IRfcommDeviceService3 = syscall.GUID{0x1C22ACE6, 0xDD44, 0x4D23,
	[8]byte{0x86, 0x6D, 0x8F, 0x34, 0x86, 0xEE, 0x64, 0x90}}

type IRfcommDeviceService3Interface interface {
	win32.IInspectableInterface
	Get_DeviceAccessInformation() *IDeviceAccessInformation
	RequestAccessAsync() *IAsyncOperation[DeviceAccessStatus]
}

type IRfcommDeviceService3Vtbl struct {
	win32.IInspectableVtbl
	Get_DeviceAccessInformation uintptr
	RequestAccessAsync          uintptr
}

type IRfcommDeviceService3 struct {
	win32.IInspectable
}

func (this *IRfcommDeviceService3) Vtbl() *IRfcommDeviceService3Vtbl {
	return (*IRfcommDeviceService3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRfcommDeviceService3) Get_DeviceAccessInformation() *IDeviceAccessInformation {
	var _result *IDeviceAccessInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceAccessInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRfcommDeviceService3) RequestAccessAsync() *IAsyncOperation[DeviceAccessStatus] {
	var _result *IAsyncOperation[DeviceAccessStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A4A149EF-626D-41AC-B253-87AC5C27E28A
var IID_IRfcommDeviceServiceStatics = syscall.GUID{0xA4A149EF, 0x626D, 0x41AC,
	[8]byte{0xB2, 0x53, 0x87, 0xAC, 0x5C, 0x27, 0xE2, 0x8A}}

type IRfcommDeviceServiceStaticsInterface interface {
	win32.IInspectableInterface
	FromIdAsync(deviceId string) *IAsyncOperation[*IRfcommDeviceService]
	GetDeviceSelector(serviceId *IRfcommServiceId) string
}

type IRfcommDeviceServiceStaticsVtbl struct {
	win32.IInspectableVtbl
	FromIdAsync       uintptr
	GetDeviceSelector uintptr
}

type IRfcommDeviceServiceStatics struct {
	win32.IInspectable
}

func (this *IRfcommDeviceServiceStatics) Vtbl() *IRfcommDeviceServiceStaticsVtbl {
	return (*IRfcommDeviceServiceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRfcommDeviceServiceStatics) FromIdAsync(deviceId string) *IAsyncOperation[*IRfcommDeviceService] {
	var _result *IAsyncOperation[*IRfcommDeviceService]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRfcommDeviceServiceStatics) GetDeviceSelector(serviceId *IRfcommServiceId) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(serviceId)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// AA8CB1C9-E78D-4BE4-8076-0A3D87A0A05F
var IID_IRfcommDeviceServiceStatics2 = syscall.GUID{0xAA8CB1C9, 0xE78D, 0x4BE4,
	[8]byte{0x80, 0x76, 0x0A, 0x3D, 0x87, 0xA0, 0xA0, 0x5F}}

type IRfcommDeviceServiceStatics2Interface interface {
	win32.IInspectableInterface
	GetDeviceSelectorForBluetoothDevice(bluetoothDevice *IBluetoothDevice) string
	GetDeviceSelectorForBluetoothDeviceWithCacheMode(bluetoothDevice *IBluetoothDevice, cacheMode BluetoothCacheMode) string
	GetDeviceSelectorForBluetoothDeviceAndServiceId(bluetoothDevice *IBluetoothDevice, serviceId *IRfcommServiceId) string
	GetDeviceSelectorForBluetoothDeviceAndServiceIdWithCacheMode(bluetoothDevice *IBluetoothDevice, serviceId *IRfcommServiceId, cacheMode BluetoothCacheMode) string
}

type IRfcommDeviceServiceStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelectorForBluetoothDevice                          uintptr
	GetDeviceSelectorForBluetoothDeviceWithCacheMode             uintptr
	GetDeviceSelectorForBluetoothDeviceAndServiceId              uintptr
	GetDeviceSelectorForBluetoothDeviceAndServiceIdWithCacheMode uintptr
}

type IRfcommDeviceServiceStatics2 struct {
	win32.IInspectable
}

func (this *IRfcommDeviceServiceStatics2) Vtbl() *IRfcommDeviceServiceStatics2Vtbl {
	return (*IRfcommDeviceServiceStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRfcommDeviceServiceStatics2) GetDeviceSelectorForBluetoothDevice(bluetoothDevice *IBluetoothDevice) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorForBluetoothDevice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bluetoothDevice)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IRfcommDeviceServiceStatics2) GetDeviceSelectorForBluetoothDeviceWithCacheMode(bluetoothDevice *IBluetoothDevice, cacheMode BluetoothCacheMode) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorForBluetoothDeviceWithCacheMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bluetoothDevice)), uintptr(cacheMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IRfcommDeviceServiceStatics2) GetDeviceSelectorForBluetoothDeviceAndServiceId(bluetoothDevice *IBluetoothDevice, serviceId *IRfcommServiceId) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorForBluetoothDeviceAndServiceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bluetoothDevice)), uintptr(unsafe.Pointer(serviceId)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IRfcommDeviceServiceStatics2) GetDeviceSelectorForBluetoothDeviceAndServiceIdWithCacheMode(bluetoothDevice *IBluetoothDevice, serviceId *IRfcommServiceId, cacheMode BluetoothCacheMode) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorForBluetoothDeviceAndServiceIdWithCacheMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bluetoothDevice)), uintptr(unsafe.Pointer(serviceId)), uintptr(cacheMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 3B48388C-7CCF-488E-9625-D259A5732D55
var IID_IRfcommDeviceServicesResult = syscall.GUID{0x3B48388C, 0x7CCF, 0x488E,
	[8]byte{0x96, 0x25, 0xD2, 0x59, 0xA5, 0x73, 0x2D, 0x55}}

type IRfcommDeviceServicesResultInterface interface {
	win32.IInspectableInterface
	Get_Error() BluetoothError
	Get_Services() *IVectorView[*IRfcommDeviceService]
}

type IRfcommDeviceServicesResultVtbl struct {
	win32.IInspectableVtbl
	Get_Error    uintptr
	Get_Services uintptr
}

type IRfcommDeviceServicesResult struct {
	win32.IInspectable
}

func (this *IRfcommDeviceServicesResult) Vtbl() *IRfcommDeviceServicesResultVtbl {
	return (*IRfcommDeviceServicesResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRfcommDeviceServicesResult) Get_Error() BluetoothError {
	var _result BluetoothError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Error, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRfcommDeviceServicesResult) Get_Services() *IVectorView[*IRfcommDeviceService] {
	var _result *IVectorView[*IRfcommDeviceService]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Services, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 22629204-7E02-4017-8136-DA1B6A1B9BBF
var IID_IRfcommServiceId = syscall.GUID{0x22629204, 0x7E02, 0x4017,
	[8]byte{0x81, 0x36, 0xDA, 0x1B, 0x6A, 0x1B, 0x9B, 0xBF}}

type IRfcommServiceIdInterface interface {
	win32.IInspectableInterface
	Get_Uuid() syscall.GUID
	AsShortId() uint32
	AsString() string
}

type IRfcommServiceIdVtbl struct {
	win32.IInspectableVtbl
	Get_Uuid  uintptr
	AsShortId uintptr
	AsString  uintptr
}

type IRfcommServiceId struct {
	win32.IInspectable
}

func (this *IRfcommServiceId) Vtbl() *IRfcommServiceIdVtbl {
	return (*IRfcommServiceIdVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRfcommServiceId) Get_Uuid() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Uuid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRfcommServiceId) AsShortId() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AsShortId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRfcommServiceId) AsString() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AsString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 2A179EBA-A975-46E3-B56B-08FFD783A5FE
var IID_IRfcommServiceIdStatics = syscall.GUID{0x2A179EBA, 0xA975, 0x46E3,
	[8]byte{0xB5, 0x6B, 0x08, 0xFF, 0xD7, 0x83, 0xA5, 0xFE}}

type IRfcommServiceIdStaticsInterface interface {
	win32.IInspectableInterface
	FromUuid(uuid syscall.GUID) *IRfcommServiceId
	FromShortId(shortId uint32) *IRfcommServiceId
	Get_SerialPort() *IRfcommServiceId
	Get_ObexObjectPush() *IRfcommServiceId
	Get_ObexFileTransfer() *IRfcommServiceId
	Get_PhoneBookAccessPce() *IRfcommServiceId
	Get_PhoneBookAccessPse() *IRfcommServiceId
	Get_GenericFileTransfer() *IRfcommServiceId
}

type IRfcommServiceIdStaticsVtbl struct {
	win32.IInspectableVtbl
	FromUuid                uintptr
	FromShortId             uintptr
	Get_SerialPort          uintptr
	Get_ObexObjectPush      uintptr
	Get_ObexFileTransfer    uintptr
	Get_PhoneBookAccessPce  uintptr
	Get_PhoneBookAccessPse  uintptr
	Get_GenericFileTransfer uintptr
}

type IRfcommServiceIdStatics struct {
	win32.IInspectable
}

func (this *IRfcommServiceIdStatics) Vtbl() *IRfcommServiceIdStaticsVtbl {
	return (*IRfcommServiceIdStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRfcommServiceIdStatics) FromUuid(uuid syscall.GUID) *IRfcommServiceId {
	var _result *IRfcommServiceId
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromUuid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&uuid)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRfcommServiceIdStatics) FromShortId(shortId uint32) *IRfcommServiceId {
	var _result *IRfcommServiceId
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromShortId, uintptr(unsafe.Pointer(this)), uintptr(shortId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRfcommServiceIdStatics) Get_SerialPort() *IRfcommServiceId {
	var _result *IRfcommServiceId
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SerialPort, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRfcommServiceIdStatics) Get_ObexObjectPush() *IRfcommServiceId {
	var _result *IRfcommServiceId
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ObexObjectPush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRfcommServiceIdStatics) Get_ObexFileTransfer() *IRfcommServiceId {
	var _result *IRfcommServiceId
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ObexFileTransfer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRfcommServiceIdStatics) Get_PhoneBookAccessPce() *IRfcommServiceId {
	var _result *IRfcommServiceId
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhoneBookAccessPce, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRfcommServiceIdStatics) Get_PhoneBookAccessPse() *IRfcommServiceId {
	var _result *IRfcommServiceId
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhoneBookAccessPse, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRfcommServiceIdStatics) Get_GenericFileTransfer() *IRfcommServiceId {
	var _result *IRfcommServiceId
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GenericFileTransfer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EADBFDC4-B1F6-44FF-9F7C-E7A82AB86821
var IID_IRfcommServiceProvider = syscall.GUID{0xEADBFDC4, 0xB1F6, 0x44FF,
	[8]byte{0x9F, 0x7C, 0xE7, 0xA8, 0x2A, 0xB8, 0x68, 0x21}}

type IRfcommServiceProviderInterface interface {
	win32.IInspectableInterface
	Get_ServiceId() *IRfcommServiceId
	Get_SdpRawAttributes() *IMap[uint32, *IBuffer]
	StartAdvertising(listener *IStreamSocketListener)
	StopAdvertising()
}

type IRfcommServiceProviderVtbl struct {
	win32.IInspectableVtbl
	Get_ServiceId        uintptr
	Get_SdpRawAttributes uintptr
	StartAdvertising     uintptr
	StopAdvertising      uintptr
}

type IRfcommServiceProvider struct {
	win32.IInspectable
}

func (this *IRfcommServiceProvider) Vtbl() *IRfcommServiceProviderVtbl {
	return (*IRfcommServiceProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRfcommServiceProvider) Get_ServiceId() *IRfcommServiceId {
	var _result *IRfcommServiceId
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRfcommServiceProvider) Get_SdpRawAttributes() *IMap[uint32, *IBuffer] {
	var _result *IMap[uint32, *IBuffer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SdpRawAttributes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRfcommServiceProvider) StartAdvertising(listener *IStreamSocketListener) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartAdvertising, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(listener)))
	_ = _hr
}

func (this *IRfcommServiceProvider) StopAdvertising() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopAdvertising, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 736BDFC6-3C81-4D1E-BAF2-DDBB81284512
var IID_IRfcommServiceProvider2 = syscall.GUID{0x736BDFC6, 0x3C81, 0x4D1E,
	[8]byte{0xBA, 0xF2, 0xDD, 0xBB, 0x81, 0x28, 0x45, 0x12}}

type IRfcommServiceProvider2Interface interface {
	win32.IInspectableInterface
	StartAdvertisingWithRadioDiscoverability(listener *IStreamSocketListener, radioDiscoverable bool)
}

type IRfcommServiceProvider2Vtbl struct {
	win32.IInspectableVtbl
	StartAdvertisingWithRadioDiscoverability uintptr
}

type IRfcommServiceProvider2 struct {
	win32.IInspectable
}

func (this *IRfcommServiceProvider2) Vtbl() *IRfcommServiceProvider2Vtbl {
	return (*IRfcommServiceProvider2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRfcommServiceProvider2) StartAdvertisingWithRadioDiscoverability(listener *IStreamSocketListener, radioDiscoverable bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartAdvertisingWithRadioDiscoverability, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(listener)), uintptr(*(*byte)(unsafe.Pointer(&radioDiscoverable))))
	_ = _hr
}

// 98888303-69CA-413A-84F7-4344C7292997
var IID_IRfcommServiceProviderStatics = syscall.GUID{0x98888303, 0x69CA, 0x413A,
	[8]byte{0x84, 0xF7, 0x43, 0x44, 0xC7, 0x29, 0x29, 0x97}}

type IRfcommServiceProviderStaticsInterface interface {
	win32.IInspectableInterface
	CreateAsync(serviceId *IRfcommServiceId) *IAsyncOperation[*IRfcommServiceProvider]
}

type IRfcommServiceProviderStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateAsync uintptr
}

type IRfcommServiceProviderStatics struct {
	win32.IInspectable
}

func (this *IRfcommServiceProviderStatics) Vtbl() *IRfcommServiceProviderStaticsVtbl {
	return (*IRfcommServiceProviderStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRfcommServiceProviderStatics) CreateAsync(serviceId *IRfcommServiceId) *IAsyncOperation[*IRfcommServiceProvider] {
	var _result *IAsyncOperation[*IRfcommServiceProvider]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(serviceId)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type RfcommDeviceService struct {
	RtClass
	*IRfcommDeviceService
}

func NewIRfcommDeviceServiceStatics() *IRfcommDeviceServiceStatics {
	var p *IRfcommDeviceServiceStatics
	hs := NewHStr("Windows.Devices.Bluetooth.Rfcomm.RfcommDeviceService")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IRfcommDeviceServiceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIRfcommDeviceServiceStatics2() *IRfcommDeviceServiceStatics2 {
	var p *IRfcommDeviceServiceStatics2
	hs := NewHStr("Windows.Devices.Bluetooth.Rfcomm.RfcommDeviceService")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IRfcommDeviceServiceStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type RfcommDeviceServicesResult struct {
	RtClass
	*IRfcommDeviceServicesResult
}

type RfcommServiceId struct {
	RtClass
	*IRfcommServiceId
}

func NewIRfcommServiceIdStatics() *IRfcommServiceIdStatics {
	var p *IRfcommServiceIdStatics
	hs := NewHStr("Windows.Devices.Bluetooth.Rfcomm.RfcommServiceId")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IRfcommServiceIdStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type RfcommServiceProvider struct {
	RtClass
	*IRfcommServiceProvider
}

func NewIRfcommServiceProviderStatics() *IRfcommServiceProviderStatics {
	var p *IRfcommServiceProviderStatics
	hs := NewHStr("Windows.Devices.Bluetooth.Rfcomm.RfcommServiceProvider")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IRfcommServiceProviderStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
