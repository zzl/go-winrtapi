package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type NamedPolicyKind int32

const (
	NamedPolicyKind_Invalid NamedPolicyKind = 0
	NamedPolicyKind_Binary  NamedPolicyKind = 1
	NamedPolicyKind_Boolean NamedPolicyKind = 2
	NamedPolicyKind_Int32   NamedPolicyKind = 3
	NamedPolicyKind_Int64   NamedPolicyKind = 4
	NamedPolicyKind_String  NamedPolicyKind = 5
)

// interfaces

// 38DCB198-95AC-4077-A643-8078CAE26400
var IID_INamedPolicyData = syscall.GUID{0x38DCB198, 0x95AC, 0x4077,
	[8]byte{0xA6, 0x43, 0x80, 0x78, 0xCA, 0xE2, 0x64, 0x00}}

type INamedPolicyDataInterface interface {
	win32.IInspectableInterface
	Get_Area() string
	Get_Name() string
	Get_Kind() NamedPolicyKind
	Get_IsManaged() bool
	Get_IsUserPolicy() bool
	Get_User() *IUser
	GetBoolean() bool
	GetBinary() *IBuffer
	GetInt32() int32
	GetInt64() int64
	GetString() string
	Add_Changed(changedHandler TypedEventHandler[*INamedPolicyData, interface{}]) EventRegistrationToken
	Remove_Changed(cookie EventRegistrationToken)
}

type INamedPolicyDataVtbl struct {
	win32.IInspectableVtbl
	Get_Area         uintptr
	Get_Name         uintptr
	Get_Kind         uintptr
	Get_IsManaged    uintptr
	Get_IsUserPolicy uintptr
	Get_User         uintptr
	GetBoolean       uintptr
	GetBinary        uintptr
	GetInt32         uintptr
	GetInt64         uintptr
	GetString        uintptr
	Add_Changed      uintptr
	Remove_Changed   uintptr
}

type INamedPolicyData struct {
	win32.IInspectable
}

func (this *INamedPolicyData) Vtbl() *INamedPolicyDataVtbl {
	return (*INamedPolicyDataVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INamedPolicyData) Get_Area() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Area, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INamedPolicyData) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INamedPolicyData) Get_Kind() NamedPolicyKind {
	var _result NamedPolicyKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Kind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INamedPolicyData) Get_IsManaged() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsManaged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INamedPolicyData) Get_IsUserPolicy() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsUserPolicy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INamedPolicyData) Get_User() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_User, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INamedPolicyData) GetBoolean() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetBoolean, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INamedPolicyData) GetBinary() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetBinary, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INamedPolicyData) GetInt32() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetInt32, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INamedPolicyData) GetInt64() int64 {
	var _result int64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetInt64, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INamedPolicyData) GetString() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INamedPolicyData) Add_Changed(changedHandler TypedEventHandler[*INamedPolicyData, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Changed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(changedHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INamedPolicyData) Remove_Changed(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Changed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

// 7F793BE7-76C4-4058-8CAD-67662CD05F0D
var IID_INamedPolicyStatics = syscall.GUID{0x7F793BE7, 0x76C4, 0x4058,
	[8]byte{0x8C, 0xAD, 0x67, 0x66, 0x2C, 0xD0, 0x5F, 0x0D}}

type INamedPolicyStaticsInterface interface {
	win32.IInspectableInterface
	GetPolicyFromPath(area string, name string) *INamedPolicyData
	GetPolicyFromPathForUser(user *IUser, area string, name string) *INamedPolicyData
}

type INamedPolicyStaticsVtbl struct {
	win32.IInspectableVtbl
	GetPolicyFromPath        uintptr
	GetPolicyFromPathForUser uintptr
}

type INamedPolicyStatics struct {
	win32.IInspectable
}

func (this *INamedPolicyStatics) Vtbl() *INamedPolicyStaticsVtbl {
	return (*INamedPolicyStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INamedPolicyStatics) GetPolicyFromPath(area string, name string) *INamedPolicyData {
	var _result *INamedPolicyData
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPolicyFromPath, uintptr(unsafe.Pointer(this)), NewHStr(area).Ptr, NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INamedPolicyStatics) GetPolicyFromPathForUser(user *IUser, area string, name string) *INamedPolicyData {
	var _result *INamedPolicyData
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPolicyFromPathForUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), NewHStr(area).Ptr, NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type NamedPolicy struct {
	RtClass
}

func NewINamedPolicyStatics() *INamedPolicyStatics {
	var p *INamedPolicyStatics
	hs := NewHStr("Windows.Management.Policies.NamedPolicy")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_INamedPolicyStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type NamedPolicyData struct {
	RtClass
	*INamedPolicyData
}
