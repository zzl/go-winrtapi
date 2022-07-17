package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
// flags
type AttributeTargets uint32

const (
	AttributeTargets_All           AttributeTargets = 4294967295
	AttributeTargets_Delegate      AttributeTargets = 1
	AttributeTargets_Enum          AttributeTargets = 2
	AttributeTargets_Event         AttributeTargets = 4
	AttributeTargets_Field         AttributeTargets = 8
	AttributeTargets_Interface     AttributeTargets = 16
	AttributeTargets_Method        AttributeTargets = 64
	AttributeTargets_Parameter     AttributeTargets = 128
	AttributeTargets_Property      AttributeTargets = 256
	AttributeTargets_RuntimeClass  AttributeTargets = 512
	AttributeTargets_Struct        AttributeTargets = 1024
	AttributeTargets_InterfaceImpl AttributeTargets = 2048
	AttributeTargets_ApiContract   AttributeTargets = 8192
)

// enum
type CompositionType int32

const (
	CompositionType_Protected CompositionType = 1
	CompositionType_Public    CompositionType = 2
)

// enum
type DeprecationType int32

const (
	DeprecationType_Deprecate DeprecationType = 0
	DeprecationType_Remove    DeprecationType = 1
)

// enum
type FeatureStage int32

const (
	FeatureStage_AlwaysDisabled    FeatureStage = 0
	FeatureStage_DisabledByDefault FeatureStage = 1
	FeatureStage_EnabledByDefault  FeatureStage = 2
	FeatureStage_AlwaysEnabled     FeatureStage = 3
)

// enum
type GCPressureAmount int32

const (
	GCPressureAmount_Low    GCPressureAmount = 0
	GCPressureAmount_Medium GCPressureAmount = 1
	GCPressureAmount_High   GCPressureAmount = 2
)

// enum
type MarshalingType int32

const (
	MarshalingType_None              MarshalingType = 1
	MarshalingType_Agile             MarshalingType = 2
	MarshalingType_Standard          MarshalingType = 3
	MarshalingType_InvalidMarshaling MarshalingType = 0
)

// enum
type Platform int32

const (
	Platform_Windows      Platform = 0
	Platform_WindowsPhone Platform = 1
)

// enum
type ThreadingModel int32

const (
	ThreadingModel_STA              ThreadingModel = 1
	ThreadingModel_MTA              ThreadingModel = 2
	ThreadingModel_Both             ThreadingModel = 3
	ThreadingModel_InvalidThreading ThreadingModel = 0
)

// interfaces

// 997439FE-F681-4A11-B416-C13A47E8BA36
var IID_IApiInformationStatics = syscall.GUID{0x997439FE, 0xF681, 0x4A11,
	[8]byte{0xB4, 0x16, 0xC1, 0x3A, 0x47, 0xE8, 0xBA, 0x36}}

type IApiInformationStaticsInterface interface {
	win32.IInspectableInterface
	IsTypePresent(typeName string) bool
	IsMethodPresent(typeName string, methodName string) bool
	IsMethodPresentWithArity(typeName string, methodName string, inputParameterCount uint32) bool
	IsEventPresent(typeName string, eventName string) bool
	IsPropertyPresent(typeName string, propertyName string) bool
	IsReadOnlyPropertyPresent(typeName string, propertyName string) bool
	IsWriteablePropertyPresent(typeName string, propertyName string) bool
	IsEnumNamedValuePresent(enumTypeName string, valueName string) bool
	IsApiContractPresentByMajor(contractName string, majorVersion uint16) bool
	IsApiContractPresentByMajorAndMinor(contractName string, majorVersion uint16, minorVersion uint16) bool
}

type IApiInformationStaticsVtbl struct {
	win32.IInspectableVtbl
	IsTypePresent                       uintptr
	IsMethodPresent                     uintptr
	IsMethodPresentWithArity            uintptr
	IsEventPresent                      uintptr
	IsPropertyPresent                   uintptr
	IsReadOnlyPropertyPresent           uintptr
	IsWriteablePropertyPresent          uintptr
	IsEnumNamedValuePresent             uintptr
	IsApiContractPresentByMajor         uintptr
	IsApiContractPresentByMajorAndMinor uintptr
}

type IApiInformationStatics struct {
	win32.IInspectable
}

func (this *IApiInformationStatics) Vtbl() *IApiInformationStaticsVtbl {
	return (*IApiInformationStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApiInformationStatics) IsTypePresent(typeName string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsTypePresent, uintptr(unsafe.Pointer(this)), NewHStr(typeName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApiInformationStatics) IsMethodPresent(typeName string, methodName string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsMethodPresent, uintptr(unsafe.Pointer(this)), NewHStr(typeName).Ptr, NewHStr(methodName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApiInformationStatics) IsMethodPresentWithArity(typeName string, methodName string, inputParameterCount uint32) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsMethodPresentWithArity, uintptr(unsafe.Pointer(this)), NewHStr(typeName).Ptr, NewHStr(methodName).Ptr, uintptr(inputParameterCount), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApiInformationStatics) IsEventPresent(typeName string, eventName string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsEventPresent, uintptr(unsafe.Pointer(this)), NewHStr(typeName).Ptr, NewHStr(eventName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApiInformationStatics) IsPropertyPresent(typeName string, propertyName string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsPropertyPresent, uintptr(unsafe.Pointer(this)), NewHStr(typeName).Ptr, NewHStr(propertyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApiInformationStatics) IsReadOnlyPropertyPresent(typeName string, propertyName string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsReadOnlyPropertyPresent, uintptr(unsafe.Pointer(this)), NewHStr(typeName).Ptr, NewHStr(propertyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApiInformationStatics) IsWriteablePropertyPresent(typeName string, propertyName string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsWriteablePropertyPresent, uintptr(unsafe.Pointer(this)), NewHStr(typeName).Ptr, NewHStr(propertyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApiInformationStatics) IsEnumNamedValuePresent(enumTypeName string, valueName string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsEnumNamedValuePresent, uintptr(unsafe.Pointer(this)), NewHStr(enumTypeName).Ptr, NewHStr(valueName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApiInformationStatics) IsApiContractPresentByMajor(contractName string, majorVersion uint16) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsApiContractPresentByMajor, uintptr(unsafe.Pointer(this)), NewHStr(contractName).Ptr, uintptr(majorVersion), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApiInformationStatics) IsApiContractPresentByMajorAndMinor(contractName string, majorVersion uint16, minorVersion uint16) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsApiContractPresentByMajorAndMinor, uintptr(unsafe.Pointer(this)), NewHStr(contractName).Ptr, uintptr(majorVersion), uintptr(minorVersion), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type ApiInformation struct {
	RtClass
}

func NewIApiInformationStatics() *IApiInformationStatics {
	var p *IApiInformationStatics
	hs := NewHStr("Windows.Foundation.Metadata.ApiInformation")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IApiInformationStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
