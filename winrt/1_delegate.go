package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

type IDelegate struct {
	win32.IUnknown
}

type IDelegateInterface interface {
	win32.IUnknownInterface
	Invoke(arg1, arg2, arg3 uintptr) com.Error
}

type IDelegateImpl struct {
	com.IUnknownImpl
	RealObject IDelegateInterface
	iid        *syscall.GUID
}

func (this *IDelegateImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(IDelegateInterface)
}

func (this *IDelegateImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	iid := *riid
	if this.iid == nil && !IsSysGuid(iid) {
		this.iid = &iid
	}
	//guess generic iid..
	if this.iid != nil && iid == *this.iid || iid == win32.IID_IAgileObject {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *IDelegateImpl) Invoke(arg1, arg2, arg3 uintptr) com.Error {
	var ret com.Error
	return ret
}

type IDelegateVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type IDelegateComObj struct {
	com.IUnknownComObj
}

func (this *IDelegateComObj) impl() IDelegateInterface {
	return this.Impl().(IDelegateInterface)
}

func (this *IDelegateComObj) Invoke(arg1, arg2, arg3 uintptr) uintptr {
	return (uintptr)(this.impl().Invoke(arg1, arg2, arg3))
}

var _pIDelegateVtbl *IDelegateVtbl

func (this *IDelegateComObj) BuildVtbl(lock bool) *IDelegateVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pIDelegateVtbl != nil {
		return _pIDelegateVtbl
	}
	_pIDelegateVtbl = &IDelegateVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:       syscall.NewCallback((*IDelegateComObj).Invoke),
	}
	return _pIDelegateVtbl
}

func (this *IDelegateComObj) IDelegate() *IDelegate {
	return (*IDelegate)(unsafe.Pointer(this))
}

func (this *IDelegateComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewIDelegateComObj(impl IDelegateInterface) *IDelegateComObj {
	tid := win32.GetCurrentThreadId()
	println(tid)
	comObj := com.NewComObj[IDelegateComObj](impl)
	com.AddToScope(comObj)
	return comObj
}

func NewIDelegate(impl IDelegateInterface) *IDelegate {
	return NewIDelegateComObj(impl).IDelegate()
}

type IDelegateByFuncImpl struct {
	IDelegateImpl
	handlerFunc func(arg1, arg2, arg3 uintptr) com.Error
}

func (this *IDelegateByFuncImpl) Invoke(arg1, arg2, arg3 uintptr) com.Error {
	return this.handlerFunc(arg1, arg2, arg3)
}

func NewIDelegateByFunc(handlerFunc func(arg1, arg2, arg3 uintptr) com.Error) *IDelegate {
	impl := &IDelegateByFuncImpl{handlerFunc: handlerFunc}
	return NewIDelegateComObj(impl).IDelegate()
}

func NewNoArgFuncDelegate(anyFunc func() com.Error) *IDelegate {
	return NewIDelegateByFunc(func(arg1, arg2, arg3 uintptr) com.Error {
		return anyFunc()
	})
}

func NewOneArgFuncDelegate[TArg any](anyFunc func(arg TArg) com.Error) *IDelegate {
	return NewIDelegateByFunc(func(arg1, arg2, arg3 uintptr) com.Error {
		return anyFunc(*(*TArg)(unsafe.Pointer(&arg1)))
	})
}

func NewTwoArgFuncDelegate[TArg1 any, TArg2 any](anyFunc func(arg1 TArg1, arg2 TArg2) com.Error) *IDelegate {
	return NewIDelegateByFunc(func(arg1, arg2, arg3 uintptr) com.Error {
		return anyFunc(*(*TArg1)(unsafe.Pointer(&arg1)), *(*TArg2)(unsafe.Pointer(&arg2)))
	})
}

func NewThreeArgFuncDelegate[TArg1 any, TArg2 any, TArg3 any](anyFunc func(arg1 TArg1, arg2 TArg2, arg3 TArg3) com.Error) *IDelegate {
	return NewIDelegateByFunc(func(arg1, arg2, arg3 uintptr) com.Error {
		return anyFunc(*(*TArg1)(unsafe.Pointer(&arg1)), *(*TArg2)(unsafe.Pointer(&arg2)), *(*TArg3)(unsafe.Pointer(&arg3)))
	})
}
