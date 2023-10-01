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
type AlternateNormalizationFormat int32

const (
	AlternateNormalizationFormat_NotNormalized AlternateNormalizationFormat = 0
	AlternateNormalizationFormat_Number        AlternateNormalizationFormat = 1
	AlternateNormalizationFormat_Currency      AlternateNormalizationFormat = 3
	AlternateNormalizationFormat_Date          AlternateNormalizationFormat = 4
	AlternateNormalizationFormat_Time          AlternateNormalizationFormat = 5
)

// enum
// flags
type TextPredictionOptions uint32

const (
	TextPredictionOptions_None        TextPredictionOptions = 0
	TextPredictionOptions_Predictions TextPredictionOptions = 1
	TextPredictionOptions_Corrections TextPredictionOptions = 2
)

// enum
type UnicodeGeneralCategory int32

const (
	UnicodeGeneralCategory_UppercaseLetter         UnicodeGeneralCategory = 0
	UnicodeGeneralCategory_LowercaseLetter         UnicodeGeneralCategory = 1
	UnicodeGeneralCategory_TitlecaseLetter         UnicodeGeneralCategory = 2
	UnicodeGeneralCategory_ModifierLetter          UnicodeGeneralCategory = 3
	UnicodeGeneralCategory_OtherLetter             UnicodeGeneralCategory = 4
	UnicodeGeneralCategory_NonspacingMark          UnicodeGeneralCategory = 5
	UnicodeGeneralCategory_SpacingCombiningMark    UnicodeGeneralCategory = 6
	UnicodeGeneralCategory_EnclosingMark           UnicodeGeneralCategory = 7
	UnicodeGeneralCategory_DecimalDigitNumber      UnicodeGeneralCategory = 8
	UnicodeGeneralCategory_LetterNumber            UnicodeGeneralCategory = 9
	UnicodeGeneralCategory_OtherNumber             UnicodeGeneralCategory = 10
	UnicodeGeneralCategory_SpaceSeparator          UnicodeGeneralCategory = 11
	UnicodeGeneralCategory_LineSeparator           UnicodeGeneralCategory = 12
	UnicodeGeneralCategory_ParagraphSeparator      UnicodeGeneralCategory = 13
	UnicodeGeneralCategory_Control                 UnicodeGeneralCategory = 14
	UnicodeGeneralCategory_Format                  UnicodeGeneralCategory = 15
	UnicodeGeneralCategory_Surrogate               UnicodeGeneralCategory = 16
	UnicodeGeneralCategory_PrivateUse              UnicodeGeneralCategory = 17
	UnicodeGeneralCategory_ConnectorPunctuation    UnicodeGeneralCategory = 18
	UnicodeGeneralCategory_DashPunctuation         UnicodeGeneralCategory = 19
	UnicodeGeneralCategory_OpenPunctuation         UnicodeGeneralCategory = 20
	UnicodeGeneralCategory_ClosePunctuation        UnicodeGeneralCategory = 21
	UnicodeGeneralCategory_InitialQuotePunctuation UnicodeGeneralCategory = 22
	UnicodeGeneralCategory_FinalQuotePunctuation   UnicodeGeneralCategory = 23
	UnicodeGeneralCategory_OtherPunctuation        UnicodeGeneralCategory = 24
	UnicodeGeneralCategory_MathSymbol              UnicodeGeneralCategory = 25
	UnicodeGeneralCategory_CurrencySymbol          UnicodeGeneralCategory = 26
	UnicodeGeneralCategory_ModifierSymbol          UnicodeGeneralCategory = 27
	UnicodeGeneralCategory_OtherSymbol             UnicodeGeneralCategory = 28
	UnicodeGeneralCategory_NotAssigned             UnicodeGeneralCategory = 29
)

// enum
type UnicodeNumericType int32

const (
	UnicodeNumericType_None    UnicodeNumericType = 0
	UnicodeNumericType_Decimal UnicodeNumericType = 1
	UnicodeNumericType_Digit   UnicodeNumericType = 2
	UnicodeNumericType_Numeric UnicodeNumericType = 3
)

// structs

type TextSegment struct {
	StartPosition uint32
	Length        uint32
}

// func types

// 3A3DFC9C-AEDE-4DC7-9E6C-41C044BD3592
type SelectableWordSegmentsTokenizingHandler func(precedingWords *IIterable[*ISelectableWordSegment], words *IIterable[*ISelectableWordSegment]) com.Error

// A5DD6357-BF2A-4C4F-A31F-29E71C6F8B35
type WordSegmentsTokenizingHandler func(precedingWords *IIterable[*IWordSegment], words *IIterable[*IWordSegment]) com.Error

// interfaces

// 47396C1E-51B9-4207-9146-248E636A1D1D
var IID_IAlternateWordForm = syscall.GUID{0x47396C1E, 0x51B9, 0x4207,
	[8]byte{0x91, 0x46, 0x24, 0x8E, 0x63, 0x6A, 0x1D, 0x1D}}

type IAlternateWordFormInterface interface {
	win32.IInspectableInterface
	Get_SourceTextSegment() TextSegment
	Get_AlternateText() string
	Get_NormalizationFormat() AlternateNormalizationFormat
}

type IAlternateWordFormVtbl struct {
	win32.IInspectableVtbl
	Get_SourceTextSegment   uintptr
	Get_AlternateText       uintptr
	Get_NormalizationFormat uintptr
}

type IAlternateWordForm struct {
	win32.IInspectable
}

func (this *IAlternateWordForm) Vtbl() *IAlternateWordFormVtbl {
	return (*IAlternateWordFormVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAlternateWordForm) Get_SourceTextSegment() TextSegment {
	var _result TextSegment
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SourceTextSegment, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAlternateWordForm) Get_AlternateText() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlternateText, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAlternateWordForm) Get_NormalizationFormat() AlternateNormalizationFormat {
	var _result AlternateNormalizationFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NormalizationFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 916A4CB7-8AA7-4C78-B374-5DEDB752E60B
var IID_ISelectableWordSegment = syscall.GUID{0x916A4CB7, 0x8AA7, 0x4C78,
	[8]byte{0xB3, 0x74, 0x5D, 0xED, 0xB7, 0x52, 0xE6, 0x0B}}

type ISelectableWordSegmentInterface interface {
	win32.IInspectableInterface
	Get_Text() string
	Get_SourceTextSegment() TextSegment
}

type ISelectableWordSegmentVtbl struct {
	win32.IInspectableVtbl
	Get_Text              uintptr
	Get_SourceTextSegment uintptr
}

type ISelectableWordSegment struct {
	win32.IInspectable
}

func (this *ISelectableWordSegment) Vtbl() *ISelectableWordSegmentVtbl {
	return (*ISelectableWordSegmentVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISelectableWordSegment) Get_Text() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Text, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISelectableWordSegment) Get_SourceTextSegment() TextSegment {
	var _result TextSegment
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SourceTextSegment, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F6DC31E7-4B13-45C5-8897-7D71269E085D
var IID_ISelectableWordsSegmenter = syscall.GUID{0xF6DC31E7, 0x4B13, 0x45C5,
	[8]byte{0x88, 0x97, 0x7D, 0x71, 0x26, 0x9E, 0x08, 0x5D}}

type ISelectableWordsSegmenterInterface interface {
	win32.IInspectableInterface
	Get_ResolvedLanguage() string
	GetTokenAt(text string, startIndex uint32) *ISelectableWordSegment
	GetTokens(text string) *IVectorView[*ISelectableWordSegment]
	Tokenize(text string, startIndex uint32, handler SelectableWordSegmentsTokenizingHandler)
}

type ISelectableWordsSegmenterVtbl struct {
	win32.IInspectableVtbl
	Get_ResolvedLanguage uintptr
	GetTokenAt           uintptr
	GetTokens            uintptr
	Tokenize             uintptr
}

type ISelectableWordsSegmenter struct {
	win32.IInspectable
}

func (this *ISelectableWordsSegmenter) Vtbl() *ISelectableWordsSegmenterVtbl {
	return (*ISelectableWordsSegmenterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISelectableWordsSegmenter) Get_ResolvedLanguage() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResolvedLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISelectableWordsSegmenter) GetTokenAt(text string, startIndex uint32) *ISelectableWordSegment {
	var _result *ISelectableWordSegment
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTokenAt, uintptr(unsafe.Pointer(this)), NewHStr(text).Ptr, uintptr(startIndex), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISelectableWordsSegmenter) GetTokens(text string) *IVectorView[*ISelectableWordSegment] {
	var _result *IVectorView[*ISelectableWordSegment]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTokens, uintptr(unsafe.Pointer(this)), NewHStr(text).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISelectableWordsSegmenter) Tokenize(text string, startIndex uint32, handler SelectableWordSegmentsTokenizingHandler) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Tokenize, uintptr(unsafe.Pointer(this)), NewHStr(text).Ptr, uintptr(startIndex), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))))
	_ = _hr
}

// 8C7A7648-6057-4339-BC70-F210010A4150
var IID_ISelectableWordsSegmenterFactory = syscall.GUID{0x8C7A7648, 0x6057, 0x4339,
	[8]byte{0xBC, 0x70, 0xF2, 0x10, 0x01, 0x0A, 0x41, 0x50}}

type ISelectableWordsSegmenterFactoryInterface interface {
	win32.IInspectableInterface
	CreateWithLanguage(language string) *ISelectableWordsSegmenter
}

type ISelectableWordsSegmenterFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateWithLanguage uintptr
}

type ISelectableWordsSegmenterFactory struct {
	win32.IInspectable
}

func (this *ISelectableWordsSegmenterFactory) Vtbl() *ISelectableWordsSegmenterFactoryVtbl {
	return (*ISelectableWordsSegmenterFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISelectableWordsSegmenterFactory) CreateWithLanguage(language string) *ISelectableWordsSegmenter {
	var _result *ISelectableWordsSegmenter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithLanguage, uintptr(unsafe.Pointer(this)), NewHStr(language).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6A1CAB51-1FB2-4909-80B8-35731A2B3E7F
var IID_ISemanticTextQuery = syscall.GUID{0x6A1CAB51, 0x1FB2, 0x4909,
	[8]byte{0x80, 0xB8, 0x35, 0x73, 0x1A, 0x2B, 0x3E, 0x7F}}

type ISemanticTextQueryInterface interface {
	win32.IInspectableInterface
	Find(content string) *IVectorView[TextSegment]
	FindInProperty(propertyContent string, propertyName string) *IVectorView[TextSegment]
}

type ISemanticTextQueryVtbl struct {
	win32.IInspectableVtbl
	Find           uintptr
	FindInProperty uintptr
}

type ISemanticTextQuery struct {
	win32.IInspectable
}

func (this *ISemanticTextQuery) Vtbl() *ISemanticTextQueryVtbl {
	return (*ISemanticTextQueryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISemanticTextQuery) Find(content string) *IVectorView[TextSegment] {
	var _result *IVectorView[TextSegment]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Find, uintptr(unsafe.Pointer(this)), NewHStr(content).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISemanticTextQuery) FindInProperty(propertyContent string, propertyName string) *IVectorView[TextSegment] {
	var _result *IVectorView[TextSegment]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindInProperty, uintptr(unsafe.Pointer(this)), NewHStr(propertyContent).Ptr, NewHStr(propertyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 238C0503-F995-4587-8777-A2B7D80ACFEF
var IID_ISemanticTextQueryFactory = syscall.GUID{0x238C0503, 0xF995, 0x4587,
	[8]byte{0x87, 0x77, 0xA2, 0xB7, 0xD8, 0x0A, 0xCF, 0xEF}}

type ISemanticTextQueryFactoryInterface interface {
	win32.IInspectableInterface
	Create(aqsFilter string) *ISemanticTextQuery
	CreateWithLanguage(aqsFilter string, filterLanguage string) *ISemanticTextQuery
}

type ISemanticTextQueryFactoryVtbl struct {
	win32.IInspectableVtbl
	Create             uintptr
	CreateWithLanguage uintptr
}

type ISemanticTextQueryFactory struct {
	win32.IInspectable
}

func (this *ISemanticTextQueryFactory) Vtbl() *ISemanticTextQueryFactoryVtbl {
	return (*ISemanticTextQueryFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISemanticTextQueryFactory) Create(aqsFilter string) *ISemanticTextQuery {
	var _result *ISemanticTextQuery
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(aqsFilter).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISemanticTextQueryFactory) CreateWithLanguage(aqsFilter string, filterLanguage string) *ISemanticTextQuery {
	var _result *ISemanticTextQuery
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithLanguage, uintptr(unsafe.Pointer(this)), NewHStr(aqsFilter).Ptr, NewHStr(filterLanguage).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 03606A5E-2AA9-4AB6-AF8B-A562B63A8992
var IID_ITextConversionGenerator = syscall.GUID{0x03606A5E, 0x2AA9, 0x4AB6,
	[8]byte{0xAF, 0x8B, 0xA5, 0x62, 0xB6, 0x3A, 0x89, 0x92}}

type ITextConversionGeneratorInterface interface {
	win32.IInspectableInterface
	Get_ResolvedLanguage() string
	Get_LanguageAvailableButNotInstalled() bool
	GetCandidatesAsync(input string) *IAsyncOperation[*IVectorView[string]]
	GetCandidatesWithMaxCountAsync(input string, maxCandidates uint32) *IAsyncOperation[*IVectorView[string]]
}

type ITextConversionGeneratorVtbl struct {
	win32.IInspectableVtbl
	Get_ResolvedLanguage                 uintptr
	Get_LanguageAvailableButNotInstalled uintptr
	GetCandidatesAsync                   uintptr
	GetCandidatesWithMaxCountAsync       uintptr
}

type ITextConversionGenerator struct {
	win32.IInspectable
}

func (this *ITextConversionGenerator) Vtbl() *ITextConversionGeneratorVtbl {
	return (*ITextConversionGeneratorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITextConversionGenerator) Get_ResolvedLanguage() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResolvedLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ITextConversionGenerator) Get_LanguageAvailableButNotInstalled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LanguageAvailableButNotInstalled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITextConversionGenerator) GetCandidatesAsync(input string) *IAsyncOperation[*IVectorView[string]] {
	var _result *IAsyncOperation[*IVectorView[string]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCandidatesAsync, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITextConversionGenerator) GetCandidatesWithMaxCountAsync(input string, maxCandidates uint32) *IAsyncOperation[*IVectorView[string]] {
	var _result *IAsyncOperation[*IVectorView[string]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCandidatesWithMaxCountAsync, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(maxCandidates), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FCAA3781-3083-49AB-BE15-56DFBBB74D6F
var IID_ITextConversionGeneratorFactory = syscall.GUID{0xFCAA3781, 0x3083, 0x49AB,
	[8]byte{0xBE, 0x15, 0x56, 0xDF, 0xBB, 0xB7, 0x4D, 0x6F}}

type ITextConversionGeneratorFactoryInterface interface {
	win32.IInspectableInterface
	Create(languageTag string) *ITextConversionGenerator
}

type ITextConversionGeneratorFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ITextConversionGeneratorFactory struct {
	win32.IInspectable
}

func (this *ITextConversionGeneratorFactory) Vtbl() *ITextConversionGeneratorFactoryVtbl {
	return (*ITextConversionGeneratorFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITextConversionGeneratorFactory) Create(languageTag string) *ITextConversionGenerator {
	var _result *ITextConversionGenerator
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(languageTag).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9362A40A-9B7A-4569-94CF-D84F2F38CF9B
var IID_ITextPhoneme = syscall.GUID{0x9362A40A, 0x9B7A, 0x4569,
	[8]byte{0x94, 0xCF, 0xD8, 0x4F, 0x2F, 0x38, 0xCF, 0x9B}}

type ITextPhonemeInterface interface {
	win32.IInspectableInterface
	Get_DisplayText() string
	Get_ReadingText() string
}

type ITextPhonemeVtbl struct {
	win32.IInspectableVtbl
	Get_DisplayText uintptr
	Get_ReadingText uintptr
}

type ITextPhoneme struct {
	win32.IInspectable
}

func (this *ITextPhoneme) Vtbl() *ITextPhonemeVtbl {
	return (*ITextPhonemeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITextPhoneme) Get_DisplayText() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayText, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ITextPhoneme) Get_ReadingText() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReadingText, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 5EACAB07-ABF1-4CB6-9D9E-326F2B468756
var IID_ITextPredictionGenerator = syscall.GUID{0x5EACAB07, 0xABF1, 0x4CB6,
	[8]byte{0x9D, 0x9E, 0x32, 0x6F, 0x2B, 0x46, 0x87, 0x56}}

type ITextPredictionGeneratorInterface interface {
	win32.IInspectableInterface
	Get_ResolvedLanguage() string
	Get_LanguageAvailableButNotInstalled() bool
	GetCandidatesAsync(input string) *IAsyncOperation[*IVectorView[string]]
	GetCandidatesWithMaxCountAsync(input string, maxCandidates uint32) *IAsyncOperation[*IVectorView[string]]
}

type ITextPredictionGeneratorVtbl struct {
	win32.IInspectableVtbl
	Get_ResolvedLanguage                 uintptr
	Get_LanguageAvailableButNotInstalled uintptr
	GetCandidatesAsync                   uintptr
	GetCandidatesWithMaxCountAsync       uintptr
}

type ITextPredictionGenerator struct {
	win32.IInspectable
}

func (this *ITextPredictionGenerator) Vtbl() *ITextPredictionGeneratorVtbl {
	return (*ITextPredictionGeneratorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITextPredictionGenerator) Get_ResolvedLanguage() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResolvedLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ITextPredictionGenerator) Get_LanguageAvailableButNotInstalled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LanguageAvailableButNotInstalled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITextPredictionGenerator) GetCandidatesAsync(input string) *IAsyncOperation[*IVectorView[string]] {
	var _result *IAsyncOperation[*IVectorView[string]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCandidatesAsync, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITextPredictionGenerator) GetCandidatesWithMaxCountAsync(input string, maxCandidates uint32) *IAsyncOperation[*IVectorView[string]] {
	var _result *IAsyncOperation[*IVectorView[string]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCandidatesWithMaxCountAsync, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(maxCandidates), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B84723B8-2C77-486A-900A-A3453EEDC15D
var IID_ITextPredictionGenerator2 = syscall.GUID{0xB84723B8, 0x2C77, 0x486A,
	[8]byte{0x90, 0x0A, 0xA3, 0x45, 0x3E, 0xED, 0xC1, 0x5D}}

type ITextPredictionGenerator2Interface interface {
	win32.IInspectableInterface
	GetCandidatesWithParametersAsync(input string, maxCandidates uint32, predictionOptions TextPredictionOptions, previousStrings *IIterable[string]) *IAsyncOperation[*IVectorView[string]]
	GetNextWordCandidatesAsync(maxCandidates uint32, previousStrings *IIterable[string]) *IAsyncOperation[*IVectorView[string]]
	Get_InputScope() unsafe.Pointer
	Put_InputScope(value unsafe.Pointer)
}

type ITextPredictionGenerator2Vtbl struct {
	win32.IInspectableVtbl
	GetCandidatesWithParametersAsync uintptr
	GetNextWordCandidatesAsync       uintptr
	Get_InputScope                   uintptr
	Put_InputScope                   uintptr
}

type ITextPredictionGenerator2 struct {
	win32.IInspectable
}

func (this *ITextPredictionGenerator2) Vtbl() *ITextPredictionGenerator2Vtbl {
	return (*ITextPredictionGenerator2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITextPredictionGenerator2) GetCandidatesWithParametersAsync(input string, maxCandidates uint32, predictionOptions TextPredictionOptions, previousStrings *IIterable[string]) *IAsyncOperation[*IVectorView[string]] {
	var _result *IAsyncOperation[*IVectorView[string]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCandidatesWithParametersAsync, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(maxCandidates), uintptr(predictionOptions), uintptr(unsafe.Pointer(previousStrings)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITextPredictionGenerator2) GetNextWordCandidatesAsync(maxCandidates uint32, previousStrings *IIterable[string]) *IAsyncOperation[*IVectorView[string]] {
	var _result *IAsyncOperation[*IVectorView[string]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNextWordCandidatesAsync, uintptr(unsafe.Pointer(this)), uintptr(maxCandidates), uintptr(unsafe.Pointer(previousStrings)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITextPredictionGenerator2) Get_InputScope() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InputScope, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITextPredictionGenerator2) Put_InputScope(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InputScope, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 7257B416-8BA2-4751-9D30-9D85435653A2
var IID_ITextPredictionGeneratorFactory = syscall.GUID{0x7257B416, 0x8BA2, 0x4751,
	[8]byte{0x9D, 0x30, 0x9D, 0x85, 0x43, 0x56, 0x53, 0xA2}}

type ITextPredictionGeneratorFactoryInterface interface {
	win32.IInspectableInterface
	Create(languageTag string) *ITextPredictionGenerator
}

type ITextPredictionGeneratorFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ITextPredictionGeneratorFactory struct {
	win32.IInspectable
}

func (this *ITextPredictionGeneratorFactory) Vtbl() *ITextPredictionGeneratorFactoryVtbl {
	return (*ITextPredictionGeneratorFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITextPredictionGeneratorFactory) Create(languageTag string) *ITextPredictionGenerator {
	var _result *ITextPredictionGenerator
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(languageTag).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 51E7F514-9C51-4D86-AE1B-B498FBAD8313
var IID_ITextReverseConversionGenerator = syscall.GUID{0x51E7F514, 0x9C51, 0x4D86,
	[8]byte{0xAE, 0x1B, 0xB4, 0x98, 0xFB, 0xAD, 0x83, 0x13}}

type ITextReverseConversionGeneratorInterface interface {
	win32.IInspectableInterface
	Get_ResolvedLanguage() string
	Get_LanguageAvailableButNotInstalled() bool
	ConvertBackAsync(input string) *IAsyncOperation[string]
}

type ITextReverseConversionGeneratorVtbl struct {
	win32.IInspectableVtbl
	Get_ResolvedLanguage                 uintptr
	Get_LanguageAvailableButNotInstalled uintptr
	ConvertBackAsync                     uintptr
}

type ITextReverseConversionGenerator struct {
	win32.IInspectable
}

func (this *ITextReverseConversionGenerator) Vtbl() *ITextReverseConversionGeneratorVtbl {
	return (*ITextReverseConversionGeneratorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITextReverseConversionGenerator) Get_ResolvedLanguage() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResolvedLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ITextReverseConversionGenerator) Get_LanguageAvailableButNotInstalled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LanguageAvailableButNotInstalled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITextReverseConversionGenerator) ConvertBackAsync(input string) *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConvertBackAsync, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1AAFD2EC-85D6-46FD-828A-3A4830FA6E18
var IID_ITextReverseConversionGenerator2 = syscall.GUID{0x1AAFD2EC, 0x85D6, 0x46FD,
	[8]byte{0x82, 0x8A, 0x3A, 0x48, 0x30, 0xFA, 0x6E, 0x18}}

type ITextReverseConversionGenerator2Interface interface {
	win32.IInspectableInterface
	GetPhonemesAsync(input string) *IAsyncOperation[*IVectorView[*ITextPhoneme]]
}

type ITextReverseConversionGenerator2Vtbl struct {
	win32.IInspectableVtbl
	GetPhonemesAsync uintptr
}

type ITextReverseConversionGenerator2 struct {
	win32.IInspectable
}

func (this *ITextReverseConversionGenerator2) Vtbl() *ITextReverseConversionGenerator2Vtbl {
	return (*ITextReverseConversionGenerator2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITextReverseConversionGenerator2) GetPhonemesAsync(input string) *IAsyncOperation[*IVectorView[*ITextPhoneme]] {
	var _result *IAsyncOperation[*IVectorView[*ITextPhoneme]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPhonemesAsync, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 63BED326-1FDA-41F6-89D5-23DDEA3C729A
var IID_ITextReverseConversionGeneratorFactory = syscall.GUID{0x63BED326, 0x1FDA, 0x41F6,
	[8]byte{0x89, 0xD5, 0x23, 0xDD, 0xEA, 0x3C, 0x72, 0x9A}}

type ITextReverseConversionGeneratorFactoryInterface interface {
	win32.IInspectableInterface
	Create(languageTag string) *ITextReverseConversionGenerator
}

type ITextReverseConversionGeneratorFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ITextReverseConversionGeneratorFactory struct {
	win32.IInspectable
}

func (this *ITextReverseConversionGeneratorFactory) Vtbl() *ITextReverseConversionGeneratorFactoryVtbl {
	return (*ITextReverseConversionGeneratorFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITextReverseConversionGeneratorFactory) Create(languageTag string) *ITextReverseConversionGenerator {
	var _result *ITextReverseConversionGenerator
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(languageTag).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 97909E87-9291-4F91-B6C8-B6E359D7A7FB
var IID_IUnicodeCharactersStatics = syscall.GUID{0x97909E87, 0x9291, 0x4F91,
	[8]byte{0xB6, 0xC8, 0xB6, 0xE3, 0x59, 0xD7, 0xA7, 0xFB}}

type IUnicodeCharactersStaticsInterface interface {
	win32.IInspectableInterface
	GetCodepointFromSurrogatePair(highSurrogate uint32, lowSurrogate uint32) uint32
	GetSurrogatePairFromCodepoint(codepoint uint32, highSurrogate uint16, lowSurrogate uint16)
	IsHighSurrogate(codepoint uint32) bool
	IsLowSurrogate(codepoint uint32) bool
	IsSupplementary(codepoint uint32) bool
	IsNoncharacter(codepoint uint32) bool
	IsWhitespace(codepoint uint32) bool
	IsAlphabetic(codepoint uint32) bool
	IsCased(codepoint uint32) bool
	IsUppercase(codepoint uint32) bool
	IsLowercase(codepoint uint32) bool
	IsIdStart(codepoint uint32) bool
	IsIdContinue(codepoint uint32) bool
	IsGraphemeBase(codepoint uint32) bool
	IsGraphemeExtend(codepoint uint32) bool
	GetNumericType(codepoint uint32) UnicodeNumericType
	GetGeneralCategory(codepoint uint32) UnicodeGeneralCategory
}

type IUnicodeCharactersStaticsVtbl struct {
	win32.IInspectableVtbl
	GetCodepointFromSurrogatePair uintptr
	GetSurrogatePairFromCodepoint uintptr
	IsHighSurrogate               uintptr
	IsLowSurrogate                uintptr
	IsSupplementary               uintptr
	IsNoncharacter                uintptr
	IsWhitespace                  uintptr
	IsAlphabetic                  uintptr
	IsCased                       uintptr
	IsUppercase                   uintptr
	IsLowercase                   uintptr
	IsIdStart                     uintptr
	IsIdContinue                  uintptr
	IsGraphemeBase                uintptr
	IsGraphemeExtend              uintptr
	GetNumericType                uintptr
	GetGeneralCategory            uintptr
}

type IUnicodeCharactersStatics struct {
	win32.IInspectable
}

func (this *IUnicodeCharactersStatics) Vtbl() *IUnicodeCharactersStaticsVtbl {
	return (*IUnicodeCharactersStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUnicodeCharactersStatics) GetCodepointFromSurrogatePair(highSurrogate uint32, lowSurrogate uint32) uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCodepointFromSurrogatePair, uintptr(unsafe.Pointer(this)), uintptr(highSurrogate), uintptr(lowSurrogate), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUnicodeCharactersStatics) GetSurrogatePairFromCodepoint(codepoint uint32, highSurrogate uint16, lowSurrogate uint16) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSurrogatePairFromCodepoint, uintptr(unsafe.Pointer(this)), uintptr(codepoint), uintptr(highSurrogate), uintptr(lowSurrogate))
	_ = _hr
}

func (this *IUnicodeCharactersStatics) IsHighSurrogate(codepoint uint32) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsHighSurrogate, uintptr(unsafe.Pointer(this)), uintptr(codepoint), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUnicodeCharactersStatics) IsLowSurrogate(codepoint uint32) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsLowSurrogate, uintptr(unsafe.Pointer(this)), uintptr(codepoint), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUnicodeCharactersStatics) IsSupplementary(codepoint uint32) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsSupplementary, uintptr(unsafe.Pointer(this)), uintptr(codepoint), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUnicodeCharactersStatics) IsNoncharacter(codepoint uint32) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsNoncharacter, uintptr(unsafe.Pointer(this)), uintptr(codepoint), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUnicodeCharactersStatics) IsWhitespace(codepoint uint32) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsWhitespace, uintptr(unsafe.Pointer(this)), uintptr(codepoint), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUnicodeCharactersStatics) IsAlphabetic(codepoint uint32) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsAlphabetic, uintptr(unsafe.Pointer(this)), uintptr(codepoint), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUnicodeCharactersStatics) IsCased(codepoint uint32) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsCased, uintptr(unsafe.Pointer(this)), uintptr(codepoint), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUnicodeCharactersStatics) IsUppercase(codepoint uint32) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsUppercase, uintptr(unsafe.Pointer(this)), uintptr(codepoint), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUnicodeCharactersStatics) IsLowercase(codepoint uint32) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsLowercase, uintptr(unsafe.Pointer(this)), uintptr(codepoint), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUnicodeCharactersStatics) IsIdStart(codepoint uint32) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsIdStart, uintptr(unsafe.Pointer(this)), uintptr(codepoint), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUnicodeCharactersStatics) IsIdContinue(codepoint uint32) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsIdContinue, uintptr(unsafe.Pointer(this)), uintptr(codepoint), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUnicodeCharactersStatics) IsGraphemeBase(codepoint uint32) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsGraphemeBase, uintptr(unsafe.Pointer(this)), uintptr(codepoint), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUnicodeCharactersStatics) IsGraphemeExtend(codepoint uint32) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsGraphemeExtend, uintptr(unsafe.Pointer(this)), uintptr(codepoint), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUnicodeCharactersStatics) GetNumericType(codepoint uint32) UnicodeNumericType {
	var _result UnicodeNumericType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNumericType, uintptr(unsafe.Pointer(this)), uintptr(codepoint), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUnicodeCharactersStatics) GetGeneralCategory(codepoint uint32) UnicodeGeneralCategory {
	var _result UnicodeGeneralCategory
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetGeneralCategory, uintptr(unsafe.Pointer(this)), uintptr(codepoint), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D2D4BA6D-987C-4CC0-B6BD-D49A11B38F9A
var IID_IWordSegment = syscall.GUID{0xD2D4BA6D, 0x987C, 0x4CC0,
	[8]byte{0xB6, 0xBD, 0xD4, 0x9A, 0x11, 0xB3, 0x8F, 0x9A}}

type IWordSegmentInterface interface {
	win32.IInspectableInterface
	Get_Text() string
	Get_SourceTextSegment() TextSegment
	Get_AlternateForms() *IVectorView[*IAlternateWordForm]
}

type IWordSegmentVtbl struct {
	win32.IInspectableVtbl
	Get_Text              uintptr
	Get_SourceTextSegment uintptr
	Get_AlternateForms    uintptr
}

type IWordSegment struct {
	win32.IInspectable
}

func (this *IWordSegment) Vtbl() *IWordSegmentVtbl {
	return (*IWordSegmentVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWordSegment) Get_Text() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Text, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWordSegment) Get_SourceTextSegment() TextSegment {
	var _result TextSegment
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SourceTextSegment, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWordSegment) Get_AlternateForms() *IVectorView[*IAlternateWordForm] {
	var _result *IVectorView[*IAlternateWordForm]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlternateForms, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 86B4D4D1-B2FE-4E34-A81D-66640300454F
var IID_IWordsSegmenter = syscall.GUID{0x86B4D4D1, 0xB2FE, 0x4E34,
	[8]byte{0xA8, 0x1D, 0x66, 0x64, 0x03, 0x00, 0x45, 0x4F}}

type IWordsSegmenterInterface interface {
	win32.IInspectableInterface
	Get_ResolvedLanguage() string
	GetTokenAt(text string, startIndex uint32) *IWordSegment
	GetTokens(text string) *IVectorView[*IWordSegment]
	Tokenize(text string, startIndex uint32, handler WordSegmentsTokenizingHandler)
}

type IWordsSegmenterVtbl struct {
	win32.IInspectableVtbl
	Get_ResolvedLanguage uintptr
	GetTokenAt           uintptr
	GetTokens            uintptr
	Tokenize             uintptr
}

type IWordsSegmenter struct {
	win32.IInspectable
}

func (this *IWordsSegmenter) Vtbl() *IWordsSegmenterVtbl {
	return (*IWordsSegmenterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWordsSegmenter) Get_ResolvedLanguage() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResolvedLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWordsSegmenter) GetTokenAt(text string, startIndex uint32) *IWordSegment {
	var _result *IWordSegment
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTokenAt, uintptr(unsafe.Pointer(this)), NewHStr(text).Ptr, uintptr(startIndex), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWordsSegmenter) GetTokens(text string) *IVectorView[*IWordSegment] {
	var _result *IVectorView[*IWordSegment]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTokens, uintptr(unsafe.Pointer(this)), NewHStr(text).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWordsSegmenter) Tokenize(text string, startIndex uint32, handler WordSegmentsTokenizingHandler) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Tokenize, uintptr(unsafe.Pointer(this)), NewHStr(text).Ptr, uintptr(startIndex), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))))
	_ = _hr
}

// E6977274-FC35-455C-8BFB-6D7F4653CA97
var IID_IWordsSegmenterFactory = syscall.GUID{0xE6977274, 0xFC35, 0x455C,
	[8]byte{0x8B, 0xFB, 0x6D, 0x7F, 0x46, 0x53, 0xCA, 0x97}}

type IWordsSegmenterFactoryInterface interface {
	win32.IInspectableInterface
	CreateWithLanguage(language string) *IWordsSegmenter
}

type IWordsSegmenterFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateWithLanguage uintptr
}

type IWordsSegmenterFactory struct {
	win32.IInspectable
}

func (this *IWordsSegmenterFactory) Vtbl() *IWordsSegmenterFactoryVtbl {
	return (*IWordsSegmenterFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWordsSegmenterFactory) CreateWithLanguage(language string) *IWordsSegmenter {
	var _result *IWordsSegmenter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithLanguage, uintptr(unsafe.Pointer(this)), NewHStr(language).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type AlternateWordForm struct {
	RtClass
	*IAlternateWordForm
}

type SelectableWordSegment struct {
	RtClass
	*ISelectableWordSegment
}

type SelectableWordsSegmenter struct {
	RtClass
	*ISelectableWordsSegmenter
}

func NewSelectableWordsSegmenter_CreateWithLanguage(language string) *SelectableWordsSegmenter {
	hs := NewHStr("Windows.Data.Text.SelectableWordsSegmenter")
	var pFac *ISelectableWordsSegmenterFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISelectableWordsSegmenterFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISelectableWordsSegmenter
	p = pFac.CreateWithLanguage(language)
	result := &SelectableWordsSegmenter{
		RtClass:                   RtClass{PInspect: &p.IInspectable},
		ISelectableWordsSegmenter: p,
	}
	com.AddToScope(result)
	return result
}

type TextPhoneme struct {
	RtClass
	*ITextPhoneme
}

type UnicodeCharacters struct {
	RtClass
}

func NewIUnicodeCharactersStatics() *IUnicodeCharactersStatics {
	var p *IUnicodeCharactersStatics
	hs := NewHStr("Windows.Data.Text.UnicodeCharacters")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IUnicodeCharactersStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type WordSegment struct {
	RtClass
	*IWordSegment
}

type WordsSegmenter struct {
	RtClass
	*IWordsSegmenter
}

func NewWordsSegmenter_CreateWithLanguage(language string) *WordsSegmenter {
	hs := NewHStr("Windows.Data.Text.WordsSegmenter")
	var pFac *IWordsSegmenterFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWordsSegmenterFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IWordsSegmenter
	p = pFac.CreateWithLanguage(language)
	result := &WordsSegmenter{
		RtClass:         RtClass{PInspect: &p.IInspectable},
		IWordsSegmenter: p,
	}
	com.AddToScope(result)
	return result
}
