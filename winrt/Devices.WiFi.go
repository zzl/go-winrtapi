package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type WiFiAccessStatus int32

const (
	WiFiAccessStatus_Unspecified    WiFiAccessStatus = 0
	WiFiAccessStatus_Allowed        WiFiAccessStatus = 1
	WiFiAccessStatus_DeniedByUser   WiFiAccessStatus = 2
	WiFiAccessStatus_DeniedBySystem WiFiAccessStatus = 3
)

// enum
type WiFiConnectionMethod int32

const (
	WiFiConnectionMethod_Default       WiFiConnectionMethod = 0
	WiFiConnectionMethod_WpsPin        WiFiConnectionMethod = 1
	WiFiConnectionMethod_WpsPushButton WiFiConnectionMethod = 2
)

// enum
type WiFiConnectionStatus int32

const (
	WiFiConnectionStatus_UnspecifiedFailure                WiFiConnectionStatus = 0
	WiFiConnectionStatus_Success                           WiFiConnectionStatus = 1
	WiFiConnectionStatus_AccessRevoked                     WiFiConnectionStatus = 2
	WiFiConnectionStatus_InvalidCredential                 WiFiConnectionStatus = 3
	WiFiConnectionStatus_NetworkNotAvailable               WiFiConnectionStatus = 4
	WiFiConnectionStatus_Timeout                           WiFiConnectionStatus = 5
	WiFiConnectionStatus_UnsupportedAuthenticationProtocol WiFiConnectionStatus = 6
)

// enum
type WiFiNetworkKind int32

const (
	WiFiNetworkKind_Any            WiFiNetworkKind = 0
	WiFiNetworkKind_Infrastructure WiFiNetworkKind = 1
	WiFiNetworkKind_Adhoc          WiFiNetworkKind = 2
)

// enum
type WiFiPhyKind int32

const (
	WiFiPhyKind_Unknown    WiFiPhyKind = 0
	WiFiPhyKind_Fhss       WiFiPhyKind = 1
	WiFiPhyKind_Dsss       WiFiPhyKind = 2
	WiFiPhyKind_IRBaseband WiFiPhyKind = 3
	WiFiPhyKind_Ofdm       WiFiPhyKind = 4
	WiFiPhyKind_Hrdsss     WiFiPhyKind = 5
	WiFiPhyKind_Erp        WiFiPhyKind = 6
	WiFiPhyKind_HT         WiFiPhyKind = 7
	WiFiPhyKind_Vht        WiFiPhyKind = 8
	WiFiPhyKind_Dmg        WiFiPhyKind = 9
	WiFiPhyKind_HE         WiFiPhyKind = 10
)

// enum
type WiFiReconnectionKind int32

const (
	WiFiReconnectionKind_Automatic WiFiReconnectionKind = 0
	WiFiReconnectionKind_Manual    WiFiReconnectionKind = 1
)

// enum
type WiFiWpsConfigurationStatus int32

const (
	WiFiWpsConfigurationStatus_UnspecifiedFailure WiFiWpsConfigurationStatus = 0
	WiFiWpsConfigurationStatus_Success            WiFiWpsConfigurationStatus = 1
	WiFiWpsConfigurationStatus_Timeout            WiFiWpsConfigurationStatus = 2
)

// enum
type WiFiWpsKind int32

const (
	WiFiWpsKind_Unknown    WiFiWpsKind = 0
	WiFiWpsKind_Pin        WiFiWpsKind = 1
	WiFiWpsKind_PushButton WiFiWpsKind = 2
	WiFiWpsKind_Nfc        WiFiWpsKind = 3
	WiFiWpsKind_Ethernet   WiFiWpsKind = 4
	WiFiWpsKind_Usb        WiFiWpsKind = 5
)

// interfaces

// A6C4E423-3D75-43A4-B9DE-11E26B72D9B0
var IID_IWiFiAdapter = syscall.GUID{0xA6C4E423, 0x3D75, 0x43A4,
	[8]byte{0xB9, 0xDE, 0x11, 0xE2, 0x6B, 0x72, 0xD9, 0xB0}}

type IWiFiAdapterInterface interface {
	win32.IInspectableInterface
	Get_NetworkAdapter() *INetworkAdapter
	ScanAsync() *IAsyncAction
	Get_NetworkReport() *IWiFiNetworkReport
	Add_AvailableNetworksChanged(args TypedEventHandler[*IWiFiAdapter, interface{}]) EventRegistrationToken
	Remove_AvailableNetworksChanged(eventCookie EventRegistrationToken)
	ConnectAsync(availableNetwork *IWiFiAvailableNetwork, reconnectionKind WiFiReconnectionKind) *IAsyncOperation[*IWiFiConnectionResult]
	ConnectWithPasswordCredentialAsync(availableNetwork *IWiFiAvailableNetwork, reconnectionKind WiFiReconnectionKind, passwordCredential *IPasswordCredential) *IAsyncOperation[*IWiFiConnectionResult]
	ConnectWithPasswordCredentialAndSsidAsync(availableNetwork *IWiFiAvailableNetwork, reconnectionKind WiFiReconnectionKind, passwordCredential *IPasswordCredential, ssid string) *IAsyncOperation[*IWiFiConnectionResult]
	Disconnect()
}

type IWiFiAdapterVtbl struct {
	win32.IInspectableVtbl
	Get_NetworkAdapter                        uintptr
	ScanAsync                                 uintptr
	Get_NetworkReport                         uintptr
	Add_AvailableNetworksChanged              uintptr
	Remove_AvailableNetworksChanged           uintptr
	ConnectAsync                              uintptr
	ConnectWithPasswordCredentialAsync        uintptr
	ConnectWithPasswordCredentialAndSsidAsync uintptr
	Disconnect                                uintptr
}

type IWiFiAdapter struct {
	win32.IInspectable
}

func (this *IWiFiAdapter) Vtbl() *IWiFiAdapterVtbl {
	return (*IWiFiAdapterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiAdapter) Get_NetworkAdapter() *INetworkAdapter {
	var _result *INetworkAdapter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkAdapter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiAdapter) ScanAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ScanAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiAdapter) Get_NetworkReport() *IWiFiNetworkReport {
	var _result *IWiFiNetworkReport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkReport, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiAdapter) Add_AvailableNetworksChanged(args TypedEventHandler[*IWiFiAdapter, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AvailableNetworksChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(args))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiAdapter) Remove_AvailableNetworksChanged(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AvailableNetworksChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

func (this *IWiFiAdapter) ConnectAsync(availableNetwork *IWiFiAvailableNetwork, reconnectionKind WiFiReconnectionKind) *IAsyncOperation[*IWiFiConnectionResult] {
	var _result *IAsyncOperation[*IWiFiConnectionResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConnectAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(availableNetwork)), uintptr(reconnectionKind), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiAdapter) ConnectWithPasswordCredentialAsync(availableNetwork *IWiFiAvailableNetwork, reconnectionKind WiFiReconnectionKind, passwordCredential *IPasswordCredential) *IAsyncOperation[*IWiFiConnectionResult] {
	var _result *IAsyncOperation[*IWiFiConnectionResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConnectWithPasswordCredentialAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(availableNetwork)), uintptr(reconnectionKind), uintptr(unsafe.Pointer(passwordCredential)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiAdapter) ConnectWithPasswordCredentialAndSsidAsync(availableNetwork *IWiFiAvailableNetwork, reconnectionKind WiFiReconnectionKind, passwordCredential *IPasswordCredential, ssid string) *IAsyncOperation[*IWiFiConnectionResult] {
	var _result *IAsyncOperation[*IWiFiConnectionResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConnectWithPasswordCredentialAndSsidAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(availableNetwork)), uintptr(reconnectionKind), uintptr(unsafe.Pointer(passwordCredential)), NewHStr(ssid).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiAdapter) Disconnect() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Disconnect, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 5BC4501D-81E4-453D-9430-1FCAFBADD6B6
var IID_IWiFiAdapter2 = syscall.GUID{0x5BC4501D, 0x81E4, 0x453D,
	[8]byte{0x94, 0x30, 0x1F, 0xCA, 0xFB, 0xAD, 0xD6, 0xB6}}

type IWiFiAdapter2Interface interface {
	win32.IInspectableInterface
	GetWpsConfigurationAsync(availableNetwork *IWiFiAvailableNetwork) *IAsyncOperation[*IWiFiWpsConfigurationResult]
	ConnectWithPasswordCredentialAndSsidAndConnectionMethodAsync(availableNetwork *IWiFiAvailableNetwork, reconnectionKind WiFiReconnectionKind, passwordCredential *IPasswordCredential, ssid string, connectionMethod WiFiConnectionMethod) *IAsyncOperation[*IWiFiConnectionResult]
}

type IWiFiAdapter2Vtbl struct {
	win32.IInspectableVtbl
	GetWpsConfigurationAsync                                     uintptr
	ConnectWithPasswordCredentialAndSsidAndConnectionMethodAsync uintptr
}

type IWiFiAdapter2 struct {
	win32.IInspectable
}

func (this *IWiFiAdapter2) Vtbl() *IWiFiAdapter2Vtbl {
	return (*IWiFiAdapter2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiAdapter2) GetWpsConfigurationAsync(availableNetwork *IWiFiAvailableNetwork) *IAsyncOperation[*IWiFiWpsConfigurationResult] {
	var _result *IAsyncOperation[*IWiFiWpsConfigurationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetWpsConfigurationAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(availableNetwork)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiAdapter2) ConnectWithPasswordCredentialAndSsidAndConnectionMethodAsync(availableNetwork *IWiFiAvailableNetwork, reconnectionKind WiFiReconnectionKind, passwordCredential *IPasswordCredential, ssid string, connectionMethod WiFiConnectionMethod) *IAsyncOperation[*IWiFiConnectionResult] {
	var _result *IAsyncOperation[*IWiFiConnectionResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConnectWithPasswordCredentialAndSsidAndConnectionMethodAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(availableNetwork)), uintptr(reconnectionKind), uintptr(unsafe.Pointer(passwordCredential)), NewHStr(ssid).Ptr, uintptr(connectionMethod), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DA25FDDD-D24C-43E3-AABD-C4659F730F99
var IID_IWiFiAdapterStatics = syscall.GUID{0xDA25FDDD, 0xD24C, 0x43E3,
	[8]byte{0xAA, 0xBD, 0xC4, 0x65, 0x9F, 0x73, 0x0F, 0x99}}

type IWiFiAdapterStaticsInterface interface {
	win32.IInspectableInterface
	FindAllAdaptersAsync() *IAsyncOperation[*IVectorView[*IWiFiAdapter]]
	GetDeviceSelector() string
	FromIdAsync(deviceId string) *IAsyncOperation[*IWiFiAdapter]
	RequestAccessAsync() *IAsyncOperation[WiFiAccessStatus]
}

type IWiFiAdapterStaticsVtbl struct {
	win32.IInspectableVtbl
	FindAllAdaptersAsync uintptr
	GetDeviceSelector    uintptr
	FromIdAsync          uintptr
	RequestAccessAsync   uintptr
}

type IWiFiAdapterStatics struct {
	win32.IInspectable
}

func (this *IWiFiAdapterStatics) Vtbl() *IWiFiAdapterStaticsVtbl {
	return (*IWiFiAdapterStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiAdapterStatics) FindAllAdaptersAsync() *IAsyncOperation[*IVectorView[*IWiFiAdapter]] {
	var _result *IAsyncOperation[*IVectorView[*IWiFiAdapter]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllAdaptersAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiAdapterStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWiFiAdapterStatics) FromIdAsync(deviceId string) *IAsyncOperation[*IWiFiAdapter] {
	var _result *IAsyncOperation[*IWiFiAdapter]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiAdapterStatics) RequestAccessAsync() *IAsyncOperation[WiFiAccessStatus] {
	var _result *IAsyncOperation[WiFiAccessStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 26E96246-183E-4704-9826-71B4A2F0F668
var IID_IWiFiAvailableNetwork = syscall.GUID{0x26E96246, 0x183E, 0x4704,
	[8]byte{0x98, 0x26, 0x71, 0xB4, 0xA2, 0xF0, 0xF6, 0x68}}

type IWiFiAvailableNetworkInterface interface {
	win32.IInspectableInterface
	Get_Uptime() TimeSpan
	Get_Ssid() string
	Get_Bssid() string
	Get_ChannelCenterFrequencyInKilohertz() int32
	Get_NetworkRssiInDecibelMilliwatts() float64
	Get_SignalBars() byte
	Get_NetworkKind() WiFiNetworkKind
	Get_PhyKind() WiFiPhyKind
	Get_SecuritySettings() *INetworkSecuritySettings
	Get_BeaconInterval() TimeSpan
	Get_IsWiFiDirect() bool
}

type IWiFiAvailableNetworkVtbl struct {
	win32.IInspectableVtbl
	Get_Uptime                            uintptr
	Get_Ssid                              uintptr
	Get_Bssid                             uintptr
	Get_ChannelCenterFrequencyInKilohertz uintptr
	Get_NetworkRssiInDecibelMilliwatts    uintptr
	Get_SignalBars                        uintptr
	Get_NetworkKind                       uintptr
	Get_PhyKind                           uintptr
	Get_SecuritySettings                  uintptr
	Get_BeaconInterval                    uintptr
	Get_IsWiFiDirect                      uintptr
}

type IWiFiAvailableNetwork struct {
	win32.IInspectable
}

func (this *IWiFiAvailableNetwork) Vtbl() *IWiFiAvailableNetworkVtbl {
	return (*IWiFiAvailableNetworkVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiAvailableNetwork) Get_Uptime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Uptime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiAvailableNetwork) Get_Ssid() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ssid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWiFiAvailableNetwork) Get_Bssid() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bssid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWiFiAvailableNetwork) Get_ChannelCenterFrequencyInKilohertz() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ChannelCenterFrequencyInKilohertz, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiAvailableNetwork) Get_NetworkRssiInDecibelMilliwatts() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkRssiInDecibelMilliwatts, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiAvailableNetwork) Get_SignalBars() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SignalBars, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiAvailableNetwork) Get_NetworkKind() WiFiNetworkKind {
	var _result WiFiNetworkKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiAvailableNetwork) Get_PhyKind() WiFiPhyKind {
	var _result WiFiPhyKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhyKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiAvailableNetwork) Get_SecuritySettings() *INetworkSecuritySettings {
	var _result *INetworkSecuritySettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SecuritySettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiAvailableNetwork) Get_BeaconInterval() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BeaconInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiAvailableNetwork) Get_IsWiFiDirect() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsWiFiDirect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 143BDFD9-C37D-40BE-A5C8-857BCE85A931
var IID_IWiFiConnectionResult = syscall.GUID{0x143BDFD9, 0xC37D, 0x40BE,
	[8]byte{0xA5, 0xC8, 0x85, 0x7B, 0xCE, 0x85, 0xA9, 0x31}}

type IWiFiConnectionResultInterface interface {
	win32.IInspectableInterface
	Get_ConnectionStatus() WiFiConnectionStatus
}

type IWiFiConnectionResultVtbl struct {
	win32.IInspectableVtbl
	Get_ConnectionStatus uintptr
}

type IWiFiConnectionResult struct {
	win32.IInspectable
}

func (this *IWiFiConnectionResult) Vtbl() *IWiFiConnectionResultVtbl {
	return (*IWiFiConnectionResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiConnectionResult) Get_ConnectionStatus() WiFiConnectionStatus {
	var _result WiFiConnectionStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConnectionStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 9524DED2-5911-445E-8194-BE4F1A704895
var IID_IWiFiNetworkReport = syscall.GUID{0x9524DED2, 0x5911, 0x445E,
	[8]byte{0x81, 0x94, 0xBE, 0x4F, 0x1A, 0x70, 0x48, 0x95}}

type IWiFiNetworkReportInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() DateTime
	Get_AvailableNetworks() *IVectorView[*IWiFiAvailableNetwork]
}

type IWiFiNetworkReportVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp         uintptr
	Get_AvailableNetworks uintptr
}

type IWiFiNetworkReport struct {
	win32.IInspectable
}

func (this *IWiFiNetworkReport) Vtbl() *IWiFiNetworkReportVtbl {
	return (*IWiFiNetworkReportVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiNetworkReport) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiNetworkReport) Get_AvailableNetworks() *IVectorView[*IWiFiAvailableNetwork] {
	var _result *IVectorView[*IWiFiAvailableNetwork]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AvailableNetworks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 67B49871-17EE-42D1-B14F-5A11F1226FB5
var IID_IWiFiWpsConfigurationResult = syscall.GUID{0x67B49871, 0x17EE, 0x42D1,
	[8]byte{0xB1, 0x4F, 0x5A, 0x11, 0xF1, 0x22, 0x6F, 0xB5}}

type IWiFiWpsConfigurationResultInterface interface {
	win32.IInspectableInterface
	Get_Status() WiFiWpsConfigurationStatus
	Get_SupportedWpsKinds() *IVectorView[WiFiWpsKind]
}

type IWiFiWpsConfigurationResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status            uintptr
	Get_SupportedWpsKinds uintptr
}

type IWiFiWpsConfigurationResult struct {
	win32.IInspectable
}

func (this *IWiFiWpsConfigurationResult) Vtbl() *IWiFiWpsConfigurationResultVtbl {
	return (*IWiFiWpsConfigurationResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiWpsConfigurationResult) Get_Status() WiFiWpsConfigurationStatus {
	var _result WiFiWpsConfigurationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiWpsConfigurationResult) Get_SupportedWpsKinds() *IVectorView[WiFiWpsKind] {
	var _result *IVectorView[WiFiWpsKind]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedWpsKinds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type WiFiAdapter struct {
	RtClass
	*IWiFiAdapter
}

func NewIWiFiAdapterStatics() *IWiFiAdapterStatics {
	var p *IWiFiAdapterStatics
	hs := NewHStr("Windows.Devices.WiFi.WiFiAdapter")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWiFiAdapterStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type WiFiAvailableNetwork struct {
	RtClass
	*IWiFiAvailableNetwork
}

type WiFiConnectionResult struct {
	RtClass
	*IWiFiConnectionResult
}

type WiFiNetworkReport struct {
	RtClass
	*IWiFiNetworkReport
}
