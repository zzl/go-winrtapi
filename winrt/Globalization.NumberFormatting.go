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
type CurrencyFormatterMode int32

const (
	CurrencyFormatterMode_UseSymbol       CurrencyFormatterMode = 0
	CurrencyFormatterMode_UseCurrencyCode CurrencyFormatterMode = 1
)

// enum
type RoundingAlgorithm int32

const (
	RoundingAlgorithm_None                  RoundingAlgorithm = 0
	RoundingAlgorithm_RoundDown             RoundingAlgorithm = 1
	RoundingAlgorithm_RoundUp               RoundingAlgorithm = 2
	RoundingAlgorithm_RoundTowardsZero      RoundingAlgorithm = 3
	RoundingAlgorithm_RoundAwayFromZero     RoundingAlgorithm = 4
	RoundingAlgorithm_RoundHalfDown         RoundingAlgorithm = 5
	RoundingAlgorithm_RoundHalfUp           RoundingAlgorithm = 6
	RoundingAlgorithm_RoundHalfTowardsZero  RoundingAlgorithm = 7
	RoundingAlgorithm_RoundHalfAwayFromZero RoundingAlgorithm = 8
	RoundingAlgorithm_RoundHalfToEven       RoundingAlgorithm = 9
	RoundingAlgorithm_RoundHalfToOdd        RoundingAlgorithm = 10
)

// interfaces

// 11730CA5-4B00-41B2-B332-73B12A497D54
var IID_ICurrencyFormatter = syscall.GUID{0x11730CA5, 0x4B00, 0x41B2,
	[8]byte{0xB3, 0x32, 0x73, 0xB1, 0x2A, 0x49, 0x7D, 0x54}}

type ICurrencyFormatterInterface interface {
	win32.IInspectableInterface
	Get_Currency() string
	Put_Currency(value string)
}

type ICurrencyFormatterVtbl struct {
	win32.IInspectableVtbl
	Get_Currency uintptr
	Put_Currency uintptr
}

type ICurrencyFormatter struct {
	win32.IInspectable
}

func (this *ICurrencyFormatter) Vtbl() *ICurrencyFormatterVtbl {
	return (*ICurrencyFormatterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICurrencyFormatter) Get_Currency() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Currency, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyFormatter) Put_Currency(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Currency, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 072C2F1D-E7BA-4197-920E-247C92F7DEA6
var IID_ICurrencyFormatter2 = syscall.GUID{0x072C2F1D, 0xE7BA, 0x4197,
	[8]byte{0x92, 0x0E, 0x24, 0x7C, 0x92, 0xF7, 0xDE, 0xA6}}

type ICurrencyFormatter2Interface interface {
	win32.IInspectableInterface
	Get_Mode() CurrencyFormatterMode
	Put_Mode(value CurrencyFormatterMode)
	ApplyRoundingForCurrency(roundingAlgorithm RoundingAlgorithm)
}

type ICurrencyFormatter2Vtbl struct {
	win32.IInspectableVtbl
	Get_Mode                 uintptr
	Put_Mode                 uintptr
	ApplyRoundingForCurrency uintptr
}

type ICurrencyFormatter2 struct {
	win32.IInspectable
}

func (this *ICurrencyFormatter2) Vtbl() *ICurrencyFormatter2Vtbl {
	return (*ICurrencyFormatter2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICurrencyFormatter2) Get_Mode() CurrencyFormatterMode {
	var _result CurrencyFormatterMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICurrencyFormatter2) Put_Mode(value CurrencyFormatterMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Mode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICurrencyFormatter2) ApplyRoundingForCurrency(roundingAlgorithm RoundingAlgorithm) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ApplyRoundingForCurrency, uintptr(unsafe.Pointer(this)), uintptr(roundingAlgorithm))
	_ = _hr
}

// 86C7537E-B938-4AA2-84B0-2C33DC5B1450
var IID_ICurrencyFormatterFactory = syscall.GUID{0x86C7537E, 0xB938, 0x4AA2,
	[8]byte{0x84, 0xB0, 0x2C, 0x33, 0xDC, 0x5B, 0x14, 0x50}}

type ICurrencyFormatterFactoryInterface interface {
	win32.IInspectableInterface
	CreateCurrencyFormatterCode(currencyCode string) *ICurrencyFormatter
	CreateCurrencyFormatterCodeContext(currencyCode string, languages *IIterable[string], geographicRegion string) *ICurrencyFormatter
}

type ICurrencyFormatterFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateCurrencyFormatterCode        uintptr
	CreateCurrencyFormatterCodeContext uintptr
}

type ICurrencyFormatterFactory struct {
	win32.IInspectable
}

func (this *ICurrencyFormatterFactory) Vtbl() *ICurrencyFormatterFactoryVtbl {
	return (*ICurrencyFormatterFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICurrencyFormatterFactory) CreateCurrencyFormatterCode(currencyCode string) *ICurrencyFormatter {
	var _result *ICurrencyFormatter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCurrencyFormatterCode, uintptr(unsafe.Pointer(this)), NewHStr(currencyCode).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICurrencyFormatterFactory) CreateCurrencyFormatterCodeContext(currencyCode string, languages *IIterable[string], geographicRegion string) *ICurrencyFormatter {
	var _result *ICurrencyFormatter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCurrencyFormatterCodeContext, uintptr(unsafe.Pointer(this)), NewHStr(currencyCode).Ptr, uintptr(unsafe.Pointer(languages)), NewHStr(geographicRegion).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0D018C9A-E393-46B8-B830-7A69C8F89FBB
var IID_IDecimalFormatterFactory = syscall.GUID{0x0D018C9A, 0xE393, 0x46B8,
	[8]byte{0xB8, 0x30, 0x7A, 0x69, 0xC8, 0xF8, 0x9F, 0xBB}}

type IDecimalFormatterFactoryInterface interface {
	win32.IInspectableInterface
	CreateDecimalFormatter(languages *IIterable[string], geographicRegion string) *INumberFormatter
}

type IDecimalFormatterFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateDecimalFormatter uintptr
}

type IDecimalFormatterFactory struct {
	win32.IInspectable
}

func (this *IDecimalFormatterFactory) Vtbl() *IDecimalFormatterFactoryVtbl {
	return (*IDecimalFormatterFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDecimalFormatterFactory) CreateDecimalFormatter(languages *IIterable[string], geographicRegion string) *INumberFormatter {
	var _result *INumberFormatter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDecimalFormatter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(languages)), NewHStr(geographicRegion).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 70A64FF8-66AB-4155-9DA1-739E46764543
var IID_IIncrementNumberRounder = syscall.GUID{0x70A64FF8, 0x66AB, 0x4155,
	[8]byte{0x9D, 0xA1, 0x73, 0x9E, 0x46, 0x76, 0x45, 0x43}}

type IIncrementNumberRounderInterface interface {
	win32.IInspectableInterface
	Get_RoundingAlgorithm() RoundingAlgorithm
	Put_RoundingAlgorithm(value RoundingAlgorithm)
	Get_Increment() float64
	Put_Increment(value float64)
}

type IIncrementNumberRounderVtbl struct {
	win32.IInspectableVtbl
	Get_RoundingAlgorithm uintptr
	Put_RoundingAlgorithm uintptr
	Get_Increment         uintptr
	Put_Increment         uintptr
}

type IIncrementNumberRounder struct {
	win32.IInspectable
}

func (this *IIncrementNumberRounder) Vtbl() *IIncrementNumberRounderVtbl {
	return (*IIncrementNumberRounderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IIncrementNumberRounder) Get_RoundingAlgorithm() RoundingAlgorithm {
	var _result RoundingAlgorithm
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RoundingAlgorithm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IIncrementNumberRounder) Put_RoundingAlgorithm(value RoundingAlgorithm) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RoundingAlgorithm, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IIncrementNumberRounder) Get_Increment() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Increment, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IIncrementNumberRounder) Put_Increment(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Increment, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// A5007C49-7676-4DB7-8631-1B6FF265CAA9
var IID_INumberFormatter = syscall.GUID{0xA5007C49, 0x7676, 0x4DB7,
	[8]byte{0x86, 0x31, 0x1B, 0x6F, 0xF2, 0x65, 0xCA, 0xA9}}

type INumberFormatterInterface interface {
	win32.IInspectableInterface
	FormatInt(value int64) string
	FormatUInt(value uint64) string
	FormatDouble(value float64) string
}

type INumberFormatterVtbl struct {
	win32.IInspectableVtbl
	FormatInt    uintptr
	FormatUInt   uintptr
	FormatDouble uintptr
}

type INumberFormatter struct {
	win32.IInspectable
}

func (this *INumberFormatter) Vtbl() *INumberFormatterVtbl {
	return (*INumberFormatterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INumberFormatter) FormatInt(value int64) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FormatInt, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumberFormatter) FormatUInt(value uint64) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FormatUInt, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumberFormatter) FormatDouble(value float64) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FormatDouble, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// D4A8C1F0-80D0-4B0D-A89E-882C1E8F8310
var IID_INumberFormatter2 = syscall.GUID{0xD4A8C1F0, 0x80D0, 0x4B0D,
	[8]byte{0xA8, 0x9E, 0x88, 0x2C, 0x1E, 0x8F, 0x83, 0x10}}

type INumberFormatter2Interface interface {
	win32.IInspectableInterface
	FormatInt(value int64) string
	FormatUInt(value uint64) string
	FormatDouble(value float64) string
}

type INumberFormatter2Vtbl struct {
	win32.IInspectableVtbl
	FormatInt    uintptr
	FormatUInt   uintptr
	FormatDouble uintptr
}

type INumberFormatter2 struct {
	win32.IInspectable
}

func (this *INumberFormatter2) Vtbl() *INumberFormatter2Vtbl {
	return (*INumberFormatter2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INumberFormatter2) FormatInt(value int64) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FormatInt, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumberFormatter2) FormatUInt(value uint64) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FormatUInt, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumberFormatter2) FormatDouble(value float64) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FormatDouble, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 80332D21-AEE1-4A39-BAA2-07ED8C96DAF6
var IID_INumberFormatterOptions = syscall.GUID{0x80332D21, 0xAEE1, 0x4A39,
	[8]byte{0xBA, 0xA2, 0x07, 0xED, 0x8C, 0x96, 0xDA, 0xF6}}

type INumberFormatterOptionsInterface interface {
	win32.IInspectableInterface
	Get_Languages() *IVectorView[string]
	Get_GeographicRegion() string
	Get_IntegerDigits() int32
	Put_IntegerDigits(value int32)
	Get_FractionDigits() int32
	Put_FractionDigits(value int32)
	Get_IsGrouped() bool
	Put_IsGrouped(value bool)
	Get_IsDecimalPointAlwaysDisplayed() bool
	Put_IsDecimalPointAlwaysDisplayed(value bool)
	Get_NumeralSystem() string
	Put_NumeralSystem(value string)
	Get_ResolvedLanguage() string
	Get_ResolvedGeographicRegion() string
}

type INumberFormatterOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_Languages                     uintptr
	Get_GeographicRegion              uintptr
	Get_IntegerDigits                 uintptr
	Put_IntegerDigits                 uintptr
	Get_FractionDigits                uintptr
	Put_FractionDigits                uintptr
	Get_IsGrouped                     uintptr
	Put_IsGrouped                     uintptr
	Get_IsDecimalPointAlwaysDisplayed uintptr
	Put_IsDecimalPointAlwaysDisplayed uintptr
	Get_NumeralSystem                 uintptr
	Put_NumeralSystem                 uintptr
	Get_ResolvedLanguage              uintptr
	Get_ResolvedGeographicRegion      uintptr
}

type INumberFormatterOptions struct {
	win32.IInspectable
}

func (this *INumberFormatterOptions) Vtbl() *INumberFormatterOptionsVtbl {
	return (*INumberFormatterOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INumberFormatterOptions) Get_Languages() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Languages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INumberFormatterOptions) Get_GeographicRegion() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GeographicRegion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumberFormatterOptions) Get_IntegerDigits() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IntegerDigits, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INumberFormatterOptions) Put_IntegerDigits(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IntegerDigits, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *INumberFormatterOptions) Get_FractionDigits() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FractionDigits, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INumberFormatterOptions) Put_FractionDigits(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FractionDigits, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *INumberFormatterOptions) Get_IsGrouped() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsGrouped, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INumberFormatterOptions) Put_IsGrouped(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsGrouped, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *INumberFormatterOptions) Get_IsDecimalPointAlwaysDisplayed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDecimalPointAlwaysDisplayed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INumberFormatterOptions) Put_IsDecimalPointAlwaysDisplayed(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsDecimalPointAlwaysDisplayed, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *INumberFormatterOptions) Get_NumeralSystem() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NumeralSystem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumberFormatterOptions) Put_NumeralSystem(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NumeralSystem, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *INumberFormatterOptions) Get_ResolvedLanguage() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResolvedLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumberFormatterOptions) Get_ResolvedGeographicRegion() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResolvedGeographicRegion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// E6659412-4A13-4A53-83A1-392FBE4CFF9F
var IID_INumberParser = syscall.GUID{0xE6659412, 0x4A13, 0x4A53,
	[8]byte{0x83, 0xA1, 0x39, 0x2F, 0xBE, 0x4C, 0xFF, 0x9F}}

type INumberParserInterface interface {
	win32.IInspectableInterface
	ParseInt(text string) *IReference[int64]
	ParseUInt(text string) *IReference[uint64]
	ParseDouble(text string) *IReference[float64]
}

type INumberParserVtbl struct {
	win32.IInspectableVtbl
	ParseInt    uintptr
	ParseUInt   uintptr
	ParseDouble uintptr
}

type INumberParser struct {
	win32.IInspectable
}

func (this *INumberParser) Vtbl() *INumberParserVtbl {
	return (*INumberParserVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INumberParser) ParseInt(text string) *IReference[int64] {
	var _result *IReference[int64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ParseInt, uintptr(unsafe.Pointer(this)), NewHStr(text).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INumberParser) ParseUInt(text string) *IReference[uint64] {
	var _result *IReference[uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ParseUInt, uintptr(unsafe.Pointer(this)), NewHStr(text).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INumberParser) ParseDouble(text string) *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ParseDouble, uintptr(unsafe.Pointer(this)), NewHStr(text).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5473C375-38ED-4631-B80C-EF34FC48B7F5
var IID_INumberRounder = syscall.GUID{0x5473C375, 0x38ED, 0x4631,
	[8]byte{0xB8, 0x0C, 0xEF, 0x34, 0xFC, 0x48, 0xB7, 0xF5}}

type INumberRounderInterface interface {
	win32.IInspectableInterface
	RoundInt32(value int32) int32
	RoundUInt32(value uint32) uint32
	RoundInt64(value int64) int64
	RoundUInt64(value uint64) uint64
	RoundSingle(value float32) float32
	RoundDouble(value float64) float64
}

type INumberRounderVtbl struct {
	win32.IInspectableVtbl
	RoundInt32  uintptr
	RoundUInt32 uintptr
	RoundInt64  uintptr
	RoundUInt64 uintptr
	RoundSingle uintptr
	RoundDouble uintptr
}

type INumberRounder struct {
	win32.IInspectable
}

func (this *INumberRounder) Vtbl() *INumberRounderVtbl {
	return (*INumberRounderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INumberRounder) RoundInt32(value int32) int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RoundInt32, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INumberRounder) RoundUInt32(value uint32) uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RoundUInt32, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INumberRounder) RoundInt64(value int64) int64 {
	var _result int64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RoundInt64, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INumberRounder) RoundUInt64(value uint64) uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RoundUInt64, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INumberRounder) RoundSingle(value float32) float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RoundSingle, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INumberRounder) RoundDouble(value float64) float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RoundDouble, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3B088433-646F-4EFE-8D48-66EB2E49E736
var IID_INumberRounderOption = syscall.GUID{0x3B088433, 0x646F, 0x4EFE,
	[8]byte{0x8D, 0x48, 0x66, 0xEB, 0x2E, 0x49, 0xE7, 0x36}}

type INumberRounderOptionInterface interface {
	win32.IInspectableInterface
	Get_NumberRounder() *INumberRounder
	Put_NumberRounder(value *INumberRounder)
}

type INumberRounderOptionVtbl struct {
	win32.IInspectableVtbl
	Get_NumberRounder uintptr
	Put_NumberRounder uintptr
}

type INumberRounderOption struct {
	win32.IInspectable
}

func (this *INumberRounderOption) Vtbl() *INumberRounderOptionVtbl {
	return (*INumberRounderOptionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INumberRounderOption) Get_NumberRounder() *INumberRounder {
	var _result *INumberRounder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NumberRounder, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INumberRounderOption) Put_NumberRounder(value *INumberRounder) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NumberRounder, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 28F5BC2C-8C23-4234-AD2E-FA5A3A426E9B
var IID_INumeralSystemTranslator = syscall.GUID{0x28F5BC2C, 0x8C23, 0x4234,
	[8]byte{0xAD, 0x2E, 0xFA, 0x5A, 0x3A, 0x42, 0x6E, 0x9B}}

type INumeralSystemTranslatorInterface interface {
	win32.IInspectableInterface
	Get_Languages() *IVectorView[string]
	Get_ResolvedLanguage() string
	Get_NumeralSystem() string
	Put_NumeralSystem(value string)
	TranslateNumerals(value string) string
}

type INumeralSystemTranslatorVtbl struct {
	win32.IInspectableVtbl
	Get_Languages        uintptr
	Get_ResolvedLanguage uintptr
	Get_NumeralSystem    uintptr
	Put_NumeralSystem    uintptr
	TranslateNumerals    uintptr
}

type INumeralSystemTranslator struct {
	win32.IInspectable
}

func (this *INumeralSystemTranslator) Vtbl() *INumeralSystemTranslatorVtbl {
	return (*INumeralSystemTranslatorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INumeralSystemTranslator) Get_Languages() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Languages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INumeralSystemTranslator) Get_ResolvedLanguage() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResolvedLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemTranslator) Get_NumeralSystem() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NumeralSystem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemTranslator) Put_NumeralSystem(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NumeralSystem, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *INumeralSystemTranslator) TranslateNumerals(value string) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TranslateNumerals, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 9630C8DA-36EF-4D88-A85C-6F0D98D620A6
var IID_INumeralSystemTranslatorFactory = syscall.GUID{0x9630C8DA, 0x36EF, 0x4D88,
	[8]byte{0xA8, 0x5C, 0x6F, 0x0D, 0x98, 0xD6, 0x20, 0xA6}}

type INumeralSystemTranslatorFactoryInterface interface {
	win32.IInspectableInterface
	Create(languages *IIterable[string]) *INumeralSystemTranslator
}

type INumeralSystemTranslatorFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type INumeralSystemTranslatorFactory struct {
	win32.IInspectable
}

func (this *INumeralSystemTranslatorFactory) Vtbl() *INumeralSystemTranslatorFactoryVtbl {
	return (*INumeralSystemTranslatorFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INumeralSystemTranslatorFactory) Create(languages *IIterable[string]) *INumeralSystemTranslator {
	var _result *INumeralSystemTranslator
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(languages)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B7828AEF-FED4-4018-A6E2-E09961E03765
var IID_IPercentFormatterFactory = syscall.GUID{0xB7828AEF, 0xFED4, 0x4018,
	[8]byte{0xA6, 0xE2, 0xE0, 0x99, 0x61, 0xE0, 0x37, 0x65}}

type IPercentFormatterFactoryInterface interface {
	win32.IInspectableInterface
	CreatePercentFormatter(languages *IIterable[string], geographicRegion string) *INumberFormatter
}

type IPercentFormatterFactoryVtbl struct {
	win32.IInspectableVtbl
	CreatePercentFormatter uintptr
}

type IPercentFormatterFactory struct {
	win32.IInspectable
}

func (this *IPercentFormatterFactory) Vtbl() *IPercentFormatterFactoryVtbl {
	return (*IPercentFormatterFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPercentFormatterFactory) CreatePercentFormatter(languages *IIterable[string], geographicRegion string) *INumberFormatter {
	var _result *INumberFormatter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreatePercentFormatter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(languages)), NewHStr(geographicRegion).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2B37B4AC-E638-4ED5-A998-62F6B06A49AE
var IID_IPermilleFormatterFactory = syscall.GUID{0x2B37B4AC, 0xE638, 0x4ED5,
	[8]byte{0xA9, 0x98, 0x62, 0xF6, 0xB0, 0x6A, 0x49, 0xAE}}

type IPermilleFormatterFactoryInterface interface {
	win32.IInspectableInterface
	CreatePermilleFormatter(languages *IIterable[string], geographicRegion string) *INumberFormatter
}

type IPermilleFormatterFactoryVtbl struct {
	win32.IInspectableVtbl
	CreatePermilleFormatter uintptr
}

type IPermilleFormatterFactory struct {
	win32.IInspectable
}

func (this *IPermilleFormatterFactory) Vtbl() *IPermilleFormatterFactoryVtbl {
	return (*IPermilleFormatterFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPermilleFormatterFactory) CreatePermilleFormatter(languages *IIterable[string], geographicRegion string) *INumberFormatter {
	var _result *INumberFormatter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreatePermilleFormatter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(languages)), NewHStr(geographicRegion).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FD1CDD31-0A3C-49C4-A642-96A1564F4F30
var IID_ISignedZeroOption = syscall.GUID{0xFD1CDD31, 0x0A3C, 0x49C4,
	[8]byte{0xA6, 0x42, 0x96, 0xA1, 0x56, 0x4F, 0x4F, 0x30}}

type ISignedZeroOptionInterface interface {
	win32.IInspectableInterface
	Get_IsZeroSigned() bool
	Put_IsZeroSigned(value bool)
}

type ISignedZeroOptionVtbl struct {
	win32.IInspectableVtbl
	Get_IsZeroSigned uintptr
	Put_IsZeroSigned uintptr
}

type ISignedZeroOption struct {
	win32.IInspectable
}

func (this *ISignedZeroOption) Vtbl() *ISignedZeroOptionVtbl {
	return (*ISignedZeroOptionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISignedZeroOption) Get_IsZeroSigned() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsZeroSigned, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISignedZeroOption) Put_IsZeroSigned(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsZeroSigned, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// F5941BCA-6646-4913-8C76-1B191FF94DFD
var IID_ISignificantDigitsNumberRounder = syscall.GUID{0xF5941BCA, 0x6646, 0x4913,
	[8]byte{0x8C, 0x76, 0x1B, 0x19, 0x1F, 0xF9, 0x4D, 0xFD}}

type ISignificantDigitsNumberRounderInterface interface {
	win32.IInspectableInterface
	Get_RoundingAlgorithm() RoundingAlgorithm
	Put_RoundingAlgorithm(value RoundingAlgorithm)
	Get_SignificantDigits() uint32
	Put_SignificantDigits(value uint32)
}

type ISignificantDigitsNumberRounderVtbl struct {
	win32.IInspectableVtbl
	Get_RoundingAlgorithm uintptr
	Put_RoundingAlgorithm uintptr
	Get_SignificantDigits uintptr
	Put_SignificantDigits uintptr
}

type ISignificantDigitsNumberRounder struct {
	win32.IInspectable
}

func (this *ISignificantDigitsNumberRounder) Vtbl() *ISignificantDigitsNumberRounderVtbl {
	return (*ISignificantDigitsNumberRounderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISignificantDigitsNumberRounder) Get_RoundingAlgorithm() RoundingAlgorithm {
	var _result RoundingAlgorithm
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RoundingAlgorithm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISignificantDigitsNumberRounder) Put_RoundingAlgorithm(value RoundingAlgorithm) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RoundingAlgorithm, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISignificantDigitsNumberRounder) Get_SignificantDigits() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SignificantDigits, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISignificantDigitsNumberRounder) Put_SignificantDigits(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SignificantDigits, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 1D4DFCDD-2D43-4EE8-BBF1-C1B26A711A58
var IID_ISignificantDigitsOption = syscall.GUID{0x1D4DFCDD, 0x2D43, 0x4EE8,
	[8]byte{0xBB, 0xF1, 0xC1, 0xB2, 0x6A, 0x71, 0x1A, 0x58}}

type ISignificantDigitsOptionInterface interface {
	win32.IInspectableInterface
	Get_SignificantDigits() int32
	Put_SignificantDigits(value int32)
}

type ISignificantDigitsOptionVtbl struct {
	win32.IInspectableVtbl
	Get_SignificantDigits uintptr
	Put_SignificantDigits uintptr
}

type ISignificantDigitsOption struct {
	win32.IInspectable
}

func (this *ISignificantDigitsOption) Vtbl() *ISignificantDigitsOptionVtbl {
	return (*ISignificantDigitsOptionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISignificantDigitsOption) Get_SignificantDigits() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SignificantDigits, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISignificantDigitsOption) Put_SignificantDigits(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SignificantDigits, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// classes

type CurrencyFormatter struct {
	RtClass
	*ICurrencyFormatter
}

func NewCurrencyFormatter_CreateCurrencyFormatterCode(currencyCode string) *CurrencyFormatter {
	hs := NewHStr("Windows.Globalization.NumberFormatting.CurrencyFormatter")
	var pFac *ICurrencyFormatterFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICurrencyFormatterFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ICurrencyFormatter
	p = pFac.CreateCurrencyFormatterCode(currencyCode)
	result := &CurrencyFormatter{
		RtClass:            RtClass{PInspect: &p.IInspectable},
		ICurrencyFormatter: p,
	}
	com.AddToScope(result)
	return result
}

func NewCurrencyFormatter_CreateCurrencyFormatterCodeContext(currencyCode string, languages *IIterable[string], geographicRegion string) *CurrencyFormatter {
	hs := NewHStr("Windows.Globalization.NumberFormatting.CurrencyFormatter")
	var pFac *ICurrencyFormatterFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICurrencyFormatterFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ICurrencyFormatter
	p = pFac.CreateCurrencyFormatterCodeContext(currencyCode, languages, geographicRegion)
	result := &CurrencyFormatter{
		RtClass:            RtClass{PInspect: &p.IInspectable},
		ICurrencyFormatter: p,
	}
	com.AddToScope(result)
	return result
}

type DecimalFormatter struct {
	RtClass
	*INumberFormatter
}

func NewDecimalFormatter() *DecimalFormatter {
	hs := NewHStr("Windows.Globalization.NumberFormatting.DecimalFormatter")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &DecimalFormatter{
		RtClass:          RtClass{PInspect: p},
		INumberFormatter: (*INumberFormatter)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewDecimalFormatter_CreateDecimalFormatter(languages *IIterable[string], geographicRegion string) *DecimalFormatter {
	hs := NewHStr("Windows.Globalization.NumberFormatting.DecimalFormatter")
	var pFac *IDecimalFormatterFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDecimalFormatterFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *INumberFormatter
	p = pFac.CreateDecimalFormatter(languages, geographicRegion)
	result := &DecimalFormatter{
		RtClass:          RtClass{PInspect: &p.IInspectable},
		INumberFormatter: p,
	}
	com.AddToScope(result)
	return result
}

type IncrementNumberRounder struct {
	RtClass
	*INumberRounder
}

func NewIncrementNumberRounder() *IncrementNumberRounder {
	hs := NewHStr("Windows.Globalization.NumberFormatting.IncrementNumberRounder")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &IncrementNumberRounder{
		RtClass:        RtClass{PInspect: p},
		INumberRounder: (*INumberRounder)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type NumeralSystemTranslator struct {
	RtClass
	*INumeralSystemTranslator
}

func NewNumeralSystemTranslator() *NumeralSystemTranslator {
	hs := NewHStr("Windows.Globalization.NumberFormatting.NumeralSystemTranslator")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &NumeralSystemTranslator{
		RtClass:                  RtClass{PInspect: p},
		INumeralSystemTranslator: (*INumeralSystemTranslator)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewNumeralSystemTranslator_Create(languages *IIterable[string]) *NumeralSystemTranslator {
	hs := NewHStr("Windows.Globalization.NumberFormatting.NumeralSystemTranslator")
	var pFac *INumeralSystemTranslatorFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_INumeralSystemTranslatorFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *INumeralSystemTranslator
	p = pFac.Create(languages)
	result := &NumeralSystemTranslator{
		RtClass:                  RtClass{PInspect: &p.IInspectable},
		INumeralSystemTranslator: p,
	}
	com.AddToScope(result)
	return result
}

type PercentFormatter struct {
	RtClass
	*INumberFormatter
}

func NewPercentFormatter() *PercentFormatter {
	hs := NewHStr("Windows.Globalization.NumberFormatting.PercentFormatter")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &PercentFormatter{
		RtClass:          RtClass{PInspect: p},
		INumberFormatter: (*INumberFormatter)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewPercentFormatter_CreatePercentFormatter(languages *IIterable[string], geographicRegion string) *PercentFormatter {
	hs := NewHStr("Windows.Globalization.NumberFormatting.PercentFormatter")
	var pFac *IPercentFormatterFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPercentFormatterFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *INumberFormatter
	p = pFac.CreatePercentFormatter(languages, geographicRegion)
	result := &PercentFormatter{
		RtClass:          RtClass{PInspect: &p.IInspectable},
		INumberFormatter: p,
	}
	com.AddToScope(result)
	return result
}

type PermilleFormatter struct {
	RtClass
	*INumberFormatter
}

func NewPermilleFormatter() *PermilleFormatter {
	hs := NewHStr("Windows.Globalization.NumberFormatting.PermilleFormatter")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &PermilleFormatter{
		RtClass:          RtClass{PInspect: p},
		INumberFormatter: (*INumberFormatter)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewPermilleFormatter_CreatePermilleFormatter(languages *IIterable[string], geographicRegion string) *PermilleFormatter {
	hs := NewHStr("Windows.Globalization.NumberFormatting.PermilleFormatter")
	var pFac *IPermilleFormatterFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPermilleFormatterFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *INumberFormatter
	p = pFac.CreatePermilleFormatter(languages, geographicRegion)
	result := &PermilleFormatter{
		RtClass:          RtClass{PInspect: &p.IInspectable},
		INumberFormatter: p,
	}
	com.AddToScope(result)
	return result
}

type SignificantDigitsNumberRounder struct {
	RtClass
	*INumberRounder
}

func NewSignificantDigitsNumberRounder() *SignificantDigitsNumberRounder {
	hs := NewHStr("Windows.Globalization.NumberFormatting.SignificantDigitsNumberRounder")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &SignificantDigitsNumberRounder{
		RtClass:        RtClass{PInspect: p},
		INumberRounder: (*INumberRounder)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}
