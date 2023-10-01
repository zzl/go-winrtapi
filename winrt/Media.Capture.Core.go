package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// interfaces

// D1EB4C5C-1B53-4E4A-8B5C-DB7887AC949B
var IID_IVariablePhotoCapturedEventArgs = syscall.GUID{0xD1EB4C5C, 0x1B53, 0x4E4A,
	[8]byte{0x8B, 0x5C, 0xDB, 0x78, 0x87, 0xAC, 0x94, 0x9B}}

type IVariablePhotoCapturedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Frame() *ICapturedFrame
	Get_CaptureTimeOffset() TimeSpan
	Get_UsedFrameControllerIndex() *IReference[uint32]
	Get_CapturedFrameControlValues() *ICapturedFrameControlValues
}

type IVariablePhotoCapturedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Frame                      uintptr
	Get_CaptureTimeOffset          uintptr
	Get_UsedFrameControllerIndex   uintptr
	Get_CapturedFrameControlValues uintptr
}

type IVariablePhotoCapturedEventArgs struct {
	win32.IInspectable
}

func (this *IVariablePhotoCapturedEventArgs) Vtbl() *IVariablePhotoCapturedEventArgsVtbl {
	return (*IVariablePhotoCapturedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVariablePhotoCapturedEventArgs) Get_Frame() *ICapturedFrame {
	var _result *ICapturedFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Frame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVariablePhotoCapturedEventArgs) Get_CaptureTimeOffset() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CaptureTimeOffset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVariablePhotoCapturedEventArgs) Get_UsedFrameControllerIndex() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UsedFrameControllerIndex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVariablePhotoCapturedEventArgs) Get_CapturedFrameControlValues() *ICapturedFrameControlValues {
	var _result *ICapturedFrameControlValues
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CapturedFrameControlValues, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D0112D1D-031E-4041-A6D6-BD742476A8EE
var IID_IVariablePhotoSequenceCapture = syscall.GUID{0xD0112D1D, 0x031E, 0x4041,
	[8]byte{0xA6, 0xD6, 0xBD, 0x74, 0x24, 0x76, 0xA8, 0xEE}}

type IVariablePhotoSequenceCaptureInterface interface {
	win32.IInspectableInterface
	StartAsync() *IAsyncAction
	StopAsync() *IAsyncAction
	FinishAsync() *IAsyncAction
	Add_PhotoCaptured(handler TypedEventHandler[*IVariablePhotoSequenceCapture, *IVariablePhotoCapturedEventArgs]) EventRegistrationToken
	Remove_PhotoCaptured(token EventRegistrationToken)
	Add_Stopped(handler TypedEventHandler[*IVariablePhotoSequenceCapture, interface{}]) EventRegistrationToken
	Remove_Stopped(token EventRegistrationToken)
}

type IVariablePhotoSequenceCaptureVtbl struct {
	win32.IInspectableVtbl
	StartAsync           uintptr
	StopAsync            uintptr
	FinishAsync          uintptr
	Add_PhotoCaptured    uintptr
	Remove_PhotoCaptured uintptr
	Add_Stopped          uintptr
	Remove_Stopped       uintptr
}

type IVariablePhotoSequenceCapture struct {
	win32.IInspectable
}

func (this *IVariablePhotoSequenceCapture) Vtbl() *IVariablePhotoSequenceCaptureVtbl {
	return (*IVariablePhotoSequenceCaptureVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVariablePhotoSequenceCapture) StartAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVariablePhotoSequenceCapture) StopAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVariablePhotoSequenceCapture) FinishAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FinishAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVariablePhotoSequenceCapture) Add_PhotoCaptured(handler TypedEventHandler[*IVariablePhotoSequenceCapture, *IVariablePhotoCapturedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PhotoCaptured, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVariablePhotoSequenceCapture) Remove_PhotoCaptured(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PhotoCaptured, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IVariablePhotoSequenceCapture) Add_Stopped(handler TypedEventHandler[*IVariablePhotoSequenceCapture, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Stopped, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVariablePhotoSequenceCapture) Remove_Stopped(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Stopped, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// FE2C62BC-50B0-43E3-917C-E3B92798942F
var IID_IVariablePhotoSequenceCapture2 = syscall.GUID{0xFE2C62BC, 0x50B0, 0x43E3,
	[8]byte{0x91, 0x7C, 0xE3, 0xB9, 0x27, 0x98, 0x94, 0x2F}}

type IVariablePhotoSequenceCapture2Interface interface {
	win32.IInspectableInterface
	UpdateSettingsAsync() *IAsyncAction
}

type IVariablePhotoSequenceCapture2Vtbl struct {
	win32.IInspectableVtbl
	UpdateSettingsAsync uintptr
}

type IVariablePhotoSequenceCapture2 struct {
	win32.IInspectable
}

func (this *IVariablePhotoSequenceCapture2) Vtbl() *IVariablePhotoSequenceCapture2Vtbl {
	return (*IVariablePhotoSequenceCapture2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVariablePhotoSequenceCapture2) UpdateSettingsAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UpdateSettingsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type VariablePhotoCapturedEventArgs struct {
	RtClass
	*IVariablePhotoCapturedEventArgs
}

type VariablePhotoSequenceCapture struct {
	RtClass
	*IVariablePhotoSequenceCapture
}
