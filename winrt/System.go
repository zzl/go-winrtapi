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
type AppDiagnosticInfoWatcherStatus int32

const (
	AppDiagnosticInfoWatcherStatus_Created              AppDiagnosticInfoWatcherStatus = 0
	AppDiagnosticInfoWatcherStatus_Started              AppDiagnosticInfoWatcherStatus = 1
	AppDiagnosticInfoWatcherStatus_EnumerationCompleted AppDiagnosticInfoWatcherStatus = 2
	AppDiagnosticInfoWatcherStatus_Stopping             AppDiagnosticInfoWatcherStatus = 3
	AppDiagnosticInfoWatcherStatus_Stopped              AppDiagnosticInfoWatcherStatus = 4
	AppDiagnosticInfoWatcherStatus_Aborted              AppDiagnosticInfoWatcherStatus = 5
)

// enum
type AppMemoryUsageLevel int32

const (
	AppMemoryUsageLevel_Low       AppMemoryUsageLevel = 0
	AppMemoryUsageLevel_Medium    AppMemoryUsageLevel = 1
	AppMemoryUsageLevel_High      AppMemoryUsageLevel = 2
	AppMemoryUsageLevel_OverLimit AppMemoryUsageLevel = 3
)

// enum
type AppResourceGroupEnergyQuotaState int32

const (
	AppResourceGroupEnergyQuotaState_Unknown AppResourceGroupEnergyQuotaState = 0
	AppResourceGroupEnergyQuotaState_Over    AppResourceGroupEnergyQuotaState = 1
	AppResourceGroupEnergyQuotaState_Under   AppResourceGroupEnergyQuotaState = 2
)

// enum
type AppResourceGroupExecutionState int32

const (
	AppResourceGroupExecutionState_Unknown    AppResourceGroupExecutionState = 0
	AppResourceGroupExecutionState_Running    AppResourceGroupExecutionState = 1
	AppResourceGroupExecutionState_Suspending AppResourceGroupExecutionState = 2
	AppResourceGroupExecutionState_Suspended  AppResourceGroupExecutionState = 3
	AppResourceGroupExecutionState_NotRunning AppResourceGroupExecutionState = 4
)

// enum
type AppResourceGroupInfoWatcherStatus int32

const (
	AppResourceGroupInfoWatcherStatus_Created              AppResourceGroupInfoWatcherStatus = 0
	AppResourceGroupInfoWatcherStatus_Started              AppResourceGroupInfoWatcherStatus = 1
	AppResourceGroupInfoWatcherStatus_EnumerationCompleted AppResourceGroupInfoWatcherStatus = 2
	AppResourceGroupInfoWatcherStatus_Stopping             AppResourceGroupInfoWatcherStatus = 3
	AppResourceGroupInfoWatcherStatus_Stopped              AppResourceGroupInfoWatcherStatus = 4
	AppResourceGroupInfoWatcherStatus_Aborted              AppResourceGroupInfoWatcherStatus = 5
)

// enum
type AutoUpdateTimeZoneStatus int32

const (
	AutoUpdateTimeZoneStatus_Attempted AutoUpdateTimeZoneStatus = 0
	AutoUpdateTimeZoneStatus_TimedOut  AutoUpdateTimeZoneStatus = 1
	AutoUpdateTimeZoneStatus_Failed    AutoUpdateTimeZoneStatus = 2
)

// enum
type DiagnosticAccessStatus int32

const (
	DiagnosticAccessStatus_Unspecified DiagnosticAccessStatus = 0
	DiagnosticAccessStatus_Denied      DiagnosticAccessStatus = 1
	DiagnosticAccessStatus_Limited     DiagnosticAccessStatus = 2
	DiagnosticAccessStatus_Allowed     DiagnosticAccessStatus = 3
)

// enum
type DispatcherQueuePriority int32

const (
	DispatcherQueuePriority_Low    DispatcherQueuePriority = -10
	DispatcherQueuePriority_Normal DispatcherQueuePriority = 0
	DispatcherQueuePriority_High   DispatcherQueuePriority = 10
)

// enum
type LaunchFileStatus int32

const (
	LaunchFileStatus_Success              LaunchFileStatus = 0
	LaunchFileStatus_AppUnavailable       LaunchFileStatus = 1
	LaunchFileStatus_DeniedByPolicy       LaunchFileStatus = 2
	LaunchFileStatus_FileTypeNotSupported LaunchFileStatus = 3
	LaunchFileStatus_Unknown              LaunchFileStatus = 4
)

// enum
type LaunchQuerySupportStatus int32

const (
	LaunchQuerySupportStatus_Available       LaunchQuerySupportStatus = 0
	LaunchQuerySupportStatus_AppNotInstalled LaunchQuerySupportStatus = 1
	LaunchQuerySupportStatus_AppUnavailable  LaunchQuerySupportStatus = 2
	LaunchQuerySupportStatus_NotSupported    LaunchQuerySupportStatus = 3
	LaunchQuerySupportStatus_Unknown         LaunchQuerySupportStatus = 4
)

// enum
type LaunchQuerySupportType int32

const (
	LaunchQuerySupportType_Uri           LaunchQuerySupportType = 0
	LaunchQuerySupportType_UriForResults LaunchQuerySupportType = 1
)

// enum
type LaunchUriStatus int32

const (
	LaunchUriStatus_Success             LaunchUriStatus = 0
	LaunchUriStatus_AppUnavailable      LaunchUriStatus = 1
	LaunchUriStatus_ProtocolUnavailable LaunchUriStatus = 2
	LaunchUriStatus_Unknown             LaunchUriStatus = 3
)

// enum
type PowerState int32

const (
	PowerState_ConnectedStandby PowerState = 0
	PowerState_SleepS3          PowerState = 1
)

// enum
type ProcessorArchitecture int32

const (
	ProcessorArchitecture_X86        ProcessorArchitecture = 0
	ProcessorArchitecture_Arm        ProcessorArchitecture = 5
	ProcessorArchitecture_X64        ProcessorArchitecture = 9
	ProcessorArchitecture_Neutral    ProcessorArchitecture = 11
	ProcessorArchitecture_Arm64      ProcessorArchitecture = 12
	ProcessorArchitecture_X86OnArm64 ProcessorArchitecture = 14
	ProcessorArchitecture_Unknown    ProcessorArchitecture = 65535
)

// enum
type RemoteLaunchUriStatus int32

const (
	RemoteLaunchUriStatus_Unknown                 RemoteLaunchUriStatus = 0
	RemoteLaunchUriStatus_Success                 RemoteLaunchUriStatus = 1
	RemoteLaunchUriStatus_AppUnavailable          RemoteLaunchUriStatus = 2
	RemoteLaunchUriStatus_ProtocolUnavailable     RemoteLaunchUriStatus = 3
	RemoteLaunchUriStatus_RemoteSystemUnavailable RemoteLaunchUriStatus = 4
	RemoteLaunchUriStatus_ValueSetTooLarge        RemoteLaunchUriStatus = 5
	RemoteLaunchUriStatus_DeniedByLocalSystem     RemoteLaunchUriStatus = 6
	RemoteLaunchUriStatus_DeniedByRemoteSystem    RemoteLaunchUriStatus = 7
)

// enum
type ShutdownKind int32

const (
	ShutdownKind_Shutdown ShutdownKind = 0
	ShutdownKind_Restart  ShutdownKind = 1
)

// enum
type UserAuthenticationStatus int32

const (
	UserAuthenticationStatus_Unauthenticated       UserAuthenticationStatus = 0
	UserAuthenticationStatus_LocallyAuthenticated  UserAuthenticationStatus = 1
	UserAuthenticationStatus_RemotelyAuthenticated UserAuthenticationStatus = 2
)

// enum
type UserPictureSize int32

const (
	UserPictureSize_Size64x64     UserPictureSize = 0
	UserPictureSize_Size208x208   UserPictureSize = 1
	UserPictureSize_Size424x424   UserPictureSize = 2
	UserPictureSize_Size1080x1080 UserPictureSize = 3
)

// enum
type UserType int32

const (
	UserType_LocalUser     UserType = 0
	UserType_RemoteUser    UserType = 1
	UserType_LocalGuest    UserType = 2
	UserType_RemoteGuest   UserType = 3
	UserType_SystemManaged UserType = 4
)

// enum
type UserWatcherStatus int32

const (
	UserWatcherStatus_Created              UserWatcherStatus = 0
	UserWatcherStatus_Started              UserWatcherStatus = 1
	UserWatcherStatus_EnumerationCompleted UserWatcherStatus = 2
	UserWatcherStatus_Stopping             UserWatcherStatus = 3
	UserWatcherStatus_Stopped              UserWatcherStatus = 4
	UserWatcherStatus_Aborted              UserWatcherStatus = 5
)

// enum
type UserWatcherUpdateKind int32

const (
	UserWatcherUpdateKind_Properties UserWatcherUpdateKind = 0
	UserWatcherUpdateKind_Picture    UserWatcherUpdateKind = 1
)

// enum
type VirtualKey int32

const (
	VirtualKey_None                         VirtualKey = 0
	VirtualKey_LeftButton                   VirtualKey = 1
	VirtualKey_RightButton                  VirtualKey = 2
	VirtualKey_Cancel                       VirtualKey = 3
	VirtualKey_MiddleButton                 VirtualKey = 4
	VirtualKey_XButton1                     VirtualKey = 5
	VirtualKey_XButton2                     VirtualKey = 6
	VirtualKey_Back                         VirtualKey = 8
	VirtualKey_Tab                          VirtualKey = 9
	VirtualKey_Clear                        VirtualKey = 12
	VirtualKey_Enter                        VirtualKey = 13
	VirtualKey_Shift                        VirtualKey = 16
	VirtualKey_Control                      VirtualKey = 17
	VirtualKey_Menu                         VirtualKey = 18
	VirtualKey_Pause                        VirtualKey = 19
	VirtualKey_CapitalLock                  VirtualKey = 20
	VirtualKey_Kana                         VirtualKey = 21
	VirtualKey_Hangul                       VirtualKey = 21
	VirtualKey_ImeOn                        VirtualKey = 22
	VirtualKey_Junja                        VirtualKey = 23
	VirtualKey_Final                        VirtualKey = 24
	VirtualKey_Hanja                        VirtualKey = 25
	VirtualKey_Kanji                        VirtualKey = 25
	VirtualKey_ImeOff                       VirtualKey = 26
	VirtualKey_Escape                       VirtualKey = 27
	VirtualKey_Convert                      VirtualKey = 28
	VirtualKey_NonConvert                   VirtualKey = 29
	VirtualKey_Accept                       VirtualKey = 30
	VirtualKey_ModeChange                   VirtualKey = 31
	VirtualKey_Space                        VirtualKey = 32
	VirtualKey_PageUp                       VirtualKey = 33
	VirtualKey_PageDown                     VirtualKey = 34
	VirtualKey_End                          VirtualKey = 35
	VirtualKey_Home                         VirtualKey = 36
	VirtualKey_Left                         VirtualKey = 37
	VirtualKey_Up                           VirtualKey = 38
	VirtualKey_Right                        VirtualKey = 39
	VirtualKey_Down                         VirtualKey = 40
	VirtualKey_Select                       VirtualKey = 41
	VirtualKey_Print                        VirtualKey = 42
	VirtualKey_Execute                      VirtualKey = 43
	VirtualKey_Snapshot                     VirtualKey = 44
	VirtualKey_Insert                       VirtualKey = 45
	VirtualKey_Delete                       VirtualKey = 46
	VirtualKey_Help                         VirtualKey = 47
	VirtualKey_Number0                      VirtualKey = 48
	VirtualKey_Number1                      VirtualKey = 49
	VirtualKey_Number2                      VirtualKey = 50
	VirtualKey_Number3                      VirtualKey = 51
	VirtualKey_Number4                      VirtualKey = 52
	VirtualKey_Number5                      VirtualKey = 53
	VirtualKey_Number6                      VirtualKey = 54
	VirtualKey_Number7                      VirtualKey = 55
	VirtualKey_Number8                      VirtualKey = 56
	VirtualKey_Number9                      VirtualKey = 57
	VirtualKey_A                            VirtualKey = 65
	VirtualKey_B                            VirtualKey = 66
	VirtualKey_C                            VirtualKey = 67
	VirtualKey_D                            VirtualKey = 68
	VirtualKey_E                            VirtualKey = 69
	VirtualKey_F                            VirtualKey = 70
	VirtualKey_G                            VirtualKey = 71
	VirtualKey_H                            VirtualKey = 72
	VirtualKey_I                            VirtualKey = 73
	VirtualKey_J                            VirtualKey = 74
	VirtualKey_K                            VirtualKey = 75
	VirtualKey_L                            VirtualKey = 76
	VirtualKey_M                            VirtualKey = 77
	VirtualKey_N                            VirtualKey = 78
	VirtualKey_O                            VirtualKey = 79
	VirtualKey_P                            VirtualKey = 80
	VirtualKey_Q                            VirtualKey = 81
	VirtualKey_R                            VirtualKey = 82
	VirtualKey_S                            VirtualKey = 83
	VirtualKey_T                            VirtualKey = 84
	VirtualKey_U                            VirtualKey = 85
	VirtualKey_V                            VirtualKey = 86
	VirtualKey_W                            VirtualKey = 87
	VirtualKey_X                            VirtualKey = 88
	VirtualKey_Y                            VirtualKey = 89
	VirtualKey_Z                            VirtualKey = 90
	VirtualKey_LeftWindows                  VirtualKey = 91
	VirtualKey_RightWindows                 VirtualKey = 92
	VirtualKey_Application                  VirtualKey = 93
	VirtualKey_Sleep                        VirtualKey = 95
	VirtualKey_NumberPad0                   VirtualKey = 96
	VirtualKey_NumberPad1                   VirtualKey = 97
	VirtualKey_NumberPad2                   VirtualKey = 98
	VirtualKey_NumberPad3                   VirtualKey = 99
	VirtualKey_NumberPad4                   VirtualKey = 100
	VirtualKey_NumberPad5                   VirtualKey = 101
	VirtualKey_NumberPad6                   VirtualKey = 102
	VirtualKey_NumberPad7                   VirtualKey = 103
	VirtualKey_NumberPad8                   VirtualKey = 104
	VirtualKey_NumberPad9                   VirtualKey = 105
	VirtualKey_Multiply                     VirtualKey = 106
	VirtualKey_Add                          VirtualKey = 107
	VirtualKey_Separator                    VirtualKey = 108
	VirtualKey_Subtract                     VirtualKey = 109
	VirtualKey_Decimal                      VirtualKey = 110
	VirtualKey_Divide                       VirtualKey = 111
	VirtualKey_F1                           VirtualKey = 112
	VirtualKey_F2                           VirtualKey = 113
	VirtualKey_F3                           VirtualKey = 114
	VirtualKey_F4                           VirtualKey = 115
	VirtualKey_F5                           VirtualKey = 116
	VirtualKey_F6                           VirtualKey = 117
	VirtualKey_F7                           VirtualKey = 118
	VirtualKey_F8                           VirtualKey = 119
	VirtualKey_F9                           VirtualKey = 120
	VirtualKey_F10                          VirtualKey = 121
	VirtualKey_F11                          VirtualKey = 122
	VirtualKey_F12                          VirtualKey = 123
	VirtualKey_F13                          VirtualKey = 124
	VirtualKey_F14                          VirtualKey = 125
	VirtualKey_F15                          VirtualKey = 126
	VirtualKey_F16                          VirtualKey = 127
	VirtualKey_F17                          VirtualKey = 128
	VirtualKey_F18                          VirtualKey = 129
	VirtualKey_F19                          VirtualKey = 130
	VirtualKey_F20                          VirtualKey = 131
	VirtualKey_F21                          VirtualKey = 132
	VirtualKey_F22                          VirtualKey = 133
	VirtualKey_F23                          VirtualKey = 134
	VirtualKey_F24                          VirtualKey = 135
	VirtualKey_NavigationView               VirtualKey = 136
	VirtualKey_NavigationMenu               VirtualKey = 137
	VirtualKey_NavigationUp                 VirtualKey = 138
	VirtualKey_NavigationDown               VirtualKey = 139
	VirtualKey_NavigationLeft               VirtualKey = 140
	VirtualKey_NavigationRight              VirtualKey = 141
	VirtualKey_NavigationAccept             VirtualKey = 142
	VirtualKey_NavigationCancel             VirtualKey = 143
	VirtualKey_NumberKeyLock                VirtualKey = 144
	VirtualKey_Scroll                       VirtualKey = 145
	VirtualKey_LeftShift                    VirtualKey = 160
	VirtualKey_RightShift                   VirtualKey = 161
	VirtualKey_LeftControl                  VirtualKey = 162
	VirtualKey_RightControl                 VirtualKey = 163
	VirtualKey_LeftMenu                     VirtualKey = 164
	VirtualKey_RightMenu                    VirtualKey = 165
	VirtualKey_GoBack                       VirtualKey = 166
	VirtualKey_GoForward                    VirtualKey = 167
	VirtualKey_Refresh                      VirtualKey = 168
	VirtualKey_Stop                         VirtualKey = 169
	VirtualKey_Search                       VirtualKey = 170
	VirtualKey_Favorites                    VirtualKey = 171
	VirtualKey_GoHome                       VirtualKey = 172
	VirtualKey_GamepadA                     VirtualKey = 195
	VirtualKey_GamepadB                     VirtualKey = 196
	VirtualKey_GamepadX                     VirtualKey = 197
	VirtualKey_GamepadY                     VirtualKey = 198
	VirtualKey_GamepadRightShoulder         VirtualKey = 199
	VirtualKey_GamepadLeftShoulder          VirtualKey = 200
	VirtualKey_GamepadLeftTrigger           VirtualKey = 201
	VirtualKey_GamepadRightTrigger          VirtualKey = 202
	VirtualKey_GamepadDPadUp                VirtualKey = 203
	VirtualKey_GamepadDPadDown              VirtualKey = 204
	VirtualKey_GamepadDPadLeft              VirtualKey = 205
	VirtualKey_GamepadDPadRight             VirtualKey = 206
	VirtualKey_GamepadMenu                  VirtualKey = 207
	VirtualKey_GamepadView                  VirtualKey = 208
	VirtualKey_GamepadLeftThumbstickButton  VirtualKey = 209
	VirtualKey_GamepadRightThumbstickButton VirtualKey = 210
	VirtualKey_GamepadLeftThumbstickUp      VirtualKey = 211
	VirtualKey_GamepadLeftThumbstickDown    VirtualKey = 212
	VirtualKey_GamepadLeftThumbstickRight   VirtualKey = 213
	VirtualKey_GamepadLeftThumbstickLeft    VirtualKey = 214
	VirtualKey_GamepadRightThumbstickUp     VirtualKey = 215
	VirtualKey_GamepadRightThumbstickDown   VirtualKey = 216
	VirtualKey_GamepadRightThumbstickRight  VirtualKey = 217
	VirtualKey_GamepadRightThumbstickLeft   VirtualKey = 218
)

// enum
// flags
type VirtualKeyModifiers uint32

const (
	VirtualKeyModifiers_None    VirtualKeyModifiers = 0
	VirtualKeyModifiers_Control VirtualKeyModifiers = 1
	VirtualKeyModifiers_Menu    VirtualKeyModifiers = 2
	VirtualKeyModifiers_Shift   VirtualKeyModifiers = 4
	VirtualKeyModifiers_Windows VirtualKeyModifiers = 8
)

// structs

type SystemManagementContract struct {
}

// func types

// DFA2DC9C-1A2D-4917-98F2-939AF1D6E0C8
type DispatcherQueueHandler func() com.Error

// interfaces

// 6B528900-F46E-4EB0-AA6C-38AF557CF9ED
var IID_IAppActivationResult = syscall.GUID{0x6B528900, 0xF46E, 0x4EB0,
	[8]byte{0xAA, 0x6C, 0x38, 0xAF, 0x55, 0x7C, 0xF9, 0xED}}

type IAppActivationResultInterface interface {
	win32.IInspectableInterface
	Get_ExtendedError() HResult
	Get_AppResourceGroupInfo() *IAppResourceGroupInfo
}

type IAppActivationResultVtbl struct {
	win32.IInspectableVtbl
	Get_ExtendedError        uintptr
	Get_AppResourceGroupInfo uintptr
}

type IAppActivationResult struct {
	win32.IInspectable
}

func (this *IAppActivationResult) Vtbl() *IAppActivationResultVtbl {
	return (*IAppActivationResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppActivationResult) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppActivationResult) Get_AppResourceGroupInfo() *IAppResourceGroupInfo {
	var _result *IAppResourceGroupInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppResourceGroupInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E348A69A-8889-4CA3-BE07-D5FFFF5F0804
var IID_IAppDiagnosticInfo = syscall.GUID{0xE348A69A, 0x8889, 0x4CA3,
	[8]byte{0xBE, 0x07, 0xD5, 0xFF, 0xFF, 0x5F, 0x08, 0x04}}

type IAppDiagnosticInfoInterface interface {
	win32.IInspectableInterface
	Get_AppInfo() *IAppInfo
}

type IAppDiagnosticInfoVtbl struct {
	win32.IInspectableVtbl
	Get_AppInfo uintptr
}

type IAppDiagnosticInfo struct {
	win32.IInspectable
}

func (this *IAppDiagnosticInfo) Vtbl() *IAppDiagnosticInfoVtbl {
	return (*IAppDiagnosticInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppDiagnosticInfo) Get_AppInfo() *IAppInfo {
	var _result *IAppInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DF46FBD7-191A-446C-9473-8FBC2374A354
var IID_IAppDiagnosticInfo2 = syscall.GUID{0xDF46FBD7, 0x191A, 0x446C,
	[8]byte{0x94, 0x73, 0x8F, 0xBC, 0x23, 0x74, 0xA3, 0x54}}

type IAppDiagnosticInfo2Interface interface {
	win32.IInspectableInterface
	GetResourceGroups() *IVector[*IAppResourceGroupInfo]
	CreateResourceGroupWatcher() *IAppResourceGroupInfoWatcher
}

type IAppDiagnosticInfo2Vtbl struct {
	win32.IInspectableVtbl
	GetResourceGroups          uintptr
	CreateResourceGroupWatcher uintptr
}

type IAppDiagnosticInfo2 struct {
	win32.IInspectable
}

func (this *IAppDiagnosticInfo2) Vtbl() *IAppDiagnosticInfo2Vtbl {
	return (*IAppDiagnosticInfo2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppDiagnosticInfo2) GetResourceGroups() *IVector[*IAppResourceGroupInfo] {
	var _result *IVector[*IAppResourceGroupInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetResourceGroups, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppDiagnosticInfo2) CreateResourceGroupWatcher() *IAppResourceGroupInfoWatcher {
	var _result *IAppResourceGroupInfoWatcher
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateResourceGroupWatcher, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C895C63D-DD61-4C65-BABD-81A10B4F9815
var IID_IAppDiagnosticInfo3 = syscall.GUID{0xC895C63D, 0xDD61, 0x4C65,
	[8]byte{0xBA, 0xBD, 0x81, 0xA1, 0x0B, 0x4F, 0x98, 0x15}}

type IAppDiagnosticInfo3Interface interface {
	win32.IInspectableInterface
	LaunchAsync() *IAsyncOperation[*IAppActivationResult]
}

type IAppDiagnosticInfo3Vtbl struct {
	win32.IInspectableVtbl
	LaunchAsync uintptr
}

type IAppDiagnosticInfo3 struct {
	win32.IInspectable
}

func (this *IAppDiagnosticInfo3) Vtbl() *IAppDiagnosticInfo3Vtbl {
	return (*IAppDiagnosticInfo3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppDiagnosticInfo3) LaunchAsync() *IAsyncOperation[*IAppActivationResult] {
	var _result *IAsyncOperation[*IAppActivationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CE6925BF-10CA-40C8-A9CA-C5C96501866E
var IID_IAppDiagnosticInfoStatics = syscall.GUID{0xCE6925BF, 0x10CA, 0x40C8,
	[8]byte{0xA9, 0xCA, 0xC5, 0xC9, 0x65, 0x01, 0x86, 0x6E}}

type IAppDiagnosticInfoStaticsInterface interface {
	win32.IInspectableInterface
	RequestInfoAsync() *IAsyncOperation[*IVector[*IAppDiagnosticInfo]]
}

type IAppDiagnosticInfoStaticsVtbl struct {
	win32.IInspectableVtbl
	RequestInfoAsync uintptr
}

type IAppDiagnosticInfoStatics struct {
	win32.IInspectable
}

func (this *IAppDiagnosticInfoStatics) Vtbl() *IAppDiagnosticInfoStaticsVtbl {
	return (*IAppDiagnosticInfoStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppDiagnosticInfoStatics) RequestInfoAsync() *IAsyncOperation[*IVector[*IAppDiagnosticInfo]] {
	var _result *IAsyncOperation[*IVector[*IAppDiagnosticInfo]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestInfoAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 05B24B86-1000-4C90-BB9F-7235071C50FE
var IID_IAppDiagnosticInfoStatics2 = syscall.GUID{0x05B24B86, 0x1000, 0x4C90,
	[8]byte{0xBB, 0x9F, 0x72, 0x35, 0x07, 0x1C, 0x50, 0xFE}}

type IAppDiagnosticInfoStatics2Interface interface {
	win32.IInspectableInterface
	CreateWatcher() *IAppDiagnosticInfoWatcher
	RequestAccessAsync() *IAsyncOperation[DiagnosticAccessStatus]
	RequestInfoForPackageAsync(packageFamilyName string) *IAsyncOperation[*IVector[*IAppDiagnosticInfo]]
	RequestInfoForAppAsync() *IAsyncOperation[*IVector[*IAppDiagnosticInfo]]
	RequestInfoForAppUserModelId(appUserModelId string) *IAsyncOperation[*IVector[*IAppDiagnosticInfo]]
}

type IAppDiagnosticInfoStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateWatcher                uintptr
	RequestAccessAsync           uintptr
	RequestInfoForPackageAsync   uintptr
	RequestInfoForAppAsync       uintptr
	RequestInfoForAppUserModelId uintptr
}

type IAppDiagnosticInfoStatics2 struct {
	win32.IInspectable
}

func (this *IAppDiagnosticInfoStatics2) Vtbl() *IAppDiagnosticInfoStatics2Vtbl {
	return (*IAppDiagnosticInfoStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppDiagnosticInfoStatics2) CreateWatcher() *IAppDiagnosticInfoWatcher {
	var _result *IAppDiagnosticInfoWatcher
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWatcher, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppDiagnosticInfoStatics2) RequestAccessAsync() *IAsyncOperation[DiagnosticAccessStatus] {
	var _result *IAsyncOperation[DiagnosticAccessStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppDiagnosticInfoStatics2) RequestInfoForPackageAsync(packageFamilyName string) *IAsyncOperation[*IVector[*IAppDiagnosticInfo]] {
	var _result *IAsyncOperation[*IVector[*IAppDiagnosticInfo]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestInfoForPackageAsync, uintptr(unsafe.Pointer(this)), NewHStr(packageFamilyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppDiagnosticInfoStatics2) RequestInfoForAppAsync() *IAsyncOperation[*IVector[*IAppDiagnosticInfo]] {
	var _result *IAsyncOperation[*IVector[*IAppDiagnosticInfo]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestInfoForAppAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppDiagnosticInfoStatics2) RequestInfoForAppUserModelId(appUserModelId string) *IAsyncOperation[*IVector[*IAppDiagnosticInfo]] {
	var _result *IAsyncOperation[*IVector[*IAppDiagnosticInfo]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestInfoForAppUserModelId, uintptr(unsafe.Pointer(this)), NewHStr(appUserModelId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 75575070-01D3-489A-9325-52F9CC6EDE0A
var IID_IAppDiagnosticInfoWatcher = syscall.GUID{0x75575070, 0x01D3, 0x489A,
	[8]byte{0x93, 0x25, 0x52, 0xF9, 0xCC, 0x6E, 0xDE, 0x0A}}

type IAppDiagnosticInfoWatcherInterface interface {
	win32.IInspectableInterface
	Add_Added(handler TypedEventHandler[*IAppDiagnosticInfoWatcher, *IAppDiagnosticInfoWatcherEventArgs]) EventRegistrationToken
	Remove_Added(token EventRegistrationToken)
	Add_Removed(handler TypedEventHandler[*IAppDiagnosticInfoWatcher, *IAppDiagnosticInfoWatcherEventArgs]) EventRegistrationToken
	Remove_Removed(token EventRegistrationToken)
	Add_EnumerationCompleted(handler TypedEventHandler[*IAppDiagnosticInfoWatcher, interface{}]) EventRegistrationToken
	Remove_EnumerationCompleted(token EventRegistrationToken)
	Add_Stopped(handler TypedEventHandler[*IAppDiagnosticInfoWatcher, interface{}]) EventRegistrationToken
	Remove_Stopped(token EventRegistrationToken)
	Get_Status() AppDiagnosticInfoWatcherStatus
	Start()
	Stop()
}

type IAppDiagnosticInfoWatcherVtbl struct {
	win32.IInspectableVtbl
	Add_Added                   uintptr
	Remove_Added                uintptr
	Add_Removed                 uintptr
	Remove_Removed              uintptr
	Add_EnumerationCompleted    uintptr
	Remove_EnumerationCompleted uintptr
	Add_Stopped                 uintptr
	Remove_Stopped              uintptr
	Get_Status                  uintptr
	Start                       uintptr
	Stop                        uintptr
}

type IAppDiagnosticInfoWatcher struct {
	win32.IInspectable
}

func (this *IAppDiagnosticInfoWatcher) Vtbl() *IAppDiagnosticInfoWatcherVtbl {
	return (*IAppDiagnosticInfoWatcherVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppDiagnosticInfoWatcher) Add_Added(handler TypedEventHandler[*IAppDiagnosticInfoWatcher, *IAppDiagnosticInfoWatcherEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Added, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppDiagnosticInfoWatcher) Remove_Added(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Added, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppDiagnosticInfoWatcher) Add_Removed(handler TypedEventHandler[*IAppDiagnosticInfoWatcher, *IAppDiagnosticInfoWatcherEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Removed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppDiagnosticInfoWatcher) Remove_Removed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Removed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppDiagnosticInfoWatcher) Add_EnumerationCompleted(handler TypedEventHandler[*IAppDiagnosticInfoWatcher, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_EnumerationCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppDiagnosticInfoWatcher) Remove_EnumerationCompleted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_EnumerationCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppDiagnosticInfoWatcher) Add_Stopped(handler TypedEventHandler[*IAppDiagnosticInfoWatcher, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Stopped, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppDiagnosticInfoWatcher) Remove_Stopped(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Stopped, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppDiagnosticInfoWatcher) Get_Status() AppDiagnosticInfoWatcherStatus {
	var _result AppDiagnosticInfoWatcherStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppDiagnosticInfoWatcher) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAppDiagnosticInfoWatcher) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 7017C716-E1DA-4C65-99DF-046DFF5BE71A
var IID_IAppDiagnosticInfoWatcherEventArgs = syscall.GUID{0x7017C716, 0xE1DA, 0x4C65,
	[8]byte{0x99, 0xDF, 0x04, 0x6D, 0xFF, 0x5B, 0xE7, 0x1A}}

type IAppDiagnosticInfoWatcherEventArgsInterface interface {
	win32.IInspectableInterface
	Get_AppDiagnosticInfo() *IAppDiagnosticInfo
}

type IAppDiagnosticInfoWatcherEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_AppDiagnosticInfo uintptr
}

type IAppDiagnosticInfoWatcherEventArgs struct {
	win32.IInspectable
}

func (this *IAppDiagnosticInfoWatcherEventArgs) Vtbl() *IAppDiagnosticInfoWatcherEventArgsVtbl {
	return (*IAppDiagnosticInfoWatcherEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppDiagnosticInfoWatcherEventArgs) Get_AppDiagnosticInfo() *IAppDiagnosticInfo {
	var _result *IAppDiagnosticInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppDiagnosticInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6F039BF0-F91B-4DF8-AE77-3033CCB69114
var IID_IAppExecutionStateChangeResult = syscall.GUID{0x6F039BF0, 0xF91B, 0x4DF8,
	[8]byte{0xAE, 0x77, 0x30, 0x33, 0xCC, 0xB6, 0x91, 0x14}}

type IAppExecutionStateChangeResultInterface interface {
	win32.IInspectableInterface
	Get_ExtendedError() HResult
}

type IAppExecutionStateChangeResultVtbl struct {
	win32.IInspectableVtbl
	Get_ExtendedError uintptr
}

type IAppExecutionStateChangeResult struct {
	win32.IInspectable
}

func (this *IAppExecutionStateChangeResult) Vtbl() *IAppExecutionStateChangeResultVtbl {
	return (*IAppExecutionStateChangeResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppExecutionStateChangeResult) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6D65339B-4D6F-45BC-9C5E-E49B3FF2758D
var IID_IAppMemoryReport = syscall.GUID{0x6D65339B, 0x4D6F, 0x45BC,
	[8]byte{0x9C, 0x5E, 0xE4, 0x9B, 0x3F, 0xF2, 0x75, 0x8D}}

type IAppMemoryReportInterface interface {
	win32.IInspectableInterface
	Get_PrivateCommitUsage() uint64
	Get_PeakPrivateCommitUsage() uint64
	Get_TotalCommitUsage() uint64
	Get_TotalCommitLimit() uint64
}

type IAppMemoryReportVtbl struct {
	win32.IInspectableVtbl
	Get_PrivateCommitUsage     uintptr
	Get_PeakPrivateCommitUsage uintptr
	Get_TotalCommitUsage       uintptr
	Get_TotalCommitLimit       uintptr
}

type IAppMemoryReport struct {
	win32.IInspectable
}

func (this *IAppMemoryReport) Vtbl() *IAppMemoryReportVtbl {
	return (*IAppMemoryReportVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppMemoryReport) Get_PrivateCommitUsage() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrivateCommitUsage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppMemoryReport) Get_PeakPrivateCommitUsage() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PeakPrivateCommitUsage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppMemoryReport) Get_TotalCommitUsage() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TotalCommitUsage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppMemoryReport) Get_TotalCommitLimit() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TotalCommitLimit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5F7F3738-51B7-42DC-B7ED-79BA46D28857
var IID_IAppMemoryReport2 = syscall.GUID{0x5F7F3738, 0x51B7, 0x42DC,
	[8]byte{0xB7, 0xED, 0x79, 0xBA, 0x46, 0xD2, 0x88, 0x57}}

type IAppMemoryReport2Interface interface {
	win32.IInspectableInterface
	Get_ExpectedTotalCommitLimit() uint64
}

type IAppMemoryReport2Vtbl struct {
	win32.IInspectableVtbl
	Get_ExpectedTotalCommitLimit uintptr
}

type IAppMemoryReport2 struct {
	win32.IInspectable
}

func (this *IAppMemoryReport2) Vtbl() *IAppMemoryReport2Vtbl {
	return (*IAppMemoryReport2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppMemoryReport2) Get_ExpectedTotalCommitLimit() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExpectedTotalCommitLimit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 79F86664-FECA-4DA5-9E40-2BC63EFDC979
var IID_IAppMemoryUsageLimitChangingEventArgs = syscall.GUID{0x79F86664, 0xFECA, 0x4DA5,
	[8]byte{0x9E, 0x40, 0x2B, 0xC6, 0x3E, 0xFD, 0xC9, 0x79}}

type IAppMemoryUsageLimitChangingEventArgsInterface interface {
	win32.IInspectableInterface
	Get_OldLimit() uint64
	Get_NewLimit() uint64
}

type IAppMemoryUsageLimitChangingEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_OldLimit uintptr
	Get_NewLimit uintptr
}

type IAppMemoryUsageLimitChangingEventArgs struct {
	win32.IInspectable
}

func (this *IAppMemoryUsageLimitChangingEventArgs) Vtbl() *IAppMemoryUsageLimitChangingEventArgsVtbl {
	return (*IAppMemoryUsageLimitChangingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppMemoryUsageLimitChangingEventArgs) Get_OldLimit() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OldLimit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppMemoryUsageLimitChangingEventArgs) Get_NewLimit() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NewLimit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 2566E74E-B05D-40C2-9DC1-1A4F039EA120
var IID_IAppResourceGroupBackgroundTaskReport = syscall.GUID{0x2566E74E, 0xB05D, 0x40C2,
	[8]byte{0x9D, 0xC1, 0x1A, 0x4F, 0x03, 0x9E, 0xA1, 0x20}}

type IAppResourceGroupBackgroundTaskReportInterface interface {
	win32.IInspectableInterface
	Get_TaskId() syscall.GUID
	Get_Name() string
	Get_Trigger() string
	Get_EntryPoint() string
}

type IAppResourceGroupBackgroundTaskReportVtbl struct {
	win32.IInspectableVtbl
	Get_TaskId     uintptr
	Get_Name       uintptr
	Get_Trigger    uintptr
	Get_EntryPoint uintptr
}

type IAppResourceGroupBackgroundTaskReport struct {
	win32.IInspectable
}

func (this *IAppResourceGroupBackgroundTaskReport) Vtbl() *IAppResourceGroupBackgroundTaskReportVtbl {
	return (*IAppResourceGroupBackgroundTaskReportVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppResourceGroupBackgroundTaskReport) Get_TaskId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TaskId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppResourceGroupBackgroundTaskReport) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppResourceGroupBackgroundTaskReport) Get_Trigger() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Trigger, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppResourceGroupBackgroundTaskReport) Get_EntryPoint() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EntryPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// B913F77A-E807-49F4-845E-7B8BDCFE8EE7
var IID_IAppResourceGroupInfo = syscall.GUID{0xB913F77A, 0xE807, 0x49F4,
	[8]byte{0x84, 0x5E, 0x7B, 0x8B, 0xDC, 0xFE, 0x8E, 0xE7}}

type IAppResourceGroupInfoInterface interface {
	win32.IInspectableInterface
	Get_InstanceId() syscall.GUID
	Get_IsShared() bool
	GetBackgroundTaskReports() *IVector[*IAppResourceGroupBackgroundTaskReport]
	GetMemoryReport() *IAppResourceGroupMemoryReport
	GetProcessDiagnosticInfos() *IVector[*IProcessDiagnosticInfo]
	GetStateReport() *IAppResourceGroupStateReport
}

type IAppResourceGroupInfoVtbl struct {
	win32.IInspectableVtbl
	Get_InstanceId            uintptr
	Get_IsShared              uintptr
	GetBackgroundTaskReports  uintptr
	GetMemoryReport           uintptr
	GetProcessDiagnosticInfos uintptr
	GetStateReport            uintptr
}

type IAppResourceGroupInfo struct {
	win32.IInspectable
}

func (this *IAppResourceGroupInfo) Vtbl() *IAppResourceGroupInfoVtbl {
	return (*IAppResourceGroupInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppResourceGroupInfo) Get_InstanceId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InstanceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppResourceGroupInfo) Get_IsShared() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsShared, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppResourceGroupInfo) GetBackgroundTaskReports() *IVector[*IAppResourceGroupBackgroundTaskReport] {
	var _result *IVector[*IAppResourceGroupBackgroundTaskReport]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetBackgroundTaskReports, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppResourceGroupInfo) GetMemoryReport() *IAppResourceGroupMemoryReport {
	var _result *IAppResourceGroupMemoryReport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetMemoryReport, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppResourceGroupInfo) GetProcessDiagnosticInfos() *IVector[*IProcessDiagnosticInfo] {
	var _result *IVector[*IProcessDiagnosticInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetProcessDiagnosticInfos, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppResourceGroupInfo) GetStateReport() *IAppResourceGroupStateReport {
	var _result *IAppResourceGroupStateReport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStateReport, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EE9B236D-D305-4D6B-92F7-6AFDAD72DEDC
var IID_IAppResourceGroupInfo2 = syscall.GUID{0xEE9B236D, 0xD305, 0x4D6B,
	[8]byte{0x92, 0xF7, 0x6A, 0xFD, 0xAD, 0x72, 0xDE, 0xDC}}

type IAppResourceGroupInfo2Interface interface {
	win32.IInspectableInterface
	StartSuspendAsync() *IAsyncOperation[*IAppExecutionStateChangeResult]
	StartResumeAsync() *IAsyncOperation[*IAppExecutionStateChangeResult]
	StartTerminateAsync() *IAsyncOperation[*IAppExecutionStateChangeResult]
}

type IAppResourceGroupInfo2Vtbl struct {
	win32.IInspectableVtbl
	StartSuspendAsync   uintptr
	StartResumeAsync    uintptr
	StartTerminateAsync uintptr
}

type IAppResourceGroupInfo2 struct {
	win32.IInspectable
}

func (this *IAppResourceGroupInfo2) Vtbl() *IAppResourceGroupInfo2Vtbl {
	return (*IAppResourceGroupInfo2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppResourceGroupInfo2) StartSuspendAsync() *IAsyncOperation[*IAppExecutionStateChangeResult] {
	var _result *IAsyncOperation[*IAppExecutionStateChangeResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartSuspendAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppResourceGroupInfo2) StartResumeAsync() *IAsyncOperation[*IAppExecutionStateChangeResult] {
	var _result *IAsyncOperation[*IAppExecutionStateChangeResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartResumeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppResourceGroupInfo2) StartTerminateAsync() *IAsyncOperation[*IAppExecutionStateChangeResult] {
	var _result *IAsyncOperation[*IAppExecutionStateChangeResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartTerminateAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D9B0A0FD-6E5A-4C72-8B17-09FEC4A212BD
var IID_IAppResourceGroupInfoWatcher = syscall.GUID{0xD9B0A0FD, 0x6E5A, 0x4C72,
	[8]byte{0x8B, 0x17, 0x09, 0xFE, 0xC4, 0xA2, 0x12, 0xBD}}

type IAppResourceGroupInfoWatcherInterface interface {
	win32.IInspectableInterface
	Add_Added(handler TypedEventHandler[*IAppResourceGroupInfoWatcher, *IAppResourceGroupInfoWatcherEventArgs]) EventRegistrationToken
	Remove_Added(token EventRegistrationToken)
	Add_Removed(handler TypedEventHandler[*IAppResourceGroupInfoWatcher, *IAppResourceGroupInfoWatcherEventArgs]) EventRegistrationToken
	Remove_Removed(token EventRegistrationToken)
	Add_EnumerationCompleted(handler TypedEventHandler[*IAppResourceGroupInfoWatcher, interface{}]) EventRegistrationToken
	Remove_EnumerationCompleted(token EventRegistrationToken)
	Add_Stopped(handler TypedEventHandler[*IAppResourceGroupInfoWatcher, interface{}]) EventRegistrationToken
	Remove_Stopped(token EventRegistrationToken)
	Add_ExecutionStateChanged(handler TypedEventHandler[*IAppResourceGroupInfoWatcher, *IAppResourceGroupInfoWatcherExecutionStateChangedEventArgs]) EventRegistrationToken
	Remove_ExecutionStateChanged(token EventRegistrationToken)
	Get_Status() AppResourceGroupInfoWatcherStatus
	Start()
	Stop()
}

type IAppResourceGroupInfoWatcherVtbl struct {
	win32.IInspectableVtbl
	Add_Added                    uintptr
	Remove_Added                 uintptr
	Add_Removed                  uintptr
	Remove_Removed               uintptr
	Add_EnumerationCompleted     uintptr
	Remove_EnumerationCompleted  uintptr
	Add_Stopped                  uintptr
	Remove_Stopped               uintptr
	Add_ExecutionStateChanged    uintptr
	Remove_ExecutionStateChanged uintptr
	Get_Status                   uintptr
	Start                        uintptr
	Stop                         uintptr
}

type IAppResourceGroupInfoWatcher struct {
	win32.IInspectable
}

func (this *IAppResourceGroupInfoWatcher) Vtbl() *IAppResourceGroupInfoWatcherVtbl {
	return (*IAppResourceGroupInfoWatcherVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppResourceGroupInfoWatcher) Add_Added(handler TypedEventHandler[*IAppResourceGroupInfoWatcher, *IAppResourceGroupInfoWatcherEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Added, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppResourceGroupInfoWatcher) Remove_Added(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Added, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppResourceGroupInfoWatcher) Add_Removed(handler TypedEventHandler[*IAppResourceGroupInfoWatcher, *IAppResourceGroupInfoWatcherEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Removed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppResourceGroupInfoWatcher) Remove_Removed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Removed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppResourceGroupInfoWatcher) Add_EnumerationCompleted(handler TypedEventHandler[*IAppResourceGroupInfoWatcher, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_EnumerationCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppResourceGroupInfoWatcher) Remove_EnumerationCompleted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_EnumerationCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppResourceGroupInfoWatcher) Add_Stopped(handler TypedEventHandler[*IAppResourceGroupInfoWatcher, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Stopped, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppResourceGroupInfoWatcher) Remove_Stopped(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Stopped, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppResourceGroupInfoWatcher) Add_ExecutionStateChanged(handler TypedEventHandler[*IAppResourceGroupInfoWatcher, *IAppResourceGroupInfoWatcherExecutionStateChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ExecutionStateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppResourceGroupInfoWatcher) Remove_ExecutionStateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ExecutionStateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppResourceGroupInfoWatcher) Get_Status() AppResourceGroupInfoWatcherStatus {
	var _result AppResourceGroupInfoWatcherStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppResourceGroupInfoWatcher) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAppResourceGroupInfoWatcher) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 7A787637-6302-4D2F-BF89-1C12D0B2A6B9
var IID_IAppResourceGroupInfoWatcherEventArgs = syscall.GUID{0x7A787637, 0x6302, 0x4D2F,
	[8]byte{0xBF, 0x89, 0x1C, 0x12, 0xD0, 0xB2, 0xA6, 0xB9}}

type IAppResourceGroupInfoWatcherEventArgsInterface interface {
	win32.IInspectableInterface
	Get_AppDiagnosticInfos() *IVectorView[*IAppDiagnosticInfo]
	Get_AppResourceGroupInfo() *IAppResourceGroupInfo
}

type IAppResourceGroupInfoWatcherEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_AppDiagnosticInfos   uintptr
	Get_AppResourceGroupInfo uintptr
}

type IAppResourceGroupInfoWatcherEventArgs struct {
	win32.IInspectable
}

func (this *IAppResourceGroupInfoWatcherEventArgs) Vtbl() *IAppResourceGroupInfoWatcherEventArgsVtbl {
	return (*IAppResourceGroupInfoWatcherEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppResourceGroupInfoWatcherEventArgs) Get_AppDiagnosticInfos() *IVectorView[*IAppDiagnosticInfo] {
	var _result *IVectorView[*IAppDiagnosticInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppDiagnosticInfos, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppResourceGroupInfoWatcherEventArgs) Get_AppResourceGroupInfo() *IAppResourceGroupInfo {
	var _result *IAppResourceGroupInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppResourceGroupInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1BDBEDD7-FEE6-4FD4-98DD-E92A2CC299F3
var IID_IAppResourceGroupInfoWatcherExecutionStateChangedEventArgs = syscall.GUID{0x1BDBEDD7, 0xFEE6, 0x4FD4,
	[8]byte{0x98, 0xDD, 0xE9, 0x2A, 0x2C, 0xC2, 0x99, 0xF3}}

type IAppResourceGroupInfoWatcherExecutionStateChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_AppDiagnosticInfos() *IVectorView[*IAppDiagnosticInfo]
	Get_AppResourceGroupInfo() *IAppResourceGroupInfo
}

type IAppResourceGroupInfoWatcherExecutionStateChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_AppDiagnosticInfos   uintptr
	Get_AppResourceGroupInfo uintptr
}

type IAppResourceGroupInfoWatcherExecutionStateChangedEventArgs struct {
	win32.IInspectable
}

func (this *IAppResourceGroupInfoWatcherExecutionStateChangedEventArgs) Vtbl() *IAppResourceGroupInfoWatcherExecutionStateChangedEventArgsVtbl {
	return (*IAppResourceGroupInfoWatcherExecutionStateChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppResourceGroupInfoWatcherExecutionStateChangedEventArgs) Get_AppDiagnosticInfos() *IVectorView[*IAppDiagnosticInfo] {
	var _result *IVectorView[*IAppDiagnosticInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppDiagnosticInfos, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppResourceGroupInfoWatcherExecutionStateChangedEventArgs) Get_AppResourceGroupInfo() *IAppResourceGroupInfo {
	var _result *IAppResourceGroupInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppResourceGroupInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2C8C06B1-7DB1-4C51-A225-7FAE2D49E431
var IID_IAppResourceGroupMemoryReport = syscall.GUID{0x2C8C06B1, 0x7DB1, 0x4C51,
	[8]byte{0xA2, 0x25, 0x7F, 0xAE, 0x2D, 0x49, 0xE4, 0x31}}

type IAppResourceGroupMemoryReportInterface interface {
	win32.IInspectableInterface
	Get_CommitUsageLimit() uint64
	Get_CommitUsageLevel() AppMemoryUsageLevel
	Get_PrivateCommitUsage() uint64
	Get_TotalCommitUsage() uint64
}

type IAppResourceGroupMemoryReportVtbl struct {
	win32.IInspectableVtbl
	Get_CommitUsageLimit   uintptr
	Get_CommitUsageLevel   uintptr
	Get_PrivateCommitUsage uintptr
	Get_TotalCommitUsage   uintptr
}

type IAppResourceGroupMemoryReport struct {
	win32.IInspectable
}

func (this *IAppResourceGroupMemoryReport) Vtbl() *IAppResourceGroupMemoryReportVtbl {
	return (*IAppResourceGroupMemoryReportVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppResourceGroupMemoryReport) Get_CommitUsageLimit() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CommitUsageLimit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppResourceGroupMemoryReport) Get_CommitUsageLevel() AppMemoryUsageLevel {
	var _result AppMemoryUsageLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CommitUsageLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppResourceGroupMemoryReport) Get_PrivateCommitUsage() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrivateCommitUsage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppResourceGroupMemoryReport) Get_TotalCommitUsage() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TotalCommitUsage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 52849F18-2F70-4236-AB40-D04DB0C7B931
var IID_IAppResourceGroupStateReport = syscall.GUID{0x52849F18, 0x2F70, 0x4236,
	[8]byte{0xAB, 0x40, 0xD0, 0x4D, 0xB0, 0xC7, 0xB9, 0x31}}

type IAppResourceGroupStateReportInterface interface {
	win32.IInspectableInterface
	Get_ExecutionState() AppResourceGroupExecutionState
	Get_EnergyQuotaState() AppResourceGroupEnergyQuotaState
}

type IAppResourceGroupStateReportVtbl struct {
	win32.IInspectableVtbl
	Get_ExecutionState   uintptr
	Get_EnergyQuotaState uintptr
}

type IAppResourceGroupStateReport struct {
	win32.IInspectable
}

func (this *IAppResourceGroupStateReport) Vtbl() *IAppResourceGroupStateReportVtbl {
	return (*IAppResourceGroupStateReportVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppResourceGroupStateReport) Get_ExecutionState() AppResourceGroupExecutionState {
	var _result AppResourceGroupExecutionState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExecutionState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppResourceGroupStateReport) Get_EnergyQuotaState() AppResourceGroupEnergyQuotaState {
	var _result AppResourceGroupEnergyQuotaState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EnergyQuotaState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5D50CAC5-92D2-5409-B56F-7F73E10EA4C3
var IID_IAppUriHandlerHost = syscall.GUID{0x5D50CAC5, 0x92D2, 0x5409,
	[8]byte{0xB5, 0x6F, 0x7F, 0x73, 0xE1, 0x0E, 0xA4, 0xC3}}

type IAppUriHandlerHostInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Put_Name(value string)
}

type IAppUriHandlerHostVtbl struct {
	win32.IInspectableVtbl
	Get_Name uintptr
	Put_Name uintptr
}

type IAppUriHandlerHost struct {
	win32.IInspectable
}

func (this *IAppUriHandlerHost) Vtbl() *IAppUriHandlerHostVtbl {
	return (*IAppUriHandlerHostVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppUriHandlerHost) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppUriHandlerHost) Put_Name(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Name, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 257C3C96-CE04-5F98-96BB-3EBD3E9275BB
var IID_IAppUriHandlerHostFactory = syscall.GUID{0x257C3C96, 0xCE04, 0x5F98,
	[8]byte{0x96, 0xBB, 0x3E, 0xBD, 0x3E, 0x92, 0x75, 0xBB}}

type IAppUriHandlerHostFactoryInterface interface {
	win32.IInspectableInterface
	CreateInstance(name string) *IAppUriHandlerHost
}

type IAppUriHandlerHostFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateInstance uintptr
}

type IAppUriHandlerHostFactory struct {
	win32.IInspectable
}

func (this *IAppUriHandlerHostFactory) Vtbl() *IAppUriHandlerHostFactoryVtbl {
	return (*IAppUriHandlerHostFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppUriHandlerHostFactory) CreateInstance(name string) *IAppUriHandlerHost {
	var _result *IAppUriHandlerHost
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateInstance, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6F73AEB1-4569-5C3F-9BA0-99123EEA32C3
var IID_IAppUriHandlerRegistration = syscall.GUID{0x6F73AEB1, 0x4569, 0x5C3F,
	[8]byte{0x9B, 0xA0, 0x99, 0x12, 0x3E, 0xEA, 0x32, 0xC3}}

type IAppUriHandlerRegistrationInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Get_User() *IUser
	GetAppAddedHostsAsync() *IAsyncOperation[*IVector[*IAppUriHandlerHost]]
	SetAppAddedHostsAsync(hosts *IIterable[*IAppUriHandlerHost]) *IAsyncAction
}

type IAppUriHandlerRegistrationVtbl struct {
	win32.IInspectableVtbl
	Get_Name              uintptr
	Get_User              uintptr
	GetAppAddedHostsAsync uintptr
	SetAppAddedHostsAsync uintptr
}

type IAppUriHandlerRegistration struct {
	win32.IInspectable
}

func (this *IAppUriHandlerRegistration) Vtbl() *IAppUriHandlerRegistrationVtbl {
	return (*IAppUriHandlerRegistrationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppUriHandlerRegistration) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppUriHandlerRegistration) Get_User() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_User, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppUriHandlerRegistration) GetAppAddedHostsAsync() *IAsyncOperation[*IVector[*IAppUriHandlerHost]] {
	var _result *IAsyncOperation[*IVector[*IAppUriHandlerHost]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAppAddedHostsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppUriHandlerRegistration) SetAppAddedHostsAsync(hosts *IIterable[*IAppUriHandlerHost]) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetAppAddedHostsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(hosts)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E62C9A52-AC94-5750-AC1B-6CFB6F250263
var IID_IAppUriHandlerRegistrationManager = syscall.GUID{0xE62C9A52, 0xAC94, 0x5750,
	[8]byte{0xAC, 0x1B, 0x6C, 0xFB, 0x6F, 0x25, 0x02, 0x63}}

type IAppUriHandlerRegistrationManagerInterface interface {
	win32.IInspectableInterface
	Get_User() *IUser
	TryGetRegistration(name string) *IAppUriHandlerRegistration
}

type IAppUriHandlerRegistrationManagerVtbl struct {
	win32.IInspectableVtbl
	Get_User           uintptr
	TryGetRegistration uintptr
}

type IAppUriHandlerRegistrationManager struct {
	win32.IInspectable
}

func (this *IAppUriHandlerRegistrationManager) Vtbl() *IAppUriHandlerRegistrationManagerVtbl {
	return (*IAppUriHandlerRegistrationManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppUriHandlerRegistrationManager) Get_User() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_User, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppUriHandlerRegistrationManager) TryGetRegistration(name string) *IAppUriHandlerRegistration {
	var _result *IAppUriHandlerRegistration
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetRegistration, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D5CEDD9F-5729-5B76-A1D4-0285F295C124
var IID_IAppUriHandlerRegistrationManagerStatics = syscall.GUID{0xD5CEDD9F, 0x5729, 0x5B76,
	[8]byte{0xA1, 0xD4, 0x02, 0x85, 0xF2, 0x95, 0xC1, 0x24}}

type IAppUriHandlerRegistrationManagerStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *IAppUriHandlerRegistrationManager
	GetForUser(user *IUser) *IAppUriHandlerRegistrationManager
}

type IAppUriHandlerRegistrationManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
	GetForUser uintptr
}

type IAppUriHandlerRegistrationManagerStatics struct {
	win32.IInspectable
}

func (this *IAppUriHandlerRegistrationManagerStatics) Vtbl() *IAppUriHandlerRegistrationManagerStaticsVtbl {
	return (*IAppUriHandlerRegistrationManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppUriHandlerRegistrationManagerStatics) GetDefault() *IAppUriHandlerRegistrationManager {
	var _result *IAppUriHandlerRegistrationManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppUriHandlerRegistrationManagerStatics) GetForUser(user *IUser) *IAppUriHandlerRegistrationManager {
	var _result *IAppUriHandlerRegistrationManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5D2150D1-47EE-48AB-A52B-9F1954278D82
var IID_IDateTimeSettingsStatics = syscall.GUID{0x5D2150D1, 0x47EE, 0x48AB,
	[8]byte{0xA5, 0x2B, 0x9F, 0x19, 0x54, 0x27, 0x8D, 0x82}}

type IDateTimeSettingsStaticsInterface interface {
	win32.IInspectableInterface
	SetSystemDateTime(utcDateTime DateTime)
}

type IDateTimeSettingsStaticsVtbl struct {
	win32.IInspectableVtbl
	SetSystemDateTime uintptr
}

type IDateTimeSettingsStatics struct {
	win32.IInspectable
}

func (this *IDateTimeSettingsStatics) Vtbl() *IDateTimeSettingsStaticsVtbl {
	return (*IDateTimeSettingsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDateTimeSettingsStatics) SetSystemDateTime(utcDateTime DateTime) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetSystemDateTime, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&utcDateTime)))
	_ = _hr
}

// 603E88E4-A338-4FFE-A457-A5CFB9CEB899
var IID_IDispatcherQueue = syscall.GUID{0x603E88E4, 0xA338, 0x4FFE,
	[8]byte{0xA4, 0x57, 0xA5, 0xCF, 0xB9, 0xCE, 0xB8, 0x99}}

type IDispatcherQueueInterface interface {
	win32.IInspectableInterface
	CreateTimer() *IDispatcherQueueTimer
	TryEnqueue(callback DispatcherQueueHandler) bool
	TryEnqueueWithPriority(priority DispatcherQueuePriority, callback DispatcherQueueHandler) bool
	Add_ShutdownStarting(handler TypedEventHandler[*IDispatcherQueue, *IDispatcherQueueShutdownStartingEventArgs]) EventRegistrationToken
	Remove_ShutdownStarting(token EventRegistrationToken)
	Add_ShutdownCompleted(handler TypedEventHandler[*IDispatcherQueue, interface{}]) EventRegistrationToken
	Remove_ShutdownCompleted(token EventRegistrationToken)
}

type IDispatcherQueueVtbl struct {
	win32.IInspectableVtbl
	CreateTimer              uintptr
	TryEnqueue               uintptr
	TryEnqueueWithPriority   uintptr
	Add_ShutdownStarting     uintptr
	Remove_ShutdownStarting  uintptr
	Add_ShutdownCompleted    uintptr
	Remove_ShutdownCompleted uintptr
}

type IDispatcherQueue struct {
	win32.IInspectable
}

func (this *IDispatcherQueue) Vtbl() *IDispatcherQueueVtbl {
	return (*IDispatcherQueueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDispatcherQueue) CreateTimer() *IDispatcherQueueTimer {
	var _result *IDispatcherQueueTimer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateTimer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDispatcherQueue) TryEnqueue(callback DispatcherQueueHandler) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryEnqueue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewNoArgFuncDelegate(callback))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDispatcherQueue) TryEnqueueWithPriority(priority DispatcherQueuePriority, callback DispatcherQueueHandler) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryEnqueueWithPriority, uintptr(unsafe.Pointer(this)), uintptr(priority), uintptr(unsafe.Pointer(NewNoArgFuncDelegate(callback))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDispatcherQueue) Add_ShutdownStarting(handler TypedEventHandler[*IDispatcherQueue, *IDispatcherQueueShutdownStartingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ShutdownStarting, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDispatcherQueue) Remove_ShutdownStarting(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ShutdownStarting, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IDispatcherQueue) Add_ShutdownCompleted(handler TypedEventHandler[*IDispatcherQueue, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ShutdownCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDispatcherQueue) Remove_ShutdownCompleted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ShutdownCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// C822C647-30EF-506E-BD1E-A647AE6675FF
var IID_IDispatcherQueue2 = syscall.GUID{0xC822C647, 0x30EF, 0x506E,
	[8]byte{0xBD, 0x1E, 0xA6, 0x47, 0xAE, 0x66, 0x75, 0xFF}}

type IDispatcherQueue2Interface interface {
	win32.IInspectableInterface
	Get_HasThreadAccess() bool
}

type IDispatcherQueue2Vtbl struct {
	win32.IInspectableVtbl
	Get_HasThreadAccess uintptr
}

type IDispatcherQueue2 struct {
	win32.IInspectable
}

func (this *IDispatcherQueue2) Vtbl() *IDispatcherQueue2Vtbl {
	return (*IDispatcherQueue2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDispatcherQueue2) Get_HasThreadAccess() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasThreadAccess, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 22F34E66-50DB-4E36-A98D-61C01B384D20
var IID_IDispatcherQueueController = syscall.GUID{0x22F34E66, 0x50DB, 0x4E36,
	[8]byte{0xA9, 0x8D, 0x61, 0xC0, 0x1B, 0x38, 0x4D, 0x20}}

type IDispatcherQueueControllerInterface interface {
	win32.IInspectableInterface
	Get_DispatcherQueue() *IDispatcherQueue
	ShutdownQueueAsync() *IAsyncAction
}

type IDispatcherQueueControllerVtbl struct {
	win32.IInspectableVtbl
	Get_DispatcherQueue uintptr
	ShutdownQueueAsync  uintptr
}

type IDispatcherQueueController struct {
	win32.IInspectable
}

func (this *IDispatcherQueueController) Vtbl() *IDispatcherQueueControllerVtbl {
	return (*IDispatcherQueueControllerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDispatcherQueueController) Get_DispatcherQueue() *IDispatcherQueue {
	var _result *IDispatcherQueue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DispatcherQueue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDispatcherQueueController) ShutdownQueueAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShutdownQueueAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0A6C98E0-5198-49A2-A313-3F70D1F13C27
var IID_IDispatcherQueueControllerStatics = syscall.GUID{0x0A6C98E0, 0x5198, 0x49A2,
	[8]byte{0xA3, 0x13, 0x3F, 0x70, 0xD1, 0xF1, 0x3C, 0x27}}

type IDispatcherQueueControllerStaticsInterface interface {
	win32.IInspectableInterface
	CreateOnDedicatedThread() *IDispatcherQueueController
}

type IDispatcherQueueControllerStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateOnDedicatedThread uintptr
}

type IDispatcherQueueControllerStatics struct {
	win32.IInspectable
}

func (this *IDispatcherQueueControllerStatics) Vtbl() *IDispatcherQueueControllerStaticsVtbl {
	return (*IDispatcherQueueControllerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDispatcherQueueControllerStatics) CreateOnDedicatedThread() *IDispatcherQueueController {
	var _result *IDispatcherQueueController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateOnDedicatedThread, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C4724C4C-FF97-40C0-A226-CC0AAA545E89
var IID_IDispatcherQueueShutdownStartingEventArgs = syscall.GUID{0xC4724C4C, 0xFF97, 0x40C0,
	[8]byte{0xA2, 0x26, 0xCC, 0x0A, 0xAA, 0x54, 0x5E, 0x89}}

type IDispatcherQueueShutdownStartingEventArgsInterface interface {
	win32.IInspectableInterface
	GetDeferral() *IDeferral
}

type IDispatcherQueueShutdownStartingEventArgsVtbl struct {
	win32.IInspectableVtbl
	GetDeferral uintptr
}

type IDispatcherQueueShutdownStartingEventArgs struct {
	win32.IInspectable
}

func (this *IDispatcherQueueShutdownStartingEventArgs) Vtbl() *IDispatcherQueueShutdownStartingEventArgsVtbl {
	return (*IDispatcherQueueShutdownStartingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDispatcherQueueShutdownStartingEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A96D83D7-9371-4517-9245-D0824AC12C74
var IID_IDispatcherQueueStatics = syscall.GUID{0xA96D83D7, 0x9371, 0x4517,
	[8]byte{0x92, 0x45, 0xD0, 0x82, 0x4A, 0xC1, 0x2C, 0x74}}

type IDispatcherQueueStaticsInterface interface {
	win32.IInspectableInterface
	GetForCurrentThread() *IDispatcherQueue
}

type IDispatcherQueueStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForCurrentThread uintptr
}

type IDispatcherQueueStatics struct {
	win32.IInspectable
}

func (this *IDispatcherQueueStatics) Vtbl() *IDispatcherQueueStaticsVtbl {
	return (*IDispatcherQueueStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDispatcherQueueStatics) GetForCurrentThread() *IDispatcherQueue {
	var _result *IDispatcherQueue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentThread, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5FEABB1D-A31C-4727-B1AC-37454649D56A
var IID_IDispatcherQueueTimer = syscall.GUID{0x5FEABB1D, 0xA31C, 0x4727,
	[8]byte{0xB1, 0xAC, 0x37, 0x45, 0x46, 0x49, 0xD5, 0x6A}}

type IDispatcherQueueTimerInterface interface {
	win32.IInspectableInterface
	Get_Interval() TimeSpan
	Put_Interval(value TimeSpan)
	Get_IsRunning() bool
	Get_IsRepeating() bool
	Put_IsRepeating(value bool)
	Start()
	Stop()
	Add_Tick(handler TypedEventHandler[*IDispatcherQueueTimer, interface{}]) EventRegistrationToken
	Remove_Tick(token EventRegistrationToken)
}

type IDispatcherQueueTimerVtbl struct {
	win32.IInspectableVtbl
	Get_Interval    uintptr
	Put_Interval    uintptr
	Get_IsRunning   uintptr
	Get_IsRepeating uintptr
	Put_IsRepeating uintptr
	Start           uintptr
	Stop            uintptr
	Add_Tick        uintptr
	Remove_Tick     uintptr
}

type IDispatcherQueueTimer struct {
	win32.IInspectable
}

func (this *IDispatcherQueueTimer) Vtbl() *IDispatcherQueueTimerVtbl {
	return (*IDispatcherQueueTimerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDispatcherQueueTimer) Get_Interval() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Interval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDispatcherQueueTimer) Put_Interval(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Interval, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IDispatcherQueueTimer) Get_IsRunning() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsRunning, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDispatcherQueueTimer) Get_IsRepeating() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsRepeating, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDispatcherQueueTimer) Put_IsRepeating(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsRepeating, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IDispatcherQueueTimer) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IDispatcherQueueTimer) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IDispatcherQueueTimer) Add_Tick(handler TypedEventHandler[*IDispatcherQueueTimer, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Tick, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDispatcherQueueTimer) Remove_Tick(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Tick, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// BB91C27D-6B87-432A-BD04-776C6F5FB2AB
var IID_IFolderLauncherOptions = syscall.GUID{0xBB91C27D, 0x6B87, 0x432A,
	[8]byte{0xBD, 0x04, 0x77, 0x6C, 0x6F, 0x5F, 0xB2, 0xAB}}

type IFolderLauncherOptionsInterface interface {
	win32.IInspectableInterface
	Get_ItemsToSelect() *IVector[*IStorageItem]
}

type IFolderLauncherOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_ItemsToSelect uintptr
}

type IFolderLauncherOptions struct {
	win32.IInspectable
}

func (this *IFolderLauncherOptions) Vtbl() *IFolderLauncherOptionsVtbl {
	return (*IFolderLauncherOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFolderLauncherOptions) Get_ItemsToSelect() *IVector[*IStorageItem] {
	var _result *IVector[*IStorageItem]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ItemsToSelect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7755911A-70C5-48E5-B637-5BA3441E4EE4
var IID_IKnownUserPropertiesStatics = syscall.GUID{0x7755911A, 0x70C5, 0x48E5,
	[8]byte{0xB6, 0x37, 0x5B, 0xA3, 0x44, 0x1E, 0x4E, 0xE4}}

type IKnownUserPropertiesStaticsInterface interface {
	win32.IInspectableInterface
	Get_DisplayName() string
	Get_FirstName() string
	Get_LastName() string
	Get_ProviderName() string
	Get_AccountName() string
	Get_GuestHost() string
	Get_PrincipalName() string
	Get_DomainName() string
	Get_SessionInitiationProtocolUri() string
}

type IKnownUserPropertiesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_DisplayName                  uintptr
	Get_FirstName                    uintptr
	Get_LastName                     uintptr
	Get_ProviderName                 uintptr
	Get_AccountName                  uintptr
	Get_GuestHost                    uintptr
	Get_PrincipalName                uintptr
	Get_DomainName                   uintptr
	Get_SessionInitiationProtocolUri uintptr
}

type IKnownUserPropertiesStatics struct {
	win32.IInspectable
}

func (this *IKnownUserPropertiesStatics) Vtbl() *IKnownUserPropertiesStaticsVtbl {
	return (*IKnownUserPropertiesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownUserPropertiesStatics) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownUserPropertiesStatics) Get_FirstName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FirstName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownUserPropertiesStatics) Get_LastName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LastName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownUserPropertiesStatics) Get_ProviderName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownUserPropertiesStatics) Get_AccountName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AccountName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownUserPropertiesStatics) Get_GuestHost() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GuestHost, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownUserPropertiesStatics) Get_PrincipalName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrincipalName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownUserPropertiesStatics) Get_DomainName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DomainName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownUserPropertiesStatics) Get_SessionInitiationProtocolUri() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SessionInitiationProtocolUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// EC27A8DF-F6D5-45CA-913A-70A40C5C8221
var IID_ILaunchUriResult = syscall.GUID{0xEC27A8DF, 0xF6D5, 0x45CA,
	[8]byte{0x91, 0x3A, 0x70, 0xA4, 0x0C, 0x5C, 0x82, 0x21}}

type ILaunchUriResultInterface interface {
	win32.IInspectableInterface
	Get_Status() LaunchUriStatus
	Get_Result() *IPropertySet
}

type ILaunchUriResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
	Get_Result uintptr
}

type ILaunchUriResult struct {
	win32.IInspectable
}

func (this *ILaunchUriResult) Vtbl() *ILaunchUriResultVtbl {
	return (*ILaunchUriResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILaunchUriResult) Get_Status() LaunchUriStatus {
	var _result LaunchUriStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILaunchUriResult) Get_Result() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Result, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BAFA21D8-B071-4CD8-853E-341203E557D3
var IID_ILauncherOptions = syscall.GUID{0xBAFA21D8, 0xB071, 0x4CD8,
	[8]byte{0x85, 0x3E, 0x34, 0x12, 0x03, 0xE5, 0x57, 0xD3}}

type ILauncherOptionsInterface interface {
	win32.IInspectableInterface
	Get_TreatAsUntrusted() bool
	Put_TreatAsUntrusted(value bool)
	Get_DisplayApplicationPicker() bool
	Put_DisplayApplicationPicker(value bool)
	Get_UI() *ILauncherUIOptions
	Get_PreferredApplicationPackageFamilyName() string
	Put_PreferredApplicationPackageFamilyName(value string)
	Get_PreferredApplicationDisplayName() string
	Put_PreferredApplicationDisplayName(value string)
	Get_FallbackUri() *IUriRuntimeClass
	Put_FallbackUri(value *IUriRuntimeClass)
	Get_ContentType() string
	Put_ContentType(value string)
}

type ILauncherOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_TreatAsUntrusted                      uintptr
	Put_TreatAsUntrusted                      uintptr
	Get_DisplayApplicationPicker              uintptr
	Put_DisplayApplicationPicker              uintptr
	Get_UI                                    uintptr
	Get_PreferredApplicationPackageFamilyName uintptr
	Put_PreferredApplicationPackageFamilyName uintptr
	Get_PreferredApplicationDisplayName       uintptr
	Put_PreferredApplicationDisplayName       uintptr
	Get_FallbackUri                           uintptr
	Put_FallbackUri                           uintptr
	Get_ContentType                           uintptr
	Put_ContentType                           uintptr
}

type ILauncherOptions struct {
	win32.IInspectable
}

func (this *ILauncherOptions) Vtbl() *ILauncherOptionsVtbl {
	return (*ILauncherOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILauncherOptions) Get_TreatAsUntrusted() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TreatAsUntrusted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILauncherOptions) Put_TreatAsUntrusted(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TreatAsUntrusted, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ILauncherOptions) Get_DisplayApplicationPicker() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayApplicationPicker, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILauncherOptions) Put_DisplayApplicationPicker(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DisplayApplicationPicker, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ILauncherOptions) Get_UI() *ILauncherUIOptions {
	var _result *ILauncherUIOptions
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UI, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherOptions) Get_PreferredApplicationPackageFamilyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PreferredApplicationPackageFamilyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILauncherOptions) Put_PreferredApplicationPackageFamilyName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PreferredApplicationPackageFamilyName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ILauncherOptions) Get_PreferredApplicationDisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PreferredApplicationDisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILauncherOptions) Put_PreferredApplicationDisplayName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PreferredApplicationDisplayName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ILauncherOptions) Get_FallbackUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FallbackUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherOptions) Put_FallbackUri(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FallbackUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ILauncherOptions) Get_ContentType() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILauncherOptions) Put_ContentType(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ContentType, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 3BA08EB4-6E40-4DCE-A1A3-2F53950AFB49
var IID_ILauncherOptions2 = syscall.GUID{0x3BA08EB4, 0x6E40, 0x4DCE,
	[8]byte{0xA1, 0xA3, 0x2F, 0x53, 0x95, 0x0A, 0xFB, 0x49}}

type ILauncherOptions2Interface interface {
	win32.IInspectableInterface
	Get_TargetApplicationPackageFamilyName() string
	Put_TargetApplicationPackageFamilyName(value string)
	Get_NeighboringFilesQuery() unsafe.Pointer
	Put_NeighboringFilesQuery(value unsafe.Pointer)
}

type ILauncherOptions2Vtbl struct {
	win32.IInspectableVtbl
	Get_TargetApplicationPackageFamilyName uintptr
	Put_TargetApplicationPackageFamilyName uintptr
	Get_NeighboringFilesQuery              uintptr
	Put_NeighboringFilesQuery              uintptr
}

type ILauncherOptions2 struct {
	win32.IInspectable
}

func (this *ILauncherOptions2) Vtbl() *ILauncherOptions2Vtbl {
	return (*ILauncherOptions2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILauncherOptions2) Get_TargetApplicationPackageFamilyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TargetApplicationPackageFamilyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILauncherOptions2) Put_TargetApplicationPackageFamilyName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TargetApplicationPackageFamilyName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ILauncherOptions2) Get_NeighboringFilesQuery() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NeighboringFilesQuery, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILauncherOptions2) Put_NeighboringFilesQuery(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NeighboringFilesQuery, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// F0770655-4B63-4E3A-9107-4E687841923A
var IID_ILauncherOptions3 = syscall.GUID{0xF0770655, 0x4B63, 0x4E3A,
	[8]byte{0x91, 0x07, 0x4E, 0x68, 0x78, 0x41, 0x92, 0x3A}}

type ILauncherOptions3Interface interface {
	win32.IInspectableInterface
	Get_IgnoreAppUriHandlers() bool
	Put_IgnoreAppUriHandlers(value bool)
}

type ILauncherOptions3Vtbl struct {
	win32.IInspectableVtbl
	Get_IgnoreAppUriHandlers uintptr
	Put_IgnoreAppUriHandlers uintptr
}

type ILauncherOptions3 struct {
	win32.IInspectable
}

func (this *ILauncherOptions3) Vtbl() *ILauncherOptions3Vtbl {
	return (*ILauncherOptions3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILauncherOptions3) Get_IgnoreAppUriHandlers() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IgnoreAppUriHandlers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILauncherOptions3) Put_IgnoreAppUriHandlers(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IgnoreAppUriHandlers, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// EF6FD10E-E6FB-4814-A44E-57E8B9D9A01B
var IID_ILauncherOptions4 = syscall.GUID{0xEF6FD10E, 0xE6FB, 0x4814,
	[8]byte{0xA4, 0x4E, 0x57, 0xE8, 0xB9, 0xD9, 0xA0, 0x1B}}

type ILauncherOptions4Interface interface {
	win32.IInspectableInterface
	Get_LimitPickerToCurrentAppAndAppUriHandlers() bool
	Put_LimitPickerToCurrentAppAndAppUriHandlers(value bool)
}

type ILauncherOptions4Vtbl struct {
	win32.IInspectableVtbl
	Get_LimitPickerToCurrentAppAndAppUriHandlers uintptr
	Put_LimitPickerToCurrentAppAndAppUriHandlers uintptr
}

type ILauncherOptions4 struct {
	win32.IInspectable
}

func (this *ILauncherOptions4) Vtbl() *ILauncherOptions4Vtbl {
	return (*ILauncherOptions4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILauncherOptions4) Get_LimitPickerToCurrentAppAndAppUriHandlers() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LimitPickerToCurrentAppAndAppUriHandlers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILauncherOptions4) Put_LimitPickerToCurrentAppAndAppUriHandlers(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LimitPickerToCurrentAppAndAppUriHandlers, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 277151C3-9E3E-42F6-91A4-5DFDEB232451
var IID_ILauncherStatics = syscall.GUID{0x277151C3, 0x9E3E, 0x42F6,
	[8]byte{0x91, 0xA4, 0x5D, 0xFD, 0xEB, 0x23, 0x24, 0x51}}

type ILauncherStaticsInterface interface {
	win32.IInspectableInterface
	LaunchFileAsync(file *IStorageFile) *IAsyncOperation[bool]
	LaunchFileWithOptionsAsync(file *IStorageFile, options *ILauncherOptions) *IAsyncOperation[bool]
	LaunchUriAsync(uri *IUriRuntimeClass) *IAsyncOperation[bool]
	LaunchUriWithOptionsAsync(uri *IUriRuntimeClass, options *ILauncherOptions) *IAsyncOperation[bool]
}

type ILauncherStaticsVtbl struct {
	win32.IInspectableVtbl
	LaunchFileAsync            uintptr
	LaunchFileWithOptionsAsync uintptr
	LaunchUriAsync             uintptr
	LaunchUriWithOptionsAsync  uintptr
}

type ILauncherStatics struct {
	win32.IInspectable
}

func (this *ILauncherStatics) Vtbl() *ILauncherStaticsVtbl {
	return (*ILauncherStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILauncherStatics) LaunchFileAsync(file *IStorageFile) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchFileAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherStatics) LaunchFileWithOptionsAsync(file *IStorageFile, options *ILauncherOptions) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchFileWithOptionsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherStatics) LaunchUriAsync(uri *IUriRuntimeClass) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchUriAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherStatics) LaunchUriWithOptionsAsync(uri *IUriRuntimeClass, options *ILauncherOptions) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchUriWithOptionsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 59BA2FBB-24CB-4C02-A4C4-8294569D54F1
var IID_ILauncherStatics2 = syscall.GUID{0x59BA2FBB, 0x24CB, 0x4C02,
	[8]byte{0xA4, 0xC4, 0x82, 0x94, 0x56, 0x9D, 0x54, 0xF1}}

type ILauncherStatics2Interface interface {
	win32.IInspectableInterface
	LaunchUriForResultsAsync(uri *IUriRuntimeClass, options *ILauncherOptions) *IAsyncOperation[*ILaunchUriResult]
	LaunchUriForResultsWithDataAsync(uri *IUriRuntimeClass, options *ILauncherOptions, inputData *IPropertySet) *IAsyncOperation[*ILaunchUriResult]
	LaunchUriWithDataAsync(uri *IUriRuntimeClass, options *ILauncherOptions, inputData *IPropertySet) *IAsyncOperation[bool]
	QueryUriSupportAsync(uri *IUriRuntimeClass, launchQuerySupportType LaunchQuerySupportType) *IAsyncOperation[LaunchQuerySupportStatus]
	QueryUriSupportWithPackageFamilyNameAsync(uri *IUriRuntimeClass, launchQuerySupportType LaunchQuerySupportType, packageFamilyName string) *IAsyncOperation[LaunchQuerySupportStatus]
	QueryFileSupportAsync(file *IStorageFile) *IAsyncOperation[LaunchQuerySupportStatus]
	QueryFileSupportWithPackageFamilyNameAsync(file *IStorageFile, packageFamilyName string) *IAsyncOperation[LaunchQuerySupportStatus]
	FindUriSchemeHandlersAsync(scheme string) *IAsyncOperation[*IVectorView[*IAppInfo]]
	FindUriSchemeHandlersWithLaunchUriTypeAsync(scheme string, launchQuerySupportType LaunchQuerySupportType) *IAsyncOperation[*IVectorView[*IAppInfo]]
	FindFileHandlersAsync(extension string) *IAsyncOperation[*IVectorView[*IAppInfo]]
}

type ILauncherStatics2Vtbl struct {
	win32.IInspectableVtbl
	LaunchUriForResultsAsync                    uintptr
	LaunchUriForResultsWithDataAsync            uintptr
	LaunchUriWithDataAsync                      uintptr
	QueryUriSupportAsync                        uintptr
	QueryUriSupportWithPackageFamilyNameAsync   uintptr
	QueryFileSupportAsync                       uintptr
	QueryFileSupportWithPackageFamilyNameAsync  uintptr
	FindUriSchemeHandlersAsync                  uintptr
	FindUriSchemeHandlersWithLaunchUriTypeAsync uintptr
	FindFileHandlersAsync                       uintptr
}

type ILauncherStatics2 struct {
	win32.IInspectable
}

func (this *ILauncherStatics2) Vtbl() *ILauncherStatics2Vtbl {
	return (*ILauncherStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILauncherStatics2) LaunchUriForResultsAsync(uri *IUriRuntimeClass, options *ILauncherOptions) *IAsyncOperation[*ILaunchUriResult] {
	var _result *IAsyncOperation[*ILaunchUriResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchUriForResultsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherStatics2) LaunchUriForResultsWithDataAsync(uri *IUriRuntimeClass, options *ILauncherOptions, inputData *IPropertySet) *IAsyncOperation[*ILaunchUriResult] {
	var _result *IAsyncOperation[*ILaunchUriResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchUriForResultsWithDataAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(inputData)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherStatics2) LaunchUriWithDataAsync(uri *IUriRuntimeClass, options *ILauncherOptions, inputData *IPropertySet) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchUriWithDataAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(inputData)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherStatics2) QueryUriSupportAsync(uri *IUriRuntimeClass, launchQuerySupportType LaunchQuerySupportType) *IAsyncOperation[LaunchQuerySupportStatus] {
	var _result *IAsyncOperation[LaunchQuerySupportStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().QueryUriSupportAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(launchQuerySupportType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherStatics2) QueryUriSupportWithPackageFamilyNameAsync(uri *IUriRuntimeClass, launchQuerySupportType LaunchQuerySupportType, packageFamilyName string) *IAsyncOperation[LaunchQuerySupportStatus] {
	var _result *IAsyncOperation[LaunchQuerySupportStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().QueryUriSupportWithPackageFamilyNameAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(launchQuerySupportType), NewHStr(packageFamilyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherStatics2) QueryFileSupportAsync(file *IStorageFile) *IAsyncOperation[LaunchQuerySupportStatus] {
	var _result *IAsyncOperation[LaunchQuerySupportStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().QueryFileSupportAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherStatics2) QueryFileSupportWithPackageFamilyNameAsync(file *IStorageFile, packageFamilyName string) *IAsyncOperation[LaunchQuerySupportStatus] {
	var _result *IAsyncOperation[LaunchQuerySupportStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().QueryFileSupportWithPackageFamilyNameAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), NewHStr(packageFamilyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherStatics2) FindUriSchemeHandlersAsync(scheme string) *IAsyncOperation[*IVectorView[*IAppInfo]] {
	var _result *IAsyncOperation[*IVectorView[*IAppInfo]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindUriSchemeHandlersAsync, uintptr(unsafe.Pointer(this)), NewHStr(scheme).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherStatics2) FindUriSchemeHandlersWithLaunchUriTypeAsync(scheme string, launchQuerySupportType LaunchQuerySupportType) *IAsyncOperation[*IVectorView[*IAppInfo]] {
	var _result *IAsyncOperation[*IVectorView[*IAppInfo]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindUriSchemeHandlersWithLaunchUriTypeAsync, uintptr(unsafe.Pointer(this)), NewHStr(scheme).Ptr, uintptr(launchQuerySupportType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherStatics2) FindFileHandlersAsync(extension string) *IAsyncOperation[*IVectorView[*IAppInfo]] {
	var _result *IAsyncOperation[*IVectorView[*IAppInfo]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindFileHandlersAsync, uintptr(unsafe.Pointer(this)), NewHStr(extension).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 234261A8-9DB3-4683-AA42-DC6F51D33847
var IID_ILauncherStatics3 = syscall.GUID{0x234261A8, 0x9DB3, 0x4683,
	[8]byte{0xAA, 0x42, 0xDC, 0x6F, 0x51, 0xD3, 0x38, 0x47}}

type ILauncherStatics3Interface interface {
	win32.IInspectableInterface
	LaunchFolderAsync(folder *IStorageFolder) *IAsyncOperation[bool]
	LaunchFolderWithOptionsAsync(folder *IStorageFolder, options *IFolderLauncherOptions) *IAsyncOperation[bool]
}

type ILauncherStatics3Vtbl struct {
	win32.IInspectableVtbl
	LaunchFolderAsync            uintptr
	LaunchFolderWithOptionsAsync uintptr
}

type ILauncherStatics3 struct {
	win32.IInspectable
}

func (this *ILauncherStatics3) Vtbl() *ILauncherStatics3Vtbl {
	return (*ILauncherStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILauncherStatics3) LaunchFolderAsync(folder *IStorageFolder) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchFolderAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(folder)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherStatics3) LaunchFolderWithOptionsAsync(folder *IStorageFolder, options *IFolderLauncherOptions) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchFolderWithOptionsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(folder)), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B9EC819F-B5A5-41C6-B3B3-DD1B3178BCF2
var IID_ILauncherStatics4 = syscall.GUID{0xB9EC819F, 0xB5A5, 0x41C6,
	[8]byte{0xB3, 0xB3, 0xDD, 0x1B, 0x31, 0x78, 0xBC, 0xF2}}

type ILauncherStatics4Interface interface {
	win32.IInspectableInterface
	QueryAppUriSupportAsync(uri *IUriRuntimeClass) *IAsyncOperation[LaunchQuerySupportStatus]
	QueryAppUriSupportWithPackageFamilyNameAsync(uri *IUriRuntimeClass, packageFamilyName string) *IAsyncOperation[LaunchQuerySupportStatus]
	FindAppUriHandlersAsync(uri *IUriRuntimeClass) *IAsyncOperation[*IVectorView[*IAppInfo]]
	LaunchUriForUserAsync(user *IUser, uri *IUriRuntimeClass) *IAsyncOperation[LaunchUriStatus]
	LaunchUriWithOptionsForUserAsync(user *IUser, uri *IUriRuntimeClass, options *ILauncherOptions) *IAsyncOperation[LaunchUriStatus]
	LaunchUriWithDataForUserAsync(user *IUser, uri *IUriRuntimeClass, options *ILauncherOptions, inputData *IPropertySet) *IAsyncOperation[LaunchUriStatus]
	LaunchUriForResultsForUserAsync(user *IUser, uri *IUriRuntimeClass, options *ILauncherOptions) *IAsyncOperation[*ILaunchUriResult]
	LaunchUriForResultsWithDataForUserAsync(user *IUser, uri *IUriRuntimeClass, options *ILauncherOptions, inputData *IPropertySet) *IAsyncOperation[*ILaunchUriResult]
}

type ILauncherStatics4Vtbl struct {
	win32.IInspectableVtbl
	QueryAppUriSupportAsync                      uintptr
	QueryAppUriSupportWithPackageFamilyNameAsync uintptr
	FindAppUriHandlersAsync                      uintptr
	LaunchUriForUserAsync                        uintptr
	LaunchUriWithOptionsForUserAsync             uintptr
	LaunchUriWithDataForUserAsync                uintptr
	LaunchUriForResultsForUserAsync              uintptr
	LaunchUriForResultsWithDataForUserAsync      uintptr
}

type ILauncherStatics4 struct {
	win32.IInspectable
}

func (this *ILauncherStatics4) Vtbl() *ILauncherStatics4Vtbl {
	return (*ILauncherStatics4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILauncherStatics4) QueryAppUriSupportAsync(uri *IUriRuntimeClass) *IAsyncOperation[LaunchQuerySupportStatus] {
	var _result *IAsyncOperation[LaunchQuerySupportStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().QueryAppUriSupportAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherStatics4) QueryAppUriSupportWithPackageFamilyNameAsync(uri *IUriRuntimeClass, packageFamilyName string) *IAsyncOperation[LaunchQuerySupportStatus] {
	var _result *IAsyncOperation[LaunchQuerySupportStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().QueryAppUriSupportWithPackageFamilyNameAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), NewHStr(packageFamilyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherStatics4) FindAppUriHandlersAsync(uri *IUriRuntimeClass) *IAsyncOperation[*IVectorView[*IAppInfo]] {
	var _result *IAsyncOperation[*IVectorView[*IAppInfo]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAppUriHandlersAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherStatics4) LaunchUriForUserAsync(user *IUser, uri *IUriRuntimeClass) *IAsyncOperation[LaunchUriStatus] {
	var _result *IAsyncOperation[LaunchUriStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchUriForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherStatics4) LaunchUriWithOptionsForUserAsync(user *IUser, uri *IUriRuntimeClass, options *ILauncherOptions) *IAsyncOperation[LaunchUriStatus] {
	var _result *IAsyncOperation[LaunchUriStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchUriWithOptionsForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherStatics4) LaunchUriWithDataForUserAsync(user *IUser, uri *IUriRuntimeClass, options *ILauncherOptions, inputData *IPropertySet) *IAsyncOperation[LaunchUriStatus] {
	var _result *IAsyncOperation[LaunchUriStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchUriWithDataForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(inputData)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherStatics4) LaunchUriForResultsForUserAsync(user *IUser, uri *IUriRuntimeClass, options *ILauncherOptions) *IAsyncOperation[*ILaunchUriResult] {
	var _result *IAsyncOperation[*ILaunchUriResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchUriForResultsForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherStatics4) LaunchUriForResultsWithDataForUserAsync(user *IUser, uri *IUriRuntimeClass, options *ILauncherOptions, inputData *IPropertySet) *IAsyncOperation[*ILaunchUriResult] {
	var _result *IAsyncOperation[*ILaunchUriResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchUriForResultsWithDataForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(inputData)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5B24EF84-D895-5FEA-9153-1AC49AED9BA9
var IID_ILauncherStatics5 = syscall.GUID{0x5B24EF84, 0xD895, 0x5FEA,
	[8]byte{0x91, 0x53, 0x1A, 0xC4, 0x9A, 0xED, 0x9B, 0xA9}}

type ILauncherStatics5Interface interface {
	win32.IInspectableInterface
	LaunchFolderPathAsync(path string) *IAsyncOperation[bool]
	LaunchFolderPathWithOptionsAsync(path string, options *IFolderLauncherOptions) *IAsyncOperation[bool]
	LaunchFolderPathForUserAsync(user *IUser, path string) *IAsyncOperation[bool]
	LaunchFolderPathWithOptionsForUserAsync(user *IUser, path string, options *IFolderLauncherOptions) *IAsyncOperation[bool]
}

type ILauncherStatics5Vtbl struct {
	win32.IInspectableVtbl
	LaunchFolderPathAsync                   uintptr
	LaunchFolderPathWithOptionsAsync        uintptr
	LaunchFolderPathForUserAsync            uintptr
	LaunchFolderPathWithOptionsForUserAsync uintptr
}

type ILauncherStatics5 struct {
	win32.IInspectable
}

func (this *ILauncherStatics5) Vtbl() *ILauncherStatics5Vtbl {
	return (*ILauncherStatics5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILauncherStatics5) LaunchFolderPathAsync(path string) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchFolderPathAsync, uintptr(unsafe.Pointer(this)), NewHStr(path).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherStatics5) LaunchFolderPathWithOptionsAsync(path string, options *IFolderLauncherOptions) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchFolderPathWithOptionsAsync, uintptr(unsafe.Pointer(this)), NewHStr(path).Ptr, uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherStatics5) LaunchFolderPathForUserAsync(user *IUser, path string) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchFolderPathForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), NewHStr(path).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherStatics5) LaunchFolderPathWithOptionsForUserAsync(user *IUser, path string, options *IFolderLauncherOptions) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchFolderPathWithOptionsForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), NewHStr(path).Ptr, uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1B25DA6E-8AA6-41E9-8251-4165F5985F49
var IID_ILauncherUIOptions = syscall.GUID{0x1B25DA6E, 0x8AA6, 0x41E9,
	[8]byte{0x82, 0x51, 0x41, 0x65, 0xF5, 0x98, 0x5F, 0x49}}

type ILauncherUIOptionsInterface interface {
	win32.IInspectableInterface
	Get_InvocationPoint() *IReference[Point]
	Put_InvocationPoint(value *IReference[Point])
	Get_SelectionRect() *IReference[Rect]
	Put_SelectionRect(value *IReference[Rect])
	Get_PreferredPlacement() Placement
	Put_PreferredPlacement(value Placement)
}

type ILauncherUIOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_InvocationPoint    uintptr
	Put_InvocationPoint    uintptr
	Get_SelectionRect      uintptr
	Put_SelectionRect      uintptr
	Get_PreferredPlacement uintptr
	Put_PreferredPlacement uintptr
}

type ILauncherUIOptions struct {
	win32.IInspectable
}

func (this *ILauncherUIOptions) Vtbl() *ILauncherUIOptionsVtbl {
	return (*ILauncherUIOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILauncherUIOptions) Get_InvocationPoint() *IReference[Point] {
	var _result *IReference[Point]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InvocationPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherUIOptions) Put_InvocationPoint(value *IReference[Point]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InvocationPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ILauncherUIOptions) Get_SelectionRect() *IReference[Rect] {
	var _result *IReference[Rect]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectionRect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILauncherUIOptions) Put_SelectionRect(value *IReference[Rect]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SelectionRect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ILauncherUIOptions) Get_PreferredPlacement() Placement {
	var _result Placement
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PreferredPlacement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILauncherUIOptions) Put_PreferredPlacement(value Placement) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PreferredPlacement, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 8A9B29F1-7CA7-49DE-9BD3-3C5B7184F616
var IID_ILauncherViewOptions = syscall.GUID{0x8A9B29F1, 0x7CA7, 0x49DE,
	[8]byte{0x9B, 0xD3, 0x3C, 0x5B, 0x71, 0x84, 0xF6, 0x16}}

type ILauncherViewOptionsInterface interface {
	win32.IInspectableInterface
	Get_DesiredRemainingView() ViewSizePreference
	Put_DesiredRemainingView(value ViewSizePreference)
}

type ILauncherViewOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_DesiredRemainingView uintptr
	Put_DesiredRemainingView uintptr
}

type ILauncherViewOptions struct {
	win32.IInspectable
}

func (this *ILauncherViewOptions) Vtbl() *ILauncherViewOptionsVtbl {
	return (*ILauncherViewOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILauncherViewOptions) Get_DesiredRemainingView() ViewSizePreference {
	var _result ViewSizePreference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesiredRemainingView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILauncherViewOptions) Put_DesiredRemainingView(value ViewSizePreference) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DesiredRemainingView, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 5C6C279C-D7CA-4779-9188-4057219CE64C
var IID_IMemoryManagerStatics = syscall.GUID{0x5C6C279C, 0xD7CA, 0x4779,
	[8]byte{0x91, 0x88, 0x40, 0x57, 0x21, 0x9C, 0xE6, 0x4C}}

type IMemoryManagerStaticsInterface interface {
	win32.IInspectableInterface
	Get_AppMemoryUsage() uint64
	Get_AppMemoryUsageLimit() uint64
	Get_AppMemoryUsageLevel() AppMemoryUsageLevel
	Add_AppMemoryUsageIncreased(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_AppMemoryUsageIncreased(token EventRegistrationToken)
	Add_AppMemoryUsageDecreased(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_AppMemoryUsageDecreased(token EventRegistrationToken)
	Add_AppMemoryUsageLimitChanging(handler EventHandler[*IAppMemoryUsageLimitChangingEventArgs]) EventRegistrationToken
	Remove_AppMemoryUsageLimitChanging(token EventRegistrationToken)
}

type IMemoryManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_AppMemoryUsage                 uintptr
	Get_AppMemoryUsageLimit            uintptr
	Get_AppMemoryUsageLevel            uintptr
	Add_AppMemoryUsageIncreased        uintptr
	Remove_AppMemoryUsageIncreased     uintptr
	Add_AppMemoryUsageDecreased        uintptr
	Remove_AppMemoryUsageDecreased     uintptr
	Add_AppMemoryUsageLimitChanging    uintptr
	Remove_AppMemoryUsageLimitChanging uintptr
}

type IMemoryManagerStatics struct {
	win32.IInspectable
}

func (this *IMemoryManagerStatics) Vtbl() *IMemoryManagerStaticsVtbl {
	return (*IMemoryManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMemoryManagerStatics) Get_AppMemoryUsage() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppMemoryUsage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMemoryManagerStatics) Get_AppMemoryUsageLimit() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppMemoryUsageLimit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMemoryManagerStatics) Get_AppMemoryUsageLevel() AppMemoryUsageLevel {
	var _result AppMemoryUsageLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppMemoryUsageLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMemoryManagerStatics) Add_AppMemoryUsageIncreased(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AppMemoryUsageIncreased, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMemoryManagerStatics) Remove_AppMemoryUsageIncreased(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AppMemoryUsageIncreased, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMemoryManagerStatics) Add_AppMemoryUsageDecreased(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AppMemoryUsageDecreased, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMemoryManagerStatics) Remove_AppMemoryUsageDecreased(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AppMemoryUsageDecreased, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMemoryManagerStatics) Add_AppMemoryUsageLimitChanging(handler EventHandler[*IAppMemoryUsageLimitChangingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AppMemoryUsageLimitChanging, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMemoryManagerStatics) Remove_AppMemoryUsageLimitChanging(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AppMemoryUsageLimitChanging, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 6EEE351F-6D62-423F-9479-B01F9C9F7669
var IID_IMemoryManagerStatics2 = syscall.GUID{0x6EEE351F, 0x6D62, 0x423F,
	[8]byte{0x94, 0x79, 0xB0, 0x1F, 0x9C, 0x9F, 0x76, 0x69}}

type IMemoryManagerStatics2Interface interface {
	win32.IInspectableInterface
	GetAppMemoryReport() *IAppMemoryReport
	GetProcessMemoryReport() *IProcessMemoryReport
}

type IMemoryManagerStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetAppMemoryReport     uintptr
	GetProcessMemoryReport uintptr
}

type IMemoryManagerStatics2 struct {
	win32.IInspectable
}

func (this *IMemoryManagerStatics2) Vtbl() *IMemoryManagerStatics2Vtbl {
	return (*IMemoryManagerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMemoryManagerStatics2) GetAppMemoryReport() *IAppMemoryReport {
	var _result *IAppMemoryReport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAppMemoryReport, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMemoryManagerStatics2) GetProcessMemoryReport() *IProcessMemoryReport {
	var _result *IProcessMemoryReport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetProcessMemoryReport, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 149B59CE-92AD-4E35-89EB-50DFB4C0D91C
var IID_IMemoryManagerStatics3 = syscall.GUID{0x149B59CE, 0x92AD, 0x4E35,
	[8]byte{0x89, 0xEB, 0x50, 0xDF, 0xB4, 0xC0, 0xD9, 0x1C}}

type IMemoryManagerStatics3Interface interface {
	win32.IInspectableInterface
	TrySetAppMemoryUsageLimit(value uint64) bool
}

type IMemoryManagerStatics3Vtbl struct {
	win32.IInspectableVtbl
	TrySetAppMemoryUsageLimit uintptr
}

type IMemoryManagerStatics3 struct {
	win32.IInspectable
}

func (this *IMemoryManagerStatics3) Vtbl() *IMemoryManagerStatics3Vtbl {
	return (*IMemoryManagerStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMemoryManagerStatics3) TrySetAppMemoryUsageLimit(value uint64) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySetAppMemoryUsageLimit, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C5A94828-E84E-4886-8A0D-44B3190E3B72
var IID_IMemoryManagerStatics4 = syscall.GUID{0xC5A94828, 0xE84E, 0x4886,
	[8]byte{0x8A, 0x0D, 0x44, 0xB3, 0x19, 0x0E, 0x3B, 0x72}}

type IMemoryManagerStatics4Interface interface {
	win32.IInspectableInterface
	Get_ExpectedAppMemoryUsageLimit() uint64
}

type IMemoryManagerStatics4Vtbl struct {
	win32.IInspectableVtbl
	Get_ExpectedAppMemoryUsageLimit uintptr
}

type IMemoryManagerStatics4 struct {
	win32.IInspectable
}

func (this *IMemoryManagerStatics4) Vtbl() *IMemoryManagerStatics4Vtbl {
	return (*IMemoryManagerStatics4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMemoryManagerStatics4) Get_ExpectedAppMemoryUsageLimit() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExpectedAppMemoryUsageLimit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3080B9CF-F444-4A83-BEAF-A549A0F3229C
var IID_IProcessLauncherOptions = syscall.GUID{0x3080B9CF, 0xF444, 0x4A83,
	[8]byte{0xBE, 0xAF, 0xA5, 0x49, 0xA0, 0xF3, 0x22, 0x9C}}

type IProcessLauncherOptionsInterface interface {
	win32.IInspectableInterface
	Get_StandardInput() *IInputStream
	Put_StandardInput(value *IInputStream)
	Get_StandardOutput() *IOutputStream
	Put_StandardOutput(value *IOutputStream)
	Get_StandardError() *IOutputStream
	Put_StandardError(value *IOutputStream)
	Get_WorkingDirectory() string
	Put_WorkingDirectory(value string)
}

type IProcessLauncherOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_StandardInput    uintptr
	Put_StandardInput    uintptr
	Get_StandardOutput   uintptr
	Put_StandardOutput   uintptr
	Get_StandardError    uintptr
	Put_StandardError    uintptr
	Get_WorkingDirectory uintptr
	Put_WorkingDirectory uintptr
}

type IProcessLauncherOptions struct {
	win32.IInspectable
}

func (this *IProcessLauncherOptions) Vtbl() *IProcessLauncherOptionsVtbl {
	return (*IProcessLauncherOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProcessLauncherOptions) Get_StandardInput() *IInputStream {
	var _result *IInputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StandardInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProcessLauncherOptions) Put_StandardInput(value *IInputStream) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StandardInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IProcessLauncherOptions) Get_StandardOutput() *IOutputStream {
	var _result *IOutputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StandardOutput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProcessLauncherOptions) Put_StandardOutput(value *IOutputStream) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StandardOutput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IProcessLauncherOptions) Get_StandardError() *IOutputStream {
	var _result *IOutputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StandardError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProcessLauncherOptions) Put_StandardError(value *IOutputStream) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StandardError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IProcessLauncherOptions) Get_WorkingDirectory() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WorkingDirectory, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IProcessLauncherOptions) Put_WorkingDirectory(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_WorkingDirectory, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 544C8934-86D8-4991-8E75-ECE8A43B6B6D
var IID_IProcessLauncherResult = syscall.GUID{0x544C8934, 0x86D8, 0x4991,
	[8]byte{0x8E, 0x75, 0xEC, 0xE8, 0xA4, 0x3B, 0x6B, 0x6D}}

type IProcessLauncherResultInterface interface {
	win32.IInspectableInterface
	Get_ExitCode() uint32
}

type IProcessLauncherResultVtbl struct {
	win32.IInspectableVtbl
	Get_ExitCode uintptr
}

type IProcessLauncherResult struct {
	win32.IInspectable
}

func (this *IProcessLauncherResult) Vtbl() *IProcessLauncherResultVtbl {
	return (*IProcessLauncherResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProcessLauncherResult) Get_ExitCode() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExitCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 33AB66E7-2D0E-448B-A6A0-C13C3836D09C
var IID_IProcessLauncherStatics = syscall.GUID{0x33AB66E7, 0x2D0E, 0x448B,
	[8]byte{0xA6, 0xA0, 0xC1, 0x3C, 0x38, 0x36, 0xD0, 0x9C}}

type IProcessLauncherStaticsInterface interface {
	win32.IInspectableInterface
	RunToCompletionAsync(fileName string, args string) *IAsyncOperation[*IProcessLauncherResult]
	RunToCompletionAsyncWithOptions(fileName string, args string, options *IProcessLauncherOptions) *IAsyncOperation[*IProcessLauncherResult]
}

type IProcessLauncherStaticsVtbl struct {
	win32.IInspectableVtbl
	RunToCompletionAsync            uintptr
	RunToCompletionAsyncWithOptions uintptr
}

type IProcessLauncherStatics struct {
	win32.IInspectable
}

func (this *IProcessLauncherStatics) Vtbl() *IProcessLauncherStaticsVtbl {
	return (*IProcessLauncherStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProcessLauncherStatics) RunToCompletionAsync(fileName string, args string) *IAsyncOperation[*IProcessLauncherResult] {
	var _result *IAsyncOperation[*IProcessLauncherResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RunToCompletionAsync, uintptr(unsafe.Pointer(this)), NewHStr(fileName).Ptr, NewHStr(args).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProcessLauncherStatics) RunToCompletionAsyncWithOptions(fileName string, args string, options *IProcessLauncherOptions) *IAsyncOperation[*IProcessLauncherResult] {
	var _result *IAsyncOperation[*IProcessLauncherResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RunToCompletionAsyncWithOptions, uintptr(unsafe.Pointer(this)), NewHStr(fileName).Ptr, NewHStr(args).Ptr, uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 087305A8-9B70-4782-8741-3A982B6CE5E4
var IID_IProcessMemoryReport = syscall.GUID{0x087305A8, 0x9B70, 0x4782,
	[8]byte{0x87, 0x41, 0x3A, 0x98, 0x2B, 0x6C, 0xE5, 0xE4}}

type IProcessMemoryReportInterface interface {
	win32.IInspectableInterface
	Get_PrivateWorkingSetUsage() uint64
	Get_TotalWorkingSetUsage() uint64
}

type IProcessMemoryReportVtbl struct {
	win32.IInspectableVtbl
	Get_PrivateWorkingSetUsage uintptr
	Get_TotalWorkingSetUsage   uintptr
}

type IProcessMemoryReport struct {
	win32.IInspectable
}

func (this *IProcessMemoryReport) Vtbl() *IProcessMemoryReportVtbl {
	return (*IProcessMemoryReportVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProcessMemoryReport) Get_PrivateWorkingSetUsage() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrivateWorkingSetUsage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProcessMemoryReport) Get_TotalWorkingSetUsage() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TotalWorkingSetUsage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D581293A-6DE9-4D28-9378-F86782E182BB
var IID_IProtocolForResultsOperation = syscall.GUID{0xD581293A, 0x6DE9, 0x4D28,
	[8]byte{0x93, 0x78, 0xF8, 0x67, 0x82, 0xE1, 0x82, 0xBB}}

type IProtocolForResultsOperationInterface interface {
	win32.IInspectableInterface
	ReportCompleted(data *IPropertySet)
}

type IProtocolForResultsOperationVtbl struct {
	win32.IInspectableVtbl
	ReportCompleted uintptr
}

type IProtocolForResultsOperation struct {
	win32.IInspectable
}

func (this *IProtocolForResultsOperation) Vtbl() *IProtocolForResultsOperationVtbl {
	return (*IProtocolForResultsOperationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProtocolForResultsOperation) ReportCompleted(data *IPropertySet) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReportCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(data)))
	_ = _hr
}

// 9E3A2788-2891-4CDF-A2D6-9DFF7D02E693
var IID_IRemoteLauncherOptions = syscall.GUID{0x9E3A2788, 0x2891, 0x4CDF,
	[8]byte{0xA2, 0xD6, 0x9D, 0xFF, 0x7D, 0x02, 0xE6, 0x93}}

type IRemoteLauncherOptionsInterface interface {
	win32.IInspectableInterface
	Get_FallbackUri() *IUriRuntimeClass
	Put_FallbackUri(value *IUriRuntimeClass)
	Get_PreferredAppIds() *IVector[string]
}

type IRemoteLauncherOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_FallbackUri     uintptr
	Put_FallbackUri     uintptr
	Get_PreferredAppIds uintptr
}

type IRemoteLauncherOptions struct {
	win32.IInspectable
}

func (this *IRemoteLauncherOptions) Vtbl() *IRemoteLauncherOptionsVtbl {
	return (*IRemoteLauncherOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRemoteLauncherOptions) Get_FallbackUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FallbackUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRemoteLauncherOptions) Put_FallbackUri(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FallbackUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IRemoteLauncherOptions) Get_PreferredAppIds() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PreferredAppIds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D7DB7A93-A30C-48B7-9F21-051026A4E517
var IID_IRemoteLauncherStatics = syscall.GUID{0xD7DB7A93, 0xA30C, 0x48B7,
	[8]byte{0x9F, 0x21, 0x05, 0x10, 0x26, 0xA4, 0xE5, 0x17}}

type IRemoteLauncherStaticsInterface interface {
	win32.IInspectableInterface
	LaunchUriAsync(remoteSystemConnectionRequest unsafe.Pointer, uri *IUriRuntimeClass) *IAsyncOperation[RemoteLaunchUriStatus]
	LaunchUriWithOptionsAsync(remoteSystemConnectionRequest unsafe.Pointer, uri *IUriRuntimeClass, options *IRemoteLauncherOptions) *IAsyncOperation[RemoteLaunchUriStatus]
	LaunchUriWithDataAsync(remoteSystemConnectionRequest unsafe.Pointer, uri *IUriRuntimeClass, options *IRemoteLauncherOptions, inputData *IPropertySet) *IAsyncOperation[RemoteLaunchUriStatus]
}

type IRemoteLauncherStaticsVtbl struct {
	win32.IInspectableVtbl
	LaunchUriAsync            uintptr
	LaunchUriWithOptionsAsync uintptr
	LaunchUriWithDataAsync    uintptr
}

type IRemoteLauncherStatics struct {
	win32.IInspectable
}

func (this *IRemoteLauncherStatics) Vtbl() *IRemoteLauncherStaticsVtbl {
	return (*IRemoteLauncherStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRemoteLauncherStatics) LaunchUriAsync(remoteSystemConnectionRequest unsafe.Pointer, uri *IUriRuntimeClass) *IAsyncOperation[RemoteLaunchUriStatus] {
	var _result *IAsyncOperation[RemoteLaunchUriStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchUriAsync, uintptr(unsafe.Pointer(this)), uintptr(remoteSystemConnectionRequest), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRemoteLauncherStatics) LaunchUriWithOptionsAsync(remoteSystemConnectionRequest unsafe.Pointer, uri *IUriRuntimeClass, options *IRemoteLauncherOptions) *IAsyncOperation[RemoteLaunchUriStatus] {
	var _result *IAsyncOperation[RemoteLaunchUriStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchUriWithOptionsAsync, uintptr(unsafe.Pointer(this)), uintptr(remoteSystemConnectionRequest), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRemoteLauncherStatics) LaunchUriWithDataAsync(remoteSystemConnectionRequest unsafe.Pointer, uri *IUriRuntimeClass, options *IRemoteLauncherOptions, inputData *IPropertySet) *IAsyncOperation[RemoteLaunchUriStatus] {
	var _result *IAsyncOperation[RemoteLaunchUriStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchUriWithDataAsync, uintptr(unsafe.Pointer(this)), uintptr(remoteSystemConnectionRequest), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(inputData)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 72E247ED-DD5B-4D6C-B1D0-C57A7BBB5F94
var IID_IShutdownManagerStatics = syscall.GUID{0x72E247ED, 0xDD5B, 0x4D6C,
	[8]byte{0xB1, 0xD0, 0xC5, 0x7A, 0x7B, 0xBB, 0x5F, 0x94}}

type IShutdownManagerStaticsInterface interface {
	win32.IInspectableInterface
	BeginShutdown(shutdownKind ShutdownKind, timeout TimeSpan)
	CancelShutdown()
}

type IShutdownManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	BeginShutdown  uintptr
	CancelShutdown uintptr
}

type IShutdownManagerStatics struct {
	win32.IInspectable
}

func (this *IShutdownManagerStatics) Vtbl() *IShutdownManagerStaticsVtbl {
	return (*IShutdownManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IShutdownManagerStatics) BeginShutdown(shutdownKind ShutdownKind, timeout TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().BeginShutdown, uintptr(unsafe.Pointer(this)), uintptr(shutdownKind), *(*uintptr)(unsafe.Pointer(&timeout)))
	_ = _hr
}

func (this *IShutdownManagerStatics) CancelShutdown() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CancelShutdown, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 0F69A02F-9C34-43C7-A8C3-70B30A7F7504
var IID_IShutdownManagerStatics2 = syscall.GUID{0x0F69A02F, 0x9C34, 0x43C7,
	[8]byte{0xA8, 0xC3, 0x70, 0xB3, 0x0A, 0x7F, 0x75, 0x04}}

type IShutdownManagerStatics2Interface interface {
	win32.IInspectableInterface
	IsPowerStateSupported(powerState PowerState) bool
	EnterPowerState(powerState PowerState)
	EnterPowerStateWithTimeSpan(powerState PowerState, wakeUpAfter TimeSpan)
}

type IShutdownManagerStatics2Vtbl struct {
	win32.IInspectableVtbl
	IsPowerStateSupported       uintptr
	EnterPowerState             uintptr
	EnterPowerStateWithTimeSpan uintptr
}

type IShutdownManagerStatics2 struct {
	win32.IInspectable
}

func (this *IShutdownManagerStatics2) Vtbl() *IShutdownManagerStatics2Vtbl {
	return (*IShutdownManagerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IShutdownManagerStatics2) IsPowerStateSupported(powerState PowerState) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsPowerStateSupported, uintptr(unsafe.Pointer(this)), uintptr(powerState), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IShutdownManagerStatics2) EnterPowerState(powerState PowerState) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnterPowerState, uintptr(unsafe.Pointer(this)), uintptr(powerState))
	_ = _hr
}

func (this *IShutdownManagerStatics2) EnterPowerStateWithTimeSpan(powerState PowerState, wakeUpAfter TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnterPowerStateWithTimeSpan, uintptr(unsafe.Pointer(this)), uintptr(powerState), *(*uintptr)(unsafe.Pointer(&wakeUpAfter)))
	_ = _hr
}

// 9B3B2BEA-A101-41AE-9FBD-028728BAB73D
var IID_ITimeZoneSettingsStatics = syscall.GUID{0x9B3B2BEA, 0xA101, 0x41AE,
	[8]byte{0x9F, 0xBD, 0x02, 0x87, 0x28, 0xBA, 0xB7, 0x3D}}

type ITimeZoneSettingsStaticsInterface interface {
	win32.IInspectableInterface
	Get_CurrentTimeZoneDisplayName() string
	Get_SupportedTimeZoneDisplayNames() *IVectorView[string]
	Get_CanChangeTimeZone() bool
	ChangeTimeZoneByDisplayName(timeZoneDisplayName string)
}

type ITimeZoneSettingsStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_CurrentTimeZoneDisplayName    uintptr
	Get_SupportedTimeZoneDisplayNames uintptr
	Get_CanChangeTimeZone             uintptr
	ChangeTimeZoneByDisplayName       uintptr
}

type ITimeZoneSettingsStatics struct {
	win32.IInspectable
}

func (this *ITimeZoneSettingsStatics) Vtbl() *ITimeZoneSettingsStaticsVtbl {
	return (*ITimeZoneSettingsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimeZoneSettingsStatics) Get_CurrentTimeZoneDisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentTimeZoneDisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ITimeZoneSettingsStatics) Get_SupportedTimeZoneDisplayNames() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedTimeZoneDisplayNames, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITimeZoneSettingsStatics) Get_CanChangeTimeZone() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanChangeTimeZone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimeZoneSettingsStatics) ChangeTimeZoneByDisplayName(timeZoneDisplayName string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ChangeTimeZoneByDisplayName, uintptr(unsafe.Pointer(this)), NewHStr(timeZoneDisplayName).Ptr)
	_ = _hr
}

// 555C0DB8-39A8-49FA-B4F6-A2C7FC2842EC
var IID_ITimeZoneSettingsStatics2 = syscall.GUID{0x555C0DB8, 0x39A8, 0x49FA,
	[8]byte{0xB4, 0xF6, 0xA2, 0xC7, 0xFC, 0x28, 0x42, 0xEC}}

type ITimeZoneSettingsStatics2Interface interface {
	win32.IInspectableInterface
	AutoUpdateTimeZoneAsync(timeout TimeSpan) *IAsyncOperation[AutoUpdateTimeZoneStatus]
}

type ITimeZoneSettingsStatics2Vtbl struct {
	win32.IInspectableVtbl
	AutoUpdateTimeZoneAsync uintptr
}

type ITimeZoneSettingsStatics2 struct {
	win32.IInspectable
}

func (this *ITimeZoneSettingsStatics2) Vtbl() *ITimeZoneSettingsStatics2Vtbl {
	return (*ITimeZoneSettingsStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimeZoneSettingsStatics2) AutoUpdateTimeZoneAsync(timeout TimeSpan) *IAsyncOperation[AutoUpdateTimeZoneStatus] {
	var _result *IAsyncOperation[AutoUpdateTimeZoneStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AutoUpdateTimeZoneAsync, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&timeout)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DF9A26C6-E746-4BCD-B5D4-120103C4209B
var IID_IUser = syscall.GUID{0xDF9A26C6, 0xE746, 0x4BCD,
	[8]byte{0xB5, 0xD4, 0x12, 0x01, 0x03, 0xC4, 0x20, 0x9B}}

type IUserInterface interface {
	win32.IInspectableInterface
	Get_NonRoamableId() string
	Get_AuthenticationStatus() UserAuthenticationStatus
	Get_Type() UserType
	GetPropertyAsync(value string) *IAsyncOperation[interface{}]
	GetPropertiesAsync(values *IVectorView[string]) *IAsyncOperation[*IPropertySet]
	GetPictureAsync(desiredSize UserPictureSize) *IAsyncOperation[*IRandomAccessStreamReference]
}

type IUserVtbl struct {
	win32.IInspectableVtbl
	Get_NonRoamableId        uintptr
	Get_AuthenticationStatus uintptr
	Get_Type                 uintptr
	GetPropertyAsync         uintptr
	GetPropertiesAsync       uintptr
	GetPictureAsync          uintptr
}

type IUser struct {
	win32.IInspectable
}

func (this *IUser) Vtbl() *IUserVtbl {
	return (*IUserVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUser) Get_NonRoamableId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NonRoamableId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUser) Get_AuthenticationStatus() UserAuthenticationStatus {
	var _result UserAuthenticationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AuthenticationStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUser) Get_Type() UserType {
	var _result UserType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Type, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUser) GetPropertyAsync(value string) *IAsyncOperation[interface{}] {
	var _result *IAsyncOperation[interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPropertyAsync, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUser) GetPropertiesAsync(values *IVectorView[string]) *IAsyncOperation[*IPropertySet] {
	var _result *IAsyncOperation[*IPropertySet]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPropertiesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(values)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUser) GetPictureAsync(desiredSize UserPictureSize) *IAsyncOperation[*IRandomAccessStreamReference] {
	var _result *IAsyncOperation[*IRandomAccessStreamReference]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPictureAsync, uintptr(unsafe.Pointer(this)), uintptr(desiredSize), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 88B59568-BB30-42FB-A270-E9902E40EFA7
var IID_IUserAuthenticationStatusChangeDeferral = syscall.GUID{0x88B59568, 0xBB30, 0x42FB,
	[8]byte{0xA2, 0x70, 0xE9, 0x90, 0x2E, 0x40, 0xEF, 0xA7}}

type IUserAuthenticationStatusChangeDeferralInterface interface {
	win32.IInspectableInterface
	Complete()
}

type IUserAuthenticationStatusChangeDeferralVtbl struct {
	win32.IInspectableVtbl
	Complete uintptr
}

type IUserAuthenticationStatusChangeDeferral struct {
	win32.IInspectable
}

func (this *IUserAuthenticationStatusChangeDeferral) Vtbl() *IUserAuthenticationStatusChangeDeferralVtbl {
	return (*IUserAuthenticationStatusChangeDeferralVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUserAuthenticationStatusChangeDeferral) Complete() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Complete, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 8C030F28-A711-4C1E-AB48-04179C15938F
var IID_IUserAuthenticationStatusChangingEventArgs = syscall.GUID{0x8C030F28, 0xA711, 0x4C1E,
	[8]byte{0xAB, 0x48, 0x04, 0x17, 0x9C, 0x15, 0x93, 0x8F}}

type IUserAuthenticationStatusChangingEventArgsInterface interface {
	win32.IInspectableInterface
	GetDeferral() *IUserAuthenticationStatusChangeDeferral
	Get_User() *IUser
	Get_NewStatus() UserAuthenticationStatus
	Get_CurrentStatus() UserAuthenticationStatus
}

type IUserAuthenticationStatusChangingEventArgsVtbl struct {
	win32.IInspectableVtbl
	GetDeferral       uintptr
	Get_User          uintptr
	Get_NewStatus     uintptr
	Get_CurrentStatus uintptr
}

type IUserAuthenticationStatusChangingEventArgs struct {
	win32.IInspectable
}

func (this *IUserAuthenticationStatusChangingEventArgs) Vtbl() *IUserAuthenticationStatusChangingEventArgsVtbl {
	return (*IUserAuthenticationStatusChangingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUserAuthenticationStatusChangingEventArgs) GetDeferral() *IUserAuthenticationStatusChangeDeferral {
	var _result *IUserAuthenticationStatusChangeDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserAuthenticationStatusChangingEventArgs) Get_User() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_User, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserAuthenticationStatusChangingEventArgs) Get_NewStatus() UserAuthenticationStatus {
	var _result UserAuthenticationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NewStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUserAuthenticationStatusChangingEventArgs) Get_CurrentStatus() UserAuthenticationStatus {
	var _result UserAuthenticationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 086459DC-18C6-48DB-BC99-724FB9203CCC
var IID_IUserChangedEventArgs = syscall.GUID{0x086459DC, 0x18C6, 0x48DB,
	[8]byte{0xBC, 0x99, 0x72, 0x4F, 0xB9, 0x20, 0x3C, 0xCC}}

type IUserChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_User() *IUser
}

type IUserChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_User uintptr
}

type IUserChangedEventArgs struct {
	win32.IInspectable
}

func (this *IUserChangedEventArgs) Vtbl() *IUserChangedEventArgsVtbl {
	return (*IUserChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUserChangedEventArgs) Get_User() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_User, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6B2CCB44-6F01-560C-97AD-FC7F32EC581F
var IID_IUserChangedEventArgs2 = syscall.GUID{0x6B2CCB44, 0x6F01, 0x560C,
	[8]byte{0x97, 0xAD, 0xFC, 0x7F, 0x32, 0xEC, 0x58, 0x1F}}

type IUserChangedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_ChangedPropertyKinds() *IVectorView[UserWatcherUpdateKind]
}

type IUserChangedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_ChangedPropertyKinds uintptr
}

type IUserChangedEventArgs2 struct {
	win32.IInspectable
}

func (this *IUserChangedEventArgs2) Vtbl() *IUserChangedEventArgs2Vtbl {
	return (*IUserChangedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUserChangedEventArgs2) Get_ChangedPropertyKinds() *IVectorView[UserWatcherUpdateKind] {
	var _result *IVectorView[UserWatcherUpdateKind]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ChangedPropertyKinds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BD1F6F6C-BB5D-4D7B-A5F0-C8CD11A38D42
var IID_IUserDeviceAssociationChangedEventArgs = syscall.GUID{0xBD1F6F6C, 0xBB5D, 0x4D7B,
	[8]byte{0xA5, 0xF0, 0xC8, 0xCD, 0x11, 0xA3, 0x8D, 0x42}}

type IUserDeviceAssociationChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_NewUser() *IUser
	Get_OldUser() *IUser
}

type IUserDeviceAssociationChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId uintptr
	Get_NewUser  uintptr
	Get_OldUser  uintptr
}

type IUserDeviceAssociationChangedEventArgs struct {
	win32.IInspectable
}

func (this *IUserDeviceAssociationChangedEventArgs) Vtbl() *IUserDeviceAssociationChangedEventArgsVtbl {
	return (*IUserDeviceAssociationChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUserDeviceAssociationChangedEventArgs) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUserDeviceAssociationChangedEventArgs) Get_NewUser() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NewUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserDeviceAssociationChangedEventArgs) Get_OldUser() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OldUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7E491E14-F85A-4C07-8DA9-7FE3D0542343
var IID_IUserDeviceAssociationStatics = syscall.GUID{0x7E491E14, 0xF85A, 0x4C07,
	[8]byte{0x8D, 0xA9, 0x7F, 0xE3, 0xD0, 0x54, 0x23, 0x43}}

type IUserDeviceAssociationStaticsInterface interface {
	win32.IInspectableInterface
	FindUserFromDeviceId(deviceId string) *IUser
	Add_UserDeviceAssociationChanged(handler EventHandler[*IUserDeviceAssociationChangedEventArgs]) EventRegistrationToken
	Remove_UserDeviceAssociationChanged(token EventRegistrationToken)
}

type IUserDeviceAssociationStaticsVtbl struct {
	win32.IInspectableVtbl
	FindUserFromDeviceId                uintptr
	Add_UserDeviceAssociationChanged    uintptr
	Remove_UserDeviceAssociationChanged uintptr
}

type IUserDeviceAssociationStatics struct {
	win32.IInspectable
}

func (this *IUserDeviceAssociationStatics) Vtbl() *IUserDeviceAssociationStaticsVtbl {
	return (*IUserDeviceAssociationStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUserDeviceAssociationStatics) FindUserFromDeviceId(deviceId string) *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindUserFromDeviceId, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserDeviceAssociationStatics) Add_UserDeviceAssociationChanged(handler EventHandler[*IUserDeviceAssociationChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_UserDeviceAssociationChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUserDeviceAssociationStatics) Remove_UserDeviceAssociationChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_UserDeviceAssociationChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 7D548008-F1E3-4A6C-8DDC-A9BB0F488AED
var IID_IUserPicker = syscall.GUID{0x7D548008, 0xF1E3, 0x4A6C,
	[8]byte{0x8D, 0xDC, 0xA9, 0xBB, 0x0F, 0x48, 0x8A, 0xED}}

type IUserPickerInterface interface {
	win32.IInspectableInterface
	Get_AllowGuestAccounts() bool
	Put_AllowGuestAccounts(value bool)
	Get_SuggestedSelectedUser() *IUser
	Put_SuggestedSelectedUser(value *IUser)
	PickSingleUserAsync() *IAsyncOperation[*IUser]
}

type IUserPickerVtbl struct {
	win32.IInspectableVtbl
	Get_AllowGuestAccounts    uintptr
	Put_AllowGuestAccounts    uintptr
	Get_SuggestedSelectedUser uintptr
	Put_SuggestedSelectedUser uintptr
	PickSingleUserAsync       uintptr
}

type IUserPicker struct {
	win32.IInspectable
}

func (this *IUserPicker) Vtbl() *IUserPickerVtbl {
	return (*IUserPickerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUserPicker) Get_AllowGuestAccounts() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllowGuestAccounts, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUserPicker) Put_AllowGuestAccounts(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AllowGuestAccounts, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IUserPicker) Get_SuggestedSelectedUser() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SuggestedSelectedUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserPicker) Put_SuggestedSelectedUser(value *IUser) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SuggestedSelectedUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IUserPicker) PickSingleUserAsync() *IAsyncOperation[*IUser] {
	var _result *IAsyncOperation[*IUser]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PickSingleUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DE3290DC-7E73-4DF6-A1AE-4D7ECA82B40D
var IID_IUserPickerStatics = syscall.GUID{0xDE3290DC, 0x7E73, 0x4DF6,
	[8]byte{0xA1, 0xAE, 0x4D, 0x7E, 0xCA, 0x82, 0xB4, 0x0D}}

type IUserPickerStaticsInterface interface {
	win32.IInspectableInterface
	IsSupported() bool
}

type IUserPickerStaticsVtbl struct {
	win32.IInspectableVtbl
	IsSupported uintptr
}

type IUserPickerStatics struct {
	win32.IInspectable
}

func (this *IUserPickerStatics) Vtbl() *IUserPickerStaticsVtbl {
	return (*IUserPickerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUserPickerStatics) IsSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 155EB23B-242A-45E0-A2E9-3171FC6A7FDD
var IID_IUserStatics = syscall.GUID{0x155EB23B, 0x242A, 0x45E0,
	[8]byte{0xA2, 0xE9, 0x31, 0x71, 0xFC, 0x6A, 0x7F, 0xDD}}

type IUserStaticsInterface interface {
	win32.IInspectableInterface
	CreateWatcher() *IUserWatcher
	FindAllAsync() *IAsyncOperation[*IVectorView[*IUser]]
	FindAllAsyncByType(type_ UserType) *IAsyncOperation[*IVectorView[*IUser]]
	FindAllAsyncByTypeAndStatus(type_ UserType, status UserAuthenticationStatus) *IAsyncOperation[*IVectorView[*IUser]]
	GetFromId(nonRoamableId string) *IUser
}

type IUserStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateWatcher               uintptr
	FindAllAsync                uintptr
	FindAllAsyncByType          uintptr
	FindAllAsyncByTypeAndStatus uintptr
	GetFromId                   uintptr
}

type IUserStatics struct {
	win32.IInspectable
}

func (this *IUserStatics) Vtbl() *IUserStaticsVtbl {
	return (*IUserStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUserStatics) CreateWatcher() *IUserWatcher {
	var _result *IUserWatcher
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWatcher, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserStatics) FindAllAsync() *IAsyncOperation[*IVectorView[*IUser]] {
	var _result *IAsyncOperation[*IVectorView[*IUser]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserStatics) FindAllAsyncByType(type_ UserType) *IAsyncOperation[*IVectorView[*IUser]] {
	var _result *IAsyncOperation[*IVectorView[*IUser]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllAsyncByType, uintptr(unsafe.Pointer(this)), uintptr(type_), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserStatics) FindAllAsyncByTypeAndStatus(type_ UserType, status UserAuthenticationStatus) *IAsyncOperation[*IVectorView[*IUser]] {
	var _result *IAsyncOperation[*IVectorView[*IUser]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllAsyncByTypeAndStatus, uintptr(unsafe.Pointer(this)), uintptr(type_), uintptr(status), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserStatics) GetFromId(nonRoamableId string) *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFromId, uintptr(unsafe.Pointer(this)), NewHStr(nonRoamableId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 74A37E11-2EB5-4487-B0D5-2C6790E013E9
var IID_IUserStatics2 = syscall.GUID{0x74A37E11, 0x2EB5, 0x4487,
	[8]byte{0xB0, 0xD5, 0x2C, 0x67, 0x90, 0xE0, 0x13, 0xE9}}

type IUserStatics2Interface interface {
	win32.IInspectableInterface
	GetDefault() *IUser
}

type IUserStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
}

type IUserStatics2 struct {
	win32.IInspectable
}

func (this *IUserStatics2) Vtbl() *IUserStatics2Vtbl {
	return (*IUserStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUserStatics2) GetDefault() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 155EB23B-242A-45E0-A2E9-3171FC6A7FBB
var IID_IUserWatcher = syscall.GUID{0x155EB23B, 0x242A, 0x45E0,
	[8]byte{0xA2, 0xE9, 0x31, 0x71, 0xFC, 0x6A, 0x7F, 0xBB}}

type IUserWatcherInterface interface {
	win32.IInspectableInterface
	Get_Status() UserWatcherStatus
	Start()
	Stop()
	Add_Added(handler TypedEventHandler[*IUserWatcher, *IUserChangedEventArgs]) EventRegistrationToken
	Remove_Added(token EventRegistrationToken)
	Add_Removed(handler TypedEventHandler[*IUserWatcher, *IUserChangedEventArgs]) EventRegistrationToken
	Remove_Removed(token EventRegistrationToken)
	Add_Updated(handler TypedEventHandler[*IUserWatcher, *IUserChangedEventArgs]) EventRegistrationToken
	Remove_Updated(token EventRegistrationToken)
	Add_AuthenticationStatusChanged(handler TypedEventHandler[*IUserWatcher, *IUserChangedEventArgs]) EventRegistrationToken
	Remove_AuthenticationStatusChanged(token EventRegistrationToken)
	Add_AuthenticationStatusChanging(handler TypedEventHandler[*IUserWatcher, *IUserAuthenticationStatusChangingEventArgs]) EventRegistrationToken
	Remove_AuthenticationStatusChanging(token EventRegistrationToken)
	Add_EnumerationCompleted(handler TypedEventHandler[*IUserWatcher, interface{}]) EventRegistrationToken
	Remove_EnumerationCompleted(token EventRegistrationToken)
	Add_Stopped(handler TypedEventHandler[*IUserWatcher, interface{}]) EventRegistrationToken
	Remove_Stopped(token EventRegistrationToken)
}

type IUserWatcherVtbl struct {
	win32.IInspectableVtbl
	Get_Status                          uintptr
	Start                               uintptr
	Stop                                uintptr
	Add_Added                           uintptr
	Remove_Added                        uintptr
	Add_Removed                         uintptr
	Remove_Removed                      uintptr
	Add_Updated                         uintptr
	Remove_Updated                      uintptr
	Add_AuthenticationStatusChanged     uintptr
	Remove_AuthenticationStatusChanged  uintptr
	Add_AuthenticationStatusChanging    uintptr
	Remove_AuthenticationStatusChanging uintptr
	Add_EnumerationCompleted            uintptr
	Remove_EnumerationCompleted         uintptr
	Add_Stopped                         uintptr
	Remove_Stopped                      uintptr
}

type IUserWatcher struct {
	win32.IInspectable
}

func (this *IUserWatcher) Vtbl() *IUserWatcherVtbl {
	return (*IUserWatcherVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUserWatcher) Get_Status() UserWatcherStatus {
	var _result UserWatcherStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUserWatcher) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IUserWatcher) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IUserWatcher) Add_Added(handler TypedEventHandler[*IUserWatcher, *IUserChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Added, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUserWatcher) Remove_Added(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Added, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IUserWatcher) Add_Removed(handler TypedEventHandler[*IUserWatcher, *IUserChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Removed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUserWatcher) Remove_Removed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Removed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IUserWatcher) Add_Updated(handler TypedEventHandler[*IUserWatcher, *IUserChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Updated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUserWatcher) Remove_Updated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Updated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IUserWatcher) Add_AuthenticationStatusChanged(handler TypedEventHandler[*IUserWatcher, *IUserChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AuthenticationStatusChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUserWatcher) Remove_AuthenticationStatusChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AuthenticationStatusChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IUserWatcher) Add_AuthenticationStatusChanging(handler TypedEventHandler[*IUserWatcher, *IUserAuthenticationStatusChangingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AuthenticationStatusChanging, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUserWatcher) Remove_AuthenticationStatusChanging(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AuthenticationStatusChanging, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IUserWatcher) Add_EnumerationCompleted(handler TypedEventHandler[*IUserWatcher, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_EnumerationCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUserWatcher) Remove_EnumerationCompleted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_EnumerationCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IUserWatcher) Add_Stopped(handler TypedEventHandler[*IUserWatcher, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Stopped, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUserWatcher) Remove_Stopped(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Stopped, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// classes

type AppActivationResult struct {
	RtClass
	*IAppActivationResult
}

type AppDiagnosticInfo struct {
	RtClass
	*IAppDiagnosticInfo
}

func NewIAppDiagnosticInfoStatics2() *IAppDiagnosticInfoStatics2 {
	var p *IAppDiagnosticInfoStatics2
	hs := NewHStr("Windows.System.AppDiagnosticInfo")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAppDiagnosticInfoStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIAppDiagnosticInfoStatics() *IAppDiagnosticInfoStatics {
	var p *IAppDiagnosticInfoStatics
	hs := NewHStr("Windows.System.AppDiagnosticInfo")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAppDiagnosticInfoStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type AppDiagnosticInfoWatcher struct {
	RtClass
	*IAppDiagnosticInfoWatcher
}

type AppDiagnosticInfoWatcherEventArgs struct {
	RtClass
	*IAppDiagnosticInfoWatcherEventArgs
}

type AppExecutionStateChangeResult struct {
	RtClass
	*IAppExecutionStateChangeResult
}

type AppResourceGroupBackgroundTaskReport struct {
	RtClass
	*IAppResourceGroupBackgroundTaskReport
}

type AppResourceGroupInfo struct {
	RtClass
	*IAppResourceGroupInfo
}

type AppResourceGroupInfoWatcher struct {
	RtClass
	*IAppResourceGroupInfoWatcher
}

type AppResourceGroupInfoWatcherEventArgs struct {
	RtClass
	*IAppResourceGroupInfoWatcherEventArgs
}

type AppResourceGroupInfoWatcherExecutionStateChangedEventArgs struct {
	RtClass
	*IAppResourceGroupInfoWatcherExecutionStateChangedEventArgs
}

type AppResourceGroupMemoryReport struct {
	RtClass
	*IAppResourceGroupMemoryReport
}

type AppResourceGroupStateReport struct {
	RtClass
	*IAppResourceGroupStateReport
}

type DateTimeSettings struct {
	RtClass
}

func NewIDateTimeSettingsStatics() *IDateTimeSettingsStatics {
	var p *IDateTimeSettingsStatics
	hs := NewHStr("Windows.System.DateTimeSettings")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDateTimeSettingsStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type DispatcherQueue struct {
	RtClass
	*IDispatcherQueue
}

func NewIDispatcherQueueStatics() *IDispatcherQueueStatics {
	var p *IDispatcherQueueStatics
	hs := NewHStr("Windows.System.DispatcherQueue")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDispatcherQueueStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type DispatcherQueueController struct {
	RtClass
	*IDispatcherQueueController
}

func NewIDispatcherQueueControllerStatics() *IDispatcherQueueControllerStatics {
	var p *IDispatcherQueueControllerStatics
	hs := NewHStr("Windows.System.DispatcherQueueController")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDispatcherQueueControllerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type DispatcherQueueShutdownStartingEventArgs struct {
	RtClass
	*IDispatcherQueueShutdownStartingEventArgs
}

type DispatcherQueueTimer struct {
	RtClass
	*IDispatcherQueueTimer
}

type KnownUserProperties struct {
	RtClass
}

func NewIKnownUserPropertiesStatics() *IKnownUserPropertiesStatics {
	var p *IKnownUserPropertiesStatics
	hs := NewHStr("Windows.System.KnownUserProperties")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKnownUserPropertiesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type ProcessLauncher struct {
	RtClass
}

func NewIProcessLauncherStatics() *IProcessLauncherStatics {
	var p *IProcessLauncherStatics
	hs := NewHStr("Windows.System.ProcessLauncher")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IProcessLauncherStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type ProcessLauncherOptions struct {
	RtClass
	*IProcessLauncherOptions
}

func NewProcessLauncherOptions() *ProcessLauncherOptions {
	hs := NewHStr("Windows.System.ProcessLauncherOptions")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &ProcessLauncherOptions{
		RtClass:                 RtClass{PInspect: p},
		IProcessLauncherOptions: (*IProcessLauncherOptions)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type ProcessLauncherResult struct {
	RtClass
	*IProcessLauncherResult
}

type ShutdownManager struct {
	RtClass
}

func NewIShutdownManagerStatics2() *IShutdownManagerStatics2 {
	var p *IShutdownManagerStatics2
	hs := NewHStr("Windows.System.ShutdownManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IShutdownManagerStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIShutdownManagerStatics() *IShutdownManagerStatics {
	var p *IShutdownManagerStatics
	hs := NewHStr("Windows.System.ShutdownManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IShutdownManagerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type User struct {
	RtClass
	*IUser
}

func NewIUserStatics2() *IUserStatics2 {
	var p *IUserStatics2
	hs := NewHStr("Windows.System.User")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IUserStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIUserStatics() *IUserStatics {
	var p *IUserStatics
	hs := NewHStr("Windows.System.User")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IUserStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type UserAuthenticationStatusChangeDeferral struct {
	RtClass
	*IUserAuthenticationStatusChangeDeferral
}

type UserAuthenticationStatusChangingEventArgs struct {
	RtClass
	*IUserAuthenticationStatusChangingEventArgs
}

type UserChangedEventArgs struct {
	RtClass
	*IUserChangedEventArgs
}

type UserDeviceAssociation struct {
	RtClass
}

func NewIUserDeviceAssociationStatics() *IUserDeviceAssociationStatics {
	var p *IUserDeviceAssociationStatics
	hs := NewHStr("Windows.System.UserDeviceAssociation")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IUserDeviceAssociationStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type UserDeviceAssociationChangedEventArgs struct {
	RtClass
	*IUserDeviceAssociationChangedEventArgs
}

type UserPicker struct {
	RtClass
	*IUserPicker
}

func NewUserPicker() *UserPicker {
	hs := NewHStr("Windows.System.UserPicker")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &UserPicker{
		RtClass:     RtClass{PInspect: p},
		IUserPicker: (*IUserPicker)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewIUserPickerStatics() *IUserPickerStatics {
	var p *IUserPickerStatics
	hs := NewHStr("Windows.System.UserPicker")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IUserPickerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type UserWatcher struct {
	RtClass
	*IUserWatcher
}
