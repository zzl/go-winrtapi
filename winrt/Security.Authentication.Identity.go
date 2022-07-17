package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// interfaces

// 38321ACC-672B-4823-B603-6B3C753DAF97
var IID_IEnterpriseKeyCredentialRegistrationInfo = syscall.GUID{0x38321ACC, 0x672B, 0x4823,
	[8]byte{0xB6, 0x03, 0x6B, 0x3C, 0x75, 0x3D, 0xAF, 0x97}}

type IEnterpriseKeyCredentialRegistrationInfoInterface interface {
	win32.IInspectableInterface
	Get_TenantId() string
	Get_TenantName() string
	Get_Subject() string
	Get_KeyId() string
	Get_KeyName() string
}

type IEnterpriseKeyCredentialRegistrationInfoVtbl struct {
	win32.IInspectableVtbl
	Get_TenantId   uintptr
	Get_TenantName uintptr
	Get_Subject    uintptr
	Get_KeyId      uintptr
	Get_KeyName    uintptr
}

type IEnterpriseKeyCredentialRegistrationInfo struct {
	win32.IInspectable
}

func (this *IEnterpriseKeyCredentialRegistrationInfo) Vtbl() *IEnterpriseKeyCredentialRegistrationInfoVtbl {
	return (*IEnterpriseKeyCredentialRegistrationInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEnterpriseKeyCredentialRegistrationInfo) Get_TenantId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TenantId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEnterpriseKeyCredentialRegistrationInfo) Get_TenantName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TenantName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEnterpriseKeyCredentialRegistrationInfo) Get_Subject() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Subject, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEnterpriseKeyCredentialRegistrationInfo) Get_KeyId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KeyId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEnterpriseKeyCredentialRegistrationInfo) Get_KeyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KeyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 83F3BE3F-A25F-4CBA-BB8E-BDC32D03C297
var IID_IEnterpriseKeyCredentialRegistrationManager = syscall.GUID{0x83F3BE3F, 0xA25F, 0x4CBA,
	[8]byte{0xBB, 0x8E, 0xBD, 0xC3, 0x2D, 0x03, 0xC2, 0x97}}

type IEnterpriseKeyCredentialRegistrationManagerInterface interface {
	win32.IInspectableInterface
	GetRegistrationsAsync() *IAsyncOperation[*IVectorView[*IEnterpriseKeyCredentialRegistrationInfo]]
}

type IEnterpriseKeyCredentialRegistrationManagerVtbl struct {
	win32.IInspectableVtbl
	GetRegistrationsAsync uintptr
}

type IEnterpriseKeyCredentialRegistrationManager struct {
	win32.IInspectable
}

func (this *IEnterpriseKeyCredentialRegistrationManager) Vtbl() *IEnterpriseKeyCredentialRegistrationManagerVtbl {
	return (*IEnterpriseKeyCredentialRegistrationManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEnterpriseKeyCredentialRegistrationManager) GetRegistrationsAsync() *IAsyncOperation[*IVectorView[*IEnterpriseKeyCredentialRegistrationInfo]] {
	var _result *IAsyncOperation[*IVectorView[*IEnterpriseKeyCredentialRegistrationInfo]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetRegistrationsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 77B85E9E-ACF4-4BC0-BAC2-40BB46EFBB3F
var IID_IEnterpriseKeyCredentialRegistrationManagerStatics = syscall.GUID{0x77B85E9E, 0xACF4, 0x4BC0,
	[8]byte{0xBA, 0xC2, 0x40, 0xBB, 0x46, 0xEF, 0xBB, 0x3F}}

type IEnterpriseKeyCredentialRegistrationManagerStaticsInterface interface {
	win32.IInspectableInterface
	Get_Current() *IEnterpriseKeyCredentialRegistrationManager
}

type IEnterpriseKeyCredentialRegistrationManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Current uintptr
}

type IEnterpriseKeyCredentialRegistrationManagerStatics struct {
	win32.IInspectable
}

func (this *IEnterpriseKeyCredentialRegistrationManagerStatics) Vtbl() *IEnterpriseKeyCredentialRegistrationManagerStaticsVtbl {
	return (*IEnterpriseKeyCredentialRegistrationManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEnterpriseKeyCredentialRegistrationManagerStatics) Get_Current() *IEnterpriseKeyCredentialRegistrationManager {
	var _result *IEnterpriseKeyCredentialRegistrationManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Current, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type EnterpriseKeyCredentialRegistrationInfo struct {
	RtClass
	*IEnterpriseKeyCredentialRegistrationInfo
}

type EnterpriseKeyCredentialRegistrationManager struct {
	RtClass
	*IEnterpriseKeyCredentialRegistrationManager
}

func NewIEnterpriseKeyCredentialRegistrationManagerStatics() *IEnterpriseKeyCredentialRegistrationManagerStatics {
	var p *IEnterpriseKeyCredentialRegistrationManagerStatics
	hs := NewHStr("Windows.Security.Authentication.Identity.EnterpriseKeyCredentialRegistrationManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IEnterpriseKeyCredentialRegistrationManagerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
