package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"log"
	"syscall"
	"unsafe"
)

// interfaces

// B12E5C3A-B76D-459B-BEEB-901151CD77D1
var IID_ILanguageFont = syscall.GUID{0xB12E5C3A, 0xB76D, 0x459B,
	[8]byte{0xBE, 0xEB, 0x90, 0x11, 0x51, 0xCD, 0x77, 0xD1}}

type ILanguageFontInterface interface {
	win32.IInspectableInterface
	Get_FontFamily() string
	Get_FontWeight() unsafe.Pointer
	Get_FontStretch() unsafe.Pointer
	Get_FontStyle() unsafe.Pointer
	Get_ScaleFactor() float64
}

type ILanguageFontVtbl struct {
	win32.IInspectableVtbl
	Get_FontFamily  uintptr
	Get_FontWeight  uintptr
	Get_FontStretch uintptr
	Get_FontStyle   uintptr
	Get_ScaleFactor uintptr
}

type ILanguageFont struct {
	win32.IInspectable
}

func (this *ILanguageFont) Vtbl() *ILanguageFontVtbl {
	return (*ILanguageFontVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILanguageFont) Get_FontFamily() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FontFamily, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILanguageFont) Get_FontWeight() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FontWeight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILanguageFont) Get_FontStretch() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FontStretch, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILanguageFont) Get_FontStyle() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FontStyle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILanguageFont) Get_ScaleFactor() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScaleFactor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F33A7FC3-3A5C-4AEA-B9FF-B39FB242F7F6
var IID_ILanguageFontGroup = syscall.GUID{0xF33A7FC3, 0x3A5C, 0x4AEA,
	[8]byte{0xB9, 0xFF, 0xB3, 0x9F, 0xB2, 0x42, 0xF7, 0xF6}}

type ILanguageFontGroupInterface interface {
	win32.IInspectableInterface
	Get_UITextFont() *ILanguageFont
	Get_UIHeadingFont() *ILanguageFont
	Get_UITitleFont() *ILanguageFont
	Get_UICaptionFont() *ILanguageFont
	Get_UINotificationHeadingFont() *ILanguageFont
	Get_TraditionalDocumentFont() *ILanguageFont
	Get_ModernDocumentFont() *ILanguageFont
	Get_DocumentHeadingFont() *ILanguageFont
	Get_FixedWidthTextFont() *ILanguageFont
	Get_DocumentAlternate1Font() *ILanguageFont
	Get_DocumentAlternate2Font() *ILanguageFont
}

type ILanguageFontGroupVtbl struct {
	win32.IInspectableVtbl
	Get_UITextFont                uintptr
	Get_UIHeadingFont             uintptr
	Get_UITitleFont               uintptr
	Get_UICaptionFont             uintptr
	Get_UINotificationHeadingFont uintptr
	Get_TraditionalDocumentFont   uintptr
	Get_ModernDocumentFont        uintptr
	Get_DocumentHeadingFont       uintptr
	Get_FixedWidthTextFont        uintptr
	Get_DocumentAlternate1Font    uintptr
	Get_DocumentAlternate2Font    uintptr
}

type ILanguageFontGroup struct {
	win32.IInspectable
}

func (this *ILanguageFontGroup) Vtbl() *ILanguageFontGroupVtbl {
	return (*ILanguageFontGroupVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILanguageFontGroup) Get_UITextFont() *ILanguageFont {
	var _result *ILanguageFont
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UITextFont, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILanguageFontGroup) Get_UIHeadingFont() *ILanguageFont {
	var _result *ILanguageFont
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UIHeadingFont, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILanguageFontGroup) Get_UITitleFont() *ILanguageFont {
	var _result *ILanguageFont
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UITitleFont, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILanguageFontGroup) Get_UICaptionFont() *ILanguageFont {
	var _result *ILanguageFont
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UICaptionFont, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILanguageFontGroup) Get_UINotificationHeadingFont() *ILanguageFont {
	var _result *ILanguageFont
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UINotificationHeadingFont, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILanguageFontGroup) Get_TraditionalDocumentFont() *ILanguageFont {
	var _result *ILanguageFont
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TraditionalDocumentFont, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILanguageFontGroup) Get_ModernDocumentFont() *ILanguageFont {
	var _result *ILanguageFont
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ModernDocumentFont, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILanguageFontGroup) Get_DocumentHeadingFont() *ILanguageFont {
	var _result *ILanguageFont
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DocumentHeadingFont, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILanguageFontGroup) Get_FixedWidthTextFont() *ILanguageFont {
	var _result *ILanguageFont
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FixedWidthTextFont, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILanguageFontGroup) Get_DocumentAlternate1Font() *ILanguageFont {
	var _result *ILanguageFont
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DocumentAlternate1Font, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILanguageFontGroup) Get_DocumentAlternate2Font() *ILanguageFont {
	var _result *ILanguageFont
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DocumentAlternate2Font, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FCAEAC67-4E77-49C7-B856-DDE934FC735B
var IID_ILanguageFontGroupFactory = syscall.GUID{0xFCAEAC67, 0x4E77, 0x49C7,
	[8]byte{0xB8, 0x56, 0xDD, 0xE9, 0x34, 0xFC, 0x73, 0x5B}}

type ILanguageFontGroupFactoryInterface interface {
	win32.IInspectableInterface
	CreateLanguageFontGroup(languageTag string) *ILanguageFontGroup
}

type ILanguageFontGroupFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateLanguageFontGroup uintptr
}

type ILanguageFontGroupFactory struct {
	win32.IInspectable
}

func (this *ILanguageFontGroupFactory) Vtbl() *ILanguageFontGroupFactoryVtbl {
	return (*ILanguageFontGroupFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILanguageFontGroupFactory) CreateLanguageFontGroup(languageTag string) *ILanguageFontGroup {
	var _result *ILanguageFontGroup
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateLanguageFontGroup, uintptr(unsafe.Pointer(this)), NewHStr(languageTag).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type LanguageFont struct {
	RtClass
	*ILanguageFont
}

type LanguageFontGroup struct {
	RtClass
	*ILanguageFontGroup
}

func NewLanguageFontGroup_CreateLanguageFontGroup(languageTag string) *LanguageFontGroup {
	hs := NewHStr("Windows.Globalization.Fonts.LanguageFontGroup")
	var pFac *ILanguageFontGroupFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ILanguageFontGroupFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ILanguageFontGroup
	p = pFac.CreateLanguageFontGroup(languageTag)
	result := &LanguageFontGroup{
		RtClass:            RtClass{PInspect: &p.IInspectable},
		ILanguageFontGroup: p,
	}
	com.AddToScope(result)
	return result
}
