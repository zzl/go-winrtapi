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
// flags
type GattCharacteristicProperties uint32

const (
	GattCharacteristicProperties_None                      GattCharacteristicProperties = 0
	GattCharacteristicProperties_Broadcast                 GattCharacteristicProperties = 1
	GattCharacteristicProperties_Read                      GattCharacteristicProperties = 2
	GattCharacteristicProperties_WriteWithoutResponse      GattCharacteristicProperties = 4
	GattCharacteristicProperties_Write                     GattCharacteristicProperties = 8
	GattCharacteristicProperties_Notify                    GattCharacteristicProperties = 16
	GattCharacteristicProperties_Indicate                  GattCharacteristicProperties = 32
	GattCharacteristicProperties_AuthenticatedSignedWrites GattCharacteristicProperties = 64
	GattCharacteristicProperties_ExtendedProperties        GattCharacteristicProperties = 128
	GattCharacteristicProperties_ReliableWrites            GattCharacteristicProperties = 256
	GattCharacteristicProperties_WritableAuxiliaries       GattCharacteristicProperties = 512
)

// enum
type GattClientCharacteristicConfigurationDescriptorValue int32

const (
	GattClientCharacteristicConfigurationDescriptorValue_None     GattClientCharacteristicConfigurationDescriptorValue = 0
	GattClientCharacteristicConfigurationDescriptorValue_Notify   GattClientCharacteristicConfigurationDescriptorValue = 1
	GattClientCharacteristicConfigurationDescriptorValue_Indicate GattClientCharacteristicConfigurationDescriptorValue = 2
)

// enum
type GattCommunicationStatus int32

const (
	GattCommunicationStatus_Success       GattCommunicationStatus = 0
	GattCommunicationStatus_Unreachable   GattCommunicationStatus = 1
	GattCommunicationStatus_ProtocolError GattCommunicationStatus = 2
	GattCommunicationStatus_AccessDenied  GattCommunicationStatus = 3
)

// enum
type GattOpenStatus int32

const (
	GattOpenStatus_Unspecified      GattOpenStatus = 0
	GattOpenStatus_Success          GattOpenStatus = 1
	GattOpenStatus_AlreadyOpened    GattOpenStatus = 2
	GattOpenStatus_NotFound         GattOpenStatus = 3
	GattOpenStatus_SharingViolation GattOpenStatus = 4
	GattOpenStatus_AccessDenied     GattOpenStatus = 5
)

// enum
type GattProtectionLevel int32

const (
	GattProtectionLevel_Plain                               GattProtectionLevel = 0
	GattProtectionLevel_AuthenticationRequired              GattProtectionLevel = 1
	GattProtectionLevel_EncryptionRequired                  GattProtectionLevel = 2
	GattProtectionLevel_EncryptionAndAuthenticationRequired GattProtectionLevel = 3
)

// enum
type GattRequestState int32

const (
	GattRequestState_Pending   GattRequestState = 0
	GattRequestState_Completed GattRequestState = 1
	GattRequestState_Canceled  GattRequestState = 2
)

// enum
type GattServiceProviderAdvertisementStatus int32

const (
	GattServiceProviderAdvertisementStatus_Created                            GattServiceProviderAdvertisementStatus = 0
	GattServiceProviderAdvertisementStatus_Stopped                            GattServiceProviderAdvertisementStatus = 1
	GattServiceProviderAdvertisementStatus_Started                            GattServiceProviderAdvertisementStatus = 2
	GattServiceProviderAdvertisementStatus_Aborted                            GattServiceProviderAdvertisementStatus = 3
	GattServiceProviderAdvertisementStatus_StartedWithoutAllAdvertisementData GattServiceProviderAdvertisementStatus = 4
)

// enum
type GattSessionStatus int32

const (
	GattSessionStatus_Closed GattSessionStatus = 0
	GattSessionStatus_Active GattSessionStatus = 1
)

// enum
type GattSharingMode int32

const (
	GattSharingMode_Unspecified        GattSharingMode = 0
	GattSharingMode_Exclusive          GattSharingMode = 1
	GattSharingMode_SharedReadOnly     GattSharingMode = 2
	GattSharingMode_SharedReadAndWrite GattSharingMode = 3
)

// enum
type GattWriteOption int32

const (
	GattWriteOption_WriteWithResponse    GattWriteOption = 0
	GattWriteOption_WriteWithoutResponse GattWriteOption = 1
)

// interfaces

// 59CB50C1-5934-4F68-A198-EB864FA44E6B
var IID_IGattCharacteristic = syscall.GUID{0x59CB50C1, 0x5934, 0x4F68,
	[8]byte{0xA1, 0x98, 0xEB, 0x86, 0x4F, 0xA4, 0x4E, 0x6B}}

type IGattCharacteristicInterface interface {
	win32.IInspectableInterface
	GetDescriptors(descriptorUuid syscall.GUID) *IVectorView[*IGattDescriptor]
	Get_CharacteristicProperties() GattCharacteristicProperties
	Get_ProtectionLevel() GattProtectionLevel
	Put_ProtectionLevel(value GattProtectionLevel)
	Get_UserDescription() string
	Get_Uuid() syscall.GUID
	Get_AttributeHandle() uint16
	Get_PresentationFormats() *IVectorView[*IGattPresentationFormat]
	ReadValueAsync() *IAsyncOperation[*IGattReadResult]
	ReadValueWithCacheModeAsync(cacheMode BluetoothCacheMode) *IAsyncOperation[*IGattReadResult]
	WriteValueAsync(value *IBuffer) *IAsyncOperation[GattCommunicationStatus]
	WriteValueWithOptionAsync(value *IBuffer, writeOption GattWriteOption) *IAsyncOperation[GattCommunicationStatus]
	ReadClientCharacteristicConfigurationDescriptorAsync() *IAsyncOperation[*IGattReadClientCharacteristicConfigurationDescriptorResult]
	WriteClientCharacteristicConfigurationDescriptorAsync(clientCharacteristicConfigurationDescriptorValue GattClientCharacteristicConfigurationDescriptorValue) *IAsyncOperation[GattCommunicationStatus]
	Add_ValueChanged(valueChangedHandler TypedEventHandler[*IGattCharacteristic, *IGattValueChangedEventArgs]) EventRegistrationToken
	Remove_ValueChanged(valueChangedEventCookie EventRegistrationToken)
}

type IGattCharacteristicVtbl struct {
	win32.IInspectableVtbl
	GetDescriptors                                        uintptr
	Get_CharacteristicProperties                          uintptr
	Get_ProtectionLevel                                   uintptr
	Put_ProtectionLevel                                   uintptr
	Get_UserDescription                                   uintptr
	Get_Uuid                                              uintptr
	Get_AttributeHandle                                   uintptr
	Get_PresentationFormats                               uintptr
	ReadValueAsync                                        uintptr
	ReadValueWithCacheModeAsync                           uintptr
	WriteValueAsync                                       uintptr
	WriteValueWithOptionAsync                             uintptr
	ReadClientCharacteristicConfigurationDescriptorAsync  uintptr
	WriteClientCharacteristicConfigurationDescriptorAsync uintptr
	Add_ValueChanged                                      uintptr
	Remove_ValueChanged                                   uintptr
}

type IGattCharacteristic struct {
	win32.IInspectable
}

func (this *IGattCharacteristic) Vtbl() *IGattCharacteristicVtbl {
	return (*IGattCharacteristicVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattCharacteristic) GetDescriptors(descriptorUuid syscall.GUID) *IVectorView[*IGattDescriptor] {
	var _result *IVectorView[*IGattDescriptor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDescriptors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&descriptorUuid)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattCharacteristic) Get_CharacteristicProperties() GattCharacteristicProperties {
	var _result GattCharacteristicProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CharacteristicProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristic) Get_ProtectionLevel() GattProtectionLevel {
	var _result GattProtectionLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristic) Put_ProtectionLevel(value GattProtectionLevel) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGattCharacteristic) Get_UserDescription() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGattCharacteristic) Get_Uuid() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Uuid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristic) Get_AttributeHandle() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AttributeHandle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristic) Get_PresentationFormats() *IVectorView[*IGattPresentationFormat] {
	var _result *IVectorView[*IGattPresentationFormat]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PresentationFormats, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattCharacteristic) ReadValueAsync() *IAsyncOperation[*IGattReadResult] {
	var _result *IAsyncOperation[*IGattReadResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadValueAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattCharacteristic) ReadValueWithCacheModeAsync(cacheMode BluetoothCacheMode) *IAsyncOperation[*IGattReadResult] {
	var _result *IAsyncOperation[*IGattReadResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadValueWithCacheModeAsync, uintptr(unsafe.Pointer(this)), uintptr(cacheMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattCharacteristic) WriteValueAsync(value *IBuffer) *IAsyncOperation[GattCommunicationStatus] {
	var _result *IAsyncOperation[GattCommunicationStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteValueAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattCharacteristic) WriteValueWithOptionAsync(value *IBuffer, writeOption GattWriteOption) *IAsyncOperation[GattCommunicationStatus] {
	var _result *IAsyncOperation[GattCommunicationStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteValueWithOptionAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)), uintptr(writeOption), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattCharacteristic) ReadClientCharacteristicConfigurationDescriptorAsync() *IAsyncOperation[*IGattReadClientCharacteristicConfigurationDescriptorResult] {
	var _result *IAsyncOperation[*IGattReadClientCharacteristicConfigurationDescriptorResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadClientCharacteristicConfigurationDescriptorAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattCharacteristic) WriteClientCharacteristicConfigurationDescriptorAsync(clientCharacteristicConfigurationDescriptorValue GattClientCharacteristicConfigurationDescriptorValue) *IAsyncOperation[GattCommunicationStatus] {
	var _result *IAsyncOperation[GattCommunicationStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteClientCharacteristicConfigurationDescriptorAsync, uintptr(unsafe.Pointer(this)), uintptr(clientCharacteristicConfigurationDescriptorValue), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattCharacteristic) Add_ValueChanged(valueChangedHandler TypedEventHandler[*IGattCharacteristic, *IGattValueChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ValueChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(valueChangedHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristic) Remove_ValueChanged(valueChangedEventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ValueChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&valueChangedEventCookie)))
	_ = _hr
}

// AE1AB578-EC06-4764-B780-9835A1D35D6E
var IID_IGattCharacteristic2 = syscall.GUID{0xAE1AB578, 0xEC06, 0x4764,
	[8]byte{0xB7, 0x80, 0x98, 0x35, 0xA1, 0xD3, 0x5D, 0x6E}}

type IGattCharacteristic2Interface interface {
	win32.IInspectableInterface
	Get_Service() *IGattDeviceService
	GetAllDescriptors() *IVectorView[*IGattDescriptor]
}

type IGattCharacteristic2Vtbl struct {
	win32.IInspectableVtbl
	Get_Service       uintptr
	GetAllDescriptors uintptr
}

type IGattCharacteristic2 struct {
	win32.IInspectable
}

func (this *IGattCharacteristic2) Vtbl() *IGattCharacteristic2Vtbl {
	return (*IGattCharacteristic2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattCharacteristic2) Get_Service() *IGattDeviceService {
	var _result *IGattDeviceService
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Service, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattCharacteristic2) GetAllDescriptors() *IVectorView[*IGattDescriptor] {
	var _result *IVectorView[*IGattDescriptor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAllDescriptors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3F3C663E-93D4-406B-B817-DB81F8ED53B3
var IID_IGattCharacteristic3 = syscall.GUID{0x3F3C663E, 0x93D4, 0x406B,
	[8]byte{0xB8, 0x17, 0xDB, 0x81, 0xF8, 0xED, 0x53, 0xB3}}

type IGattCharacteristic3Interface interface {
	win32.IInspectableInterface
	GetDescriptorsAsync() *IAsyncOperation[*IGattDescriptorsResult]
	GetDescriptorsWithCacheModeAsync(cacheMode BluetoothCacheMode) *IAsyncOperation[*IGattDescriptorsResult]
	GetDescriptorsForUuidAsync(descriptorUuid syscall.GUID) *IAsyncOperation[*IGattDescriptorsResult]
	GetDescriptorsForUuidWithCacheModeAsync(descriptorUuid syscall.GUID, cacheMode BluetoothCacheMode) *IAsyncOperation[*IGattDescriptorsResult]
	WriteValueWithResultAsync(value *IBuffer) *IAsyncOperation[*IGattWriteResult]
	WriteValueWithResultAndOptionAsync(value *IBuffer, writeOption GattWriteOption) *IAsyncOperation[*IGattWriteResult]
	WriteClientCharacteristicConfigurationDescriptorWithResultAsync(clientCharacteristicConfigurationDescriptorValue GattClientCharacteristicConfigurationDescriptorValue) *IAsyncOperation[*IGattWriteResult]
}

type IGattCharacteristic3Vtbl struct {
	win32.IInspectableVtbl
	GetDescriptorsAsync                                             uintptr
	GetDescriptorsWithCacheModeAsync                                uintptr
	GetDescriptorsForUuidAsync                                      uintptr
	GetDescriptorsForUuidWithCacheModeAsync                         uintptr
	WriteValueWithResultAsync                                       uintptr
	WriteValueWithResultAndOptionAsync                              uintptr
	WriteClientCharacteristicConfigurationDescriptorWithResultAsync uintptr
}

type IGattCharacteristic3 struct {
	win32.IInspectable
}

func (this *IGattCharacteristic3) Vtbl() *IGattCharacteristic3Vtbl {
	return (*IGattCharacteristic3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattCharacteristic3) GetDescriptorsAsync() *IAsyncOperation[*IGattDescriptorsResult] {
	var _result *IAsyncOperation[*IGattDescriptorsResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDescriptorsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattCharacteristic3) GetDescriptorsWithCacheModeAsync(cacheMode BluetoothCacheMode) *IAsyncOperation[*IGattDescriptorsResult] {
	var _result *IAsyncOperation[*IGattDescriptorsResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDescriptorsWithCacheModeAsync, uintptr(unsafe.Pointer(this)), uintptr(cacheMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattCharacteristic3) GetDescriptorsForUuidAsync(descriptorUuid syscall.GUID) *IAsyncOperation[*IGattDescriptorsResult] {
	var _result *IAsyncOperation[*IGattDescriptorsResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDescriptorsForUuidAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&descriptorUuid)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattCharacteristic3) GetDescriptorsForUuidWithCacheModeAsync(descriptorUuid syscall.GUID, cacheMode BluetoothCacheMode) *IAsyncOperation[*IGattDescriptorsResult] {
	var _result *IAsyncOperation[*IGattDescriptorsResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDescriptorsForUuidWithCacheModeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&descriptorUuid)), uintptr(cacheMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattCharacteristic3) WriteValueWithResultAsync(value *IBuffer) *IAsyncOperation[*IGattWriteResult] {
	var _result *IAsyncOperation[*IGattWriteResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteValueWithResultAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattCharacteristic3) WriteValueWithResultAndOptionAsync(value *IBuffer, writeOption GattWriteOption) *IAsyncOperation[*IGattWriteResult] {
	var _result *IAsyncOperation[*IGattWriteResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteValueWithResultAndOptionAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)), uintptr(writeOption), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattCharacteristic3) WriteClientCharacteristicConfigurationDescriptorWithResultAsync(clientCharacteristicConfigurationDescriptorValue GattClientCharacteristicConfigurationDescriptorValue) *IAsyncOperation[*IGattWriteResult] {
	var _result *IAsyncOperation[*IGattWriteResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteClientCharacteristicConfigurationDescriptorWithResultAsync, uintptr(unsafe.Pointer(this)), uintptr(clientCharacteristicConfigurationDescriptorValue), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 59CB50C3-5934-4F68-A198-EB864FA44E6B
var IID_IGattCharacteristicStatics = syscall.GUID{0x59CB50C3, 0x5934, 0x4F68,
	[8]byte{0xA1, 0x98, 0xEB, 0x86, 0x4F, 0xA4, 0x4E, 0x6B}}

type IGattCharacteristicStaticsInterface interface {
	win32.IInspectableInterface
	ConvertShortIdToUuid(shortId uint16) syscall.GUID
}

type IGattCharacteristicStaticsVtbl struct {
	win32.IInspectableVtbl
	ConvertShortIdToUuid uintptr
}

type IGattCharacteristicStatics struct {
	win32.IInspectable
}

func (this *IGattCharacteristicStatics) Vtbl() *IGattCharacteristicStaticsVtbl {
	return (*IGattCharacteristicStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattCharacteristicStatics) ConvertShortIdToUuid(shortId uint16) syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConvertShortIdToUuid, uintptr(unsafe.Pointer(this)), uintptr(shortId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 58FA4586-B1DE-470C-B7DE-0D11FF44F4B7
var IID_IGattCharacteristicUuidsStatics = syscall.GUID{0x58FA4586, 0xB1DE, 0x470C,
	[8]byte{0xB7, 0xDE, 0x0D, 0x11, 0xFF, 0x44, 0xF4, 0xB7}}

type IGattCharacteristicUuidsStaticsInterface interface {
	win32.IInspectableInterface
	Get_BatteryLevel() syscall.GUID
	Get_BloodPressureFeature() syscall.GUID
	Get_BloodPressureMeasurement() syscall.GUID
	Get_BodySensorLocation() syscall.GUID
	Get_CscFeature() syscall.GUID
	Get_CscMeasurement() syscall.GUID
	Get_GlucoseFeature() syscall.GUID
	Get_GlucoseMeasurement() syscall.GUID
	Get_GlucoseMeasurementContext() syscall.GUID
	Get_HeartRateControlPoint() syscall.GUID
	Get_HeartRateMeasurement() syscall.GUID
	Get_IntermediateCuffPressure() syscall.GUID
	Get_IntermediateTemperature() syscall.GUID
	Get_MeasurementInterval() syscall.GUID
	Get_RecordAccessControlPoint() syscall.GUID
	Get_RscFeature() syscall.GUID
	Get_RscMeasurement() syscall.GUID
	Get_SCControlPoint() syscall.GUID
	Get_SensorLocation() syscall.GUID
	Get_TemperatureMeasurement() syscall.GUID
	Get_TemperatureType() syscall.GUID
}

type IGattCharacteristicUuidsStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_BatteryLevel              uintptr
	Get_BloodPressureFeature      uintptr
	Get_BloodPressureMeasurement  uintptr
	Get_BodySensorLocation        uintptr
	Get_CscFeature                uintptr
	Get_CscMeasurement            uintptr
	Get_GlucoseFeature            uintptr
	Get_GlucoseMeasurement        uintptr
	Get_GlucoseMeasurementContext uintptr
	Get_HeartRateControlPoint     uintptr
	Get_HeartRateMeasurement      uintptr
	Get_IntermediateCuffPressure  uintptr
	Get_IntermediateTemperature   uintptr
	Get_MeasurementInterval       uintptr
	Get_RecordAccessControlPoint  uintptr
	Get_RscFeature                uintptr
	Get_RscMeasurement            uintptr
	Get_SCControlPoint            uintptr
	Get_SensorLocation            uintptr
	Get_TemperatureMeasurement    uintptr
	Get_TemperatureType           uintptr
}

type IGattCharacteristicUuidsStatics struct {
	win32.IInspectable
}

func (this *IGattCharacteristicUuidsStatics) Vtbl() *IGattCharacteristicUuidsStaticsVtbl {
	return (*IGattCharacteristicUuidsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattCharacteristicUuidsStatics) Get_BatteryLevel() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BatteryLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics) Get_BloodPressureFeature() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BloodPressureFeature, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics) Get_BloodPressureMeasurement() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BloodPressureMeasurement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics) Get_BodySensorLocation() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BodySensorLocation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics) Get_CscFeature() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CscFeature, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics) Get_CscMeasurement() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CscMeasurement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics) Get_GlucoseFeature() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GlucoseFeature, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics) Get_GlucoseMeasurement() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GlucoseMeasurement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics) Get_GlucoseMeasurementContext() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GlucoseMeasurementContext, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics) Get_HeartRateControlPoint() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HeartRateControlPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics) Get_HeartRateMeasurement() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HeartRateMeasurement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics) Get_IntermediateCuffPressure() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IntermediateCuffPressure, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics) Get_IntermediateTemperature() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IntermediateTemperature, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics) Get_MeasurementInterval() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MeasurementInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics) Get_RecordAccessControlPoint() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RecordAccessControlPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics) Get_RscFeature() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RscFeature, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics) Get_RscMeasurement() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RscMeasurement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics) Get_SCControlPoint() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SCControlPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics) Get_SensorLocation() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SensorLocation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics) Get_TemperatureMeasurement() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TemperatureMeasurement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics) Get_TemperatureType() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TemperatureType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1855B425-D46E-4A2C-9C3F-ED6DEA29E7BE
var IID_IGattCharacteristicUuidsStatics2 = syscall.GUID{0x1855B425, 0xD46E, 0x4A2C,
	[8]byte{0x9C, 0x3F, 0xED, 0x6D, 0xEA, 0x29, 0xE7, 0xBE}}

type IGattCharacteristicUuidsStatics2Interface interface {
	win32.IInspectableInterface
	Get_AlertCategoryId() syscall.GUID
	Get_AlertCategoryIdBitMask() syscall.GUID
	Get_AlertLevel() syscall.GUID
	Get_AlertNotificationControlPoint() syscall.GUID
	Get_AlertStatus() syscall.GUID
	Get_GapAppearance() syscall.GUID
	Get_BootKeyboardInputReport() syscall.GUID
	Get_BootKeyboardOutputReport() syscall.GUID
	Get_BootMouseInputReport() syscall.GUID
	Get_CurrentTime() syscall.GUID
	Get_CyclingPowerControlPoint() syscall.GUID
	Get_CyclingPowerFeature() syscall.GUID
	Get_CyclingPowerMeasurement() syscall.GUID
	Get_CyclingPowerVector() syscall.GUID
	Get_DateTime() syscall.GUID
	Get_DayDateTime() syscall.GUID
	Get_DayOfWeek() syscall.GUID
	Get_GapDeviceName() syscall.GUID
	Get_DstOffset() syscall.GUID
	Get_ExactTime256() syscall.GUID
	Get_FirmwareRevisionString() syscall.GUID
	Get_HardwareRevisionString() syscall.GUID
	Get_HidControlPoint() syscall.GUID
	Get_HidInformation() syscall.GUID
	Get_Ieee1107320601RegulatoryCertificationDataList() syscall.GUID
	Get_LnControlPoint() syscall.GUID
	Get_LnFeature() syscall.GUID
	Get_LocalTimeInformation() syscall.GUID
	Get_LocationAndSpeed() syscall.GUID
	Get_ManufacturerNameString() syscall.GUID
	Get_ModelNumberString() syscall.GUID
	Get_Navigation() syscall.GUID
	Get_NewAlert() syscall.GUID
	Get_GapPeripheralPreferredConnectionParameters() syscall.GUID
	Get_GapPeripheralPrivacyFlag() syscall.GUID
	Get_PnpId() syscall.GUID
	Get_PositionQuality() syscall.GUID
	Get_ProtocolMode() syscall.GUID
	Get_GapReconnectionAddress() syscall.GUID
	Get_ReferenceTimeInformation() syscall.GUID
	Get_Report() syscall.GUID
	Get_ReportMap() syscall.GUID
	Get_RingerControlPoint() syscall.GUID
	Get_RingerSetting() syscall.GUID
	Get_ScanIntervalWindow() syscall.GUID
	Get_ScanRefresh() syscall.GUID
	Get_SerialNumberString() syscall.GUID
	Get_GattServiceChanged() syscall.GUID
	Get_SoftwareRevisionString() syscall.GUID
	Get_SupportedNewAlertCategory() syscall.GUID
	Get_SupportUnreadAlertCategory() syscall.GUID
	Get_SystemId() syscall.GUID
	Get_TimeAccuracy() syscall.GUID
	Get_TimeSource() syscall.GUID
	Get_TimeUpdateControlPoint() syscall.GUID
	Get_TimeUpdateState() syscall.GUID
	Get_TimeWithDst() syscall.GUID
	Get_TimeZone() syscall.GUID
	Get_TxPowerLevel() syscall.GUID
	Get_UnreadAlertStatus() syscall.GUID
}

type IGattCharacteristicUuidsStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_AlertCategoryId                               uintptr
	Get_AlertCategoryIdBitMask                        uintptr
	Get_AlertLevel                                    uintptr
	Get_AlertNotificationControlPoint                 uintptr
	Get_AlertStatus                                   uintptr
	Get_GapAppearance                                 uintptr
	Get_BootKeyboardInputReport                       uintptr
	Get_BootKeyboardOutputReport                      uintptr
	Get_BootMouseInputReport                          uintptr
	Get_CurrentTime                                   uintptr
	Get_CyclingPowerControlPoint                      uintptr
	Get_CyclingPowerFeature                           uintptr
	Get_CyclingPowerMeasurement                       uintptr
	Get_CyclingPowerVector                            uintptr
	Get_DateTime                                      uintptr
	Get_DayDateTime                                   uintptr
	Get_DayOfWeek                                     uintptr
	Get_GapDeviceName                                 uintptr
	Get_DstOffset                                     uintptr
	Get_ExactTime256                                  uintptr
	Get_FirmwareRevisionString                        uintptr
	Get_HardwareRevisionString                        uintptr
	Get_HidControlPoint                               uintptr
	Get_HidInformation                                uintptr
	Get_Ieee1107320601RegulatoryCertificationDataList uintptr
	Get_LnControlPoint                                uintptr
	Get_LnFeature                                     uintptr
	Get_LocalTimeInformation                          uintptr
	Get_LocationAndSpeed                              uintptr
	Get_ManufacturerNameString                        uintptr
	Get_ModelNumberString                             uintptr
	Get_Navigation                                    uintptr
	Get_NewAlert                                      uintptr
	Get_GapPeripheralPreferredConnectionParameters    uintptr
	Get_GapPeripheralPrivacyFlag                      uintptr
	Get_PnpId                                         uintptr
	Get_PositionQuality                               uintptr
	Get_ProtocolMode                                  uintptr
	Get_GapReconnectionAddress                        uintptr
	Get_ReferenceTimeInformation                      uintptr
	Get_Report                                        uintptr
	Get_ReportMap                                     uintptr
	Get_RingerControlPoint                            uintptr
	Get_RingerSetting                                 uintptr
	Get_ScanIntervalWindow                            uintptr
	Get_ScanRefresh                                   uintptr
	Get_SerialNumberString                            uintptr
	Get_GattServiceChanged                            uintptr
	Get_SoftwareRevisionString                        uintptr
	Get_SupportedNewAlertCategory                     uintptr
	Get_SupportUnreadAlertCategory                    uintptr
	Get_SystemId                                      uintptr
	Get_TimeAccuracy                                  uintptr
	Get_TimeSource                                    uintptr
	Get_TimeUpdateControlPoint                        uintptr
	Get_TimeUpdateState                               uintptr
	Get_TimeWithDst                                   uintptr
	Get_TimeZone                                      uintptr
	Get_TxPowerLevel                                  uintptr
	Get_UnreadAlertStatus                             uintptr
}

type IGattCharacteristicUuidsStatics2 struct {
	win32.IInspectable
}

func (this *IGattCharacteristicUuidsStatics2) Vtbl() *IGattCharacteristicUuidsStatics2Vtbl {
	return (*IGattCharacteristicUuidsStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattCharacteristicUuidsStatics2) Get_AlertCategoryId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlertCategoryId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_AlertCategoryIdBitMask() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlertCategoryIdBitMask, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_AlertLevel() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlertLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_AlertNotificationControlPoint() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlertNotificationControlPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_AlertStatus() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlertStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_GapAppearance() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GapAppearance, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_BootKeyboardInputReport() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BootKeyboardInputReport, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_BootKeyboardOutputReport() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BootKeyboardOutputReport, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_BootMouseInputReport() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BootMouseInputReport, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_CurrentTime() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_CyclingPowerControlPoint() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CyclingPowerControlPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_CyclingPowerFeature() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CyclingPowerFeature, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_CyclingPowerMeasurement() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CyclingPowerMeasurement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_CyclingPowerVector() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CyclingPowerVector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_DateTime() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DateTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_DayDateTime() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DayDateTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_DayOfWeek() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DayOfWeek, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_GapDeviceName() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GapDeviceName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_DstOffset() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DstOffset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_ExactTime256() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExactTime256, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_FirmwareRevisionString() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FirmwareRevisionString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_HardwareRevisionString() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HardwareRevisionString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_HidControlPoint() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HidControlPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_HidInformation() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HidInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_Ieee1107320601RegulatoryCertificationDataList() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ieee1107320601RegulatoryCertificationDataList, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_LnControlPoint() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LnControlPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_LnFeature() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LnFeature, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_LocalTimeInformation() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalTimeInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_LocationAndSpeed() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocationAndSpeed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_ManufacturerNameString() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ManufacturerNameString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_ModelNumberString() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ModelNumberString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_Navigation() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Navigation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_NewAlert() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NewAlert, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_GapPeripheralPreferredConnectionParameters() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GapPeripheralPreferredConnectionParameters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_GapPeripheralPrivacyFlag() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GapPeripheralPrivacyFlag, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_PnpId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PnpId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_PositionQuality() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PositionQuality, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_ProtocolMode() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtocolMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_GapReconnectionAddress() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GapReconnectionAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_ReferenceTimeInformation() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReferenceTimeInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_Report() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Report, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_ReportMap() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportMap, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_RingerControlPoint() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RingerControlPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_RingerSetting() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RingerSetting, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_ScanIntervalWindow() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScanIntervalWindow, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_ScanRefresh() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScanRefresh, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_SerialNumberString() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SerialNumberString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_GattServiceChanged() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GattServiceChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_SoftwareRevisionString() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SoftwareRevisionString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_SupportedNewAlertCategory() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedNewAlertCategory, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_SupportUnreadAlertCategory() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportUnreadAlertCategory, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_SystemId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SystemId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_TimeAccuracy() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimeAccuracy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_TimeSource() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimeSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_TimeUpdateControlPoint() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimeUpdateControlPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_TimeUpdateState() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimeUpdateState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_TimeWithDst() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimeWithDst, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_TimeZone() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimeZone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_TxPowerLevel() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TxPowerLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicUuidsStatics2) Get_UnreadAlertStatus() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UnreadAlertStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1194945C-B257-4F3E-9DB7-F68BC9A9AEF2
var IID_IGattCharacteristicsResult = syscall.GUID{0x1194945C, 0xB257, 0x4F3E,
	[8]byte{0x9D, 0xB7, 0xF6, 0x8B, 0xC9, 0xA9, 0xAE, 0xF2}}

type IGattCharacteristicsResultInterface interface {
	win32.IInspectableInterface
	Get_Status() GattCommunicationStatus
	Get_ProtocolError() *IReference[byte]
	Get_Characteristics() *IVectorView[*IGattCharacteristic]
}

type IGattCharacteristicsResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status          uintptr
	Get_ProtocolError   uintptr
	Get_Characteristics uintptr
}

type IGattCharacteristicsResult struct {
	win32.IInspectable
}

func (this *IGattCharacteristicsResult) Vtbl() *IGattCharacteristicsResultVtbl {
	return (*IGattCharacteristicsResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattCharacteristicsResult) Get_Status() GattCommunicationStatus {
	var _result GattCommunicationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattCharacteristicsResult) Get_ProtocolError() *IReference[byte] {
	var _result *IReference[byte]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtocolError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattCharacteristicsResult) Get_Characteristics() *IVectorView[*IGattCharacteristic] {
	var _result *IVectorView[*IGattCharacteristic]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Characteristics, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 506D5599-0112-419A-8E3B-AE21AFABD2C2
var IID_IGattClientNotificationResult = syscall.GUID{0x506D5599, 0x0112, 0x419A,
	[8]byte{0x8E, 0x3B, 0xAE, 0x21, 0xAF, 0xAB, 0xD2, 0xC2}}

type IGattClientNotificationResultInterface interface {
	win32.IInspectableInterface
	Get_SubscribedClient() *IGattSubscribedClient
	Get_Status() GattCommunicationStatus
	Get_ProtocolError() *IReference[byte]
}

type IGattClientNotificationResultVtbl struct {
	win32.IInspectableVtbl
	Get_SubscribedClient uintptr
	Get_Status           uintptr
	Get_ProtocolError    uintptr
}

type IGattClientNotificationResult struct {
	win32.IInspectable
}

func (this *IGattClientNotificationResult) Vtbl() *IGattClientNotificationResultVtbl {
	return (*IGattClientNotificationResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattClientNotificationResult) Get_SubscribedClient() *IGattSubscribedClient {
	var _result *IGattSubscribedClient
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SubscribedClient, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattClientNotificationResult) Get_Status() GattCommunicationStatus {
	var _result GattCommunicationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattClientNotificationResult) Get_ProtocolError() *IReference[byte] {
	var _result *IReference[byte]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtocolError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8FAEC497-45E0-497E-9582-29A1FE281AD5
var IID_IGattClientNotificationResult2 = syscall.GUID{0x8FAEC497, 0x45E0, 0x497E,
	[8]byte{0x95, 0x82, 0x29, 0xA1, 0xFE, 0x28, 0x1A, 0xD5}}

type IGattClientNotificationResult2Interface interface {
	win32.IInspectableInterface
	Get_BytesSent() uint16
}

type IGattClientNotificationResult2Vtbl struct {
	win32.IInspectableVtbl
	Get_BytesSent uintptr
}

type IGattClientNotificationResult2 struct {
	win32.IInspectable
}

func (this *IGattClientNotificationResult2) Vtbl() *IGattClientNotificationResult2Vtbl {
	return (*IGattClientNotificationResult2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattClientNotificationResult2) Get_BytesSent() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BytesSent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 92055F2B-8084-4344-B4C2-284DE19A8506
var IID_IGattDescriptor = syscall.GUID{0x92055F2B, 0x8084, 0x4344,
	[8]byte{0xB4, 0xC2, 0x28, 0x4D, 0xE1, 0x9A, 0x85, 0x06}}

type IGattDescriptorInterface interface {
	win32.IInspectableInterface
	Get_ProtectionLevel() GattProtectionLevel
	Put_ProtectionLevel(value GattProtectionLevel)
	Get_Uuid() syscall.GUID
	Get_AttributeHandle() uint16
	ReadValueAsync() *IAsyncOperation[*IGattReadResult]
	ReadValueWithCacheModeAsync(cacheMode BluetoothCacheMode) *IAsyncOperation[*IGattReadResult]
	WriteValueAsync(value *IBuffer) *IAsyncOperation[GattCommunicationStatus]
}

type IGattDescriptorVtbl struct {
	win32.IInspectableVtbl
	Get_ProtectionLevel         uintptr
	Put_ProtectionLevel         uintptr
	Get_Uuid                    uintptr
	Get_AttributeHandle         uintptr
	ReadValueAsync              uintptr
	ReadValueWithCacheModeAsync uintptr
	WriteValueAsync             uintptr
}

type IGattDescriptor struct {
	win32.IInspectable
}

func (this *IGattDescriptor) Vtbl() *IGattDescriptorVtbl {
	return (*IGattDescriptorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattDescriptor) Get_ProtectionLevel() GattProtectionLevel {
	var _result GattProtectionLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattDescriptor) Put_ProtectionLevel(value GattProtectionLevel) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGattDescriptor) Get_Uuid() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Uuid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattDescriptor) Get_AttributeHandle() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AttributeHandle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattDescriptor) ReadValueAsync() *IAsyncOperation[*IGattReadResult] {
	var _result *IAsyncOperation[*IGattReadResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadValueAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattDescriptor) ReadValueWithCacheModeAsync(cacheMode BluetoothCacheMode) *IAsyncOperation[*IGattReadResult] {
	var _result *IAsyncOperation[*IGattReadResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadValueWithCacheModeAsync, uintptr(unsafe.Pointer(this)), uintptr(cacheMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattDescriptor) WriteValueAsync(value *IBuffer) *IAsyncOperation[GattCommunicationStatus] {
	var _result *IAsyncOperation[GattCommunicationStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteValueAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8F563D39-D630-406C-BA11-10CDD16B0E5E
var IID_IGattDescriptor2 = syscall.GUID{0x8F563D39, 0xD630, 0x406C,
	[8]byte{0xBA, 0x11, 0x10, 0xCD, 0xD1, 0x6B, 0x0E, 0x5E}}

type IGattDescriptor2Interface interface {
	win32.IInspectableInterface
	WriteValueWithResultAsync(value *IBuffer) *IAsyncOperation[*IGattWriteResult]
}

type IGattDescriptor2Vtbl struct {
	win32.IInspectableVtbl
	WriteValueWithResultAsync uintptr
}

type IGattDescriptor2 struct {
	win32.IInspectable
}

func (this *IGattDescriptor2) Vtbl() *IGattDescriptor2Vtbl {
	return (*IGattDescriptor2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattDescriptor2) WriteValueWithResultAsync(value *IBuffer) *IAsyncOperation[*IGattWriteResult] {
	var _result *IAsyncOperation[*IGattWriteResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteValueWithResultAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 92055F2D-8084-4344-B4C2-284DE19A8506
var IID_IGattDescriptorStatics = syscall.GUID{0x92055F2D, 0x8084, 0x4344,
	[8]byte{0xB4, 0xC2, 0x28, 0x4D, 0xE1, 0x9A, 0x85, 0x06}}

type IGattDescriptorStaticsInterface interface {
	win32.IInspectableInterface
	ConvertShortIdToUuid(shortId uint16) syscall.GUID
}

type IGattDescriptorStaticsVtbl struct {
	win32.IInspectableVtbl
	ConvertShortIdToUuid uintptr
}

type IGattDescriptorStatics struct {
	win32.IInspectable
}

func (this *IGattDescriptorStatics) Vtbl() *IGattDescriptorStaticsVtbl {
	return (*IGattDescriptorStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattDescriptorStatics) ConvertShortIdToUuid(shortId uint16) syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConvertShortIdToUuid, uintptr(unsafe.Pointer(this)), uintptr(shortId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A6F862CE-9CFC-42F1-9185-FF37B75181D3
var IID_IGattDescriptorUuidsStatics = syscall.GUID{0xA6F862CE, 0x9CFC, 0x42F1,
	[8]byte{0x91, 0x85, 0xFF, 0x37, 0xB7, 0x51, 0x81, 0xD3}}

type IGattDescriptorUuidsStaticsInterface interface {
	win32.IInspectableInterface
	Get_CharacteristicAggregateFormat() syscall.GUID
	Get_CharacteristicExtendedProperties() syscall.GUID
	Get_CharacteristicPresentationFormat() syscall.GUID
	Get_CharacteristicUserDescription() syscall.GUID
	Get_ClientCharacteristicConfiguration() syscall.GUID
	Get_ServerCharacteristicConfiguration() syscall.GUID
}

type IGattDescriptorUuidsStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_CharacteristicAggregateFormat     uintptr
	Get_CharacteristicExtendedProperties  uintptr
	Get_CharacteristicPresentationFormat  uintptr
	Get_CharacteristicUserDescription     uintptr
	Get_ClientCharacteristicConfiguration uintptr
	Get_ServerCharacteristicConfiguration uintptr
}

type IGattDescriptorUuidsStatics struct {
	win32.IInspectable
}

func (this *IGattDescriptorUuidsStatics) Vtbl() *IGattDescriptorUuidsStaticsVtbl {
	return (*IGattDescriptorUuidsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattDescriptorUuidsStatics) Get_CharacteristicAggregateFormat() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CharacteristicAggregateFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattDescriptorUuidsStatics) Get_CharacteristicExtendedProperties() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CharacteristicExtendedProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattDescriptorUuidsStatics) Get_CharacteristicPresentationFormat() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CharacteristicPresentationFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattDescriptorUuidsStatics) Get_CharacteristicUserDescription() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CharacteristicUserDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattDescriptorUuidsStatics) Get_ClientCharacteristicConfiguration() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ClientCharacteristicConfiguration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattDescriptorUuidsStatics) Get_ServerCharacteristicConfiguration() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerCharacteristicConfiguration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 9BC091F3-95E7-4489-8D25-FF81955A57B9
var IID_IGattDescriptorsResult = syscall.GUID{0x9BC091F3, 0x95E7, 0x4489,
	[8]byte{0x8D, 0x25, 0xFF, 0x81, 0x95, 0x5A, 0x57, 0xB9}}

type IGattDescriptorsResultInterface interface {
	win32.IInspectableInterface
	Get_Status() GattCommunicationStatus
	Get_ProtocolError() *IReference[byte]
	Get_Descriptors() *IVectorView[*IGattDescriptor]
}

type IGattDescriptorsResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status        uintptr
	Get_ProtocolError uintptr
	Get_Descriptors   uintptr
}

type IGattDescriptorsResult struct {
	win32.IInspectable
}

func (this *IGattDescriptorsResult) Vtbl() *IGattDescriptorsResultVtbl {
	return (*IGattDescriptorsResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattDescriptorsResult) Get_Status() GattCommunicationStatus {
	var _result GattCommunicationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattDescriptorsResult) Get_ProtocolError() *IReference[byte] {
	var _result *IReference[byte]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtocolError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattDescriptorsResult) Get_Descriptors() *IVectorView[*IGattDescriptor] {
	var _result *IVectorView[*IGattDescriptor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Descriptors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AC7B7C05-B33C-47CF-990F-6B8F5577DF71
var IID_IGattDeviceService = syscall.GUID{0xAC7B7C05, 0xB33C, 0x47CF,
	[8]byte{0x99, 0x0F, 0x6B, 0x8F, 0x55, 0x77, 0xDF, 0x71}}

type IGattDeviceServiceInterface interface {
	win32.IInspectableInterface
	GetCharacteristics(characteristicUuid syscall.GUID) *IVectorView[*IGattCharacteristic]
	GetIncludedServices(serviceUuid syscall.GUID) *IVectorView[*IGattDeviceService]
	Get_DeviceId() string
	Get_Uuid() syscall.GUID
	Get_AttributeHandle() uint16
}

type IGattDeviceServiceVtbl struct {
	win32.IInspectableVtbl
	GetCharacteristics  uintptr
	GetIncludedServices uintptr
	Get_DeviceId        uintptr
	Get_Uuid            uintptr
	Get_AttributeHandle uintptr
}

type IGattDeviceService struct {
	win32.IInspectable
}

func (this *IGattDeviceService) Vtbl() *IGattDeviceServiceVtbl {
	return (*IGattDeviceServiceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattDeviceService) GetCharacteristics(characteristicUuid syscall.GUID) *IVectorView[*IGattCharacteristic] {
	var _result *IVectorView[*IGattCharacteristic]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCharacteristics, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&characteristicUuid)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattDeviceService) GetIncludedServices(serviceUuid syscall.GUID) *IVectorView[*IGattDeviceService] {
	var _result *IVectorView[*IGattDeviceService]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetIncludedServices, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&serviceUuid)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattDeviceService) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGattDeviceService) Get_Uuid() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Uuid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattDeviceService) Get_AttributeHandle() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AttributeHandle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// FC54520B-0B0D-4708-BAE0-9FFD9489BC59
var IID_IGattDeviceService2 = syscall.GUID{0xFC54520B, 0x0B0D, 0x4708,
	[8]byte{0xBA, 0xE0, 0x9F, 0xFD, 0x94, 0x89, 0xBC, 0x59}}

type IGattDeviceService2Interface interface {
	win32.IInspectableInterface
	Get_Device() *IBluetoothLEDevice
	Get_ParentServices() *IVectorView[*IGattDeviceService]
	GetAllCharacteristics() *IVectorView[*IGattCharacteristic]
	GetAllIncludedServices() *IVectorView[*IGattDeviceService]
}

type IGattDeviceService2Vtbl struct {
	win32.IInspectableVtbl
	Get_Device             uintptr
	Get_ParentServices     uintptr
	GetAllCharacteristics  uintptr
	GetAllIncludedServices uintptr
}

type IGattDeviceService2 struct {
	win32.IInspectable
}

func (this *IGattDeviceService2) Vtbl() *IGattDeviceService2Vtbl {
	return (*IGattDeviceService2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattDeviceService2) Get_Device() *IBluetoothLEDevice {
	var _result *IBluetoothLEDevice
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Device, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattDeviceService2) Get_ParentServices() *IVectorView[*IGattDeviceService] {
	var _result *IVectorView[*IGattDeviceService]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ParentServices, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattDeviceService2) GetAllCharacteristics() *IVectorView[*IGattCharacteristic] {
	var _result *IVectorView[*IGattCharacteristic]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAllCharacteristics, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattDeviceService2) GetAllIncludedServices() *IVectorView[*IGattDeviceService] {
	var _result *IVectorView[*IGattDeviceService]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAllIncludedServices, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B293A950-0C53-437C-A9B3-5C3210C6E569
var IID_IGattDeviceService3 = syscall.GUID{0xB293A950, 0x0C53, 0x437C,
	[8]byte{0xA9, 0xB3, 0x5C, 0x32, 0x10, 0xC6, 0xE5, 0x69}}

type IGattDeviceService3Interface interface {
	win32.IInspectableInterface
	Get_DeviceAccessInformation() *IDeviceAccessInformation
	Get_Session() *IGattSession
	Get_SharingMode() GattSharingMode
	RequestAccessAsync() *IAsyncOperation[DeviceAccessStatus]
	OpenAsync(sharingMode GattSharingMode) *IAsyncOperation[GattOpenStatus]
	GetCharacteristicsAsync() *IAsyncOperation[*IGattCharacteristicsResult]
	GetCharacteristicsWithCacheModeAsync(cacheMode BluetoothCacheMode) *IAsyncOperation[*IGattCharacteristicsResult]
	GetCharacteristicsForUuidAsync(characteristicUuid syscall.GUID) *IAsyncOperation[*IGattCharacteristicsResult]
	GetCharacteristicsForUuidWithCacheModeAsync(characteristicUuid syscall.GUID, cacheMode BluetoothCacheMode) *IAsyncOperation[*IGattCharacteristicsResult]
	GetIncludedServicesAsync() *IAsyncOperation[*IGattDeviceServicesResult]
	GetIncludedServicesWithCacheModeAsync(cacheMode BluetoothCacheMode) *IAsyncOperation[*IGattDeviceServicesResult]
	GetIncludedServicesForUuidAsync(serviceUuid syscall.GUID) *IAsyncOperation[*IGattDeviceServicesResult]
	GetIncludedServicesForUuidWithCacheModeAsync(serviceUuid syscall.GUID, cacheMode BluetoothCacheMode) *IAsyncOperation[*IGattDeviceServicesResult]
}

type IGattDeviceService3Vtbl struct {
	win32.IInspectableVtbl
	Get_DeviceAccessInformation                  uintptr
	Get_Session                                  uintptr
	Get_SharingMode                              uintptr
	RequestAccessAsync                           uintptr
	OpenAsync                                    uintptr
	GetCharacteristicsAsync                      uintptr
	GetCharacteristicsWithCacheModeAsync         uintptr
	GetCharacteristicsForUuidAsync               uintptr
	GetCharacteristicsForUuidWithCacheModeAsync  uintptr
	GetIncludedServicesAsync                     uintptr
	GetIncludedServicesWithCacheModeAsync        uintptr
	GetIncludedServicesForUuidAsync              uintptr
	GetIncludedServicesForUuidWithCacheModeAsync uintptr
}

type IGattDeviceService3 struct {
	win32.IInspectable
}

func (this *IGattDeviceService3) Vtbl() *IGattDeviceService3Vtbl {
	return (*IGattDeviceService3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattDeviceService3) Get_DeviceAccessInformation() *IDeviceAccessInformation {
	var _result *IDeviceAccessInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceAccessInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattDeviceService3) Get_Session() *IGattSession {
	var _result *IGattSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Session, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattDeviceService3) Get_SharingMode() GattSharingMode {
	var _result GattSharingMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SharingMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattDeviceService3) RequestAccessAsync() *IAsyncOperation[DeviceAccessStatus] {
	var _result *IAsyncOperation[DeviceAccessStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattDeviceService3) OpenAsync(sharingMode GattSharingMode) *IAsyncOperation[GattOpenStatus] {
	var _result *IAsyncOperation[GattOpenStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenAsync, uintptr(unsafe.Pointer(this)), uintptr(sharingMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattDeviceService3) GetCharacteristicsAsync() *IAsyncOperation[*IGattCharacteristicsResult] {
	var _result *IAsyncOperation[*IGattCharacteristicsResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCharacteristicsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattDeviceService3) GetCharacteristicsWithCacheModeAsync(cacheMode BluetoothCacheMode) *IAsyncOperation[*IGattCharacteristicsResult] {
	var _result *IAsyncOperation[*IGattCharacteristicsResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCharacteristicsWithCacheModeAsync, uintptr(unsafe.Pointer(this)), uintptr(cacheMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattDeviceService3) GetCharacteristicsForUuidAsync(characteristicUuid syscall.GUID) *IAsyncOperation[*IGattCharacteristicsResult] {
	var _result *IAsyncOperation[*IGattCharacteristicsResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCharacteristicsForUuidAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&characteristicUuid)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattDeviceService3) GetCharacteristicsForUuidWithCacheModeAsync(characteristicUuid syscall.GUID, cacheMode BluetoothCacheMode) *IAsyncOperation[*IGattCharacteristicsResult] {
	var _result *IAsyncOperation[*IGattCharacteristicsResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCharacteristicsForUuidWithCacheModeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&characteristicUuid)), uintptr(cacheMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattDeviceService3) GetIncludedServicesAsync() *IAsyncOperation[*IGattDeviceServicesResult] {
	var _result *IAsyncOperation[*IGattDeviceServicesResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetIncludedServicesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattDeviceService3) GetIncludedServicesWithCacheModeAsync(cacheMode BluetoothCacheMode) *IAsyncOperation[*IGattDeviceServicesResult] {
	var _result *IAsyncOperation[*IGattDeviceServicesResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetIncludedServicesWithCacheModeAsync, uintptr(unsafe.Pointer(this)), uintptr(cacheMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattDeviceService3) GetIncludedServicesForUuidAsync(serviceUuid syscall.GUID) *IAsyncOperation[*IGattDeviceServicesResult] {
	var _result *IAsyncOperation[*IGattDeviceServicesResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetIncludedServicesForUuidAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&serviceUuid)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattDeviceService3) GetIncludedServicesForUuidWithCacheModeAsync(serviceUuid syscall.GUID, cacheMode BluetoothCacheMode) *IAsyncOperation[*IGattDeviceServicesResult] {
	var _result *IAsyncOperation[*IGattDeviceServicesResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetIncludedServicesForUuidWithCacheModeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&serviceUuid)), uintptr(cacheMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 196D0022-FAAD-45DC-AE5B-2AC3184E84DB
var IID_IGattDeviceServiceStatics = syscall.GUID{0x196D0022, 0xFAAD, 0x45DC,
	[8]byte{0xAE, 0x5B, 0x2A, 0xC3, 0x18, 0x4E, 0x84, 0xDB}}

type IGattDeviceServiceStaticsInterface interface {
	win32.IInspectableInterface
	FromIdAsync(deviceId string) *IAsyncOperation[*IGattDeviceService]
	GetDeviceSelectorFromUuid(serviceUuid syscall.GUID) string
	GetDeviceSelectorFromShortId(serviceShortId uint16) string
	ConvertShortIdToUuid(shortId uint16) syscall.GUID
}

type IGattDeviceServiceStaticsVtbl struct {
	win32.IInspectableVtbl
	FromIdAsync                  uintptr
	GetDeviceSelectorFromUuid    uintptr
	GetDeviceSelectorFromShortId uintptr
	ConvertShortIdToUuid         uintptr
}

type IGattDeviceServiceStatics struct {
	win32.IInspectable
}

func (this *IGattDeviceServiceStatics) Vtbl() *IGattDeviceServiceStaticsVtbl {
	return (*IGattDeviceServiceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattDeviceServiceStatics) FromIdAsync(deviceId string) *IAsyncOperation[*IGattDeviceService] {
	var _result *IAsyncOperation[*IGattDeviceService]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattDeviceServiceStatics) GetDeviceSelectorFromUuid(serviceUuid syscall.GUID) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorFromUuid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&serviceUuid)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGattDeviceServiceStatics) GetDeviceSelectorFromShortId(serviceShortId uint16) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorFromShortId, uintptr(unsafe.Pointer(this)), uintptr(serviceShortId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGattDeviceServiceStatics) ConvertShortIdToUuid(shortId uint16) syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConvertShortIdToUuid, uintptr(unsafe.Pointer(this)), uintptr(shortId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 0604186E-24A6-4B0D-A2F2-30CC01545D25
var IID_IGattDeviceServiceStatics2 = syscall.GUID{0x0604186E, 0x24A6, 0x4B0D,
	[8]byte{0xA2, 0xF2, 0x30, 0xCC, 0x01, 0x54, 0x5D, 0x25}}

type IGattDeviceServiceStatics2Interface interface {
	win32.IInspectableInterface
	FromIdWithSharingModeAsync(deviceId string, sharingMode GattSharingMode) *IAsyncOperation[*IGattDeviceService]
	GetDeviceSelectorForBluetoothDeviceId(bluetoothDeviceId *IBluetoothDeviceId) string
	GetDeviceSelectorForBluetoothDeviceIdWithCacheMode(bluetoothDeviceId *IBluetoothDeviceId, cacheMode BluetoothCacheMode) string
	GetDeviceSelectorForBluetoothDeviceIdAndUuid(bluetoothDeviceId *IBluetoothDeviceId, serviceUuid syscall.GUID) string
	GetDeviceSelectorForBluetoothDeviceIdAndUuidWithCacheMode(bluetoothDeviceId *IBluetoothDeviceId, serviceUuid syscall.GUID, cacheMode BluetoothCacheMode) string
}

type IGattDeviceServiceStatics2Vtbl struct {
	win32.IInspectableVtbl
	FromIdWithSharingModeAsync                                uintptr
	GetDeviceSelectorForBluetoothDeviceId                     uintptr
	GetDeviceSelectorForBluetoothDeviceIdWithCacheMode        uintptr
	GetDeviceSelectorForBluetoothDeviceIdAndUuid              uintptr
	GetDeviceSelectorForBluetoothDeviceIdAndUuidWithCacheMode uintptr
}

type IGattDeviceServiceStatics2 struct {
	win32.IInspectable
}

func (this *IGattDeviceServiceStatics2) Vtbl() *IGattDeviceServiceStatics2Vtbl {
	return (*IGattDeviceServiceStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattDeviceServiceStatics2) FromIdWithSharingModeAsync(deviceId string, sharingMode GattSharingMode) *IAsyncOperation[*IGattDeviceService] {
	var _result *IAsyncOperation[*IGattDeviceService]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdWithSharingModeAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(sharingMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattDeviceServiceStatics2) GetDeviceSelectorForBluetoothDeviceId(bluetoothDeviceId *IBluetoothDeviceId) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorForBluetoothDeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bluetoothDeviceId)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGattDeviceServiceStatics2) GetDeviceSelectorForBluetoothDeviceIdWithCacheMode(bluetoothDeviceId *IBluetoothDeviceId, cacheMode BluetoothCacheMode) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorForBluetoothDeviceIdWithCacheMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bluetoothDeviceId)), uintptr(cacheMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGattDeviceServiceStatics2) GetDeviceSelectorForBluetoothDeviceIdAndUuid(bluetoothDeviceId *IBluetoothDeviceId, serviceUuid syscall.GUID) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorForBluetoothDeviceIdAndUuid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bluetoothDeviceId)), uintptr(unsafe.Pointer(&serviceUuid)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGattDeviceServiceStatics2) GetDeviceSelectorForBluetoothDeviceIdAndUuidWithCacheMode(bluetoothDeviceId *IBluetoothDeviceId, serviceUuid syscall.GUID, cacheMode BluetoothCacheMode) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorForBluetoothDeviceIdAndUuidWithCacheMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bluetoothDeviceId)), uintptr(unsafe.Pointer(&serviceUuid)), uintptr(cacheMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 171DD3EE-016D-419D-838A-576CF475A3D8
var IID_IGattDeviceServicesResult = syscall.GUID{0x171DD3EE, 0x016D, 0x419D,
	[8]byte{0x83, 0x8A, 0x57, 0x6C, 0xF4, 0x75, 0xA3, 0xD8}}

type IGattDeviceServicesResultInterface interface {
	win32.IInspectableInterface
	Get_Status() GattCommunicationStatus
	Get_ProtocolError() *IReference[byte]
	Get_Services() *IVectorView[*IGattDeviceService]
}

type IGattDeviceServicesResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status        uintptr
	Get_ProtocolError uintptr
	Get_Services      uintptr
}

type IGattDeviceServicesResult struct {
	win32.IInspectable
}

func (this *IGattDeviceServicesResult) Vtbl() *IGattDeviceServicesResultVtbl {
	return (*IGattDeviceServicesResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattDeviceServicesResult) Get_Status() GattCommunicationStatus {
	var _result GattCommunicationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattDeviceServicesResult) Get_ProtocolError() *IReference[byte] {
	var _result *IReference[byte]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtocolError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattDeviceServicesResult) Get_Services() *IVectorView[*IGattDeviceService] {
	var _result *IVectorView[*IGattDeviceService]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Services, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AEDE376D-5412-4D74-92A8-8DEB8526829C
var IID_IGattLocalCharacteristic = syscall.GUID{0xAEDE376D, 0x5412, 0x4D74,
	[8]byte{0x92, 0xA8, 0x8D, 0xEB, 0x85, 0x26, 0x82, 0x9C}}

type IGattLocalCharacteristicInterface interface {
	win32.IInspectableInterface
	Get_Uuid() syscall.GUID
	Get_StaticValue() *IBuffer
	Get_CharacteristicProperties() GattCharacteristicProperties
	Get_ReadProtectionLevel() GattProtectionLevel
	Get_WriteProtectionLevel() GattProtectionLevel
	CreateDescriptorAsync(descriptorUuid syscall.GUID, parameters *IGattLocalDescriptorParameters) *IAsyncOperation[*IGattLocalDescriptorResult]
	Get_Descriptors() *IVectorView[*IGattLocalDescriptor]
	Get_UserDescription() string
	Get_PresentationFormats() *IVectorView[*IGattPresentationFormat]
	Get_SubscribedClients() *IVectorView[*IGattSubscribedClient]
	Add_SubscribedClientsChanged(handler TypedEventHandler[*IGattLocalCharacteristic, interface{}]) EventRegistrationToken
	Remove_SubscribedClientsChanged(token EventRegistrationToken)
	Add_ReadRequested(handler TypedEventHandler[*IGattLocalCharacteristic, *IGattReadRequestedEventArgs]) EventRegistrationToken
	Remove_ReadRequested(token EventRegistrationToken)
	Add_WriteRequested(handler TypedEventHandler[*IGattLocalCharacteristic, *IGattWriteRequestedEventArgs]) EventRegistrationToken
	Remove_WriteRequested(token EventRegistrationToken)
	NotifyValueAsync(value *IBuffer) *IAsyncOperation[*IVectorView[*IGattClientNotificationResult]]
	NotifyValueForSubscribedClientAsync(value *IBuffer, subscribedClient *IGattSubscribedClient) *IAsyncOperation[*IGattClientNotificationResult]
}

type IGattLocalCharacteristicVtbl struct {
	win32.IInspectableVtbl
	Get_Uuid                            uintptr
	Get_StaticValue                     uintptr
	Get_CharacteristicProperties        uintptr
	Get_ReadProtectionLevel             uintptr
	Get_WriteProtectionLevel            uintptr
	CreateDescriptorAsync               uintptr
	Get_Descriptors                     uintptr
	Get_UserDescription                 uintptr
	Get_PresentationFormats             uintptr
	Get_SubscribedClients               uintptr
	Add_SubscribedClientsChanged        uintptr
	Remove_SubscribedClientsChanged     uintptr
	Add_ReadRequested                   uintptr
	Remove_ReadRequested                uintptr
	Add_WriteRequested                  uintptr
	Remove_WriteRequested               uintptr
	NotifyValueAsync                    uintptr
	NotifyValueForSubscribedClientAsync uintptr
}

type IGattLocalCharacteristic struct {
	win32.IInspectable
}

func (this *IGattLocalCharacteristic) Vtbl() *IGattLocalCharacteristicVtbl {
	return (*IGattLocalCharacteristicVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattLocalCharacteristic) Get_Uuid() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Uuid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattLocalCharacteristic) Get_StaticValue() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StaticValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattLocalCharacteristic) Get_CharacteristicProperties() GattCharacteristicProperties {
	var _result GattCharacteristicProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CharacteristicProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattLocalCharacteristic) Get_ReadProtectionLevel() GattProtectionLevel {
	var _result GattProtectionLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReadProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattLocalCharacteristic) Get_WriteProtectionLevel() GattProtectionLevel {
	var _result GattProtectionLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WriteProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattLocalCharacteristic) CreateDescriptorAsync(descriptorUuid syscall.GUID, parameters *IGattLocalDescriptorParameters) *IAsyncOperation[*IGattLocalDescriptorResult] {
	var _result *IAsyncOperation[*IGattLocalDescriptorResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDescriptorAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&descriptorUuid)), uintptr(unsafe.Pointer(parameters)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattLocalCharacteristic) Get_Descriptors() *IVectorView[*IGattLocalDescriptor] {
	var _result *IVectorView[*IGattLocalDescriptor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Descriptors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattLocalCharacteristic) Get_UserDescription() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGattLocalCharacteristic) Get_PresentationFormats() *IVectorView[*IGattPresentationFormat] {
	var _result *IVectorView[*IGattPresentationFormat]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PresentationFormats, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattLocalCharacteristic) Get_SubscribedClients() *IVectorView[*IGattSubscribedClient] {
	var _result *IVectorView[*IGattSubscribedClient]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SubscribedClients, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattLocalCharacteristic) Add_SubscribedClientsChanged(handler TypedEventHandler[*IGattLocalCharacteristic, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SubscribedClientsChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattLocalCharacteristic) Remove_SubscribedClientsChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SubscribedClientsChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGattLocalCharacteristic) Add_ReadRequested(handler TypedEventHandler[*IGattLocalCharacteristic, *IGattReadRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ReadRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattLocalCharacteristic) Remove_ReadRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ReadRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGattLocalCharacteristic) Add_WriteRequested(handler TypedEventHandler[*IGattLocalCharacteristic, *IGattWriteRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_WriteRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattLocalCharacteristic) Remove_WriteRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_WriteRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGattLocalCharacteristic) NotifyValueAsync(value *IBuffer) *IAsyncOperation[*IVectorView[*IGattClientNotificationResult]] {
	var _result *IAsyncOperation[*IVectorView[*IGattClientNotificationResult]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().NotifyValueAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattLocalCharacteristic) NotifyValueForSubscribedClientAsync(value *IBuffer, subscribedClient *IGattSubscribedClient) *IAsyncOperation[*IGattClientNotificationResult] {
	var _result *IAsyncOperation[*IGattClientNotificationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().NotifyValueForSubscribedClientAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(subscribedClient)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FAF73DB4-4CFF-44C7-8445-040E6EAD0063
var IID_IGattLocalCharacteristicParameters = syscall.GUID{0xFAF73DB4, 0x4CFF, 0x44C7,
	[8]byte{0x84, 0x45, 0x04, 0x0E, 0x6E, 0xAD, 0x00, 0x63}}

type IGattLocalCharacteristicParametersInterface interface {
	win32.IInspectableInterface
	Put_StaticValue(value *IBuffer)
	Get_StaticValue() *IBuffer
	Put_CharacteristicProperties(value GattCharacteristicProperties)
	Get_CharacteristicProperties() GattCharacteristicProperties
	Put_ReadProtectionLevel(value GattProtectionLevel)
	Get_ReadProtectionLevel() GattProtectionLevel
	Put_WriteProtectionLevel(value GattProtectionLevel)
	Get_WriteProtectionLevel() GattProtectionLevel
	Put_UserDescription(value string)
	Get_UserDescription() string
	Get_PresentationFormats() *IVector[*IGattPresentationFormat]
}

type IGattLocalCharacteristicParametersVtbl struct {
	win32.IInspectableVtbl
	Put_StaticValue              uintptr
	Get_StaticValue              uintptr
	Put_CharacteristicProperties uintptr
	Get_CharacteristicProperties uintptr
	Put_ReadProtectionLevel      uintptr
	Get_ReadProtectionLevel      uintptr
	Put_WriteProtectionLevel     uintptr
	Get_WriteProtectionLevel     uintptr
	Put_UserDescription          uintptr
	Get_UserDescription          uintptr
	Get_PresentationFormats      uintptr
}

type IGattLocalCharacteristicParameters struct {
	win32.IInspectable
}

func (this *IGattLocalCharacteristicParameters) Vtbl() *IGattLocalCharacteristicParametersVtbl {
	return (*IGattLocalCharacteristicParametersVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattLocalCharacteristicParameters) Put_StaticValue(value *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StaticValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IGattLocalCharacteristicParameters) Get_StaticValue() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StaticValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattLocalCharacteristicParameters) Put_CharacteristicProperties(value GattCharacteristicProperties) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CharacteristicProperties, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGattLocalCharacteristicParameters) Get_CharacteristicProperties() GattCharacteristicProperties {
	var _result GattCharacteristicProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CharacteristicProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattLocalCharacteristicParameters) Put_ReadProtectionLevel(value GattProtectionLevel) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReadProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGattLocalCharacteristicParameters) Get_ReadProtectionLevel() GattProtectionLevel {
	var _result GattProtectionLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReadProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattLocalCharacteristicParameters) Put_WriteProtectionLevel(value GattProtectionLevel) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_WriteProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGattLocalCharacteristicParameters) Get_WriteProtectionLevel() GattProtectionLevel {
	var _result GattProtectionLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WriteProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattLocalCharacteristicParameters) Put_UserDescription(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_UserDescription, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IGattLocalCharacteristicParameters) Get_UserDescription() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGattLocalCharacteristicParameters) Get_PresentationFormats() *IVector[*IGattPresentationFormat] {
	var _result *IVector[*IGattPresentationFormat]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PresentationFormats, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7975DE9B-0170-4397-9666-92F863F12EE6
var IID_IGattLocalCharacteristicResult = syscall.GUID{0x7975DE9B, 0x0170, 0x4397,
	[8]byte{0x96, 0x66, 0x92, 0xF8, 0x63, 0xF1, 0x2E, 0xE6}}

type IGattLocalCharacteristicResultInterface interface {
	win32.IInspectableInterface
	Get_Characteristic() *IGattLocalCharacteristic
	Get_Error() BluetoothError
}

type IGattLocalCharacteristicResultVtbl struct {
	win32.IInspectableVtbl
	Get_Characteristic uintptr
	Get_Error          uintptr
}

type IGattLocalCharacteristicResult struct {
	win32.IInspectable
}

func (this *IGattLocalCharacteristicResult) Vtbl() *IGattLocalCharacteristicResultVtbl {
	return (*IGattLocalCharacteristicResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattLocalCharacteristicResult) Get_Characteristic() *IGattLocalCharacteristic {
	var _result *IGattLocalCharacteristic
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Characteristic, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattLocalCharacteristicResult) Get_Error() BluetoothError {
	var _result BluetoothError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Error, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F48EBE06-789D-4A4B-8652-BD017B5D2FC6
var IID_IGattLocalDescriptor = syscall.GUID{0xF48EBE06, 0x789D, 0x4A4B,
	[8]byte{0x86, 0x52, 0xBD, 0x01, 0x7B, 0x5D, 0x2F, 0xC6}}

type IGattLocalDescriptorInterface interface {
	win32.IInspectableInterface
	Get_Uuid() syscall.GUID
	Get_StaticValue() *IBuffer
	Get_ReadProtectionLevel() GattProtectionLevel
	Get_WriteProtectionLevel() GattProtectionLevel
	Add_ReadRequested(handler TypedEventHandler[*IGattLocalDescriptor, *IGattReadRequestedEventArgs]) EventRegistrationToken
	Remove_ReadRequested(token EventRegistrationToken)
	Add_WriteRequested(handler TypedEventHandler[*IGattLocalDescriptor, *IGattWriteRequestedEventArgs]) EventRegistrationToken
	Remove_WriteRequested(token EventRegistrationToken)
}

type IGattLocalDescriptorVtbl struct {
	win32.IInspectableVtbl
	Get_Uuid                 uintptr
	Get_StaticValue          uintptr
	Get_ReadProtectionLevel  uintptr
	Get_WriteProtectionLevel uintptr
	Add_ReadRequested        uintptr
	Remove_ReadRequested     uintptr
	Add_WriteRequested       uintptr
	Remove_WriteRequested    uintptr
}

type IGattLocalDescriptor struct {
	win32.IInspectable
}

func (this *IGattLocalDescriptor) Vtbl() *IGattLocalDescriptorVtbl {
	return (*IGattLocalDescriptorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattLocalDescriptor) Get_Uuid() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Uuid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattLocalDescriptor) Get_StaticValue() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StaticValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattLocalDescriptor) Get_ReadProtectionLevel() GattProtectionLevel {
	var _result GattProtectionLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReadProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattLocalDescriptor) Get_WriteProtectionLevel() GattProtectionLevel {
	var _result GattProtectionLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WriteProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattLocalDescriptor) Add_ReadRequested(handler TypedEventHandler[*IGattLocalDescriptor, *IGattReadRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ReadRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattLocalDescriptor) Remove_ReadRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ReadRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGattLocalDescriptor) Add_WriteRequested(handler TypedEventHandler[*IGattLocalDescriptor, *IGattWriteRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_WriteRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattLocalDescriptor) Remove_WriteRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_WriteRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 5FDEDE6A-F3C1-4B66-8C4B-E3D2293B40E9
var IID_IGattLocalDescriptorParameters = syscall.GUID{0x5FDEDE6A, 0xF3C1, 0x4B66,
	[8]byte{0x8C, 0x4B, 0xE3, 0xD2, 0x29, 0x3B, 0x40, 0xE9}}

type IGattLocalDescriptorParametersInterface interface {
	win32.IInspectableInterface
	Put_StaticValue(value *IBuffer)
	Get_StaticValue() *IBuffer
	Put_ReadProtectionLevel(value GattProtectionLevel)
	Get_ReadProtectionLevel() GattProtectionLevel
	Put_WriteProtectionLevel(value GattProtectionLevel)
	Get_WriteProtectionLevel() GattProtectionLevel
}

type IGattLocalDescriptorParametersVtbl struct {
	win32.IInspectableVtbl
	Put_StaticValue          uintptr
	Get_StaticValue          uintptr
	Put_ReadProtectionLevel  uintptr
	Get_ReadProtectionLevel  uintptr
	Put_WriteProtectionLevel uintptr
	Get_WriteProtectionLevel uintptr
}

type IGattLocalDescriptorParameters struct {
	win32.IInspectable
}

func (this *IGattLocalDescriptorParameters) Vtbl() *IGattLocalDescriptorParametersVtbl {
	return (*IGattLocalDescriptorParametersVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattLocalDescriptorParameters) Put_StaticValue(value *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StaticValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IGattLocalDescriptorParameters) Get_StaticValue() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StaticValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattLocalDescriptorParameters) Put_ReadProtectionLevel(value GattProtectionLevel) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReadProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGattLocalDescriptorParameters) Get_ReadProtectionLevel() GattProtectionLevel {
	var _result GattProtectionLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReadProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattLocalDescriptorParameters) Put_WriteProtectionLevel(value GattProtectionLevel) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_WriteProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGattLocalDescriptorParameters) Get_WriteProtectionLevel() GattProtectionLevel {
	var _result GattProtectionLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WriteProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 375791BE-321F-4366-BFC1-3BC6B82C79F8
var IID_IGattLocalDescriptorResult = syscall.GUID{0x375791BE, 0x321F, 0x4366,
	[8]byte{0xBF, 0xC1, 0x3B, 0xC6, 0xB8, 0x2C, 0x79, 0xF8}}

type IGattLocalDescriptorResultInterface interface {
	win32.IInspectableInterface
	Get_Descriptor() *IGattLocalDescriptor
	Get_Error() BluetoothError
}

type IGattLocalDescriptorResultVtbl struct {
	win32.IInspectableVtbl
	Get_Descriptor uintptr
	Get_Error      uintptr
}

type IGattLocalDescriptorResult struct {
	win32.IInspectable
}

func (this *IGattLocalDescriptorResult) Vtbl() *IGattLocalDescriptorResultVtbl {
	return (*IGattLocalDescriptorResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattLocalDescriptorResult) Get_Descriptor() *IGattLocalDescriptor {
	var _result *IGattLocalDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Descriptor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattLocalDescriptorResult) Get_Error() BluetoothError {
	var _result BluetoothError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Error, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F513E258-F7F7-4902-B803-57FCC7D6FE83
var IID_IGattLocalService = syscall.GUID{0xF513E258, 0xF7F7, 0x4902,
	[8]byte{0xB8, 0x03, 0x57, 0xFC, 0xC7, 0xD6, 0xFE, 0x83}}

type IGattLocalServiceInterface interface {
	win32.IInspectableInterface
	Get_Uuid() syscall.GUID
	CreateCharacteristicAsync(characteristicUuid syscall.GUID, parameters *IGattLocalCharacteristicParameters) *IAsyncOperation[*IGattLocalCharacteristicResult]
	Get_Characteristics() *IVectorView[*IGattLocalCharacteristic]
}

type IGattLocalServiceVtbl struct {
	win32.IInspectableVtbl
	Get_Uuid                  uintptr
	CreateCharacteristicAsync uintptr
	Get_Characteristics       uintptr
}

type IGattLocalService struct {
	win32.IInspectable
}

func (this *IGattLocalService) Vtbl() *IGattLocalServiceVtbl {
	return (*IGattLocalServiceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattLocalService) Get_Uuid() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Uuid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattLocalService) CreateCharacteristicAsync(characteristicUuid syscall.GUID, parameters *IGattLocalCharacteristicParameters) *IAsyncOperation[*IGattLocalCharacteristicResult] {
	var _result *IAsyncOperation[*IGattLocalCharacteristicResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCharacteristicAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&characteristicUuid)), uintptr(unsafe.Pointer(parameters)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattLocalService) Get_Characteristics() *IVectorView[*IGattLocalCharacteristic] {
	var _result *IVectorView[*IGattLocalCharacteristic]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Characteristics, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 196D0021-FAAD-45DC-AE5B-2AC3184E84DB
var IID_IGattPresentationFormat = syscall.GUID{0x196D0021, 0xFAAD, 0x45DC,
	[8]byte{0xAE, 0x5B, 0x2A, 0xC3, 0x18, 0x4E, 0x84, 0xDB}}

type IGattPresentationFormatInterface interface {
	win32.IInspectableInterface
	Get_FormatType() byte
	Get_Exponent() int32
	Get_Unit() uint16
	Get_Namespace() byte
	Get_Description() uint16
}

type IGattPresentationFormatVtbl struct {
	win32.IInspectableVtbl
	Get_FormatType  uintptr
	Get_Exponent    uintptr
	Get_Unit        uintptr
	Get_Namespace   uintptr
	Get_Description uintptr
}

type IGattPresentationFormat struct {
	win32.IInspectable
}

func (this *IGattPresentationFormat) Vtbl() *IGattPresentationFormatVtbl {
	return (*IGattPresentationFormatVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattPresentationFormat) Get_FormatType() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FormatType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormat) Get_Exponent() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Exponent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormat) Get_Unit() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Unit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormat) Get_Namespace() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Namespace, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormat) Get_Description() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Description, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 196D0020-FAAD-45DC-AE5B-2AC3184E84DB
var IID_IGattPresentationFormatStatics = syscall.GUID{0x196D0020, 0xFAAD, 0x45DC,
	[8]byte{0xAE, 0x5B, 0x2A, 0xC3, 0x18, 0x4E, 0x84, 0xDB}}

type IGattPresentationFormatStaticsInterface interface {
	win32.IInspectableInterface
	Get_BluetoothSigAssignedNumbers() byte
}

type IGattPresentationFormatStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_BluetoothSigAssignedNumbers uintptr
}

type IGattPresentationFormatStatics struct {
	win32.IInspectable
}

func (this *IGattPresentationFormatStatics) Vtbl() *IGattPresentationFormatStaticsVtbl {
	return (*IGattPresentationFormatStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattPresentationFormatStatics) Get_BluetoothSigAssignedNumbers() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BluetoothSigAssignedNumbers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A9C21713-B82F-435E-B634-21FD85A43C07
var IID_IGattPresentationFormatStatics2 = syscall.GUID{0xA9C21713, 0xB82F, 0x435E,
	[8]byte{0xB6, 0x34, 0x21, 0xFD, 0x85, 0xA4, 0x3C, 0x07}}

type IGattPresentationFormatStatics2Interface interface {
	win32.IInspectableInterface
	FromParts(formatType byte, exponent int32, unit uint16, namespaceId byte, description uint16) *IGattPresentationFormat
}

type IGattPresentationFormatStatics2Vtbl struct {
	win32.IInspectableVtbl
	FromParts uintptr
}

type IGattPresentationFormatStatics2 struct {
	win32.IInspectable
}

func (this *IGattPresentationFormatStatics2) Vtbl() *IGattPresentationFormatStatics2Vtbl {
	return (*IGattPresentationFormatStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattPresentationFormatStatics2) FromParts(formatType byte, exponent int32, unit uint16, namespaceId byte, description uint16) *IGattPresentationFormat {
	var _result *IGattPresentationFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromParts, uintptr(unsafe.Pointer(this)), uintptr(formatType), uintptr(exponent), uintptr(unit), uintptr(namespaceId), uintptr(description), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FAF1BA0A-30BA-409C-BEF7-CFFB6D03B8FB
var IID_IGattPresentationFormatTypesStatics = syscall.GUID{0xFAF1BA0A, 0x30BA, 0x409C,
	[8]byte{0xBE, 0xF7, 0xCF, 0xFB, 0x6D, 0x03, 0xB8, 0xFB}}

type IGattPresentationFormatTypesStaticsInterface interface {
	win32.IInspectableInterface
	Get_Boolean() byte
	Get_Bit2() byte
	Get_Nibble() byte
	Get_UInt8() byte
	Get_UInt12() byte
	Get_UInt16() byte
	Get_UInt24() byte
	Get_UInt32() byte
	Get_UInt48() byte
	Get_UInt64() byte
	Get_UInt128() byte
	Get_SInt8() byte
	Get_SInt12() byte
	Get_SInt16() byte
	Get_SInt24() byte
	Get_SInt32() byte
	Get_SInt48() byte
	Get_SInt64() byte
	Get_SInt128() byte
	Get_Float32() byte
	Get_Float64() byte
	Get_SFloat() byte
	Get_Float() byte
	Get_DUInt16() byte
	Get_Utf8() byte
	Get_Utf16() byte
	Get_Struct() byte
}

type IGattPresentationFormatTypesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Boolean uintptr
	Get_Bit2    uintptr
	Get_Nibble  uintptr
	Get_UInt8   uintptr
	Get_UInt12  uintptr
	Get_UInt16  uintptr
	Get_UInt24  uintptr
	Get_UInt32  uintptr
	Get_UInt48  uintptr
	Get_UInt64  uintptr
	Get_UInt128 uintptr
	Get_SInt8   uintptr
	Get_SInt12  uintptr
	Get_SInt16  uintptr
	Get_SInt24  uintptr
	Get_SInt32  uintptr
	Get_SInt48  uintptr
	Get_SInt64  uintptr
	Get_SInt128 uintptr
	Get_Float32 uintptr
	Get_Float64 uintptr
	Get_SFloat  uintptr
	Get_Float   uintptr
	Get_DUInt16 uintptr
	Get_Utf8    uintptr
	Get_Utf16   uintptr
	Get_Struct  uintptr
}

type IGattPresentationFormatTypesStatics struct {
	win32.IInspectable
}

func (this *IGattPresentationFormatTypesStatics) Vtbl() *IGattPresentationFormatTypesStaticsVtbl {
	return (*IGattPresentationFormatTypesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattPresentationFormatTypesStatics) Get_Boolean() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Boolean, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_Bit2() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bit2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_Nibble() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Nibble, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_UInt8() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UInt8, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_UInt12() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UInt12, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_UInt16() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UInt16, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_UInt24() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UInt24, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_UInt32() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UInt32, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_UInt48() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UInt48, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_UInt64() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UInt64, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_UInt128() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UInt128, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_SInt8() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SInt8, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_SInt12() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SInt12, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_SInt16() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SInt16, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_SInt24() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SInt24, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_SInt32() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SInt32, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_SInt48() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SInt48, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_SInt64() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SInt64, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_SInt128() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SInt128, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_Float32() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Float32, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_Float64() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Float64, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_SFloat() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SFloat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_Float() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Float, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_DUInt16() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DUInt16, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_Utf8() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Utf8, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_Utf16() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Utf16, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattPresentationFormatTypesStatics) Get_Struct() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Struct, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// CA46C5C5-0ECC-4809-BEA3-CF79BC991E37
var IID_IGattProtocolErrorStatics = syscall.GUID{0xCA46C5C5, 0x0ECC, 0x4809,
	[8]byte{0xBE, 0xA3, 0xCF, 0x79, 0xBC, 0x99, 0x1E, 0x37}}

type IGattProtocolErrorStaticsInterface interface {
	win32.IInspectableInterface
	Get_InvalidHandle() byte
	Get_ReadNotPermitted() byte
	Get_WriteNotPermitted() byte
	Get_InvalidPdu() byte
	Get_InsufficientAuthentication() byte
	Get_RequestNotSupported() byte
	Get_InvalidOffset() byte
	Get_InsufficientAuthorization() byte
	Get_PrepareQueueFull() byte
	Get_AttributeNotFound() byte
	Get_AttributeNotLong() byte
	Get_InsufficientEncryptionKeySize() byte
	Get_InvalidAttributeValueLength() byte
	Get_UnlikelyError() byte
	Get_InsufficientEncryption() byte
	Get_UnsupportedGroupType() byte
	Get_InsufficientResources() byte
}

type IGattProtocolErrorStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_InvalidHandle                 uintptr
	Get_ReadNotPermitted              uintptr
	Get_WriteNotPermitted             uintptr
	Get_InvalidPdu                    uintptr
	Get_InsufficientAuthentication    uintptr
	Get_RequestNotSupported           uintptr
	Get_InvalidOffset                 uintptr
	Get_InsufficientAuthorization     uintptr
	Get_PrepareQueueFull              uintptr
	Get_AttributeNotFound             uintptr
	Get_AttributeNotLong              uintptr
	Get_InsufficientEncryptionKeySize uintptr
	Get_InvalidAttributeValueLength   uintptr
	Get_UnlikelyError                 uintptr
	Get_InsufficientEncryption        uintptr
	Get_UnsupportedGroupType          uintptr
	Get_InsufficientResources         uintptr
}

type IGattProtocolErrorStatics struct {
	win32.IInspectable
}

func (this *IGattProtocolErrorStatics) Vtbl() *IGattProtocolErrorStaticsVtbl {
	return (*IGattProtocolErrorStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattProtocolErrorStatics) Get_InvalidHandle() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InvalidHandle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattProtocolErrorStatics) Get_ReadNotPermitted() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReadNotPermitted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattProtocolErrorStatics) Get_WriteNotPermitted() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WriteNotPermitted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattProtocolErrorStatics) Get_InvalidPdu() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InvalidPdu, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattProtocolErrorStatics) Get_InsufficientAuthentication() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InsufficientAuthentication, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattProtocolErrorStatics) Get_RequestNotSupported() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestNotSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattProtocolErrorStatics) Get_InvalidOffset() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InvalidOffset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattProtocolErrorStatics) Get_InsufficientAuthorization() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InsufficientAuthorization, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattProtocolErrorStatics) Get_PrepareQueueFull() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrepareQueueFull, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattProtocolErrorStatics) Get_AttributeNotFound() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AttributeNotFound, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattProtocolErrorStatics) Get_AttributeNotLong() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AttributeNotLong, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattProtocolErrorStatics) Get_InsufficientEncryptionKeySize() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InsufficientEncryptionKeySize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattProtocolErrorStatics) Get_InvalidAttributeValueLength() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InvalidAttributeValueLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattProtocolErrorStatics) Get_UnlikelyError() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UnlikelyError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattProtocolErrorStatics) Get_InsufficientEncryption() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InsufficientEncryption, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattProtocolErrorStatics) Get_UnsupportedGroupType() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UnsupportedGroupType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattProtocolErrorStatics) Get_InsufficientResources() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InsufficientResources, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 63A66F09-1AEA-4C4C-A50F-97BAE474B348
var IID_IGattReadClientCharacteristicConfigurationDescriptorResult = syscall.GUID{0x63A66F09, 0x1AEA, 0x4C4C,
	[8]byte{0xA5, 0x0F, 0x97, 0xBA, 0xE4, 0x74, 0xB3, 0x48}}

type IGattReadClientCharacteristicConfigurationDescriptorResultInterface interface {
	win32.IInspectableInterface
	Get_Status() GattCommunicationStatus
	Get_ClientCharacteristicConfigurationDescriptor() GattClientCharacteristicConfigurationDescriptorValue
}

type IGattReadClientCharacteristicConfigurationDescriptorResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status                                      uintptr
	Get_ClientCharacteristicConfigurationDescriptor uintptr
}

type IGattReadClientCharacteristicConfigurationDescriptorResult struct {
	win32.IInspectable
}

func (this *IGattReadClientCharacteristicConfigurationDescriptorResult) Vtbl() *IGattReadClientCharacteristicConfigurationDescriptorResultVtbl {
	return (*IGattReadClientCharacteristicConfigurationDescriptorResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattReadClientCharacteristicConfigurationDescriptorResult) Get_Status() GattCommunicationStatus {
	var _result GattCommunicationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattReadClientCharacteristicConfigurationDescriptorResult) Get_ClientCharacteristicConfigurationDescriptor() GattClientCharacteristicConfigurationDescriptorValue {
	var _result GattClientCharacteristicConfigurationDescriptorValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ClientCharacteristicConfigurationDescriptor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1BF1A59D-BA4D-4622-8651-F4EE150D0A5D
var IID_IGattReadClientCharacteristicConfigurationDescriptorResult2 = syscall.GUID{0x1BF1A59D, 0xBA4D, 0x4622,
	[8]byte{0x86, 0x51, 0xF4, 0xEE, 0x15, 0x0D, 0x0A, 0x5D}}

type IGattReadClientCharacteristicConfigurationDescriptorResult2Interface interface {
	win32.IInspectableInterface
	Get_ProtocolError() *IReference[byte]
}

type IGattReadClientCharacteristicConfigurationDescriptorResult2Vtbl struct {
	win32.IInspectableVtbl
	Get_ProtocolError uintptr
}

type IGattReadClientCharacteristicConfigurationDescriptorResult2 struct {
	win32.IInspectable
}

func (this *IGattReadClientCharacteristicConfigurationDescriptorResult2) Vtbl() *IGattReadClientCharacteristicConfigurationDescriptorResult2Vtbl {
	return (*IGattReadClientCharacteristicConfigurationDescriptorResult2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattReadClientCharacteristicConfigurationDescriptorResult2) Get_ProtocolError() *IReference[byte] {
	var _result *IReference[byte]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtocolError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F1DD6535-6ACD-42A6-A4BB-D789DAE0043E
var IID_IGattReadRequest = syscall.GUID{0xF1DD6535, 0x6ACD, 0x42A6,
	[8]byte{0xA4, 0xBB, 0xD7, 0x89, 0xDA, 0xE0, 0x04, 0x3E}}

type IGattReadRequestInterface interface {
	win32.IInspectableInterface
	Get_Offset() uint32
	Get_Length() uint32
	Get_State() GattRequestState
	Add_StateChanged(handler TypedEventHandler[*IGattReadRequest, *IGattRequestStateChangedEventArgs]) EventRegistrationToken
	Remove_StateChanged(token EventRegistrationToken)
	RespondWithValue(value *IBuffer)
	RespondWithProtocolError(protocolError byte)
}

type IGattReadRequestVtbl struct {
	win32.IInspectableVtbl
	Get_Offset               uintptr
	Get_Length               uintptr
	Get_State                uintptr
	Add_StateChanged         uintptr
	Remove_StateChanged      uintptr
	RespondWithValue         uintptr
	RespondWithProtocolError uintptr
}

type IGattReadRequest struct {
	win32.IInspectable
}

func (this *IGattReadRequest) Vtbl() *IGattReadRequestVtbl {
	return (*IGattReadRequestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattReadRequest) Get_Offset() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Offset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattReadRequest) Get_Length() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Length, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattReadRequest) Get_State() GattRequestState {
	var _result GattRequestState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattReadRequest) Add_StateChanged(handler TypedEventHandler[*IGattReadRequest, *IGattRequestStateChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattReadRequest) Remove_StateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGattReadRequest) RespondWithValue(value *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RespondWithValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IGattReadRequest) RespondWithProtocolError(protocolError byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RespondWithProtocolError, uintptr(unsafe.Pointer(this)), uintptr(protocolError))
	_ = _hr
}

// 93497243-F39C-484B-8AB6-996BA486CFA3
var IID_IGattReadRequestedEventArgs = syscall.GUID{0x93497243, 0xF39C, 0x484B,
	[8]byte{0x8A, 0xB6, 0x99, 0x6B, 0xA4, 0x86, 0xCF, 0xA3}}

type IGattReadRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Session() *IGattSession
	GetDeferral() *IDeferral
	GetRequestAsync() *IAsyncOperation[*IGattReadRequest]
}

type IGattReadRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Session     uintptr
	GetDeferral     uintptr
	GetRequestAsync uintptr
}

type IGattReadRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IGattReadRequestedEventArgs) Vtbl() *IGattReadRequestedEventArgsVtbl {
	return (*IGattReadRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattReadRequestedEventArgs) Get_Session() *IGattSession {
	var _result *IGattSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Session, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattReadRequestedEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattReadRequestedEventArgs) GetRequestAsync() *IAsyncOperation[*IGattReadRequest] {
	var _result *IAsyncOperation[*IGattReadRequest]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetRequestAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 63A66F08-1AEA-4C4C-A50F-97BAE474B348
var IID_IGattReadResult = syscall.GUID{0x63A66F08, 0x1AEA, 0x4C4C,
	[8]byte{0xA5, 0x0F, 0x97, 0xBA, 0xE4, 0x74, 0xB3, 0x48}}

type IGattReadResultInterface interface {
	win32.IInspectableInterface
	Get_Status() GattCommunicationStatus
	Get_Value() *IBuffer
}

type IGattReadResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
	Get_Value  uintptr
}

type IGattReadResult struct {
	win32.IInspectable
}

func (this *IGattReadResult) Vtbl() *IGattReadResultVtbl {
	return (*IGattReadResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattReadResult) Get_Status() GattCommunicationStatus {
	var _result GattCommunicationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattReadResult) Get_Value() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A10F50A0-FB43-48AF-BAAA-638A5C6329FE
var IID_IGattReadResult2 = syscall.GUID{0xA10F50A0, 0xFB43, 0x48AF,
	[8]byte{0xBA, 0xAA, 0x63, 0x8A, 0x5C, 0x63, 0x29, 0xFE}}

type IGattReadResult2Interface interface {
	win32.IInspectableInterface
	Get_ProtocolError() *IReference[byte]
}

type IGattReadResult2Vtbl struct {
	win32.IInspectableVtbl
	Get_ProtocolError uintptr
}

type IGattReadResult2 struct {
	win32.IInspectable
}

func (this *IGattReadResult2) Vtbl() *IGattReadResult2Vtbl {
	return (*IGattReadResult2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattReadResult2) Get_ProtocolError() *IReference[byte] {
	var _result *IReference[byte]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtocolError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 63A66F07-1AEA-4C4C-A50F-97BAE474B348
var IID_IGattReliableWriteTransaction = syscall.GUID{0x63A66F07, 0x1AEA, 0x4C4C,
	[8]byte{0xA5, 0x0F, 0x97, 0xBA, 0xE4, 0x74, 0xB3, 0x48}}

type IGattReliableWriteTransactionInterface interface {
	win32.IInspectableInterface
	WriteValue(characteristic *IGattCharacteristic, value *IBuffer)
	CommitAsync() *IAsyncOperation[GattCommunicationStatus]
}

type IGattReliableWriteTransactionVtbl struct {
	win32.IInspectableVtbl
	WriteValue  uintptr
	CommitAsync uintptr
}

type IGattReliableWriteTransaction struct {
	win32.IInspectable
}

func (this *IGattReliableWriteTransaction) Vtbl() *IGattReliableWriteTransactionVtbl {
	return (*IGattReliableWriteTransactionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattReliableWriteTransaction) WriteValue(characteristic *IGattCharacteristic, value *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(characteristic)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IGattReliableWriteTransaction) CommitAsync() *IAsyncOperation[GattCommunicationStatus] {
	var _result *IAsyncOperation[GattCommunicationStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CommitAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 51113987-EF12-462F-9FB2-A1A43A679416
var IID_IGattReliableWriteTransaction2 = syscall.GUID{0x51113987, 0xEF12, 0x462F,
	[8]byte{0x9F, 0xB2, 0xA1, 0xA4, 0x3A, 0x67, 0x94, 0x16}}

type IGattReliableWriteTransaction2Interface interface {
	win32.IInspectableInterface
	CommitWithResultAsync() *IAsyncOperation[*IGattWriteResult]
}

type IGattReliableWriteTransaction2Vtbl struct {
	win32.IInspectableVtbl
	CommitWithResultAsync uintptr
}

type IGattReliableWriteTransaction2 struct {
	win32.IInspectable
}

func (this *IGattReliableWriteTransaction2) Vtbl() *IGattReliableWriteTransaction2Vtbl {
	return (*IGattReliableWriteTransaction2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattReliableWriteTransaction2) CommitWithResultAsync() *IAsyncOperation[*IGattWriteResult] {
	var _result *IAsyncOperation[*IGattWriteResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CommitWithResultAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E834D92C-27BE-44B3-9D0D-4FC6E808DD3F
var IID_IGattRequestStateChangedEventArgs = syscall.GUID{0xE834D92C, 0x27BE, 0x44B3,
	[8]byte{0x9D, 0x0D, 0x4F, 0xC6, 0xE8, 0x08, 0xDD, 0x3F}}

type IGattRequestStateChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_State() GattRequestState
	Get_Error() BluetoothError
}

type IGattRequestStateChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_State uintptr
	Get_Error uintptr
}

type IGattRequestStateChangedEventArgs struct {
	win32.IInspectable
}

func (this *IGattRequestStateChangedEventArgs) Vtbl() *IGattRequestStateChangedEventArgsVtbl {
	return (*IGattRequestStateChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattRequestStateChangedEventArgs) Get_State() GattRequestState {
	var _result GattRequestState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattRequestStateChangedEventArgs) Get_Error() BluetoothError {
	var _result BluetoothError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Error, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7822B3CD-2889-4F86-A051-3F0AED1C2760
var IID_IGattServiceProvider = syscall.GUID{0x7822B3CD, 0x2889, 0x4F86,
	[8]byte{0xA0, 0x51, 0x3F, 0x0A, 0xED, 0x1C, 0x27, 0x60}}

type IGattServiceProviderInterface interface {
	win32.IInspectableInterface
	Get_Service() *IGattLocalService
	Get_AdvertisementStatus() GattServiceProviderAdvertisementStatus
	Add_AdvertisementStatusChanged(handler TypedEventHandler[*IGattServiceProvider, *IGattServiceProviderAdvertisementStatusChangedEventArgs]) EventRegistrationToken
	Remove_AdvertisementStatusChanged(token EventRegistrationToken)
	StartAdvertising()
	StartAdvertisingWithParameters(parameters *IGattServiceProviderAdvertisingParameters)
	StopAdvertising()
}

type IGattServiceProviderVtbl struct {
	win32.IInspectableVtbl
	Get_Service                       uintptr
	Get_AdvertisementStatus           uintptr
	Add_AdvertisementStatusChanged    uintptr
	Remove_AdvertisementStatusChanged uintptr
	StartAdvertising                  uintptr
	StartAdvertisingWithParameters    uintptr
	StopAdvertising                   uintptr
}

type IGattServiceProvider struct {
	win32.IInspectable
}

func (this *IGattServiceProvider) Vtbl() *IGattServiceProviderVtbl {
	return (*IGattServiceProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattServiceProvider) Get_Service() *IGattLocalService {
	var _result *IGattLocalService
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Service, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattServiceProvider) Get_AdvertisementStatus() GattServiceProviderAdvertisementStatus {
	var _result GattServiceProviderAdvertisementStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AdvertisementStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattServiceProvider) Add_AdvertisementStatusChanged(handler TypedEventHandler[*IGattServiceProvider, *IGattServiceProviderAdvertisementStatusChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AdvertisementStatusChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattServiceProvider) Remove_AdvertisementStatusChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AdvertisementStatusChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGattServiceProvider) StartAdvertising() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartAdvertising, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IGattServiceProvider) StartAdvertisingWithParameters(parameters *IGattServiceProviderAdvertisingParameters) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartAdvertisingWithParameters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(parameters)))
	_ = _hr
}

func (this *IGattServiceProvider) StopAdvertising() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopAdvertising, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 59A5AA65-FA21-4FFC-B155-04D928012686
var IID_IGattServiceProviderAdvertisementStatusChangedEventArgs = syscall.GUID{0x59A5AA65, 0xFA21, 0x4FFC,
	[8]byte{0xB1, 0x55, 0x04, 0xD9, 0x28, 0x01, 0x26, 0x86}}

type IGattServiceProviderAdvertisementStatusChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Error() BluetoothError
	Get_Status() GattServiceProviderAdvertisementStatus
}

type IGattServiceProviderAdvertisementStatusChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Error  uintptr
	Get_Status uintptr
}

type IGattServiceProviderAdvertisementStatusChangedEventArgs struct {
	win32.IInspectable
}

func (this *IGattServiceProviderAdvertisementStatusChangedEventArgs) Vtbl() *IGattServiceProviderAdvertisementStatusChangedEventArgsVtbl {
	return (*IGattServiceProviderAdvertisementStatusChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattServiceProviderAdvertisementStatusChangedEventArgs) Get_Error() BluetoothError {
	var _result BluetoothError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Error, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattServiceProviderAdvertisementStatusChangedEventArgs) Get_Status() GattServiceProviderAdvertisementStatus {
	var _result GattServiceProviderAdvertisementStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// E2CE31AB-6315-4C22-9BD7-781DBC3D8D82
var IID_IGattServiceProviderAdvertisingParameters = syscall.GUID{0xE2CE31AB, 0x6315, 0x4C22,
	[8]byte{0x9B, 0xD7, 0x78, 0x1D, 0xBC, 0x3D, 0x8D, 0x82}}

type IGattServiceProviderAdvertisingParametersInterface interface {
	win32.IInspectableInterface
	Put_IsConnectable(value bool)
	Get_IsConnectable() bool
	Put_IsDiscoverable(value bool)
	Get_IsDiscoverable() bool
}

type IGattServiceProviderAdvertisingParametersVtbl struct {
	win32.IInspectableVtbl
	Put_IsConnectable  uintptr
	Get_IsConnectable  uintptr
	Put_IsDiscoverable uintptr
	Get_IsDiscoverable uintptr
}

type IGattServiceProviderAdvertisingParameters struct {
	win32.IInspectable
}

func (this *IGattServiceProviderAdvertisingParameters) Vtbl() *IGattServiceProviderAdvertisingParametersVtbl {
	return (*IGattServiceProviderAdvertisingParametersVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattServiceProviderAdvertisingParameters) Put_IsConnectable(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsConnectable, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IGattServiceProviderAdvertisingParameters) Get_IsConnectable() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsConnectable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattServiceProviderAdvertisingParameters) Put_IsDiscoverable(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsDiscoverable, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IGattServiceProviderAdvertisingParameters) Get_IsDiscoverable() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDiscoverable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// FF68468D-CA92-4434-9743-0E90988AD879
var IID_IGattServiceProviderAdvertisingParameters2 = syscall.GUID{0xFF68468D, 0xCA92, 0x4434,
	[8]byte{0x97, 0x43, 0x0E, 0x90, 0x98, 0x8A, 0xD8, 0x79}}

type IGattServiceProviderAdvertisingParameters2Interface interface {
	win32.IInspectableInterface
	Put_ServiceData(value *IBuffer)
	Get_ServiceData() *IBuffer
}

type IGattServiceProviderAdvertisingParameters2Vtbl struct {
	win32.IInspectableVtbl
	Put_ServiceData uintptr
	Get_ServiceData uintptr
}

type IGattServiceProviderAdvertisingParameters2 struct {
	win32.IInspectable
}

func (this *IGattServiceProviderAdvertisingParameters2) Vtbl() *IGattServiceProviderAdvertisingParameters2Vtbl {
	return (*IGattServiceProviderAdvertisingParameters2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattServiceProviderAdvertisingParameters2) Put_ServiceData(value *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ServiceData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IGattServiceProviderAdvertisingParameters2) Get_ServiceData() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 764696D8-C53E-428C-8A48-67AFE02C3AE6
var IID_IGattServiceProviderResult = syscall.GUID{0x764696D8, 0xC53E, 0x428C,
	[8]byte{0x8A, 0x48, 0x67, 0xAF, 0xE0, 0x2C, 0x3A, 0xE6}}

type IGattServiceProviderResultInterface interface {
	win32.IInspectableInterface
	Get_Error() BluetoothError
	Get_ServiceProvider() *IGattServiceProvider
}

type IGattServiceProviderResultVtbl struct {
	win32.IInspectableVtbl
	Get_Error           uintptr
	Get_ServiceProvider uintptr
}

type IGattServiceProviderResult struct {
	win32.IInspectable
}

func (this *IGattServiceProviderResult) Vtbl() *IGattServiceProviderResultVtbl {
	return (*IGattServiceProviderResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattServiceProviderResult) Get_Error() BluetoothError {
	var _result BluetoothError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Error, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattServiceProviderResult) Get_ServiceProvider() *IGattServiceProvider {
	var _result *IGattServiceProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 31794063-5256-4054-A4F4-7BBE7755A57E
var IID_IGattServiceProviderStatics = syscall.GUID{0x31794063, 0x5256, 0x4054,
	[8]byte{0xA4, 0xF4, 0x7B, 0xBE, 0x77, 0x55, 0xA5, 0x7E}}

type IGattServiceProviderStaticsInterface interface {
	win32.IInspectableInterface
	CreateAsync(serviceUuid syscall.GUID) *IAsyncOperation[*IGattServiceProviderResult]
}

type IGattServiceProviderStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateAsync uintptr
}

type IGattServiceProviderStatics struct {
	win32.IInspectable
}

func (this *IGattServiceProviderStatics) Vtbl() *IGattServiceProviderStaticsVtbl {
	return (*IGattServiceProviderStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattServiceProviderStatics) CreateAsync(serviceUuid syscall.GUID) *IAsyncOperation[*IGattServiceProviderResult] {
	var _result *IAsyncOperation[*IGattServiceProviderResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&serviceUuid)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6DC57058-9ABA-4417-B8F2-DCE016D34EE2
var IID_IGattServiceUuidsStatics = syscall.GUID{0x6DC57058, 0x9ABA, 0x4417,
	[8]byte{0xB8, 0xF2, 0xDC, 0xE0, 0x16, 0xD3, 0x4E, 0xE2}}

type IGattServiceUuidsStaticsInterface interface {
	win32.IInspectableInterface
	Get_Battery() syscall.GUID
	Get_BloodPressure() syscall.GUID
	Get_CyclingSpeedAndCadence() syscall.GUID
	Get_GenericAccess() syscall.GUID
	Get_GenericAttribute() syscall.GUID
	Get_Glucose() syscall.GUID
	Get_HealthThermometer() syscall.GUID
	Get_HeartRate() syscall.GUID
	Get_RunningSpeedAndCadence() syscall.GUID
}

type IGattServiceUuidsStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Battery                uintptr
	Get_BloodPressure          uintptr
	Get_CyclingSpeedAndCadence uintptr
	Get_GenericAccess          uintptr
	Get_GenericAttribute       uintptr
	Get_Glucose                uintptr
	Get_HealthThermometer      uintptr
	Get_HeartRate              uintptr
	Get_RunningSpeedAndCadence uintptr
}

type IGattServiceUuidsStatics struct {
	win32.IInspectable
}

func (this *IGattServiceUuidsStatics) Vtbl() *IGattServiceUuidsStaticsVtbl {
	return (*IGattServiceUuidsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattServiceUuidsStatics) Get_Battery() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Battery, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattServiceUuidsStatics) Get_BloodPressure() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BloodPressure, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattServiceUuidsStatics) Get_CyclingSpeedAndCadence() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CyclingSpeedAndCadence, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattServiceUuidsStatics) Get_GenericAccess() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GenericAccess, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattServiceUuidsStatics) Get_GenericAttribute() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GenericAttribute, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattServiceUuidsStatics) Get_Glucose() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Glucose, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattServiceUuidsStatics) Get_HealthThermometer() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HealthThermometer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattServiceUuidsStatics) Get_HeartRate() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HeartRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattServiceUuidsStatics) Get_RunningSpeedAndCadence() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RunningSpeedAndCadence, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D2AE94F5-3D15-4F79-9C0C-EAAFA675155C
var IID_IGattServiceUuidsStatics2 = syscall.GUID{0xD2AE94F5, 0x3D15, 0x4F79,
	[8]byte{0x9C, 0x0C, 0xEA, 0xAF, 0xA6, 0x75, 0x15, 0x5C}}

type IGattServiceUuidsStatics2Interface interface {
	win32.IInspectableInterface
	Get_AlertNotification() syscall.GUID
	Get_CurrentTime() syscall.GUID
	Get_CyclingPower() syscall.GUID
	Get_DeviceInformation() syscall.GUID
	Get_HumanInterfaceDevice() syscall.GUID
	Get_ImmediateAlert() syscall.GUID
	Get_LinkLoss() syscall.GUID
	Get_LocationAndNavigation() syscall.GUID
	Get_NextDstChange() syscall.GUID
	Get_PhoneAlertStatus() syscall.GUID
	Get_ReferenceTimeUpdate() syscall.GUID
	Get_ScanParameters() syscall.GUID
	Get_TxPower() syscall.GUID
}

type IGattServiceUuidsStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_AlertNotification     uintptr
	Get_CurrentTime           uintptr
	Get_CyclingPower          uintptr
	Get_DeviceInformation     uintptr
	Get_HumanInterfaceDevice  uintptr
	Get_ImmediateAlert        uintptr
	Get_LinkLoss              uintptr
	Get_LocationAndNavigation uintptr
	Get_NextDstChange         uintptr
	Get_PhoneAlertStatus      uintptr
	Get_ReferenceTimeUpdate   uintptr
	Get_ScanParameters        uintptr
	Get_TxPower               uintptr
}

type IGattServiceUuidsStatics2 struct {
	win32.IInspectable
}

func (this *IGattServiceUuidsStatics2) Vtbl() *IGattServiceUuidsStatics2Vtbl {
	return (*IGattServiceUuidsStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattServiceUuidsStatics2) Get_AlertNotification() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlertNotification, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattServiceUuidsStatics2) Get_CurrentTime() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattServiceUuidsStatics2) Get_CyclingPower() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CyclingPower, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattServiceUuidsStatics2) Get_DeviceInformation() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattServiceUuidsStatics2) Get_HumanInterfaceDevice() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HumanInterfaceDevice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattServiceUuidsStatics2) Get_ImmediateAlert() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ImmediateAlert, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattServiceUuidsStatics2) Get_LinkLoss() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LinkLoss, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattServiceUuidsStatics2) Get_LocationAndNavigation() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocationAndNavigation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattServiceUuidsStatics2) Get_NextDstChange() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NextDstChange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattServiceUuidsStatics2) Get_PhoneAlertStatus() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhoneAlertStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattServiceUuidsStatics2) Get_ReferenceTimeUpdate() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReferenceTimeUpdate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattServiceUuidsStatics2) Get_ScanParameters() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScanParameters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattServiceUuidsStatics2) Get_TxPower() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TxPower, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D23B5143-E04E-4C24-999C-9C256F9856B1
var IID_IGattSession = syscall.GUID{0xD23B5143, 0xE04E, 0x4C24,
	[8]byte{0x99, 0x9C, 0x9C, 0x25, 0x6F, 0x98, 0x56, 0xB1}}

type IGattSessionInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() *IBluetoothDeviceId
	Get_CanMaintainConnection() bool
	Put_MaintainConnection(value bool)
	Get_MaintainConnection() bool
	Get_MaxPduSize() uint16
	Get_SessionStatus() GattSessionStatus
	Add_MaxPduSizeChanged(handler TypedEventHandler[*IGattSession, interface{}]) EventRegistrationToken
	Remove_MaxPduSizeChanged(token EventRegistrationToken)
	Add_SessionStatusChanged(handler TypedEventHandler[*IGattSession, *IGattSessionStatusChangedEventArgs]) EventRegistrationToken
	Remove_SessionStatusChanged(token EventRegistrationToken)
}

type IGattSessionVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId                uintptr
	Get_CanMaintainConnection   uintptr
	Put_MaintainConnection      uintptr
	Get_MaintainConnection      uintptr
	Get_MaxPduSize              uintptr
	Get_SessionStatus           uintptr
	Add_MaxPduSizeChanged       uintptr
	Remove_MaxPduSizeChanged    uintptr
	Add_SessionStatusChanged    uintptr
	Remove_SessionStatusChanged uintptr
}

type IGattSession struct {
	win32.IInspectable
}

func (this *IGattSession) Vtbl() *IGattSessionVtbl {
	return (*IGattSessionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattSession) Get_DeviceId() *IBluetoothDeviceId {
	var _result *IBluetoothDeviceId
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattSession) Get_CanMaintainConnection() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanMaintainConnection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattSession) Put_MaintainConnection(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaintainConnection, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IGattSession) Get_MaintainConnection() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaintainConnection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattSession) Get_MaxPduSize() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxPduSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattSession) Get_SessionStatus() GattSessionStatus {
	var _result GattSessionStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SessionStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattSession) Add_MaxPduSizeChanged(handler TypedEventHandler[*IGattSession, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_MaxPduSizeChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattSession) Remove_MaxPduSizeChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_MaxPduSizeChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGattSession) Add_SessionStatusChanged(handler TypedEventHandler[*IGattSession, *IGattSessionStatusChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SessionStatusChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattSession) Remove_SessionStatusChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SessionStatusChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 2E65B95C-539F-4DB7-82A8-73BDBBF73EBF
var IID_IGattSessionStatics = syscall.GUID{0x2E65B95C, 0x539F, 0x4DB7,
	[8]byte{0x82, 0xA8, 0x73, 0xBD, 0xBB, 0xF7, 0x3E, 0xBF}}

type IGattSessionStaticsInterface interface {
	win32.IInspectableInterface
	FromDeviceIdAsync(deviceId *IBluetoothDeviceId) *IAsyncOperation[*IGattSession]
}

type IGattSessionStaticsVtbl struct {
	win32.IInspectableVtbl
	FromDeviceIdAsync uintptr
}

type IGattSessionStatics struct {
	win32.IInspectable
}

func (this *IGattSessionStatics) Vtbl() *IGattSessionStaticsVtbl {
	return (*IGattSessionStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattSessionStatics) FromDeviceIdAsync(deviceId *IBluetoothDeviceId) *IAsyncOperation[*IGattSession] {
	var _result *IAsyncOperation[*IGattSession]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromDeviceIdAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deviceId)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7605B72E-837F-404C-AB34-3163F39DDF32
var IID_IGattSessionStatusChangedEventArgs = syscall.GUID{0x7605B72E, 0x837F, 0x404C,
	[8]byte{0xAB, 0x34, 0x31, 0x63, 0xF3, 0x9D, 0xDF, 0x32}}

type IGattSessionStatusChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Error() BluetoothError
	Get_Status() GattSessionStatus
}

type IGattSessionStatusChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Error  uintptr
	Get_Status uintptr
}

type IGattSessionStatusChangedEventArgs struct {
	win32.IInspectable
}

func (this *IGattSessionStatusChangedEventArgs) Vtbl() *IGattSessionStatusChangedEventArgsVtbl {
	return (*IGattSessionStatusChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattSessionStatusChangedEventArgs) Get_Error() BluetoothError {
	var _result BluetoothError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Error, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattSessionStatusChangedEventArgs) Get_Status() GattSessionStatus {
	var _result GattSessionStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 736E9001-15A4-4EC2-9248-E3F20D463BE9
var IID_IGattSubscribedClient = syscall.GUID{0x736E9001, 0x15A4, 0x4EC2,
	[8]byte{0x92, 0x48, 0xE3, 0xF2, 0x0D, 0x46, 0x3B, 0xE9}}

type IGattSubscribedClientInterface interface {
	win32.IInspectableInterface
	Get_Session() *IGattSession
	Get_MaxNotificationSize() uint16
	Add_MaxNotificationSizeChanged(handler TypedEventHandler[*IGattSubscribedClient, interface{}]) EventRegistrationToken
	Remove_MaxNotificationSizeChanged(token EventRegistrationToken)
}

type IGattSubscribedClientVtbl struct {
	win32.IInspectableVtbl
	Get_Session                       uintptr
	Get_MaxNotificationSize           uintptr
	Add_MaxNotificationSizeChanged    uintptr
	Remove_MaxNotificationSizeChanged uintptr
}

type IGattSubscribedClient struct {
	win32.IInspectable
}

func (this *IGattSubscribedClient) Vtbl() *IGattSubscribedClientVtbl {
	return (*IGattSubscribedClientVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattSubscribedClient) Get_Session() *IGattSession {
	var _result *IGattSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Session, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattSubscribedClient) Get_MaxNotificationSize() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxNotificationSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattSubscribedClient) Add_MaxNotificationSizeChanged(handler TypedEventHandler[*IGattSubscribedClient, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_MaxNotificationSizeChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattSubscribedClient) Remove_MaxNotificationSizeChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_MaxNotificationSizeChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// D21BDB54-06E3-4ED8-A263-ACFAC8BA7313
var IID_IGattValueChangedEventArgs = syscall.GUID{0xD21BDB54, 0x06E3, 0x4ED8,
	[8]byte{0xA2, 0x63, 0xAC, 0xFA, 0xC8, 0xBA, 0x73, 0x13}}

type IGattValueChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_CharacteristicValue() *IBuffer
	Get_Timestamp() DateTime
}

type IGattValueChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_CharacteristicValue uintptr
	Get_Timestamp           uintptr
}

type IGattValueChangedEventArgs struct {
	win32.IInspectable
}

func (this *IGattValueChangedEventArgs) Vtbl() *IGattValueChangedEventArgsVtbl {
	return (*IGattValueChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattValueChangedEventArgs) Get_CharacteristicValue() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CharacteristicValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattValueChangedEventArgs) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// AEB6A9ED-DE2F-4FC2-A9A8-94EA7844F13D
var IID_IGattWriteRequest = syscall.GUID{0xAEB6A9ED, 0xDE2F, 0x4FC2,
	[8]byte{0xA9, 0xA8, 0x94, 0xEA, 0x78, 0x44, 0xF1, 0x3D}}

type IGattWriteRequestInterface interface {
	win32.IInspectableInterface
	Get_Value() *IBuffer
	Get_Offset() uint32
	Get_Option() GattWriteOption
	Get_State() GattRequestState
	Add_StateChanged(handler TypedEventHandler[*IGattWriteRequest, *IGattRequestStateChangedEventArgs]) EventRegistrationToken
	Remove_StateChanged(token EventRegistrationToken)
	Respond()
	RespondWithProtocolError(protocolError byte)
}

type IGattWriteRequestVtbl struct {
	win32.IInspectableVtbl
	Get_Value                uintptr
	Get_Offset               uintptr
	Get_Option               uintptr
	Get_State                uintptr
	Add_StateChanged         uintptr
	Remove_StateChanged      uintptr
	Respond                  uintptr
	RespondWithProtocolError uintptr
}

type IGattWriteRequest struct {
	win32.IInspectable
}

func (this *IGattWriteRequest) Vtbl() *IGattWriteRequestVtbl {
	return (*IGattWriteRequestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattWriteRequest) Get_Value() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattWriteRequest) Get_Offset() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Offset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattWriteRequest) Get_Option() GattWriteOption {
	var _result GattWriteOption
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Option, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattWriteRequest) Get_State() GattRequestState {
	var _result GattRequestState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattWriteRequest) Add_StateChanged(handler TypedEventHandler[*IGattWriteRequest, *IGattRequestStateChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattWriteRequest) Remove_StateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGattWriteRequest) Respond() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Respond, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IGattWriteRequest) RespondWithProtocolError(protocolError byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RespondWithProtocolError, uintptr(unsafe.Pointer(this)), uintptr(protocolError))
	_ = _hr
}

// 2DEC8BBE-A73A-471A-94D5-037DEADD0806
var IID_IGattWriteRequestedEventArgs = syscall.GUID{0x2DEC8BBE, 0xA73A, 0x471A,
	[8]byte{0x94, 0xD5, 0x03, 0x7D, 0xEA, 0xDD, 0x08, 0x06}}

type IGattWriteRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Session() *IGattSession
	GetDeferral() *IDeferral
	GetRequestAsync() *IAsyncOperation[*IGattWriteRequest]
}

type IGattWriteRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Session     uintptr
	GetDeferral     uintptr
	GetRequestAsync uintptr
}

type IGattWriteRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IGattWriteRequestedEventArgs) Vtbl() *IGattWriteRequestedEventArgsVtbl {
	return (*IGattWriteRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattWriteRequestedEventArgs) Get_Session() *IGattSession {
	var _result *IGattSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Session, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattWriteRequestedEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattWriteRequestedEventArgs) GetRequestAsync() *IAsyncOperation[*IGattWriteRequest] {
	var _result *IAsyncOperation[*IGattWriteRequest]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetRequestAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4991DDB1-CB2B-44F7-99FC-D29A2871DC9B
var IID_IGattWriteResult = syscall.GUID{0x4991DDB1, 0xCB2B, 0x44F7,
	[8]byte{0x99, 0xFC, 0xD2, 0x9A, 0x28, 0x71, 0xDC, 0x9B}}

type IGattWriteResultInterface interface {
	win32.IInspectableInterface
	Get_Status() GattCommunicationStatus
	Get_ProtocolError() *IReference[byte]
}

type IGattWriteResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status        uintptr
	Get_ProtocolError uintptr
}

type IGattWriteResult struct {
	win32.IInspectable
}

func (this *IGattWriteResult) Vtbl() *IGattWriteResultVtbl {
	return (*IGattWriteResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattWriteResult) Get_Status() GattCommunicationStatus {
	var _result GattCommunicationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGattWriteResult) Get_ProtocolError() *IReference[byte] {
	var _result *IReference[byte]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtocolError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type GattCharacteristic struct {
	RtClass
	*IGattCharacteristic
}

func NewIGattCharacteristicStatics() *IGattCharacteristicStatics {
	var p *IGattCharacteristicStatics
	hs := NewHStr("Windows.Devices.Bluetooth.GenericAttributeProfile.GattCharacteristic")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGattCharacteristicStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type GattCharacteristicUuids struct {
	RtClass
}

func NewIGattCharacteristicUuidsStatics2() *IGattCharacteristicUuidsStatics2 {
	var p *IGattCharacteristicUuidsStatics2
	hs := NewHStr("Windows.Devices.Bluetooth.GenericAttributeProfile.GattCharacteristicUuids")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGattCharacteristicUuidsStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIGattCharacteristicUuidsStatics() *IGattCharacteristicUuidsStatics {
	var p *IGattCharacteristicUuidsStatics
	hs := NewHStr("Windows.Devices.Bluetooth.GenericAttributeProfile.GattCharacteristicUuids")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGattCharacteristicUuidsStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type GattCharacteristicsResult struct {
	RtClass
	*IGattCharacteristicsResult
}

type GattClientNotificationResult struct {
	RtClass
	*IGattClientNotificationResult
}

type GattDescriptor struct {
	RtClass
	*IGattDescriptor
}

func NewIGattDescriptorStatics() *IGattDescriptorStatics {
	var p *IGattDescriptorStatics
	hs := NewHStr("Windows.Devices.Bluetooth.GenericAttributeProfile.GattDescriptor")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGattDescriptorStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type GattDescriptorUuids struct {
	RtClass
}

func NewIGattDescriptorUuidsStatics() *IGattDescriptorUuidsStatics {
	var p *IGattDescriptorUuidsStatics
	hs := NewHStr("Windows.Devices.Bluetooth.GenericAttributeProfile.GattDescriptorUuids")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGattDescriptorUuidsStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type GattDescriptorsResult struct {
	RtClass
	*IGattDescriptorsResult
}

type GattDeviceService struct {
	RtClass
	*IGattDeviceService
}

func NewIGattDeviceServiceStatics2() *IGattDeviceServiceStatics2 {
	var p *IGattDeviceServiceStatics2
	hs := NewHStr("Windows.Devices.Bluetooth.GenericAttributeProfile.GattDeviceService")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGattDeviceServiceStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIGattDeviceServiceStatics() *IGattDeviceServiceStatics {
	var p *IGattDeviceServiceStatics
	hs := NewHStr("Windows.Devices.Bluetooth.GenericAttributeProfile.GattDeviceService")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGattDeviceServiceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type GattDeviceServicesResult struct {
	RtClass
	*IGattDeviceServicesResult
}

type GattLocalCharacteristic struct {
	RtClass
	*IGattLocalCharacteristic
}

type GattLocalCharacteristicParameters struct {
	RtClass
	*IGattLocalCharacteristicParameters
}

func NewGattLocalCharacteristicParameters() *GattLocalCharacteristicParameters {
	hs := NewHStr("Windows.Devices.Bluetooth.GenericAttributeProfile.GattLocalCharacteristicParameters")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &GattLocalCharacteristicParameters{
		RtClass:                            RtClass{PInspect: p},
		IGattLocalCharacteristicParameters: (*IGattLocalCharacteristicParameters)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type GattLocalCharacteristicResult struct {
	RtClass
	*IGattLocalCharacteristicResult
}

type GattLocalDescriptor struct {
	RtClass
	*IGattLocalDescriptor
}

type GattLocalDescriptorParameters struct {
	RtClass
	*IGattLocalDescriptorParameters
}

func NewGattLocalDescriptorParameters() *GattLocalDescriptorParameters {
	hs := NewHStr("Windows.Devices.Bluetooth.GenericAttributeProfile.GattLocalDescriptorParameters")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &GattLocalDescriptorParameters{
		RtClass:                        RtClass{PInspect: p},
		IGattLocalDescriptorParameters: (*IGattLocalDescriptorParameters)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type GattLocalDescriptorResult struct {
	RtClass
	*IGattLocalDescriptorResult
}

type GattLocalService struct {
	RtClass
	*IGattLocalService
}

type GattPresentationFormat struct {
	RtClass
	*IGattPresentationFormat
}

func NewIGattPresentationFormatStatics() *IGattPresentationFormatStatics {
	var p *IGattPresentationFormatStatics
	hs := NewHStr("Windows.Devices.Bluetooth.GenericAttributeProfile.GattPresentationFormat")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGattPresentationFormatStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIGattPresentationFormatStatics2() *IGattPresentationFormatStatics2 {
	var p *IGattPresentationFormatStatics2
	hs := NewHStr("Windows.Devices.Bluetooth.GenericAttributeProfile.GattPresentationFormat")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGattPresentationFormatStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type GattPresentationFormatTypes struct {
	RtClass
}

func NewIGattPresentationFormatTypesStatics() *IGattPresentationFormatTypesStatics {
	var p *IGattPresentationFormatTypesStatics
	hs := NewHStr("Windows.Devices.Bluetooth.GenericAttributeProfile.GattPresentationFormatTypes")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGattPresentationFormatTypesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type GattProtocolError struct {
	RtClass
}

func NewIGattProtocolErrorStatics() *IGattProtocolErrorStatics {
	var p *IGattProtocolErrorStatics
	hs := NewHStr("Windows.Devices.Bluetooth.GenericAttributeProfile.GattProtocolError")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGattProtocolErrorStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type GattReadClientCharacteristicConfigurationDescriptorResult struct {
	RtClass
	*IGattReadClientCharacteristicConfigurationDescriptorResult
}

type GattReadRequest struct {
	RtClass
	*IGattReadRequest
}

type GattReadRequestedEventArgs struct {
	RtClass
	*IGattReadRequestedEventArgs
}

type GattReadResult struct {
	RtClass
	*IGattReadResult
}

type GattReliableWriteTransaction struct {
	RtClass
	*IGattReliableWriteTransaction
}

func NewGattReliableWriteTransaction() *GattReliableWriteTransaction {
	hs := NewHStr("Windows.Devices.Bluetooth.GenericAttributeProfile.GattReliableWriteTransaction")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &GattReliableWriteTransaction{
		RtClass:                       RtClass{PInspect: p},
		IGattReliableWriteTransaction: (*IGattReliableWriteTransaction)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type GattRequestStateChangedEventArgs struct {
	RtClass
	*IGattRequestStateChangedEventArgs
}

type GattServiceProvider struct {
	RtClass
	*IGattServiceProvider
}

func NewIGattServiceProviderStatics() *IGattServiceProviderStatics {
	var p *IGattServiceProviderStatics
	hs := NewHStr("Windows.Devices.Bluetooth.GenericAttributeProfile.GattServiceProvider")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGattServiceProviderStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type GattServiceProviderAdvertisementStatusChangedEventArgs struct {
	RtClass
	*IGattServiceProviderAdvertisementStatusChangedEventArgs
}

type GattServiceProviderAdvertisingParameters struct {
	RtClass
	*IGattServiceProviderAdvertisingParameters
}

func NewGattServiceProviderAdvertisingParameters() *GattServiceProviderAdvertisingParameters {
	hs := NewHStr("Windows.Devices.Bluetooth.GenericAttributeProfile.GattServiceProviderAdvertisingParameters")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &GattServiceProviderAdvertisingParameters{
		RtClass: RtClass{PInspect: p},
		IGattServiceProviderAdvertisingParameters: (*IGattServiceProviderAdvertisingParameters)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type GattServiceProviderResult struct {
	RtClass
	*IGattServiceProviderResult
}

type GattServiceUuids struct {
	RtClass
}

func NewIGattServiceUuidsStatics2() *IGattServiceUuidsStatics2 {
	var p *IGattServiceUuidsStatics2
	hs := NewHStr("Windows.Devices.Bluetooth.GenericAttributeProfile.GattServiceUuids")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGattServiceUuidsStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIGattServiceUuidsStatics() *IGattServiceUuidsStatics {
	var p *IGattServiceUuidsStatics
	hs := NewHStr("Windows.Devices.Bluetooth.GenericAttributeProfile.GattServiceUuids")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGattServiceUuidsStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type GattSession struct {
	RtClass
	*IGattSession
}

func NewIGattSessionStatics() *IGattSessionStatics {
	var p *IGattSessionStatics
	hs := NewHStr("Windows.Devices.Bluetooth.GenericAttributeProfile.GattSession")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGattSessionStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type GattSessionStatusChangedEventArgs struct {
	RtClass
	*IGattSessionStatusChangedEventArgs
}

type GattSubscribedClient struct {
	RtClass
	*IGattSubscribedClient
}

type GattValueChangedEventArgs struct {
	RtClass
	*IGattValueChangedEventArgs
}

type GattWriteRequest struct {
	RtClass
	*IGattWriteRequest
}

type GattWriteRequestedEventArgs struct {
	RtClass
	*IGattWriteRequestedEventArgs
}

type GattWriteResult struct {
	RtClass
	*IGattWriteResult
}
