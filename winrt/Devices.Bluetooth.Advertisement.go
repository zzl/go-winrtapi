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
type BluetoothLEAdvertisementFlags uint32

const (
	BluetoothLEAdvertisementFlags_None                      BluetoothLEAdvertisementFlags = 0
	BluetoothLEAdvertisementFlags_LimitedDiscoverableMode   BluetoothLEAdvertisementFlags = 1
	BluetoothLEAdvertisementFlags_GeneralDiscoverableMode   BluetoothLEAdvertisementFlags = 2
	BluetoothLEAdvertisementFlags_ClassicNotSupported       BluetoothLEAdvertisementFlags = 4
	BluetoothLEAdvertisementFlags_DualModeControllerCapable BluetoothLEAdvertisementFlags = 8
	BluetoothLEAdvertisementFlags_DualModeHostCapable       BluetoothLEAdvertisementFlags = 16
)

// enum
type BluetoothLEAdvertisementPublisherStatus int32

const (
	BluetoothLEAdvertisementPublisherStatus_Created  BluetoothLEAdvertisementPublisherStatus = 0
	BluetoothLEAdvertisementPublisherStatus_Waiting  BluetoothLEAdvertisementPublisherStatus = 1
	BluetoothLEAdvertisementPublisherStatus_Started  BluetoothLEAdvertisementPublisherStatus = 2
	BluetoothLEAdvertisementPublisherStatus_Stopping BluetoothLEAdvertisementPublisherStatus = 3
	BluetoothLEAdvertisementPublisherStatus_Stopped  BluetoothLEAdvertisementPublisherStatus = 4
	BluetoothLEAdvertisementPublisherStatus_Aborted  BluetoothLEAdvertisementPublisherStatus = 5
)

// enum
type BluetoothLEAdvertisementType int32

const (
	BluetoothLEAdvertisementType_ConnectableUndirected    BluetoothLEAdvertisementType = 0
	BluetoothLEAdvertisementType_ConnectableDirected      BluetoothLEAdvertisementType = 1
	BluetoothLEAdvertisementType_ScannableUndirected      BluetoothLEAdvertisementType = 2
	BluetoothLEAdvertisementType_NonConnectableUndirected BluetoothLEAdvertisementType = 3
	BluetoothLEAdvertisementType_ScanResponse             BluetoothLEAdvertisementType = 4
	BluetoothLEAdvertisementType_Extended                 BluetoothLEAdvertisementType = 5
)

// enum
type BluetoothLEAdvertisementWatcherStatus int32

const (
	BluetoothLEAdvertisementWatcherStatus_Created  BluetoothLEAdvertisementWatcherStatus = 0
	BluetoothLEAdvertisementWatcherStatus_Started  BluetoothLEAdvertisementWatcherStatus = 1
	BluetoothLEAdvertisementWatcherStatus_Stopping BluetoothLEAdvertisementWatcherStatus = 2
	BluetoothLEAdvertisementWatcherStatus_Stopped  BluetoothLEAdvertisementWatcherStatus = 3
	BluetoothLEAdvertisementWatcherStatus_Aborted  BluetoothLEAdvertisementWatcherStatus = 4
)

// enum
type BluetoothLEScanningMode int32

const (
	BluetoothLEScanningMode_Passive BluetoothLEScanningMode = 0
	BluetoothLEScanningMode_Active  BluetoothLEScanningMode = 1
	BluetoothLEScanningMode_None    BluetoothLEScanningMode = 2
)

// interfaces

// 066FB2B7-33D1-4E7D-8367-CF81D0F79653
var IID_IBluetoothLEAdvertisement = syscall.GUID{0x066FB2B7, 0x33D1, 0x4E7D,
	[8]byte{0x83, 0x67, 0xCF, 0x81, 0xD0, 0xF7, 0x96, 0x53}}

type IBluetoothLEAdvertisementInterface interface {
	win32.IInspectableInterface
	Get_Flags() *IReference[BluetoothLEAdvertisementFlags]
	Put_Flags(value *IReference[BluetoothLEAdvertisementFlags])
	Get_LocalName() string
	Put_LocalName(value string)
	Get_ServiceUuids() *IVector[syscall.GUID]
	Get_ManufacturerData() *IVector[*IBluetoothLEManufacturerData]
	Get_DataSections() *IVector[*IBluetoothLEAdvertisementDataSection]
	GetManufacturerDataByCompanyId(companyId uint16) *IVectorView[*IBluetoothLEManufacturerData]
	GetSectionsByType(type_ byte) *IVectorView[*IBluetoothLEAdvertisementDataSection]
}

type IBluetoothLEAdvertisementVtbl struct {
	win32.IInspectableVtbl
	Get_Flags                      uintptr
	Put_Flags                      uintptr
	Get_LocalName                  uintptr
	Put_LocalName                  uintptr
	Get_ServiceUuids               uintptr
	Get_ManufacturerData           uintptr
	Get_DataSections               uintptr
	GetManufacturerDataByCompanyId uintptr
	GetSectionsByType              uintptr
}

type IBluetoothLEAdvertisement struct {
	win32.IInspectable
}

func (this *IBluetoothLEAdvertisement) Vtbl() *IBluetoothLEAdvertisementVtbl {
	return (*IBluetoothLEAdvertisementVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAdvertisement) Get_Flags() *IReference[BluetoothLEAdvertisementFlags] {
	var _result *IReference[BluetoothLEAdvertisementFlags]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Flags, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEAdvertisement) Put_Flags(value *IReference[BluetoothLEAdvertisementFlags]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Flags, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IBluetoothLEAdvertisement) Get_LocalName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBluetoothLEAdvertisement) Put_LocalName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LocalName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IBluetoothLEAdvertisement) Get_ServiceUuids() *IVector[syscall.GUID] {
	var _result *IVector[syscall.GUID]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceUuids, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEAdvertisement) Get_ManufacturerData() *IVector[*IBluetoothLEManufacturerData] {
	var _result *IVector[*IBluetoothLEManufacturerData]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ManufacturerData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEAdvertisement) Get_DataSections() *IVector[*IBluetoothLEAdvertisementDataSection] {
	var _result *IVector[*IBluetoothLEAdvertisementDataSection]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DataSections, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEAdvertisement) GetManufacturerDataByCompanyId(companyId uint16) *IVectorView[*IBluetoothLEManufacturerData] {
	var _result *IVectorView[*IBluetoothLEManufacturerData]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetManufacturerDataByCompanyId, uintptr(unsafe.Pointer(this)), uintptr(companyId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEAdvertisement) GetSectionsByType(type_ byte) *IVectorView[*IBluetoothLEAdvertisementDataSection] {
	var _result *IVectorView[*IBluetoothLEAdvertisementDataSection]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSectionsByType, uintptr(unsafe.Pointer(this)), uintptr(type_), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FBFAD7F2-B9C5-4A08-BC51-502F8EF68A79
var IID_IBluetoothLEAdvertisementBytePattern = syscall.GUID{0xFBFAD7F2, 0xB9C5, 0x4A08,
	[8]byte{0xBC, 0x51, 0x50, 0x2F, 0x8E, 0xF6, 0x8A, 0x79}}

type IBluetoothLEAdvertisementBytePatternInterface interface {
	win32.IInspectableInterface
	Get_DataType() byte
	Put_DataType(value byte)
	Get_Offset() int16
	Put_Offset(value int16)
	Get_Data() *IBuffer
	Put_Data(value *IBuffer)
}

type IBluetoothLEAdvertisementBytePatternVtbl struct {
	win32.IInspectableVtbl
	Get_DataType uintptr
	Put_DataType uintptr
	Get_Offset   uintptr
	Put_Offset   uintptr
	Get_Data     uintptr
	Put_Data     uintptr
}

type IBluetoothLEAdvertisementBytePattern struct {
	win32.IInspectable
}

func (this *IBluetoothLEAdvertisementBytePattern) Vtbl() *IBluetoothLEAdvertisementBytePatternVtbl {
	return (*IBluetoothLEAdvertisementBytePatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAdvertisementBytePattern) Get_DataType() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DataType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementBytePattern) Put_DataType(value byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DataType, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IBluetoothLEAdvertisementBytePattern) Get_Offset() int16 {
	var _result int16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Offset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementBytePattern) Put_Offset(value int16) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Offset, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IBluetoothLEAdvertisementBytePattern) Get_Data() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEAdvertisementBytePattern) Put_Data(value *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// C2E24D73-FD5C-4EC3-BE2A-9CA6FA11B7BD
var IID_IBluetoothLEAdvertisementBytePatternFactory = syscall.GUID{0xC2E24D73, 0xFD5C, 0x4EC3,
	[8]byte{0xBE, 0x2A, 0x9C, 0xA6, 0xFA, 0x11, 0xB7, 0xBD}}

type IBluetoothLEAdvertisementBytePatternFactoryInterface interface {
	win32.IInspectableInterface
	Create(dataType byte, offset int16, data *IBuffer) *IBluetoothLEAdvertisementBytePattern
}

type IBluetoothLEAdvertisementBytePatternFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IBluetoothLEAdvertisementBytePatternFactory struct {
	win32.IInspectable
}

func (this *IBluetoothLEAdvertisementBytePatternFactory) Vtbl() *IBluetoothLEAdvertisementBytePatternFactoryVtbl {
	return (*IBluetoothLEAdvertisementBytePatternFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAdvertisementBytePatternFactory) Create(dataType byte, offset int16, data *IBuffer) *IBluetoothLEAdvertisementBytePattern {
	var _result *IBluetoothLEAdvertisementBytePattern
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(dataType), uintptr(offset), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D7213314-3A43-40F9-B6F0-92BFEFC34AE3
var IID_IBluetoothLEAdvertisementDataSection = syscall.GUID{0xD7213314, 0x3A43, 0x40F9,
	[8]byte{0xB6, 0xF0, 0x92, 0xBF, 0xEF, 0xC3, 0x4A, 0xE3}}

type IBluetoothLEAdvertisementDataSectionInterface interface {
	win32.IInspectableInterface
	Get_DataType() byte
	Put_DataType(value byte)
	Get_Data() *IBuffer
	Put_Data(value *IBuffer)
}

type IBluetoothLEAdvertisementDataSectionVtbl struct {
	win32.IInspectableVtbl
	Get_DataType uintptr
	Put_DataType uintptr
	Get_Data     uintptr
	Put_Data     uintptr
}

type IBluetoothLEAdvertisementDataSection struct {
	win32.IInspectable
}

func (this *IBluetoothLEAdvertisementDataSection) Vtbl() *IBluetoothLEAdvertisementDataSectionVtbl {
	return (*IBluetoothLEAdvertisementDataSectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAdvertisementDataSection) Get_DataType() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DataType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementDataSection) Put_DataType(value byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DataType, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IBluetoothLEAdvertisementDataSection) Get_Data() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEAdvertisementDataSection) Put_Data(value *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// E7A40942-A845-4045-BF7E-3E9971DB8A6B
var IID_IBluetoothLEAdvertisementDataSectionFactory = syscall.GUID{0xE7A40942, 0xA845, 0x4045,
	[8]byte{0xBF, 0x7E, 0x3E, 0x99, 0x71, 0xDB, 0x8A, 0x6B}}

type IBluetoothLEAdvertisementDataSectionFactoryInterface interface {
	win32.IInspectableInterface
	Create(dataType byte, data *IBuffer) *IBluetoothLEAdvertisementDataSection
}

type IBluetoothLEAdvertisementDataSectionFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IBluetoothLEAdvertisementDataSectionFactory struct {
	win32.IInspectable
}

func (this *IBluetoothLEAdvertisementDataSectionFactory) Vtbl() *IBluetoothLEAdvertisementDataSectionFactoryVtbl {
	return (*IBluetoothLEAdvertisementDataSectionFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAdvertisementDataSectionFactory) Create(dataType byte, data *IBuffer) *IBluetoothLEAdvertisementDataSection {
	var _result *IBluetoothLEAdvertisementDataSection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(dataType), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3BB6472F-0606-434B-A76E-74159F0684D3
var IID_IBluetoothLEAdvertisementDataTypesStatics = syscall.GUID{0x3BB6472F, 0x0606, 0x434B,
	[8]byte{0xA7, 0x6E, 0x74, 0x15, 0x9F, 0x06, 0x84, 0xD3}}

type IBluetoothLEAdvertisementDataTypesStaticsInterface interface {
	win32.IInspectableInterface
	Get_Flags() byte
	Get_IncompleteService16BitUuids() byte
	Get_CompleteService16BitUuids() byte
	Get_IncompleteService32BitUuids() byte
	Get_CompleteService32BitUuids() byte
	Get_IncompleteService128BitUuids() byte
	Get_CompleteService128BitUuids() byte
	Get_ShortenedLocalName() byte
	Get_CompleteLocalName() byte
	Get_TxPowerLevel() byte
	Get_PeripheralConnectionIntervalRange() byte
	Get_ServiceSolicitation16BitUuids() byte
	Get_ServiceSolicitation32BitUuids() byte
	Get_ServiceSolicitation128BitUuids() byte
	Get_ServiceData16BitUuids() byte
	Get_ServiceData32BitUuids() byte
	Get_ServiceData128BitUuids() byte
	Get_PublicTargetAddress() byte
	Get_RandomTargetAddress() byte
	Get_Appearance() byte
	Get_AdvertisingInterval() byte
	Get_ManufacturerSpecificData() byte
}

type IBluetoothLEAdvertisementDataTypesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Flags                             uintptr
	Get_IncompleteService16BitUuids       uintptr
	Get_CompleteService16BitUuids         uintptr
	Get_IncompleteService32BitUuids       uintptr
	Get_CompleteService32BitUuids         uintptr
	Get_IncompleteService128BitUuids      uintptr
	Get_CompleteService128BitUuids        uintptr
	Get_ShortenedLocalName                uintptr
	Get_CompleteLocalName                 uintptr
	Get_TxPowerLevel                      uintptr
	Get_PeripheralConnectionIntervalRange uintptr
	Get_ServiceSolicitation16BitUuids     uintptr
	Get_ServiceSolicitation32BitUuids     uintptr
	Get_ServiceSolicitation128BitUuids    uintptr
	Get_ServiceData16BitUuids             uintptr
	Get_ServiceData32BitUuids             uintptr
	Get_ServiceData128BitUuids            uintptr
	Get_PublicTargetAddress               uintptr
	Get_RandomTargetAddress               uintptr
	Get_Appearance                        uintptr
	Get_AdvertisingInterval               uintptr
	Get_ManufacturerSpecificData          uintptr
}

type IBluetoothLEAdvertisementDataTypesStatics struct {
	win32.IInspectable
}

func (this *IBluetoothLEAdvertisementDataTypesStatics) Vtbl() *IBluetoothLEAdvertisementDataTypesStaticsVtbl {
	return (*IBluetoothLEAdvertisementDataTypesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAdvertisementDataTypesStatics) Get_Flags() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Flags, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementDataTypesStatics) Get_IncompleteService16BitUuids() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IncompleteService16BitUuids, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementDataTypesStatics) Get_CompleteService16BitUuids() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CompleteService16BitUuids, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementDataTypesStatics) Get_IncompleteService32BitUuids() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IncompleteService32BitUuids, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementDataTypesStatics) Get_CompleteService32BitUuids() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CompleteService32BitUuids, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementDataTypesStatics) Get_IncompleteService128BitUuids() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IncompleteService128BitUuids, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementDataTypesStatics) Get_CompleteService128BitUuids() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CompleteService128BitUuids, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementDataTypesStatics) Get_ShortenedLocalName() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShortenedLocalName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementDataTypesStatics) Get_CompleteLocalName() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CompleteLocalName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementDataTypesStatics) Get_TxPowerLevel() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TxPowerLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementDataTypesStatics) Get_PeripheralConnectionIntervalRange() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PeripheralConnectionIntervalRange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementDataTypesStatics) Get_ServiceSolicitation16BitUuids() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceSolicitation16BitUuids, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementDataTypesStatics) Get_ServiceSolicitation32BitUuids() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceSolicitation32BitUuids, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementDataTypesStatics) Get_ServiceSolicitation128BitUuids() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceSolicitation128BitUuids, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementDataTypesStatics) Get_ServiceData16BitUuids() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceData16BitUuids, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementDataTypesStatics) Get_ServiceData32BitUuids() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceData32BitUuids, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementDataTypesStatics) Get_ServiceData128BitUuids() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceData128BitUuids, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementDataTypesStatics) Get_PublicTargetAddress() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PublicTargetAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementDataTypesStatics) Get_RandomTargetAddress() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RandomTargetAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementDataTypesStatics) Get_Appearance() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Appearance, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementDataTypesStatics) Get_AdvertisingInterval() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AdvertisingInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementDataTypesStatics) Get_ManufacturerSpecificData() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ManufacturerSpecificData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 131EB0D3-D04E-47B1-837E-49405BF6F80F
var IID_IBluetoothLEAdvertisementFilter = syscall.GUID{0x131EB0D3, 0xD04E, 0x47B1,
	[8]byte{0x83, 0x7E, 0x49, 0x40, 0x5B, 0xF6, 0xF8, 0x0F}}

type IBluetoothLEAdvertisementFilterInterface interface {
	win32.IInspectableInterface
	Get_Advertisement() *IBluetoothLEAdvertisement
	Put_Advertisement(value *IBluetoothLEAdvertisement)
	Get_BytePatterns() *IVector[*IBluetoothLEAdvertisementBytePattern]
}

type IBluetoothLEAdvertisementFilterVtbl struct {
	win32.IInspectableVtbl
	Get_Advertisement uintptr
	Put_Advertisement uintptr
	Get_BytePatterns  uintptr
}

type IBluetoothLEAdvertisementFilter struct {
	win32.IInspectable
}

func (this *IBluetoothLEAdvertisementFilter) Vtbl() *IBluetoothLEAdvertisementFilterVtbl {
	return (*IBluetoothLEAdvertisementFilterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAdvertisementFilter) Get_Advertisement() *IBluetoothLEAdvertisement {
	var _result *IBluetoothLEAdvertisement
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Advertisement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEAdvertisementFilter) Put_Advertisement(value *IBluetoothLEAdvertisement) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Advertisement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IBluetoothLEAdvertisementFilter) Get_BytePatterns() *IVector[*IBluetoothLEAdvertisementBytePattern] {
	var _result *IVector[*IBluetoothLEAdvertisementBytePattern]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BytePatterns, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CDE820F9-D9FA-43D6-A264-DDD8B7DA8B78
var IID_IBluetoothLEAdvertisementPublisher = syscall.GUID{0xCDE820F9, 0xD9FA, 0x43D6,
	[8]byte{0xA2, 0x64, 0xDD, 0xD8, 0xB7, 0xDA, 0x8B, 0x78}}

type IBluetoothLEAdvertisementPublisherInterface interface {
	win32.IInspectableInterface
	Get_Status() BluetoothLEAdvertisementPublisherStatus
	Get_Advertisement() *IBluetoothLEAdvertisement
	Start()
	Stop()
	Add_StatusChanged(handler TypedEventHandler[*IBluetoothLEAdvertisementPublisher, *IBluetoothLEAdvertisementPublisherStatusChangedEventArgs]) EventRegistrationToken
	Remove_StatusChanged(token EventRegistrationToken)
}

type IBluetoothLEAdvertisementPublisherVtbl struct {
	win32.IInspectableVtbl
	Get_Status           uintptr
	Get_Advertisement    uintptr
	Start                uintptr
	Stop                 uintptr
	Add_StatusChanged    uintptr
	Remove_StatusChanged uintptr
}

type IBluetoothLEAdvertisementPublisher struct {
	win32.IInspectable
}

func (this *IBluetoothLEAdvertisementPublisher) Vtbl() *IBluetoothLEAdvertisementPublisherVtbl {
	return (*IBluetoothLEAdvertisementPublisherVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAdvertisementPublisher) Get_Status() BluetoothLEAdvertisementPublisherStatus {
	var _result BluetoothLEAdvertisementPublisherStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementPublisher) Get_Advertisement() *IBluetoothLEAdvertisement {
	var _result *IBluetoothLEAdvertisement
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Advertisement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEAdvertisementPublisher) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IBluetoothLEAdvertisementPublisher) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IBluetoothLEAdvertisementPublisher) Add_StatusChanged(handler TypedEventHandler[*IBluetoothLEAdvertisementPublisher, *IBluetoothLEAdvertisementPublisherStatusChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StatusChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementPublisher) Remove_StatusChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StatusChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// FBDB545E-56F1-510F-A434-217FBD9E7BD2
var IID_IBluetoothLEAdvertisementPublisher2 = syscall.GUID{0xFBDB545E, 0x56F1, 0x510F,
	[8]byte{0xA4, 0x34, 0x21, 0x7F, 0xBD, 0x9E, 0x7B, 0xD2}}

type IBluetoothLEAdvertisementPublisher2Interface interface {
	win32.IInspectableInterface
	Get_PreferredTransmitPowerLevelInDBm() *IReference[int16]
	Put_PreferredTransmitPowerLevelInDBm(value *IReference[int16])
	Get_UseExtendedAdvertisement() bool
	Put_UseExtendedAdvertisement(value bool)
	Get_IsAnonymous() bool
	Put_IsAnonymous(value bool)
	Get_IncludeTransmitPowerLevel() bool
	Put_IncludeTransmitPowerLevel(value bool)
}

type IBluetoothLEAdvertisementPublisher2Vtbl struct {
	win32.IInspectableVtbl
	Get_PreferredTransmitPowerLevelInDBm uintptr
	Put_PreferredTransmitPowerLevelInDBm uintptr
	Get_UseExtendedAdvertisement         uintptr
	Put_UseExtendedAdvertisement         uintptr
	Get_IsAnonymous                      uintptr
	Put_IsAnonymous                      uintptr
	Get_IncludeTransmitPowerLevel        uintptr
	Put_IncludeTransmitPowerLevel        uintptr
}

type IBluetoothLEAdvertisementPublisher2 struct {
	win32.IInspectable
}

func (this *IBluetoothLEAdvertisementPublisher2) Vtbl() *IBluetoothLEAdvertisementPublisher2Vtbl {
	return (*IBluetoothLEAdvertisementPublisher2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAdvertisementPublisher2) Get_PreferredTransmitPowerLevelInDBm() *IReference[int16] {
	var _result *IReference[int16]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PreferredTransmitPowerLevelInDBm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEAdvertisementPublisher2) Put_PreferredTransmitPowerLevelInDBm(value *IReference[int16]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PreferredTransmitPowerLevelInDBm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IBluetoothLEAdvertisementPublisher2) Get_UseExtendedAdvertisement() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UseExtendedAdvertisement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementPublisher2) Put_UseExtendedAdvertisement(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_UseExtendedAdvertisement, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IBluetoothLEAdvertisementPublisher2) Get_IsAnonymous() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAnonymous, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementPublisher2) Put_IsAnonymous(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsAnonymous, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IBluetoothLEAdvertisementPublisher2) Get_IncludeTransmitPowerLevel() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IncludeTransmitPowerLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementPublisher2) Put_IncludeTransmitPowerLevel(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IncludeTransmitPowerLevel, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 5C5F065E-B863-4981-A1AF-1C544D8B0C0D
var IID_IBluetoothLEAdvertisementPublisherFactory = syscall.GUID{0x5C5F065E, 0xB863, 0x4981,
	[8]byte{0xA1, 0xAF, 0x1C, 0x54, 0x4D, 0x8B, 0x0C, 0x0D}}

type IBluetoothLEAdvertisementPublisherFactoryInterface interface {
	win32.IInspectableInterface
	Create(advertisement *IBluetoothLEAdvertisement) *IBluetoothLEAdvertisementPublisher
}

type IBluetoothLEAdvertisementPublisherFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IBluetoothLEAdvertisementPublisherFactory struct {
	win32.IInspectable
}

func (this *IBluetoothLEAdvertisementPublisherFactory) Vtbl() *IBluetoothLEAdvertisementPublisherFactoryVtbl {
	return (*IBluetoothLEAdvertisementPublisherFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAdvertisementPublisherFactory) Create(advertisement *IBluetoothLEAdvertisement) *IBluetoothLEAdvertisementPublisher {
	var _result *IBluetoothLEAdvertisementPublisher
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(advertisement)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 09C2BD9F-2DFF-4B23-86EE-0D14FB94AEAE
var IID_IBluetoothLEAdvertisementPublisherStatusChangedEventArgs = syscall.GUID{0x09C2BD9F, 0x2DFF, 0x4B23,
	[8]byte{0x86, 0xEE, 0x0D, 0x14, 0xFB, 0x94, 0xAE, 0xAE}}

type IBluetoothLEAdvertisementPublisherStatusChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Status() BluetoothLEAdvertisementPublisherStatus
	Get_Error() BluetoothError
}

type IBluetoothLEAdvertisementPublisherStatusChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
	Get_Error  uintptr
}

type IBluetoothLEAdvertisementPublisherStatusChangedEventArgs struct {
	win32.IInspectable
}

func (this *IBluetoothLEAdvertisementPublisherStatusChangedEventArgs) Vtbl() *IBluetoothLEAdvertisementPublisherStatusChangedEventArgsVtbl {
	return (*IBluetoothLEAdvertisementPublisherStatusChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAdvertisementPublisherStatusChangedEventArgs) Get_Status() BluetoothLEAdvertisementPublisherStatus {
	var _result BluetoothLEAdvertisementPublisherStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementPublisherStatusChangedEventArgs) Get_Error() BluetoothError {
	var _result BluetoothError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Error, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 8F62790E-DC88-5C8B-B34E-10B321850F88
var IID_IBluetoothLEAdvertisementPublisherStatusChangedEventArgs2 = syscall.GUID{0x8F62790E, 0xDC88, 0x5C8B,
	[8]byte{0xB3, 0x4E, 0x10, 0xB3, 0x21, 0x85, 0x0F, 0x88}}

type IBluetoothLEAdvertisementPublisherStatusChangedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_SelectedTransmitPowerLevelInDBm() *IReference[int16]
}

type IBluetoothLEAdvertisementPublisherStatusChangedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_SelectedTransmitPowerLevelInDBm uintptr
}

type IBluetoothLEAdvertisementPublisherStatusChangedEventArgs2 struct {
	win32.IInspectable
}

func (this *IBluetoothLEAdvertisementPublisherStatusChangedEventArgs2) Vtbl() *IBluetoothLEAdvertisementPublisherStatusChangedEventArgs2Vtbl {
	return (*IBluetoothLEAdvertisementPublisherStatusChangedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAdvertisementPublisherStatusChangedEventArgs2) Get_SelectedTransmitPowerLevelInDBm() *IReference[int16] {
	var _result *IReference[int16]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectedTransmitPowerLevelInDBm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 27987DDF-E596-41BE-8D43-9E6731D4A913
var IID_IBluetoothLEAdvertisementReceivedEventArgs = syscall.GUID{0x27987DDF, 0xE596, 0x41BE,
	[8]byte{0x8D, 0x43, 0x9E, 0x67, 0x31, 0xD4, 0xA9, 0x13}}

type IBluetoothLEAdvertisementReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_RawSignalStrengthInDBm() int16
	Get_BluetoothAddress() uint64
	Get_AdvertisementType() BluetoothLEAdvertisementType
	Get_Timestamp() DateTime
	Get_Advertisement() *IBluetoothLEAdvertisement
}

type IBluetoothLEAdvertisementReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_RawSignalStrengthInDBm uintptr
	Get_BluetoothAddress       uintptr
	Get_AdvertisementType      uintptr
	Get_Timestamp              uintptr
	Get_Advertisement          uintptr
}

type IBluetoothLEAdvertisementReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IBluetoothLEAdvertisementReceivedEventArgs) Vtbl() *IBluetoothLEAdvertisementReceivedEventArgsVtbl {
	return (*IBluetoothLEAdvertisementReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAdvertisementReceivedEventArgs) Get_RawSignalStrengthInDBm() int16 {
	var _result int16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RawSignalStrengthInDBm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementReceivedEventArgs) Get_BluetoothAddress() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BluetoothAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementReceivedEventArgs) Get_AdvertisementType() BluetoothLEAdvertisementType {
	var _result BluetoothLEAdvertisementType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AdvertisementType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementReceivedEventArgs) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementReceivedEventArgs) Get_Advertisement() *IBluetoothLEAdvertisement {
	var _result *IBluetoothLEAdvertisement
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Advertisement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 12D9C87B-0399-5F0E-A348-53B02B6B162E
var IID_IBluetoothLEAdvertisementReceivedEventArgs2 = syscall.GUID{0x12D9C87B, 0x0399, 0x5F0E,
	[8]byte{0xA3, 0x48, 0x53, 0xB0, 0x2B, 0x6B, 0x16, 0x2E}}

type IBluetoothLEAdvertisementReceivedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_BluetoothAddressType() BluetoothAddressType
	Get_TransmitPowerLevelInDBm() *IReference[int16]
	Get_IsAnonymous() bool
	Get_IsConnectable() bool
	Get_IsScannable() bool
	Get_IsDirected() bool
	Get_IsScanResponse() bool
}

type IBluetoothLEAdvertisementReceivedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_BluetoothAddressType    uintptr
	Get_TransmitPowerLevelInDBm uintptr
	Get_IsAnonymous             uintptr
	Get_IsConnectable           uintptr
	Get_IsScannable             uintptr
	Get_IsDirected              uintptr
	Get_IsScanResponse          uintptr
}

type IBluetoothLEAdvertisementReceivedEventArgs2 struct {
	win32.IInspectable
}

func (this *IBluetoothLEAdvertisementReceivedEventArgs2) Vtbl() *IBluetoothLEAdvertisementReceivedEventArgs2Vtbl {
	return (*IBluetoothLEAdvertisementReceivedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAdvertisementReceivedEventArgs2) Get_BluetoothAddressType() BluetoothAddressType {
	var _result BluetoothAddressType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BluetoothAddressType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementReceivedEventArgs2) Get_TransmitPowerLevelInDBm() *IReference[int16] {
	var _result *IReference[int16]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TransmitPowerLevelInDBm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEAdvertisementReceivedEventArgs2) Get_IsAnonymous() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAnonymous, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementReceivedEventArgs2) Get_IsConnectable() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsConnectable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementReceivedEventArgs2) Get_IsScannable() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsScannable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementReceivedEventArgs2) Get_IsDirected() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDirected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementReceivedEventArgs2) Get_IsScanResponse() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsScanResponse, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A6AC336F-F3D3-4297-8D6C-C81EA6623F40
var IID_IBluetoothLEAdvertisementWatcher = syscall.GUID{0xA6AC336F, 0xF3D3, 0x4297,
	[8]byte{0x8D, 0x6C, 0xC8, 0x1E, 0xA6, 0x62, 0x3F, 0x40}}

type IBluetoothLEAdvertisementWatcherInterface interface {
	win32.IInspectableInterface
	Get_MinSamplingInterval() TimeSpan
	Get_MaxSamplingInterval() TimeSpan
	Get_MinOutOfRangeTimeout() TimeSpan
	Get_MaxOutOfRangeTimeout() TimeSpan
	Get_Status() BluetoothLEAdvertisementWatcherStatus
	Get_ScanningMode() BluetoothLEScanningMode
	Put_ScanningMode(value BluetoothLEScanningMode)
	Get_SignalStrengthFilter() *IBluetoothSignalStrengthFilter
	Put_SignalStrengthFilter(value *IBluetoothSignalStrengthFilter)
	Get_AdvertisementFilter() *IBluetoothLEAdvertisementFilter
	Put_AdvertisementFilter(value *IBluetoothLEAdvertisementFilter)
	Start()
	Stop()
	Add_Received(handler TypedEventHandler[*IBluetoothLEAdvertisementWatcher, *IBluetoothLEAdvertisementReceivedEventArgs]) EventRegistrationToken
	Remove_Received(token EventRegistrationToken)
	Add_Stopped(handler TypedEventHandler[*IBluetoothLEAdvertisementWatcher, *IBluetoothLEAdvertisementWatcherStoppedEventArgs]) EventRegistrationToken
	Remove_Stopped(token EventRegistrationToken)
}

type IBluetoothLEAdvertisementWatcherVtbl struct {
	win32.IInspectableVtbl
	Get_MinSamplingInterval  uintptr
	Get_MaxSamplingInterval  uintptr
	Get_MinOutOfRangeTimeout uintptr
	Get_MaxOutOfRangeTimeout uintptr
	Get_Status               uintptr
	Get_ScanningMode         uintptr
	Put_ScanningMode         uintptr
	Get_SignalStrengthFilter uintptr
	Put_SignalStrengthFilter uintptr
	Get_AdvertisementFilter  uintptr
	Put_AdvertisementFilter  uintptr
	Start                    uintptr
	Stop                     uintptr
	Add_Received             uintptr
	Remove_Received          uintptr
	Add_Stopped              uintptr
	Remove_Stopped           uintptr
}

type IBluetoothLEAdvertisementWatcher struct {
	win32.IInspectable
}

func (this *IBluetoothLEAdvertisementWatcher) Vtbl() *IBluetoothLEAdvertisementWatcherVtbl {
	return (*IBluetoothLEAdvertisementWatcherVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAdvertisementWatcher) Get_MinSamplingInterval() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinSamplingInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementWatcher) Get_MaxSamplingInterval() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxSamplingInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementWatcher) Get_MinOutOfRangeTimeout() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinOutOfRangeTimeout, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementWatcher) Get_MaxOutOfRangeTimeout() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxOutOfRangeTimeout, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementWatcher) Get_Status() BluetoothLEAdvertisementWatcherStatus {
	var _result BluetoothLEAdvertisementWatcherStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementWatcher) Get_ScanningMode() BluetoothLEScanningMode {
	var _result BluetoothLEScanningMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScanningMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementWatcher) Put_ScanningMode(value BluetoothLEScanningMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ScanningMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IBluetoothLEAdvertisementWatcher) Get_SignalStrengthFilter() *IBluetoothSignalStrengthFilter {
	var _result *IBluetoothSignalStrengthFilter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SignalStrengthFilter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEAdvertisementWatcher) Put_SignalStrengthFilter(value *IBluetoothSignalStrengthFilter) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SignalStrengthFilter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IBluetoothLEAdvertisementWatcher) Get_AdvertisementFilter() *IBluetoothLEAdvertisementFilter {
	var _result *IBluetoothLEAdvertisementFilter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AdvertisementFilter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEAdvertisementWatcher) Put_AdvertisementFilter(value *IBluetoothLEAdvertisementFilter) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AdvertisementFilter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IBluetoothLEAdvertisementWatcher) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IBluetoothLEAdvertisementWatcher) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IBluetoothLEAdvertisementWatcher) Add_Received(handler TypedEventHandler[*IBluetoothLEAdvertisementWatcher, *IBluetoothLEAdvertisementReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Received, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementWatcher) Remove_Received(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Received, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IBluetoothLEAdvertisementWatcher) Add_Stopped(handler TypedEventHandler[*IBluetoothLEAdvertisementWatcher, *IBluetoothLEAdvertisementWatcherStoppedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Stopped, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementWatcher) Remove_Stopped(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Stopped, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 01BF26BC-B164-5805-90A3-E8A7997FF225
var IID_IBluetoothLEAdvertisementWatcher2 = syscall.GUID{0x01BF26BC, 0xB164, 0x5805,
	[8]byte{0x90, 0xA3, 0xE8, 0xA7, 0x99, 0x7F, 0xF2, 0x25}}

type IBluetoothLEAdvertisementWatcher2Interface interface {
	win32.IInspectableInterface
	Get_AllowExtendedAdvertisements() bool
	Put_AllowExtendedAdvertisements(value bool)
}

type IBluetoothLEAdvertisementWatcher2Vtbl struct {
	win32.IInspectableVtbl
	Get_AllowExtendedAdvertisements uintptr
	Put_AllowExtendedAdvertisements uintptr
}

type IBluetoothLEAdvertisementWatcher2 struct {
	win32.IInspectable
}

func (this *IBluetoothLEAdvertisementWatcher2) Vtbl() *IBluetoothLEAdvertisementWatcher2Vtbl {
	return (*IBluetoothLEAdvertisementWatcher2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAdvertisementWatcher2) Get_AllowExtendedAdvertisements() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllowExtendedAdvertisements, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementWatcher2) Put_AllowExtendedAdvertisements(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AllowExtendedAdvertisements, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 9AAF2D56-39AC-453E-B32A-85C657E017F1
var IID_IBluetoothLEAdvertisementWatcherFactory = syscall.GUID{0x9AAF2D56, 0x39AC, 0x453E,
	[8]byte{0xB3, 0x2A, 0x85, 0xC6, 0x57, 0xE0, 0x17, 0xF1}}

type IBluetoothLEAdvertisementWatcherFactoryInterface interface {
	win32.IInspectableInterface
	Create(advertisementFilter *IBluetoothLEAdvertisementFilter) *IBluetoothLEAdvertisementWatcher
}

type IBluetoothLEAdvertisementWatcherFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IBluetoothLEAdvertisementWatcherFactory struct {
	win32.IInspectable
}

func (this *IBluetoothLEAdvertisementWatcherFactory) Vtbl() *IBluetoothLEAdvertisementWatcherFactoryVtbl {
	return (*IBluetoothLEAdvertisementWatcherFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAdvertisementWatcherFactory) Create(advertisementFilter *IBluetoothLEAdvertisementFilter) *IBluetoothLEAdvertisementWatcher {
	var _result *IBluetoothLEAdvertisementWatcher
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(advertisementFilter)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DD40F84D-E7B9-43E3-9C04-0685D085FD8C
var IID_IBluetoothLEAdvertisementWatcherStoppedEventArgs = syscall.GUID{0xDD40F84D, 0xE7B9, 0x43E3,
	[8]byte{0x9C, 0x04, 0x06, 0x85, 0xD0, 0x85, 0xFD, 0x8C}}

type IBluetoothLEAdvertisementWatcherStoppedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Error() BluetoothError
}

type IBluetoothLEAdvertisementWatcherStoppedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Error uintptr
}

type IBluetoothLEAdvertisementWatcherStoppedEventArgs struct {
	win32.IInspectable
}

func (this *IBluetoothLEAdvertisementWatcherStoppedEventArgs) Vtbl() *IBluetoothLEAdvertisementWatcherStoppedEventArgsVtbl {
	return (*IBluetoothLEAdvertisementWatcherStoppedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAdvertisementWatcherStoppedEventArgs) Get_Error() BluetoothError {
	var _result BluetoothError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Error, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 912DBA18-6963-4533-B061-4694DAFB34E5
var IID_IBluetoothLEManufacturerData = syscall.GUID{0x912DBA18, 0x6963, 0x4533,
	[8]byte{0xB0, 0x61, 0x46, 0x94, 0xDA, 0xFB, 0x34, 0xE5}}

type IBluetoothLEManufacturerDataInterface interface {
	win32.IInspectableInterface
	Get_CompanyId() uint16
	Put_CompanyId(value uint16)
	Get_Data() *IBuffer
	Put_Data(value *IBuffer)
}

type IBluetoothLEManufacturerDataVtbl struct {
	win32.IInspectableVtbl
	Get_CompanyId uintptr
	Put_CompanyId uintptr
	Get_Data      uintptr
	Put_Data      uintptr
}

type IBluetoothLEManufacturerData struct {
	win32.IInspectable
}

func (this *IBluetoothLEManufacturerData) Vtbl() *IBluetoothLEManufacturerDataVtbl {
	return (*IBluetoothLEManufacturerDataVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEManufacturerData) Get_CompanyId() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CompanyId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEManufacturerData) Put_CompanyId(value uint16) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CompanyId, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IBluetoothLEManufacturerData) Get_Data() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEManufacturerData) Put_Data(value *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// C09B39F8-319A-441E-8DE5-66A81E877A6C
var IID_IBluetoothLEManufacturerDataFactory = syscall.GUID{0xC09B39F8, 0x319A, 0x441E,
	[8]byte{0x8D, 0xE5, 0x66, 0xA8, 0x1E, 0x87, 0x7A, 0x6C}}

type IBluetoothLEManufacturerDataFactoryInterface interface {
	win32.IInspectableInterface
	Create(companyId uint16, data *IBuffer) *IBluetoothLEManufacturerData
}

type IBluetoothLEManufacturerDataFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IBluetoothLEManufacturerDataFactory struct {
	win32.IInspectable
}

func (this *IBluetoothLEManufacturerDataFactory) Vtbl() *IBluetoothLEManufacturerDataFactoryVtbl {
	return (*IBluetoothLEManufacturerDataFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEManufacturerDataFactory) Create(companyId uint16, data *IBuffer) *IBluetoothLEManufacturerData {
	var _result *IBluetoothLEManufacturerData
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(companyId), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type BluetoothLEAdvertisement struct {
	RtClass
	*IBluetoothLEAdvertisement
}

func NewBluetoothLEAdvertisement() *BluetoothLEAdvertisement {
	hs := NewHStr("Windows.Devices.Bluetooth.Advertisement.BluetoothLEAdvertisement")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &BluetoothLEAdvertisement{
		RtClass:                   RtClass{PInspect: p},
		IBluetoothLEAdvertisement: (*IBluetoothLEAdvertisement)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type BluetoothLEAdvertisementBytePattern struct {
	RtClass
	*IBluetoothLEAdvertisementBytePattern
}

func NewBluetoothLEAdvertisementBytePattern() *BluetoothLEAdvertisementBytePattern {
	hs := NewHStr("Windows.Devices.Bluetooth.Advertisement.BluetoothLEAdvertisementBytePattern")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &BluetoothLEAdvertisementBytePattern{
		RtClass:                              RtClass{PInspect: p},
		IBluetoothLEAdvertisementBytePattern: (*IBluetoothLEAdvertisementBytePattern)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewBluetoothLEAdvertisementBytePattern_Create(dataType byte, offset int16, data *IBuffer) *BluetoothLEAdvertisementBytePattern {
	hs := NewHStr("Windows.Devices.Bluetooth.Advertisement.BluetoothLEAdvertisementBytePattern")
	var pFac *IBluetoothLEAdvertisementBytePatternFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IBluetoothLEAdvertisementBytePatternFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IBluetoothLEAdvertisementBytePattern
	p = pFac.Create(dataType, offset, data)
	result := &BluetoothLEAdvertisementBytePattern{
		RtClass:                              RtClass{PInspect: &p.IInspectable},
		IBluetoothLEAdvertisementBytePattern: p,
	}
	com.AddToScope(result)
	return result
}

type BluetoothLEAdvertisementDataSection struct {
	RtClass
	*IBluetoothLEAdvertisementDataSection
}

func NewBluetoothLEAdvertisementDataSection() *BluetoothLEAdvertisementDataSection {
	hs := NewHStr("Windows.Devices.Bluetooth.Advertisement.BluetoothLEAdvertisementDataSection")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &BluetoothLEAdvertisementDataSection{
		RtClass:                              RtClass{PInspect: p},
		IBluetoothLEAdvertisementDataSection: (*IBluetoothLEAdvertisementDataSection)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewBluetoothLEAdvertisementDataSection_Create(dataType byte, data *IBuffer) *BluetoothLEAdvertisementDataSection {
	hs := NewHStr("Windows.Devices.Bluetooth.Advertisement.BluetoothLEAdvertisementDataSection")
	var pFac *IBluetoothLEAdvertisementDataSectionFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IBluetoothLEAdvertisementDataSectionFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IBluetoothLEAdvertisementDataSection
	p = pFac.Create(dataType, data)
	result := &BluetoothLEAdvertisementDataSection{
		RtClass:                              RtClass{PInspect: &p.IInspectable},
		IBluetoothLEAdvertisementDataSection: p,
	}
	com.AddToScope(result)
	return result
}

type BluetoothLEAdvertisementDataTypes struct {
	RtClass
}

func NewIBluetoothLEAdvertisementDataTypesStatics() *IBluetoothLEAdvertisementDataTypesStatics {
	var p *IBluetoothLEAdvertisementDataTypesStatics
	hs := NewHStr("Windows.Devices.Bluetooth.Advertisement.BluetoothLEAdvertisementDataTypes")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IBluetoothLEAdvertisementDataTypesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type BluetoothLEAdvertisementFilter struct {
	RtClass
	*IBluetoothLEAdvertisementFilter
}

func NewBluetoothLEAdvertisementFilter() *BluetoothLEAdvertisementFilter {
	hs := NewHStr("Windows.Devices.Bluetooth.Advertisement.BluetoothLEAdvertisementFilter")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &BluetoothLEAdvertisementFilter{
		RtClass:                         RtClass{PInspect: p},
		IBluetoothLEAdvertisementFilter: (*IBluetoothLEAdvertisementFilter)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type BluetoothLEAdvertisementPublisher struct {
	RtClass
	*IBluetoothLEAdvertisementPublisher
}

func NewBluetoothLEAdvertisementPublisher() *BluetoothLEAdvertisementPublisher {
	hs := NewHStr("Windows.Devices.Bluetooth.Advertisement.BluetoothLEAdvertisementPublisher")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &BluetoothLEAdvertisementPublisher{
		RtClass:                            RtClass{PInspect: p},
		IBluetoothLEAdvertisementPublisher: (*IBluetoothLEAdvertisementPublisher)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewBluetoothLEAdvertisementPublisher_Create(advertisement *IBluetoothLEAdvertisement) *BluetoothLEAdvertisementPublisher {
	hs := NewHStr("Windows.Devices.Bluetooth.Advertisement.BluetoothLEAdvertisementPublisher")
	var pFac *IBluetoothLEAdvertisementPublisherFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IBluetoothLEAdvertisementPublisherFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IBluetoothLEAdvertisementPublisher
	p = pFac.Create(advertisement)
	result := &BluetoothLEAdvertisementPublisher{
		RtClass:                            RtClass{PInspect: &p.IInspectable},
		IBluetoothLEAdvertisementPublisher: p,
	}
	com.AddToScope(result)
	return result
}

type BluetoothLEAdvertisementPublisherStatusChangedEventArgs struct {
	RtClass
	*IBluetoothLEAdvertisementPublisherStatusChangedEventArgs
}

type BluetoothLEAdvertisementReceivedEventArgs struct {
	RtClass
	*IBluetoothLEAdvertisementReceivedEventArgs
}

type BluetoothLEAdvertisementWatcher struct {
	RtClass
	*IBluetoothLEAdvertisementWatcher
}

func NewBluetoothLEAdvertisementWatcher() *BluetoothLEAdvertisementWatcher {
	hs := NewHStr("Windows.Devices.Bluetooth.Advertisement.BluetoothLEAdvertisementWatcher")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &BluetoothLEAdvertisementWatcher{
		RtClass:                          RtClass{PInspect: p},
		IBluetoothLEAdvertisementWatcher: (*IBluetoothLEAdvertisementWatcher)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewBluetoothLEAdvertisementWatcher_Create(advertisementFilter *IBluetoothLEAdvertisementFilter) *BluetoothLEAdvertisementWatcher {
	hs := NewHStr("Windows.Devices.Bluetooth.Advertisement.BluetoothLEAdvertisementWatcher")
	var pFac *IBluetoothLEAdvertisementWatcherFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IBluetoothLEAdvertisementWatcherFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IBluetoothLEAdvertisementWatcher
	p = pFac.Create(advertisementFilter)
	result := &BluetoothLEAdvertisementWatcher{
		RtClass:                          RtClass{PInspect: &p.IInspectable},
		IBluetoothLEAdvertisementWatcher: p,
	}
	com.AddToScope(result)
	return result
}

type BluetoothLEAdvertisementWatcherStoppedEventArgs struct {
	RtClass
	*IBluetoothLEAdvertisementWatcherStoppedEventArgs
}

type BluetoothLEManufacturerData struct {
	RtClass
	*IBluetoothLEManufacturerData
}

func NewBluetoothLEManufacturerData() *BluetoothLEManufacturerData {
	hs := NewHStr("Windows.Devices.Bluetooth.Advertisement.BluetoothLEManufacturerData")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &BluetoothLEManufacturerData{
		RtClass:                      RtClass{PInspect: p},
		IBluetoothLEManufacturerData: (*IBluetoothLEManufacturerData)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewBluetoothLEManufacturerData_Create(companyId uint16, data *IBuffer) *BluetoothLEManufacturerData {
	hs := NewHStr("Windows.Devices.Bluetooth.Advertisement.BluetoothLEManufacturerData")
	var pFac *IBluetoothLEManufacturerDataFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IBluetoothLEManufacturerDataFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IBluetoothLEManufacturerData
	p = pFac.Create(companyId, data)
	result := &BluetoothLEManufacturerData{
		RtClass:                      RtClass{PInspect: &p.IInspectable},
		IBluetoothLEManufacturerData: p,
	}
	com.AddToScope(result)
	return result
}
