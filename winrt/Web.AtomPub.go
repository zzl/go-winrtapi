package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"log"
	"syscall"
	"unsafe"
)

// interfaces

// 35392C38-CDED-4D4C-9637-05F15C1C9406
var IID_IAtomPubClient = syscall.GUID{0x35392C38, 0xCDED, 0x4D4C,
	[8]byte{0x96, 0x37, 0x05, 0xF1, 0x5C, 0x1C, 0x94, 0x06}}

type IAtomPubClientInterface interface {
	win32.IInspectableInterface
	RetrieveServiceDocumentAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*IServiceDocument, RetrievalProgress]
	RetrieveMediaResourceAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*IInputStream, RetrievalProgress]
	RetrieveResourceAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*ISyndicationItem, RetrievalProgress]
	CreateResourceAsync(uri *IUriRuntimeClass, description string, item *ISyndicationItem) *IAsyncOperationWithProgress[*ISyndicationItem, TransferProgress]
	CreateMediaResourceAsync(uri *IUriRuntimeClass, mediaType string, description string, mediaStream *IInputStream) *IAsyncOperationWithProgress[*ISyndicationItem, TransferProgress]
	UpdateMediaResourceAsync(uri *IUriRuntimeClass, mediaType string, mediaStream *IInputStream) *IAsyncActionWithProgress[TransferProgress]
	UpdateResourceAsync(uri *IUriRuntimeClass, item *ISyndicationItem) *IAsyncActionWithProgress[TransferProgress]
	UpdateResourceItemAsync(item *ISyndicationItem) *IAsyncActionWithProgress[TransferProgress]
	DeleteResourceAsync(uri *IUriRuntimeClass) *IAsyncActionWithProgress[TransferProgress]
	DeleteResourceItemAsync(item *ISyndicationItem) *IAsyncActionWithProgress[TransferProgress]
	CancelAsyncOperations()
}

type IAtomPubClientVtbl struct {
	win32.IInspectableVtbl
	RetrieveServiceDocumentAsync uintptr
	RetrieveMediaResourceAsync   uintptr
	RetrieveResourceAsync        uintptr
	CreateResourceAsync          uintptr
	CreateMediaResourceAsync     uintptr
	UpdateMediaResourceAsync     uintptr
	UpdateResourceAsync          uintptr
	UpdateResourceItemAsync      uintptr
	DeleteResourceAsync          uintptr
	DeleteResourceItemAsync      uintptr
	CancelAsyncOperations        uintptr
}

type IAtomPubClient struct {
	win32.IInspectable
}

func (this *IAtomPubClient) Vtbl() *IAtomPubClientVtbl {
	return (*IAtomPubClientVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAtomPubClient) RetrieveServiceDocumentAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*IServiceDocument, RetrievalProgress] {
	var _result *IAsyncOperationWithProgress[*IServiceDocument, RetrievalProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RetrieveServiceDocumentAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAtomPubClient) RetrieveMediaResourceAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*IInputStream, RetrievalProgress] {
	var _result *IAsyncOperationWithProgress[*IInputStream, RetrievalProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RetrieveMediaResourceAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAtomPubClient) RetrieveResourceAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*ISyndicationItem, RetrievalProgress] {
	var _result *IAsyncOperationWithProgress[*ISyndicationItem, RetrievalProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RetrieveResourceAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAtomPubClient) CreateResourceAsync(uri *IUriRuntimeClass, description string, item *ISyndicationItem) *IAsyncOperationWithProgress[*ISyndicationItem, TransferProgress] {
	var _result *IAsyncOperationWithProgress[*ISyndicationItem, TransferProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateResourceAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), NewHStr(description).Ptr, uintptr(unsafe.Pointer(item)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAtomPubClient) CreateMediaResourceAsync(uri *IUriRuntimeClass, mediaType string, description string, mediaStream *IInputStream) *IAsyncOperationWithProgress[*ISyndicationItem, TransferProgress] {
	var _result *IAsyncOperationWithProgress[*ISyndicationItem, TransferProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateMediaResourceAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), NewHStr(mediaType).Ptr, NewHStr(description).Ptr, uintptr(unsafe.Pointer(mediaStream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAtomPubClient) UpdateMediaResourceAsync(uri *IUriRuntimeClass, mediaType string, mediaStream *IInputStream) *IAsyncActionWithProgress[TransferProgress] {
	var _result *IAsyncActionWithProgress[TransferProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UpdateMediaResourceAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), NewHStr(mediaType).Ptr, uintptr(unsafe.Pointer(mediaStream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAtomPubClient) UpdateResourceAsync(uri *IUriRuntimeClass, item *ISyndicationItem) *IAsyncActionWithProgress[TransferProgress] {
	var _result *IAsyncActionWithProgress[TransferProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UpdateResourceAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(item)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAtomPubClient) UpdateResourceItemAsync(item *ISyndicationItem) *IAsyncActionWithProgress[TransferProgress] {
	var _result *IAsyncActionWithProgress[TransferProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UpdateResourceItemAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(item)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAtomPubClient) DeleteResourceAsync(uri *IUriRuntimeClass) *IAsyncActionWithProgress[TransferProgress] {
	var _result *IAsyncActionWithProgress[TransferProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DeleteResourceAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAtomPubClient) DeleteResourceItemAsync(item *ISyndicationItem) *IAsyncActionWithProgress[TransferProgress] {
	var _result *IAsyncActionWithProgress[TransferProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DeleteResourceItemAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(item)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAtomPubClient) CancelAsyncOperations() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CancelAsyncOperations, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 49D55012-57CB-4BDE-AB9F-2610B172777B
var IID_IAtomPubClientFactory = syscall.GUID{0x49D55012, 0x57CB, 0x4BDE,
	[8]byte{0xAB, 0x9F, 0x26, 0x10, 0xB1, 0x72, 0x77, 0x7B}}

type IAtomPubClientFactoryInterface interface {
	win32.IInspectableInterface
	CreateAtomPubClientWithCredentials(serverCredential *IPasswordCredential) *IAtomPubClient
}

type IAtomPubClientFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateAtomPubClientWithCredentials uintptr
}

type IAtomPubClientFactory struct {
	win32.IInspectable
}

func (this *IAtomPubClientFactory) Vtbl() *IAtomPubClientFactoryVtbl {
	return (*IAtomPubClientFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAtomPubClientFactory) CreateAtomPubClientWithCredentials(serverCredential *IPasswordCredential) *IAtomPubClient {
	var _result *IAtomPubClient
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAtomPubClientWithCredentials, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(serverCredential)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7F5FD609-BC88-41D4-88FA-3DE6704D428E
var IID_IResourceCollection = syscall.GUID{0x7F5FD609, 0xBC88, 0x41D4,
	[8]byte{0x88, 0xFA, 0x3D, 0xE6, 0x70, 0x4D, 0x42, 0x8E}}

type IResourceCollectionInterface interface {
	win32.IInspectableInterface
	Get_Title() *ISyndicationText
	Get_Uri() *IUriRuntimeClass
	Get_Categories() *IVectorView[*ISyndicationCategory]
	Get_Accepts() *IVectorView[string]
}

type IResourceCollectionVtbl struct {
	win32.IInspectableVtbl
	Get_Title      uintptr
	Get_Uri        uintptr
	Get_Categories uintptr
	Get_Accepts    uintptr
}

type IResourceCollection struct {
	win32.IInspectable
}

func (this *IResourceCollection) Vtbl() *IResourceCollectionVtbl {
	return (*IResourceCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IResourceCollection) Get_Title() *ISyndicationText {
	var _result *ISyndicationText
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IResourceCollection) Get_Uri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Uri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IResourceCollection) Get_Categories() *IVectorView[*ISyndicationCategory] {
	var _result *IVectorView[*ISyndicationCategory]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Categories, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IResourceCollection) Get_Accepts() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Accepts, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8B7EC771-2AB3-4DBE-8BCC-778F92B75E51
var IID_IServiceDocument = syscall.GUID{0x8B7EC771, 0x2AB3, 0x4DBE,
	[8]byte{0x8B, 0xCC, 0x77, 0x8F, 0x92, 0xB7, 0x5E, 0x51}}

type IServiceDocumentInterface interface {
	win32.IInspectableInterface
	Get_Workspaces() *IVectorView[*IWorkspace]
}

type IServiceDocumentVtbl struct {
	win32.IInspectableVtbl
	Get_Workspaces uintptr
}

type IServiceDocument struct {
	win32.IInspectable
}

func (this *IServiceDocument) Vtbl() *IServiceDocumentVtbl {
	return (*IServiceDocumentVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IServiceDocument) Get_Workspaces() *IVectorView[*IWorkspace] {
	var _result *IVectorView[*IWorkspace]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Workspaces, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B41DA63B-A4B8-4036-89C5-83C31266BA49
var IID_IWorkspace = syscall.GUID{0xB41DA63B, 0xA4B8, 0x4036,
	[8]byte{0x89, 0xC5, 0x83, 0xC3, 0x12, 0x66, 0xBA, 0x49}}

type IWorkspaceInterface interface {
	win32.IInspectableInterface
	Get_Title() *ISyndicationText
	Get_Collections() *IVectorView[*IResourceCollection]
}

type IWorkspaceVtbl struct {
	win32.IInspectableVtbl
	Get_Title       uintptr
	Get_Collections uintptr
}

type IWorkspace struct {
	win32.IInspectable
}

func (this *IWorkspace) Vtbl() *IWorkspaceVtbl {
	return (*IWorkspaceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWorkspace) Get_Title() *ISyndicationText {
	var _result *ISyndicationText
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWorkspace) Get_Collections() *IVectorView[*IResourceCollection] {
	var _result *IVectorView[*IResourceCollection]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Collections, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type AtomPubClient struct {
	RtClass
	*IAtomPubClient
}

func NewAtomPubClient() *AtomPubClient {
	hs := NewHStr("Windows.Web.AtomPub.AtomPubClient")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &AtomPubClient{
		RtClass:        RtClass{PInspect: p},
		IAtomPubClient: (*IAtomPubClient)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewAtomPubClient_CreateAtomPubClientWithCredentials(serverCredential *IPasswordCredential) *AtomPubClient {
	hs := NewHStr("Windows.Web.AtomPub.AtomPubClient")
	var pFac *IAtomPubClientFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAtomPubClientFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IAtomPubClient
	p = pFac.CreateAtomPubClientWithCredentials(serverCredential)
	result := &AtomPubClient{
		RtClass:        RtClass{PInspect: &p.IInspectable},
		IAtomPubClient: p,
	}
	com.AddToScope(result)
	return result
}

type ResourceCollection struct {
	RtClass
	*IResourceCollection
}

type ServiceDocument struct {
	RtClass
	*IServiceDocument
}

type Workspace struct {
	RtClass
	*IWorkspace
}
