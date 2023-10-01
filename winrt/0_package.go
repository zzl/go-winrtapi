package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"log"
	"reflect"
	"runtime"
	"syscall"
	"unsafe"
)

var GlobalScope *com.Scope

func Initialize() Initialized {
	runtime.LockOSThread()
	com.InitializeContext()
	win32.RoInitialize(win32.RO_INIT_SINGLETHREADED)
	GlobalScope = com.NewScope()
	return Initialized{}
}

func InitializeMt() Initialized {
	runtime.LockOSThread()
	com.InitializeContext()
	win32.RoInitialize(win32.RO_INIT_MULTITHREADED)
	GlobalScope = com.NewScope()
	return Initialized{}
}

type Initialized struct {
}

func (me Initialized) Uninitialize() {
	Uninitialize()
}

func Uninitialize() {
	GlobalScope.Leave()
	win32.RoUninitialize()
	com.UninitializeContext()
	runtime.UnlockOSThread()
}

type HStr struct {
	Ptr win32.HSTRING
}

func NewHStr(s string) *HStr {
	wsz, _ := syscall.UTF16FromString(s)
	hString := &HStr{}
	hr := win32.WindowsCreateString(&wsz[0], uint32(len(wsz)-1), &hString.Ptr)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	runtime.SetFinalizer(hString, (*HStr).Dispose)
	return hString
}

func (this *HStr) Dispose() {
	win32.WindowsDeleteString(this.Ptr)
	this.Ptr = 0
}

func HStringToStrAndFree(hs win32.HSTRING) string {
	var cc uint32
	pwsz := win32.WindowsGetStringRawBuffer(hs, &cc)
	wsz := unsafe.Slice(pwsz, cc)
	str := syscall.UTF16ToString(wsz)
	win32.WindowsDeleteString(hs)
	return str
}

func PPInspectable(pInterface interface{}) **win32.IInspectable {
	return (**win32.IInspectable)(reflect.ValueOf(pInterface).UnsafePointer())
}

type RtClass struct {
	PInspect *win32.IInspectable
}
