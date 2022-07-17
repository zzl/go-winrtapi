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
type SettingsEdgeLocation int32

const (
	SettingsEdgeLocation_Right SettingsEdgeLocation = 0
	SettingsEdgeLocation_Left  SettingsEdgeLocation = 1
)

// enum
// flags
type SupportedWebAccountActions uint32

const (
	SupportedWebAccountActions_None        SupportedWebAccountActions = 0
	SupportedWebAccountActions_Reconnect   SupportedWebAccountActions = 1
	SupportedWebAccountActions_Remove      SupportedWebAccountActions = 2
	SupportedWebAccountActions_ViewDetails SupportedWebAccountActions = 4
	SupportedWebAccountActions_Manage      SupportedWebAccountActions = 8
	SupportedWebAccountActions_More        SupportedWebAccountActions = 16
)

// enum
type WebAccountAction int32

const (
	WebAccountAction_Reconnect   WebAccountAction = 0
	WebAccountAction_Remove      WebAccountAction = 1
	WebAccountAction_ViewDetails WebAccountAction = 2
	WebAccountAction_Manage      WebAccountAction = 3
	WebAccountAction_More        WebAccountAction = 4
)

// structs

type ApplicationsSettingsContract struct {
}

// func types

//61C0E185-0977-4678-B4E2-98727AFBEED9
type CredentialCommandCredentialDeletedHandler func(command *ICredentialCommand) com.Error

//1EE6E459-1705-4A9A-B599-A0C3D6921973
type WebAccountCommandInvokedHandler func(command *IWebAccountCommand, args *IWebAccountInvokedArgs) com.Error

//B7DE5527-4C8F-42DD-84DA-5EC493ABDB9A
type WebAccountProviderCommandInvokedHandler func(command *IWebAccountProviderCommand) com.Error

// interfaces

// 81EA942C-4F09-4406-A538-838D9B14B7E6
var IID_IAccountsSettingsPane = syscall.GUID{0x81EA942C, 0x4F09, 0x4406,
	[8]byte{0xA5, 0x38, 0x83, 0x8D, 0x9B, 0x14, 0xB7, 0xE6}}

type IAccountsSettingsPaneInterface interface {
	win32.IInspectableInterface
	Add_AccountCommandsRequested(handler TypedEventHandler[*IAccountsSettingsPane, *IAccountsSettingsPaneCommandsRequestedEventArgs]) EventRegistrationToken
	Remove_AccountCommandsRequested(cookie EventRegistrationToken)
}

type IAccountsSettingsPaneVtbl struct {
	win32.IInspectableVtbl
	Add_AccountCommandsRequested    uintptr
	Remove_AccountCommandsRequested uintptr
}

type IAccountsSettingsPane struct {
	win32.IInspectable
}

func (this *IAccountsSettingsPane) Vtbl() *IAccountsSettingsPaneVtbl {
	return (*IAccountsSettingsPaneVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccountsSettingsPane) Add_AccountCommandsRequested(handler TypedEventHandler[*IAccountsSettingsPane, *IAccountsSettingsPaneCommandsRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AccountCommandsRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAccountsSettingsPane) Remove_AccountCommandsRequested(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AccountCommandsRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

// 3B68C099-DB19-45D0-9ABF-95D3773C9330
var IID_IAccountsSettingsPaneCommandsRequestedEventArgs = syscall.GUID{0x3B68C099, 0xDB19, 0x45D0,
	[8]byte{0x9A, 0xBF, 0x95, 0xD3, 0x77, 0x3C, 0x93, 0x30}}

type IAccountsSettingsPaneCommandsRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_WebAccountProviderCommands() *IVector[*IWebAccountProviderCommand]
	Get_WebAccountCommands() *IVector[*IWebAccountCommand]
	Get_CredentialCommands() *IVector[*ICredentialCommand]
	Get_Commands() *IVector[*IUICommand]
	Get_HeaderText() string
	Put_HeaderText(value string)
	GetDeferral() *IAccountsSettingsPaneEventDeferral
}

type IAccountsSettingsPaneCommandsRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_WebAccountProviderCommands uintptr
	Get_WebAccountCommands         uintptr
	Get_CredentialCommands         uintptr
	Get_Commands                   uintptr
	Get_HeaderText                 uintptr
	Put_HeaderText                 uintptr
	GetDeferral                    uintptr
}

type IAccountsSettingsPaneCommandsRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IAccountsSettingsPaneCommandsRequestedEventArgs) Vtbl() *IAccountsSettingsPaneCommandsRequestedEventArgsVtbl {
	return (*IAccountsSettingsPaneCommandsRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccountsSettingsPaneCommandsRequestedEventArgs) Get_WebAccountProviderCommands() *IVector[*IWebAccountProviderCommand] {
	var _result *IVector[*IWebAccountProviderCommand]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WebAccountProviderCommands, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAccountsSettingsPaneCommandsRequestedEventArgs) Get_WebAccountCommands() *IVector[*IWebAccountCommand] {
	var _result *IVector[*IWebAccountCommand]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WebAccountCommands, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAccountsSettingsPaneCommandsRequestedEventArgs) Get_CredentialCommands() *IVector[*ICredentialCommand] {
	var _result *IVector[*ICredentialCommand]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CredentialCommands, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAccountsSettingsPaneCommandsRequestedEventArgs) Get_Commands() *IVector[*IUICommand] {
	var _result *IVector[*IUICommand]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Commands, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAccountsSettingsPaneCommandsRequestedEventArgs) Get_HeaderText() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HeaderText, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAccountsSettingsPaneCommandsRequestedEventArgs) Put_HeaderText(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_HeaderText, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IAccountsSettingsPaneCommandsRequestedEventArgs) GetDeferral() *IAccountsSettingsPaneEventDeferral {
	var _result *IAccountsSettingsPaneEventDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 362F7BAD-4E37-4967-8C40-E78EE7A1E5BB
var IID_IAccountsSettingsPaneCommandsRequestedEventArgs2 = syscall.GUID{0x362F7BAD, 0x4E37, 0x4967,
	[8]byte{0x8C, 0x40, 0xE7, 0x8E, 0xE7, 0xA1, 0xE5, 0xBB}}

type IAccountsSettingsPaneCommandsRequestedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_User() *IUser
}

type IAccountsSettingsPaneCommandsRequestedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_User uintptr
}

type IAccountsSettingsPaneCommandsRequestedEventArgs2 struct {
	win32.IInspectable
}

func (this *IAccountsSettingsPaneCommandsRequestedEventArgs2) Vtbl() *IAccountsSettingsPaneCommandsRequestedEventArgs2Vtbl {
	return (*IAccountsSettingsPaneCommandsRequestedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccountsSettingsPaneCommandsRequestedEventArgs2) Get_User() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_User, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CBF25D3F-E5BA-40EF-93DA-65E096E5FB04
var IID_IAccountsSettingsPaneEventDeferral = syscall.GUID{0xCBF25D3F, 0xE5BA, 0x40EF,
	[8]byte{0x93, 0xDA, 0x65, 0xE0, 0x96, 0xE5, 0xFB, 0x04}}

type IAccountsSettingsPaneEventDeferralInterface interface {
	win32.IInspectableInterface
	Complete()
}

type IAccountsSettingsPaneEventDeferralVtbl struct {
	win32.IInspectableVtbl
	Complete uintptr
}

type IAccountsSettingsPaneEventDeferral struct {
	win32.IInspectable
}

func (this *IAccountsSettingsPaneEventDeferral) Vtbl() *IAccountsSettingsPaneEventDeferralVtbl {
	return (*IAccountsSettingsPaneEventDeferralVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccountsSettingsPaneEventDeferral) Complete() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Complete, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 561F8B60-B0EC-4150-A8DC-208EE44B068A
var IID_IAccountsSettingsPaneStatics = syscall.GUID{0x561F8B60, 0xB0EC, 0x4150,
	[8]byte{0xA8, 0xDC, 0x20, 0x8E, 0xE4, 0x4B, 0x06, 0x8A}}

type IAccountsSettingsPaneStaticsInterface interface {
	win32.IInspectableInterface
	GetForCurrentView() *IAccountsSettingsPane
	Show()
}

type IAccountsSettingsPaneStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForCurrentView uintptr
	Show              uintptr
}

type IAccountsSettingsPaneStatics struct {
	win32.IInspectable
}

func (this *IAccountsSettingsPaneStatics) Vtbl() *IAccountsSettingsPaneStaticsVtbl {
	return (*IAccountsSettingsPaneStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccountsSettingsPaneStatics) GetForCurrentView() *IAccountsSettingsPane {
	var _result *IAccountsSettingsPane
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAccountsSettingsPaneStatics) Show() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Show, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// D21DF7C2-CE0D-484F-B8E8-E823C215765E
var IID_IAccountsSettingsPaneStatics2 = syscall.GUID{0xD21DF7C2, 0xCE0D, 0x484F,
	[8]byte{0xB8, 0xE8, 0xE8, 0x23, 0xC2, 0x15, 0x76, 0x5E}}

type IAccountsSettingsPaneStatics2Interface interface {
	win32.IInspectableInterface
	ShowManageAccountsAsync() *IAsyncAction
	ShowAddAccountAsync() *IAsyncAction
}

type IAccountsSettingsPaneStatics2Vtbl struct {
	win32.IInspectableVtbl
	ShowManageAccountsAsync uintptr
	ShowAddAccountAsync     uintptr
}

type IAccountsSettingsPaneStatics2 struct {
	win32.IInspectable
}

func (this *IAccountsSettingsPaneStatics2) Vtbl() *IAccountsSettingsPaneStatics2Vtbl {
	return (*IAccountsSettingsPaneStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccountsSettingsPaneStatics2) ShowManageAccountsAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowManageAccountsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAccountsSettingsPaneStatics2) ShowAddAccountAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowAddAccountAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 08410458-A2BA-4C6F-B4AC-48F514331216
var IID_IAccountsSettingsPaneStatics3 = syscall.GUID{0x08410458, 0xA2BA, 0x4C6F,
	[8]byte{0xB4, 0xAC, 0x48, 0xF5, 0x14, 0x33, 0x12, 0x16}}

type IAccountsSettingsPaneStatics3Interface interface {
	win32.IInspectableInterface
	ShowManageAccountsForUserAsync(user *IUser) *IAsyncAction
	ShowAddAccountForUserAsync(user *IUser) *IAsyncAction
}

type IAccountsSettingsPaneStatics3Vtbl struct {
	win32.IInspectableVtbl
	ShowManageAccountsForUserAsync uintptr
	ShowAddAccountForUserAsync     uintptr
}

type IAccountsSettingsPaneStatics3 struct {
	win32.IInspectable
}

func (this *IAccountsSettingsPaneStatics3) Vtbl() *IAccountsSettingsPaneStatics3Vtbl {
	return (*IAccountsSettingsPaneStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccountsSettingsPaneStatics3) ShowManageAccountsForUserAsync(user *IUser) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowManageAccountsForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAccountsSettingsPaneStatics3) ShowAddAccountForUserAsync(user *IUser) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowAddAccountForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A5F665E6-6143-4A7A-A971-B017BA978CE2
var IID_ICredentialCommand = syscall.GUID{0xA5F665E6, 0x6143, 0x4A7A,
	[8]byte{0xA9, 0x71, 0xB0, 0x17, 0xBA, 0x97, 0x8C, 0xE2}}

type ICredentialCommandInterface interface {
	win32.IInspectableInterface
	Get_PasswordCredential() *IPasswordCredential
	Get_CredentialDeleted() CredentialCommandCredentialDeletedHandler
}

type ICredentialCommandVtbl struct {
	win32.IInspectableVtbl
	Get_PasswordCredential uintptr
	Get_CredentialDeleted  uintptr
}

type ICredentialCommand struct {
	win32.IInspectable
}

func (this *ICredentialCommand) Vtbl() *ICredentialCommandVtbl {
	return (*ICredentialCommandVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICredentialCommand) Get_PasswordCredential() *IPasswordCredential {
	var _result *IPasswordCredential
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PasswordCredential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICredentialCommand) Get_CredentialDeleted() CredentialCommandCredentialDeletedHandler {
	var _result CredentialCommandCredentialDeletedHandler
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CredentialDeleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 27E88C17-BC3E-4B80-9495-4ED720E48A91
var IID_ICredentialCommandFactory = syscall.GUID{0x27E88C17, 0xBC3E, 0x4B80,
	[8]byte{0x94, 0x95, 0x4E, 0xD7, 0x20, 0xE4, 0x8A, 0x91}}

type ICredentialCommandFactoryInterface interface {
	win32.IInspectableInterface
	CreateCredentialCommand(passwordCredential *IPasswordCredential) *ICredentialCommand
	CreateCredentialCommandWithHandler(passwordCredential *IPasswordCredential, deleted CredentialCommandCredentialDeletedHandler) *ICredentialCommand
}

type ICredentialCommandFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateCredentialCommand            uintptr
	CreateCredentialCommandWithHandler uintptr
}

type ICredentialCommandFactory struct {
	win32.IInspectable
}

func (this *ICredentialCommandFactory) Vtbl() *ICredentialCommandFactoryVtbl {
	return (*ICredentialCommandFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICredentialCommandFactory) CreateCredentialCommand(passwordCredential *IPasswordCredential) *ICredentialCommand {
	var _result *ICredentialCommand
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCredentialCommand, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(passwordCredential)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICredentialCommandFactory) CreateCredentialCommandWithHandler(passwordCredential *IPasswordCredential, deleted CredentialCommandCredentialDeletedHandler) *ICredentialCommand {
	var _result *ICredentialCommand
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCredentialCommandWithHandler, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(passwordCredential)), uintptr(unsafe.Pointer(NewOneArgFuncDelegate(deleted))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 68E15B33-1C83-433A-AA5A-CEEEA5BD4764
var IID_ISettingsCommandFactory = syscall.GUID{0x68E15B33, 0x1C83, 0x433A,
	[8]byte{0xAA, 0x5A, 0xCE, 0xEE, 0xA5, 0xBD, 0x47, 0x64}}

type ISettingsCommandFactoryInterface interface {
	win32.IInspectableInterface
	CreateSettingsCommand(settingsCommandId interface{}, label string, handler UICommandInvokedHandler) *IUICommand
}

type ISettingsCommandFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateSettingsCommand uintptr
}

type ISettingsCommandFactory struct {
	win32.IInspectable
}

func (this *ISettingsCommandFactory) Vtbl() *ISettingsCommandFactoryVtbl {
	return (*ISettingsCommandFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISettingsCommandFactory) CreateSettingsCommand(settingsCommandId interface{}, label string, handler UICommandInvokedHandler) *IUICommand {
	var _result *IUICommand
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSettingsCommand, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&settingsCommandId)), NewHStr(label).Ptr, uintptr(unsafe.Pointer(NewOneArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 749AE954-2F69-4B17-8ABA-D05CE5778E46
var IID_ISettingsCommandStatics = syscall.GUID{0x749AE954, 0x2F69, 0x4B17,
	[8]byte{0x8A, 0xBA, 0xD0, 0x5C, 0xE5, 0x77, 0x8E, 0x46}}

type ISettingsCommandStaticsInterface interface {
	win32.IInspectableInterface
	Get_AccountsCommand() *IUICommand
}

type ISettingsCommandStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_AccountsCommand uintptr
}

type ISettingsCommandStatics struct {
	win32.IInspectable
}

func (this *ISettingsCommandStatics) Vtbl() *ISettingsCommandStaticsVtbl {
	return (*ISettingsCommandStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISettingsCommandStatics) Get_AccountsCommand() *IUICommand {
	var _result *IUICommand
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AccountsCommand, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B1CD0932-4570-4C69-8D38-89446561ACE0
var IID_ISettingsPane = syscall.GUID{0xB1CD0932, 0x4570, 0x4C69,
	[8]byte{0x8D, 0x38, 0x89, 0x44, 0x65, 0x61, 0xAC, 0xE0}}

type ISettingsPaneInterface interface {
	win32.IInspectableInterface
	Add_CommandsRequested(handler TypedEventHandler[*ISettingsPane, *ISettingsPaneCommandsRequestedEventArgs]) EventRegistrationToken
	Remove_CommandsRequested(cookie EventRegistrationToken)
}

type ISettingsPaneVtbl struct {
	win32.IInspectableVtbl
	Add_CommandsRequested    uintptr
	Remove_CommandsRequested uintptr
}

type ISettingsPane struct {
	win32.IInspectable
}

func (this *ISettingsPane) Vtbl() *ISettingsPaneVtbl {
	return (*ISettingsPaneVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISettingsPane) Add_CommandsRequested(handler TypedEventHandler[*ISettingsPane, *ISettingsPaneCommandsRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_CommandsRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISettingsPane) Remove_CommandsRequested(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_CommandsRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

// 44DF23AE-5D6E-4068-A168-F47643182114
var IID_ISettingsPaneCommandsRequest = syscall.GUID{0x44DF23AE, 0x5D6E, 0x4068,
	[8]byte{0xA1, 0x68, 0xF4, 0x76, 0x43, 0x18, 0x21, 0x14}}

type ISettingsPaneCommandsRequestInterface interface {
	win32.IInspectableInterface
	Get_ApplicationCommands() *IVector[*IUICommand]
}

type ISettingsPaneCommandsRequestVtbl struct {
	win32.IInspectableVtbl
	Get_ApplicationCommands uintptr
}

type ISettingsPaneCommandsRequest struct {
	win32.IInspectable
}

func (this *ISettingsPaneCommandsRequest) Vtbl() *ISettingsPaneCommandsRequestVtbl {
	return (*ISettingsPaneCommandsRequestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISettingsPaneCommandsRequest) Get_ApplicationCommands() *IVector[*IUICommand] {
	var _result *IVector[*IUICommand]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ApplicationCommands, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 205F5D24-1B48-4629-A6CA-2FDFEDAFB75D
var IID_ISettingsPaneCommandsRequestedEventArgs = syscall.GUID{0x205F5D24, 0x1B48, 0x4629,
	[8]byte{0xA6, 0xCA, 0x2F, 0xDF, 0xED, 0xAF, 0xB7, 0x5D}}

type ISettingsPaneCommandsRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Request() *ISettingsPaneCommandsRequest
}

type ISettingsPaneCommandsRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Request uintptr
}

type ISettingsPaneCommandsRequestedEventArgs struct {
	win32.IInspectable
}

func (this *ISettingsPaneCommandsRequestedEventArgs) Vtbl() *ISettingsPaneCommandsRequestedEventArgsVtbl {
	return (*ISettingsPaneCommandsRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISettingsPaneCommandsRequestedEventArgs) Get_Request() *ISettingsPaneCommandsRequest {
	var _result *ISettingsPaneCommandsRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Request, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1C6A52C5-FF19-471B-BA6B-F8F35694AD9A
var IID_ISettingsPaneStatics = syscall.GUID{0x1C6A52C5, 0xFF19, 0x471B,
	[8]byte{0xBA, 0x6B, 0xF8, 0xF3, 0x56, 0x94, 0xAD, 0x9A}}

type ISettingsPaneStaticsInterface interface {
	win32.IInspectableInterface
	GetForCurrentView() *ISettingsPane
	Show()
	Get_Edge() SettingsEdgeLocation
}

type ISettingsPaneStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForCurrentView uintptr
	Show              uintptr
	Get_Edge          uintptr
}

type ISettingsPaneStatics struct {
	win32.IInspectable
}

func (this *ISettingsPaneStatics) Vtbl() *ISettingsPaneStaticsVtbl {
	return (*ISettingsPaneStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISettingsPaneStatics) GetForCurrentView() *ISettingsPane {
	var _result *ISettingsPane
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISettingsPaneStatics) Show() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Show, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ISettingsPaneStatics) Get_Edge() SettingsEdgeLocation {
	var _result SettingsEdgeLocation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Edge, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// CAA39398-9CFA-4246-B0C4-A913A3896541
var IID_IWebAccountCommand = syscall.GUID{0xCAA39398, 0x9CFA, 0x4246,
	[8]byte{0xB0, 0xC4, 0xA9, 0x13, 0xA3, 0x89, 0x65, 0x41}}

type IWebAccountCommandInterface interface {
	win32.IInspectableInterface
	Get_WebAccount() *IWebAccount
	Get_Invoked() WebAccountCommandInvokedHandler
	Get_Actions() SupportedWebAccountActions
}

type IWebAccountCommandVtbl struct {
	win32.IInspectableVtbl
	Get_WebAccount uintptr
	Get_Invoked    uintptr
	Get_Actions    uintptr
}

type IWebAccountCommand struct {
	win32.IInspectable
}

func (this *IWebAccountCommand) Vtbl() *IWebAccountCommandVtbl {
	return (*IWebAccountCommandVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountCommand) Get_WebAccount() *IWebAccount {
	var _result *IWebAccount
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WebAccount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountCommand) Get_Invoked() WebAccountCommandInvokedHandler {
	var _result WebAccountCommandInvokedHandler
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Invoked, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWebAccountCommand) Get_Actions() SupportedWebAccountActions {
	var _result SupportedWebAccountActions
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Actions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// BFA6CDFF-2F2D-42F5-81DE-1D56BAFC496D
var IID_IWebAccountCommandFactory = syscall.GUID{0xBFA6CDFF, 0x2F2D, 0x42F5,
	[8]byte{0x81, 0xDE, 0x1D, 0x56, 0xBA, 0xFC, 0x49, 0x6D}}

type IWebAccountCommandFactoryInterface interface {
	win32.IInspectableInterface
	CreateWebAccountCommand(webAccount *IWebAccount, invoked WebAccountCommandInvokedHandler, actions SupportedWebAccountActions) *IWebAccountCommand
}

type IWebAccountCommandFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateWebAccountCommand uintptr
}

type IWebAccountCommandFactory struct {
	win32.IInspectable
}

func (this *IWebAccountCommandFactory) Vtbl() *IWebAccountCommandFactoryVtbl {
	return (*IWebAccountCommandFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountCommandFactory) CreateWebAccountCommand(webAccount *IWebAccount, invoked WebAccountCommandInvokedHandler, actions SupportedWebAccountActions) *IWebAccountCommand {
	var _result *IWebAccountCommand
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWebAccountCommand, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(webAccount)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(invoked))), uintptr(actions), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E7ABCC40-A1D8-4C5D-9A7F-1D34B2F90AD2
var IID_IWebAccountInvokedArgs = syscall.GUID{0xE7ABCC40, 0xA1D8, 0x4C5D,
	[8]byte{0x9A, 0x7F, 0x1D, 0x34, 0xB2, 0xF9, 0x0A, 0xD2}}

type IWebAccountInvokedArgsInterface interface {
	win32.IInspectableInterface
	Get_Action() WebAccountAction
}

type IWebAccountInvokedArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Action uintptr
}

type IWebAccountInvokedArgs struct {
	win32.IInspectable
}

func (this *IWebAccountInvokedArgs) Vtbl() *IWebAccountInvokedArgsVtbl {
	return (*IWebAccountInvokedArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountInvokedArgs) Get_Action() WebAccountAction {
	var _result WebAccountAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Action, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D69BDD9A-A0A6-4E9B-88DC-C71E757A3501
var IID_IWebAccountProviderCommand = syscall.GUID{0xD69BDD9A, 0xA0A6, 0x4E9B,
	[8]byte{0x88, 0xDC, 0xC7, 0x1E, 0x75, 0x7A, 0x35, 0x01}}

type IWebAccountProviderCommandInterface interface {
	win32.IInspectableInterface
	Get_WebAccountProvider() *IWebAccountProvider
	Get_Invoked() WebAccountProviderCommandInvokedHandler
}

type IWebAccountProviderCommandVtbl struct {
	win32.IInspectableVtbl
	Get_WebAccountProvider uintptr
	Get_Invoked            uintptr
}

type IWebAccountProviderCommand struct {
	win32.IInspectable
}

func (this *IWebAccountProviderCommand) Vtbl() *IWebAccountProviderCommandVtbl {
	return (*IWebAccountProviderCommandVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountProviderCommand) Get_WebAccountProvider() *IWebAccountProvider {
	var _result *IWebAccountProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WebAccountProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountProviderCommand) Get_Invoked() WebAccountProviderCommandInvokedHandler {
	var _result WebAccountProviderCommandInvokedHandler
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Invoked, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D5658A1B-B176-4776-8469-A9D3FF0B3F59
var IID_IWebAccountProviderCommandFactory = syscall.GUID{0xD5658A1B, 0xB176, 0x4776,
	[8]byte{0x84, 0x69, 0xA9, 0xD3, 0xFF, 0x0B, 0x3F, 0x59}}

type IWebAccountProviderCommandFactoryInterface interface {
	win32.IInspectableInterface
	CreateWebAccountProviderCommand(webAccountProvider *IWebAccountProvider, invoked WebAccountProviderCommandInvokedHandler) *IWebAccountProviderCommand
}

type IWebAccountProviderCommandFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateWebAccountProviderCommand uintptr
}

type IWebAccountProviderCommandFactory struct {
	win32.IInspectable
}

func (this *IWebAccountProviderCommandFactory) Vtbl() *IWebAccountProviderCommandFactoryVtbl {
	return (*IWebAccountProviderCommandFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountProviderCommandFactory) CreateWebAccountProviderCommand(webAccountProvider *IWebAccountProvider, invoked WebAccountProviderCommandInvokedHandler) *IWebAccountProviderCommand {
	var _result *IWebAccountProviderCommand
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWebAccountProviderCommand, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(webAccountProvider)), uintptr(unsafe.Pointer(NewOneArgFuncDelegate(invoked))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type AccountsSettingsPane struct {
	RtClass
	*IAccountsSettingsPane
}

func NewIAccountsSettingsPaneStatics2() *IAccountsSettingsPaneStatics2 {
	var p *IAccountsSettingsPaneStatics2
	hs := NewHStr("Windows.UI.ApplicationSettings.AccountsSettingsPane")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAccountsSettingsPaneStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIAccountsSettingsPaneStatics() *IAccountsSettingsPaneStatics {
	var p *IAccountsSettingsPaneStatics
	hs := NewHStr("Windows.UI.ApplicationSettings.AccountsSettingsPane")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAccountsSettingsPaneStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIAccountsSettingsPaneStatics3() *IAccountsSettingsPaneStatics3 {
	var p *IAccountsSettingsPaneStatics3
	hs := NewHStr("Windows.UI.ApplicationSettings.AccountsSettingsPane")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAccountsSettingsPaneStatics3, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type AccountsSettingsPaneCommandsRequestedEventArgs struct {
	RtClass
	*IAccountsSettingsPaneCommandsRequestedEventArgs
}

type AccountsSettingsPaneEventDeferral struct {
	RtClass
	*IAccountsSettingsPaneEventDeferral
}

type SettingsCommand struct {
	RtClass
	*IUICommand
}

func NewSettingsCommand_CreateSettingsCommand(settingsCommandId interface{}, label string, handler UICommandInvokedHandler) *SettingsCommand {
	hs := NewHStr("Windows.UI.ApplicationSettings.SettingsCommand")
	var pFac *ISettingsCommandFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISettingsCommandFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IUICommand
	p = pFac.CreateSettingsCommand(settingsCommandId, label, handler)
	result := &SettingsCommand{
		RtClass:    RtClass{PInspect: &p.IInspectable},
		IUICommand: p,
	}
	com.AddToScope(result)
	return result
}

func NewISettingsCommandStatics() *ISettingsCommandStatics {
	var p *ISettingsCommandStatics
	hs := NewHStr("Windows.UI.ApplicationSettings.SettingsCommand")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISettingsCommandStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type WebAccountCommand struct {
	RtClass
	*IWebAccountCommand
}

func NewWebAccountCommand_CreateWebAccountCommand(webAccount *IWebAccount, invoked WebAccountCommandInvokedHandler, actions SupportedWebAccountActions) *WebAccountCommand {
	hs := NewHStr("Windows.UI.ApplicationSettings.WebAccountCommand")
	var pFac *IWebAccountCommandFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebAccountCommandFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IWebAccountCommand
	p = pFac.CreateWebAccountCommand(webAccount, invoked, actions)
	result := &WebAccountCommand{
		RtClass:            RtClass{PInspect: &p.IInspectable},
		IWebAccountCommand: p,
	}
	com.AddToScope(result)
	return result
}

type WebAccountInvokedArgs struct {
	RtClass
	*IWebAccountInvokedArgs
}

type WebAccountProviderCommand struct {
	RtClass
	*IWebAccountProviderCommand
}

func NewWebAccountProviderCommand_CreateWebAccountProviderCommand(webAccountProvider *IWebAccountProvider, invoked WebAccountProviderCommandInvokedHandler) *WebAccountProviderCommand {
	hs := NewHStr("Windows.UI.ApplicationSettings.WebAccountProviderCommand")
	var pFac *IWebAccountProviderCommandFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebAccountProviderCommandFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IWebAccountProviderCommand
	p = pFac.CreateWebAccountProviderCommand(webAccountProvider, invoked)
	result := &WebAccountProviderCommand{
		RtClass:                    RtClass{PInspect: &p.IInspectable},
		IWebAccountProviderCommand: p,
	}
	com.AddToScope(result)
	return result
}
