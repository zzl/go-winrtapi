package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type VibrationAccessStatus int32

const (
	VibrationAccessStatus_Allowed             VibrationAccessStatus = 0
	VibrationAccessStatus_DeniedByUser        VibrationAccessStatus = 1
	VibrationAccessStatus_DeniedBySystem      VibrationAccessStatus = 2
	VibrationAccessStatus_DeniedByEnergySaver VibrationAccessStatus = 3
)

// interfaces

// 3D577EF7-4CEE-11E6-B535-001BDC06AB3B
var IID_IKnownSimpleHapticsControllerWaveformsStatics = syscall.GUID{0x3D577EF7, 0x4CEE, 0x11E6,
	[8]byte{0xB5, 0x35, 0x00, 0x1B, 0xDC, 0x06, 0xAB, 0x3B}}

type IKnownSimpleHapticsControllerWaveformsStaticsInterface interface {
	win32.IInspectableInterface
	Get_Click() uint16
	Get_BuzzContinuous() uint16
	Get_RumbleContinuous() uint16
	Get_Press() uint16
	Get_Release() uint16
}

type IKnownSimpleHapticsControllerWaveformsStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Click            uintptr
	Get_BuzzContinuous   uintptr
	Get_RumbleContinuous uintptr
	Get_Press            uintptr
	Get_Release          uintptr
}

type IKnownSimpleHapticsControllerWaveformsStatics struct {
	win32.IInspectable
}

func (this *IKnownSimpleHapticsControllerWaveformsStatics) Vtbl() *IKnownSimpleHapticsControllerWaveformsStaticsVtbl {
	return (*IKnownSimpleHapticsControllerWaveformsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownSimpleHapticsControllerWaveformsStatics) Get_Click() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Click, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IKnownSimpleHapticsControllerWaveformsStatics) Get_BuzzContinuous() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BuzzContinuous, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IKnownSimpleHapticsControllerWaveformsStatics) Get_RumbleContinuous() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RumbleContinuous, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IKnownSimpleHapticsControllerWaveformsStatics) Get_Press() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Press, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IKnownSimpleHapticsControllerWaveformsStatics) Get_Release() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Release, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3D577EF9-4CEE-11E6-B535-001BDC06AB3B
var IID_ISimpleHapticsController = syscall.GUID{0x3D577EF9, 0x4CEE, 0x11E6,
	[8]byte{0xB5, 0x35, 0x00, 0x1B, 0xDC, 0x06, 0xAB, 0x3B}}

type ISimpleHapticsControllerInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Get_SupportedFeedback() *IVectorView[*ISimpleHapticsControllerFeedback]
	Get_IsIntensitySupported() bool
	Get_IsPlayCountSupported() bool
	Get_IsPlayDurationSupported() bool
	Get_IsReplayPauseIntervalSupported() bool
	StopFeedback()
	SendHapticFeedback(feedback *ISimpleHapticsControllerFeedback)
	SendHapticFeedbackWithIntensity(feedback *ISimpleHapticsControllerFeedback, intensity float64)
	SendHapticFeedbackForDuration(feedback *ISimpleHapticsControllerFeedback, intensity float64, playDuration TimeSpan)
	SendHapticFeedbackForPlayCount(feedback *ISimpleHapticsControllerFeedback, intensity float64, playCount int32, replayPauseInterval TimeSpan)
}

type ISimpleHapticsControllerVtbl struct {
	win32.IInspectableVtbl
	Get_Id                             uintptr
	Get_SupportedFeedback              uintptr
	Get_IsIntensitySupported           uintptr
	Get_IsPlayCountSupported           uintptr
	Get_IsPlayDurationSupported        uintptr
	Get_IsReplayPauseIntervalSupported uintptr
	StopFeedback                       uintptr
	SendHapticFeedback                 uintptr
	SendHapticFeedbackWithIntensity    uintptr
	SendHapticFeedbackForDuration      uintptr
	SendHapticFeedbackForPlayCount     uintptr
}

type ISimpleHapticsController struct {
	win32.IInspectable
}

func (this *ISimpleHapticsController) Vtbl() *ISimpleHapticsControllerVtbl {
	return (*ISimpleHapticsControllerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISimpleHapticsController) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISimpleHapticsController) Get_SupportedFeedback() *IVectorView[*ISimpleHapticsControllerFeedback] {
	var _result *IVectorView[*ISimpleHapticsControllerFeedback]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedFeedback, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISimpleHapticsController) Get_IsIntensitySupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsIntensitySupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISimpleHapticsController) Get_IsPlayCountSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPlayCountSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISimpleHapticsController) Get_IsPlayDurationSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPlayDurationSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISimpleHapticsController) Get_IsReplayPauseIntervalSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsReplayPauseIntervalSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISimpleHapticsController) StopFeedback() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopFeedback, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ISimpleHapticsController) SendHapticFeedback(feedback *ISimpleHapticsControllerFeedback) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendHapticFeedback, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(feedback)))
	_ = _hr
}

func (this *ISimpleHapticsController) SendHapticFeedbackWithIntensity(feedback *ISimpleHapticsControllerFeedback, intensity float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendHapticFeedbackWithIntensity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(feedback)), uintptr(intensity))
	_ = _hr
}

func (this *ISimpleHapticsController) SendHapticFeedbackForDuration(feedback *ISimpleHapticsControllerFeedback, intensity float64, playDuration TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendHapticFeedbackForDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(feedback)), uintptr(intensity), *(*uintptr)(unsafe.Pointer(&playDuration)))
	_ = _hr
}

func (this *ISimpleHapticsController) SendHapticFeedbackForPlayCount(feedback *ISimpleHapticsControllerFeedback, intensity float64, playCount int32, replayPauseInterval TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendHapticFeedbackForPlayCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(feedback)), uintptr(intensity), uintptr(playCount), *(*uintptr)(unsafe.Pointer(&replayPauseInterval)))
	_ = _hr
}

// 3D577EF8-4CEE-11E6-B535-001BDC06AB3B
var IID_ISimpleHapticsControllerFeedback = syscall.GUID{0x3D577EF8, 0x4CEE, 0x11E6,
	[8]byte{0xB5, 0x35, 0x00, 0x1B, 0xDC, 0x06, 0xAB, 0x3B}}

type ISimpleHapticsControllerFeedbackInterface interface {
	win32.IInspectableInterface
	Get_Waveform() uint16
	Get_Duration() TimeSpan
}

type ISimpleHapticsControllerFeedbackVtbl struct {
	win32.IInspectableVtbl
	Get_Waveform uintptr
	Get_Duration uintptr
}

type ISimpleHapticsControllerFeedback struct {
	win32.IInspectable
}

func (this *ISimpleHapticsControllerFeedback) Vtbl() *ISimpleHapticsControllerFeedbackVtbl {
	return (*ISimpleHapticsControllerFeedbackVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISimpleHapticsControllerFeedback) Get_Waveform() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Waveform, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISimpleHapticsControllerFeedback) Get_Duration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 40F21A3E-8844-47FF-B312-06185A3844DA
var IID_IVibrationDevice = syscall.GUID{0x40F21A3E, 0x8844, 0x47FF,
	[8]byte{0xB3, 0x12, 0x06, 0x18, 0x5A, 0x38, 0x44, 0xDA}}

type IVibrationDeviceInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Get_SimpleHapticsController() *ISimpleHapticsController
}

type IVibrationDeviceVtbl struct {
	win32.IInspectableVtbl
	Get_Id                      uintptr
	Get_SimpleHapticsController uintptr
}

type IVibrationDevice struct {
	win32.IInspectable
}

func (this *IVibrationDevice) Vtbl() *IVibrationDeviceVtbl {
	return (*IVibrationDeviceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVibrationDevice) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IVibrationDevice) Get_SimpleHapticsController() *ISimpleHapticsController {
	var _result *ISimpleHapticsController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SimpleHapticsController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 53E2EDED-2290-4AC9-8EB3-1A84122EB71C
var IID_IVibrationDeviceStatics = syscall.GUID{0x53E2EDED, 0x2290, 0x4AC9,
	[8]byte{0x8E, 0xB3, 0x1A, 0x84, 0x12, 0x2E, 0xB7, 0x1C}}

type IVibrationDeviceStaticsInterface interface {
	win32.IInspectableInterface
	RequestAccessAsync() *IAsyncOperation[VibrationAccessStatus]
	GetDeviceSelector() string
	FromIdAsync(deviceId string) *IAsyncOperation[*IVibrationDevice]
	GetDefaultAsync() *IAsyncOperation[*IVibrationDevice]
	FindAllAsync() *IAsyncOperation[*IVectorView[*IVibrationDevice]]
}

type IVibrationDeviceStaticsVtbl struct {
	win32.IInspectableVtbl
	RequestAccessAsync uintptr
	GetDeviceSelector  uintptr
	FromIdAsync        uintptr
	GetDefaultAsync    uintptr
	FindAllAsync       uintptr
}

type IVibrationDeviceStatics struct {
	win32.IInspectable
}

func (this *IVibrationDeviceStatics) Vtbl() *IVibrationDeviceStaticsVtbl {
	return (*IVibrationDeviceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVibrationDeviceStatics) RequestAccessAsync() *IAsyncOperation[VibrationAccessStatus] {
	var _result *IAsyncOperation[VibrationAccessStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVibrationDeviceStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IVibrationDeviceStatics) FromIdAsync(deviceId string) *IAsyncOperation[*IVibrationDevice] {
	var _result *IAsyncOperation[*IVibrationDevice]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVibrationDeviceStatics) GetDefaultAsync() *IAsyncOperation[*IVibrationDevice] {
	var _result *IAsyncOperation[*IVibrationDevice]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVibrationDeviceStatics) FindAllAsync() *IAsyncOperation[*IVectorView[*IVibrationDevice]] {
	var _result *IAsyncOperation[*IVectorView[*IVibrationDevice]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type KnownSimpleHapticsControllerWaveforms struct {
	RtClass
}

func NewIKnownSimpleHapticsControllerWaveformsStatics() *IKnownSimpleHapticsControllerWaveformsStatics {
	var p *IKnownSimpleHapticsControllerWaveformsStatics
	hs := NewHStr("Windows.Devices.Haptics.KnownSimpleHapticsControllerWaveforms")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKnownSimpleHapticsControllerWaveformsStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SimpleHapticsController struct {
	RtClass
	*ISimpleHapticsController
}

type SimpleHapticsControllerFeedback struct {
	RtClass
	*ISimpleHapticsControllerFeedback
}

type VibrationDevice struct {
	RtClass
	*IVibrationDevice
}

func NewIVibrationDeviceStatics() *IVibrationDeviceStatics {
	var p *IVibrationDeviceStatics
	hs := NewHStr("Windows.Devices.Haptics.VibrationDevice")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IVibrationDeviceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
