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
type AltitudeReferenceSystem int32

const (
	AltitudeReferenceSystem_Unspecified AltitudeReferenceSystem = 0
	AltitudeReferenceSystem_Terrain     AltitudeReferenceSystem = 1
	AltitudeReferenceSystem_Ellipsoid   AltitudeReferenceSystem = 2
	AltitudeReferenceSystem_Geoid       AltitudeReferenceSystem = 3
	AltitudeReferenceSystem_Surface     AltitudeReferenceSystem = 4
)

// enum
type GeolocationAccessStatus int32

const (
	GeolocationAccessStatus_Unspecified GeolocationAccessStatus = 0
	GeolocationAccessStatus_Allowed     GeolocationAccessStatus = 1
	GeolocationAccessStatus_Denied      GeolocationAccessStatus = 2
)

// enum
type GeoshapeType int32

const (
	GeoshapeType_Geopoint       GeoshapeType = 0
	GeoshapeType_Geocircle      GeoshapeType = 1
	GeoshapeType_Geopath        GeoshapeType = 2
	GeoshapeType_GeoboundingBox GeoshapeType = 3
)

// enum
type PositionAccuracy int32

const (
	PositionAccuracy_Default PositionAccuracy = 0
	PositionAccuracy_High    PositionAccuracy = 1
)

// enum
type PositionSource int32

const (
	PositionSource_Cellular   PositionSource = 0
	PositionSource_Satellite  PositionSource = 1
	PositionSource_WiFi       PositionSource = 2
	PositionSource_IPAddress  PositionSource = 3
	PositionSource_Unknown    PositionSource = 4
	PositionSource_Default    PositionSource = 5
	PositionSource_Obfuscated PositionSource = 6
)

// enum
type PositionStatus int32

const (
	PositionStatus_Ready          PositionStatus = 0
	PositionStatus_Initializing   PositionStatus = 1
	PositionStatus_NoData         PositionStatus = 2
	PositionStatus_Disabled       PositionStatus = 3
	PositionStatus_NotInitialized PositionStatus = 4
	PositionStatus_NotAvailable   PositionStatus = 5
)

// enum
type VisitMonitoringScope int32

const (
	VisitMonitoringScope_Venue VisitMonitoringScope = 0
	VisitMonitoringScope_City  VisitMonitoringScope = 1
)

// enum
type VisitStateChange int32

const (
	VisitStateChange_TrackingLost  VisitStateChange = 0
	VisitStateChange_Arrived       VisitStateChange = 1
	VisitStateChange_Departed      VisitStateChange = 2
	VisitStateChange_OtherMovement VisitStateChange = 3
)

// structs

type BasicGeoposition struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
}

// interfaces

// A8567A1A-64F4-4D48-BCEA-F6B008ECA34C
var IID_ICivicAddress = syscall.GUID{0xA8567A1A, 0x64F4, 0x4D48,
	[8]byte{0xBC, 0xEA, 0xF6, 0xB0, 0x08, 0xEC, 0xA3, 0x4C}}

type ICivicAddressInterface interface {
	win32.IInspectableInterface
	Get_Country() string
	Get_State() string
	Get_City() string
	Get_PostalCode() string
	Get_Timestamp() DateTime
}

type ICivicAddressVtbl struct {
	win32.IInspectableVtbl
	Get_Country    uintptr
	Get_State      uintptr
	Get_City       uintptr
	Get_PostalCode uintptr
	Get_Timestamp  uintptr
}

type ICivicAddress struct {
	win32.IInspectable
}

func (this *ICivicAddress) Vtbl() *ICivicAddressVtbl {
	return (*ICivicAddressVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICivicAddress) Get_Country() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Country, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICivicAddress) Get_State() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICivicAddress) Get_City() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_City, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICivicAddress) Get_PostalCode() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PostalCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICivicAddress) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 0896C80B-274F-43DA-9A06-CBFCDAEB4EC2
var IID_IGeoboundingBox = syscall.GUID{0x0896C80B, 0x274F, 0x43DA,
	[8]byte{0x9A, 0x06, 0xCB, 0xFC, 0xDA, 0xEB, 0x4E, 0xC2}}

type IGeoboundingBoxInterface interface {
	win32.IInspectableInterface
	Get_NorthwestCorner() BasicGeoposition
	Get_SoutheastCorner() BasicGeoposition
	Get_Center() BasicGeoposition
	Get_MinAltitude() float64
	Get_MaxAltitude() float64
}

type IGeoboundingBoxVtbl struct {
	win32.IInspectableVtbl
	Get_NorthwestCorner uintptr
	Get_SoutheastCorner uintptr
	Get_Center          uintptr
	Get_MinAltitude     uintptr
	Get_MaxAltitude     uintptr
}

type IGeoboundingBox struct {
	win32.IInspectable
}

func (this *IGeoboundingBox) Vtbl() *IGeoboundingBoxVtbl {
	return (*IGeoboundingBoxVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeoboundingBox) Get_NorthwestCorner() BasicGeoposition {
	var _result BasicGeoposition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NorthwestCorner, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGeoboundingBox) Get_SoutheastCorner() BasicGeoposition {
	var _result BasicGeoposition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SoutheastCorner, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGeoboundingBox) Get_Center() BasicGeoposition {
	var _result BasicGeoposition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Center, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGeoboundingBox) Get_MinAltitude() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinAltitude, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGeoboundingBox) Get_MaxAltitude() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxAltitude, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4DFBA589-0411-4ABC-B3B5-5BBCCB57D98C
var IID_IGeoboundingBoxFactory = syscall.GUID{0x4DFBA589, 0x0411, 0x4ABC,
	[8]byte{0xB3, 0xB5, 0x5B, 0xBC, 0xCB, 0x57, 0xD9, 0x8C}}

type IGeoboundingBoxFactoryInterface interface {
	win32.IInspectableInterface
	Create(northwestCorner BasicGeoposition, southeastCorner BasicGeoposition) *IGeoboundingBox
	CreateWithAltitudeReference(northwestCorner BasicGeoposition, southeastCorner BasicGeoposition, altitudeReferenceSystem AltitudeReferenceSystem) *IGeoboundingBox
	CreateWithAltitudeReferenceAndSpatialReference(northwestCorner BasicGeoposition, southeastCorner BasicGeoposition, altitudeReferenceSystem AltitudeReferenceSystem, spatialReferenceId uint32) *IGeoboundingBox
}

type IGeoboundingBoxFactoryVtbl struct {
	win32.IInspectableVtbl
	Create                                         uintptr
	CreateWithAltitudeReference                    uintptr
	CreateWithAltitudeReferenceAndSpatialReference uintptr
}

type IGeoboundingBoxFactory struct {
	win32.IInspectable
}

func (this *IGeoboundingBoxFactory) Vtbl() *IGeoboundingBoxFactoryVtbl {
	return (*IGeoboundingBoxFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeoboundingBoxFactory) Create(northwestCorner BasicGeoposition, southeastCorner BasicGeoposition) *IGeoboundingBox {
	var _result *IGeoboundingBox
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&northwestCorner)), uintptr(unsafe.Pointer(&southeastCorner)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeoboundingBoxFactory) CreateWithAltitudeReference(northwestCorner BasicGeoposition, southeastCorner BasicGeoposition, altitudeReferenceSystem AltitudeReferenceSystem) *IGeoboundingBox {
	var _result *IGeoboundingBox
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithAltitudeReference, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&northwestCorner)), uintptr(unsafe.Pointer(&southeastCorner)), uintptr(altitudeReferenceSystem), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeoboundingBoxFactory) CreateWithAltitudeReferenceAndSpatialReference(northwestCorner BasicGeoposition, southeastCorner BasicGeoposition, altitudeReferenceSystem AltitudeReferenceSystem, spatialReferenceId uint32) *IGeoboundingBox {
	var _result *IGeoboundingBox
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithAltitudeReferenceAndSpatialReference, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&northwestCorner)), uintptr(unsafe.Pointer(&southeastCorner)), uintptr(altitudeReferenceSystem), uintptr(spatialReferenceId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 67B80708-E61A-4CD0-841B-93233792B5CA
var IID_IGeoboundingBoxStatics = syscall.GUID{0x67B80708, 0xE61A, 0x4CD0,
	[8]byte{0x84, 0x1B, 0x93, 0x23, 0x37, 0x92, 0xB5, 0xCA}}

type IGeoboundingBoxStaticsInterface interface {
	win32.IInspectableInterface
	TryCompute(positions *IIterable[BasicGeoposition]) *IGeoboundingBox
	TryComputeWithAltitudeReference(positions *IIterable[BasicGeoposition], altitudeRefSystem AltitudeReferenceSystem) *IGeoboundingBox
	TryComputeWithAltitudeReferenceAndSpatialReference(positions *IIterable[BasicGeoposition], altitudeRefSystem AltitudeReferenceSystem, spatialReferenceId uint32) *IGeoboundingBox
}

type IGeoboundingBoxStaticsVtbl struct {
	win32.IInspectableVtbl
	TryCompute                                         uintptr
	TryComputeWithAltitudeReference                    uintptr
	TryComputeWithAltitudeReferenceAndSpatialReference uintptr
}

type IGeoboundingBoxStatics struct {
	win32.IInspectable
}

func (this *IGeoboundingBoxStatics) Vtbl() *IGeoboundingBoxStaticsVtbl {
	return (*IGeoboundingBoxStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeoboundingBoxStatics) TryCompute(positions *IIterable[BasicGeoposition]) *IGeoboundingBox {
	var _result *IGeoboundingBox
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryCompute, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(positions)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeoboundingBoxStatics) TryComputeWithAltitudeReference(positions *IIterable[BasicGeoposition], altitudeRefSystem AltitudeReferenceSystem) *IGeoboundingBox {
	var _result *IGeoboundingBox
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryComputeWithAltitudeReference, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(positions)), uintptr(altitudeRefSystem), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeoboundingBoxStatics) TryComputeWithAltitudeReferenceAndSpatialReference(positions *IIterable[BasicGeoposition], altitudeRefSystem AltitudeReferenceSystem, spatialReferenceId uint32) *IGeoboundingBox {
	var _result *IGeoboundingBox
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryComputeWithAltitudeReferenceAndSpatialReference, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(positions)), uintptr(altitudeRefSystem), uintptr(spatialReferenceId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 39E45843-A7F9-4E63-92A7-BA0C28D124B1
var IID_IGeocircle = syscall.GUID{0x39E45843, 0xA7F9, 0x4E63,
	[8]byte{0x92, 0xA7, 0xBA, 0x0C, 0x28, 0xD1, 0x24, 0xB1}}

type IGeocircleInterface interface {
	win32.IInspectableInterface
	Get_Center() BasicGeoposition
	Get_Radius() float64
}

type IGeocircleVtbl struct {
	win32.IInspectableVtbl
	Get_Center uintptr
	Get_Radius uintptr
}

type IGeocircle struct {
	win32.IInspectable
}

func (this *IGeocircle) Vtbl() *IGeocircleVtbl {
	return (*IGeocircleVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeocircle) Get_Center() BasicGeoposition {
	var _result BasicGeoposition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Center, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGeocircle) Get_Radius() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Radius, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// AFD6531F-72B1-4F7D-87CC-4ED4C9849C05
var IID_IGeocircleFactory = syscall.GUID{0xAFD6531F, 0x72B1, 0x4F7D,
	[8]byte{0x87, 0xCC, 0x4E, 0xD4, 0xC9, 0x84, 0x9C, 0x05}}

type IGeocircleFactoryInterface interface {
	win32.IInspectableInterface
	Create(position BasicGeoposition, radius float64) *IGeocircle
	CreateWithAltitudeReferenceSystem(position BasicGeoposition, radius float64, altitudeReferenceSystem AltitudeReferenceSystem) *IGeocircle
	CreateWithAltitudeReferenceSystemAndSpatialReferenceId(position BasicGeoposition, radius float64, altitudeReferenceSystem AltitudeReferenceSystem, spatialReferenceId uint32) *IGeocircle
}

type IGeocircleFactoryVtbl struct {
	win32.IInspectableVtbl
	Create                                                 uintptr
	CreateWithAltitudeReferenceSystem                      uintptr
	CreateWithAltitudeReferenceSystemAndSpatialReferenceId uintptr
}

type IGeocircleFactory struct {
	win32.IInspectable
}

func (this *IGeocircleFactory) Vtbl() *IGeocircleFactoryVtbl {
	return (*IGeocircleFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeocircleFactory) Create(position BasicGeoposition, radius float64) *IGeocircle {
	var _result *IGeocircle
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&position)), uintptr(radius), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeocircleFactory) CreateWithAltitudeReferenceSystem(position BasicGeoposition, radius float64, altitudeReferenceSystem AltitudeReferenceSystem) *IGeocircle {
	var _result *IGeocircle
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithAltitudeReferenceSystem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&position)), uintptr(radius), uintptr(altitudeReferenceSystem), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeocircleFactory) CreateWithAltitudeReferenceSystemAndSpatialReferenceId(position BasicGeoposition, radius float64, altitudeReferenceSystem AltitudeReferenceSystem, spatialReferenceId uint32) *IGeocircle {
	var _result *IGeocircle
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithAltitudeReferenceSystemAndSpatialReferenceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&position)), uintptr(radius), uintptr(altitudeReferenceSystem), uintptr(spatialReferenceId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EE21A3AA-976A-4C70-803D-083EA55BCBC4
var IID_IGeocoordinate = syscall.GUID{0xEE21A3AA, 0x976A, 0x4C70,
	[8]byte{0x80, 0x3D, 0x08, 0x3E, 0xA5, 0x5B, 0xCB, 0xC4}}

type IGeocoordinateInterface interface {
	win32.IInspectableInterface
	Get_Latitude() float64
	Get_Longitude() float64
	Get_Altitude() *IReference[float64]
	Get_Accuracy() float64
	Get_AltitudeAccuracy() *IReference[float64]
	Get_Heading() *IReference[float64]
	Get_Speed() *IReference[float64]
	Get_Timestamp() DateTime
}

type IGeocoordinateVtbl struct {
	win32.IInspectableVtbl
	Get_Latitude         uintptr
	Get_Longitude        uintptr
	Get_Altitude         uintptr
	Get_Accuracy         uintptr
	Get_AltitudeAccuracy uintptr
	Get_Heading          uintptr
	Get_Speed            uintptr
	Get_Timestamp        uintptr
}

type IGeocoordinate struct {
	win32.IInspectable
}

func (this *IGeocoordinate) Vtbl() *IGeocoordinateVtbl {
	return (*IGeocoordinateVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeocoordinate) Get_Latitude() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Latitude, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGeocoordinate) Get_Longitude() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Longitude, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGeocoordinate) Get_Altitude() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Altitude, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeocoordinate) Get_Accuracy() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Accuracy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGeocoordinate) Get_AltitudeAccuracy() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AltitudeAccuracy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeocoordinate) Get_Heading() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Heading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeocoordinate) Get_Speed() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Speed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeocoordinate) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C32A74D9-2608-474C-912C-06DD490F4AF7
var IID_IGeocoordinateSatelliteData = syscall.GUID{0xC32A74D9, 0x2608, 0x474C,
	[8]byte{0x91, 0x2C, 0x06, 0xDD, 0x49, 0x0F, 0x4A, 0xF7}}

type IGeocoordinateSatelliteDataInterface interface {
	win32.IInspectableInterface
	Get_PositionDilutionOfPrecision() *IReference[float64]
	Get_HorizontalDilutionOfPrecision() *IReference[float64]
	Get_VerticalDilutionOfPrecision() *IReference[float64]
}

type IGeocoordinateSatelliteDataVtbl struct {
	win32.IInspectableVtbl
	Get_PositionDilutionOfPrecision   uintptr
	Get_HorizontalDilutionOfPrecision uintptr
	Get_VerticalDilutionOfPrecision   uintptr
}

type IGeocoordinateSatelliteData struct {
	win32.IInspectable
}

func (this *IGeocoordinateSatelliteData) Vtbl() *IGeocoordinateSatelliteDataVtbl {
	return (*IGeocoordinateSatelliteDataVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeocoordinateSatelliteData) Get_PositionDilutionOfPrecision() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PositionDilutionOfPrecision, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeocoordinateSatelliteData) Get_HorizontalDilutionOfPrecision() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HorizontalDilutionOfPrecision, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeocoordinateSatelliteData) Get_VerticalDilutionOfPrecision() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VerticalDilutionOfPrecision, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 761C8CFD-A19D-5A51-80F5-71676115483E
var IID_IGeocoordinateSatelliteData2 = syscall.GUID{0x761C8CFD, 0xA19D, 0x5A51,
	[8]byte{0x80, 0xF5, 0x71, 0x67, 0x61, 0x15, 0x48, 0x3E}}

type IGeocoordinateSatelliteData2Interface interface {
	win32.IInspectableInterface
	Get_GeometricDilutionOfPrecision() *IReference[float64]
	Get_TimeDilutionOfPrecision() *IReference[float64]
}

type IGeocoordinateSatelliteData2Vtbl struct {
	win32.IInspectableVtbl
	Get_GeometricDilutionOfPrecision uintptr
	Get_TimeDilutionOfPrecision      uintptr
}

type IGeocoordinateSatelliteData2 struct {
	win32.IInspectable
}

func (this *IGeocoordinateSatelliteData2) Vtbl() *IGeocoordinateSatelliteData2Vtbl {
	return (*IGeocoordinateSatelliteData2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeocoordinateSatelliteData2) Get_GeometricDilutionOfPrecision() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GeometricDilutionOfPrecision, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeocoordinateSatelliteData2) Get_TimeDilutionOfPrecision() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimeDilutionOfPrecision, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FEEA0525-D22C-4D46-B527-0B96066FC7DB
var IID_IGeocoordinateWithPoint = syscall.GUID{0xFEEA0525, 0xD22C, 0x4D46,
	[8]byte{0xB5, 0x27, 0x0B, 0x96, 0x06, 0x6F, 0xC7, 0xDB}}

type IGeocoordinateWithPointInterface interface {
	win32.IInspectableInterface
	Get_Point() *IGeopoint
}

type IGeocoordinateWithPointVtbl struct {
	win32.IInspectableVtbl
	Get_Point uintptr
}

type IGeocoordinateWithPoint struct {
	win32.IInspectable
}

func (this *IGeocoordinateWithPoint) Vtbl() *IGeocoordinateWithPointVtbl {
	return (*IGeocoordinateWithPointVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeocoordinateWithPoint) Get_Point() *IGeopoint {
	var _result *IGeopoint
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Point, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 95E634BE-DBD6-40AC-B8F2-A65C0340D9A6
var IID_IGeocoordinateWithPositionData = syscall.GUID{0x95E634BE, 0xDBD6, 0x40AC,
	[8]byte{0xB8, 0xF2, 0xA6, 0x5C, 0x03, 0x40, 0xD9, 0xA6}}

type IGeocoordinateWithPositionDataInterface interface {
	win32.IInspectableInterface
	Get_PositionSource() PositionSource
	Get_SatelliteData() *IGeocoordinateSatelliteData
}

type IGeocoordinateWithPositionDataVtbl struct {
	win32.IInspectableVtbl
	Get_PositionSource uintptr
	Get_SatelliteData  uintptr
}

type IGeocoordinateWithPositionData struct {
	win32.IInspectable
}

func (this *IGeocoordinateWithPositionData) Vtbl() *IGeocoordinateWithPositionDataVtbl {
	return (*IGeocoordinateWithPositionDataVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeocoordinateWithPositionData) Get_PositionSource() PositionSource {
	var _result PositionSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PositionSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGeocoordinateWithPositionData) Get_SatelliteData() *IGeocoordinateSatelliteData {
	var _result *IGeocoordinateSatelliteData
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SatelliteData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8543FC02-C9F1-4610-AFE0-8BC3A6A87036
var IID_IGeocoordinateWithPositionSourceTimestamp = syscall.GUID{0x8543FC02, 0xC9F1, 0x4610,
	[8]byte{0xAF, 0xE0, 0x8B, 0xC3, 0xA6, 0xA8, 0x70, 0x36}}

type IGeocoordinateWithPositionSourceTimestampInterface interface {
	win32.IInspectableInterface
	Get_PositionSourceTimestamp() *IReference[DateTime]
}

type IGeocoordinateWithPositionSourceTimestampVtbl struct {
	win32.IInspectableVtbl
	Get_PositionSourceTimestamp uintptr
}

type IGeocoordinateWithPositionSourceTimestamp struct {
	win32.IInspectable
}

func (this *IGeocoordinateWithPositionSourceTimestamp) Vtbl() *IGeocoordinateWithPositionSourceTimestampVtbl {
	return (*IGeocoordinateWithPositionSourceTimestampVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeocoordinateWithPositionSourceTimestamp) Get_PositionSourceTimestamp() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PositionSourceTimestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 397CEBD7-EE38-5F3B-8900-C4A7BC9CF953
var IID_IGeocoordinateWithRemoteSource = syscall.GUID{0x397CEBD7, 0xEE38, 0x5F3B,
	[8]byte{0x89, 0x00, 0xC4, 0xA7, 0xBC, 0x9C, 0xF9, 0x53}}

type IGeocoordinateWithRemoteSourceInterface interface {
	win32.IInspectableInterface
	Get_IsRemoteSource() bool
}

type IGeocoordinateWithRemoteSourceVtbl struct {
	win32.IInspectableVtbl
	Get_IsRemoteSource uintptr
}

type IGeocoordinateWithRemoteSource struct {
	win32.IInspectable
}

func (this *IGeocoordinateWithRemoteSource) Vtbl() *IGeocoordinateWithRemoteSourceVtbl {
	return (*IGeocoordinateWithRemoteSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeocoordinateWithRemoteSource) Get_IsRemoteSource() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsRemoteSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A9C3BF62-4524-4989-8AA9-DE019D2E551F
var IID_IGeolocator = syscall.GUID{0xA9C3BF62, 0x4524, 0x4989,
	[8]byte{0x8A, 0xA9, 0xDE, 0x01, 0x9D, 0x2E, 0x55, 0x1F}}

type IGeolocatorInterface interface {
	win32.IInspectableInterface
	Get_DesiredAccuracy() PositionAccuracy
	Put_DesiredAccuracy(value PositionAccuracy)
	Get_MovementThreshold() float64
	Put_MovementThreshold(value float64)
	Get_ReportInterval() uint32
	Put_ReportInterval(value uint32)
	Get_LocationStatus() PositionStatus
	GetGeopositionAsync() *IAsyncOperation[*IGeoposition]
	GetGeopositionAsyncWithAgeAndTimeout(maximumAge TimeSpan, timeout TimeSpan) *IAsyncOperation[*IGeoposition]
	Add_PositionChanged(handler TypedEventHandler[*IGeolocator, *IPositionChangedEventArgs]) EventRegistrationToken
	Remove_PositionChanged(token EventRegistrationToken)
	Add_StatusChanged(handler TypedEventHandler[*IGeolocator, *IStatusChangedEventArgs]) EventRegistrationToken
	Remove_StatusChanged(token EventRegistrationToken)
}

type IGeolocatorVtbl struct {
	win32.IInspectableVtbl
	Get_DesiredAccuracy                  uintptr
	Put_DesiredAccuracy                  uintptr
	Get_MovementThreshold                uintptr
	Put_MovementThreshold                uintptr
	Get_ReportInterval                   uintptr
	Put_ReportInterval                   uintptr
	Get_LocationStatus                   uintptr
	GetGeopositionAsync                  uintptr
	GetGeopositionAsyncWithAgeAndTimeout uintptr
	Add_PositionChanged                  uintptr
	Remove_PositionChanged               uintptr
	Add_StatusChanged                    uintptr
	Remove_StatusChanged                 uintptr
}

type IGeolocator struct {
	win32.IInspectable
}

func (this *IGeolocator) Vtbl() *IGeolocatorVtbl {
	return (*IGeolocatorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeolocator) Get_DesiredAccuracy() PositionAccuracy {
	var _result PositionAccuracy
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesiredAccuracy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGeolocator) Put_DesiredAccuracy(value PositionAccuracy) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DesiredAccuracy, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGeolocator) Get_MovementThreshold() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MovementThreshold, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGeolocator) Put_MovementThreshold(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MovementThreshold, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGeolocator) Get_ReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGeolocator) Put_ReportInterval(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReportInterval, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGeolocator) Get_LocationStatus() PositionStatus {
	var _result PositionStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocationStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGeolocator) GetGeopositionAsync() *IAsyncOperation[*IGeoposition] {
	var _result *IAsyncOperation[*IGeoposition]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetGeopositionAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeolocator) GetGeopositionAsyncWithAgeAndTimeout(maximumAge TimeSpan, timeout TimeSpan) *IAsyncOperation[*IGeoposition] {
	var _result *IAsyncOperation[*IGeoposition]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetGeopositionAsyncWithAgeAndTimeout, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&maximumAge)), *(*uintptr)(unsafe.Pointer(&timeout)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeolocator) Add_PositionChanged(handler TypedEventHandler[*IGeolocator, *IPositionChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PositionChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGeolocator) Remove_PositionChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PositionChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGeolocator) Add_StatusChanged(handler TypedEventHandler[*IGeolocator, *IStatusChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StatusChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGeolocator) Remove_StatusChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StatusChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// D1B42E6D-8891-43B4-AD36-27C6FE9A97B1
var IID_IGeolocator2 = syscall.GUID{0xD1B42E6D, 0x8891, 0x43B4,
	[8]byte{0xAD, 0x36, 0x27, 0xC6, 0xFE, 0x9A, 0x97, 0xB1}}

type IGeolocator2Interface interface {
	win32.IInspectableInterface
	AllowFallbackToConsentlessPositions()
}

type IGeolocator2Vtbl struct {
	win32.IInspectableVtbl
	AllowFallbackToConsentlessPositions uintptr
}

type IGeolocator2 struct {
	win32.IInspectable
}

func (this *IGeolocator2) Vtbl() *IGeolocator2Vtbl {
	return (*IGeolocator2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeolocator2) AllowFallbackToConsentlessPositions() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AllowFallbackToConsentlessPositions, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 9A8E7571-2DF5-4591-9F87-EB5FD894E9B7
var IID_IGeolocatorStatics = syscall.GUID{0x9A8E7571, 0x2DF5, 0x4591,
	[8]byte{0x9F, 0x87, 0xEB, 0x5F, 0xD8, 0x94, 0xE9, 0xB7}}

type IGeolocatorStaticsInterface interface {
	win32.IInspectableInterface
	RequestAccessAsync() *IAsyncOperation[GeolocationAccessStatus]
	GetGeopositionHistoryAsync(startTime DateTime) *IAsyncOperation[*IVectorView[*IGeoposition]]
	GetGeopositionHistoryWithDurationAsync(startTime DateTime, duration TimeSpan) *IAsyncOperation[*IVectorView[*IGeoposition]]
}

type IGeolocatorStaticsVtbl struct {
	win32.IInspectableVtbl
	RequestAccessAsync                     uintptr
	GetGeopositionHistoryAsync             uintptr
	GetGeopositionHistoryWithDurationAsync uintptr
}

type IGeolocatorStatics struct {
	win32.IInspectable
}

func (this *IGeolocatorStatics) Vtbl() *IGeolocatorStaticsVtbl {
	return (*IGeolocatorStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeolocatorStatics) RequestAccessAsync() *IAsyncOperation[GeolocationAccessStatus] {
	var _result *IAsyncOperation[GeolocationAccessStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeolocatorStatics) GetGeopositionHistoryAsync(startTime DateTime) *IAsyncOperation[*IVectorView[*IGeoposition]] {
	var _result *IAsyncOperation[*IVectorView[*IGeoposition]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetGeopositionHistoryAsync, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&startTime)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeolocatorStatics) GetGeopositionHistoryWithDurationAsync(startTime DateTime, duration TimeSpan) *IAsyncOperation[*IVectorView[*IGeoposition]] {
	var _result *IAsyncOperation[*IVectorView[*IGeoposition]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetGeopositionHistoryWithDurationAsync, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&startTime)), *(*uintptr)(unsafe.Pointer(&duration)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 993011A2-FA1C-4631-A71D-0DBEB1250D9C
var IID_IGeolocatorStatics2 = syscall.GUID{0x993011A2, 0xFA1C, 0x4631,
	[8]byte{0xA7, 0x1D, 0x0D, 0xBE, 0xB1, 0x25, 0x0D, 0x9C}}

type IGeolocatorStatics2Interface interface {
	win32.IInspectableInterface
	Get_IsDefaultGeopositionRecommended() bool
	Put_DefaultGeoposition(value *IReference[BasicGeoposition])
	Get_DefaultGeoposition() *IReference[BasicGeoposition]
}

type IGeolocatorStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsDefaultGeopositionRecommended uintptr
	Put_DefaultGeoposition              uintptr
	Get_DefaultGeoposition              uintptr
}

type IGeolocatorStatics2 struct {
	win32.IInspectable
}

func (this *IGeolocatorStatics2) Vtbl() *IGeolocatorStatics2Vtbl {
	return (*IGeolocatorStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeolocatorStatics2) Get_IsDefaultGeopositionRecommended() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDefaultGeopositionRecommended, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGeolocatorStatics2) Put_DefaultGeoposition(value *IReference[BasicGeoposition]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DefaultGeoposition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IGeolocatorStatics2) Get_DefaultGeoposition() *IReference[BasicGeoposition] {
	var _result *IReference[BasicGeoposition]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DefaultGeoposition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 96F5D3C1-B80F-460A-994D-A96C47A51AA4
var IID_IGeolocatorWithScalarAccuracy = syscall.GUID{0x96F5D3C1, 0xB80F, 0x460A,
	[8]byte{0x99, 0x4D, 0xA9, 0x6C, 0x47, 0xA5, 0x1A, 0xA4}}

type IGeolocatorWithScalarAccuracyInterface interface {
	win32.IInspectableInterface
	Get_DesiredAccuracyInMeters() *IReference[uint32]
	Put_DesiredAccuracyInMeters(value *IReference[uint32])
}

type IGeolocatorWithScalarAccuracyVtbl struct {
	win32.IInspectableVtbl
	Get_DesiredAccuracyInMeters uintptr
	Put_DesiredAccuracyInMeters uintptr
}

type IGeolocatorWithScalarAccuracy struct {
	win32.IInspectable
}

func (this *IGeolocatorWithScalarAccuracy) Vtbl() *IGeolocatorWithScalarAccuracyVtbl {
	return (*IGeolocatorWithScalarAccuracyVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeolocatorWithScalarAccuracy) Get_DesiredAccuracyInMeters() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesiredAccuracyInMeters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeolocatorWithScalarAccuracy) Put_DesiredAccuracyInMeters(value *IReference[uint32]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DesiredAccuracyInMeters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// E53FD7B9-2DA4-4714-A652-DE8593289898
var IID_IGeopath = syscall.GUID{0xE53FD7B9, 0x2DA4, 0x4714,
	[8]byte{0xA6, 0x52, 0xDE, 0x85, 0x93, 0x28, 0x98, 0x98}}

type IGeopathInterface interface {
	win32.IInspectableInterface
	Get_Positions() *IVectorView[BasicGeoposition]
}

type IGeopathVtbl struct {
	win32.IInspectableVtbl
	Get_Positions uintptr
}

type IGeopath struct {
	win32.IInspectable
}

func (this *IGeopath) Vtbl() *IGeopathVtbl {
	return (*IGeopathVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeopath) Get_Positions() *IVectorView[BasicGeoposition] {
	var _result *IVectorView[BasicGeoposition]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Positions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 27BEA9C8-C7E7-4359-9B9B-FCA3E05EF593
var IID_IGeopathFactory = syscall.GUID{0x27BEA9C8, 0xC7E7, 0x4359,
	[8]byte{0x9B, 0x9B, 0xFC, 0xA3, 0xE0, 0x5E, 0xF5, 0x93}}

type IGeopathFactoryInterface interface {
	win32.IInspectableInterface
	Create(positions *IIterable[BasicGeoposition]) *IGeopath
	CreateWithAltitudeReference(positions *IIterable[BasicGeoposition], altitudeReferenceSystem AltitudeReferenceSystem) *IGeopath
	CreateWithAltitudeReferenceAndSpatialReference(positions *IIterable[BasicGeoposition], altitudeReferenceSystem AltitudeReferenceSystem, spatialReferenceId uint32) *IGeopath
}

type IGeopathFactoryVtbl struct {
	win32.IInspectableVtbl
	Create                                         uintptr
	CreateWithAltitudeReference                    uintptr
	CreateWithAltitudeReferenceAndSpatialReference uintptr
}

type IGeopathFactory struct {
	win32.IInspectable
}

func (this *IGeopathFactory) Vtbl() *IGeopathFactoryVtbl {
	return (*IGeopathFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeopathFactory) Create(positions *IIterable[BasicGeoposition]) *IGeopath {
	var _result *IGeopath
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(positions)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeopathFactory) CreateWithAltitudeReference(positions *IIterable[BasicGeoposition], altitudeReferenceSystem AltitudeReferenceSystem) *IGeopath {
	var _result *IGeopath
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithAltitudeReference, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(positions)), uintptr(altitudeReferenceSystem), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeopathFactory) CreateWithAltitudeReferenceAndSpatialReference(positions *IIterable[BasicGeoposition], altitudeReferenceSystem AltitudeReferenceSystem, spatialReferenceId uint32) *IGeopath {
	var _result *IGeopath
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithAltitudeReferenceAndSpatialReference, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(positions)), uintptr(altitudeReferenceSystem), uintptr(spatialReferenceId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6BFA00EB-E56E-49BB-9CAF-CBAA78A8BCEF
var IID_IGeopoint = syscall.GUID{0x6BFA00EB, 0xE56E, 0x49BB,
	[8]byte{0x9C, 0xAF, 0xCB, 0xAA, 0x78, 0xA8, 0xBC, 0xEF}}

type IGeopointInterface interface {
	win32.IInspectableInterface
	Get_Position() BasicGeoposition
}

type IGeopointVtbl struct {
	win32.IInspectableVtbl
	Get_Position uintptr
}

type IGeopoint struct {
	win32.IInspectable
}

func (this *IGeopoint) Vtbl() *IGeopointVtbl {
	return (*IGeopointVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeopoint) Get_Position() BasicGeoposition {
	var _result BasicGeoposition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// DB6B8D33-76BD-4E30-8AF7-A844DC37B7A0
var IID_IGeopointFactory = syscall.GUID{0xDB6B8D33, 0x76BD, 0x4E30,
	[8]byte{0x8A, 0xF7, 0xA8, 0x44, 0xDC, 0x37, 0xB7, 0xA0}}

type IGeopointFactoryInterface interface {
	win32.IInspectableInterface
	Create(position BasicGeoposition) *IGeopoint
	CreateWithAltitudeReferenceSystem(position BasicGeoposition, altitudeReferenceSystem AltitudeReferenceSystem) *IGeopoint
	CreateWithAltitudeReferenceSystemAndSpatialReferenceId(position BasicGeoposition, altitudeReferenceSystem AltitudeReferenceSystem, spatialReferenceId uint32) *IGeopoint
}

type IGeopointFactoryVtbl struct {
	win32.IInspectableVtbl
	Create                                                 uintptr
	CreateWithAltitudeReferenceSystem                      uintptr
	CreateWithAltitudeReferenceSystemAndSpatialReferenceId uintptr
}

type IGeopointFactory struct {
	win32.IInspectable
}

func (this *IGeopointFactory) Vtbl() *IGeopointFactoryVtbl {
	return (*IGeopointFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeopointFactory) Create(position BasicGeoposition) *IGeopoint {
	var _result *IGeopoint
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&position)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeopointFactory) CreateWithAltitudeReferenceSystem(position BasicGeoposition, altitudeReferenceSystem AltitudeReferenceSystem) *IGeopoint {
	var _result *IGeopoint
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithAltitudeReferenceSystem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&position)), uintptr(altitudeReferenceSystem), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeopointFactory) CreateWithAltitudeReferenceSystemAndSpatialReferenceId(position BasicGeoposition, altitudeReferenceSystem AltitudeReferenceSystem, spatialReferenceId uint32) *IGeopoint {
	var _result *IGeopoint
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithAltitudeReferenceSystemAndSpatialReferenceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&position)), uintptr(altitudeReferenceSystem), uintptr(spatialReferenceId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C18D0454-7D41-4FF7-A957-9DFFB4EF7F5B
var IID_IGeoposition = syscall.GUID{0xC18D0454, 0x7D41, 0x4FF7,
	[8]byte{0xA9, 0x57, 0x9D, 0xFF, 0xB4, 0xEF, 0x7F, 0x5B}}

type IGeopositionInterface interface {
	win32.IInspectableInterface
	Get_Coordinate() *IGeocoordinate
	Get_CivicAddress() *ICivicAddress
}

type IGeopositionVtbl struct {
	win32.IInspectableVtbl
	Get_Coordinate   uintptr
	Get_CivicAddress uintptr
}

type IGeoposition struct {
	win32.IInspectable
}

func (this *IGeoposition) Vtbl() *IGeopositionVtbl {
	return (*IGeopositionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeoposition) Get_Coordinate() *IGeocoordinate {
	var _result *IGeocoordinate
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Coordinate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeoposition) Get_CivicAddress() *ICivicAddress {
	var _result *ICivicAddress
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CivicAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7F62F697-8671-4B0D-86F8-474A8496187C
var IID_IGeoposition2 = syscall.GUID{0x7F62F697, 0x8671, 0x4B0D,
	[8]byte{0x86, 0xF8, 0x47, 0x4A, 0x84, 0x96, 0x18, 0x7C}}

type IGeoposition2Interface interface {
	win32.IInspectableInterface
	Get_VenueData() *IVenueData
}

type IGeoposition2Vtbl struct {
	win32.IInspectableVtbl
	Get_VenueData uintptr
}

type IGeoposition2 struct {
	win32.IInspectable
}

func (this *IGeoposition2) Vtbl() *IGeoposition2Vtbl {
	return (*IGeoposition2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeoposition2) Get_VenueData() *IVenueData {
	var _result *IVenueData
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VenueData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C99CA2AF-C729-43C1-8FAB-D6DEC914DF7E
var IID_IGeoshape = syscall.GUID{0xC99CA2AF, 0xC729, 0x43C1,
	[8]byte{0x8F, 0xAB, 0xD6, 0xDE, 0xC9, 0x14, 0xDF, 0x7E}}

type IGeoshapeInterface interface {
	win32.IInspectableInterface
	Get_GeoshapeType() GeoshapeType
	Get_SpatialReferenceId() uint32
	Get_AltitudeReferenceSystem() AltitudeReferenceSystem
}

type IGeoshapeVtbl struct {
	win32.IInspectableVtbl
	Get_GeoshapeType            uintptr
	Get_SpatialReferenceId      uintptr
	Get_AltitudeReferenceSystem uintptr
}

type IGeoshape struct {
	win32.IInspectable
}

func (this *IGeoshape) Vtbl() *IGeoshapeVtbl {
	return (*IGeoshapeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeoshape) Get_GeoshapeType() GeoshapeType {
	var _result GeoshapeType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GeoshapeType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGeoshape) Get_SpatialReferenceId() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SpatialReferenceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGeoshape) Get_AltitudeReferenceSystem() AltitudeReferenceSystem {
	var _result AltitudeReferenceSystem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AltitudeReferenceSystem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// B1877A76-9EF6-41AB-A0DD-793ECE76E2DE
var IID_IGeovisit = syscall.GUID{0xB1877A76, 0x9EF6, 0x41AB,
	[8]byte{0xA0, 0xDD, 0x79, 0x3E, 0xCE, 0x76, 0xE2, 0xDE}}

type IGeovisitInterface interface {
	win32.IInspectableInterface
	Get_Position() *IGeoposition
	Get_StateChange() VisitStateChange
	Get_Timestamp() DateTime
}

type IGeovisitVtbl struct {
	win32.IInspectableVtbl
	Get_Position    uintptr
	Get_StateChange uintptr
	Get_Timestamp   uintptr
}

type IGeovisit struct {
	win32.IInspectable
}

func (this *IGeovisit) Vtbl() *IGeovisitVtbl {
	return (*IGeovisitVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeovisit) Get_Position() *IGeoposition {
	var _result *IGeoposition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeovisit) Get_StateChange() VisitStateChange {
	var _result VisitStateChange
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StateChange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGeovisit) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 80118AAF-5944-4591-83C1-396647F54F2C
var IID_IGeovisitMonitor = syscall.GUID{0x80118AAF, 0x5944, 0x4591,
	[8]byte{0x83, 0xC1, 0x39, 0x66, 0x47, 0xF5, 0x4F, 0x2C}}

type IGeovisitMonitorInterface interface {
	win32.IInspectableInterface
	Get_MonitoringScope() VisitMonitoringScope
	Start(value VisitMonitoringScope)
	Stop()
	Add_VisitStateChanged(handler TypedEventHandler[*IGeovisitMonitor, *IGeovisitStateChangedEventArgs]) EventRegistrationToken
	Remove_VisitStateChanged(token EventRegistrationToken)
}

type IGeovisitMonitorVtbl struct {
	win32.IInspectableVtbl
	Get_MonitoringScope      uintptr
	Start                    uintptr
	Stop                     uintptr
	Add_VisitStateChanged    uintptr
	Remove_VisitStateChanged uintptr
}

type IGeovisitMonitor struct {
	win32.IInspectable
}

func (this *IGeovisitMonitor) Vtbl() *IGeovisitMonitorVtbl {
	return (*IGeovisitMonitorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeovisitMonitor) Get_MonitoringScope() VisitMonitoringScope {
	var _result VisitMonitoringScope
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MonitoringScope, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGeovisitMonitor) Start(value VisitMonitoringScope) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGeovisitMonitor) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IGeovisitMonitor) Add_VisitStateChanged(handler TypedEventHandler[*IGeovisitMonitor, *IGeovisitStateChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_VisitStateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGeovisitMonitor) Remove_VisitStateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_VisitStateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// BCF976A7-BBF2-4CDD-95CF-554C82EDFB87
var IID_IGeovisitMonitorStatics = syscall.GUID{0xBCF976A7, 0xBBF2, 0x4CDD,
	[8]byte{0x95, 0xCF, 0x55, 0x4C, 0x82, 0xED, 0xFB, 0x87}}

type IGeovisitMonitorStaticsInterface interface {
	win32.IInspectableInterface
	GetLastReportAsync() *IAsyncOperation[*IGeovisit]
}

type IGeovisitMonitorStaticsVtbl struct {
	win32.IInspectableVtbl
	GetLastReportAsync uintptr
}

type IGeovisitMonitorStatics struct {
	win32.IInspectable
}

func (this *IGeovisitMonitorStatics) Vtbl() *IGeovisitMonitorStaticsVtbl {
	return (*IGeovisitMonitorStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeovisitMonitorStatics) GetLastReportAsync() *IAsyncOperation[*IGeovisit] {
	var _result *IAsyncOperation[*IGeovisit]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetLastReportAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CEB4D1FF-8B53-4968-BEED-4CECD029CE15
var IID_IGeovisitStateChangedEventArgs = syscall.GUID{0xCEB4D1FF, 0x8B53, 0x4968,
	[8]byte{0xBE, 0xED, 0x4C, 0xEC, 0xD0, 0x29, 0xCE, 0x15}}

type IGeovisitStateChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Visit() *IGeovisit
}

type IGeovisitStateChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Visit uintptr
}

type IGeovisitStateChangedEventArgs struct {
	win32.IInspectable
}

func (this *IGeovisitStateChangedEventArgs) Vtbl() *IGeovisitStateChangedEventArgsVtbl {
	return (*IGeovisitStateChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeovisitStateChangedEventArgs) Get_Visit() *IGeovisit {
	var _result *IGeovisit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Visit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EA770D9E-D1C9-454B-99B7-B2F8CDD2482F
var IID_IGeovisitTriggerDetails = syscall.GUID{0xEA770D9E, 0xD1C9, 0x454B,
	[8]byte{0x99, 0xB7, 0xB2, 0xF8, 0xCD, 0xD2, 0x48, 0x2F}}

type IGeovisitTriggerDetailsInterface interface {
	win32.IInspectableInterface
	ReadReports() *IVectorView[*IGeovisit]
}

type IGeovisitTriggerDetailsVtbl struct {
	win32.IInspectableVtbl
	ReadReports uintptr
}

type IGeovisitTriggerDetails struct {
	win32.IInspectable
}

func (this *IGeovisitTriggerDetails) Vtbl() *IGeovisitTriggerDetailsVtbl {
	return (*IGeovisitTriggerDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeovisitTriggerDetails) ReadReports() *IVectorView[*IGeovisit] {
	var _result *IVectorView[*IGeovisit]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadReports, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 37859CE5-9D1E-46C5-BF3B-6AD8CAC1A093
var IID_IPositionChangedEventArgs = syscall.GUID{0x37859CE5, 0x9D1E, 0x46C5,
	[8]byte{0xBF, 0x3B, 0x6A, 0xD8, 0xCA, 0xC1, 0xA0, 0x93}}

type IPositionChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Position() *IGeoposition
}

type IPositionChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Position uintptr
}

type IPositionChangedEventArgs struct {
	win32.IInspectable
}

func (this *IPositionChangedEventArgs) Vtbl() *IPositionChangedEventArgsVtbl {
	return (*IPositionChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPositionChangedEventArgs) Get_Position() *IGeoposition {
	var _result *IGeoposition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3453D2DA-8C93-4111-A205-9AECFC9BE5C0
var IID_IStatusChangedEventArgs = syscall.GUID{0x3453D2DA, 0x8C93, 0x4111,
	[8]byte{0xA2, 0x05, 0x9A, 0xEC, 0xFC, 0x9B, 0xE5, 0xC0}}

type IStatusChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Status() PositionStatus
}

type IStatusChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
}

type IStatusChangedEventArgs struct {
	win32.IInspectable
}

func (this *IStatusChangedEventArgs) Vtbl() *IStatusChangedEventArgsVtbl {
	return (*IStatusChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStatusChangedEventArgs) Get_Status() PositionStatus {
	var _result PositionStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 66F39187-60E3-4B2F-B527-4F53F1C3C677
var IID_IVenueData = syscall.GUID{0x66F39187, 0x60E3, 0x4B2F,
	[8]byte{0xB5, 0x27, 0x4F, 0x53, 0xF1, 0xC3, 0xC6, 0x77}}

type IVenueDataInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Get_Level() string
}

type IVenueDataVtbl struct {
	win32.IInspectableVtbl
	Get_Id    uintptr
	Get_Level uintptr
}

type IVenueData struct {
	win32.IInspectable
}

func (this *IVenueData) Vtbl() *IVenueDataVtbl {
	return (*IVenueDataVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVenueData) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IVenueData) Get_Level() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Level, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// classes

type CivicAddress struct {
	RtClass
	*ICivicAddress
}

type GeoboundingBox struct {
	RtClass
	*IGeoboundingBox
}

func NewGeoboundingBox_Create(northwestCorner BasicGeoposition, southeastCorner BasicGeoposition) *GeoboundingBox {
	hs := NewHStr("Windows.Devices.Geolocation.GeoboundingBox")
	var pFac *IGeoboundingBoxFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGeoboundingBoxFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IGeoboundingBox
	p = pFac.Create(northwestCorner, southeastCorner)
	result := &GeoboundingBox{
		RtClass:         RtClass{PInspect: &p.IInspectable},
		IGeoboundingBox: p,
	}
	com.AddToScope(result)
	return result
}

func NewGeoboundingBox_CreateWithAltitudeReference(northwestCorner BasicGeoposition, southeastCorner BasicGeoposition, altitudeReferenceSystem AltitudeReferenceSystem) *GeoboundingBox {
	hs := NewHStr("Windows.Devices.Geolocation.GeoboundingBox")
	var pFac *IGeoboundingBoxFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGeoboundingBoxFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IGeoboundingBox
	p = pFac.CreateWithAltitudeReference(northwestCorner, southeastCorner, altitudeReferenceSystem)
	result := &GeoboundingBox{
		RtClass:         RtClass{PInspect: &p.IInspectable},
		IGeoboundingBox: p,
	}
	com.AddToScope(result)
	return result
}

func NewGeoboundingBox_CreateWithAltitudeReferenceAndSpatialReference(northwestCorner BasicGeoposition, southeastCorner BasicGeoposition, altitudeReferenceSystem AltitudeReferenceSystem, spatialReferenceId uint32) *GeoboundingBox {
	hs := NewHStr("Windows.Devices.Geolocation.GeoboundingBox")
	var pFac *IGeoboundingBoxFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGeoboundingBoxFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IGeoboundingBox
	p = pFac.CreateWithAltitudeReferenceAndSpatialReference(northwestCorner, southeastCorner, altitudeReferenceSystem, spatialReferenceId)
	result := &GeoboundingBox{
		RtClass:         RtClass{PInspect: &p.IInspectable},
		IGeoboundingBox: p,
	}
	com.AddToScope(result)
	return result
}

func NewIGeoboundingBoxStatics() *IGeoboundingBoxStatics {
	var p *IGeoboundingBoxStatics
	hs := NewHStr("Windows.Devices.Geolocation.GeoboundingBox")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGeoboundingBoxStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type Geocircle struct {
	RtClass
	*IGeocircle
}

func NewGeocircle_Create(position BasicGeoposition, radius float64) *Geocircle {
	hs := NewHStr("Windows.Devices.Geolocation.Geocircle")
	var pFac *IGeocircleFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGeocircleFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IGeocircle
	p = pFac.Create(position, radius)
	result := &Geocircle{
		RtClass:    RtClass{PInspect: &p.IInspectable},
		IGeocircle: p,
	}
	com.AddToScope(result)
	return result
}

func NewGeocircle_CreateWithAltitudeReferenceSystem(position BasicGeoposition, radius float64, altitudeReferenceSystem AltitudeReferenceSystem) *Geocircle {
	hs := NewHStr("Windows.Devices.Geolocation.Geocircle")
	var pFac *IGeocircleFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGeocircleFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IGeocircle
	p = pFac.CreateWithAltitudeReferenceSystem(position, radius, altitudeReferenceSystem)
	result := &Geocircle{
		RtClass:    RtClass{PInspect: &p.IInspectable},
		IGeocircle: p,
	}
	com.AddToScope(result)
	return result
}

func NewGeocircle_CreateWithAltitudeReferenceSystemAndSpatialReferenceId(position BasicGeoposition, radius float64, altitudeReferenceSystem AltitudeReferenceSystem, spatialReferenceId uint32) *Geocircle {
	hs := NewHStr("Windows.Devices.Geolocation.Geocircle")
	var pFac *IGeocircleFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGeocircleFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IGeocircle
	p = pFac.CreateWithAltitudeReferenceSystemAndSpatialReferenceId(position, radius, altitudeReferenceSystem, spatialReferenceId)
	result := &Geocircle{
		RtClass:    RtClass{PInspect: &p.IInspectable},
		IGeocircle: p,
	}
	com.AddToScope(result)
	return result
}

type Geocoordinate struct {
	RtClass
	*IGeocoordinate
}

type GeocoordinateSatelliteData struct {
	RtClass
	*IGeocoordinateSatelliteData
}

type Geolocator struct {
	RtClass
	*IGeolocator
}

func NewGeolocator() *Geolocator {
	hs := NewHStr("Windows.Devices.Geolocation.Geolocator")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &Geolocator{
		RtClass:     RtClass{PInspect: p},
		IGeolocator: (*IGeolocator)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewIGeolocatorStatics() *IGeolocatorStatics {
	var p *IGeolocatorStatics
	hs := NewHStr("Windows.Devices.Geolocation.Geolocator")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGeolocatorStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIGeolocatorStatics2() *IGeolocatorStatics2 {
	var p *IGeolocatorStatics2
	hs := NewHStr("Windows.Devices.Geolocation.Geolocator")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGeolocatorStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type Geopath struct {
	RtClass
	*IGeopath
}

func NewGeopath_Create(positions *IIterable[BasicGeoposition]) *Geopath {
	hs := NewHStr("Windows.Devices.Geolocation.Geopath")
	var pFac *IGeopathFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGeopathFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IGeopath
	p = pFac.Create(positions)
	result := &Geopath{
		RtClass:  RtClass{PInspect: &p.IInspectable},
		IGeopath: p,
	}
	com.AddToScope(result)
	return result
}

func NewGeopath_CreateWithAltitudeReference(positions *IIterable[BasicGeoposition], altitudeReferenceSystem AltitudeReferenceSystem) *Geopath {
	hs := NewHStr("Windows.Devices.Geolocation.Geopath")
	var pFac *IGeopathFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGeopathFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IGeopath
	p = pFac.CreateWithAltitudeReference(positions, altitudeReferenceSystem)
	result := &Geopath{
		RtClass:  RtClass{PInspect: &p.IInspectable},
		IGeopath: p,
	}
	com.AddToScope(result)
	return result
}

func NewGeopath_CreateWithAltitudeReferenceAndSpatialReference(positions *IIterable[BasicGeoposition], altitudeReferenceSystem AltitudeReferenceSystem, spatialReferenceId uint32) *Geopath {
	hs := NewHStr("Windows.Devices.Geolocation.Geopath")
	var pFac *IGeopathFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGeopathFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IGeopath
	p = pFac.CreateWithAltitudeReferenceAndSpatialReference(positions, altitudeReferenceSystem, spatialReferenceId)
	result := &Geopath{
		RtClass:  RtClass{PInspect: &p.IInspectable},
		IGeopath: p,
	}
	com.AddToScope(result)
	return result
}

type Geopoint struct {
	RtClass
	*IGeopoint
}

func NewGeopoint_Create(position BasicGeoposition) *Geopoint {
	hs := NewHStr("Windows.Devices.Geolocation.Geopoint")
	var pFac *IGeopointFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGeopointFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IGeopoint
	p = pFac.Create(position)
	result := &Geopoint{
		RtClass:   RtClass{PInspect: &p.IInspectable},
		IGeopoint: p,
	}
	com.AddToScope(result)
	return result
}

func NewGeopoint_CreateWithAltitudeReferenceSystem(position BasicGeoposition, altitudeReferenceSystem AltitudeReferenceSystem) *Geopoint {
	hs := NewHStr("Windows.Devices.Geolocation.Geopoint")
	var pFac *IGeopointFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGeopointFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IGeopoint
	p = pFac.CreateWithAltitudeReferenceSystem(position, altitudeReferenceSystem)
	result := &Geopoint{
		RtClass:   RtClass{PInspect: &p.IInspectable},
		IGeopoint: p,
	}
	com.AddToScope(result)
	return result
}

func NewGeopoint_CreateWithAltitudeReferenceSystemAndSpatialReferenceId(position BasicGeoposition, altitudeReferenceSystem AltitudeReferenceSystem, spatialReferenceId uint32) *Geopoint {
	hs := NewHStr("Windows.Devices.Geolocation.Geopoint")
	var pFac *IGeopointFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGeopointFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IGeopoint
	p = pFac.CreateWithAltitudeReferenceSystemAndSpatialReferenceId(position, altitudeReferenceSystem, spatialReferenceId)
	result := &Geopoint{
		RtClass:   RtClass{PInspect: &p.IInspectable},
		IGeopoint: p,
	}
	com.AddToScope(result)
	return result
}

type Geoposition struct {
	RtClass
	*IGeoposition
}

type Geovisit struct {
	RtClass
	*IGeovisit
}

type GeovisitMonitor struct {
	RtClass
	*IGeovisitMonitor
}

func NewGeovisitMonitor() *GeovisitMonitor {
	hs := NewHStr("Windows.Devices.Geolocation.GeovisitMonitor")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &GeovisitMonitor{
		RtClass:          RtClass{PInspect: p},
		IGeovisitMonitor: (*IGeovisitMonitor)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewIGeovisitMonitorStatics() *IGeovisitMonitorStatics {
	var p *IGeovisitMonitorStatics
	hs := NewHStr("Windows.Devices.Geolocation.GeovisitMonitor")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGeovisitMonitorStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type GeovisitStateChangedEventArgs struct {
	RtClass
	*IGeovisitStateChangedEventArgs
}

type GeovisitTriggerDetails struct {
	RtClass
	*IGeovisitTriggerDetails
}

type PositionChangedEventArgs struct {
	RtClass
	*IPositionChangedEventArgs
}

type StatusChangedEventArgs struct {
	RtClass
	*IStatusChangedEventArgs
}

type VenueData struct {
	RtClass
	*IVenueData
}
