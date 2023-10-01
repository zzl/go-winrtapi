package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"log"
	"syscall"
	"unsafe"
)

// interfaces

// 2D75F35A-70A7-4395-BA2D-CEF0B18399F9
var IID_ICompositorController = syscall.GUID{0x2D75F35A, 0x70A7, 0x4395,
	[8]byte{0xBA, 0x2D, 0xCE, 0xF0, 0xB1, 0x83, 0x99, 0xF9}}

type ICompositorControllerInterface interface {
	win32.IInspectableInterface
	Get_Compositor() *ICompositor
	Commit()
	EnsurePreviousCommitCompletedAsync() *IAsyncAction
	Add_CommitNeeded(handler TypedEventHandler[*ICompositorController, interface{}]) EventRegistrationToken
	Remove_CommitNeeded(token EventRegistrationToken)
}

type ICompositorControllerVtbl struct {
	win32.IInspectableVtbl
	Get_Compositor                     uintptr
	Commit                             uintptr
	EnsurePreviousCommitCompletedAsync uintptr
	Add_CommitNeeded                   uintptr
	Remove_CommitNeeded                uintptr
}

type ICompositorController struct {
	win32.IInspectable
}

func (this *ICompositorController) Vtbl() *ICompositorControllerVtbl {
	return (*ICompositorControllerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositorController) Get_Compositor() *ICompositor {
	var _result *ICompositor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Compositor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositorController) Commit() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Commit, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ICompositorController) EnsurePreviousCommitCompletedAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnsurePreviousCommitCompletedAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositorController) Add_CommitNeeded(handler TypedEventHandler[*ICompositorController, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_CommitNeeded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositorController) Remove_CommitNeeded(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_CommitNeeded, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// classes

type CompositorController struct {
	RtClass
	*ICompositorController
}

func NewCompositorController() *CompositorController {
	hs := NewHStr("Windows.UI.Composition.Core.CompositorController")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &CompositorController{
		RtClass:               RtClass{PInspect: p},
		ICompositorController: (*ICompositorController)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}
