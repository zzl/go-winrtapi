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
type CellularClass int32

const (
	CellularClass_None CellularClass = 0
	CellularClass_Gsm  CellularClass = 1
	CellularClass_Cdma CellularClass = 2
)

// enum
type SmsBroadcastType int32

const (
	SmsBroadcastType_Other                    SmsBroadcastType = 0
	SmsBroadcastType_CmasPresidential         SmsBroadcastType = 1
	SmsBroadcastType_CmasExtreme              SmsBroadcastType = 2
	SmsBroadcastType_CmasSevere               SmsBroadcastType = 3
	SmsBroadcastType_CmasAmber                SmsBroadcastType = 4
	SmsBroadcastType_CmasTest                 SmsBroadcastType = 5
	SmsBroadcastType_EUAlert1                 SmsBroadcastType = 6
	SmsBroadcastType_EUAlert2                 SmsBroadcastType = 7
	SmsBroadcastType_EUAlert3                 SmsBroadcastType = 8
	SmsBroadcastType_EUAlertAmber             SmsBroadcastType = 9
	SmsBroadcastType_EUAlertInfo              SmsBroadcastType = 10
	SmsBroadcastType_EtwsEarthquake           SmsBroadcastType = 11
	SmsBroadcastType_EtwsTsunami              SmsBroadcastType = 12
	SmsBroadcastType_EtwsTsunamiAndEarthquake SmsBroadcastType = 13
	SmsBroadcastType_LatAlertLocal            SmsBroadcastType = 14
)

// enum
type SmsDataFormat int32

const (
	SmsDataFormat_Unknown     SmsDataFormat = 0
	SmsDataFormat_CdmaSubmit  SmsDataFormat = 1
	SmsDataFormat_GsmSubmit   SmsDataFormat = 2
	SmsDataFormat_CdmaDeliver SmsDataFormat = 3
	SmsDataFormat_GsmDeliver  SmsDataFormat = 4
)

// enum
type SmsDeviceStatus int32

const (
	SmsDeviceStatus_Off                      SmsDeviceStatus = 0
	SmsDeviceStatus_Ready                    SmsDeviceStatus = 1
	SmsDeviceStatus_SimNotInserted           SmsDeviceStatus = 2
	SmsDeviceStatus_BadSim                   SmsDeviceStatus = 3
	SmsDeviceStatus_DeviceFailure            SmsDeviceStatus = 4
	SmsDeviceStatus_SubscriptionNotActivated SmsDeviceStatus = 5
	SmsDeviceStatus_DeviceLocked             SmsDeviceStatus = 6
	SmsDeviceStatus_DeviceBlocked            SmsDeviceStatus = 7
)

// enum
type SmsEncoding int32

const (
	SmsEncoding_Unknown       SmsEncoding = 0
	SmsEncoding_Optimal       SmsEncoding = 1
	SmsEncoding_SevenBitAscii SmsEncoding = 2
	SmsEncoding_Unicode       SmsEncoding = 3
	SmsEncoding_GsmSevenBit   SmsEncoding = 4
	SmsEncoding_EightBit      SmsEncoding = 5
	SmsEncoding_Latin         SmsEncoding = 6
	SmsEncoding_Korean        SmsEncoding = 7
	SmsEncoding_IA5           SmsEncoding = 8
	SmsEncoding_ShiftJis      SmsEncoding = 9
	SmsEncoding_LatinHebrew   SmsEncoding = 10
)

// enum
type SmsFilterActionType int32

const (
	SmsFilterActionType_AcceptImmediately SmsFilterActionType = 0
	SmsFilterActionType_Drop              SmsFilterActionType = 1
	SmsFilterActionType_Peek              SmsFilterActionType = 2
	SmsFilterActionType_Accept            SmsFilterActionType = 3
)

// enum
type SmsGeographicalScope int32

const (
	SmsGeographicalScope_None                     SmsGeographicalScope = 0
	SmsGeographicalScope_CellWithImmediateDisplay SmsGeographicalScope = 1
	SmsGeographicalScope_LocationArea             SmsGeographicalScope = 2
	SmsGeographicalScope_Plmn                     SmsGeographicalScope = 3
	SmsGeographicalScope_Cell                     SmsGeographicalScope = 4
)

// enum
type SmsMessageClass int32

const (
	SmsMessageClass_None   SmsMessageClass = 0
	SmsMessageClass_Class0 SmsMessageClass = 1
	SmsMessageClass_Class1 SmsMessageClass = 2
	SmsMessageClass_Class2 SmsMessageClass = 3
	SmsMessageClass_Class3 SmsMessageClass = 4
)

// enum
type SmsMessageFilter int32

const (
	SmsMessageFilter_All    SmsMessageFilter = 0
	SmsMessageFilter_Unread SmsMessageFilter = 1
	SmsMessageFilter_Read   SmsMessageFilter = 2
	SmsMessageFilter_Sent   SmsMessageFilter = 3
	SmsMessageFilter_Draft  SmsMessageFilter = 4
)

// enum
type SmsMessageType int32

const (
	SmsMessageType_Binary    SmsMessageType = 0
	SmsMessageType_Text      SmsMessageType = 1
	SmsMessageType_Wap       SmsMessageType = 2
	SmsMessageType_App       SmsMessageType = 3
	SmsMessageType_Broadcast SmsMessageType = 4
	SmsMessageType_Voicemail SmsMessageType = 5
	SmsMessageType_Status    SmsMessageType = 6
)

// enum
type SmsModemErrorCode int32

const (
	SmsModemErrorCode_Other                            SmsModemErrorCode = 0
	SmsModemErrorCode_MessagingNetworkError            SmsModemErrorCode = 1
	SmsModemErrorCode_SmsOperationNotSupportedByDevice SmsModemErrorCode = 2
	SmsModemErrorCode_SmsServiceNotSupportedByNetwork  SmsModemErrorCode = 3
	SmsModemErrorCode_DeviceFailure                    SmsModemErrorCode = 4
	SmsModemErrorCode_MessageNotEncodedProperly        SmsModemErrorCode = 5
	SmsModemErrorCode_MessageTooLarge                  SmsModemErrorCode = 6
	SmsModemErrorCode_DeviceNotReady                   SmsModemErrorCode = 7
	SmsModemErrorCode_NetworkNotReady                  SmsModemErrorCode = 8
	SmsModemErrorCode_InvalidSmscAddress               SmsModemErrorCode = 9
	SmsModemErrorCode_NetworkFailure                   SmsModemErrorCode = 10
	SmsModemErrorCode_FixedDialingNumberRestricted     SmsModemErrorCode = 11
)

// structs

type LegacySmsApiContract struct {
}

type SmsEncodedLength struct {
	SegmentCount              uint32
	CharacterCountLastSegment uint32
	CharactersPerSegment      uint32
	ByteCountLastSegment      uint32
	BytesPerSegment           uint32
}

// func types

// 982B1162-3DD7-4618-AF89-0C272D5D06D8
type SmsDeviceStatusChangedEventHandler func(sender *ISmsDevice) com.Error

// 0B7AD409-EC2D-47CE-A253-732BEEEBCACD
type SmsMessageReceivedEventHandler func(sender *ISmsDevice, e *ISmsMessageReceivedEventArgs) com.Error

// interfaces

// E8BB8494-D3A0-4A0A-86D7-291033A8CF54
var IID_ISmsAppMessage = syscall.GUID{0xE8BB8494, 0xD3A0, 0x4A0A,
	[8]byte{0x86, 0xD7, 0x29, 0x10, 0x33, 0xA8, 0xCF, 0x54}}

type ISmsAppMessageInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() DateTime
	Get_To() string
	Put_To(value string)
	Get_From() string
	Get_Body() string
	Put_Body(value string)
	Get_CallbackNumber() string
	Put_CallbackNumber(value string)
	Get_IsDeliveryNotificationEnabled() bool
	Put_IsDeliveryNotificationEnabled(value bool)
	Get_RetryAttemptCount() int32
	Put_RetryAttemptCount(value int32)
	Get_Encoding() SmsEncoding
	Put_Encoding(value SmsEncoding)
	Get_PortNumber() int32
	Put_PortNumber(value int32)
	Get_TeleserviceId() int32
	Put_TeleserviceId(value int32)
	Get_ProtocolId() int32
	Put_ProtocolId(value int32)
	Get_BinaryBody() *IBuffer
	Put_BinaryBody(value *IBuffer)
}

type ISmsAppMessageVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp                     uintptr
	Get_To                            uintptr
	Put_To                            uintptr
	Get_From                          uintptr
	Get_Body                          uintptr
	Put_Body                          uintptr
	Get_CallbackNumber                uintptr
	Put_CallbackNumber                uintptr
	Get_IsDeliveryNotificationEnabled uintptr
	Put_IsDeliveryNotificationEnabled uintptr
	Get_RetryAttemptCount             uintptr
	Put_RetryAttemptCount             uintptr
	Get_Encoding                      uintptr
	Put_Encoding                      uintptr
	Get_PortNumber                    uintptr
	Put_PortNumber                    uintptr
	Get_TeleserviceId                 uintptr
	Put_TeleserviceId                 uintptr
	Get_ProtocolId                    uintptr
	Put_ProtocolId                    uintptr
	Get_BinaryBody                    uintptr
	Put_BinaryBody                    uintptr
}

type ISmsAppMessage struct {
	win32.IInspectable
}

func (this *ISmsAppMessage) Vtbl() *ISmsAppMessageVtbl {
	return (*ISmsAppMessageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsAppMessage) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsAppMessage) Get_To() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_To, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsAppMessage) Put_To(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_To, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISmsAppMessage) Get_From() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_From, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsAppMessage) Get_Body() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Body, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsAppMessage) Put_Body(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Body, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISmsAppMessage) Get_CallbackNumber() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CallbackNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsAppMessage) Put_CallbackNumber(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CallbackNumber, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISmsAppMessage) Get_IsDeliveryNotificationEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDeliveryNotificationEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsAppMessage) Put_IsDeliveryNotificationEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsDeliveryNotificationEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ISmsAppMessage) Get_RetryAttemptCount() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RetryAttemptCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsAppMessage) Put_RetryAttemptCount(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RetryAttemptCount, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISmsAppMessage) Get_Encoding() SmsEncoding {
	var _result SmsEncoding
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Encoding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsAppMessage) Put_Encoding(value SmsEncoding) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Encoding, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISmsAppMessage) Get_PortNumber() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PortNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsAppMessage) Put_PortNumber(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PortNumber, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISmsAppMessage) Get_TeleserviceId() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TeleserviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsAppMessage) Put_TeleserviceId(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TeleserviceId, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISmsAppMessage) Get_ProtocolId() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtocolId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsAppMessage) Put_ProtocolId(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProtocolId, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISmsAppMessage) Get_BinaryBody() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BinaryBody, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsAppMessage) Put_BinaryBody(value *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BinaryBody, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 5BF4E813-3B53-4C6E-B61A-D86A63755650
var IID_ISmsBinaryMessage = syscall.GUID{0x5BF4E813, 0x3B53, 0x4C6E,
	[8]byte{0xB6, 0x1A, 0xD8, 0x6A, 0x63, 0x75, 0x56, 0x50}}

type ISmsBinaryMessageInterface interface {
	win32.IInspectableInterface
	Get_Format() SmsDataFormat
	Put_Format(value SmsDataFormat)
	GetData() []byte
	SetData(valueLength uint32, value *byte)
}

type ISmsBinaryMessageVtbl struct {
	win32.IInspectableVtbl
	Get_Format uintptr
	Put_Format uintptr
	GetData    uintptr
	SetData    uintptr
}

type ISmsBinaryMessage struct {
	win32.IInspectable
}

func (this *ISmsBinaryMessage) Vtbl() *ISmsBinaryMessageVtbl {
	return (*ISmsBinaryMessageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsBinaryMessage) Get_Format() SmsDataFormat {
	var _result SmsDataFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Format, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsBinaryMessage) Put_Format(value SmsDataFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Format, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISmsBinaryMessage) GetData() []byte {
	var _result []byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsBinaryMessage) SetData(valueLength uint32, value *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetData, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 75AEBBF1-E4B7-4874-A09C-2956E592F957
var IID_ISmsBroadcastMessage = syscall.GUID{0x75AEBBF1, 0xE4B7, 0x4874,
	[8]byte{0xA0, 0x9C, 0x29, 0x56, 0xE5, 0x92, 0xF9, 0x57}}

type ISmsBroadcastMessageInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() DateTime
	Get_To() string
	Get_Body() string
	Get_Channel() int32
	Get_GeographicalScope() SmsGeographicalScope
	Get_MessageCode() int32
	Get_UpdateNumber() int32
	Get_BroadcastType() SmsBroadcastType
	Get_IsEmergencyAlert() bool
	Get_IsUserPopupRequested() bool
}

type ISmsBroadcastMessageVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp            uintptr
	Get_To                   uintptr
	Get_Body                 uintptr
	Get_Channel              uintptr
	Get_GeographicalScope    uintptr
	Get_MessageCode          uintptr
	Get_UpdateNumber         uintptr
	Get_BroadcastType        uintptr
	Get_IsEmergencyAlert     uintptr
	Get_IsUserPopupRequested uintptr
}

type ISmsBroadcastMessage struct {
	win32.IInspectable
}

func (this *ISmsBroadcastMessage) Vtbl() *ISmsBroadcastMessageVtbl {
	return (*ISmsBroadcastMessageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsBroadcastMessage) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsBroadcastMessage) Get_To() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_To, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsBroadcastMessage) Get_Body() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Body, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsBroadcastMessage) Get_Channel() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Channel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsBroadcastMessage) Get_GeographicalScope() SmsGeographicalScope {
	var _result SmsGeographicalScope
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GeographicalScope, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsBroadcastMessage) Get_MessageCode() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MessageCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsBroadcastMessage) Get_UpdateNumber() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UpdateNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsBroadcastMessage) Get_BroadcastType() SmsBroadcastType {
	var _result SmsBroadcastType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BroadcastType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsBroadcastMessage) Get_IsEmergencyAlert() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEmergencyAlert, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsBroadcastMessage) Get_IsUserPopupRequested() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsUserPopupRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 091791ED-872B-4EEC-9C72-AB11627B34EC
var IID_ISmsDevice = syscall.GUID{0x091791ED, 0x872B, 0x4EEC,
	[8]byte{0x9C, 0x72, 0xAB, 0x11, 0x62, 0x7B, 0x34, 0xEC}}

type ISmsDeviceInterface interface {
	win32.IInspectableInterface
	SendMessageAsync(message *ISmsMessage) *IAsyncAction
	CalculateLength(message *ISmsTextMessage) SmsEncodedLength
	Get_AccountPhoneNumber() string
	Get_CellularClass() CellularClass
	Get_MessageStore() *ISmsDeviceMessageStore
	Get_DeviceStatus() SmsDeviceStatus
	Add_SmsMessageReceived(eventHandler SmsMessageReceivedEventHandler) EventRegistrationToken
	Remove_SmsMessageReceived(eventCookie EventRegistrationToken)
	Add_SmsDeviceStatusChanged(eventHandler SmsDeviceStatusChangedEventHandler) EventRegistrationToken
	Remove_SmsDeviceStatusChanged(eventCookie EventRegistrationToken)
}

type ISmsDeviceVtbl struct {
	win32.IInspectableVtbl
	SendMessageAsync              uintptr
	CalculateLength               uintptr
	Get_AccountPhoneNumber        uintptr
	Get_CellularClass             uintptr
	Get_MessageStore              uintptr
	Get_DeviceStatus              uintptr
	Add_SmsMessageReceived        uintptr
	Remove_SmsMessageReceived     uintptr
	Add_SmsDeviceStatusChanged    uintptr
	Remove_SmsDeviceStatusChanged uintptr
}

type ISmsDevice struct {
	win32.IInspectable
}

func (this *ISmsDevice) Vtbl() *ISmsDeviceVtbl {
	return (*ISmsDeviceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsDevice) SendMessageAsync(message *ISmsMessage) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendMessageAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(message)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsDevice) CalculateLength(message *ISmsTextMessage) SmsEncodedLength {
	var _result SmsEncodedLength
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CalculateLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(message)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsDevice) Get_AccountPhoneNumber() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AccountPhoneNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsDevice) Get_CellularClass() CellularClass {
	var _result CellularClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CellularClass, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsDevice) Get_MessageStore() *ISmsDeviceMessageStore {
	var _result *ISmsDeviceMessageStore
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MessageStore, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsDevice) Get_DeviceStatus() SmsDeviceStatus {
	var _result SmsDeviceStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsDevice) Add_SmsMessageReceived(eventHandler SmsMessageReceivedEventHandler) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SmsMessageReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(eventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsDevice) Remove_SmsMessageReceived(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SmsMessageReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

func (this *ISmsDevice) Add_SmsDeviceStatusChanged(eventHandler SmsDeviceStatusChangedEventHandler) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SmsDeviceStatusChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewOneArgFuncDelegate(eventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsDevice) Remove_SmsDeviceStatusChanged(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SmsDeviceStatusChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

// BD8A5C13-E522-46CB-B8D5-9EAD30FB6C47
var IID_ISmsDevice2 = syscall.GUID{0xBD8A5C13, 0xE522, 0x46CB,
	[8]byte{0xB8, 0xD5, 0x9E, 0xAD, 0x30, 0xFB, 0x6C, 0x47}}

type ISmsDevice2Interface interface {
	win32.IInspectableInterface
	Get_SmscAddress() string
	Put_SmscAddress(value string)
	Get_DeviceId() string
	Get_ParentDeviceId() string
	Get_AccountPhoneNumber() string
	Get_CellularClass() CellularClass
	Get_DeviceStatus() SmsDeviceStatus
	CalculateLength(message *ISmsMessageBase) SmsEncodedLength
	SendMessageAndGetResultAsync(message *ISmsMessageBase) *IAsyncOperation[*ISmsSendMessageResult]
	Add_DeviceStatusChanged(eventHandler TypedEventHandler[*ISmsDevice2, interface{}]) EventRegistrationToken
	Remove_DeviceStatusChanged(eventCookie EventRegistrationToken)
}

type ISmsDevice2Vtbl struct {
	win32.IInspectableVtbl
	Get_SmscAddress              uintptr
	Put_SmscAddress              uintptr
	Get_DeviceId                 uintptr
	Get_ParentDeviceId           uintptr
	Get_AccountPhoneNumber       uintptr
	Get_CellularClass            uintptr
	Get_DeviceStatus             uintptr
	CalculateLength              uintptr
	SendMessageAndGetResultAsync uintptr
	Add_DeviceStatusChanged      uintptr
	Remove_DeviceStatusChanged   uintptr
}

type ISmsDevice2 struct {
	win32.IInspectable
}

func (this *ISmsDevice2) Vtbl() *ISmsDevice2Vtbl {
	return (*ISmsDevice2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsDevice2) Get_SmscAddress() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SmscAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsDevice2) Put_SmscAddress(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SmscAddress, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISmsDevice2) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsDevice2) Get_ParentDeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ParentDeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsDevice2) Get_AccountPhoneNumber() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AccountPhoneNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsDevice2) Get_CellularClass() CellularClass {
	var _result CellularClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CellularClass, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsDevice2) Get_DeviceStatus() SmsDeviceStatus {
	var _result SmsDeviceStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsDevice2) CalculateLength(message *ISmsMessageBase) SmsEncodedLength {
	var _result SmsEncodedLength
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CalculateLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(message)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsDevice2) SendMessageAndGetResultAsync(message *ISmsMessageBase) *IAsyncOperation[*ISmsSendMessageResult] {
	var _result *IAsyncOperation[*ISmsSendMessageResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendMessageAndGetResultAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(message)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsDevice2) Add_DeviceStatusChanged(eventHandler TypedEventHandler[*ISmsDevice2, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DeviceStatusChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(eventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsDevice2) Remove_DeviceStatusChanged(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DeviceStatusChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

// 65C78325-1031-491E-8FB6-EF9991AFE363
var IID_ISmsDevice2Statics = syscall.GUID{0x65C78325, 0x1031, 0x491E,
	[8]byte{0x8F, 0xB6, 0xEF, 0x99, 0x91, 0xAF, 0xE3, 0x63}}

type ISmsDevice2StaticsInterface interface {
	win32.IInspectableInterface
	GetDeviceSelector() string
	FromId(deviceId string) *ISmsDevice2
	GetDefault() *ISmsDevice2
	FromParentId(parentDeviceId string) *ISmsDevice2
}

type ISmsDevice2StaticsVtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector uintptr
	FromId            uintptr
	GetDefault        uintptr
	FromParentId      uintptr
}

type ISmsDevice2Statics struct {
	win32.IInspectable
}

func (this *ISmsDevice2Statics) Vtbl() *ISmsDevice2StaticsVtbl {
	return (*ISmsDevice2StaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsDevice2Statics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsDevice2Statics) FromId(deviceId string) *ISmsDevice2 {
	var _result *ISmsDevice2
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromId, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsDevice2Statics) GetDefault() *ISmsDevice2 {
	var _result *ISmsDevice2
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsDevice2Statics) FromParentId(parentDeviceId string) *ISmsDevice2 {
	var _result *ISmsDevice2
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromParentId, uintptr(unsafe.Pointer(this)), NewHStr(parentDeviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9889F253-F188-4427-8D54-CE0C2423C5C1
var IID_ISmsDeviceMessageStore = syscall.GUID{0x9889F253, 0xF188, 0x4427,
	[8]byte{0x8D, 0x54, 0xCE, 0x0C, 0x24, 0x23, 0xC5, 0xC1}}

type ISmsDeviceMessageStoreInterface interface {
	win32.IInspectableInterface
	DeleteMessageAsync(messageId uint32) *IAsyncAction
	DeleteMessagesAsync(messageFilter SmsMessageFilter) *IAsyncAction
	GetMessageAsync(messageId uint32) *IAsyncOperation[*ISmsMessage]
	GetMessagesAsync(messageFilter SmsMessageFilter) *IAsyncOperationWithProgress[*IVectorView[*ISmsMessage], int32]
	Get_MaxMessages() uint32
}

type ISmsDeviceMessageStoreVtbl struct {
	win32.IInspectableVtbl
	DeleteMessageAsync  uintptr
	DeleteMessagesAsync uintptr
	GetMessageAsync     uintptr
	GetMessagesAsync    uintptr
	Get_MaxMessages     uintptr
}

type ISmsDeviceMessageStore struct {
	win32.IInspectable
}

func (this *ISmsDeviceMessageStore) Vtbl() *ISmsDeviceMessageStoreVtbl {
	return (*ISmsDeviceMessageStoreVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsDeviceMessageStore) DeleteMessageAsync(messageId uint32) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DeleteMessageAsync, uintptr(unsafe.Pointer(this)), uintptr(messageId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsDeviceMessageStore) DeleteMessagesAsync(messageFilter SmsMessageFilter) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DeleteMessagesAsync, uintptr(unsafe.Pointer(this)), uintptr(messageFilter), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsDeviceMessageStore) GetMessageAsync(messageId uint32) *IAsyncOperation[*ISmsMessage] {
	var _result *IAsyncOperation[*ISmsMessage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetMessageAsync, uintptr(unsafe.Pointer(this)), uintptr(messageId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsDeviceMessageStore) GetMessagesAsync(messageFilter SmsMessageFilter) *IAsyncOperationWithProgress[*IVectorView[*ISmsMessage], int32] {
	var _result *IAsyncOperationWithProgress[*IVectorView[*ISmsMessage], int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetMessagesAsync, uintptr(unsafe.Pointer(this)), uintptr(messageFilter), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsDeviceMessageStore) Get_MaxMessages() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxMessages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F88D07EA-D815-4DD1-A234-4520CE4604A4
var IID_ISmsDeviceStatics = syscall.GUID{0xF88D07EA, 0xD815, 0x4DD1,
	[8]byte{0xA2, 0x34, 0x45, 0x20, 0xCE, 0x46, 0x04, 0xA4}}

type ISmsDeviceStaticsInterface interface {
	win32.IInspectableInterface
	GetDeviceSelector() string
	FromIdAsync(deviceId string) *IAsyncOperation[*ISmsDevice]
	GetDefaultAsync() *IAsyncOperation[*ISmsDevice]
}

type ISmsDeviceStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector uintptr
	FromIdAsync       uintptr
	GetDefaultAsync   uintptr
}

type ISmsDeviceStatics struct {
	win32.IInspectable
}

func (this *ISmsDeviceStatics) Vtbl() *ISmsDeviceStaticsVtbl {
	return (*ISmsDeviceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsDeviceStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsDeviceStatics) FromIdAsync(deviceId string) *IAsyncOperation[*ISmsDevice] {
	var _result *IAsyncOperation[*ISmsDevice]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsDeviceStatics) GetDefaultAsync() *IAsyncOperation[*ISmsDevice] {
	var _result *IAsyncOperation[*ISmsDevice]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2CA11C87-0873-4CAF-8A7D-BD471E8586D1
var IID_ISmsDeviceStatics2 = syscall.GUID{0x2CA11C87, 0x0873, 0x4CAF,
	[8]byte{0x8A, 0x7D, 0xBD, 0x47, 0x1E, 0x85, 0x86, 0xD1}}

type ISmsDeviceStatics2Interface interface {
	win32.IInspectableInterface
	FromNetworkAccountIdAsync(networkAccountId string) *IAsyncOperation[*ISmsDevice]
}

type ISmsDeviceStatics2Vtbl struct {
	win32.IInspectableVtbl
	FromNetworkAccountIdAsync uintptr
}

type ISmsDeviceStatics2 struct {
	win32.IInspectable
}

func (this *ISmsDeviceStatics2) Vtbl() *ISmsDeviceStatics2Vtbl {
	return (*ISmsDeviceStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsDeviceStatics2) FromNetworkAccountIdAsync(networkAccountId string) *IAsyncOperation[*ISmsDevice] {
	var _result *IAsyncOperation[*ISmsDevice]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromNetworkAccountIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(networkAccountId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 40E32FAE-B049-4FBC-AFE9-E2A610EFF55C
var IID_ISmsFilterRule = syscall.GUID{0x40E32FAE, 0xB049, 0x4FBC,
	[8]byte{0xAF, 0xE9, 0xE2, 0xA6, 0x10, 0xEF, 0xF5, 0x5C}}

type ISmsFilterRuleInterface interface {
	win32.IInspectableInterface
	Get_MessageType() SmsMessageType
	Get_ImsiPrefixes() *IVector[string]
	Get_DeviceIds() *IVector[string]
	Get_SenderNumbers() *IVector[string]
	Get_TextMessagePrefixes() *IVector[string]
	Get_PortNumbers() *IVector[int32]
	Get_CellularClass() CellularClass
	Put_CellularClass(value CellularClass)
	Get_ProtocolIds() *IVector[int32]
	Get_TeleserviceIds() *IVector[int32]
	Get_WapApplicationIds() *IVector[string]
	Get_WapContentTypes() *IVector[string]
	Get_BroadcastTypes() *IVector[SmsBroadcastType]
	Get_BroadcastChannels() *IVector[int32]
}

type ISmsFilterRuleVtbl struct {
	win32.IInspectableVtbl
	Get_MessageType         uintptr
	Get_ImsiPrefixes        uintptr
	Get_DeviceIds           uintptr
	Get_SenderNumbers       uintptr
	Get_TextMessagePrefixes uintptr
	Get_PortNumbers         uintptr
	Get_CellularClass       uintptr
	Put_CellularClass       uintptr
	Get_ProtocolIds         uintptr
	Get_TeleserviceIds      uintptr
	Get_WapApplicationIds   uintptr
	Get_WapContentTypes     uintptr
	Get_BroadcastTypes      uintptr
	Get_BroadcastChannels   uintptr
}

type ISmsFilterRule struct {
	win32.IInspectable
}

func (this *ISmsFilterRule) Vtbl() *ISmsFilterRuleVtbl {
	return (*ISmsFilterRuleVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsFilterRule) Get_MessageType() SmsMessageType {
	var _result SmsMessageType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MessageType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsFilterRule) Get_ImsiPrefixes() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ImsiPrefixes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsFilterRule) Get_DeviceIds() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceIds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsFilterRule) Get_SenderNumbers() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SenderNumbers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsFilterRule) Get_TextMessagePrefixes() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TextMessagePrefixes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsFilterRule) Get_PortNumbers() *IVector[int32] {
	var _result *IVector[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PortNumbers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsFilterRule) Get_CellularClass() CellularClass {
	var _result CellularClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CellularClass, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsFilterRule) Put_CellularClass(value CellularClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CellularClass, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISmsFilterRule) Get_ProtocolIds() *IVector[int32] {
	var _result *IVector[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtocolIds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsFilterRule) Get_TeleserviceIds() *IVector[int32] {
	var _result *IVector[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TeleserviceIds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsFilterRule) Get_WapApplicationIds() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WapApplicationIds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsFilterRule) Get_WapContentTypes() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WapContentTypes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsFilterRule) Get_BroadcastTypes() *IVector[SmsBroadcastType] {
	var _result *IVector[SmsBroadcastType]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BroadcastTypes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsFilterRule) Get_BroadcastChannels() *IVector[int32] {
	var _result *IVector[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BroadcastChannels, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 00C36508-6296-4F29-9AAD-8920CEBA3CE8
var IID_ISmsFilterRuleFactory = syscall.GUID{0x00C36508, 0x6296, 0x4F29,
	[8]byte{0x9A, 0xAD, 0x89, 0x20, 0xCE, 0xBA, 0x3C, 0xE8}}

type ISmsFilterRuleFactoryInterface interface {
	win32.IInspectableInterface
	CreateFilterRule(messageType SmsMessageType) *ISmsFilterRule
}

type ISmsFilterRuleFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateFilterRule uintptr
}

type ISmsFilterRuleFactory struct {
	win32.IInspectable
}

func (this *ISmsFilterRuleFactory) Vtbl() *ISmsFilterRuleFactoryVtbl {
	return (*ISmsFilterRuleFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsFilterRuleFactory) CreateFilterRule(messageType SmsMessageType) *ISmsFilterRule {
	var _result *ISmsFilterRule
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFilterRule, uintptr(unsafe.Pointer(this)), uintptr(messageType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4E47EAFB-79CD-4881-9894-55A4135B23FA
var IID_ISmsFilterRules = syscall.GUID{0x4E47EAFB, 0x79CD, 0x4881,
	[8]byte{0x98, 0x94, 0x55, 0xA4, 0x13, 0x5B, 0x23, 0xFA}}

type ISmsFilterRulesInterface interface {
	win32.IInspectableInterface
	Get_ActionType() SmsFilterActionType
	Get_Rules() *IVector[*ISmsFilterRule]
}

type ISmsFilterRulesVtbl struct {
	win32.IInspectableVtbl
	Get_ActionType uintptr
	Get_Rules      uintptr
}

type ISmsFilterRules struct {
	win32.IInspectable
}

func (this *ISmsFilterRules) Vtbl() *ISmsFilterRulesVtbl {
	return (*ISmsFilterRulesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsFilterRules) Get_ActionType() SmsFilterActionType {
	var _result SmsFilterActionType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActionType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsFilterRules) Get_Rules() *IVector[*ISmsFilterRule] {
	var _result *IVector[*ISmsFilterRule]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Rules, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A09924ED-6E2E-4530-9FDE-465D02EED00E
var IID_ISmsFilterRulesFactory = syscall.GUID{0xA09924ED, 0x6E2E, 0x4530,
	[8]byte{0x9F, 0xDE, 0x46, 0x5D, 0x02, 0xEE, 0xD0, 0x0E}}

type ISmsFilterRulesFactoryInterface interface {
	win32.IInspectableInterface
	CreateFilterRules(actionType SmsFilterActionType) *ISmsFilterRules
}

type ISmsFilterRulesFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateFilterRules uintptr
}

type ISmsFilterRulesFactory struct {
	win32.IInspectable
}

func (this *ISmsFilterRulesFactory) Vtbl() *ISmsFilterRulesFactoryVtbl {
	return (*ISmsFilterRulesFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsFilterRulesFactory) CreateFilterRules(actionType SmsFilterActionType) *ISmsFilterRules {
	var _result *ISmsFilterRules
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFilterRules, uintptr(unsafe.Pointer(this)), uintptr(actionType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// ED3C5E28-6984-4B07-811D-8D5906ED3CEA
var IID_ISmsMessage = syscall.GUID{0xED3C5E28, 0x6984, 0x4B07,
	[8]byte{0x81, 0x1D, 0x8D, 0x59, 0x06, 0xED, 0x3C, 0xEA}}

type ISmsMessageInterface interface {
	win32.IInspectableInterface
	Get_Id() uint32
	Get_MessageClass() SmsMessageClass
}

type ISmsMessageVtbl struct {
	win32.IInspectableVtbl
	Get_Id           uintptr
	Get_MessageClass uintptr
}

type ISmsMessage struct {
	win32.IInspectable
}

func (this *ISmsMessage) Vtbl() *ISmsMessageVtbl {
	return (*ISmsMessageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsMessage) Get_Id() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsMessage) Get_MessageClass() SmsMessageClass {
	var _result SmsMessageClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MessageClass, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 2CF0FE30-FE50-4FC6-AA88-4CCFE27A29EA
var IID_ISmsMessageBase = syscall.GUID{0x2CF0FE30, 0xFE50, 0x4FC6,
	[8]byte{0xAA, 0x88, 0x4C, 0xCF, 0xE2, 0x7A, 0x29, 0xEA}}

type ISmsMessageBaseInterface interface {
	win32.IInspectableInterface
	Get_MessageType() SmsMessageType
	Get_DeviceId() string
	Get_CellularClass() CellularClass
	Get_MessageClass() SmsMessageClass
	Get_SimIccId() string
}

type ISmsMessageBaseVtbl struct {
	win32.IInspectableVtbl
	Get_MessageType   uintptr
	Get_DeviceId      uintptr
	Get_CellularClass uintptr
	Get_MessageClass  uintptr
	Get_SimIccId      uintptr
}

type ISmsMessageBase struct {
	win32.IInspectable
}

func (this *ISmsMessageBase) Vtbl() *ISmsMessageBaseVtbl {
	return (*ISmsMessageBaseVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsMessageBase) Get_MessageType() SmsMessageType {
	var _result SmsMessageType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MessageType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsMessageBase) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsMessageBase) Get_CellularClass() CellularClass {
	var _result CellularClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CellularClass, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsMessageBase) Get_MessageClass() SmsMessageClass {
	var _result SmsMessageClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MessageClass, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsMessageBase) Get_SimIccId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SimIccId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 08E80A98-B8E5-41C1-A3D8-D3ABFAE22675
var IID_ISmsMessageReceivedEventArgs = syscall.GUID{0x08E80A98, 0xB8E5, 0x41C1,
	[8]byte{0xA3, 0xD8, 0xD3, 0xAB, 0xFA, 0xE2, 0x26, 0x75}}

type ISmsMessageReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_TextMessage() *ISmsTextMessage
	Get_BinaryMessage() *ISmsBinaryMessage
}

type ISmsMessageReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_TextMessage   uintptr
	Get_BinaryMessage uintptr
}

type ISmsMessageReceivedEventArgs struct {
	win32.IInspectable
}

func (this *ISmsMessageReceivedEventArgs) Vtbl() *ISmsMessageReceivedEventArgsVtbl {
	return (*ISmsMessageReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsMessageReceivedEventArgs) Get_TextMessage() *ISmsTextMessage {
	var _result *ISmsTextMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TextMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsMessageReceivedEventArgs) Get_BinaryMessage() *ISmsBinaryMessage {
	var _result *ISmsBinaryMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BinaryMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2BCFCBD4-2657-4128-AD5F-E3877132BDB1
var IID_ISmsMessageReceivedTriggerDetails = syscall.GUID{0x2BCFCBD4, 0x2657, 0x4128,
	[8]byte{0xAD, 0x5F, 0xE3, 0x87, 0x71, 0x32, 0xBD, 0xB1}}

type ISmsMessageReceivedTriggerDetailsInterface interface {
	win32.IInspectableInterface
	Get_MessageType() SmsMessageType
	Get_TextMessage() *ISmsTextMessage2
	Get_WapMessage() *ISmsWapMessage
	Get_AppMessage() *ISmsAppMessage
	Get_BroadcastMessage() *ISmsBroadcastMessage
	Get_VoicemailMessage() *ISmsVoicemailMessage
	Get_StatusMessage() *ISmsStatusMessage
	Drop()
	Accept()
}

type ISmsMessageReceivedTriggerDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_MessageType      uintptr
	Get_TextMessage      uintptr
	Get_WapMessage       uintptr
	Get_AppMessage       uintptr
	Get_BroadcastMessage uintptr
	Get_VoicemailMessage uintptr
	Get_StatusMessage    uintptr
	Drop                 uintptr
	Accept               uintptr
}

type ISmsMessageReceivedTriggerDetails struct {
	win32.IInspectable
}

func (this *ISmsMessageReceivedTriggerDetails) Vtbl() *ISmsMessageReceivedTriggerDetailsVtbl {
	return (*ISmsMessageReceivedTriggerDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsMessageReceivedTriggerDetails) Get_MessageType() SmsMessageType {
	var _result SmsMessageType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MessageType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsMessageReceivedTriggerDetails) Get_TextMessage() *ISmsTextMessage2 {
	var _result *ISmsTextMessage2
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TextMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsMessageReceivedTriggerDetails) Get_WapMessage() *ISmsWapMessage {
	var _result *ISmsWapMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WapMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsMessageReceivedTriggerDetails) Get_AppMessage() *ISmsAppMessage {
	var _result *ISmsAppMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsMessageReceivedTriggerDetails) Get_BroadcastMessage() *ISmsBroadcastMessage {
	var _result *ISmsBroadcastMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BroadcastMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsMessageReceivedTriggerDetails) Get_VoicemailMessage() *ISmsVoicemailMessage {
	var _result *ISmsVoicemailMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VoicemailMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsMessageReceivedTriggerDetails) Get_StatusMessage() *ISmsStatusMessage {
	var _result *ISmsStatusMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StatusMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsMessageReceivedTriggerDetails) Drop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Drop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ISmsMessageReceivedTriggerDetails) Accept() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Accept, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 1720503E-F34F-446B-83B3-0FF19923B409
var IID_ISmsMessageRegistration = syscall.GUID{0x1720503E, 0xF34F, 0x446B,
	[8]byte{0x83, 0xB3, 0x0F, 0xF1, 0x99, 0x23, 0xB4, 0x09}}

type ISmsMessageRegistrationInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Unregister()
	Add_MessageReceived(eventHandler TypedEventHandler[*ISmsMessageRegistration, *ISmsMessageReceivedTriggerDetails]) EventRegistrationToken
	Remove_MessageReceived(eventCookie EventRegistrationToken)
}

type ISmsMessageRegistrationVtbl struct {
	win32.IInspectableVtbl
	Get_Id                 uintptr
	Unregister             uintptr
	Add_MessageReceived    uintptr
	Remove_MessageReceived uintptr
}

type ISmsMessageRegistration struct {
	win32.IInspectable
}

func (this *ISmsMessageRegistration) Vtbl() *ISmsMessageRegistrationVtbl {
	return (*ISmsMessageRegistrationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsMessageRegistration) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsMessageRegistration) Unregister() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Unregister, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ISmsMessageRegistration) Add_MessageReceived(eventHandler TypedEventHandler[*ISmsMessageRegistration, *ISmsMessageReceivedTriggerDetails]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_MessageReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(eventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsMessageRegistration) Remove_MessageReceived(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_MessageReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

// 63A05464-2898-4778-A03C-6F994907D63A
var IID_ISmsMessageRegistrationStatics = syscall.GUID{0x63A05464, 0x2898, 0x4778,
	[8]byte{0xA0, 0x3C, 0x6F, 0x99, 0x49, 0x07, 0xD6, 0x3A}}

type ISmsMessageRegistrationStaticsInterface interface {
	win32.IInspectableInterface
	Get_AllRegistrations() *IVectorView[*ISmsMessageRegistration]
	Register(id string, filterRules *ISmsFilterRules) *ISmsMessageRegistration
}

type ISmsMessageRegistrationStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_AllRegistrations uintptr
	Register             uintptr
}

type ISmsMessageRegistrationStatics struct {
	win32.IInspectable
}

func (this *ISmsMessageRegistrationStatics) Vtbl() *ISmsMessageRegistrationStaticsVtbl {
	return (*ISmsMessageRegistrationStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsMessageRegistrationStatics) Get_AllRegistrations() *IVectorView[*ISmsMessageRegistration] {
	var _result *IVectorView[*ISmsMessageRegistration]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllRegistrations, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsMessageRegistrationStatics) Register(id string, filterRules *ISmsFilterRules) *ISmsMessageRegistration {
	var _result *ISmsMessageRegistration
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Register, uintptr(unsafe.Pointer(this)), NewHStr(id).Ptr, uintptr(unsafe.Pointer(filterRules)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5BB50F15-E46D-4C82-847D-5A0304C1D53D
var IID_ISmsReceivedEventDetails = syscall.GUID{0x5BB50F15, 0xE46D, 0x4C82,
	[8]byte{0x84, 0x7D, 0x5A, 0x03, 0x04, 0xC1, 0xD5, 0x3D}}

type ISmsReceivedEventDetailsInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_MessageIndex() uint32
}

type ISmsReceivedEventDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId     uintptr
	Get_MessageIndex uintptr
}

type ISmsReceivedEventDetails struct {
	win32.IInspectable
}

func (this *ISmsReceivedEventDetails) Vtbl() *ISmsReceivedEventDetailsVtbl {
	return (*ISmsReceivedEventDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsReceivedEventDetails) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsReceivedEventDetails) Get_MessageIndex() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MessageIndex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 40E05C86-A7B4-4771-9AE7-0B5FFB12C03A
var IID_ISmsReceivedEventDetails2 = syscall.GUID{0x40E05C86, 0xA7B4, 0x4771,
	[8]byte{0x9A, 0xE7, 0x0B, 0x5F, 0xFB, 0x12, 0xC0, 0x3A}}

type ISmsReceivedEventDetails2Interface interface {
	win32.IInspectableInterface
	Get_MessageClass() SmsMessageClass
	Get_BinaryMessage() *ISmsBinaryMessage
}

type ISmsReceivedEventDetails2Vtbl struct {
	win32.IInspectableVtbl
	Get_MessageClass  uintptr
	Get_BinaryMessage uintptr
}

type ISmsReceivedEventDetails2 struct {
	win32.IInspectable
}

func (this *ISmsReceivedEventDetails2) Vtbl() *ISmsReceivedEventDetails2Vtbl {
	return (*ISmsReceivedEventDetails2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsReceivedEventDetails2) Get_MessageClass() SmsMessageClass {
	var _result SmsMessageClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MessageClass, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsReceivedEventDetails2) Get_BinaryMessage() *ISmsBinaryMessage {
	var _result *ISmsBinaryMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BinaryMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DB139AF2-78C9-4FEB-9622-452328088D62
var IID_ISmsSendMessageResult = syscall.GUID{0xDB139AF2, 0x78C9, 0x4FEB,
	[8]byte{0x96, 0x22, 0x45, 0x23, 0x28, 0x08, 0x8D, 0x62}}

type ISmsSendMessageResultInterface interface {
	win32.IInspectableInterface
	Get_IsSuccessful() bool
	Get_MessageReferenceNumbers() *IVectorView[int32]
	Get_CellularClass() CellularClass
	Get_ModemErrorCode() SmsModemErrorCode
	Get_IsErrorTransient() bool
	Get_NetworkCauseCode() int32
	Get_TransportFailureCause() int32
}

type ISmsSendMessageResultVtbl struct {
	win32.IInspectableVtbl
	Get_IsSuccessful            uintptr
	Get_MessageReferenceNumbers uintptr
	Get_CellularClass           uintptr
	Get_ModemErrorCode          uintptr
	Get_IsErrorTransient        uintptr
	Get_NetworkCauseCode        uintptr
	Get_TransportFailureCause   uintptr
}

type ISmsSendMessageResult struct {
	win32.IInspectable
}

func (this *ISmsSendMessageResult) Vtbl() *ISmsSendMessageResultVtbl {
	return (*ISmsSendMessageResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsSendMessageResult) Get_IsSuccessful() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSuccessful, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsSendMessageResult) Get_MessageReferenceNumbers() *IVectorView[int32] {
	var _result *IVectorView[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MessageReferenceNumbers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsSendMessageResult) Get_CellularClass() CellularClass {
	var _result CellularClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CellularClass, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsSendMessageResult) Get_ModemErrorCode() SmsModemErrorCode {
	var _result SmsModemErrorCode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ModemErrorCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsSendMessageResult) Get_IsErrorTransient() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsErrorTransient, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsSendMessageResult) Get_NetworkCauseCode() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkCauseCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsSendMessageResult) Get_TransportFailureCause() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TransportFailureCause, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// E6D28342-B70B-4677-9379-C9783FDFF8F4
var IID_ISmsStatusMessage = syscall.GUID{0xE6D28342, 0xB70B, 0x4677,
	[8]byte{0x93, 0x79, 0xC9, 0x78, 0x3F, 0xDF, 0xF8, 0xF4}}

type ISmsStatusMessageInterface interface {
	win32.IInspectableInterface
	Get_To() string
	Get_From() string
	Get_Body() string
	Get_Status() int32
	Get_MessageReferenceNumber() int32
	Get_ServiceCenterTimestamp() DateTime
	Get_DischargeTime() DateTime
}

type ISmsStatusMessageVtbl struct {
	win32.IInspectableVtbl
	Get_To                     uintptr
	Get_From                   uintptr
	Get_Body                   uintptr
	Get_Status                 uintptr
	Get_MessageReferenceNumber uintptr
	Get_ServiceCenterTimestamp uintptr
	Get_DischargeTime          uintptr
}

type ISmsStatusMessage struct {
	win32.IInspectable
}

func (this *ISmsStatusMessage) Vtbl() *ISmsStatusMessageVtbl {
	return (*ISmsStatusMessageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsStatusMessage) Get_To() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_To, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsStatusMessage) Get_From() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_From, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsStatusMessage) Get_Body() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Body, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsStatusMessage) Get_Status() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsStatusMessage) Get_MessageReferenceNumber() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MessageReferenceNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsStatusMessage) Get_ServiceCenterTimestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceCenterTimestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsStatusMessage) Get_DischargeTime() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DischargeTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D61C904C-A495-487F-9A6F-971548C5BC9F
var IID_ISmsTextMessage = syscall.GUID{0xD61C904C, 0xA495, 0x487F,
	[8]byte{0x9A, 0x6F, 0x97, 0x15, 0x48, 0xC5, 0xBC, 0x9F}}

type ISmsTextMessageInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() DateTime
	Get_PartReferenceId() uint32
	Get_PartNumber() uint32
	Get_PartCount() uint32
	Get_To() string
	Put_To(value string)
	Get_From() string
	Put_From(value string)
	Get_Body() string
	Put_Body(value string)
	Get_Encoding() SmsEncoding
	Put_Encoding(value SmsEncoding)
	ToBinaryMessages(format SmsDataFormat) *IVectorView[*ISmsBinaryMessage]
}

type ISmsTextMessageVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp       uintptr
	Get_PartReferenceId uintptr
	Get_PartNumber      uintptr
	Get_PartCount       uintptr
	Get_To              uintptr
	Put_To              uintptr
	Get_From            uintptr
	Put_From            uintptr
	Get_Body            uintptr
	Put_Body            uintptr
	Get_Encoding        uintptr
	Put_Encoding        uintptr
	ToBinaryMessages    uintptr
}

type ISmsTextMessage struct {
	win32.IInspectable
}

func (this *ISmsTextMessage) Vtbl() *ISmsTextMessageVtbl {
	return (*ISmsTextMessageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsTextMessage) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsTextMessage) Get_PartReferenceId() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PartReferenceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsTextMessage) Get_PartNumber() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PartNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsTextMessage) Get_PartCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PartCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsTextMessage) Get_To() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_To, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsTextMessage) Put_To(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_To, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISmsTextMessage) Get_From() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_From, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsTextMessage) Put_From(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_From, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISmsTextMessage) Get_Body() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Body, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsTextMessage) Put_Body(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Body, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISmsTextMessage) Get_Encoding() SmsEncoding {
	var _result SmsEncoding
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Encoding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsTextMessage) Put_Encoding(value SmsEncoding) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Encoding, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISmsTextMessage) ToBinaryMessages(format SmsDataFormat) *IVectorView[*ISmsBinaryMessage] {
	var _result *IVectorView[*ISmsBinaryMessage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ToBinaryMessages, uintptr(unsafe.Pointer(this)), uintptr(format), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 22A0D893-4555-4755-B5A1-E7FD84955F8D
var IID_ISmsTextMessage2 = syscall.GUID{0x22A0D893, 0x4555, 0x4755,
	[8]byte{0xB5, 0xA1, 0xE7, 0xFD, 0x84, 0x95, 0x5F, 0x8D}}

type ISmsTextMessage2Interface interface {
	win32.IInspectableInterface
	Get_Timestamp() DateTime
	Get_To() string
	Put_To(value string)
	Get_From() string
	Get_Body() string
	Put_Body(value string)
	Get_Encoding() SmsEncoding
	Put_Encoding(value SmsEncoding)
	Get_CallbackNumber() string
	Put_CallbackNumber(value string)
	Get_IsDeliveryNotificationEnabled() bool
	Put_IsDeliveryNotificationEnabled(value bool)
	Get_RetryAttemptCount() int32
	Put_RetryAttemptCount(value int32)
	Get_TeleserviceId() int32
	Get_ProtocolId() int32
}

type ISmsTextMessage2Vtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp                     uintptr
	Get_To                            uintptr
	Put_To                            uintptr
	Get_From                          uintptr
	Get_Body                          uintptr
	Put_Body                          uintptr
	Get_Encoding                      uintptr
	Put_Encoding                      uintptr
	Get_CallbackNumber                uintptr
	Put_CallbackNumber                uintptr
	Get_IsDeliveryNotificationEnabled uintptr
	Put_IsDeliveryNotificationEnabled uintptr
	Get_RetryAttemptCount             uintptr
	Put_RetryAttemptCount             uintptr
	Get_TeleserviceId                 uintptr
	Get_ProtocolId                    uintptr
}

type ISmsTextMessage2 struct {
	win32.IInspectable
}

func (this *ISmsTextMessage2) Vtbl() *ISmsTextMessage2Vtbl {
	return (*ISmsTextMessage2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsTextMessage2) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsTextMessage2) Get_To() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_To, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsTextMessage2) Put_To(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_To, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISmsTextMessage2) Get_From() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_From, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsTextMessage2) Get_Body() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Body, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsTextMessage2) Put_Body(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Body, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISmsTextMessage2) Get_Encoding() SmsEncoding {
	var _result SmsEncoding
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Encoding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsTextMessage2) Put_Encoding(value SmsEncoding) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Encoding, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISmsTextMessage2) Get_CallbackNumber() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CallbackNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsTextMessage2) Put_CallbackNumber(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CallbackNumber, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISmsTextMessage2) Get_IsDeliveryNotificationEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDeliveryNotificationEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsTextMessage2) Put_IsDeliveryNotificationEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsDeliveryNotificationEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ISmsTextMessage2) Get_RetryAttemptCount() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RetryAttemptCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsTextMessage2) Put_RetryAttemptCount(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RetryAttemptCount, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISmsTextMessage2) Get_TeleserviceId() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TeleserviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsTextMessage2) Get_ProtocolId() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtocolId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7F68C5ED-3CCC-47A3-8C55-380D3B010892
var IID_ISmsTextMessageStatics = syscall.GUID{0x7F68C5ED, 0x3CCC, 0x47A3,
	[8]byte{0x8C, 0x55, 0x38, 0x0D, 0x3B, 0x01, 0x08, 0x92}}

type ISmsTextMessageStaticsInterface interface {
	win32.IInspectableInterface
	FromBinaryMessage(binaryMessage *ISmsBinaryMessage) *ISmsTextMessage
	FromBinaryData(format SmsDataFormat, valueLength uint32, value *byte) *ISmsTextMessage
}

type ISmsTextMessageStaticsVtbl struct {
	win32.IInspectableVtbl
	FromBinaryMessage uintptr
	FromBinaryData    uintptr
}

type ISmsTextMessageStatics struct {
	win32.IInspectable
}

func (this *ISmsTextMessageStatics) Vtbl() *ISmsTextMessageStaticsVtbl {
	return (*ISmsTextMessageStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsTextMessageStatics) FromBinaryMessage(binaryMessage *ISmsBinaryMessage) *ISmsTextMessage {
	var _result *ISmsTextMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromBinaryMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(binaryMessage)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsTextMessageStatics) FromBinaryData(format SmsDataFormat, valueLength uint32, value *byte) *ISmsTextMessage {
	var _result *ISmsTextMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromBinaryData, uintptr(unsafe.Pointer(this)), uintptr(format), uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 271AA0A6-95B1-44FF-BCB8-B8FDD7E08BC3
var IID_ISmsVoicemailMessage = syscall.GUID{0x271AA0A6, 0x95B1, 0x44FF,
	[8]byte{0xBC, 0xB8, 0xB8, 0xFD, 0xD7, 0xE0, 0x8B, 0xC3}}

type ISmsVoicemailMessageInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() DateTime
	Get_To() string
	Get_Body() string
	Get_MessageCount() *IReference[int32]
}

type ISmsVoicemailMessageVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp    uintptr
	Get_To           uintptr
	Get_Body         uintptr
	Get_MessageCount uintptr
}

type ISmsVoicemailMessage struct {
	win32.IInspectable
}

func (this *ISmsVoicemailMessage) Vtbl() *ISmsVoicemailMessageVtbl {
	return (*ISmsVoicemailMessageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsVoicemailMessage) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsVoicemailMessage) Get_To() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_To, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsVoicemailMessage) Get_Body() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Body, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsVoicemailMessage) Get_MessageCount() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MessageCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CD937743-7A55-4D3B-9021-F22E022D09C5
var IID_ISmsWapMessage = syscall.GUID{0xCD937743, 0x7A55, 0x4D3B,
	[8]byte{0x90, 0x21, 0xF2, 0x2E, 0x02, 0x2D, 0x09, 0xC5}}

type ISmsWapMessageInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() DateTime
	Get_To() string
	Get_From() string
	Get_ApplicationId() string
	Get_ContentType() string
	Get_BinaryBody() *IBuffer
	Get_Headers() *IMap[string, string]
}

type ISmsWapMessageVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp     uintptr
	Get_To            uintptr
	Get_From          uintptr
	Get_ApplicationId uintptr
	Get_ContentType   uintptr
	Get_BinaryBody    uintptr
	Get_Headers       uintptr
}

type ISmsWapMessage struct {
	win32.IInspectable
}

func (this *ISmsWapMessage) Vtbl() *ISmsWapMessageVtbl {
	return (*ISmsWapMessageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsWapMessage) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISmsWapMessage) Get_To() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_To, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsWapMessage) Get_From() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_From, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsWapMessage) Get_ApplicationId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ApplicationId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsWapMessage) Get_ContentType() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISmsWapMessage) Get_BinaryBody() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BinaryBody, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISmsWapMessage) Get_Headers() *IMap[string, string] {
	var _result *IMap[string, string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Headers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type SmsAppMessage struct {
	RtClass
	*ISmsAppMessage
}

func NewSmsAppMessage() *SmsAppMessage {
	hs := NewHStr("Windows.Devices.Sms.SmsAppMessage")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &SmsAppMessage{
		RtClass:        RtClass{PInspect: p},
		ISmsAppMessage: (*ISmsAppMessage)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type SmsBinaryMessage struct {
	RtClass
	*ISmsBinaryMessage
}

func NewSmsBinaryMessage() *SmsBinaryMessage {
	hs := NewHStr("Windows.Devices.Sms.SmsBinaryMessage")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &SmsBinaryMessage{
		RtClass:           RtClass{PInspect: p},
		ISmsBinaryMessage: (*ISmsBinaryMessage)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type SmsBroadcastMessage struct {
	RtClass
	*ISmsBroadcastMessage
}

type SmsDevice struct {
	RtClass
	*ISmsDevice
}

func NewISmsDeviceStatics2() *ISmsDeviceStatics2 {
	var p *ISmsDeviceStatics2
	hs := NewHStr("Windows.Devices.Sms.SmsDevice")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISmsDeviceStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewISmsDeviceStatics() *ISmsDeviceStatics {
	var p *ISmsDeviceStatics
	hs := NewHStr("Windows.Devices.Sms.SmsDevice")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISmsDeviceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SmsDevice2 struct {
	RtClass
	*ISmsDevice2
}

func NewISmsDevice2Statics() *ISmsDevice2Statics {
	var p *ISmsDevice2Statics
	hs := NewHStr("Windows.Devices.Sms.SmsDevice2")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISmsDevice2Statics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SmsDeviceMessageStore struct {
	RtClass
	*ISmsDeviceMessageStore
}

type SmsFilterRule struct {
	RtClass
	*ISmsFilterRule
}

func NewSmsFilterRule_CreateFilterRule(messageType SmsMessageType) *SmsFilterRule {
	hs := NewHStr("Windows.Devices.Sms.SmsFilterRule")
	var pFac *ISmsFilterRuleFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISmsFilterRuleFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISmsFilterRule
	p = pFac.CreateFilterRule(messageType)
	result := &SmsFilterRule{
		RtClass:        RtClass{PInspect: &p.IInspectable},
		ISmsFilterRule: p,
	}
	com.AddToScope(result)
	return result
}

type SmsFilterRules struct {
	RtClass
	*ISmsFilterRules
}

func NewSmsFilterRules_CreateFilterRules(actionType SmsFilterActionType) *SmsFilterRules {
	hs := NewHStr("Windows.Devices.Sms.SmsFilterRules")
	var pFac *ISmsFilterRulesFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISmsFilterRulesFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISmsFilterRules
	p = pFac.CreateFilterRules(actionType)
	result := &SmsFilterRules{
		RtClass:         RtClass{PInspect: &p.IInspectable},
		ISmsFilterRules: p,
	}
	com.AddToScope(result)
	return result
}

type SmsMessageReceivedTriggerDetails struct {
	RtClass
	*ISmsMessageReceivedTriggerDetails
}

type SmsMessageRegistration struct {
	RtClass
	*ISmsMessageRegistration
}

func NewISmsMessageRegistrationStatics() *ISmsMessageRegistrationStatics {
	var p *ISmsMessageRegistrationStatics
	hs := NewHStr("Windows.Devices.Sms.SmsMessageRegistration")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISmsMessageRegistrationStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SmsReceivedEventDetails struct {
	RtClass
	*ISmsReceivedEventDetails
}

type SmsSendMessageResult struct {
	RtClass
	*ISmsSendMessageResult
}

type SmsStatusMessage struct {
	RtClass
	*ISmsStatusMessage
}

type SmsTextMessage struct {
	RtClass
	*ISmsTextMessage
}

func NewSmsTextMessage() *SmsTextMessage {
	hs := NewHStr("Windows.Devices.Sms.SmsTextMessage")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &SmsTextMessage{
		RtClass:         RtClass{PInspect: p},
		ISmsTextMessage: (*ISmsTextMessage)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewISmsTextMessageStatics() *ISmsTextMessageStatics {
	var p *ISmsTextMessageStatics
	hs := NewHStr("Windows.Devices.Sms.SmsTextMessage")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISmsTextMessageStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SmsTextMessage2 struct {
	RtClass
	*ISmsTextMessage2
}

func NewSmsTextMessage2() *SmsTextMessage2 {
	hs := NewHStr("Windows.Devices.Sms.SmsTextMessage2")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &SmsTextMessage2{
		RtClass:          RtClass{PInspect: p},
		ISmsTextMessage2: (*ISmsTextMessage2)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type SmsVoicemailMessage struct {
	RtClass
	*ISmsVoicemailMessage
}

type SmsWapMessage struct {
	RtClass
	*ISmsWapMessage
}
