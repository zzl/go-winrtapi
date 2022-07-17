package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type MessagingSyncPolicy int32

const (
	MessagingSyncPolicy_Disallowed MessagingSyncPolicy = 0
	MessagingSyncPolicy_Allowed    MessagingSyncPolicy = 1
	MessagingSyncPolicy_Required   MessagingSyncPolicy = 2
)

// structs

type WorkplaceSettingsContract struct {
}

// interfaces

// C39709E7-741C-41F2-A4B6-314C31502586
var IID_IMdmAllowPolicyStatics = syscall.GUID{0xC39709E7, 0x741C, 0x41F2,
	[8]byte{0xA4, 0xB6, 0x31, 0x4C, 0x31, 0x50, 0x25, 0x86}}

type IMdmAllowPolicyStaticsInterface interface {
	win32.IInspectableInterface
	IsBrowserAllowed() bool
	IsCameraAllowed() bool
	IsMicrosoftAccountAllowed() bool
	IsStoreAllowed() bool
}

type IMdmAllowPolicyStaticsVtbl struct {
	win32.IInspectableVtbl
	IsBrowserAllowed          uintptr
	IsCameraAllowed           uintptr
	IsMicrosoftAccountAllowed uintptr
	IsStoreAllowed            uintptr
}

type IMdmAllowPolicyStatics struct {
	win32.IInspectable
}

func (this *IMdmAllowPolicyStatics) Vtbl() *IMdmAllowPolicyStaticsVtbl {
	return (*IMdmAllowPolicyStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMdmAllowPolicyStatics) IsBrowserAllowed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsBrowserAllowed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMdmAllowPolicyStatics) IsCameraAllowed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsCameraAllowed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMdmAllowPolicyStatics) IsMicrosoftAccountAllowed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsMicrosoftAccountAllowed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMdmAllowPolicyStatics) IsStoreAllowed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsStoreAllowed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C99C7526-03D4-49F9-A993-43EFCCD265C4
var IID_IMdmPolicyStatics2 = syscall.GUID{0xC99C7526, 0x03D4, 0x49F9,
	[8]byte{0xA9, 0x93, 0x43, 0xEF, 0xCC, 0xD2, 0x65, 0xC4}}

type IMdmPolicyStatics2Interface interface {
	win32.IInspectableInterface
	GetMessagingSyncPolicy() MessagingSyncPolicy
}

type IMdmPolicyStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetMessagingSyncPolicy uintptr
}

type IMdmPolicyStatics2 struct {
	win32.IInspectable
}

func (this *IMdmPolicyStatics2) Vtbl() *IMdmPolicyStatics2Vtbl {
	return (*IMdmPolicyStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMdmPolicyStatics2) GetMessagingSyncPolicy() MessagingSyncPolicy {
	var _result MessagingSyncPolicy
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetMessagingSyncPolicy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// E4676FFD-2D92-4C08-BAD4-F6590B54A6D3
var IID_IWorkplaceSettingsStatics = syscall.GUID{0xE4676FFD, 0x2D92, 0x4C08,
	[8]byte{0xBA, 0xD4, 0xF6, 0x59, 0x0B, 0x54, 0xA6, 0xD3}}

type IWorkplaceSettingsStaticsInterface interface {
	win32.IInspectableInterface
	Get_IsMicrosoftAccountOptional() bool
}

type IWorkplaceSettingsStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_IsMicrosoftAccountOptional uintptr
}

type IWorkplaceSettingsStatics struct {
	win32.IInspectable
}

func (this *IWorkplaceSettingsStatics) Vtbl() *IWorkplaceSettingsStaticsVtbl {
	return (*IWorkplaceSettingsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWorkplaceSettingsStatics) Get_IsMicrosoftAccountOptional() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsMicrosoftAccountOptional, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type MdmPolicy struct {
	RtClass
}

func NewIMdmAllowPolicyStatics() *IMdmAllowPolicyStatics {
	var p *IMdmAllowPolicyStatics
	hs := NewHStr("Windows.Management.Workplace.MdmPolicy")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMdmAllowPolicyStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIMdmPolicyStatics2() *IMdmPolicyStatics2 {
	var p *IMdmPolicyStatics2
	hs := NewHStr("Windows.Management.Workplace.MdmPolicy")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMdmPolicyStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
