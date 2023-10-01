package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"log"
	"syscall"
	"unsafe"
)

// enums

// enum
type JsonErrorStatus int32

const (
	JsonErrorStatus_Unknown             JsonErrorStatus = 0
	JsonErrorStatus_InvalidJsonString   JsonErrorStatus = 1
	JsonErrorStatus_InvalidJsonNumber   JsonErrorStatus = 2
	JsonErrorStatus_JsonValueNotFound   JsonErrorStatus = 3
	JsonErrorStatus_ImplementationLimit JsonErrorStatus = 4
)

// enum
type JsonValueType int32

const (
	JsonValueType_Null    JsonValueType = 0
	JsonValueType_Boolean JsonValueType = 1
	JsonValueType_Number  JsonValueType = 2
	JsonValueType_String  JsonValueType = 3
	JsonValueType_Array   JsonValueType = 4
	JsonValueType_Object  JsonValueType = 5
)

// interfaces

// 08C1DDB6-0CBD-4A9A-B5D3-2F852DC37E81
var IID_IJsonArray = syscall.GUID{0x08C1DDB6, 0x0CBD, 0x4A9A,
	[8]byte{0xB5, 0xD3, 0x2F, 0x85, 0x2D, 0xC3, 0x7E, 0x81}}

type IJsonArrayInterface interface {
	win32.IInspectableInterface
	GetObjectAt(index uint32) *IJsonObject
	GetArrayAt(index uint32) *IJsonArray
	GetStringAt(index uint32) string
	GetNumberAt(index uint32) float64
	GetBooleanAt(index uint32) bool
}

type IJsonArrayVtbl struct {
	win32.IInspectableVtbl
	GetObjectAt  uintptr
	GetArrayAt   uintptr
	GetStringAt  uintptr
	GetNumberAt  uintptr
	GetBooleanAt uintptr
}

type IJsonArray struct {
	win32.IInspectable
}

func (this *IJsonArray) Vtbl() *IJsonArrayVtbl {
	return (*IJsonArrayVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IJsonArray) GetObjectAt(index uint32) *IJsonObject {
	var _result *IJsonObject
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetObjectAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IJsonArray) GetArrayAt(index uint32) *IJsonArray {
	var _result *IJsonArray
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetArrayAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IJsonArray) GetStringAt(index uint32) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStringAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IJsonArray) GetNumberAt(index uint32) float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNumberAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IJsonArray) GetBooleanAt(index uint32) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetBooleanAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// DB1434A9-E164-499F-93E2-8A8F49BB90BA
var IID_IJsonArrayStatics = syscall.GUID{0xDB1434A9, 0xE164, 0x499F,
	[8]byte{0x93, 0xE2, 0x8A, 0x8F, 0x49, 0xBB, 0x90, 0xBA}}

type IJsonArrayStaticsInterface interface {
	win32.IInspectableInterface
	Parse(input string) *IJsonArray
	TryParse(input string, result *IJsonArray) bool
}

type IJsonArrayStaticsVtbl struct {
	win32.IInspectableVtbl
	Parse    uintptr
	TryParse uintptr
}

type IJsonArrayStatics struct {
	win32.IInspectable
}

func (this *IJsonArrayStatics) Vtbl() *IJsonArrayStaticsVtbl {
	return (*IJsonArrayStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IJsonArrayStatics) Parse(input string) *IJsonArray {
	var _result *IJsonArray
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Parse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IJsonArrayStatics) TryParse(input string, result *IJsonArray) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(result)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 404030DA-87D0-436C-83AB-FC7B12C0CC26
var IID_IJsonErrorStatics2 = syscall.GUID{0x404030DA, 0x87D0, 0x436C,
	[8]byte{0x83, 0xAB, 0xFC, 0x7B, 0x12, 0xC0, 0xCC, 0x26}}

type IJsonErrorStatics2Interface interface {
	win32.IInspectableInterface
	GetJsonStatus(hresult int32) JsonErrorStatus
}

type IJsonErrorStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetJsonStatus uintptr
}

type IJsonErrorStatics2 struct {
	win32.IInspectable
}

func (this *IJsonErrorStatics2) Vtbl() *IJsonErrorStatics2Vtbl {
	return (*IJsonErrorStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IJsonErrorStatics2) GetJsonStatus(hresult int32) JsonErrorStatus {
	var _result JsonErrorStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetJsonStatus, uintptr(unsafe.Pointer(this)), uintptr(hresult), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 064E24DD-29C2-4F83-9AC1-9EE11578BEB3
var IID_IJsonObject = syscall.GUID{0x064E24DD, 0x29C2, 0x4F83,
	[8]byte{0x9A, 0xC1, 0x9E, 0xE1, 0x15, 0x78, 0xBE, 0xB3}}

type IJsonObjectInterface interface {
	win32.IInspectableInterface
	GetNamedValue(name string) *IJsonValue
	SetNamedValue(name string, value *IJsonValue)
	GetNamedObject(name string) *IJsonObject
	GetNamedArray(name string) *IJsonArray
	GetNamedString(name string) string
	GetNamedNumber(name string) float64
	GetNamedBoolean(name string) bool
}

type IJsonObjectVtbl struct {
	win32.IInspectableVtbl
	GetNamedValue   uintptr
	SetNamedValue   uintptr
	GetNamedObject  uintptr
	GetNamedArray   uintptr
	GetNamedString  uintptr
	GetNamedNumber  uintptr
	GetNamedBoolean uintptr
}

type IJsonObject struct {
	win32.IInspectable
}

func (this *IJsonObject) Vtbl() *IJsonObjectVtbl {
	return (*IJsonObjectVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IJsonObject) GetNamedValue(name string) *IJsonValue {
	var _result *IJsonValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNamedValue, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IJsonObject) SetNamedValue(name string, value *IJsonValue) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetNamedValue, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IJsonObject) GetNamedObject(name string) *IJsonObject {
	var _result *IJsonObject
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNamedObject, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IJsonObject) GetNamedArray(name string) *IJsonArray {
	var _result *IJsonArray
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNamedArray, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IJsonObject) GetNamedString(name string) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNamedString, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IJsonObject) GetNamedNumber(name string) float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNamedNumber, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IJsonObject) GetNamedBoolean(name string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNamedBoolean, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 2289F159-54DE-45D8-ABCC-22603FA066A0
var IID_IJsonObjectStatics = syscall.GUID{0x2289F159, 0x54DE, 0x45D8,
	[8]byte{0xAB, 0xCC, 0x22, 0x60, 0x3F, 0xA0, 0x66, 0xA0}}

type IJsonObjectStaticsInterface interface {
	win32.IInspectableInterface
	Parse(input string) *IJsonObject
	TryParse(input string, result *IJsonObject) bool
}

type IJsonObjectStaticsVtbl struct {
	win32.IInspectableVtbl
	Parse    uintptr
	TryParse uintptr
}

type IJsonObjectStatics struct {
	win32.IInspectable
}

func (this *IJsonObjectStatics) Vtbl() *IJsonObjectStaticsVtbl {
	return (*IJsonObjectStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IJsonObjectStatics) Parse(input string) *IJsonObject {
	var _result *IJsonObject
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Parse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IJsonObjectStatics) TryParse(input string, result *IJsonObject) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(result)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D960D2A2-B7F0-4F00-8E44-D82CF415EA13
var IID_IJsonObjectWithDefaultValues = syscall.GUID{0xD960D2A2, 0xB7F0, 0x4F00,
	[8]byte{0x8E, 0x44, 0xD8, 0x2C, 0xF4, 0x15, 0xEA, 0x13}}

type IJsonObjectWithDefaultValuesInterface interface {
	win32.IInspectableInterface
	GetNamedValueOrDefault(name string, defaultValue *IJsonValue) *IJsonValue
	GetNamedObjectOrDefault(name string, defaultValue *IJsonObject) *IJsonObject
	GetNamedStringOrDefault(name string, defaultValue string) string
	GetNamedArrayOrDefault(name string, defaultValue *IJsonArray) *IJsonArray
	GetNamedNumberOrDefault(name string, defaultValue float64) float64
	GetNamedBooleanOrDefault(name string, defaultValue bool) bool
}

type IJsonObjectWithDefaultValuesVtbl struct {
	win32.IInspectableVtbl
	GetNamedValueOrDefault   uintptr
	GetNamedObjectOrDefault  uintptr
	GetNamedStringOrDefault  uintptr
	GetNamedArrayOrDefault   uintptr
	GetNamedNumberOrDefault  uintptr
	GetNamedBooleanOrDefault uintptr
}

type IJsonObjectWithDefaultValues struct {
	win32.IInspectable
}

func (this *IJsonObjectWithDefaultValues) Vtbl() *IJsonObjectWithDefaultValuesVtbl {
	return (*IJsonObjectWithDefaultValuesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IJsonObjectWithDefaultValues) GetNamedValueOrDefault(name string, defaultValue *IJsonValue) *IJsonValue {
	var _result *IJsonValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNamedValueOrDefault, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(defaultValue)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IJsonObjectWithDefaultValues) GetNamedObjectOrDefault(name string, defaultValue *IJsonObject) *IJsonObject {
	var _result *IJsonObject
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNamedObjectOrDefault, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(defaultValue)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IJsonObjectWithDefaultValues) GetNamedStringOrDefault(name string, defaultValue string) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNamedStringOrDefault, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, NewHStr(defaultValue).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IJsonObjectWithDefaultValues) GetNamedArrayOrDefault(name string, defaultValue *IJsonArray) *IJsonArray {
	var _result *IJsonArray
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNamedArrayOrDefault, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(defaultValue)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IJsonObjectWithDefaultValues) GetNamedNumberOrDefault(name string, defaultValue float64) float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNamedNumberOrDefault, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(defaultValue), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IJsonObjectWithDefaultValues) GetNamedBooleanOrDefault(name string, defaultValue bool) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNamedBooleanOrDefault, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(*(*byte)(unsafe.Pointer(&defaultValue))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A3219ECB-F0B3-4DCD-BEEE-19D48CD3ED1E
var IID_IJsonValue = syscall.GUID{0xA3219ECB, 0xF0B3, 0x4DCD,
	[8]byte{0xBE, 0xEE, 0x19, 0xD4, 0x8C, 0xD3, 0xED, 0x1E}}

type IJsonValueInterface interface {
	win32.IInspectableInterface
	Get_ValueType() JsonValueType
	Stringify() string
	GetString() string
	GetNumber() float64
	GetBoolean() bool
	GetArray() *IJsonArray
	GetObject() *IJsonObject
}

type IJsonValueVtbl struct {
	win32.IInspectableVtbl
	Get_ValueType uintptr
	Stringify     uintptr
	GetString     uintptr
	GetNumber     uintptr
	GetBoolean    uintptr
	GetArray      uintptr
	GetObject     uintptr
}

type IJsonValue struct {
	win32.IInspectable
}

func (this *IJsonValue) Vtbl() *IJsonValueVtbl {
	return (*IJsonValueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IJsonValue) Get_ValueType() JsonValueType {
	var _result JsonValueType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ValueType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IJsonValue) Stringify() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stringify, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IJsonValue) GetString() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IJsonValue) GetNumber() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IJsonValue) GetBoolean() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetBoolean, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IJsonValue) GetArray() *IJsonArray {
	var _result *IJsonArray
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IJsonValue) GetObject() *IJsonObject {
	var _result *IJsonObject
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetObject, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5F6B544A-2F53-48E1-91A3-F78B50A6345C
var IID_IJsonValueStatics = syscall.GUID{0x5F6B544A, 0x2F53, 0x48E1,
	[8]byte{0x91, 0xA3, 0xF7, 0x8B, 0x50, 0xA6, 0x34, 0x5C}}

type IJsonValueStaticsInterface interface {
	win32.IInspectableInterface
	Parse(input string) *IJsonValue
	TryParse(input string, result *IJsonValue) bool
	CreateBooleanValue(input bool) *IJsonValue
	CreateNumberValue(input float64) *IJsonValue
	CreateStringValue(input string) *IJsonValue
}

type IJsonValueStaticsVtbl struct {
	win32.IInspectableVtbl
	Parse              uintptr
	TryParse           uintptr
	CreateBooleanValue uintptr
	CreateNumberValue  uintptr
	CreateStringValue  uintptr
}

type IJsonValueStatics struct {
	win32.IInspectable
}

func (this *IJsonValueStatics) Vtbl() *IJsonValueStaticsVtbl {
	return (*IJsonValueStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IJsonValueStatics) Parse(input string) *IJsonValue {
	var _result *IJsonValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Parse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IJsonValueStatics) TryParse(input string, result *IJsonValue) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(result)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IJsonValueStatics) CreateBooleanValue(input bool) *IJsonValue {
	var _result *IJsonValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateBooleanValue, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&input))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IJsonValueStatics) CreateNumberValue(input float64) *IJsonValue {
	var _result *IJsonValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateNumberValue, uintptr(unsafe.Pointer(this)), uintptr(input), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IJsonValueStatics) CreateStringValue(input string) *IJsonValue {
	var _result *IJsonValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateStringValue, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1D9ECBE4-3FE8-4335-8392-93D8E36865F0
var IID_IJsonValueStatics2 = syscall.GUID{0x1D9ECBE4, 0x3FE8, 0x4335,
	[8]byte{0x83, 0x92, 0x93, 0xD8, 0xE3, 0x68, 0x65, 0xF0}}

type IJsonValueStatics2Interface interface {
	win32.IInspectableInterface
	CreateNullValue() *IJsonValue
}

type IJsonValueStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateNullValue uintptr
}

type IJsonValueStatics2 struct {
	win32.IInspectable
}

func (this *IJsonValueStatics2) Vtbl() *IJsonValueStatics2Vtbl {
	return (*IJsonValueStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IJsonValueStatics2) CreateNullValue() *IJsonValue {
	var _result *IJsonValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateNullValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type JsonArray struct {
	RtClass
	*IJsonArray
}

func NewJsonArray() *JsonArray {
	hs := NewHStr("Windows.Data.Json.JsonArray")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &JsonArray{
		RtClass:    RtClass{PInspect: p},
		IJsonArray: (*IJsonArray)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewIJsonArrayStatics() *IJsonArrayStatics {
	var p *IJsonArrayStatics
	hs := NewHStr("Windows.Data.Json.JsonArray")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IJsonArrayStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type JsonError struct {
	RtClass
}

func NewIJsonErrorStatics2() *IJsonErrorStatics2 {
	var p *IJsonErrorStatics2
	hs := NewHStr("Windows.Data.Json.JsonError")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IJsonErrorStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type JsonObject struct {
	RtClass
	*IJsonObject
}

func NewJsonObject() *JsonObject {
	hs := NewHStr("Windows.Data.Json.JsonObject")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &JsonObject{
		RtClass:     RtClass{PInspect: p},
		IJsonObject: (*IJsonObject)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewIJsonObjectStatics() *IJsonObjectStatics {
	var p *IJsonObjectStatics
	hs := NewHStr("Windows.Data.Json.JsonObject")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IJsonObjectStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type JsonValue struct {
	RtClass
	*IJsonValue
}

func NewIJsonValueStatics2() *IJsonValueStatics2 {
	var p *IJsonValueStatics2
	hs := NewHStr("Windows.Data.Json.JsonValue")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IJsonValueStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIJsonValueStatics() *IJsonValueStatics {
	var p *IJsonValueStatics
	hs := NewHStr("Windows.Data.Json.JsonValue")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IJsonValueStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
