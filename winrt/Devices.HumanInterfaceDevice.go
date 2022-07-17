package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type HidCollectionType int32

const (
	HidCollectionType_Physical      HidCollectionType = 0
	HidCollectionType_Application   HidCollectionType = 1
	HidCollectionType_Logical       HidCollectionType = 2
	HidCollectionType_Report        HidCollectionType = 3
	HidCollectionType_NamedArray    HidCollectionType = 4
	HidCollectionType_UsageSwitch   HidCollectionType = 5
	HidCollectionType_UsageModifier HidCollectionType = 6
	HidCollectionType_Other         HidCollectionType = 7
)

// enum
type HidReportType int32

const (
	HidReportType_Input   HidReportType = 0
	HidReportType_Output  HidReportType = 1
	HidReportType_Feature HidReportType = 2
)

// interfaces

// 524DF48A-3695-408C-BBA2-E2EB5ABFBC20
var IID_IHidBooleanControl = syscall.GUID{0x524DF48A, 0x3695, 0x408C,
	[8]byte{0xBB, 0xA2, 0xE2, 0xEB, 0x5A, 0xBF, 0xBC, 0x20}}

type IHidBooleanControlInterface interface {
	win32.IInspectableInterface
	Get_Id() uint32
	Get_UsagePage() uint16
	Get_UsageId() uint16
	Get_IsActive() bool
	Put_IsActive(value bool)
	Get_ControlDescription() *IHidBooleanControlDescription
}

type IHidBooleanControlVtbl struct {
	win32.IInspectableVtbl
	Get_Id                 uintptr
	Get_UsagePage          uintptr
	Get_UsageId            uintptr
	Get_IsActive           uintptr
	Put_IsActive           uintptr
	Get_ControlDescription uintptr
}

type IHidBooleanControl struct {
	win32.IInspectable
}

func (this *IHidBooleanControl) Vtbl() *IHidBooleanControlVtbl {
	return (*IHidBooleanControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHidBooleanControl) Get_Id() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidBooleanControl) Get_UsagePage() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UsagePage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidBooleanControl) Get_UsageId() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UsageId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidBooleanControl) Get_IsActive() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsActive, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidBooleanControl) Put_IsActive(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsActive, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IHidBooleanControl) Get_ControlDescription() *IHidBooleanControlDescription {
	var _result *IHidBooleanControlDescription
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ControlDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6196E543-29D8-4A2A-8683-849E207BBE31
var IID_IHidBooleanControlDescription = syscall.GUID{0x6196E543, 0x29D8, 0x4A2A,
	[8]byte{0x86, 0x83, 0x84, 0x9E, 0x20, 0x7B, 0xBE, 0x31}}

type IHidBooleanControlDescriptionInterface interface {
	win32.IInspectableInterface
	Get_Id() uint32
	Get_ReportId() uint16
	Get_ReportType() HidReportType
	Get_UsagePage() uint16
	Get_UsageId() uint16
	Get_ParentCollections() *IVectorView[*IHidCollection]
}

type IHidBooleanControlDescriptionVtbl struct {
	win32.IInspectableVtbl
	Get_Id                uintptr
	Get_ReportId          uintptr
	Get_ReportType        uintptr
	Get_UsagePage         uintptr
	Get_UsageId           uintptr
	Get_ParentCollections uintptr
}

type IHidBooleanControlDescription struct {
	win32.IInspectable
}

func (this *IHidBooleanControlDescription) Vtbl() *IHidBooleanControlDescriptionVtbl {
	return (*IHidBooleanControlDescriptionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHidBooleanControlDescription) Get_Id() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidBooleanControlDescription) Get_ReportId() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidBooleanControlDescription) Get_ReportType() HidReportType {
	var _result HidReportType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidBooleanControlDescription) Get_UsagePage() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UsagePage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidBooleanControlDescription) Get_UsageId() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UsageId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidBooleanControlDescription) Get_ParentCollections() *IVectorView[*IHidCollection] {
	var _result *IVectorView[*IHidCollection]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ParentCollections, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C8EED2EA-8A77-4C36-AA00-5FF0449D3E73
var IID_IHidBooleanControlDescription2 = syscall.GUID{0xC8EED2EA, 0x8A77, 0x4C36,
	[8]byte{0xAA, 0x00, 0x5F, 0xF0, 0x44, 0x9D, 0x3E, 0x73}}

type IHidBooleanControlDescription2Interface interface {
	win32.IInspectableInterface
	Get_IsAbsolute() bool
}

type IHidBooleanControlDescription2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsAbsolute uintptr
}

type IHidBooleanControlDescription2 struct {
	win32.IInspectable
}

func (this *IHidBooleanControlDescription2) Vtbl() *IHidBooleanControlDescription2Vtbl {
	return (*IHidBooleanControlDescription2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHidBooleanControlDescription2) Get_IsAbsolute() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAbsolute, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7189F5A3-32F1-46E3-BEFD-44D2663B7E6A
var IID_IHidCollection = syscall.GUID{0x7189F5A3, 0x32F1, 0x46E3,
	[8]byte{0xBE, 0xFD, 0x44, 0xD2, 0x66, 0x3B, 0x7E, 0x6A}}

type IHidCollectionInterface interface {
	win32.IInspectableInterface
	Get_Id() uint32
	Get_Type() HidCollectionType
	Get_UsagePage() uint32
	Get_UsageId() uint32
}

type IHidCollectionVtbl struct {
	win32.IInspectableVtbl
	Get_Id        uintptr
	Get_Type      uintptr
	Get_UsagePage uintptr
	Get_UsageId   uintptr
}

type IHidCollection struct {
	win32.IInspectable
}

func (this *IHidCollection) Vtbl() *IHidCollectionVtbl {
	return (*IHidCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHidCollection) Get_Id() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidCollection) Get_Type() HidCollectionType {
	var _result HidCollectionType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Type, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidCollection) Get_UsagePage() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UsagePage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidCollection) Get_UsageId() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UsageId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5F8A14E7-2200-432E-95DA-D09B87D574A8
var IID_IHidDevice = syscall.GUID{0x5F8A14E7, 0x2200, 0x432E,
	[8]byte{0x95, 0xDA, 0xD0, 0x9B, 0x87, 0xD5, 0x74, 0xA8}}

type IHidDeviceInterface interface {
	win32.IInspectableInterface
	Get_VendorId() uint16
	Get_ProductId() uint16
	Get_Version() uint16
	Get_UsagePage() uint16
	Get_UsageId() uint16
	GetInputReportAsync() *IAsyncOperation[*IHidInputReport]
	GetInputReportByIdAsync(reportId uint16) *IAsyncOperation[*IHidInputReport]
	GetFeatureReportAsync() *IAsyncOperation[*IHidFeatureReport]
	GetFeatureReportByIdAsync(reportId uint16) *IAsyncOperation[*IHidFeatureReport]
	CreateOutputReport() *IHidOutputReport
	CreateOutputReportById(reportId uint16) *IHidOutputReport
	CreateFeatureReport() *IHidFeatureReport
	CreateFeatureReportById(reportId uint16) *IHidFeatureReport
	SendOutputReportAsync(outputReport *IHidOutputReport) *IAsyncOperation[uint32]
	SendFeatureReportAsync(featureReport *IHidFeatureReport) *IAsyncOperation[uint32]
	GetBooleanControlDescriptions(reportType HidReportType, usagePage uint16, usageId uint16) *IVectorView[*IHidBooleanControlDescription]
	GetNumericControlDescriptions(reportType HidReportType, usagePage uint16, usageId uint16) *IVectorView[*IHidNumericControlDescription]
	Add_InputReportReceived(reportHandler TypedEventHandler[*IHidDevice, *IHidInputReportReceivedEventArgs]) EventRegistrationToken
	Remove_InputReportReceived(token EventRegistrationToken)
}

type IHidDeviceVtbl struct {
	win32.IInspectableVtbl
	Get_VendorId                  uintptr
	Get_ProductId                 uintptr
	Get_Version                   uintptr
	Get_UsagePage                 uintptr
	Get_UsageId                   uintptr
	GetInputReportAsync           uintptr
	GetInputReportByIdAsync       uintptr
	GetFeatureReportAsync         uintptr
	GetFeatureReportByIdAsync     uintptr
	CreateOutputReport            uintptr
	CreateOutputReportById        uintptr
	CreateFeatureReport           uintptr
	CreateFeatureReportById       uintptr
	SendOutputReportAsync         uintptr
	SendFeatureReportAsync        uintptr
	GetBooleanControlDescriptions uintptr
	GetNumericControlDescriptions uintptr
	Add_InputReportReceived       uintptr
	Remove_InputReportReceived    uintptr
}

type IHidDevice struct {
	win32.IInspectable
}

func (this *IHidDevice) Vtbl() *IHidDeviceVtbl {
	return (*IHidDeviceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHidDevice) Get_VendorId() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VendorId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidDevice) Get_ProductId() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProductId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidDevice) Get_Version() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Version, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidDevice) Get_UsagePage() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UsagePage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidDevice) Get_UsageId() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UsageId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidDevice) GetInputReportAsync() *IAsyncOperation[*IHidInputReport] {
	var _result *IAsyncOperation[*IHidInputReport]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetInputReportAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidDevice) GetInputReportByIdAsync(reportId uint16) *IAsyncOperation[*IHidInputReport] {
	var _result *IAsyncOperation[*IHidInputReport]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetInputReportByIdAsync, uintptr(unsafe.Pointer(this)), uintptr(reportId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidDevice) GetFeatureReportAsync() *IAsyncOperation[*IHidFeatureReport] {
	var _result *IAsyncOperation[*IHidFeatureReport]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFeatureReportAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidDevice) GetFeatureReportByIdAsync(reportId uint16) *IAsyncOperation[*IHidFeatureReport] {
	var _result *IAsyncOperation[*IHidFeatureReport]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFeatureReportByIdAsync, uintptr(unsafe.Pointer(this)), uintptr(reportId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidDevice) CreateOutputReport() *IHidOutputReport {
	var _result *IHidOutputReport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateOutputReport, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidDevice) CreateOutputReportById(reportId uint16) *IHidOutputReport {
	var _result *IHidOutputReport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateOutputReportById, uintptr(unsafe.Pointer(this)), uintptr(reportId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidDevice) CreateFeatureReport() *IHidFeatureReport {
	var _result *IHidFeatureReport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFeatureReport, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidDevice) CreateFeatureReportById(reportId uint16) *IHidFeatureReport {
	var _result *IHidFeatureReport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFeatureReportById, uintptr(unsafe.Pointer(this)), uintptr(reportId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidDevice) SendOutputReportAsync(outputReport *IHidOutputReport) *IAsyncOperation[uint32] {
	var _result *IAsyncOperation[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendOutputReportAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(outputReport)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidDevice) SendFeatureReportAsync(featureReport *IHidFeatureReport) *IAsyncOperation[uint32] {
	var _result *IAsyncOperation[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendFeatureReportAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(featureReport)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidDevice) GetBooleanControlDescriptions(reportType HidReportType, usagePage uint16, usageId uint16) *IVectorView[*IHidBooleanControlDescription] {
	var _result *IVectorView[*IHidBooleanControlDescription]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetBooleanControlDescriptions, uintptr(unsafe.Pointer(this)), uintptr(reportType), uintptr(usagePage), uintptr(usageId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidDevice) GetNumericControlDescriptions(reportType HidReportType, usagePage uint16, usageId uint16) *IVectorView[*IHidNumericControlDescription] {
	var _result *IVectorView[*IHidNumericControlDescription]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNumericControlDescriptions, uintptr(unsafe.Pointer(this)), uintptr(reportType), uintptr(usagePage), uintptr(usageId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidDevice) Add_InputReportReceived(reportHandler TypedEventHandler[*IHidDevice, *IHidInputReportReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_InputReportReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(reportHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidDevice) Remove_InputReportReceived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_InputReportReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 9E5981E4-9856-418C-9F73-77DE0CD85754
var IID_IHidDeviceStatics = syscall.GUID{0x9E5981E4, 0x9856, 0x418C,
	[8]byte{0x9F, 0x73, 0x77, 0xDE, 0x0C, 0xD8, 0x57, 0x54}}

type IHidDeviceStaticsInterface interface {
	win32.IInspectableInterface
	GetDeviceSelector(usagePage uint16, usageId uint16) string
	GetDeviceSelectorVidPid(usagePage uint16, usageId uint16, vendorId uint16, productId uint16) string
	FromIdAsync(deviceId string, accessMode FileAccessMode) *IAsyncOperation[*IHidDevice]
}

type IHidDeviceStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector       uintptr
	GetDeviceSelectorVidPid uintptr
	FromIdAsync             uintptr
}

type IHidDeviceStatics struct {
	win32.IInspectable
}

func (this *IHidDeviceStatics) Vtbl() *IHidDeviceStaticsVtbl {
	return (*IHidDeviceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHidDeviceStatics) GetDeviceSelector(usagePage uint16, usageId uint16) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(usagePage), uintptr(usageId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHidDeviceStatics) GetDeviceSelectorVidPid(usagePage uint16, usageId uint16, vendorId uint16, productId uint16) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorVidPid, uintptr(unsafe.Pointer(this)), uintptr(usagePage), uintptr(usageId), uintptr(vendorId), uintptr(productId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHidDeviceStatics) FromIdAsync(deviceId string, accessMode FileAccessMode) *IAsyncOperation[*IHidDevice] {
	var _result *IAsyncOperation[*IHidDevice]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(accessMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 841D9B79-5AE5-46E3-82EF-1FEC5C8942F4
var IID_IHidFeatureReport = syscall.GUID{0x841D9B79, 0x5AE5, 0x46E3,
	[8]byte{0x82, 0xEF, 0x1F, 0xEC, 0x5C, 0x89, 0x42, 0xF4}}

type IHidFeatureReportInterface interface {
	win32.IInspectableInterface
	Get_Id() uint16
	Get_Data() *IBuffer
	Put_Data(value *IBuffer)
	GetBooleanControl(usagePage uint16, usageId uint16) *IHidBooleanControl
	GetBooleanControlByDescription(controlDescription *IHidBooleanControlDescription) *IHidBooleanControl
	GetNumericControl(usagePage uint16, usageId uint16) *IHidNumericControl
	GetNumericControlByDescription(controlDescription *IHidNumericControlDescription) *IHidNumericControl
}

type IHidFeatureReportVtbl struct {
	win32.IInspectableVtbl
	Get_Id                         uintptr
	Get_Data                       uintptr
	Put_Data                       uintptr
	GetBooleanControl              uintptr
	GetBooleanControlByDescription uintptr
	GetNumericControl              uintptr
	GetNumericControlByDescription uintptr
}

type IHidFeatureReport struct {
	win32.IInspectable
}

func (this *IHidFeatureReport) Vtbl() *IHidFeatureReportVtbl {
	return (*IHidFeatureReportVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHidFeatureReport) Get_Id() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidFeatureReport) Get_Data() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidFeatureReport) Put_Data(value *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHidFeatureReport) GetBooleanControl(usagePage uint16, usageId uint16) *IHidBooleanControl {
	var _result *IHidBooleanControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetBooleanControl, uintptr(unsafe.Pointer(this)), uintptr(usagePage), uintptr(usageId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidFeatureReport) GetBooleanControlByDescription(controlDescription *IHidBooleanControlDescription) *IHidBooleanControl {
	var _result *IHidBooleanControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetBooleanControlByDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(controlDescription)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidFeatureReport) GetNumericControl(usagePage uint16, usageId uint16) *IHidNumericControl {
	var _result *IHidNumericControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNumericControl, uintptr(unsafe.Pointer(this)), uintptr(usagePage), uintptr(usageId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidFeatureReport) GetNumericControlByDescription(controlDescription *IHidNumericControlDescription) *IHidNumericControl {
	var _result *IHidNumericControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNumericControlByDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(controlDescription)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C35D0E50-F7E7-4E8D-B23E-CABBE56B90E9
var IID_IHidInputReport = syscall.GUID{0xC35D0E50, 0xF7E7, 0x4E8D,
	[8]byte{0xB2, 0x3E, 0xCA, 0xBB, 0xE5, 0x6B, 0x90, 0xE9}}

type IHidInputReportInterface interface {
	win32.IInspectableInterface
	Get_Id() uint16
	Get_Data() *IBuffer
	Get_ActivatedBooleanControls() *IVectorView[*IHidBooleanControl]
	Get_TransitionedBooleanControls() *IVectorView[*IHidBooleanControl]
	GetBooleanControl(usagePage uint16, usageId uint16) *IHidBooleanControl
	GetBooleanControlByDescription(controlDescription *IHidBooleanControlDescription) *IHidBooleanControl
	GetNumericControl(usagePage uint16, usageId uint16) *IHidNumericControl
	GetNumericControlByDescription(controlDescription *IHidNumericControlDescription) *IHidNumericControl
}

type IHidInputReportVtbl struct {
	win32.IInspectableVtbl
	Get_Id                          uintptr
	Get_Data                        uintptr
	Get_ActivatedBooleanControls    uintptr
	Get_TransitionedBooleanControls uintptr
	GetBooleanControl               uintptr
	GetBooleanControlByDescription  uintptr
	GetNumericControl               uintptr
	GetNumericControlByDescription  uintptr
}

type IHidInputReport struct {
	win32.IInspectable
}

func (this *IHidInputReport) Vtbl() *IHidInputReportVtbl {
	return (*IHidInputReportVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHidInputReport) Get_Id() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidInputReport) Get_Data() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidInputReport) Get_ActivatedBooleanControls() *IVectorView[*IHidBooleanControl] {
	var _result *IVectorView[*IHidBooleanControl]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActivatedBooleanControls, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidInputReport) Get_TransitionedBooleanControls() *IVectorView[*IHidBooleanControl] {
	var _result *IVectorView[*IHidBooleanControl]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TransitionedBooleanControls, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidInputReport) GetBooleanControl(usagePage uint16, usageId uint16) *IHidBooleanControl {
	var _result *IHidBooleanControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetBooleanControl, uintptr(unsafe.Pointer(this)), uintptr(usagePage), uintptr(usageId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidInputReport) GetBooleanControlByDescription(controlDescription *IHidBooleanControlDescription) *IHidBooleanControl {
	var _result *IHidBooleanControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetBooleanControlByDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(controlDescription)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidInputReport) GetNumericControl(usagePage uint16, usageId uint16) *IHidNumericControl {
	var _result *IHidNumericControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNumericControl, uintptr(unsafe.Pointer(this)), uintptr(usagePage), uintptr(usageId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidInputReport) GetNumericControlByDescription(controlDescription *IHidNumericControlDescription) *IHidNumericControl {
	var _result *IHidNumericControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNumericControlByDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(controlDescription)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7059C5CB-59B2-4DC2-985C-0ADC6136FA2D
var IID_IHidInputReportReceivedEventArgs = syscall.GUID{0x7059C5CB, 0x59B2, 0x4DC2,
	[8]byte{0x98, 0x5C, 0x0A, 0xDC, 0x61, 0x36, 0xFA, 0x2D}}

type IHidInputReportReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Report() *IHidInputReport
}

type IHidInputReportReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Report uintptr
}

type IHidInputReportReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IHidInputReportReceivedEventArgs) Vtbl() *IHidInputReportReceivedEventArgsVtbl {
	return (*IHidInputReportReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHidInputReportReceivedEventArgs) Get_Report() *IHidInputReport {
	var _result *IHidInputReport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Report, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E38A12A5-35A7-4B75-89C8-FB1F28B10823
var IID_IHidNumericControl = syscall.GUID{0xE38A12A5, 0x35A7, 0x4B75,
	[8]byte{0x89, 0xC8, 0xFB, 0x1F, 0x28, 0xB1, 0x08, 0x23}}

type IHidNumericControlInterface interface {
	win32.IInspectableInterface
	Get_Id() uint32
	Get_IsGrouped() bool
	Get_UsagePage() uint16
	Get_UsageId() uint16
	Get_Value() int64
	Put_Value(value int64)
	Get_ScaledValue() int64
	Put_ScaledValue(value int64)
	Get_ControlDescription() *IHidNumericControlDescription
}

type IHidNumericControlVtbl struct {
	win32.IInspectableVtbl
	Get_Id                 uintptr
	Get_IsGrouped          uintptr
	Get_UsagePage          uintptr
	Get_UsageId            uintptr
	Get_Value              uintptr
	Put_Value              uintptr
	Get_ScaledValue        uintptr
	Put_ScaledValue        uintptr
	Get_ControlDescription uintptr
}

type IHidNumericControl struct {
	win32.IInspectable
}

func (this *IHidNumericControl) Vtbl() *IHidNumericControlVtbl {
	return (*IHidNumericControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHidNumericControl) Get_Id() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidNumericControl) Get_IsGrouped() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsGrouped, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidNumericControl) Get_UsagePage() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UsagePage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidNumericControl) Get_UsageId() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UsageId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidNumericControl) Get_Value() int64 {
	var _result int64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidNumericControl) Put_Value(value int64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Value, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IHidNumericControl) Get_ScaledValue() int64 {
	var _result int64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScaledValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidNumericControl) Put_ScaledValue(value int64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ScaledValue, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IHidNumericControl) Get_ControlDescription() *IHidNumericControlDescription {
	var _result *IHidNumericControlDescription
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ControlDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 638D5E86-1D97-4C75-927F-5FF58BA05E32
var IID_IHidNumericControlDescription = syscall.GUID{0x638D5E86, 0x1D97, 0x4C75,
	[8]byte{0x92, 0x7F, 0x5F, 0xF5, 0x8B, 0xA0, 0x5E, 0x32}}

type IHidNumericControlDescriptionInterface interface {
	win32.IInspectableInterface
	Get_Id() uint32
	Get_ReportId() uint16
	Get_ReportType() HidReportType
	Get_ReportSize() uint32
	Get_ReportCount() uint32
	Get_UsagePage() uint16
	Get_UsageId() uint16
	Get_LogicalMinimum() int32
	Get_LogicalMaximum() int32
	Get_PhysicalMinimum() int32
	Get_PhysicalMaximum() int32
	Get_UnitExponent() uint32
	Get_Unit() uint32
	Get_IsAbsolute() bool
	Get_HasNull() bool
	Get_ParentCollections() *IVectorView[*IHidCollection]
}

type IHidNumericControlDescriptionVtbl struct {
	win32.IInspectableVtbl
	Get_Id                uintptr
	Get_ReportId          uintptr
	Get_ReportType        uintptr
	Get_ReportSize        uintptr
	Get_ReportCount       uintptr
	Get_UsagePage         uintptr
	Get_UsageId           uintptr
	Get_LogicalMinimum    uintptr
	Get_LogicalMaximum    uintptr
	Get_PhysicalMinimum   uintptr
	Get_PhysicalMaximum   uintptr
	Get_UnitExponent      uintptr
	Get_Unit              uintptr
	Get_IsAbsolute        uintptr
	Get_HasNull           uintptr
	Get_ParentCollections uintptr
}

type IHidNumericControlDescription struct {
	win32.IInspectable
}

func (this *IHidNumericControlDescription) Vtbl() *IHidNumericControlDescriptionVtbl {
	return (*IHidNumericControlDescriptionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHidNumericControlDescription) Get_Id() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidNumericControlDescription) Get_ReportId() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidNumericControlDescription) Get_ReportType() HidReportType {
	var _result HidReportType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidNumericControlDescription) Get_ReportSize() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidNumericControlDescription) Get_ReportCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidNumericControlDescription) Get_UsagePage() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UsagePage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidNumericControlDescription) Get_UsageId() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UsageId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidNumericControlDescription) Get_LogicalMinimum() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LogicalMinimum, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidNumericControlDescription) Get_LogicalMaximum() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LogicalMaximum, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidNumericControlDescription) Get_PhysicalMinimum() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhysicalMinimum, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidNumericControlDescription) Get_PhysicalMaximum() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhysicalMaximum, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidNumericControlDescription) Get_UnitExponent() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UnitExponent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidNumericControlDescription) Get_Unit() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Unit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidNumericControlDescription) Get_IsAbsolute() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAbsolute, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidNumericControlDescription) Get_HasNull() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasNull, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidNumericControlDescription) Get_ParentCollections() *IVectorView[*IHidCollection] {
	var _result *IVectorView[*IHidCollection]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ParentCollections, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 62CB2544-C896-4463-93C1-DF9DB053C450
var IID_IHidOutputReport = syscall.GUID{0x62CB2544, 0xC896, 0x4463,
	[8]byte{0x93, 0xC1, 0xDF, 0x9D, 0xB0, 0x53, 0xC4, 0x50}}

type IHidOutputReportInterface interface {
	win32.IInspectableInterface
	Get_Id() uint16
	Get_Data() *IBuffer
	Put_Data(value *IBuffer)
	GetBooleanControl(usagePage uint16, usageId uint16) *IHidBooleanControl
	GetBooleanControlByDescription(controlDescription *IHidBooleanControlDescription) *IHidBooleanControl
	GetNumericControl(usagePage uint16, usageId uint16) *IHidNumericControl
	GetNumericControlByDescription(controlDescription *IHidNumericControlDescription) *IHidNumericControl
}

type IHidOutputReportVtbl struct {
	win32.IInspectableVtbl
	Get_Id                         uintptr
	Get_Data                       uintptr
	Put_Data                       uintptr
	GetBooleanControl              uintptr
	GetBooleanControlByDescription uintptr
	GetNumericControl              uintptr
	GetNumericControlByDescription uintptr
}

type IHidOutputReport struct {
	win32.IInspectable
}

func (this *IHidOutputReport) Vtbl() *IHidOutputReportVtbl {
	return (*IHidOutputReportVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHidOutputReport) Get_Id() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHidOutputReport) Get_Data() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidOutputReport) Put_Data(value *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHidOutputReport) GetBooleanControl(usagePage uint16, usageId uint16) *IHidBooleanControl {
	var _result *IHidBooleanControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetBooleanControl, uintptr(unsafe.Pointer(this)), uintptr(usagePage), uintptr(usageId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidOutputReport) GetBooleanControlByDescription(controlDescription *IHidBooleanControlDescription) *IHidBooleanControl {
	var _result *IHidBooleanControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetBooleanControlByDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(controlDescription)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidOutputReport) GetNumericControl(usagePage uint16, usageId uint16) *IHidNumericControl {
	var _result *IHidNumericControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNumericControl, uintptr(unsafe.Pointer(this)), uintptr(usagePage), uintptr(usageId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHidOutputReport) GetNumericControlByDescription(controlDescription *IHidNumericControlDescription) *IHidNumericControl {
	var _result *IHidNumericControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNumericControlByDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(controlDescription)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type HidBooleanControl struct {
	RtClass
	*IHidBooleanControl
}

type HidBooleanControlDescription struct {
	RtClass
	*IHidBooleanControlDescription
}

type HidCollection struct {
	RtClass
	*IHidCollection
}

type HidDevice struct {
	RtClass
	*IHidDevice
}

func NewIHidDeviceStatics() *IHidDeviceStatics {
	var p *IHidDeviceStatics
	hs := NewHStr("Windows.Devices.HumanInterfaceDevice.HidDevice")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHidDeviceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type HidFeatureReport struct {
	RtClass
	*IHidFeatureReport
}

type HidInputReport struct {
	RtClass
	*IHidInputReport
}

type HidInputReportReceivedEventArgs struct {
	RtClass
	*IHidInputReportReceivedEventArgs
}

type HidNumericControl struct {
	RtClass
	*IHidNumericControl
}

type HidNumericControlDescription struct {
	RtClass
	*IHidNumericControlDescription
}

type HidOutputReport struct {
	RtClass
	*IHidOutputReport
}
