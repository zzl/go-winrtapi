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
type MidiMessageType int32

const (
	MidiMessageType_None                  MidiMessageType = 0
	MidiMessageType_NoteOff               MidiMessageType = 128
	MidiMessageType_NoteOn                MidiMessageType = 144
	MidiMessageType_PolyphonicKeyPressure MidiMessageType = 160
	MidiMessageType_ControlChange         MidiMessageType = 176
	MidiMessageType_ProgramChange         MidiMessageType = 192
	MidiMessageType_ChannelPressure       MidiMessageType = 208
	MidiMessageType_PitchBendChange       MidiMessageType = 224
	MidiMessageType_SystemExclusive       MidiMessageType = 240
	MidiMessageType_MidiTimeCode          MidiMessageType = 241
	MidiMessageType_SongPositionPointer   MidiMessageType = 242
	MidiMessageType_SongSelect            MidiMessageType = 243
	MidiMessageType_TuneRequest           MidiMessageType = 246
	MidiMessageType_EndSystemExclusive    MidiMessageType = 247
	MidiMessageType_TimingClock           MidiMessageType = 248
	MidiMessageType_Start                 MidiMessageType = 250
	MidiMessageType_Continue              MidiMessageType = 251
	MidiMessageType_Stop                  MidiMessageType = 252
	MidiMessageType_ActiveSensing         MidiMessageType = 254
	MidiMessageType_SystemReset           MidiMessageType = 255
)

// interfaces

// BE1FA860-62B4-4D52-A37E-92E54D35B909
var IID_IMidiChannelPressureMessage = syscall.GUID{0xBE1FA860, 0x62B4, 0x4D52,
	[8]byte{0xA3, 0x7E, 0x92, 0xE5, 0x4D, 0x35, 0xB9, 0x09}}

type IMidiChannelPressureMessageInterface interface {
	win32.IInspectableInterface
	Get_Channel() byte
	Get_Pressure() byte
}

type IMidiChannelPressureMessageVtbl struct {
	win32.IInspectableVtbl
	Get_Channel  uintptr
	Get_Pressure uintptr
}

type IMidiChannelPressureMessage struct {
	win32.IInspectable
}

func (this *IMidiChannelPressureMessage) Vtbl() *IMidiChannelPressureMessageVtbl {
	return (*IMidiChannelPressureMessageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiChannelPressureMessage) Get_Channel() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Channel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMidiChannelPressureMessage) Get_Pressure() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Pressure, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6218ED2F-2284-412A-94CF-10FB04842C6C
var IID_IMidiChannelPressureMessageFactory = syscall.GUID{0x6218ED2F, 0x2284, 0x412A,
	[8]byte{0x94, 0xCF, 0x10, 0xFB, 0x04, 0x84, 0x2C, 0x6C}}

type IMidiChannelPressureMessageFactoryInterface interface {
	win32.IInspectableInterface
	CreateMidiChannelPressureMessage(channel byte, pressure byte) *IMidiChannelPressureMessage
}

type IMidiChannelPressureMessageFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateMidiChannelPressureMessage uintptr
}

type IMidiChannelPressureMessageFactory struct {
	win32.IInspectable
}

func (this *IMidiChannelPressureMessageFactory) Vtbl() *IMidiChannelPressureMessageFactoryVtbl {
	return (*IMidiChannelPressureMessageFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiChannelPressureMessageFactory) CreateMidiChannelPressureMessage(channel byte, pressure byte) *IMidiChannelPressureMessage {
	var _result *IMidiChannelPressureMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateMidiChannelPressureMessage, uintptr(unsafe.Pointer(this)), uintptr(channel), uintptr(pressure), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B7E15F83-780D-405F-B781-3E1598C97F40
var IID_IMidiControlChangeMessage = syscall.GUID{0xB7E15F83, 0x780D, 0x405F,
	[8]byte{0xB7, 0x81, 0x3E, 0x15, 0x98, 0xC9, 0x7F, 0x40}}

type IMidiControlChangeMessageInterface interface {
	win32.IInspectableInterface
	Get_Channel() byte
	Get_Controller() byte
	Get_ControlValue() byte
}

type IMidiControlChangeMessageVtbl struct {
	win32.IInspectableVtbl
	Get_Channel      uintptr
	Get_Controller   uintptr
	Get_ControlValue uintptr
}

type IMidiControlChangeMessage struct {
	win32.IInspectable
}

func (this *IMidiControlChangeMessage) Vtbl() *IMidiControlChangeMessageVtbl {
	return (*IMidiControlChangeMessageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiControlChangeMessage) Get_Channel() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Channel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMidiControlChangeMessage) Get_Controller() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Controller, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMidiControlChangeMessage) Get_ControlValue() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ControlValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 2AB14321-956C-46AD-9752-F87F55052FE3
var IID_IMidiControlChangeMessageFactory = syscall.GUID{0x2AB14321, 0x956C, 0x46AD,
	[8]byte{0x97, 0x52, 0xF8, 0x7F, 0x55, 0x05, 0x2F, 0xE3}}

type IMidiControlChangeMessageFactoryInterface interface {
	win32.IInspectableInterface
	CreateMidiControlChangeMessage(channel byte, controller byte, controlValue byte) *IMidiControlChangeMessage
}

type IMidiControlChangeMessageFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateMidiControlChangeMessage uintptr
}

type IMidiControlChangeMessageFactory struct {
	win32.IInspectable
}

func (this *IMidiControlChangeMessageFactory) Vtbl() *IMidiControlChangeMessageFactoryVtbl {
	return (*IMidiControlChangeMessageFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiControlChangeMessageFactory) CreateMidiControlChangeMessage(channel byte, controller byte, controlValue byte) *IMidiControlChangeMessage {
	var _result *IMidiControlChangeMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateMidiControlChangeMessage, uintptr(unsafe.Pointer(this)), uintptr(channel), uintptr(controller), uintptr(controlValue), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D5C1D9DB-971A-4EAF-A23D-EA19FE607FF9
var IID_IMidiInPort = syscall.GUID{0xD5C1D9DB, 0x971A, 0x4EAF,
	[8]byte{0xA2, 0x3D, 0xEA, 0x19, 0xFE, 0x60, 0x7F, 0xF9}}

type IMidiInPortInterface interface {
	win32.IInspectableInterface
	Add_MessageReceived(handler TypedEventHandler[*IMidiInPort, *IMidiMessageReceivedEventArgs]) EventRegistrationToken
	Remove_MessageReceived(token EventRegistrationToken)
	Get_DeviceId() string
}

type IMidiInPortVtbl struct {
	win32.IInspectableVtbl
	Add_MessageReceived    uintptr
	Remove_MessageReceived uintptr
	Get_DeviceId           uintptr
}

type IMidiInPort struct {
	win32.IInspectable
}

func (this *IMidiInPort) Vtbl() *IMidiInPortVtbl {
	return (*IMidiInPortVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiInPort) Add_MessageReceived(handler TypedEventHandler[*IMidiInPort, *IMidiMessageReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_MessageReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMidiInPort) Remove_MessageReceived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_MessageReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMidiInPort) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 44C439DC-67FF-4A6E-8BAC-FDB6610CF296
var IID_IMidiInPortStatics = syscall.GUID{0x44C439DC, 0x67FF, 0x4A6E,
	[8]byte{0x8B, 0xAC, 0xFD, 0xB6, 0x61, 0x0C, 0xF2, 0x96}}

type IMidiInPortStaticsInterface interface {
	win32.IInspectableInterface
	FromIdAsync(deviceId string) *IAsyncOperation[*IMidiInPort]
	GetDeviceSelector() string
}

type IMidiInPortStaticsVtbl struct {
	win32.IInspectableVtbl
	FromIdAsync       uintptr
	GetDeviceSelector uintptr
}

type IMidiInPortStatics struct {
	win32.IInspectable
}

func (this *IMidiInPortStatics) Vtbl() *IMidiInPortStaticsVtbl {
	return (*IMidiInPortStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiInPortStatics) FromIdAsync(deviceId string) *IAsyncOperation[*IMidiInPort] {
	var _result *IAsyncOperation[*IMidiInPort]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMidiInPortStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 79767945-1094-4283-9BE0-289FC0EE8334
var IID_IMidiMessage = syscall.GUID{0x79767945, 0x1094, 0x4283,
	[8]byte{0x9B, 0xE0, 0x28, 0x9F, 0xC0, 0xEE, 0x83, 0x34}}

type IMidiMessageInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() TimeSpan
	Get_RawData() *IBuffer
	Get_Type() MidiMessageType
}

type IMidiMessageVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp uintptr
	Get_RawData   uintptr
	Get_Type      uintptr
}

type IMidiMessage struct {
	win32.IInspectable
}

func (this *IMidiMessage) Vtbl() *IMidiMessageVtbl {
	return (*IMidiMessageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiMessage) Get_Timestamp() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMidiMessage) Get_RawData() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RawData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMidiMessage) Get_Type() MidiMessageType {
	var _result MidiMessageType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Type, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 76566E56-F328-4B51-907D-B3A8CE96BF80
var IID_IMidiMessageReceivedEventArgs = syscall.GUID{0x76566E56, 0xF328, 0x4B51,
	[8]byte{0x90, 0x7D, 0xB3, 0xA8, 0xCE, 0x96, 0xBF, 0x80}}

type IMidiMessageReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Message() *IMidiMessage
}

type IMidiMessageReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Message uintptr
}

type IMidiMessageReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IMidiMessageReceivedEventArgs) Vtbl() *IMidiMessageReceivedEventArgsVtbl {
	return (*IMidiMessageReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiMessageReceivedEventArgs) Get_Message() *IMidiMessage {
	var _result *IMidiMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Message, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 16FD8AF4-198E-4D8F-A654-D305A293548F
var IID_IMidiNoteOffMessage = syscall.GUID{0x16FD8AF4, 0x198E, 0x4D8F,
	[8]byte{0xA6, 0x54, 0xD3, 0x05, 0xA2, 0x93, 0x54, 0x8F}}

type IMidiNoteOffMessageInterface interface {
	win32.IInspectableInterface
	Get_Channel() byte
	Get_Note() byte
	Get_Velocity() byte
}

type IMidiNoteOffMessageVtbl struct {
	win32.IInspectableVtbl
	Get_Channel  uintptr
	Get_Note     uintptr
	Get_Velocity uintptr
}

type IMidiNoteOffMessage struct {
	win32.IInspectable
}

func (this *IMidiNoteOffMessage) Vtbl() *IMidiNoteOffMessageVtbl {
	return (*IMidiNoteOffMessageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiNoteOffMessage) Get_Channel() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Channel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMidiNoteOffMessage) Get_Note() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Note, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMidiNoteOffMessage) Get_Velocity() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Velocity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A6B240E0-A749-425F-8AF4-A4D979CC15B5
var IID_IMidiNoteOffMessageFactory = syscall.GUID{0xA6B240E0, 0xA749, 0x425F,
	[8]byte{0x8A, 0xF4, 0xA4, 0xD9, 0x79, 0xCC, 0x15, 0xB5}}

type IMidiNoteOffMessageFactoryInterface interface {
	win32.IInspectableInterface
	CreateMidiNoteOffMessage(channel byte, note byte, velocity byte) *IMidiNoteOffMessage
}

type IMidiNoteOffMessageFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateMidiNoteOffMessage uintptr
}

type IMidiNoteOffMessageFactory struct {
	win32.IInspectable
}

func (this *IMidiNoteOffMessageFactory) Vtbl() *IMidiNoteOffMessageFactoryVtbl {
	return (*IMidiNoteOffMessageFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiNoteOffMessageFactory) CreateMidiNoteOffMessage(channel byte, note byte, velocity byte) *IMidiNoteOffMessage {
	var _result *IMidiNoteOffMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateMidiNoteOffMessage, uintptr(unsafe.Pointer(this)), uintptr(channel), uintptr(note), uintptr(velocity), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E0224AF5-6181-46DD-AFA2-410004C057AA
var IID_IMidiNoteOnMessage = syscall.GUID{0xE0224AF5, 0x6181, 0x46DD,
	[8]byte{0xAF, 0xA2, 0x41, 0x00, 0x04, 0xC0, 0x57, 0xAA}}

type IMidiNoteOnMessageInterface interface {
	win32.IInspectableInterface
	Get_Channel() byte
	Get_Note() byte
	Get_Velocity() byte
}

type IMidiNoteOnMessageVtbl struct {
	win32.IInspectableVtbl
	Get_Channel  uintptr
	Get_Note     uintptr
	Get_Velocity uintptr
}

type IMidiNoteOnMessage struct {
	win32.IInspectable
}

func (this *IMidiNoteOnMessage) Vtbl() *IMidiNoteOnMessageVtbl {
	return (*IMidiNoteOnMessageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiNoteOnMessage) Get_Channel() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Channel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMidiNoteOnMessage) Get_Note() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Note, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMidiNoteOnMessage) Get_Velocity() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Velocity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 9B4280A0-59C1-420E-B517-15A10AA9606B
var IID_IMidiNoteOnMessageFactory = syscall.GUID{0x9B4280A0, 0x59C1, 0x420E,
	[8]byte{0xB5, 0x17, 0x15, 0xA1, 0x0A, 0xA9, 0x60, 0x6B}}

type IMidiNoteOnMessageFactoryInterface interface {
	win32.IInspectableInterface
	CreateMidiNoteOnMessage(channel byte, note byte, velocity byte) *IMidiNoteOnMessage
}

type IMidiNoteOnMessageFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateMidiNoteOnMessage uintptr
}

type IMidiNoteOnMessageFactory struct {
	win32.IInspectable
}

func (this *IMidiNoteOnMessageFactory) Vtbl() *IMidiNoteOnMessageFactoryVtbl {
	return (*IMidiNoteOnMessageFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiNoteOnMessageFactory) CreateMidiNoteOnMessage(channel byte, note byte, velocity byte) *IMidiNoteOnMessage {
	var _result *IMidiNoteOnMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateMidiNoteOnMessage, uintptr(unsafe.Pointer(this)), uintptr(channel), uintptr(note), uintptr(velocity), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 931D6D9F-57A2-4A3A-ADB8-4640886F6693
var IID_IMidiOutPort = syscall.GUID{0x931D6D9F, 0x57A2, 0x4A3A,
	[8]byte{0xAD, 0xB8, 0x46, 0x40, 0x88, 0x6F, 0x66, 0x93}}

type IMidiOutPortInterface interface {
	win32.IInspectableInterface
	SendMessage(midiMessage *IMidiMessage)
	SendBuffer(midiData *IBuffer)
	Get_DeviceId() string
}

type IMidiOutPortVtbl struct {
	win32.IInspectableVtbl
	SendMessage  uintptr
	SendBuffer   uintptr
	Get_DeviceId uintptr
}

type IMidiOutPort struct {
	win32.IInspectable
}

func (this *IMidiOutPort) Vtbl() *IMidiOutPortVtbl {
	return (*IMidiOutPortVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiOutPort) SendMessage(midiMessage *IMidiMessage) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(midiMessage)))
	_ = _hr
}

func (this *IMidiOutPort) SendBuffer(midiData *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(midiData)))
	_ = _hr
}

func (this *IMidiOutPort) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 065CC3E9-0F88-448B-9B64-A95826C65B8F
var IID_IMidiOutPortStatics = syscall.GUID{0x065CC3E9, 0x0F88, 0x448B,
	[8]byte{0x9B, 0x64, 0xA9, 0x58, 0x26, 0xC6, 0x5B, 0x8F}}

type IMidiOutPortStaticsInterface interface {
	win32.IInspectableInterface
	FromIdAsync(deviceId string) *IAsyncOperation[*IMidiOutPort]
	GetDeviceSelector() string
}

type IMidiOutPortStaticsVtbl struct {
	win32.IInspectableVtbl
	FromIdAsync       uintptr
	GetDeviceSelector uintptr
}

type IMidiOutPortStatics struct {
	win32.IInspectable
}

func (this *IMidiOutPortStatics) Vtbl() *IMidiOutPortStaticsVtbl {
	return (*IMidiOutPortStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiOutPortStatics) FromIdAsync(deviceId string) *IAsyncOperation[*IMidiOutPort] {
	var _result *IAsyncOperation[*IMidiOutPort]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMidiOutPortStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 29DF4CB1-2E9F-4FAF-8C2B-9CB82A9079CA
var IID_IMidiPitchBendChangeMessage = syscall.GUID{0x29DF4CB1, 0x2E9F, 0x4FAF,
	[8]byte{0x8C, 0x2B, 0x9C, 0xB8, 0x2A, 0x90, 0x79, 0xCA}}

type IMidiPitchBendChangeMessageInterface interface {
	win32.IInspectableInterface
	Get_Channel() byte
	Get_Bend() uint16
}

type IMidiPitchBendChangeMessageVtbl struct {
	win32.IInspectableVtbl
	Get_Channel uintptr
	Get_Bend    uintptr
}

type IMidiPitchBendChangeMessage struct {
	win32.IInspectable
}

func (this *IMidiPitchBendChangeMessage) Vtbl() *IMidiPitchBendChangeMessageVtbl {
	return (*IMidiPitchBendChangeMessageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiPitchBendChangeMessage) Get_Channel() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Channel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMidiPitchBendChangeMessage) Get_Bend() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bend, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F5EEDF55-CFC8-4926-B30E-A3622393306C
var IID_IMidiPitchBendChangeMessageFactory = syscall.GUID{0xF5EEDF55, 0xCFC8, 0x4926,
	[8]byte{0xB3, 0x0E, 0xA3, 0x62, 0x23, 0x93, 0x30, 0x6C}}

type IMidiPitchBendChangeMessageFactoryInterface interface {
	win32.IInspectableInterface
	CreateMidiPitchBendChangeMessage(channel byte, bend uint16) *IMidiPitchBendChangeMessage
}

type IMidiPitchBendChangeMessageFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateMidiPitchBendChangeMessage uintptr
}

type IMidiPitchBendChangeMessageFactory struct {
	win32.IInspectable
}

func (this *IMidiPitchBendChangeMessageFactory) Vtbl() *IMidiPitchBendChangeMessageFactoryVtbl {
	return (*IMidiPitchBendChangeMessageFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiPitchBendChangeMessageFactory) CreateMidiPitchBendChangeMessage(channel byte, bend uint16) *IMidiPitchBendChangeMessage {
	var _result *IMidiPitchBendChangeMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateMidiPitchBendChangeMessage, uintptr(unsafe.Pointer(this)), uintptr(channel), uintptr(bend), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1F7337FE-ACE8-48A0-868E-7CDBF20F04D6
var IID_IMidiPolyphonicKeyPressureMessage = syscall.GUID{0x1F7337FE, 0xACE8, 0x48A0,
	[8]byte{0x86, 0x8E, 0x7C, 0xDB, 0xF2, 0x0F, 0x04, 0xD6}}

type IMidiPolyphonicKeyPressureMessageInterface interface {
	win32.IInspectableInterface
	Get_Channel() byte
	Get_Note() byte
	Get_Pressure() byte
}

type IMidiPolyphonicKeyPressureMessageVtbl struct {
	win32.IInspectableVtbl
	Get_Channel  uintptr
	Get_Note     uintptr
	Get_Pressure uintptr
}

type IMidiPolyphonicKeyPressureMessage struct {
	win32.IInspectable
}

func (this *IMidiPolyphonicKeyPressureMessage) Vtbl() *IMidiPolyphonicKeyPressureMessageVtbl {
	return (*IMidiPolyphonicKeyPressureMessageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiPolyphonicKeyPressureMessage) Get_Channel() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Channel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMidiPolyphonicKeyPressureMessage) Get_Note() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Note, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMidiPolyphonicKeyPressureMessage) Get_Pressure() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Pressure, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// E98F483E-C4B3-4DD2-917C-E349815A1B3B
var IID_IMidiPolyphonicKeyPressureMessageFactory = syscall.GUID{0xE98F483E, 0xC4B3, 0x4DD2,
	[8]byte{0x91, 0x7C, 0xE3, 0x49, 0x81, 0x5A, 0x1B, 0x3B}}

type IMidiPolyphonicKeyPressureMessageFactoryInterface interface {
	win32.IInspectableInterface
	CreateMidiPolyphonicKeyPressureMessage(channel byte, note byte, pressure byte) *IMidiPolyphonicKeyPressureMessage
}

type IMidiPolyphonicKeyPressureMessageFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateMidiPolyphonicKeyPressureMessage uintptr
}

type IMidiPolyphonicKeyPressureMessageFactory struct {
	win32.IInspectable
}

func (this *IMidiPolyphonicKeyPressureMessageFactory) Vtbl() *IMidiPolyphonicKeyPressureMessageFactoryVtbl {
	return (*IMidiPolyphonicKeyPressureMessageFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiPolyphonicKeyPressureMessageFactory) CreateMidiPolyphonicKeyPressureMessage(channel byte, note byte, pressure byte) *IMidiPolyphonicKeyPressureMessage {
	var _result *IMidiPolyphonicKeyPressureMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateMidiPolyphonicKeyPressureMessage, uintptr(unsafe.Pointer(this)), uintptr(channel), uintptr(note), uintptr(pressure), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9CBB3C78-7A3E-4327-AA98-20B8E4485AF8
var IID_IMidiProgramChangeMessage = syscall.GUID{0x9CBB3C78, 0x7A3E, 0x4327,
	[8]byte{0xAA, 0x98, 0x20, 0xB8, 0xE4, 0x48, 0x5A, 0xF8}}

type IMidiProgramChangeMessageInterface interface {
	win32.IInspectableInterface
	Get_Channel() byte
	Get_Program() byte
}

type IMidiProgramChangeMessageVtbl struct {
	win32.IInspectableVtbl
	Get_Channel uintptr
	Get_Program uintptr
}

type IMidiProgramChangeMessage struct {
	win32.IInspectable
}

func (this *IMidiProgramChangeMessage) Vtbl() *IMidiProgramChangeMessageVtbl {
	return (*IMidiProgramChangeMessageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiProgramChangeMessage) Get_Channel() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Channel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMidiProgramChangeMessage) Get_Program() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Program, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D6B04387-524B-4104-9C99-6572BFD2E261
var IID_IMidiProgramChangeMessageFactory = syscall.GUID{0xD6B04387, 0x524B, 0x4104,
	[8]byte{0x9C, 0x99, 0x65, 0x72, 0xBF, 0xD2, 0xE2, 0x61}}

type IMidiProgramChangeMessageFactoryInterface interface {
	win32.IInspectableInterface
	CreateMidiProgramChangeMessage(channel byte, program byte) *IMidiProgramChangeMessage
}

type IMidiProgramChangeMessageFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateMidiProgramChangeMessage uintptr
}

type IMidiProgramChangeMessageFactory struct {
	win32.IInspectable
}

func (this *IMidiProgramChangeMessageFactory) Vtbl() *IMidiProgramChangeMessageFactoryVtbl {
	return (*IMidiProgramChangeMessageFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiProgramChangeMessageFactory) CreateMidiProgramChangeMessage(channel byte, program byte) *IMidiProgramChangeMessage {
	var _result *IMidiProgramChangeMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateMidiProgramChangeMessage, uintptr(unsafe.Pointer(this)), uintptr(channel), uintptr(program), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4CA50C56-EC5E-4AE4-A115-88DC57CC2B79
var IID_IMidiSongPositionPointerMessage = syscall.GUID{0x4CA50C56, 0xEC5E, 0x4AE4,
	[8]byte{0xA1, 0x15, 0x88, 0xDC, 0x57, 0xCC, 0x2B, 0x79}}

type IMidiSongPositionPointerMessageInterface interface {
	win32.IInspectableInterface
	Get_Beats() uint16
}

type IMidiSongPositionPointerMessageVtbl struct {
	win32.IInspectableVtbl
	Get_Beats uintptr
}

type IMidiSongPositionPointerMessage struct {
	win32.IInspectable
}

func (this *IMidiSongPositionPointerMessage) Vtbl() *IMidiSongPositionPointerMessageVtbl {
	return (*IMidiSongPositionPointerMessageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiSongPositionPointerMessage) Get_Beats() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Beats, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 9C00E996-F10B-4FEA-B395-F5D6CF80F64E
var IID_IMidiSongPositionPointerMessageFactory = syscall.GUID{0x9C00E996, 0xF10B, 0x4FEA,
	[8]byte{0xB3, 0x95, 0xF5, 0xD6, 0xCF, 0x80, 0xF6, 0x4E}}

type IMidiSongPositionPointerMessageFactoryInterface interface {
	win32.IInspectableInterface
	CreateMidiSongPositionPointerMessage(beats uint16) *IMidiSongPositionPointerMessage
}

type IMidiSongPositionPointerMessageFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateMidiSongPositionPointerMessage uintptr
}

type IMidiSongPositionPointerMessageFactory struct {
	win32.IInspectable
}

func (this *IMidiSongPositionPointerMessageFactory) Vtbl() *IMidiSongPositionPointerMessageFactoryVtbl {
	return (*IMidiSongPositionPointerMessageFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiSongPositionPointerMessageFactory) CreateMidiSongPositionPointerMessage(beats uint16) *IMidiSongPositionPointerMessage {
	var _result *IMidiSongPositionPointerMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateMidiSongPositionPointerMessage, uintptr(unsafe.Pointer(this)), uintptr(beats), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 49F0F27F-6D83-4741-A5BF-4629F6BE974F
var IID_IMidiSongSelectMessage = syscall.GUID{0x49F0F27F, 0x6D83, 0x4741,
	[8]byte{0xA5, 0xBF, 0x46, 0x29, 0xF6, 0xBE, 0x97, 0x4F}}

type IMidiSongSelectMessageInterface interface {
	win32.IInspectableInterface
	Get_Song() byte
}

type IMidiSongSelectMessageVtbl struct {
	win32.IInspectableVtbl
	Get_Song uintptr
}

type IMidiSongSelectMessage struct {
	win32.IInspectable
}

func (this *IMidiSongSelectMessage) Vtbl() *IMidiSongSelectMessageVtbl {
	return (*IMidiSongSelectMessageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiSongSelectMessage) Get_Song() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Song, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 848878E4-8748-4129-A66C-A05493F75DAA
var IID_IMidiSongSelectMessageFactory = syscall.GUID{0x848878E4, 0x8748, 0x4129,
	[8]byte{0xA6, 0x6C, 0xA0, 0x54, 0x93, 0xF7, 0x5D, 0xAA}}

type IMidiSongSelectMessageFactoryInterface interface {
	win32.IInspectableInterface
	CreateMidiSongSelectMessage(song byte) *IMidiSongSelectMessage
}

type IMidiSongSelectMessageFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateMidiSongSelectMessage uintptr
}

type IMidiSongSelectMessageFactory struct {
	win32.IInspectable
}

func (this *IMidiSongSelectMessageFactory) Vtbl() *IMidiSongSelectMessageFactoryVtbl {
	return (*IMidiSongSelectMessageFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiSongSelectMessageFactory) CreateMidiSongSelectMessage(song byte) *IMidiSongSelectMessage {
	var _result *IMidiSongSelectMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateMidiSongSelectMessage, uintptr(unsafe.Pointer(this)), uintptr(song), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F0DA155E-DB90-405F-B8AE-21D2E17F2E45
var IID_IMidiSynthesizer = syscall.GUID{0xF0DA155E, 0xDB90, 0x405F,
	[8]byte{0xB8, 0xAE, 0x21, 0xD2, 0xE1, 0x7F, 0x2E, 0x45}}

type IMidiSynthesizerInterface interface {
	win32.IInspectableInterface
	Get_AudioDevice() *IDeviceInformation
	Get_Volume() float64
	Put_Volume(value float64)
}

type IMidiSynthesizerVtbl struct {
	win32.IInspectableVtbl
	Get_AudioDevice uintptr
	Get_Volume      uintptr
	Put_Volume      uintptr
}

type IMidiSynthesizer struct {
	win32.IInspectable
}

func (this *IMidiSynthesizer) Vtbl() *IMidiSynthesizerVtbl {
	return (*IMidiSynthesizerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiSynthesizer) Get_AudioDevice() *IDeviceInformation {
	var _result *IDeviceInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioDevice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMidiSynthesizer) Get_Volume() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Volume, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMidiSynthesizer) Put_Volume(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Volume, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 4224EAA8-6629-4D6B-AA8F-D4521A5A31CE
var IID_IMidiSynthesizerStatics = syscall.GUID{0x4224EAA8, 0x6629, 0x4D6B,
	[8]byte{0xAA, 0x8F, 0xD4, 0x52, 0x1A, 0x5A, 0x31, 0xCE}}

type IMidiSynthesizerStaticsInterface interface {
	win32.IInspectableInterface
	CreateAsync() *IAsyncOperation[*IMidiSynthesizer]
	CreateFromAudioDeviceAsync(audioDevice *IDeviceInformation) *IAsyncOperation[*IMidiSynthesizer]
	IsSynthesizer(midiDevice *IDeviceInformation) bool
}

type IMidiSynthesizerStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateAsync                uintptr
	CreateFromAudioDeviceAsync uintptr
	IsSynthesizer              uintptr
}

type IMidiSynthesizerStatics struct {
	win32.IInspectable
}

func (this *IMidiSynthesizerStatics) Vtbl() *IMidiSynthesizerStaticsVtbl {
	return (*IMidiSynthesizerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiSynthesizerStatics) CreateAsync() *IAsyncOperation[*IMidiSynthesizer] {
	var _result *IAsyncOperation[*IMidiSynthesizer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMidiSynthesizerStatics) CreateFromAudioDeviceAsync(audioDevice *IDeviceInformation) *IAsyncOperation[*IMidiSynthesizer] {
	var _result *IAsyncOperation[*IMidiSynthesizer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromAudioDeviceAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(audioDevice)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMidiSynthesizerStatics) IsSynthesizer(midiDevice *IDeviceInformation) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsSynthesizer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(midiDevice)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 083DE222-3B74-4320-9B42-0CA8545F8A24
var IID_IMidiSystemExclusiveMessageFactory = syscall.GUID{0x083DE222, 0x3B74, 0x4320,
	[8]byte{0x9B, 0x42, 0x0C, 0xA8, 0x54, 0x5F, 0x8A, 0x24}}

type IMidiSystemExclusiveMessageFactoryInterface interface {
	win32.IInspectableInterface
	CreateMidiSystemExclusiveMessage(rawData *IBuffer) *IMidiMessage
}

type IMidiSystemExclusiveMessageFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateMidiSystemExclusiveMessage uintptr
}

type IMidiSystemExclusiveMessageFactory struct {
	win32.IInspectable
}

func (this *IMidiSystemExclusiveMessageFactory) Vtbl() *IMidiSystemExclusiveMessageFactoryVtbl {
	return (*IMidiSystemExclusiveMessageFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiSystemExclusiveMessageFactory) CreateMidiSystemExclusiveMessage(rawData *IBuffer) *IMidiMessage {
	var _result *IMidiMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateMidiSystemExclusiveMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(rawData)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0BF7087D-FA63-4A1C-8DEB-C0E87796A6D7
var IID_IMidiTimeCodeMessage = syscall.GUID{0x0BF7087D, 0xFA63, 0x4A1C,
	[8]byte{0x8D, 0xEB, 0xC0, 0xE8, 0x77, 0x96, 0xA6, 0xD7}}

type IMidiTimeCodeMessageInterface interface {
	win32.IInspectableInterface
	Get_FrameType() byte
	Get_Values() byte
}

type IMidiTimeCodeMessageVtbl struct {
	win32.IInspectableVtbl
	Get_FrameType uintptr
	Get_Values    uintptr
}

type IMidiTimeCodeMessage struct {
	win32.IInspectable
}

func (this *IMidiTimeCodeMessage) Vtbl() *IMidiTimeCodeMessageVtbl {
	return (*IMidiTimeCodeMessageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiTimeCodeMessage) Get_FrameType() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMidiTimeCodeMessage) Get_Values() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Values, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// EB3099C5-771C-40DE-B961-175A7489A85E
var IID_IMidiTimeCodeMessageFactory = syscall.GUID{0xEB3099C5, 0x771C, 0x40DE,
	[8]byte{0xB9, 0x61, 0x17, 0x5A, 0x74, 0x89, 0xA8, 0x5E}}

type IMidiTimeCodeMessageFactoryInterface interface {
	win32.IInspectableInterface
	CreateMidiTimeCodeMessage(frameType byte, values byte) *IMidiTimeCodeMessage
}

type IMidiTimeCodeMessageFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateMidiTimeCodeMessage uintptr
}

type IMidiTimeCodeMessageFactory struct {
	win32.IInspectable
}

func (this *IMidiTimeCodeMessageFactory) Vtbl() *IMidiTimeCodeMessageFactoryVtbl {
	return (*IMidiTimeCodeMessageFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMidiTimeCodeMessageFactory) CreateMidiTimeCodeMessage(frameType byte, values byte) *IMidiTimeCodeMessage {
	var _result *IMidiTimeCodeMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateMidiTimeCodeMessage, uintptr(unsafe.Pointer(this)), uintptr(frameType), uintptr(values), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type MidiActiveSensingMessage struct {
	RtClass
	*IMidiMessage
}

func NewMidiActiveSensingMessage() *MidiActiveSensingMessage {
	hs := NewHStr("Windows.Devices.Midi.MidiActiveSensingMessage")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &MidiActiveSensingMessage{
		RtClass:      RtClass{PInspect: p},
		IMidiMessage: (*IMidiMessage)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type MidiChannelPressureMessage struct {
	RtClass
	*IMidiChannelPressureMessage
}

func NewMidiChannelPressureMessage_CreateMidiChannelPressureMessage(channel byte, pressure byte) *MidiChannelPressureMessage {
	hs := NewHStr("Windows.Devices.Midi.MidiChannelPressureMessage")
	var pFac *IMidiChannelPressureMessageFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMidiChannelPressureMessageFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IMidiChannelPressureMessage
	p = pFac.CreateMidiChannelPressureMessage(channel, pressure)
	result := &MidiChannelPressureMessage{
		RtClass:                     RtClass{PInspect: &p.IInspectable},
		IMidiChannelPressureMessage: p,
	}
	com.AddToScope(result)
	return result
}

type MidiContinueMessage struct {
	RtClass
	*IMidiMessage
}

func NewMidiContinueMessage() *MidiContinueMessage {
	hs := NewHStr("Windows.Devices.Midi.MidiContinueMessage")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &MidiContinueMessage{
		RtClass:      RtClass{PInspect: p},
		IMidiMessage: (*IMidiMessage)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type MidiControlChangeMessage struct {
	RtClass
	*IMidiControlChangeMessage
}

func NewMidiControlChangeMessage_CreateMidiControlChangeMessage(channel byte, controller byte, controlValue byte) *MidiControlChangeMessage {
	hs := NewHStr("Windows.Devices.Midi.MidiControlChangeMessage")
	var pFac *IMidiControlChangeMessageFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMidiControlChangeMessageFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IMidiControlChangeMessage
	p = pFac.CreateMidiControlChangeMessage(channel, controller, controlValue)
	result := &MidiControlChangeMessage{
		RtClass:                   RtClass{PInspect: &p.IInspectable},
		IMidiControlChangeMessage: p,
	}
	com.AddToScope(result)
	return result
}

type MidiInPort struct {
	RtClass
	*IMidiInPort
}

func NewIMidiInPortStatics() *IMidiInPortStatics {
	var p *IMidiInPortStatics
	hs := NewHStr("Windows.Devices.Midi.MidiInPort")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMidiInPortStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type MidiMessageReceivedEventArgs struct {
	RtClass
	*IMidiMessageReceivedEventArgs
}

type MidiNoteOffMessage struct {
	RtClass
	*IMidiNoteOffMessage
}

func NewMidiNoteOffMessage_CreateMidiNoteOffMessage(channel byte, note byte, velocity byte) *MidiNoteOffMessage {
	hs := NewHStr("Windows.Devices.Midi.MidiNoteOffMessage")
	var pFac *IMidiNoteOffMessageFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMidiNoteOffMessageFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IMidiNoteOffMessage
	p = pFac.CreateMidiNoteOffMessage(channel, note, velocity)
	result := &MidiNoteOffMessage{
		RtClass:             RtClass{PInspect: &p.IInspectable},
		IMidiNoteOffMessage: p,
	}
	com.AddToScope(result)
	return result
}

type MidiNoteOnMessage struct {
	RtClass
	*IMidiNoteOnMessage
}

func NewMidiNoteOnMessage_CreateMidiNoteOnMessage(channel byte, note byte, velocity byte) *MidiNoteOnMessage {
	hs := NewHStr("Windows.Devices.Midi.MidiNoteOnMessage")
	var pFac *IMidiNoteOnMessageFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMidiNoteOnMessageFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IMidiNoteOnMessage
	p = pFac.CreateMidiNoteOnMessage(channel, note, velocity)
	result := &MidiNoteOnMessage{
		RtClass:            RtClass{PInspect: &p.IInspectable},
		IMidiNoteOnMessage: p,
	}
	com.AddToScope(result)
	return result
}

type MidiOutPort struct {
	RtClass
	*IMidiOutPort
}

func NewIMidiOutPortStatics() *IMidiOutPortStatics {
	var p *IMidiOutPortStatics
	hs := NewHStr("Windows.Devices.Midi.MidiOutPort")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMidiOutPortStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type MidiPitchBendChangeMessage struct {
	RtClass
	*IMidiPitchBendChangeMessage
}

func NewMidiPitchBendChangeMessage_CreateMidiPitchBendChangeMessage(channel byte, bend uint16) *MidiPitchBendChangeMessage {
	hs := NewHStr("Windows.Devices.Midi.MidiPitchBendChangeMessage")
	var pFac *IMidiPitchBendChangeMessageFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMidiPitchBendChangeMessageFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IMidiPitchBendChangeMessage
	p = pFac.CreateMidiPitchBendChangeMessage(channel, bend)
	result := &MidiPitchBendChangeMessage{
		RtClass:                     RtClass{PInspect: &p.IInspectable},
		IMidiPitchBendChangeMessage: p,
	}
	com.AddToScope(result)
	return result
}

type MidiPolyphonicKeyPressureMessage struct {
	RtClass
	*IMidiPolyphonicKeyPressureMessage
}

func NewMidiPolyphonicKeyPressureMessage_CreateMidiPolyphonicKeyPressureMessage(channel byte, note byte, pressure byte) *MidiPolyphonicKeyPressureMessage {
	hs := NewHStr("Windows.Devices.Midi.MidiPolyphonicKeyPressureMessage")
	var pFac *IMidiPolyphonicKeyPressureMessageFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMidiPolyphonicKeyPressureMessageFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IMidiPolyphonicKeyPressureMessage
	p = pFac.CreateMidiPolyphonicKeyPressureMessage(channel, note, pressure)
	result := &MidiPolyphonicKeyPressureMessage{
		RtClass:                           RtClass{PInspect: &p.IInspectable},
		IMidiPolyphonicKeyPressureMessage: p,
	}
	com.AddToScope(result)
	return result
}

type MidiProgramChangeMessage struct {
	RtClass
	*IMidiProgramChangeMessage
}

func NewMidiProgramChangeMessage_CreateMidiProgramChangeMessage(channel byte, program byte) *MidiProgramChangeMessage {
	hs := NewHStr("Windows.Devices.Midi.MidiProgramChangeMessage")
	var pFac *IMidiProgramChangeMessageFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMidiProgramChangeMessageFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IMidiProgramChangeMessage
	p = pFac.CreateMidiProgramChangeMessage(channel, program)
	result := &MidiProgramChangeMessage{
		RtClass:                   RtClass{PInspect: &p.IInspectable},
		IMidiProgramChangeMessage: p,
	}
	com.AddToScope(result)
	return result
}

type MidiSongPositionPointerMessage struct {
	RtClass
	*IMidiSongPositionPointerMessage
}

func NewMidiSongPositionPointerMessage_CreateMidiSongPositionPointerMessage(beats uint16) *MidiSongPositionPointerMessage {
	hs := NewHStr("Windows.Devices.Midi.MidiSongPositionPointerMessage")
	var pFac *IMidiSongPositionPointerMessageFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMidiSongPositionPointerMessageFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IMidiSongPositionPointerMessage
	p = pFac.CreateMidiSongPositionPointerMessage(beats)
	result := &MidiSongPositionPointerMessage{
		RtClass:                         RtClass{PInspect: &p.IInspectable},
		IMidiSongPositionPointerMessage: p,
	}
	com.AddToScope(result)
	return result
}

type MidiSongSelectMessage struct {
	RtClass
	*IMidiSongSelectMessage
}

func NewMidiSongSelectMessage_CreateMidiSongSelectMessage(song byte) *MidiSongSelectMessage {
	hs := NewHStr("Windows.Devices.Midi.MidiSongSelectMessage")
	var pFac *IMidiSongSelectMessageFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMidiSongSelectMessageFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IMidiSongSelectMessage
	p = pFac.CreateMidiSongSelectMessage(song)
	result := &MidiSongSelectMessage{
		RtClass:                RtClass{PInspect: &p.IInspectable},
		IMidiSongSelectMessage: p,
	}
	com.AddToScope(result)
	return result
}

type MidiStartMessage struct {
	RtClass
	*IMidiMessage
}

func NewMidiStartMessage() *MidiStartMessage {
	hs := NewHStr("Windows.Devices.Midi.MidiStartMessage")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &MidiStartMessage{
		RtClass:      RtClass{PInspect: p},
		IMidiMessage: (*IMidiMessage)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type MidiStopMessage struct {
	RtClass
	*IMidiMessage
}

func NewMidiStopMessage() *MidiStopMessage {
	hs := NewHStr("Windows.Devices.Midi.MidiStopMessage")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &MidiStopMessage{
		RtClass:      RtClass{PInspect: p},
		IMidiMessage: (*IMidiMessage)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type MidiSynthesizer struct {
	RtClass
	*IMidiSynthesizer
}

func NewIMidiSynthesizerStatics() *IMidiSynthesizerStatics {
	var p *IMidiSynthesizerStatics
	hs := NewHStr("Windows.Devices.Midi.MidiSynthesizer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMidiSynthesizerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type MidiSystemExclusiveMessage struct {
	RtClass
	*IMidiMessage
}

func NewMidiSystemExclusiveMessage_CreateMidiSystemExclusiveMessage(rawData *IBuffer) *MidiSystemExclusiveMessage {
	hs := NewHStr("Windows.Devices.Midi.MidiSystemExclusiveMessage")
	var pFac *IMidiSystemExclusiveMessageFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMidiSystemExclusiveMessageFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IMidiMessage
	p = pFac.CreateMidiSystemExclusiveMessage(rawData)
	result := &MidiSystemExclusiveMessage{
		RtClass:      RtClass{PInspect: &p.IInspectable},
		IMidiMessage: p,
	}
	com.AddToScope(result)
	return result
}

type MidiSystemResetMessage struct {
	RtClass
	*IMidiMessage
}

func NewMidiSystemResetMessage() *MidiSystemResetMessage {
	hs := NewHStr("Windows.Devices.Midi.MidiSystemResetMessage")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &MidiSystemResetMessage{
		RtClass:      RtClass{PInspect: p},
		IMidiMessage: (*IMidiMessage)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type MidiTimeCodeMessage struct {
	RtClass
	*IMidiTimeCodeMessage
}

func NewMidiTimeCodeMessage_CreateMidiTimeCodeMessage(frameType byte, values byte) *MidiTimeCodeMessage {
	hs := NewHStr("Windows.Devices.Midi.MidiTimeCodeMessage")
	var pFac *IMidiTimeCodeMessageFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMidiTimeCodeMessageFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IMidiTimeCodeMessage
	p = pFac.CreateMidiTimeCodeMessage(frameType, values)
	result := &MidiTimeCodeMessage{
		RtClass:              RtClass{PInspect: &p.IInspectable},
		IMidiTimeCodeMessage: p,
	}
	com.AddToScope(result)
	return result
}

type MidiTimingClockMessage struct {
	RtClass
	*IMidiMessage
}

func NewMidiTimingClockMessage() *MidiTimingClockMessage {
	hs := NewHStr("Windows.Devices.Midi.MidiTimingClockMessage")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &MidiTimingClockMessage{
		RtClass:      RtClass{PInspect: p},
		IMidiMessage: (*IMidiMessage)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type MidiTuneRequestMessage struct {
	RtClass
	*IMidiMessage
}

func NewMidiTuneRequestMessage() *MidiTuneRequestMessage {
	hs := NewHStr("Windows.Devices.Midi.MidiTuneRequestMessage")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &MidiTuneRequestMessage{
		RtClass:      RtClass{PInspect: p},
		IMidiMessage: (*IMidiMessage)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}
