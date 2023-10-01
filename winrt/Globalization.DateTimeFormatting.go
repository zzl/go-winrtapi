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
type DayFormat int32

const (
	DayFormat_None    DayFormat = 0
	DayFormat_Default DayFormat = 1
)

// enum
type DayOfWeekFormat int32

const (
	DayOfWeekFormat_None        DayOfWeekFormat = 0
	DayOfWeekFormat_Default     DayOfWeekFormat = 1
	DayOfWeekFormat_Abbreviated DayOfWeekFormat = 2
	DayOfWeekFormat_Full        DayOfWeekFormat = 3
)

// enum
type HourFormat int32

const (
	HourFormat_None    HourFormat = 0
	HourFormat_Default HourFormat = 1
)

// enum
type MinuteFormat int32

const (
	MinuteFormat_None    MinuteFormat = 0
	MinuteFormat_Default MinuteFormat = 1
)

// enum
type MonthFormat int32

const (
	MonthFormat_None        MonthFormat = 0
	MonthFormat_Default     MonthFormat = 1
	MonthFormat_Abbreviated MonthFormat = 2
	MonthFormat_Full        MonthFormat = 3
	MonthFormat_Numeric     MonthFormat = 4
)

// enum
type SecondFormat int32

const (
	SecondFormat_None    SecondFormat = 0
	SecondFormat_Default SecondFormat = 1
)

// enum
type YearFormat int32

const (
	YearFormat_None        YearFormat = 0
	YearFormat_Default     YearFormat = 1
	YearFormat_Abbreviated YearFormat = 2
	YearFormat_Full        YearFormat = 3
)

// interfaces

// 95EECA10-73E0-4E4B-A183-3D6AD0BA35EC
var IID_IDateTimeFormatter = syscall.GUID{0x95EECA10, 0x73E0, 0x4E4B,
	[8]byte{0xA1, 0x83, 0x3D, 0x6A, 0xD0, 0xBA, 0x35, 0xEC}}

type IDateTimeFormatterInterface interface {
	win32.IInspectableInterface
	Get_Languages() *IVectorView[string]
	Get_GeographicRegion() string
	Get_Calendar() string
	Get_Clock() string
	Get_NumeralSystem() string
	Put_NumeralSystem(value string)
	Get_Patterns() *IVectorView[string]
	Get_Template() string
	Format(value DateTime) string
	Get_IncludeYear() YearFormat
	Get_IncludeMonth() MonthFormat
	Get_IncludeDayOfWeek() DayOfWeekFormat
	Get_IncludeDay() DayFormat
	Get_IncludeHour() HourFormat
	Get_IncludeMinute() MinuteFormat
	Get_IncludeSecond() SecondFormat
	Get_ResolvedLanguage() string
	Get_ResolvedGeographicRegion() string
}

type IDateTimeFormatterVtbl struct {
	win32.IInspectableVtbl
	Get_Languages                uintptr
	Get_GeographicRegion         uintptr
	Get_Calendar                 uintptr
	Get_Clock                    uintptr
	Get_NumeralSystem            uintptr
	Put_NumeralSystem            uintptr
	Get_Patterns                 uintptr
	Get_Template                 uintptr
	Format                       uintptr
	Get_IncludeYear              uintptr
	Get_IncludeMonth             uintptr
	Get_IncludeDayOfWeek         uintptr
	Get_IncludeDay               uintptr
	Get_IncludeHour              uintptr
	Get_IncludeMinute            uintptr
	Get_IncludeSecond            uintptr
	Get_ResolvedLanguage         uintptr
	Get_ResolvedGeographicRegion uintptr
}

type IDateTimeFormatter struct {
	win32.IInspectable
}

func (this *IDateTimeFormatter) Vtbl() *IDateTimeFormatterVtbl {
	return (*IDateTimeFormatterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDateTimeFormatter) Get_Languages() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Languages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDateTimeFormatter) Get_GeographicRegion() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GeographicRegion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDateTimeFormatter) Get_Calendar() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Calendar, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDateTimeFormatter) Get_Clock() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Clock, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDateTimeFormatter) Get_NumeralSystem() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NumeralSystem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDateTimeFormatter) Put_NumeralSystem(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NumeralSystem, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IDateTimeFormatter) Get_Patterns() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Patterns, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDateTimeFormatter) Get_Template() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Template, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDateTimeFormatter) Format(value DateTime) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Format, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDateTimeFormatter) Get_IncludeYear() YearFormat {
	var _result YearFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IncludeYear, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDateTimeFormatter) Get_IncludeMonth() MonthFormat {
	var _result MonthFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IncludeMonth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDateTimeFormatter) Get_IncludeDayOfWeek() DayOfWeekFormat {
	var _result DayOfWeekFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IncludeDayOfWeek, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDateTimeFormatter) Get_IncludeDay() DayFormat {
	var _result DayFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IncludeDay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDateTimeFormatter) Get_IncludeHour() HourFormat {
	var _result HourFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IncludeHour, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDateTimeFormatter) Get_IncludeMinute() MinuteFormat {
	var _result MinuteFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IncludeMinute, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDateTimeFormatter) Get_IncludeSecond() SecondFormat {
	var _result SecondFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IncludeSecond, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDateTimeFormatter) Get_ResolvedLanguage() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResolvedLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDateTimeFormatter) Get_ResolvedGeographicRegion() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResolvedGeographicRegion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 27C91A86-BDAA-4FD0-9E36-671D5AA5EE03
var IID_IDateTimeFormatter2 = syscall.GUID{0x27C91A86, 0xBDAA, 0x4FD0,
	[8]byte{0x9E, 0x36, 0x67, 0x1D, 0x5A, 0xA5, 0xEE, 0x03}}

type IDateTimeFormatter2Interface interface {
	win32.IInspectableInterface
	FormatUsingTimeZone(datetime DateTime, timeZoneId string) string
}

type IDateTimeFormatter2Vtbl struct {
	win32.IInspectableVtbl
	FormatUsingTimeZone uintptr
}

type IDateTimeFormatter2 struct {
	win32.IInspectable
}

func (this *IDateTimeFormatter2) Vtbl() *IDateTimeFormatter2Vtbl {
	return (*IDateTimeFormatter2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDateTimeFormatter2) FormatUsingTimeZone(datetime DateTime, timeZoneId string) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FormatUsingTimeZone, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&datetime)), NewHStr(timeZoneId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// EC8D8A53-1A2E-412D-8815-3B745FB1A2A0
var IID_IDateTimeFormatterFactory = syscall.GUID{0xEC8D8A53, 0x1A2E, 0x412D,
	[8]byte{0x88, 0x15, 0x3B, 0x74, 0x5F, 0xB1, 0xA2, 0xA0}}

type IDateTimeFormatterFactoryInterface interface {
	win32.IInspectableInterface
	CreateDateTimeFormatter(formatTemplate string) *IDateTimeFormatter
	CreateDateTimeFormatterLanguages(formatTemplate string, languages *IIterable[string]) *IDateTimeFormatter
	CreateDateTimeFormatterContext(formatTemplate string, languages *IIterable[string], geographicRegion string, calendar string, clock string) *IDateTimeFormatter
	CreateDateTimeFormatterDate(yearFormat YearFormat, monthFormat MonthFormat, dayFormat DayFormat, dayOfWeekFormat DayOfWeekFormat) *IDateTimeFormatter
	CreateDateTimeFormatterTime(hourFormat HourFormat, minuteFormat MinuteFormat, secondFormat SecondFormat) *IDateTimeFormatter
	CreateDateTimeFormatterDateTimeLanguages(yearFormat YearFormat, monthFormat MonthFormat, dayFormat DayFormat, dayOfWeekFormat DayOfWeekFormat, hourFormat HourFormat, minuteFormat MinuteFormat, secondFormat SecondFormat, languages *IIterable[string]) *IDateTimeFormatter
	CreateDateTimeFormatterDateTimeContext(yearFormat YearFormat, monthFormat MonthFormat, dayFormat DayFormat, dayOfWeekFormat DayOfWeekFormat, hourFormat HourFormat, minuteFormat MinuteFormat, secondFormat SecondFormat, languages *IIterable[string], geographicRegion string, calendar string, clock string) *IDateTimeFormatter
}

type IDateTimeFormatterFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateDateTimeFormatter                  uintptr
	CreateDateTimeFormatterLanguages         uintptr
	CreateDateTimeFormatterContext           uintptr
	CreateDateTimeFormatterDate              uintptr
	CreateDateTimeFormatterTime              uintptr
	CreateDateTimeFormatterDateTimeLanguages uintptr
	CreateDateTimeFormatterDateTimeContext   uintptr
}

type IDateTimeFormatterFactory struct {
	win32.IInspectable
}

func (this *IDateTimeFormatterFactory) Vtbl() *IDateTimeFormatterFactoryVtbl {
	return (*IDateTimeFormatterFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDateTimeFormatterFactory) CreateDateTimeFormatter(formatTemplate string) *IDateTimeFormatter {
	var _result *IDateTimeFormatter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDateTimeFormatter, uintptr(unsafe.Pointer(this)), NewHStr(formatTemplate).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDateTimeFormatterFactory) CreateDateTimeFormatterLanguages(formatTemplate string, languages *IIterable[string]) *IDateTimeFormatter {
	var _result *IDateTimeFormatter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDateTimeFormatterLanguages, uintptr(unsafe.Pointer(this)), NewHStr(formatTemplate).Ptr, uintptr(unsafe.Pointer(languages)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDateTimeFormatterFactory) CreateDateTimeFormatterContext(formatTemplate string, languages *IIterable[string], geographicRegion string, calendar string, clock string) *IDateTimeFormatter {
	var _result *IDateTimeFormatter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDateTimeFormatterContext, uintptr(unsafe.Pointer(this)), NewHStr(formatTemplate).Ptr, uintptr(unsafe.Pointer(languages)), NewHStr(geographicRegion).Ptr, NewHStr(calendar).Ptr, NewHStr(clock).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDateTimeFormatterFactory) CreateDateTimeFormatterDate(yearFormat YearFormat, monthFormat MonthFormat, dayFormat DayFormat, dayOfWeekFormat DayOfWeekFormat) *IDateTimeFormatter {
	var _result *IDateTimeFormatter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDateTimeFormatterDate, uintptr(unsafe.Pointer(this)), uintptr(yearFormat), uintptr(monthFormat), uintptr(dayFormat), uintptr(dayOfWeekFormat), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDateTimeFormatterFactory) CreateDateTimeFormatterTime(hourFormat HourFormat, minuteFormat MinuteFormat, secondFormat SecondFormat) *IDateTimeFormatter {
	var _result *IDateTimeFormatter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDateTimeFormatterTime, uintptr(unsafe.Pointer(this)), uintptr(hourFormat), uintptr(minuteFormat), uintptr(secondFormat), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDateTimeFormatterFactory) CreateDateTimeFormatterDateTimeLanguages(yearFormat YearFormat, monthFormat MonthFormat, dayFormat DayFormat, dayOfWeekFormat DayOfWeekFormat, hourFormat HourFormat, minuteFormat MinuteFormat, secondFormat SecondFormat, languages *IIterable[string]) *IDateTimeFormatter {
	var _result *IDateTimeFormatter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDateTimeFormatterDateTimeLanguages, uintptr(unsafe.Pointer(this)), uintptr(yearFormat), uintptr(monthFormat), uintptr(dayFormat), uintptr(dayOfWeekFormat), uintptr(hourFormat), uintptr(minuteFormat), uintptr(secondFormat), uintptr(unsafe.Pointer(languages)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDateTimeFormatterFactory) CreateDateTimeFormatterDateTimeContext(yearFormat YearFormat, monthFormat MonthFormat, dayFormat DayFormat, dayOfWeekFormat DayOfWeekFormat, hourFormat HourFormat, minuteFormat MinuteFormat, secondFormat SecondFormat, languages *IIterable[string], geographicRegion string, calendar string, clock string) *IDateTimeFormatter {
	var _result *IDateTimeFormatter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDateTimeFormatterDateTimeContext, uintptr(unsafe.Pointer(this)), uintptr(yearFormat), uintptr(monthFormat), uintptr(dayFormat), uintptr(dayOfWeekFormat), uintptr(hourFormat), uintptr(minuteFormat), uintptr(secondFormat), uintptr(unsafe.Pointer(languages)), NewHStr(geographicRegion).Ptr, NewHStr(calendar).Ptr, NewHStr(clock).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BFCDE7C0-DF4C-4A2E-9012-F47DAF3F1212
var IID_IDateTimeFormatterStatics = syscall.GUID{0xBFCDE7C0, 0xDF4C, 0x4A2E,
	[8]byte{0x90, 0x12, 0xF4, 0x7D, 0xAF, 0x3F, 0x12, 0x12}}

type IDateTimeFormatterStaticsInterface interface {
	win32.IInspectableInterface
	Get_LongDate() *IDateTimeFormatter
	Get_LongTime() *IDateTimeFormatter
	Get_ShortDate() *IDateTimeFormatter
	Get_ShortTime() *IDateTimeFormatter
}

type IDateTimeFormatterStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_LongDate  uintptr
	Get_LongTime  uintptr
	Get_ShortDate uintptr
	Get_ShortTime uintptr
}

type IDateTimeFormatterStatics struct {
	win32.IInspectable
}

func (this *IDateTimeFormatterStatics) Vtbl() *IDateTimeFormatterStaticsVtbl {
	return (*IDateTimeFormatterStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDateTimeFormatterStatics) Get_LongDate() *IDateTimeFormatter {
	var _result *IDateTimeFormatter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LongDate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDateTimeFormatterStatics) Get_LongTime() *IDateTimeFormatter {
	var _result *IDateTimeFormatter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LongTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDateTimeFormatterStatics) Get_ShortDate() *IDateTimeFormatter {
	var _result *IDateTimeFormatter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShortDate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDateTimeFormatterStatics) Get_ShortTime() *IDateTimeFormatter {
	var _result *IDateTimeFormatter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShortTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type DateTimeFormatter struct {
	RtClass
	*IDateTimeFormatter
}

func NewDateTimeFormatter_CreateDateTimeFormatter(formatTemplate string) *DateTimeFormatter {
	hs := NewHStr("Windows.Globalization.DateTimeFormatting.DateTimeFormatter")
	var pFac *IDateTimeFormatterFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDateTimeFormatterFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IDateTimeFormatter
	p = pFac.CreateDateTimeFormatter(formatTemplate)
	result := &DateTimeFormatter{
		RtClass:            RtClass{PInspect: &p.IInspectable},
		IDateTimeFormatter: p,
	}
	com.AddToScope(result)
	return result
}

func NewDateTimeFormatter_CreateDateTimeFormatterLanguages(formatTemplate string, languages *IIterable[string]) *DateTimeFormatter {
	hs := NewHStr("Windows.Globalization.DateTimeFormatting.DateTimeFormatter")
	var pFac *IDateTimeFormatterFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDateTimeFormatterFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IDateTimeFormatter
	p = pFac.CreateDateTimeFormatterLanguages(formatTemplate, languages)
	result := &DateTimeFormatter{
		RtClass:            RtClass{PInspect: &p.IInspectable},
		IDateTimeFormatter: p,
	}
	com.AddToScope(result)
	return result
}

func NewDateTimeFormatter_CreateDateTimeFormatterContext(formatTemplate string, languages *IIterable[string], geographicRegion string, calendar string, clock string) *DateTimeFormatter {
	hs := NewHStr("Windows.Globalization.DateTimeFormatting.DateTimeFormatter")
	var pFac *IDateTimeFormatterFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDateTimeFormatterFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IDateTimeFormatter
	p = pFac.CreateDateTimeFormatterContext(formatTemplate, languages, geographicRegion, calendar, clock)
	result := &DateTimeFormatter{
		RtClass:            RtClass{PInspect: &p.IInspectable},
		IDateTimeFormatter: p,
	}
	com.AddToScope(result)
	return result
}

func NewDateTimeFormatter_CreateDateTimeFormatterDate(yearFormat YearFormat, monthFormat MonthFormat, dayFormat DayFormat, dayOfWeekFormat DayOfWeekFormat) *DateTimeFormatter {
	hs := NewHStr("Windows.Globalization.DateTimeFormatting.DateTimeFormatter")
	var pFac *IDateTimeFormatterFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDateTimeFormatterFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IDateTimeFormatter
	p = pFac.CreateDateTimeFormatterDate(yearFormat, monthFormat, dayFormat, dayOfWeekFormat)
	result := &DateTimeFormatter{
		RtClass:            RtClass{PInspect: &p.IInspectable},
		IDateTimeFormatter: p,
	}
	com.AddToScope(result)
	return result
}

func NewDateTimeFormatter_CreateDateTimeFormatterTime(hourFormat HourFormat, minuteFormat MinuteFormat, secondFormat SecondFormat) *DateTimeFormatter {
	hs := NewHStr("Windows.Globalization.DateTimeFormatting.DateTimeFormatter")
	var pFac *IDateTimeFormatterFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDateTimeFormatterFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IDateTimeFormatter
	p = pFac.CreateDateTimeFormatterTime(hourFormat, minuteFormat, secondFormat)
	result := &DateTimeFormatter{
		RtClass:            RtClass{PInspect: &p.IInspectable},
		IDateTimeFormatter: p,
	}
	com.AddToScope(result)
	return result
}

func NewDateTimeFormatter_CreateDateTimeFormatterDateTimeLanguages(yearFormat YearFormat, monthFormat MonthFormat, dayFormat DayFormat, dayOfWeekFormat DayOfWeekFormat, hourFormat HourFormat, minuteFormat MinuteFormat, secondFormat SecondFormat, languages *IIterable[string]) *DateTimeFormatter {
	hs := NewHStr("Windows.Globalization.DateTimeFormatting.DateTimeFormatter")
	var pFac *IDateTimeFormatterFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDateTimeFormatterFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IDateTimeFormatter
	p = pFac.CreateDateTimeFormatterDateTimeLanguages(yearFormat, monthFormat, dayFormat, dayOfWeekFormat, hourFormat, minuteFormat, secondFormat, languages)
	result := &DateTimeFormatter{
		RtClass:            RtClass{PInspect: &p.IInspectable},
		IDateTimeFormatter: p,
	}
	com.AddToScope(result)
	return result
}

func NewDateTimeFormatter_CreateDateTimeFormatterDateTimeContext(yearFormat YearFormat, monthFormat MonthFormat, dayFormat DayFormat, dayOfWeekFormat DayOfWeekFormat, hourFormat HourFormat, minuteFormat MinuteFormat, secondFormat SecondFormat, languages *IIterable[string], geographicRegion string, calendar string, clock string) *DateTimeFormatter {
	hs := NewHStr("Windows.Globalization.DateTimeFormatting.DateTimeFormatter")
	var pFac *IDateTimeFormatterFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDateTimeFormatterFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IDateTimeFormatter
	p = pFac.CreateDateTimeFormatterDateTimeContext(yearFormat, monthFormat, dayFormat, dayOfWeekFormat, hourFormat, minuteFormat, secondFormat, languages, geographicRegion, calendar, clock)
	result := &DateTimeFormatter{
		RtClass:            RtClass{PInspect: &p.IInspectable},
		IDateTimeFormatter: p,
	}
	com.AddToScope(result)
	return result
}

func NewIDateTimeFormatterStatics() *IDateTimeFormatterStatics {
	var p *IDateTimeFormatterStatics
	hs := NewHStr("Windows.Globalization.DateTimeFormatting.DateTimeFormatter")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDateTimeFormatterStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
