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
type PaymentCanMakePaymentResultStatus int32

const (
	PaymentCanMakePaymentResultStatus_Unknown                               PaymentCanMakePaymentResultStatus = 0
	PaymentCanMakePaymentResultStatus_Yes                                   PaymentCanMakePaymentResultStatus = 1
	PaymentCanMakePaymentResultStatus_No                                    PaymentCanMakePaymentResultStatus = 2
	PaymentCanMakePaymentResultStatus_NotAllowed                            PaymentCanMakePaymentResultStatus = 3
	PaymentCanMakePaymentResultStatus_UserNotSignedIn                       PaymentCanMakePaymentResultStatus = 4
	PaymentCanMakePaymentResultStatus_SpecifiedPaymentMethodIdsNotSupported PaymentCanMakePaymentResultStatus = 5
	PaymentCanMakePaymentResultStatus_NoQualifyingCardOnFile                PaymentCanMakePaymentResultStatus = 6
)

// enum
type PaymentOptionPresence int32

const (
	PaymentOptionPresence_None     PaymentOptionPresence = 0
	PaymentOptionPresence_Optional PaymentOptionPresence = 1
	PaymentOptionPresence_Required PaymentOptionPresence = 2
)

// enum
type PaymentRequestChangeKind int32

const (
	PaymentRequestChangeKind_ShippingOption  PaymentRequestChangeKind = 0
	PaymentRequestChangeKind_ShippingAddress PaymentRequestChangeKind = 1
)

// enum
type PaymentRequestCompletionStatus int32

const (
	PaymentRequestCompletionStatus_Succeeded PaymentRequestCompletionStatus = 0
	PaymentRequestCompletionStatus_Failed    PaymentRequestCompletionStatus = 1
	PaymentRequestCompletionStatus_Unknown   PaymentRequestCompletionStatus = 2
)

// enum
type PaymentRequestStatus int32

const (
	PaymentRequestStatus_Succeeded PaymentRequestStatus = 0
	PaymentRequestStatus_Failed    PaymentRequestStatus = 1
	PaymentRequestStatus_Canceled  PaymentRequestStatus = 2
)

// enum
type PaymentShippingType int32

const (
	PaymentShippingType_Shipping PaymentShippingType = 0
	PaymentShippingType_Delivery PaymentShippingType = 1
	PaymentShippingType_Pickup   PaymentShippingType = 2
)

// func types

// 5078B9E1-F398-4F2C-A27E-94D371CF6C7D
type PaymentRequestChangedHandler func(paymentRequest *IPaymentRequest, args *IPaymentRequestChangedArgs) com.Error

// interfaces

// 5F2264E9-6F3A-4166-A018-0A0B06BB32B5
var IID_IPaymentAddress = syscall.GUID{0x5F2264E9, 0x6F3A, 0x4166,
	[8]byte{0xA0, 0x18, 0x0A, 0x0B, 0x06, 0xBB, 0x32, 0xB5}}

type IPaymentAddressInterface interface {
	win32.IInspectableInterface
	Get_Country() string
	Put_Country(value string)
	Get_AddressLines() *IVectorView[string]
	Put_AddressLines(value *IVectorView[string])
	Get_Region() string
	Put_Region(value string)
	Get_City() string
	Put_City(value string)
	Get_DependentLocality() string
	Put_DependentLocality(value string)
	Get_PostalCode() string
	Put_PostalCode(value string)
	Get_SortingCode() string
	Put_SortingCode(value string)
	Get_LanguageCode() string
	Put_LanguageCode(value string)
	Get_Organization() string
	Put_Organization(value string)
	Get_Recipient() string
	Put_Recipient(value string)
	Get_PhoneNumber() string
	Put_PhoneNumber(value string)
	Get_Properties() *IPropertySet
}

type IPaymentAddressVtbl struct {
	win32.IInspectableVtbl
	Get_Country           uintptr
	Put_Country           uintptr
	Get_AddressLines      uintptr
	Put_AddressLines      uintptr
	Get_Region            uintptr
	Put_Region            uintptr
	Get_City              uintptr
	Put_City              uintptr
	Get_DependentLocality uintptr
	Put_DependentLocality uintptr
	Get_PostalCode        uintptr
	Put_PostalCode        uintptr
	Get_SortingCode       uintptr
	Put_SortingCode       uintptr
	Get_LanguageCode      uintptr
	Put_LanguageCode      uintptr
	Get_Organization      uintptr
	Put_Organization      uintptr
	Get_Recipient         uintptr
	Put_Recipient         uintptr
	Get_PhoneNumber       uintptr
	Put_PhoneNumber       uintptr
	Get_Properties        uintptr
}

type IPaymentAddress struct {
	win32.IInspectable
}

func (this *IPaymentAddress) Vtbl() *IPaymentAddressVtbl {
	return (*IPaymentAddressVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentAddress) Get_Country() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Country, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPaymentAddress) Put_Country(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Country, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPaymentAddress) Get_AddressLines() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AddressLines, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentAddress) Put_AddressLines(value *IVectorView[string]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AddressLines, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPaymentAddress) Get_Region() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Region, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPaymentAddress) Put_Region(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Region, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPaymentAddress) Get_City() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_City, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPaymentAddress) Put_City(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_City, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPaymentAddress) Get_DependentLocality() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DependentLocality, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPaymentAddress) Put_DependentLocality(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DependentLocality, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPaymentAddress) Get_PostalCode() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PostalCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPaymentAddress) Put_PostalCode(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PostalCode, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPaymentAddress) Get_SortingCode() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SortingCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPaymentAddress) Put_SortingCode(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SortingCode, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPaymentAddress) Get_LanguageCode() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LanguageCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPaymentAddress) Put_LanguageCode(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LanguageCode, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPaymentAddress) Get_Organization() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Organization, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPaymentAddress) Put_Organization(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Organization, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPaymentAddress) Get_Recipient() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Recipient, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPaymentAddress) Put_Recipient(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Recipient, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPaymentAddress) Get_PhoneNumber() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhoneNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPaymentAddress) Put_PhoneNumber(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PhoneNumber, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPaymentAddress) Get_Properties() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7696FE55-D5D3-4D3D-B345-45591759C510
var IID_IPaymentCanMakePaymentResult = syscall.GUID{0x7696FE55, 0xD5D3, 0x4D3D,
	[8]byte{0xB3, 0x45, 0x45, 0x59, 0x17, 0x59, 0xC5, 0x10}}

type IPaymentCanMakePaymentResultInterface interface {
	win32.IInspectableInterface
	Get_Status() PaymentCanMakePaymentResultStatus
}

type IPaymentCanMakePaymentResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
}

type IPaymentCanMakePaymentResult struct {
	win32.IInspectable
}

func (this *IPaymentCanMakePaymentResult) Vtbl() *IPaymentCanMakePaymentResultVtbl {
	return (*IPaymentCanMakePaymentResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentCanMakePaymentResult) Get_Status() PaymentCanMakePaymentResultStatus {
	var _result PaymentCanMakePaymentResultStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// BBDCAA3E-7D49-4F69-AA53-2A0F8164B7C9
var IID_IPaymentCanMakePaymentResultFactory = syscall.GUID{0xBBDCAA3E, 0x7D49, 0x4F69,
	[8]byte{0xAA, 0x53, 0x2A, 0x0F, 0x81, 0x64, 0xB7, 0xC9}}

type IPaymentCanMakePaymentResultFactoryInterface interface {
	win32.IInspectableInterface
	Create(value PaymentCanMakePaymentResultStatus) *IPaymentCanMakePaymentResult
}

type IPaymentCanMakePaymentResultFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IPaymentCanMakePaymentResultFactory struct {
	win32.IInspectable
}

func (this *IPaymentCanMakePaymentResultFactory) Vtbl() *IPaymentCanMakePaymentResultFactoryVtbl {
	return (*IPaymentCanMakePaymentResultFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentCanMakePaymentResultFactory) Create(value PaymentCanMakePaymentResultStatus) *IPaymentCanMakePaymentResult {
	var _result *IPaymentCanMakePaymentResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E3A3E9E0-B41F-4987-BDCB-071331F2DAA4
var IID_IPaymentCurrencyAmount = syscall.GUID{0xE3A3E9E0, 0xB41F, 0x4987,
	[8]byte{0xBD, 0xCB, 0x07, 0x13, 0x31, 0xF2, 0xDA, 0xA4}}

type IPaymentCurrencyAmountInterface interface {
	win32.IInspectableInterface
	Get_Currency() string
	Put_Currency(value string)
	Get_CurrencySystem() string
	Put_CurrencySystem(value string)
	Get_Value() string
	Put_Value(value string)
}

type IPaymentCurrencyAmountVtbl struct {
	win32.IInspectableVtbl
	Get_Currency       uintptr
	Put_Currency       uintptr
	Get_CurrencySystem uintptr
	Put_CurrencySystem uintptr
	Get_Value          uintptr
	Put_Value          uintptr
}

type IPaymentCurrencyAmount struct {
	win32.IInspectable
}

func (this *IPaymentCurrencyAmount) Vtbl() *IPaymentCurrencyAmountVtbl {
	return (*IPaymentCurrencyAmountVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentCurrencyAmount) Get_Currency() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Currency, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPaymentCurrencyAmount) Put_Currency(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Currency, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPaymentCurrencyAmount) Get_CurrencySystem() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrencySystem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPaymentCurrencyAmount) Put_CurrencySystem(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CurrencySystem, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPaymentCurrencyAmount) Get_Value() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPaymentCurrencyAmount) Put_Value(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Value, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 3257D338-140C-4575-8535-F773178C09A7
var IID_IPaymentCurrencyAmountFactory = syscall.GUID{0x3257D338, 0x140C, 0x4575,
	[8]byte{0x85, 0x35, 0xF7, 0x73, 0x17, 0x8C, 0x09, 0xA7}}

type IPaymentCurrencyAmountFactoryInterface interface {
	win32.IInspectableInterface
	Create(value string, currency string) *IPaymentCurrencyAmount
	CreateWithCurrencySystem(value string, currency string, currencySystem string) *IPaymentCurrencyAmount
}

type IPaymentCurrencyAmountFactoryVtbl struct {
	win32.IInspectableVtbl
	Create                   uintptr
	CreateWithCurrencySystem uintptr
}

type IPaymentCurrencyAmountFactory struct {
	win32.IInspectable
}

func (this *IPaymentCurrencyAmountFactory) Vtbl() *IPaymentCurrencyAmountFactoryVtbl {
	return (*IPaymentCurrencyAmountFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentCurrencyAmountFactory) Create(value string, currency string) *IPaymentCurrencyAmount {
	var _result *IPaymentCurrencyAmount
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr, NewHStr(currency).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentCurrencyAmountFactory) CreateWithCurrencySystem(value string, currency string, currencySystem string) *IPaymentCurrencyAmount {
	var _result *IPaymentCurrencyAmount
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithCurrencySystem, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr, NewHStr(currency).Ptr, NewHStr(currencySystem).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 53BB2D7D-E0EB-4053-8EAE-CE7C48E02945
var IID_IPaymentDetails = syscall.GUID{0x53BB2D7D, 0xE0EB, 0x4053,
	[8]byte{0x8E, 0xAE, 0xCE, 0x7C, 0x48, 0xE0, 0x29, 0x45}}

type IPaymentDetailsInterface interface {
	win32.IInspectableInterface
	Get_Total() *IPaymentItem
	Put_Total(value *IPaymentItem)
	Get_DisplayItems() *IVectorView[*IPaymentItem]
	Put_DisplayItems(value *IVectorView[*IPaymentItem])
	Get_ShippingOptions() *IVectorView[*IPaymentShippingOption]
	Put_ShippingOptions(value *IVectorView[*IPaymentShippingOption])
	Get_Modifiers() *IVectorView[*IPaymentDetailsModifier]
	Put_Modifiers(value *IVectorView[*IPaymentDetailsModifier])
}

type IPaymentDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_Total           uintptr
	Put_Total           uintptr
	Get_DisplayItems    uintptr
	Put_DisplayItems    uintptr
	Get_ShippingOptions uintptr
	Put_ShippingOptions uintptr
	Get_Modifiers       uintptr
	Put_Modifiers       uintptr
}

type IPaymentDetails struct {
	win32.IInspectable
}

func (this *IPaymentDetails) Vtbl() *IPaymentDetailsVtbl {
	return (*IPaymentDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentDetails) Get_Total() *IPaymentItem {
	var _result *IPaymentItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Total, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentDetails) Put_Total(value *IPaymentItem) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Total, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPaymentDetails) Get_DisplayItems() *IVectorView[*IPaymentItem] {
	var _result *IVectorView[*IPaymentItem]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayItems, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentDetails) Put_DisplayItems(value *IVectorView[*IPaymentItem]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DisplayItems, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPaymentDetails) Get_ShippingOptions() *IVectorView[*IPaymentShippingOption] {
	var _result *IVectorView[*IPaymentShippingOption]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShippingOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentDetails) Put_ShippingOptions(value *IVectorView[*IPaymentShippingOption]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ShippingOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPaymentDetails) Get_Modifiers() *IVectorView[*IPaymentDetailsModifier] {
	var _result *IVectorView[*IPaymentDetailsModifier]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Modifiers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentDetails) Put_Modifiers(value *IVectorView[*IPaymentDetailsModifier]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Modifiers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// CFE8AFEE-C0EA-4CA1-8BC7-6DE67B1F3763
var IID_IPaymentDetailsFactory = syscall.GUID{0xCFE8AFEE, 0xC0EA, 0x4CA1,
	[8]byte{0x8B, 0xC7, 0x6D, 0xE6, 0x7B, 0x1F, 0x37, 0x63}}

type IPaymentDetailsFactoryInterface interface {
	win32.IInspectableInterface
	Create(total *IPaymentItem) *IPaymentDetails
	CreateWithDisplayItems(total *IPaymentItem, displayItems *IIterable[*IPaymentItem]) *IPaymentDetails
}

type IPaymentDetailsFactoryVtbl struct {
	win32.IInspectableVtbl
	Create                 uintptr
	CreateWithDisplayItems uintptr
}

type IPaymentDetailsFactory struct {
	win32.IInspectable
}

func (this *IPaymentDetailsFactory) Vtbl() *IPaymentDetailsFactoryVtbl {
	return (*IPaymentDetailsFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentDetailsFactory) Create(total *IPaymentItem) *IPaymentDetails {
	var _result *IPaymentDetails
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(total)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentDetailsFactory) CreateWithDisplayItems(total *IPaymentItem, displayItems *IIterable[*IPaymentItem]) *IPaymentDetails {
	var _result *IPaymentDetails
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithDisplayItems, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(total)), uintptr(unsafe.Pointer(displayItems)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BE1C7D65-4323-41D7-B305-DFCB765F69DE
var IID_IPaymentDetailsModifier = syscall.GUID{0xBE1C7D65, 0x4323, 0x41D7,
	[8]byte{0xB3, 0x05, 0xDF, 0xCB, 0x76, 0x5F, 0x69, 0xDE}}

type IPaymentDetailsModifierInterface interface {
	win32.IInspectableInterface
	Get_JsonData() string
	Get_SupportedMethodIds() *IVectorView[string]
	Get_Total() *IPaymentItem
	Get_AdditionalDisplayItems() *IVectorView[*IPaymentItem]
}

type IPaymentDetailsModifierVtbl struct {
	win32.IInspectableVtbl
	Get_JsonData               uintptr
	Get_SupportedMethodIds     uintptr
	Get_Total                  uintptr
	Get_AdditionalDisplayItems uintptr
}

type IPaymentDetailsModifier struct {
	win32.IInspectable
}

func (this *IPaymentDetailsModifier) Vtbl() *IPaymentDetailsModifierVtbl {
	return (*IPaymentDetailsModifierVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentDetailsModifier) Get_JsonData() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_JsonData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPaymentDetailsModifier) Get_SupportedMethodIds() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedMethodIds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentDetailsModifier) Get_Total() *IPaymentItem {
	var _result *IPaymentItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Total, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentDetailsModifier) Get_AdditionalDisplayItems() *IVectorView[*IPaymentItem] {
	var _result *IVectorView[*IPaymentItem]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AdditionalDisplayItems, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 79005286-54DE-429C-9E4F-5DCE6E10EBCE
var IID_IPaymentDetailsModifierFactory = syscall.GUID{0x79005286, 0x54DE, 0x429C,
	[8]byte{0x9E, 0x4F, 0x5D, 0xCE, 0x6E, 0x10, 0xEB, 0xCE}}

type IPaymentDetailsModifierFactoryInterface interface {
	win32.IInspectableInterface
	Create(supportedMethodIds *IIterable[string], total *IPaymentItem) *IPaymentDetailsModifier
	CreateWithAdditionalDisplayItems(supportedMethodIds *IIterable[string], total *IPaymentItem, additionalDisplayItems *IIterable[*IPaymentItem]) *IPaymentDetailsModifier
	CreateWithAdditionalDisplayItemsAndJsonData(supportedMethodIds *IIterable[string], total *IPaymentItem, additionalDisplayItems *IIterable[*IPaymentItem], jsonData string) *IPaymentDetailsModifier
}

type IPaymentDetailsModifierFactoryVtbl struct {
	win32.IInspectableVtbl
	Create                                      uintptr
	CreateWithAdditionalDisplayItems            uintptr
	CreateWithAdditionalDisplayItemsAndJsonData uintptr
}

type IPaymentDetailsModifierFactory struct {
	win32.IInspectable
}

func (this *IPaymentDetailsModifierFactory) Vtbl() *IPaymentDetailsModifierFactoryVtbl {
	return (*IPaymentDetailsModifierFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentDetailsModifierFactory) Create(supportedMethodIds *IIterable[string], total *IPaymentItem) *IPaymentDetailsModifier {
	var _result *IPaymentDetailsModifier
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(supportedMethodIds)), uintptr(unsafe.Pointer(total)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentDetailsModifierFactory) CreateWithAdditionalDisplayItems(supportedMethodIds *IIterable[string], total *IPaymentItem, additionalDisplayItems *IIterable[*IPaymentItem]) *IPaymentDetailsModifier {
	var _result *IPaymentDetailsModifier
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithAdditionalDisplayItems, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(supportedMethodIds)), uintptr(unsafe.Pointer(total)), uintptr(unsafe.Pointer(additionalDisplayItems)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentDetailsModifierFactory) CreateWithAdditionalDisplayItemsAndJsonData(supportedMethodIds *IIterable[string], total *IPaymentItem, additionalDisplayItems *IIterable[*IPaymentItem], jsonData string) *IPaymentDetailsModifier {
	var _result *IPaymentDetailsModifier
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithAdditionalDisplayItemsAndJsonData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(supportedMethodIds)), uintptr(unsafe.Pointer(total)), uintptr(unsafe.Pointer(additionalDisplayItems)), NewHStr(jsonData).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 685AC88B-79B2-4B76-9E03-A876223DFE72
var IID_IPaymentItem = syscall.GUID{0x685AC88B, 0x79B2, 0x4B76,
	[8]byte{0x9E, 0x03, 0xA8, 0x76, 0x22, 0x3D, 0xFE, 0x72}}

type IPaymentItemInterface interface {
	win32.IInspectableInterface
	Get_Label() string
	Put_Label(value string)
	Get_Amount() *IPaymentCurrencyAmount
	Put_Amount(value *IPaymentCurrencyAmount)
	Get_Pending() bool
	Put_Pending(value bool)
}

type IPaymentItemVtbl struct {
	win32.IInspectableVtbl
	Get_Label   uintptr
	Put_Label   uintptr
	Get_Amount  uintptr
	Put_Amount  uintptr
	Get_Pending uintptr
	Put_Pending uintptr
}

type IPaymentItem struct {
	win32.IInspectable
}

func (this *IPaymentItem) Vtbl() *IPaymentItemVtbl {
	return (*IPaymentItemVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentItem) Get_Label() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Label, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPaymentItem) Put_Label(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Label, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPaymentItem) Get_Amount() *IPaymentCurrencyAmount {
	var _result *IPaymentCurrencyAmount
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Amount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentItem) Put_Amount(value *IPaymentCurrencyAmount) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Amount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPaymentItem) Get_Pending() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Pending, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPaymentItem) Put_Pending(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Pending, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// C6AB7AD8-2503-4D1D-A778-02B2E5927B2C
var IID_IPaymentItemFactory = syscall.GUID{0xC6AB7AD8, 0x2503, 0x4D1D,
	[8]byte{0xA7, 0x78, 0x02, 0xB2, 0xE5, 0x92, 0x7B, 0x2C}}

type IPaymentItemFactoryInterface interface {
	win32.IInspectableInterface
	Create(label string, amount *IPaymentCurrencyAmount) *IPaymentItem
}

type IPaymentItemFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IPaymentItemFactory struct {
	win32.IInspectable
}

func (this *IPaymentItemFactory) Vtbl() *IPaymentItemFactoryVtbl {
	return (*IPaymentItemFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentItemFactory) Create(label string, amount *IPaymentCurrencyAmount) *IPaymentItem {
	var _result *IPaymentItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(label).Ptr, uintptr(unsafe.Pointer(amount)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FB0EE829-EC0C-449A-83DA-7AE3073365A2
var IID_IPaymentMediator = syscall.GUID{0xFB0EE829, 0xEC0C, 0x449A,
	[8]byte{0x83, 0xDA, 0x7A, 0xE3, 0x07, 0x33, 0x65, 0xA2}}

type IPaymentMediatorInterface interface {
	win32.IInspectableInterface
	GetSupportedMethodIdsAsync() *IAsyncOperation[*IVectorView[string]]
	SubmitPaymentRequestAsync(paymentRequest *IPaymentRequest) *IAsyncOperation[*IPaymentRequestSubmitResult]
	SubmitPaymentRequestWithChangeHandlerAsync(paymentRequest *IPaymentRequest, changeHandler PaymentRequestChangedHandler) *IAsyncOperation[*IPaymentRequestSubmitResult]
}

type IPaymentMediatorVtbl struct {
	win32.IInspectableVtbl
	GetSupportedMethodIdsAsync                 uintptr
	SubmitPaymentRequestAsync                  uintptr
	SubmitPaymentRequestWithChangeHandlerAsync uintptr
}

type IPaymentMediator struct {
	win32.IInspectable
}

func (this *IPaymentMediator) Vtbl() *IPaymentMediatorVtbl {
	return (*IPaymentMediatorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentMediator) GetSupportedMethodIdsAsync() *IAsyncOperation[*IVectorView[string]] {
	var _result *IAsyncOperation[*IVectorView[string]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSupportedMethodIdsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentMediator) SubmitPaymentRequestAsync(paymentRequest *IPaymentRequest) *IAsyncOperation[*IPaymentRequestSubmitResult] {
	var _result *IAsyncOperation[*IPaymentRequestSubmitResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SubmitPaymentRequestAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(paymentRequest)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentMediator) SubmitPaymentRequestWithChangeHandlerAsync(paymentRequest *IPaymentRequest, changeHandler PaymentRequestChangedHandler) *IAsyncOperation[*IPaymentRequestSubmitResult] {
	var _result *IAsyncOperation[*IPaymentRequestSubmitResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SubmitPaymentRequestWithChangeHandlerAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(paymentRequest)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(changeHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CEEF98F1-E407-4128-8E73-D93D5F822786
var IID_IPaymentMediator2 = syscall.GUID{0xCEEF98F1, 0xE407, 0x4128,
	[8]byte{0x8E, 0x73, 0xD9, 0x3D, 0x5F, 0x82, 0x27, 0x86}}

type IPaymentMediator2Interface interface {
	win32.IInspectableInterface
	CanMakePaymentAsync(paymentRequest *IPaymentRequest) *IAsyncOperation[*IPaymentCanMakePaymentResult]
}

type IPaymentMediator2Vtbl struct {
	win32.IInspectableVtbl
	CanMakePaymentAsync uintptr
}

type IPaymentMediator2 struct {
	win32.IInspectable
}

func (this *IPaymentMediator2) Vtbl() *IPaymentMediator2Vtbl {
	return (*IPaymentMediator2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentMediator2) CanMakePaymentAsync(paymentRequest *IPaymentRequest) *IAsyncOperation[*IPaymentCanMakePaymentResult] {
	var _result *IAsyncOperation[*IPaymentCanMakePaymentResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CanMakePaymentAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(paymentRequest)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 63445050-0E94-4ED6-AACB-E6012BD327A7
var IID_IPaymentMerchantInfo = syscall.GUID{0x63445050, 0x0E94, 0x4ED6,
	[8]byte{0xAA, 0xCB, 0xE6, 0x01, 0x2B, 0xD3, 0x27, 0xA7}}

type IPaymentMerchantInfoInterface interface {
	win32.IInspectableInterface
	Get_PackageFullName() string
	Get_Uri() *IUriRuntimeClass
}

type IPaymentMerchantInfoVtbl struct {
	win32.IInspectableVtbl
	Get_PackageFullName uintptr
	Get_Uri             uintptr
}

type IPaymentMerchantInfo struct {
	win32.IInspectable
}

func (this *IPaymentMerchantInfo) Vtbl() *IPaymentMerchantInfoVtbl {
	return (*IPaymentMerchantInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentMerchantInfo) Get_PackageFullName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PackageFullName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPaymentMerchantInfo) Get_Uri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Uri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9E89CED3-CCB7-4167-A8EC-E10AE96DBCD1
var IID_IPaymentMerchantInfoFactory = syscall.GUID{0x9E89CED3, 0xCCB7, 0x4167,
	[8]byte{0xA8, 0xEC, 0xE1, 0x0A, 0xE9, 0x6D, 0xBC, 0xD1}}

type IPaymentMerchantInfoFactoryInterface interface {
	win32.IInspectableInterface
	Create(uri *IUriRuntimeClass) *IPaymentMerchantInfo
}

type IPaymentMerchantInfoFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IPaymentMerchantInfoFactory struct {
	win32.IInspectable
}

func (this *IPaymentMerchantInfoFactory) Vtbl() *IPaymentMerchantInfoFactoryVtbl {
	return (*IPaymentMerchantInfoFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentMerchantInfoFactory) Create(uri *IUriRuntimeClass) *IPaymentMerchantInfo {
	var _result *IPaymentMerchantInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D1D3CAF4-DE98-4129-B1B7-C3AD86237BF4
var IID_IPaymentMethodData = syscall.GUID{0xD1D3CAF4, 0xDE98, 0x4129,
	[8]byte{0xB1, 0xB7, 0xC3, 0xAD, 0x86, 0x23, 0x7B, 0xF4}}

type IPaymentMethodDataInterface interface {
	win32.IInspectableInterface
	Get_SupportedMethodIds() *IVectorView[string]
	Get_JsonData() string
}

type IPaymentMethodDataVtbl struct {
	win32.IInspectableVtbl
	Get_SupportedMethodIds uintptr
	Get_JsonData           uintptr
}

type IPaymentMethodData struct {
	win32.IInspectable
}

func (this *IPaymentMethodData) Vtbl() *IPaymentMethodDataVtbl {
	return (*IPaymentMethodDataVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentMethodData) Get_SupportedMethodIds() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedMethodIds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentMethodData) Get_JsonData() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_JsonData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 8ADDD27F-9BAA-4A82-8342-A8210992A36B
var IID_IPaymentMethodDataFactory = syscall.GUID{0x8ADDD27F, 0x9BAA, 0x4A82,
	[8]byte{0x83, 0x42, 0xA8, 0x21, 0x09, 0x92, 0xA3, 0x6B}}

type IPaymentMethodDataFactoryInterface interface {
	win32.IInspectableInterface
	Create(supportedMethodIds *IIterable[string]) *IPaymentMethodData
	CreateWithJsonData(supportedMethodIds *IIterable[string], jsonData string) *IPaymentMethodData
}

type IPaymentMethodDataFactoryVtbl struct {
	win32.IInspectableVtbl
	Create             uintptr
	CreateWithJsonData uintptr
}

type IPaymentMethodDataFactory struct {
	win32.IInspectable
}

func (this *IPaymentMethodDataFactory) Vtbl() *IPaymentMethodDataFactoryVtbl {
	return (*IPaymentMethodDataFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentMethodDataFactory) Create(supportedMethodIds *IIterable[string]) *IPaymentMethodData {
	var _result *IPaymentMethodData
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(supportedMethodIds)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentMethodDataFactory) CreateWithJsonData(supportedMethodIds *IIterable[string], jsonData string) *IPaymentMethodData {
	var _result *IPaymentMethodData
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithJsonData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(supportedMethodIds)), NewHStr(jsonData).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AAA30854-1F2B-4365-8251-01B58915A5BC
var IID_IPaymentOptions = syscall.GUID{0xAAA30854, 0x1F2B, 0x4365,
	[8]byte{0x82, 0x51, 0x01, 0xB5, 0x89, 0x15, 0xA5, 0xBC}}

type IPaymentOptionsInterface interface {
	win32.IInspectableInterface
	Get_RequestPayerEmail() PaymentOptionPresence
	Put_RequestPayerEmail(value PaymentOptionPresence)
	Get_RequestPayerName() PaymentOptionPresence
	Put_RequestPayerName(value PaymentOptionPresence)
	Get_RequestPayerPhoneNumber() PaymentOptionPresence
	Put_RequestPayerPhoneNumber(value PaymentOptionPresence)
	Get_RequestShipping() bool
	Put_RequestShipping(value bool)
	Get_ShippingType() PaymentShippingType
	Put_ShippingType(value PaymentShippingType)
}

type IPaymentOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_RequestPayerEmail       uintptr
	Put_RequestPayerEmail       uintptr
	Get_RequestPayerName        uintptr
	Put_RequestPayerName        uintptr
	Get_RequestPayerPhoneNumber uintptr
	Put_RequestPayerPhoneNumber uintptr
	Get_RequestShipping         uintptr
	Put_RequestShipping         uintptr
	Get_ShippingType            uintptr
	Put_ShippingType            uintptr
}

type IPaymentOptions struct {
	win32.IInspectable
}

func (this *IPaymentOptions) Vtbl() *IPaymentOptionsVtbl {
	return (*IPaymentOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentOptions) Get_RequestPayerEmail() PaymentOptionPresence {
	var _result PaymentOptionPresence
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestPayerEmail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPaymentOptions) Put_RequestPayerEmail(value PaymentOptionPresence) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RequestPayerEmail, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPaymentOptions) Get_RequestPayerName() PaymentOptionPresence {
	var _result PaymentOptionPresence
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestPayerName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPaymentOptions) Put_RequestPayerName(value PaymentOptionPresence) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RequestPayerName, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPaymentOptions) Get_RequestPayerPhoneNumber() PaymentOptionPresence {
	var _result PaymentOptionPresence
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestPayerPhoneNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPaymentOptions) Put_RequestPayerPhoneNumber(value PaymentOptionPresence) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RequestPayerPhoneNumber, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPaymentOptions) Get_RequestShipping() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestShipping, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPaymentOptions) Put_RequestShipping(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RequestShipping, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPaymentOptions) Get_ShippingType() PaymentShippingType {
	var _result PaymentShippingType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShippingType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPaymentOptions) Put_ShippingType(value PaymentShippingType) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ShippingType, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// B74942E1-ED7B-47EB-BC08-78CC5D6896B6
var IID_IPaymentRequest = syscall.GUID{0xB74942E1, 0xED7B, 0x47EB,
	[8]byte{0xBC, 0x08, 0x78, 0xCC, 0x5D, 0x68, 0x96, 0xB6}}

type IPaymentRequestInterface interface {
	win32.IInspectableInterface
	Get_MerchantInfo() *IPaymentMerchantInfo
	Get_Details() *IPaymentDetails
	Get_MethodData() *IVectorView[*IPaymentMethodData]
	Get_Options() *IPaymentOptions
}

type IPaymentRequestVtbl struct {
	win32.IInspectableVtbl
	Get_MerchantInfo uintptr
	Get_Details      uintptr
	Get_MethodData   uintptr
	Get_Options      uintptr
}

type IPaymentRequest struct {
	win32.IInspectable
}

func (this *IPaymentRequest) Vtbl() *IPaymentRequestVtbl {
	return (*IPaymentRequestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentRequest) Get_MerchantInfo() *IPaymentMerchantInfo {
	var _result *IPaymentMerchantInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MerchantInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentRequest) Get_Details() *IPaymentDetails {
	var _result *IPaymentDetails
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Details, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentRequest) Get_MethodData() *IVectorView[*IPaymentMethodData] {
	var _result *IVectorView[*IPaymentMethodData]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MethodData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentRequest) Get_Options() *IPaymentOptions {
	var _result *IPaymentOptions
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Options, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B63CCFB5-5998-493E-A04C-67048A50F141
var IID_IPaymentRequest2 = syscall.GUID{0xB63CCFB5, 0x5998, 0x493E,
	[8]byte{0xA0, 0x4C, 0x67, 0x04, 0x8A, 0x50, 0xF1, 0x41}}

type IPaymentRequest2Interface interface {
	win32.IInspectableInterface
	Get_Id() string
}

type IPaymentRequest2Vtbl struct {
	win32.IInspectableVtbl
	Get_Id uintptr
}

type IPaymentRequest2 struct {
	win32.IInspectable
}

func (this *IPaymentRequest2) Vtbl() *IPaymentRequest2Vtbl {
	return (*IPaymentRequest2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentRequest2) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// C6145E44-CD8B-4BE4-B555-27C99194C0C5
var IID_IPaymentRequestChangedArgs = syscall.GUID{0xC6145E44, 0xCD8B, 0x4BE4,
	[8]byte{0xB5, 0x55, 0x27, 0xC9, 0x91, 0x94, 0xC0, 0xC5}}

type IPaymentRequestChangedArgsInterface interface {
	win32.IInspectableInterface
	Get_ChangeKind() PaymentRequestChangeKind
	Get_ShippingAddress() *IPaymentAddress
	Get_SelectedShippingOption() *IPaymentShippingOption
	Acknowledge(changeResult *IPaymentRequestChangedResult)
}

type IPaymentRequestChangedArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ChangeKind             uintptr
	Get_ShippingAddress        uintptr
	Get_SelectedShippingOption uintptr
	Acknowledge                uintptr
}

type IPaymentRequestChangedArgs struct {
	win32.IInspectable
}

func (this *IPaymentRequestChangedArgs) Vtbl() *IPaymentRequestChangedArgsVtbl {
	return (*IPaymentRequestChangedArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentRequestChangedArgs) Get_ChangeKind() PaymentRequestChangeKind {
	var _result PaymentRequestChangeKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ChangeKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPaymentRequestChangedArgs) Get_ShippingAddress() *IPaymentAddress {
	var _result *IPaymentAddress
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShippingAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentRequestChangedArgs) Get_SelectedShippingOption() *IPaymentShippingOption {
	var _result *IPaymentShippingOption
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectedShippingOption, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentRequestChangedArgs) Acknowledge(changeResult *IPaymentRequestChangedResult) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Acknowledge, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(changeResult)))
	_ = _hr
}

// DF699E5C-16C4-47AD-9401-8440EC0757DB
var IID_IPaymentRequestChangedResult = syscall.GUID{0xDF699E5C, 0x16C4, 0x47AD,
	[8]byte{0x94, 0x01, 0x84, 0x40, 0xEC, 0x07, 0x57, 0xDB}}

type IPaymentRequestChangedResultInterface interface {
	win32.IInspectableInterface
	Get_ChangeAcceptedByMerchant() bool
	Put_ChangeAcceptedByMerchant(value bool)
	Get_Message() string
	Put_Message(value string)
	Get_UpdatedPaymentDetails() *IPaymentDetails
	Put_UpdatedPaymentDetails(value *IPaymentDetails)
}

type IPaymentRequestChangedResultVtbl struct {
	win32.IInspectableVtbl
	Get_ChangeAcceptedByMerchant uintptr
	Put_ChangeAcceptedByMerchant uintptr
	Get_Message                  uintptr
	Put_Message                  uintptr
	Get_UpdatedPaymentDetails    uintptr
	Put_UpdatedPaymentDetails    uintptr
}

type IPaymentRequestChangedResult struct {
	win32.IInspectable
}

func (this *IPaymentRequestChangedResult) Vtbl() *IPaymentRequestChangedResultVtbl {
	return (*IPaymentRequestChangedResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentRequestChangedResult) Get_ChangeAcceptedByMerchant() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ChangeAcceptedByMerchant, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPaymentRequestChangedResult) Put_ChangeAcceptedByMerchant(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ChangeAcceptedByMerchant, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPaymentRequestChangedResult) Get_Message() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Message, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPaymentRequestChangedResult) Put_Message(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Message, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPaymentRequestChangedResult) Get_UpdatedPaymentDetails() *IPaymentDetails {
	var _result *IPaymentDetails
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UpdatedPaymentDetails, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentRequestChangedResult) Put_UpdatedPaymentDetails(value *IPaymentDetails) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_UpdatedPaymentDetails, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 08740F56-1D33-4431-814B-67EA24BF21DB
var IID_IPaymentRequestChangedResultFactory = syscall.GUID{0x08740F56, 0x1D33, 0x4431,
	[8]byte{0x81, 0x4B, 0x67, 0xEA, 0x24, 0xBF, 0x21, 0xDB}}

type IPaymentRequestChangedResultFactoryInterface interface {
	win32.IInspectableInterface
	Create(changeAcceptedByMerchant bool) *IPaymentRequestChangedResult
	CreateWithPaymentDetails(changeAcceptedByMerchant bool, updatedPaymentDetails *IPaymentDetails) *IPaymentRequestChangedResult
}

type IPaymentRequestChangedResultFactoryVtbl struct {
	win32.IInspectableVtbl
	Create                   uintptr
	CreateWithPaymentDetails uintptr
}

type IPaymentRequestChangedResultFactory struct {
	win32.IInspectable
}

func (this *IPaymentRequestChangedResultFactory) Vtbl() *IPaymentRequestChangedResultFactoryVtbl {
	return (*IPaymentRequestChangedResultFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentRequestChangedResultFactory) Create(changeAcceptedByMerchant bool) *IPaymentRequestChangedResult {
	var _result *IPaymentRequestChangedResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&changeAcceptedByMerchant))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentRequestChangedResultFactory) CreateWithPaymentDetails(changeAcceptedByMerchant bool, updatedPaymentDetails *IPaymentDetails) *IPaymentRequestChangedResult {
	var _result *IPaymentRequestChangedResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithPaymentDetails, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&changeAcceptedByMerchant))), uintptr(unsafe.Pointer(updatedPaymentDetails)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3E8A79DC-6B74-42D3-B103-F0DE35FB1848
var IID_IPaymentRequestFactory = syscall.GUID{0x3E8A79DC, 0x6B74, 0x42D3,
	[8]byte{0xB1, 0x03, 0xF0, 0xDE, 0x35, 0xFB, 0x18, 0x48}}

type IPaymentRequestFactoryInterface interface {
	win32.IInspectableInterface
	Create(details *IPaymentDetails, methodData *IIterable[*IPaymentMethodData]) *IPaymentRequest
	CreateWithMerchantInfo(details *IPaymentDetails, methodData *IIterable[*IPaymentMethodData], merchantInfo *IPaymentMerchantInfo) *IPaymentRequest
	CreateWithMerchantInfoAndOptions(details *IPaymentDetails, methodData *IIterable[*IPaymentMethodData], merchantInfo *IPaymentMerchantInfo, options *IPaymentOptions) *IPaymentRequest
}

type IPaymentRequestFactoryVtbl struct {
	win32.IInspectableVtbl
	Create                           uintptr
	CreateWithMerchantInfo           uintptr
	CreateWithMerchantInfoAndOptions uintptr
}

type IPaymentRequestFactory struct {
	win32.IInspectable
}

func (this *IPaymentRequestFactory) Vtbl() *IPaymentRequestFactoryVtbl {
	return (*IPaymentRequestFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentRequestFactory) Create(details *IPaymentDetails, methodData *IIterable[*IPaymentMethodData]) *IPaymentRequest {
	var _result *IPaymentRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(details)), uintptr(unsafe.Pointer(methodData)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentRequestFactory) CreateWithMerchantInfo(details *IPaymentDetails, methodData *IIterable[*IPaymentMethodData], merchantInfo *IPaymentMerchantInfo) *IPaymentRequest {
	var _result *IPaymentRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithMerchantInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(details)), uintptr(unsafe.Pointer(methodData)), uintptr(unsafe.Pointer(merchantInfo)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentRequestFactory) CreateWithMerchantInfoAndOptions(details *IPaymentDetails, methodData *IIterable[*IPaymentMethodData], merchantInfo *IPaymentMerchantInfo, options *IPaymentOptions) *IPaymentRequest {
	var _result *IPaymentRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithMerchantInfoAndOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(details)), uintptr(unsafe.Pointer(methodData)), uintptr(unsafe.Pointer(merchantInfo)), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E6CE1325-A506-4372-B7EF-1A031D5662D1
var IID_IPaymentRequestFactory2 = syscall.GUID{0xE6CE1325, 0xA506, 0x4372,
	[8]byte{0xB7, 0xEF, 0x1A, 0x03, 0x1D, 0x56, 0x62, 0xD1}}

type IPaymentRequestFactory2Interface interface {
	win32.IInspectableInterface
	CreateWithMerchantInfoOptionsAndId(details *IPaymentDetails, methodData *IIterable[*IPaymentMethodData], merchantInfo *IPaymentMerchantInfo, options *IPaymentOptions, id string) *IPaymentRequest
}

type IPaymentRequestFactory2Vtbl struct {
	win32.IInspectableVtbl
	CreateWithMerchantInfoOptionsAndId uintptr
}

type IPaymentRequestFactory2 struct {
	win32.IInspectable
}

func (this *IPaymentRequestFactory2) Vtbl() *IPaymentRequestFactory2Vtbl {
	return (*IPaymentRequestFactory2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentRequestFactory2) CreateWithMerchantInfoOptionsAndId(details *IPaymentDetails, methodData *IIterable[*IPaymentMethodData], merchantInfo *IPaymentMerchantInfo, options *IPaymentOptions, id string) *IPaymentRequest {
	var _result *IPaymentRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithMerchantInfoOptionsAndId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(details)), uintptr(unsafe.Pointer(methodData)), uintptr(unsafe.Pointer(merchantInfo)), uintptr(unsafe.Pointer(options)), NewHStr(id).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7B9C3912-30F2-4E90-B249-8CE7D78FFE56
var IID_IPaymentRequestSubmitResult = syscall.GUID{0x7B9C3912, 0x30F2, 0x4E90,
	[8]byte{0xB2, 0x49, 0x8C, 0xE7, 0xD7, 0x8F, 0xFE, 0x56}}

type IPaymentRequestSubmitResultInterface interface {
	win32.IInspectableInterface
	Get_Status() PaymentRequestStatus
	Get_Response() *IPaymentResponse
}

type IPaymentRequestSubmitResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status   uintptr
	Get_Response uintptr
}

type IPaymentRequestSubmitResult struct {
	win32.IInspectable
}

func (this *IPaymentRequestSubmitResult) Vtbl() *IPaymentRequestSubmitResultVtbl {
	return (*IPaymentRequestSubmitResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentRequestSubmitResult) Get_Status() PaymentRequestStatus {
	var _result PaymentRequestStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPaymentRequestSubmitResult) Get_Response() *IPaymentResponse {
	var _result *IPaymentResponse
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Response, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E1389457-8BD2-4888-9FA8-97985545108E
var IID_IPaymentResponse = syscall.GUID{0xE1389457, 0x8BD2, 0x4888,
	[8]byte{0x9F, 0xA8, 0x97, 0x98, 0x55, 0x45, 0x10, 0x8E}}

type IPaymentResponseInterface interface {
	win32.IInspectableInterface
	Get_PaymentToken() *IPaymentToken
	Get_ShippingOption() *IPaymentShippingOption
	Get_ShippingAddress() *IPaymentAddress
	Get_PayerEmail() string
	Get_PayerName() string
	Get_PayerPhoneNumber() string
	CompleteAsync(status PaymentRequestCompletionStatus) *IAsyncAction
}

type IPaymentResponseVtbl struct {
	win32.IInspectableVtbl
	Get_PaymentToken     uintptr
	Get_ShippingOption   uintptr
	Get_ShippingAddress  uintptr
	Get_PayerEmail       uintptr
	Get_PayerName        uintptr
	Get_PayerPhoneNumber uintptr
	CompleteAsync        uintptr
}

type IPaymentResponse struct {
	win32.IInspectable
}

func (this *IPaymentResponse) Vtbl() *IPaymentResponseVtbl {
	return (*IPaymentResponseVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentResponse) Get_PaymentToken() *IPaymentToken {
	var _result *IPaymentToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PaymentToken, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentResponse) Get_ShippingOption() *IPaymentShippingOption {
	var _result *IPaymentShippingOption
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShippingOption, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentResponse) Get_ShippingAddress() *IPaymentAddress {
	var _result *IPaymentAddress
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShippingAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentResponse) Get_PayerEmail() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PayerEmail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPaymentResponse) Get_PayerName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PayerName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPaymentResponse) Get_PayerPhoneNumber() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PayerPhoneNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPaymentResponse) CompleteAsync(status PaymentRequestCompletionStatus) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CompleteAsync, uintptr(unsafe.Pointer(this)), uintptr(status), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 13372ADA-9753-4574-8966-93145A76C7F9
var IID_IPaymentShippingOption = syscall.GUID{0x13372ADA, 0x9753, 0x4574,
	[8]byte{0x89, 0x66, 0x93, 0x14, 0x5A, 0x76, 0xC7, 0xF9}}

type IPaymentShippingOptionInterface interface {
	win32.IInspectableInterface
	Get_Label() string
	Put_Label(value string)
	Get_Amount() *IPaymentCurrencyAmount
	Put_Amount(value *IPaymentCurrencyAmount)
	Get_Tag() string
	Put_Tag(value string)
	Get_IsSelected() bool
	Put_IsSelected(value bool)
}

type IPaymentShippingOptionVtbl struct {
	win32.IInspectableVtbl
	Get_Label      uintptr
	Put_Label      uintptr
	Get_Amount     uintptr
	Put_Amount     uintptr
	Get_Tag        uintptr
	Put_Tag        uintptr
	Get_IsSelected uintptr
	Put_IsSelected uintptr
}

type IPaymentShippingOption struct {
	win32.IInspectable
}

func (this *IPaymentShippingOption) Vtbl() *IPaymentShippingOptionVtbl {
	return (*IPaymentShippingOptionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentShippingOption) Get_Label() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Label, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPaymentShippingOption) Put_Label(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Label, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPaymentShippingOption) Get_Amount() *IPaymentCurrencyAmount {
	var _result *IPaymentCurrencyAmount
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Amount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentShippingOption) Put_Amount(value *IPaymentCurrencyAmount) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Amount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPaymentShippingOption) Get_Tag() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Tag, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPaymentShippingOption) Put_Tag(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Tag, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPaymentShippingOption) Get_IsSelected() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSelected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPaymentShippingOption) Put_IsSelected(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsSelected, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 5DE5F917-B2D7-446B-9D73-6123FBCA3BC6
var IID_IPaymentShippingOptionFactory = syscall.GUID{0x5DE5F917, 0xB2D7, 0x446B,
	[8]byte{0x9D, 0x73, 0x61, 0x23, 0xFB, 0xCA, 0x3B, 0xC6}}

type IPaymentShippingOptionFactoryInterface interface {
	win32.IInspectableInterface
	Create(label string, amount *IPaymentCurrencyAmount) *IPaymentShippingOption
	CreateWithSelected(label string, amount *IPaymentCurrencyAmount, selected bool) *IPaymentShippingOption
	CreateWithSelectedAndTag(label string, amount *IPaymentCurrencyAmount, selected bool, tag string) *IPaymentShippingOption
}

type IPaymentShippingOptionFactoryVtbl struct {
	win32.IInspectableVtbl
	Create                   uintptr
	CreateWithSelected       uintptr
	CreateWithSelectedAndTag uintptr
}

type IPaymentShippingOptionFactory struct {
	win32.IInspectable
}

func (this *IPaymentShippingOptionFactory) Vtbl() *IPaymentShippingOptionFactoryVtbl {
	return (*IPaymentShippingOptionFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentShippingOptionFactory) Create(label string, amount *IPaymentCurrencyAmount) *IPaymentShippingOption {
	var _result *IPaymentShippingOption
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(label).Ptr, uintptr(unsafe.Pointer(amount)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentShippingOptionFactory) CreateWithSelected(label string, amount *IPaymentCurrencyAmount, selected bool) *IPaymentShippingOption {
	var _result *IPaymentShippingOption
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithSelected, uintptr(unsafe.Pointer(this)), NewHStr(label).Ptr, uintptr(unsafe.Pointer(amount)), uintptr(*(*byte)(unsafe.Pointer(&selected))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentShippingOptionFactory) CreateWithSelectedAndTag(label string, amount *IPaymentCurrencyAmount, selected bool, tag string) *IPaymentShippingOption {
	var _result *IPaymentShippingOption
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithSelectedAndTag, uintptr(unsafe.Pointer(this)), NewHStr(label).Ptr, uintptr(unsafe.Pointer(amount)), uintptr(*(*byte)(unsafe.Pointer(&selected))), NewHStr(tag).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BBCAC013-CCD0-41F2-B2A1-0A2E4B5DCE25
var IID_IPaymentToken = syscall.GUID{0xBBCAC013, 0xCCD0, 0x41F2,
	[8]byte{0xB2, 0xA1, 0x0A, 0x2E, 0x4B, 0x5D, 0xCE, 0x25}}

type IPaymentTokenInterface interface {
	win32.IInspectableInterface
	Get_PaymentMethodId() string
	Get_JsonDetails() string
}

type IPaymentTokenVtbl struct {
	win32.IInspectableVtbl
	Get_PaymentMethodId uintptr
	Get_JsonDetails     uintptr
}

type IPaymentToken struct {
	win32.IInspectable
}

func (this *IPaymentToken) Vtbl() *IPaymentTokenVtbl {
	return (*IPaymentTokenVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentToken) Get_PaymentMethodId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PaymentMethodId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPaymentToken) Get_JsonDetails() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_JsonDetails, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 988CD7AA-4753-4904-8373-DD7B08B995C1
var IID_IPaymentTokenFactory = syscall.GUID{0x988CD7AA, 0x4753, 0x4904,
	[8]byte{0x83, 0x73, 0xDD, 0x7B, 0x08, 0xB9, 0x95, 0xC1}}

type IPaymentTokenFactoryInterface interface {
	win32.IInspectableInterface
	Create(paymentMethodId string) *IPaymentToken
	CreateWithJsonDetails(paymentMethodId string, jsonDetails string) *IPaymentToken
}

type IPaymentTokenFactoryVtbl struct {
	win32.IInspectableVtbl
	Create                uintptr
	CreateWithJsonDetails uintptr
}

type IPaymentTokenFactory struct {
	win32.IInspectable
}

func (this *IPaymentTokenFactory) Vtbl() *IPaymentTokenFactoryVtbl {
	return (*IPaymentTokenFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPaymentTokenFactory) Create(paymentMethodId string) *IPaymentToken {
	var _result *IPaymentToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(paymentMethodId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPaymentTokenFactory) CreateWithJsonDetails(paymentMethodId string, jsonDetails string) *IPaymentToken {
	var _result *IPaymentToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithJsonDetails, uintptr(unsafe.Pointer(this)), NewHStr(paymentMethodId).Ptr, NewHStr(jsonDetails).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type PaymentMediator struct {
	RtClass
	*IPaymentMediator
}

func NewPaymentMediator() *PaymentMediator {
	hs := NewHStr("Windows.ApplicationModel.Payments.PaymentMediator")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &PaymentMediator{
		RtClass:          RtClass{PInspect: p},
		IPaymentMediator: (*IPaymentMediator)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}
