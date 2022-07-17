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
type MdmAlertDataType int32

const (
	MdmAlertDataType_String  MdmAlertDataType = 0
	MdmAlertDataType_Base64  MdmAlertDataType = 1
	MdmAlertDataType_Boolean MdmAlertDataType = 2
	MdmAlertDataType_Integer MdmAlertDataType = 3
)

// enum
type MdmAlertMark int32

const (
	MdmAlertMark_None          MdmAlertMark = 0
	MdmAlertMark_Fatal         MdmAlertMark = 1
	MdmAlertMark_Critical      MdmAlertMark = 2
	MdmAlertMark_Warning       MdmAlertMark = 3
	MdmAlertMark_Informational MdmAlertMark = 4
)

// enum
type MdmSessionState int32

const (
	MdmSessionState_NotStarted           MdmSessionState = 0
	MdmSessionState_Starting             MdmSessionState = 1
	MdmSessionState_Connecting           MdmSessionState = 2
	MdmSessionState_Communicating        MdmSessionState = 3
	MdmSessionState_AlertStatusAvailable MdmSessionState = 4
	MdmSessionState_Retrying             MdmSessionState = 5
	MdmSessionState_Completed            MdmSessionState = 6
)

// interfaces

// B0FBC327-28C1-4B52-A548-C5807CAF70B6
var IID_IMdmAlert = syscall.GUID{0xB0FBC327, 0x28C1, 0x4B52,
	[8]byte{0xA5, 0x48, 0xC5, 0x80, 0x7C, 0xAF, 0x70, 0xB6}}

type IMdmAlertInterface interface {
	win32.IInspectableInterface
	Get_Data() string
	Put_Data(value string)
	Get_Format() MdmAlertDataType
	Put_Format(value MdmAlertDataType)
	Get_Mark() MdmAlertMark
	Put_Mark(value MdmAlertMark)
	Get_Source() string
	Put_Source(value string)
	Get_Status() uint32
	Get_Target() string
	Put_Target(value string)
	Get_Type() string
	Put_Type(value string)
}

type IMdmAlertVtbl struct {
	win32.IInspectableVtbl
	Get_Data   uintptr
	Put_Data   uintptr
	Get_Format uintptr
	Put_Format uintptr
	Get_Mark   uintptr
	Put_Mark   uintptr
	Get_Source uintptr
	Put_Source uintptr
	Get_Status uintptr
	Get_Target uintptr
	Put_Target uintptr
	Get_Type   uintptr
	Put_Type   uintptr
}

type IMdmAlert struct {
	win32.IInspectable
}

func (this *IMdmAlert) Vtbl() *IMdmAlertVtbl {
	return (*IMdmAlertVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMdmAlert) Get_Data() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMdmAlert) Put_Data(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Data, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMdmAlert) Get_Format() MdmAlertDataType {
	var _result MdmAlertDataType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Format, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMdmAlert) Put_Format(value MdmAlertDataType) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Format, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMdmAlert) Get_Mark() MdmAlertMark {
	var _result MdmAlertMark
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mark, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMdmAlert) Put_Mark(value MdmAlertMark) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Mark, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMdmAlert) Get_Source() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMdmAlert) Put_Source(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Source, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMdmAlert) Get_Status() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMdmAlert) Get_Target() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Target, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMdmAlert) Put_Target(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Target, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMdmAlert) Get_Type() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Type, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMdmAlert) Put_Type(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Type, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// FE89314C-8F64-4797-A9D7-9D88F86AE166
var IID_IMdmSession = syscall.GUID{0xFE89314C, 0x8F64, 0x4797,
	[8]byte{0xA9, 0xD7, 0x9D, 0x88, 0xF8, 0x6A, 0xE1, 0x66}}

type IMdmSessionInterface interface {
	win32.IInspectableInterface
	Get_Alerts() *IVectorView[*IMdmAlert]
	Get_ExtendedError() HResult
	Get_Id() string
	Get_State() MdmSessionState
	AttachAsync() *IAsyncAction
	Delete()
	StartAsync() *IAsyncAction
	StartWithAlertsAsync(alerts *IIterable[*IMdmAlert]) *IAsyncAction
}

type IMdmSessionVtbl struct {
	win32.IInspectableVtbl
	Get_Alerts           uintptr
	Get_ExtendedError    uintptr
	Get_Id               uintptr
	Get_State            uintptr
	AttachAsync          uintptr
	Delete               uintptr
	StartAsync           uintptr
	StartWithAlertsAsync uintptr
}

type IMdmSession struct {
	win32.IInspectable
}

func (this *IMdmSession) Vtbl() *IMdmSessionVtbl {
	return (*IMdmSessionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMdmSession) Get_Alerts() *IVectorView[*IMdmAlert] {
	var _result *IVectorView[*IMdmAlert]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Alerts, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMdmSession) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMdmSession) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMdmSession) Get_State() MdmSessionState {
	var _result MdmSessionState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMdmSession) AttachAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AttachAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMdmSession) Delete() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Delete, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IMdmSession) StartAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMdmSession) StartWithAlertsAsync(alerts *IIterable[*IMdmAlert]) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartWithAlertsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(alerts)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CF4AD959-F745-4B79-9B5C-DE0BF8EFE44B
var IID_IMdmSessionManagerStatics = syscall.GUID{0xCF4AD959, 0xF745, 0x4B79,
	[8]byte{0x9B, 0x5C, 0xDE, 0x0B, 0xF8, 0xEF, 0xE4, 0x4B}}

type IMdmSessionManagerStaticsInterface interface {
	win32.IInspectableInterface
	Get_SessionIds() *IVectorView[string]
	TryCreateSession() *IMdmSession
	DeleteSessionById(sessionId string)
	GetSessionById(sessionId string) *IMdmSession
}

type IMdmSessionManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_SessionIds    uintptr
	TryCreateSession  uintptr
	DeleteSessionById uintptr
	GetSessionById    uintptr
}

type IMdmSessionManagerStatics struct {
	win32.IInspectable
}

func (this *IMdmSessionManagerStatics) Vtbl() *IMdmSessionManagerStaticsVtbl {
	return (*IMdmSessionManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMdmSessionManagerStatics) Get_SessionIds() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SessionIds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMdmSessionManagerStatics) TryCreateSession() *IMdmSession {
	var _result *IMdmSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryCreateSession, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMdmSessionManagerStatics) DeleteSessionById(sessionId string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DeleteSessionById, uintptr(unsafe.Pointer(this)), NewHStr(sessionId).Ptr)
	_ = _hr
}

func (this *IMdmSessionManagerStatics) GetSessionById(sessionId string) *IMdmSession {
	var _result *IMdmSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSessionById, uintptr(unsafe.Pointer(this)), NewHStr(sessionId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type MdmAlert struct {
	RtClass
	*IMdmAlert
}

func NewMdmAlert() *MdmAlert {
	hs := NewHStr("Windows.Management.MdmAlert")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &MdmAlert{
		RtClass:   RtClass{PInspect: p},
		IMdmAlert: (*IMdmAlert)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type MdmSession struct {
	RtClass
	*IMdmSession
}

type MdmSessionManager struct {
	RtClass
}

func NewIMdmSessionManagerStatics() *IMdmSessionManagerStatics {
	var p *IMdmSessionManagerStatics
	hs := NewHStr("Windows.Management.MdmSessionManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMdmSessionManagerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
