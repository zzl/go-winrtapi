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
type GraphicsCaptureAccessKind int32

const (
	GraphicsCaptureAccessKind_Borderless GraphicsCaptureAccessKind = 0
)

// interfaces

// FA50C623-38DA-4B32-ACF3-FA9734AD800E
var IID_IDirect3D11CaptureFrame = syscall.GUID{0xFA50C623, 0x38DA, 0x4B32,
	[8]byte{0xAC, 0xF3, 0xFA, 0x97, 0x34, 0xAD, 0x80, 0x0E}}

type IDirect3D11CaptureFrameInterface interface {
	win32.IInspectableInterface
	Get_Surface() unsafe.Pointer
	Get_SystemRelativeTime() TimeSpan
	Get_ContentSize() unsafe.Pointer
}

type IDirect3D11CaptureFrameVtbl struct {
	win32.IInspectableVtbl
	Get_Surface            uintptr
	Get_SystemRelativeTime uintptr
	Get_ContentSize        uintptr
}

type IDirect3D11CaptureFrame struct {
	win32.IInspectable
}

func (this *IDirect3D11CaptureFrame) Vtbl() *IDirect3D11CaptureFrameVtbl {
	return (*IDirect3D11CaptureFrameVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDirect3D11CaptureFrame) Get_Surface() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Surface, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDirect3D11CaptureFrame) Get_SystemRelativeTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SystemRelativeTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDirect3D11CaptureFrame) Get_ContentSize() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 24EB6D22-1975-422E-82E7-780DBD8DDF24
var IID_IDirect3D11CaptureFramePool = syscall.GUID{0x24EB6D22, 0x1975, 0x422E,
	[8]byte{0x82, 0xE7, 0x78, 0x0D, 0xBD, 0x8D, 0xDF, 0x24}}

type IDirect3D11CaptureFramePoolInterface interface {
	win32.IInspectableInterface
	Recreate(device unsafe.Pointer, pixelFormat unsafe.Pointer, numberOfBuffers int32, size unsafe.Pointer)
	TryGetNextFrame() *IDirect3D11CaptureFrame
	Add_FrameArrived(handler TypedEventHandler[*IDirect3D11CaptureFramePool, interface{}]) EventRegistrationToken
	Remove_FrameArrived(token EventRegistrationToken)
	CreateCaptureSession(item *IGraphicsCaptureItem) *IGraphicsCaptureSession
	Get_DispatcherQueue() *IDispatcherQueue
}

type IDirect3D11CaptureFramePoolVtbl struct {
	win32.IInspectableVtbl
	Recreate             uintptr
	TryGetNextFrame      uintptr
	Add_FrameArrived     uintptr
	Remove_FrameArrived  uintptr
	CreateCaptureSession uintptr
	Get_DispatcherQueue  uintptr
}

type IDirect3D11CaptureFramePool struct {
	win32.IInspectable
}

func (this *IDirect3D11CaptureFramePool) Vtbl() *IDirect3D11CaptureFramePoolVtbl {
	return (*IDirect3D11CaptureFramePoolVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDirect3D11CaptureFramePool) Recreate(device unsafe.Pointer, pixelFormat unsafe.Pointer, numberOfBuffers int32, size unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Recreate, uintptr(unsafe.Pointer(this)), uintptr(device), uintptr(pixelFormat), uintptr(numberOfBuffers), uintptr(size))
	_ = _hr
}

func (this *IDirect3D11CaptureFramePool) TryGetNextFrame() *IDirect3D11CaptureFrame {
	var _result *IDirect3D11CaptureFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetNextFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDirect3D11CaptureFramePool) Add_FrameArrived(handler TypedEventHandler[*IDirect3D11CaptureFramePool, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_FrameArrived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDirect3D11CaptureFramePool) Remove_FrameArrived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_FrameArrived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IDirect3D11CaptureFramePool) CreateCaptureSession(item *IGraphicsCaptureItem) *IGraphicsCaptureSession {
	var _result *IGraphicsCaptureSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCaptureSession, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(item)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDirect3D11CaptureFramePool) Get_DispatcherQueue() *IDispatcherQueue {
	var _result *IDispatcherQueue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DispatcherQueue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7784056A-67AA-4D53-AE54-1088D5A8CA21
var IID_IDirect3D11CaptureFramePoolStatics = syscall.GUID{0x7784056A, 0x67AA, 0x4D53,
	[8]byte{0xAE, 0x54, 0x10, 0x88, 0xD5, 0xA8, 0xCA, 0x21}}

type IDirect3D11CaptureFramePoolStaticsInterface interface {
	win32.IInspectableInterface
	Create(device unsafe.Pointer, pixelFormat unsafe.Pointer, numberOfBuffers int32, size unsafe.Pointer) *IDirect3D11CaptureFramePool
}

type IDirect3D11CaptureFramePoolStaticsVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IDirect3D11CaptureFramePoolStatics struct {
	win32.IInspectable
}

func (this *IDirect3D11CaptureFramePoolStatics) Vtbl() *IDirect3D11CaptureFramePoolStaticsVtbl {
	return (*IDirect3D11CaptureFramePoolStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDirect3D11CaptureFramePoolStatics) Create(device unsafe.Pointer, pixelFormat unsafe.Pointer, numberOfBuffers int32, size unsafe.Pointer) *IDirect3D11CaptureFramePool {
	var _result *IDirect3D11CaptureFramePool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(device), uintptr(pixelFormat), uintptr(numberOfBuffers), uintptr(size), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 589B103F-6BBC-5DF5-A991-02E28B3B66D5
var IID_IDirect3D11CaptureFramePoolStatics2 = syscall.GUID{0x589B103F, 0x6BBC, 0x5DF5,
	[8]byte{0xA9, 0x91, 0x02, 0xE2, 0x8B, 0x3B, 0x66, 0xD5}}

type IDirect3D11CaptureFramePoolStatics2Interface interface {
	win32.IInspectableInterface
	CreateFreeThreaded(device unsafe.Pointer, pixelFormat unsafe.Pointer, numberOfBuffers int32, size unsafe.Pointer) *IDirect3D11CaptureFramePool
}

type IDirect3D11CaptureFramePoolStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateFreeThreaded uintptr
}

type IDirect3D11CaptureFramePoolStatics2 struct {
	win32.IInspectable
}

func (this *IDirect3D11CaptureFramePoolStatics2) Vtbl() *IDirect3D11CaptureFramePoolStatics2Vtbl {
	return (*IDirect3D11CaptureFramePoolStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDirect3D11CaptureFramePoolStatics2) CreateFreeThreaded(device unsafe.Pointer, pixelFormat unsafe.Pointer, numberOfBuffers int32, size unsafe.Pointer) *IDirect3D11CaptureFramePool {
	var _result *IDirect3D11CaptureFramePool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFreeThreaded, uintptr(unsafe.Pointer(this)), uintptr(device), uintptr(pixelFormat), uintptr(numberOfBuffers), uintptr(size), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 743ED370-06EC-5040-A58A-901F0F757095
var IID_IGraphicsCaptureAccessStatics = syscall.GUID{0x743ED370, 0x06EC, 0x5040,
	[8]byte{0xA5, 0x8A, 0x90, 0x1F, 0x0F, 0x75, 0x70, 0x95}}

type IGraphicsCaptureAccessStaticsInterface interface {
	win32.IInspectableInterface
	RequestAccessAsync(request GraphicsCaptureAccessKind) *IAsyncOperation[unsafe.Pointer]
}

type IGraphicsCaptureAccessStaticsVtbl struct {
	win32.IInspectableVtbl
	RequestAccessAsync uintptr
}

type IGraphicsCaptureAccessStatics struct {
	win32.IInspectable
}

func (this *IGraphicsCaptureAccessStatics) Vtbl() *IGraphicsCaptureAccessStaticsVtbl {
	return (*IGraphicsCaptureAccessStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGraphicsCaptureAccessStatics) RequestAccessAsync(request GraphicsCaptureAccessKind) *IAsyncOperation[unsafe.Pointer] {
	var _result *IAsyncOperation[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessAsync, uintptr(unsafe.Pointer(this)), uintptr(request), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 79C3F95B-31F7-4EC2-A464-632EF5D30760
var IID_IGraphicsCaptureItem = syscall.GUID{0x79C3F95B, 0x31F7, 0x4EC2,
	[8]byte{0xA4, 0x64, 0x63, 0x2E, 0xF5, 0xD3, 0x07, 0x60}}

type IGraphicsCaptureItemInterface interface {
	win32.IInspectableInterface
	Get_DisplayName() string
	Get_Size() unsafe.Pointer
	Add_Closed(handler TypedEventHandler[*IGraphicsCaptureItem, interface{}]) EventRegistrationToken
	Remove_Closed(token EventRegistrationToken)
}

type IGraphicsCaptureItemVtbl struct {
	win32.IInspectableVtbl
	Get_DisplayName uintptr
	Get_Size        uintptr
	Add_Closed      uintptr
	Remove_Closed   uintptr
}

type IGraphicsCaptureItem struct {
	win32.IInspectable
}

func (this *IGraphicsCaptureItem) Vtbl() *IGraphicsCaptureItemVtbl {
	return (*IGraphicsCaptureItemVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGraphicsCaptureItem) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGraphicsCaptureItem) Get_Size() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Size, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGraphicsCaptureItem) Add_Closed(handler TypedEventHandler[*IGraphicsCaptureItem, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Closed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGraphicsCaptureItem) Remove_Closed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Closed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// A87EBEA5-457C-5788-AB47-0CF1D3637E74
var IID_IGraphicsCaptureItemStatics = syscall.GUID{0xA87EBEA5, 0x457C, 0x5788,
	[8]byte{0xAB, 0x47, 0x0C, 0xF1, 0xD3, 0x63, 0x7E, 0x74}}

type IGraphicsCaptureItemStaticsInterface interface {
	win32.IInspectableInterface
	CreateFromVisual(visual *IVisual) *IGraphicsCaptureItem
}

type IGraphicsCaptureItemStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateFromVisual uintptr
}

type IGraphicsCaptureItemStatics struct {
	win32.IInspectable
}

func (this *IGraphicsCaptureItemStatics) Vtbl() *IGraphicsCaptureItemStaticsVtbl {
	return (*IGraphicsCaptureItemStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGraphicsCaptureItemStatics) CreateFromVisual(visual *IVisual) *IGraphicsCaptureItem {
	var _result *IGraphicsCaptureItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromVisual, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(visual)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3B92ACC9-E584-5862-BF5C-9C316C6D2DBB
var IID_IGraphicsCaptureItemStatics2 = syscall.GUID{0x3B92ACC9, 0xE584, 0x5862,
	[8]byte{0xBF, 0x5C, 0x9C, 0x31, 0x6C, 0x6D, 0x2D, 0xBB}}

type IGraphicsCaptureItemStatics2Interface interface {
	win32.IInspectableInterface
	TryCreateFromWindowId(windowId unsafe.Pointer) *IGraphicsCaptureItem
	TryCreateFromDisplayId(displayId unsafe.Pointer) *IGraphicsCaptureItem
}

type IGraphicsCaptureItemStatics2Vtbl struct {
	win32.IInspectableVtbl
	TryCreateFromWindowId  uintptr
	TryCreateFromDisplayId uintptr
}

type IGraphicsCaptureItemStatics2 struct {
	win32.IInspectable
}

func (this *IGraphicsCaptureItemStatics2) Vtbl() *IGraphicsCaptureItemStatics2Vtbl {
	return (*IGraphicsCaptureItemStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGraphicsCaptureItemStatics2) TryCreateFromWindowId(windowId unsafe.Pointer) *IGraphicsCaptureItem {
	var _result *IGraphicsCaptureItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryCreateFromWindowId, uintptr(unsafe.Pointer(this)), uintptr(windowId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGraphicsCaptureItemStatics2) TryCreateFromDisplayId(displayId unsafe.Pointer) *IGraphicsCaptureItem {
	var _result *IGraphicsCaptureItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryCreateFromDisplayId, uintptr(unsafe.Pointer(this)), uintptr(displayId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5A1711B3-AD79-4B4A-9336-1318FDDE3539
var IID_IGraphicsCapturePicker = syscall.GUID{0x5A1711B3, 0xAD79, 0x4B4A,
	[8]byte{0x93, 0x36, 0x13, 0x18, 0xFD, 0xDE, 0x35, 0x39}}

type IGraphicsCapturePickerInterface interface {
	win32.IInspectableInterface
	PickSingleItemAsync() *IAsyncOperation[*IGraphicsCaptureItem]
}

type IGraphicsCapturePickerVtbl struct {
	win32.IInspectableVtbl
	PickSingleItemAsync uintptr
}

type IGraphicsCapturePicker struct {
	win32.IInspectable
}

func (this *IGraphicsCapturePicker) Vtbl() *IGraphicsCapturePickerVtbl {
	return (*IGraphicsCapturePickerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGraphicsCapturePicker) PickSingleItemAsync() *IAsyncOperation[*IGraphicsCaptureItem] {
	var _result *IAsyncOperation[*IGraphicsCaptureItem]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PickSingleItemAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 814E42A9-F70F-4AD7-939B-FDDCC6EB880D
var IID_IGraphicsCaptureSession = syscall.GUID{0x814E42A9, 0xF70F, 0x4AD7,
	[8]byte{0x93, 0x9B, 0xFD, 0xDC, 0xC6, 0xEB, 0x88, 0x0D}}

type IGraphicsCaptureSessionInterface interface {
	win32.IInspectableInterface
	StartCapture()
}

type IGraphicsCaptureSessionVtbl struct {
	win32.IInspectableVtbl
	StartCapture uintptr
}

type IGraphicsCaptureSession struct {
	win32.IInspectable
}

func (this *IGraphicsCaptureSession) Vtbl() *IGraphicsCaptureSessionVtbl {
	return (*IGraphicsCaptureSessionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGraphicsCaptureSession) StartCapture() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartCapture, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 2C39AE40-7D2E-5044-804E-8B6799D4CF9E
var IID_IGraphicsCaptureSession2 = syscall.GUID{0x2C39AE40, 0x7D2E, 0x5044,
	[8]byte{0x80, 0x4E, 0x8B, 0x67, 0x99, 0xD4, 0xCF, 0x9E}}

type IGraphicsCaptureSession2Interface interface {
	win32.IInspectableInterface
	Get_IsCursorCaptureEnabled() bool
	Put_IsCursorCaptureEnabled(value bool)
}

type IGraphicsCaptureSession2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsCursorCaptureEnabled uintptr
	Put_IsCursorCaptureEnabled uintptr
}

type IGraphicsCaptureSession2 struct {
	win32.IInspectable
}

func (this *IGraphicsCaptureSession2) Vtbl() *IGraphicsCaptureSession2Vtbl {
	return (*IGraphicsCaptureSession2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGraphicsCaptureSession2) Get_IsCursorCaptureEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCursorCaptureEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGraphicsCaptureSession2) Put_IsCursorCaptureEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsCursorCaptureEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// F2CDD966-22AE-5EA1-9596-3A289344C3BE
var IID_IGraphicsCaptureSession3 = syscall.GUID{0xF2CDD966, 0x22AE, 0x5EA1,
	[8]byte{0x95, 0x96, 0x3A, 0x28, 0x93, 0x44, 0xC3, 0xBE}}

type IGraphicsCaptureSession3Interface interface {
	win32.IInspectableInterface
	Get_IsBorderRequired() bool
	Put_IsBorderRequired(value bool)
}

type IGraphicsCaptureSession3Vtbl struct {
	win32.IInspectableVtbl
	Get_IsBorderRequired uintptr
	Put_IsBorderRequired uintptr
}

type IGraphicsCaptureSession3 struct {
	win32.IInspectable
}

func (this *IGraphicsCaptureSession3) Vtbl() *IGraphicsCaptureSession3Vtbl {
	return (*IGraphicsCaptureSession3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGraphicsCaptureSession3) Get_IsBorderRequired() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsBorderRequired, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGraphicsCaptureSession3) Put_IsBorderRequired(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsBorderRequired, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 2224A540-5974-49AA-B232-0882536F4CB5
var IID_IGraphicsCaptureSessionStatics = syscall.GUID{0x2224A540, 0x5974, 0x49AA,
	[8]byte{0xB2, 0x32, 0x08, 0x82, 0x53, 0x6F, 0x4C, 0xB5}}

type IGraphicsCaptureSessionStaticsInterface interface {
	win32.IInspectableInterface
	IsSupported() bool
}

type IGraphicsCaptureSessionStaticsVtbl struct {
	win32.IInspectableVtbl
	IsSupported uintptr
}

type IGraphicsCaptureSessionStatics struct {
	win32.IInspectable
}

func (this *IGraphicsCaptureSessionStatics) Vtbl() *IGraphicsCaptureSessionStaticsVtbl {
	return (*IGraphicsCaptureSessionStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGraphicsCaptureSessionStatics) IsSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type Direct3D11CaptureFrame struct {
	RtClass
	*IDirect3D11CaptureFrame
}

type Direct3D11CaptureFramePool struct {
	RtClass
	*IDirect3D11CaptureFramePool
}

func NewIDirect3D11CaptureFramePoolStatics2() *IDirect3D11CaptureFramePoolStatics2 {
	var p *IDirect3D11CaptureFramePoolStatics2
	hs := NewHStr("Windows.Graphics.Capture.Direct3D11CaptureFramePool")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDirect3D11CaptureFramePoolStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIDirect3D11CaptureFramePoolStatics() *IDirect3D11CaptureFramePoolStatics {
	var p *IDirect3D11CaptureFramePoolStatics
	hs := NewHStr("Windows.Graphics.Capture.Direct3D11CaptureFramePool")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDirect3D11CaptureFramePoolStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type GraphicsCaptureItem struct {
	RtClass
	*IGraphicsCaptureItem
}

func NewIGraphicsCaptureItemStatics2() *IGraphicsCaptureItemStatics2 {
	var p *IGraphicsCaptureItemStatics2
	hs := NewHStr("Windows.Graphics.Capture.GraphicsCaptureItem")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGraphicsCaptureItemStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIGraphicsCaptureItemStatics() *IGraphicsCaptureItemStatics {
	var p *IGraphicsCaptureItemStatics
	hs := NewHStr("Windows.Graphics.Capture.GraphicsCaptureItem")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGraphicsCaptureItemStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type GraphicsCapturePicker struct {
	RtClass
	*IGraphicsCapturePicker
}

func NewGraphicsCapturePicker() *GraphicsCapturePicker {
	hs := NewHStr("Windows.Graphics.Capture.GraphicsCapturePicker")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &GraphicsCapturePicker{
		RtClass:                RtClass{PInspect: p},
		IGraphicsCapturePicker: (*IGraphicsCapturePicker)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type GraphicsCaptureSession struct {
	RtClass
	*IGraphicsCaptureSession
}

func NewIGraphicsCaptureSessionStatics() *IGraphicsCaptureSessionStatics {
	var p *IGraphicsCaptureSessionStatics
	hs := NewHStr("Windows.Graphics.Capture.GraphicsCaptureSession")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGraphicsCaptureSessionStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
