package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// interfaces

// 87C24804-A22E-4ADB-BA26-D78EF639BCF4
var IID_IPerceptionTimestamp = syscall.GUID{0x87C24804, 0xA22E, 0x4ADB,
	[8]byte{0xBA, 0x26, 0xD7, 0x8E, 0xF6, 0x39, 0xBC, 0xF4}}

type IPerceptionTimestampInterface interface {
	win32.IInspectableInterface
	Get_TargetTime() DateTime
	Get_PredictionAmount() TimeSpan
}

type IPerceptionTimestampVtbl struct {
	win32.IInspectableVtbl
	Get_TargetTime       uintptr
	Get_PredictionAmount uintptr
}

type IPerceptionTimestamp struct {
	win32.IInspectable
}

func (this *IPerceptionTimestamp) Vtbl() *IPerceptionTimestampVtbl {
	return (*IPerceptionTimestampVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionTimestamp) Get_TargetTime() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TargetTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionTimestamp) Get_PredictionAmount() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PredictionAmount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// E354B7ED-2BD1-41B7-9ED0-74A15C354537
var IID_IPerceptionTimestamp2 = syscall.GUID{0xE354B7ED, 0x2BD1, 0x41B7,
	[8]byte{0x9E, 0xD0, 0x74, 0xA1, 0x5C, 0x35, 0x45, 0x37}}

type IPerceptionTimestamp2Interface interface {
	win32.IInspectableInterface
	Get_SystemRelativeTargetTime() TimeSpan
}

type IPerceptionTimestamp2Vtbl struct {
	win32.IInspectableVtbl
	Get_SystemRelativeTargetTime uintptr
}

type IPerceptionTimestamp2 struct {
	win32.IInspectable
}

func (this *IPerceptionTimestamp2) Vtbl() *IPerceptionTimestamp2Vtbl {
	return (*IPerceptionTimestamp2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionTimestamp2) Get_SystemRelativeTargetTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SystemRelativeTargetTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 47A611D4-A9DF-4EDC-855D-F4D339D967AC
var IID_IPerceptionTimestampHelperStatics = syscall.GUID{0x47A611D4, 0xA9DF, 0x4EDC,
	[8]byte{0x85, 0x5D, 0xF4, 0xD3, 0x39, 0xD9, 0x67, 0xAC}}

type IPerceptionTimestampHelperStaticsInterface interface {
	win32.IInspectableInterface
	FromHistoricalTargetTime(targetTime DateTime) *IPerceptionTimestamp
}

type IPerceptionTimestampHelperStaticsVtbl struct {
	win32.IInspectableVtbl
	FromHistoricalTargetTime uintptr
}

type IPerceptionTimestampHelperStatics struct {
	win32.IInspectable
}

func (this *IPerceptionTimestampHelperStatics) Vtbl() *IPerceptionTimestampHelperStaticsVtbl {
	return (*IPerceptionTimestampHelperStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionTimestampHelperStatics) FromHistoricalTargetTime(targetTime DateTime) *IPerceptionTimestamp {
	var _result *IPerceptionTimestamp
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromHistoricalTargetTime, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&targetTime)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 73D1A7FE-3FB9-4571-87D4-3C920A5E86EB
var IID_IPerceptionTimestampHelperStatics2 = syscall.GUID{0x73D1A7FE, 0x3FB9, 0x4571,
	[8]byte{0x87, 0xD4, 0x3C, 0x92, 0x0A, 0x5E, 0x86, 0xEB}}

type IPerceptionTimestampHelperStatics2Interface interface {
	win32.IInspectableInterface
	FromSystemRelativeTargetTime(targetTime TimeSpan) *IPerceptionTimestamp
}

type IPerceptionTimestampHelperStatics2Vtbl struct {
	win32.IInspectableVtbl
	FromSystemRelativeTargetTime uintptr
}

type IPerceptionTimestampHelperStatics2 struct {
	win32.IInspectable
}

func (this *IPerceptionTimestampHelperStatics2) Vtbl() *IPerceptionTimestampHelperStatics2Vtbl {
	return (*IPerceptionTimestampHelperStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionTimestampHelperStatics2) FromSystemRelativeTargetTime(targetTime TimeSpan) *IPerceptionTimestamp {
	var _result *IPerceptionTimestamp
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromSystemRelativeTargetTime, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&targetTime)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type PerceptionTimestampHelper struct {
	RtClass
}

func NewIPerceptionTimestampHelperStatics() *IPerceptionTimestampHelperStatics {
	var p *IPerceptionTimestampHelperStatics
	hs := NewHStr("Windows.Perception.PerceptionTimestampHelper")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPerceptionTimestampHelperStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIPerceptionTimestampHelperStatics2() *IPerceptionTimestampHelperStatics2 {
	var p *IPerceptionTimestampHelperStatics2
	hs := NewHStr("Windows.Perception.PerceptionTimestampHelper")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPerceptionTimestampHelperStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
