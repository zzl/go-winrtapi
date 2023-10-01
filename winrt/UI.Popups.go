package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
// flags
type MessageDialogOptions uint32

const (
	MessageDialogOptions_None                      MessageDialogOptions = 0
	MessageDialogOptions_AcceptUserInputAfterDelay MessageDialogOptions = 1
)

// enum
type Placement int32

const (
	Placement_Default Placement = 0
	Placement_Above   Placement = 1
	Placement_Below   Placement = 2
	Placement_Left    Placement = 3
	Placement_Right   Placement = 4
)

// func types

// DAF77A4F-C27A-4298-9AC6-2922C45E7DA6
type UICommandInvokedHandler func(command *IUICommand) com.Error

// interfaces

// 33F59B01-5325-43AB-9AB3-BDAE440E4121
var IID_IMessageDialog = syscall.GUID{0x33F59B01, 0x5325, 0x43AB,
	[8]byte{0x9A, 0xB3, 0xBD, 0xAE, 0x44, 0x0E, 0x41, 0x21}}

type IMessageDialogInterface interface {
	win32.IInspectableInterface
	Get_Title() string
	Put_Title(value string)
	Get_Commands() *IVector[*IUICommand]
	Get_DefaultCommandIndex() uint32
	Put_DefaultCommandIndex(value uint32)
	Get_CancelCommandIndex() uint32
	Put_CancelCommandIndex(value uint32)
	Get_Content() string
	Put_Content(value string)
	ShowAsync() *IAsyncOperation[*IUICommand]
	Get_Options() MessageDialogOptions
	Put_Options(value MessageDialogOptions)
}

type IMessageDialogVtbl struct {
	win32.IInspectableVtbl
	Get_Title               uintptr
	Put_Title               uintptr
	Get_Commands            uintptr
	Get_DefaultCommandIndex uintptr
	Put_DefaultCommandIndex uintptr
	Get_CancelCommandIndex  uintptr
	Put_CancelCommandIndex  uintptr
	Get_Content             uintptr
	Put_Content             uintptr
	ShowAsync               uintptr
	Get_Options             uintptr
	Put_Options             uintptr
}

type IMessageDialog struct {
	win32.IInspectable
}

func (this *IMessageDialog) Vtbl() *IMessageDialogVtbl {
	return (*IMessageDialogVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMessageDialog) Get_Title() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMessageDialog) Put_Title(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Title, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMessageDialog) Get_Commands() *IVector[*IUICommand] {
	var _result *IVector[*IUICommand]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Commands, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMessageDialog) Get_DefaultCommandIndex() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DefaultCommandIndex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMessageDialog) Put_DefaultCommandIndex(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DefaultCommandIndex, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMessageDialog) Get_CancelCommandIndex() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CancelCommandIndex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMessageDialog) Put_CancelCommandIndex(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CancelCommandIndex, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMessageDialog) Get_Content() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Content, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMessageDialog) Put_Content(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Content, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMessageDialog) ShowAsync() *IAsyncOperation[*IUICommand] {
	var _result *IAsyncOperation[*IUICommand]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMessageDialog) Get_Options() MessageDialogOptions {
	var _result MessageDialogOptions
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Options, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMessageDialog) Put_Options(value MessageDialogOptions) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Options, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 2D161777-A66F-4EA5-BB87-793FFA4941F2
var IID_IMessageDialogFactory = syscall.GUID{0x2D161777, 0xA66F, 0x4EA5,
	[8]byte{0xBB, 0x87, 0x79, 0x3F, 0xFA, 0x49, 0x41, 0xF2}}

type IMessageDialogFactoryInterface interface {
	win32.IInspectableInterface
	Create(content string) *IMessageDialog
	CreateWithTitle(content string, title string) *IMessageDialog
}

type IMessageDialogFactoryVtbl struct {
	win32.IInspectableVtbl
	Create          uintptr
	CreateWithTitle uintptr
}

type IMessageDialogFactory struct {
	win32.IInspectable
}

func (this *IMessageDialogFactory) Vtbl() *IMessageDialogFactoryVtbl {
	return (*IMessageDialogFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMessageDialogFactory) Create(content string) *IMessageDialog {
	var _result *IMessageDialog
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(content).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMessageDialogFactory) CreateWithTitle(content string, title string) *IMessageDialog {
	var _result *IMessageDialog
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithTitle, uintptr(unsafe.Pointer(this)), NewHStr(content).Ptr, NewHStr(title).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4E9BC6DC-880D-47FC-A0A1-72B639E62559
var IID_IPopupMenu = syscall.GUID{0x4E9BC6DC, 0x880D, 0x47FC,
	[8]byte{0xA0, 0xA1, 0x72, 0xB6, 0x39, 0xE6, 0x25, 0x59}}

type IPopupMenuInterface interface {
	win32.IInspectableInterface
	Get_Commands() *IVector[*IUICommand]
	ShowAsync(invocationPoint Point) *IAsyncOperation[*IUICommand]
	ShowAsyncWithRect(selection Rect) *IAsyncOperation[*IUICommand]
	ShowAsyncWithRectAndPlacement(selection Rect, preferredPlacement Placement) *IAsyncOperation[*IUICommand]
}

type IPopupMenuVtbl struct {
	win32.IInspectableVtbl
	Get_Commands                  uintptr
	ShowAsync                     uintptr
	ShowAsyncWithRect             uintptr
	ShowAsyncWithRectAndPlacement uintptr
}

type IPopupMenu struct {
	win32.IInspectable
}

func (this *IPopupMenu) Vtbl() *IPopupMenuVtbl {
	return (*IPopupMenuVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPopupMenu) Get_Commands() *IVector[*IUICommand] {
	var _result *IVector[*IUICommand]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Commands, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPopupMenu) ShowAsync(invocationPoint Point) *IAsyncOperation[*IUICommand] {
	var _result *IAsyncOperation[*IUICommand]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowAsync, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&invocationPoint)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPopupMenu) ShowAsyncWithRect(selection Rect) *IAsyncOperation[*IUICommand] {
	var _result *IAsyncOperation[*IUICommand]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowAsyncWithRect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&selection)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPopupMenu) ShowAsyncWithRectAndPlacement(selection Rect, preferredPlacement Placement) *IAsyncOperation[*IUICommand] {
	var _result *IAsyncOperation[*IUICommand]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowAsyncWithRectAndPlacement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&selection)), uintptr(preferredPlacement), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4FF93A75-4145-47FF-AC7F-DFF1C1FA5B0F
var IID_IUICommand = syscall.GUID{0x4FF93A75, 0x4145, 0x47FF,
	[8]byte{0xAC, 0x7F, 0xDF, 0xF1, 0xC1, 0xFA, 0x5B, 0x0F}}

type IUICommandInterface interface {
	win32.IInspectableInterface
	Get_Label() string
	Put_Label(value string)
	Get_Invoked() UICommandInvokedHandler
	Put_Invoked(value UICommandInvokedHandler)
	Get_Id() interface{}
	Put_Id(value interface{})
}

type IUICommandVtbl struct {
	win32.IInspectableVtbl
	Get_Label   uintptr
	Put_Label   uintptr
	Get_Invoked uintptr
	Put_Invoked uintptr
	Get_Id      uintptr
	Put_Id      uintptr
}

type IUICommand struct {
	win32.IInspectable
}

func (this *IUICommand) Vtbl() *IUICommandVtbl {
	return (*IUICommandVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUICommand) Get_Label() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Label, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUICommand) Put_Label(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Label, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IUICommand) Get_Invoked() UICommandInvokedHandler {
	var _result UICommandInvokedHandler
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Invoked, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUICommand) Put_Invoked(value UICommandInvokedHandler) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Invoked, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewOneArgFuncDelegate(value))))
	_ = _hr
}

func (this *IUICommand) Get_Id() interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUICommand) Put_Id(value interface{}) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Id, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

// A21A8189-26B0-4676-AE94-54041BC125E8
var IID_IUICommandFactory = syscall.GUID{0xA21A8189, 0x26B0, 0x4676,
	[8]byte{0xAE, 0x94, 0x54, 0x04, 0x1B, 0xC1, 0x25, 0xE8}}

type IUICommandFactoryInterface interface {
	win32.IInspectableInterface
	Create(label string) *IUICommand
	CreateWithHandler(label string, action UICommandInvokedHandler) *IUICommand
	CreateWithHandlerAndId(label string, action UICommandInvokedHandler, commandId interface{}) *IUICommand
}

type IUICommandFactoryVtbl struct {
	win32.IInspectableVtbl
	Create                 uintptr
	CreateWithHandler      uintptr
	CreateWithHandlerAndId uintptr
}

type IUICommandFactory struct {
	win32.IInspectable
}

func (this *IUICommandFactory) Vtbl() *IUICommandFactoryVtbl {
	return (*IUICommandFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUICommandFactory) Create(label string) *IUICommand {
	var _result *IUICommand
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(label).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUICommandFactory) CreateWithHandler(label string, action UICommandInvokedHandler) *IUICommand {
	var _result *IUICommand
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithHandler, uintptr(unsafe.Pointer(this)), NewHStr(label).Ptr, uintptr(unsafe.Pointer(NewOneArgFuncDelegate(action))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUICommandFactory) CreateWithHandlerAndId(label string, action UICommandInvokedHandler, commandId interface{}) *IUICommand {
	var _result *IUICommand
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithHandlerAndId, uintptr(unsafe.Pointer(this)), NewHStr(label).Ptr, uintptr(unsafe.Pointer(NewOneArgFuncDelegate(action))), *(*uintptr)(unsafe.Pointer(&commandId)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}
