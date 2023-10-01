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
type AdaptiveNotificationContentKind int32

const (
	AdaptiveNotificationContentKind_Text AdaptiveNotificationContentKind = 0
)

// enum
type BadgeTemplateType int32

const (
	BadgeTemplateType_BadgeGlyph  BadgeTemplateType = 0
	BadgeTemplateType_BadgeNumber BadgeTemplateType = 1
)

// enum
// flags
type NotificationKinds uint32

const (
	NotificationKinds_Unknown NotificationKinds = 0
	NotificationKinds_Toast   NotificationKinds = 1
)

// enum
type NotificationMirroring int32

const (
	NotificationMirroring_Allowed  NotificationMirroring = 0
	NotificationMirroring_Disabled NotificationMirroring = 1
)

// enum
type NotificationSetting int32

const (
	NotificationSetting_Enabled                NotificationSetting = 0
	NotificationSetting_DisabledForApplication NotificationSetting = 1
	NotificationSetting_DisabledForUser        NotificationSetting = 2
	NotificationSetting_DisabledByGroupPolicy  NotificationSetting = 3
	NotificationSetting_DisabledByManifest     NotificationSetting = 4
)

// enum
type NotificationUpdateResult int32

const (
	NotificationUpdateResult_Succeeded            NotificationUpdateResult = 0
	NotificationUpdateResult_Failed               NotificationUpdateResult = 1
	NotificationUpdateResult_NotificationNotFound NotificationUpdateResult = 2
)

// enum
type PeriodicUpdateRecurrence int32

const (
	PeriodicUpdateRecurrence_HalfHour    PeriodicUpdateRecurrence = 0
	PeriodicUpdateRecurrence_Hour        PeriodicUpdateRecurrence = 1
	PeriodicUpdateRecurrence_SixHours    PeriodicUpdateRecurrence = 2
	PeriodicUpdateRecurrence_TwelveHours PeriodicUpdateRecurrence = 3
	PeriodicUpdateRecurrence_Daily       PeriodicUpdateRecurrence = 4
)

// enum
type TileFlyoutTemplateType int32

const (
	TileFlyoutTemplateType_TileFlyoutTemplate01 TileFlyoutTemplateType = 0
)

// enum
type TileTemplateType int32

const (
	TileTemplateType_TileSquareImage                           TileTemplateType = 0
	TileTemplateType_TileSquareBlock                           TileTemplateType = 1
	TileTemplateType_TileSquareText01                          TileTemplateType = 2
	TileTemplateType_TileSquareText02                          TileTemplateType = 3
	TileTemplateType_TileSquareText03                          TileTemplateType = 4
	TileTemplateType_TileSquareText04                          TileTemplateType = 5
	TileTemplateType_TileSquarePeekImageAndText01              TileTemplateType = 6
	TileTemplateType_TileSquarePeekImageAndText02              TileTemplateType = 7
	TileTemplateType_TileSquarePeekImageAndText03              TileTemplateType = 8
	TileTemplateType_TileSquarePeekImageAndText04              TileTemplateType = 9
	TileTemplateType_TileWideImage                             TileTemplateType = 10
	TileTemplateType_TileWideImageCollection                   TileTemplateType = 11
	TileTemplateType_TileWideImageAndText01                    TileTemplateType = 12
	TileTemplateType_TileWideImageAndText02                    TileTemplateType = 13
	TileTemplateType_TileWideBlockAndText01                    TileTemplateType = 14
	TileTemplateType_TileWideBlockAndText02                    TileTemplateType = 15
	TileTemplateType_TileWidePeekImageCollection01             TileTemplateType = 16
	TileTemplateType_TileWidePeekImageCollection02             TileTemplateType = 17
	TileTemplateType_TileWidePeekImageCollection03             TileTemplateType = 18
	TileTemplateType_TileWidePeekImageCollection04             TileTemplateType = 19
	TileTemplateType_TileWidePeekImageCollection05             TileTemplateType = 20
	TileTemplateType_TileWidePeekImageCollection06             TileTemplateType = 21
	TileTemplateType_TileWidePeekImageAndText01                TileTemplateType = 22
	TileTemplateType_TileWidePeekImageAndText02                TileTemplateType = 23
	TileTemplateType_TileWidePeekImage01                       TileTemplateType = 24
	TileTemplateType_TileWidePeekImage02                       TileTemplateType = 25
	TileTemplateType_TileWidePeekImage03                       TileTemplateType = 26
	TileTemplateType_TileWidePeekImage04                       TileTemplateType = 27
	TileTemplateType_TileWidePeekImage05                       TileTemplateType = 28
	TileTemplateType_TileWidePeekImage06                       TileTemplateType = 29
	TileTemplateType_TileWideSmallImageAndText01               TileTemplateType = 30
	TileTemplateType_TileWideSmallImageAndText02               TileTemplateType = 31
	TileTemplateType_TileWideSmallImageAndText03               TileTemplateType = 32
	TileTemplateType_TileWideSmallImageAndText04               TileTemplateType = 33
	TileTemplateType_TileWideSmallImageAndText05               TileTemplateType = 34
	TileTemplateType_TileWideText01                            TileTemplateType = 35
	TileTemplateType_TileWideText02                            TileTemplateType = 36
	TileTemplateType_TileWideText03                            TileTemplateType = 37
	TileTemplateType_TileWideText04                            TileTemplateType = 38
	TileTemplateType_TileWideText05                            TileTemplateType = 39
	TileTemplateType_TileWideText06                            TileTemplateType = 40
	TileTemplateType_TileWideText07                            TileTemplateType = 41
	TileTemplateType_TileWideText08                            TileTemplateType = 42
	TileTemplateType_TileWideText09                            TileTemplateType = 43
	TileTemplateType_TileWideText10                            TileTemplateType = 44
	TileTemplateType_TileWideText11                            TileTemplateType = 45
	TileTemplateType_TileSquare150x150Image                    TileTemplateType = 0
	TileTemplateType_TileSquare150x150Block                    TileTemplateType = 1
	TileTemplateType_TileSquare150x150Text01                   TileTemplateType = 2
	TileTemplateType_TileSquare150x150Text02                   TileTemplateType = 3
	TileTemplateType_TileSquare150x150Text03                   TileTemplateType = 4
	TileTemplateType_TileSquare150x150Text04                   TileTemplateType = 5
	TileTemplateType_TileSquare150x150PeekImageAndText01       TileTemplateType = 6
	TileTemplateType_TileSquare150x150PeekImageAndText02       TileTemplateType = 7
	TileTemplateType_TileSquare150x150PeekImageAndText03       TileTemplateType = 8
	TileTemplateType_TileSquare150x150PeekImageAndText04       TileTemplateType = 9
	TileTemplateType_TileWide310x150Image                      TileTemplateType = 10
	TileTemplateType_TileWide310x150ImageCollection            TileTemplateType = 11
	TileTemplateType_TileWide310x150ImageAndText01             TileTemplateType = 12
	TileTemplateType_TileWide310x150ImageAndText02             TileTemplateType = 13
	TileTemplateType_TileWide310x150BlockAndText01             TileTemplateType = 14
	TileTemplateType_TileWide310x150BlockAndText02             TileTemplateType = 15
	TileTemplateType_TileWide310x150PeekImageCollection01      TileTemplateType = 16
	TileTemplateType_TileWide310x150PeekImageCollection02      TileTemplateType = 17
	TileTemplateType_TileWide310x150PeekImageCollection03      TileTemplateType = 18
	TileTemplateType_TileWide310x150PeekImageCollection04      TileTemplateType = 19
	TileTemplateType_TileWide310x150PeekImageCollection05      TileTemplateType = 20
	TileTemplateType_TileWide310x150PeekImageCollection06      TileTemplateType = 21
	TileTemplateType_TileWide310x150PeekImageAndText01         TileTemplateType = 22
	TileTemplateType_TileWide310x150PeekImageAndText02         TileTemplateType = 23
	TileTemplateType_TileWide310x150PeekImage01                TileTemplateType = 24
	TileTemplateType_TileWide310x150PeekImage02                TileTemplateType = 25
	TileTemplateType_TileWide310x150PeekImage03                TileTemplateType = 26
	TileTemplateType_TileWide310x150PeekImage04                TileTemplateType = 27
	TileTemplateType_TileWide310x150PeekImage05                TileTemplateType = 28
	TileTemplateType_TileWide310x150PeekImage06                TileTemplateType = 29
	TileTemplateType_TileWide310x150SmallImageAndText01        TileTemplateType = 30
	TileTemplateType_TileWide310x150SmallImageAndText02        TileTemplateType = 31
	TileTemplateType_TileWide310x150SmallImageAndText03        TileTemplateType = 32
	TileTemplateType_TileWide310x150SmallImageAndText04        TileTemplateType = 33
	TileTemplateType_TileWide310x150SmallImageAndText05        TileTemplateType = 34
	TileTemplateType_TileWide310x150Text01                     TileTemplateType = 35
	TileTemplateType_TileWide310x150Text02                     TileTemplateType = 36
	TileTemplateType_TileWide310x150Text03                     TileTemplateType = 37
	TileTemplateType_TileWide310x150Text04                     TileTemplateType = 38
	TileTemplateType_TileWide310x150Text05                     TileTemplateType = 39
	TileTemplateType_TileWide310x150Text06                     TileTemplateType = 40
	TileTemplateType_TileWide310x150Text07                     TileTemplateType = 41
	TileTemplateType_TileWide310x150Text08                     TileTemplateType = 42
	TileTemplateType_TileWide310x150Text09                     TileTemplateType = 43
	TileTemplateType_TileWide310x150Text10                     TileTemplateType = 44
	TileTemplateType_TileWide310x150Text11                     TileTemplateType = 45
	TileTemplateType_TileSquare310x310BlockAndText01           TileTemplateType = 46
	TileTemplateType_TileSquare310x310BlockAndText02           TileTemplateType = 47
	TileTemplateType_TileSquare310x310Image                    TileTemplateType = 48
	TileTemplateType_TileSquare310x310ImageAndText01           TileTemplateType = 49
	TileTemplateType_TileSquare310x310ImageAndText02           TileTemplateType = 50
	TileTemplateType_TileSquare310x310ImageAndTextOverlay01    TileTemplateType = 51
	TileTemplateType_TileSquare310x310ImageAndTextOverlay02    TileTemplateType = 52
	TileTemplateType_TileSquare310x310ImageAndTextOverlay03    TileTemplateType = 53
	TileTemplateType_TileSquare310x310ImageCollectionAndText01 TileTemplateType = 54
	TileTemplateType_TileSquare310x310ImageCollectionAndText02 TileTemplateType = 55
	TileTemplateType_TileSquare310x310ImageCollection          TileTemplateType = 56
	TileTemplateType_TileSquare310x310SmallImagesAndTextList01 TileTemplateType = 57
	TileTemplateType_TileSquare310x310SmallImagesAndTextList02 TileTemplateType = 58
	TileTemplateType_TileSquare310x310SmallImagesAndTextList03 TileTemplateType = 59
	TileTemplateType_TileSquare310x310SmallImagesAndTextList04 TileTemplateType = 60
	TileTemplateType_TileSquare310x310Text01                   TileTemplateType = 61
	TileTemplateType_TileSquare310x310Text02                   TileTemplateType = 62
	TileTemplateType_TileSquare310x310Text03                   TileTemplateType = 63
	TileTemplateType_TileSquare310x310Text04                   TileTemplateType = 64
	TileTemplateType_TileSquare310x310Text05                   TileTemplateType = 65
	TileTemplateType_TileSquare310x310Text06                   TileTemplateType = 66
	TileTemplateType_TileSquare310x310Text07                   TileTemplateType = 67
	TileTemplateType_TileSquare310x310Text08                   TileTemplateType = 68
	TileTemplateType_TileSquare310x310TextList01               TileTemplateType = 69
	TileTemplateType_TileSquare310x310TextList02               TileTemplateType = 70
	TileTemplateType_TileSquare310x310TextList03               TileTemplateType = 71
	TileTemplateType_TileSquare310x310SmallImageAndText01      TileTemplateType = 72
	TileTemplateType_TileSquare310x310SmallImagesAndTextList05 TileTemplateType = 73
	TileTemplateType_TileSquare310x310Text09                   TileTemplateType = 74
	TileTemplateType_TileSquare71x71IconWithBadge              TileTemplateType = 75
	TileTemplateType_TileSquare150x150IconWithBadge            TileTemplateType = 76
	TileTemplateType_TileWide310x150IconWithBadgeAndText       TileTemplateType = 77
	TileTemplateType_TileSquare71x71Image                      TileTemplateType = 78
	TileTemplateType_TileTall150x310Image                      TileTemplateType = 79
)

// enum
type ToastDismissalReason int32

const (
	ToastDismissalReason_UserCanceled      ToastDismissalReason = 0
	ToastDismissalReason_ApplicationHidden ToastDismissalReason = 1
	ToastDismissalReason_TimedOut          ToastDismissalReason = 2
)

// enum
type ToastHistoryChangedType int32

const (
	ToastHistoryChangedType_Cleared ToastHistoryChangedType = 0
	ToastHistoryChangedType_Removed ToastHistoryChangedType = 1
	ToastHistoryChangedType_Expired ToastHistoryChangedType = 2
	ToastHistoryChangedType_Added   ToastHistoryChangedType = 3
)

// enum
type ToastNotificationPriority int32

const (
	ToastNotificationPriority_Default ToastNotificationPriority = 0
	ToastNotificationPriority_High    ToastNotificationPriority = 1
)

// enum
type ToastTemplateType int32

const (
	ToastTemplateType_ToastImageAndText01 ToastTemplateType = 0
	ToastTemplateType_ToastImageAndText02 ToastTemplateType = 1
	ToastTemplateType_ToastImageAndText03 ToastTemplateType = 2
	ToastTemplateType_ToastImageAndText04 ToastTemplateType = 3
	ToastTemplateType_ToastText01         ToastTemplateType = 4
	ToastTemplateType_ToastText02         ToastTemplateType = 5
	ToastTemplateType_ToastText03         ToastTemplateType = 6
	ToastTemplateType_ToastText04         ToastTemplateType = 7
)

// enum
type UserNotificationChangedKind int32

const (
	UserNotificationChangedKind_Added   UserNotificationChangedKind = 0
	UserNotificationChangedKind_Removed UserNotificationChangedKind = 1
)

// interfaces

// EB0DBE66-7448-448D-9DB8-D78ACD2ABBA9
var IID_IAdaptiveNotificationContent = syscall.GUID{0xEB0DBE66, 0x7448, 0x448D,
	[8]byte{0x9D, 0xB8, 0xD7, 0x8A, 0xCD, 0x2A, 0xBB, 0xA9}}

type IAdaptiveNotificationContentInterface interface {
	win32.IInspectableInterface
	Get_Kind() AdaptiveNotificationContentKind
	Get_Hints() *IMap[string, string]
}

type IAdaptiveNotificationContentVtbl struct {
	win32.IInspectableVtbl
	Get_Kind  uintptr
	Get_Hints uintptr
}

type IAdaptiveNotificationContent struct {
	win32.IInspectable
}

func (this *IAdaptiveNotificationContent) Vtbl() *IAdaptiveNotificationContentVtbl {
	return (*IAdaptiveNotificationContentVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveNotificationContent) Get_Kind() AdaptiveNotificationContentKind {
	var _result AdaptiveNotificationContentKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Kind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveNotificationContent) Get_Hints() *IMap[string, string] {
	var _result *IMap[string, string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Hints, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 46D4A3BE-609A-4326-A40B-BFDE872034A3
var IID_IAdaptiveNotificationText = syscall.GUID{0x46D4A3BE, 0x609A, 0x4326,
	[8]byte{0xA4, 0x0B, 0xBF, 0xDE, 0x87, 0x20, 0x34, 0xA3}}

type IAdaptiveNotificationTextInterface interface {
	win32.IInspectableInterface
	Get_Text() string
	Put_Text(value string)
	Get_Language() string
	Put_Language(value string)
}

type IAdaptiveNotificationTextVtbl struct {
	win32.IInspectableVtbl
	Get_Text     uintptr
	Put_Text     uintptr
	Get_Language uintptr
	Put_Language uintptr
}

type IAdaptiveNotificationText struct {
	win32.IInspectable
}

func (this *IAdaptiveNotificationText) Vtbl() *IAdaptiveNotificationTextVtbl {
	return (*IAdaptiveNotificationTextVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveNotificationText) Get_Text() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Text, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAdaptiveNotificationText) Put_Text(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Text, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IAdaptiveNotificationText) Get_Language() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Language, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAdaptiveNotificationText) Put_Language(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Language, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 075CB4CA-D08A-4E2F-9233-7E289C1F7722
var IID_IBadgeNotification = syscall.GUID{0x075CB4CA, 0xD08A, 0x4E2F,
	[8]byte{0x92, 0x33, 0x7E, 0x28, 0x9C, 0x1F, 0x77, 0x22}}

type IBadgeNotificationInterface interface {
	win32.IInspectableInterface
	Get_Content() unsafe.Pointer
	Put_ExpirationTime(value *IReference[DateTime])
	Get_ExpirationTime() *IReference[DateTime]
}

type IBadgeNotificationVtbl struct {
	win32.IInspectableVtbl
	Get_Content        uintptr
	Put_ExpirationTime uintptr
	Get_ExpirationTime uintptr
}

type IBadgeNotification struct {
	win32.IInspectable
}

func (this *IBadgeNotification) Vtbl() *IBadgeNotificationVtbl {
	return (*IBadgeNotificationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBadgeNotification) Get_Content() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Content, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBadgeNotification) Put_ExpirationTime(value *IReference[DateTime]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ExpirationTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IBadgeNotification) Get_ExpirationTime() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExpirationTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EDF255CE-0618-4D59-948A-5A61040C52F9
var IID_IBadgeNotificationFactory = syscall.GUID{0xEDF255CE, 0x0618, 0x4D59,
	[8]byte{0x94, 0x8A, 0x5A, 0x61, 0x04, 0x0C, 0x52, 0xF9}}

type IBadgeNotificationFactoryInterface interface {
	win32.IInspectableInterface
	CreateBadgeNotification(content unsafe.Pointer) *IBadgeNotification
}

type IBadgeNotificationFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateBadgeNotification uintptr
}

type IBadgeNotificationFactory struct {
	win32.IInspectable
}

func (this *IBadgeNotificationFactory) Vtbl() *IBadgeNotificationFactoryVtbl {
	return (*IBadgeNotificationFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBadgeNotificationFactory) CreateBadgeNotification(content unsafe.Pointer) *IBadgeNotification {
	var _result *IBadgeNotification
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateBadgeNotification, uintptr(unsafe.Pointer(this)), uintptr(content), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 996B21BC-0386-44E5-BA8D-0C1077A62E92
var IID_IBadgeUpdateManagerForUser = syscall.GUID{0x996B21BC, 0x0386, 0x44E5,
	[8]byte{0xBA, 0x8D, 0x0C, 0x10, 0x77, 0xA6, 0x2E, 0x92}}

type IBadgeUpdateManagerForUserInterface interface {
	win32.IInspectableInterface
	CreateBadgeUpdaterForApplication() *IBadgeUpdater
	CreateBadgeUpdaterForApplicationWithId(applicationId string) *IBadgeUpdater
	CreateBadgeUpdaterForSecondaryTile(tileId string) *IBadgeUpdater
	Get_User() *IUser
}

type IBadgeUpdateManagerForUserVtbl struct {
	win32.IInspectableVtbl
	CreateBadgeUpdaterForApplication       uintptr
	CreateBadgeUpdaterForApplicationWithId uintptr
	CreateBadgeUpdaterForSecondaryTile     uintptr
	Get_User                               uintptr
}

type IBadgeUpdateManagerForUser struct {
	win32.IInspectable
}

func (this *IBadgeUpdateManagerForUser) Vtbl() *IBadgeUpdateManagerForUserVtbl {
	return (*IBadgeUpdateManagerForUserVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBadgeUpdateManagerForUser) CreateBadgeUpdaterForApplication() *IBadgeUpdater {
	var _result *IBadgeUpdater
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateBadgeUpdaterForApplication, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBadgeUpdateManagerForUser) CreateBadgeUpdaterForApplicationWithId(applicationId string) *IBadgeUpdater {
	var _result *IBadgeUpdater
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateBadgeUpdaterForApplicationWithId, uintptr(unsafe.Pointer(this)), NewHStr(applicationId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBadgeUpdateManagerForUser) CreateBadgeUpdaterForSecondaryTile(tileId string) *IBadgeUpdater {
	var _result *IBadgeUpdater
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateBadgeUpdaterForSecondaryTile, uintptr(unsafe.Pointer(this)), NewHStr(tileId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBadgeUpdateManagerForUser) Get_User() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_User, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 33400FAA-6DD5-4105-AEBC-9B50FCA492DA
var IID_IBadgeUpdateManagerStatics = syscall.GUID{0x33400FAA, 0x6DD5, 0x4105,
	[8]byte{0xAE, 0xBC, 0x9B, 0x50, 0xFC, 0xA4, 0x92, 0xDA}}

type IBadgeUpdateManagerStaticsInterface interface {
	win32.IInspectableInterface
	CreateBadgeUpdaterForApplication() *IBadgeUpdater
	CreateBadgeUpdaterForApplicationWithId(applicationId string) *IBadgeUpdater
	CreateBadgeUpdaterForSecondaryTile(tileId string) *IBadgeUpdater
	GetTemplateContent(type_ BadgeTemplateType) unsafe.Pointer
}

type IBadgeUpdateManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateBadgeUpdaterForApplication       uintptr
	CreateBadgeUpdaterForApplicationWithId uintptr
	CreateBadgeUpdaterForSecondaryTile     uintptr
	GetTemplateContent                     uintptr
}

type IBadgeUpdateManagerStatics struct {
	win32.IInspectable
}

func (this *IBadgeUpdateManagerStatics) Vtbl() *IBadgeUpdateManagerStaticsVtbl {
	return (*IBadgeUpdateManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBadgeUpdateManagerStatics) CreateBadgeUpdaterForApplication() *IBadgeUpdater {
	var _result *IBadgeUpdater
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateBadgeUpdaterForApplication, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBadgeUpdateManagerStatics) CreateBadgeUpdaterForApplicationWithId(applicationId string) *IBadgeUpdater {
	var _result *IBadgeUpdater
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateBadgeUpdaterForApplicationWithId, uintptr(unsafe.Pointer(this)), NewHStr(applicationId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBadgeUpdateManagerStatics) CreateBadgeUpdaterForSecondaryTile(tileId string) *IBadgeUpdater {
	var _result *IBadgeUpdater
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateBadgeUpdaterForSecondaryTile, uintptr(unsafe.Pointer(this)), NewHStr(tileId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBadgeUpdateManagerStatics) GetTemplateContent(type_ BadgeTemplateType) unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTemplateContent, uintptr(unsafe.Pointer(this)), uintptr(type_), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 979A35CE-F940-48BF-94E8-CA244D400B41
var IID_IBadgeUpdateManagerStatics2 = syscall.GUID{0x979A35CE, 0xF940, 0x48BF,
	[8]byte{0x94, 0xE8, 0xCA, 0x24, 0x4D, 0x40, 0x0B, 0x41}}

type IBadgeUpdateManagerStatics2Interface interface {
	win32.IInspectableInterface
	GetForUser(user *IUser) *IBadgeUpdateManagerForUser
}

type IBadgeUpdateManagerStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetForUser uintptr
}

type IBadgeUpdateManagerStatics2 struct {
	win32.IInspectable
}

func (this *IBadgeUpdateManagerStatics2) Vtbl() *IBadgeUpdateManagerStatics2Vtbl {
	return (*IBadgeUpdateManagerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBadgeUpdateManagerStatics2) GetForUser(user *IUser) *IBadgeUpdateManagerForUser {
	var _result *IBadgeUpdateManagerForUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B5FA1FD4-7562-4F6C-BFA3-1B6ED2E57F2F
var IID_IBadgeUpdater = syscall.GUID{0xB5FA1FD4, 0x7562, 0x4F6C,
	[8]byte{0xBF, 0xA3, 0x1B, 0x6E, 0xD2, 0xE5, 0x7F, 0x2F}}

type IBadgeUpdaterInterface interface {
	win32.IInspectableInterface
	Update(notification *IBadgeNotification)
	Clear()
	StartPeriodicUpdate(badgeContent *IUriRuntimeClass, requestedInterval PeriodicUpdateRecurrence)
	StartPeriodicUpdateAtTime(badgeContent *IUriRuntimeClass, startTime DateTime, requestedInterval PeriodicUpdateRecurrence)
	StopPeriodicUpdate()
}

type IBadgeUpdaterVtbl struct {
	win32.IInspectableVtbl
	Update                    uintptr
	Clear                     uintptr
	StartPeriodicUpdate       uintptr
	StartPeriodicUpdateAtTime uintptr
	StopPeriodicUpdate        uintptr
}

type IBadgeUpdater struct {
	win32.IInspectable
}

func (this *IBadgeUpdater) Vtbl() *IBadgeUpdaterVtbl {
	return (*IBadgeUpdaterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBadgeUpdater) Update(notification *IBadgeNotification) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Update, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(notification)))
	_ = _hr
}

func (this *IBadgeUpdater) Clear() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Clear, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IBadgeUpdater) StartPeriodicUpdate(badgeContent *IUriRuntimeClass, requestedInterval PeriodicUpdateRecurrence) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartPeriodicUpdate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(badgeContent)), uintptr(requestedInterval))
	_ = _hr
}

func (this *IBadgeUpdater) StartPeriodicUpdateAtTime(badgeContent *IUriRuntimeClass, startTime DateTime, requestedInterval PeriodicUpdateRecurrence) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartPeriodicUpdateAtTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(badgeContent)), *(*uintptr)(unsafe.Pointer(&startTime)), uintptr(requestedInterval))
	_ = _hr
}

func (this *IBadgeUpdater) StopPeriodicUpdate() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopPeriodicUpdate, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 06206598-D496-497D-8692-4F7D7C2770DF
var IID_IKnownAdaptiveNotificationHintsStatics = syscall.GUID{0x06206598, 0xD496, 0x497D,
	[8]byte{0x86, 0x92, 0x4F, 0x7D, 0x7C, 0x27, 0x70, 0xDF}}

type IKnownAdaptiveNotificationHintsStaticsInterface interface {
	win32.IInspectableInterface
	Get_Style() string
	Get_Wrap() string
	Get_MaxLines() string
	Get_MinLines() string
	Get_TextStacking() string
	Get_Align() string
}

type IKnownAdaptiveNotificationHintsStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Style        uintptr
	Get_Wrap         uintptr
	Get_MaxLines     uintptr
	Get_MinLines     uintptr
	Get_TextStacking uintptr
	Get_Align        uintptr
}

type IKnownAdaptiveNotificationHintsStatics struct {
	win32.IInspectable
}

func (this *IKnownAdaptiveNotificationHintsStatics) Vtbl() *IKnownAdaptiveNotificationHintsStaticsVtbl {
	return (*IKnownAdaptiveNotificationHintsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownAdaptiveNotificationHintsStatics) Get_Style() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Style, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownAdaptiveNotificationHintsStatics) Get_Wrap() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Wrap, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownAdaptiveNotificationHintsStatics) Get_MaxLines() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxLines, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownAdaptiveNotificationHintsStatics) Get_MinLines() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinLines, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownAdaptiveNotificationHintsStatics) Get_TextStacking() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TextStacking, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownAdaptiveNotificationHintsStatics) Get_Align() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Align, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 202192D7-8996-45AA-8BA1-D461D72C2A1B
var IID_IKnownAdaptiveNotificationTextStylesStatics = syscall.GUID{0x202192D7, 0x8996, 0x45AA,
	[8]byte{0x8B, 0xA1, 0xD4, 0x61, 0xD7, 0x2C, 0x2A, 0x1B}}

type IKnownAdaptiveNotificationTextStylesStaticsInterface interface {
	win32.IInspectableInterface
	Get_Caption() string
	Get_Body() string
	Get_Base() string
	Get_Subtitle() string
	Get_Title() string
	Get_Subheader() string
	Get_Header() string
	Get_TitleNumeral() string
	Get_SubheaderNumeral() string
	Get_HeaderNumeral() string
	Get_CaptionSubtle() string
	Get_BodySubtle() string
	Get_BaseSubtle() string
	Get_SubtitleSubtle() string
	Get_TitleSubtle() string
	Get_SubheaderSubtle() string
	Get_SubheaderNumeralSubtle() string
	Get_HeaderSubtle() string
	Get_HeaderNumeralSubtle() string
}

type IKnownAdaptiveNotificationTextStylesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Caption                uintptr
	Get_Body                   uintptr
	Get_Base                   uintptr
	Get_Subtitle               uintptr
	Get_Title                  uintptr
	Get_Subheader              uintptr
	Get_Header                 uintptr
	Get_TitleNumeral           uintptr
	Get_SubheaderNumeral       uintptr
	Get_HeaderNumeral          uintptr
	Get_CaptionSubtle          uintptr
	Get_BodySubtle             uintptr
	Get_BaseSubtle             uintptr
	Get_SubtitleSubtle         uintptr
	Get_TitleSubtle            uintptr
	Get_SubheaderSubtle        uintptr
	Get_SubheaderNumeralSubtle uintptr
	Get_HeaderSubtle           uintptr
	Get_HeaderNumeralSubtle    uintptr
}

type IKnownAdaptiveNotificationTextStylesStatics struct {
	win32.IInspectable
}

func (this *IKnownAdaptiveNotificationTextStylesStatics) Vtbl() *IKnownAdaptiveNotificationTextStylesStaticsVtbl {
	return (*IKnownAdaptiveNotificationTextStylesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownAdaptiveNotificationTextStylesStatics) Get_Caption() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Caption, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownAdaptiveNotificationTextStylesStatics) Get_Body() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Body, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownAdaptiveNotificationTextStylesStatics) Get_Base() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Base, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownAdaptiveNotificationTextStylesStatics) Get_Subtitle() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Subtitle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownAdaptiveNotificationTextStylesStatics) Get_Title() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownAdaptiveNotificationTextStylesStatics) Get_Subheader() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Subheader, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownAdaptiveNotificationTextStylesStatics) Get_Header() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Header, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownAdaptiveNotificationTextStylesStatics) Get_TitleNumeral() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TitleNumeral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownAdaptiveNotificationTextStylesStatics) Get_SubheaderNumeral() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SubheaderNumeral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownAdaptiveNotificationTextStylesStatics) Get_HeaderNumeral() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HeaderNumeral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownAdaptiveNotificationTextStylesStatics) Get_CaptionSubtle() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CaptionSubtle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownAdaptiveNotificationTextStylesStatics) Get_BodySubtle() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BodySubtle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownAdaptiveNotificationTextStylesStatics) Get_BaseSubtle() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BaseSubtle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownAdaptiveNotificationTextStylesStatics) Get_SubtitleSubtle() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SubtitleSubtle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownAdaptiveNotificationTextStylesStatics) Get_TitleSubtle() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TitleSubtle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownAdaptiveNotificationTextStylesStatics) Get_SubheaderSubtle() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SubheaderSubtle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownAdaptiveNotificationTextStylesStatics) Get_SubheaderNumeralSubtle() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SubheaderNumeralSubtle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownAdaptiveNotificationTextStylesStatics) Get_HeaderSubtle() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HeaderSubtle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownAdaptiveNotificationTextStylesStatics) Get_HeaderNumeralSubtle() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HeaderNumeralSubtle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 79427BAE-A8B7-4D58-89EA-76A7B7BCCDED
var IID_IKnownNotificationBindingsStatics = syscall.GUID{0x79427BAE, 0xA8B7, 0x4D58,
	[8]byte{0x89, 0xEA, 0x76, 0xA7, 0xB7, 0xBC, 0xCD, 0xED}}

type IKnownNotificationBindingsStaticsInterface interface {
	win32.IInspectableInterface
	Get_ToastGeneric() string
}

type IKnownNotificationBindingsStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_ToastGeneric uintptr
}

type IKnownNotificationBindingsStatics struct {
	win32.IInspectable
}

func (this *IKnownNotificationBindingsStatics) Vtbl() *IKnownNotificationBindingsStaticsVtbl {
	return (*IKnownNotificationBindingsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownNotificationBindingsStatics) Get_ToastGeneric() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ToastGeneric, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 108037FE-EB76-4F82-97BC-DA07530A2E20
var IID_INotification = syscall.GUID{0x108037FE, 0xEB76, 0x4F82,
	[8]byte{0x97, 0xBC, 0xDA, 0x07, 0x53, 0x0A, 0x2E, 0x20}}

type INotificationInterface interface {
	win32.IInspectableInterface
	Get_ExpirationTime() *IReference[DateTime]
	Put_ExpirationTime(value *IReference[DateTime])
	Get_Visual() *INotificationVisual
	Put_Visual(value *INotificationVisual)
}

type INotificationVtbl struct {
	win32.IInspectableVtbl
	Get_ExpirationTime uintptr
	Put_ExpirationTime uintptr
	Get_Visual         uintptr
	Put_Visual         uintptr
}

type INotification struct {
	win32.IInspectable
}

func (this *INotification) Vtbl() *INotificationVtbl {
	return (*INotificationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INotification) Get_ExpirationTime() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExpirationTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INotification) Put_ExpirationTime(value *IReference[DateTime]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ExpirationTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *INotification) Get_Visual() *INotificationVisual {
	var _result *INotificationVisual
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Visual, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INotification) Put_Visual(value *INotificationVisual) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Visual, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// F29E4B85-0370-4AD3-B4EA-DA9E35E7EABF
var IID_INotificationBinding = syscall.GUID{0xF29E4B85, 0x0370, 0x4AD3,
	[8]byte{0xB4, 0xEA, 0xDA, 0x9E, 0x35, 0xE7, 0xEA, 0xBF}}

type INotificationBindingInterface interface {
	win32.IInspectableInterface
	Get_Template() string
	Put_Template(value string)
	Get_Language() string
	Put_Language(value string)
	Get_Hints() *IMap[string, string]
	GetTextElements() *IVectorView[*IAdaptiveNotificationText]
}

type INotificationBindingVtbl struct {
	win32.IInspectableVtbl
	Get_Template    uintptr
	Put_Template    uintptr
	Get_Language    uintptr
	Put_Language    uintptr
	Get_Hints       uintptr
	GetTextElements uintptr
}

type INotificationBinding struct {
	win32.IInspectable
}

func (this *INotificationBinding) Vtbl() *INotificationBindingVtbl {
	return (*INotificationBindingVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INotificationBinding) Get_Template() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Template, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INotificationBinding) Put_Template(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Template, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *INotificationBinding) Get_Language() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Language, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INotificationBinding) Put_Language(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Language, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *INotificationBinding) Get_Hints() *IMap[string, string] {
	var _result *IMap[string, string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Hints, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INotificationBinding) GetTextElements() *IVectorView[*IAdaptiveNotificationText] {
	var _result *IVectorView[*IAdaptiveNotificationText]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTextElements, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9FFD2312-9D6A-4AAF-B6AC-FF17F0C1F280
var IID_INotificationData = syscall.GUID{0x9FFD2312, 0x9D6A, 0x4AAF,
	[8]byte{0xB6, 0xAC, 0xFF, 0x17, 0xF0, 0xC1, 0xF2, 0x80}}

type INotificationDataInterface interface {
	win32.IInspectableInterface
	Get_Values() *IMap[string, string]
	Get_SequenceNumber() uint32
	Put_SequenceNumber(value uint32)
}

type INotificationDataVtbl struct {
	win32.IInspectableVtbl
	Get_Values         uintptr
	Get_SequenceNumber uintptr
	Put_SequenceNumber uintptr
}

type INotificationData struct {
	win32.IInspectable
}

func (this *INotificationData) Vtbl() *INotificationDataVtbl {
	return (*INotificationDataVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INotificationData) Get_Values() *IMap[string, string] {
	var _result *IMap[string, string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Values, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INotificationData) Get_SequenceNumber() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SequenceNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INotificationData) Put_SequenceNumber(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SequenceNumber, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 23C1E33A-1C10-46FB-8040-DEC384621CF8
var IID_INotificationDataFactory = syscall.GUID{0x23C1E33A, 0x1C10, 0x46FB,
	[8]byte{0x80, 0x40, 0xDE, 0xC3, 0x84, 0x62, 0x1C, 0xF8}}

type INotificationDataFactoryInterface interface {
	win32.IInspectableInterface
	CreateNotificationDataWithValuesAndSequenceNumber(initialValues *IIterable[*IKeyValuePair[string, string]], sequenceNumber uint32) *INotificationData
	CreateNotificationDataWithValues(initialValues *IIterable[*IKeyValuePair[string, string]]) *INotificationData
}

type INotificationDataFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateNotificationDataWithValuesAndSequenceNumber uintptr
	CreateNotificationDataWithValues                  uintptr
}

type INotificationDataFactory struct {
	win32.IInspectable
}

func (this *INotificationDataFactory) Vtbl() *INotificationDataFactoryVtbl {
	return (*INotificationDataFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INotificationDataFactory) CreateNotificationDataWithValuesAndSequenceNumber(initialValues *IIterable[*IKeyValuePair[string, string]], sequenceNumber uint32) *INotificationData {
	var _result *INotificationData
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateNotificationDataWithValuesAndSequenceNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(initialValues)), uintptr(sequenceNumber), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INotificationDataFactory) CreateNotificationDataWithValues(initialValues *IIterable[*IKeyValuePair[string, string]]) *INotificationData {
	var _result *INotificationData
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateNotificationDataWithValues, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(initialValues)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 68835B8E-AA56-4E11-86D3-5F9A6957BC5B
var IID_INotificationVisual = syscall.GUID{0x68835B8E, 0xAA56, 0x4E11,
	[8]byte{0x86, 0xD3, 0x5F, 0x9A, 0x69, 0x57, 0xBC, 0x5B}}

type INotificationVisualInterface interface {
	win32.IInspectableInterface
	Get_Language() string
	Put_Language(value string)
	Get_Bindings() *IVector[*INotificationBinding]
	GetBinding(templateName string) *INotificationBinding
}

type INotificationVisualVtbl struct {
	win32.IInspectableVtbl
	Get_Language uintptr
	Put_Language uintptr
	Get_Bindings uintptr
	GetBinding   uintptr
}

type INotificationVisual struct {
	win32.IInspectable
}

func (this *INotificationVisual) Vtbl() *INotificationVisualVtbl {
	return (*INotificationVisualVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INotificationVisual) Get_Language() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Language, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INotificationVisual) Put_Language(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Language, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *INotificationVisual) Get_Bindings() *IVector[*INotificationBinding] {
	var _result *IVector[*INotificationBinding]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bindings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INotificationVisual) GetBinding(templateName string) *INotificationBinding {
	var _result *INotificationBinding
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetBinding, uintptr(unsafe.Pointer(this)), NewHStr(templateName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0ABCA6D5-99DC-4C78-A11C-C9E7F86D7EF7
var IID_IScheduledTileNotification = syscall.GUID{0x0ABCA6D5, 0x99DC, 0x4C78,
	[8]byte{0xA1, 0x1C, 0xC9, 0xE7, 0xF8, 0x6D, 0x7E, 0xF7}}

type IScheduledTileNotificationInterface interface {
	win32.IInspectableInterface
	Get_Content() unsafe.Pointer
	Get_DeliveryTime() DateTime
	Put_ExpirationTime(value *IReference[DateTime])
	Get_ExpirationTime() *IReference[DateTime]
	Put_Tag(value string)
	Get_Tag() string
	Put_Id(value string)
	Get_Id() string
}

type IScheduledTileNotificationVtbl struct {
	win32.IInspectableVtbl
	Get_Content        uintptr
	Get_DeliveryTime   uintptr
	Put_ExpirationTime uintptr
	Get_ExpirationTime uintptr
	Put_Tag            uintptr
	Get_Tag            uintptr
	Put_Id             uintptr
	Get_Id             uintptr
}

type IScheduledTileNotification struct {
	win32.IInspectable
}

func (this *IScheduledTileNotification) Vtbl() *IScheduledTileNotificationVtbl {
	return (*IScheduledTileNotificationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IScheduledTileNotification) Get_Content() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Content, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IScheduledTileNotification) Get_DeliveryTime() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeliveryTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IScheduledTileNotification) Put_ExpirationTime(value *IReference[DateTime]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ExpirationTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IScheduledTileNotification) Get_ExpirationTime() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExpirationTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IScheduledTileNotification) Put_Tag(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Tag, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IScheduledTileNotification) Get_Tag() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Tag, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IScheduledTileNotification) Put_Id(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Id, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IScheduledTileNotification) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 3383138A-98C0-4C3B-BBD6-4A633C7CFC29
var IID_IScheduledTileNotificationFactory = syscall.GUID{0x3383138A, 0x98C0, 0x4C3B,
	[8]byte{0xBB, 0xD6, 0x4A, 0x63, 0x3C, 0x7C, 0xFC, 0x29}}

type IScheduledTileNotificationFactoryInterface interface {
	win32.IInspectableInterface
	CreateScheduledTileNotification(content unsafe.Pointer, deliveryTime DateTime) *IScheduledTileNotification
}

type IScheduledTileNotificationFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateScheduledTileNotification uintptr
}

type IScheduledTileNotificationFactory struct {
	win32.IInspectable
}

func (this *IScheduledTileNotificationFactory) Vtbl() *IScheduledTileNotificationFactoryVtbl {
	return (*IScheduledTileNotificationFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IScheduledTileNotificationFactory) CreateScheduledTileNotification(content unsafe.Pointer, deliveryTime DateTime) *IScheduledTileNotification {
	var _result *IScheduledTileNotification
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateScheduledTileNotification, uintptr(unsafe.Pointer(this)), uintptr(content), *(*uintptr)(unsafe.Pointer(&deliveryTime)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 79F577F8-0DE7-48CD-9740-9B370490C838
var IID_IScheduledToastNotification = syscall.GUID{0x79F577F8, 0x0DE7, 0x48CD,
	[8]byte{0x97, 0x40, 0x9B, 0x37, 0x04, 0x90, 0xC8, 0x38}}

type IScheduledToastNotificationInterface interface {
	win32.IInspectableInterface
	Get_Content() unsafe.Pointer
	Get_DeliveryTime() DateTime
	Get_SnoozeInterval() *IReference[TimeSpan]
	Get_MaximumSnoozeCount() uint32
	Put_Id(value string)
	Get_Id() string
}

type IScheduledToastNotificationVtbl struct {
	win32.IInspectableVtbl
	Get_Content            uintptr
	Get_DeliveryTime       uintptr
	Get_SnoozeInterval     uintptr
	Get_MaximumSnoozeCount uintptr
	Put_Id                 uintptr
	Get_Id                 uintptr
}

type IScheduledToastNotification struct {
	win32.IInspectable
}

func (this *IScheduledToastNotification) Vtbl() *IScheduledToastNotificationVtbl {
	return (*IScheduledToastNotificationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IScheduledToastNotification) Get_Content() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Content, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IScheduledToastNotification) Get_DeliveryTime() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeliveryTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IScheduledToastNotification) Get_SnoozeInterval() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SnoozeInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IScheduledToastNotification) Get_MaximumSnoozeCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaximumSnoozeCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IScheduledToastNotification) Put_Id(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Id, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IScheduledToastNotification) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// A66EA09C-31B4-43B0-B5DD-7A40E85363B1
var IID_IScheduledToastNotification2 = syscall.GUID{0xA66EA09C, 0x31B4, 0x43B0,
	[8]byte{0xB5, 0xDD, 0x7A, 0x40, 0xE8, 0x53, 0x63, 0xB1}}

type IScheduledToastNotification2Interface interface {
	win32.IInspectableInterface
	Put_Tag(value string)
	Get_Tag() string
	Put_Group(value string)
	Get_Group() string
	Put_SuppressPopup(value bool)
	Get_SuppressPopup() bool
}

type IScheduledToastNotification2Vtbl struct {
	win32.IInspectableVtbl
	Put_Tag           uintptr
	Get_Tag           uintptr
	Put_Group         uintptr
	Get_Group         uintptr
	Put_SuppressPopup uintptr
	Get_SuppressPopup uintptr
}

type IScheduledToastNotification2 struct {
	win32.IInspectable
}

func (this *IScheduledToastNotification2) Vtbl() *IScheduledToastNotification2Vtbl {
	return (*IScheduledToastNotification2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IScheduledToastNotification2) Put_Tag(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Tag, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IScheduledToastNotification2) Get_Tag() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Tag, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IScheduledToastNotification2) Put_Group(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Group, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IScheduledToastNotification2) Get_Group() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Group, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IScheduledToastNotification2) Put_SuppressPopup(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SuppressPopup, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IScheduledToastNotification2) Get_SuppressPopup() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SuppressPopup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 98429E8B-BD32-4A3B-9D15-22AEA49462A1
var IID_IScheduledToastNotification3 = syscall.GUID{0x98429E8B, 0xBD32, 0x4A3B,
	[8]byte{0x9D, 0x15, 0x22, 0xAE, 0xA4, 0x94, 0x62, 0xA1}}

type IScheduledToastNotification3Interface interface {
	win32.IInspectableInterface
	Get_NotificationMirroring() NotificationMirroring
	Put_NotificationMirroring(value NotificationMirroring)
	Get_RemoteId() string
	Put_RemoteId(value string)
}

type IScheduledToastNotification3Vtbl struct {
	win32.IInspectableVtbl
	Get_NotificationMirroring uintptr
	Put_NotificationMirroring uintptr
	Get_RemoteId              uintptr
	Put_RemoteId              uintptr
}

type IScheduledToastNotification3 struct {
	win32.IInspectable
}

func (this *IScheduledToastNotification3) Vtbl() *IScheduledToastNotification3Vtbl {
	return (*IScheduledToastNotification3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IScheduledToastNotification3) Get_NotificationMirroring() NotificationMirroring {
	var _result NotificationMirroring
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NotificationMirroring, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IScheduledToastNotification3) Put_NotificationMirroring(value NotificationMirroring) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NotificationMirroring, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IScheduledToastNotification3) Get_RemoteId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemoteId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IScheduledToastNotification3) Put_RemoteId(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RemoteId, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 1D4761FD-BDEF-4E4A-96BE-0101369B58D2
var IID_IScheduledToastNotification4 = syscall.GUID{0x1D4761FD, 0xBDEF, 0x4E4A,
	[8]byte{0x96, 0xBE, 0x01, 0x01, 0x36, 0x9B, 0x58, 0xD2}}

type IScheduledToastNotification4Interface interface {
	win32.IInspectableInterface
	Get_ExpirationTime() *IReference[DateTime]
	Put_ExpirationTime(value *IReference[DateTime])
}

type IScheduledToastNotification4Vtbl struct {
	win32.IInspectableVtbl
	Get_ExpirationTime uintptr
	Put_ExpirationTime uintptr
}

type IScheduledToastNotification4 struct {
	win32.IInspectable
}

func (this *IScheduledToastNotification4) Vtbl() *IScheduledToastNotification4Vtbl {
	return (*IScheduledToastNotification4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IScheduledToastNotification4) Get_ExpirationTime() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExpirationTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IScheduledToastNotification4) Put_ExpirationTime(value *IReference[DateTime]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ExpirationTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// E7BED191-0BB9-4189-8394-31761B476FD7
var IID_IScheduledToastNotificationFactory = syscall.GUID{0xE7BED191, 0x0BB9, 0x4189,
	[8]byte{0x83, 0x94, 0x31, 0x76, 0x1B, 0x47, 0x6F, 0xD7}}

type IScheduledToastNotificationFactoryInterface interface {
	win32.IInspectableInterface
	CreateScheduledToastNotification(content unsafe.Pointer, deliveryTime DateTime) *IScheduledToastNotification
	CreateScheduledToastNotificationRecurring(content unsafe.Pointer, deliveryTime DateTime, snoozeInterval TimeSpan, maximumSnoozeCount uint32) *IScheduledToastNotification
}

type IScheduledToastNotificationFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateScheduledToastNotification          uintptr
	CreateScheduledToastNotificationRecurring uintptr
}

type IScheduledToastNotificationFactory struct {
	win32.IInspectable
}

func (this *IScheduledToastNotificationFactory) Vtbl() *IScheduledToastNotificationFactoryVtbl {
	return (*IScheduledToastNotificationFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IScheduledToastNotificationFactory) CreateScheduledToastNotification(content unsafe.Pointer, deliveryTime DateTime) *IScheduledToastNotification {
	var _result *IScheduledToastNotification
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateScheduledToastNotification, uintptr(unsafe.Pointer(this)), uintptr(content), *(*uintptr)(unsafe.Pointer(&deliveryTime)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IScheduledToastNotificationFactory) CreateScheduledToastNotificationRecurring(content unsafe.Pointer, deliveryTime DateTime, snoozeInterval TimeSpan, maximumSnoozeCount uint32) *IScheduledToastNotification {
	var _result *IScheduledToastNotification
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateScheduledToastNotificationRecurring, uintptr(unsafe.Pointer(this)), uintptr(content), *(*uintptr)(unsafe.Pointer(&deliveryTime)), *(*uintptr)(unsafe.Pointer(&snoozeInterval)), uintptr(maximumSnoozeCount), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6173F6B4-412A-5E2C-A6ED-A0209AEF9A09
var IID_IScheduledToastNotificationShowingEventArgs = syscall.GUID{0x6173F6B4, 0x412A, 0x5E2C,
	[8]byte{0xA6, 0xED, 0xA0, 0x20, 0x9A, 0xEF, 0x9A, 0x09}}

type IScheduledToastNotificationShowingEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Cancel() bool
	Put_Cancel(value bool)
	Get_ScheduledToastNotification() *IScheduledToastNotification
	GetDeferral() *IDeferral
}

type IScheduledToastNotificationShowingEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Cancel                     uintptr
	Put_Cancel                     uintptr
	Get_ScheduledToastNotification uintptr
	GetDeferral                    uintptr
}

type IScheduledToastNotificationShowingEventArgs struct {
	win32.IInspectable
}

func (this *IScheduledToastNotificationShowingEventArgs) Vtbl() *IScheduledToastNotificationShowingEventArgsVtbl {
	return (*IScheduledToastNotificationShowingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IScheduledToastNotificationShowingEventArgs) Get_Cancel() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Cancel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IScheduledToastNotificationShowingEventArgs) Put_Cancel(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Cancel, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IScheduledToastNotificationShowingEventArgs) Get_ScheduledToastNotification() *IScheduledToastNotification {
	var _result *IScheduledToastNotification
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScheduledToastNotification, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IScheduledToastNotificationShowingEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 342D8988-5AF2-481A-A6A3-F2FDC78DE88E
var IID_IShownTileNotification = syscall.GUID{0x342D8988, 0x5AF2, 0x481A,
	[8]byte{0xA6, 0xA3, 0xF2, 0xFD, 0xC7, 0x8D, 0xE8, 0x8E}}

type IShownTileNotificationInterface interface {
	win32.IInspectableInterface
	Get_Arguments() string
}

type IShownTileNotificationVtbl struct {
	win32.IInspectableVtbl
	Get_Arguments uintptr
}

type IShownTileNotification struct {
	win32.IInspectable
}

func (this *IShownTileNotification) Vtbl() *IShownTileNotificationVtbl {
	return (*IShownTileNotificationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IShownTileNotification) Get_Arguments() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Arguments, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 9A53B261-C70C-42BE-B2F3-F42AA97D34E5
var IID_ITileFlyoutNotification = syscall.GUID{0x9A53B261, 0xC70C, 0x42BE,
	[8]byte{0xB2, 0xF3, 0xF4, 0x2A, 0xA9, 0x7D, 0x34, 0xE5}}

type ITileFlyoutNotificationInterface interface {
	win32.IInspectableInterface
	Get_Content() unsafe.Pointer
	Put_ExpirationTime(value *IReference[DateTime])
	Get_ExpirationTime() *IReference[DateTime]
}

type ITileFlyoutNotificationVtbl struct {
	win32.IInspectableVtbl
	Get_Content        uintptr
	Put_ExpirationTime uintptr
	Get_ExpirationTime uintptr
}

type ITileFlyoutNotification struct {
	win32.IInspectable
}

func (this *ITileFlyoutNotification) Vtbl() *ITileFlyoutNotificationVtbl {
	return (*ITileFlyoutNotificationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITileFlyoutNotification) Get_Content() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Content, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITileFlyoutNotification) Put_ExpirationTime(value *IReference[DateTime]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ExpirationTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ITileFlyoutNotification) Get_ExpirationTime() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExpirationTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EF556FF5-5226-4F2B-B278-88A35DFE569F
var IID_ITileFlyoutNotificationFactory = syscall.GUID{0xEF556FF5, 0x5226, 0x4F2B,
	[8]byte{0xB2, 0x78, 0x88, 0xA3, 0x5D, 0xFE, 0x56, 0x9F}}

type ITileFlyoutNotificationFactoryInterface interface {
	win32.IInspectableInterface
	CreateTileFlyoutNotification(content unsafe.Pointer) *ITileFlyoutNotification
}

type ITileFlyoutNotificationFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateTileFlyoutNotification uintptr
}

type ITileFlyoutNotificationFactory struct {
	win32.IInspectable
}

func (this *ITileFlyoutNotificationFactory) Vtbl() *ITileFlyoutNotificationFactoryVtbl {
	return (*ITileFlyoutNotificationFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITileFlyoutNotificationFactory) CreateTileFlyoutNotification(content unsafe.Pointer) *ITileFlyoutNotification {
	var _result *ITileFlyoutNotification
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateTileFlyoutNotification, uintptr(unsafe.Pointer(this)), uintptr(content), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 04363B0B-1AC0-4B99-88E7-ADA83E953D48
var IID_ITileFlyoutUpdateManagerStatics = syscall.GUID{0x04363B0B, 0x1AC0, 0x4B99,
	[8]byte{0x88, 0xE7, 0xAD, 0xA8, 0x3E, 0x95, 0x3D, 0x48}}

type ITileFlyoutUpdateManagerStaticsInterface interface {
	win32.IInspectableInterface
	CreateTileFlyoutUpdaterForApplication() *ITileFlyoutUpdater
	CreateTileFlyoutUpdaterForApplicationWithId(applicationId string) *ITileFlyoutUpdater
	CreateTileFlyoutUpdaterForSecondaryTile(tileId string) *ITileFlyoutUpdater
	GetTemplateContent(type_ TileFlyoutTemplateType) unsafe.Pointer
}

type ITileFlyoutUpdateManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateTileFlyoutUpdaterForApplication       uintptr
	CreateTileFlyoutUpdaterForApplicationWithId uintptr
	CreateTileFlyoutUpdaterForSecondaryTile     uintptr
	GetTemplateContent                          uintptr
}

type ITileFlyoutUpdateManagerStatics struct {
	win32.IInspectable
}

func (this *ITileFlyoutUpdateManagerStatics) Vtbl() *ITileFlyoutUpdateManagerStaticsVtbl {
	return (*ITileFlyoutUpdateManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITileFlyoutUpdateManagerStatics) CreateTileFlyoutUpdaterForApplication() *ITileFlyoutUpdater {
	var _result *ITileFlyoutUpdater
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateTileFlyoutUpdaterForApplication, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITileFlyoutUpdateManagerStatics) CreateTileFlyoutUpdaterForApplicationWithId(applicationId string) *ITileFlyoutUpdater {
	var _result *ITileFlyoutUpdater
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateTileFlyoutUpdaterForApplicationWithId, uintptr(unsafe.Pointer(this)), NewHStr(applicationId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITileFlyoutUpdateManagerStatics) CreateTileFlyoutUpdaterForSecondaryTile(tileId string) *ITileFlyoutUpdater {
	var _result *ITileFlyoutUpdater
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateTileFlyoutUpdaterForSecondaryTile, uintptr(unsafe.Pointer(this)), NewHStr(tileId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITileFlyoutUpdateManagerStatics) GetTemplateContent(type_ TileFlyoutTemplateType) unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTemplateContent, uintptr(unsafe.Pointer(this)), uintptr(type_), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 8D40C76A-C465-4052-A740-5C2654C1A089
var IID_ITileFlyoutUpdater = syscall.GUID{0x8D40C76A, 0xC465, 0x4052,
	[8]byte{0xA7, 0x40, 0x5C, 0x26, 0x54, 0xC1, 0xA0, 0x89}}

type ITileFlyoutUpdaterInterface interface {
	win32.IInspectableInterface
	Update(notification *ITileFlyoutNotification)
	Clear()
	StartPeriodicUpdate(tileFlyoutContent *IUriRuntimeClass, requestedInterval PeriodicUpdateRecurrence)
	StartPeriodicUpdateAtTime(tileFlyoutContent *IUriRuntimeClass, startTime DateTime, requestedInterval PeriodicUpdateRecurrence)
	StopPeriodicUpdate()
	Get_Setting() NotificationSetting
}

type ITileFlyoutUpdaterVtbl struct {
	win32.IInspectableVtbl
	Update                    uintptr
	Clear                     uintptr
	StartPeriodicUpdate       uintptr
	StartPeriodicUpdateAtTime uintptr
	StopPeriodicUpdate        uintptr
	Get_Setting               uintptr
}

type ITileFlyoutUpdater struct {
	win32.IInspectable
}

func (this *ITileFlyoutUpdater) Vtbl() *ITileFlyoutUpdaterVtbl {
	return (*ITileFlyoutUpdaterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITileFlyoutUpdater) Update(notification *ITileFlyoutNotification) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Update, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(notification)))
	_ = _hr
}

func (this *ITileFlyoutUpdater) Clear() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Clear, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ITileFlyoutUpdater) StartPeriodicUpdate(tileFlyoutContent *IUriRuntimeClass, requestedInterval PeriodicUpdateRecurrence) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartPeriodicUpdate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(tileFlyoutContent)), uintptr(requestedInterval))
	_ = _hr
}

func (this *ITileFlyoutUpdater) StartPeriodicUpdateAtTime(tileFlyoutContent *IUriRuntimeClass, startTime DateTime, requestedInterval PeriodicUpdateRecurrence) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartPeriodicUpdateAtTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(tileFlyoutContent)), *(*uintptr)(unsafe.Pointer(&startTime)), uintptr(requestedInterval))
	_ = _hr
}

func (this *ITileFlyoutUpdater) StopPeriodicUpdate() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopPeriodicUpdate, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ITileFlyoutUpdater) Get_Setting() NotificationSetting {
	var _result NotificationSetting
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Setting, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// EBAEC8FA-50EC-4C18-B4D0-3AF02E5540AB
var IID_ITileNotification = syscall.GUID{0xEBAEC8FA, 0x50EC, 0x4C18,
	[8]byte{0xB4, 0xD0, 0x3A, 0xF0, 0x2E, 0x55, 0x40, 0xAB}}

type ITileNotificationInterface interface {
	win32.IInspectableInterface
	Get_Content() unsafe.Pointer
	Put_ExpirationTime(value *IReference[DateTime])
	Get_ExpirationTime() *IReference[DateTime]
	Put_Tag(value string)
	Get_Tag() string
}

type ITileNotificationVtbl struct {
	win32.IInspectableVtbl
	Get_Content        uintptr
	Put_ExpirationTime uintptr
	Get_ExpirationTime uintptr
	Put_Tag            uintptr
	Get_Tag            uintptr
}

type ITileNotification struct {
	win32.IInspectable
}

func (this *ITileNotification) Vtbl() *ITileNotificationVtbl {
	return (*ITileNotificationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITileNotification) Get_Content() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Content, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITileNotification) Put_ExpirationTime(value *IReference[DateTime]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ExpirationTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ITileNotification) Get_ExpirationTime() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExpirationTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITileNotification) Put_Tag(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Tag, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ITileNotification) Get_Tag() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Tag, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// C6ABDD6E-4928-46C8-BDBF-81A047DEA0D4
var IID_ITileNotificationFactory = syscall.GUID{0xC6ABDD6E, 0x4928, 0x46C8,
	[8]byte{0xBD, 0xBF, 0x81, 0xA0, 0x47, 0xDE, 0xA0, 0xD4}}

type ITileNotificationFactoryInterface interface {
	win32.IInspectableInterface
	CreateTileNotification(content unsafe.Pointer) *ITileNotification
}

type ITileNotificationFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateTileNotification uintptr
}

type ITileNotificationFactory struct {
	win32.IInspectable
}

func (this *ITileNotificationFactory) Vtbl() *ITileNotificationFactoryVtbl {
	return (*ITileNotificationFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITileNotificationFactory) CreateTileNotification(content unsafe.Pointer) *ITileNotification {
	var _result *ITileNotification
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateTileNotification, uintptr(unsafe.Pointer(this)), uintptr(content), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 55141348-2EE2-4E2D-9CC1-216A20DECC9F
var IID_ITileUpdateManagerForUser = syscall.GUID{0x55141348, 0x2EE2, 0x4E2D,
	[8]byte{0x9C, 0xC1, 0x21, 0x6A, 0x20, 0xDE, 0xCC, 0x9F}}

type ITileUpdateManagerForUserInterface interface {
	win32.IInspectableInterface
	CreateTileUpdaterForApplication() *ITileUpdater
	CreateTileUpdaterForApplicationWithId(applicationId string) *ITileUpdater
	CreateTileUpdaterForSecondaryTile(tileId string) *ITileUpdater
	Get_User() *IUser
}

type ITileUpdateManagerForUserVtbl struct {
	win32.IInspectableVtbl
	CreateTileUpdaterForApplication       uintptr
	CreateTileUpdaterForApplicationWithId uintptr
	CreateTileUpdaterForSecondaryTile     uintptr
	Get_User                              uintptr
}

type ITileUpdateManagerForUser struct {
	win32.IInspectable
}

func (this *ITileUpdateManagerForUser) Vtbl() *ITileUpdateManagerForUserVtbl {
	return (*ITileUpdateManagerForUserVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITileUpdateManagerForUser) CreateTileUpdaterForApplication() *ITileUpdater {
	var _result *ITileUpdater
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateTileUpdaterForApplication, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITileUpdateManagerForUser) CreateTileUpdaterForApplicationWithId(applicationId string) *ITileUpdater {
	var _result *ITileUpdater
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateTileUpdaterForApplicationWithId, uintptr(unsafe.Pointer(this)), NewHStr(applicationId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITileUpdateManagerForUser) CreateTileUpdaterForSecondaryTile(tileId string) *ITileUpdater {
	var _result *ITileUpdater
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateTileUpdaterForSecondaryTile, uintptr(unsafe.Pointer(this)), NewHStr(tileId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITileUpdateManagerForUser) Get_User() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_User, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DA159E5D-3EA9-4986-8D84-B09D5E12276D
var IID_ITileUpdateManagerStatics = syscall.GUID{0xDA159E5D, 0x3EA9, 0x4986,
	[8]byte{0x8D, 0x84, 0xB0, 0x9D, 0x5E, 0x12, 0x27, 0x6D}}

type ITileUpdateManagerStaticsInterface interface {
	win32.IInspectableInterface
	CreateTileUpdaterForApplication() *ITileUpdater
	CreateTileUpdaterForApplicationWithId(applicationId string) *ITileUpdater
	CreateTileUpdaterForSecondaryTile(tileId string) *ITileUpdater
	GetTemplateContent(type_ TileTemplateType) unsafe.Pointer
}

type ITileUpdateManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateTileUpdaterForApplication       uintptr
	CreateTileUpdaterForApplicationWithId uintptr
	CreateTileUpdaterForSecondaryTile     uintptr
	GetTemplateContent                    uintptr
}

type ITileUpdateManagerStatics struct {
	win32.IInspectable
}

func (this *ITileUpdateManagerStatics) Vtbl() *ITileUpdateManagerStaticsVtbl {
	return (*ITileUpdateManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITileUpdateManagerStatics) CreateTileUpdaterForApplication() *ITileUpdater {
	var _result *ITileUpdater
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateTileUpdaterForApplication, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITileUpdateManagerStatics) CreateTileUpdaterForApplicationWithId(applicationId string) *ITileUpdater {
	var _result *ITileUpdater
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateTileUpdaterForApplicationWithId, uintptr(unsafe.Pointer(this)), NewHStr(applicationId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITileUpdateManagerStatics) CreateTileUpdaterForSecondaryTile(tileId string) *ITileUpdater {
	var _result *ITileUpdater
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateTileUpdaterForSecondaryTile, uintptr(unsafe.Pointer(this)), NewHStr(tileId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITileUpdateManagerStatics) GetTemplateContent(type_ TileTemplateType) unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTemplateContent, uintptr(unsafe.Pointer(this)), uintptr(type_), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 731C1DDC-8E14-4B7C-A34B-9D22DE76C84D
var IID_ITileUpdateManagerStatics2 = syscall.GUID{0x731C1DDC, 0x8E14, 0x4B7C,
	[8]byte{0xA3, 0x4B, 0x9D, 0x22, 0xDE, 0x76, 0xC8, 0x4D}}

type ITileUpdateManagerStatics2Interface interface {
	win32.IInspectableInterface
	GetForUser(user *IUser) *ITileUpdateManagerForUser
}

type ITileUpdateManagerStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetForUser uintptr
}

type ITileUpdateManagerStatics2 struct {
	win32.IInspectable
}

func (this *ITileUpdateManagerStatics2) Vtbl() *ITileUpdateManagerStatics2Vtbl {
	return (*ITileUpdateManagerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITileUpdateManagerStatics2) GetForUser(user *IUser) *ITileUpdateManagerForUser {
	var _result *ITileUpdateManagerForUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0942A48B-1D91-44EC-9243-C1E821C29A20
var IID_ITileUpdater = syscall.GUID{0x0942A48B, 0x1D91, 0x44EC,
	[8]byte{0x92, 0x43, 0xC1, 0xE8, 0x21, 0xC2, 0x9A, 0x20}}

type ITileUpdaterInterface interface {
	win32.IInspectableInterface
	Update(notification *ITileNotification)
	Clear()
	EnableNotificationQueue(enable bool)
	Get_Setting() NotificationSetting
	AddToSchedule(scheduledTile *IScheduledTileNotification)
	RemoveFromSchedule(scheduledTile *IScheduledTileNotification)
	GetScheduledTileNotifications() *IVectorView[*IScheduledTileNotification]
	StartPeriodicUpdate(tileContent *IUriRuntimeClass, requestedInterval PeriodicUpdateRecurrence)
	StartPeriodicUpdateAtTime(tileContent *IUriRuntimeClass, startTime DateTime, requestedInterval PeriodicUpdateRecurrence)
	StopPeriodicUpdate()
	StartPeriodicUpdateBatch(tileContents *IIterable[*IUriRuntimeClass], requestedInterval PeriodicUpdateRecurrence)
	StartPeriodicUpdateBatchAtTime(tileContents *IIterable[*IUriRuntimeClass], startTime DateTime, requestedInterval PeriodicUpdateRecurrence)
}

type ITileUpdaterVtbl struct {
	win32.IInspectableVtbl
	Update                         uintptr
	Clear                          uintptr
	EnableNotificationQueue        uintptr
	Get_Setting                    uintptr
	AddToSchedule                  uintptr
	RemoveFromSchedule             uintptr
	GetScheduledTileNotifications  uintptr
	StartPeriodicUpdate            uintptr
	StartPeriodicUpdateAtTime      uintptr
	StopPeriodicUpdate             uintptr
	StartPeriodicUpdateBatch       uintptr
	StartPeriodicUpdateBatchAtTime uintptr
}

type ITileUpdater struct {
	win32.IInspectable
}

func (this *ITileUpdater) Vtbl() *ITileUpdaterVtbl {
	return (*ITileUpdaterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITileUpdater) Update(notification *ITileNotification) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Update, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(notification)))
	_ = _hr
}

func (this *ITileUpdater) Clear() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Clear, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ITileUpdater) EnableNotificationQueue(enable bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnableNotificationQueue, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&enable))))
	_ = _hr
}

func (this *ITileUpdater) Get_Setting() NotificationSetting {
	var _result NotificationSetting
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Setting, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITileUpdater) AddToSchedule(scheduledTile *IScheduledTileNotification) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddToSchedule, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(scheduledTile)))
	_ = _hr
}

func (this *ITileUpdater) RemoveFromSchedule(scheduledTile *IScheduledTileNotification) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveFromSchedule, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(scheduledTile)))
	_ = _hr
}

func (this *ITileUpdater) GetScheduledTileNotifications() *IVectorView[*IScheduledTileNotification] {
	var _result *IVectorView[*IScheduledTileNotification]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetScheduledTileNotifications, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITileUpdater) StartPeriodicUpdate(tileContent *IUriRuntimeClass, requestedInterval PeriodicUpdateRecurrence) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartPeriodicUpdate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(tileContent)), uintptr(requestedInterval))
	_ = _hr
}

func (this *ITileUpdater) StartPeriodicUpdateAtTime(tileContent *IUriRuntimeClass, startTime DateTime, requestedInterval PeriodicUpdateRecurrence) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartPeriodicUpdateAtTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(tileContent)), *(*uintptr)(unsafe.Pointer(&startTime)), uintptr(requestedInterval))
	_ = _hr
}

func (this *ITileUpdater) StopPeriodicUpdate() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopPeriodicUpdate, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ITileUpdater) StartPeriodicUpdateBatch(tileContents *IIterable[*IUriRuntimeClass], requestedInterval PeriodicUpdateRecurrence) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartPeriodicUpdateBatch, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(tileContents)), uintptr(requestedInterval))
	_ = _hr
}

func (this *ITileUpdater) StartPeriodicUpdateBatchAtTime(tileContents *IIterable[*IUriRuntimeClass], startTime DateTime, requestedInterval PeriodicUpdateRecurrence) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartPeriodicUpdateBatchAtTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(tileContents)), *(*uintptr)(unsafe.Pointer(&startTime)), uintptr(requestedInterval))
	_ = _hr
}

// A2266E12-15EE-43ED-83F5-65B352BB1A84
var IID_ITileUpdater2 = syscall.GUID{0xA2266E12, 0x15EE, 0x43ED,
	[8]byte{0x83, 0xF5, 0x65, 0xB3, 0x52, 0xBB, 0x1A, 0x84}}

type ITileUpdater2Interface interface {
	win32.IInspectableInterface
	EnableNotificationQueueForSquare150x150(enable bool)
	EnableNotificationQueueForWide310x150(enable bool)
	EnableNotificationQueueForSquare310x310(enable bool)
}

type ITileUpdater2Vtbl struct {
	win32.IInspectableVtbl
	EnableNotificationQueueForSquare150x150 uintptr
	EnableNotificationQueueForWide310x150   uintptr
	EnableNotificationQueueForSquare310x310 uintptr
}

type ITileUpdater2 struct {
	win32.IInspectable
}

func (this *ITileUpdater2) Vtbl() *ITileUpdater2Vtbl {
	return (*ITileUpdater2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITileUpdater2) EnableNotificationQueueForSquare150x150(enable bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnableNotificationQueueForSquare150x150, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&enable))))
	_ = _hr
}

func (this *ITileUpdater2) EnableNotificationQueueForWide310x150(enable bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnableNotificationQueueForWide310x150, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&enable))))
	_ = _hr
}

func (this *ITileUpdater2) EnableNotificationQueueForSquare310x310(enable bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnableNotificationQueueForSquare310x310, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&enable))))
	_ = _hr
}

// E3BF92F3-C197-436F-8265-0625824F8DAC
var IID_IToastActivatedEventArgs = syscall.GUID{0xE3BF92F3, 0xC197, 0x436F,
	[8]byte{0x82, 0x65, 0x06, 0x25, 0x82, 0x4F, 0x8D, 0xAC}}

type IToastActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Arguments() string
}

type IToastActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Arguments uintptr
}

type IToastActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IToastActivatedEventArgs) Vtbl() *IToastActivatedEventArgsVtbl {
	return (*IToastActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastActivatedEventArgs) Get_Arguments() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Arguments, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// AB7DA512-CC61-568E-81BE-304AC31038FA
var IID_IToastActivatedEventArgs2 = syscall.GUID{0xAB7DA512, 0xCC61, 0x568E,
	[8]byte{0x81, 0xBE, 0x30, 0x4A, 0xC3, 0x10, 0x38, 0xFA}}

type IToastActivatedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_UserInput() *IPropertySet
}

type IToastActivatedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_UserInput uintptr
}

type IToastActivatedEventArgs2 struct {
	win32.IInspectable
}

func (this *IToastActivatedEventArgs2) Vtbl() *IToastActivatedEventArgs2Vtbl {
	return (*IToastActivatedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastActivatedEventArgs2) Get_UserInput() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0A8BC3B0-E0BE-4858-BC2A-89DFE0B32863
var IID_IToastCollection = syscall.GUID{0x0A8BC3B0, 0xE0BE, 0x4858,
	[8]byte{0xBC, 0x2A, 0x89, 0xDF, 0xE0, 0xB3, 0x28, 0x63}}

type IToastCollectionInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Get_DisplayName() string
	Put_DisplayName(value string)
	Get_LaunchArgs() string
	Put_LaunchArgs(value string)
	Get_Icon() *IUriRuntimeClass
	Put_Icon(value *IUriRuntimeClass)
}

type IToastCollectionVtbl struct {
	win32.IInspectableVtbl
	Get_Id          uintptr
	Get_DisplayName uintptr
	Put_DisplayName uintptr
	Get_LaunchArgs  uintptr
	Put_LaunchArgs  uintptr
	Get_Icon        uintptr
	Put_Icon        uintptr
}

type IToastCollection struct {
	win32.IInspectable
}

func (this *IToastCollection) Vtbl() *IToastCollectionVtbl {
	return (*IToastCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastCollection) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IToastCollection) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IToastCollection) Put_DisplayName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DisplayName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IToastCollection) Get_LaunchArgs() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LaunchArgs, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IToastCollection) Put_LaunchArgs(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LaunchArgs, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IToastCollection) Get_Icon() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Icon, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IToastCollection) Put_Icon(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Icon, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 164DD3D7-73C4-44F7-B4FF-FB6D4BF1F4C6
var IID_IToastCollectionFactory = syscall.GUID{0x164DD3D7, 0x73C4, 0x44F7,
	[8]byte{0xB4, 0xFF, 0xFB, 0x6D, 0x4B, 0xF1, 0xF4, 0xC6}}

type IToastCollectionFactoryInterface interface {
	win32.IInspectableInterface
	CreateInstance(collectionId string, displayName string, launchArgs string, iconUri *IUriRuntimeClass) *IToastCollection
}

type IToastCollectionFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateInstance uintptr
}

type IToastCollectionFactory struct {
	win32.IInspectable
}

func (this *IToastCollectionFactory) Vtbl() *IToastCollectionFactoryVtbl {
	return (*IToastCollectionFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastCollectionFactory) CreateInstance(collectionId string, displayName string, launchArgs string, iconUri *IUriRuntimeClass) *IToastCollection {
	var _result *IToastCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateInstance, uintptr(unsafe.Pointer(this)), NewHStr(collectionId).Ptr, NewHStr(displayName).Ptr, NewHStr(launchArgs).Ptr, uintptr(unsafe.Pointer(iconUri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2A1821FE-179D-49BC-B79D-A527920D3665
var IID_IToastCollectionManager = syscall.GUID{0x2A1821FE, 0x179D, 0x49BC,
	[8]byte{0xB7, 0x9D, 0xA5, 0x27, 0x92, 0x0D, 0x36, 0x65}}

type IToastCollectionManagerInterface interface {
	win32.IInspectableInterface
	SaveToastCollectionAsync(collection *IToastCollection) *IAsyncAction
	FindAllToastCollectionsAsync() *IAsyncOperation[*IVectorView[*IToastCollection]]
	GetToastCollectionAsync(collectionId string) *IAsyncOperation[*IToastCollection]
	RemoveToastCollectionAsync(collectionId string) *IAsyncAction
	RemoveAllToastCollectionsAsync() *IAsyncAction
	Get_User() *IUser
	Get_AppId() string
}

type IToastCollectionManagerVtbl struct {
	win32.IInspectableVtbl
	SaveToastCollectionAsync       uintptr
	FindAllToastCollectionsAsync   uintptr
	GetToastCollectionAsync        uintptr
	RemoveToastCollectionAsync     uintptr
	RemoveAllToastCollectionsAsync uintptr
	Get_User                       uintptr
	Get_AppId                      uintptr
}

type IToastCollectionManager struct {
	win32.IInspectable
}

func (this *IToastCollectionManager) Vtbl() *IToastCollectionManagerVtbl {
	return (*IToastCollectionManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastCollectionManager) SaveToastCollectionAsync(collection *IToastCollection) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SaveToastCollectionAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(collection)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IToastCollectionManager) FindAllToastCollectionsAsync() *IAsyncOperation[*IVectorView[*IToastCollection]] {
	var _result *IAsyncOperation[*IVectorView[*IToastCollection]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllToastCollectionsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IToastCollectionManager) GetToastCollectionAsync(collectionId string) *IAsyncOperation[*IToastCollection] {
	var _result *IAsyncOperation[*IToastCollection]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetToastCollectionAsync, uintptr(unsafe.Pointer(this)), NewHStr(collectionId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IToastCollectionManager) RemoveToastCollectionAsync(collectionId string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveToastCollectionAsync, uintptr(unsafe.Pointer(this)), NewHStr(collectionId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IToastCollectionManager) RemoveAllToastCollectionsAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveAllToastCollectionsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IToastCollectionManager) Get_User() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_User, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IToastCollectionManager) Get_AppId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 3F89D935-D9CB-4538-A0F0-FFE7659938F8
var IID_IToastDismissedEventArgs = syscall.GUID{0x3F89D935, 0xD9CB, 0x4538,
	[8]byte{0xA0, 0xF0, 0xFF, 0xE7, 0x65, 0x99, 0x38, 0xF8}}

type IToastDismissedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Reason() ToastDismissalReason
}

type IToastDismissedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Reason uintptr
}

type IToastDismissedEventArgs struct {
	win32.IInspectable
}

func (this *IToastDismissedEventArgs) Vtbl() *IToastDismissedEventArgsVtbl {
	return (*IToastDismissedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastDismissedEventArgs) Get_Reason() ToastDismissalReason {
	var _result ToastDismissalReason
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reason, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 35176862-CFD4-44F8-AD64-F500FD896C3B
var IID_IToastFailedEventArgs = syscall.GUID{0x35176862, 0xCFD4, 0x44F8,
	[8]byte{0xAD, 0x64, 0xF5, 0x00, 0xFD, 0x89, 0x6C, 0x3B}}

type IToastFailedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ErrorCode() HResult
}

type IToastFailedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ErrorCode uintptr
}

type IToastFailedEventArgs struct {
	win32.IInspectable
}

func (this *IToastFailedEventArgs) Vtbl() *IToastFailedEventArgsVtbl {
	return (*IToastFailedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastFailedEventArgs) Get_ErrorCode() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 997E2675-059E-4E60-8B06-1760917C8B80
var IID_IToastNotification = syscall.GUID{0x997E2675, 0x059E, 0x4E60,
	[8]byte{0x8B, 0x06, 0x17, 0x60, 0x91, 0x7C, 0x8B, 0x80}}

type IToastNotificationInterface interface {
	win32.IInspectableInterface
	Get_Content() unsafe.Pointer
	Put_ExpirationTime(value *IReference[DateTime])
	Get_ExpirationTime() *IReference[DateTime]
	Add_Dismissed(handler TypedEventHandler[*IToastNotification, *IToastDismissedEventArgs]) EventRegistrationToken
	Remove_Dismissed(token EventRegistrationToken)
	Add_Activated(handler TypedEventHandler[*IToastNotification, interface{}]) EventRegistrationToken
	Remove_Activated(token EventRegistrationToken)
	Add_Failed(handler TypedEventHandler[*IToastNotification, *IToastFailedEventArgs]) EventRegistrationToken
	Remove_Failed(token EventRegistrationToken)
}

type IToastNotificationVtbl struct {
	win32.IInspectableVtbl
	Get_Content        uintptr
	Put_ExpirationTime uintptr
	Get_ExpirationTime uintptr
	Add_Dismissed      uintptr
	Remove_Dismissed   uintptr
	Add_Activated      uintptr
	Remove_Activated   uintptr
	Add_Failed         uintptr
	Remove_Failed      uintptr
}

type IToastNotification struct {
	win32.IInspectable
}

func (this *IToastNotification) Vtbl() *IToastNotificationVtbl {
	return (*IToastNotificationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastNotification) Get_Content() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Content, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IToastNotification) Put_ExpirationTime(value *IReference[DateTime]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ExpirationTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IToastNotification) Get_ExpirationTime() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExpirationTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IToastNotification) Add_Dismissed(handler TypedEventHandler[*IToastNotification, *IToastDismissedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Dismissed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IToastNotification) Remove_Dismissed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Dismissed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IToastNotification) Add_Activated(handler TypedEventHandler[*IToastNotification, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Activated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IToastNotification) Remove_Activated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Activated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IToastNotification) Add_Failed(handler TypedEventHandler[*IToastNotification, *IToastFailedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Failed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IToastNotification) Remove_Failed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Failed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 9DFB9FD1-143A-490E-90BF-B9FBA7132DE7
var IID_IToastNotification2 = syscall.GUID{0x9DFB9FD1, 0x143A, 0x490E,
	[8]byte{0x90, 0xBF, 0xB9, 0xFB, 0xA7, 0x13, 0x2D, 0xE7}}

type IToastNotification2Interface interface {
	win32.IInspectableInterface
	Put_Tag(value string)
	Get_Tag() string
	Put_Group(value string)
	Get_Group() string
	Put_SuppressPopup(value bool)
	Get_SuppressPopup() bool
}

type IToastNotification2Vtbl struct {
	win32.IInspectableVtbl
	Put_Tag           uintptr
	Get_Tag           uintptr
	Put_Group         uintptr
	Get_Group         uintptr
	Put_SuppressPopup uintptr
	Get_SuppressPopup uintptr
}

type IToastNotification2 struct {
	win32.IInspectable
}

func (this *IToastNotification2) Vtbl() *IToastNotification2Vtbl {
	return (*IToastNotification2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastNotification2) Put_Tag(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Tag, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IToastNotification2) Get_Tag() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Tag, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IToastNotification2) Put_Group(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Group, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IToastNotification2) Get_Group() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Group, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IToastNotification2) Put_SuppressPopup(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SuppressPopup, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IToastNotification2) Get_SuppressPopup() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SuppressPopup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 31E8AED8-8141-4F99-BC0A-C4ED21297D77
var IID_IToastNotification3 = syscall.GUID{0x31E8AED8, 0x8141, 0x4F99,
	[8]byte{0xBC, 0x0A, 0xC4, 0xED, 0x21, 0x29, 0x7D, 0x77}}

type IToastNotification3Interface interface {
	win32.IInspectableInterface
	Get_NotificationMirroring() NotificationMirroring
	Put_NotificationMirroring(value NotificationMirroring)
	Get_RemoteId() string
	Put_RemoteId(value string)
}

type IToastNotification3Vtbl struct {
	win32.IInspectableVtbl
	Get_NotificationMirroring uintptr
	Put_NotificationMirroring uintptr
	Get_RemoteId              uintptr
	Put_RemoteId              uintptr
}

type IToastNotification3 struct {
	win32.IInspectable
}

func (this *IToastNotification3) Vtbl() *IToastNotification3Vtbl {
	return (*IToastNotification3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastNotification3) Get_NotificationMirroring() NotificationMirroring {
	var _result NotificationMirroring
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NotificationMirroring, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IToastNotification3) Put_NotificationMirroring(value NotificationMirroring) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NotificationMirroring, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IToastNotification3) Get_RemoteId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemoteId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IToastNotification3) Put_RemoteId(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RemoteId, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 15154935-28EA-4727-88E9-C58680E2D118
var IID_IToastNotification4 = syscall.GUID{0x15154935, 0x28EA, 0x4727,
	[8]byte{0x88, 0xE9, 0xC5, 0x86, 0x80, 0xE2, 0xD1, 0x18}}

type IToastNotification4Interface interface {
	win32.IInspectableInterface
	Get_Data() *INotificationData
	Put_Data(value *INotificationData)
	Get_Priority() ToastNotificationPriority
	Put_Priority(value ToastNotificationPriority)
}

type IToastNotification4Vtbl struct {
	win32.IInspectableVtbl
	Get_Data     uintptr
	Put_Data     uintptr
	Get_Priority uintptr
	Put_Priority uintptr
}

type IToastNotification4 struct {
	win32.IInspectable
}

func (this *IToastNotification4) Vtbl() *IToastNotification4Vtbl {
	return (*IToastNotification4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastNotification4) Get_Data() *INotificationData {
	var _result *INotificationData
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IToastNotification4) Put_Data(value *INotificationData) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IToastNotification4) Get_Priority() ToastNotificationPriority {
	var _result ToastNotificationPriority
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Priority, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IToastNotification4) Put_Priority(value ToastNotificationPriority) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Priority, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 43EBFE53-89AE-5C1E-A279-3AECFE9B6F54
var IID_IToastNotification6 = syscall.GUID{0x43EBFE53, 0x89AE, 0x5C1E,
	[8]byte{0xA2, 0x79, 0x3A, 0xEC, 0xFE, 0x9B, 0x6F, 0x54}}

type IToastNotification6Interface interface {
	win32.IInspectableInterface
	Get_ExpiresOnReboot() bool
	Put_ExpiresOnReboot(value bool)
}

type IToastNotification6Vtbl struct {
	win32.IInspectableVtbl
	Get_ExpiresOnReboot uintptr
	Put_ExpiresOnReboot uintptr
}

type IToastNotification6 struct {
	win32.IInspectable
}

func (this *IToastNotification6) Vtbl() *IToastNotification6Vtbl {
	return (*IToastNotification6Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastNotification6) Get_ExpiresOnReboot() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExpiresOnReboot, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IToastNotification6) Put_ExpiresOnReboot(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ExpiresOnReboot, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 9445135A-38F3-42F6-96AA-7955B0F03DA2
var IID_IToastNotificationActionTriggerDetail = syscall.GUID{0x9445135A, 0x38F3, 0x42F6,
	[8]byte{0x96, 0xAA, 0x79, 0x55, 0xB0, 0xF0, 0x3D, 0xA2}}

type IToastNotificationActionTriggerDetailInterface interface {
	win32.IInspectableInterface
	Get_Argument() string
	Get_UserInput() *IPropertySet
}

type IToastNotificationActionTriggerDetailVtbl struct {
	win32.IInspectableVtbl
	Get_Argument  uintptr
	Get_UserInput uintptr
}

type IToastNotificationActionTriggerDetail struct {
	win32.IInspectable
}

func (this *IToastNotificationActionTriggerDetail) Vtbl() *IToastNotificationActionTriggerDetailVtbl {
	return (*IToastNotificationActionTriggerDetailVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastNotificationActionTriggerDetail) Get_Argument() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Argument, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IToastNotificationActionTriggerDetail) Get_UserInput() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 04124B20-82C6-4229-B109-FD9ED4662B53
var IID_IToastNotificationFactory = syscall.GUID{0x04124B20, 0x82C6, 0x4229,
	[8]byte{0xB1, 0x09, 0xFD, 0x9E, 0xD4, 0x66, 0x2B, 0x53}}

type IToastNotificationFactoryInterface interface {
	win32.IInspectableInterface
	CreateToastNotification(content unsafe.Pointer) *IToastNotification
}

type IToastNotificationFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateToastNotification uintptr
}

type IToastNotificationFactory struct {
	win32.IInspectable
}

func (this *IToastNotificationFactory) Vtbl() *IToastNotificationFactoryVtbl {
	return (*IToastNotificationFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastNotificationFactory) CreateToastNotification(content unsafe.Pointer) *IToastNotification {
	var _result *IToastNotification
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateToastNotification, uintptr(unsafe.Pointer(this)), uintptr(content), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5CADDC63-01D3-4C97-986F-0533483FEE14
var IID_IToastNotificationHistory = syscall.GUID{0x5CADDC63, 0x01D3, 0x4C97,
	[8]byte{0x98, 0x6F, 0x05, 0x33, 0x48, 0x3F, 0xEE, 0x14}}

type IToastNotificationHistoryInterface interface {
	win32.IInspectableInterface
	RemoveGroup(group string)
	RemoveGroupWithId(group string, applicationId string)
	RemoveGroupedTagWithId(tag string, group string, applicationId string)
	RemoveGroupedTag(tag string, group string)
	Remove(tag string)
	Clear()
	ClearWithId(applicationId string)
}

type IToastNotificationHistoryVtbl struct {
	win32.IInspectableVtbl
	RemoveGroup            uintptr
	RemoveGroupWithId      uintptr
	RemoveGroupedTagWithId uintptr
	RemoveGroupedTag       uintptr
	Remove                 uintptr
	Clear                  uintptr
	ClearWithId            uintptr
}

type IToastNotificationHistory struct {
	win32.IInspectable
}

func (this *IToastNotificationHistory) Vtbl() *IToastNotificationHistoryVtbl {
	return (*IToastNotificationHistoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastNotificationHistory) RemoveGroup(group string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveGroup, uintptr(unsafe.Pointer(this)), NewHStr(group).Ptr)
	_ = _hr
}

func (this *IToastNotificationHistory) RemoveGroupWithId(group string, applicationId string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveGroupWithId, uintptr(unsafe.Pointer(this)), NewHStr(group).Ptr, NewHStr(applicationId).Ptr)
	_ = _hr
}

func (this *IToastNotificationHistory) RemoveGroupedTagWithId(tag string, group string, applicationId string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveGroupedTagWithId, uintptr(unsafe.Pointer(this)), NewHStr(tag).Ptr, NewHStr(group).Ptr, NewHStr(applicationId).Ptr)
	_ = _hr
}

func (this *IToastNotificationHistory) RemoveGroupedTag(tag string, group string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveGroupedTag, uintptr(unsafe.Pointer(this)), NewHStr(tag).Ptr, NewHStr(group).Ptr)
	_ = _hr
}

func (this *IToastNotificationHistory) Remove(tag string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove, uintptr(unsafe.Pointer(this)), NewHStr(tag).Ptr)
	_ = _hr
}

func (this *IToastNotificationHistory) Clear() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Clear, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IToastNotificationHistory) ClearWithId(applicationId string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearWithId, uintptr(unsafe.Pointer(this)), NewHStr(applicationId).Ptr)
	_ = _hr
}

// 3BC3D253-2F31-4092-9129-8AD5ABF067DA
var IID_IToastNotificationHistory2 = syscall.GUID{0x3BC3D253, 0x2F31, 0x4092,
	[8]byte{0x91, 0x29, 0x8A, 0xD5, 0xAB, 0xF0, 0x67, 0xDA}}

type IToastNotificationHistory2Interface interface {
	win32.IInspectableInterface
	GetHistory() *IVectorView[*IToastNotification]
	GetHistoryWithId(applicationId string) *IVectorView[*IToastNotification]
}

type IToastNotificationHistory2Vtbl struct {
	win32.IInspectableVtbl
	GetHistory       uintptr
	GetHistoryWithId uintptr
}

type IToastNotificationHistory2 struct {
	win32.IInspectable
}

func (this *IToastNotificationHistory2) Vtbl() *IToastNotificationHistory2Vtbl {
	return (*IToastNotificationHistory2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastNotificationHistory2) GetHistory() *IVectorView[*IToastNotification] {
	var _result *IVectorView[*IToastNotification]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetHistory, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IToastNotificationHistory2) GetHistoryWithId(applicationId string) *IVectorView[*IToastNotification] {
	var _result *IVectorView[*IToastNotification]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetHistoryWithId, uintptr(unsafe.Pointer(this)), NewHStr(applicationId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DB037FFA-0068-412C-9C83-267C37F65670
var IID_IToastNotificationHistoryChangedTriggerDetail = syscall.GUID{0xDB037FFA, 0x0068, 0x412C,
	[8]byte{0x9C, 0x83, 0x26, 0x7C, 0x37, 0xF6, 0x56, 0x70}}

type IToastNotificationHistoryChangedTriggerDetailInterface interface {
	win32.IInspectableInterface
	Get_ChangeType() ToastHistoryChangedType
}

type IToastNotificationHistoryChangedTriggerDetailVtbl struct {
	win32.IInspectableVtbl
	Get_ChangeType uintptr
}

type IToastNotificationHistoryChangedTriggerDetail struct {
	win32.IInspectable
}

func (this *IToastNotificationHistoryChangedTriggerDetail) Vtbl() *IToastNotificationHistoryChangedTriggerDetailVtbl {
	return (*IToastNotificationHistoryChangedTriggerDetailVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastNotificationHistoryChangedTriggerDetail) Get_ChangeType() ToastHistoryChangedType {
	var _result ToastHistoryChangedType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ChangeType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 0B36E982-C871-49FB-BABB-25BDBC4CC45B
var IID_IToastNotificationHistoryChangedTriggerDetail2 = syscall.GUID{0x0B36E982, 0xC871, 0x49FB,
	[8]byte{0xBA, 0xBB, 0x25, 0xBD, 0xBC, 0x4C, 0xC4, 0x5B}}

type IToastNotificationHistoryChangedTriggerDetail2Interface interface {
	win32.IInspectableInterface
	Get_CollectionId() string
}

type IToastNotificationHistoryChangedTriggerDetail2Vtbl struct {
	win32.IInspectableVtbl
	Get_CollectionId uintptr
}

type IToastNotificationHistoryChangedTriggerDetail2 struct {
	win32.IInspectable
}

func (this *IToastNotificationHistoryChangedTriggerDetail2) Vtbl() *IToastNotificationHistoryChangedTriggerDetail2Vtbl {
	return (*IToastNotificationHistoryChangedTriggerDetail2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastNotificationHistoryChangedTriggerDetail2) Get_CollectionId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CollectionId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 79AB57F6-43FE-487B-8A7F-99567200AE94
var IID_IToastNotificationManagerForUser = syscall.GUID{0x79AB57F6, 0x43FE, 0x487B,
	[8]byte{0x8A, 0x7F, 0x99, 0x56, 0x72, 0x00, 0xAE, 0x94}}

type IToastNotificationManagerForUserInterface interface {
	win32.IInspectableInterface
	CreateToastNotifier() *IToastNotifier
	CreateToastNotifierWithId(applicationId string) *IToastNotifier
	Get_History() *IToastNotificationHistory
	Get_User() *IUser
}

type IToastNotificationManagerForUserVtbl struct {
	win32.IInspectableVtbl
	CreateToastNotifier       uintptr
	CreateToastNotifierWithId uintptr
	Get_History               uintptr
	Get_User                  uintptr
}

type IToastNotificationManagerForUser struct {
	win32.IInspectable
}

func (this *IToastNotificationManagerForUser) Vtbl() *IToastNotificationManagerForUserVtbl {
	return (*IToastNotificationManagerForUserVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastNotificationManagerForUser) CreateToastNotifier() *IToastNotifier {
	var _result *IToastNotifier
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateToastNotifier, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IToastNotificationManagerForUser) CreateToastNotifierWithId(applicationId string) *IToastNotifier {
	var _result *IToastNotifier
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateToastNotifierWithId, uintptr(unsafe.Pointer(this)), NewHStr(applicationId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IToastNotificationManagerForUser) Get_History() *IToastNotificationHistory {
	var _result *IToastNotificationHistory
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_History, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IToastNotificationManagerForUser) Get_User() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_User, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 679C64B7-81AB-42C2-8819-C958767753F4
var IID_IToastNotificationManagerForUser2 = syscall.GUID{0x679C64B7, 0x81AB, 0x42C2,
	[8]byte{0x88, 0x19, 0xC9, 0x58, 0x76, 0x77, 0x53, 0xF4}}

type IToastNotificationManagerForUser2Interface interface {
	win32.IInspectableInterface
	GetToastNotifierForToastCollectionIdAsync(collectionId string) *IAsyncOperation[*IToastNotifier]
	GetHistoryForToastCollectionIdAsync(collectionId string) *IAsyncOperation[*IToastNotificationHistory]
	GetToastCollectionManager() *IToastCollectionManager
	GetToastCollectionManagerWithAppId(appId string) *IToastCollectionManager
}

type IToastNotificationManagerForUser2Vtbl struct {
	win32.IInspectableVtbl
	GetToastNotifierForToastCollectionIdAsync uintptr
	GetHistoryForToastCollectionIdAsync       uintptr
	GetToastCollectionManager                 uintptr
	GetToastCollectionManagerWithAppId        uintptr
}

type IToastNotificationManagerForUser2 struct {
	win32.IInspectable
}

func (this *IToastNotificationManagerForUser2) Vtbl() *IToastNotificationManagerForUser2Vtbl {
	return (*IToastNotificationManagerForUser2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastNotificationManagerForUser2) GetToastNotifierForToastCollectionIdAsync(collectionId string) *IAsyncOperation[*IToastNotifier] {
	var _result *IAsyncOperation[*IToastNotifier]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetToastNotifierForToastCollectionIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(collectionId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IToastNotificationManagerForUser2) GetHistoryForToastCollectionIdAsync(collectionId string) *IAsyncOperation[*IToastNotificationHistory] {
	var _result *IAsyncOperation[*IToastNotificationHistory]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetHistoryForToastCollectionIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(collectionId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IToastNotificationManagerForUser2) GetToastCollectionManager() *IToastCollectionManager {
	var _result *IToastCollectionManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetToastCollectionManager, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IToastNotificationManagerForUser2) GetToastCollectionManagerWithAppId(appId string) *IToastCollectionManager {
	var _result *IToastCollectionManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetToastCollectionManagerWithAppId, uintptr(unsafe.Pointer(this)), NewHStr(appId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 50AC103F-D235-4598-BBEF-98FE4D1A3AD4
var IID_IToastNotificationManagerStatics = syscall.GUID{0x50AC103F, 0xD235, 0x4598,
	[8]byte{0xBB, 0xEF, 0x98, 0xFE, 0x4D, 0x1A, 0x3A, 0xD4}}

type IToastNotificationManagerStaticsInterface interface {
	win32.IInspectableInterface
	CreateToastNotifier() *IToastNotifier
	CreateToastNotifierWithId(applicationId string) *IToastNotifier
	GetTemplateContent(type_ ToastTemplateType) unsafe.Pointer
}

type IToastNotificationManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateToastNotifier       uintptr
	CreateToastNotifierWithId uintptr
	GetTemplateContent        uintptr
}

type IToastNotificationManagerStatics struct {
	win32.IInspectable
}

func (this *IToastNotificationManagerStatics) Vtbl() *IToastNotificationManagerStaticsVtbl {
	return (*IToastNotificationManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastNotificationManagerStatics) CreateToastNotifier() *IToastNotifier {
	var _result *IToastNotifier
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateToastNotifier, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IToastNotificationManagerStatics) CreateToastNotifierWithId(applicationId string) *IToastNotifier {
	var _result *IToastNotifier
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateToastNotifierWithId, uintptr(unsafe.Pointer(this)), NewHStr(applicationId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IToastNotificationManagerStatics) GetTemplateContent(type_ ToastTemplateType) unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTemplateContent, uintptr(unsafe.Pointer(this)), uintptr(type_), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7AB93C52-0E48-4750-BA9D-1A4113981847
var IID_IToastNotificationManagerStatics2 = syscall.GUID{0x7AB93C52, 0x0E48, 0x4750,
	[8]byte{0xBA, 0x9D, 0x1A, 0x41, 0x13, 0x98, 0x18, 0x47}}

type IToastNotificationManagerStatics2Interface interface {
	win32.IInspectableInterface
	Get_History() *IToastNotificationHistory
}

type IToastNotificationManagerStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_History uintptr
}

type IToastNotificationManagerStatics2 struct {
	win32.IInspectable
}

func (this *IToastNotificationManagerStatics2) Vtbl() *IToastNotificationManagerStatics2Vtbl {
	return (*IToastNotificationManagerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastNotificationManagerStatics2) Get_History() *IToastNotificationHistory {
	var _result *IToastNotificationHistory
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_History, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8F993FD3-E516-45FB-8130-398E93FA52C3
var IID_IToastNotificationManagerStatics4 = syscall.GUID{0x8F993FD3, 0xE516, 0x45FB,
	[8]byte{0x81, 0x30, 0x39, 0x8E, 0x93, 0xFA, 0x52, 0xC3}}

type IToastNotificationManagerStatics4Interface interface {
	win32.IInspectableInterface
	GetForUser(user *IUser) *IToastNotificationManagerForUser
	ConfigureNotificationMirroring(value NotificationMirroring)
}

type IToastNotificationManagerStatics4Vtbl struct {
	win32.IInspectableVtbl
	GetForUser                     uintptr
	ConfigureNotificationMirroring uintptr
}

type IToastNotificationManagerStatics4 struct {
	win32.IInspectable
}

func (this *IToastNotificationManagerStatics4) Vtbl() *IToastNotificationManagerStatics4Vtbl {
	return (*IToastNotificationManagerStatics4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastNotificationManagerStatics4) GetForUser(user *IUser) *IToastNotificationManagerForUser {
	var _result *IToastNotificationManagerForUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IToastNotificationManagerStatics4) ConfigureNotificationMirroring(value NotificationMirroring) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConfigureNotificationMirroring, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// D6F5F569-D40D-407C-8989-88CAB42CFD14
var IID_IToastNotificationManagerStatics5 = syscall.GUID{0xD6F5F569, 0xD40D, 0x407C,
	[8]byte{0x89, 0x89, 0x88, 0xCA, 0xB4, 0x2C, 0xFD, 0x14}}

type IToastNotificationManagerStatics5Interface interface {
	win32.IInspectableInterface
	GetDefault() *IToastNotificationManagerForUser
}

type IToastNotificationManagerStatics5Vtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
}

type IToastNotificationManagerStatics5 struct {
	win32.IInspectable
}

func (this *IToastNotificationManagerStatics5) Vtbl() *IToastNotificationManagerStatics5Vtbl {
	return (*IToastNotificationManagerStatics5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastNotificationManagerStatics5) GetDefault() *IToastNotificationManagerForUser {
	var _result *IToastNotificationManagerForUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 75927B93-03F3-41EC-91D3-6E5BAC1B38E7
var IID_IToastNotifier = syscall.GUID{0x75927B93, 0x03F3, 0x41EC,
	[8]byte{0x91, 0xD3, 0x6E, 0x5B, 0xAC, 0x1B, 0x38, 0xE7}}

type IToastNotifierInterface interface {
	win32.IInspectableInterface
	Show(notification *IToastNotification)
	Hide(notification *IToastNotification)
	Get_Setting() NotificationSetting
	AddToSchedule(scheduledToast *IScheduledToastNotification)
	RemoveFromSchedule(scheduledToast *IScheduledToastNotification)
	GetScheduledToastNotifications() *IVectorView[*IScheduledToastNotification]
}

type IToastNotifierVtbl struct {
	win32.IInspectableVtbl
	Show                           uintptr
	Hide                           uintptr
	Get_Setting                    uintptr
	AddToSchedule                  uintptr
	RemoveFromSchedule             uintptr
	GetScheduledToastNotifications uintptr
}

type IToastNotifier struct {
	win32.IInspectable
}

func (this *IToastNotifier) Vtbl() *IToastNotifierVtbl {
	return (*IToastNotifierVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastNotifier) Show(notification *IToastNotification) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Show, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(notification)))
	_ = _hr
}

func (this *IToastNotifier) Hide(notification *IToastNotification) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Hide, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(notification)))
	_ = _hr
}

func (this *IToastNotifier) Get_Setting() NotificationSetting {
	var _result NotificationSetting
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Setting, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IToastNotifier) AddToSchedule(scheduledToast *IScheduledToastNotification) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddToSchedule, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(scheduledToast)))
	_ = _hr
}

func (this *IToastNotifier) RemoveFromSchedule(scheduledToast *IScheduledToastNotification) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveFromSchedule, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(scheduledToast)))
	_ = _hr
}

func (this *IToastNotifier) GetScheduledToastNotifications() *IVectorView[*IScheduledToastNotification] {
	var _result *IVectorView[*IScheduledToastNotification]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetScheduledToastNotifications, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 354389C6-7C01-4BD5-9C20-604340CD2B74
var IID_IToastNotifier2 = syscall.GUID{0x354389C6, 0x7C01, 0x4BD5,
	[8]byte{0x9C, 0x20, 0x60, 0x43, 0x40, 0xCD, 0x2B, 0x74}}

type IToastNotifier2Interface interface {
	win32.IInspectableInterface
	UpdateWithTagAndGroup(data *INotificationData, tag string, group string) NotificationUpdateResult
	UpdateWithTag(data *INotificationData, tag string) NotificationUpdateResult
}

type IToastNotifier2Vtbl struct {
	win32.IInspectableVtbl
	UpdateWithTagAndGroup uintptr
	UpdateWithTag         uintptr
}

type IToastNotifier2 struct {
	win32.IInspectable
}

func (this *IToastNotifier2) Vtbl() *IToastNotifier2Vtbl {
	return (*IToastNotifier2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastNotifier2) UpdateWithTagAndGroup(data *INotificationData, tag string, group string) NotificationUpdateResult {
	var _result NotificationUpdateResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UpdateWithTagAndGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(data)), NewHStr(tag).Ptr, NewHStr(group).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IToastNotifier2) UpdateWithTag(data *INotificationData, tag string) NotificationUpdateResult {
	var _result NotificationUpdateResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UpdateWithTag, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(data)), NewHStr(tag).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// AE75A04A-3B0C-51AD-B7E8-B08AB6052549
var IID_IToastNotifier3 = syscall.GUID{0xAE75A04A, 0x3B0C, 0x51AD,
	[8]byte{0xB7, 0xE8, 0xB0, 0x8A, 0xB6, 0x05, 0x25, 0x49}}

type IToastNotifier3Interface interface {
	win32.IInspectableInterface
	Add_ScheduledToastNotificationShowing(handler TypedEventHandler[*IToastNotifier, *IScheduledToastNotificationShowingEventArgs]) EventRegistrationToken
	Remove_ScheduledToastNotificationShowing(token EventRegistrationToken)
}

type IToastNotifier3Vtbl struct {
	win32.IInspectableVtbl
	Add_ScheduledToastNotificationShowing    uintptr
	Remove_ScheduledToastNotificationShowing uintptr
}

type IToastNotifier3 struct {
	win32.IInspectable
}

func (this *IToastNotifier3) Vtbl() *IToastNotifier3Vtbl {
	return (*IToastNotifier3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastNotifier3) Add_ScheduledToastNotificationShowing(handler TypedEventHandler[*IToastNotifier, *IScheduledToastNotificationShowingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ScheduledToastNotificationShowing, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IToastNotifier3) Remove_ScheduledToastNotificationShowing(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ScheduledToastNotificationShowing, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// ADF7E52F-4E53-42D5-9C33-EB5EA515B23E
var IID_IUserNotification = syscall.GUID{0xADF7E52F, 0x4E53, 0x42D5,
	[8]byte{0x9C, 0x33, 0xEB, 0x5E, 0xA5, 0x15, 0xB2, 0x3E}}

type IUserNotificationInterface interface {
	win32.IInspectableInterface
	Get_Notification() *INotification
	Get_AppInfo() *IAppInfo
	Get_Id() uint32
	Get_CreationTime() DateTime
}

type IUserNotificationVtbl struct {
	win32.IInspectableVtbl
	Get_Notification uintptr
	Get_AppInfo      uintptr
	Get_Id           uintptr
	Get_CreationTime uintptr
}

type IUserNotification struct {
	win32.IInspectable
}

func (this *IUserNotification) Vtbl() *IUserNotificationVtbl {
	return (*IUserNotificationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUserNotification) Get_Notification() *INotification {
	var _result *INotification
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Notification, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserNotification) Get_AppInfo() *IAppInfo {
	var _result *IAppInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserNotification) Get_Id() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUserNotification) Get_CreationTime() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CreationTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// B6BD6839-79CF-4B25-82C0-0CE1EEF81F8C
var IID_IUserNotificationChangedEventArgs = syscall.GUID{0xB6BD6839, 0x79CF, 0x4B25,
	[8]byte{0x82, 0xC0, 0x0C, 0xE1, 0xEE, 0xF8, 0x1F, 0x8C}}

type IUserNotificationChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ChangeKind() UserNotificationChangedKind
	Get_UserNotificationId() uint32
}

type IUserNotificationChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ChangeKind         uintptr
	Get_UserNotificationId uintptr
}

type IUserNotificationChangedEventArgs struct {
	win32.IInspectable
}

func (this *IUserNotificationChangedEventArgs) Vtbl() *IUserNotificationChangedEventArgsVtbl {
	return (*IUserNotificationChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUserNotificationChangedEventArgs) Get_ChangeKind() UserNotificationChangedKind {
	var _result UserNotificationChangedKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ChangeKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUserNotificationChangedEventArgs) Get_UserNotificationId() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserNotificationId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type ToastActivatedEventArgs struct {
	RtClass
	*IToastActivatedEventArgs
}

type ToastCollection struct {
	RtClass
	*IToastCollection
}

func NewToastCollection_CreateInstance(collectionId string, displayName string, launchArgs string, iconUri *IUriRuntimeClass) *ToastCollection {
	hs := NewHStr("Windows.UI.Notifications.ToastCollection")
	var pFac *IToastCollectionFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IToastCollectionFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IToastCollection
	p = pFac.CreateInstance(collectionId, displayName, launchArgs, iconUri)
	result := &ToastCollection{
		RtClass:          RtClass{PInspect: &p.IInspectable},
		IToastCollection: p,
	}
	com.AddToScope(result)
	return result
}

type ToastCollectionManager struct {
	RtClass
	*IToastCollectionManager
}

type ToastDismissedEventArgs struct {
	RtClass
	*IToastDismissedEventArgs
}

type ToastFailedEventArgs struct {
	RtClass
	*IToastFailedEventArgs
}

type ToastNotification struct {
	RtClass
	*IToastNotification
}

func NewToastNotification_CreateToastNotification(content unsafe.Pointer) *ToastNotification {
	hs := NewHStr("Windows.UI.Notifications.ToastNotification")
	var pFac *IToastNotificationFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IToastNotificationFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IToastNotification
	p = pFac.CreateToastNotification(content)
	result := &ToastNotification{
		RtClass:            RtClass{PInspect: &p.IInspectable},
		IToastNotification: p,
	}
	com.AddToScope(result)
	return result
}

type ToastNotificationManager struct {
	RtClass
}

func NewIToastNotificationManagerStatics4() *IToastNotificationManagerStatics4 {
	var p *IToastNotificationManagerStatics4
	hs := NewHStr("Windows.UI.Notifications.ToastNotificationManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IToastNotificationManagerStatics4, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIToastNotificationManagerStatics5() *IToastNotificationManagerStatics5 {
	var p *IToastNotificationManagerStatics5
	hs := NewHStr("Windows.UI.Notifications.ToastNotificationManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IToastNotificationManagerStatics5, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIToastNotificationManagerStatics2() *IToastNotificationManagerStatics2 {
	var p *IToastNotificationManagerStatics2
	hs := NewHStr("Windows.UI.Notifications.ToastNotificationManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IToastNotificationManagerStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIToastNotificationManagerStatics() *IToastNotificationManagerStatics {
	var p *IToastNotificationManagerStatics
	hs := NewHStr("Windows.UI.Notifications.ToastNotificationManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IToastNotificationManagerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type ToastNotifier struct {
	RtClass
	*IToastNotifier
}
