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
type BarcodeScannerStatus int32

const (
	BarcodeScannerStatus_Online       BarcodeScannerStatus = 0
	BarcodeScannerStatus_Off          BarcodeScannerStatus = 1
	BarcodeScannerStatus_Offline      BarcodeScannerStatus = 2
	BarcodeScannerStatus_OffOrOffline BarcodeScannerStatus = 3
	BarcodeScannerStatus_Extended     BarcodeScannerStatus = 4
)

// enum
type BarcodeSymbologyDecodeLengthKind int32

const (
	BarcodeSymbologyDecodeLengthKind_AnyLength BarcodeSymbologyDecodeLengthKind = 0
	BarcodeSymbologyDecodeLengthKind_Discrete  BarcodeSymbologyDecodeLengthKind = 1
	BarcodeSymbologyDecodeLengthKind_Range     BarcodeSymbologyDecodeLengthKind = 2
)

// enum
type CashDrawerStatusKind int32

const (
	CashDrawerStatusKind_Online       CashDrawerStatusKind = 0
	CashDrawerStatusKind_Off          CashDrawerStatusKind = 1
	CashDrawerStatusKind_Offline      CashDrawerStatusKind = 2
	CashDrawerStatusKind_OffOrOffline CashDrawerStatusKind = 3
	CashDrawerStatusKind_Extended     CashDrawerStatusKind = 4
)

// enum
type LineDisplayCursorType int32

const (
	LineDisplayCursorType_None      LineDisplayCursorType = 0
	LineDisplayCursorType_Block     LineDisplayCursorType = 1
	LineDisplayCursorType_HalfBlock LineDisplayCursorType = 2
	LineDisplayCursorType_Underline LineDisplayCursorType = 3
	LineDisplayCursorType_Reverse   LineDisplayCursorType = 4
	LineDisplayCursorType_Other     LineDisplayCursorType = 5
)

// enum
type LineDisplayDescriptorState int32

const (
	LineDisplayDescriptorState_Off   LineDisplayDescriptorState = 0
	LineDisplayDescriptorState_On    LineDisplayDescriptorState = 1
	LineDisplayDescriptorState_Blink LineDisplayDescriptorState = 2
)

// enum
type LineDisplayHorizontalAlignment int32

const (
	LineDisplayHorizontalAlignment_Left   LineDisplayHorizontalAlignment = 0
	LineDisplayHorizontalAlignment_Center LineDisplayHorizontalAlignment = 1
	LineDisplayHorizontalAlignment_Right  LineDisplayHorizontalAlignment = 2
)

// enum
type LineDisplayMarqueeFormat int32

const (
	LineDisplayMarqueeFormat_None  LineDisplayMarqueeFormat = 0
	LineDisplayMarqueeFormat_Walk  LineDisplayMarqueeFormat = 1
	LineDisplayMarqueeFormat_Place LineDisplayMarqueeFormat = 2
)

// enum
type LineDisplayPowerStatus int32

const (
	LineDisplayPowerStatus_Unknown      LineDisplayPowerStatus = 0
	LineDisplayPowerStatus_Online       LineDisplayPowerStatus = 1
	LineDisplayPowerStatus_Off          LineDisplayPowerStatus = 2
	LineDisplayPowerStatus_Offline      LineDisplayPowerStatus = 3
	LineDisplayPowerStatus_OffOrOffline LineDisplayPowerStatus = 4
)

// enum
type LineDisplayScrollDirection int32

const (
	LineDisplayScrollDirection_Up    LineDisplayScrollDirection = 0
	LineDisplayScrollDirection_Down  LineDisplayScrollDirection = 1
	LineDisplayScrollDirection_Left  LineDisplayScrollDirection = 2
	LineDisplayScrollDirection_Right LineDisplayScrollDirection = 3
)

// enum
type LineDisplayTextAttribute int32

const (
	LineDisplayTextAttribute_Normal       LineDisplayTextAttribute = 0
	LineDisplayTextAttribute_Blink        LineDisplayTextAttribute = 1
	LineDisplayTextAttribute_Reverse      LineDisplayTextAttribute = 2
	LineDisplayTextAttribute_ReverseBlink LineDisplayTextAttribute = 3
)

// enum
type LineDisplayTextAttributeGranularity int32

const (
	LineDisplayTextAttributeGranularity_NotSupported  LineDisplayTextAttributeGranularity = 0
	LineDisplayTextAttributeGranularity_EntireDisplay LineDisplayTextAttributeGranularity = 1
	LineDisplayTextAttributeGranularity_PerCharacter  LineDisplayTextAttributeGranularity = 2
)

// enum
type LineDisplayVerticalAlignment int32

const (
	LineDisplayVerticalAlignment_Top    LineDisplayVerticalAlignment = 0
	LineDisplayVerticalAlignment_Center LineDisplayVerticalAlignment = 1
	LineDisplayVerticalAlignment_Bottom LineDisplayVerticalAlignment = 2
)

// enum
type MagneticStripeReaderAuthenticationLevel int32

const (
	MagneticStripeReaderAuthenticationLevel_NotSupported MagneticStripeReaderAuthenticationLevel = 0
	MagneticStripeReaderAuthenticationLevel_Optional     MagneticStripeReaderAuthenticationLevel = 1
	MagneticStripeReaderAuthenticationLevel_Required     MagneticStripeReaderAuthenticationLevel = 2
)

// enum
type MagneticStripeReaderAuthenticationProtocol int32

const (
	MagneticStripeReaderAuthenticationProtocol_None              MagneticStripeReaderAuthenticationProtocol = 0
	MagneticStripeReaderAuthenticationProtocol_ChallengeResponse MagneticStripeReaderAuthenticationProtocol = 1
)

// enum
type MagneticStripeReaderErrorReportingType int32

const (
	MagneticStripeReaderErrorReportingType_CardLevel  MagneticStripeReaderErrorReportingType = 0
	MagneticStripeReaderErrorReportingType_TrackLevel MagneticStripeReaderErrorReportingType = 1
)

// enum
type MagneticStripeReaderStatus int32

const (
	MagneticStripeReaderStatus_Unauthenticated MagneticStripeReaderStatus = 0
	MagneticStripeReaderStatus_Authenticated   MagneticStripeReaderStatus = 1
	MagneticStripeReaderStatus_Extended        MagneticStripeReaderStatus = 2
)

// enum
type MagneticStripeReaderTrackErrorType int32

const (
	MagneticStripeReaderTrackErrorType_None               MagneticStripeReaderTrackErrorType = 0
	MagneticStripeReaderTrackErrorType_StartSentinelError MagneticStripeReaderTrackErrorType = 1
	MagneticStripeReaderTrackErrorType_EndSentinelError   MagneticStripeReaderTrackErrorType = 2
	MagneticStripeReaderTrackErrorType_ParityError        MagneticStripeReaderTrackErrorType = 3
	MagneticStripeReaderTrackErrorType_LrcError           MagneticStripeReaderTrackErrorType = 4
	MagneticStripeReaderTrackErrorType_Unknown            MagneticStripeReaderTrackErrorType = -1
)

// enum
type MagneticStripeReaderTrackIds int32

const (
	MagneticStripeReaderTrackIds_None   MagneticStripeReaderTrackIds = 0
	MagneticStripeReaderTrackIds_Track1 MagneticStripeReaderTrackIds = 1
	MagneticStripeReaderTrackIds_Track2 MagneticStripeReaderTrackIds = 2
	MagneticStripeReaderTrackIds_Track3 MagneticStripeReaderTrackIds = 4
	MagneticStripeReaderTrackIds_Track4 MagneticStripeReaderTrackIds = 8
)

// enum
// flags
type PosConnectionTypes uint32

const (
	PosConnectionTypes_Local     PosConnectionTypes = 1
	PosConnectionTypes_IP        PosConnectionTypes = 2
	PosConnectionTypes_Bluetooth PosConnectionTypes = 4
	PosConnectionTypes_All       PosConnectionTypes = 4294967295
)

// enum
type PosPrinterAlignment int32

const (
	PosPrinterAlignment_Left   PosPrinterAlignment = 0
	PosPrinterAlignment_Center PosPrinterAlignment = 1
	PosPrinterAlignment_Right  PosPrinterAlignment = 2
)

// enum
type PosPrinterBarcodeTextPosition int32

const (
	PosPrinterBarcodeTextPosition_None  PosPrinterBarcodeTextPosition = 0
	PosPrinterBarcodeTextPosition_Above PosPrinterBarcodeTextPosition = 1
	PosPrinterBarcodeTextPosition_Below PosPrinterBarcodeTextPosition = 2
)

// enum
// flags
type PosPrinterCartridgeSensors uint32

const (
	PosPrinterCartridgeSensors_None         PosPrinterCartridgeSensors = 0
	PosPrinterCartridgeSensors_Removed      PosPrinterCartridgeSensors = 1
	PosPrinterCartridgeSensors_Empty        PosPrinterCartridgeSensors = 2
	PosPrinterCartridgeSensors_HeadCleaning PosPrinterCartridgeSensors = 4
	PosPrinterCartridgeSensors_NearEnd      PosPrinterCartridgeSensors = 8
)

// enum
// flags
type PosPrinterColorCapabilities uint32

const (
	PosPrinterColorCapabilities_None    PosPrinterColorCapabilities = 0
	PosPrinterColorCapabilities_Primary PosPrinterColorCapabilities = 1
	PosPrinterColorCapabilities_Custom1 PosPrinterColorCapabilities = 2
	PosPrinterColorCapabilities_Custom2 PosPrinterColorCapabilities = 4
	PosPrinterColorCapabilities_Custom3 PosPrinterColorCapabilities = 8
	PosPrinterColorCapabilities_Custom4 PosPrinterColorCapabilities = 16
	PosPrinterColorCapabilities_Custom5 PosPrinterColorCapabilities = 32
	PosPrinterColorCapabilities_Custom6 PosPrinterColorCapabilities = 64
	PosPrinterColorCapabilities_Cyan    PosPrinterColorCapabilities = 128
	PosPrinterColorCapabilities_Magenta PosPrinterColorCapabilities = 256
	PosPrinterColorCapabilities_Yellow  PosPrinterColorCapabilities = 512
	PosPrinterColorCapabilities_Full    PosPrinterColorCapabilities = 1024
)

// enum
type PosPrinterColorCartridge int32

const (
	PosPrinterColorCartridge_Unknown PosPrinterColorCartridge = 0
	PosPrinterColorCartridge_Primary PosPrinterColorCartridge = 1
	PosPrinterColorCartridge_Custom1 PosPrinterColorCartridge = 2
	PosPrinterColorCartridge_Custom2 PosPrinterColorCartridge = 3
	PosPrinterColorCartridge_Custom3 PosPrinterColorCartridge = 4
	PosPrinterColorCartridge_Custom4 PosPrinterColorCartridge = 5
	PosPrinterColorCartridge_Custom5 PosPrinterColorCartridge = 6
	PosPrinterColorCartridge_Custom6 PosPrinterColorCartridge = 7
	PosPrinterColorCartridge_Cyan    PosPrinterColorCartridge = 8
	PosPrinterColorCartridge_Magenta PosPrinterColorCartridge = 9
	PosPrinterColorCartridge_Yellow  PosPrinterColorCartridge = 10
)

// enum
type PosPrinterLineDirection int32

const (
	PosPrinterLineDirection_Horizontal PosPrinterLineDirection = 0
	PosPrinterLineDirection_Vertical   PosPrinterLineDirection = 1
)

// enum
type PosPrinterLineStyle int32

const (
	PosPrinterLineStyle_SingleSolid PosPrinterLineStyle = 0
	PosPrinterLineStyle_DoubleSolid PosPrinterLineStyle = 1
	PosPrinterLineStyle_Broken      PosPrinterLineStyle = 2
	PosPrinterLineStyle_Chain       PosPrinterLineStyle = 3
)

// enum
type PosPrinterMapMode int32

const (
	PosPrinterMapMode_Dots    PosPrinterMapMode = 0
	PosPrinterMapMode_Twips   PosPrinterMapMode = 1
	PosPrinterMapMode_English PosPrinterMapMode = 2
	PosPrinterMapMode_Metric  PosPrinterMapMode = 3
)

// enum
// flags
type PosPrinterMarkFeedCapabilities uint32

const (
	PosPrinterMarkFeedCapabilities_None               PosPrinterMarkFeedCapabilities = 0
	PosPrinterMarkFeedCapabilities_ToTakeUp           PosPrinterMarkFeedCapabilities = 1
	PosPrinterMarkFeedCapabilities_ToCutter           PosPrinterMarkFeedCapabilities = 2
	PosPrinterMarkFeedCapabilities_ToCurrentTopOfForm PosPrinterMarkFeedCapabilities = 4
	PosPrinterMarkFeedCapabilities_ToNextTopOfForm    PosPrinterMarkFeedCapabilities = 8
)

// enum
type PosPrinterMarkFeedKind int32

const (
	PosPrinterMarkFeedKind_ToTakeUp           PosPrinterMarkFeedKind = 0
	PosPrinterMarkFeedKind_ToCutter           PosPrinterMarkFeedKind = 1
	PosPrinterMarkFeedKind_ToCurrentTopOfForm PosPrinterMarkFeedKind = 2
	PosPrinterMarkFeedKind_ToNextTopOfForm    PosPrinterMarkFeedKind = 3
)

// enum
type PosPrinterPrintSide int32

const (
	PosPrinterPrintSide_Unknown PosPrinterPrintSide = 0
	PosPrinterPrintSide_Side1   PosPrinterPrintSide = 1
	PosPrinterPrintSide_Side2   PosPrinterPrintSide = 2
)

// enum
type PosPrinterRotation int32

const (
	PosPrinterRotation_Normal    PosPrinterRotation = 0
	PosPrinterRotation_Right90   PosPrinterRotation = 1
	PosPrinterRotation_Left90    PosPrinterRotation = 2
	PosPrinterRotation_Rotate180 PosPrinterRotation = 3
)

// enum
// flags
type PosPrinterRuledLineCapabilities uint32

const (
	PosPrinterRuledLineCapabilities_None       PosPrinterRuledLineCapabilities = 0
	PosPrinterRuledLineCapabilities_Horizontal PosPrinterRuledLineCapabilities = 1
	PosPrinterRuledLineCapabilities_Vertical   PosPrinterRuledLineCapabilities = 2
)

// enum
type PosPrinterStatusKind int32

const (
	PosPrinterStatusKind_Online       PosPrinterStatusKind = 0
	PosPrinterStatusKind_Off          PosPrinterStatusKind = 1
	PosPrinterStatusKind_Offline      PosPrinterStatusKind = 2
	PosPrinterStatusKind_OffOrOffline PosPrinterStatusKind = 3
	PosPrinterStatusKind_Extended     PosPrinterStatusKind = 4
)

// enum
type UnifiedPosErrorReason int32

const (
	UnifiedPosErrorReason_UnknownErrorReason UnifiedPosErrorReason = 0
	UnifiedPosErrorReason_NoService          UnifiedPosErrorReason = 1
	UnifiedPosErrorReason_Disabled           UnifiedPosErrorReason = 2
	UnifiedPosErrorReason_Illegal            UnifiedPosErrorReason = 3
	UnifiedPosErrorReason_NoHardware         UnifiedPosErrorReason = 4
	UnifiedPosErrorReason_Closed             UnifiedPosErrorReason = 5
	UnifiedPosErrorReason_Offline            UnifiedPosErrorReason = 6
	UnifiedPosErrorReason_Failure            UnifiedPosErrorReason = 7
	UnifiedPosErrorReason_Timeout            UnifiedPosErrorReason = 8
	UnifiedPosErrorReason_Busy               UnifiedPosErrorReason = 9
	UnifiedPosErrorReason_Extended           UnifiedPosErrorReason = 10
)

// enum
type UnifiedPosErrorSeverity int32

const (
	UnifiedPosErrorSeverity_UnknownErrorSeverity UnifiedPosErrorSeverity = 0
	UnifiedPosErrorSeverity_Warning              UnifiedPosErrorSeverity = 1
	UnifiedPosErrorSeverity_Recoverable          UnifiedPosErrorSeverity = 2
	UnifiedPosErrorSeverity_Unrecoverable        UnifiedPosErrorSeverity = 3
	UnifiedPosErrorSeverity_AssistanceRequired   UnifiedPosErrorSeverity = 4
	UnifiedPosErrorSeverity_Fatal                UnifiedPosErrorSeverity = 5
)

// enum
type UnifiedPosHealthCheckLevel int32

const (
	UnifiedPosHealthCheckLevel_UnknownHealthCheckLevel UnifiedPosHealthCheckLevel = 0
	UnifiedPosHealthCheckLevel_POSInternal             UnifiedPosHealthCheckLevel = 1
	UnifiedPosHealthCheckLevel_External                UnifiedPosHealthCheckLevel = 2
	UnifiedPosHealthCheckLevel_Interactive             UnifiedPosHealthCheckLevel = 3
)

// enum
type UnifiedPosPowerReportingType int32

const (
	UnifiedPosPowerReportingType_UnknownPowerReportingType UnifiedPosPowerReportingType = 0
	UnifiedPosPowerReportingType_Standard                  UnifiedPosPowerReportingType = 1
	UnifiedPosPowerReportingType_Advanced                  UnifiedPosPowerReportingType = 2
)

// structs

type SizeUInt32 struct {
	Width  uint32
	Height uint32
}

// interfaces

// BEA33E06-B264-4F03-A9C1-45B20F01134F
var IID_IBarcodeScanner = syscall.GUID{0xBEA33E06, 0xB264, 0x4F03,
	[8]byte{0xA9, 0xC1, 0x45, 0xB2, 0x0F, 0x01, 0x13, 0x4F}}

type IBarcodeScannerInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_Capabilities() *IBarcodeScannerCapabilities
	ClaimScannerAsync() *IAsyncOperation[*IClaimedBarcodeScanner]
	CheckHealthAsync(level UnifiedPosHealthCheckLevel) *IAsyncOperation[string]
	GetSupportedSymbologiesAsync() *IAsyncOperation[*IVectorView[uint32]]
	IsSymbologySupportedAsync(barcodeSymbology uint32) *IAsyncOperation[bool]
	RetrieveStatisticsAsync(statisticsCategories *IIterable[string]) *IAsyncOperation[*IBuffer]
	GetSupportedProfiles() *IVectorView[string]
	IsProfileSupported(profile string) bool
	Add_StatusUpdated(handler TypedEventHandler[*IBarcodeScanner, *IBarcodeScannerStatusUpdatedEventArgs]) EventRegistrationToken
	Remove_StatusUpdated(token EventRegistrationToken)
}

type IBarcodeScannerVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId                 uintptr
	Get_Capabilities             uintptr
	ClaimScannerAsync            uintptr
	CheckHealthAsync             uintptr
	GetSupportedSymbologiesAsync uintptr
	IsSymbologySupportedAsync    uintptr
	RetrieveStatisticsAsync      uintptr
	GetSupportedProfiles         uintptr
	IsProfileSupported           uintptr
	Add_StatusUpdated            uintptr
	Remove_StatusUpdated         uintptr
}

type IBarcodeScanner struct {
	win32.IInspectable
}

func (this *IBarcodeScanner) Vtbl() *IBarcodeScannerVtbl {
	return (*IBarcodeScannerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarcodeScanner) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBarcodeScanner) Get_Capabilities() *IBarcodeScannerCapabilities {
	var _result *IBarcodeScannerCapabilities
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Capabilities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBarcodeScanner) ClaimScannerAsync() *IAsyncOperation[*IClaimedBarcodeScanner] {
	var _result *IAsyncOperation[*IClaimedBarcodeScanner]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClaimScannerAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBarcodeScanner) CheckHealthAsync(level UnifiedPosHealthCheckLevel) *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CheckHealthAsync, uintptr(unsafe.Pointer(this)), uintptr(level), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBarcodeScanner) GetSupportedSymbologiesAsync() *IAsyncOperation[*IVectorView[uint32]] {
	var _result *IAsyncOperation[*IVectorView[uint32]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSupportedSymbologiesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBarcodeScanner) IsSymbologySupportedAsync(barcodeSymbology uint32) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsSymbologySupportedAsync, uintptr(unsafe.Pointer(this)), uintptr(barcodeSymbology), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBarcodeScanner) RetrieveStatisticsAsync(statisticsCategories *IIterable[string]) *IAsyncOperation[*IBuffer] {
	var _result *IAsyncOperation[*IBuffer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RetrieveStatisticsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(statisticsCategories)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBarcodeScanner) GetSupportedProfiles() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSupportedProfiles, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBarcodeScanner) IsProfileSupported(profile string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsProfileSupported, uintptr(unsafe.Pointer(this)), NewHStr(profile).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeScanner) Add_StatusUpdated(handler TypedEventHandler[*IBarcodeScanner, *IBarcodeScannerStatusUpdatedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StatusUpdated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeScanner) Remove_StatusUpdated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StatusUpdated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 89215167-8CEE-436D-89AB-8DFB43BB4286
var IID_IBarcodeScanner2 = syscall.GUID{0x89215167, 0x8CEE, 0x436D,
	[8]byte{0x89, 0xAB, 0x8D, 0xFB, 0x43, 0xBB, 0x42, 0x86}}

type IBarcodeScanner2Interface interface {
	win32.IInspectableInterface
	Get_VideoDeviceId() string
}

type IBarcodeScanner2Vtbl struct {
	win32.IInspectableVtbl
	Get_VideoDeviceId uintptr
}

type IBarcodeScanner2 struct {
	win32.IInspectable
}

func (this *IBarcodeScanner2) Vtbl() *IBarcodeScanner2Vtbl {
	return (*IBarcodeScanner2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarcodeScanner2) Get_VideoDeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoDeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// C60691E4-F2C8-4420-A307-B12EF6622857
var IID_IBarcodeScannerCapabilities = syscall.GUID{0xC60691E4, 0xF2C8, 0x4420,
	[8]byte{0xA3, 0x07, 0xB1, 0x2E, 0xF6, 0x62, 0x28, 0x57}}

type IBarcodeScannerCapabilitiesInterface interface {
	win32.IInspectableInterface
	Get_PowerReportingType() UnifiedPosPowerReportingType
	Get_IsStatisticsReportingSupported() bool
	Get_IsStatisticsUpdatingSupported() bool
	Get_IsImagePreviewSupported() bool
}

type IBarcodeScannerCapabilitiesVtbl struct {
	win32.IInspectableVtbl
	Get_PowerReportingType             uintptr
	Get_IsStatisticsReportingSupported uintptr
	Get_IsStatisticsUpdatingSupported  uintptr
	Get_IsImagePreviewSupported        uintptr
}

type IBarcodeScannerCapabilities struct {
	win32.IInspectable
}

func (this *IBarcodeScannerCapabilities) Vtbl() *IBarcodeScannerCapabilitiesVtbl {
	return (*IBarcodeScannerCapabilitiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarcodeScannerCapabilities) Get_PowerReportingType() UnifiedPosPowerReportingType {
	var _result UnifiedPosPowerReportingType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PowerReportingType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeScannerCapabilities) Get_IsStatisticsReportingSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsStatisticsReportingSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeScannerCapabilities) Get_IsStatisticsUpdatingSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsStatisticsUpdatingSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeScannerCapabilities) Get_IsImagePreviewSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsImagePreviewSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 8E5AB3E9-0E2C-472F-A1CC-EE8054B6A684
var IID_IBarcodeScannerCapabilities1 = syscall.GUID{0x8E5AB3E9, 0x0E2C, 0x472F,
	[8]byte{0xA1, 0xCC, 0xEE, 0x80, 0x54, 0xB6, 0xA6, 0x84}}

type IBarcodeScannerCapabilities1Interface interface {
	win32.IInspectableInterface
	Get_IsSoftwareTriggerSupported() bool
}

type IBarcodeScannerCapabilities1Vtbl struct {
	win32.IInspectableVtbl
	Get_IsSoftwareTriggerSupported uintptr
}

type IBarcodeScannerCapabilities1 struct {
	win32.IInspectable
}

func (this *IBarcodeScannerCapabilities1) Vtbl() *IBarcodeScannerCapabilities1Vtbl {
	return (*IBarcodeScannerCapabilities1Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarcodeScannerCapabilities1) Get_IsSoftwareTriggerSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSoftwareTriggerSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F211CFEC-E1A1-4EA8-9ABC-92B1596270AB
var IID_IBarcodeScannerCapabilities2 = syscall.GUID{0xF211CFEC, 0xE1A1, 0x4EA8,
	[8]byte{0x9A, 0xBC, 0x92, 0xB1, 0x59, 0x62, 0x70, 0xAB}}

type IBarcodeScannerCapabilities2Interface interface {
	win32.IInspectableInterface
	Get_IsVideoPreviewSupported() bool
}

type IBarcodeScannerCapabilities2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsVideoPreviewSupported uintptr
}

type IBarcodeScannerCapabilities2 struct {
	win32.IInspectable
}

func (this *IBarcodeScannerCapabilities2) Vtbl() *IBarcodeScannerCapabilities2Vtbl {
	return (*IBarcodeScannerCapabilities2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarcodeScannerCapabilities2) Get_IsVideoPreviewSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsVideoPreviewSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4234A7E2-ED97-467D-AD2B-01E44313A929
var IID_IBarcodeScannerDataReceivedEventArgs = syscall.GUID{0x4234A7E2, 0xED97, 0x467D,
	[8]byte{0xAD, 0x2B, 0x01, 0xE4, 0x43, 0x13, 0xA9, 0x29}}

type IBarcodeScannerDataReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Report() *IBarcodeScannerReport
}

type IBarcodeScannerDataReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Report uintptr
}

type IBarcodeScannerDataReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IBarcodeScannerDataReceivedEventArgs) Vtbl() *IBarcodeScannerDataReceivedEventArgsVtbl {
	return (*IBarcodeScannerDataReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarcodeScannerDataReceivedEventArgs) Get_Report() *IBarcodeScannerReport {
	var _result *IBarcodeScannerReport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Report, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2CD2602F-CF3A-4002-A75A-C5EC468F0A20
var IID_IBarcodeScannerErrorOccurredEventArgs = syscall.GUID{0x2CD2602F, 0xCF3A, 0x4002,
	[8]byte{0xA7, 0x5A, 0xC5, 0xEC, 0x46, 0x8F, 0x0A, 0x20}}

type IBarcodeScannerErrorOccurredEventArgsInterface interface {
	win32.IInspectableInterface
	Get_PartialInputData() *IBarcodeScannerReport
	Get_IsRetriable() bool
	Get_ErrorData() *IUnifiedPosErrorData
}

type IBarcodeScannerErrorOccurredEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_PartialInputData uintptr
	Get_IsRetriable      uintptr
	Get_ErrorData        uintptr
}

type IBarcodeScannerErrorOccurredEventArgs struct {
	win32.IInspectable
}

func (this *IBarcodeScannerErrorOccurredEventArgs) Vtbl() *IBarcodeScannerErrorOccurredEventArgsVtbl {
	return (*IBarcodeScannerErrorOccurredEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarcodeScannerErrorOccurredEventArgs) Get_PartialInputData() *IBarcodeScannerReport {
	var _result *IBarcodeScannerReport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PartialInputData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBarcodeScannerErrorOccurredEventArgs) Get_IsRetriable() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsRetriable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeScannerErrorOccurredEventArgs) Get_ErrorData() *IUnifiedPosErrorData {
	var _result *IUnifiedPosErrorData
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F3B7DE85-6E8B-434E-9F58-06EF26BC4BAF
var IID_IBarcodeScannerImagePreviewReceivedEventArgs = syscall.GUID{0xF3B7DE85, 0x6E8B, 0x434E,
	[8]byte{0x9F, 0x58, 0x06, 0xEF, 0x26, 0xBC, 0x4B, 0xAF}}

type IBarcodeScannerImagePreviewReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Preview() *IRandomAccessStreamWithContentType
}

type IBarcodeScannerImagePreviewReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Preview uintptr
}

type IBarcodeScannerImagePreviewReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IBarcodeScannerImagePreviewReceivedEventArgs) Vtbl() *IBarcodeScannerImagePreviewReceivedEventArgsVtbl {
	return (*IBarcodeScannerImagePreviewReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarcodeScannerImagePreviewReceivedEventArgs) Get_Preview() *IRandomAccessStreamWithContentType {
	var _result *IRandomAccessStreamWithContentType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Preview, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5CE4D8B0-A489-4B96-86C4-F0BF8A37753D
var IID_IBarcodeScannerReport = syscall.GUID{0x5CE4D8B0, 0xA489, 0x4B96,
	[8]byte{0x86, 0xC4, 0xF0, 0xBF, 0x8A, 0x37, 0x75, 0x3D}}

type IBarcodeScannerReportInterface interface {
	win32.IInspectableInterface
	Get_ScanDataType() uint32
	Get_ScanData() *IBuffer
	Get_ScanDataLabel() *IBuffer
}

type IBarcodeScannerReportVtbl struct {
	win32.IInspectableVtbl
	Get_ScanDataType  uintptr
	Get_ScanData      uintptr
	Get_ScanDataLabel uintptr
}

type IBarcodeScannerReport struct {
	win32.IInspectable
}

func (this *IBarcodeScannerReport) Vtbl() *IBarcodeScannerReportVtbl {
	return (*IBarcodeScannerReportVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarcodeScannerReport) Get_ScanDataType() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScanDataType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeScannerReport) Get_ScanData() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScanData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBarcodeScannerReport) Get_ScanDataLabel() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScanDataLabel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A2547326-2013-457C-8963-49C15DCA78CE
var IID_IBarcodeScannerReportFactory = syscall.GUID{0xA2547326, 0x2013, 0x457C,
	[8]byte{0x89, 0x63, 0x49, 0xC1, 0x5D, 0xCA, 0x78, 0xCE}}

type IBarcodeScannerReportFactoryInterface interface {
	win32.IInspectableInterface
	CreateInstance(scanDataType uint32, scanData *IBuffer, scanDataLabel *IBuffer) *IBarcodeScannerReport
}

type IBarcodeScannerReportFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateInstance uintptr
}

type IBarcodeScannerReportFactory struct {
	win32.IInspectable
}

func (this *IBarcodeScannerReportFactory) Vtbl() *IBarcodeScannerReportFactoryVtbl {
	return (*IBarcodeScannerReportFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarcodeScannerReportFactory) CreateInstance(scanDataType uint32, scanData *IBuffer, scanDataLabel *IBuffer) *IBarcodeScannerReport {
	var _result *IBarcodeScannerReport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateInstance, uintptr(unsafe.Pointer(this)), uintptr(scanDataType), uintptr(unsafe.Pointer(scanData)), uintptr(unsafe.Pointer(scanDataLabel)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5D115F6F-DA49-41E8-8C8C-F0CB62A9C4FC
var IID_IBarcodeScannerStatics = syscall.GUID{0x5D115F6F, 0xDA49, 0x41E8,
	[8]byte{0x8C, 0x8C, 0xF0, 0xCB, 0x62, 0xA9, 0xC4, 0xFC}}

type IBarcodeScannerStaticsInterface interface {
	win32.IInspectableInterface
	GetDefaultAsync() *IAsyncOperation[*IBarcodeScanner]
	FromIdAsync(deviceId string) *IAsyncOperation[*IBarcodeScanner]
	GetDeviceSelector() string
}

type IBarcodeScannerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefaultAsync   uintptr
	FromIdAsync       uintptr
	GetDeviceSelector uintptr
}

type IBarcodeScannerStatics struct {
	win32.IInspectable
}

func (this *IBarcodeScannerStatics) Vtbl() *IBarcodeScannerStaticsVtbl {
	return (*IBarcodeScannerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarcodeScannerStatics) GetDefaultAsync() *IAsyncOperation[*IBarcodeScanner] {
	var _result *IAsyncOperation[*IBarcodeScanner]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBarcodeScannerStatics) FromIdAsync(deviceId string) *IAsyncOperation[*IBarcodeScanner] {
	var _result *IAsyncOperation[*IBarcodeScanner]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBarcodeScannerStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// B8652473-A36F-4007-B1D0-279EBE92A656
var IID_IBarcodeScannerStatics2 = syscall.GUID{0xB8652473, 0xA36F, 0x4007,
	[8]byte{0xB1, 0xD0, 0x27, 0x9E, 0xBE, 0x92, 0xA6, 0x56}}

type IBarcodeScannerStatics2Interface interface {
	win32.IInspectableInterface
	GetDeviceSelectorWithConnectionTypes(connectionTypes PosConnectionTypes) string
}

type IBarcodeScannerStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelectorWithConnectionTypes uintptr
}

type IBarcodeScannerStatics2 struct {
	win32.IInspectable
}

func (this *IBarcodeScannerStatics2) Vtbl() *IBarcodeScannerStatics2Vtbl {
	return (*IBarcodeScannerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarcodeScannerStatics2) GetDeviceSelectorWithConnectionTypes(connectionTypes PosConnectionTypes) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorWithConnectionTypes, uintptr(unsafe.Pointer(this)), uintptr(connectionTypes), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 355D8586-9C43-462B-A91A-816DC97F452C
var IID_IBarcodeScannerStatusUpdatedEventArgs = syscall.GUID{0x355D8586, 0x9C43, 0x462B,
	[8]byte{0xA9, 0x1A, 0x81, 0x6D, 0xC9, 0x7F, 0x45, 0x2C}}

type IBarcodeScannerStatusUpdatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Status() BarcodeScannerStatus
	Get_ExtendedStatus() uint32
}

type IBarcodeScannerStatusUpdatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Status         uintptr
	Get_ExtendedStatus uintptr
}

type IBarcodeScannerStatusUpdatedEventArgs struct {
	win32.IInspectable
}

func (this *IBarcodeScannerStatusUpdatedEventArgs) Vtbl() *IBarcodeScannerStatusUpdatedEventArgsVtbl {
	return (*IBarcodeScannerStatusUpdatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarcodeScannerStatusUpdatedEventArgs) Get_Status() BarcodeScannerStatus {
	var _result BarcodeScannerStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeScannerStatusUpdatedEventArgs) Get_ExtendedStatus() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// CA8549BB-06D2-43F4-A44B-C620679FD8D0
var IID_IBarcodeSymbologiesStatics = syscall.GUID{0xCA8549BB, 0x06D2, 0x43F4,
	[8]byte{0xA4, 0x4B, 0xC6, 0x20, 0x67, 0x9F, 0xD8, 0xD0}}

type IBarcodeSymbologiesStaticsInterface interface {
	win32.IInspectableInterface
	Get_Unknown() uint32
	Get_Ean8() uint32
	Get_Ean8Add2() uint32
	Get_Ean8Add5() uint32
	Get_Eanv() uint32
	Get_EanvAdd2() uint32
	Get_EanvAdd5() uint32
	Get_Ean13() uint32
	Get_Ean13Add2() uint32
	Get_Ean13Add5() uint32
	Get_Isbn() uint32
	Get_IsbnAdd5() uint32
	Get_Ismn() uint32
	Get_IsmnAdd2() uint32
	Get_IsmnAdd5() uint32
	Get_Issn() uint32
	Get_IssnAdd2() uint32
	Get_IssnAdd5() uint32
	Get_Ean99() uint32
	Get_Ean99Add2() uint32
	Get_Ean99Add5() uint32
	Get_Upca() uint32
	Get_UpcaAdd2() uint32
	Get_UpcaAdd5() uint32
	Get_Upce() uint32
	Get_UpceAdd2() uint32
	Get_UpceAdd5() uint32
	Get_UpcCoupon() uint32
	Get_TfStd() uint32
	Get_TfDis() uint32
	Get_TfInt() uint32
	Get_TfInd() uint32
	Get_TfMat() uint32
	Get_TfIata() uint32
	Get_Gs1DatabarType1() uint32
	Get_Gs1DatabarType2() uint32
	Get_Gs1DatabarType3() uint32
	Get_Code39() uint32
	Get_Code39Ex() uint32
	Get_Trioptic39() uint32
	Get_Code32() uint32
	Get_Pzn() uint32
	Get_Code93() uint32
	Get_Code93Ex() uint32
	Get_Code128() uint32
	Get_Gs1128() uint32
	Get_Gs1128Coupon() uint32
	Get_UccEan128() uint32
	Get_Sisac() uint32
	Get_Isbt() uint32
	Get_Codabar() uint32
	Get_Code11() uint32
	Get_Msi() uint32
	Get_Plessey() uint32
	Get_Telepen() uint32
	Get_Code16k() uint32
	Get_CodablockA() uint32
	Get_CodablockF() uint32
	Get_Codablock128() uint32
	Get_Code49() uint32
	Get_Aztec() uint32
	Get_DataCode() uint32
	Get_DataMatrix() uint32
	Get_HanXin() uint32
	Get_Maxicode() uint32
	Get_MicroPdf417() uint32
	Get_MicroQr() uint32
	Get_Pdf417() uint32
	Get_Qr() uint32
	Get_MsTag() uint32
	Get_Ccab() uint32
	Get_Ccc() uint32
	Get_Tlc39() uint32
	Get_AusPost() uint32
	Get_CanPost() uint32
	Get_ChinaPost() uint32
	Get_DutchKix() uint32
	Get_InfoMail() uint32
	Get_ItalianPost25() uint32
	Get_ItalianPost39() uint32
	Get_JapanPost() uint32
	Get_KoreanPost() uint32
	Get_SwedenPost() uint32
	Get_UkPost() uint32
	Get_UsIntelligent() uint32
	Get_UsIntelligentPkg() uint32
	Get_UsPlanet() uint32
	Get_UsPostNet() uint32
	Get_Us4StateFics() uint32
	Get_OcrA() uint32
	Get_OcrB() uint32
	Get_Micr() uint32
	Get_ExtendedBase() uint32
	GetName(scanDataType uint32) string
}

type IBarcodeSymbologiesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Unknown          uintptr
	Get_Ean8             uintptr
	Get_Ean8Add2         uintptr
	Get_Ean8Add5         uintptr
	Get_Eanv             uintptr
	Get_EanvAdd2         uintptr
	Get_EanvAdd5         uintptr
	Get_Ean13            uintptr
	Get_Ean13Add2        uintptr
	Get_Ean13Add5        uintptr
	Get_Isbn             uintptr
	Get_IsbnAdd5         uintptr
	Get_Ismn             uintptr
	Get_IsmnAdd2         uintptr
	Get_IsmnAdd5         uintptr
	Get_Issn             uintptr
	Get_IssnAdd2         uintptr
	Get_IssnAdd5         uintptr
	Get_Ean99            uintptr
	Get_Ean99Add2        uintptr
	Get_Ean99Add5        uintptr
	Get_Upca             uintptr
	Get_UpcaAdd2         uintptr
	Get_UpcaAdd5         uintptr
	Get_Upce             uintptr
	Get_UpceAdd2         uintptr
	Get_UpceAdd5         uintptr
	Get_UpcCoupon        uintptr
	Get_TfStd            uintptr
	Get_TfDis            uintptr
	Get_TfInt            uintptr
	Get_TfInd            uintptr
	Get_TfMat            uintptr
	Get_TfIata           uintptr
	Get_Gs1DatabarType1  uintptr
	Get_Gs1DatabarType2  uintptr
	Get_Gs1DatabarType3  uintptr
	Get_Code39           uintptr
	Get_Code39Ex         uintptr
	Get_Trioptic39       uintptr
	Get_Code32           uintptr
	Get_Pzn              uintptr
	Get_Code93           uintptr
	Get_Code93Ex         uintptr
	Get_Code128          uintptr
	Get_Gs1128           uintptr
	Get_Gs1128Coupon     uintptr
	Get_UccEan128        uintptr
	Get_Sisac            uintptr
	Get_Isbt             uintptr
	Get_Codabar          uintptr
	Get_Code11           uintptr
	Get_Msi              uintptr
	Get_Plessey          uintptr
	Get_Telepen          uintptr
	Get_Code16k          uintptr
	Get_CodablockA       uintptr
	Get_CodablockF       uintptr
	Get_Codablock128     uintptr
	Get_Code49           uintptr
	Get_Aztec            uintptr
	Get_DataCode         uintptr
	Get_DataMatrix       uintptr
	Get_HanXin           uintptr
	Get_Maxicode         uintptr
	Get_MicroPdf417      uintptr
	Get_MicroQr          uintptr
	Get_Pdf417           uintptr
	Get_Qr               uintptr
	Get_MsTag            uintptr
	Get_Ccab             uintptr
	Get_Ccc              uintptr
	Get_Tlc39            uintptr
	Get_AusPost          uintptr
	Get_CanPost          uintptr
	Get_ChinaPost        uintptr
	Get_DutchKix         uintptr
	Get_InfoMail         uintptr
	Get_ItalianPost25    uintptr
	Get_ItalianPost39    uintptr
	Get_JapanPost        uintptr
	Get_KoreanPost       uintptr
	Get_SwedenPost       uintptr
	Get_UkPost           uintptr
	Get_UsIntelligent    uintptr
	Get_UsIntelligentPkg uintptr
	Get_UsPlanet         uintptr
	Get_UsPostNet        uintptr
	Get_Us4StateFics     uintptr
	Get_OcrA             uintptr
	Get_OcrB             uintptr
	Get_Micr             uintptr
	Get_ExtendedBase     uintptr
	GetName              uintptr
}

type IBarcodeSymbologiesStatics struct {
	win32.IInspectable
}

func (this *IBarcodeSymbologiesStatics) Vtbl() *IBarcodeSymbologiesStaticsVtbl {
	return (*IBarcodeSymbologiesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarcodeSymbologiesStatics) Get_Unknown() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Unknown, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Ean8() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ean8, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Ean8Add2() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ean8Add2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Ean8Add5() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ean8Add5, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Eanv() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Eanv, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_EanvAdd2() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EanvAdd2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_EanvAdd5() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EanvAdd5, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Ean13() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ean13, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Ean13Add2() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ean13Add2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Ean13Add5() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ean13Add5, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Isbn() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Isbn, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_IsbnAdd5() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsbnAdd5, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Ismn() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ismn, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_IsmnAdd2() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsmnAdd2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_IsmnAdd5() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsmnAdd5, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Issn() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Issn, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_IssnAdd2() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IssnAdd2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_IssnAdd5() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IssnAdd5, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Ean99() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ean99, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Ean99Add2() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ean99Add2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Ean99Add5() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ean99Add5, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Upca() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Upca, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_UpcaAdd2() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UpcaAdd2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_UpcaAdd5() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UpcaAdd5, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Upce() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Upce, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_UpceAdd2() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UpceAdd2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_UpceAdd5() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UpceAdd5, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_UpcCoupon() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UpcCoupon, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_TfStd() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TfStd, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_TfDis() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TfDis, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_TfInt() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TfInt, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_TfInd() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TfInd, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_TfMat() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TfMat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_TfIata() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TfIata, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Gs1DatabarType1() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Gs1DatabarType1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Gs1DatabarType2() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Gs1DatabarType2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Gs1DatabarType3() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Gs1DatabarType3, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Code39() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Code39, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Code39Ex() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Code39Ex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Trioptic39() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Trioptic39, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Code32() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Code32, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Pzn() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Pzn, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Code93() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Code93, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Code93Ex() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Code93Ex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Code128() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Code128, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Gs1128() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Gs1128, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Gs1128Coupon() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Gs1128Coupon, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_UccEan128() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UccEan128, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Sisac() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Sisac, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Isbt() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Isbt, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Codabar() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Codabar, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Code11() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Code11, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Msi() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Msi, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Plessey() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Plessey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Telepen() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Telepen, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Code16k() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Code16k, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_CodablockA() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CodablockA, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_CodablockF() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CodablockF, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Codablock128() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Codablock128, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Code49() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Code49, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Aztec() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Aztec, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_DataCode() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DataCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_DataMatrix() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DataMatrix, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_HanXin() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HanXin, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Maxicode() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Maxicode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_MicroPdf417() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MicroPdf417, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_MicroQr() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MicroQr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Pdf417() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Pdf417, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Qr() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Qr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_MsTag() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MsTag, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Ccab() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ccab, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Ccc() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ccc, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Tlc39() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Tlc39, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_AusPost() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AusPost, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_CanPost() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanPost, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_ChinaPost() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ChinaPost, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_DutchKix() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DutchKix, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_InfoMail() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InfoMail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_ItalianPost25() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ItalianPost25, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_ItalianPost39() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ItalianPost39, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_JapanPost() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_JapanPost, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_KoreanPost() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KoreanPost, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_SwedenPost() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SwedenPost, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_UkPost() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UkPost, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_UsIntelligent() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UsIntelligent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_UsIntelligentPkg() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UsIntelligentPkg, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_UsPlanet() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UsPlanet, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_UsPostNet() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UsPostNet, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Us4StateFics() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Us4StateFics, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_OcrA() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OcrA, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_OcrB() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OcrB, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_Micr() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Micr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) Get_ExtendedBase() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedBase, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologiesStatics) GetName(scanDataType uint32) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetName, uintptr(unsafe.Pointer(this)), uintptr(scanDataType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 8B7518F4-99D0-40BF-9424-B91D6DD4C6E0
var IID_IBarcodeSymbologiesStatics2 = syscall.GUID{0x8B7518F4, 0x99D0, 0x40BF,
	[8]byte{0x94, 0x24, 0xB9, 0x1D, 0x6D, 0xD4, 0xC6, 0xE0}}

type IBarcodeSymbologiesStatics2Interface interface {
	win32.IInspectableInterface
	Get_Gs1DWCode() uint32
}

type IBarcodeSymbologiesStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_Gs1DWCode uintptr
}

type IBarcodeSymbologiesStatics2 struct {
	win32.IInspectable
}

func (this *IBarcodeSymbologiesStatics2) Vtbl() *IBarcodeSymbologiesStatics2Vtbl {
	return (*IBarcodeSymbologiesStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarcodeSymbologiesStatics2) Get_Gs1DWCode() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Gs1DWCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 66413A78-AB7A-4ADA-8ECE-936014B2EAD7
var IID_IBarcodeSymbologyAttributes = syscall.GUID{0x66413A78, 0xAB7A, 0x4ADA,
	[8]byte{0x8E, 0xCE, 0x93, 0x60, 0x14, 0xB2, 0xEA, 0xD7}}

type IBarcodeSymbologyAttributesInterface interface {
	win32.IInspectableInterface
	Get_IsCheckDigitValidationEnabled() bool
	Put_IsCheckDigitValidationEnabled(value bool)
	Get_IsCheckDigitValidationSupported() bool
	Get_IsCheckDigitTransmissionEnabled() bool
	Put_IsCheckDigitTransmissionEnabled(value bool)
	Get_IsCheckDigitTransmissionSupported() bool
	Get_DecodeLength1() uint32
	Put_DecodeLength1(value uint32)
	Get_DecodeLength2() uint32
	Put_DecodeLength2(value uint32)
	Get_DecodeLengthKind() BarcodeSymbologyDecodeLengthKind
	Put_DecodeLengthKind(value BarcodeSymbologyDecodeLengthKind)
	Get_IsDecodeLengthSupported() bool
}

type IBarcodeSymbologyAttributesVtbl struct {
	win32.IInspectableVtbl
	Get_IsCheckDigitValidationEnabled     uintptr
	Put_IsCheckDigitValidationEnabled     uintptr
	Get_IsCheckDigitValidationSupported   uintptr
	Get_IsCheckDigitTransmissionEnabled   uintptr
	Put_IsCheckDigitTransmissionEnabled   uintptr
	Get_IsCheckDigitTransmissionSupported uintptr
	Get_DecodeLength1                     uintptr
	Put_DecodeLength1                     uintptr
	Get_DecodeLength2                     uintptr
	Put_DecodeLength2                     uintptr
	Get_DecodeLengthKind                  uintptr
	Put_DecodeLengthKind                  uintptr
	Get_IsDecodeLengthSupported           uintptr
}

type IBarcodeSymbologyAttributes struct {
	win32.IInspectable
}

func (this *IBarcodeSymbologyAttributes) Vtbl() *IBarcodeSymbologyAttributesVtbl {
	return (*IBarcodeSymbologyAttributesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarcodeSymbologyAttributes) Get_IsCheckDigitValidationEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCheckDigitValidationEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologyAttributes) Put_IsCheckDigitValidationEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsCheckDigitValidationEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IBarcodeSymbologyAttributes) Get_IsCheckDigitValidationSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCheckDigitValidationSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologyAttributes) Get_IsCheckDigitTransmissionEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCheckDigitTransmissionEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologyAttributes) Put_IsCheckDigitTransmissionEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsCheckDigitTransmissionEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IBarcodeSymbologyAttributes) Get_IsCheckDigitTransmissionSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCheckDigitTransmissionSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologyAttributes) Get_DecodeLength1() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DecodeLength1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologyAttributes) Put_DecodeLength1(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DecodeLength1, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IBarcodeSymbologyAttributes) Get_DecodeLength2() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DecodeLength2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologyAttributes) Put_DecodeLength2(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DecodeLength2, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IBarcodeSymbologyAttributes) Get_DecodeLengthKind() BarcodeSymbologyDecodeLengthKind {
	var _result BarcodeSymbologyDecodeLengthKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DecodeLengthKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarcodeSymbologyAttributes) Put_DecodeLengthKind(value BarcodeSymbologyDecodeLengthKind) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DecodeLengthKind, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IBarcodeSymbologyAttributes) Get_IsDecodeLengthSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDecodeLengthSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 9F88F5C8-DE54-4AEE-A890-920BCBFE30FC
var IID_ICashDrawer = syscall.GUID{0x9F88F5C8, 0xDE54, 0x4AEE,
	[8]byte{0xA8, 0x90, 0x92, 0x0B, 0xCB, 0xFE, 0x30, 0xFC}}

type ICashDrawerInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_Capabilities() *ICashDrawerCapabilities
	Get_Status() *ICashDrawerStatus
	Get_IsDrawerOpen() bool
	Get_DrawerEventSource() *ICashDrawerEventSource
	ClaimDrawerAsync() *IAsyncOperation[*IClaimedCashDrawer]
	CheckHealthAsync(level UnifiedPosHealthCheckLevel) *IAsyncOperation[string]
	GetStatisticsAsync(statisticsCategories *IIterable[string]) *IAsyncOperation[string]
	Add_StatusUpdated(handler TypedEventHandler[*ICashDrawer, *ICashDrawerStatusUpdatedEventArgs]) EventRegistrationToken
	Remove_StatusUpdated(token EventRegistrationToken)
}

type ICashDrawerVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId          uintptr
	Get_Capabilities      uintptr
	Get_Status            uintptr
	Get_IsDrawerOpen      uintptr
	Get_DrawerEventSource uintptr
	ClaimDrawerAsync      uintptr
	CheckHealthAsync      uintptr
	GetStatisticsAsync    uintptr
	Add_StatusUpdated     uintptr
	Remove_StatusUpdated  uintptr
}

type ICashDrawer struct {
	win32.IInspectable
}

func (this *ICashDrawer) Vtbl() *ICashDrawerVtbl {
	return (*ICashDrawerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICashDrawer) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICashDrawer) Get_Capabilities() *ICashDrawerCapabilities {
	var _result *ICashDrawerCapabilities
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Capabilities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICashDrawer) Get_Status() *ICashDrawerStatus {
	var _result *ICashDrawerStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICashDrawer) Get_IsDrawerOpen() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDrawerOpen, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICashDrawer) Get_DrawerEventSource() *ICashDrawerEventSource {
	var _result *ICashDrawerEventSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DrawerEventSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICashDrawer) ClaimDrawerAsync() *IAsyncOperation[*IClaimedCashDrawer] {
	var _result *IAsyncOperation[*IClaimedCashDrawer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClaimDrawerAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICashDrawer) CheckHealthAsync(level UnifiedPosHealthCheckLevel) *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CheckHealthAsync, uintptr(unsafe.Pointer(this)), uintptr(level), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICashDrawer) GetStatisticsAsync(statisticsCategories *IIterable[string]) *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStatisticsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(statisticsCategories)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICashDrawer) Add_StatusUpdated(handler TypedEventHandler[*ICashDrawer, *ICashDrawerStatusUpdatedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StatusUpdated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICashDrawer) Remove_StatusUpdated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StatusUpdated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 0BC6DE0B-E8E7-4B1F-B1D1-3E501AD08247
var IID_ICashDrawerCapabilities = syscall.GUID{0x0BC6DE0B, 0xE8E7, 0x4B1F,
	[8]byte{0xB1, 0xD1, 0x3E, 0x50, 0x1A, 0xD0, 0x82, 0x47}}

type ICashDrawerCapabilitiesInterface interface {
	win32.IInspectableInterface
	Get_PowerReportingType() UnifiedPosPowerReportingType
	Get_IsStatisticsReportingSupported() bool
	Get_IsStatisticsUpdatingSupported() bool
	Get_IsStatusReportingSupported() bool
	Get_IsStatusMultiDrawerDetectSupported() bool
	Get_IsDrawerOpenSensorAvailable() bool
}

type ICashDrawerCapabilitiesVtbl struct {
	win32.IInspectableVtbl
	Get_PowerReportingType                 uintptr
	Get_IsStatisticsReportingSupported     uintptr
	Get_IsStatisticsUpdatingSupported      uintptr
	Get_IsStatusReportingSupported         uintptr
	Get_IsStatusMultiDrawerDetectSupported uintptr
	Get_IsDrawerOpenSensorAvailable        uintptr
}

type ICashDrawerCapabilities struct {
	win32.IInspectable
}

func (this *ICashDrawerCapabilities) Vtbl() *ICashDrawerCapabilitiesVtbl {
	return (*ICashDrawerCapabilitiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICashDrawerCapabilities) Get_PowerReportingType() UnifiedPosPowerReportingType {
	var _result UnifiedPosPowerReportingType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PowerReportingType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICashDrawerCapabilities) Get_IsStatisticsReportingSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsStatisticsReportingSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICashDrawerCapabilities) Get_IsStatisticsUpdatingSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsStatisticsUpdatingSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICashDrawerCapabilities) Get_IsStatusReportingSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsStatusReportingSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICashDrawerCapabilities) Get_IsStatusMultiDrawerDetectSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsStatusMultiDrawerDetectSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICashDrawerCapabilities) Get_IsDrawerOpenSensorAvailable() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDrawerOpenSensorAvailable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6BF88CC7-6F63-430E-AB3B-95D75FFBE87F
var IID_ICashDrawerCloseAlarm = syscall.GUID{0x6BF88CC7, 0x6F63, 0x430E,
	[8]byte{0xAB, 0x3B, 0x95, 0xD7, 0x5F, 0xFB, 0xE8, 0x7F}}

type ICashDrawerCloseAlarmInterface interface {
	win32.IInspectableInterface
	Put_AlarmTimeout(value TimeSpan)
	Get_AlarmTimeout() TimeSpan
	Put_BeepFrequency(value uint32)
	Get_BeepFrequency() uint32
	Put_BeepDuration(value TimeSpan)
	Get_BeepDuration() TimeSpan
	Put_BeepDelay(value TimeSpan)
	Get_BeepDelay() TimeSpan
	Add_AlarmTimeoutExpired(handler TypedEventHandler[*ICashDrawerCloseAlarm, interface{}]) EventRegistrationToken
	Remove_AlarmTimeoutExpired(token EventRegistrationToken)
	StartAsync() *IAsyncOperation[bool]
}

type ICashDrawerCloseAlarmVtbl struct {
	win32.IInspectableVtbl
	Put_AlarmTimeout           uintptr
	Get_AlarmTimeout           uintptr
	Put_BeepFrequency          uintptr
	Get_BeepFrequency          uintptr
	Put_BeepDuration           uintptr
	Get_BeepDuration           uintptr
	Put_BeepDelay              uintptr
	Get_BeepDelay              uintptr
	Add_AlarmTimeoutExpired    uintptr
	Remove_AlarmTimeoutExpired uintptr
	StartAsync                 uintptr
}

type ICashDrawerCloseAlarm struct {
	win32.IInspectable
}

func (this *ICashDrawerCloseAlarm) Vtbl() *ICashDrawerCloseAlarmVtbl {
	return (*ICashDrawerCloseAlarmVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICashDrawerCloseAlarm) Put_AlarmTimeout(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AlarmTimeout, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ICashDrawerCloseAlarm) Get_AlarmTimeout() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlarmTimeout, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICashDrawerCloseAlarm) Put_BeepFrequency(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BeepFrequency, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICashDrawerCloseAlarm) Get_BeepFrequency() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BeepFrequency, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICashDrawerCloseAlarm) Put_BeepDuration(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BeepDuration, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ICashDrawerCloseAlarm) Get_BeepDuration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BeepDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICashDrawerCloseAlarm) Put_BeepDelay(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BeepDelay, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ICashDrawerCloseAlarm) Get_BeepDelay() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BeepDelay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICashDrawerCloseAlarm) Add_AlarmTimeoutExpired(handler TypedEventHandler[*ICashDrawerCloseAlarm, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AlarmTimeoutExpired, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICashDrawerCloseAlarm) Remove_AlarmTimeoutExpired(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AlarmTimeoutExpired, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICashDrawerCloseAlarm) StartAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E006E46C-F2F9-442F-8DD6-06C10A4227BA
var IID_ICashDrawerEventSource = syscall.GUID{0xE006E46C, 0xF2F9, 0x442F,
	[8]byte{0x8D, 0xD6, 0x06, 0xC1, 0x0A, 0x42, 0x27, 0xBA}}

type ICashDrawerEventSourceInterface interface {
	win32.IInspectableInterface
	Add_DrawerClosed(handler TypedEventHandler[*ICashDrawerEventSource, *ICashDrawerEventSourceEventArgs]) EventRegistrationToken
	Remove_DrawerClosed(token EventRegistrationToken)
	Add_DrawerOpened(handler TypedEventHandler[*ICashDrawerEventSource, *ICashDrawerEventSourceEventArgs]) EventRegistrationToken
	Remove_DrawerOpened(token EventRegistrationToken)
}

type ICashDrawerEventSourceVtbl struct {
	win32.IInspectableVtbl
	Add_DrawerClosed    uintptr
	Remove_DrawerClosed uintptr
	Add_DrawerOpened    uintptr
	Remove_DrawerOpened uintptr
}

type ICashDrawerEventSource struct {
	win32.IInspectable
}

func (this *ICashDrawerEventSource) Vtbl() *ICashDrawerEventSourceVtbl {
	return (*ICashDrawerEventSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICashDrawerEventSource) Add_DrawerClosed(handler TypedEventHandler[*ICashDrawerEventSource, *ICashDrawerEventSourceEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DrawerClosed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICashDrawerEventSource) Remove_DrawerClosed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DrawerClosed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICashDrawerEventSource) Add_DrawerOpened(handler TypedEventHandler[*ICashDrawerEventSource, *ICashDrawerEventSourceEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DrawerOpened, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICashDrawerEventSource) Remove_DrawerOpened(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DrawerOpened, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 69CB3BC1-147F-421C-9C23-090123BB786C
var IID_ICashDrawerEventSourceEventArgs = syscall.GUID{0x69CB3BC1, 0x147F, 0x421C,
	[8]byte{0x9C, 0x23, 0x09, 0x01, 0x23, 0xBB, 0x78, 0x6C}}

type ICashDrawerEventSourceEventArgsInterface interface {
	win32.IInspectableInterface
	Get_CashDrawer() *ICashDrawer
}

type ICashDrawerEventSourceEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_CashDrawer uintptr
}

type ICashDrawerEventSourceEventArgs struct {
	win32.IInspectable
}

func (this *ICashDrawerEventSourceEventArgs) Vtbl() *ICashDrawerEventSourceEventArgsVtbl {
	return (*ICashDrawerEventSourceEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICashDrawerEventSourceEventArgs) Get_CashDrawer() *ICashDrawer {
	var _result *ICashDrawer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CashDrawer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DFA0955A-D437-4FFF-B547-DDA969A4F883
var IID_ICashDrawerStatics = syscall.GUID{0xDFA0955A, 0xD437, 0x4FFF,
	[8]byte{0xB5, 0x47, 0xDD, 0xA9, 0x69, 0xA4, 0xF8, 0x83}}

type ICashDrawerStaticsInterface interface {
	win32.IInspectableInterface
	GetDefaultAsync() *IAsyncOperation[*ICashDrawer]
	FromIdAsync(deviceId string) *IAsyncOperation[*ICashDrawer]
	GetDeviceSelector() string
}

type ICashDrawerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefaultAsync   uintptr
	FromIdAsync       uintptr
	GetDeviceSelector uintptr
}

type ICashDrawerStatics struct {
	win32.IInspectable
}

func (this *ICashDrawerStatics) Vtbl() *ICashDrawerStaticsVtbl {
	return (*ICashDrawerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICashDrawerStatics) GetDefaultAsync() *IAsyncOperation[*ICashDrawer] {
	var _result *IAsyncOperation[*ICashDrawer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICashDrawerStatics) FromIdAsync(deviceId string) *IAsyncOperation[*ICashDrawer] {
	var _result *IAsyncOperation[*ICashDrawer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICashDrawerStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 3E818121-8C42-40E8-9C0E-40297048104C
var IID_ICashDrawerStatics2 = syscall.GUID{0x3E818121, 0x8C42, 0x40E8,
	[8]byte{0x9C, 0x0E, 0x40, 0x29, 0x70, 0x48, 0x10, 0x4C}}

type ICashDrawerStatics2Interface interface {
	win32.IInspectableInterface
	GetDeviceSelectorWithConnectionTypes(connectionTypes PosConnectionTypes) string
}

type ICashDrawerStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelectorWithConnectionTypes uintptr
}

type ICashDrawerStatics2 struct {
	win32.IInspectable
}

func (this *ICashDrawerStatics2) Vtbl() *ICashDrawerStatics2Vtbl {
	return (*ICashDrawerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICashDrawerStatics2) GetDeviceSelectorWithConnectionTypes(connectionTypes PosConnectionTypes) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorWithConnectionTypes, uintptr(unsafe.Pointer(this)), uintptr(connectionTypes), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 6BBD78BF-DCA1-4E06-99EB-5AF6A5AEC108
var IID_ICashDrawerStatus = syscall.GUID{0x6BBD78BF, 0xDCA1, 0x4E06,
	[8]byte{0x99, 0xEB, 0x5A, 0xF6, 0xA5, 0xAE, 0xC1, 0x08}}

type ICashDrawerStatusInterface interface {
	win32.IInspectableInterface
	Get_StatusKind() CashDrawerStatusKind
	Get_ExtendedStatus() uint32
}

type ICashDrawerStatusVtbl struct {
	win32.IInspectableVtbl
	Get_StatusKind     uintptr
	Get_ExtendedStatus uintptr
}

type ICashDrawerStatus struct {
	win32.IInspectable
}

func (this *ICashDrawerStatus) Vtbl() *ICashDrawerStatusVtbl {
	return (*ICashDrawerStatusVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICashDrawerStatus) Get_StatusKind() CashDrawerStatusKind {
	var _result CashDrawerStatusKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StatusKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICashDrawerStatus) Get_ExtendedStatus() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 30AAE98A-0D70-459C-9553-87E124C52488
var IID_ICashDrawerStatusUpdatedEventArgs = syscall.GUID{0x30AAE98A, 0x0D70, 0x459C,
	[8]byte{0x95, 0x53, 0x87, 0xE1, 0x24, 0xC5, 0x24, 0x88}}

type ICashDrawerStatusUpdatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Status() *ICashDrawerStatus
}

type ICashDrawerStatusUpdatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
}

type ICashDrawerStatusUpdatedEventArgs struct {
	win32.IInspectable
}

func (this *ICashDrawerStatusUpdatedEventArgs) Vtbl() *ICashDrawerStatusUpdatedEventArgsVtbl {
	return (*ICashDrawerStatusUpdatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICashDrawerStatusUpdatedEventArgs) Get_Status() *ICashDrawerStatus {
	var _result *ICashDrawerStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4A63B49C-8FA4-4332-BB26-945D11D81E0F
var IID_IClaimedBarcodeScanner = syscall.GUID{0x4A63B49C, 0x8FA4, 0x4332,
	[8]byte{0xBB, 0x26, 0x94, 0x5D, 0x11, 0xD8, 0x1E, 0x0F}}

type IClaimedBarcodeScannerInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_IsEnabled() bool
	Put_IsDisabledOnDataReceived(value bool)
	Get_IsDisabledOnDataReceived() bool
	Put_IsDecodeDataEnabled(value bool)
	Get_IsDecodeDataEnabled() bool
	EnableAsync() *IAsyncAction
	DisableAsync() *IAsyncAction
	RetainDevice()
	SetActiveSymbologiesAsync(symbologies *IIterable[uint32]) *IAsyncAction
	ResetStatisticsAsync(statisticsCategories *IIterable[string]) *IAsyncAction
	UpdateStatisticsAsync(statistics *IIterable[*IKeyValuePair[string, string]]) *IAsyncAction
	SetActiveProfileAsync(profile string) *IAsyncAction
	Add_DataReceived(handler TypedEventHandler[*IClaimedBarcodeScanner, *IBarcodeScannerDataReceivedEventArgs]) EventRegistrationToken
	Remove_DataReceived(token EventRegistrationToken)
	Add_TriggerPressed(handler EventHandler[*IClaimedBarcodeScanner]) EventRegistrationToken
	Remove_TriggerPressed(token EventRegistrationToken)
	Add_TriggerReleased(handler EventHandler[*IClaimedBarcodeScanner]) EventRegistrationToken
	Remove_TriggerReleased(token EventRegistrationToken)
	Add_ReleaseDeviceRequested(handler EventHandler[*IClaimedBarcodeScanner]) EventRegistrationToken
	Remove_ReleaseDeviceRequested(token EventRegistrationToken)
	Add_ImagePreviewReceived(handler TypedEventHandler[*IClaimedBarcodeScanner, *IBarcodeScannerImagePreviewReceivedEventArgs]) EventRegistrationToken
	Remove_ImagePreviewReceived(token EventRegistrationToken)
	Add_ErrorOccurred(handler TypedEventHandler[*IClaimedBarcodeScanner, *IBarcodeScannerErrorOccurredEventArgs]) EventRegistrationToken
	Remove_ErrorOccurred(token EventRegistrationToken)
}

type IClaimedBarcodeScannerVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId                  uintptr
	Get_IsEnabled                 uintptr
	Put_IsDisabledOnDataReceived  uintptr
	Get_IsDisabledOnDataReceived  uintptr
	Put_IsDecodeDataEnabled       uintptr
	Get_IsDecodeDataEnabled       uintptr
	EnableAsync                   uintptr
	DisableAsync                  uintptr
	RetainDevice                  uintptr
	SetActiveSymbologiesAsync     uintptr
	ResetStatisticsAsync          uintptr
	UpdateStatisticsAsync         uintptr
	SetActiveProfileAsync         uintptr
	Add_DataReceived              uintptr
	Remove_DataReceived           uintptr
	Add_TriggerPressed            uintptr
	Remove_TriggerPressed         uintptr
	Add_TriggerReleased           uintptr
	Remove_TriggerReleased        uintptr
	Add_ReleaseDeviceRequested    uintptr
	Remove_ReleaseDeviceRequested uintptr
	Add_ImagePreviewReceived      uintptr
	Remove_ImagePreviewReceived   uintptr
	Add_ErrorOccurred             uintptr
	Remove_ErrorOccurred          uintptr
}

type IClaimedBarcodeScanner struct {
	win32.IInspectable
}

func (this *IClaimedBarcodeScanner) Vtbl() *IClaimedBarcodeScannerVtbl {
	return (*IClaimedBarcodeScannerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClaimedBarcodeScanner) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IClaimedBarcodeScanner) Get_IsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedBarcodeScanner) Put_IsDisabledOnDataReceived(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsDisabledOnDataReceived, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IClaimedBarcodeScanner) Get_IsDisabledOnDataReceived() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDisabledOnDataReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedBarcodeScanner) Put_IsDecodeDataEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsDecodeDataEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IClaimedBarcodeScanner) Get_IsDecodeDataEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDecodeDataEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedBarcodeScanner) EnableAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnableAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedBarcodeScanner) DisableAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DisableAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedBarcodeScanner) RetainDevice() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RetainDevice, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IClaimedBarcodeScanner) SetActiveSymbologiesAsync(symbologies *IIterable[uint32]) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetActiveSymbologiesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(symbologies)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedBarcodeScanner) ResetStatisticsAsync(statisticsCategories *IIterable[string]) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ResetStatisticsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(statisticsCategories)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedBarcodeScanner) UpdateStatisticsAsync(statistics *IIterable[*IKeyValuePair[string, string]]) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UpdateStatisticsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(statistics)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedBarcodeScanner) SetActiveProfileAsync(profile string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetActiveProfileAsync, uintptr(unsafe.Pointer(this)), NewHStr(profile).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedBarcodeScanner) Add_DataReceived(handler TypedEventHandler[*IClaimedBarcodeScanner, *IBarcodeScannerDataReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DataReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedBarcodeScanner) Remove_DataReceived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DataReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IClaimedBarcodeScanner) Add_TriggerPressed(handler EventHandler[*IClaimedBarcodeScanner]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_TriggerPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedBarcodeScanner) Remove_TriggerPressed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_TriggerPressed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IClaimedBarcodeScanner) Add_TriggerReleased(handler EventHandler[*IClaimedBarcodeScanner]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_TriggerReleased, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedBarcodeScanner) Remove_TriggerReleased(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_TriggerReleased, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IClaimedBarcodeScanner) Add_ReleaseDeviceRequested(handler EventHandler[*IClaimedBarcodeScanner]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ReleaseDeviceRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedBarcodeScanner) Remove_ReleaseDeviceRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ReleaseDeviceRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IClaimedBarcodeScanner) Add_ImagePreviewReceived(handler TypedEventHandler[*IClaimedBarcodeScanner, *IBarcodeScannerImagePreviewReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ImagePreviewReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedBarcodeScanner) Remove_ImagePreviewReceived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ImagePreviewReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IClaimedBarcodeScanner) Add_ErrorOccurred(handler TypedEventHandler[*IClaimedBarcodeScanner, *IBarcodeScannerErrorOccurredEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ErrorOccurred, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedBarcodeScanner) Remove_ErrorOccurred(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ErrorOccurred, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// F61AAD0C-8551-42B4-998C-970C20210A22
var IID_IClaimedBarcodeScanner1 = syscall.GUID{0xF61AAD0C, 0x8551, 0x42B4,
	[8]byte{0x99, 0x8C, 0x97, 0x0C, 0x20, 0x21, 0x0A, 0x22}}

type IClaimedBarcodeScanner1Interface interface {
	win32.IInspectableInterface
	StartSoftwareTriggerAsync() *IAsyncAction
	StopSoftwareTriggerAsync() *IAsyncAction
}

type IClaimedBarcodeScanner1Vtbl struct {
	win32.IInspectableVtbl
	StartSoftwareTriggerAsync uintptr
	StopSoftwareTriggerAsync  uintptr
}

type IClaimedBarcodeScanner1 struct {
	win32.IInspectable
}

func (this *IClaimedBarcodeScanner1) Vtbl() *IClaimedBarcodeScanner1Vtbl {
	return (*IClaimedBarcodeScanner1Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClaimedBarcodeScanner1) StartSoftwareTriggerAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartSoftwareTriggerAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedBarcodeScanner1) StopSoftwareTriggerAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopSoftwareTriggerAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E3B59E8C-2D8B-4F70-8AF3-3448BEDD5FE2
var IID_IClaimedBarcodeScanner2 = syscall.GUID{0xE3B59E8C, 0x2D8B, 0x4F70,
	[8]byte{0x8A, 0xF3, 0x34, 0x48, 0xBE, 0xDD, 0x5F, 0xE2}}

type IClaimedBarcodeScanner2Interface interface {
	win32.IInspectableInterface
	GetSymbologyAttributesAsync(barcodeSymbology uint32) *IAsyncOperation[*IBarcodeSymbologyAttributes]
	SetSymbologyAttributesAsync(barcodeSymbology uint32, attributes *IBarcodeSymbologyAttributes) *IAsyncOperation[bool]
}

type IClaimedBarcodeScanner2Vtbl struct {
	win32.IInspectableVtbl
	GetSymbologyAttributesAsync uintptr
	SetSymbologyAttributesAsync uintptr
}

type IClaimedBarcodeScanner2 struct {
	win32.IInspectable
}

func (this *IClaimedBarcodeScanner2) Vtbl() *IClaimedBarcodeScanner2Vtbl {
	return (*IClaimedBarcodeScanner2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClaimedBarcodeScanner2) GetSymbologyAttributesAsync(barcodeSymbology uint32) *IAsyncOperation[*IBarcodeSymbologyAttributes] {
	var _result *IAsyncOperation[*IBarcodeSymbologyAttributes]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSymbologyAttributesAsync, uintptr(unsafe.Pointer(this)), uintptr(barcodeSymbology), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedBarcodeScanner2) SetSymbologyAttributesAsync(barcodeSymbology uint32, attributes *IBarcodeSymbologyAttributes) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetSymbologyAttributesAsync, uintptr(unsafe.Pointer(this)), uintptr(barcodeSymbology), uintptr(unsafe.Pointer(attributes)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E6CEB430-712E-45FC-8B86-CD55F5AEF79D
var IID_IClaimedBarcodeScanner3 = syscall.GUID{0xE6CEB430, 0x712E, 0x45FC,
	[8]byte{0x8B, 0x86, 0xCD, 0x55, 0xF5, 0xAE, 0xF7, 0x9D}}

type IClaimedBarcodeScanner3Interface interface {
	win32.IInspectableInterface
	ShowVideoPreviewAsync() *IAsyncOperation[bool]
	HideVideoPreview()
	Put_IsVideoPreviewShownOnEnable(value bool)
	Get_IsVideoPreviewShownOnEnable() bool
}

type IClaimedBarcodeScanner3Vtbl struct {
	win32.IInspectableVtbl
	ShowVideoPreviewAsync           uintptr
	HideVideoPreview                uintptr
	Put_IsVideoPreviewShownOnEnable uintptr
	Get_IsVideoPreviewShownOnEnable uintptr
}

type IClaimedBarcodeScanner3 struct {
	win32.IInspectable
}

func (this *IClaimedBarcodeScanner3) Vtbl() *IClaimedBarcodeScanner3Vtbl {
	return (*IClaimedBarcodeScanner3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClaimedBarcodeScanner3) ShowVideoPreviewAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowVideoPreviewAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedBarcodeScanner3) HideVideoPreview() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().HideVideoPreview, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IClaimedBarcodeScanner3) Put_IsVideoPreviewShownOnEnable(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsVideoPreviewShownOnEnable, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IClaimedBarcodeScanner3) Get_IsVideoPreviewShownOnEnable() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsVideoPreviewShownOnEnable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5D501F97-376A-41A8-A230-2F37C1949DDE
var IID_IClaimedBarcodeScanner4 = syscall.GUID{0x5D501F97, 0x376A, 0x41A8,
	[8]byte{0xA2, 0x30, 0x2F, 0x37, 0xC1, 0x94, 0x9D, 0xDE}}

type IClaimedBarcodeScanner4Interface interface {
	win32.IInspectableInterface
	Add_Closed(handler TypedEventHandler[*IClaimedBarcodeScanner, *IClaimedBarcodeScannerClosedEventArgs]) EventRegistrationToken
	Remove_Closed(token EventRegistrationToken)
}

type IClaimedBarcodeScanner4Vtbl struct {
	win32.IInspectableVtbl
	Add_Closed    uintptr
	Remove_Closed uintptr
}

type IClaimedBarcodeScanner4 struct {
	win32.IInspectable
}

func (this *IClaimedBarcodeScanner4) Vtbl() *IClaimedBarcodeScanner4Vtbl {
	return (*IClaimedBarcodeScanner4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClaimedBarcodeScanner4) Add_Closed(handler TypedEventHandler[*IClaimedBarcodeScanner, *IClaimedBarcodeScannerClosedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Closed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedBarcodeScanner4) Remove_Closed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Closed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// CF7D5489-A22C-4C65-A901-88D77D833954
var IID_IClaimedBarcodeScannerClosedEventArgs = syscall.GUID{0xCF7D5489, 0xA22C, 0x4C65,
	[8]byte{0xA9, 0x01, 0x88, 0xD7, 0x7D, 0x83, 0x39, 0x54}}

type IClaimedBarcodeScannerClosedEventArgsInterface interface {
	win32.IInspectableInterface
}

type IClaimedBarcodeScannerClosedEventArgsVtbl struct {
	win32.IInspectableVtbl
}

type IClaimedBarcodeScannerClosedEventArgs struct {
	win32.IInspectable
}

func (this *IClaimedBarcodeScannerClosedEventArgs) Vtbl() *IClaimedBarcodeScannerClosedEventArgsVtbl {
	return (*IClaimedBarcodeScannerClosedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// CA3F99AF-ABB8-42C1-8A84-5C66512F5A75
var IID_IClaimedCashDrawer = syscall.GUID{0xCA3F99AF, 0xABB8, 0x42C1,
	[8]byte{0x8A, 0x84, 0x5C, 0x66, 0x51, 0x2F, 0x5A, 0x75}}

type IClaimedCashDrawerInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_IsEnabled() bool
	Get_IsDrawerOpen() bool
	Get_CloseAlarm() *ICashDrawerCloseAlarm
	OpenDrawerAsync() *IAsyncOperation[bool]
	EnableAsync() *IAsyncOperation[bool]
	DisableAsync() *IAsyncOperation[bool]
	RetainDeviceAsync() *IAsyncOperation[bool]
	ResetStatisticsAsync(statisticsCategories *IIterable[string]) *IAsyncOperation[bool]
	UpdateStatisticsAsync(statistics *IIterable[*IKeyValuePair[string, string]]) *IAsyncOperation[bool]
	Add_ReleaseDeviceRequested(handler TypedEventHandler[*IClaimedCashDrawer, interface{}]) EventRegistrationToken
	Remove_ReleaseDeviceRequested(token EventRegistrationToken)
}

type IClaimedCashDrawerVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId                  uintptr
	Get_IsEnabled                 uintptr
	Get_IsDrawerOpen              uintptr
	Get_CloseAlarm                uintptr
	OpenDrawerAsync               uintptr
	EnableAsync                   uintptr
	DisableAsync                  uintptr
	RetainDeviceAsync             uintptr
	ResetStatisticsAsync          uintptr
	UpdateStatisticsAsync         uintptr
	Add_ReleaseDeviceRequested    uintptr
	Remove_ReleaseDeviceRequested uintptr
}

type IClaimedCashDrawer struct {
	win32.IInspectable
}

func (this *IClaimedCashDrawer) Vtbl() *IClaimedCashDrawerVtbl {
	return (*IClaimedCashDrawerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClaimedCashDrawer) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IClaimedCashDrawer) Get_IsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedCashDrawer) Get_IsDrawerOpen() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDrawerOpen, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedCashDrawer) Get_CloseAlarm() *ICashDrawerCloseAlarm {
	var _result *ICashDrawerCloseAlarm
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CloseAlarm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedCashDrawer) OpenDrawerAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenDrawerAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedCashDrawer) EnableAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnableAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedCashDrawer) DisableAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DisableAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedCashDrawer) RetainDeviceAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RetainDeviceAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedCashDrawer) ResetStatisticsAsync(statisticsCategories *IIterable[string]) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ResetStatisticsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(statisticsCategories)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedCashDrawer) UpdateStatisticsAsync(statistics *IIterable[*IKeyValuePair[string, string]]) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UpdateStatisticsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(statistics)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedCashDrawer) Add_ReleaseDeviceRequested(handler TypedEventHandler[*IClaimedCashDrawer, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ReleaseDeviceRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedCashDrawer) Remove_ReleaseDeviceRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ReleaseDeviceRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 9CBAB5A2-DE42-4D5B-B0C1-9B57A2BA89C3
var IID_IClaimedCashDrawer2 = syscall.GUID{0x9CBAB5A2, 0xDE42, 0x4D5B,
	[8]byte{0xB0, 0xC1, 0x9B, 0x57, 0xA2, 0xBA, 0x89, 0xC3}}

type IClaimedCashDrawer2Interface interface {
	win32.IInspectableInterface
	Add_Closed(handler TypedEventHandler[*IClaimedCashDrawer, *IClaimedCashDrawerClosedEventArgs]) EventRegistrationToken
	Remove_Closed(token EventRegistrationToken)
}

type IClaimedCashDrawer2Vtbl struct {
	win32.IInspectableVtbl
	Add_Closed    uintptr
	Remove_Closed uintptr
}

type IClaimedCashDrawer2 struct {
	win32.IInspectable
}

func (this *IClaimedCashDrawer2) Vtbl() *IClaimedCashDrawer2Vtbl {
	return (*IClaimedCashDrawer2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClaimedCashDrawer2) Add_Closed(handler TypedEventHandler[*IClaimedCashDrawer, *IClaimedCashDrawerClosedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Closed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedCashDrawer2) Remove_Closed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Closed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// CC573F33-3F34-4C5C-BAAE-DEADF16CD7FA
var IID_IClaimedCashDrawerClosedEventArgs = syscall.GUID{0xCC573F33, 0x3F34, 0x4C5C,
	[8]byte{0xBA, 0xAE, 0xDE, 0xAD, 0xF1, 0x6C, 0xD7, 0xFA}}

type IClaimedCashDrawerClosedEventArgsInterface interface {
	win32.IInspectableInterface
}

type IClaimedCashDrawerClosedEventArgsVtbl struct {
	win32.IInspectableVtbl
}

type IClaimedCashDrawerClosedEventArgs struct {
	win32.IInspectable
}

func (this *IClaimedCashDrawerClosedEventArgs) Vtbl() *IClaimedCashDrawerClosedEventArgsVtbl {
	return (*IClaimedCashDrawerClosedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 67EA0630-517D-487F-9FDF-D2E0A0A264A5
var IID_IClaimedJournalPrinter = syscall.GUID{0x67EA0630, 0x517D, 0x487F,
	[8]byte{0x9F, 0xDF, 0xD2, 0xE0, 0xA0, 0xA2, 0x64, 0xA5}}

type IClaimedJournalPrinterInterface interface {
	win32.IInspectableInterface
	CreateJob() *IPosPrinterJob
}

type IClaimedJournalPrinterVtbl struct {
	win32.IInspectableVtbl
	CreateJob uintptr
}

type IClaimedJournalPrinter struct {
	win32.IInspectable
}

func (this *IClaimedJournalPrinter) Vtbl() *IClaimedJournalPrinterVtbl {
	return (*IClaimedJournalPrinterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClaimedJournalPrinter) CreateJob() *IPosPrinterJob {
	var _result *IPosPrinterJob
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateJob, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 120AC970-9A75-4ACF-AAE7-09972BCF8794
var IID_IClaimedLineDisplay = syscall.GUID{0x120AC970, 0x9A75, 0x4ACF,
	[8]byte{0xAA, 0xE7, 0x09, 0x97, 0x2B, 0xCF, 0x87, 0x94}}

type IClaimedLineDisplayInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_Capabilities() *ILineDisplayCapabilities
	Get_PhysicalDeviceName() string
	Get_PhysicalDeviceDescription() string
	Get_DeviceControlDescription() string
	Get_DeviceControlVersion() string
	Get_DeviceServiceVersion() string
	Get_DefaultWindow() *ILineDisplayWindow
	RetainDevice()
	Add_ReleaseDeviceRequested(handler TypedEventHandler[*IClaimedLineDisplay, interface{}]) EventRegistrationToken
	Remove_ReleaseDeviceRequested(token EventRegistrationToken)
}

type IClaimedLineDisplayVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId                  uintptr
	Get_Capabilities              uintptr
	Get_PhysicalDeviceName        uintptr
	Get_PhysicalDeviceDescription uintptr
	Get_DeviceControlDescription  uintptr
	Get_DeviceControlVersion      uintptr
	Get_DeviceServiceVersion      uintptr
	Get_DefaultWindow             uintptr
	RetainDevice                  uintptr
	Add_ReleaseDeviceRequested    uintptr
	Remove_ReleaseDeviceRequested uintptr
}

type IClaimedLineDisplay struct {
	win32.IInspectable
}

func (this *IClaimedLineDisplay) Vtbl() *IClaimedLineDisplayVtbl {
	return (*IClaimedLineDisplayVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClaimedLineDisplay) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IClaimedLineDisplay) Get_Capabilities() *ILineDisplayCapabilities {
	var _result *ILineDisplayCapabilities
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Capabilities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedLineDisplay) Get_PhysicalDeviceName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhysicalDeviceName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IClaimedLineDisplay) Get_PhysicalDeviceDescription() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhysicalDeviceDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IClaimedLineDisplay) Get_DeviceControlDescription() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceControlDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IClaimedLineDisplay) Get_DeviceControlVersion() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceControlVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IClaimedLineDisplay) Get_DeviceServiceVersion() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceServiceVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IClaimedLineDisplay) Get_DefaultWindow() *ILineDisplayWindow {
	var _result *ILineDisplayWindow
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DefaultWindow, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedLineDisplay) RetainDevice() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RetainDevice, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IClaimedLineDisplay) Add_ReleaseDeviceRequested(handler TypedEventHandler[*IClaimedLineDisplay, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ReleaseDeviceRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedLineDisplay) Remove_ReleaseDeviceRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ReleaseDeviceRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// A31C75ED-41F5-4E76-A074-795E47A46E97
var IID_IClaimedLineDisplay2 = syscall.GUID{0xA31C75ED, 0x41F5, 0x4E76,
	[8]byte{0xA0, 0x74, 0x79, 0x5E, 0x47, 0xA4, 0x6E, 0x97}}

type IClaimedLineDisplay2Interface interface {
	win32.IInspectableInterface
	GetStatisticsAsync(statisticsCategories *IIterable[string]) *IAsyncOperation[string]
	CheckHealthAsync(level UnifiedPosHealthCheckLevel) *IAsyncOperation[string]
	CheckPowerStatusAsync() *IAsyncOperation[LineDisplayPowerStatus]
	Add_StatusUpdated(handler TypedEventHandler[*IClaimedLineDisplay, *ILineDisplayStatusUpdatedEventArgs]) EventRegistrationToken
	Remove_StatusUpdated(token EventRegistrationToken)
	Get_SupportedScreenSizesInCharacters() *IVectorView[Size]
	Get_MaxBitmapSizeInPixels() Size
	Get_SupportedCharacterSets() *IVectorView[int32]
	Get_CustomGlyphs() *ILineDisplayCustomGlyphs
	GetAttributes() *ILineDisplayAttributes
	TryUpdateAttributesAsync(attributes *ILineDisplayAttributes) *IAsyncOperation[bool]
	TrySetDescriptorAsync(descriptor uint32, descriptorState LineDisplayDescriptorState) *IAsyncOperation[bool]
	TryClearDescriptorsAsync() *IAsyncOperation[bool]
	TryCreateWindowAsync(viewport Rect, windowSize Size) *IAsyncOperation[*ILineDisplayWindow]
	TryStoreStorageFileBitmapAsync(bitmap *IStorageFile) *IAsyncOperation[*ILineDisplayStoredBitmap]
	TryStoreStorageFileBitmapWithAlignmentAsync(bitmap *IStorageFile, horizontalAlignment LineDisplayHorizontalAlignment, verticalAlignment LineDisplayVerticalAlignment) *IAsyncOperation[*ILineDisplayStoredBitmap]
	TryStoreStorageFileBitmapWithAlignmentAndWidthAsync(bitmap *IStorageFile, horizontalAlignment LineDisplayHorizontalAlignment, verticalAlignment LineDisplayVerticalAlignment, widthInPixels int32) *IAsyncOperation[*ILineDisplayStoredBitmap]
}

type IClaimedLineDisplay2Vtbl struct {
	win32.IInspectableVtbl
	GetStatisticsAsync                                  uintptr
	CheckHealthAsync                                    uintptr
	CheckPowerStatusAsync                               uintptr
	Add_StatusUpdated                                   uintptr
	Remove_StatusUpdated                                uintptr
	Get_SupportedScreenSizesInCharacters                uintptr
	Get_MaxBitmapSizeInPixels                           uintptr
	Get_SupportedCharacterSets                          uintptr
	Get_CustomGlyphs                                    uintptr
	GetAttributes                                       uintptr
	TryUpdateAttributesAsync                            uintptr
	TrySetDescriptorAsync                               uintptr
	TryClearDescriptorsAsync                            uintptr
	TryCreateWindowAsync                                uintptr
	TryStoreStorageFileBitmapAsync                      uintptr
	TryStoreStorageFileBitmapWithAlignmentAsync         uintptr
	TryStoreStorageFileBitmapWithAlignmentAndWidthAsync uintptr
}

type IClaimedLineDisplay2 struct {
	win32.IInspectable
}

func (this *IClaimedLineDisplay2) Vtbl() *IClaimedLineDisplay2Vtbl {
	return (*IClaimedLineDisplay2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClaimedLineDisplay2) GetStatisticsAsync(statisticsCategories *IIterable[string]) *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStatisticsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(statisticsCategories)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedLineDisplay2) CheckHealthAsync(level UnifiedPosHealthCheckLevel) *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CheckHealthAsync, uintptr(unsafe.Pointer(this)), uintptr(level), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedLineDisplay2) CheckPowerStatusAsync() *IAsyncOperation[LineDisplayPowerStatus] {
	var _result *IAsyncOperation[LineDisplayPowerStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CheckPowerStatusAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedLineDisplay2) Add_StatusUpdated(handler TypedEventHandler[*IClaimedLineDisplay, *ILineDisplayStatusUpdatedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StatusUpdated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedLineDisplay2) Remove_StatusUpdated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StatusUpdated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IClaimedLineDisplay2) Get_SupportedScreenSizesInCharacters() *IVectorView[Size] {
	var _result *IVectorView[Size]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedScreenSizesInCharacters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedLineDisplay2) Get_MaxBitmapSizeInPixels() Size {
	var _result Size
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxBitmapSizeInPixels, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedLineDisplay2) Get_SupportedCharacterSets() *IVectorView[int32] {
	var _result *IVectorView[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedCharacterSets, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedLineDisplay2) Get_CustomGlyphs() *ILineDisplayCustomGlyphs {
	var _result *ILineDisplayCustomGlyphs
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CustomGlyphs, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedLineDisplay2) GetAttributes() *ILineDisplayAttributes {
	var _result *ILineDisplayAttributes
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAttributes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedLineDisplay2) TryUpdateAttributesAsync(attributes *ILineDisplayAttributes) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryUpdateAttributesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(attributes)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedLineDisplay2) TrySetDescriptorAsync(descriptor uint32, descriptorState LineDisplayDescriptorState) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySetDescriptorAsync, uintptr(unsafe.Pointer(this)), uintptr(descriptor), uintptr(descriptorState), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedLineDisplay2) TryClearDescriptorsAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryClearDescriptorsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedLineDisplay2) TryCreateWindowAsync(viewport Rect, windowSize Size) *IAsyncOperation[*ILineDisplayWindow] {
	var _result *IAsyncOperation[*ILineDisplayWindow]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryCreateWindowAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&viewport)), *(*uintptr)(unsafe.Pointer(&windowSize)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedLineDisplay2) TryStoreStorageFileBitmapAsync(bitmap *IStorageFile) *IAsyncOperation[*ILineDisplayStoredBitmap] {
	var _result *IAsyncOperation[*ILineDisplayStoredBitmap]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryStoreStorageFileBitmapAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bitmap)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedLineDisplay2) TryStoreStorageFileBitmapWithAlignmentAsync(bitmap *IStorageFile, horizontalAlignment LineDisplayHorizontalAlignment, verticalAlignment LineDisplayVerticalAlignment) *IAsyncOperation[*ILineDisplayStoredBitmap] {
	var _result *IAsyncOperation[*ILineDisplayStoredBitmap]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryStoreStorageFileBitmapWithAlignmentAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bitmap)), uintptr(horizontalAlignment), uintptr(verticalAlignment), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedLineDisplay2) TryStoreStorageFileBitmapWithAlignmentAndWidthAsync(bitmap *IStorageFile, horizontalAlignment LineDisplayHorizontalAlignment, verticalAlignment LineDisplayVerticalAlignment, widthInPixels int32) *IAsyncOperation[*ILineDisplayStoredBitmap] {
	var _result *IAsyncOperation[*ILineDisplayStoredBitmap]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryStoreStorageFileBitmapWithAlignmentAndWidthAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bitmap)), uintptr(horizontalAlignment), uintptr(verticalAlignment), uintptr(widthInPixels), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 642ECD92-E9D4-4ECC-AF75-329C274CD18F
var IID_IClaimedLineDisplay3 = syscall.GUID{0x642ECD92, 0xE9D4, 0x4ECC,
	[8]byte{0xAF, 0x75, 0x32, 0x9C, 0x27, 0x4C, 0xD1, 0x8F}}

type IClaimedLineDisplay3Interface interface {
	win32.IInspectableInterface
	Add_Closed(handler TypedEventHandler[*IClaimedLineDisplay, *IClaimedLineDisplayClosedEventArgs]) EventRegistrationToken
	Remove_Closed(token EventRegistrationToken)
}

type IClaimedLineDisplay3Vtbl struct {
	win32.IInspectableVtbl
	Add_Closed    uintptr
	Remove_Closed uintptr
}

type IClaimedLineDisplay3 struct {
	win32.IInspectable
}

func (this *IClaimedLineDisplay3) Vtbl() *IClaimedLineDisplay3Vtbl {
	return (*IClaimedLineDisplay3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClaimedLineDisplay3) Add_Closed(handler TypedEventHandler[*IClaimedLineDisplay, *IClaimedLineDisplayClosedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Closed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedLineDisplay3) Remove_Closed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Closed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// F915F364-D3D5-4F10-B511-90939EDFACD8
var IID_IClaimedLineDisplayClosedEventArgs = syscall.GUID{0xF915F364, 0xD3D5, 0x4F10,
	[8]byte{0xB5, 0x11, 0x90, 0x93, 0x9E, 0xDF, 0xAC, 0xD8}}

type IClaimedLineDisplayClosedEventArgsInterface interface {
	win32.IInspectableInterface
}

type IClaimedLineDisplayClosedEventArgsVtbl struct {
	win32.IInspectableVtbl
}

type IClaimedLineDisplayClosedEventArgs struct {
	win32.IInspectable
}

func (this *IClaimedLineDisplayClosedEventArgs) Vtbl() *IClaimedLineDisplayClosedEventArgsVtbl {
	return (*IClaimedLineDisplayClosedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 78CA98FB-8B6B-4973-86F0-3E570C351825
var IID_IClaimedLineDisplayStatics = syscall.GUID{0x78CA98FB, 0x8B6B, 0x4973,
	[8]byte{0x86, 0xF0, 0x3E, 0x57, 0x0C, 0x35, 0x18, 0x25}}

type IClaimedLineDisplayStaticsInterface interface {
	win32.IInspectableInterface
	FromIdAsync(deviceId string) *IAsyncOperation[*IClaimedLineDisplay]
	GetDeviceSelector() string
	GetDeviceSelectorWithConnectionTypes(connectionTypes PosConnectionTypes) string
}

type IClaimedLineDisplayStaticsVtbl struct {
	win32.IInspectableVtbl
	FromIdAsync                          uintptr
	GetDeviceSelector                    uintptr
	GetDeviceSelectorWithConnectionTypes uintptr
}

type IClaimedLineDisplayStatics struct {
	win32.IInspectable
}

func (this *IClaimedLineDisplayStatics) Vtbl() *IClaimedLineDisplayStaticsVtbl {
	return (*IClaimedLineDisplayStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClaimedLineDisplayStatics) FromIdAsync(deviceId string) *IAsyncOperation[*IClaimedLineDisplay] {
	var _result *IAsyncOperation[*IClaimedLineDisplay]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedLineDisplayStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IClaimedLineDisplayStatics) GetDeviceSelectorWithConnectionTypes(connectionTypes PosConnectionTypes) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorWithConnectionTypes, uintptr(unsafe.Pointer(this)), uintptr(connectionTypes), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 475CA8F3-9417-48BC-B9D7-4163A7844C02
var IID_IClaimedMagneticStripeReader = syscall.GUID{0x475CA8F3, 0x9417, 0x48BC,
	[8]byte{0xB9, 0xD7, 0x41, 0x63, 0xA7, 0x84, 0x4C, 0x02}}

type IClaimedMagneticStripeReaderInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_IsEnabled() bool
	Put_IsDisabledOnDataReceived(value bool)
	Get_IsDisabledOnDataReceived() bool
	Put_IsDecodeDataEnabled(value bool)
	Get_IsDecodeDataEnabled() bool
	Get_IsDeviceAuthenticated() bool
	Put_DataEncryptionAlgorithm(value uint32)
	Get_DataEncryptionAlgorithm() uint32
	Put_TracksToRead(value MagneticStripeReaderTrackIds)
	Get_TracksToRead() MagneticStripeReaderTrackIds
	Put_IsTransmitSentinelsEnabled(value bool)
	Get_IsTransmitSentinelsEnabled() bool
	EnableAsync() *IAsyncAction
	DisableAsync() *IAsyncAction
	RetainDevice()
	SetErrorReportingType(value MagneticStripeReaderErrorReportingType)
	RetrieveDeviceAuthenticationDataAsync() *IAsyncOperation[*IBuffer]
	AuthenticateDeviceAsync(responseTokenLength uint32, responseToken *byte) *IAsyncAction
	DeAuthenticateDeviceAsync(responseTokenLength uint32, responseToken *byte) *IAsyncAction
	UpdateKeyAsync(key string, keyName string) *IAsyncAction
	ResetStatisticsAsync(statisticsCategories *IIterable[string]) *IAsyncAction
	UpdateStatisticsAsync(statistics *IIterable[*IKeyValuePair[string, string]]) *IAsyncAction
	Add_BankCardDataReceived(handler TypedEventHandler[*IClaimedMagneticStripeReader, *IMagneticStripeReaderBankCardDataReceivedEventArgs]) EventRegistrationToken
	Remove_BankCardDataReceived(token EventRegistrationToken)
	Add_AamvaCardDataReceived(handler TypedEventHandler[*IClaimedMagneticStripeReader, *IMagneticStripeReaderAamvaCardDataReceivedEventArgs]) EventRegistrationToken
	Remove_AamvaCardDataReceived(token EventRegistrationToken)
	Add_VendorSpecificDataReceived(handler TypedEventHandler[*IClaimedMagneticStripeReader, *IMagneticStripeReaderVendorSpecificCardDataReceivedEventArgs]) EventRegistrationToken
	Remove_VendorSpecificDataReceived(token EventRegistrationToken)
	Add_ReleaseDeviceRequested(handler EventHandler[*IClaimedMagneticStripeReader]) EventRegistrationToken
	Remove_ReleaseDeviceRequested(token EventRegistrationToken)
	Add_ErrorOccurred(handler TypedEventHandler[*IClaimedMagneticStripeReader, *IMagneticStripeReaderErrorOccurredEventArgs]) EventRegistrationToken
	Remove_ErrorOccurred(token EventRegistrationToken)
}

type IClaimedMagneticStripeReaderVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId                          uintptr
	Get_IsEnabled                         uintptr
	Put_IsDisabledOnDataReceived          uintptr
	Get_IsDisabledOnDataReceived          uintptr
	Put_IsDecodeDataEnabled               uintptr
	Get_IsDecodeDataEnabled               uintptr
	Get_IsDeviceAuthenticated             uintptr
	Put_DataEncryptionAlgorithm           uintptr
	Get_DataEncryptionAlgorithm           uintptr
	Put_TracksToRead                      uintptr
	Get_TracksToRead                      uintptr
	Put_IsTransmitSentinelsEnabled        uintptr
	Get_IsTransmitSentinelsEnabled        uintptr
	EnableAsync                           uintptr
	DisableAsync                          uintptr
	RetainDevice                          uintptr
	SetErrorReportingType                 uintptr
	RetrieveDeviceAuthenticationDataAsync uintptr
	AuthenticateDeviceAsync               uintptr
	DeAuthenticateDeviceAsync             uintptr
	UpdateKeyAsync                        uintptr
	ResetStatisticsAsync                  uintptr
	UpdateStatisticsAsync                 uintptr
	Add_BankCardDataReceived              uintptr
	Remove_BankCardDataReceived           uintptr
	Add_AamvaCardDataReceived             uintptr
	Remove_AamvaCardDataReceived          uintptr
	Add_VendorSpecificDataReceived        uintptr
	Remove_VendorSpecificDataReceived     uintptr
	Add_ReleaseDeviceRequested            uintptr
	Remove_ReleaseDeviceRequested         uintptr
	Add_ErrorOccurred                     uintptr
	Remove_ErrorOccurred                  uintptr
}

type IClaimedMagneticStripeReader struct {
	win32.IInspectable
}

func (this *IClaimedMagneticStripeReader) Vtbl() *IClaimedMagneticStripeReaderVtbl {
	return (*IClaimedMagneticStripeReaderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClaimedMagneticStripeReader) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IClaimedMagneticStripeReader) Get_IsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedMagneticStripeReader) Put_IsDisabledOnDataReceived(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsDisabledOnDataReceived, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IClaimedMagneticStripeReader) Get_IsDisabledOnDataReceived() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDisabledOnDataReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedMagneticStripeReader) Put_IsDecodeDataEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsDecodeDataEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IClaimedMagneticStripeReader) Get_IsDecodeDataEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDecodeDataEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedMagneticStripeReader) Get_IsDeviceAuthenticated() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDeviceAuthenticated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedMagneticStripeReader) Put_DataEncryptionAlgorithm(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DataEncryptionAlgorithm, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IClaimedMagneticStripeReader) Get_DataEncryptionAlgorithm() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DataEncryptionAlgorithm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedMagneticStripeReader) Put_TracksToRead(value MagneticStripeReaderTrackIds) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TracksToRead, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IClaimedMagneticStripeReader) Get_TracksToRead() MagneticStripeReaderTrackIds {
	var _result MagneticStripeReaderTrackIds
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TracksToRead, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedMagneticStripeReader) Put_IsTransmitSentinelsEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsTransmitSentinelsEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IClaimedMagneticStripeReader) Get_IsTransmitSentinelsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsTransmitSentinelsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedMagneticStripeReader) EnableAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnableAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedMagneticStripeReader) DisableAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DisableAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedMagneticStripeReader) RetainDevice() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RetainDevice, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IClaimedMagneticStripeReader) SetErrorReportingType(value MagneticStripeReaderErrorReportingType) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetErrorReportingType, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IClaimedMagneticStripeReader) RetrieveDeviceAuthenticationDataAsync() *IAsyncOperation[*IBuffer] {
	var _result *IAsyncOperation[*IBuffer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RetrieveDeviceAuthenticationDataAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedMagneticStripeReader) AuthenticateDeviceAsync(responseTokenLength uint32, responseToken *byte) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AuthenticateDeviceAsync, uintptr(unsafe.Pointer(this)), uintptr(responseTokenLength), uintptr(unsafe.Pointer(responseToken)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedMagneticStripeReader) DeAuthenticateDeviceAsync(responseTokenLength uint32, responseToken *byte) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DeAuthenticateDeviceAsync, uintptr(unsafe.Pointer(this)), uintptr(responseTokenLength), uintptr(unsafe.Pointer(responseToken)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedMagneticStripeReader) UpdateKeyAsync(key string, keyName string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UpdateKeyAsync, uintptr(unsafe.Pointer(this)), NewHStr(key).Ptr, NewHStr(keyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedMagneticStripeReader) ResetStatisticsAsync(statisticsCategories *IIterable[string]) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ResetStatisticsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(statisticsCategories)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedMagneticStripeReader) UpdateStatisticsAsync(statistics *IIterable[*IKeyValuePair[string, string]]) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UpdateStatisticsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(statistics)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedMagneticStripeReader) Add_BankCardDataReceived(handler TypedEventHandler[*IClaimedMagneticStripeReader, *IMagneticStripeReaderBankCardDataReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_BankCardDataReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedMagneticStripeReader) Remove_BankCardDataReceived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_BankCardDataReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IClaimedMagneticStripeReader) Add_AamvaCardDataReceived(handler TypedEventHandler[*IClaimedMagneticStripeReader, *IMagneticStripeReaderAamvaCardDataReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AamvaCardDataReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedMagneticStripeReader) Remove_AamvaCardDataReceived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AamvaCardDataReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IClaimedMagneticStripeReader) Add_VendorSpecificDataReceived(handler TypedEventHandler[*IClaimedMagneticStripeReader, *IMagneticStripeReaderVendorSpecificCardDataReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_VendorSpecificDataReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedMagneticStripeReader) Remove_VendorSpecificDataReceived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_VendorSpecificDataReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IClaimedMagneticStripeReader) Add_ReleaseDeviceRequested(handler EventHandler[*IClaimedMagneticStripeReader]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ReleaseDeviceRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedMagneticStripeReader) Remove_ReleaseDeviceRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ReleaseDeviceRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IClaimedMagneticStripeReader) Add_ErrorOccurred(handler TypedEventHandler[*IClaimedMagneticStripeReader, *IMagneticStripeReaderErrorOccurredEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ErrorOccurred, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedMagneticStripeReader) Remove_ErrorOccurred(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ErrorOccurred, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 236FAFDF-E2DC-4D7D-9C78-060DF2BF2928
var IID_IClaimedMagneticStripeReader2 = syscall.GUID{0x236FAFDF, 0xE2DC, 0x4D7D,
	[8]byte{0x9C, 0x78, 0x06, 0x0D, 0xF2, 0xBF, 0x29, 0x28}}

type IClaimedMagneticStripeReader2Interface interface {
	win32.IInspectableInterface
	Add_Closed(handler TypedEventHandler[*IClaimedMagneticStripeReader, *IClaimedMagneticStripeReaderClosedEventArgs]) EventRegistrationToken
	Remove_Closed(token EventRegistrationToken)
}

type IClaimedMagneticStripeReader2Vtbl struct {
	win32.IInspectableVtbl
	Add_Closed    uintptr
	Remove_Closed uintptr
}

type IClaimedMagneticStripeReader2 struct {
	win32.IInspectable
}

func (this *IClaimedMagneticStripeReader2) Vtbl() *IClaimedMagneticStripeReader2Vtbl {
	return (*IClaimedMagneticStripeReader2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClaimedMagneticStripeReader2) Add_Closed(handler TypedEventHandler[*IClaimedMagneticStripeReader, *IClaimedMagneticStripeReaderClosedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Closed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedMagneticStripeReader2) Remove_Closed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Closed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 14ADA93A-ADCD-4C80-ACDA-C3EAED2647E1
var IID_IClaimedMagneticStripeReaderClosedEventArgs = syscall.GUID{0x14ADA93A, 0xADCD, 0x4C80,
	[8]byte{0xAC, 0xDA, 0xC3, 0xEA, 0xED, 0x26, 0x47, 0xE1}}

type IClaimedMagneticStripeReaderClosedEventArgsInterface interface {
	win32.IInspectableInterface
}

type IClaimedMagneticStripeReaderClosedEventArgsVtbl struct {
	win32.IInspectableVtbl
}

type IClaimedMagneticStripeReaderClosedEventArgs struct {
	win32.IInspectable
}

func (this *IClaimedMagneticStripeReaderClosedEventArgs) Vtbl() *IClaimedMagneticStripeReaderClosedEventArgsVtbl {
	return (*IClaimedMagneticStripeReaderClosedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 6D64CE0C-E03E-4B14-A38E-C28C34B86353
var IID_IClaimedPosPrinter = syscall.GUID{0x6D64CE0C, 0xE03E, 0x4B14,
	[8]byte{0xA3, 0x8E, 0xC2, 0x8C, 0x34, 0xB8, 0x63, 0x53}}

type IClaimedPosPrinterInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_IsEnabled() bool
	Put_CharacterSet(value uint32)
	Get_CharacterSet() uint32
	Get_IsCoverOpen() bool
	Put_IsCharacterSetMappingEnabled(value bool)
	Get_IsCharacterSetMappingEnabled() bool
	Put_MapMode(value PosPrinterMapMode)
	Get_MapMode() PosPrinterMapMode
	Get_Receipt() *IClaimedReceiptPrinter
	Get_Slip() *IClaimedSlipPrinter
	Get_Journal() *IClaimedJournalPrinter
	EnableAsync() *IAsyncOperation[bool]
	DisableAsync() *IAsyncOperation[bool]
	RetainDeviceAsync() *IAsyncOperation[bool]
	ResetStatisticsAsync(statisticsCategories *IIterable[string]) *IAsyncOperation[bool]
	UpdateStatisticsAsync(statistics *IIterable[*IKeyValuePair[string, string]]) *IAsyncOperation[bool]
	Add_ReleaseDeviceRequested(handler TypedEventHandler[*IClaimedPosPrinter, *IPosPrinterReleaseDeviceRequestedEventArgs]) EventRegistrationToken
	Remove_ReleaseDeviceRequested(token EventRegistrationToken)
}

type IClaimedPosPrinterVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId                     uintptr
	Get_IsEnabled                    uintptr
	Put_CharacterSet                 uintptr
	Get_CharacterSet                 uintptr
	Get_IsCoverOpen                  uintptr
	Put_IsCharacterSetMappingEnabled uintptr
	Get_IsCharacterSetMappingEnabled uintptr
	Put_MapMode                      uintptr
	Get_MapMode                      uintptr
	Get_Receipt                      uintptr
	Get_Slip                         uintptr
	Get_Journal                      uintptr
	EnableAsync                      uintptr
	DisableAsync                     uintptr
	RetainDeviceAsync                uintptr
	ResetStatisticsAsync             uintptr
	UpdateStatisticsAsync            uintptr
	Add_ReleaseDeviceRequested       uintptr
	Remove_ReleaseDeviceRequested    uintptr
}

type IClaimedPosPrinter struct {
	win32.IInspectable
}

func (this *IClaimedPosPrinter) Vtbl() *IClaimedPosPrinterVtbl {
	return (*IClaimedPosPrinterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClaimedPosPrinter) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IClaimedPosPrinter) Get_IsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedPosPrinter) Put_CharacterSet(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CharacterSet, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IClaimedPosPrinter) Get_CharacterSet() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CharacterSet, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedPosPrinter) Get_IsCoverOpen() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCoverOpen, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedPosPrinter) Put_IsCharacterSetMappingEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsCharacterSetMappingEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IClaimedPosPrinter) Get_IsCharacterSetMappingEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCharacterSetMappingEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedPosPrinter) Put_MapMode(value PosPrinterMapMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MapMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IClaimedPosPrinter) Get_MapMode() PosPrinterMapMode {
	var _result PosPrinterMapMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MapMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedPosPrinter) Get_Receipt() *IClaimedReceiptPrinter {
	var _result *IClaimedReceiptPrinter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Receipt, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedPosPrinter) Get_Slip() *IClaimedSlipPrinter {
	var _result *IClaimedSlipPrinter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Slip, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedPosPrinter) Get_Journal() *IClaimedJournalPrinter {
	var _result *IClaimedJournalPrinter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Journal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedPosPrinter) EnableAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnableAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedPosPrinter) DisableAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DisableAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedPosPrinter) RetainDeviceAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RetainDeviceAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedPosPrinter) ResetStatisticsAsync(statisticsCategories *IIterable[string]) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ResetStatisticsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(statisticsCategories)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedPosPrinter) UpdateStatisticsAsync(statistics *IIterable[*IKeyValuePair[string, string]]) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UpdateStatisticsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(statistics)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedPosPrinter) Add_ReleaseDeviceRequested(handler TypedEventHandler[*IClaimedPosPrinter, *IPosPrinterReleaseDeviceRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ReleaseDeviceRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedPosPrinter) Remove_ReleaseDeviceRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ReleaseDeviceRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 5BF7A3D5-5198-437A-82DF-589993FA77E1
var IID_IClaimedPosPrinter2 = syscall.GUID{0x5BF7A3D5, 0x5198, 0x437A,
	[8]byte{0x82, 0xDF, 0x58, 0x99, 0x93, 0xFA, 0x77, 0xE1}}

type IClaimedPosPrinter2Interface interface {
	win32.IInspectableInterface
	Add_Closed(handler TypedEventHandler[*IClaimedPosPrinter, *IClaimedPosPrinterClosedEventArgs]) EventRegistrationToken
	Remove_Closed(token EventRegistrationToken)
}

type IClaimedPosPrinter2Vtbl struct {
	win32.IInspectableVtbl
	Add_Closed    uintptr
	Remove_Closed uintptr
}

type IClaimedPosPrinter2 struct {
	win32.IInspectable
}

func (this *IClaimedPosPrinter2) Vtbl() *IClaimedPosPrinter2Vtbl {
	return (*IClaimedPosPrinter2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClaimedPosPrinter2) Add_Closed(handler TypedEventHandler[*IClaimedPosPrinter, *IClaimedPosPrinterClosedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Closed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedPosPrinter2) Remove_Closed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Closed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// E2B7A27B-4D40-471D-92ED-63375B18C788
var IID_IClaimedPosPrinterClosedEventArgs = syscall.GUID{0xE2B7A27B, 0x4D40, 0x471D,
	[8]byte{0x92, 0xED, 0x63, 0x37, 0x5B, 0x18, 0xC7, 0x88}}

type IClaimedPosPrinterClosedEventArgsInterface interface {
	win32.IInspectableInterface
}

type IClaimedPosPrinterClosedEventArgsVtbl struct {
	win32.IInspectableVtbl
}

type IClaimedPosPrinterClosedEventArgs struct {
	win32.IInspectable
}

func (this *IClaimedPosPrinterClosedEventArgs) Vtbl() *IClaimedPosPrinterClosedEventArgsVtbl {
	return (*IClaimedPosPrinterClosedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 9AD27A74-DD61-4EE2-9837-5B5D72D538B9
var IID_IClaimedReceiptPrinter = syscall.GUID{0x9AD27A74, 0xDD61, 0x4EE2,
	[8]byte{0x98, 0x37, 0x5B, 0x5D, 0x72, 0xD5, 0x38, 0xB9}}

type IClaimedReceiptPrinterInterface interface {
	win32.IInspectableInterface
	Get_SidewaysMaxLines() uint32
	Get_SidewaysMaxChars() uint32
	Get_LinesToPaperCut() uint32
	Get_PageSize() Size
	Get_PrintArea() Rect
	CreateJob() *IReceiptPrintJob
}

type IClaimedReceiptPrinterVtbl struct {
	win32.IInspectableVtbl
	Get_SidewaysMaxLines uintptr
	Get_SidewaysMaxChars uintptr
	Get_LinesToPaperCut  uintptr
	Get_PageSize         uintptr
	Get_PrintArea        uintptr
	CreateJob            uintptr
}

type IClaimedReceiptPrinter struct {
	win32.IInspectable
}

func (this *IClaimedReceiptPrinter) Vtbl() *IClaimedReceiptPrinterVtbl {
	return (*IClaimedReceiptPrinterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClaimedReceiptPrinter) Get_SidewaysMaxLines() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SidewaysMaxLines, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedReceiptPrinter) Get_SidewaysMaxChars() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SidewaysMaxChars, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedReceiptPrinter) Get_LinesToPaperCut() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LinesToPaperCut, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedReceiptPrinter) Get_PageSize() Size {
	var _result Size
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PageSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedReceiptPrinter) Get_PrintArea() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrintArea, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedReceiptPrinter) CreateJob() *IReceiptPrintJob {
	var _result *IReceiptPrintJob
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateJob, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BD5DEFF2-AF90-4E8A-B77B-E3AE9CA63A7F
var IID_IClaimedSlipPrinter = syscall.GUID{0xBD5DEFF2, 0xAF90, 0x4E8A,
	[8]byte{0xB7, 0x7B, 0xE3, 0xAE, 0x9C, 0xA6, 0x3A, 0x7F}}

type IClaimedSlipPrinterInterface interface {
	win32.IInspectableInterface
	Get_SidewaysMaxLines() uint32
	Get_SidewaysMaxChars() uint32
	Get_MaxLines() uint32
	Get_LinesNearEndToEnd() uint32
	Get_PrintSide() PosPrinterPrintSide
	Get_PageSize() Size
	Get_PrintArea() Rect
	OpenJaws()
	CloseJaws()
	InsertSlipAsync(timeout TimeSpan) *IAsyncOperation[bool]
	RemoveSlipAsync(timeout TimeSpan) *IAsyncOperation[bool]
	ChangePrintSide(printSide PosPrinterPrintSide)
	CreateJob() *IReceiptOrSlipJob
}

type IClaimedSlipPrinterVtbl struct {
	win32.IInspectableVtbl
	Get_SidewaysMaxLines  uintptr
	Get_SidewaysMaxChars  uintptr
	Get_MaxLines          uintptr
	Get_LinesNearEndToEnd uintptr
	Get_PrintSide         uintptr
	Get_PageSize          uintptr
	Get_PrintArea         uintptr
	OpenJaws              uintptr
	CloseJaws             uintptr
	InsertSlipAsync       uintptr
	RemoveSlipAsync       uintptr
	ChangePrintSide       uintptr
	CreateJob             uintptr
}

type IClaimedSlipPrinter struct {
	win32.IInspectable
}

func (this *IClaimedSlipPrinter) Vtbl() *IClaimedSlipPrinterVtbl {
	return (*IClaimedSlipPrinterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClaimedSlipPrinter) Get_SidewaysMaxLines() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SidewaysMaxLines, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedSlipPrinter) Get_SidewaysMaxChars() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SidewaysMaxChars, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedSlipPrinter) Get_MaxLines() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxLines, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedSlipPrinter) Get_LinesNearEndToEnd() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LinesNearEndToEnd, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedSlipPrinter) Get_PrintSide() PosPrinterPrintSide {
	var _result PosPrinterPrintSide
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrintSide, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedSlipPrinter) Get_PageSize() Size {
	var _result Size
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PageSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedSlipPrinter) Get_PrintArea() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrintArea, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IClaimedSlipPrinter) OpenJaws() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenJaws, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IClaimedSlipPrinter) CloseJaws() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CloseJaws, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IClaimedSlipPrinter) InsertSlipAsync(timeout TimeSpan) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertSlipAsync, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&timeout)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedSlipPrinter) RemoveSlipAsync(timeout TimeSpan) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveSlipAsync, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&timeout)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IClaimedSlipPrinter) ChangePrintSide(printSide PosPrinterPrintSide) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ChangePrintSide, uintptr(unsafe.Pointer(this)), uintptr(printSide))
	_ = _hr
}

func (this *IClaimedSlipPrinter) CreateJob() *IReceiptOrSlipJob {
	var _result *IReceiptOrSlipJob
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateJob, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B7EB66A8-FE8A-4CFB-8B42-E35B280CB27C
var IID_ICommonClaimedPosPrinterStation = syscall.GUID{0xB7EB66A8, 0xFE8A, 0x4CFB,
	[8]byte{0x8B, 0x42, 0xE3, 0x5B, 0x28, 0x0C, 0xB2, 0x7C}}

type ICommonClaimedPosPrinterStationInterface interface {
	win32.IInspectableInterface
	Put_CharactersPerLine(value uint32)
	Get_CharactersPerLine() uint32
	Put_LineHeight(value uint32)
	Get_LineHeight() uint32
	Put_LineSpacing(value uint32)
	Get_LineSpacing() uint32
	Get_LineWidth() uint32
	Put_IsLetterQuality(value bool)
	Get_IsLetterQuality() bool
	Get_IsPaperNearEnd() bool
	Put_ColorCartridge(value PosPrinterColorCartridge)
	Get_ColorCartridge() PosPrinterColorCartridge
	Get_IsCoverOpen() bool
	Get_IsCartridgeRemoved() bool
	Get_IsCartridgeEmpty() bool
	Get_IsHeadCleaning() bool
	Get_IsPaperEmpty() bool
	Get_IsReadyToPrint() bool
	ValidateData(data string) bool
}

type ICommonClaimedPosPrinterStationVtbl struct {
	win32.IInspectableVtbl
	Put_CharactersPerLine  uintptr
	Get_CharactersPerLine  uintptr
	Put_LineHeight         uintptr
	Get_LineHeight         uintptr
	Put_LineSpacing        uintptr
	Get_LineSpacing        uintptr
	Get_LineWidth          uintptr
	Put_IsLetterQuality    uintptr
	Get_IsLetterQuality    uintptr
	Get_IsPaperNearEnd     uintptr
	Put_ColorCartridge     uintptr
	Get_ColorCartridge     uintptr
	Get_IsCoverOpen        uintptr
	Get_IsCartridgeRemoved uintptr
	Get_IsCartridgeEmpty   uintptr
	Get_IsHeadCleaning     uintptr
	Get_IsPaperEmpty       uintptr
	Get_IsReadyToPrint     uintptr
	ValidateData           uintptr
}

type ICommonClaimedPosPrinterStation struct {
	win32.IInspectable
}

func (this *ICommonClaimedPosPrinterStation) Vtbl() *ICommonClaimedPosPrinterStationVtbl {
	return (*ICommonClaimedPosPrinterStationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICommonClaimedPosPrinterStation) Put_CharactersPerLine(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CharactersPerLine, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICommonClaimedPosPrinterStation) Get_CharactersPerLine() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CharactersPerLine, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonClaimedPosPrinterStation) Put_LineHeight(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LineHeight, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICommonClaimedPosPrinterStation) Get_LineHeight() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LineHeight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonClaimedPosPrinterStation) Put_LineSpacing(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LineSpacing, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICommonClaimedPosPrinterStation) Get_LineSpacing() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LineSpacing, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonClaimedPosPrinterStation) Get_LineWidth() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LineWidth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonClaimedPosPrinterStation) Put_IsLetterQuality(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsLetterQuality, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICommonClaimedPosPrinterStation) Get_IsLetterQuality() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsLetterQuality, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonClaimedPosPrinterStation) Get_IsPaperNearEnd() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPaperNearEnd, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonClaimedPosPrinterStation) Put_ColorCartridge(value PosPrinterColorCartridge) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ColorCartridge, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICommonClaimedPosPrinterStation) Get_ColorCartridge() PosPrinterColorCartridge {
	var _result PosPrinterColorCartridge
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ColorCartridge, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonClaimedPosPrinterStation) Get_IsCoverOpen() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCoverOpen, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonClaimedPosPrinterStation) Get_IsCartridgeRemoved() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCartridgeRemoved, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonClaimedPosPrinterStation) Get_IsCartridgeEmpty() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCartridgeEmpty, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonClaimedPosPrinterStation) Get_IsHeadCleaning() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsHeadCleaning, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonClaimedPosPrinterStation) Get_IsPaperEmpty() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPaperEmpty, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonClaimedPosPrinterStation) Get_IsReadyToPrint() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsReadyToPrint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonClaimedPosPrinterStation) ValidateData(data string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ValidateData, uintptr(unsafe.Pointer(this)), NewHStr(data).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// DE5B52CA-E02E-40E9-9E5E-1B488E6AACFC
var IID_ICommonPosPrintStationCapabilities = syscall.GUID{0xDE5B52CA, 0xE02E, 0x40E9,
	[8]byte{0x9E, 0x5E, 0x1B, 0x48, 0x8E, 0x6A, 0xAC, 0xFC}}

type ICommonPosPrintStationCapabilitiesInterface interface {
	win32.IInspectableInterface
	Get_IsPrinterPresent() bool
	Get_IsDualColorSupported() bool
	Get_ColorCartridgeCapabilities() PosPrinterColorCapabilities
	Get_CartridgeSensors() PosPrinterCartridgeSensors
	Get_IsBoldSupported() bool
	Get_IsItalicSupported() bool
	Get_IsUnderlineSupported() bool
	Get_IsDoubleHighPrintSupported() bool
	Get_IsDoubleWidePrintSupported() bool
	Get_IsDoubleHighDoubleWidePrintSupported() bool
	Get_IsPaperEmptySensorSupported() bool
	Get_IsPaperNearEndSensorSupported() bool
	Get_SupportedCharactersPerLine() *IVectorView[uint32]
}

type ICommonPosPrintStationCapabilitiesVtbl struct {
	win32.IInspectableVtbl
	Get_IsPrinterPresent                     uintptr
	Get_IsDualColorSupported                 uintptr
	Get_ColorCartridgeCapabilities           uintptr
	Get_CartridgeSensors                     uintptr
	Get_IsBoldSupported                      uintptr
	Get_IsItalicSupported                    uintptr
	Get_IsUnderlineSupported                 uintptr
	Get_IsDoubleHighPrintSupported           uintptr
	Get_IsDoubleWidePrintSupported           uintptr
	Get_IsDoubleHighDoubleWidePrintSupported uintptr
	Get_IsPaperEmptySensorSupported          uintptr
	Get_IsPaperNearEndSensorSupported        uintptr
	Get_SupportedCharactersPerLine           uintptr
}

type ICommonPosPrintStationCapabilities struct {
	win32.IInspectable
}

func (this *ICommonPosPrintStationCapabilities) Vtbl() *ICommonPosPrintStationCapabilitiesVtbl {
	return (*ICommonPosPrintStationCapabilitiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICommonPosPrintStationCapabilities) Get_IsPrinterPresent() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPrinterPresent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonPosPrintStationCapabilities) Get_IsDualColorSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDualColorSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonPosPrintStationCapabilities) Get_ColorCartridgeCapabilities() PosPrinterColorCapabilities {
	var _result PosPrinterColorCapabilities
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ColorCartridgeCapabilities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonPosPrintStationCapabilities) Get_CartridgeSensors() PosPrinterCartridgeSensors {
	var _result PosPrinterCartridgeSensors
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CartridgeSensors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonPosPrintStationCapabilities) Get_IsBoldSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsBoldSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonPosPrintStationCapabilities) Get_IsItalicSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsItalicSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonPosPrintStationCapabilities) Get_IsUnderlineSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsUnderlineSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonPosPrintStationCapabilities) Get_IsDoubleHighPrintSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDoubleHighPrintSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonPosPrintStationCapabilities) Get_IsDoubleWidePrintSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDoubleWidePrintSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonPosPrintStationCapabilities) Get_IsDoubleHighDoubleWidePrintSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDoubleHighDoubleWidePrintSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonPosPrintStationCapabilities) Get_IsPaperEmptySensorSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPaperEmptySensorSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonPosPrintStationCapabilities) Get_IsPaperNearEndSensorSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPaperNearEndSensorSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonPosPrintStationCapabilities) Get_SupportedCharactersPerLine() *IVectorView[uint32] {
	var _result *IVectorView[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedCharactersPerLine, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 09286B8B-9873-4D05-BFBE-4727A6038F69
var IID_ICommonReceiptSlipCapabilities = syscall.GUID{0x09286B8B, 0x9873, 0x4D05,
	[8]byte{0xBF, 0xBE, 0x47, 0x27, 0xA6, 0x03, 0x8F, 0x69}}

type ICommonReceiptSlipCapabilitiesInterface interface {
	win32.IInspectableInterface
	Get_IsBarcodeSupported() bool
	Get_IsBitmapSupported() bool
	Get_IsLeft90RotationSupported() bool
	Get_IsRight90RotationSupported() bool
	Get_Is180RotationSupported() bool
	Get_IsPrintAreaSupported() bool
	Get_RuledLineCapabilities() PosPrinterRuledLineCapabilities
	Get_SupportedBarcodeRotations() *IVectorView[PosPrinterRotation]
	Get_SupportedBitmapRotations() *IVectorView[PosPrinterRotation]
}

type ICommonReceiptSlipCapabilitiesVtbl struct {
	win32.IInspectableVtbl
	Get_IsBarcodeSupported         uintptr
	Get_IsBitmapSupported          uintptr
	Get_IsLeft90RotationSupported  uintptr
	Get_IsRight90RotationSupported uintptr
	Get_Is180RotationSupported     uintptr
	Get_IsPrintAreaSupported       uintptr
	Get_RuledLineCapabilities      uintptr
	Get_SupportedBarcodeRotations  uintptr
	Get_SupportedBitmapRotations   uintptr
}

type ICommonReceiptSlipCapabilities struct {
	win32.IInspectable
}

func (this *ICommonReceiptSlipCapabilities) Vtbl() *ICommonReceiptSlipCapabilitiesVtbl {
	return (*ICommonReceiptSlipCapabilitiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICommonReceiptSlipCapabilities) Get_IsBarcodeSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsBarcodeSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonReceiptSlipCapabilities) Get_IsBitmapSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsBitmapSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonReceiptSlipCapabilities) Get_IsLeft90RotationSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsLeft90RotationSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonReceiptSlipCapabilities) Get_IsRight90RotationSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsRight90RotationSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonReceiptSlipCapabilities) Get_Is180RotationSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Is180RotationSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonReceiptSlipCapabilities) Get_IsPrintAreaSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPrintAreaSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonReceiptSlipCapabilities) Get_RuledLineCapabilities() PosPrinterRuledLineCapabilities {
	var _result PosPrinterRuledLineCapabilities
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RuledLineCapabilities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICommonReceiptSlipCapabilities) Get_SupportedBarcodeRotations() *IVectorView[PosPrinterRotation] {
	var _result *IVectorView[PosPrinterRotation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedBarcodeRotations, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICommonReceiptSlipCapabilities) Get_SupportedBitmapRotations() *IVectorView[PosPrinterRotation] {
	var _result *IVectorView[PosPrinterRotation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedBitmapRotations, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9F4F2864-F3F0-55D0-8C39-74CC91783EED
var IID_IJournalPrintJob = syscall.GUID{0x9F4F2864, 0xF3F0, 0x55D0,
	[8]byte{0x8C, 0x39, 0x74, 0xCC, 0x91, 0x78, 0x3E, 0xED}}

type IJournalPrintJobInterface interface {
	win32.IInspectableInterface
	Print(data string, printOptions *IPosPrinterPrintOptions)
	FeedPaperByLine(lineCount int32)
	FeedPaperByMapModeUnit(distance int32)
}

type IJournalPrintJobVtbl struct {
	win32.IInspectableVtbl
	Print                  uintptr
	FeedPaperByLine        uintptr
	FeedPaperByMapModeUnit uintptr
}

type IJournalPrintJob struct {
	win32.IInspectable
}

func (this *IJournalPrintJob) Vtbl() *IJournalPrintJobVtbl {
	return (*IJournalPrintJobVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IJournalPrintJob) Print(data string, printOptions *IPosPrinterPrintOptions) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Print, uintptr(unsafe.Pointer(this)), NewHStr(data).Ptr, uintptr(unsafe.Pointer(printOptions)))
	_ = _hr
}

func (this *IJournalPrintJob) FeedPaperByLine(lineCount int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FeedPaperByLine, uintptr(unsafe.Pointer(this)), uintptr(lineCount))
	_ = _hr
}

func (this *IJournalPrintJob) FeedPaperByMapModeUnit(distance int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FeedPaperByMapModeUnit, uintptr(unsafe.Pointer(this)), uintptr(distance))
	_ = _hr
}

// 3B5CCC43-E047-4463-BB58-17B5BA1D8056
var IID_IJournalPrinterCapabilities = syscall.GUID{0x3B5CCC43, 0xE047, 0x4463,
	[8]byte{0xBB, 0x58, 0x17, 0xB5, 0xBA, 0x1D, 0x80, 0x56}}

type IJournalPrinterCapabilitiesInterface interface {
	win32.IInspectableInterface
}

type IJournalPrinterCapabilitiesVtbl struct {
	win32.IInspectableVtbl
}

type IJournalPrinterCapabilities struct {
	win32.IInspectable
}

func (this *IJournalPrinterCapabilities) Vtbl() *IJournalPrinterCapabilitiesVtbl {
	return (*IJournalPrinterCapabilitiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 03B0B645-33B8-533B-BAAA-A4389283AB0A
var IID_IJournalPrinterCapabilities2 = syscall.GUID{0x03B0B645, 0x33B8, 0x533B,
	[8]byte{0xBA, 0xAA, 0xA4, 0x38, 0x92, 0x83, 0xAB, 0x0A}}

type IJournalPrinterCapabilities2Interface interface {
	win32.IInspectableInterface
	Get_IsReverseVideoSupported() bool
	Get_IsStrikethroughSupported() bool
	Get_IsSuperscriptSupported() bool
	Get_IsSubscriptSupported() bool
	Get_IsReversePaperFeedByLineSupported() bool
	Get_IsReversePaperFeedByMapModeUnitSupported() bool
}

type IJournalPrinterCapabilities2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsReverseVideoSupported                  uintptr
	Get_IsStrikethroughSupported                 uintptr
	Get_IsSuperscriptSupported                   uintptr
	Get_IsSubscriptSupported                     uintptr
	Get_IsReversePaperFeedByLineSupported        uintptr
	Get_IsReversePaperFeedByMapModeUnitSupported uintptr
}

type IJournalPrinterCapabilities2 struct {
	win32.IInspectable
}

func (this *IJournalPrinterCapabilities2) Vtbl() *IJournalPrinterCapabilities2Vtbl {
	return (*IJournalPrinterCapabilities2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IJournalPrinterCapabilities2) Get_IsReverseVideoSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsReverseVideoSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IJournalPrinterCapabilities2) Get_IsStrikethroughSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsStrikethroughSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IJournalPrinterCapabilities2) Get_IsSuperscriptSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSuperscriptSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IJournalPrinterCapabilities2) Get_IsSubscriptSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSubscriptSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IJournalPrinterCapabilities2) Get_IsReversePaperFeedByLineSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsReversePaperFeedByLineSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IJournalPrinterCapabilities2) Get_IsReversePaperFeedByMapModeUnitSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsReversePaperFeedByMapModeUnitSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 24F5DF4E-3C99-44E2-B73F-E51BE3637A8C
var IID_ILineDisplay = syscall.GUID{0x24F5DF4E, 0x3C99, 0x44E2,
	[8]byte{0xB7, 0x3F, 0xE5, 0x1B, 0xE3, 0x63, 0x7A, 0x8C}}

type ILineDisplayInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_Capabilities() *ILineDisplayCapabilities
	Get_PhysicalDeviceName() string
	Get_PhysicalDeviceDescription() string
	Get_DeviceControlDescription() string
	Get_DeviceControlVersion() string
	Get_DeviceServiceVersion() string
	ClaimAsync() *IAsyncOperation[*IClaimedLineDisplay]
}

type ILineDisplayVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId                  uintptr
	Get_Capabilities              uintptr
	Get_PhysicalDeviceName        uintptr
	Get_PhysicalDeviceDescription uintptr
	Get_DeviceControlDescription  uintptr
	Get_DeviceControlVersion      uintptr
	Get_DeviceServiceVersion      uintptr
	ClaimAsync                    uintptr
}

type ILineDisplay struct {
	win32.IInspectable
}

func (this *ILineDisplay) Vtbl() *ILineDisplayVtbl {
	return (*ILineDisplayVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILineDisplay) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILineDisplay) Get_Capabilities() *ILineDisplayCapabilities {
	var _result *ILineDisplayCapabilities
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Capabilities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILineDisplay) Get_PhysicalDeviceName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhysicalDeviceName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILineDisplay) Get_PhysicalDeviceDescription() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhysicalDeviceDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILineDisplay) Get_DeviceControlDescription() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceControlDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILineDisplay) Get_DeviceControlVersion() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceControlVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILineDisplay) Get_DeviceServiceVersion() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceServiceVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILineDisplay) ClaimAsync() *IAsyncOperation[*IClaimedLineDisplay] {
	var _result *IAsyncOperation[*IClaimedLineDisplay]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClaimAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C296A628-EF44-40F3-BD1C-B04C6A5CDC7D
var IID_ILineDisplay2 = syscall.GUID{0xC296A628, 0xEF44, 0x40F3,
	[8]byte{0xBD, 0x1C, 0xB0, 0x4C, 0x6A, 0x5C, 0xDC, 0x7D}}

type ILineDisplay2Interface interface {
	win32.IInspectableInterface
	CheckPowerStatusAsync() *IAsyncOperation[LineDisplayPowerStatus]
}

type ILineDisplay2Vtbl struct {
	win32.IInspectableVtbl
	CheckPowerStatusAsync uintptr
}

type ILineDisplay2 struct {
	win32.IInspectable
}

func (this *ILineDisplay2) Vtbl() *ILineDisplay2Vtbl {
	return (*ILineDisplay2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILineDisplay2) CheckPowerStatusAsync() *IAsyncOperation[LineDisplayPowerStatus] {
	var _result *IAsyncOperation[LineDisplayPowerStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CheckPowerStatusAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C17DE99C-229A-4C14-A6F1-B4E4B1FEAD92
var IID_ILineDisplayAttributes = syscall.GUID{0xC17DE99C, 0x229A, 0x4C14,
	[8]byte{0xA6, 0xF1, 0xB4, 0xE4, 0xB1, 0xFE, 0xAD, 0x92}}

type ILineDisplayAttributesInterface interface {
	win32.IInspectableInterface
	Get_IsPowerNotifyEnabled() bool
	Put_IsPowerNotifyEnabled(value bool)
	Get_Brightness() int32
	Put_Brightness(value int32)
	Get_BlinkRate() TimeSpan
	Put_BlinkRate(value TimeSpan)
	Get_ScreenSizeInCharacters() Size
	Put_ScreenSizeInCharacters(value Size)
	Get_CharacterSet() int32
	Put_CharacterSet(value int32)
	Get_IsCharacterSetMappingEnabled() bool
	Put_IsCharacterSetMappingEnabled(value bool)
	Get_CurrentWindow() *ILineDisplayWindow
	Put_CurrentWindow(value *ILineDisplayWindow)
}

type ILineDisplayAttributesVtbl struct {
	win32.IInspectableVtbl
	Get_IsPowerNotifyEnabled         uintptr
	Put_IsPowerNotifyEnabled         uintptr
	Get_Brightness                   uintptr
	Put_Brightness                   uintptr
	Get_BlinkRate                    uintptr
	Put_BlinkRate                    uintptr
	Get_ScreenSizeInCharacters       uintptr
	Put_ScreenSizeInCharacters       uintptr
	Get_CharacterSet                 uintptr
	Put_CharacterSet                 uintptr
	Get_IsCharacterSetMappingEnabled uintptr
	Put_IsCharacterSetMappingEnabled uintptr
	Get_CurrentWindow                uintptr
	Put_CurrentWindow                uintptr
}

type ILineDisplayAttributes struct {
	win32.IInspectable
}

func (this *ILineDisplayAttributes) Vtbl() *ILineDisplayAttributesVtbl {
	return (*ILineDisplayAttributesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILineDisplayAttributes) Get_IsPowerNotifyEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPowerNotifyEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayAttributes) Put_IsPowerNotifyEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsPowerNotifyEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ILineDisplayAttributes) Get_Brightness() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Brightness, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayAttributes) Put_Brightness(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Brightness, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ILineDisplayAttributes) Get_BlinkRate() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BlinkRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayAttributes) Put_BlinkRate(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BlinkRate, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ILineDisplayAttributes) Get_ScreenSizeInCharacters() Size {
	var _result Size
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScreenSizeInCharacters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayAttributes) Put_ScreenSizeInCharacters(value Size) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ScreenSizeInCharacters, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ILineDisplayAttributes) Get_CharacterSet() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CharacterSet, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayAttributes) Put_CharacterSet(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CharacterSet, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ILineDisplayAttributes) Get_IsCharacterSetMappingEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCharacterSetMappingEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayAttributes) Put_IsCharacterSetMappingEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsCharacterSetMappingEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ILineDisplayAttributes) Get_CurrentWindow() *ILineDisplayWindow {
	var _result *ILineDisplayWindow
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentWindow, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILineDisplayAttributes) Put_CurrentWindow(value *ILineDisplayWindow) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CurrentWindow, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 5A15B5D1-8DC5-4B9C-9172-303E47B70C55
var IID_ILineDisplayCapabilities = syscall.GUID{0x5A15B5D1, 0x8DC5, 0x4B9C,
	[8]byte{0x91, 0x72, 0x30, 0x3E, 0x47, 0xB7, 0x0C, 0x55}}

type ILineDisplayCapabilitiesInterface interface {
	win32.IInspectableInterface
	Get_IsStatisticsReportingSupported() bool
	Get_IsStatisticsUpdatingSupported() bool
	Get_PowerReportingType() UnifiedPosPowerReportingType
	Get_CanChangeScreenSize() bool
	Get_CanDisplayBitmaps() bool
	Get_CanReadCharacterAtCursor() bool
	Get_CanMapCharacterSets() bool
	Get_CanDisplayCustomGlyphs() bool
	Get_CanReverse() LineDisplayTextAttributeGranularity
	Get_CanBlink() LineDisplayTextAttributeGranularity
	Get_CanChangeBlinkRate() bool
	Get_IsBrightnessSupported() bool
	Get_IsCursorSupported() bool
	Get_IsHorizontalMarqueeSupported() bool
	Get_IsVerticalMarqueeSupported() bool
	Get_IsInterCharacterWaitSupported() bool
	Get_SupportedDescriptors() uint32
	Get_SupportedWindows() uint32
}

type ILineDisplayCapabilitiesVtbl struct {
	win32.IInspectableVtbl
	Get_IsStatisticsReportingSupported uintptr
	Get_IsStatisticsUpdatingSupported  uintptr
	Get_PowerReportingType             uintptr
	Get_CanChangeScreenSize            uintptr
	Get_CanDisplayBitmaps              uintptr
	Get_CanReadCharacterAtCursor       uintptr
	Get_CanMapCharacterSets            uintptr
	Get_CanDisplayCustomGlyphs         uintptr
	Get_CanReverse                     uintptr
	Get_CanBlink                       uintptr
	Get_CanChangeBlinkRate             uintptr
	Get_IsBrightnessSupported          uintptr
	Get_IsCursorSupported              uintptr
	Get_IsHorizontalMarqueeSupported   uintptr
	Get_IsVerticalMarqueeSupported     uintptr
	Get_IsInterCharacterWaitSupported  uintptr
	Get_SupportedDescriptors           uintptr
	Get_SupportedWindows               uintptr
}

type ILineDisplayCapabilities struct {
	win32.IInspectable
}

func (this *ILineDisplayCapabilities) Vtbl() *ILineDisplayCapabilitiesVtbl {
	return (*ILineDisplayCapabilitiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILineDisplayCapabilities) Get_IsStatisticsReportingSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsStatisticsReportingSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCapabilities) Get_IsStatisticsUpdatingSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsStatisticsUpdatingSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCapabilities) Get_PowerReportingType() UnifiedPosPowerReportingType {
	var _result UnifiedPosPowerReportingType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PowerReportingType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCapabilities) Get_CanChangeScreenSize() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanChangeScreenSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCapabilities) Get_CanDisplayBitmaps() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanDisplayBitmaps, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCapabilities) Get_CanReadCharacterAtCursor() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanReadCharacterAtCursor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCapabilities) Get_CanMapCharacterSets() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanMapCharacterSets, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCapabilities) Get_CanDisplayCustomGlyphs() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanDisplayCustomGlyphs, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCapabilities) Get_CanReverse() LineDisplayTextAttributeGranularity {
	var _result LineDisplayTextAttributeGranularity
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanReverse, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCapabilities) Get_CanBlink() LineDisplayTextAttributeGranularity {
	var _result LineDisplayTextAttributeGranularity
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanBlink, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCapabilities) Get_CanChangeBlinkRate() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanChangeBlinkRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCapabilities) Get_IsBrightnessSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsBrightnessSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCapabilities) Get_IsCursorSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCursorSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCapabilities) Get_IsHorizontalMarqueeSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsHorizontalMarqueeSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCapabilities) Get_IsVerticalMarqueeSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsVerticalMarqueeSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCapabilities) Get_IsInterCharacterWaitSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsInterCharacterWaitSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCapabilities) Get_SupportedDescriptors() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedDescriptors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCapabilities) Get_SupportedWindows() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedWindows, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// ECDFFC45-754A-4E3B-AB2B-151181085605
var IID_ILineDisplayCursor = syscall.GUID{0xECDFFC45, 0x754A, 0x4E3B,
	[8]byte{0xAB, 0x2B, 0x15, 0x11, 0x81, 0x08, 0x56, 0x05}}

type ILineDisplayCursorInterface interface {
	win32.IInspectableInterface
	Get_CanCustomize() bool
	Get_IsBlinkSupported() bool
	Get_IsBlockSupported() bool
	Get_IsHalfBlockSupported() bool
	Get_IsUnderlineSupported() bool
	Get_IsReverseSupported() bool
	Get_IsOtherSupported() bool
	GetAttributes() *ILineDisplayCursorAttributes
	TryUpdateAttributesAsync(attributes *ILineDisplayCursorAttributes) *IAsyncOperation[bool]
}

type ILineDisplayCursorVtbl struct {
	win32.IInspectableVtbl
	Get_CanCustomize         uintptr
	Get_IsBlinkSupported     uintptr
	Get_IsBlockSupported     uintptr
	Get_IsHalfBlockSupported uintptr
	Get_IsUnderlineSupported uintptr
	Get_IsReverseSupported   uintptr
	Get_IsOtherSupported     uintptr
	GetAttributes            uintptr
	TryUpdateAttributesAsync uintptr
}

type ILineDisplayCursor struct {
	win32.IInspectable
}

func (this *ILineDisplayCursor) Vtbl() *ILineDisplayCursorVtbl {
	return (*ILineDisplayCursorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILineDisplayCursor) Get_CanCustomize() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanCustomize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCursor) Get_IsBlinkSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsBlinkSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCursor) Get_IsBlockSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsBlockSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCursor) Get_IsHalfBlockSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsHalfBlockSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCursor) Get_IsUnderlineSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsUnderlineSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCursor) Get_IsReverseSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsReverseSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCursor) Get_IsOtherSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsOtherSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCursor) GetAttributes() *ILineDisplayCursorAttributes {
	var _result *ILineDisplayCursorAttributes
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAttributes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILineDisplayCursor) TryUpdateAttributesAsync(attributes *ILineDisplayCursorAttributes) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryUpdateAttributesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(attributes)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4E2D54FE-4FFD-4190-AAE1-CE285F20C896
var IID_ILineDisplayCursorAttributes = syscall.GUID{0x4E2D54FE, 0x4FFD, 0x4190,
	[8]byte{0xAA, 0xE1, 0xCE, 0x28, 0x5F, 0x20, 0xC8, 0x96}}

type ILineDisplayCursorAttributesInterface interface {
	win32.IInspectableInterface
	Get_IsBlinkEnabled() bool
	Put_IsBlinkEnabled(value bool)
	Get_CursorType() LineDisplayCursorType
	Put_CursorType(value LineDisplayCursorType)
	Get_IsAutoAdvanceEnabled() bool
	Put_IsAutoAdvanceEnabled(value bool)
	Get_Position() Point
	Put_Position(value Point)
}

type ILineDisplayCursorAttributesVtbl struct {
	win32.IInspectableVtbl
	Get_IsBlinkEnabled       uintptr
	Put_IsBlinkEnabled       uintptr
	Get_CursorType           uintptr
	Put_CursorType           uintptr
	Get_IsAutoAdvanceEnabled uintptr
	Put_IsAutoAdvanceEnabled uintptr
	Get_Position             uintptr
	Put_Position             uintptr
}

type ILineDisplayCursorAttributes struct {
	win32.IInspectable
}

func (this *ILineDisplayCursorAttributes) Vtbl() *ILineDisplayCursorAttributesVtbl {
	return (*ILineDisplayCursorAttributesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILineDisplayCursorAttributes) Get_IsBlinkEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsBlinkEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCursorAttributes) Put_IsBlinkEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsBlinkEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ILineDisplayCursorAttributes) Get_CursorType() LineDisplayCursorType {
	var _result LineDisplayCursorType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CursorType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCursorAttributes) Put_CursorType(value LineDisplayCursorType) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CursorType, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ILineDisplayCursorAttributes) Get_IsAutoAdvanceEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAutoAdvanceEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCursorAttributes) Put_IsAutoAdvanceEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsAutoAdvanceEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ILineDisplayCursorAttributes) Get_Position() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCursorAttributes) Put_Position(value Point) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Position, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

// 2257F63C-F263-44F1-A1A0-E750A6A0EC54
var IID_ILineDisplayCustomGlyphs = syscall.GUID{0x2257F63C, 0xF263, 0x44F1,
	[8]byte{0xA1, 0xA0, 0xE7, 0x50, 0xA6, 0xA0, 0xEC, 0x54}}

type ILineDisplayCustomGlyphsInterface interface {
	win32.IInspectableInterface
	Get_SizeInPixels() Size
	Get_SupportedGlyphCodes() *IVectorView[uint32]
	TryRedefineAsync(glyphCode uint32, glyphData *IBuffer) *IAsyncOperation[bool]
}

type ILineDisplayCustomGlyphsVtbl struct {
	win32.IInspectableVtbl
	Get_SizeInPixels        uintptr
	Get_SupportedGlyphCodes uintptr
	TryRedefineAsync        uintptr
}

type ILineDisplayCustomGlyphs struct {
	win32.IInspectable
}

func (this *ILineDisplayCustomGlyphs) Vtbl() *ILineDisplayCustomGlyphsVtbl {
	return (*ILineDisplayCustomGlyphsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILineDisplayCustomGlyphs) Get_SizeInPixels() Size {
	var _result Size
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SizeInPixels, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayCustomGlyphs) Get_SupportedGlyphCodes() *IVectorView[uint32] {
	var _result *IVectorView[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedGlyphCodes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILineDisplayCustomGlyphs) TryRedefineAsync(glyphCode uint32, glyphData *IBuffer) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryRedefineAsync, uintptr(unsafe.Pointer(this)), uintptr(glyphCode), uintptr(unsafe.Pointer(glyphData)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A3D33E3E-F46A-4B7A-BC21-53EB3B57F8B4
var IID_ILineDisplayMarquee = syscall.GUID{0xA3D33E3E, 0xF46A, 0x4B7A,
	[8]byte{0xBC, 0x21, 0x53, 0xEB, 0x3B, 0x57, 0xF8, 0xB4}}

type ILineDisplayMarqueeInterface interface {
	win32.IInspectableInterface
	Get_Format() LineDisplayMarqueeFormat
	Put_Format(value LineDisplayMarqueeFormat)
	Get_RepeatWaitInterval() TimeSpan
	Put_RepeatWaitInterval(value TimeSpan)
	Get_ScrollWaitInterval() TimeSpan
	Put_ScrollWaitInterval(value TimeSpan)
	TryStartScrollingAsync(direction LineDisplayScrollDirection) *IAsyncOperation[bool]
	TryStopScrollingAsync() *IAsyncOperation[bool]
}

type ILineDisplayMarqueeVtbl struct {
	win32.IInspectableVtbl
	Get_Format             uintptr
	Put_Format             uintptr
	Get_RepeatWaitInterval uintptr
	Put_RepeatWaitInterval uintptr
	Get_ScrollWaitInterval uintptr
	Put_ScrollWaitInterval uintptr
	TryStartScrollingAsync uintptr
	TryStopScrollingAsync  uintptr
}

type ILineDisplayMarquee struct {
	win32.IInspectable
}

func (this *ILineDisplayMarquee) Vtbl() *ILineDisplayMarqueeVtbl {
	return (*ILineDisplayMarqueeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILineDisplayMarquee) Get_Format() LineDisplayMarqueeFormat {
	var _result LineDisplayMarqueeFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Format, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayMarquee) Put_Format(value LineDisplayMarqueeFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Format, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ILineDisplayMarquee) Get_RepeatWaitInterval() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RepeatWaitInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayMarquee) Put_RepeatWaitInterval(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RepeatWaitInterval, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ILineDisplayMarquee) Get_ScrollWaitInterval() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScrollWaitInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayMarquee) Put_ScrollWaitInterval(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ScrollWaitInterval, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ILineDisplayMarquee) TryStartScrollingAsync(direction LineDisplayScrollDirection) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryStartScrollingAsync, uintptr(unsafe.Pointer(this)), uintptr(direction), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILineDisplayMarquee) TryStopScrollingAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryStopScrollingAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 022DC0B6-11B0-4690-9547-0B39C5AF2114
var IID_ILineDisplayStatics = syscall.GUID{0x022DC0B6, 0x11B0, 0x4690,
	[8]byte{0x95, 0x47, 0x0B, 0x39, 0xC5, 0xAF, 0x21, 0x14}}

type ILineDisplayStaticsInterface interface {
	win32.IInspectableInterface
	FromIdAsync(deviceId string) *IAsyncOperation[*ILineDisplay]
	GetDefaultAsync() *IAsyncOperation[*ILineDisplay]
	GetDeviceSelector() string
	GetDeviceSelectorWithConnectionTypes(connectionTypes PosConnectionTypes) string
}

type ILineDisplayStaticsVtbl struct {
	win32.IInspectableVtbl
	FromIdAsync                          uintptr
	GetDefaultAsync                      uintptr
	GetDeviceSelector                    uintptr
	GetDeviceSelectorWithConnectionTypes uintptr
}

type ILineDisplayStatics struct {
	win32.IInspectable
}

func (this *ILineDisplayStatics) Vtbl() *ILineDisplayStaticsVtbl {
	return (*ILineDisplayStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILineDisplayStatics) FromIdAsync(deviceId string) *IAsyncOperation[*ILineDisplay] {
	var _result *IAsyncOperation[*ILineDisplay]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILineDisplayStatics) GetDefaultAsync() *IAsyncOperation[*ILineDisplay] {
	var _result *IAsyncOperation[*ILineDisplay]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILineDisplayStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILineDisplayStatics) GetDeviceSelectorWithConnectionTypes(connectionTypes PosConnectionTypes) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorWithConnectionTypes, uintptr(unsafe.Pointer(this)), uintptr(connectionTypes), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 600C3F1C-77AB-4968-A7DE-C02FF169F2CC
var IID_ILineDisplayStatics2 = syscall.GUID{0x600C3F1C, 0x77AB, 0x4968,
	[8]byte{0xA7, 0xDE, 0xC0, 0x2F, 0xF1, 0x69, 0xF2, 0xCC}}

type ILineDisplayStatics2Interface interface {
	win32.IInspectableInterface
	Get_StatisticsCategorySelector() *ILineDisplayStatisticsCategorySelector
}

type ILineDisplayStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_StatisticsCategorySelector uintptr
}

type ILineDisplayStatics2 struct {
	win32.IInspectable
}

func (this *ILineDisplayStatics2) Vtbl() *ILineDisplayStatics2Vtbl {
	return (*ILineDisplayStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILineDisplayStatics2) Get_StatisticsCategorySelector() *ILineDisplayStatisticsCategorySelector {
	var _result *ILineDisplayStatisticsCategorySelector
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StatisticsCategorySelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B521C46B-9274-4D24-94F3-B6017B832444
var IID_ILineDisplayStatisticsCategorySelector = syscall.GUID{0xB521C46B, 0x9274, 0x4D24,
	[8]byte{0x94, 0xF3, 0xB6, 0x01, 0x7B, 0x83, 0x24, 0x44}}

type ILineDisplayStatisticsCategorySelectorInterface interface {
	win32.IInspectableInterface
	Get_AllStatistics() string
	Get_UnifiedPosStatistics() string
	Get_ManufacturerStatistics() string
}

type ILineDisplayStatisticsCategorySelectorVtbl struct {
	win32.IInspectableVtbl
	Get_AllStatistics          uintptr
	Get_UnifiedPosStatistics   uintptr
	Get_ManufacturerStatistics uintptr
}

type ILineDisplayStatisticsCategorySelector struct {
	win32.IInspectable
}

func (this *ILineDisplayStatisticsCategorySelector) Vtbl() *ILineDisplayStatisticsCategorySelectorVtbl {
	return (*ILineDisplayStatisticsCategorySelectorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILineDisplayStatisticsCategorySelector) Get_AllStatistics() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllStatistics, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILineDisplayStatisticsCategorySelector) Get_UnifiedPosStatistics() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UnifiedPosStatistics, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILineDisplayStatisticsCategorySelector) Get_ManufacturerStatistics() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ManufacturerStatistics, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// DDD57C1A-86FB-4EBA-93D1-6F5EDA52B752
var IID_ILineDisplayStatusUpdatedEventArgs = syscall.GUID{0xDDD57C1A, 0x86FB, 0x4EBA,
	[8]byte{0x93, 0xD1, 0x6F, 0x5E, 0xDA, 0x52, 0xB7, 0x52}}

type ILineDisplayStatusUpdatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Status() LineDisplayPowerStatus
}

type ILineDisplayStatusUpdatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
}

type ILineDisplayStatusUpdatedEventArgs struct {
	win32.IInspectable
}

func (this *ILineDisplayStatusUpdatedEventArgs) Vtbl() *ILineDisplayStatusUpdatedEventArgsVtbl {
	return (*ILineDisplayStatusUpdatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILineDisplayStatusUpdatedEventArgs) Get_Status() LineDisplayPowerStatus {
	var _result LineDisplayPowerStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F621515B-D81E-43BA-BF1B-BCFA3C785BA0
var IID_ILineDisplayStoredBitmap = syscall.GUID{0xF621515B, 0xD81E, 0x43BA,
	[8]byte{0xBF, 0x1B, 0xBC, 0xFA, 0x3C, 0x78, 0x5B, 0xA0}}

type ILineDisplayStoredBitmapInterface interface {
	win32.IInspectableInterface
	Get_EscapeSequence() string
	TryDeleteAsync() *IAsyncOperation[bool]
}

type ILineDisplayStoredBitmapVtbl struct {
	win32.IInspectableVtbl
	Get_EscapeSequence uintptr
	TryDeleteAsync     uintptr
}

type ILineDisplayStoredBitmap struct {
	win32.IInspectable
}

func (this *ILineDisplayStoredBitmap) Vtbl() *ILineDisplayStoredBitmapVtbl {
	return (*ILineDisplayStoredBitmapVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILineDisplayStoredBitmap) Get_EscapeSequence() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EscapeSequence, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILineDisplayStoredBitmap) TryDeleteAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryDeleteAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D21FEEF4-2364-4BE5-BEE1-851680AF4964
var IID_ILineDisplayWindow = syscall.GUID{0xD21FEEF4, 0x2364, 0x4BE5,
	[8]byte{0xBE, 0xE1, 0x85, 0x16, 0x80, 0xAF, 0x49, 0x64}}

type ILineDisplayWindowInterface interface {
	win32.IInspectableInterface
	Get_SizeInCharacters() Size
	Get_InterCharacterWaitInterval() TimeSpan
	Put_InterCharacterWaitInterval(value TimeSpan)
	TryRefreshAsync() *IAsyncOperation[bool]
	TryDisplayTextAsync(text string, displayAttribute LineDisplayTextAttribute) *IAsyncOperation[bool]
	TryDisplayTextAtPositionAsync(text string, displayAttribute LineDisplayTextAttribute, startPosition Point) *IAsyncOperation[bool]
	TryDisplayTextNormalAsync(text string) *IAsyncOperation[bool]
	TryScrollTextAsync(direction LineDisplayScrollDirection, numberOfColumnsOrRows uint32) *IAsyncOperation[bool]
	TryClearTextAsync() *IAsyncOperation[bool]
}

type ILineDisplayWindowVtbl struct {
	win32.IInspectableVtbl
	Get_SizeInCharacters           uintptr
	Get_InterCharacterWaitInterval uintptr
	Put_InterCharacterWaitInterval uintptr
	TryRefreshAsync                uintptr
	TryDisplayTextAsync            uintptr
	TryDisplayTextAtPositionAsync  uintptr
	TryDisplayTextNormalAsync      uintptr
	TryScrollTextAsync             uintptr
	TryClearTextAsync              uintptr
}

type ILineDisplayWindow struct {
	win32.IInspectable
}

func (this *ILineDisplayWindow) Vtbl() *ILineDisplayWindowVtbl {
	return (*ILineDisplayWindowVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILineDisplayWindow) Get_SizeInCharacters() Size {
	var _result Size
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SizeInCharacters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayWindow) Get_InterCharacterWaitInterval() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InterCharacterWaitInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILineDisplayWindow) Put_InterCharacterWaitInterval(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InterCharacterWaitInterval, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ILineDisplayWindow) TryRefreshAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryRefreshAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILineDisplayWindow) TryDisplayTextAsync(text string, displayAttribute LineDisplayTextAttribute) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryDisplayTextAsync, uintptr(unsafe.Pointer(this)), NewHStr(text).Ptr, uintptr(displayAttribute), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILineDisplayWindow) TryDisplayTextAtPositionAsync(text string, displayAttribute LineDisplayTextAttribute, startPosition Point) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryDisplayTextAtPositionAsync, uintptr(unsafe.Pointer(this)), NewHStr(text).Ptr, uintptr(displayAttribute), *(*uintptr)(unsafe.Pointer(&startPosition)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILineDisplayWindow) TryDisplayTextNormalAsync(text string) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryDisplayTextNormalAsync, uintptr(unsafe.Pointer(this)), NewHStr(text).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILineDisplayWindow) TryScrollTextAsync(direction LineDisplayScrollDirection, numberOfColumnsOrRows uint32) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryScrollTextAsync, uintptr(unsafe.Pointer(this)), uintptr(direction), uintptr(numberOfColumnsOrRows), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILineDisplayWindow) TryClearTextAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryClearTextAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A95CE2E6-BDD8-4365-8E11-DE94DE8DFF02
var IID_ILineDisplayWindow2 = syscall.GUID{0xA95CE2E6, 0xBDD8, 0x4365,
	[8]byte{0x8E, 0x11, 0xDE, 0x94, 0xDE, 0x8D, 0xFF, 0x02}}

type ILineDisplayWindow2Interface interface {
	win32.IInspectableInterface
	Get_Cursor() *ILineDisplayCursor
	Get_Marquee() *ILineDisplayMarquee
	ReadCharacterAtCursorAsync() *IAsyncOperation[uint32]
	TryDisplayStoredBitmapAtCursorAsync(bitmap *ILineDisplayStoredBitmap) *IAsyncOperation[bool]
	TryDisplayStorageFileBitmapAtCursorAsync(bitmap *IStorageFile) *IAsyncOperation[bool]
	TryDisplayStorageFileBitmapAtCursorWithAlignmentAsync(bitmap *IStorageFile, horizontalAlignment LineDisplayHorizontalAlignment, verticalAlignment LineDisplayVerticalAlignment) *IAsyncOperation[bool]
	TryDisplayStorageFileBitmapAtCursorWithAlignmentAndWidthAsync(bitmap *IStorageFile, horizontalAlignment LineDisplayHorizontalAlignment, verticalAlignment LineDisplayVerticalAlignment, widthInPixels int32) *IAsyncOperation[bool]
	TryDisplayStorageFileBitmapAtPointAsync(bitmap *IStorageFile, offsetInPixels Point) *IAsyncOperation[bool]
	TryDisplayStorageFileBitmapAtPointWithWidthAsync(bitmap *IStorageFile, offsetInPixels Point, widthInPixels int32) *IAsyncOperation[bool]
}

type ILineDisplayWindow2Vtbl struct {
	win32.IInspectableVtbl
	Get_Cursor                                                    uintptr
	Get_Marquee                                                   uintptr
	ReadCharacterAtCursorAsync                                    uintptr
	TryDisplayStoredBitmapAtCursorAsync                           uintptr
	TryDisplayStorageFileBitmapAtCursorAsync                      uintptr
	TryDisplayStorageFileBitmapAtCursorWithAlignmentAsync         uintptr
	TryDisplayStorageFileBitmapAtCursorWithAlignmentAndWidthAsync uintptr
	TryDisplayStorageFileBitmapAtPointAsync                       uintptr
	TryDisplayStorageFileBitmapAtPointWithWidthAsync              uintptr
}

type ILineDisplayWindow2 struct {
	win32.IInspectable
}

func (this *ILineDisplayWindow2) Vtbl() *ILineDisplayWindow2Vtbl {
	return (*ILineDisplayWindow2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILineDisplayWindow2) Get_Cursor() *ILineDisplayCursor {
	var _result *ILineDisplayCursor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Cursor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILineDisplayWindow2) Get_Marquee() *ILineDisplayMarquee {
	var _result *ILineDisplayMarquee
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Marquee, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILineDisplayWindow2) ReadCharacterAtCursorAsync() *IAsyncOperation[uint32] {
	var _result *IAsyncOperation[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadCharacterAtCursorAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILineDisplayWindow2) TryDisplayStoredBitmapAtCursorAsync(bitmap *ILineDisplayStoredBitmap) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryDisplayStoredBitmapAtCursorAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bitmap)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILineDisplayWindow2) TryDisplayStorageFileBitmapAtCursorAsync(bitmap *IStorageFile) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryDisplayStorageFileBitmapAtCursorAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bitmap)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILineDisplayWindow2) TryDisplayStorageFileBitmapAtCursorWithAlignmentAsync(bitmap *IStorageFile, horizontalAlignment LineDisplayHorizontalAlignment, verticalAlignment LineDisplayVerticalAlignment) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryDisplayStorageFileBitmapAtCursorWithAlignmentAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bitmap)), uintptr(horizontalAlignment), uintptr(verticalAlignment), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILineDisplayWindow2) TryDisplayStorageFileBitmapAtCursorWithAlignmentAndWidthAsync(bitmap *IStorageFile, horizontalAlignment LineDisplayHorizontalAlignment, verticalAlignment LineDisplayVerticalAlignment, widthInPixels int32) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryDisplayStorageFileBitmapAtCursorWithAlignmentAndWidthAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bitmap)), uintptr(horizontalAlignment), uintptr(verticalAlignment), uintptr(widthInPixels), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILineDisplayWindow2) TryDisplayStorageFileBitmapAtPointAsync(bitmap *IStorageFile, offsetInPixels Point) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryDisplayStorageFileBitmapAtPointAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bitmap)), *(*uintptr)(unsafe.Pointer(&offsetInPixels)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILineDisplayWindow2) TryDisplayStorageFileBitmapAtPointWithWidthAsync(bitmap *IStorageFile, offsetInPixels Point, widthInPixels int32) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryDisplayStorageFileBitmapAtPointWithWidthAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bitmap)), *(*uintptr)(unsafe.Pointer(&offsetInPixels)), uintptr(widthInPixels), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1A92B015-47C3-468A-9333-0C6517574883
var IID_IMagneticStripeReader = syscall.GUID{0x1A92B015, 0x47C3, 0x468A,
	[8]byte{0x93, 0x33, 0x0C, 0x65, 0x17, 0x57, 0x48, 0x83}}

type IMagneticStripeReaderInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_Capabilities() *IMagneticStripeReaderCapabilities
	Get_SupportedCardTypes() []uint32
	Get_DeviceAuthenticationProtocol() MagneticStripeReaderAuthenticationProtocol
	CheckHealthAsync(level UnifiedPosHealthCheckLevel) *IAsyncOperation[string]
	ClaimReaderAsync() *IAsyncOperation[*IClaimedMagneticStripeReader]
	RetrieveStatisticsAsync(statisticsCategories *IIterable[string]) *IAsyncOperation[*IBuffer]
	GetErrorReportingType() MagneticStripeReaderErrorReportingType
	Add_StatusUpdated(handler TypedEventHandler[*IMagneticStripeReader, *IMagneticStripeReaderStatusUpdatedEventArgs]) EventRegistrationToken
	Remove_StatusUpdated(token EventRegistrationToken)
}

type IMagneticStripeReaderVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId                     uintptr
	Get_Capabilities                 uintptr
	Get_SupportedCardTypes           uintptr
	Get_DeviceAuthenticationProtocol uintptr
	CheckHealthAsync                 uintptr
	ClaimReaderAsync                 uintptr
	RetrieveStatisticsAsync          uintptr
	GetErrorReportingType            uintptr
	Add_StatusUpdated                uintptr
	Remove_StatusUpdated             uintptr
}

type IMagneticStripeReader struct {
	win32.IInspectable
}

func (this *IMagneticStripeReader) Vtbl() *IMagneticStripeReaderVtbl {
	return (*IMagneticStripeReaderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMagneticStripeReader) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReader) Get_Capabilities() *IMagneticStripeReaderCapabilities {
	var _result *IMagneticStripeReaderCapabilities
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Capabilities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMagneticStripeReader) Get_SupportedCardTypes() []uint32 {
	var _result []uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedCardTypes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagneticStripeReader) Get_DeviceAuthenticationProtocol() MagneticStripeReaderAuthenticationProtocol {
	var _result MagneticStripeReaderAuthenticationProtocol
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceAuthenticationProtocol, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagneticStripeReader) CheckHealthAsync(level UnifiedPosHealthCheckLevel) *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CheckHealthAsync, uintptr(unsafe.Pointer(this)), uintptr(level), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMagneticStripeReader) ClaimReaderAsync() *IAsyncOperation[*IClaimedMagneticStripeReader] {
	var _result *IAsyncOperation[*IClaimedMagneticStripeReader]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClaimReaderAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMagneticStripeReader) RetrieveStatisticsAsync(statisticsCategories *IIterable[string]) *IAsyncOperation[*IBuffer] {
	var _result *IAsyncOperation[*IBuffer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RetrieveStatisticsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(statisticsCategories)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMagneticStripeReader) GetErrorReportingType() MagneticStripeReaderErrorReportingType {
	var _result MagneticStripeReaderErrorReportingType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetErrorReportingType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagneticStripeReader) Add_StatusUpdated(handler TypedEventHandler[*IMagneticStripeReader, *IMagneticStripeReaderStatusUpdatedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StatusUpdated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagneticStripeReader) Remove_StatusUpdated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StatusUpdated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 0A4BBD51-C316-4910-87F3-7A62BA862D31
var IID_IMagneticStripeReaderAamvaCardDataReceivedEventArgs = syscall.GUID{0x0A4BBD51, 0xC316, 0x4910,
	[8]byte{0x87, 0xF3, 0x7A, 0x62, 0xBA, 0x86, 0x2D, 0x31}}

type IMagneticStripeReaderAamvaCardDataReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Report() *IMagneticStripeReaderReport
	Get_LicenseNumber() string
	Get_ExpirationDate() string
	Get_Restrictions() string
	Get_Class() string
	Get_Endorsements() string
	Get_BirthDate() string
	Get_FirstName() string
	Get_Surname() string
	Get_Suffix() string
	Get_Gender() string
	Get_HairColor() string
	Get_EyeColor() string
	Get_Height() string
	Get_Weight() string
	Get_Address() string
	Get_City() string
	Get_State() string
	Get_PostalCode() string
}

type IMagneticStripeReaderAamvaCardDataReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Report         uintptr
	Get_LicenseNumber  uintptr
	Get_ExpirationDate uintptr
	Get_Restrictions   uintptr
	Get_Class          uintptr
	Get_Endorsements   uintptr
	Get_BirthDate      uintptr
	Get_FirstName      uintptr
	Get_Surname        uintptr
	Get_Suffix         uintptr
	Get_Gender         uintptr
	Get_HairColor      uintptr
	Get_EyeColor       uintptr
	Get_Height         uintptr
	Get_Weight         uintptr
	Get_Address        uintptr
	Get_City           uintptr
	Get_State          uintptr
	Get_PostalCode     uintptr
}

type IMagneticStripeReaderAamvaCardDataReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IMagneticStripeReaderAamvaCardDataReceivedEventArgs) Vtbl() *IMagneticStripeReaderAamvaCardDataReceivedEventArgsVtbl {
	return (*IMagneticStripeReaderAamvaCardDataReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMagneticStripeReaderAamvaCardDataReceivedEventArgs) Get_Report() *IMagneticStripeReaderReport {
	var _result *IMagneticStripeReaderReport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Report, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMagneticStripeReaderAamvaCardDataReceivedEventArgs) Get_LicenseNumber() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LicenseNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReaderAamvaCardDataReceivedEventArgs) Get_ExpirationDate() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExpirationDate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReaderAamvaCardDataReceivedEventArgs) Get_Restrictions() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Restrictions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReaderAamvaCardDataReceivedEventArgs) Get_Class() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Class, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReaderAamvaCardDataReceivedEventArgs) Get_Endorsements() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Endorsements, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReaderAamvaCardDataReceivedEventArgs) Get_BirthDate() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BirthDate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReaderAamvaCardDataReceivedEventArgs) Get_FirstName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FirstName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReaderAamvaCardDataReceivedEventArgs) Get_Surname() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Surname, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReaderAamvaCardDataReceivedEventArgs) Get_Suffix() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Suffix, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReaderAamvaCardDataReceivedEventArgs) Get_Gender() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Gender, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReaderAamvaCardDataReceivedEventArgs) Get_HairColor() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HairColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReaderAamvaCardDataReceivedEventArgs) Get_EyeColor() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EyeColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReaderAamvaCardDataReceivedEventArgs) Get_Height() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Height, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReaderAamvaCardDataReceivedEventArgs) Get_Weight() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Weight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReaderAamvaCardDataReceivedEventArgs) Get_Address() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Address, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReaderAamvaCardDataReceivedEventArgs) Get_City() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_City, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReaderAamvaCardDataReceivedEventArgs) Get_State() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReaderAamvaCardDataReceivedEventArgs) Get_PostalCode() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PostalCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 2E958823-A31A-4763-882C-23725E39B08E
var IID_IMagneticStripeReaderBankCardDataReceivedEventArgs = syscall.GUID{0x2E958823, 0xA31A, 0x4763,
	[8]byte{0x88, 0x2C, 0x23, 0x72, 0x5E, 0x39, 0xB0, 0x8E}}

type IMagneticStripeReaderBankCardDataReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Report() *IMagneticStripeReaderReport
	Get_AccountNumber() string
	Get_ExpirationDate() string
	Get_ServiceCode() string
	Get_Title() string
	Get_FirstName() string
	Get_MiddleInitial() string
	Get_Surname() string
	Get_Suffix() string
}

type IMagneticStripeReaderBankCardDataReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Report         uintptr
	Get_AccountNumber  uintptr
	Get_ExpirationDate uintptr
	Get_ServiceCode    uintptr
	Get_Title          uintptr
	Get_FirstName      uintptr
	Get_MiddleInitial  uintptr
	Get_Surname        uintptr
	Get_Suffix         uintptr
}

type IMagneticStripeReaderBankCardDataReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IMagneticStripeReaderBankCardDataReceivedEventArgs) Vtbl() *IMagneticStripeReaderBankCardDataReceivedEventArgsVtbl {
	return (*IMagneticStripeReaderBankCardDataReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMagneticStripeReaderBankCardDataReceivedEventArgs) Get_Report() *IMagneticStripeReaderReport {
	var _result *IMagneticStripeReaderReport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Report, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMagneticStripeReaderBankCardDataReceivedEventArgs) Get_AccountNumber() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AccountNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReaderBankCardDataReceivedEventArgs) Get_ExpirationDate() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExpirationDate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReaderBankCardDataReceivedEventArgs) Get_ServiceCode() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReaderBankCardDataReceivedEventArgs) Get_Title() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReaderBankCardDataReceivedEventArgs) Get_FirstName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FirstName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReaderBankCardDataReceivedEventArgs) Get_MiddleInitial() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MiddleInitial, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReaderBankCardDataReceivedEventArgs) Get_Surname() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Surname, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReaderBankCardDataReceivedEventArgs) Get_Suffix() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Suffix, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 7128809C-C440-44A2-A467-469175D02896
var IID_IMagneticStripeReaderCapabilities = syscall.GUID{0x7128809C, 0xC440, 0x44A2,
	[8]byte{0xA4, 0x67, 0x46, 0x91, 0x75, 0xD0, 0x28, 0x96}}

type IMagneticStripeReaderCapabilitiesInterface interface {
	win32.IInspectableInterface
	Get_CardAuthentication() string
	Get_SupportedEncryptionAlgorithms() uint32
	Get_AuthenticationLevel() MagneticStripeReaderAuthenticationLevel
	Get_IsIsoSupported() bool
	Get_IsJisOneSupported() bool
	Get_IsJisTwoSupported() bool
	Get_PowerReportingType() UnifiedPosPowerReportingType
	Get_IsStatisticsReportingSupported() bool
	Get_IsStatisticsUpdatingSupported() bool
	Get_IsTrackDataMaskingSupported() bool
	Get_IsTransmitSentinelsSupported() bool
}

type IMagneticStripeReaderCapabilitiesVtbl struct {
	win32.IInspectableVtbl
	Get_CardAuthentication             uintptr
	Get_SupportedEncryptionAlgorithms  uintptr
	Get_AuthenticationLevel            uintptr
	Get_IsIsoSupported                 uintptr
	Get_IsJisOneSupported              uintptr
	Get_IsJisTwoSupported              uintptr
	Get_PowerReportingType             uintptr
	Get_IsStatisticsReportingSupported uintptr
	Get_IsStatisticsUpdatingSupported  uintptr
	Get_IsTrackDataMaskingSupported    uintptr
	Get_IsTransmitSentinelsSupported   uintptr
}

type IMagneticStripeReaderCapabilities struct {
	win32.IInspectable
}

func (this *IMagneticStripeReaderCapabilities) Vtbl() *IMagneticStripeReaderCapabilitiesVtbl {
	return (*IMagneticStripeReaderCapabilitiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMagneticStripeReaderCapabilities) Get_CardAuthentication() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CardAuthentication, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagneticStripeReaderCapabilities) Get_SupportedEncryptionAlgorithms() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedEncryptionAlgorithms, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagneticStripeReaderCapabilities) Get_AuthenticationLevel() MagneticStripeReaderAuthenticationLevel {
	var _result MagneticStripeReaderAuthenticationLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AuthenticationLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagneticStripeReaderCapabilities) Get_IsIsoSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsIsoSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagneticStripeReaderCapabilities) Get_IsJisOneSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsJisOneSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagneticStripeReaderCapabilities) Get_IsJisTwoSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsJisTwoSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagneticStripeReaderCapabilities) Get_PowerReportingType() UnifiedPosPowerReportingType {
	var _result UnifiedPosPowerReportingType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PowerReportingType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagneticStripeReaderCapabilities) Get_IsStatisticsReportingSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsStatisticsReportingSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagneticStripeReaderCapabilities) Get_IsStatisticsUpdatingSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsStatisticsUpdatingSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagneticStripeReaderCapabilities) Get_IsTrackDataMaskingSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsTrackDataMaskingSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagneticStripeReaderCapabilities) Get_IsTransmitSentinelsSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsTransmitSentinelsSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 528F2C5D-2986-474F-8454-7CCD05928D5F
var IID_IMagneticStripeReaderCardTypesStatics = syscall.GUID{0x528F2C5D, 0x2986, 0x474F,
	[8]byte{0x84, 0x54, 0x7C, 0xCD, 0x05, 0x92, 0x8D, 0x5F}}

type IMagneticStripeReaderCardTypesStaticsInterface interface {
	win32.IInspectableInterface
	Get_Unknown() uint32
	Get_Bank() uint32
	Get_Aamva() uint32
	Get_ExtendedBase() uint32
}

type IMagneticStripeReaderCardTypesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Unknown      uintptr
	Get_Bank         uintptr
	Get_Aamva        uintptr
	Get_ExtendedBase uintptr
}

type IMagneticStripeReaderCardTypesStatics struct {
	win32.IInspectable
}

func (this *IMagneticStripeReaderCardTypesStatics) Vtbl() *IMagneticStripeReaderCardTypesStaticsVtbl {
	return (*IMagneticStripeReaderCardTypesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMagneticStripeReaderCardTypesStatics) Get_Unknown() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Unknown, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagneticStripeReaderCardTypesStatics) Get_Bank() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bank, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagneticStripeReaderCardTypesStatics) Get_Aamva() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Aamva, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagneticStripeReaderCardTypesStatics) Get_ExtendedBase() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedBase, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 53B57350-C3DB-4754-9C00-41392374A109
var IID_IMagneticStripeReaderEncryptionAlgorithmsStatics = syscall.GUID{0x53B57350, 0xC3DB, 0x4754,
	[8]byte{0x9C, 0x00, 0x41, 0x39, 0x23, 0x74, 0xA1, 0x09}}

type IMagneticStripeReaderEncryptionAlgorithmsStaticsInterface interface {
	win32.IInspectableInterface
	Get_None() uint32
	Get_TripleDesDukpt() uint32
	Get_ExtendedBase() uint32
}

type IMagneticStripeReaderEncryptionAlgorithmsStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_None           uintptr
	Get_TripleDesDukpt uintptr
	Get_ExtendedBase   uintptr
}

type IMagneticStripeReaderEncryptionAlgorithmsStatics struct {
	win32.IInspectable
}

func (this *IMagneticStripeReaderEncryptionAlgorithmsStatics) Vtbl() *IMagneticStripeReaderEncryptionAlgorithmsStaticsVtbl {
	return (*IMagneticStripeReaderEncryptionAlgorithmsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMagneticStripeReaderEncryptionAlgorithmsStatics) Get_None() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_None, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagneticStripeReaderEncryptionAlgorithmsStatics) Get_TripleDesDukpt() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TripleDesDukpt, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagneticStripeReaderEncryptionAlgorithmsStatics) Get_ExtendedBase() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedBase, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1FEDF95D-2C84-41AD-B778-F2356A789AB1
var IID_IMagneticStripeReaderErrorOccurredEventArgs = syscall.GUID{0x1FEDF95D, 0x2C84, 0x41AD,
	[8]byte{0xB7, 0x78, 0xF2, 0x35, 0x6A, 0x78, 0x9A, 0xB1}}

type IMagneticStripeReaderErrorOccurredEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Track1Status() MagneticStripeReaderTrackErrorType
	Get_Track2Status() MagneticStripeReaderTrackErrorType
	Get_Track3Status() MagneticStripeReaderTrackErrorType
	Get_Track4Status() MagneticStripeReaderTrackErrorType
	Get_ErrorData() *IUnifiedPosErrorData
	Get_PartialInputData() *IMagneticStripeReaderReport
}

type IMagneticStripeReaderErrorOccurredEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Track1Status     uintptr
	Get_Track2Status     uintptr
	Get_Track3Status     uintptr
	Get_Track4Status     uintptr
	Get_ErrorData        uintptr
	Get_PartialInputData uintptr
}

type IMagneticStripeReaderErrorOccurredEventArgs struct {
	win32.IInspectable
}

func (this *IMagneticStripeReaderErrorOccurredEventArgs) Vtbl() *IMagneticStripeReaderErrorOccurredEventArgsVtbl {
	return (*IMagneticStripeReaderErrorOccurredEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMagneticStripeReaderErrorOccurredEventArgs) Get_Track1Status() MagneticStripeReaderTrackErrorType {
	var _result MagneticStripeReaderTrackErrorType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Track1Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagneticStripeReaderErrorOccurredEventArgs) Get_Track2Status() MagneticStripeReaderTrackErrorType {
	var _result MagneticStripeReaderTrackErrorType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Track2Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagneticStripeReaderErrorOccurredEventArgs) Get_Track3Status() MagneticStripeReaderTrackErrorType {
	var _result MagneticStripeReaderTrackErrorType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Track3Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagneticStripeReaderErrorOccurredEventArgs) Get_Track4Status() MagneticStripeReaderTrackErrorType {
	var _result MagneticStripeReaderTrackErrorType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Track4Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagneticStripeReaderErrorOccurredEventArgs) Get_ErrorData() *IUnifiedPosErrorData {
	var _result *IUnifiedPosErrorData
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMagneticStripeReaderErrorOccurredEventArgs) Get_PartialInputData() *IMagneticStripeReaderReport {
	var _result *IMagneticStripeReaderReport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PartialInputData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6A5B6047-99B0-4188-BEF1-EDDF79F78FE6
var IID_IMagneticStripeReaderReport = syscall.GUID{0x6A5B6047, 0x99B0, 0x4188,
	[8]byte{0xBE, 0xF1, 0xED, 0xDF, 0x79, 0xF7, 0x8F, 0xE6}}

type IMagneticStripeReaderReportInterface interface {
	win32.IInspectableInterface
	Get_CardType() uint32
	Get_Track1() *IMagneticStripeReaderTrackData
	Get_Track2() *IMagneticStripeReaderTrackData
	Get_Track3() *IMagneticStripeReaderTrackData
	Get_Track4() *IMagneticStripeReaderTrackData
	Get_Properties() *IMapView[string, string]
	Get_CardAuthenticationData() *IBuffer
	Get_CardAuthenticationDataLength() uint32
	Get_AdditionalSecurityInformation() *IBuffer
}

type IMagneticStripeReaderReportVtbl struct {
	win32.IInspectableVtbl
	Get_CardType                      uintptr
	Get_Track1                        uintptr
	Get_Track2                        uintptr
	Get_Track3                        uintptr
	Get_Track4                        uintptr
	Get_Properties                    uintptr
	Get_CardAuthenticationData        uintptr
	Get_CardAuthenticationDataLength  uintptr
	Get_AdditionalSecurityInformation uintptr
}

type IMagneticStripeReaderReport struct {
	win32.IInspectable
}

func (this *IMagneticStripeReaderReport) Vtbl() *IMagneticStripeReaderReportVtbl {
	return (*IMagneticStripeReaderReportVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMagneticStripeReaderReport) Get_CardType() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CardType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagneticStripeReaderReport) Get_Track1() *IMagneticStripeReaderTrackData {
	var _result *IMagneticStripeReaderTrackData
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Track1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMagneticStripeReaderReport) Get_Track2() *IMagneticStripeReaderTrackData {
	var _result *IMagneticStripeReaderTrackData
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Track2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMagneticStripeReaderReport) Get_Track3() *IMagneticStripeReaderTrackData {
	var _result *IMagneticStripeReaderTrackData
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Track3, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMagneticStripeReaderReport) Get_Track4() *IMagneticStripeReaderTrackData {
	var _result *IMagneticStripeReaderTrackData
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Track4, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMagneticStripeReaderReport) Get_Properties() *IMapView[string, string] {
	var _result *IMapView[string, string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMagneticStripeReaderReport) Get_CardAuthenticationData() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CardAuthenticationData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMagneticStripeReaderReport) Get_CardAuthenticationDataLength() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CardAuthenticationDataLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagneticStripeReaderReport) Get_AdditionalSecurityInformation() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AdditionalSecurityInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C45FAB4A-EFD7-4760-A5CE-15B0E47E94EB
var IID_IMagneticStripeReaderStatics = syscall.GUID{0xC45FAB4A, 0xEFD7, 0x4760,
	[8]byte{0xA5, 0xCE, 0x15, 0xB0, 0xE4, 0x7E, 0x94, 0xEB}}

type IMagneticStripeReaderStaticsInterface interface {
	win32.IInspectableInterface
	GetDefaultAsync() *IAsyncOperation[*IMagneticStripeReader]
	FromIdAsync(deviceId string) *IAsyncOperation[*IMagneticStripeReader]
	GetDeviceSelector() string
}

type IMagneticStripeReaderStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefaultAsync   uintptr
	FromIdAsync       uintptr
	GetDeviceSelector uintptr
}

type IMagneticStripeReaderStatics struct {
	win32.IInspectable
}

func (this *IMagneticStripeReaderStatics) Vtbl() *IMagneticStripeReaderStaticsVtbl {
	return (*IMagneticStripeReaderStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMagneticStripeReaderStatics) GetDefaultAsync() *IAsyncOperation[*IMagneticStripeReader] {
	var _result *IAsyncOperation[*IMagneticStripeReader]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMagneticStripeReaderStatics) FromIdAsync(deviceId string) *IAsyncOperation[*IMagneticStripeReader] {
	var _result *IAsyncOperation[*IMagneticStripeReader]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMagneticStripeReaderStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 8CADC362-D667-48FA-86BC-F5AE1189262B
var IID_IMagneticStripeReaderStatics2 = syscall.GUID{0x8CADC362, 0xD667, 0x48FA,
	[8]byte{0x86, 0xBC, 0xF5, 0xAE, 0x11, 0x89, 0x26, 0x2B}}

type IMagneticStripeReaderStatics2Interface interface {
	win32.IInspectableInterface
	GetDeviceSelectorWithConnectionTypes(connectionTypes PosConnectionTypes) string
}

type IMagneticStripeReaderStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelectorWithConnectionTypes uintptr
}

type IMagneticStripeReaderStatics2 struct {
	win32.IInspectable
}

func (this *IMagneticStripeReaderStatics2) Vtbl() *IMagneticStripeReaderStatics2Vtbl {
	return (*IMagneticStripeReaderStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMagneticStripeReaderStatics2) GetDeviceSelectorWithConnectionTypes(connectionTypes PosConnectionTypes) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorWithConnectionTypes, uintptr(unsafe.Pointer(this)), uintptr(connectionTypes), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 09CC6BB0-3262-401D-9E8A-E80D6358906B
var IID_IMagneticStripeReaderStatusUpdatedEventArgs = syscall.GUID{0x09CC6BB0, 0x3262, 0x401D,
	[8]byte{0x9E, 0x8A, 0xE8, 0x0D, 0x63, 0x58, 0x90, 0x6B}}

type IMagneticStripeReaderStatusUpdatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Status() MagneticStripeReaderStatus
	Get_ExtendedStatus() uint32
}

type IMagneticStripeReaderStatusUpdatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Status         uintptr
	Get_ExtendedStatus uintptr
}

type IMagneticStripeReaderStatusUpdatedEventArgs struct {
	win32.IInspectable
}

func (this *IMagneticStripeReaderStatusUpdatedEventArgs) Vtbl() *IMagneticStripeReaderStatusUpdatedEventArgsVtbl {
	return (*IMagneticStripeReaderStatusUpdatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMagneticStripeReaderStatusUpdatedEventArgs) Get_Status() MagneticStripeReaderStatus {
	var _result MagneticStripeReaderStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagneticStripeReaderStatusUpdatedEventArgs) Get_ExtendedStatus() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 104CF671-4A9D-446E-ABC5-20402307BA36
var IID_IMagneticStripeReaderTrackData = syscall.GUID{0x104CF671, 0x4A9D, 0x446E,
	[8]byte{0xAB, 0xC5, 0x20, 0x40, 0x23, 0x07, 0xBA, 0x36}}

type IMagneticStripeReaderTrackDataInterface interface {
	win32.IInspectableInterface
	Get_Data() *IBuffer
	Get_DiscretionaryData() *IBuffer
	Get_EncryptedData() *IBuffer
}

type IMagneticStripeReaderTrackDataVtbl struct {
	win32.IInspectableVtbl
	Get_Data              uintptr
	Get_DiscretionaryData uintptr
	Get_EncryptedData     uintptr
}

type IMagneticStripeReaderTrackData struct {
	win32.IInspectable
}

func (this *IMagneticStripeReaderTrackData) Vtbl() *IMagneticStripeReaderTrackDataVtbl {
	return (*IMagneticStripeReaderTrackDataVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMagneticStripeReaderTrackData) Get_Data() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMagneticStripeReaderTrackData) Get_DiscretionaryData() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DiscretionaryData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMagneticStripeReaderTrackData) Get_EncryptedData() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EncryptedData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AF0A5514-59CC-4A60-99E8-99A53DACE5AA
var IID_IMagneticStripeReaderVendorSpecificCardDataReceivedEventArgs = syscall.GUID{0xAF0A5514, 0x59CC, 0x4A60,
	[8]byte{0x99, 0xE8, 0x99, 0xA5, 0x3D, 0xAC, 0xE5, 0xAA}}

type IMagneticStripeReaderVendorSpecificCardDataReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Report() *IMagneticStripeReaderReport
}

type IMagneticStripeReaderVendorSpecificCardDataReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Report uintptr
}

type IMagneticStripeReaderVendorSpecificCardDataReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IMagneticStripeReaderVendorSpecificCardDataReceivedEventArgs) Vtbl() *IMagneticStripeReaderVendorSpecificCardDataReceivedEventArgsVtbl {
	return (*IMagneticStripeReaderVendorSpecificCardDataReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMagneticStripeReaderVendorSpecificCardDataReceivedEventArgs) Get_Report() *IMagneticStripeReaderReport {
	var _result *IMagneticStripeReaderReport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Report, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2A03C10E-9A19-4A01-994F-12DFAD6ADCBF
var IID_IPosPrinter = syscall.GUID{0x2A03C10E, 0x9A19, 0x4A01,
	[8]byte{0x99, 0x4F, 0x12, 0xDF, 0xAD, 0x6A, 0xDC, 0xBF}}

type IPosPrinterInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_Capabilities() *IPosPrinterCapabilities
	Get_SupportedCharacterSets() *IVectorView[uint32]
	Get_SupportedTypeFaces() *IVectorView[string]
	Get_Status() *IPosPrinterStatus
	ClaimPrinterAsync() *IAsyncOperation[*IClaimedPosPrinter]
	CheckHealthAsync(level UnifiedPosHealthCheckLevel) *IAsyncOperation[string]
	GetStatisticsAsync(statisticsCategories *IIterable[string]) *IAsyncOperation[string]
	Add_StatusUpdated(handler TypedEventHandler[*IPosPrinter, *IPosPrinterStatusUpdatedEventArgs]) EventRegistrationToken
	Remove_StatusUpdated(token EventRegistrationToken)
}

type IPosPrinterVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId               uintptr
	Get_Capabilities           uintptr
	Get_SupportedCharacterSets uintptr
	Get_SupportedTypeFaces     uintptr
	Get_Status                 uintptr
	ClaimPrinterAsync          uintptr
	CheckHealthAsync           uintptr
	GetStatisticsAsync         uintptr
	Add_StatusUpdated          uintptr
	Remove_StatusUpdated       uintptr
}

type IPosPrinter struct {
	win32.IInspectable
}

func (this *IPosPrinter) Vtbl() *IPosPrinterVtbl {
	return (*IPosPrinterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPosPrinter) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPosPrinter) Get_Capabilities() *IPosPrinterCapabilities {
	var _result *IPosPrinterCapabilities
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Capabilities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPosPrinter) Get_SupportedCharacterSets() *IVectorView[uint32] {
	var _result *IVectorView[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedCharacterSets, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPosPrinter) Get_SupportedTypeFaces() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedTypeFaces, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPosPrinter) Get_Status() *IPosPrinterStatus {
	var _result *IPosPrinterStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPosPrinter) ClaimPrinterAsync() *IAsyncOperation[*IClaimedPosPrinter] {
	var _result *IAsyncOperation[*IClaimedPosPrinter]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClaimPrinterAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPosPrinter) CheckHealthAsync(level UnifiedPosHealthCheckLevel) *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CheckHealthAsync, uintptr(unsafe.Pointer(this)), uintptr(level), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPosPrinter) GetStatisticsAsync(statisticsCategories *IIterable[string]) *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStatisticsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(statisticsCategories)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPosPrinter) Add_StatusUpdated(handler TypedEventHandler[*IPosPrinter, *IPosPrinterStatusUpdatedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StatusUpdated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPosPrinter) Remove_StatusUpdated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StatusUpdated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 248475E8-8B98-5517-8E48-760E86F68987
var IID_IPosPrinter2 = syscall.GUID{0x248475E8, 0x8B98, 0x5517,
	[8]byte{0x8E, 0x48, 0x76, 0x0E, 0x86, 0xF6, 0x89, 0x87}}

type IPosPrinter2Interface interface {
	win32.IInspectableInterface
	Get_SupportedBarcodeSymbologies() *IVectorView[uint32]
	GetFontProperty(typeface string) *IPosPrinterFontProperty
}

type IPosPrinter2Vtbl struct {
	win32.IInspectableVtbl
	Get_SupportedBarcodeSymbologies uintptr
	GetFontProperty                 uintptr
}

type IPosPrinter2 struct {
	win32.IInspectable
}

func (this *IPosPrinter2) Vtbl() *IPosPrinter2Vtbl {
	return (*IPosPrinter2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPosPrinter2) Get_SupportedBarcodeSymbologies() *IVectorView[uint32] {
	var _result *IVectorView[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedBarcodeSymbologies, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPosPrinter2) GetFontProperty(typeface string) *IPosPrinterFontProperty {
	var _result *IPosPrinterFontProperty
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFontProperty, uintptr(unsafe.Pointer(this)), NewHStr(typeface).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CDE95721-4380-4985-ADC5-39DB30CD93BC
var IID_IPosPrinterCapabilities = syscall.GUID{0xCDE95721, 0x4380, 0x4985,
	[8]byte{0xAD, 0xC5, 0x39, 0xDB, 0x30, 0xCD, 0x93, 0xBC}}

type IPosPrinterCapabilitiesInterface interface {
	win32.IInspectableInterface
	Get_PowerReportingType() UnifiedPosPowerReportingType
	Get_IsStatisticsReportingSupported() bool
	Get_IsStatisticsUpdatingSupported() bool
	Get_DefaultCharacterSet() uint32
	Get_HasCoverSensor() bool
	Get_CanMapCharacterSet() bool
	Get_IsTransactionSupported() bool
	Get_Receipt() *IReceiptPrinterCapabilities
	Get_Slip() *ISlipPrinterCapabilities
	Get_Journal() *IJournalPrinterCapabilities
}

type IPosPrinterCapabilitiesVtbl struct {
	win32.IInspectableVtbl
	Get_PowerReportingType             uintptr
	Get_IsStatisticsReportingSupported uintptr
	Get_IsStatisticsUpdatingSupported  uintptr
	Get_DefaultCharacterSet            uintptr
	Get_HasCoverSensor                 uintptr
	Get_CanMapCharacterSet             uintptr
	Get_IsTransactionSupported         uintptr
	Get_Receipt                        uintptr
	Get_Slip                           uintptr
	Get_Journal                        uintptr
}

type IPosPrinterCapabilities struct {
	win32.IInspectable
}

func (this *IPosPrinterCapabilities) Vtbl() *IPosPrinterCapabilitiesVtbl {
	return (*IPosPrinterCapabilitiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPosPrinterCapabilities) Get_PowerReportingType() UnifiedPosPowerReportingType {
	var _result UnifiedPosPowerReportingType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PowerReportingType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPosPrinterCapabilities) Get_IsStatisticsReportingSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsStatisticsReportingSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPosPrinterCapabilities) Get_IsStatisticsUpdatingSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsStatisticsUpdatingSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPosPrinterCapabilities) Get_DefaultCharacterSet() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DefaultCharacterSet, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPosPrinterCapabilities) Get_HasCoverSensor() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasCoverSensor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPosPrinterCapabilities) Get_CanMapCharacterSet() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanMapCharacterSet, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPosPrinterCapabilities) Get_IsTransactionSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsTransactionSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPosPrinterCapabilities) Get_Receipt() *IReceiptPrinterCapabilities {
	var _result *IReceiptPrinterCapabilities
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Receipt, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPosPrinterCapabilities) Get_Slip() *ISlipPrinterCapabilities {
	var _result *ISlipPrinterCapabilities
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Slip, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPosPrinterCapabilities) Get_Journal() *IJournalPrinterCapabilities {
	var _result *IJournalPrinterCapabilities
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Journal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5C709EFF-709A-4FE7-B215-06A748A38B39
var IID_IPosPrinterCharacterSetIdsStatics = syscall.GUID{0x5C709EFF, 0x709A, 0x4FE7,
	[8]byte{0xB2, 0x15, 0x06, 0xA7, 0x48, 0xA3, 0x8B, 0x39}}

type IPosPrinterCharacterSetIdsStaticsInterface interface {
	win32.IInspectableInterface
	Get_Utf16LE() uint32
	Get_Ascii() uint32
	Get_Ansi() uint32
}

type IPosPrinterCharacterSetIdsStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Utf16LE uintptr
	Get_Ascii   uintptr
	Get_Ansi    uintptr
}

type IPosPrinterCharacterSetIdsStatics struct {
	win32.IInspectable
}

func (this *IPosPrinterCharacterSetIdsStatics) Vtbl() *IPosPrinterCharacterSetIdsStaticsVtbl {
	return (*IPosPrinterCharacterSetIdsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPosPrinterCharacterSetIdsStatics) Get_Utf16LE() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Utf16LE, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPosPrinterCharacterSetIdsStatics) Get_Ascii() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ascii, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPosPrinterCharacterSetIdsStatics) Get_Ansi() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ansi, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A7F4E93A-F8AC-5F04-84D2-29B16D8A633C
var IID_IPosPrinterFontProperty = syscall.GUID{0xA7F4E93A, 0xF8AC, 0x5F04,
	[8]byte{0x84, 0xD2, 0x29, 0xB1, 0x6D, 0x8A, 0x63, 0x3C}}

type IPosPrinterFontPropertyInterface interface {
	win32.IInspectableInterface
	Get_TypeFace() string
	Get_IsScalableToAnySize() bool
	Get_CharacterSizes() *IVectorView[SizeUInt32]
}

type IPosPrinterFontPropertyVtbl struct {
	win32.IInspectableVtbl
	Get_TypeFace            uintptr
	Get_IsScalableToAnySize uintptr
	Get_CharacterSizes      uintptr
}

type IPosPrinterFontProperty struct {
	win32.IInspectable
}

func (this *IPosPrinterFontProperty) Vtbl() *IPosPrinterFontPropertyVtbl {
	return (*IPosPrinterFontPropertyVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPosPrinterFontProperty) Get_TypeFace() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TypeFace, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPosPrinterFontProperty) Get_IsScalableToAnySize() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsScalableToAnySize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPosPrinterFontProperty) Get_CharacterSizes() *IVectorView[SizeUInt32] {
	var _result *IVectorView[SizeUInt32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CharacterSizes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9A94005C-0615-4591-A58F-30F87EDFE2E4
var IID_IPosPrinterJob = syscall.GUID{0x9A94005C, 0x0615, 0x4591,
	[8]byte{0xA5, 0x8F, 0x30, 0xF8, 0x7E, 0xDF, 0xE2, 0xE4}}

type IPosPrinterJobInterface interface {
	win32.IInspectableInterface
	Print(data string)
	PrintLine(data string)
	PrintNewline()
	ExecuteAsync() *IAsyncOperation[bool]
}

type IPosPrinterJobVtbl struct {
	win32.IInspectableVtbl
	Print        uintptr
	PrintLine    uintptr
	PrintNewline uintptr
	ExecuteAsync uintptr
}

type IPosPrinterJob struct {
	win32.IInspectable
}

func (this *IPosPrinterJob) Vtbl() *IPosPrinterJobVtbl {
	return (*IPosPrinterJobVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPosPrinterJob) Print(data string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Print, uintptr(unsafe.Pointer(this)), NewHStr(data).Ptr)
	_ = _hr
}

func (this *IPosPrinterJob) PrintLine(data string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PrintLine, uintptr(unsafe.Pointer(this)), NewHStr(data).Ptr)
	_ = _hr
}

func (this *IPosPrinterJob) PrintNewline() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PrintNewline, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IPosPrinterJob) ExecuteAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ExecuteAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0A2E16FD-1D02-5A58-9D59-BFCDE76FDE86
var IID_IPosPrinterPrintOptions = syscall.GUID{0x0A2E16FD, 0x1D02, 0x5A58,
	[8]byte{0x9D, 0x59, 0xBF, 0xCD, 0xE7, 0x6F, 0xDE, 0x86}}

type IPosPrinterPrintOptionsInterface interface {
	win32.IInspectableInterface
	Get_TypeFace() string
	Put_TypeFace(value string)
	Get_CharacterHeight() uint32
	Put_CharacterHeight(value uint32)
	Get_Bold() bool
	Put_Bold(value bool)
	Get_Italic() bool
	Put_Italic(value bool)
	Get_Underline() bool
	Put_Underline(value bool)
	Get_ReverseVideo() bool
	Put_ReverseVideo(value bool)
	Get_Strikethrough() bool
	Put_Strikethrough(value bool)
	Get_Superscript() bool
	Put_Superscript(value bool)
	Get_Subscript() bool
	Put_Subscript(value bool)
	Get_DoubleWide() bool
	Put_DoubleWide(value bool)
	Get_DoubleHigh() bool
	Put_DoubleHigh(value bool)
	Get_Alignment() PosPrinterAlignment
	Put_Alignment(value PosPrinterAlignment)
	Get_CharacterSet() uint32
	Put_CharacterSet(value uint32)
}

type IPosPrinterPrintOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_TypeFace        uintptr
	Put_TypeFace        uintptr
	Get_CharacterHeight uintptr
	Put_CharacterHeight uintptr
	Get_Bold            uintptr
	Put_Bold            uintptr
	Get_Italic          uintptr
	Put_Italic          uintptr
	Get_Underline       uintptr
	Put_Underline       uintptr
	Get_ReverseVideo    uintptr
	Put_ReverseVideo    uintptr
	Get_Strikethrough   uintptr
	Put_Strikethrough   uintptr
	Get_Superscript     uintptr
	Put_Superscript     uintptr
	Get_Subscript       uintptr
	Put_Subscript       uintptr
	Get_DoubleWide      uintptr
	Put_DoubleWide      uintptr
	Get_DoubleHigh      uintptr
	Put_DoubleHigh      uintptr
	Get_Alignment       uintptr
	Put_Alignment       uintptr
	Get_CharacterSet    uintptr
	Put_CharacterSet    uintptr
}

type IPosPrinterPrintOptions struct {
	win32.IInspectable
}

func (this *IPosPrinterPrintOptions) Vtbl() *IPosPrinterPrintOptionsVtbl {
	return (*IPosPrinterPrintOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPosPrinterPrintOptions) Get_TypeFace() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TypeFace, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPosPrinterPrintOptions) Put_TypeFace(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TypeFace, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPosPrinterPrintOptions) Get_CharacterHeight() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CharacterHeight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPosPrinterPrintOptions) Put_CharacterHeight(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CharacterHeight, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPosPrinterPrintOptions) Get_Bold() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bold, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPosPrinterPrintOptions) Put_Bold(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Bold, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPosPrinterPrintOptions) Get_Italic() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Italic, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPosPrinterPrintOptions) Put_Italic(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Italic, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPosPrinterPrintOptions) Get_Underline() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Underline, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPosPrinterPrintOptions) Put_Underline(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Underline, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPosPrinterPrintOptions) Get_ReverseVideo() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReverseVideo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPosPrinterPrintOptions) Put_ReverseVideo(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReverseVideo, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPosPrinterPrintOptions) Get_Strikethrough() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Strikethrough, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPosPrinterPrintOptions) Put_Strikethrough(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Strikethrough, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPosPrinterPrintOptions) Get_Superscript() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Superscript, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPosPrinterPrintOptions) Put_Superscript(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Superscript, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPosPrinterPrintOptions) Get_Subscript() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Subscript, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPosPrinterPrintOptions) Put_Subscript(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Subscript, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPosPrinterPrintOptions) Get_DoubleWide() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DoubleWide, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPosPrinterPrintOptions) Put_DoubleWide(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DoubleWide, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPosPrinterPrintOptions) Get_DoubleHigh() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DoubleHigh, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPosPrinterPrintOptions) Put_DoubleHigh(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DoubleHigh, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPosPrinterPrintOptions) Get_Alignment() PosPrinterAlignment {
	var _result PosPrinterAlignment
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Alignment, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPosPrinterPrintOptions) Put_Alignment(value PosPrinterAlignment) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Alignment, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPosPrinterPrintOptions) Get_CharacterSet() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CharacterSet, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPosPrinterPrintOptions) Put_CharacterSet(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CharacterSet, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 2BCBA359-1CEF-40B2-9ECB-F927F856AE3C
var IID_IPosPrinterReleaseDeviceRequestedEventArgs = syscall.GUID{0x2BCBA359, 0x1CEF, 0x40B2,
	[8]byte{0x9E, 0xCB, 0xF9, 0x27, 0xF8, 0x56, 0xAE, 0x3C}}

type IPosPrinterReleaseDeviceRequestedEventArgsInterface interface {
	win32.IInspectableInterface
}

type IPosPrinterReleaseDeviceRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
}

type IPosPrinterReleaseDeviceRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IPosPrinterReleaseDeviceRequestedEventArgs) Vtbl() *IPosPrinterReleaseDeviceRequestedEventArgsVtbl {
	return (*IPosPrinterReleaseDeviceRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 8CE0D4EA-132F-4CDF-A64A-2D0D7C96A85B
var IID_IPosPrinterStatics = syscall.GUID{0x8CE0D4EA, 0x132F, 0x4CDF,
	[8]byte{0xA6, 0x4A, 0x2D, 0x0D, 0x7C, 0x96, 0xA8, 0x5B}}

type IPosPrinterStaticsInterface interface {
	win32.IInspectableInterface
	GetDefaultAsync() *IAsyncOperation[*IPosPrinter]
	FromIdAsync(deviceId string) *IAsyncOperation[*IPosPrinter]
	GetDeviceSelector() string
}

type IPosPrinterStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefaultAsync   uintptr
	FromIdAsync       uintptr
	GetDeviceSelector uintptr
}

type IPosPrinterStatics struct {
	win32.IInspectable
}

func (this *IPosPrinterStatics) Vtbl() *IPosPrinterStaticsVtbl {
	return (*IPosPrinterStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPosPrinterStatics) GetDefaultAsync() *IAsyncOperation[*IPosPrinter] {
	var _result *IAsyncOperation[*IPosPrinter]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPosPrinterStatics) FromIdAsync(deviceId string) *IAsyncOperation[*IPosPrinter] {
	var _result *IAsyncOperation[*IPosPrinter]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPosPrinterStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// EECD2C1C-B0D0-42E7-B137-B89B16244D41
var IID_IPosPrinterStatics2 = syscall.GUID{0xEECD2C1C, 0xB0D0, 0x42E7,
	[8]byte{0xB1, 0x37, 0xB8, 0x9B, 0x16, 0x24, 0x4D, 0x41}}

type IPosPrinterStatics2Interface interface {
	win32.IInspectableInterface
	GetDeviceSelectorWithConnectionTypes(connectionTypes PosConnectionTypes) string
}

type IPosPrinterStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelectorWithConnectionTypes uintptr
}

type IPosPrinterStatics2 struct {
	win32.IInspectable
}

func (this *IPosPrinterStatics2) Vtbl() *IPosPrinterStatics2Vtbl {
	return (*IPosPrinterStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPosPrinterStatics2) GetDeviceSelectorWithConnectionTypes(connectionTypes PosConnectionTypes) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorWithConnectionTypes, uintptr(unsafe.Pointer(this)), uintptr(connectionTypes), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// D1F0C730-DA40-4328-BF76-5156FA33B747
var IID_IPosPrinterStatus = syscall.GUID{0xD1F0C730, 0xDA40, 0x4328,
	[8]byte{0xBF, 0x76, 0x51, 0x56, 0xFA, 0x33, 0xB7, 0x47}}

type IPosPrinterStatusInterface interface {
	win32.IInspectableInterface
	Get_StatusKind() PosPrinterStatusKind
	Get_ExtendedStatus() uint32
}

type IPosPrinterStatusVtbl struct {
	win32.IInspectableVtbl
	Get_StatusKind     uintptr
	Get_ExtendedStatus uintptr
}

type IPosPrinterStatus struct {
	win32.IInspectable
}

func (this *IPosPrinterStatus) Vtbl() *IPosPrinterStatusVtbl {
	return (*IPosPrinterStatusVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPosPrinterStatus) Get_StatusKind() PosPrinterStatusKind {
	var _result PosPrinterStatusKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StatusKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPosPrinterStatus) Get_ExtendedStatus() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 2EDB87DF-13A6-428D-BA81-B0E7C3E5A3CD
var IID_IPosPrinterStatusUpdatedEventArgs = syscall.GUID{0x2EDB87DF, 0x13A6, 0x428D,
	[8]byte{0xBA, 0x81, 0xB0, 0xE7, 0xC3, 0xE5, 0xA3, 0xCD}}

type IPosPrinterStatusUpdatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Status() *IPosPrinterStatus
}

type IPosPrinterStatusUpdatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
}

type IPosPrinterStatusUpdatedEventArgs struct {
	win32.IInspectable
}

func (this *IPosPrinterStatusUpdatedEventArgs) Vtbl() *IPosPrinterStatusUpdatedEventArgsVtbl {
	return (*IPosPrinterStatusUpdatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPosPrinterStatusUpdatedEventArgs) Get_Status() *IPosPrinterStatus {
	var _result *IPosPrinterStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 532199BE-C8C3-4DC2-89E9-5C4A37B34DDC
var IID_IReceiptOrSlipJob = syscall.GUID{0x532199BE, 0xC8C3, 0x4DC2,
	[8]byte{0x89, 0xE9, 0x5C, 0x4A, 0x37, 0xB3, 0x4D, 0xDC}}

type IReceiptOrSlipJobInterface interface {
	win32.IInspectableInterface
	SetBarcodeRotation(value PosPrinterRotation)
	SetPrintRotation(value PosPrinterRotation, includeBitmaps bool)
	SetPrintArea(value Rect)
	SetBitmap(bitmapNumber uint32, bitmap *IBitmapFrame, alignment PosPrinterAlignment)
	SetBitmapCustomWidthStandardAlign(bitmapNumber uint32, bitmap *IBitmapFrame, alignment PosPrinterAlignment, width uint32)
	SetCustomAlignedBitmap(bitmapNumber uint32, bitmap *IBitmapFrame, alignmentDistance uint32)
	SetBitmapCustomWidthCustomAlign(bitmapNumber uint32, bitmap *IBitmapFrame, alignmentDistance uint32, width uint32)
	PrintSavedBitmap(bitmapNumber uint32)
	DrawRuledLine(positionList string, lineDirection PosPrinterLineDirection, lineWidth uint32, lineStyle PosPrinterLineStyle, lineColor uint32)
	PrintBarcode(data string, symbology uint32, height uint32, width uint32, textPosition PosPrinterBarcodeTextPosition, alignment PosPrinterAlignment)
	PrintBarcodeCustomAlign(data string, symbology uint32, height uint32, width uint32, textPosition PosPrinterBarcodeTextPosition, alignmentDistance uint32)
	PrintBitmap(bitmap *IBitmapFrame, alignment PosPrinterAlignment)
	PrintBitmapCustomWidthStandardAlign(bitmap *IBitmapFrame, alignment PosPrinterAlignment, width uint32)
	PrintCustomAlignedBitmap(bitmap *IBitmapFrame, alignmentDistance uint32)
	PrintBitmapCustomWidthCustomAlign(bitmap *IBitmapFrame, alignmentDistance uint32, width uint32)
}

type IReceiptOrSlipJobVtbl struct {
	win32.IInspectableVtbl
	SetBarcodeRotation                  uintptr
	SetPrintRotation                    uintptr
	SetPrintArea                        uintptr
	SetBitmap                           uintptr
	SetBitmapCustomWidthStandardAlign   uintptr
	SetCustomAlignedBitmap              uintptr
	SetBitmapCustomWidthCustomAlign     uintptr
	PrintSavedBitmap                    uintptr
	DrawRuledLine                       uintptr
	PrintBarcode                        uintptr
	PrintBarcodeCustomAlign             uintptr
	PrintBitmap                         uintptr
	PrintBitmapCustomWidthStandardAlign uintptr
	PrintCustomAlignedBitmap            uintptr
	PrintBitmapCustomWidthCustomAlign   uintptr
}

type IReceiptOrSlipJob struct {
	win32.IInspectable
}

func (this *IReceiptOrSlipJob) Vtbl() *IReceiptOrSlipJobVtbl {
	return (*IReceiptOrSlipJobVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IReceiptOrSlipJob) SetBarcodeRotation(value PosPrinterRotation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetBarcodeRotation, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IReceiptOrSlipJob) SetPrintRotation(value PosPrinterRotation, includeBitmaps bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPrintRotation, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(*(*byte)(unsafe.Pointer(&includeBitmaps))))
	_ = _hr
}

func (this *IReceiptOrSlipJob) SetPrintArea(value Rect) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPrintArea, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IReceiptOrSlipJob) SetBitmap(bitmapNumber uint32, bitmap *IBitmapFrame, alignment PosPrinterAlignment) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetBitmap, uintptr(unsafe.Pointer(this)), uintptr(bitmapNumber), uintptr(unsafe.Pointer(bitmap)), uintptr(alignment))
	_ = _hr
}

func (this *IReceiptOrSlipJob) SetBitmapCustomWidthStandardAlign(bitmapNumber uint32, bitmap *IBitmapFrame, alignment PosPrinterAlignment, width uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetBitmapCustomWidthStandardAlign, uintptr(unsafe.Pointer(this)), uintptr(bitmapNumber), uintptr(unsafe.Pointer(bitmap)), uintptr(alignment), uintptr(width))
	_ = _hr
}

func (this *IReceiptOrSlipJob) SetCustomAlignedBitmap(bitmapNumber uint32, bitmap *IBitmapFrame, alignmentDistance uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetCustomAlignedBitmap, uintptr(unsafe.Pointer(this)), uintptr(bitmapNumber), uintptr(unsafe.Pointer(bitmap)), uintptr(alignmentDistance))
	_ = _hr
}

func (this *IReceiptOrSlipJob) SetBitmapCustomWidthCustomAlign(bitmapNumber uint32, bitmap *IBitmapFrame, alignmentDistance uint32, width uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetBitmapCustomWidthCustomAlign, uintptr(unsafe.Pointer(this)), uintptr(bitmapNumber), uintptr(unsafe.Pointer(bitmap)), uintptr(alignmentDistance), uintptr(width))
	_ = _hr
}

func (this *IReceiptOrSlipJob) PrintSavedBitmap(bitmapNumber uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PrintSavedBitmap, uintptr(unsafe.Pointer(this)), uintptr(bitmapNumber))
	_ = _hr
}

func (this *IReceiptOrSlipJob) DrawRuledLine(positionList string, lineDirection PosPrinterLineDirection, lineWidth uint32, lineStyle PosPrinterLineStyle, lineColor uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DrawRuledLine, uintptr(unsafe.Pointer(this)), NewHStr(positionList).Ptr, uintptr(lineDirection), uintptr(lineWidth), uintptr(lineStyle), uintptr(lineColor))
	_ = _hr
}

func (this *IReceiptOrSlipJob) PrintBarcode(data string, symbology uint32, height uint32, width uint32, textPosition PosPrinterBarcodeTextPosition, alignment PosPrinterAlignment) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PrintBarcode, uintptr(unsafe.Pointer(this)), NewHStr(data).Ptr, uintptr(symbology), uintptr(height), uintptr(width), uintptr(textPosition), uintptr(alignment))
	_ = _hr
}

func (this *IReceiptOrSlipJob) PrintBarcodeCustomAlign(data string, symbology uint32, height uint32, width uint32, textPosition PosPrinterBarcodeTextPosition, alignmentDistance uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PrintBarcodeCustomAlign, uintptr(unsafe.Pointer(this)), NewHStr(data).Ptr, uintptr(symbology), uintptr(height), uintptr(width), uintptr(textPosition), uintptr(alignmentDistance))
	_ = _hr
}

func (this *IReceiptOrSlipJob) PrintBitmap(bitmap *IBitmapFrame, alignment PosPrinterAlignment) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PrintBitmap, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bitmap)), uintptr(alignment))
	_ = _hr
}

func (this *IReceiptOrSlipJob) PrintBitmapCustomWidthStandardAlign(bitmap *IBitmapFrame, alignment PosPrinterAlignment, width uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PrintBitmapCustomWidthStandardAlign, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bitmap)), uintptr(alignment), uintptr(width))
	_ = _hr
}

func (this *IReceiptOrSlipJob) PrintCustomAlignedBitmap(bitmap *IBitmapFrame, alignmentDistance uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PrintCustomAlignedBitmap, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bitmap)), uintptr(alignmentDistance))
	_ = _hr
}

func (this *IReceiptOrSlipJob) PrintBitmapCustomWidthCustomAlign(bitmap *IBitmapFrame, alignmentDistance uint32, width uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PrintBitmapCustomWidthCustomAlign, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bitmap)), uintptr(alignmentDistance), uintptr(width))
	_ = _hr
}

// AA96066E-ACAD-4B79-9D0F-C0CFC08DC77B
var IID_IReceiptPrintJob = syscall.GUID{0xAA96066E, 0xACAD, 0x4B79,
	[8]byte{0x9D, 0x0F, 0xC0, 0xCF, 0xC0, 0x8D, 0xC7, 0x7B}}

type IReceiptPrintJobInterface interface {
	win32.IInspectableInterface
	MarkFeed(kind PosPrinterMarkFeedKind)
	CutPaper(percentage float64)
	CutPaperDefault()
}

type IReceiptPrintJobVtbl struct {
	win32.IInspectableVtbl
	MarkFeed        uintptr
	CutPaper        uintptr
	CutPaperDefault uintptr
}

type IReceiptPrintJob struct {
	win32.IInspectable
}

func (this *IReceiptPrintJob) Vtbl() *IReceiptPrintJobVtbl {
	return (*IReceiptPrintJobVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IReceiptPrintJob) MarkFeed(kind PosPrinterMarkFeedKind) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MarkFeed, uintptr(unsafe.Pointer(this)), uintptr(kind))
	_ = _hr
}

func (this *IReceiptPrintJob) CutPaper(percentage float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CutPaper, uintptr(unsafe.Pointer(this)), uintptr(percentage))
	_ = _hr
}

func (this *IReceiptPrintJob) CutPaperDefault() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CutPaperDefault, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 0CBC12E3-9E29-5179-BCD8-1811D3B9A10E
var IID_IReceiptPrintJob2 = syscall.GUID{0x0CBC12E3, 0x9E29, 0x5179,
	[8]byte{0xBC, 0xD8, 0x18, 0x11, 0xD3, 0xB9, 0xA1, 0x0E}}

type IReceiptPrintJob2Interface interface {
	win32.IInspectableInterface
	StampPaper()
	Print(data string, printOptions *IPosPrinterPrintOptions)
	FeedPaperByLine(lineCount int32)
	FeedPaperByMapModeUnit(distance int32)
}

type IReceiptPrintJob2Vtbl struct {
	win32.IInspectableVtbl
	StampPaper             uintptr
	Print                  uintptr
	FeedPaperByLine        uintptr
	FeedPaperByMapModeUnit uintptr
}

type IReceiptPrintJob2 struct {
	win32.IInspectable
}

func (this *IReceiptPrintJob2) Vtbl() *IReceiptPrintJob2Vtbl {
	return (*IReceiptPrintJob2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IReceiptPrintJob2) StampPaper() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StampPaper, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IReceiptPrintJob2) Print(data string, printOptions *IPosPrinterPrintOptions) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Print, uintptr(unsafe.Pointer(this)), NewHStr(data).Ptr, uintptr(unsafe.Pointer(printOptions)))
	_ = _hr
}

func (this *IReceiptPrintJob2) FeedPaperByLine(lineCount int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FeedPaperByLine, uintptr(unsafe.Pointer(this)), uintptr(lineCount))
	_ = _hr
}

func (this *IReceiptPrintJob2) FeedPaperByMapModeUnit(distance int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FeedPaperByMapModeUnit, uintptr(unsafe.Pointer(this)), uintptr(distance))
	_ = _hr
}

// B8F0B58F-51A8-43FC-9BD5-8DE272A6415B
var IID_IReceiptPrinterCapabilities = syscall.GUID{0xB8F0B58F, 0x51A8, 0x43FC,
	[8]byte{0x9B, 0xD5, 0x8D, 0xE2, 0x72, 0xA6, 0x41, 0x5B}}

type IReceiptPrinterCapabilitiesInterface interface {
	win32.IInspectableInterface
	Get_CanCutPaper() bool
	Get_IsStampSupported() bool
	Get_MarkFeedCapabilities() PosPrinterMarkFeedCapabilities
}

type IReceiptPrinterCapabilitiesVtbl struct {
	win32.IInspectableVtbl
	Get_CanCutPaper          uintptr
	Get_IsStampSupported     uintptr
	Get_MarkFeedCapabilities uintptr
}

type IReceiptPrinterCapabilities struct {
	win32.IInspectable
}

func (this *IReceiptPrinterCapabilities) Vtbl() *IReceiptPrinterCapabilitiesVtbl {
	return (*IReceiptPrinterCapabilitiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IReceiptPrinterCapabilities) Get_CanCutPaper() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanCutPaper, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReceiptPrinterCapabilities) Get_IsStampSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsStampSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReceiptPrinterCapabilities) Get_MarkFeedCapabilities() PosPrinterMarkFeedCapabilities {
	var _result PosPrinterMarkFeedCapabilities
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MarkFeedCapabilities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 20030638-8A2C-55AC-9A7B-7576D8869E99
var IID_IReceiptPrinterCapabilities2 = syscall.GUID{0x20030638, 0x8A2C, 0x55AC,
	[8]byte{0x9A, 0x7B, 0x75, 0x76, 0xD8, 0x86, 0x9E, 0x99}}

type IReceiptPrinterCapabilities2Interface interface {
	win32.IInspectableInterface
	Get_IsReverseVideoSupported() bool
	Get_IsStrikethroughSupported() bool
	Get_IsSuperscriptSupported() bool
	Get_IsSubscriptSupported() bool
	Get_IsReversePaperFeedByLineSupported() bool
	Get_IsReversePaperFeedByMapModeUnitSupported() bool
}

type IReceiptPrinterCapabilities2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsReverseVideoSupported                  uintptr
	Get_IsStrikethroughSupported                 uintptr
	Get_IsSuperscriptSupported                   uintptr
	Get_IsSubscriptSupported                     uintptr
	Get_IsReversePaperFeedByLineSupported        uintptr
	Get_IsReversePaperFeedByMapModeUnitSupported uintptr
}

type IReceiptPrinterCapabilities2 struct {
	win32.IInspectable
}

func (this *IReceiptPrinterCapabilities2) Vtbl() *IReceiptPrinterCapabilities2Vtbl {
	return (*IReceiptPrinterCapabilities2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IReceiptPrinterCapabilities2) Get_IsReverseVideoSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsReverseVideoSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReceiptPrinterCapabilities2) Get_IsStrikethroughSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsStrikethroughSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReceiptPrinterCapabilities2) Get_IsSuperscriptSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSuperscriptSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReceiptPrinterCapabilities2) Get_IsSubscriptSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSubscriptSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReceiptPrinterCapabilities2) Get_IsReversePaperFeedByLineSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsReversePaperFeedByLineSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReceiptPrinterCapabilities2) Get_IsReversePaperFeedByMapModeUnitSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsReversePaperFeedByMapModeUnitSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5D88F95D-6131-5A4B-B7D5-8EF2DA7B4165
var IID_ISlipPrintJob = syscall.GUID{0x5D88F95D, 0x6131, 0x5A4B,
	[8]byte{0xB7, 0xD5, 0x8E, 0xF2, 0xDA, 0x7B, 0x41, 0x65}}

type ISlipPrintJobInterface interface {
	win32.IInspectableInterface
	Print(data string, printOptions *IPosPrinterPrintOptions)
	FeedPaperByLine(lineCount int32)
	FeedPaperByMapModeUnit(distance int32)
}

type ISlipPrintJobVtbl struct {
	win32.IInspectableVtbl
	Print                  uintptr
	FeedPaperByLine        uintptr
	FeedPaperByMapModeUnit uintptr
}

type ISlipPrintJob struct {
	win32.IInspectable
}

func (this *ISlipPrintJob) Vtbl() *ISlipPrintJobVtbl {
	return (*ISlipPrintJobVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISlipPrintJob) Print(data string, printOptions *IPosPrinterPrintOptions) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Print, uintptr(unsafe.Pointer(this)), NewHStr(data).Ptr, uintptr(unsafe.Pointer(printOptions)))
	_ = _hr
}

func (this *ISlipPrintJob) FeedPaperByLine(lineCount int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FeedPaperByLine, uintptr(unsafe.Pointer(this)), uintptr(lineCount))
	_ = _hr
}

func (this *ISlipPrintJob) FeedPaperByMapModeUnit(distance int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FeedPaperByMapModeUnit, uintptr(unsafe.Pointer(this)), uintptr(distance))
	_ = _hr
}

// 99B16399-488C-4157-8AC2-9F57F708D3DB
var IID_ISlipPrinterCapabilities = syscall.GUID{0x99B16399, 0x488C, 0x4157,
	[8]byte{0x8A, 0xC2, 0x9F, 0x57, 0xF7, 0x08, 0xD3, 0xDB}}

type ISlipPrinterCapabilitiesInterface interface {
	win32.IInspectableInterface
	Get_IsFullLengthSupported() bool
	Get_IsBothSidesPrintingSupported() bool
}

type ISlipPrinterCapabilitiesVtbl struct {
	win32.IInspectableVtbl
	Get_IsFullLengthSupported        uintptr
	Get_IsBothSidesPrintingSupported uintptr
}

type ISlipPrinterCapabilities struct {
	win32.IInspectable
}

func (this *ISlipPrinterCapabilities) Vtbl() *ISlipPrinterCapabilitiesVtbl {
	return (*ISlipPrinterCapabilitiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISlipPrinterCapabilities) Get_IsFullLengthSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsFullLengthSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISlipPrinterCapabilities) Get_IsBothSidesPrintingSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsBothSidesPrintingSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6FF89671-2D1A-5000-87C2-B0851BFDF07E
var IID_ISlipPrinterCapabilities2 = syscall.GUID{0x6FF89671, 0x2D1A, 0x5000,
	[8]byte{0x87, 0xC2, 0xB0, 0x85, 0x1B, 0xFD, 0xF0, 0x7E}}

type ISlipPrinterCapabilities2Interface interface {
	win32.IInspectableInterface
	Get_IsReverseVideoSupported() bool
	Get_IsStrikethroughSupported() bool
	Get_IsSuperscriptSupported() bool
	Get_IsSubscriptSupported() bool
	Get_IsReversePaperFeedByLineSupported() bool
	Get_IsReversePaperFeedByMapModeUnitSupported() bool
}

type ISlipPrinterCapabilities2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsReverseVideoSupported                  uintptr
	Get_IsStrikethroughSupported                 uintptr
	Get_IsSuperscriptSupported                   uintptr
	Get_IsSubscriptSupported                     uintptr
	Get_IsReversePaperFeedByLineSupported        uintptr
	Get_IsReversePaperFeedByMapModeUnitSupported uintptr
}

type ISlipPrinterCapabilities2 struct {
	win32.IInspectable
}

func (this *ISlipPrinterCapabilities2) Vtbl() *ISlipPrinterCapabilities2Vtbl {
	return (*ISlipPrinterCapabilities2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISlipPrinterCapabilities2) Get_IsReverseVideoSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsReverseVideoSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISlipPrinterCapabilities2) Get_IsStrikethroughSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsStrikethroughSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISlipPrinterCapabilities2) Get_IsSuperscriptSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSuperscriptSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISlipPrinterCapabilities2) Get_IsSubscriptSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSubscriptSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISlipPrinterCapabilities2) Get_IsReversePaperFeedByLineSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsReversePaperFeedByLineSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISlipPrinterCapabilities2) Get_IsReversePaperFeedByMapModeUnitSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsReversePaperFeedByMapModeUnitSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 2B998C3A-555C-4889-8ED8-C599BB3A712A
var IID_IUnifiedPosErrorData = syscall.GUID{0x2B998C3A, 0x555C, 0x4889,
	[8]byte{0x8E, 0xD8, 0xC5, 0x99, 0xBB, 0x3A, 0x71, 0x2A}}

type IUnifiedPosErrorDataInterface interface {
	win32.IInspectableInterface
	Get_Message() string
	Get_Severity() UnifiedPosErrorSeverity
	Get_Reason() UnifiedPosErrorReason
	Get_ExtendedReason() uint32
}

type IUnifiedPosErrorDataVtbl struct {
	win32.IInspectableVtbl
	Get_Message        uintptr
	Get_Severity       uintptr
	Get_Reason         uintptr
	Get_ExtendedReason uintptr
}

type IUnifiedPosErrorData struct {
	win32.IInspectable
}

func (this *IUnifiedPosErrorData) Vtbl() *IUnifiedPosErrorDataVtbl {
	return (*IUnifiedPosErrorDataVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUnifiedPosErrorData) Get_Message() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Message, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUnifiedPosErrorData) Get_Severity() UnifiedPosErrorSeverity {
	var _result UnifiedPosErrorSeverity
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Severity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUnifiedPosErrorData) Get_Reason() UnifiedPosErrorReason {
	var _result UnifiedPosErrorReason
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reason, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUnifiedPosErrorData) Get_ExtendedReason() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedReason, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4B982551-1FFE-451B-A368-63E0CE465F5A
var IID_IUnifiedPosErrorDataFactory = syscall.GUID{0x4B982551, 0x1FFE, 0x451B,
	[8]byte{0xA3, 0x68, 0x63, 0xE0, 0xCE, 0x46, 0x5F, 0x5A}}

type IUnifiedPosErrorDataFactoryInterface interface {
	win32.IInspectableInterface
	CreateInstance(message string, severity UnifiedPosErrorSeverity, reason UnifiedPosErrorReason, extendedReason uint32) *IUnifiedPosErrorData
}

type IUnifiedPosErrorDataFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateInstance uintptr
}

type IUnifiedPosErrorDataFactory struct {
	win32.IInspectable
}

func (this *IUnifiedPosErrorDataFactory) Vtbl() *IUnifiedPosErrorDataFactoryVtbl {
	return (*IUnifiedPosErrorDataFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUnifiedPosErrorDataFactory) CreateInstance(message string, severity UnifiedPosErrorSeverity, reason UnifiedPosErrorReason, extendedReason uint32) *IUnifiedPosErrorData {
	var _result *IUnifiedPosErrorData
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateInstance, uintptr(unsafe.Pointer(this)), NewHStr(message).Ptr, uintptr(severity), uintptr(reason), uintptr(extendedReason), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type BarcodeSymbologyAttributes struct {
	RtClass
	*IBarcodeSymbologyAttributes
}

type CashDrawer struct {
	RtClass
	*ICashDrawer
}

func NewICashDrawerStatics() *ICashDrawerStatics {
	var p *ICashDrawerStatics
	hs := NewHStr("Windows.Devices.PointOfService.CashDrawer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICashDrawerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewICashDrawerStatics2() *ICashDrawerStatics2 {
	var p *ICashDrawerStatics2
	hs := NewHStr("Windows.Devices.PointOfService.CashDrawer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICashDrawerStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type CashDrawerCapabilities struct {
	RtClass
	*ICashDrawerCapabilities
}

type CashDrawerCloseAlarm struct {
	RtClass
	*ICashDrawerCloseAlarm
}

type CashDrawerClosedEventArgs struct {
	RtClass
	*ICashDrawerEventSourceEventArgs
}

type CashDrawerEventSource struct {
	RtClass
	*ICashDrawerEventSource
}

type CashDrawerOpenedEventArgs struct {
	RtClass
	*ICashDrawerEventSourceEventArgs
}

type CashDrawerStatus struct {
	RtClass
	*ICashDrawerStatus
}

type CashDrawerStatusUpdatedEventArgs struct {
	RtClass
	*ICashDrawerStatusUpdatedEventArgs
}

type ClaimedCashDrawer struct {
	RtClass
	*IClaimedCashDrawer
}

type ClaimedJournalPrinter struct {
	RtClass
	*IClaimedJournalPrinter
}

type ClaimedLineDisplay struct {
	RtClass
	*IClaimedLineDisplay
}

func NewIClaimedLineDisplayStatics() *IClaimedLineDisplayStatics {
	var p *IClaimedLineDisplayStatics
	hs := NewHStr("Windows.Devices.PointOfService.ClaimedLineDisplay")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IClaimedLineDisplayStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type ClaimedPosPrinter struct {
	RtClass
	*IClaimedPosPrinter
}

type ClaimedReceiptPrinter struct {
	RtClass
	*IClaimedReceiptPrinter
}

type ClaimedSlipPrinter struct {
	RtClass
	*IClaimedSlipPrinter
}

type JournalPrintJob struct {
	RtClass
	*IPosPrinterJob
}

type JournalPrinterCapabilities struct {
	RtClass
	*IJournalPrinterCapabilities
}

type LineDisplay struct {
	RtClass
	*ILineDisplay
}

func NewILineDisplayStatics2() *ILineDisplayStatics2 {
	var p *ILineDisplayStatics2
	hs := NewHStr("Windows.Devices.PointOfService.LineDisplay")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ILineDisplayStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewILineDisplayStatics() *ILineDisplayStatics {
	var p *ILineDisplayStatics
	hs := NewHStr("Windows.Devices.PointOfService.LineDisplay")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ILineDisplayStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type LineDisplayAttributes struct {
	RtClass
	*ILineDisplayAttributes
}

type LineDisplayCapabilities struct {
	RtClass
	*ILineDisplayCapabilities
}

type LineDisplayCursor struct {
	RtClass
	*ILineDisplayCursor
}

type LineDisplayCursorAttributes struct {
	RtClass
	*ILineDisplayCursorAttributes
}

type LineDisplayCustomGlyphs struct {
	RtClass
	*ILineDisplayCustomGlyphs
}

type LineDisplayMarquee struct {
	RtClass
	*ILineDisplayMarquee
}

type LineDisplayStatisticsCategorySelector struct {
	RtClass
	*ILineDisplayStatisticsCategorySelector
}

type LineDisplayStatusUpdatedEventArgs struct {
	RtClass
	*ILineDisplayStatusUpdatedEventArgs
}

type LineDisplayStoredBitmap struct {
	RtClass
	*ILineDisplayStoredBitmap
}

type LineDisplayWindow struct {
	RtClass
	*ILineDisplayWindow
}

type PosPrinter struct {
	RtClass
	*IPosPrinter
}

func NewIPosPrinterStatics() *IPosPrinterStatics {
	var p *IPosPrinterStatics
	hs := NewHStr("Windows.Devices.PointOfService.PosPrinter")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPosPrinterStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIPosPrinterStatics2() *IPosPrinterStatics2 {
	var p *IPosPrinterStatics2
	hs := NewHStr("Windows.Devices.PointOfService.PosPrinter")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPosPrinterStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type PosPrinterCapabilities struct {
	RtClass
	*IPosPrinterCapabilities
}

type PosPrinterCharacterSetIds struct {
	RtClass
}

func NewIPosPrinterCharacterSetIdsStatics() *IPosPrinterCharacterSetIdsStatics {
	var p *IPosPrinterCharacterSetIdsStatics
	hs := NewHStr("Windows.Devices.PointOfService.PosPrinterCharacterSetIds")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPosPrinterCharacterSetIdsStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type PosPrinterFontProperty struct {
	RtClass
	*IPosPrinterFontProperty
}

type PosPrinterPrintOptions struct {
	RtClass
	*IPosPrinterPrintOptions
}

func NewPosPrinterPrintOptions() *PosPrinterPrintOptions {
	hs := NewHStr("Windows.Devices.PointOfService.PosPrinterPrintOptions")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &PosPrinterPrintOptions{
		RtClass:                 RtClass{PInspect: p},
		IPosPrinterPrintOptions: (*IPosPrinterPrintOptions)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type PosPrinterReleaseDeviceRequestedEventArgs struct {
	RtClass
	*IPosPrinterReleaseDeviceRequestedEventArgs
}

type PosPrinterStatus struct {
	RtClass
	*IPosPrinterStatus
}

type PosPrinterStatusUpdatedEventArgs struct {
	RtClass
	*IPosPrinterStatusUpdatedEventArgs
}

type ReceiptPrintJob struct {
	RtClass
	*IReceiptPrintJob
}

type ReceiptPrinterCapabilities struct {
	RtClass
	*IReceiptPrinterCapabilities
}

type SlipPrintJob struct {
	RtClass
	*IReceiptOrSlipJob
}

type SlipPrinterCapabilities struct {
	RtClass
	*ISlipPrinterCapabilities
}
