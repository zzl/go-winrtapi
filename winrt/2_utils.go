package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"reflect"
	"syscall"
	"unsafe"
)

func IsSysGuid(guid syscall.GUID) bool {
	if guid.Data2 == 0 && guid.Data3 == 0 &&
		guid.Data4 == [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46} {
		return true
	}
	if guid == win32.IID_INoMarshal ||
		guid == win32.IID_IAgileObject {
		return true
	}
	return false
}

func PostProcessGenericResult[TResult any](result TResult) TResult {
	switch v := any(result).(type) {
	case string:
		hs := *(*win32.HSTRING)(unsafe.Pointer(&result))
		result = any(HStringToStrAndFree(hs)).(TResult)
	case win32.IUnknownObject:
		com.AddToScope(v)
		break
	}
	return result
}

const PtrSize = unsafe.Sizeof(uintptr(0))

func CastArgToPointer[TArg any](arg TArg) unsafe.Pointer {
	switch v := any(arg).(type) {
	case string:
		return unsafe.Pointer(NewHStr(v).Ptr)
	case win32.IUnknownObject:
		return unsafe.Pointer(v.GetIUnknown())
	default:
		size := reflect.TypeOf(arg).Size()
		if size > PtrSize {
			return unsafe.Pointer(&arg)
		} else {
			return unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&arg)))
		}
	}
}
