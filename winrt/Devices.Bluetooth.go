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
type BluetoothAddressType int32

const (
	BluetoothAddressType_Public      BluetoothAddressType = 0
	BluetoothAddressType_Random      BluetoothAddressType = 1
	BluetoothAddressType_Unspecified BluetoothAddressType = 2
)

// enum
type BluetoothCacheMode int32

const (
	BluetoothCacheMode_Cached   BluetoothCacheMode = 0
	BluetoothCacheMode_Uncached BluetoothCacheMode = 1
)

// enum
type BluetoothConnectionStatus int32

const (
	BluetoothConnectionStatus_Disconnected BluetoothConnectionStatus = 0
	BluetoothConnectionStatus_Connected    BluetoothConnectionStatus = 1
)

// enum
type BluetoothError int32

const (
	BluetoothError_Success               BluetoothError = 0
	BluetoothError_RadioNotAvailable     BluetoothError = 1
	BluetoothError_ResourceInUse         BluetoothError = 2
	BluetoothError_DeviceNotConnected    BluetoothError = 3
	BluetoothError_OtherError            BluetoothError = 4
	BluetoothError_DisabledByPolicy      BluetoothError = 5
	BluetoothError_NotSupported          BluetoothError = 6
	BluetoothError_DisabledByUser        BluetoothError = 7
	BluetoothError_ConsentRequired       BluetoothError = 8
	BluetoothError_TransportNotSupported BluetoothError = 9
)

// enum
type BluetoothLEPreferredConnectionParametersRequestStatus int32

const (
	BluetoothLEPreferredConnectionParametersRequestStatus_Unspecified        BluetoothLEPreferredConnectionParametersRequestStatus = 0
	BluetoothLEPreferredConnectionParametersRequestStatus_Success            BluetoothLEPreferredConnectionParametersRequestStatus = 1
	BluetoothLEPreferredConnectionParametersRequestStatus_DeviceNotAvailable BluetoothLEPreferredConnectionParametersRequestStatus = 2
	BluetoothLEPreferredConnectionParametersRequestStatus_AccessDenied       BluetoothLEPreferredConnectionParametersRequestStatus = 3
)

// enum
type BluetoothMajorClass int32

const (
	BluetoothMajorClass_Miscellaneous      BluetoothMajorClass = 0
	BluetoothMajorClass_Computer           BluetoothMajorClass = 1
	BluetoothMajorClass_Phone              BluetoothMajorClass = 2
	BluetoothMajorClass_NetworkAccessPoint BluetoothMajorClass = 3
	BluetoothMajorClass_AudioVideo         BluetoothMajorClass = 4
	BluetoothMajorClass_Peripheral         BluetoothMajorClass = 5
	BluetoothMajorClass_Imaging            BluetoothMajorClass = 6
	BluetoothMajorClass_Wearable           BluetoothMajorClass = 7
	BluetoothMajorClass_Toy                BluetoothMajorClass = 8
	BluetoothMajorClass_Health             BluetoothMajorClass = 9
)

// enum
type BluetoothMinorClass int32

const (
	BluetoothMinorClass_Uncategorized                        BluetoothMinorClass = 0
	BluetoothMinorClass_ComputerDesktop                      BluetoothMinorClass = 1
	BluetoothMinorClass_ComputerServer                       BluetoothMinorClass = 2
	BluetoothMinorClass_ComputerLaptop                       BluetoothMinorClass = 3
	BluetoothMinorClass_ComputerHandheld                     BluetoothMinorClass = 4
	BluetoothMinorClass_ComputerPalmSize                     BluetoothMinorClass = 5
	BluetoothMinorClass_ComputerWearable                     BluetoothMinorClass = 6
	BluetoothMinorClass_ComputerTablet                       BluetoothMinorClass = 7
	BluetoothMinorClass_PhoneCellular                        BluetoothMinorClass = 1
	BluetoothMinorClass_PhoneCordless                        BluetoothMinorClass = 2
	BluetoothMinorClass_PhoneSmartPhone                      BluetoothMinorClass = 3
	BluetoothMinorClass_PhoneWired                           BluetoothMinorClass = 4
	BluetoothMinorClass_PhoneIsdn                            BluetoothMinorClass = 5
	BluetoothMinorClass_NetworkFullyAvailable                BluetoothMinorClass = 0
	BluetoothMinorClass_NetworkUsed01To17Percent             BluetoothMinorClass = 8
	BluetoothMinorClass_NetworkUsed17To33Percent             BluetoothMinorClass = 16
	BluetoothMinorClass_NetworkUsed33To50Percent             BluetoothMinorClass = 24
	BluetoothMinorClass_NetworkUsed50To67Percent             BluetoothMinorClass = 32
	BluetoothMinorClass_NetworkUsed67To83Percent             BluetoothMinorClass = 40
	BluetoothMinorClass_NetworkUsed83To99Percent             BluetoothMinorClass = 48
	BluetoothMinorClass_NetworkNoServiceAvailable            BluetoothMinorClass = 56
	BluetoothMinorClass_AudioVideoWearableHeadset            BluetoothMinorClass = 1
	BluetoothMinorClass_AudioVideoHandsFree                  BluetoothMinorClass = 2
	BluetoothMinorClass_AudioVideoMicrophone                 BluetoothMinorClass = 4
	BluetoothMinorClass_AudioVideoLoudspeaker                BluetoothMinorClass = 5
	BluetoothMinorClass_AudioVideoHeadphones                 BluetoothMinorClass = 6
	BluetoothMinorClass_AudioVideoPortableAudio              BluetoothMinorClass = 7
	BluetoothMinorClass_AudioVideoCarAudio                   BluetoothMinorClass = 8
	BluetoothMinorClass_AudioVideoSetTopBox                  BluetoothMinorClass = 9
	BluetoothMinorClass_AudioVideoHifiAudioDevice            BluetoothMinorClass = 10
	BluetoothMinorClass_AudioVideoVcr                        BluetoothMinorClass = 11
	BluetoothMinorClass_AudioVideoVideoCamera                BluetoothMinorClass = 12
	BluetoothMinorClass_AudioVideoCamcorder                  BluetoothMinorClass = 13
	BluetoothMinorClass_AudioVideoVideoMonitor               BluetoothMinorClass = 14
	BluetoothMinorClass_AudioVideoVideoDisplayAndLoudspeaker BluetoothMinorClass = 15
	BluetoothMinorClass_AudioVideoVideoConferencing          BluetoothMinorClass = 16
	BluetoothMinorClass_AudioVideoGamingOrToy                BluetoothMinorClass = 18
	BluetoothMinorClass_PeripheralJoystick                   BluetoothMinorClass = 1
	BluetoothMinorClass_PeripheralGamepad                    BluetoothMinorClass = 2
	BluetoothMinorClass_PeripheralRemoteControl              BluetoothMinorClass = 3
	BluetoothMinorClass_PeripheralSensing                    BluetoothMinorClass = 4
	BluetoothMinorClass_PeripheralDigitizerTablet            BluetoothMinorClass = 5
	BluetoothMinorClass_PeripheralCardReader                 BluetoothMinorClass = 6
	BluetoothMinorClass_PeripheralDigitalPen                 BluetoothMinorClass = 7
	BluetoothMinorClass_PeripheralHandheldScanner            BluetoothMinorClass = 8
	BluetoothMinorClass_PeripheralHandheldGesture            BluetoothMinorClass = 9
	BluetoothMinorClass_WearableWristwatch                   BluetoothMinorClass = 1
	BluetoothMinorClass_WearablePager                        BluetoothMinorClass = 2
	BluetoothMinorClass_WearableJacket                       BluetoothMinorClass = 3
	BluetoothMinorClass_WearableHelmet                       BluetoothMinorClass = 4
	BluetoothMinorClass_WearableGlasses                      BluetoothMinorClass = 5
	BluetoothMinorClass_ToyRobot                             BluetoothMinorClass = 1
	BluetoothMinorClass_ToyVehicle                           BluetoothMinorClass = 2
	BluetoothMinorClass_ToyDoll                              BluetoothMinorClass = 3
	BluetoothMinorClass_ToyController                        BluetoothMinorClass = 4
	BluetoothMinorClass_ToyGame                              BluetoothMinorClass = 5
	BluetoothMinorClass_HealthBloodPressureMonitor           BluetoothMinorClass = 1
	BluetoothMinorClass_HealthThermometer                    BluetoothMinorClass = 2
	BluetoothMinorClass_HealthWeighingScale                  BluetoothMinorClass = 3
	BluetoothMinorClass_HealthGlucoseMeter                   BluetoothMinorClass = 4
	BluetoothMinorClass_HealthPulseOximeter                  BluetoothMinorClass = 5
	BluetoothMinorClass_HealthHeartRateMonitor               BluetoothMinorClass = 6
	BluetoothMinorClass_HealthHealthDataDisplay              BluetoothMinorClass = 7
	BluetoothMinorClass_HealthStepCounter                    BluetoothMinorClass = 8
	BluetoothMinorClass_HealthBodyCompositionAnalyzer        BluetoothMinorClass = 9
	BluetoothMinorClass_HealthPeakFlowMonitor                BluetoothMinorClass = 10
	BluetoothMinorClass_HealthMedicationMonitor              BluetoothMinorClass = 11
	BluetoothMinorClass_HealthKneeProsthesis                 BluetoothMinorClass = 12
	BluetoothMinorClass_HealthAnkleProsthesis                BluetoothMinorClass = 13
	BluetoothMinorClass_HealthGenericHealthManager           BluetoothMinorClass = 14
	BluetoothMinorClass_HealthPersonalMobilityDevice         BluetoothMinorClass = 15
)

// enum
// flags
type BluetoothServiceCapabilities uint32

const (
	BluetoothServiceCapabilities_None                    BluetoothServiceCapabilities = 0
	BluetoothServiceCapabilities_LimitedDiscoverableMode BluetoothServiceCapabilities = 1
	BluetoothServiceCapabilities_PositioningService      BluetoothServiceCapabilities = 8
	BluetoothServiceCapabilities_NetworkingService       BluetoothServiceCapabilities = 16
	BluetoothServiceCapabilities_RenderingService        BluetoothServiceCapabilities = 32
	BluetoothServiceCapabilities_CapturingService        BluetoothServiceCapabilities = 64
	BluetoothServiceCapabilities_ObjectTransferService   BluetoothServiceCapabilities = 128
	BluetoothServiceCapabilities_AudioService            BluetoothServiceCapabilities = 256
	BluetoothServiceCapabilities_TelephoneService        BluetoothServiceCapabilities = 512
	BluetoothServiceCapabilities_InformationService      BluetoothServiceCapabilities = 1024
)

// interfaces

// 7974F04C-5F7A-4A34-9225-A855F84B1A8B
var IID_IBluetoothAdapter = syscall.GUID{0x7974F04C, 0x5F7A, 0x4A34,
	[8]byte{0x92, 0x25, 0xA8, 0x55, 0xF8, 0x4B, 0x1A, 0x8B}}

type IBluetoothAdapterInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_BluetoothAddress() uint64
	Get_IsClassicSupported() bool
	Get_IsLowEnergySupported() bool
	Get_IsPeripheralRoleSupported() bool
	Get_IsCentralRoleSupported() bool
	Get_IsAdvertisementOffloadSupported() bool
	GetRadioAsync() *IAsyncOperation[*IRadio]
}

type IBluetoothAdapterVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId                        uintptr
	Get_BluetoothAddress                uintptr
	Get_IsClassicSupported              uintptr
	Get_IsLowEnergySupported            uintptr
	Get_IsPeripheralRoleSupported       uintptr
	Get_IsCentralRoleSupported          uintptr
	Get_IsAdvertisementOffloadSupported uintptr
	GetRadioAsync                       uintptr
}

type IBluetoothAdapter struct {
	win32.IInspectable
}

func (this *IBluetoothAdapter) Vtbl() *IBluetoothAdapterVtbl {
	return (*IBluetoothAdapterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothAdapter) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBluetoothAdapter) Get_BluetoothAddress() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BluetoothAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothAdapter) Get_IsClassicSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsClassicSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothAdapter) Get_IsLowEnergySupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsLowEnergySupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothAdapter) Get_IsPeripheralRoleSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPeripheralRoleSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothAdapter) Get_IsCentralRoleSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCentralRoleSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothAdapter) Get_IsAdvertisementOffloadSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAdvertisementOffloadSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothAdapter) GetRadioAsync() *IAsyncOperation[*IRadio] {
	var _result *IAsyncOperation[*IRadio]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetRadioAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AC94CECC-24D5-41B3-916D-1097C50B102B
var IID_IBluetoothAdapter2 = syscall.GUID{0xAC94CECC, 0x24D5, 0x41B3,
	[8]byte{0x91, 0x6D, 0x10, 0x97, 0xC5, 0x0B, 0x10, 0x2B}}

type IBluetoothAdapter2Interface interface {
	win32.IInspectableInterface
	Get_AreClassicSecureConnectionsSupported() bool
	Get_AreLowEnergySecureConnectionsSupported() bool
}

type IBluetoothAdapter2Vtbl struct {
	win32.IInspectableVtbl
	Get_AreClassicSecureConnectionsSupported   uintptr
	Get_AreLowEnergySecureConnectionsSupported uintptr
}

type IBluetoothAdapter2 struct {
	win32.IInspectable
}

func (this *IBluetoothAdapter2) Vtbl() *IBluetoothAdapter2Vtbl {
	return (*IBluetoothAdapter2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothAdapter2) Get_AreClassicSecureConnectionsSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AreClassicSecureConnectionsSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothAdapter2) Get_AreLowEnergySecureConnectionsSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AreLowEnergySecureConnectionsSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 8F8624E0-CBA9-5211-9F89-3AAC62B4C6B8
var IID_IBluetoothAdapter3 = syscall.GUID{0x8F8624E0, 0xCBA9, 0x5211,
	[8]byte{0x9F, 0x89, 0x3A, 0xAC, 0x62, 0xB4, 0xC6, 0xB8}}

type IBluetoothAdapter3Interface interface {
	win32.IInspectableInterface
	Get_IsExtendedAdvertisingSupported() bool
	Get_MaxAdvertisementDataLength() uint32
}

type IBluetoothAdapter3Vtbl struct {
	win32.IInspectableVtbl
	Get_IsExtendedAdvertisingSupported uintptr
	Get_MaxAdvertisementDataLength     uintptr
}

type IBluetoothAdapter3 struct {
	win32.IInspectable
}

func (this *IBluetoothAdapter3) Vtbl() *IBluetoothAdapter3Vtbl {
	return (*IBluetoothAdapter3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothAdapter3) Get_IsExtendedAdvertisingSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsExtendedAdvertisingSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothAdapter3) Get_MaxAdvertisementDataLength() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxAdvertisementDataLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 8B02FB6A-AC4C-4741-8661-8EAB7D17EA9F
var IID_IBluetoothAdapterStatics = syscall.GUID{0x8B02FB6A, 0xAC4C, 0x4741,
	[8]byte{0x86, 0x61, 0x8E, 0xAB, 0x7D, 0x17, 0xEA, 0x9F}}

type IBluetoothAdapterStaticsInterface interface {
	win32.IInspectableInterface
	GetDeviceSelector() string
	FromIdAsync(deviceId string) *IAsyncOperation[*IBluetoothAdapter]
	GetDefaultAsync() *IAsyncOperation[*IBluetoothAdapter]
}

type IBluetoothAdapterStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector uintptr
	FromIdAsync       uintptr
	GetDefaultAsync   uintptr
}

type IBluetoothAdapterStatics struct {
	win32.IInspectable
}

func (this *IBluetoothAdapterStatics) Vtbl() *IBluetoothAdapterStaticsVtbl {
	return (*IBluetoothAdapterStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothAdapterStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBluetoothAdapterStatics) FromIdAsync(deviceId string) *IAsyncOperation[*IBluetoothAdapter] {
	var _result *IAsyncOperation[*IBluetoothAdapter]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothAdapterStatics) GetDefaultAsync() *IAsyncOperation[*IBluetoothAdapter] {
	var _result *IAsyncOperation[*IBluetoothAdapter]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D640227E-D7D7-4661-9454-65039CA17A2B
var IID_IBluetoothClassOfDevice = syscall.GUID{0xD640227E, 0xD7D7, 0x4661,
	[8]byte{0x94, 0x54, 0x65, 0x03, 0x9C, 0xA1, 0x7A, 0x2B}}

type IBluetoothClassOfDeviceInterface interface {
	win32.IInspectableInterface
	Get_RawValue() uint32
	Get_MajorClass() BluetoothMajorClass
	Get_MinorClass() BluetoothMinorClass
	Get_ServiceCapabilities() BluetoothServiceCapabilities
}

type IBluetoothClassOfDeviceVtbl struct {
	win32.IInspectableVtbl
	Get_RawValue            uintptr
	Get_MajorClass          uintptr
	Get_MinorClass          uintptr
	Get_ServiceCapabilities uintptr
}

type IBluetoothClassOfDevice struct {
	win32.IInspectable
}

func (this *IBluetoothClassOfDevice) Vtbl() *IBluetoothClassOfDeviceVtbl {
	return (*IBluetoothClassOfDeviceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothClassOfDevice) Get_RawValue() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RawValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothClassOfDevice) Get_MajorClass() BluetoothMajorClass {
	var _result BluetoothMajorClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MajorClass, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothClassOfDevice) Get_MinorClass() BluetoothMinorClass {
	var _result BluetoothMinorClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinorClass, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothClassOfDevice) Get_ServiceCapabilities() BluetoothServiceCapabilities {
	var _result BluetoothServiceCapabilities
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceCapabilities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// E46135BD-0FA2-416C-91B4-C1E48CA061C1
var IID_IBluetoothClassOfDeviceStatics = syscall.GUID{0xE46135BD, 0x0FA2, 0x416C,
	[8]byte{0x91, 0xB4, 0xC1, 0xE4, 0x8C, 0xA0, 0x61, 0xC1}}

type IBluetoothClassOfDeviceStaticsInterface interface {
	win32.IInspectableInterface
	FromRawValue(rawValue uint32) *IBluetoothClassOfDevice
	FromParts(majorClass BluetoothMajorClass, minorClass BluetoothMinorClass, serviceCapabilities BluetoothServiceCapabilities) *IBluetoothClassOfDevice
}

type IBluetoothClassOfDeviceStaticsVtbl struct {
	win32.IInspectableVtbl
	FromRawValue uintptr
	FromParts    uintptr
}

type IBluetoothClassOfDeviceStatics struct {
	win32.IInspectable
}

func (this *IBluetoothClassOfDeviceStatics) Vtbl() *IBluetoothClassOfDeviceStaticsVtbl {
	return (*IBluetoothClassOfDeviceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothClassOfDeviceStatics) FromRawValue(rawValue uint32) *IBluetoothClassOfDevice {
	var _result *IBluetoothClassOfDevice
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromRawValue, uintptr(unsafe.Pointer(this)), uintptr(rawValue), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothClassOfDeviceStatics) FromParts(majorClass BluetoothMajorClass, minorClass BluetoothMinorClass, serviceCapabilities BluetoothServiceCapabilities) *IBluetoothClassOfDevice {
	var _result *IBluetoothClassOfDevice
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromParts, uintptr(unsafe.Pointer(this)), uintptr(majorClass), uintptr(minorClass), uintptr(serviceCapabilities), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2335B156-90D2-4A04-AEF5-0E20B9E6B707
var IID_IBluetoothDevice = syscall.GUID{0x2335B156, 0x90D2, 0x4A04,
	[8]byte{0xAE, 0xF5, 0x0E, 0x20, 0xB9, 0xE6, 0xB7, 0x07}}

type IBluetoothDeviceInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_HostName() *IHostName
	Get_Name() string
	Get_ClassOfDevice() *IBluetoothClassOfDevice
	Get_SdpRecords() *IVectorView[*IBuffer]
	Get_RfcommServices() *IVectorView[*IRfcommDeviceService]
	Get_ConnectionStatus() BluetoothConnectionStatus
	Get_BluetoothAddress() uint64
	Add_NameChanged(handler TypedEventHandler[*IBluetoothDevice, interface{}]) EventRegistrationToken
	Remove_NameChanged(token EventRegistrationToken)
	Add_SdpRecordsChanged(handler TypedEventHandler[*IBluetoothDevice, interface{}]) EventRegistrationToken
	Remove_SdpRecordsChanged(token EventRegistrationToken)
	Add_ConnectionStatusChanged(handler TypedEventHandler[*IBluetoothDevice, interface{}]) EventRegistrationToken
	Remove_ConnectionStatusChanged(token EventRegistrationToken)
}

type IBluetoothDeviceVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId                   uintptr
	Get_HostName                   uintptr
	Get_Name                       uintptr
	Get_ClassOfDevice              uintptr
	Get_SdpRecords                 uintptr
	Get_RfcommServices             uintptr
	Get_ConnectionStatus           uintptr
	Get_BluetoothAddress           uintptr
	Add_NameChanged                uintptr
	Remove_NameChanged             uintptr
	Add_SdpRecordsChanged          uintptr
	Remove_SdpRecordsChanged       uintptr
	Add_ConnectionStatusChanged    uintptr
	Remove_ConnectionStatusChanged uintptr
}

type IBluetoothDevice struct {
	win32.IInspectable
}

func (this *IBluetoothDevice) Vtbl() *IBluetoothDeviceVtbl {
	return (*IBluetoothDeviceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothDevice) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBluetoothDevice) Get_HostName() *IHostName {
	var _result *IHostName
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HostName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothDevice) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBluetoothDevice) Get_ClassOfDevice() *IBluetoothClassOfDevice {
	var _result *IBluetoothClassOfDevice
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ClassOfDevice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothDevice) Get_SdpRecords() *IVectorView[*IBuffer] {
	var _result *IVectorView[*IBuffer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SdpRecords, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothDevice) Get_RfcommServices() *IVectorView[*IRfcommDeviceService] {
	var _result *IVectorView[*IRfcommDeviceService]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RfcommServices, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothDevice) Get_ConnectionStatus() BluetoothConnectionStatus {
	var _result BluetoothConnectionStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConnectionStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothDevice) Get_BluetoothAddress() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BluetoothAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothDevice) Add_NameChanged(handler TypedEventHandler[*IBluetoothDevice, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_NameChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothDevice) Remove_NameChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_NameChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IBluetoothDevice) Add_SdpRecordsChanged(handler TypedEventHandler[*IBluetoothDevice, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SdpRecordsChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothDevice) Remove_SdpRecordsChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SdpRecordsChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IBluetoothDevice) Add_ConnectionStatusChanged(handler TypedEventHandler[*IBluetoothDevice, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ConnectionStatusChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothDevice) Remove_ConnectionStatusChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ConnectionStatusChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 0133F954-B156-4DD0-B1F5-C11BC31A5163
var IID_IBluetoothDevice2 = syscall.GUID{0x0133F954, 0xB156, 0x4DD0,
	[8]byte{0xB1, 0xF5, 0xC1, 0x1B, 0xC3, 0x1A, 0x51, 0x63}}

type IBluetoothDevice2Interface interface {
	win32.IInspectableInterface
	Get_DeviceInformation() *IDeviceInformation
}

type IBluetoothDevice2Vtbl struct {
	win32.IInspectableVtbl
	Get_DeviceInformation uintptr
}

type IBluetoothDevice2 struct {
	win32.IInspectable
}

func (this *IBluetoothDevice2) Vtbl() *IBluetoothDevice2Vtbl {
	return (*IBluetoothDevice2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothDevice2) Get_DeviceInformation() *IDeviceInformation {
	var _result *IDeviceInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 57FFF78B-651A-4454-B90F-EB21EF0B0D71
var IID_IBluetoothDevice3 = syscall.GUID{0x57FFF78B, 0x651A, 0x4454,
	[8]byte{0xB9, 0x0F, 0xEB, 0x21, 0xEF, 0x0B, 0x0D, 0x71}}

type IBluetoothDevice3Interface interface {
	win32.IInspectableInterface
	Get_DeviceAccessInformation() *IDeviceAccessInformation
	RequestAccessAsync() *IAsyncOperation[DeviceAccessStatus]
	GetRfcommServicesAsync() *IAsyncOperation[*IRfcommDeviceServicesResult]
	GetRfcommServicesWithCacheModeAsync(cacheMode BluetoothCacheMode) *IAsyncOperation[*IRfcommDeviceServicesResult]
	GetRfcommServicesForIdAsync(serviceId *IRfcommServiceId) *IAsyncOperation[*IRfcommDeviceServicesResult]
	GetRfcommServicesForIdWithCacheModeAsync(serviceId *IRfcommServiceId, cacheMode BluetoothCacheMode) *IAsyncOperation[*IRfcommDeviceServicesResult]
}

type IBluetoothDevice3Vtbl struct {
	win32.IInspectableVtbl
	Get_DeviceAccessInformation              uintptr
	RequestAccessAsync                       uintptr
	GetRfcommServicesAsync                   uintptr
	GetRfcommServicesWithCacheModeAsync      uintptr
	GetRfcommServicesForIdAsync              uintptr
	GetRfcommServicesForIdWithCacheModeAsync uintptr
}

type IBluetoothDevice3 struct {
	win32.IInspectable
}

func (this *IBluetoothDevice3) Vtbl() *IBluetoothDevice3Vtbl {
	return (*IBluetoothDevice3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothDevice3) Get_DeviceAccessInformation() *IDeviceAccessInformation {
	var _result *IDeviceAccessInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceAccessInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothDevice3) RequestAccessAsync() *IAsyncOperation[DeviceAccessStatus] {
	var _result *IAsyncOperation[DeviceAccessStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothDevice3) GetRfcommServicesAsync() *IAsyncOperation[*IRfcommDeviceServicesResult] {
	var _result *IAsyncOperation[*IRfcommDeviceServicesResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetRfcommServicesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothDevice3) GetRfcommServicesWithCacheModeAsync(cacheMode BluetoothCacheMode) *IAsyncOperation[*IRfcommDeviceServicesResult] {
	var _result *IAsyncOperation[*IRfcommDeviceServicesResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetRfcommServicesWithCacheModeAsync, uintptr(unsafe.Pointer(this)), uintptr(cacheMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothDevice3) GetRfcommServicesForIdAsync(serviceId *IRfcommServiceId) *IAsyncOperation[*IRfcommDeviceServicesResult] {
	var _result *IAsyncOperation[*IRfcommDeviceServicesResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetRfcommServicesForIdAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(serviceId)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothDevice3) GetRfcommServicesForIdWithCacheModeAsync(serviceId *IRfcommServiceId, cacheMode BluetoothCacheMode) *IAsyncOperation[*IRfcommDeviceServicesResult] {
	var _result *IAsyncOperation[*IRfcommDeviceServicesResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetRfcommServicesForIdWithCacheModeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(serviceId)), uintptr(cacheMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 817C34AD-0E9C-42B2-A8DC-3E8094940D12
var IID_IBluetoothDevice4 = syscall.GUID{0x817C34AD, 0x0E9C, 0x42B2,
	[8]byte{0xA8, 0xDC, 0x3E, 0x80, 0x94, 0x94, 0x0D, 0x12}}

type IBluetoothDevice4Interface interface {
	win32.IInspectableInterface
	Get_BluetoothDeviceId() *IBluetoothDeviceId
}

type IBluetoothDevice4Vtbl struct {
	win32.IInspectableVtbl
	Get_BluetoothDeviceId uintptr
}

type IBluetoothDevice4 struct {
	win32.IInspectable
}

func (this *IBluetoothDevice4) Vtbl() *IBluetoothDevice4Vtbl {
	return (*IBluetoothDevice4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothDevice4) Get_BluetoothDeviceId() *IBluetoothDeviceId {
	var _result *IBluetoothDeviceId
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BluetoothDeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B5E0B385-5E85-4559-A10D-1C7281379F96
var IID_IBluetoothDevice5 = syscall.GUID{0xB5E0B385, 0x5E85, 0x4559,
	[8]byte{0xA1, 0x0D, 0x1C, 0x72, 0x81, 0x37, 0x9F, 0x96}}

type IBluetoothDevice5Interface interface {
	win32.IInspectableInterface
	Get_WasSecureConnectionUsedForPairing() bool
}

type IBluetoothDevice5Vtbl struct {
	win32.IInspectableVtbl
	Get_WasSecureConnectionUsedForPairing uintptr
}

type IBluetoothDevice5 struct {
	win32.IInspectable
}

func (this *IBluetoothDevice5) Vtbl() *IBluetoothDevice5Vtbl {
	return (*IBluetoothDevice5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothDevice5) Get_WasSecureConnectionUsedForPairing() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WasSecureConnectionUsedForPairing, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C17949AF-57C1-4642-BCCE-E6C06B20AE76
var IID_IBluetoothDeviceId = syscall.GUID{0xC17949AF, 0x57C1, 0x4642,
	[8]byte{0xBC, 0xCE, 0xE6, 0xC0, 0x6B, 0x20, 0xAE, 0x76}}

type IBluetoothDeviceIdInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Get_IsClassicDevice() bool
	Get_IsLowEnergyDevice() bool
}

type IBluetoothDeviceIdVtbl struct {
	win32.IInspectableVtbl
	Get_Id                uintptr
	Get_IsClassicDevice   uintptr
	Get_IsLowEnergyDevice uintptr
}

type IBluetoothDeviceId struct {
	win32.IInspectable
}

func (this *IBluetoothDeviceId) Vtbl() *IBluetoothDeviceIdVtbl {
	return (*IBluetoothDeviceIdVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothDeviceId) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBluetoothDeviceId) Get_IsClassicDevice() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsClassicDevice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothDeviceId) Get_IsLowEnergyDevice() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsLowEnergyDevice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A7884E67-3EFB-4F31-BBC2-810E09977404
var IID_IBluetoothDeviceIdStatics = syscall.GUID{0xA7884E67, 0x3EFB, 0x4F31,
	[8]byte{0xBB, 0xC2, 0x81, 0x0E, 0x09, 0x97, 0x74, 0x04}}

type IBluetoothDeviceIdStaticsInterface interface {
	win32.IInspectableInterface
	FromId(deviceId string) *IBluetoothDeviceId
}

type IBluetoothDeviceIdStaticsVtbl struct {
	win32.IInspectableVtbl
	FromId uintptr
}

type IBluetoothDeviceIdStatics struct {
	win32.IInspectable
}

func (this *IBluetoothDeviceIdStatics) Vtbl() *IBluetoothDeviceIdStaticsVtbl {
	return (*IBluetoothDeviceIdStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothDeviceIdStatics) FromId(deviceId string) *IBluetoothDeviceId {
	var _result *IBluetoothDeviceId
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromId, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0991DF51-57DB-4725-BBD7-84F64327EC2C
var IID_IBluetoothDeviceStatics = syscall.GUID{0x0991DF51, 0x57DB, 0x4725,
	[8]byte{0xBB, 0xD7, 0x84, 0xF6, 0x43, 0x27, 0xEC, 0x2C}}

type IBluetoothDeviceStaticsInterface interface {
	win32.IInspectableInterface
	FromIdAsync(deviceId string) *IAsyncOperation[*IBluetoothDevice]
	FromHostNameAsync(hostName *IHostName) *IAsyncOperation[*IBluetoothDevice]
	FromBluetoothAddressAsync(address uint64) *IAsyncOperation[*IBluetoothDevice]
	GetDeviceSelector() string
}

type IBluetoothDeviceStaticsVtbl struct {
	win32.IInspectableVtbl
	FromIdAsync               uintptr
	FromHostNameAsync         uintptr
	FromBluetoothAddressAsync uintptr
	GetDeviceSelector         uintptr
}

type IBluetoothDeviceStatics struct {
	win32.IInspectable
}

func (this *IBluetoothDeviceStatics) Vtbl() *IBluetoothDeviceStaticsVtbl {
	return (*IBluetoothDeviceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothDeviceStatics) FromIdAsync(deviceId string) *IAsyncOperation[*IBluetoothDevice] {
	var _result *IAsyncOperation[*IBluetoothDevice]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothDeviceStatics) FromHostNameAsync(hostName *IHostName) *IAsyncOperation[*IBluetoothDevice] {
	var _result *IAsyncOperation[*IBluetoothDevice]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromHostNameAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(hostName)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothDeviceStatics) FromBluetoothAddressAsync(address uint64) *IAsyncOperation[*IBluetoothDevice] {
	var _result *IAsyncOperation[*IBluetoothDevice]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromBluetoothAddressAsync, uintptr(unsafe.Pointer(this)), uintptr(address), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothDeviceStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// C29E8E2F-4E14-4477-AA1B-B8B47E5B7ECE
var IID_IBluetoothDeviceStatics2 = syscall.GUID{0xC29E8E2F, 0x4E14, 0x4477,
	[8]byte{0xAA, 0x1B, 0xB8, 0xB4, 0x7E, 0x5B, 0x7E, 0xCE}}

type IBluetoothDeviceStatics2Interface interface {
	win32.IInspectableInterface
	GetDeviceSelectorFromPairingState(pairingState bool) string
	GetDeviceSelectorFromConnectionStatus(connectionStatus BluetoothConnectionStatus) string
	GetDeviceSelectorFromDeviceName(deviceName string) string
	GetDeviceSelectorFromBluetoothAddress(bluetoothAddress uint64) string
	GetDeviceSelectorFromClassOfDevice(classOfDevice *IBluetoothClassOfDevice) string
}

type IBluetoothDeviceStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelectorFromPairingState     uintptr
	GetDeviceSelectorFromConnectionStatus uintptr
	GetDeviceSelectorFromDeviceName       uintptr
	GetDeviceSelectorFromBluetoothAddress uintptr
	GetDeviceSelectorFromClassOfDevice    uintptr
}

type IBluetoothDeviceStatics2 struct {
	win32.IInspectable
}

func (this *IBluetoothDeviceStatics2) Vtbl() *IBluetoothDeviceStatics2Vtbl {
	return (*IBluetoothDeviceStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothDeviceStatics2) GetDeviceSelectorFromPairingState(pairingState bool) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorFromPairingState, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&pairingState))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBluetoothDeviceStatics2) GetDeviceSelectorFromConnectionStatus(connectionStatus BluetoothConnectionStatus) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorFromConnectionStatus, uintptr(unsafe.Pointer(this)), uintptr(connectionStatus), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBluetoothDeviceStatics2) GetDeviceSelectorFromDeviceName(deviceName string) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorFromDeviceName, uintptr(unsafe.Pointer(this)), NewHStr(deviceName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBluetoothDeviceStatics2) GetDeviceSelectorFromBluetoothAddress(bluetoothAddress uint64) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorFromBluetoothAddress, uintptr(unsafe.Pointer(this)), uintptr(bluetoothAddress), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBluetoothDeviceStatics2) GetDeviceSelectorFromClassOfDevice(classOfDevice *IBluetoothClassOfDevice) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorFromClassOfDevice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(classOfDevice)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 5D2079F2-66A8-4258-985E-02B4D9509F18
var IID_IBluetoothLEAppearance = syscall.GUID{0x5D2079F2, 0x66A8, 0x4258,
	[8]byte{0x98, 0x5E, 0x02, 0xB4, 0xD9, 0x50, 0x9F, 0x18}}

type IBluetoothLEAppearanceInterface interface {
	win32.IInspectableInterface
	Get_RawValue() uint16
	Get_Category() uint16
	Get_SubCategory() uint16
}

type IBluetoothLEAppearanceVtbl struct {
	win32.IInspectableVtbl
	Get_RawValue    uintptr
	Get_Category    uintptr
	Get_SubCategory uintptr
}

type IBluetoothLEAppearance struct {
	win32.IInspectable
}

func (this *IBluetoothLEAppearance) Vtbl() *IBluetoothLEAppearanceVtbl {
	return (*IBluetoothLEAppearanceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAppearance) Get_RawValue() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RawValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearance) Get_Category() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Category, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearance) Get_SubCategory() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SubCategory, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6D4D54FE-046A-4185-AAB6-824CF0610861
var IID_IBluetoothLEAppearanceCategoriesStatics = syscall.GUID{0x6D4D54FE, 0x046A, 0x4185,
	[8]byte{0xAA, 0xB6, 0x82, 0x4C, 0xF0, 0x61, 0x08, 0x61}}

type IBluetoothLEAppearanceCategoriesStaticsInterface interface {
	win32.IInspectableInterface
	Get_Uncategorized() uint16
	Get_Phone() uint16
	Get_Computer() uint16
	Get_Watch() uint16
	Get_Clock() uint16
	Get_Display() uint16
	Get_RemoteControl() uint16
	Get_EyeGlasses() uint16
	Get_Tag() uint16
	Get_Keyring() uint16
	Get_MediaPlayer() uint16
	Get_BarcodeScanner() uint16
	Get_Thermometer() uint16
	Get_HeartRate() uint16
	Get_BloodPressure() uint16
	Get_HumanInterfaceDevice() uint16
	Get_GlucoseMeter() uint16
	Get_RunningWalking() uint16
	Get_Cycling() uint16
	Get_PulseOximeter() uint16
	Get_WeightScale() uint16
	Get_OutdoorSportActivity() uint16
}

type IBluetoothLEAppearanceCategoriesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Uncategorized        uintptr
	Get_Phone                uintptr
	Get_Computer             uintptr
	Get_Watch                uintptr
	Get_Clock                uintptr
	Get_Display              uintptr
	Get_RemoteControl        uintptr
	Get_EyeGlasses           uintptr
	Get_Tag                  uintptr
	Get_Keyring              uintptr
	Get_MediaPlayer          uintptr
	Get_BarcodeScanner       uintptr
	Get_Thermometer          uintptr
	Get_HeartRate            uintptr
	Get_BloodPressure        uintptr
	Get_HumanInterfaceDevice uintptr
	Get_GlucoseMeter         uintptr
	Get_RunningWalking       uintptr
	Get_Cycling              uintptr
	Get_PulseOximeter        uintptr
	Get_WeightScale          uintptr
	Get_OutdoorSportActivity uintptr
}

type IBluetoothLEAppearanceCategoriesStatics struct {
	win32.IInspectable
}

func (this *IBluetoothLEAppearanceCategoriesStatics) Vtbl() *IBluetoothLEAppearanceCategoriesStaticsVtbl {
	return (*IBluetoothLEAppearanceCategoriesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAppearanceCategoriesStatics) Get_Uncategorized() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Uncategorized, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceCategoriesStatics) Get_Phone() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Phone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceCategoriesStatics) Get_Computer() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Computer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceCategoriesStatics) Get_Watch() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Watch, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceCategoriesStatics) Get_Clock() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Clock, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceCategoriesStatics) Get_Display() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Display, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceCategoriesStatics) Get_RemoteControl() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemoteControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceCategoriesStatics) Get_EyeGlasses() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EyeGlasses, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceCategoriesStatics) Get_Tag() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Tag, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceCategoriesStatics) Get_Keyring() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Keyring, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceCategoriesStatics) Get_MediaPlayer() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaPlayer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceCategoriesStatics) Get_BarcodeScanner() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BarcodeScanner, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceCategoriesStatics) Get_Thermometer() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Thermometer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceCategoriesStatics) Get_HeartRate() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HeartRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceCategoriesStatics) Get_BloodPressure() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BloodPressure, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceCategoriesStatics) Get_HumanInterfaceDevice() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HumanInterfaceDevice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceCategoriesStatics) Get_GlucoseMeter() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GlucoseMeter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceCategoriesStatics) Get_RunningWalking() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RunningWalking, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceCategoriesStatics) Get_Cycling() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Cycling, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceCategoriesStatics) Get_PulseOximeter() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PulseOximeter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceCategoriesStatics) Get_WeightScale() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WeightScale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceCategoriesStatics) Get_OutdoorSportActivity() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutdoorSportActivity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A193C0C7-4504-4F4A-9BA5-CD1054E5E065
var IID_IBluetoothLEAppearanceStatics = syscall.GUID{0xA193C0C7, 0x4504, 0x4F4A,
	[8]byte{0x9B, 0xA5, 0xCD, 0x10, 0x54, 0xE5, 0xE0, 0x65}}

type IBluetoothLEAppearanceStaticsInterface interface {
	win32.IInspectableInterface
	FromRawValue(rawValue uint16) *IBluetoothLEAppearance
	FromParts(appearanceCategory uint16, appearanceSubCategory uint16) *IBluetoothLEAppearance
}

type IBluetoothLEAppearanceStaticsVtbl struct {
	win32.IInspectableVtbl
	FromRawValue uintptr
	FromParts    uintptr
}

type IBluetoothLEAppearanceStatics struct {
	win32.IInspectable
}

func (this *IBluetoothLEAppearanceStatics) Vtbl() *IBluetoothLEAppearanceStaticsVtbl {
	return (*IBluetoothLEAppearanceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAppearanceStatics) FromRawValue(rawValue uint16) *IBluetoothLEAppearance {
	var _result *IBluetoothLEAppearance
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromRawValue, uintptr(unsafe.Pointer(this)), uintptr(rawValue), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEAppearanceStatics) FromParts(appearanceCategory uint16, appearanceSubCategory uint16) *IBluetoothLEAppearance {
	var _result *IBluetoothLEAppearance
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromParts, uintptr(unsafe.Pointer(this)), uintptr(appearanceCategory), uintptr(appearanceSubCategory), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E57BA606-2144-415A-8312-71CCF291F8D1
var IID_IBluetoothLEAppearanceSubcategoriesStatics = syscall.GUID{0xE57BA606, 0x2144, 0x415A,
	[8]byte{0x83, 0x12, 0x71, 0xCC, 0xF2, 0x91, 0xF8, 0xD1}}

type IBluetoothLEAppearanceSubcategoriesStaticsInterface interface {
	win32.IInspectableInterface
	Get_Generic() uint16
	Get_SportsWatch() uint16
	Get_ThermometerEar() uint16
	Get_HeartRateBelt() uint16
	Get_BloodPressureArm() uint16
	Get_BloodPressureWrist() uint16
	Get_Keyboard() uint16
	Get_Mouse() uint16
	Get_Joystick() uint16
	Get_Gamepad() uint16
	Get_DigitizerTablet() uint16
	Get_CardReader() uint16
	Get_DigitalPen() uint16
	Get_BarcodeScanner() uint16
	Get_RunningWalkingInShoe() uint16
	Get_RunningWalkingOnShoe() uint16
	Get_RunningWalkingOnHip() uint16
	Get_CyclingComputer() uint16
	Get_CyclingSpeedSensor() uint16
	Get_CyclingCadenceSensor() uint16
	Get_CyclingPowerSensor() uint16
	Get_CyclingSpeedCadenceSensor() uint16
	Get_OximeterFingertip() uint16
	Get_OximeterWristWorn() uint16
	Get_LocationDisplay() uint16
	Get_LocationNavigationDisplay() uint16
	Get_LocationPod() uint16
	Get_LocationNavigationPod() uint16
}

type IBluetoothLEAppearanceSubcategoriesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Generic                   uintptr
	Get_SportsWatch               uintptr
	Get_ThermometerEar            uintptr
	Get_HeartRateBelt             uintptr
	Get_BloodPressureArm          uintptr
	Get_BloodPressureWrist        uintptr
	Get_Keyboard                  uintptr
	Get_Mouse                     uintptr
	Get_Joystick                  uintptr
	Get_Gamepad                   uintptr
	Get_DigitizerTablet           uintptr
	Get_CardReader                uintptr
	Get_DigitalPen                uintptr
	Get_BarcodeScanner            uintptr
	Get_RunningWalkingInShoe      uintptr
	Get_RunningWalkingOnShoe      uintptr
	Get_RunningWalkingOnHip       uintptr
	Get_CyclingComputer           uintptr
	Get_CyclingSpeedSensor        uintptr
	Get_CyclingCadenceSensor      uintptr
	Get_CyclingPowerSensor        uintptr
	Get_CyclingSpeedCadenceSensor uintptr
	Get_OximeterFingertip         uintptr
	Get_OximeterWristWorn         uintptr
	Get_LocationDisplay           uintptr
	Get_LocationNavigationDisplay uintptr
	Get_LocationPod               uintptr
	Get_LocationNavigationPod     uintptr
}

type IBluetoothLEAppearanceSubcategoriesStatics struct {
	win32.IInspectable
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Vtbl() *IBluetoothLEAppearanceSubcategoriesStaticsVtbl {
	return (*IBluetoothLEAppearanceSubcategoriesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_Generic() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Generic, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_SportsWatch() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SportsWatch, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_ThermometerEar() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ThermometerEar, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_HeartRateBelt() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HeartRateBelt, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_BloodPressureArm() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BloodPressureArm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_BloodPressureWrist() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BloodPressureWrist, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_Keyboard() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Keyboard, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_Mouse() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mouse, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_Joystick() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Joystick, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_Gamepad() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Gamepad, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_DigitizerTablet() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DigitizerTablet, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_CardReader() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CardReader, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_DigitalPen() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DigitalPen, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_BarcodeScanner() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BarcodeScanner, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_RunningWalkingInShoe() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RunningWalkingInShoe, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_RunningWalkingOnShoe() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RunningWalkingOnShoe, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_RunningWalkingOnHip() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RunningWalkingOnHip, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_CyclingComputer() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CyclingComputer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_CyclingSpeedSensor() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CyclingSpeedSensor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_CyclingCadenceSensor() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CyclingCadenceSensor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_CyclingPowerSensor() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CyclingPowerSensor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_CyclingSpeedCadenceSensor() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CyclingSpeedCadenceSensor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_OximeterFingertip() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OximeterFingertip, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_OximeterWristWorn() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OximeterWristWorn, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_LocationDisplay() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocationDisplay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_LocationNavigationDisplay() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocationNavigationDisplay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_LocationPod() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocationPod, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAppearanceSubcategoriesStatics) Get_LocationNavigationPod() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocationNavigationPod, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 33CB0771-8DA9-508F-A366-1CA388C929AB
var IID_IBluetoothLEConnectionParameters = syscall.GUID{0x33CB0771, 0x8DA9, 0x508F,
	[8]byte{0xA3, 0x66, 0x1C, 0xA3, 0x88, 0xC9, 0x29, 0xAB}}

type IBluetoothLEConnectionParametersInterface interface {
	win32.IInspectableInterface
	Get_LinkTimeout() uint16
	Get_ConnectionLatency() uint16
	Get_ConnectionInterval() uint16
}

type IBluetoothLEConnectionParametersVtbl struct {
	win32.IInspectableVtbl
	Get_LinkTimeout        uintptr
	Get_ConnectionLatency  uintptr
	Get_ConnectionInterval uintptr
}

type IBluetoothLEConnectionParameters struct {
	win32.IInspectable
}

func (this *IBluetoothLEConnectionParameters) Vtbl() *IBluetoothLEConnectionParametersVtbl {
	return (*IBluetoothLEConnectionParametersVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEConnectionParameters) Get_LinkTimeout() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LinkTimeout, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEConnectionParameters) Get_ConnectionLatency() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConnectionLatency, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEConnectionParameters) Get_ConnectionInterval() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConnectionInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 781E5E48-621E-5A7E-8BE6-1B9561FF63C9
var IID_IBluetoothLEConnectionPhy = syscall.GUID{0x781E5E48, 0x621E, 0x5A7E,
	[8]byte{0x8B, 0xE6, 0x1B, 0x95, 0x61, 0xFF, 0x63, 0xC9}}

type IBluetoothLEConnectionPhyInterface interface {
	win32.IInspectableInterface
	Get_TransmitInfo() *IBluetoothLEConnectionPhyInfo
	Get_ReceiveInfo() *IBluetoothLEConnectionPhyInfo
}

type IBluetoothLEConnectionPhyVtbl struct {
	win32.IInspectableVtbl
	Get_TransmitInfo uintptr
	Get_ReceiveInfo  uintptr
}

type IBluetoothLEConnectionPhy struct {
	win32.IInspectable
}

func (this *IBluetoothLEConnectionPhy) Vtbl() *IBluetoothLEConnectionPhyVtbl {
	return (*IBluetoothLEConnectionPhyVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEConnectionPhy) Get_TransmitInfo() *IBluetoothLEConnectionPhyInfo {
	var _result *IBluetoothLEConnectionPhyInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TransmitInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEConnectionPhy) Get_ReceiveInfo() *IBluetoothLEConnectionPhyInfo {
	var _result *IBluetoothLEConnectionPhyInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReceiveInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9A100BDD-602E-5C27-A1AE-B230015A6394
var IID_IBluetoothLEConnectionPhyInfo = syscall.GUID{0x9A100BDD, 0x602E, 0x5C27,
	[8]byte{0xA1, 0xAE, 0xB2, 0x30, 0x01, 0x5A, 0x63, 0x94}}

type IBluetoothLEConnectionPhyInfoInterface interface {
	win32.IInspectableInterface
	Get_IsUncoded1MPhy() bool
	Get_IsUncoded2MPhy() bool
	Get_IsCodedPhy() bool
}

type IBluetoothLEConnectionPhyInfoVtbl struct {
	win32.IInspectableVtbl
	Get_IsUncoded1MPhy uintptr
	Get_IsUncoded2MPhy uintptr
	Get_IsCodedPhy     uintptr
}

type IBluetoothLEConnectionPhyInfo struct {
	win32.IInspectable
}

func (this *IBluetoothLEConnectionPhyInfo) Vtbl() *IBluetoothLEConnectionPhyInfoVtbl {
	return (*IBluetoothLEConnectionPhyInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEConnectionPhyInfo) Get_IsUncoded1MPhy() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsUncoded1MPhy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEConnectionPhyInfo) Get_IsUncoded2MPhy() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsUncoded2MPhy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEConnectionPhyInfo) Get_IsCodedPhy() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCodedPhy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// B5EE2F7B-4AD8-4642-AC48-80A0B500E887
var IID_IBluetoothLEDevice = syscall.GUID{0xB5EE2F7B, 0x4AD8, 0x4642,
	[8]byte{0xAC, 0x48, 0x80, 0xA0, 0xB5, 0x00, 0xE8, 0x87}}

type IBluetoothLEDeviceInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_Name() string
	Get_GattServices() *IVectorView[*IGattDeviceService]
	Get_ConnectionStatus() BluetoothConnectionStatus
	Get_BluetoothAddress() uint64
	GetGattService(serviceUuid syscall.GUID) *IGattDeviceService
	Add_NameChanged(handler TypedEventHandler[*IBluetoothLEDevice, interface{}]) EventRegistrationToken
	Remove_NameChanged(token EventRegistrationToken)
	Add_GattServicesChanged(handler TypedEventHandler[*IBluetoothLEDevice, interface{}]) EventRegistrationToken
	Remove_GattServicesChanged(token EventRegistrationToken)
	Add_ConnectionStatusChanged(handler TypedEventHandler[*IBluetoothLEDevice, interface{}]) EventRegistrationToken
	Remove_ConnectionStatusChanged(token EventRegistrationToken)
}

type IBluetoothLEDeviceVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId                   uintptr
	Get_Name                       uintptr
	Get_GattServices               uintptr
	Get_ConnectionStatus           uintptr
	Get_BluetoothAddress           uintptr
	GetGattService                 uintptr
	Add_NameChanged                uintptr
	Remove_NameChanged             uintptr
	Add_GattServicesChanged        uintptr
	Remove_GattServicesChanged     uintptr
	Add_ConnectionStatusChanged    uintptr
	Remove_ConnectionStatusChanged uintptr
}

type IBluetoothLEDevice struct {
	win32.IInspectable
}

func (this *IBluetoothLEDevice) Vtbl() *IBluetoothLEDeviceVtbl {
	return (*IBluetoothLEDeviceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEDevice) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBluetoothLEDevice) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBluetoothLEDevice) Get_GattServices() *IVectorView[*IGattDeviceService] {
	var _result *IVectorView[*IGattDeviceService]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GattServices, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEDevice) Get_ConnectionStatus() BluetoothConnectionStatus {
	var _result BluetoothConnectionStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConnectionStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEDevice) Get_BluetoothAddress() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BluetoothAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEDevice) GetGattService(serviceUuid syscall.GUID) *IGattDeviceService {
	var _result *IGattDeviceService
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetGattService, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&serviceUuid)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEDevice) Add_NameChanged(handler TypedEventHandler[*IBluetoothLEDevice, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_NameChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEDevice) Remove_NameChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_NameChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IBluetoothLEDevice) Add_GattServicesChanged(handler TypedEventHandler[*IBluetoothLEDevice, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_GattServicesChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEDevice) Remove_GattServicesChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_GattServicesChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IBluetoothLEDevice) Add_ConnectionStatusChanged(handler TypedEventHandler[*IBluetoothLEDevice, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ConnectionStatusChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEDevice) Remove_ConnectionStatusChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ConnectionStatusChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 26F062B3-7AEE-4D31-BABA-B1B9775F5916
var IID_IBluetoothLEDevice2 = syscall.GUID{0x26F062B3, 0x7AEE, 0x4D31,
	[8]byte{0xBA, 0xBA, 0xB1, 0xB9, 0x77, 0x5F, 0x59, 0x16}}

type IBluetoothLEDevice2Interface interface {
	win32.IInspectableInterface
	Get_DeviceInformation() *IDeviceInformation
	Get_Appearance() *IBluetoothLEAppearance
	Get_BluetoothAddressType() BluetoothAddressType
}

type IBluetoothLEDevice2Vtbl struct {
	win32.IInspectableVtbl
	Get_DeviceInformation    uintptr
	Get_Appearance           uintptr
	Get_BluetoothAddressType uintptr
}

type IBluetoothLEDevice2 struct {
	win32.IInspectable
}

func (this *IBluetoothLEDevice2) Vtbl() *IBluetoothLEDevice2Vtbl {
	return (*IBluetoothLEDevice2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEDevice2) Get_DeviceInformation() *IDeviceInformation {
	var _result *IDeviceInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEDevice2) Get_Appearance() *IBluetoothLEAppearance {
	var _result *IBluetoothLEAppearance
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Appearance, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEDevice2) Get_BluetoothAddressType() BluetoothAddressType {
	var _result BluetoothAddressType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BluetoothAddressType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// AEE9E493-44AC-40DC-AF33-B2C13C01CA46
var IID_IBluetoothLEDevice3 = syscall.GUID{0xAEE9E493, 0x44AC, 0x40DC,
	[8]byte{0xAF, 0x33, 0xB2, 0xC1, 0x3C, 0x01, 0xCA, 0x46}}

type IBluetoothLEDevice3Interface interface {
	win32.IInspectableInterface
	Get_DeviceAccessInformation() *IDeviceAccessInformation
	RequestAccessAsync() *IAsyncOperation[DeviceAccessStatus]
	GetGattServicesAsync() *IAsyncOperation[*IGattDeviceServicesResult]
	GetGattServicesWithCacheModeAsync(cacheMode BluetoothCacheMode) *IAsyncOperation[*IGattDeviceServicesResult]
	GetGattServicesForUuidAsync(serviceUuid syscall.GUID) *IAsyncOperation[*IGattDeviceServicesResult]
	GetGattServicesForUuidWithCacheModeAsync(serviceUuid syscall.GUID, cacheMode BluetoothCacheMode) *IAsyncOperation[*IGattDeviceServicesResult]
}

type IBluetoothLEDevice3Vtbl struct {
	win32.IInspectableVtbl
	Get_DeviceAccessInformation              uintptr
	RequestAccessAsync                       uintptr
	GetGattServicesAsync                     uintptr
	GetGattServicesWithCacheModeAsync        uintptr
	GetGattServicesForUuidAsync              uintptr
	GetGattServicesForUuidWithCacheModeAsync uintptr
}

type IBluetoothLEDevice3 struct {
	win32.IInspectable
}

func (this *IBluetoothLEDevice3) Vtbl() *IBluetoothLEDevice3Vtbl {
	return (*IBluetoothLEDevice3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEDevice3) Get_DeviceAccessInformation() *IDeviceAccessInformation {
	var _result *IDeviceAccessInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceAccessInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEDevice3) RequestAccessAsync() *IAsyncOperation[DeviceAccessStatus] {
	var _result *IAsyncOperation[DeviceAccessStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEDevice3) GetGattServicesAsync() *IAsyncOperation[*IGattDeviceServicesResult] {
	var _result *IAsyncOperation[*IGattDeviceServicesResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetGattServicesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEDevice3) GetGattServicesWithCacheModeAsync(cacheMode BluetoothCacheMode) *IAsyncOperation[*IGattDeviceServicesResult] {
	var _result *IAsyncOperation[*IGattDeviceServicesResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetGattServicesWithCacheModeAsync, uintptr(unsafe.Pointer(this)), uintptr(cacheMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEDevice3) GetGattServicesForUuidAsync(serviceUuid syscall.GUID) *IAsyncOperation[*IGattDeviceServicesResult] {
	var _result *IAsyncOperation[*IGattDeviceServicesResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetGattServicesForUuidAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&serviceUuid)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEDevice3) GetGattServicesForUuidWithCacheModeAsync(serviceUuid syscall.GUID, cacheMode BluetoothCacheMode) *IAsyncOperation[*IGattDeviceServicesResult] {
	var _result *IAsyncOperation[*IGattDeviceServicesResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetGattServicesForUuidWithCacheModeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&serviceUuid)), uintptr(cacheMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2B605031-2248-4B2F-ACF0-7CEE36FC5870
var IID_IBluetoothLEDevice4 = syscall.GUID{0x2B605031, 0x2248, 0x4B2F,
	[8]byte{0xAC, 0xF0, 0x7C, 0xEE, 0x36, 0xFC, 0x58, 0x70}}

type IBluetoothLEDevice4Interface interface {
	win32.IInspectableInterface
	Get_BluetoothDeviceId() *IBluetoothDeviceId
}

type IBluetoothLEDevice4Vtbl struct {
	win32.IInspectableVtbl
	Get_BluetoothDeviceId uintptr
}

type IBluetoothLEDevice4 struct {
	win32.IInspectable
}

func (this *IBluetoothLEDevice4) Vtbl() *IBluetoothLEDevice4Vtbl {
	return (*IBluetoothLEDevice4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEDevice4) Get_BluetoothDeviceId() *IBluetoothDeviceId {
	var _result *IBluetoothDeviceId
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BluetoothDeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9D6A1260-5287-458E-95BA-17C8B7BB326E
var IID_IBluetoothLEDevice5 = syscall.GUID{0x9D6A1260, 0x5287, 0x458E,
	[8]byte{0x95, 0xBA, 0x17, 0xC8, 0xB7, 0xBB, 0x32, 0x6E}}

type IBluetoothLEDevice5Interface interface {
	win32.IInspectableInterface
	Get_WasSecureConnectionUsedForPairing() bool
}

type IBluetoothLEDevice5Vtbl struct {
	win32.IInspectableVtbl
	Get_WasSecureConnectionUsedForPairing uintptr
}

type IBluetoothLEDevice5 struct {
	win32.IInspectable
}

func (this *IBluetoothLEDevice5) Vtbl() *IBluetoothLEDevice5Vtbl {
	return (*IBluetoothLEDevice5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEDevice5) Get_WasSecureConnectionUsedForPairing() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WasSecureConnectionUsedForPairing, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// CA7190EF-0CAE-573C-A1CA-E1FC5BFC39E2
var IID_IBluetoothLEDevice6 = syscall.GUID{0xCA7190EF, 0x0CAE, 0x573C,
	[8]byte{0xA1, 0xCA, 0xE1, 0xFC, 0x5B, 0xFC, 0x39, 0xE2}}

type IBluetoothLEDevice6Interface interface {
	win32.IInspectableInterface
	GetConnectionParameters() *IBluetoothLEConnectionParameters
	GetConnectionPhy() *IBluetoothLEConnectionPhy
	RequestPreferredConnectionParameters(preferredConnectionParameters *IBluetoothLEPreferredConnectionParameters) *IBluetoothLEPreferredConnectionParametersRequest
	Add_ConnectionParametersChanged(handler TypedEventHandler[*IBluetoothLEDevice, interface{}]) EventRegistrationToken
	Remove_ConnectionParametersChanged(token EventRegistrationToken)
	Add_ConnectionPhyChanged(handler TypedEventHandler[*IBluetoothLEDevice, interface{}]) EventRegistrationToken
	Remove_ConnectionPhyChanged(token EventRegistrationToken)
}

type IBluetoothLEDevice6Vtbl struct {
	win32.IInspectableVtbl
	GetConnectionParameters              uintptr
	GetConnectionPhy                     uintptr
	RequestPreferredConnectionParameters uintptr
	Add_ConnectionParametersChanged      uintptr
	Remove_ConnectionParametersChanged   uintptr
	Add_ConnectionPhyChanged             uintptr
	Remove_ConnectionPhyChanged          uintptr
}

type IBluetoothLEDevice6 struct {
	win32.IInspectable
}

func (this *IBluetoothLEDevice6) Vtbl() *IBluetoothLEDevice6Vtbl {
	return (*IBluetoothLEDevice6Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEDevice6) GetConnectionParameters() *IBluetoothLEConnectionParameters {
	var _result *IBluetoothLEConnectionParameters
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetConnectionParameters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEDevice6) GetConnectionPhy() *IBluetoothLEConnectionPhy {
	var _result *IBluetoothLEConnectionPhy
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetConnectionPhy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEDevice6) RequestPreferredConnectionParameters(preferredConnectionParameters *IBluetoothLEPreferredConnectionParameters) *IBluetoothLEPreferredConnectionParametersRequest {
	var _result *IBluetoothLEPreferredConnectionParametersRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestPreferredConnectionParameters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(preferredConnectionParameters)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEDevice6) Add_ConnectionParametersChanged(handler TypedEventHandler[*IBluetoothLEDevice, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ConnectionParametersChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEDevice6) Remove_ConnectionParametersChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ConnectionParametersChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IBluetoothLEDevice6) Add_ConnectionPhyChanged(handler TypedEventHandler[*IBluetoothLEDevice, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ConnectionPhyChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEDevice6) Remove_ConnectionPhyChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ConnectionPhyChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// C8CF1A19-F0B6-4BF0-8689-41303DE2D9F4
var IID_IBluetoothLEDeviceStatics = syscall.GUID{0xC8CF1A19, 0xF0B6, 0x4BF0,
	[8]byte{0x86, 0x89, 0x41, 0x30, 0x3D, 0xE2, 0xD9, 0xF4}}

type IBluetoothLEDeviceStaticsInterface interface {
	win32.IInspectableInterface
	FromIdAsync(deviceId string) *IAsyncOperation[*IBluetoothLEDevice]
	FromBluetoothAddressAsync(bluetoothAddress uint64) *IAsyncOperation[*IBluetoothLEDevice]
	GetDeviceSelector() string
}

type IBluetoothLEDeviceStaticsVtbl struct {
	win32.IInspectableVtbl
	FromIdAsync               uintptr
	FromBluetoothAddressAsync uintptr
	GetDeviceSelector         uintptr
}

type IBluetoothLEDeviceStatics struct {
	win32.IInspectable
}

func (this *IBluetoothLEDeviceStatics) Vtbl() *IBluetoothLEDeviceStaticsVtbl {
	return (*IBluetoothLEDeviceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEDeviceStatics) FromIdAsync(deviceId string) *IAsyncOperation[*IBluetoothLEDevice] {
	var _result *IAsyncOperation[*IBluetoothLEDevice]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEDeviceStatics) FromBluetoothAddressAsync(bluetoothAddress uint64) *IAsyncOperation[*IBluetoothLEDevice] {
	var _result *IAsyncOperation[*IBluetoothLEDevice]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromBluetoothAddressAsync, uintptr(unsafe.Pointer(this)), uintptr(bluetoothAddress), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEDeviceStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 5F12C06B-3BAC-43E8-AD16-563271BD41C2
var IID_IBluetoothLEDeviceStatics2 = syscall.GUID{0x5F12C06B, 0x3BAC, 0x43E8,
	[8]byte{0xAD, 0x16, 0x56, 0x32, 0x71, 0xBD, 0x41, 0xC2}}

type IBluetoothLEDeviceStatics2Interface interface {
	win32.IInspectableInterface
	GetDeviceSelectorFromPairingState(pairingState bool) string
	GetDeviceSelectorFromConnectionStatus(connectionStatus BluetoothConnectionStatus) string
	GetDeviceSelectorFromDeviceName(deviceName string) string
	GetDeviceSelectorFromBluetoothAddress(bluetoothAddress uint64) string
	GetDeviceSelectorFromBluetoothAddressWithBluetoothAddressType(bluetoothAddress uint64, bluetoothAddressType BluetoothAddressType) string
	GetDeviceSelectorFromAppearance(appearance *IBluetoothLEAppearance) string
	FromBluetoothAddressWithBluetoothAddressTypeAsync(bluetoothAddress uint64, bluetoothAddressType BluetoothAddressType) *IAsyncOperation[*IBluetoothLEDevice]
}

type IBluetoothLEDeviceStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelectorFromPairingState                             uintptr
	GetDeviceSelectorFromConnectionStatus                         uintptr
	GetDeviceSelectorFromDeviceName                               uintptr
	GetDeviceSelectorFromBluetoothAddress                         uintptr
	GetDeviceSelectorFromBluetoothAddressWithBluetoothAddressType uintptr
	GetDeviceSelectorFromAppearance                               uintptr
	FromBluetoothAddressWithBluetoothAddressTypeAsync             uintptr
}

type IBluetoothLEDeviceStatics2 struct {
	win32.IInspectable
}

func (this *IBluetoothLEDeviceStatics2) Vtbl() *IBluetoothLEDeviceStatics2Vtbl {
	return (*IBluetoothLEDeviceStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEDeviceStatics2) GetDeviceSelectorFromPairingState(pairingState bool) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorFromPairingState, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&pairingState))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBluetoothLEDeviceStatics2) GetDeviceSelectorFromConnectionStatus(connectionStatus BluetoothConnectionStatus) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorFromConnectionStatus, uintptr(unsafe.Pointer(this)), uintptr(connectionStatus), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBluetoothLEDeviceStatics2) GetDeviceSelectorFromDeviceName(deviceName string) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorFromDeviceName, uintptr(unsafe.Pointer(this)), NewHStr(deviceName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBluetoothLEDeviceStatics2) GetDeviceSelectorFromBluetoothAddress(bluetoothAddress uint64) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorFromBluetoothAddress, uintptr(unsafe.Pointer(this)), uintptr(bluetoothAddress), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBluetoothLEDeviceStatics2) GetDeviceSelectorFromBluetoothAddressWithBluetoothAddressType(bluetoothAddress uint64, bluetoothAddressType BluetoothAddressType) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorFromBluetoothAddressWithBluetoothAddressType, uintptr(unsafe.Pointer(this)), uintptr(bluetoothAddress), uintptr(bluetoothAddressType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBluetoothLEDeviceStatics2) GetDeviceSelectorFromAppearance(appearance *IBluetoothLEAppearance) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorFromAppearance, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(appearance)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBluetoothLEDeviceStatics2) FromBluetoothAddressWithBluetoothAddressTypeAsync(bluetoothAddress uint64, bluetoothAddressType BluetoothAddressType) *IAsyncOperation[*IBluetoothLEDevice] {
	var _result *IAsyncOperation[*IBluetoothLEDevice]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromBluetoothAddressWithBluetoothAddressTypeAsync, uintptr(unsafe.Pointer(this)), uintptr(bluetoothAddress), uintptr(bluetoothAddressType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F2F44344-7372-5F7B-9B34-29C944F5A715
var IID_IBluetoothLEPreferredConnectionParameters = syscall.GUID{0xF2F44344, 0x7372, 0x5F7B,
	[8]byte{0x9B, 0x34, 0x29, 0xC9, 0x44, 0xF5, 0xA7, 0x15}}

type IBluetoothLEPreferredConnectionParametersInterface interface {
	win32.IInspectableInterface
	Get_LinkTimeout() uint16
	Get_ConnectionLatency() uint16
	Get_MinConnectionInterval() uint16
	Get_MaxConnectionInterval() uint16
}

type IBluetoothLEPreferredConnectionParametersVtbl struct {
	win32.IInspectableVtbl
	Get_LinkTimeout           uintptr
	Get_ConnectionLatency     uintptr
	Get_MinConnectionInterval uintptr
	Get_MaxConnectionInterval uintptr
}

type IBluetoothLEPreferredConnectionParameters struct {
	win32.IInspectable
}

func (this *IBluetoothLEPreferredConnectionParameters) Vtbl() *IBluetoothLEPreferredConnectionParametersVtbl {
	return (*IBluetoothLEPreferredConnectionParametersVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEPreferredConnectionParameters) Get_LinkTimeout() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LinkTimeout, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEPreferredConnectionParameters) Get_ConnectionLatency() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConnectionLatency, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEPreferredConnectionParameters) Get_MinConnectionInterval() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinConnectionInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEPreferredConnectionParameters) Get_MaxConnectionInterval() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxConnectionInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 8A375276-A528-5266-B661-CCE6A5FF9739
var IID_IBluetoothLEPreferredConnectionParametersRequest = syscall.GUID{0x8A375276, 0xA528, 0x5266,
	[8]byte{0xB6, 0x61, 0xCC, 0xE6, 0xA5, 0xFF, 0x97, 0x39}}

type IBluetoothLEPreferredConnectionParametersRequestInterface interface {
	win32.IInspectableInterface
	Get_Status() BluetoothLEPreferredConnectionParametersRequestStatus
}

type IBluetoothLEPreferredConnectionParametersRequestVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
}

type IBluetoothLEPreferredConnectionParametersRequest struct {
	win32.IInspectable
}

func (this *IBluetoothLEPreferredConnectionParametersRequest) Vtbl() *IBluetoothLEPreferredConnectionParametersRequestVtbl {
	return (*IBluetoothLEPreferredConnectionParametersRequestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEPreferredConnectionParametersRequest) Get_Status() BluetoothLEPreferredConnectionParametersRequestStatus {
	var _result BluetoothLEPreferredConnectionParametersRequestStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 0E3E8EDC-2751-55AA-A838-8FAEEE818D72
var IID_IBluetoothLEPreferredConnectionParametersStatics = syscall.GUID{0x0E3E8EDC, 0x2751, 0x55AA,
	[8]byte{0xA8, 0x38, 0x8F, 0xAE, 0xEE, 0x81, 0x8D, 0x72}}

type IBluetoothLEPreferredConnectionParametersStaticsInterface interface {
	win32.IInspectableInterface
	Get_Balanced() *IBluetoothLEPreferredConnectionParameters
	Get_ThroughputOptimized() *IBluetoothLEPreferredConnectionParameters
	Get_PowerOptimized() *IBluetoothLEPreferredConnectionParameters
}

type IBluetoothLEPreferredConnectionParametersStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Balanced            uintptr
	Get_ThroughputOptimized uintptr
	Get_PowerOptimized      uintptr
}

type IBluetoothLEPreferredConnectionParametersStatics struct {
	win32.IInspectable
}

func (this *IBluetoothLEPreferredConnectionParametersStatics) Vtbl() *IBluetoothLEPreferredConnectionParametersStaticsVtbl {
	return (*IBluetoothLEPreferredConnectionParametersStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEPreferredConnectionParametersStatics) Get_Balanced() *IBluetoothLEPreferredConnectionParameters {
	var _result *IBluetoothLEPreferredConnectionParameters
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Balanced, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEPreferredConnectionParametersStatics) Get_ThroughputOptimized() *IBluetoothLEPreferredConnectionParameters {
	var _result *IBluetoothLEPreferredConnectionParameters
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ThroughputOptimized, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEPreferredConnectionParametersStatics) Get_PowerOptimized() *IBluetoothLEPreferredConnectionParameters {
	var _result *IBluetoothLEPreferredConnectionParameters
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PowerOptimized, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DF7B7391-6BB5-4CFE-90B1-5D7324EDCF7F
var IID_IBluetoothSignalStrengthFilter = syscall.GUID{0xDF7B7391, 0x6BB5, 0x4CFE,
	[8]byte{0x90, 0xB1, 0x5D, 0x73, 0x24, 0xED, 0xCF, 0x7F}}

type IBluetoothSignalStrengthFilterInterface interface {
	win32.IInspectableInterface
	Get_InRangeThresholdInDBm() *IReference[int16]
	Put_InRangeThresholdInDBm(value *IReference[int16])
	Get_OutOfRangeThresholdInDBm() *IReference[int16]
	Put_OutOfRangeThresholdInDBm(value *IReference[int16])
	Get_OutOfRangeTimeout() *IReference[TimeSpan]
	Put_OutOfRangeTimeout(value *IReference[TimeSpan])
	Get_SamplingInterval() *IReference[TimeSpan]
	Put_SamplingInterval(value *IReference[TimeSpan])
}

type IBluetoothSignalStrengthFilterVtbl struct {
	win32.IInspectableVtbl
	Get_InRangeThresholdInDBm    uintptr
	Put_InRangeThresholdInDBm    uintptr
	Get_OutOfRangeThresholdInDBm uintptr
	Put_OutOfRangeThresholdInDBm uintptr
	Get_OutOfRangeTimeout        uintptr
	Put_OutOfRangeTimeout        uintptr
	Get_SamplingInterval         uintptr
	Put_SamplingInterval         uintptr
}

type IBluetoothSignalStrengthFilter struct {
	win32.IInspectable
}

func (this *IBluetoothSignalStrengthFilter) Vtbl() *IBluetoothSignalStrengthFilterVtbl {
	return (*IBluetoothSignalStrengthFilterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothSignalStrengthFilter) Get_InRangeThresholdInDBm() *IReference[int16] {
	var _result *IReference[int16]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InRangeThresholdInDBm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothSignalStrengthFilter) Put_InRangeThresholdInDBm(value *IReference[int16]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InRangeThresholdInDBm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IBluetoothSignalStrengthFilter) Get_OutOfRangeThresholdInDBm() *IReference[int16] {
	var _result *IReference[int16]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutOfRangeThresholdInDBm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothSignalStrengthFilter) Put_OutOfRangeThresholdInDBm(value *IReference[int16]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_OutOfRangeThresholdInDBm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IBluetoothSignalStrengthFilter) Get_OutOfRangeTimeout() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutOfRangeTimeout, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothSignalStrengthFilter) Put_OutOfRangeTimeout(value *IReference[TimeSpan]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_OutOfRangeTimeout, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IBluetoothSignalStrengthFilter) Get_SamplingInterval() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SamplingInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothSignalStrengthFilter) Put_SamplingInterval(value *IReference[TimeSpan]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SamplingInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 17DF0CD8-CF74-4B21-AFE6-F57A11BCDEA0
var IID_IBluetoothUuidHelperStatics = syscall.GUID{0x17DF0CD8, 0xCF74, 0x4B21,
	[8]byte{0xAF, 0xE6, 0xF5, 0x7A, 0x11, 0xBC, 0xDE, 0xA0}}

type IBluetoothUuidHelperStaticsInterface interface {
	win32.IInspectableInterface
	FromShortId(shortId uint32) syscall.GUID
	TryGetShortId(uuid syscall.GUID) *IReference[uint32]
}

type IBluetoothUuidHelperStaticsVtbl struct {
	win32.IInspectableVtbl
	FromShortId   uintptr
	TryGetShortId uintptr
}

type IBluetoothUuidHelperStatics struct {
	win32.IInspectable
}

func (this *IBluetoothUuidHelperStatics) Vtbl() *IBluetoothUuidHelperStaticsVtbl {
	return (*IBluetoothUuidHelperStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothUuidHelperStatics) FromShortId(shortId uint32) syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromShortId, uintptr(unsafe.Pointer(this)), uintptr(shortId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothUuidHelperStatics) TryGetShortId(uuid syscall.GUID) *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetShortId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&uuid)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type BluetoothAdapter struct {
	RtClass
	*IBluetoothAdapter
}

func NewIBluetoothAdapterStatics() *IBluetoothAdapterStatics {
	var p *IBluetoothAdapterStatics
	hs := NewHStr("Windows.Devices.Bluetooth.BluetoothAdapter")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IBluetoothAdapterStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type BluetoothClassOfDevice struct {
	RtClass
	*IBluetoothClassOfDevice
}

func NewIBluetoothClassOfDeviceStatics() *IBluetoothClassOfDeviceStatics {
	var p *IBluetoothClassOfDeviceStatics
	hs := NewHStr("Windows.Devices.Bluetooth.BluetoothClassOfDevice")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IBluetoothClassOfDeviceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type BluetoothDevice struct {
	RtClass
	*IBluetoothDevice
}

func NewIBluetoothDeviceStatics2() *IBluetoothDeviceStatics2 {
	var p *IBluetoothDeviceStatics2
	hs := NewHStr("Windows.Devices.Bluetooth.BluetoothDevice")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IBluetoothDeviceStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIBluetoothDeviceStatics() *IBluetoothDeviceStatics {
	var p *IBluetoothDeviceStatics
	hs := NewHStr("Windows.Devices.Bluetooth.BluetoothDevice")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IBluetoothDeviceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type BluetoothDeviceId struct {
	RtClass
	*IBluetoothDeviceId
}

func NewIBluetoothDeviceIdStatics() *IBluetoothDeviceIdStatics {
	var p *IBluetoothDeviceIdStatics
	hs := NewHStr("Windows.Devices.Bluetooth.BluetoothDeviceId")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IBluetoothDeviceIdStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type BluetoothLEAppearance struct {
	RtClass
	*IBluetoothLEAppearance
}

func NewIBluetoothLEAppearanceStatics() *IBluetoothLEAppearanceStatics {
	var p *IBluetoothLEAppearanceStatics
	hs := NewHStr("Windows.Devices.Bluetooth.BluetoothLEAppearance")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IBluetoothLEAppearanceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type BluetoothLEAppearanceCategories struct {
	RtClass
}

func NewIBluetoothLEAppearanceCategoriesStatics() *IBluetoothLEAppearanceCategoriesStatics {
	var p *IBluetoothLEAppearanceCategoriesStatics
	hs := NewHStr("Windows.Devices.Bluetooth.BluetoothLEAppearanceCategories")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IBluetoothLEAppearanceCategoriesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type BluetoothLEAppearanceSubcategories struct {
	RtClass
}

func NewIBluetoothLEAppearanceSubcategoriesStatics() *IBluetoothLEAppearanceSubcategoriesStatics {
	var p *IBluetoothLEAppearanceSubcategoriesStatics
	hs := NewHStr("Windows.Devices.Bluetooth.BluetoothLEAppearanceSubcategories")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IBluetoothLEAppearanceSubcategoriesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type BluetoothLEConnectionParameters struct {
	RtClass
	*IBluetoothLEConnectionParameters
}

type BluetoothLEConnectionPhy struct {
	RtClass
	*IBluetoothLEConnectionPhy
}

type BluetoothLEConnectionPhyInfo struct {
	RtClass
	*IBluetoothLEConnectionPhyInfo
}

type BluetoothLEDevice struct {
	RtClass
	*IBluetoothLEDevice
}

func NewIBluetoothLEDeviceStatics() *IBluetoothLEDeviceStatics {
	var p *IBluetoothLEDeviceStatics
	hs := NewHStr("Windows.Devices.Bluetooth.BluetoothLEDevice")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IBluetoothLEDeviceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIBluetoothLEDeviceStatics2() *IBluetoothLEDeviceStatics2 {
	var p *IBluetoothLEDeviceStatics2
	hs := NewHStr("Windows.Devices.Bluetooth.BluetoothLEDevice")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IBluetoothLEDeviceStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type BluetoothLEPreferredConnectionParameters struct {
	RtClass
	*IBluetoothLEPreferredConnectionParameters
}

func NewIBluetoothLEPreferredConnectionParametersStatics() *IBluetoothLEPreferredConnectionParametersStatics {
	var p *IBluetoothLEPreferredConnectionParametersStatics
	hs := NewHStr("Windows.Devices.Bluetooth.BluetoothLEPreferredConnectionParameters")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IBluetoothLEPreferredConnectionParametersStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type BluetoothLEPreferredConnectionParametersRequest struct {
	RtClass
	*IBluetoothLEPreferredConnectionParametersRequest
}

type BluetoothSignalStrengthFilter struct {
	RtClass
	*IBluetoothSignalStrengthFilter
}

func NewBluetoothSignalStrengthFilter() *BluetoothSignalStrengthFilter {
	hs := NewHStr("Windows.Devices.Bluetooth.BluetoothSignalStrengthFilter")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &BluetoothSignalStrengthFilter{
		RtClass:                        RtClass{PInspect: p},
		IBluetoothSignalStrengthFilter: (*IBluetoothSignalStrengthFilter)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type BluetoothUuidHelper struct {
	RtClass
}

func NewIBluetoothUuidHelperStatics() *IBluetoothUuidHelperStatics {
	var p *IBluetoothUuidHelperStatics
	hs := NewHStr("Windows.Devices.Bluetooth.BluetoothUuidHelper")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IBluetoothUuidHelperStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
