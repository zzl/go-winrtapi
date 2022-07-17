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
type PrintBinding int32

const (
	PrintBinding_Default          PrintBinding = 0
	PrintBinding_NotAvailable     PrintBinding = 1
	PrintBinding_PrinterCustom    PrintBinding = 2
	PrintBinding_None             PrintBinding = 3
	PrintBinding_Bale             PrintBinding = 4
	PrintBinding_BindBottom       PrintBinding = 5
	PrintBinding_BindLeft         PrintBinding = 6
	PrintBinding_BindRight        PrintBinding = 7
	PrintBinding_BindTop          PrintBinding = 8
	PrintBinding_Booklet          PrintBinding = 9
	PrintBinding_EdgeStitchBottom PrintBinding = 10
	PrintBinding_EdgeStitchLeft   PrintBinding = 11
	PrintBinding_EdgeStitchRight  PrintBinding = 12
	PrintBinding_EdgeStitchTop    PrintBinding = 13
	PrintBinding_Fold             PrintBinding = 14
	PrintBinding_JogOffset        PrintBinding = 15
	PrintBinding_Trim             PrintBinding = 16
)

// enum
type PrintBordering int32

const (
	PrintBordering_Default       PrintBordering = 0
	PrintBordering_NotAvailable  PrintBordering = 1
	PrintBordering_PrinterCustom PrintBordering = 2
	PrintBordering_Bordered      PrintBordering = 3
	PrintBordering_Borderless    PrintBordering = 4
)

// enum
type PrintCollation int32

const (
	PrintCollation_Default       PrintCollation = 0
	PrintCollation_NotAvailable  PrintCollation = 1
	PrintCollation_PrinterCustom PrintCollation = 2
	PrintCollation_Collated      PrintCollation = 3
	PrintCollation_Uncollated    PrintCollation = 4
)

// enum
type PrintColorMode int32

const (
	PrintColorMode_Default       PrintColorMode = 0
	PrintColorMode_NotAvailable  PrintColorMode = 1
	PrintColorMode_PrinterCustom PrintColorMode = 2
	PrintColorMode_Color         PrintColorMode = 3
	PrintColorMode_Grayscale     PrintColorMode = 4
	PrintColorMode_Monochrome    PrintColorMode = 5
)

// enum
type PrintDuplex int32

const (
	PrintDuplex_Default           PrintDuplex = 0
	PrintDuplex_NotAvailable      PrintDuplex = 1
	PrintDuplex_PrinterCustom     PrintDuplex = 2
	PrintDuplex_OneSided          PrintDuplex = 3
	PrintDuplex_TwoSidedShortEdge PrintDuplex = 4
	PrintDuplex_TwoSidedLongEdge  PrintDuplex = 5
)

// enum
type PrintHolePunch int32

const (
	PrintHolePunch_Default       PrintHolePunch = 0
	PrintHolePunch_NotAvailable  PrintHolePunch = 1
	PrintHolePunch_PrinterCustom PrintHolePunch = 2
	PrintHolePunch_None          PrintHolePunch = 3
	PrintHolePunch_LeftEdge      PrintHolePunch = 4
	PrintHolePunch_RightEdge     PrintHolePunch = 5
	PrintHolePunch_TopEdge       PrintHolePunch = 6
	PrintHolePunch_BottomEdge    PrintHolePunch = 7
)

// enum
type PrintMediaSize int32

const (
	PrintMediaSize_Default                             PrintMediaSize = 0
	PrintMediaSize_NotAvailable                        PrintMediaSize = 1
	PrintMediaSize_PrinterCustom                       PrintMediaSize = 2
	PrintMediaSize_BusinessCard                        PrintMediaSize = 3
	PrintMediaSize_CreditCard                          PrintMediaSize = 4
	PrintMediaSize_IsoA0                               PrintMediaSize = 5
	PrintMediaSize_IsoA1                               PrintMediaSize = 6
	PrintMediaSize_IsoA10                              PrintMediaSize = 7
	PrintMediaSize_IsoA2                               PrintMediaSize = 8
	PrintMediaSize_IsoA3                               PrintMediaSize = 9
	PrintMediaSize_IsoA3Extra                          PrintMediaSize = 10
	PrintMediaSize_IsoA3Rotated                        PrintMediaSize = 11
	PrintMediaSize_IsoA4                               PrintMediaSize = 12
	PrintMediaSize_IsoA4Extra                          PrintMediaSize = 13
	PrintMediaSize_IsoA4Rotated                        PrintMediaSize = 14
	PrintMediaSize_IsoA5                               PrintMediaSize = 15
	PrintMediaSize_IsoA5Extra                          PrintMediaSize = 16
	PrintMediaSize_IsoA5Rotated                        PrintMediaSize = 17
	PrintMediaSize_IsoA6                               PrintMediaSize = 18
	PrintMediaSize_IsoA6Rotated                        PrintMediaSize = 19
	PrintMediaSize_IsoA7                               PrintMediaSize = 20
	PrintMediaSize_IsoA8                               PrintMediaSize = 21
	PrintMediaSize_IsoA9                               PrintMediaSize = 22
	PrintMediaSize_IsoB0                               PrintMediaSize = 23
	PrintMediaSize_IsoB1                               PrintMediaSize = 24
	PrintMediaSize_IsoB10                              PrintMediaSize = 25
	PrintMediaSize_IsoB2                               PrintMediaSize = 26
	PrintMediaSize_IsoB3                               PrintMediaSize = 27
	PrintMediaSize_IsoB4                               PrintMediaSize = 28
	PrintMediaSize_IsoB4Envelope                       PrintMediaSize = 29
	PrintMediaSize_IsoB5Envelope                       PrintMediaSize = 30
	PrintMediaSize_IsoB5Extra                          PrintMediaSize = 31
	PrintMediaSize_IsoB7                               PrintMediaSize = 32
	PrintMediaSize_IsoB8                               PrintMediaSize = 33
	PrintMediaSize_IsoB9                               PrintMediaSize = 34
	PrintMediaSize_IsoC0                               PrintMediaSize = 35
	PrintMediaSize_IsoC1                               PrintMediaSize = 36
	PrintMediaSize_IsoC10                              PrintMediaSize = 37
	PrintMediaSize_IsoC2                               PrintMediaSize = 38
	PrintMediaSize_IsoC3                               PrintMediaSize = 39
	PrintMediaSize_IsoC3Envelope                       PrintMediaSize = 40
	PrintMediaSize_IsoC4                               PrintMediaSize = 41
	PrintMediaSize_IsoC4Envelope                       PrintMediaSize = 42
	PrintMediaSize_IsoC5                               PrintMediaSize = 43
	PrintMediaSize_IsoC5Envelope                       PrintMediaSize = 44
	PrintMediaSize_IsoC6                               PrintMediaSize = 45
	PrintMediaSize_IsoC6C5Envelope                     PrintMediaSize = 46
	PrintMediaSize_IsoC6Envelope                       PrintMediaSize = 47
	PrintMediaSize_IsoC7                               PrintMediaSize = 48
	PrintMediaSize_IsoC8                               PrintMediaSize = 49
	PrintMediaSize_IsoC9                               PrintMediaSize = 50
	PrintMediaSize_IsoDLEnvelope                       PrintMediaSize = 51
	PrintMediaSize_IsoDLEnvelopeRotated                PrintMediaSize = 52
	PrintMediaSize_IsoSRA3                             PrintMediaSize = 53
	PrintMediaSize_Japan2LPhoto                        PrintMediaSize = 54
	PrintMediaSize_JapanChou3Envelope                  PrintMediaSize = 55
	PrintMediaSize_JapanChou3EnvelopeRotated           PrintMediaSize = 56
	PrintMediaSize_JapanChou4Envelope                  PrintMediaSize = 57
	PrintMediaSize_JapanChou4EnvelopeRotated           PrintMediaSize = 58
	PrintMediaSize_JapanDoubleHagakiPostcard           PrintMediaSize = 59
	PrintMediaSize_JapanDoubleHagakiPostcardRotated    PrintMediaSize = 60
	PrintMediaSize_JapanHagakiPostcard                 PrintMediaSize = 61
	PrintMediaSize_JapanHagakiPostcardRotated          PrintMediaSize = 62
	PrintMediaSize_JapanKaku2Envelope                  PrintMediaSize = 63
	PrintMediaSize_JapanKaku2EnvelopeRotated           PrintMediaSize = 64
	PrintMediaSize_JapanKaku3Envelope                  PrintMediaSize = 65
	PrintMediaSize_JapanKaku3EnvelopeRotated           PrintMediaSize = 66
	PrintMediaSize_JapanLPhoto                         PrintMediaSize = 67
	PrintMediaSize_JapanQuadrupleHagakiPostcard        PrintMediaSize = 68
	PrintMediaSize_JapanYou1Envelope                   PrintMediaSize = 69
	PrintMediaSize_JapanYou2Envelope                   PrintMediaSize = 70
	PrintMediaSize_JapanYou3Envelope                   PrintMediaSize = 71
	PrintMediaSize_JapanYou4Envelope                   PrintMediaSize = 72
	PrintMediaSize_JapanYou4EnvelopeRotated            PrintMediaSize = 73
	PrintMediaSize_JapanYou6Envelope                   PrintMediaSize = 74
	PrintMediaSize_JapanYou6EnvelopeRotated            PrintMediaSize = 75
	PrintMediaSize_JisB0                               PrintMediaSize = 76
	PrintMediaSize_JisB1                               PrintMediaSize = 77
	PrintMediaSize_JisB10                              PrintMediaSize = 78
	PrintMediaSize_JisB2                               PrintMediaSize = 79
	PrintMediaSize_JisB3                               PrintMediaSize = 80
	PrintMediaSize_JisB4                               PrintMediaSize = 81
	PrintMediaSize_JisB4Rotated                        PrintMediaSize = 82
	PrintMediaSize_JisB5                               PrintMediaSize = 83
	PrintMediaSize_JisB5Rotated                        PrintMediaSize = 84
	PrintMediaSize_JisB6                               PrintMediaSize = 85
	PrintMediaSize_JisB6Rotated                        PrintMediaSize = 86
	PrintMediaSize_JisB7                               PrintMediaSize = 87
	PrintMediaSize_JisB8                               PrintMediaSize = 88
	PrintMediaSize_JisB9                               PrintMediaSize = 89
	PrintMediaSize_NorthAmerica10x11                   PrintMediaSize = 90
	PrintMediaSize_NorthAmerica10x12                   PrintMediaSize = 91
	PrintMediaSize_NorthAmerica10x14                   PrintMediaSize = 92
	PrintMediaSize_NorthAmerica11x17                   PrintMediaSize = 93
	PrintMediaSize_NorthAmerica14x17                   PrintMediaSize = 94
	PrintMediaSize_NorthAmerica4x6                     PrintMediaSize = 95
	PrintMediaSize_NorthAmerica4x8                     PrintMediaSize = 96
	PrintMediaSize_NorthAmerica5x7                     PrintMediaSize = 97
	PrintMediaSize_NorthAmerica8x10                    PrintMediaSize = 98
	PrintMediaSize_NorthAmerica9x11                    PrintMediaSize = 99
	PrintMediaSize_NorthAmericaArchitectureASheet      PrintMediaSize = 100
	PrintMediaSize_NorthAmericaArchitectureBSheet      PrintMediaSize = 101
	PrintMediaSize_NorthAmericaArchitectureCSheet      PrintMediaSize = 102
	PrintMediaSize_NorthAmericaArchitectureDSheet      PrintMediaSize = 103
	PrintMediaSize_NorthAmericaArchitectureESheet      PrintMediaSize = 104
	PrintMediaSize_NorthAmericaCSheet                  PrintMediaSize = 105
	PrintMediaSize_NorthAmericaDSheet                  PrintMediaSize = 106
	PrintMediaSize_NorthAmericaESheet                  PrintMediaSize = 107
	PrintMediaSize_NorthAmericaExecutive               PrintMediaSize = 108
	PrintMediaSize_NorthAmericaGermanLegalFanfold      PrintMediaSize = 109
	PrintMediaSize_NorthAmericaGermanStandardFanfold   PrintMediaSize = 110
	PrintMediaSize_NorthAmericaLegal                   PrintMediaSize = 111
	PrintMediaSize_NorthAmericaLegalExtra              PrintMediaSize = 112
	PrintMediaSize_NorthAmericaLetter                  PrintMediaSize = 113
	PrintMediaSize_NorthAmericaLetterExtra             PrintMediaSize = 114
	PrintMediaSize_NorthAmericaLetterPlus              PrintMediaSize = 115
	PrintMediaSize_NorthAmericaLetterRotated           PrintMediaSize = 116
	PrintMediaSize_NorthAmericaMonarchEnvelope         PrintMediaSize = 117
	PrintMediaSize_NorthAmericaNote                    PrintMediaSize = 118
	PrintMediaSize_NorthAmericaNumber10Envelope        PrintMediaSize = 119
	PrintMediaSize_NorthAmericaNumber10EnvelopeRotated PrintMediaSize = 120
	PrintMediaSize_NorthAmericaNumber11Envelope        PrintMediaSize = 121
	PrintMediaSize_NorthAmericaNumber12Envelope        PrintMediaSize = 122
	PrintMediaSize_NorthAmericaNumber14Envelope        PrintMediaSize = 123
	PrintMediaSize_NorthAmericaNumber9Envelope         PrintMediaSize = 124
	PrintMediaSize_NorthAmericaPersonalEnvelope        PrintMediaSize = 125
	PrintMediaSize_NorthAmericaQuarto                  PrintMediaSize = 126
	PrintMediaSize_NorthAmericaStatement               PrintMediaSize = 127
	PrintMediaSize_NorthAmericaSuperA                  PrintMediaSize = 128
	PrintMediaSize_NorthAmericaSuperB                  PrintMediaSize = 129
	PrintMediaSize_NorthAmericaTabloid                 PrintMediaSize = 130
	PrintMediaSize_NorthAmericaTabloidExtra            PrintMediaSize = 131
	PrintMediaSize_OtherMetricA3Plus                   PrintMediaSize = 132
	PrintMediaSize_OtherMetricA4Plus                   PrintMediaSize = 133
	PrintMediaSize_OtherMetricFolio                    PrintMediaSize = 134
	PrintMediaSize_OtherMetricInviteEnvelope           PrintMediaSize = 135
	PrintMediaSize_OtherMetricItalianEnvelope          PrintMediaSize = 136
	PrintMediaSize_Prc10Envelope                       PrintMediaSize = 137
	PrintMediaSize_Prc10EnvelopeRotated                PrintMediaSize = 138
	PrintMediaSize_Prc16K                              PrintMediaSize = 139
	PrintMediaSize_Prc16KRotated                       PrintMediaSize = 140
	PrintMediaSize_Prc1Envelope                        PrintMediaSize = 141
	PrintMediaSize_Prc1EnvelopeRotated                 PrintMediaSize = 142
	PrintMediaSize_Prc2Envelope                        PrintMediaSize = 143
	PrintMediaSize_Prc2EnvelopeRotated                 PrintMediaSize = 144
	PrintMediaSize_Prc32K                              PrintMediaSize = 145
	PrintMediaSize_Prc32KBig                           PrintMediaSize = 146
	PrintMediaSize_Prc32KRotated                       PrintMediaSize = 147
	PrintMediaSize_Prc3Envelope                        PrintMediaSize = 148
	PrintMediaSize_Prc3EnvelopeRotated                 PrintMediaSize = 149
	PrintMediaSize_Prc4Envelope                        PrintMediaSize = 150
	PrintMediaSize_Prc4EnvelopeRotated                 PrintMediaSize = 151
	PrintMediaSize_Prc5Envelope                        PrintMediaSize = 152
	PrintMediaSize_Prc5EnvelopeRotated                 PrintMediaSize = 153
	PrintMediaSize_Prc6Envelope                        PrintMediaSize = 154
	PrintMediaSize_Prc6EnvelopeRotated                 PrintMediaSize = 155
	PrintMediaSize_Prc7Envelope                        PrintMediaSize = 156
	PrintMediaSize_Prc7EnvelopeRotated                 PrintMediaSize = 157
	PrintMediaSize_Prc8Envelope                        PrintMediaSize = 158
	PrintMediaSize_Prc8EnvelopeRotated                 PrintMediaSize = 159
	PrintMediaSize_Prc9Envelope                        PrintMediaSize = 160
	PrintMediaSize_Prc9EnvelopeRotated                 PrintMediaSize = 161
	PrintMediaSize_Roll04Inch                          PrintMediaSize = 162
	PrintMediaSize_Roll06Inch                          PrintMediaSize = 163
	PrintMediaSize_Roll08Inch                          PrintMediaSize = 164
	PrintMediaSize_Roll12Inch                          PrintMediaSize = 165
	PrintMediaSize_Roll15Inch                          PrintMediaSize = 166
	PrintMediaSize_Roll18Inch                          PrintMediaSize = 167
	PrintMediaSize_Roll22Inch                          PrintMediaSize = 168
	PrintMediaSize_Roll24Inch                          PrintMediaSize = 169
	PrintMediaSize_Roll30Inch                          PrintMediaSize = 170
	PrintMediaSize_Roll36Inch                          PrintMediaSize = 171
	PrintMediaSize_Roll54Inch                          PrintMediaSize = 172
)

// enum
type PrintMediaType int32

const (
	PrintMediaType_Default               PrintMediaType = 0
	PrintMediaType_NotAvailable          PrintMediaType = 1
	PrintMediaType_PrinterCustom         PrintMediaType = 2
	PrintMediaType_AutoSelect            PrintMediaType = 3
	PrintMediaType_Archival              PrintMediaType = 4
	PrintMediaType_BackPrintFilm         PrintMediaType = 5
	PrintMediaType_Bond                  PrintMediaType = 6
	PrintMediaType_CardStock             PrintMediaType = 7
	PrintMediaType_Continuous            PrintMediaType = 8
	PrintMediaType_EnvelopePlain         PrintMediaType = 9
	PrintMediaType_EnvelopeWindow        PrintMediaType = 10
	PrintMediaType_Fabric                PrintMediaType = 11
	PrintMediaType_HighResolution        PrintMediaType = 12
	PrintMediaType_Label                 PrintMediaType = 13
	PrintMediaType_MultiLayerForm        PrintMediaType = 14
	PrintMediaType_MultiPartForm         PrintMediaType = 15
	PrintMediaType_Photographic          PrintMediaType = 16
	PrintMediaType_PhotographicFilm      PrintMediaType = 17
	PrintMediaType_PhotographicGlossy    PrintMediaType = 18
	PrintMediaType_PhotographicHighGloss PrintMediaType = 19
	PrintMediaType_PhotographicMatte     PrintMediaType = 20
	PrintMediaType_PhotographicSatin     PrintMediaType = 21
	PrintMediaType_PhotographicSemiGloss PrintMediaType = 22
	PrintMediaType_Plain                 PrintMediaType = 23
	PrintMediaType_Screen                PrintMediaType = 24
	PrintMediaType_ScreenPaged           PrintMediaType = 25
	PrintMediaType_Stationery            PrintMediaType = 26
	PrintMediaType_TabStockFull          PrintMediaType = 27
	PrintMediaType_TabStockPreCut        PrintMediaType = 28
	PrintMediaType_Transparency          PrintMediaType = 29
	PrintMediaType_TShirtTransfer        PrintMediaType = 30
	PrintMediaType_None                  PrintMediaType = 31
)

// enum
type PrintOrientation int32

const (
	PrintOrientation_Default          PrintOrientation = 0
	PrintOrientation_NotAvailable     PrintOrientation = 1
	PrintOrientation_PrinterCustom    PrintOrientation = 2
	PrintOrientation_Portrait         PrintOrientation = 3
	PrintOrientation_PortraitFlipped  PrintOrientation = 4
	PrintOrientation_Landscape        PrintOrientation = 5
	PrintOrientation_LandscapeFlipped PrintOrientation = 6
)

// enum
type PrintQuality int32

const (
	PrintQuality_Default       PrintQuality = 0
	PrintQuality_NotAvailable  PrintQuality = 1
	PrintQuality_PrinterCustom PrintQuality = 2
	PrintQuality_Automatic     PrintQuality = 3
	PrintQuality_Draft         PrintQuality = 4
	PrintQuality_Fax           PrintQuality = 5
	PrintQuality_High          PrintQuality = 6
	PrintQuality_Normal        PrintQuality = 7
	PrintQuality_Photographic  PrintQuality = 8
	PrintQuality_Text          PrintQuality = 9
)

// enum
type PrintStaple int32

const (
	PrintStaple_Default           PrintStaple = 0
	PrintStaple_NotAvailable      PrintStaple = 1
	PrintStaple_PrinterCustom     PrintStaple = 2
	PrintStaple_None              PrintStaple = 3
	PrintStaple_StapleTopLeft     PrintStaple = 4
	PrintStaple_StapleTopRight    PrintStaple = 5
	PrintStaple_StapleBottomLeft  PrintStaple = 6
	PrintStaple_StapleBottomRight PrintStaple = 7
	PrintStaple_StapleDualLeft    PrintStaple = 8
	PrintStaple_StapleDualRight   PrintStaple = 9
	PrintStaple_StapleDualTop     PrintStaple = 10
	PrintStaple_StapleDualBottom  PrintStaple = 11
	PrintStaple_SaddleStitch      PrintStaple = 12
)

// enum
type PrintTaskCompletion int32

const (
	PrintTaskCompletion_Abandoned PrintTaskCompletion = 0
	PrintTaskCompletion_Canceled  PrintTaskCompletion = 1
	PrintTaskCompletion_Failed    PrintTaskCompletion = 2
	PrintTaskCompletion_Submitted PrintTaskCompletion = 3
)

// structs

type PrintPageDescription struct {
	PageSize      Size
	ImageableRect Rect
	DpiX          uint32
	DpiY          uint32
}

// func types

//6C109FA8-5CB6-4B3A-8663-F39CB02DC9B4
type PrintTaskSourceRequestedHandler func(args *IPrintTaskSourceRequestedArgs) com.Error

// interfaces

// DEDC0C30-F1EB-47DF-AAE6-ED5427511F01
var IID_IPrintDocumentSource = syscall.GUID{0xDEDC0C30, 0xF1EB, 0x47DF,
	[8]byte{0xAA, 0xE6, 0xED, 0x54, 0x27, 0x51, 0x1F, 0x01}}

type IPrintDocumentSourceInterface interface {
	win32.IInspectableInterface
}

type IPrintDocumentSourceVtbl struct {
	win32.IInspectableVtbl
}

type IPrintDocumentSource struct {
	win32.IInspectable
}

func (this *IPrintDocumentSource) Vtbl() *IPrintDocumentSourceVtbl {
	return (*IPrintDocumentSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// FF2A9694-8C99-44FD-AE4A-19D9AA9A0F0A
var IID_IPrintManager = syscall.GUID{0xFF2A9694, 0x8C99, 0x44FD,
	[8]byte{0xAE, 0x4A, 0x19, 0xD9, 0xAA, 0x9A, 0x0F, 0x0A}}

type IPrintManagerInterface interface {
	win32.IInspectableInterface
	Add_PrintTaskRequested(eventHandler TypedEventHandler[*IPrintManager, *IPrintTaskRequestedEventArgs]) EventRegistrationToken
	Remove_PrintTaskRequested(eventCookie EventRegistrationToken)
}

type IPrintManagerVtbl struct {
	win32.IInspectableVtbl
	Add_PrintTaskRequested    uintptr
	Remove_PrintTaskRequested uintptr
}

type IPrintManager struct {
	win32.IInspectable
}

func (this *IPrintManager) Vtbl() *IPrintManagerVtbl {
	return (*IPrintManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrintManager) Add_PrintTaskRequested(eventHandler TypedEventHandler[*IPrintManager, *IPrintTaskRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PrintTaskRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(eventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintManager) Remove_PrintTaskRequested(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PrintTaskRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

// 58185DCD-E634-4654-84F0-E0152A8217AC
var IID_IPrintManagerStatic = syscall.GUID{0x58185DCD, 0xE634, 0x4654,
	[8]byte{0x84, 0xF0, 0xE0, 0x15, 0x2A, 0x82, 0x17, 0xAC}}

type IPrintManagerStaticInterface interface {
	win32.IInspectableInterface
	GetForCurrentView() *IPrintManager
	ShowPrintUIAsync() *IAsyncOperation[bool]
}

type IPrintManagerStaticVtbl struct {
	win32.IInspectableVtbl
	GetForCurrentView uintptr
	ShowPrintUIAsync  uintptr
}

type IPrintManagerStatic struct {
	win32.IInspectable
}

func (this *IPrintManagerStatic) Vtbl() *IPrintManagerStaticVtbl {
	return (*IPrintManagerStaticVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrintManagerStatic) GetForCurrentView() *IPrintManager {
	var _result *IPrintManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrintManagerStatic) ShowPrintUIAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowPrintUIAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 35A99955-E6AB-4139-9ABD-B86A729B3598
var IID_IPrintManagerStatic2 = syscall.GUID{0x35A99955, 0xE6AB, 0x4139,
	[8]byte{0x9A, 0xBD, 0xB8, 0x6A, 0x72, 0x9B, 0x35, 0x98}}

type IPrintManagerStatic2Interface interface {
	win32.IInspectableInterface
	IsSupported() bool
}

type IPrintManagerStatic2Vtbl struct {
	win32.IInspectableVtbl
	IsSupported uintptr
}

type IPrintManagerStatic2 struct {
	win32.IInspectable
}

func (this *IPrintManagerStatic2) Vtbl() *IPrintManagerStatic2Vtbl {
	return (*IPrintManagerStatic2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrintManagerStatic2) IsSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// DD4BE9C9-A6A1-4ADA-930E-DA872A4F23D3
var IID_IPrintPageInfo = syscall.GUID{0xDD4BE9C9, 0xA6A1, 0x4ADA,
	[8]byte{0x93, 0x0E, 0xDA, 0x87, 0x2A, 0x4F, 0x23, 0xD3}}

type IPrintPageInfoInterface interface {
	win32.IInspectableInterface
	Put_MediaSize(value PrintMediaSize)
	Get_MediaSize() PrintMediaSize
	Put_PageSize(value Size)
	Get_PageSize() Size
	Put_DpiX(value uint32)
	Get_DpiX() uint32
	Put_DpiY(value uint32)
	Get_DpiY() uint32
	Put_Orientation(value PrintOrientation)
	Get_Orientation() PrintOrientation
}

type IPrintPageInfoVtbl struct {
	win32.IInspectableVtbl
	Put_MediaSize   uintptr
	Get_MediaSize   uintptr
	Put_PageSize    uintptr
	Get_PageSize    uintptr
	Put_DpiX        uintptr
	Get_DpiX        uintptr
	Put_DpiY        uintptr
	Get_DpiY        uintptr
	Put_Orientation uintptr
	Get_Orientation uintptr
}

type IPrintPageInfo struct {
	win32.IInspectable
}

func (this *IPrintPageInfo) Vtbl() *IPrintPageInfoVtbl {
	return (*IPrintPageInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrintPageInfo) Put_MediaSize(value PrintMediaSize) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MediaSize, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrintPageInfo) Get_MediaSize() PrintMediaSize {
	var _result PrintMediaSize
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintPageInfo) Put_PageSize(value Size) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PageSize, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IPrintPageInfo) Get_PageSize() Size {
	var _result Size
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PageSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintPageInfo) Put_DpiX(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DpiX, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrintPageInfo) Get_DpiX() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DpiX, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintPageInfo) Put_DpiY(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DpiY, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrintPageInfo) Get_DpiY() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DpiY, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintPageInfo) Put_Orientation(value PrintOrientation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Orientation, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrintPageInfo) Get_Orientation() PrintOrientation {
	var _result PrintOrientation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Orientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F8A06C54-6E7C-51C5-57FD-0660C2D71513
var IID_IPrintPageRange = syscall.GUID{0xF8A06C54, 0x6E7C, 0x51C5,
	[8]byte{0x57, 0xFD, 0x06, 0x60, 0xC2, 0xD7, 0x15, 0x13}}

type IPrintPageRangeInterface interface {
	win32.IInspectableInterface
	Get_FirstPageNumber() int32
	Get_LastPageNumber() int32
}

type IPrintPageRangeVtbl struct {
	win32.IInspectableVtbl
	Get_FirstPageNumber uintptr
	Get_LastPageNumber  uintptr
}

type IPrintPageRange struct {
	win32.IInspectable
}

func (this *IPrintPageRange) Vtbl() *IPrintPageRangeVtbl {
	return (*IPrintPageRangeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrintPageRange) Get_FirstPageNumber() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FirstPageNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintPageRange) Get_LastPageNumber() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LastPageNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 408FD45F-E047-5F85-7129-FB085A4FAD14
var IID_IPrintPageRangeFactory = syscall.GUID{0x408FD45F, 0xE047, 0x5F85,
	[8]byte{0x71, 0x29, 0xFB, 0x08, 0x5A, 0x4F, 0xAD, 0x14}}

type IPrintPageRangeFactoryInterface interface {
	win32.IInspectableInterface
	Create(firstPage int32, lastPage int32) *IPrintPageRange
	CreateWithSinglePage(page int32) *IPrintPageRange
}

type IPrintPageRangeFactoryVtbl struct {
	win32.IInspectableVtbl
	Create               uintptr
	CreateWithSinglePage uintptr
}

type IPrintPageRangeFactory struct {
	win32.IInspectable
}

func (this *IPrintPageRangeFactory) Vtbl() *IPrintPageRangeFactoryVtbl {
	return (*IPrintPageRangeFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrintPageRangeFactory) Create(firstPage int32, lastPage int32) *IPrintPageRange {
	var _result *IPrintPageRange
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(firstPage), uintptr(lastPage), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrintPageRangeFactory) CreateWithSinglePage(page int32) *IPrintPageRange {
	var _result *IPrintPageRange
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithSinglePage, uintptr(unsafe.Pointer(this)), uintptr(page), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CE6DB728-1357-46B2-A923-79F995F448FC
var IID_IPrintPageRangeOptions = syscall.GUID{0xCE6DB728, 0x1357, 0x46B2,
	[8]byte{0xA9, 0x23, 0x79, 0xF9, 0x95, 0xF4, 0x48, 0xFC}}

type IPrintPageRangeOptionsInterface interface {
	win32.IInspectableInterface
	Put_AllowAllPages(value bool)
	Get_AllowAllPages() bool
	Put_AllowCurrentPage(value bool)
	Get_AllowCurrentPage() bool
	Put_AllowCustomSetOfPages(value bool)
	Get_AllowCustomSetOfPages() bool
}

type IPrintPageRangeOptionsVtbl struct {
	win32.IInspectableVtbl
	Put_AllowAllPages         uintptr
	Get_AllowAllPages         uintptr
	Put_AllowCurrentPage      uintptr
	Get_AllowCurrentPage      uintptr
	Put_AllowCustomSetOfPages uintptr
	Get_AllowCustomSetOfPages uintptr
}

type IPrintPageRangeOptions struct {
	win32.IInspectable
}

func (this *IPrintPageRangeOptions) Vtbl() *IPrintPageRangeOptionsVtbl {
	return (*IPrintPageRangeOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrintPageRangeOptions) Put_AllowAllPages(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AllowAllPages, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPrintPageRangeOptions) Get_AllowAllPages() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllowAllPages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintPageRangeOptions) Put_AllowCurrentPage(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AllowCurrentPage, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPrintPageRangeOptions) Get_AllowCurrentPage() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllowCurrentPage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintPageRangeOptions) Put_AllowCustomSetOfPages(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AllowCustomSetOfPages, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPrintPageRangeOptions) Get_AllowCustomSetOfPages() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllowCustomSetOfPages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 61D80247-6CF6-4FAD-84E2-A5E82E2D4CEB
var IID_IPrintTask = syscall.GUID{0x61D80247, 0x6CF6, 0x4FAD,
	[8]byte{0x84, 0xE2, 0xA5, 0xE8, 0x2E, 0x2D, 0x4C, 0xEB}}

type IPrintTaskInterface interface {
	win32.IInspectableInterface
	Get_Properties() *IDataPackagePropertySet
	Get_Source() *IPrintDocumentSource
	Get_Options() *IPrintTaskOptionsCore
	Add_Previewing(eventHandler TypedEventHandler[*IPrintTask, interface{}]) EventRegistrationToken
	Remove_Previewing(eventCookie EventRegistrationToken)
	Add_Submitting(eventHandler TypedEventHandler[*IPrintTask, interface{}]) EventRegistrationToken
	Remove_Submitting(eventCookie EventRegistrationToken)
	Add_Progressing(eventHandler TypedEventHandler[*IPrintTask, *IPrintTaskProgressingEventArgs]) EventRegistrationToken
	Remove_Progressing(eventCookie EventRegistrationToken)
	Add_Completed(eventHandler TypedEventHandler[*IPrintTask, *IPrintTaskCompletedEventArgs]) EventRegistrationToken
	Remove_Completed(eventCookie EventRegistrationToken)
}

type IPrintTaskVtbl struct {
	win32.IInspectableVtbl
	Get_Properties     uintptr
	Get_Source         uintptr
	Get_Options        uintptr
	Add_Previewing     uintptr
	Remove_Previewing  uintptr
	Add_Submitting     uintptr
	Remove_Submitting  uintptr
	Add_Progressing    uintptr
	Remove_Progressing uintptr
	Add_Completed      uintptr
	Remove_Completed   uintptr
}

type IPrintTask struct {
	win32.IInspectable
}

func (this *IPrintTask) Vtbl() *IPrintTaskVtbl {
	return (*IPrintTaskVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrintTask) Get_Properties() *IDataPackagePropertySet {
	var _result *IDataPackagePropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrintTask) Get_Source() *IPrintDocumentSource {
	var _result *IPrintDocumentSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrintTask) Get_Options() *IPrintTaskOptionsCore {
	var _result *IPrintTaskOptionsCore
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Options, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrintTask) Add_Previewing(eventHandler TypedEventHandler[*IPrintTask, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Previewing, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(eventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintTask) Remove_Previewing(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Previewing, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

func (this *IPrintTask) Add_Submitting(eventHandler TypedEventHandler[*IPrintTask, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Submitting, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(eventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintTask) Remove_Submitting(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Submitting, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

func (this *IPrintTask) Add_Progressing(eventHandler TypedEventHandler[*IPrintTask, *IPrintTaskProgressingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Progressing, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(eventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintTask) Remove_Progressing(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Progressing, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

func (this *IPrintTask) Add_Completed(eventHandler TypedEventHandler[*IPrintTask, *IPrintTaskCompletedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Completed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(eventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintTask) Remove_Completed(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Completed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

// 36234877-3E53-4D9D-8F5E-316AC8DEDAE1
var IID_IPrintTask2 = syscall.GUID{0x36234877, 0x3E53, 0x4D9D,
	[8]byte{0x8F, 0x5E, 0x31, 0x6A, 0xC8, 0xDE, 0xDA, 0xE1}}

type IPrintTask2Interface interface {
	win32.IInspectableInterface
	Put_IsPreviewEnabled(value bool)
	Get_IsPreviewEnabled() bool
}

type IPrintTask2Vtbl struct {
	win32.IInspectableVtbl
	Put_IsPreviewEnabled uintptr
	Get_IsPreviewEnabled uintptr
}

type IPrintTask2 struct {
	win32.IInspectable
}

func (this *IPrintTask2) Vtbl() *IPrintTask2Vtbl {
	return (*IPrintTask2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrintTask2) Put_IsPreviewEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsPreviewEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPrintTask2) Get_IsPreviewEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPreviewEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5BCD34AF-24E9-4C10-8D07-14C346BA3FCE
var IID_IPrintTaskCompletedEventArgs = syscall.GUID{0x5BCD34AF, 0x24E9, 0x4C10,
	[8]byte{0x8D, 0x07, 0x14, 0xC3, 0x46, 0xBA, 0x3F, 0xCE}}

type IPrintTaskCompletedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Completion() PrintTaskCompletion
}

type IPrintTaskCompletedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Completion uintptr
}

type IPrintTaskCompletedEventArgs struct {
	win32.IInspectable
}

func (this *IPrintTaskCompletedEventArgs) Vtbl() *IPrintTaskCompletedEventArgsVtbl {
	return (*IPrintTaskCompletedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrintTaskCompletedEventArgs) Get_Completion() PrintTaskCompletion {
	var _result PrintTaskCompletion
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Completion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5A0A66BB-D289-41BB-96DD-57E28338AE3F
var IID_IPrintTaskOptions = syscall.GUID{0x5A0A66BB, 0xD289, 0x41BB,
	[8]byte{0x96, 0xDD, 0x57, 0xE2, 0x83, 0x38, 0xAE, 0x3F}}

type IPrintTaskOptionsInterface interface {
	win32.IInspectableInterface
	Put_Bordering(value PrintBordering)
	Get_Bordering() PrintBordering
	GetPagePrintTicket(printPageInfo *IPrintPageInfo) *IRandomAccessStream
}

type IPrintTaskOptionsVtbl struct {
	win32.IInspectableVtbl
	Put_Bordering      uintptr
	Get_Bordering      uintptr
	GetPagePrintTicket uintptr
}

type IPrintTaskOptions struct {
	win32.IInspectable
}

func (this *IPrintTaskOptions) Vtbl() *IPrintTaskOptionsVtbl {
	return (*IPrintTaskOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrintTaskOptions) Put_Bordering(value PrintBordering) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Bordering, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrintTaskOptions) Get_Bordering() PrintBordering {
	var _result PrintBordering
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bordering, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintTaskOptions) GetPagePrintTicket(printPageInfo *IPrintPageInfo) *IRandomAccessStream {
	var _result *IRandomAccessStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPagePrintTicket, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(printPageInfo)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EB9B1606-9A36-4B59-8617-B217849262E1
var IID_IPrintTaskOptions2 = syscall.GUID{0xEB9B1606, 0x9A36, 0x4B59,
	[8]byte{0x86, 0x17, 0xB2, 0x17, 0x84, 0x92, 0x62, 0xE1}}

type IPrintTaskOptions2Interface interface {
	win32.IInspectableInterface
	Get_PageRangeOptions() *IPrintPageRangeOptions
	Get_CustomPageRanges() *IVector[*IPrintPageRange]
}

type IPrintTaskOptions2Vtbl struct {
	win32.IInspectableVtbl
	Get_PageRangeOptions uintptr
	Get_CustomPageRanges uintptr
}

type IPrintTaskOptions2 struct {
	win32.IInspectable
}

func (this *IPrintTaskOptions2) Vtbl() *IPrintTaskOptions2Vtbl {
	return (*IPrintTaskOptions2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrintTaskOptions2) Get_PageRangeOptions() *IPrintPageRangeOptions {
	var _result *IPrintPageRangeOptions
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PageRangeOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrintTaskOptions2) Get_CustomPageRanges() *IVector[*IPrintPageRange] {
	var _result *IVector[*IPrintPageRange]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CustomPageRanges, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1BDBB474-4ED1-41EB-BE3C-72D18ED67337
var IID_IPrintTaskOptionsCore = syscall.GUID{0x1BDBB474, 0x4ED1, 0x41EB,
	[8]byte{0xBE, 0x3C, 0x72, 0xD1, 0x8E, 0xD6, 0x73, 0x37}}

type IPrintTaskOptionsCoreInterface interface {
	win32.IInspectableInterface
	GetPageDescription(jobPageNumber uint32) PrintPageDescription
}

type IPrintTaskOptionsCoreVtbl struct {
	win32.IInspectableVtbl
	GetPageDescription uintptr
}

type IPrintTaskOptionsCore struct {
	win32.IInspectable
}

func (this *IPrintTaskOptionsCore) Vtbl() *IPrintTaskOptionsCoreVtbl {
	return (*IPrintTaskOptionsCoreVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrintTaskOptionsCore) GetPageDescription(jobPageNumber uint32) PrintPageDescription {
	var _result PrintPageDescription
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPageDescription, uintptr(unsafe.Pointer(this)), uintptr(jobPageNumber), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C1B71832-9E93-4E55-814B-3326A59EFCE1
var IID_IPrintTaskOptionsCoreProperties = syscall.GUID{0xC1B71832, 0x9E93, 0x4E55,
	[8]byte{0x81, 0x4B, 0x33, 0x26, 0xA5, 0x9E, 0xFC, 0xE1}}

type IPrintTaskOptionsCorePropertiesInterface interface {
	win32.IInspectableInterface
	Put_MediaSize(value PrintMediaSize)
	Get_MediaSize() PrintMediaSize
	Put_MediaType(value PrintMediaType)
	Get_MediaType() PrintMediaType
	Put_Orientation(value PrintOrientation)
	Get_Orientation() PrintOrientation
	Put_PrintQuality(value PrintQuality)
	Get_PrintQuality() PrintQuality
	Put_ColorMode(value PrintColorMode)
	Get_ColorMode() PrintColorMode
	Put_Duplex(value PrintDuplex)
	Get_Duplex() PrintDuplex
	Put_Collation(value PrintCollation)
	Get_Collation() PrintCollation
	Put_Staple(value PrintStaple)
	Get_Staple() PrintStaple
	Put_HolePunch(value PrintHolePunch)
	Get_HolePunch() PrintHolePunch
	Put_Binding(value PrintBinding)
	Get_Binding() PrintBinding
	Get_MinCopies() uint32
	Get_MaxCopies() uint32
	Put_NumberOfCopies(value uint32)
	Get_NumberOfCopies() uint32
}

type IPrintTaskOptionsCorePropertiesVtbl struct {
	win32.IInspectableVtbl
	Put_MediaSize      uintptr
	Get_MediaSize      uintptr
	Put_MediaType      uintptr
	Get_MediaType      uintptr
	Put_Orientation    uintptr
	Get_Orientation    uintptr
	Put_PrintQuality   uintptr
	Get_PrintQuality   uintptr
	Put_ColorMode      uintptr
	Get_ColorMode      uintptr
	Put_Duplex         uintptr
	Get_Duplex         uintptr
	Put_Collation      uintptr
	Get_Collation      uintptr
	Put_Staple         uintptr
	Get_Staple         uintptr
	Put_HolePunch      uintptr
	Get_HolePunch      uintptr
	Put_Binding        uintptr
	Get_Binding        uintptr
	Get_MinCopies      uintptr
	Get_MaxCopies      uintptr
	Put_NumberOfCopies uintptr
	Get_NumberOfCopies uintptr
}

type IPrintTaskOptionsCoreProperties struct {
	win32.IInspectable
}

func (this *IPrintTaskOptionsCoreProperties) Vtbl() *IPrintTaskOptionsCorePropertiesVtbl {
	return (*IPrintTaskOptionsCorePropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrintTaskOptionsCoreProperties) Put_MediaSize(value PrintMediaSize) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MediaSize, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrintTaskOptionsCoreProperties) Get_MediaSize() PrintMediaSize {
	var _result PrintMediaSize
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintTaskOptionsCoreProperties) Put_MediaType(value PrintMediaType) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MediaType, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrintTaskOptionsCoreProperties) Get_MediaType() PrintMediaType {
	var _result PrintMediaType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintTaskOptionsCoreProperties) Put_Orientation(value PrintOrientation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Orientation, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrintTaskOptionsCoreProperties) Get_Orientation() PrintOrientation {
	var _result PrintOrientation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Orientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintTaskOptionsCoreProperties) Put_PrintQuality(value PrintQuality) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PrintQuality, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrintTaskOptionsCoreProperties) Get_PrintQuality() PrintQuality {
	var _result PrintQuality
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrintQuality, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintTaskOptionsCoreProperties) Put_ColorMode(value PrintColorMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ColorMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrintTaskOptionsCoreProperties) Get_ColorMode() PrintColorMode {
	var _result PrintColorMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ColorMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintTaskOptionsCoreProperties) Put_Duplex(value PrintDuplex) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Duplex, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrintTaskOptionsCoreProperties) Get_Duplex() PrintDuplex {
	var _result PrintDuplex
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duplex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintTaskOptionsCoreProperties) Put_Collation(value PrintCollation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Collation, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrintTaskOptionsCoreProperties) Get_Collation() PrintCollation {
	var _result PrintCollation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Collation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintTaskOptionsCoreProperties) Put_Staple(value PrintStaple) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Staple, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrintTaskOptionsCoreProperties) Get_Staple() PrintStaple {
	var _result PrintStaple
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Staple, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintTaskOptionsCoreProperties) Put_HolePunch(value PrintHolePunch) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_HolePunch, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrintTaskOptionsCoreProperties) Get_HolePunch() PrintHolePunch {
	var _result PrintHolePunch
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HolePunch, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintTaskOptionsCoreProperties) Put_Binding(value PrintBinding) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Binding, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrintTaskOptionsCoreProperties) Get_Binding() PrintBinding {
	var _result PrintBinding
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Binding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintTaskOptionsCoreProperties) Get_MinCopies() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinCopies, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintTaskOptionsCoreProperties) Get_MaxCopies() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxCopies, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintTaskOptionsCoreProperties) Put_NumberOfCopies(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NumberOfCopies, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPrintTaskOptionsCoreProperties) Get_NumberOfCopies() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NumberOfCopies, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 62E69E23-9A1E-4336-B74F-3CC7F4CFF709
var IID_IPrintTaskOptionsCoreUIConfiguration = syscall.GUID{0x62E69E23, 0x9A1E, 0x4336,
	[8]byte{0xB7, 0x4F, 0x3C, 0xC7, 0xF4, 0xCF, 0xF7, 0x09}}

type IPrintTaskOptionsCoreUIConfigurationInterface interface {
	win32.IInspectableInterface
	Get_DisplayedOptions() *IVector[string]
}

type IPrintTaskOptionsCoreUIConfigurationVtbl struct {
	win32.IInspectableVtbl
	Get_DisplayedOptions uintptr
}

type IPrintTaskOptionsCoreUIConfiguration struct {
	win32.IInspectable
}

func (this *IPrintTaskOptionsCoreUIConfiguration) Vtbl() *IPrintTaskOptionsCoreUIConfigurationVtbl {
	return (*IPrintTaskOptionsCoreUIConfigurationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrintTaskOptionsCoreUIConfiguration) Get_DisplayedOptions() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayedOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 810CD3CB-B410-4282-A073-5AC378234174
var IID_IPrintTaskProgressingEventArgs = syscall.GUID{0x810CD3CB, 0xB410, 0x4282,
	[8]byte{0xA0, 0x73, 0x5A, 0xC3, 0x78, 0x23, 0x41, 0x74}}

type IPrintTaskProgressingEventArgsInterface interface {
	win32.IInspectableInterface
	Get_DocumentPageCount() uint32
}

type IPrintTaskProgressingEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_DocumentPageCount uintptr
}

type IPrintTaskProgressingEventArgs struct {
	win32.IInspectable
}

func (this *IPrintTaskProgressingEventArgs) Vtbl() *IPrintTaskProgressingEventArgsVtbl {
	return (*IPrintTaskProgressingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrintTaskProgressingEventArgs) Get_DocumentPageCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DocumentPageCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6FF61E2E-2722-4240-A67C-F364849A17F3
var IID_IPrintTaskRequest = syscall.GUID{0x6FF61E2E, 0x2722, 0x4240,
	[8]byte{0xA6, 0x7C, 0xF3, 0x64, 0x84, 0x9A, 0x17, 0xF3}}

type IPrintTaskRequestInterface interface {
	win32.IInspectableInterface
	Get_Deadline() DateTime
	CreatePrintTask(title string, handler PrintTaskSourceRequestedHandler) *IPrintTask
	GetDeferral() *IPrintTaskRequestedDeferral
}

type IPrintTaskRequestVtbl struct {
	win32.IInspectableVtbl
	Get_Deadline    uintptr
	CreatePrintTask uintptr
	GetDeferral     uintptr
}

type IPrintTaskRequest struct {
	win32.IInspectable
}

func (this *IPrintTaskRequest) Vtbl() *IPrintTaskRequestVtbl {
	return (*IPrintTaskRequestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrintTaskRequest) Get_Deadline() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Deadline, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintTaskRequest) CreatePrintTask(title string, handler PrintTaskSourceRequestedHandler) *IPrintTask {
	var _result *IPrintTask
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreatePrintTask, uintptr(unsafe.Pointer(this)), NewHStr(title).Ptr, uintptr(unsafe.Pointer(NewOneArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPrintTaskRequest) GetDeferral() *IPrintTaskRequestedDeferral {
	var _result *IPrintTaskRequestedDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CFEFB3F0-CE3E-42C7-9496-64800C622C44
var IID_IPrintTaskRequestedDeferral = syscall.GUID{0xCFEFB3F0, 0xCE3E, 0x42C7,
	[8]byte{0x94, 0x96, 0x64, 0x80, 0x0C, 0x62, 0x2C, 0x44}}

type IPrintTaskRequestedDeferralInterface interface {
	win32.IInspectableInterface
	Complete()
}

type IPrintTaskRequestedDeferralVtbl struct {
	win32.IInspectableVtbl
	Complete uintptr
}

type IPrintTaskRequestedDeferral struct {
	win32.IInspectable
}

func (this *IPrintTaskRequestedDeferral) Vtbl() *IPrintTaskRequestedDeferralVtbl {
	return (*IPrintTaskRequestedDeferralVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrintTaskRequestedDeferral) Complete() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Complete, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// D0AFF924-A31B-454C-A7B6-5D0CC522FC16
var IID_IPrintTaskRequestedEventArgs = syscall.GUID{0xD0AFF924, 0xA31B, 0x454C,
	[8]byte{0xA7, 0xB6, 0x5D, 0x0C, 0xC5, 0x22, 0xFC, 0x16}}

type IPrintTaskRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Request() *IPrintTaskRequest
}

type IPrintTaskRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Request uintptr
}

type IPrintTaskRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IPrintTaskRequestedEventArgs) Vtbl() *IPrintTaskRequestedEventArgsVtbl {
	return (*IPrintTaskRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrintTaskRequestedEventArgs) Get_Request() *IPrintTaskRequest {
	var _result *IPrintTaskRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Request, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F9F067BE-F456-41F0-9C98-5CE73E851410
var IID_IPrintTaskSourceRequestedArgs = syscall.GUID{0xF9F067BE, 0xF456, 0x41F0,
	[8]byte{0x9C, 0x98, 0x5C, 0xE7, 0x3E, 0x85, 0x14, 0x10}}

type IPrintTaskSourceRequestedArgsInterface interface {
	win32.IInspectableInterface
	Get_Deadline() DateTime
	SetSource(source *IPrintDocumentSource)
	GetDeferral() *IPrintTaskSourceRequestedDeferral
}

type IPrintTaskSourceRequestedArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Deadline uintptr
	SetSource    uintptr
	GetDeferral  uintptr
}

type IPrintTaskSourceRequestedArgs struct {
	win32.IInspectable
}

func (this *IPrintTaskSourceRequestedArgs) Vtbl() *IPrintTaskSourceRequestedArgsVtbl {
	return (*IPrintTaskSourceRequestedArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrintTaskSourceRequestedArgs) Get_Deadline() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Deadline, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintTaskSourceRequestedArgs) SetSource(source *IPrintDocumentSource) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)))
	_ = _hr
}

func (this *IPrintTaskSourceRequestedArgs) GetDeferral() *IPrintTaskSourceRequestedDeferral {
	var _result *IPrintTaskSourceRequestedDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4A1560D1-6992-4D9D-8555-4CA4563FB166
var IID_IPrintTaskSourceRequestedDeferral = syscall.GUID{0x4A1560D1, 0x6992, 0x4D9D,
	[8]byte{0x85, 0x55, 0x4C, 0xA4, 0x56, 0x3F, 0xB1, 0x66}}

type IPrintTaskSourceRequestedDeferralInterface interface {
	win32.IInspectableInterface
	Complete()
}

type IPrintTaskSourceRequestedDeferralVtbl struct {
	win32.IInspectableVtbl
	Complete uintptr
}

type IPrintTaskSourceRequestedDeferral struct {
	win32.IInspectable
}

func (this *IPrintTaskSourceRequestedDeferral) Vtbl() *IPrintTaskSourceRequestedDeferralVtbl {
	return (*IPrintTaskSourceRequestedDeferralVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrintTaskSourceRequestedDeferral) Complete() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Complete, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 295D70C0-C2CB-4B7D-B0EA-93095091A220
var IID_IPrintTaskTargetDeviceSupport = syscall.GUID{0x295D70C0, 0xC2CB, 0x4B7D,
	[8]byte{0xB0, 0xEA, 0x93, 0x09, 0x50, 0x91, 0xA2, 0x20}}

type IPrintTaskTargetDeviceSupportInterface interface {
	win32.IInspectableInterface
	Put_IsPrinterTargetEnabled(value bool)
	Get_IsPrinterTargetEnabled() bool
	Put_Is3DManufacturingTargetEnabled(value bool)
	Get_Is3DManufacturingTargetEnabled() bool
}

type IPrintTaskTargetDeviceSupportVtbl struct {
	win32.IInspectableVtbl
	Put_IsPrinterTargetEnabled         uintptr
	Get_IsPrinterTargetEnabled         uintptr
	Put_Is3DManufacturingTargetEnabled uintptr
	Get_Is3DManufacturingTargetEnabled uintptr
}

type IPrintTaskTargetDeviceSupport struct {
	win32.IInspectable
}

func (this *IPrintTaskTargetDeviceSupport) Vtbl() *IPrintTaskTargetDeviceSupportVtbl {
	return (*IPrintTaskTargetDeviceSupportVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrintTaskTargetDeviceSupport) Put_IsPrinterTargetEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsPrinterTargetEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPrintTaskTargetDeviceSupport) Get_IsPrinterTargetEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPrinterTargetEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrintTaskTargetDeviceSupport) Put_Is3DManufacturingTargetEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Is3DManufacturingTargetEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPrintTaskTargetDeviceSupport) Get_Is3DManufacturingTargetEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Is3DManufacturingTargetEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// B4483D26-0DD0-4CD4-BAFF-930FC7D6A574
var IID_IStandardPrintTaskOptionsStatic = syscall.GUID{0xB4483D26, 0x0DD0, 0x4CD4,
	[8]byte{0xBA, 0xFF, 0x93, 0x0F, 0xC7, 0xD6, 0xA5, 0x74}}

type IStandardPrintTaskOptionsStaticInterface interface {
	win32.IInspectableInterface
	Get_MediaSize() string
	Get_MediaType() string
	Get_Orientation() string
	Get_PrintQuality() string
	Get_ColorMode() string
	Get_Duplex() string
	Get_Collation() string
	Get_Staple() string
	Get_HolePunch() string
	Get_Binding() string
	Get_Copies() string
	Get_NUp() string
	Get_InputBin() string
}

type IStandardPrintTaskOptionsStaticVtbl struct {
	win32.IInspectableVtbl
	Get_MediaSize    uintptr
	Get_MediaType    uintptr
	Get_Orientation  uintptr
	Get_PrintQuality uintptr
	Get_ColorMode    uintptr
	Get_Duplex       uintptr
	Get_Collation    uintptr
	Get_Staple       uintptr
	Get_HolePunch    uintptr
	Get_Binding      uintptr
	Get_Copies       uintptr
	Get_NUp          uintptr
	Get_InputBin     uintptr
}

type IStandardPrintTaskOptionsStatic struct {
	win32.IInspectable
}

func (this *IStandardPrintTaskOptionsStatic) Vtbl() *IStandardPrintTaskOptionsStaticVtbl {
	return (*IStandardPrintTaskOptionsStaticVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStandardPrintTaskOptionsStatic) Get_MediaSize() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStandardPrintTaskOptionsStatic) Get_MediaType() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStandardPrintTaskOptionsStatic) Get_Orientation() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Orientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStandardPrintTaskOptionsStatic) Get_PrintQuality() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrintQuality, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStandardPrintTaskOptionsStatic) Get_ColorMode() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ColorMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStandardPrintTaskOptionsStatic) Get_Duplex() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duplex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStandardPrintTaskOptionsStatic) Get_Collation() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Collation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStandardPrintTaskOptionsStatic) Get_Staple() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Staple, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStandardPrintTaskOptionsStatic) Get_HolePunch() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HolePunch, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStandardPrintTaskOptionsStatic) Get_Binding() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Binding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStandardPrintTaskOptionsStatic) Get_Copies() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Copies, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStandardPrintTaskOptionsStatic) Get_NUp() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NUp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStandardPrintTaskOptionsStatic) Get_InputBin() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InputBin, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 3BE38BF4-7A44-4269-9A52-81261E289EE9
var IID_IStandardPrintTaskOptionsStatic2 = syscall.GUID{0x3BE38BF4, 0x7A44, 0x4269,
	[8]byte{0x9A, 0x52, 0x81, 0x26, 0x1E, 0x28, 0x9E, 0xE9}}

type IStandardPrintTaskOptionsStatic2Interface interface {
	win32.IInspectableInterface
	Get_Bordering() string
}

type IStandardPrintTaskOptionsStatic2Vtbl struct {
	win32.IInspectableVtbl
	Get_Bordering uintptr
}

type IStandardPrintTaskOptionsStatic2 struct {
	win32.IInspectable
}

func (this *IStandardPrintTaskOptionsStatic2) Vtbl() *IStandardPrintTaskOptionsStatic2Vtbl {
	return (*IStandardPrintTaskOptionsStatic2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStandardPrintTaskOptionsStatic2) Get_Bordering() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bordering, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// BBF68E86-3858-41B3-A799-55DD9888D475
var IID_IStandardPrintTaskOptionsStatic3 = syscall.GUID{0xBBF68E86, 0x3858, 0x41B3,
	[8]byte{0xA7, 0x99, 0x55, 0xDD, 0x98, 0x88, 0xD4, 0x75}}

type IStandardPrintTaskOptionsStatic3Interface interface {
	win32.IInspectableInterface
	Get_CustomPageRanges() string
}

type IStandardPrintTaskOptionsStatic3Vtbl struct {
	win32.IInspectableVtbl
	Get_CustomPageRanges uintptr
}

type IStandardPrintTaskOptionsStatic3 struct {
	win32.IInspectable
}

func (this *IStandardPrintTaskOptionsStatic3) Vtbl() *IStandardPrintTaskOptionsStatic3Vtbl {
	return (*IStandardPrintTaskOptionsStatic3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStandardPrintTaskOptionsStatic3) Get_CustomPageRanges() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CustomPageRanges, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// classes

type PrintPageInfo struct {
	RtClass
	*IPrintPageInfo
}

func NewPrintPageInfo() *PrintPageInfo {
	hs := NewHStr("Windows.Graphics.Printing.PrintPageInfo")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &PrintPageInfo{
		RtClass:        RtClass{PInspect: p},
		IPrintPageInfo: (*IPrintPageInfo)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type PrintPageRange struct {
	RtClass
	*IPrintPageRange
}

func NewPrintPageRange_Create(firstPage int32, lastPage int32) *PrintPageRange {
	hs := NewHStr("Windows.Graphics.Printing.PrintPageRange")
	var pFac *IPrintPageRangeFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPrintPageRangeFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IPrintPageRange
	p = pFac.Create(firstPage, lastPage)
	result := &PrintPageRange{
		RtClass:         RtClass{PInspect: &p.IInspectable},
		IPrintPageRange: p,
	}
	com.AddToScope(result)
	return result
}

func NewPrintPageRange_CreateWithSinglePage(page int32) *PrintPageRange {
	hs := NewHStr("Windows.Graphics.Printing.PrintPageRange")
	var pFac *IPrintPageRangeFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPrintPageRangeFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IPrintPageRange
	p = pFac.CreateWithSinglePage(page)
	result := &PrintPageRange{
		RtClass:         RtClass{PInspect: &p.IInspectable},
		IPrintPageRange: p,
	}
	com.AddToScope(result)
	return result
}

type StandardPrintTaskOptions struct {
	RtClass
}

func NewIStandardPrintTaskOptionsStatic() *IStandardPrintTaskOptionsStatic {
	var p *IStandardPrintTaskOptionsStatic
	hs := NewHStr("Windows.Graphics.Printing.StandardPrintTaskOptions")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IStandardPrintTaskOptionsStatic, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIStandardPrintTaskOptionsStatic2() *IStandardPrintTaskOptionsStatic2 {
	var p *IStandardPrintTaskOptionsStatic2
	hs := NewHStr("Windows.Graphics.Printing.StandardPrintTaskOptions")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IStandardPrintTaskOptionsStatic2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIStandardPrintTaskOptionsStatic3() *IStandardPrintTaskOptionsStatic3 {
	var p *IStandardPrintTaskOptionsStatic3
	hs := NewHStr("Windows.Graphics.Printing.StandardPrintTaskOptions")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IStandardPrintTaskOptionsStatic3, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
