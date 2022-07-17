package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type BluetoothEventTriggeringMode int32

const (
	BluetoothEventTriggeringMode_Serial     BluetoothEventTriggeringMode = 0
	BluetoothEventTriggeringMode_Batch      BluetoothEventTriggeringMode = 1
	BluetoothEventTriggeringMode_KeepLatest BluetoothEventTriggeringMode = 2
)

// interfaces

// 610ECA86-3480-41C9-A918-7DDADF207E00
var IID_IBluetoothLEAdvertisementPublisherTriggerDetails = syscall.GUID{0x610ECA86, 0x3480, 0x41C9,
	[8]byte{0xA9, 0x18, 0x7D, 0xDA, 0xDF, 0x20, 0x7E, 0x00}}

type IBluetoothLEAdvertisementPublisherTriggerDetailsInterface interface {
	win32.IInspectableInterface
	Get_Status() BluetoothLEAdvertisementPublisherStatus
	Get_Error() BluetoothError
}

type IBluetoothLEAdvertisementPublisherTriggerDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
	Get_Error  uintptr
}

type IBluetoothLEAdvertisementPublisherTriggerDetails struct {
	win32.IInspectable
}

func (this *IBluetoothLEAdvertisementPublisherTriggerDetails) Vtbl() *IBluetoothLEAdvertisementPublisherTriggerDetailsVtbl {
	return (*IBluetoothLEAdvertisementPublisherTriggerDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAdvertisementPublisherTriggerDetails) Get_Status() BluetoothLEAdvertisementPublisherStatus {
	var _result BluetoothLEAdvertisementPublisherStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementPublisherTriggerDetails) Get_Error() BluetoothError {
	var _result BluetoothError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Error, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D4A3D025-C601-42D6-9829-4CCB3F5CD77F
var IID_IBluetoothLEAdvertisementPublisherTriggerDetails2 = syscall.GUID{0xD4A3D025, 0xC601, 0x42D6,
	[8]byte{0x98, 0x29, 0x4C, 0xCB, 0x3F, 0x5C, 0xD7, 0x7F}}

type IBluetoothLEAdvertisementPublisherTriggerDetails2Interface interface {
	win32.IInspectableInterface
	Get_SelectedTransmitPowerLevelInDBm() *IReference[int16]
}

type IBluetoothLEAdvertisementPublisherTriggerDetails2Vtbl struct {
	win32.IInspectableVtbl
	Get_SelectedTransmitPowerLevelInDBm uintptr
}

type IBluetoothLEAdvertisementPublisherTriggerDetails2 struct {
	win32.IInspectable
}

func (this *IBluetoothLEAdvertisementPublisherTriggerDetails2) Vtbl() *IBluetoothLEAdvertisementPublisherTriggerDetails2Vtbl {
	return (*IBluetoothLEAdvertisementPublisherTriggerDetails2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAdvertisementPublisherTriggerDetails2) Get_SelectedTransmitPowerLevelInDBm() *IReference[int16] {
	var _result *IReference[int16]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectedTransmitPowerLevelInDBm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A7DB5AD7-2257-4E69-9784-FEE645C1DCE0
var IID_IBluetoothLEAdvertisementWatcherTriggerDetails = syscall.GUID{0xA7DB5AD7, 0x2257, 0x4E69,
	[8]byte{0x97, 0x84, 0xFE, 0xE6, 0x45, 0xC1, 0xDC, 0xE0}}

type IBluetoothLEAdvertisementWatcherTriggerDetailsInterface interface {
	win32.IInspectableInterface
	Get_Error() BluetoothError
	Get_Advertisements() *IVectorView[*IBluetoothLEAdvertisementReceivedEventArgs]
	Get_SignalStrengthFilter() *IBluetoothSignalStrengthFilter
}

type IBluetoothLEAdvertisementWatcherTriggerDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_Error                uintptr
	Get_Advertisements       uintptr
	Get_SignalStrengthFilter uintptr
}

type IBluetoothLEAdvertisementWatcherTriggerDetails struct {
	win32.IInspectable
}

func (this *IBluetoothLEAdvertisementWatcherTriggerDetails) Vtbl() *IBluetoothLEAdvertisementWatcherTriggerDetailsVtbl {
	return (*IBluetoothLEAdvertisementWatcherTriggerDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAdvertisementWatcherTriggerDetails) Get_Error() BluetoothError {
	var _result BluetoothError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Error, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementWatcherTriggerDetails) Get_Advertisements() *IVectorView[*IBluetoothLEAdvertisementReceivedEventArgs] {
	var _result *IVectorView[*IBluetoothLEAdvertisementReceivedEventArgs]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Advertisements, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEAdvertisementWatcherTriggerDetails) Get_SignalStrengthFilter() *IBluetoothSignalStrengthFilter {
	var _result *IBluetoothSignalStrengthFilter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SignalStrengthFilter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9BA03B18-0FEC-436A-93B1-F46C697532A2
var IID_IGattCharacteristicNotificationTriggerDetails = syscall.GUID{0x9BA03B18, 0x0FEC, 0x436A,
	[8]byte{0x93, 0xB1, 0xF4, 0x6C, 0x69, 0x75, 0x32, 0xA2}}

type IGattCharacteristicNotificationTriggerDetailsInterface interface {
	win32.IInspectableInterface
	Get_Characteristic() *IGattCharacteristic
	Get_Value() *IBuffer
}

type IGattCharacteristicNotificationTriggerDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_Characteristic uintptr
	Get_Value          uintptr
}

type IGattCharacteristicNotificationTriggerDetails struct {
	win32.IInspectable
}

func (this *IGattCharacteristicNotificationTriggerDetails) Vtbl() *IGattCharacteristicNotificationTriggerDetailsVtbl {
	return (*IGattCharacteristicNotificationTriggerDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattCharacteristicNotificationTriggerDetails) Get_Characteristic() *IGattCharacteristic {
	var _result *IGattCharacteristic
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Characteristic, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattCharacteristicNotificationTriggerDetails) Get_Value() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 727A50DC-949D-454A-B192-983467E3D50F
var IID_IGattCharacteristicNotificationTriggerDetails2 = syscall.GUID{0x727A50DC, 0x949D, 0x454A,
	[8]byte{0xB1, 0x92, 0x98, 0x34, 0x67, 0xE3, 0xD5, 0x0F}}

type IGattCharacteristicNotificationTriggerDetails2Interface interface {
	win32.IInspectableInterface
	Get_Error() BluetoothError
	Get_EventTriggeringMode() BluetoothEventTriggeringMode
	Get_ValueChangedEvents() *IVectorView[*IGattValueChangedEventArgs]
}

type IGattCharacteristicNotificationTriggerDetails2Vtbl struct {
	win32.IInspectableVtbl
	Get_Error               uintptr
	Get_EventTriggeringMode uintptr
	Get_ValueChangedEvents  uintptr
}

type IGattCharacteristicNotificationTriggerDetails2 struct {
	win32.IInspectable
}

func (this *IGattCharacteristicNotificationTriggerDetails2) Vtbl() *IGattCharacteristicNotificationTriggerDetails2Vtbl {
	return (*IGattCharacteristicNotificationTriggerDetails2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattCharacteristicNotificationTriggerDetails2) Get_Error() BluetoothError {
	var _result BluetoothError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Error, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicNotificationTriggerDetails2) Get_EventTriggeringMode() BluetoothEventTriggeringMode {
	var _result BluetoothEventTriggeringMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EventTriggeringMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicNotificationTriggerDetails2) Get_ValueChangedEvents() *IVectorView[*IGattValueChangedEventArgs] {
	var _result *IVectorView[*IGattValueChangedEventArgs]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ValueChangedEvents, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7FA1B9B9-2F13-40B5-9582-8EB78E98EF13
var IID_IGattServiceProviderConnection = syscall.GUID{0x7FA1B9B9, 0x2F13, 0x40B5,
	[8]byte{0x95, 0x82, 0x8E, 0xB7, 0x8E, 0x98, 0xEF, 0x13}}

type IGattServiceProviderConnectionInterface interface {
	win32.IInspectableInterface
	Get_TriggerId() string
	Get_Service() *IGattLocalService
	Start()
}

type IGattServiceProviderConnectionVtbl struct {
	win32.IInspectableVtbl
	Get_TriggerId uintptr
	Get_Service   uintptr
	Start         uintptr
}

type IGattServiceProviderConnection struct {
	win32.IInspectable
}

func (this *IGattServiceProviderConnection) Vtbl() *IGattServiceProviderConnectionVtbl {
	return (*IGattServiceProviderConnectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattServiceProviderConnection) Get_TriggerId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TriggerId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGattServiceProviderConnection) Get_Service() *IGattLocalService {
	var _result *IGattLocalService
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Service, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattServiceProviderConnection) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 3D509F4B-0B0E-4466-B8CD-6EBDDA1FA17D
var IID_IGattServiceProviderConnectionStatics = syscall.GUID{0x3D509F4B, 0x0B0E, 0x4466,
	[8]byte{0xB8, 0xCD, 0x6E, 0xBD, 0xDA, 0x1F, 0xA1, 0x7D}}

type IGattServiceProviderConnectionStaticsInterface interface {
	win32.IInspectableInterface
	Get_AllServices() *IMapView[string, *IGattServiceProviderConnection]
}

type IGattServiceProviderConnectionStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_AllServices uintptr
}

type IGattServiceProviderConnectionStatics struct {
	win32.IInspectable
}

func (this *IGattServiceProviderConnectionStatics) Vtbl() *IGattServiceProviderConnectionStaticsVtbl {
	return (*IGattServiceProviderConnectionStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattServiceProviderConnectionStatics) Get_AllServices() *IMapView[string, *IGattServiceProviderConnection] {
	var _result *IMapView[string, *IGattServiceProviderConnection]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllServices, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AE8C0625-05FF-4AFB-B16A-DE95F3CF0158
var IID_IGattServiceProviderTriggerDetails = syscall.GUID{0xAE8C0625, 0x05FF, 0x4AFB,
	[8]byte{0xB1, 0x6A, 0xDE, 0x95, 0xF3, 0xCF, 0x01, 0x58}}

type IGattServiceProviderTriggerDetailsInterface interface {
	win32.IInspectableInterface
	Get_Connection() *IGattServiceProviderConnection
}

type IGattServiceProviderTriggerDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_Connection uintptr
}

type IGattServiceProviderTriggerDetails struct {
	win32.IInspectable
}

func (this *IGattServiceProviderTriggerDetails) Vtbl() *IGattServiceProviderTriggerDetailsVtbl {
	return (*IGattServiceProviderTriggerDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattServiceProviderTriggerDetails) Get_Connection() *IGattServiceProviderConnection {
	var _result *IGattServiceProviderConnection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Connection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F922734D-2E3C-4EFC-AB59-FC5CF96F97E3
var IID_IRfcommConnectionTriggerDetails = syscall.GUID{0xF922734D, 0x2E3C, 0x4EFC,
	[8]byte{0xAB, 0x59, 0xFC, 0x5C, 0xF9, 0x6F, 0x97, 0xE3}}

type IRfcommConnectionTriggerDetailsInterface interface {
	win32.IInspectableInterface
	Get_Socket() *IStreamSocket
	Get_Incoming() bool
	Get_RemoteDevice() *IBluetoothDevice
}

type IRfcommConnectionTriggerDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_Socket       uintptr
	Get_Incoming     uintptr
	Get_RemoteDevice uintptr
}

type IRfcommConnectionTriggerDetails struct {
	win32.IInspectable
}

func (this *IRfcommConnectionTriggerDetails) Vtbl() *IRfcommConnectionTriggerDetailsVtbl {
	return (*IRfcommConnectionTriggerDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRfcommConnectionTriggerDetails) Get_Socket() *IStreamSocket {
	var _result *IStreamSocket
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Socket, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRfcommConnectionTriggerDetails) Get_Incoming() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Incoming, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRfcommConnectionTriggerDetails) Get_RemoteDevice() *IBluetoothDevice {
	var _result *IBluetoothDevice
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemoteDevice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6D3E75A8-5429-4059-92E3-1E8B65528707
var IID_IRfcommInboundConnectionInformation = syscall.GUID{0x6D3E75A8, 0x5429, 0x4059,
	[8]byte{0x92, 0xE3, 0x1E, 0x8B, 0x65, 0x52, 0x87, 0x07}}

type IRfcommInboundConnectionInformationInterface interface {
	win32.IInspectableInterface
	Get_SdpRecord() *IBuffer
	Put_SdpRecord(value *IBuffer)
	Get_LocalServiceId() *IRfcommServiceId
	Put_LocalServiceId(value *IRfcommServiceId)
	Get_ServiceCapabilities() BluetoothServiceCapabilities
	Put_ServiceCapabilities(value BluetoothServiceCapabilities)
}

type IRfcommInboundConnectionInformationVtbl struct {
	win32.IInspectableVtbl
	Get_SdpRecord           uintptr
	Put_SdpRecord           uintptr
	Get_LocalServiceId      uintptr
	Put_LocalServiceId      uintptr
	Get_ServiceCapabilities uintptr
	Put_ServiceCapabilities uintptr
}

type IRfcommInboundConnectionInformation struct {
	win32.IInspectable
}

func (this *IRfcommInboundConnectionInformation) Vtbl() *IRfcommInboundConnectionInformationVtbl {
	return (*IRfcommInboundConnectionInformationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRfcommInboundConnectionInformation) Get_SdpRecord() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SdpRecord, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRfcommInboundConnectionInformation) Put_SdpRecord(value *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SdpRecord, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IRfcommInboundConnectionInformation) Get_LocalServiceId() *IRfcommServiceId {
	var _result *IRfcommServiceId
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalServiceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRfcommInboundConnectionInformation) Put_LocalServiceId(value *IRfcommServiceId) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LocalServiceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IRfcommInboundConnectionInformation) Get_ServiceCapabilities() BluetoothServiceCapabilities {
	var _result BluetoothServiceCapabilities
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceCapabilities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRfcommInboundConnectionInformation) Put_ServiceCapabilities(value BluetoothServiceCapabilities) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ServiceCapabilities, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// B091227B-F434-4CB0-99B1-4AB8CEDAEDD7
var IID_IRfcommOutboundConnectionInformation = syscall.GUID{0xB091227B, 0xF434, 0x4CB0,
	[8]byte{0x99, 0xB1, 0x4A, 0xB8, 0xCE, 0xDA, 0xED, 0xD7}}

type IRfcommOutboundConnectionInformationInterface interface {
	win32.IInspectableInterface
	Get_RemoteServiceId() *IRfcommServiceId
	Put_RemoteServiceId(value *IRfcommServiceId)
}

type IRfcommOutboundConnectionInformationVtbl struct {
	win32.IInspectableVtbl
	Get_RemoteServiceId uintptr
	Put_RemoteServiceId uintptr
}

type IRfcommOutboundConnectionInformation struct {
	win32.IInspectable
}

func (this *IRfcommOutboundConnectionInformation) Vtbl() *IRfcommOutboundConnectionInformationVtbl {
	return (*IRfcommOutboundConnectionInformationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRfcommOutboundConnectionInformation) Get_RemoteServiceId() *IRfcommServiceId {
	var _result *IRfcommServiceId
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemoteServiceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRfcommOutboundConnectionInformation) Put_RemoteServiceId(value *IRfcommServiceId) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RemoteServiceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// classes

type GattServiceProviderConnection struct {
	RtClass
	*IGattServiceProviderConnection
}

func NewIGattServiceProviderConnectionStatics() *IGattServiceProviderConnectionStatics {
	var p *IGattServiceProviderConnectionStatics
	hs := NewHStr("Windows.Devices.Bluetooth.Background.GattServiceProviderConnection")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGattServiceProviderConnectionStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
