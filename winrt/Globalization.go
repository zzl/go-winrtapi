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
type DayOfWeek int32

const (
	DayOfWeek_Sunday    DayOfWeek = 0
	DayOfWeek_Monday    DayOfWeek = 1
	DayOfWeek_Tuesday   DayOfWeek = 2
	DayOfWeek_Wednesday DayOfWeek = 3
	DayOfWeek_Thursday  DayOfWeek = 4
	DayOfWeek_Friday    DayOfWeek = 5
	DayOfWeek_Saturday  DayOfWeek = 6
)

// enum
type LanguageLayoutDirection int32

const (
	LanguageLayoutDirection_Ltr    LanguageLayoutDirection = 0
	LanguageLayoutDirection_Rtl    LanguageLayoutDirection = 1
	LanguageLayoutDirection_TtbLtr LanguageLayoutDirection = 2
	LanguageLayoutDirection_TtbRtl LanguageLayoutDirection = 3
)

// structs

type GlobalizationJapanesePhoneticAnalyzerContract struct {
}

// interfaces

// 75B40847-0A4C-4A92-9565-FD63C95F7AED
var IID_IApplicationLanguagesStatics = syscall.GUID{0x75B40847, 0x0A4C, 0x4A92,
	[8]byte{0x95, 0x65, 0xFD, 0x63, 0xC9, 0x5F, 0x7A, 0xED}}

type IApplicationLanguagesStaticsInterface interface {
	win32.IInspectableInterface
	Get_PrimaryLanguageOverride() string
	Put_PrimaryLanguageOverride(value string)
	Get_Languages() *IVectorView[string]
	Get_ManifestLanguages() *IVectorView[string]
}

type IApplicationLanguagesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_PrimaryLanguageOverride uintptr
	Put_PrimaryLanguageOverride uintptr
	Get_Languages               uintptr
	Get_ManifestLanguages       uintptr
}

type IApplicationLanguagesStatics struct {
	win32.IInspectable
}

func (this *IApplicationLanguagesStatics) Vtbl() *IApplicationLanguagesStaticsVtbl {
	return (*IApplicationLanguagesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationLanguagesStatics) Get_PrimaryLanguageOverride() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrimaryLanguageOverride, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IApplicationLanguagesStatics) Put_PrimaryLanguageOverride(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PrimaryLanguageOverride, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IApplicationLanguagesStatics) Get_Languages() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Languages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationLanguagesStatics) Get_ManifestLanguages() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ManifestLanguages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1DF0DE4F-072B-4D7B-8F06-CB2DB40F2BB5
var IID_IApplicationLanguagesStatics2 = syscall.GUID{0x1DF0DE4F, 0x072B, 0x4D7B,
	[8]byte{0x8F, 0x06, 0xCB, 0x2D, 0xB4, 0x0F, 0x2B, 0xB5}}

type IApplicationLanguagesStatics2Interface interface {
	win32.IInspectableInterface
	GetLanguagesForUser(user *IUser) *IVectorView[string]
}

type IApplicationLanguagesStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetLanguagesForUser uintptr
}

type IApplicationLanguagesStatics2 struct {
	win32.IInspectable
}

func (this *IApplicationLanguagesStatics2) Vtbl() *IApplicationLanguagesStatics2Vtbl {
	return (*IApplicationLanguagesStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationLanguagesStatics2) GetLanguagesForUser(user *IUser) *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetLanguagesForUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CA30221D-86D9-40FB-A26B-D44EB7CF08EA
var IID_ICalendar = syscall.GUID{0xCA30221D, 0x86D9, 0x40FB,
	[8]byte{0xA2, 0x6B, 0xD4, 0x4E, 0xB7, 0xCF, 0x08, 0xEA}}

type ICalendarInterface interface {
	win32.IInspectableInterface
	Clone() *ICalendar
	SetToMin()
	SetToMax()
	Get_Languages() *IVectorView[string]
	Get_NumeralSystem() string
	Put_NumeralSystem(value string)
	GetCalendarSystem() string
	ChangeCalendarSystem(value string)
	GetClock() string
	ChangeClock(value string)
	GetDateTime() DateTime
	SetDateTime(value DateTime)
	SetToNow()
	Get_FirstEra() int32
	Get_LastEra() int32
	Get_NumberOfEras() int32
	Get_Era() int32
	Put_Era(value int32)
	AddEras(eras int32)
	EraAsFullString() string
	EraAsString(idealLength int32) string
	Get_FirstYearInThisEra() int32
	Get_LastYearInThisEra() int32
	Get_NumberOfYearsInThisEra() int32
	Get_Year() int32
	Put_Year(value int32)
	AddYears(years int32)
	YearAsString() string
	YearAsTruncatedString(remainingDigits int32) string
	YearAsPaddedString(minDigits int32) string
	Get_FirstMonthInThisYear() int32
	Get_LastMonthInThisYear() int32
	Get_NumberOfMonthsInThisYear() int32
	Get_Month() int32
	Put_Month(value int32)
	AddMonths(months int32)
	MonthAsFullString() string
	MonthAsString(idealLength int32) string
	MonthAsFullSoloString() string
	MonthAsSoloString(idealLength int32) string
	MonthAsNumericString() string
	MonthAsPaddedNumericString(minDigits int32) string
	AddWeeks(weeks int32)
	Get_FirstDayInThisMonth() int32
	Get_LastDayInThisMonth() int32
	Get_NumberOfDaysInThisMonth() int32
	Get_Day() int32
	Put_Day(value int32)
	AddDays(days int32)
	DayAsString() string
	DayAsPaddedString(minDigits int32) string
	Get_DayOfWeek() DayOfWeek
	DayOfWeekAsFullString() string
	DayOfWeekAsString(idealLength int32) string
	DayOfWeekAsFullSoloString() string
	DayOfWeekAsSoloString(idealLength int32) string
	Get_FirstPeriodInThisDay() int32
	Get_LastPeriodInThisDay() int32
	Get_NumberOfPeriodsInThisDay() int32
	Get_Period() int32
	Put_Period(value int32)
	AddPeriods(periods int32)
	PeriodAsFullString() string
	PeriodAsString(idealLength int32) string
	Get_FirstHourInThisPeriod() int32
	Get_LastHourInThisPeriod() int32
	Get_NumberOfHoursInThisPeriod() int32
	Get_Hour() int32
	Put_Hour(value int32)
	AddHours(hours int32)
	HourAsString() string
	HourAsPaddedString(minDigits int32) string
	Get_Minute() int32
	Put_Minute(value int32)
	AddMinutes(minutes int32)
	MinuteAsString() string
	MinuteAsPaddedString(minDigits int32) string
	Get_Second() int32
	Put_Second(value int32)
	AddSeconds(seconds int32)
	SecondAsString() string
	SecondAsPaddedString(minDigits int32) string
	Get_Nanosecond() int32
	Put_Nanosecond(value int32)
	AddNanoseconds(nanoseconds int32)
	NanosecondAsString() string
	NanosecondAsPaddedString(minDigits int32) string
	Compare(other *ICalendar) int32
	CompareDateTime(other DateTime) int32
	CopyTo(other *ICalendar)
	Get_FirstMinuteInThisHour() int32
	Get_LastMinuteInThisHour() int32
	Get_NumberOfMinutesInThisHour() int32
	Get_FirstSecondInThisMinute() int32
	Get_LastSecondInThisMinute() int32
	Get_NumberOfSecondsInThisMinute() int32
	Get_ResolvedLanguage() string
	Get_IsDaylightSavingTime() bool
}

type ICalendarVtbl struct {
	win32.IInspectableVtbl
	Clone                           uintptr
	SetToMin                        uintptr
	SetToMax                        uintptr
	Get_Languages                   uintptr
	Get_NumeralSystem               uintptr
	Put_NumeralSystem               uintptr
	GetCalendarSystem               uintptr
	ChangeCalendarSystem            uintptr
	GetClock                        uintptr
	ChangeClock                     uintptr
	GetDateTime                     uintptr
	SetDateTime                     uintptr
	SetToNow                        uintptr
	Get_FirstEra                    uintptr
	Get_LastEra                     uintptr
	Get_NumberOfEras                uintptr
	Get_Era                         uintptr
	Put_Era                         uintptr
	AddEras                         uintptr
	EraAsFullString                 uintptr
	EraAsString                     uintptr
	Get_FirstYearInThisEra          uintptr
	Get_LastYearInThisEra           uintptr
	Get_NumberOfYearsInThisEra      uintptr
	Get_Year                        uintptr
	Put_Year                        uintptr
	AddYears                        uintptr
	YearAsString                    uintptr
	YearAsTruncatedString           uintptr
	YearAsPaddedString              uintptr
	Get_FirstMonthInThisYear        uintptr
	Get_LastMonthInThisYear         uintptr
	Get_NumberOfMonthsInThisYear    uintptr
	Get_Month                       uintptr
	Put_Month                       uintptr
	AddMonths                       uintptr
	MonthAsFullString               uintptr
	MonthAsString                   uintptr
	MonthAsFullSoloString           uintptr
	MonthAsSoloString               uintptr
	MonthAsNumericString            uintptr
	MonthAsPaddedNumericString      uintptr
	AddWeeks                        uintptr
	Get_FirstDayInThisMonth         uintptr
	Get_LastDayInThisMonth          uintptr
	Get_NumberOfDaysInThisMonth     uintptr
	Get_Day                         uintptr
	Put_Day                         uintptr
	AddDays                         uintptr
	DayAsString                     uintptr
	DayAsPaddedString               uintptr
	Get_DayOfWeek                   uintptr
	DayOfWeekAsFullString           uintptr
	DayOfWeekAsString               uintptr
	DayOfWeekAsFullSoloString       uintptr
	DayOfWeekAsSoloString           uintptr
	Get_FirstPeriodInThisDay        uintptr
	Get_LastPeriodInThisDay         uintptr
	Get_NumberOfPeriodsInThisDay    uintptr
	Get_Period                      uintptr
	Put_Period                      uintptr
	AddPeriods                      uintptr
	PeriodAsFullString              uintptr
	PeriodAsString                  uintptr
	Get_FirstHourInThisPeriod       uintptr
	Get_LastHourInThisPeriod        uintptr
	Get_NumberOfHoursInThisPeriod   uintptr
	Get_Hour                        uintptr
	Put_Hour                        uintptr
	AddHours                        uintptr
	HourAsString                    uintptr
	HourAsPaddedString              uintptr
	Get_Minute                      uintptr
	Put_Minute                      uintptr
	AddMinutes                      uintptr
	MinuteAsString                  uintptr
	MinuteAsPaddedString            uintptr
	Get_Second                      uintptr
	Put_Second                      uintptr
	AddSeconds                      uintptr
	SecondAsString                  uintptr
	SecondAsPaddedString            uintptr
	Get_Nanosecond                  uintptr
	Put_Nanosecond                  uintptr
	AddNanoseconds                  uintptr
	NanosecondAsString              uintptr
	NanosecondAsPaddedString        uintptr
	Compare                         uintptr
	CompareDateTime                 uintptr
	CopyTo                          uintptr
	Get_FirstMinuteInThisHour       uintptr
	Get_LastMinuteInThisHour        uintptr
	Get_NumberOfMinutesInThisHour   uintptr
	Get_FirstSecondInThisMinute     uintptr
	Get_LastSecondInThisMinute      uintptr
	Get_NumberOfSecondsInThisMinute uintptr
	Get_ResolvedLanguage            uintptr
	Get_IsDaylightSavingTime        uintptr
}

type ICalendar struct {
	win32.IInspectable
}

func (this *ICalendar) Vtbl() *ICalendarVtbl {
	return (*ICalendarVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICalendar) Clone() *ICalendar {
	var _result *ICalendar
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICalendar) SetToMin() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetToMin, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ICalendar) SetToMax() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetToMax, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ICalendar) Get_Languages() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Languages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICalendar) Get_NumeralSystem() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NumeralSystem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) Put_NumeralSystem(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NumeralSystem, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICalendar) GetCalendarSystem() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCalendarSystem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) ChangeCalendarSystem(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ChangeCalendarSystem, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICalendar) GetClock() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetClock, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) ChangeClock(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ChangeClock, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICalendar) GetDateTime() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDateTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) SetDateTime(value DateTime) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetDateTime, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ICalendar) SetToNow() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetToNow, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ICalendar) Get_FirstEra() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FirstEra, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Get_LastEra() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LastEra, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Get_NumberOfEras() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NumberOfEras, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Get_Era() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Era, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Put_Era(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Era, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICalendar) AddEras(eras int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddEras, uintptr(unsafe.Pointer(this)), uintptr(eras))
	_ = _hr
}

func (this *ICalendar) EraAsFullString() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EraAsFullString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) EraAsString(idealLength int32) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EraAsString, uintptr(unsafe.Pointer(this)), uintptr(idealLength), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) Get_FirstYearInThisEra() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FirstYearInThisEra, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Get_LastYearInThisEra() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LastYearInThisEra, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Get_NumberOfYearsInThisEra() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NumberOfYearsInThisEra, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Get_Year() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Year, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Put_Year(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Year, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICalendar) AddYears(years int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddYears, uintptr(unsafe.Pointer(this)), uintptr(years))
	_ = _hr
}

func (this *ICalendar) YearAsString() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().YearAsString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) YearAsTruncatedString(remainingDigits int32) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().YearAsTruncatedString, uintptr(unsafe.Pointer(this)), uintptr(remainingDigits), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) YearAsPaddedString(minDigits int32) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().YearAsPaddedString, uintptr(unsafe.Pointer(this)), uintptr(minDigits), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) Get_FirstMonthInThisYear() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FirstMonthInThisYear, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Get_LastMonthInThisYear() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LastMonthInThisYear, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Get_NumberOfMonthsInThisYear() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NumberOfMonthsInThisYear, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Get_Month() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Month, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Put_Month(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Month, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICalendar) AddMonths(months int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddMonths, uintptr(unsafe.Pointer(this)), uintptr(months))
	_ = _hr
}

func (this *ICalendar) MonthAsFullString() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MonthAsFullString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) MonthAsString(idealLength int32) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MonthAsString, uintptr(unsafe.Pointer(this)), uintptr(idealLength), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) MonthAsFullSoloString() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MonthAsFullSoloString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) MonthAsSoloString(idealLength int32) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MonthAsSoloString, uintptr(unsafe.Pointer(this)), uintptr(idealLength), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) MonthAsNumericString() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MonthAsNumericString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) MonthAsPaddedNumericString(minDigits int32) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MonthAsPaddedNumericString, uintptr(unsafe.Pointer(this)), uintptr(minDigits), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) AddWeeks(weeks int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddWeeks, uintptr(unsafe.Pointer(this)), uintptr(weeks))
	_ = _hr
}

func (this *ICalendar) Get_FirstDayInThisMonth() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FirstDayInThisMonth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Get_LastDayInThisMonth() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LastDayInThisMonth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Get_NumberOfDaysInThisMonth() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NumberOfDaysInThisMonth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Get_Day() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Day, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Put_Day(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Day, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICalendar) AddDays(days int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddDays, uintptr(unsafe.Pointer(this)), uintptr(days))
	_ = _hr
}

func (this *ICalendar) DayAsString() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DayAsString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) DayAsPaddedString(minDigits int32) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DayAsPaddedString, uintptr(unsafe.Pointer(this)), uintptr(minDigits), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) Get_DayOfWeek() DayOfWeek {
	var _result DayOfWeek
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DayOfWeek, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) DayOfWeekAsFullString() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DayOfWeekAsFullString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) DayOfWeekAsString(idealLength int32) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DayOfWeekAsString, uintptr(unsafe.Pointer(this)), uintptr(idealLength), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) DayOfWeekAsFullSoloString() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DayOfWeekAsFullSoloString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) DayOfWeekAsSoloString(idealLength int32) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DayOfWeekAsSoloString, uintptr(unsafe.Pointer(this)), uintptr(idealLength), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) Get_FirstPeriodInThisDay() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FirstPeriodInThisDay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Get_LastPeriodInThisDay() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LastPeriodInThisDay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Get_NumberOfPeriodsInThisDay() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NumberOfPeriodsInThisDay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Get_Period() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Period, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Put_Period(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Period, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICalendar) AddPeriods(periods int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddPeriods, uintptr(unsafe.Pointer(this)), uintptr(periods))
	_ = _hr
}

func (this *ICalendar) PeriodAsFullString() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PeriodAsFullString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) PeriodAsString(idealLength int32) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PeriodAsString, uintptr(unsafe.Pointer(this)), uintptr(idealLength), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) Get_FirstHourInThisPeriod() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FirstHourInThisPeriod, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Get_LastHourInThisPeriod() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LastHourInThisPeriod, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Get_NumberOfHoursInThisPeriod() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NumberOfHoursInThisPeriod, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Get_Hour() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Hour, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Put_Hour(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Hour, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICalendar) AddHours(hours int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddHours, uintptr(unsafe.Pointer(this)), uintptr(hours))
	_ = _hr
}

func (this *ICalendar) HourAsString() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().HourAsString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) HourAsPaddedString(minDigits int32) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().HourAsPaddedString, uintptr(unsafe.Pointer(this)), uintptr(minDigits), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) Get_Minute() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Minute, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Put_Minute(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Minute, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICalendar) AddMinutes(minutes int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddMinutes, uintptr(unsafe.Pointer(this)), uintptr(minutes))
	_ = _hr
}

func (this *ICalendar) MinuteAsString() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MinuteAsString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) MinuteAsPaddedString(minDigits int32) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MinuteAsPaddedString, uintptr(unsafe.Pointer(this)), uintptr(minDigits), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) Get_Second() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Second, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Put_Second(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Second, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICalendar) AddSeconds(seconds int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddSeconds, uintptr(unsafe.Pointer(this)), uintptr(seconds))
	_ = _hr
}

func (this *ICalendar) SecondAsString() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SecondAsString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) SecondAsPaddedString(minDigits int32) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SecondAsPaddedString, uintptr(unsafe.Pointer(this)), uintptr(minDigits), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) Get_Nanosecond() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Nanosecond, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Put_Nanosecond(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Nanosecond, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICalendar) AddNanoseconds(nanoseconds int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddNanoseconds, uintptr(unsafe.Pointer(this)), uintptr(nanoseconds))
	_ = _hr
}

func (this *ICalendar) NanosecondAsString() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().NanosecondAsString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) NanosecondAsPaddedString(minDigits int32) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().NanosecondAsPaddedString, uintptr(unsafe.Pointer(this)), uintptr(minDigits), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) Compare(other *ICalendar) int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Compare, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(other)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) CompareDateTime(other DateTime) int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CompareDateTime, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&other)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) CopyTo(other *ICalendar) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CopyTo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(other)))
	_ = _hr
}

func (this *ICalendar) Get_FirstMinuteInThisHour() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FirstMinuteInThisHour, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Get_LastMinuteInThisHour() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LastMinuteInThisHour, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Get_NumberOfMinutesInThisHour() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NumberOfMinutesInThisHour, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Get_FirstSecondInThisMinute() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FirstSecondInThisMinute, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Get_LastSecondInThisMinute() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LastSecondInThisMinute, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Get_NumberOfSecondsInThisMinute() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NumberOfSecondsInThisMinute, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICalendar) Get_ResolvedLanguage() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResolvedLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendar) Get_IsDaylightSavingTime() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDaylightSavingTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 83F58412-E56B-4C75-A66E-0F63D57758A6
var IID_ICalendarFactory = syscall.GUID{0x83F58412, 0xE56B, 0x4C75,
	[8]byte{0xA6, 0x6E, 0x0F, 0x63, 0xD5, 0x77, 0x58, 0xA6}}

type ICalendarFactoryInterface interface {
	win32.IInspectableInterface
	CreateCalendarDefaultCalendarAndClock(languages *IIterable[string]) *ICalendar
	CreateCalendar(languages *IIterable[string], calendar string, clock string) *ICalendar
}

type ICalendarFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateCalendarDefaultCalendarAndClock uintptr
	CreateCalendar                        uintptr
}

type ICalendarFactory struct {
	win32.IInspectable
}

func (this *ICalendarFactory) Vtbl() *ICalendarFactoryVtbl {
	return (*ICalendarFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICalendarFactory) CreateCalendarDefaultCalendarAndClock(languages *IIterable[string]) *ICalendar {
	var _result *ICalendar
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCalendarDefaultCalendarAndClock, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(languages)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICalendarFactory) CreateCalendar(languages *IIterable[string], calendar string, clock string) *ICalendar {
	var _result *ICalendar
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCalendar, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(languages)), NewHStr(calendar).Ptr, NewHStr(clock).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B44B378C-CA7E-4590-9E72-EA2BEC1A5115
var IID_ICalendarFactory2 = syscall.GUID{0xB44B378C, 0xCA7E, 0x4590,
	[8]byte{0x9E, 0x72, 0xEA, 0x2B, 0xEC, 0x1A, 0x51, 0x15}}

type ICalendarFactory2Interface interface {
	win32.IInspectableInterface
	CreateCalendarWithTimeZone(languages *IIterable[string], calendar string, clock string, timeZoneId string) *ICalendar
}

type ICalendarFactory2Vtbl struct {
	win32.IInspectableVtbl
	CreateCalendarWithTimeZone uintptr
}

type ICalendarFactory2 struct {
	win32.IInspectable
}

func (this *ICalendarFactory2) Vtbl() *ICalendarFactory2Vtbl {
	return (*ICalendarFactory2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICalendarFactory2) CreateCalendarWithTimeZone(languages *IIterable[string], calendar string, clock string, timeZoneId string) *ICalendar {
	var _result *ICalendar
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCalendarWithTimeZone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(languages)), NewHStr(calendar).Ptr, NewHStr(clock).Ptr, NewHStr(timeZoneId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 80653F68-2CB2-4C1F-B590-F0F52BF4FD1A
var IID_ICalendarIdentifiersStatics = syscall.GUID{0x80653F68, 0x2CB2, 0x4C1F,
	[8]byte{0xB5, 0x90, 0xF0, 0xF5, 0x2B, 0xF4, 0xFD, 0x1A}}

type ICalendarIdentifiersStaticsInterface interface {
	win32.IInspectableInterface
	Get_Gregorian() string
	Get_Hebrew() string
	Get_Hijri() string
	Get_Japanese() string
	Get_Julian() string
	Get_Korean() string
	Get_Taiwan() string
	Get_Thai() string
	Get_UmAlQura() string
}

type ICalendarIdentifiersStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Gregorian uintptr
	Get_Hebrew    uintptr
	Get_Hijri     uintptr
	Get_Japanese  uintptr
	Get_Julian    uintptr
	Get_Korean    uintptr
	Get_Taiwan    uintptr
	Get_Thai      uintptr
	Get_UmAlQura  uintptr
}

type ICalendarIdentifiersStatics struct {
	win32.IInspectable
}

func (this *ICalendarIdentifiersStatics) Vtbl() *ICalendarIdentifiersStaticsVtbl {
	return (*ICalendarIdentifiersStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICalendarIdentifiersStatics) Get_Gregorian() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Gregorian, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendarIdentifiersStatics) Get_Hebrew() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Hebrew, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendarIdentifiersStatics) Get_Hijri() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Hijri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendarIdentifiersStatics) Get_Japanese() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Japanese, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendarIdentifiersStatics) Get_Julian() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Julian, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendarIdentifiersStatics) Get_Korean() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Korean, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendarIdentifiersStatics) Get_Taiwan() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Taiwan, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendarIdentifiersStatics) Get_Thai() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Thai, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendarIdentifiersStatics) Get_UmAlQura() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UmAlQura, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 7DF4D488-5FD0-42A7-95B5-7D98D823075F
var IID_ICalendarIdentifiersStatics2 = syscall.GUID{0x7DF4D488, 0x5FD0, 0x42A7,
	[8]byte{0x95, 0xB5, 0x7D, 0x98, 0xD8, 0x23, 0x07, 0x5F}}

type ICalendarIdentifiersStatics2Interface interface {
	win32.IInspectableInterface
	Get_Persian() string
}

type ICalendarIdentifiersStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_Persian uintptr
}

type ICalendarIdentifiersStatics2 struct {
	win32.IInspectable
}

func (this *ICalendarIdentifiersStatics2) Vtbl() *ICalendarIdentifiersStatics2Vtbl {
	return (*ICalendarIdentifiersStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICalendarIdentifiersStatics2) Get_Persian() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Persian, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 2C225423-1FAD-40C0-9334-A8EB90DB04F5
var IID_ICalendarIdentifiersStatics3 = syscall.GUID{0x2C225423, 0x1FAD, 0x40C0,
	[8]byte{0x93, 0x34, 0xA8, 0xEB, 0x90, 0xDB, 0x04, 0xF5}}

type ICalendarIdentifiersStatics3Interface interface {
	win32.IInspectableInterface
	Get_ChineseLunar() string
	Get_JapaneseLunar() string
	Get_KoreanLunar() string
	Get_TaiwanLunar() string
	Get_VietnameseLunar() string
}

type ICalendarIdentifiersStatics3Vtbl struct {
	win32.IInspectableVtbl
	Get_ChineseLunar    uintptr
	Get_JapaneseLunar   uintptr
	Get_KoreanLunar     uintptr
	Get_TaiwanLunar     uintptr
	Get_VietnameseLunar uintptr
}

type ICalendarIdentifiersStatics3 struct {
	win32.IInspectable
}

func (this *ICalendarIdentifiersStatics3) Vtbl() *ICalendarIdentifiersStatics3Vtbl {
	return (*ICalendarIdentifiersStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICalendarIdentifiersStatics3) Get_ChineseLunar() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ChineseLunar, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendarIdentifiersStatics3) Get_JapaneseLunar() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_JapaneseLunar, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendarIdentifiersStatics3) Get_KoreanLunar() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KoreanLunar, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendarIdentifiersStatics3) Get_TaiwanLunar() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TaiwanLunar, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICalendarIdentifiersStatics3) Get_VietnameseLunar() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VietnameseLunar, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 523805BB-12EC-4F83-BC31-B1B4376B0808
var IID_IClockIdentifiersStatics = syscall.GUID{0x523805BB, 0x12EC, 0x4F83,
	[8]byte{0xBC, 0x31, 0xB1, 0xB4, 0x37, 0x6B, 0x08, 0x08}}

type IClockIdentifiersStaticsInterface interface {
	win32.IInspectableInterface
	Get_TwelveHour() string
	Get_TwentyFourHour() string
}

type IClockIdentifiersStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_TwelveHour     uintptr
	Get_TwentyFourHour uintptr
}

type IClockIdentifiersStatics struct {
	win32.IInspectable
}

func (this *IClockIdentifiersStatics) Vtbl() *IClockIdentifiersStaticsVtbl {
	return (*IClockIdentifiersStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClockIdentifiersStatics) Get_TwelveHour() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TwelveHour, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IClockIdentifiersStatics) Get_TwentyFourHour() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TwentyFourHour, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 74B49942-EB75-443A-95B3-7D723F56F93C
var IID_ICurrencyAmount = syscall.GUID{0x74B49942, 0xEB75, 0x443A,
	[8]byte{0x95, 0xB3, 0x7D, 0x72, 0x3F, 0x56, 0xF9, 0x3C}}

type ICurrencyAmountInterface interface {
	win32.IInspectableInterface
	Get_Amount() string
	Get_Currency() string
}

type ICurrencyAmountVtbl struct {
	win32.IInspectableVtbl
	Get_Amount   uintptr
	Get_Currency uintptr
}

type ICurrencyAmount struct {
	win32.IInspectable
}

func (this *ICurrencyAmount) Vtbl() *ICurrencyAmountVtbl {
	return (*ICurrencyAmountVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICurrencyAmount) Get_Amount() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Amount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyAmount) Get_Currency() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Currency, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 48D7168F-EF3B-4AEE-A6A1-4B036FE03FF0
var IID_ICurrencyAmountFactory = syscall.GUID{0x48D7168F, 0xEF3B, 0x4AEE,
	[8]byte{0xA6, 0xA1, 0x4B, 0x03, 0x6F, 0xE0, 0x3F, 0xF0}}

type ICurrencyAmountFactoryInterface interface {
	win32.IInspectableInterface
	Create(amount string, currency string) *ICurrencyAmount
}

type ICurrencyAmountFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ICurrencyAmountFactory struct {
	win32.IInspectable
}

func (this *ICurrencyAmountFactory) Vtbl() *ICurrencyAmountFactoryVtbl {
	return (*ICurrencyAmountFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICurrencyAmountFactory) Create(amount string, currency string) *ICurrencyAmount {
	var _result *ICurrencyAmount
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(amount).Ptr, NewHStr(currency).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9F1D091B-D586-4913-9B6A-A9BD2DC12874
var IID_ICurrencyIdentifiersStatics = syscall.GUID{0x9F1D091B, 0xD586, 0x4913,
	[8]byte{0x9B, 0x6A, 0xA9, 0xBD, 0x2D, 0xC1, 0x28, 0x74}}

type ICurrencyIdentifiersStaticsInterface interface {
	win32.IInspectableInterface
	Get_AED() string
	Get_AFN() string
	Get_ALL() string
	Get_AMD() string
	Get_ANG() string
	Get_AOA() string
	Get_ARS() string
	Get_AUD() string
	Get_AWG() string
	Get_AZN() string
	Get_BAM() string
	Get_BBD() string
	Get_BDT() string
	Get_BGN() string
	Get_BHD() string
	Get_BIF() string
	Get_BMD() string
	Get_BND() string
	Get_BOB() string
	Get_BRL() string
	Get_BSD() string
	Get_BTN() string
	Get_BWP() string
	Get_BYR() string
	Get_BZD() string
	Get_CAD() string
	Get_CDF() string
	Get_CHF() string
	Get_CLP() string
	Get_CNY() string
	Get_COP() string
	Get_CRC() string
	Get_CUP() string
	Get_CVE() string
	Get_CZK() string
	Get_DJF() string
	Get_DKK() string
	Get_DOP() string
	Get_DZD() string
	Get_EGP() string
	Get_ERN() string
	Get_ETB() string
	Get_EUR() string
	Get_FJD() string
	Get_FKP() string
	Get_GBP() string
	Get_GEL() string
	Get_GHS() string
	Get_GIP() string
	Get_GMD() string
	Get_GNF() string
	Get_GTQ() string
	Get_GYD() string
	Get_HKD() string
	Get_HNL() string
	Get_HRK() string
	Get_HTG() string
	Get_HUF() string
	Get_IDR() string
	Get_ILS() string
	Get_INR() string
	Get_IQD() string
	Get_IRR() string
	Get_ISK() string
	Get_JMD() string
	Get_JOD() string
	Get_JPY() string
	Get_KES() string
	Get_KGS() string
	Get_KHR() string
	Get_KMF() string
	Get_KPW() string
	Get_KRW() string
	Get_KWD() string
	Get_KYD() string
	Get_KZT() string
	Get_LAK() string
	Get_LBP() string
	Get_LKR() string
	Get_LRD() string
	Get_LSL() string
	Get_LTL() string
	Get_LVL() string
	Get_LYD() string
	Get_MAD() string
	Get_MDL() string
	Get_MGA() string
	Get_MKD() string
	Get_MMK() string
	Get_MNT() string
	Get_MOP() string
	Get_MRO() string
	Get_MUR() string
	Get_MVR() string
	Get_MWK() string
	Get_MXN() string
	Get_MYR() string
	Get_MZN() string
	Get_NAD() string
	Get_NGN() string
	Get_NIO() string
	Get_NOK() string
	Get_NPR() string
	Get_NZD() string
	Get_OMR() string
	Get_PAB() string
	Get_PEN() string
	Get_PGK() string
	Get_PHP() string
	Get_PKR() string
	Get_PLN() string
	Get_PYG() string
	Get_QAR() string
	Get_RON() string
	Get_RSD() string
	Get_RUB() string
	Get_RWF() string
	Get_SAR() string
	Get_SBD() string
	Get_SCR() string
	Get_SDG() string
	Get_SEK() string
	Get_SGD() string
	Get_SHP() string
	Get_SLL() string
	Get_SOS() string
	Get_SRD() string
	Get_STD() string
	Get_SYP() string
	Get_SZL() string
	Get_THB() string
	Get_TJS() string
	Get_TMT() string
	Get_TND() string
	Get_TOP() string
	Get_TRY() string
	Get_TTD() string
	Get_TWD() string
	Get_TZS() string
	Get_UAH() string
	Get_UGX() string
	Get_USD() string
	Get_UYU() string
	Get_UZS() string
	Get_VEF() string
	Get_VND() string
	Get_VUV() string
	Get_WST() string
	Get_XAF() string
	Get_XCD() string
	Get_XOF() string
	Get_XPF() string
	Get_XXX() string
	Get_YER() string
	Get_ZAR() string
	Get_ZMW() string
	Get_ZWL() string
}

type ICurrencyIdentifiersStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_AED uintptr
	Get_AFN uintptr
	Get_ALL uintptr
	Get_AMD uintptr
	Get_ANG uintptr
	Get_AOA uintptr
	Get_ARS uintptr
	Get_AUD uintptr
	Get_AWG uintptr
	Get_AZN uintptr
	Get_BAM uintptr
	Get_BBD uintptr
	Get_BDT uintptr
	Get_BGN uintptr
	Get_BHD uintptr
	Get_BIF uintptr
	Get_BMD uintptr
	Get_BND uintptr
	Get_BOB uintptr
	Get_BRL uintptr
	Get_BSD uintptr
	Get_BTN uintptr
	Get_BWP uintptr
	Get_BYR uintptr
	Get_BZD uintptr
	Get_CAD uintptr
	Get_CDF uintptr
	Get_CHF uintptr
	Get_CLP uintptr
	Get_CNY uintptr
	Get_COP uintptr
	Get_CRC uintptr
	Get_CUP uintptr
	Get_CVE uintptr
	Get_CZK uintptr
	Get_DJF uintptr
	Get_DKK uintptr
	Get_DOP uintptr
	Get_DZD uintptr
	Get_EGP uintptr
	Get_ERN uintptr
	Get_ETB uintptr
	Get_EUR uintptr
	Get_FJD uintptr
	Get_FKP uintptr
	Get_GBP uintptr
	Get_GEL uintptr
	Get_GHS uintptr
	Get_GIP uintptr
	Get_GMD uintptr
	Get_GNF uintptr
	Get_GTQ uintptr
	Get_GYD uintptr
	Get_HKD uintptr
	Get_HNL uintptr
	Get_HRK uintptr
	Get_HTG uintptr
	Get_HUF uintptr
	Get_IDR uintptr
	Get_ILS uintptr
	Get_INR uintptr
	Get_IQD uintptr
	Get_IRR uintptr
	Get_ISK uintptr
	Get_JMD uintptr
	Get_JOD uintptr
	Get_JPY uintptr
	Get_KES uintptr
	Get_KGS uintptr
	Get_KHR uintptr
	Get_KMF uintptr
	Get_KPW uintptr
	Get_KRW uintptr
	Get_KWD uintptr
	Get_KYD uintptr
	Get_KZT uintptr
	Get_LAK uintptr
	Get_LBP uintptr
	Get_LKR uintptr
	Get_LRD uintptr
	Get_LSL uintptr
	Get_LTL uintptr
	Get_LVL uintptr
	Get_LYD uintptr
	Get_MAD uintptr
	Get_MDL uintptr
	Get_MGA uintptr
	Get_MKD uintptr
	Get_MMK uintptr
	Get_MNT uintptr
	Get_MOP uintptr
	Get_MRO uintptr
	Get_MUR uintptr
	Get_MVR uintptr
	Get_MWK uintptr
	Get_MXN uintptr
	Get_MYR uintptr
	Get_MZN uintptr
	Get_NAD uintptr
	Get_NGN uintptr
	Get_NIO uintptr
	Get_NOK uintptr
	Get_NPR uintptr
	Get_NZD uintptr
	Get_OMR uintptr
	Get_PAB uintptr
	Get_PEN uintptr
	Get_PGK uintptr
	Get_PHP uintptr
	Get_PKR uintptr
	Get_PLN uintptr
	Get_PYG uintptr
	Get_QAR uintptr
	Get_RON uintptr
	Get_RSD uintptr
	Get_RUB uintptr
	Get_RWF uintptr
	Get_SAR uintptr
	Get_SBD uintptr
	Get_SCR uintptr
	Get_SDG uintptr
	Get_SEK uintptr
	Get_SGD uintptr
	Get_SHP uintptr
	Get_SLL uintptr
	Get_SOS uintptr
	Get_SRD uintptr
	Get_STD uintptr
	Get_SYP uintptr
	Get_SZL uintptr
	Get_THB uintptr
	Get_TJS uintptr
	Get_TMT uintptr
	Get_TND uintptr
	Get_TOP uintptr
	Get_TRY uintptr
	Get_TTD uintptr
	Get_TWD uintptr
	Get_TZS uintptr
	Get_UAH uintptr
	Get_UGX uintptr
	Get_USD uintptr
	Get_UYU uintptr
	Get_UZS uintptr
	Get_VEF uintptr
	Get_VND uintptr
	Get_VUV uintptr
	Get_WST uintptr
	Get_XAF uintptr
	Get_XCD uintptr
	Get_XOF uintptr
	Get_XPF uintptr
	Get_XXX uintptr
	Get_YER uintptr
	Get_ZAR uintptr
	Get_ZMW uintptr
	Get_ZWL uintptr
}

type ICurrencyIdentifiersStatics struct {
	win32.IInspectable
}

func (this *ICurrencyIdentifiersStatics) Vtbl() *ICurrencyIdentifiersStaticsVtbl {
	return (*ICurrencyIdentifiersStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICurrencyIdentifiersStatics) Get_AED() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AED, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_AFN() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AFN, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_ALL() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ALL, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_AMD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AMD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_ANG() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ANG, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_AOA() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AOA, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_ARS() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ARS, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_AUD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AUD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_AWG() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AWG, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_AZN() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AZN, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_BAM() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BAM, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_BBD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BBD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_BDT() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BDT, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_BGN() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BGN, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_BHD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BHD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_BIF() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BIF, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_BMD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BMD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_BND() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BND, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_BOB() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BOB, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_BRL() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BRL, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_BSD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BSD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_BTN() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BTN, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_BWP() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BWP, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_BYR() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BYR, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_BZD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BZD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_CAD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CAD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_CDF() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CDF, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_CHF() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CHF, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_CLP() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CLP, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_CNY() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CNY, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_COP() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_COP, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_CRC() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CRC, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_CUP() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CUP, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_CVE() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CVE, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_CZK() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CZK, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_DJF() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DJF, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_DKK() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DKK, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_DOP() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DOP, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_DZD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DZD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_EGP() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EGP, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_ERN() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ERN, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_ETB() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ETB, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_EUR() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EUR, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_FJD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FJD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_FKP() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FKP, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_GBP() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GBP, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_GEL() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GEL, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_GHS() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GHS, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_GIP() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GIP, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_GMD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GMD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_GNF() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GNF, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_GTQ() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GTQ, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_GYD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GYD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_HKD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HKD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_HNL() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HNL, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_HRK() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HRK, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_HTG() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HTG, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_HUF() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HUF, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_IDR() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IDR, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_ILS() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ILS, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_INR() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_INR, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_IQD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IQD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_IRR() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IRR, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_ISK() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ISK, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_JMD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_JMD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_JOD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_JOD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_JPY() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_JPY, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_KES() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KES, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_KGS() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KGS, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_KHR() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KHR, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_KMF() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KMF, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_KPW() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KPW, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_KRW() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KRW, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_KWD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KWD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_KYD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KYD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_KZT() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KZT, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_LAK() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LAK, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_LBP() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LBP, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_LKR() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LKR, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_LRD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LRD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_LSL() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LSL, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_LTL() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LTL, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_LVL() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LVL, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_LYD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LYD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_MAD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MAD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_MDL() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MDL, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_MGA() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MGA, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_MKD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MKD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_MMK() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MMK, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_MNT() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MNT, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_MOP() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MOP, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_MRO() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MRO, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_MUR() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MUR, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_MVR() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MVR, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_MWK() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MWK, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_MXN() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MXN, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_MYR() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MYR, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_MZN() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MZN, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_NAD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NAD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_NGN() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NGN, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_NIO() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NIO, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_NOK() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NOK, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_NPR() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NPR, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_NZD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NZD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_OMR() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OMR, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_PAB() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PAB, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_PEN() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PEN, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_PGK() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PGK, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_PHP() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PHP, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_PKR() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PKR, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_PLN() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PLN, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_PYG() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PYG, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_QAR() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_QAR, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_RON() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RON, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_RSD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RSD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_RUB() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RUB, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_RWF() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RWF, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_SAR() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SAR, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_SBD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SBD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_SCR() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SCR, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_SDG() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SDG, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_SEK() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SEK, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_SGD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SGD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_SHP() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SHP, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_SLL() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SLL, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_SOS() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SOS, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_SRD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SRD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_STD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_STD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_SYP() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SYP, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_SZL() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SZL, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_THB() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_THB, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_TJS() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TJS, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_TMT() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TMT, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_TND() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TND, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_TOP() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TOP, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_TRY() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TRY, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_TTD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TTD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_TWD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TWD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_TZS() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TZS, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_UAH() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UAH, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_UGX() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UGX, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_USD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_USD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_UYU() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UYU, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_UZS() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UZS, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_VEF() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VEF, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_VND() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VND, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_VUV() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VUV, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_WST() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WST, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_XAF() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_XAF, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_XCD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_XCD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_XOF() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_XOF, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_XPF() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_XPF, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_XXX() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_XXX, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_YER() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_YER, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_ZAR() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ZAR, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_ZMW() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ZMW, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics) Get_ZWL() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ZWL, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 1814797F-C3B2-4C33-9591-980011950D37
var IID_ICurrencyIdentifiersStatics2 = syscall.GUID{0x1814797F, 0xC3B2, 0x4C33,
	[8]byte{0x95, 0x91, 0x98, 0x00, 0x11, 0x95, 0x0D, 0x37}}

type ICurrencyIdentifiersStatics2Interface interface {
	win32.IInspectableInterface
	Get_BYN() string
}

type ICurrencyIdentifiersStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_BYN uintptr
}

type ICurrencyIdentifiersStatics2 struct {
	win32.IInspectable
}

func (this *ICurrencyIdentifiersStatics2) Vtbl() *ICurrencyIdentifiersStatics2Vtbl {
	return (*ICurrencyIdentifiersStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICurrencyIdentifiersStatics2) Get_BYN() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BYN, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 4FB23BFA-ED25-4F4D-857F-237F1748C21C
var IID_ICurrencyIdentifiersStatics3 = syscall.GUID{0x4FB23BFA, 0xED25, 0x4F4D,
	[8]byte{0x85, 0x7F, 0x23, 0x7F, 0x17, 0x48, 0xC2, 0x1C}}

type ICurrencyIdentifiersStatics3Interface interface {
	win32.IInspectableInterface
	Get_MRU() string
	Get_SSP() string
	Get_STN() string
	Get_VES() string
}

type ICurrencyIdentifiersStatics3Vtbl struct {
	win32.IInspectableVtbl
	Get_MRU uintptr
	Get_SSP uintptr
	Get_STN uintptr
	Get_VES uintptr
}

type ICurrencyIdentifiersStatics3 struct {
	win32.IInspectable
}

func (this *ICurrencyIdentifiersStatics3) Vtbl() *ICurrencyIdentifiersStatics3Vtbl {
	return (*ICurrencyIdentifiersStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICurrencyIdentifiersStatics3) Get_MRU() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MRU, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics3) Get_SSP() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SSP, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics3) Get_STN() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_STN, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICurrencyIdentifiersStatics3) Get_VES() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VES, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 01E9A621-4A64-4ED9-954F-9EDEB07BD903
var IID_IGeographicRegion = syscall.GUID{0x01E9A621, 0x4A64, 0x4ED9,
	[8]byte{0x95, 0x4F, 0x9E, 0xDE, 0xB0, 0x7B, 0xD9, 0x03}}

type IGeographicRegionInterface interface {
	win32.IInspectableInterface
	Get_Code() string
	Get_CodeTwoLetter() string
	Get_CodeThreeLetter() string
	Get_CodeThreeDigit() string
	Get_DisplayName() string
	Get_NativeName() string
	Get_CurrenciesInUse() *IVectorView[string]
}

type IGeographicRegionVtbl struct {
	win32.IInspectableVtbl
	Get_Code            uintptr
	Get_CodeTwoLetter   uintptr
	Get_CodeThreeLetter uintptr
	Get_CodeThreeDigit  uintptr
	Get_DisplayName     uintptr
	Get_NativeName      uintptr
	Get_CurrenciesInUse uintptr
}

type IGeographicRegion struct {
	win32.IInspectable
}

func (this *IGeographicRegion) Vtbl() *IGeographicRegionVtbl {
	return (*IGeographicRegionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeographicRegion) Get_Code() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Code, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGeographicRegion) Get_CodeTwoLetter() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CodeTwoLetter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGeographicRegion) Get_CodeThreeLetter() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CodeThreeLetter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGeographicRegion) Get_CodeThreeDigit() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CodeThreeDigit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGeographicRegion) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGeographicRegion) Get_NativeName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NativeName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGeographicRegion) Get_CurrenciesInUse() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrenciesInUse, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 53425270-77B4-426B-859F-81E19D512546
var IID_IGeographicRegionFactory = syscall.GUID{0x53425270, 0x77B4, 0x426B,
	[8]byte{0x85, 0x9F, 0x81, 0xE1, 0x9D, 0x51, 0x25, 0x46}}

type IGeographicRegionFactoryInterface interface {
	win32.IInspectableInterface
	CreateGeographicRegion(geographicRegionCode string) *IGeographicRegion
}

type IGeographicRegionFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateGeographicRegion uintptr
}

type IGeographicRegionFactory struct {
	win32.IInspectable
}

func (this *IGeographicRegionFactory) Vtbl() *IGeographicRegionFactoryVtbl {
	return (*IGeographicRegionFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeographicRegionFactory) CreateGeographicRegion(geographicRegionCode string) *IGeographicRegion {
	var _result *IGeographicRegion
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateGeographicRegion, uintptr(unsafe.Pointer(this)), NewHStr(geographicRegionCode).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 29E28974-7AD9-4EF4-8799-B3B44FADEC08
var IID_IGeographicRegionStatics = syscall.GUID{0x29E28974, 0x7AD9, 0x4EF4,
	[8]byte{0x87, 0x99, 0xB3, 0xB4, 0x4F, 0xAD, 0xEC, 0x08}}

type IGeographicRegionStaticsInterface interface {
	win32.IInspectableInterface
	IsSupported(geographicRegionCode string) bool
}

type IGeographicRegionStaticsVtbl struct {
	win32.IInspectableVtbl
	IsSupported uintptr
}

type IGeographicRegionStatics struct {
	win32.IInspectable
}

func (this *IGeographicRegionStatics) Vtbl() *IGeographicRegionStaticsVtbl {
	return (*IGeographicRegionStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeographicRegionStatics) IsSupported(geographicRegionCode string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsSupported, uintptr(unsafe.Pointer(this)), NewHStr(geographicRegionCode).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 2F6A9300-E85B-43E6-897D-5D82F862DF21
var IID_IJapanesePhoneme = syscall.GUID{0x2F6A9300, 0xE85B, 0x43E6,
	[8]byte{0x89, 0x7D, 0x5D, 0x82, 0xF8, 0x62, 0xDF, 0x21}}

type IJapanesePhonemeInterface interface {
	win32.IInspectableInterface
	Get_DisplayText() string
	Get_YomiText() string
	Get_IsPhraseStart() bool
}

type IJapanesePhonemeVtbl struct {
	win32.IInspectableVtbl
	Get_DisplayText   uintptr
	Get_YomiText      uintptr
	Get_IsPhraseStart uintptr
}

type IJapanesePhoneme struct {
	win32.IInspectable
}

func (this *IJapanesePhoneme) Vtbl() *IJapanesePhonemeVtbl {
	return (*IJapanesePhonemeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IJapanesePhoneme) Get_DisplayText() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayText, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IJapanesePhoneme) Get_YomiText() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_YomiText, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IJapanesePhoneme) Get_IsPhraseStart() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPhraseStart, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 88AB9E90-93DE-41B2-B4D5-8EDB227FD1C2
var IID_IJapanesePhoneticAnalyzerStatics = syscall.GUID{0x88AB9E90, 0x93DE, 0x41B2,
	[8]byte{0xB4, 0xD5, 0x8E, 0xDB, 0x22, 0x7F, 0xD1, 0xC2}}

type IJapanesePhoneticAnalyzerStaticsInterface interface {
	win32.IInspectableInterface
	GetWords(input string) *IVectorView[*IJapanesePhoneme]
	GetWordsWithMonoRubyOption(input string, monoRuby bool) *IVectorView[*IJapanesePhoneme]
}

type IJapanesePhoneticAnalyzerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetWords                   uintptr
	GetWordsWithMonoRubyOption uintptr
}

type IJapanesePhoneticAnalyzerStatics struct {
	win32.IInspectable
}

func (this *IJapanesePhoneticAnalyzerStatics) Vtbl() *IJapanesePhoneticAnalyzerStaticsVtbl {
	return (*IJapanesePhoneticAnalyzerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IJapanesePhoneticAnalyzerStatics) GetWords(input string) *IVectorView[*IJapanesePhoneme] {
	var _result *IVectorView[*IJapanesePhoneme]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetWords, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IJapanesePhoneticAnalyzerStatics) GetWordsWithMonoRubyOption(input string, monoRuby bool) *IVectorView[*IJapanesePhoneme] {
	var _result *IVectorView[*IJapanesePhoneme]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetWordsWithMonoRubyOption, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(*(*byte)(unsafe.Pointer(&monoRuby))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EA79A752-F7C2-4265-B1BD-C4DEC4E4F080
var IID_ILanguage = syscall.GUID{0xEA79A752, 0xF7C2, 0x4265,
	[8]byte{0xB1, 0xBD, 0xC4, 0xDE, 0xC4, 0xE4, 0xF0, 0x80}}

type ILanguageInterface interface {
	win32.IInspectableInterface
	Get_LanguageTag() string
	Get_DisplayName() string
	Get_NativeName() string
	Get_Script() string
}

type ILanguageVtbl struct {
	win32.IInspectableVtbl
	Get_LanguageTag uintptr
	Get_DisplayName uintptr
	Get_NativeName  uintptr
	Get_Script      uintptr
}

type ILanguage struct {
	win32.IInspectable
}

func (this *ILanguage) Vtbl() *ILanguageVtbl {
	return (*ILanguageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILanguage) Get_LanguageTag() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LanguageTag, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILanguage) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILanguage) Get_NativeName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NativeName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILanguage) Get_Script() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Script, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 6A47E5B5-D94D-4886-A404-A5A5B9D5B494
var IID_ILanguage2 = syscall.GUID{0x6A47E5B5, 0xD94D, 0x4886,
	[8]byte{0xA4, 0x04, 0xA5, 0xA5, 0xB9, 0xD5, 0xB4, 0x94}}

type ILanguage2Interface interface {
	win32.IInspectableInterface
	Get_LayoutDirection() LanguageLayoutDirection
}

type ILanguage2Vtbl struct {
	win32.IInspectableVtbl
	Get_LayoutDirection uintptr
}

type ILanguage2 struct {
	win32.IInspectable
}

func (this *ILanguage2) Vtbl() *ILanguage2Vtbl {
	return (*ILanguage2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILanguage2) Get_LayoutDirection() LanguageLayoutDirection {
	var _result LanguageLayoutDirection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LayoutDirection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C6AF3D10-641A-5BA4-BB43-5E12AED75954
var IID_ILanguage3 = syscall.GUID{0xC6AF3D10, 0x641A, 0x5BA4,
	[8]byte{0xBB, 0x43, 0x5E, 0x12, 0xAE, 0xD7, 0x59, 0x54}}

type ILanguage3Interface interface {
	win32.IInspectableInterface
	Get_AbbreviatedName() string
}

type ILanguage3Vtbl struct {
	win32.IInspectableVtbl
	Get_AbbreviatedName uintptr
}

type ILanguage3 struct {
	win32.IInspectable
}

func (this *ILanguage3) Vtbl() *ILanguage3Vtbl {
	return (*ILanguage3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILanguage3) Get_AbbreviatedName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AbbreviatedName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 7D7DAF45-368D-4364-852B-DEC927037B85
var IID_ILanguageExtensionSubtags = syscall.GUID{0x7D7DAF45, 0x368D, 0x4364,
	[8]byte{0x85, 0x2B, 0xDE, 0xC9, 0x27, 0x03, 0x7B, 0x85}}

type ILanguageExtensionSubtagsInterface interface {
	win32.IInspectableInterface
	GetExtensionSubtags(singleton string) *IVectorView[string]
}

type ILanguageExtensionSubtagsVtbl struct {
	win32.IInspectableVtbl
	GetExtensionSubtags uintptr
}

type ILanguageExtensionSubtags struct {
	win32.IInspectable
}

func (this *ILanguageExtensionSubtags) Vtbl() *ILanguageExtensionSubtagsVtbl {
	return (*ILanguageExtensionSubtagsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILanguageExtensionSubtags) GetExtensionSubtags(singleton string) *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetExtensionSubtags, uintptr(unsafe.Pointer(this)), NewHStr(singleton).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9B0252AC-0C27-44F8-B792-9793FB66C63E
var IID_ILanguageFactory = syscall.GUID{0x9B0252AC, 0x0C27, 0x44F8,
	[8]byte{0xB7, 0x92, 0x97, 0x93, 0xFB, 0x66, 0xC6, 0x3E}}

type ILanguageFactoryInterface interface {
	win32.IInspectableInterface
	CreateLanguage(languageTag string) *ILanguage
}

type ILanguageFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateLanguage uintptr
}

type ILanguageFactory struct {
	win32.IInspectable
}

func (this *ILanguageFactory) Vtbl() *ILanguageFactoryVtbl {
	return (*ILanguageFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILanguageFactory) CreateLanguage(languageTag string) *ILanguage {
	var _result *ILanguage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateLanguage, uintptr(unsafe.Pointer(this)), NewHStr(languageTag).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B23CD557-0865-46D4-89B8-D59BE8990F0D
var IID_ILanguageStatics = syscall.GUID{0xB23CD557, 0x0865, 0x46D4,
	[8]byte{0x89, 0xB8, 0xD5, 0x9B, 0xE8, 0x99, 0x0F, 0x0D}}

type ILanguageStaticsInterface interface {
	win32.IInspectableInterface
	IsWellFormed(languageTag string) bool
	Get_CurrentInputMethodLanguageTag() string
}

type ILanguageStaticsVtbl struct {
	win32.IInspectableVtbl
	IsWellFormed                      uintptr
	Get_CurrentInputMethodLanguageTag uintptr
}

type ILanguageStatics struct {
	win32.IInspectable
}

func (this *ILanguageStatics) Vtbl() *ILanguageStaticsVtbl {
	return (*ILanguageStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILanguageStatics) IsWellFormed(languageTag string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsWellFormed, uintptr(unsafe.Pointer(this)), NewHStr(languageTag).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILanguageStatics) Get_CurrentInputMethodLanguageTag() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentInputMethodLanguageTag, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 30199F6E-914B-4B2A-9D6E-E3B0E27DBE4F
var IID_ILanguageStatics2 = syscall.GUID{0x30199F6E, 0x914B, 0x4B2A,
	[8]byte{0x9D, 0x6E, 0xE3, 0xB0, 0xE2, 0x7D, 0xBE, 0x4F}}

type ILanguageStatics2Interface interface {
	win32.IInspectableInterface
	TrySetInputMethodLanguageTag(languageTag string) bool
}

type ILanguageStatics2Vtbl struct {
	win32.IInspectableVtbl
	TrySetInputMethodLanguageTag uintptr
}

type ILanguageStatics2 struct {
	win32.IInspectable
}

func (this *ILanguageStatics2) Vtbl() *ILanguageStatics2Vtbl {
	return (*ILanguageStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILanguageStatics2) TrySetInputMethodLanguageTag(languageTag string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySetInputMethodLanguageTag, uintptr(unsafe.Pointer(this)), NewHStr(languageTag).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D15ECB5A-71DE-5752-9542-FAC5B4F27261
var IID_ILanguageStatics3 = syscall.GUID{0xD15ECB5A, 0x71DE, 0x5752,
	[8]byte{0x95, 0x42, 0xFA, 0xC5, 0xB4, 0xF2, 0x72, 0x61}}

type ILanguageStatics3Interface interface {
	win32.IInspectableInterface
	GetMuiCompatibleLanguageListFromLanguageTags(languageTags *IIterable[string]) *IVector[string]
}

type ILanguageStatics3Vtbl struct {
	win32.IInspectableVtbl
	GetMuiCompatibleLanguageListFromLanguageTags uintptr
}

type ILanguageStatics3 struct {
	win32.IInspectable
}

func (this *ILanguageStatics3) Vtbl() *ILanguageStatics3Vtbl {
	return (*ILanguageStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILanguageStatics3) GetMuiCompatibleLanguageListFromLanguageTags(languageTags *IIterable[string]) *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetMuiCompatibleLanguageListFromLanguageTags, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(languageTags)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A5C662C3-68C9-4D3D-B765-972029E21DEC
var IID_INumeralSystemIdentifiersStatics = syscall.GUID{0xA5C662C3, 0x68C9, 0x4D3D,
	[8]byte{0xB7, 0x65, 0x97, 0x20, 0x29, 0xE2, 0x1D, 0xEC}}

type INumeralSystemIdentifiersStaticsInterface interface {
	win32.IInspectableInterface
	Get_Arab() string
	Get_ArabExt() string
	Get_Bali() string
	Get_Beng() string
	Get_Cham() string
	Get_Deva() string
	Get_FullWide() string
	Get_Gujr() string
	Get_Guru() string
	Get_HaniDec() string
	Get_Java() string
	Get_Kali() string
	Get_Khmr() string
	Get_Knda() string
	Get_Lana() string
	Get_LanaTham() string
	Get_Laoo() string
	Get_Latn() string
	Get_Lepc() string
	Get_Limb() string
	Get_Mlym() string
	Get_Mong() string
	Get_Mtei() string
	Get_Mymr() string
	Get_MymrShan() string
	Get_Nkoo() string
	Get_Olck() string
	Get_Orya() string
	Get_Saur() string
	Get_Sund() string
	Get_Talu() string
	Get_TamlDec() string
	Get_Telu() string
	Get_Thai() string
	Get_Tibt() string
	Get_Vaii() string
}

type INumeralSystemIdentifiersStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Arab     uintptr
	Get_ArabExt  uintptr
	Get_Bali     uintptr
	Get_Beng     uintptr
	Get_Cham     uintptr
	Get_Deva     uintptr
	Get_FullWide uintptr
	Get_Gujr     uintptr
	Get_Guru     uintptr
	Get_HaniDec  uintptr
	Get_Java     uintptr
	Get_Kali     uintptr
	Get_Khmr     uintptr
	Get_Knda     uintptr
	Get_Lana     uintptr
	Get_LanaTham uintptr
	Get_Laoo     uintptr
	Get_Latn     uintptr
	Get_Lepc     uintptr
	Get_Limb     uintptr
	Get_Mlym     uintptr
	Get_Mong     uintptr
	Get_Mtei     uintptr
	Get_Mymr     uintptr
	Get_MymrShan uintptr
	Get_Nkoo     uintptr
	Get_Olck     uintptr
	Get_Orya     uintptr
	Get_Saur     uintptr
	Get_Sund     uintptr
	Get_Talu     uintptr
	Get_TamlDec  uintptr
	Get_Telu     uintptr
	Get_Thai     uintptr
	Get_Tibt     uintptr
	Get_Vaii     uintptr
}

type INumeralSystemIdentifiersStatics struct {
	win32.IInspectable
}

func (this *INumeralSystemIdentifiersStatics) Vtbl() *INumeralSystemIdentifiersStaticsVtbl {
	return (*INumeralSystemIdentifiersStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INumeralSystemIdentifiersStatics) Get_Arab() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Arab, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_ArabExt() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ArabExt, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Bali() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bali, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Beng() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Beng, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Cham() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Cham, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Deva() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Deva, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_FullWide() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FullWide, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Gujr() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Gujr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Guru() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Guru, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_HaniDec() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HaniDec, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Java() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Java, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Kali() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Kali, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Khmr() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Khmr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Knda() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Knda, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Lana() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Lana, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_LanaTham() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LanaTham, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Laoo() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Laoo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Latn() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Latn, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Lepc() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Lepc, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Limb() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Limb, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Mlym() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mlym, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Mong() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mong, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Mtei() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mtei, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Mymr() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mymr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_MymrShan() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MymrShan, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Nkoo() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Nkoo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Olck() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Olck, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Orya() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Orya, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Saur() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Saur, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Sund() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Sund, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Talu() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Talu, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_TamlDec() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TamlDec, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Telu() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Telu, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Thai() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Thai, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Tibt() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Tibt, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics) Get_Vaii() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Vaii, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 7F003228-9DDB-4A34-9104-0260C091A7C7
var IID_INumeralSystemIdentifiersStatics2 = syscall.GUID{0x7F003228, 0x9DDB, 0x4A34,
	[8]byte{0x91, 0x04, 0x02, 0x60, 0xC0, 0x91, 0xA7, 0xC7}}

type INumeralSystemIdentifiersStatics2Interface interface {
	win32.IInspectableInterface
	Get_Brah() string
	Get_Osma() string
	Get_MathBold() string
	Get_MathDbl() string
	Get_MathSans() string
	Get_MathSanb() string
	Get_MathMono() string
	Get_ZmthBold() string
	Get_ZmthDbl() string
	Get_ZmthSans() string
	Get_ZmthSanb() string
	Get_ZmthMono() string
}

type INumeralSystemIdentifiersStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_Brah     uintptr
	Get_Osma     uintptr
	Get_MathBold uintptr
	Get_MathDbl  uintptr
	Get_MathSans uintptr
	Get_MathSanb uintptr
	Get_MathMono uintptr
	Get_ZmthBold uintptr
	Get_ZmthDbl  uintptr
	Get_ZmthSans uintptr
	Get_ZmthSanb uintptr
	Get_ZmthMono uintptr
}

type INumeralSystemIdentifiersStatics2 struct {
	win32.IInspectable
}

func (this *INumeralSystemIdentifiersStatics2) Vtbl() *INumeralSystemIdentifiersStatics2Vtbl {
	return (*INumeralSystemIdentifiersStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INumeralSystemIdentifiersStatics2) Get_Brah() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Brah, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics2) Get_Osma() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Osma, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics2) Get_MathBold() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MathBold, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics2) Get_MathDbl() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MathDbl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics2) Get_MathSans() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MathSans, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics2) Get_MathSanb() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MathSanb, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics2) Get_MathMono() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MathMono, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics2) Get_ZmthBold() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ZmthBold, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics2) Get_ZmthDbl() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ZmthDbl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics2) Get_ZmthSans() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ZmthSans, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics2) Get_ZmthSanb() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ZmthSanb, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INumeralSystemIdentifiersStatics2) Get_ZmthMono() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ZmthMono, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// BB3C25E5-46CF-4317-A3F5-02621AD54478
var IID_ITimeZoneOnCalendar = syscall.GUID{0xBB3C25E5, 0x46CF, 0x4317,
	[8]byte{0xA3, 0xF5, 0x02, 0x62, 0x1A, 0xD5, 0x44, 0x78}}

type ITimeZoneOnCalendarInterface interface {
	win32.IInspectableInterface
	GetTimeZone() string
	ChangeTimeZone(timeZoneId string)
	TimeZoneAsFullString() string
	TimeZoneAsString(idealLength int32) string
}

type ITimeZoneOnCalendarVtbl struct {
	win32.IInspectableVtbl
	GetTimeZone          uintptr
	ChangeTimeZone       uintptr
	TimeZoneAsFullString uintptr
	TimeZoneAsString     uintptr
}

type ITimeZoneOnCalendar struct {
	win32.IInspectable
}

func (this *ITimeZoneOnCalendar) Vtbl() *ITimeZoneOnCalendarVtbl {
	return (*ITimeZoneOnCalendarVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimeZoneOnCalendar) GetTimeZone() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTimeZone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ITimeZoneOnCalendar) ChangeTimeZone(timeZoneId string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ChangeTimeZone, uintptr(unsafe.Pointer(this)), NewHStr(timeZoneId).Ptr)
	_ = _hr
}

func (this *ITimeZoneOnCalendar) TimeZoneAsFullString() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TimeZoneAsFullString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ITimeZoneOnCalendar) TimeZoneAsString(idealLength int32) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TimeZoneAsString, uintptr(unsafe.Pointer(this)), uintptr(idealLength), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// classes

type Calendar struct {
	RtClass
	*ICalendar
}

func NewCalendar() *Calendar {
	hs := NewHStr("Windows.Globalization.Calendar")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &Calendar{
		RtClass:   RtClass{PInspect: p},
		ICalendar: (*ICalendar)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewCalendar_CreateCalendarWithTimeZone(languages *IIterable[string], calendar string, clock string, timeZoneId string) *Calendar {
	hs := NewHStr("Windows.Globalization.Calendar")
	var pFac *ICalendarFactory2
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICalendarFactory2, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ICalendar
	p = pFac.CreateCalendarWithTimeZone(languages, calendar, clock, timeZoneId)
	result := &Calendar{
		RtClass:   RtClass{PInspect: &p.IInspectable},
		ICalendar: p,
	}
	com.AddToScope(result)
	return result
}

type CalendarIdentifiers struct {
	RtClass
}

func NewICalendarIdentifiersStatics() *ICalendarIdentifiersStatics {
	var p *ICalendarIdentifiersStatics
	hs := NewHStr("Windows.Globalization.CalendarIdentifiers")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICalendarIdentifiersStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewICalendarIdentifiersStatics2() *ICalendarIdentifiersStatics2 {
	var p *ICalendarIdentifiersStatics2
	hs := NewHStr("Windows.Globalization.CalendarIdentifiers")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICalendarIdentifiersStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewICalendarIdentifiersStatics3() *ICalendarIdentifiersStatics3 {
	var p *ICalendarIdentifiersStatics3
	hs := NewHStr("Windows.Globalization.CalendarIdentifiers")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICalendarIdentifiersStatics3, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type ClockIdentifiers struct {
	RtClass
}

func NewIClockIdentifiersStatics() *IClockIdentifiersStatics {
	var p *IClockIdentifiersStatics
	hs := NewHStr("Windows.Globalization.ClockIdentifiers")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IClockIdentifiersStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type CurrencyIdentifiers struct {
	RtClass
}

func NewICurrencyIdentifiersStatics() *ICurrencyIdentifiersStatics {
	var p *ICurrencyIdentifiersStatics
	hs := NewHStr("Windows.Globalization.CurrencyIdentifiers")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICurrencyIdentifiersStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewICurrencyIdentifiersStatics2() *ICurrencyIdentifiersStatics2 {
	var p *ICurrencyIdentifiersStatics2
	hs := NewHStr("Windows.Globalization.CurrencyIdentifiers")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICurrencyIdentifiersStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewICurrencyIdentifiersStatics3() *ICurrencyIdentifiersStatics3 {
	var p *ICurrencyIdentifiersStatics3
	hs := NewHStr("Windows.Globalization.CurrencyIdentifiers")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICurrencyIdentifiersStatics3, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type GeographicRegion struct {
	RtClass
	*IGeographicRegion
}

func NewGeographicRegion() *GeographicRegion {
	hs := NewHStr("Windows.Globalization.GeographicRegion")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &GeographicRegion{
		RtClass:           RtClass{PInspect: p},
		IGeographicRegion: (*IGeographicRegion)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewGeographicRegion_CreateGeographicRegion(geographicRegionCode string) *GeographicRegion {
	hs := NewHStr("Windows.Globalization.GeographicRegion")
	var pFac *IGeographicRegionFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGeographicRegionFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IGeographicRegion
	p = pFac.CreateGeographicRegion(geographicRegionCode)
	result := &GeographicRegion{
		RtClass:           RtClass{PInspect: &p.IInspectable},
		IGeographicRegion: p,
	}
	com.AddToScope(result)
	return result
}

func NewIGeographicRegionStatics() *IGeographicRegionStatics {
	var p *IGeographicRegionStatics
	hs := NewHStr("Windows.Globalization.GeographicRegion")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGeographicRegionStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type Language struct {
	RtClass
	*ILanguage
}

func NewLanguage_CreateLanguage(languageTag string) *Language {
	hs := NewHStr("Windows.Globalization.Language")
	var pFac *ILanguageFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ILanguageFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ILanguage
	p = pFac.CreateLanguage(languageTag)
	result := &Language{
		RtClass:   RtClass{PInspect: &p.IInspectable},
		ILanguage: p,
	}
	com.AddToScope(result)
	return result
}

func NewILanguageStatics() *ILanguageStatics {
	var p *ILanguageStatics
	hs := NewHStr("Windows.Globalization.Language")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ILanguageStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewILanguageStatics2() *ILanguageStatics2 {
	var p *ILanguageStatics2
	hs := NewHStr("Windows.Globalization.Language")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ILanguageStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewILanguageStatics3() *ILanguageStatics3 {
	var p *ILanguageStatics3
	hs := NewHStr("Windows.Globalization.Language")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ILanguageStatics3, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type NumeralSystemIdentifiers struct {
	RtClass
}

func NewINumeralSystemIdentifiersStatics2() *INumeralSystemIdentifiersStatics2 {
	var p *INumeralSystemIdentifiersStatics2
	hs := NewHStr("Windows.Globalization.NumeralSystemIdentifiers")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_INumeralSystemIdentifiersStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewINumeralSystemIdentifiersStatics() *INumeralSystemIdentifiersStatics {
	var p *INumeralSystemIdentifiersStatics
	hs := NewHStr("Windows.Globalization.NumeralSystemIdentifiers")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_INumeralSystemIdentifiersStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
