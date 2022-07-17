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
type UsbControlRecipient int32

const (
	UsbControlRecipient_Device             UsbControlRecipient = 0
	UsbControlRecipient_SpecifiedInterface UsbControlRecipient = 1
	UsbControlRecipient_Endpoint           UsbControlRecipient = 2
	UsbControlRecipient_Other              UsbControlRecipient = 3
	UsbControlRecipient_DefaultInterface   UsbControlRecipient = 4
)

// enum
type UsbControlTransferType int32

const (
	UsbControlTransferType_Standard UsbControlTransferType = 0
	UsbControlTransferType_Class    UsbControlTransferType = 1
	UsbControlTransferType_Vendor   UsbControlTransferType = 2
)

// enum
type UsbEndpointType int32

const (
	UsbEndpointType_Control     UsbEndpointType = 0
	UsbEndpointType_Isochronous UsbEndpointType = 1
	UsbEndpointType_Bulk        UsbEndpointType = 2
	UsbEndpointType_Interrupt   UsbEndpointType = 3
)

// enum
// flags
type UsbReadOptions uint32

const (
	UsbReadOptions_None                              UsbReadOptions = 0
	UsbReadOptions_AutoClearStall                    UsbReadOptions = 1
	UsbReadOptions_OverrideAutomaticBufferManagement UsbReadOptions = 2
	UsbReadOptions_IgnoreShortPacket                 UsbReadOptions = 4
	UsbReadOptions_AllowPartialReads                 UsbReadOptions = 8
)

// enum
type UsbTransferDirection int32

const (
	UsbTransferDirection_Out UsbTransferDirection = 0
	UsbTransferDirection_In  UsbTransferDirection = 1
)

// enum
// flags
type UsbWriteOptions uint32

const (
	UsbWriteOptions_None                 UsbWriteOptions = 0
	UsbWriteOptions_AutoClearStall       UsbWriteOptions = 1
	UsbWriteOptions_ShortPacketTerminate UsbWriteOptions = 2
)

// interfaces

// 3C6E4846-06CF-42A9-9DC2-971C1B14B6E3
var IID_IUsbBulkInEndpointDescriptor = syscall.GUID{0x3C6E4846, 0x06CF, 0x42A9,
	[8]byte{0x9D, 0xC2, 0x97, 0x1C, 0x1B, 0x14, 0xB6, 0xE3}}

type IUsbBulkInEndpointDescriptorInterface interface {
	win32.IInspectableInterface
	Get_MaxPacketSize() uint32
	Get_EndpointNumber() byte
	Get_Pipe() *IUsbBulkInPipe
}

type IUsbBulkInEndpointDescriptorVtbl struct {
	win32.IInspectableVtbl
	Get_MaxPacketSize  uintptr
	Get_EndpointNumber uintptr
	Get_Pipe           uintptr
}

type IUsbBulkInEndpointDescriptor struct {
	win32.IInspectable
}

func (this *IUsbBulkInEndpointDescriptor) Vtbl() *IUsbBulkInEndpointDescriptorVtbl {
	return (*IUsbBulkInEndpointDescriptorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbBulkInEndpointDescriptor) Get_MaxPacketSize() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxPacketSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbBulkInEndpointDescriptor) Get_EndpointNumber() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EndpointNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbBulkInEndpointDescriptor) Get_Pipe() *IUsbBulkInPipe {
	var _result *IUsbBulkInPipe
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Pipe, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F01D2D3B-4548-4D50-B326-D82CDABE1220
var IID_IUsbBulkInPipe = syscall.GUID{0xF01D2D3B, 0x4548, 0x4D50,
	[8]byte{0xB3, 0x26, 0xD8, 0x2C, 0xDA, 0xBE, 0x12, 0x20}}

type IUsbBulkInPipeInterface interface {
	win32.IInspectableInterface
	Get_MaxTransferSizeBytes() uint32
	Get_EndpointDescriptor() *IUsbBulkInEndpointDescriptor
	ClearStallAsync() *IAsyncAction
	Put_ReadOptions(value UsbReadOptions)
	Get_ReadOptions() UsbReadOptions
	FlushBuffer()
	Get_InputStream() *IInputStream
}

type IUsbBulkInPipeVtbl struct {
	win32.IInspectableVtbl
	Get_MaxTransferSizeBytes uintptr
	Get_EndpointDescriptor   uintptr
	ClearStallAsync          uintptr
	Put_ReadOptions          uintptr
	Get_ReadOptions          uintptr
	FlushBuffer              uintptr
	Get_InputStream          uintptr
}

type IUsbBulkInPipe struct {
	win32.IInspectable
}

func (this *IUsbBulkInPipe) Vtbl() *IUsbBulkInPipeVtbl {
	return (*IUsbBulkInPipeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbBulkInPipe) Get_MaxTransferSizeBytes() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxTransferSizeBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbBulkInPipe) Get_EndpointDescriptor() *IUsbBulkInEndpointDescriptor {
	var _result *IUsbBulkInEndpointDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EndpointDescriptor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbBulkInPipe) ClearStallAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearStallAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbBulkInPipe) Put_ReadOptions(value UsbReadOptions) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReadOptions, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IUsbBulkInPipe) Get_ReadOptions() UsbReadOptions {
	var _result UsbReadOptions
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReadOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbBulkInPipe) FlushBuffer() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FlushBuffer, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IUsbBulkInPipe) Get_InputStream() *IInputStream {
	var _result *IInputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InputStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2820847A-FFEE-4F60-9BE1-956CAC3ECB65
var IID_IUsbBulkOutEndpointDescriptor = syscall.GUID{0x2820847A, 0xFFEE, 0x4F60,
	[8]byte{0x9B, 0xE1, 0x95, 0x6C, 0xAC, 0x3E, 0xCB, 0x65}}

type IUsbBulkOutEndpointDescriptorInterface interface {
	win32.IInspectableInterface
	Get_MaxPacketSize() uint32
	Get_EndpointNumber() byte
	Get_Pipe() *IUsbBulkOutPipe
}

type IUsbBulkOutEndpointDescriptorVtbl struct {
	win32.IInspectableVtbl
	Get_MaxPacketSize  uintptr
	Get_EndpointNumber uintptr
	Get_Pipe           uintptr
}

type IUsbBulkOutEndpointDescriptor struct {
	win32.IInspectable
}

func (this *IUsbBulkOutEndpointDescriptor) Vtbl() *IUsbBulkOutEndpointDescriptorVtbl {
	return (*IUsbBulkOutEndpointDescriptorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbBulkOutEndpointDescriptor) Get_MaxPacketSize() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxPacketSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbBulkOutEndpointDescriptor) Get_EndpointNumber() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EndpointNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbBulkOutEndpointDescriptor) Get_Pipe() *IUsbBulkOutPipe {
	var _result *IUsbBulkOutPipe
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Pipe, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A8E9EE6E-0115-45AA-8B21-37B225BCCEE7
var IID_IUsbBulkOutPipe = syscall.GUID{0xA8E9EE6E, 0x0115, 0x45AA,
	[8]byte{0x8B, 0x21, 0x37, 0xB2, 0x25, 0xBC, 0xCE, 0xE7}}

type IUsbBulkOutPipeInterface interface {
	win32.IInspectableInterface
	Get_EndpointDescriptor() *IUsbBulkOutEndpointDescriptor
	ClearStallAsync() *IAsyncAction
	Put_WriteOptions(value UsbWriteOptions)
	Get_WriteOptions() UsbWriteOptions
	Get_OutputStream() *IOutputStream
}

type IUsbBulkOutPipeVtbl struct {
	win32.IInspectableVtbl
	Get_EndpointDescriptor uintptr
	ClearStallAsync        uintptr
	Put_WriteOptions       uintptr
	Get_WriteOptions       uintptr
	Get_OutputStream       uintptr
}

type IUsbBulkOutPipe struct {
	win32.IInspectable
}

func (this *IUsbBulkOutPipe) Vtbl() *IUsbBulkOutPipeVtbl {
	return (*IUsbBulkOutPipeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbBulkOutPipe) Get_EndpointDescriptor() *IUsbBulkOutEndpointDescriptor {
	var _result *IUsbBulkOutEndpointDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EndpointDescriptor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbBulkOutPipe) ClearStallAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearStallAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbBulkOutPipe) Put_WriteOptions(value UsbWriteOptions) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_WriteOptions, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IUsbBulkOutPipe) Get_WriteOptions() UsbWriteOptions {
	var _result UsbWriteOptions
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WriteOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbBulkOutPipe) Get_OutputStream() *IOutputStream {
	var _result *IOutputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutputStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 68177429-36A9-46D7-B873-FC689251EC30
var IID_IUsbConfiguration = syscall.GUID{0x68177429, 0x36A9, 0x46D7,
	[8]byte{0xB8, 0x73, 0xFC, 0x68, 0x92, 0x51, 0xEC, 0x30}}

type IUsbConfigurationInterface interface {
	win32.IInspectableInterface
	Get_UsbInterfaces() *IVectorView[*IUsbInterface]
	Get_ConfigurationDescriptor() *IUsbConfigurationDescriptor
	Get_Descriptors() *IVectorView[*IUsbDescriptor]
}

type IUsbConfigurationVtbl struct {
	win32.IInspectableVtbl
	Get_UsbInterfaces           uintptr
	Get_ConfigurationDescriptor uintptr
	Get_Descriptors             uintptr
}

type IUsbConfiguration struct {
	win32.IInspectable
}

func (this *IUsbConfiguration) Vtbl() *IUsbConfigurationVtbl {
	return (*IUsbConfigurationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbConfiguration) Get_UsbInterfaces() *IVectorView[*IUsbInterface] {
	var _result *IVectorView[*IUsbInterface]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UsbInterfaces, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbConfiguration) Get_ConfigurationDescriptor() *IUsbConfigurationDescriptor {
	var _result *IUsbConfigurationDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConfigurationDescriptor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbConfiguration) Get_Descriptors() *IVectorView[*IUsbDescriptor] {
	var _result *IVectorView[*IUsbDescriptor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Descriptors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F2176D92-B442-407A-8207-7D646C0385F3
var IID_IUsbConfigurationDescriptor = syscall.GUID{0xF2176D92, 0xB442, 0x407A,
	[8]byte{0x82, 0x07, 0x7D, 0x64, 0x6C, 0x03, 0x85, 0xF3}}

type IUsbConfigurationDescriptorInterface interface {
	win32.IInspectableInterface
	Get_ConfigurationValue() byte
	Get_MaxPowerMilliamps() uint32
	Get_SelfPowered() bool
	Get_RemoteWakeup() bool
}

type IUsbConfigurationDescriptorVtbl struct {
	win32.IInspectableVtbl
	Get_ConfigurationValue uintptr
	Get_MaxPowerMilliamps  uintptr
	Get_SelfPowered        uintptr
	Get_RemoteWakeup       uintptr
}

type IUsbConfigurationDescriptor struct {
	win32.IInspectable
}

func (this *IUsbConfigurationDescriptor) Vtbl() *IUsbConfigurationDescriptorVtbl {
	return (*IUsbConfigurationDescriptorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbConfigurationDescriptor) Get_ConfigurationValue() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConfigurationValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbConfigurationDescriptor) Get_MaxPowerMilliamps() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxPowerMilliamps, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbConfigurationDescriptor) Get_SelfPowered() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelfPowered, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbConfigurationDescriptor) Get_RemoteWakeup() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemoteWakeup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 424CED93-E740-40A1-92BD-DA120EA04914
var IID_IUsbConfigurationDescriptorStatics = syscall.GUID{0x424CED93, 0xE740, 0x40A1,
	[8]byte{0x92, 0xBD, 0xDA, 0x12, 0x0E, 0xA0, 0x49, 0x14}}

type IUsbConfigurationDescriptorStaticsInterface interface {
	win32.IInspectableInterface
	TryParse(descriptor *IUsbDescriptor, parsed *IUsbConfigurationDescriptor) bool
	Parse(descriptor *IUsbDescriptor) *IUsbConfigurationDescriptor
}

type IUsbConfigurationDescriptorStaticsVtbl struct {
	win32.IInspectableVtbl
	TryParse uintptr
	Parse    uintptr
}

type IUsbConfigurationDescriptorStatics struct {
	win32.IInspectable
}

func (this *IUsbConfigurationDescriptorStatics) Vtbl() *IUsbConfigurationDescriptorStaticsVtbl {
	return (*IUsbConfigurationDescriptorStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbConfigurationDescriptorStatics) TryParse(descriptor *IUsbDescriptor, parsed *IUsbConfigurationDescriptor) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParse, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(descriptor)), uintptr(unsafe.Pointer(parsed)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbConfigurationDescriptorStatics) Parse(descriptor *IUsbDescriptor) *IUsbConfigurationDescriptor {
	var _result *IUsbConfigurationDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Parse, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(descriptor)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8E9465A6-D73D-46DE-94BE-AAE7F07C0F5C
var IID_IUsbControlRequestType = syscall.GUID{0x8E9465A6, 0xD73D, 0x46DE,
	[8]byte{0x94, 0xBE, 0xAA, 0xE7, 0xF0, 0x7C, 0x0F, 0x5C}}

type IUsbControlRequestTypeInterface interface {
	win32.IInspectableInterface
	Get_Direction() UsbTransferDirection
	Put_Direction(value UsbTransferDirection)
	Get_ControlTransferType() UsbControlTransferType
	Put_ControlTransferType(value UsbControlTransferType)
	Get_Recipient() UsbControlRecipient
	Put_Recipient(value UsbControlRecipient)
	Get_AsByte() byte
	Put_AsByte(value byte)
}

type IUsbControlRequestTypeVtbl struct {
	win32.IInspectableVtbl
	Get_Direction           uintptr
	Put_Direction           uintptr
	Get_ControlTransferType uintptr
	Put_ControlTransferType uintptr
	Get_Recipient           uintptr
	Put_Recipient           uintptr
	Get_AsByte              uintptr
	Put_AsByte              uintptr
}

type IUsbControlRequestType struct {
	win32.IInspectable
}

func (this *IUsbControlRequestType) Vtbl() *IUsbControlRequestTypeVtbl {
	return (*IUsbControlRequestTypeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbControlRequestType) Get_Direction() UsbTransferDirection {
	var _result UsbTransferDirection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Direction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbControlRequestType) Put_Direction(value UsbTransferDirection) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Direction, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IUsbControlRequestType) Get_ControlTransferType() UsbControlTransferType {
	var _result UsbControlTransferType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ControlTransferType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbControlRequestType) Put_ControlTransferType(value UsbControlTransferType) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ControlTransferType, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IUsbControlRequestType) Get_Recipient() UsbControlRecipient {
	var _result UsbControlRecipient
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Recipient, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbControlRequestType) Put_Recipient(value UsbControlRecipient) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Recipient, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IUsbControlRequestType) Get_AsByte() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AsByte, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbControlRequestType) Put_AsByte(value byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AsByte, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 0A89F216-5F9D-4874-8904-DA9AD3F5528F
var IID_IUsbDescriptor = syscall.GUID{0x0A89F216, 0x5F9D, 0x4874,
	[8]byte{0x89, 0x04, 0xDA, 0x9A, 0xD3, 0xF5, 0x52, 0x8F}}

type IUsbDescriptorInterface interface {
	win32.IInspectableInterface
	Get_Length() byte
	Get_DescriptorType() byte
	ReadDescriptorBuffer(buffer *IBuffer)
}

type IUsbDescriptorVtbl struct {
	win32.IInspectableVtbl
	Get_Length           uintptr
	Get_DescriptorType   uintptr
	ReadDescriptorBuffer uintptr
}

type IUsbDescriptor struct {
	win32.IInspectable
}

func (this *IUsbDescriptor) Vtbl() *IUsbDescriptorVtbl {
	return (*IUsbDescriptorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbDescriptor) Get_Length() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Length, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbDescriptor) Get_DescriptorType() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DescriptorType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbDescriptor) ReadDescriptorBuffer(buffer *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadDescriptorBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(buffer)))
	_ = _hr
}

// 5249B992-C456-44D5-AD5E-24F5A089F63B
var IID_IUsbDevice = syscall.GUID{0x5249B992, 0xC456, 0x44D5,
	[8]byte{0xAD, 0x5E, 0x24, 0xF5, 0xA0, 0x89, 0xF6, 0x3B}}

type IUsbDeviceInterface interface {
	win32.IInspectableInterface
	SendControlOutTransferAsync(setupPacket *IUsbSetupPacket, buffer *IBuffer) *IAsyncOperation[uint32]
	SendControlOutTransferAsyncNoBuffer(setupPacket *IUsbSetupPacket) *IAsyncOperation[uint32]
	SendControlInTransferAsync(setupPacket *IUsbSetupPacket, buffer *IBuffer) *IAsyncOperation[*IBuffer]
	SendControlInTransferAsyncNoBuffer(setupPacket *IUsbSetupPacket) *IAsyncOperation[*IBuffer]
	Get_DefaultInterface() *IUsbInterface
	Get_DeviceDescriptor() *IUsbDeviceDescriptor
	Get_Configuration() *IUsbConfiguration
}

type IUsbDeviceVtbl struct {
	win32.IInspectableVtbl
	SendControlOutTransferAsync         uintptr
	SendControlOutTransferAsyncNoBuffer uintptr
	SendControlInTransferAsync          uintptr
	SendControlInTransferAsyncNoBuffer  uintptr
	Get_DefaultInterface                uintptr
	Get_DeviceDescriptor                uintptr
	Get_Configuration                   uintptr
}

type IUsbDevice struct {
	win32.IInspectable
}

func (this *IUsbDevice) Vtbl() *IUsbDeviceVtbl {
	return (*IUsbDeviceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbDevice) SendControlOutTransferAsync(setupPacket *IUsbSetupPacket, buffer *IBuffer) *IAsyncOperation[uint32] {
	var _result *IAsyncOperation[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendControlOutTransferAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(setupPacket)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbDevice) SendControlOutTransferAsyncNoBuffer(setupPacket *IUsbSetupPacket) *IAsyncOperation[uint32] {
	var _result *IAsyncOperation[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendControlOutTransferAsyncNoBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(setupPacket)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbDevice) SendControlInTransferAsync(setupPacket *IUsbSetupPacket, buffer *IBuffer) *IAsyncOperation[*IBuffer] {
	var _result *IAsyncOperation[*IBuffer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendControlInTransferAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(setupPacket)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbDevice) SendControlInTransferAsyncNoBuffer(setupPacket *IUsbSetupPacket) *IAsyncOperation[*IBuffer] {
	var _result *IAsyncOperation[*IBuffer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendControlInTransferAsyncNoBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(setupPacket)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbDevice) Get_DefaultInterface() *IUsbInterface {
	var _result *IUsbInterface
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DefaultInterface, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbDevice) Get_DeviceDescriptor() *IUsbDeviceDescriptor {
	var _result *IUsbDeviceDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceDescriptor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbDevice) Get_Configuration() *IUsbConfiguration {
	var _result *IUsbConfiguration
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Configuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 051942F9-845E-47EB-B12A-38F2F617AFE7
var IID_IUsbDeviceClass = syscall.GUID{0x051942F9, 0x845E, 0x47EB,
	[8]byte{0xB1, 0x2A, 0x38, 0xF2, 0xF6, 0x17, 0xAF, 0xE7}}

type IUsbDeviceClassInterface interface {
	win32.IInspectableInterface
	Get_ClassCode() byte
	Put_ClassCode(value byte)
	Get_SubclassCode() *IReference[byte]
	Put_SubclassCode(value *IReference[byte])
	Get_ProtocolCode() *IReference[byte]
	Put_ProtocolCode(value *IReference[byte])
}

type IUsbDeviceClassVtbl struct {
	win32.IInspectableVtbl
	Get_ClassCode    uintptr
	Put_ClassCode    uintptr
	Get_SubclassCode uintptr
	Put_SubclassCode uintptr
	Get_ProtocolCode uintptr
	Put_ProtocolCode uintptr
}

type IUsbDeviceClass struct {
	win32.IInspectable
}

func (this *IUsbDeviceClass) Vtbl() *IUsbDeviceClassVtbl {
	return (*IUsbDeviceClassVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbDeviceClass) Get_ClassCode() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ClassCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbDeviceClass) Put_ClassCode(value byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ClassCode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IUsbDeviceClass) Get_SubclassCode() *IReference[byte] {
	var _result *IReference[byte]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SubclassCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbDeviceClass) Put_SubclassCode(value *IReference[byte]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SubclassCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IUsbDeviceClass) Get_ProtocolCode() *IReference[byte] {
	var _result *IReference[byte]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtocolCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbDeviceClass) Put_ProtocolCode(value *IReference[byte]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProtocolCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 686F955D-9B92-4B30-9781-C22C55AC35CB
var IID_IUsbDeviceClasses = syscall.GUID{0x686F955D, 0x9B92, 0x4B30,
	[8]byte{0x97, 0x81, 0xC2, 0x2C, 0x55, 0xAC, 0x35, 0xCB}}

type IUsbDeviceClassesInterface interface {
	win32.IInspectableInterface
}

type IUsbDeviceClassesVtbl struct {
	win32.IInspectableVtbl
}

type IUsbDeviceClasses struct {
	win32.IInspectable
}

func (this *IUsbDeviceClasses) Vtbl() *IUsbDeviceClassesVtbl {
	return (*IUsbDeviceClassesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// B20B0527-C580-4599-A165-981B4FD03230
var IID_IUsbDeviceClassesStatics = syscall.GUID{0xB20B0527, 0xC580, 0x4599,
	[8]byte{0xA1, 0x65, 0x98, 0x1B, 0x4F, 0xD0, 0x32, 0x30}}

type IUsbDeviceClassesStaticsInterface interface {
	win32.IInspectableInterface
	Get_CdcControl() *IUsbDeviceClass
	Get_Physical() *IUsbDeviceClass
	Get_PersonalHealthcare() *IUsbDeviceClass
	Get_ActiveSync() *IUsbDeviceClass
	Get_PalmSync() *IUsbDeviceClass
	Get_DeviceFirmwareUpdate() *IUsbDeviceClass
	Get_Irda() *IUsbDeviceClass
	Get_Measurement() *IUsbDeviceClass
	Get_VendorSpecific() *IUsbDeviceClass
}

type IUsbDeviceClassesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_CdcControl           uintptr
	Get_Physical             uintptr
	Get_PersonalHealthcare   uintptr
	Get_ActiveSync           uintptr
	Get_PalmSync             uintptr
	Get_DeviceFirmwareUpdate uintptr
	Get_Irda                 uintptr
	Get_Measurement          uintptr
	Get_VendorSpecific       uintptr
}

type IUsbDeviceClassesStatics struct {
	win32.IInspectable
}

func (this *IUsbDeviceClassesStatics) Vtbl() *IUsbDeviceClassesStaticsVtbl {
	return (*IUsbDeviceClassesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbDeviceClassesStatics) Get_CdcControl() *IUsbDeviceClass {
	var _result *IUsbDeviceClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CdcControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbDeviceClassesStatics) Get_Physical() *IUsbDeviceClass {
	var _result *IUsbDeviceClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Physical, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbDeviceClassesStatics) Get_PersonalHealthcare() *IUsbDeviceClass {
	var _result *IUsbDeviceClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PersonalHealthcare, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbDeviceClassesStatics) Get_ActiveSync() *IUsbDeviceClass {
	var _result *IUsbDeviceClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActiveSync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbDeviceClassesStatics) Get_PalmSync() *IUsbDeviceClass {
	var _result *IUsbDeviceClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PalmSync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbDeviceClassesStatics) Get_DeviceFirmwareUpdate() *IUsbDeviceClass {
	var _result *IUsbDeviceClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceFirmwareUpdate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbDeviceClassesStatics) Get_Irda() *IUsbDeviceClass {
	var _result *IUsbDeviceClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Irda, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbDeviceClassesStatics) Get_Measurement() *IUsbDeviceClass {
	var _result *IUsbDeviceClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Measurement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbDeviceClassesStatics) Get_VendorSpecific() *IUsbDeviceClass {
	var _result *IUsbDeviceClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VendorSpecific, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1F48D1F6-BA97-4322-B92C-B5B189216588
var IID_IUsbDeviceDescriptor = syscall.GUID{0x1F48D1F6, 0xBA97, 0x4322,
	[8]byte{0xB9, 0x2C, 0xB5, 0xB1, 0x89, 0x21, 0x65, 0x88}}

type IUsbDeviceDescriptorInterface interface {
	win32.IInspectableInterface
	Get_BcdUsb() uint32
	Get_MaxPacketSize0() byte
	Get_VendorId() uint32
	Get_ProductId() uint32
	Get_BcdDeviceRevision() uint32
	Get_NumberOfConfigurations() byte
}

type IUsbDeviceDescriptorVtbl struct {
	win32.IInspectableVtbl
	Get_BcdUsb                 uintptr
	Get_MaxPacketSize0         uintptr
	Get_VendorId               uintptr
	Get_ProductId              uintptr
	Get_BcdDeviceRevision      uintptr
	Get_NumberOfConfigurations uintptr
}

type IUsbDeviceDescriptor struct {
	win32.IInspectable
}

func (this *IUsbDeviceDescriptor) Vtbl() *IUsbDeviceDescriptorVtbl {
	return (*IUsbDeviceDescriptorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbDeviceDescriptor) Get_BcdUsb() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BcdUsb, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbDeviceDescriptor) Get_MaxPacketSize0() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxPacketSize0, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbDeviceDescriptor) Get_VendorId() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VendorId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbDeviceDescriptor) Get_ProductId() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProductId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbDeviceDescriptor) Get_BcdDeviceRevision() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BcdDeviceRevision, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbDeviceDescriptor) Get_NumberOfConfigurations() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NumberOfConfigurations, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 066B85A2-09B7-4446-8502-6FE6DCAA7309
var IID_IUsbDeviceStatics = syscall.GUID{0x066B85A2, 0x09B7, 0x4446,
	[8]byte{0x85, 0x02, 0x6F, 0xE6, 0xDC, 0xAA, 0x73, 0x09}}

type IUsbDeviceStaticsInterface interface {
	win32.IInspectableInterface
	GetDeviceSelector(vendorId uint32, productId uint32, winUsbInterfaceClass syscall.GUID) string
	GetDeviceSelectorGuidOnly(winUsbInterfaceClass syscall.GUID) string
	GetDeviceSelectorVidPidOnly(vendorId uint32, productId uint32) string
	GetDeviceClassSelector(usbClass *IUsbDeviceClass) string
	FromIdAsync(deviceId string) *IAsyncOperation[*IUsbDevice]
}

type IUsbDeviceStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector           uintptr
	GetDeviceSelectorGuidOnly   uintptr
	GetDeviceSelectorVidPidOnly uintptr
	GetDeviceClassSelector      uintptr
	FromIdAsync                 uintptr
}

type IUsbDeviceStatics struct {
	win32.IInspectable
}

func (this *IUsbDeviceStatics) Vtbl() *IUsbDeviceStaticsVtbl {
	return (*IUsbDeviceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbDeviceStatics) GetDeviceSelector(vendorId uint32, productId uint32, winUsbInterfaceClass syscall.GUID) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(vendorId), uintptr(productId), uintptr(unsafe.Pointer(&winUsbInterfaceClass)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUsbDeviceStatics) GetDeviceSelectorGuidOnly(winUsbInterfaceClass syscall.GUID) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorGuidOnly, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&winUsbInterfaceClass)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUsbDeviceStatics) GetDeviceSelectorVidPidOnly(vendorId uint32, productId uint32) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorVidPidOnly, uintptr(unsafe.Pointer(this)), uintptr(vendorId), uintptr(productId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUsbDeviceStatics) GetDeviceClassSelector(usbClass *IUsbDeviceClass) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceClassSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(usbClass)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUsbDeviceStatics) FromIdAsync(deviceId string) *IAsyncOperation[*IUsbDevice] {
	var _result *IAsyncOperation[*IUsbDevice]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6B4862D9-8DF7-4B40-AC83-578F139F0575
var IID_IUsbEndpointDescriptor = syscall.GUID{0x6B4862D9, 0x8DF7, 0x4B40,
	[8]byte{0xAC, 0x83, 0x57, 0x8F, 0x13, 0x9F, 0x05, 0x75}}

type IUsbEndpointDescriptorInterface interface {
	win32.IInspectableInterface
	Get_EndpointNumber() byte
	Get_Direction() UsbTransferDirection
	Get_EndpointType() UsbEndpointType
	Get_AsBulkInEndpointDescriptor() *IUsbBulkInEndpointDescriptor
	Get_AsInterruptInEndpointDescriptor() *IUsbInterruptInEndpointDescriptor
	Get_AsBulkOutEndpointDescriptor() *IUsbBulkOutEndpointDescriptor
	Get_AsInterruptOutEndpointDescriptor() *IUsbInterruptOutEndpointDescriptor
}

type IUsbEndpointDescriptorVtbl struct {
	win32.IInspectableVtbl
	Get_EndpointNumber                   uintptr
	Get_Direction                        uintptr
	Get_EndpointType                     uintptr
	Get_AsBulkInEndpointDescriptor       uintptr
	Get_AsInterruptInEndpointDescriptor  uintptr
	Get_AsBulkOutEndpointDescriptor      uintptr
	Get_AsInterruptOutEndpointDescriptor uintptr
}

type IUsbEndpointDescriptor struct {
	win32.IInspectable
}

func (this *IUsbEndpointDescriptor) Vtbl() *IUsbEndpointDescriptorVtbl {
	return (*IUsbEndpointDescriptorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbEndpointDescriptor) Get_EndpointNumber() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EndpointNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbEndpointDescriptor) Get_Direction() UsbTransferDirection {
	var _result UsbTransferDirection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Direction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbEndpointDescriptor) Get_EndpointType() UsbEndpointType {
	var _result UsbEndpointType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EndpointType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbEndpointDescriptor) Get_AsBulkInEndpointDescriptor() *IUsbBulkInEndpointDescriptor {
	var _result *IUsbBulkInEndpointDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AsBulkInEndpointDescriptor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbEndpointDescriptor) Get_AsInterruptInEndpointDescriptor() *IUsbInterruptInEndpointDescriptor {
	var _result *IUsbInterruptInEndpointDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AsInterruptInEndpointDescriptor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbEndpointDescriptor) Get_AsBulkOutEndpointDescriptor() *IUsbBulkOutEndpointDescriptor {
	var _result *IUsbBulkOutEndpointDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AsBulkOutEndpointDescriptor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbEndpointDescriptor) Get_AsInterruptOutEndpointDescriptor() *IUsbInterruptOutEndpointDescriptor {
	var _result *IUsbInterruptOutEndpointDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AsInterruptOutEndpointDescriptor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C890B201-9A6A-495E-A82C-295B9E708106
var IID_IUsbEndpointDescriptorStatics = syscall.GUID{0xC890B201, 0x9A6A, 0x495E,
	[8]byte{0xA8, 0x2C, 0x29, 0x5B, 0x9E, 0x70, 0x81, 0x06}}

type IUsbEndpointDescriptorStaticsInterface interface {
	win32.IInspectableInterface
	TryParse(descriptor *IUsbDescriptor, parsed *IUsbEndpointDescriptor) bool
	Parse(descriptor *IUsbDescriptor) *IUsbEndpointDescriptor
}

type IUsbEndpointDescriptorStaticsVtbl struct {
	win32.IInspectableVtbl
	TryParse uintptr
	Parse    uintptr
}

type IUsbEndpointDescriptorStatics struct {
	win32.IInspectable
}

func (this *IUsbEndpointDescriptorStatics) Vtbl() *IUsbEndpointDescriptorStaticsVtbl {
	return (*IUsbEndpointDescriptorStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbEndpointDescriptorStatics) TryParse(descriptor *IUsbDescriptor, parsed *IUsbEndpointDescriptor) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParse, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(descriptor)), uintptr(unsafe.Pointer(parsed)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbEndpointDescriptorStatics) Parse(descriptor *IUsbDescriptor) *IUsbEndpointDescriptor {
	var _result *IUsbEndpointDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Parse, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(descriptor)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A0322B95-7F47-48AB-A727-678C25BE2112
var IID_IUsbInterface = syscall.GUID{0xA0322B95, 0x7F47, 0x48AB,
	[8]byte{0xA7, 0x27, 0x67, 0x8C, 0x25, 0xBE, 0x21, 0x12}}

type IUsbInterfaceInterface interface {
	win32.IInspectableInterface
	Get_BulkInPipes() *IVectorView[*IUsbBulkInPipe]
	Get_InterruptInPipes() *IVectorView[*IUsbInterruptInPipe]
	Get_BulkOutPipes() *IVectorView[*IUsbBulkOutPipe]
	Get_InterruptOutPipes() *IVectorView[*IUsbInterruptOutPipe]
	Get_InterfaceSettings() *IVectorView[*IUsbInterfaceSetting]
	Get_InterfaceNumber() byte
	Get_Descriptors() *IVectorView[*IUsbDescriptor]
}

type IUsbInterfaceVtbl struct {
	win32.IInspectableVtbl
	Get_BulkInPipes       uintptr
	Get_InterruptInPipes  uintptr
	Get_BulkOutPipes      uintptr
	Get_InterruptOutPipes uintptr
	Get_InterfaceSettings uintptr
	Get_InterfaceNumber   uintptr
	Get_Descriptors       uintptr
}

type IUsbInterface struct {
	win32.IInspectable
}

func (this *IUsbInterface) Vtbl() *IUsbInterfaceVtbl {
	return (*IUsbInterfaceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbInterface) Get_BulkInPipes() *IVectorView[*IUsbBulkInPipe] {
	var _result *IVectorView[*IUsbBulkInPipe]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BulkInPipes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbInterface) Get_InterruptInPipes() *IVectorView[*IUsbInterruptInPipe] {
	var _result *IVectorView[*IUsbInterruptInPipe]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InterruptInPipes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbInterface) Get_BulkOutPipes() *IVectorView[*IUsbBulkOutPipe] {
	var _result *IVectorView[*IUsbBulkOutPipe]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BulkOutPipes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbInterface) Get_InterruptOutPipes() *IVectorView[*IUsbInterruptOutPipe] {
	var _result *IVectorView[*IUsbInterruptOutPipe]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InterruptOutPipes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbInterface) Get_InterfaceSettings() *IVectorView[*IUsbInterfaceSetting] {
	var _result *IVectorView[*IUsbInterfaceSetting]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InterfaceSettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbInterface) Get_InterfaceNumber() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InterfaceNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbInterface) Get_Descriptors() *IVectorView[*IUsbDescriptor] {
	var _result *IVectorView[*IUsbDescriptor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Descriptors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 199670C7-B7EE-4F90-8CD5-94A2E257598A
var IID_IUsbInterfaceDescriptor = syscall.GUID{0x199670C7, 0xB7EE, 0x4F90,
	[8]byte{0x8C, 0xD5, 0x94, 0xA2, 0xE2, 0x57, 0x59, 0x8A}}

type IUsbInterfaceDescriptorInterface interface {
	win32.IInspectableInterface
	Get_ClassCode() byte
	Get_SubclassCode() byte
	Get_ProtocolCode() byte
	Get_AlternateSettingNumber() byte
	Get_InterfaceNumber() byte
}

type IUsbInterfaceDescriptorVtbl struct {
	win32.IInspectableVtbl
	Get_ClassCode              uintptr
	Get_SubclassCode           uintptr
	Get_ProtocolCode           uintptr
	Get_AlternateSettingNumber uintptr
	Get_InterfaceNumber        uintptr
}

type IUsbInterfaceDescriptor struct {
	win32.IInspectable
}

func (this *IUsbInterfaceDescriptor) Vtbl() *IUsbInterfaceDescriptorVtbl {
	return (*IUsbInterfaceDescriptorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbInterfaceDescriptor) Get_ClassCode() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ClassCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbInterfaceDescriptor) Get_SubclassCode() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SubclassCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbInterfaceDescriptor) Get_ProtocolCode() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtocolCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbInterfaceDescriptor) Get_AlternateSettingNumber() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlternateSettingNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbInterfaceDescriptor) Get_InterfaceNumber() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InterfaceNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// E34A9FF5-77D6-48B6-B0BE-16C6422316FE
var IID_IUsbInterfaceDescriptorStatics = syscall.GUID{0xE34A9FF5, 0x77D6, 0x48B6,
	[8]byte{0xB0, 0xBE, 0x16, 0xC6, 0x42, 0x23, 0x16, 0xFE}}

type IUsbInterfaceDescriptorStaticsInterface interface {
	win32.IInspectableInterface
	TryParse(descriptor *IUsbDescriptor, parsed *IUsbInterfaceDescriptor) bool
	Parse(descriptor *IUsbDescriptor) *IUsbInterfaceDescriptor
}

type IUsbInterfaceDescriptorStaticsVtbl struct {
	win32.IInspectableVtbl
	TryParse uintptr
	Parse    uintptr
}

type IUsbInterfaceDescriptorStatics struct {
	win32.IInspectable
}

func (this *IUsbInterfaceDescriptorStatics) Vtbl() *IUsbInterfaceDescriptorStaticsVtbl {
	return (*IUsbInterfaceDescriptorStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbInterfaceDescriptorStatics) TryParse(descriptor *IUsbDescriptor, parsed *IUsbInterfaceDescriptor) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParse, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(descriptor)), uintptr(unsafe.Pointer(parsed)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbInterfaceDescriptorStatics) Parse(descriptor *IUsbDescriptor) *IUsbInterfaceDescriptor {
	var _result *IUsbInterfaceDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Parse, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(descriptor)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1827BBA7-8DA7-4AF7-8F4C-7F3032E781F5
var IID_IUsbInterfaceSetting = syscall.GUID{0x1827BBA7, 0x8DA7, 0x4AF7,
	[8]byte{0x8F, 0x4C, 0x7F, 0x30, 0x32, 0xE7, 0x81, 0xF5}}

type IUsbInterfaceSettingInterface interface {
	win32.IInspectableInterface
	Get_BulkInEndpoints() *IVectorView[*IUsbBulkInEndpointDescriptor]
	Get_InterruptInEndpoints() *IVectorView[*IUsbInterruptInEndpointDescriptor]
	Get_BulkOutEndpoints() *IVectorView[*IUsbBulkOutEndpointDescriptor]
	Get_InterruptOutEndpoints() *IVectorView[*IUsbInterruptOutEndpointDescriptor]
	Get_Selected() bool
	SelectSettingAsync() *IAsyncAction
	Get_InterfaceDescriptor() *IUsbInterfaceDescriptor
	Get_Descriptors() *IVectorView[*IUsbDescriptor]
}

type IUsbInterfaceSettingVtbl struct {
	win32.IInspectableVtbl
	Get_BulkInEndpoints       uintptr
	Get_InterruptInEndpoints  uintptr
	Get_BulkOutEndpoints      uintptr
	Get_InterruptOutEndpoints uintptr
	Get_Selected              uintptr
	SelectSettingAsync        uintptr
	Get_InterfaceDescriptor   uintptr
	Get_Descriptors           uintptr
}

type IUsbInterfaceSetting struct {
	win32.IInspectable
}

func (this *IUsbInterfaceSetting) Vtbl() *IUsbInterfaceSettingVtbl {
	return (*IUsbInterfaceSettingVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbInterfaceSetting) Get_BulkInEndpoints() *IVectorView[*IUsbBulkInEndpointDescriptor] {
	var _result *IVectorView[*IUsbBulkInEndpointDescriptor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BulkInEndpoints, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbInterfaceSetting) Get_InterruptInEndpoints() *IVectorView[*IUsbInterruptInEndpointDescriptor] {
	var _result *IVectorView[*IUsbInterruptInEndpointDescriptor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InterruptInEndpoints, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbInterfaceSetting) Get_BulkOutEndpoints() *IVectorView[*IUsbBulkOutEndpointDescriptor] {
	var _result *IVectorView[*IUsbBulkOutEndpointDescriptor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BulkOutEndpoints, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbInterfaceSetting) Get_InterruptOutEndpoints() *IVectorView[*IUsbInterruptOutEndpointDescriptor] {
	var _result *IVectorView[*IUsbInterruptOutEndpointDescriptor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InterruptOutEndpoints, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbInterfaceSetting) Get_Selected() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Selected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbInterfaceSetting) SelectSettingAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SelectSettingAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbInterfaceSetting) Get_InterfaceDescriptor() *IUsbInterfaceDescriptor {
	var _result *IUsbInterfaceDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InterfaceDescriptor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbInterfaceSetting) Get_Descriptors() *IVectorView[*IUsbDescriptor] {
	var _result *IVectorView[*IUsbDescriptor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Descriptors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C0528967-C911-4C3A-86B2-419C2DA89039
var IID_IUsbInterruptInEndpointDescriptor = syscall.GUID{0xC0528967, 0xC911, 0x4C3A,
	[8]byte{0x86, 0xB2, 0x41, 0x9C, 0x2D, 0xA8, 0x90, 0x39}}

type IUsbInterruptInEndpointDescriptorInterface interface {
	win32.IInspectableInterface
	Get_MaxPacketSize() uint32
	Get_EndpointNumber() byte
	Get_Interval() TimeSpan
	Get_Pipe() *IUsbInterruptInPipe
}

type IUsbInterruptInEndpointDescriptorVtbl struct {
	win32.IInspectableVtbl
	Get_MaxPacketSize  uintptr
	Get_EndpointNumber uintptr
	Get_Interval       uintptr
	Get_Pipe           uintptr
}

type IUsbInterruptInEndpointDescriptor struct {
	win32.IInspectable
}

func (this *IUsbInterruptInEndpointDescriptor) Vtbl() *IUsbInterruptInEndpointDescriptorVtbl {
	return (*IUsbInterruptInEndpointDescriptorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbInterruptInEndpointDescriptor) Get_MaxPacketSize() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxPacketSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbInterruptInEndpointDescriptor) Get_EndpointNumber() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EndpointNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbInterruptInEndpointDescriptor) Get_Interval() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Interval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbInterruptInEndpointDescriptor) Get_Pipe() *IUsbInterruptInPipe {
	var _result *IUsbInterruptInPipe
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Pipe, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B7B04092-1418-4936-8209-299CF5605583
var IID_IUsbInterruptInEventArgs = syscall.GUID{0xB7B04092, 0x1418, 0x4936,
	[8]byte{0x82, 0x09, 0x29, 0x9C, 0xF5, 0x60, 0x55, 0x83}}

type IUsbInterruptInEventArgsInterface interface {
	win32.IInspectableInterface
	Get_InterruptData() *IBuffer
}

type IUsbInterruptInEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_InterruptData uintptr
}

type IUsbInterruptInEventArgs struct {
	win32.IInspectable
}

func (this *IUsbInterruptInEventArgs) Vtbl() *IUsbInterruptInEventArgsVtbl {
	return (*IUsbInterruptInEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbInterruptInEventArgs) Get_InterruptData() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InterruptData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FA007116-84D7-48C7-8A3F-4C0B235F2EA6
var IID_IUsbInterruptInPipe = syscall.GUID{0xFA007116, 0x84D7, 0x48C7,
	[8]byte{0x8A, 0x3F, 0x4C, 0x0B, 0x23, 0x5F, 0x2E, 0xA6}}

type IUsbInterruptInPipeInterface interface {
	win32.IInspectableInterface
	Get_EndpointDescriptor() *IUsbInterruptInEndpointDescriptor
	ClearStallAsync() *IAsyncAction
	Add_DataReceived(handler TypedEventHandler[*IUsbInterruptInPipe, *IUsbInterruptInEventArgs]) EventRegistrationToken
	Remove_DataReceived(token EventRegistrationToken)
}

type IUsbInterruptInPipeVtbl struct {
	win32.IInspectableVtbl
	Get_EndpointDescriptor uintptr
	ClearStallAsync        uintptr
	Add_DataReceived       uintptr
	Remove_DataReceived    uintptr
}

type IUsbInterruptInPipe struct {
	win32.IInspectable
}

func (this *IUsbInterruptInPipe) Vtbl() *IUsbInterruptInPipeVtbl {
	return (*IUsbInterruptInPipeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbInterruptInPipe) Get_EndpointDescriptor() *IUsbInterruptInEndpointDescriptor {
	var _result *IUsbInterruptInEndpointDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EndpointDescriptor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbInterruptInPipe) ClearStallAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearStallAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbInterruptInPipe) Add_DataReceived(handler TypedEventHandler[*IUsbInterruptInPipe, *IUsbInterruptInEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DataReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbInterruptInPipe) Remove_DataReceived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DataReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// CC9FED81-10CA-4533-952D-9E278341E80F
var IID_IUsbInterruptOutEndpointDescriptor = syscall.GUID{0xCC9FED81, 0x10CA, 0x4533,
	[8]byte{0x95, 0x2D, 0x9E, 0x27, 0x83, 0x41, 0xE8, 0x0F}}

type IUsbInterruptOutEndpointDescriptorInterface interface {
	win32.IInspectableInterface
	Get_MaxPacketSize() uint32
	Get_EndpointNumber() byte
	Get_Interval() TimeSpan
	Get_Pipe() *IUsbInterruptOutPipe
}

type IUsbInterruptOutEndpointDescriptorVtbl struct {
	win32.IInspectableVtbl
	Get_MaxPacketSize  uintptr
	Get_EndpointNumber uintptr
	Get_Interval       uintptr
	Get_Pipe           uintptr
}

type IUsbInterruptOutEndpointDescriptor struct {
	win32.IInspectable
}

func (this *IUsbInterruptOutEndpointDescriptor) Vtbl() *IUsbInterruptOutEndpointDescriptorVtbl {
	return (*IUsbInterruptOutEndpointDescriptorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbInterruptOutEndpointDescriptor) Get_MaxPacketSize() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxPacketSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbInterruptOutEndpointDescriptor) Get_EndpointNumber() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EndpointNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbInterruptOutEndpointDescriptor) Get_Interval() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Interval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbInterruptOutEndpointDescriptor) Get_Pipe() *IUsbInterruptOutPipe {
	var _result *IUsbInterruptOutPipe
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Pipe, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E984C8A9-AAF9-49D0-B96C-F661AB4A7F95
var IID_IUsbInterruptOutPipe = syscall.GUID{0xE984C8A9, 0xAAF9, 0x49D0,
	[8]byte{0xB9, 0x6C, 0xF6, 0x61, 0xAB, 0x4A, 0x7F, 0x95}}

type IUsbInterruptOutPipeInterface interface {
	win32.IInspectableInterface
	Get_EndpointDescriptor() *IUsbInterruptOutEndpointDescriptor
	ClearStallAsync() *IAsyncAction
	Put_WriteOptions(value UsbWriteOptions)
	Get_WriteOptions() UsbWriteOptions
	Get_OutputStream() *IOutputStream
}

type IUsbInterruptOutPipeVtbl struct {
	win32.IInspectableVtbl
	Get_EndpointDescriptor uintptr
	ClearStallAsync        uintptr
	Put_WriteOptions       uintptr
	Get_WriteOptions       uintptr
	Get_OutputStream       uintptr
}

type IUsbInterruptOutPipe struct {
	win32.IInspectable
}

func (this *IUsbInterruptOutPipe) Vtbl() *IUsbInterruptOutPipeVtbl {
	return (*IUsbInterruptOutPipeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbInterruptOutPipe) Get_EndpointDescriptor() *IUsbInterruptOutEndpointDescriptor {
	var _result *IUsbInterruptOutEndpointDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EndpointDescriptor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbInterruptOutPipe) ClearStallAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearStallAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbInterruptOutPipe) Put_WriteOptions(value UsbWriteOptions) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_WriteOptions, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IUsbInterruptOutPipe) Get_WriteOptions() UsbWriteOptions {
	var _result UsbWriteOptions
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WriteOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbInterruptOutPipe) Get_OutputStream() *IOutputStream {
	var _result *IOutputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutputStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 104BA132-C78F-4C51-B654-E49D02F2CB03
var IID_IUsbSetupPacket = syscall.GUID{0x104BA132, 0xC78F, 0x4C51,
	[8]byte{0xB6, 0x54, 0xE4, 0x9D, 0x02, 0xF2, 0xCB, 0x03}}

type IUsbSetupPacketInterface interface {
	win32.IInspectableInterface
	Get_RequestType() *IUsbControlRequestType
	Put_RequestType(value *IUsbControlRequestType)
	Get_Request() byte
	Put_Request(value byte)
	Get_Value() uint32
	Put_Value(value uint32)
	Get_Index() uint32
	Put_Index(value uint32)
	Get_Length() uint32
	Put_Length(value uint32)
}

type IUsbSetupPacketVtbl struct {
	win32.IInspectableVtbl
	Get_RequestType uintptr
	Put_RequestType uintptr
	Get_Request     uintptr
	Put_Request     uintptr
	Get_Value       uintptr
	Put_Value       uintptr
	Get_Index       uintptr
	Put_Index       uintptr
	Get_Length      uintptr
	Put_Length      uintptr
}

type IUsbSetupPacket struct {
	win32.IInspectable
}

func (this *IUsbSetupPacket) Vtbl() *IUsbSetupPacketVtbl {
	return (*IUsbSetupPacketVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbSetupPacket) Get_RequestType() *IUsbControlRequestType {
	var _result *IUsbControlRequestType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUsbSetupPacket) Put_RequestType(value *IUsbControlRequestType) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RequestType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IUsbSetupPacket) Get_Request() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Request, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbSetupPacket) Put_Request(value byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Request, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IUsbSetupPacket) Get_Value() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbSetupPacket) Put_Value(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Value, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IUsbSetupPacket) Get_Index() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Index, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbSetupPacket) Put_Index(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Index, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IUsbSetupPacket) Get_Length() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Length, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUsbSetupPacket) Put_Length(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Length, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// C9257D50-1B2E-4A41-A2A7-338F0CEF3C14
var IID_IUsbSetupPacketFactory = syscall.GUID{0xC9257D50, 0x1B2E, 0x4A41,
	[8]byte{0xA2, 0xA7, 0x33, 0x8F, 0x0C, 0xEF, 0x3C, 0x14}}

type IUsbSetupPacketFactoryInterface interface {
	win32.IInspectableInterface
	CreateWithEightByteBuffer(eightByteBuffer *IBuffer) *IUsbSetupPacket
}

type IUsbSetupPacketFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateWithEightByteBuffer uintptr
}

type IUsbSetupPacketFactory struct {
	win32.IInspectable
}

func (this *IUsbSetupPacketFactory) Vtbl() *IUsbSetupPacketFactoryVtbl {
	return (*IUsbSetupPacketFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUsbSetupPacketFactory) CreateWithEightByteBuffer(eightByteBuffer *IBuffer) *IUsbSetupPacket {
	var _result *IUsbSetupPacket
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithEightByteBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eightByteBuffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type UsbBulkInEndpointDescriptor struct {
	RtClass
	*IUsbBulkInEndpointDescriptor
}

type UsbBulkInPipe struct {
	RtClass
	*IUsbBulkInPipe
}

type UsbBulkOutEndpointDescriptor struct {
	RtClass
	*IUsbBulkOutEndpointDescriptor
}

type UsbBulkOutPipe struct {
	RtClass
	*IUsbBulkOutPipe
}

type UsbConfiguration struct {
	RtClass
	*IUsbConfiguration
}

type UsbConfigurationDescriptor struct {
	RtClass
	*IUsbConfigurationDescriptor
}

func NewIUsbConfigurationDescriptorStatics() *IUsbConfigurationDescriptorStatics {
	var p *IUsbConfigurationDescriptorStatics
	hs := NewHStr("Windows.Devices.Usb.UsbConfigurationDescriptor")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IUsbConfigurationDescriptorStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type UsbControlRequestType struct {
	RtClass
	*IUsbControlRequestType
}

func NewUsbControlRequestType() *UsbControlRequestType {
	hs := NewHStr("Windows.Devices.Usb.UsbControlRequestType")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &UsbControlRequestType{
		RtClass:                RtClass{PInspect: p},
		IUsbControlRequestType: (*IUsbControlRequestType)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type UsbDescriptor struct {
	RtClass
	*IUsbDescriptor
}

type UsbDevice struct {
	RtClass
	*IUsbDevice
}

func NewIUsbDeviceStatics() *IUsbDeviceStatics {
	var p *IUsbDeviceStatics
	hs := NewHStr("Windows.Devices.Usb.UsbDevice")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IUsbDeviceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type UsbDeviceClass struct {
	RtClass
	*IUsbDeviceClass
}

func NewUsbDeviceClass() *UsbDeviceClass {
	hs := NewHStr("Windows.Devices.Usb.UsbDeviceClass")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &UsbDeviceClass{
		RtClass:         RtClass{PInspect: p},
		IUsbDeviceClass: (*IUsbDeviceClass)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type UsbDeviceClasses struct {
	RtClass
	*IUsbDeviceClasses
}

func NewIUsbDeviceClassesStatics() *IUsbDeviceClassesStatics {
	var p *IUsbDeviceClassesStatics
	hs := NewHStr("Windows.Devices.Usb.UsbDeviceClasses")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IUsbDeviceClassesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type UsbDeviceDescriptor struct {
	RtClass
	*IUsbDeviceDescriptor
}

type UsbEndpointDescriptor struct {
	RtClass
	*IUsbEndpointDescriptor
}

func NewIUsbEndpointDescriptorStatics() *IUsbEndpointDescriptorStatics {
	var p *IUsbEndpointDescriptorStatics
	hs := NewHStr("Windows.Devices.Usb.UsbEndpointDescriptor")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IUsbEndpointDescriptorStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type UsbInterface struct {
	RtClass
	*IUsbInterface
}

type UsbInterfaceDescriptor struct {
	RtClass
	*IUsbInterfaceDescriptor
}

func NewIUsbInterfaceDescriptorStatics() *IUsbInterfaceDescriptorStatics {
	var p *IUsbInterfaceDescriptorStatics
	hs := NewHStr("Windows.Devices.Usb.UsbInterfaceDescriptor")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IUsbInterfaceDescriptorStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type UsbInterfaceSetting struct {
	RtClass
	*IUsbInterfaceSetting
}

type UsbInterruptInEndpointDescriptor struct {
	RtClass
	*IUsbInterruptInEndpointDescriptor
}

type UsbInterruptInEventArgs struct {
	RtClass
	*IUsbInterruptInEventArgs
}

type UsbInterruptInPipe struct {
	RtClass
	*IUsbInterruptInPipe
}

type UsbInterruptOutEndpointDescriptor struct {
	RtClass
	*IUsbInterruptOutEndpointDescriptor
}

type UsbInterruptOutPipe struct {
	RtClass
	*IUsbInterruptOutPipe
}

type UsbSetupPacket struct {
	RtClass
	*IUsbSetupPacket
}

func NewUsbSetupPacket() *UsbSetupPacket {
	hs := NewHStr("Windows.Devices.Usb.UsbSetupPacket")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &UsbSetupPacket{
		RtClass:         RtClass{PInspect: p},
		IUsbSetupPacket: (*IUsbSetupPacket)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewUsbSetupPacket_CreateWithEightByteBuffer(eightByteBuffer *IBuffer) *UsbSetupPacket {
	hs := NewHStr("Windows.Devices.Usb.UsbSetupPacket")
	var pFac *IUsbSetupPacketFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IUsbSetupPacketFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IUsbSetupPacket
	p = pFac.CreateWithEightByteBuffer(eightByteBuffer)
	result := &UsbSetupPacket{
		RtClass:         RtClass{PInspect: &p.IInspectable},
		IUsbSetupPacket: p,
	}
	com.AddToScope(result)
	return result
}
