package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"log"
	"syscall"
	"unsafe"
)

// interfaces

// FAE761BB-805D-4BB0-95BB-C1F7C3E8EB8E
var IID_ICharacterGrouping = syscall.GUID{0xFAE761BB, 0x805D, 0x4BB0,
	[8]byte{0x95, 0xBB, 0xC1, 0xF7, 0xC3, 0xE8, 0xEB, 0x8E}}

type ICharacterGroupingInterface interface {
	win32.IInspectableInterface
	Get_First() string
	Get_Label() string
}

type ICharacterGroupingVtbl struct {
	win32.IInspectableVtbl
	Get_First uintptr
	Get_Label uintptr
}

type ICharacterGrouping struct {
	win32.IInspectable
}

func (this *ICharacterGrouping) Vtbl() *ICharacterGroupingVtbl {
	return (*ICharacterGroupingVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICharacterGrouping) Get_First() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_First, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICharacterGrouping) Get_Label() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Label, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// B8D20A75-D4CF-4055-80E5-CE169C226496
var IID_ICharacterGroupings = syscall.GUID{0xB8D20A75, 0xD4CF, 0x4055,
	[8]byte{0x80, 0xE5, 0xCE, 0x16, 0x9C, 0x22, 0x64, 0x96}}

type ICharacterGroupingsInterface interface {
	win32.IInspectableInterface
	Lookup(text string) string
}

type ICharacterGroupingsVtbl struct {
	win32.IInspectableVtbl
	Lookup uintptr
}

type ICharacterGroupings struct {
	win32.IInspectable
}

func (this *ICharacterGroupings) Vtbl() *ICharacterGroupingsVtbl {
	return (*ICharacterGroupingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICharacterGroupings) Lookup(text string) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Lookup, uintptr(unsafe.Pointer(this)), NewHStr(text).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 99EA9FD9-886D-4401-9F98-69C82D4C2F78
var IID_ICharacterGroupingsFactory = syscall.GUID{0x99EA9FD9, 0x886D, 0x4401,
	[8]byte{0x9F, 0x98, 0x69, 0xC8, 0x2D, 0x4C, 0x2F, 0x78}}

type ICharacterGroupingsFactoryInterface interface {
	win32.IInspectableInterface
	Create(language string) *ICharacterGroupings
}

type ICharacterGroupingsFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ICharacterGroupingsFactory struct {
	win32.IInspectable
}

func (this *ICharacterGroupingsFactory) Vtbl() *ICharacterGroupingsFactoryVtbl {
	return (*ICharacterGroupingsFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICharacterGroupingsFactory) Create(language string) *ICharacterGroupings {
	var _result *ICharacterGroupings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(language).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type CharacterGrouping struct {
	RtClass
	*ICharacterGrouping
}

type CharacterGroupings struct {
	RtClass
	*ICharacterGroupings
}

func NewCharacterGroupings() *CharacterGroupings {
	hs := NewHStr("Windows.Globalization.Collation.CharacterGroupings")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &CharacterGroupings{
		RtClass:             RtClass{PInspect: p},
		ICharacterGroupings: (*ICharacterGroupings)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewCharacterGroupings_Create(language string) *CharacterGroupings {
	hs := NewHStr("Windows.Globalization.Collation.CharacterGroupings")
	var pFac *ICharacterGroupingsFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICharacterGroupingsFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ICharacterGroupings
	p = pFac.Create(language)
	result := &CharacterGroupings{
		RtClass:             RtClass{PInspect: &p.IInspectable},
		ICharacterGroupings: p,
	}
	com.AddToScope(result)
	return result
}
