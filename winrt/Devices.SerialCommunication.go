package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type SerialError int32

const (
	SerialError_Frame         SerialError = 0
	SerialError_BufferOverrun SerialError = 1
	SerialError_ReceiveFull   SerialError = 2
	SerialError_ReceiveParity SerialError = 3
	SerialError_TransmitFull  SerialError = 4
)

// enum
type SerialHandshake int32

const (
	SerialHandshake_None                 SerialHandshake = 0
	SerialHandshake_RequestToSend        SerialHandshake = 1
	SerialHandshake_XOnXOff              SerialHandshake = 2
	SerialHandshake_RequestToSendXOnXOff SerialHandshake = 3
)

// enum
type SerialParity int32

const (
	SerialParity_None  SerialParity = 0
	SerialParity_Odd   SerialParity = 1
	SerialParity_Even  SerialParity = 2
	SerialParity_Mark  SerialParity = 3
	SerialParity_Space SerialParity = 4
)

// enum
type SerialPinChange int32

const (
	SerialPinChange_BreakSignal   SerialPinChange = 0
	SerialPinChange_CarrierDetect SerialPinChange = 1
	SerialPinChange_ClearToSend   SerialPinChange = 2
	SerialPinChange_DataSetReady  SerialPinChange = 3
	SerialPinChange_RingIndicator SerialPinChange = 4
)

// enum
type SerialStopBitCount int32

const (
	SerialStopBitCount_One          SerialStopBitCount = 0
	SerialStopBitCount_OnePointFive SerialStopBitCount = 1
	SerialStopBitCount_Two          SerialStopBitCount = 2
)

// interfaces

// FCC6BF59-1283-4D8A-BFDF-566B33DDB28F
var IID_IErrorReceivedEventArgs = syscall.GUID{0xFCC6BF59, 0x1283, 0x4D8A,
	[8]byte{0xBF, 0xDF, 0x56, 0x6B, 0x33, 0xDD, 0xB2, 0x8F}}

type IErrorReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Error() SerialError
}

type IErrorReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Error uintptr
}

type IErrorReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IErrorReceivedEventArgs) Vtbl() *IErrorReceivedEventArgsVtbl {
	return (*IErrorReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IErrorReceivedEventArgs) Get_Error() SerialError {
	var _result SerialError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Error, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A2BF1DB0-FC9C-4607-93D0-FA5E8343EE22
var IID_IPinChangedEventArgs = syscall.GUID{0xA2BF1DB0, 0xFC9C, 0x4607,
	[8]byte{0x93, 0xD0, 0xFA, 0x5E, 0x83, 0x43, 0xEE, 0x22}}

type IPinChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_PinChange() SerialPinChange
}

type IPinChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_PinChange uintptr
}

type IPinChangedEventArgs struct {
	win32.IInspectable
}

func (this *IPinChangedEventArgs) Vtbl() *IPinChangedEventArgsVtbl {
	return (*IPinChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPinChangedEventArgs) Get_PinChange() SerialPinChange {
	var _result SerialPinChange
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PinChange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// E187CCC6-2210-414F-B65A-F5553A03372A
var IID_ISerialDevice = syscall.GUID{0xE187CCC6, 0x2210, 0x414F,
	[8]byte{0xB6, 0x5A, 0xF5, 0x55, 0x3A, 0x03, 0x37, 0x2A}}

type ISerialDeviceInterface interface {
	win32.IInspectableInterface
	Get_BaudRate() uint32
	Put_BaudRate(value uint32)
	Get_BreakSignalState() bool
	Put_BreakSignalState(value bool)
	Get_BytesReceived() uint32
	Get_CarrierDetectState() bool
	Get_ClearToSendState() bool
	Get_DataBits() uint16
	Put_DataBits(value uint16)
	Get_DataSetReadyState() bool
	Get_Handshake() SerialHandshake
	Put_Handshake(value SerialHandshake)
	Get_IsDataTerminalReadyEnabled() bool
	Put_IsDataTerminalReadyEnabled(value bool)
	Get_IsRequestToSendEnabled() bool
	Put_IsRequestToSendEnabled(value bool)
	Get_Parity() SerialParity
	Put_Parity(value SerialParity)
	Get_PortName() string
	Get_ReadTimeout() TimeSpan
	Put_ReadTimeout(value TimeSpan)
	Get_StopBits() SerialStopBitCount
	Put_StopBits(value SerialStopBitCount)
	Get_UsbVendorId() uint16
	Get_UsbProductId() uint16
	Get_WriteTimeout() TimeSpan
	Put_WriteTimeout(value TimeSpan)
	Get_InputStream() *IInputStream
	Get_OutputStream() *IOutputStream
	Add_ErrorReceived(reportHandler TypedEventHandler[*ISerialDevice, *IErrorReceivedEventArgs]) EventRegistrationToken
	Remove_ErrorReceived(token EventRegistrationToken)
	Add_PinChanged(reportHandler TypedEventHandler[*ISerialDevice, *IPinChangedEventArgs]) EventRegistrationToken
	Remove_PinChanged(token EventRegistrationToken)
}

type ISerialDeviceVtbl struct {
	win32.IInspectableVtbl
	Get_BaudRate                   uintptr
	Put_BaudRate                   uintptr
	Get_BreakSignalState           uintptr
	Put_BreakSignalState           uintptr
	Get_BytesReceived              uintptr
	Get_CarrierDetectState         uintptr
	Get_ClearToSendState           uintptr
	Get_DataBits                   uintptr
	Put_DataBits                   uintptr
	Get_DataSetReadyState          uintptr
	Get_Handshake                  uintptr
	Put_Handshake                  uintptr
	Get_IsDataTerminalReadyEnabled uintptr
	Put_IsDataTerminalReadyEnabled uintptr
	Get_IsRequestToSendEnabled     uintptr
	Put_IsRequestToSendEnabled     uintptr
	Get_Parity                     uintptr
	Put_Parity                     uintptr
	Get_PortName                   uintptr
	Get_ReadTimeout                uintptr
	Put_ReadTimeout                uintptr
	Get_StopBits                   uintptr
	Put_StopBits                   uintptr
	Get_UsbVendorId                uintptr
	Get_UsbProductId               uintptr
	Get_WriteTimeout               uintptr
	Put_WriteTimeout               uintptr
	Get_InputStream                uintptr
	Get_OutputStream               uintptr
	Add_ErrorReceived              uintptr
	Remove_ErrorReceived           uintptr
	Add_PinChanged                 uintptr
	Remove_PinChanged              uintptr
}

type ISerialDevice struct {
	win32.IInspectable
}

func (this *ISerialDevice) Vtbl() *ISerialDeviceVtbl {
	return (*ISerialDeviceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISerialDevice) Get_BaudRate() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BaudRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISerialDevice) Put_BaudRate(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BaudRate, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISerialDevice) Get_BreakSignalState() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BreakSignalState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISerialDevice) Put_BreakSignalState(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BreakSignalState, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ISerialDevice) Get_BytesReceived() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BytesReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISerialDevice) Get_CarrierDetectState() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CarrierDetectState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISerialDevice) Get_ClearToSendState() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ClearToSendState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISerialDevice) Get_DataBits() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DataBits, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISerialDevice) Put_DataBits(value uint16) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DataBits, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISerialDevice) Get_DataSetReadyState() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DataSetReadyState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISerialDevice) Get_Handshake() SerialHandshake {
	var _result SerialHandshake
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Handshake, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISerialDevice) Put_Handshake(value SerialHandshake) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Handshake, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISerialDevice) Get_IsDataTerminalReadyEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDataTerminalReadyEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISerialDevice) Put_IsDataTerminalReadyEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsDataTerminalReadyEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ISerialDevice) Get_IsRequestToSendEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsRequestToSendEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISerialDevice) Put_IsRequestToSendEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsRequestToSendEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ISerialDevice) Get_Parity() SerialParity {
	var _result SerialParity
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Parity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISerialDevice) Put_Parity(value SerialParity) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Parity, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISerialDevice) Get_PortName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PortName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISerialDevice) Get_ReadTimeout() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReadTimeout, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISerialDevice) Put_ReadTimeout(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReadTimeout, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ISerialDevice) Get_StopBits() SerialStopBitCount {
	var _result SerialStopBitCount
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StopBits, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISerialDevice) Put_StopBits(value SerialStopBitCount) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StopBits, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISerialDevice) Get_UsbVendorId() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UsbVendorId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISerialDevice) Get_UsbProductId() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UsbProductId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISerialDevice) Get_WriteTimeout() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WriteTimeout, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISerialDevice) Put_WriteTimeout(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_WriteTimeout, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ISerialDevice) Get_InputStream() *IInputStream {
	var _result *IInputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InputStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISerialDevice) Get_OutputStream() *IOutputStream {
	var _result *IOutputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutputStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISerialDevice) Add_ErrorReceived(reportHandler TypedEventHandler[*ISerialDevice, *IErrorReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ErrorReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(reportHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISerialDevice) Remove_ErrorReceived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ErrorReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISerialDevice) Add_PinChanged(reportHandler TypedEventHandler[*ISerialDevice, *IPinChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PinChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(reportHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISerialDevice) Remove_PinChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PinChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 058C4A70-0836-4993-AE1A-B61AE3BE056B
var IID_ISerialDeviceStatics = syscall.GUID{0x058C4A70, 0x0836, 0x4993,
	[8]byte{0xAE, 0x1A, 0xB6, 0x1A, 0xE3, 0xBE, 0x05, 0x6B}}

type ISerialDeviceStaticsInterface interface {
	win32.IInspectableInterface
	GetDeviceSelector() string
	GetDeviceSelectorFromPortName(portName string) string
	GetDeviceSelectorFromUsbVidPid(vendorId uint16, productId uint16) string
	FromIdAsync(deviceId string) *IAsyncOperation[*ISerialDevice]
}

type ISerialDeviceStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector              uintptr
	GetDeviceSelectorFromPortName  uintptr
	GetDeviceSelectorFromUsbVidPid uintptr
	FromIdAsync                    uintptr
}

type ISerialDeviceStatics struct {
	win32.IInspectable
}

func (this *ISerialDeviceStatics) Vtbl() *ISerialDeviceStaticsVtbl {
	return (*ISerialDeviceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISerialDeviceStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISerialDeviceStatics) GetDeviceSelectorFromPortName(portName string) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorFromPortName, uintptr(unsafe.Pointer(this)), NewHStr(portName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISerialDeviceStatics) GetDeviceSelectorFromUsbVidPid(vendorId uint16, productId uint16) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorFromUsbVidPid, uintptr(unsafe.Pointer(this)), uintptr(vendorId), uintptr(productId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISerialDeviceStatics) FromIdAsync(deviceId string) *IAsyncOperation[*ISerialDevice] {
	var _result *IAsyncOperation[*ISerialDevice]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type ErrorReceivedEventArgs struct {
	RtClass
	*IErrorReceivedEventArgs
}

type PinChangedEventArgs struct {
	RtClass
	*IPinChangedEventArgs
}

type SerialDevice struct {
	RtClass
	*ISerialDevice
}

func NewISerialDeviceStatics() *ISerialDeviceStatics {
	var p *ISerialDeviceStatics
	hs := NewHStr("Windows.Devices.SerialCommunication.SerialDevice")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISerialDeviceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
