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
type DomainNameType int32

const (
	DomainNameType_Suffix         DomainNameType = 0
	DomainNameType_FullyQualified DomainNameType = 1
)

// enum
// flags
type HostNameSortOptions uint32

const (
	HostNameSortOptions_None                       HostNameSortOptions = 0
	HostNameSortOptions_OptimizeForLongConnections HostNameSortOptions = 2
)

// enum
type HostNameType int32

const (
	HostNameType_DomainName HostNameType = 0
	HostNameType_Ipv4       HostNameType = 1
	HostNameType_Ipv6       HostNameType = 2
	HostNameType_Bluetooth  HostNameType = 3
)

// interfaces

// 33A0AA36-F8FA-4B30-B856-76517C3BD06D
var IID_IEndpointPair = syscall.GUID{0x33A0AA36, 0xF8FA, 0x4B30,
	[8]byte{0xB8, 0x56, 0x76, 0x51, 0x7C, 0x3B, 0xD0, 0x6D}}

type IEndpointPairInterface interface {
	win32.IInspectableInterface
	Get_LocalHostName() *IHostName
	Put_LocalHostName(value *IHostName)
	Get_LocalServiceName() string
	Put_LocalServiceName(value string)
	Get_RemoteHostName() *IHostName
	Put_RemoteHostName(value *IHostName)
	Get_RemoteServiceName() string
	Put_RemoteServiceName(value string)
}

type IEndpointPairVtbl struct {
	win32.IInspectableVtbl
	Get_LocalHostName     uintptr
	Put_LocalHostName     uintptr
	Get_LocalServiceName  uintptr
	Put_LocalServiceName  uintptr
	Get_RemoteHostName    uintptr
	Put_RemoteHostName    uintptr
	Get_RemoteServiceName uintptr
	Put_RemoteServiceName uintptr
}

type IEndpointPair struct {
	win32.IInspectable
}

func (this *IEndpointPair) Vtbl() *IEndpointPairVtbl {
	return (*IEndpointPairVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEndpointPair) Get_LocalHostName() *IHostName {
	var _result *IHostName
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalHostName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IEndpointPair) Put_LocalHostName(value *IHostName) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LocalHostName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IEndpointPair) Get_LocalServiceName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalServiceName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEndpointPair) Put_LocalServiceName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LocalServiceName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IEndpointPair) Get_RemoteHostName() *IHostName {
	var _result *IHostName
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemoteHostName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IEndpointPair) Put_RemoteHostName(value *IHostName) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RemoteHostName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IEndpointPair) Get_RemoteServiceName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemoteServiceName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEndpointPair) Put_RemoteServiceName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RemoteServiceName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// B609D971-64E0-442B-AA6F-CC8C8F181F78
var IID_IEndpointPairFactory = syscall.GUID{0xB609D971, 0x64E0, 0x442B,
	[8]byte{0xAA, 0x6F, 0xCC, 0x8C, 0x8F, 0x18, 0x1F, 0x78}}

type IEndpointPairFactoryInterface interface {
	win32.IInspectableInterface
	CreateEndpointPair(localHostName *IHostName, localServiceName string, remoteHostName *IHostName, remoteServiceName string) *IEndpointPair
}

type IEndpointPairFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateEndpointPair uintptr
}

type IEndpointPairFactory struct {
	win32.IInspectable
}

func (this *IEndpointPairFactory) Vtbl() *IEndpointPairFactoryVtbl {
	return (*IEndpointPairFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEndpointPairFactory) CreateEndpointPair(localHostName *IHostName, localServiceName string, remoteHostName *IHostName, remoteServiceName string) *IEndpointPair {
	var _result *IEndpointPair
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateEndpointPair, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(localHostName)), NewHStr(localServiceName).Ptr, uintptr(unsafe.Pointer(remoteHostName)), NewHStr(remoteServiceName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BF8ECAAD-ED96-49A7-9084-D416CAE88DCB
var IID_IHostName = syscall.GUID{0xBF8ECAAD, 0xED96, 0x49A7,
	[8]byte{0x90, 0x84, 0xD4, 0x16, 0xCA, 0xE8, 0x8D, 0xCB}}

type IHostNameInterface interface {
	win32.IInspectableInterface
	Get_IPInformation() *IIPInformation
	Get_RawName() string
	Get_DisplayName() string
	Get_CanonicalName() string
	Get_Type() HostNameType
	IsEqual(hostName *IHostName) bool
}

type IHostNameVtbl struct {
	win32.IInspectableVtbl
	Get_IPInformation uintptr
	Get_RawName       uintptr
	Get_DisplayName   uintptr
	Get_CanonicalName uintptr
	Get_Type          uintptr
	IsEqual           uintptr
}

type IHostName struct {
	win32.IInspectable
}

func (this *IHostName) Vtbl() *IHostNameVtbl {
	return (*IHostNameVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHostName) Get_IPInformation() *IIPInformation {
	var _result *IIPInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IPInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHostName) Get_RawName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RawName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHostName) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHostName) Get_CanonicalName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanonicalName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHostName) Get_Type() HostNameType {
	var _result HostNameType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Type, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHostName) IsEqual(hostName *IHostName) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsEqual, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(hostName)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 458C23ED-712F-4576-ADF1-C20B2C643558
var IID_IHostNameFactory = syscall.GUID{0x458C23ED, 0x712F, 0x4576,
	[8]byte{0xAD, 0xF1, 0xC2, 0x0B, 0x2C, 0x64, 0x35, 0x58}}

type IHostNameFactoryInterface interface {
	win32.IInspectableInterface
	CreateHostName(hostName string) *IHostName
}

type IHostNameFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateHostName uintptr
}

type IHostNameFactory struct {
	win32.IInspectable
}

func (this *IHostNameFactory) Vtbl() *IHostNameFactoryVtbl {
	return (*IHostNameFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHostNameFactory) CreateHostName(hostName string) *IHostName {
	var _result *IHostName
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateHostName, uintptr(unsafe.Pointer(this)), NewHStr(hostName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F68CD4BF-A388-4E8B-91EA-54DD6DD901C0
var IID_IHostNameStatics = syscall.GUID{0xF68CD4BF, 0xA388, 0x4E8B,
	[8]byte{0x91, 0xEA, 0x54, 0xDD, 0x6D, 0xD9, 0x01, 0xC0}}

type IHostNameStaticsInterface interface {
	win32.IInspectableInterface
	Compare(value1 string, value2 string) int32
}

type IHostNameStaticsVtbl struct {
	win32.IInspectableVtbl
	Compare uintptr
}

type IHostNameStatics struct {
	win32.IInspectable
}

func (this *IHostNameStatics) Vtbl() *IHostNameStaticsVtbl {
	return (*IHostNameStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHostNameStatics) Compare(value1 string, value2 string) int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Compare, uintptr(unsafe.Pointer(this)), NewHStr(value1).Ptr, NewHStr(value2).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type EndpointPair struct {
	RtClass
	*IEndpointPair
}

func NewEndpointPair_CreateEndpointPair(localHostName *IHostName, localServiceName string, remoteHostName *IHostName, remoteServiceName string) *EndpointPair {
	hs := NewHStr("Windows.Networking.EndpointPair")
	var pFac *IEndpointPairFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IEndpointPairFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IEndpointPair
	p = pFac.CreateEndpointPair(localHostName, localServiceName, remoteHostName, remoteServiceName)
	result := &EndpointPair{
		RtClass:       RtClass{PInspect: &p.IInspectable},
		IEndpointPair: p,
	}
	com.AddToScope(result)
	return result
}

type HostName struct {
	RtClass
	*IHostName
}

func NewHostName_CreateHostName(hostName string) *HostName {
	hs := NewHStr("Windows.Networking.HostName")
	var pFac *IHostNameFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHostNameFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHostName
	p = pFac.CreateHostName(hostName)
	result := &HostName{
		RtClass:   RtClass{PInspect: &p.IInspectable},
		IHostName: p,
	}
	com.AddToScope(result)
	return result
}

func NewIHostNameStatics() *IHostNameStatics {
	var p *IHostNameStatics
	hs := NewHStr("Windows.Networking.HostName")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHostNameStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
