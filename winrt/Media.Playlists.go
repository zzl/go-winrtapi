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
type PlaylistFormat int32

const (
	PlaylistFormat_WindowsMedia PlaylistFormat = 0
	PlaylistFormat_Zune         PlaylistFormat = 1
	PlaylistFormat_M3u          PlaylistFormat = 2
)

// structs

type PlaylistsContract struct {
}

// interfaces

// 803736F5-CF44-4D97-83B3-7A089E9AB663
var IID_IPlaylist = syscall.GUID{0x803736F5, 0xCF44, 0x4D97,
	[8]byte{0x83, 0xB3, 0x7A, 0x08, 0x9E, 0x9A, 0xB6, 0x63}}

type IPlaylistInterface interface {
	win32.IInspectableInterface
	Get_Files() *IVector[*IStorageFile]
	SaveAsync() *IAsyncAction
	SaveAsAsync(saveLocation *IStorageFolder, desiredName string, option NameCollisionOption) *IAsyncOperation[*IStorageFile]
	SaveAsWithFormatAsync(saveLocation *IStorageFolder, desiredName string, option NameCollisionOption, playlistFormat PlaylistFormat) *IAsyncOperation[*IStorageFile]
}

type IPlaylistVtbl struct {
	win32.IInspectableVtbl
	Get_Files             uintptr
	SaveAsync             uintptr
	SaveAsAsync           uintptr
	SaveAsWithFormatAsync uintptr
}

type IPlaylist struct {
	win32.IInspectable
}

func (this *IPlaylist) Vtbl() *IPlaylistVtbl {
	return (*IPlaylistVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPlaylist) Get_Files() *IVector[*IStorageFile] {
	var _result *IVector[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Files, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPlaylist) SaveAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SaveAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPlaylist) SaveAsAsync(saveLocation *IStorageFolder, desiredName string, option NameCollisionOption) *IAsyncOperation[*IStorageFile] {
	var _result *IAsyncOperation[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SaveAsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(saveLocation)), NewHStr(desiredName).Ptr, uintptr(option), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPlaylist) SaveAsWithFormatAsync(saveLocation *IStorageFolder, desiredName string, option NameCollisionOption, playlistFormat PlaylistFormat) *IAsyncOperation[*IStorageFile] {
	var _result *IAsyncOperation[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SaveAsWithFormatAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(saveLocation)), NewHStr(desiredName).Ptr, uintptr(option), uintptr(playlistFormat), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C5C331CD-81F9-4FF3-95B9-70B6FF046B68
var IID_IPlaylistStatics = syscall.GUID{0xC5C331CD, 0x81F9, 0x4FF3,
	[8]byte{0x95, 0xB9, 0x70, 0xB6, 0xFF, 0x04, 0x6B, 0x68}}

type IPlaylistStaticsInterface interface {
	win32.IInspectableInterface
	LoadAsync(file *IStorageFile) *IAsyncOperation[*IPlaylist]
}

type IPlaylistStaticsVtbl struct {
	win32.IInspectableVtbl
	LoadAsync uintptr
}

type IPlaylistStatics struct {
	win32.IInspectable
}

func (this *IPlaylistStatics) Vtbl() *IPlaylistStaticsVtbl {
	return (*IPlaylistStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPlaylistStatics) LoadAsync(file *IStorageFile) *IAsyncOperation[*IPlaylist] {
	var _result *IAsyncOperation[*IPlaylist]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LoadAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type Playlist struct {
	RtClass
	*IPlaylist
}

func NewPlaylist() *Playlist {
	hs := NewHStr("Windows.Media.Playlists.Playlist")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &Playlist{
		RtClass:   RtClass{PInspect: p},
		IPlaylist: (*IPlaylist)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewIPlaylistStatics() *IPlaylistStatics {
	var p *IPlaylistStatics
	hs := NewHStr("Windows.Media.Playlists.Playlist")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPlaylistStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
